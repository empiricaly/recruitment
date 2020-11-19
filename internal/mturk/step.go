package mturk

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aymerick/raymond"
	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/participation"
	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	stepRunModel "github.com/empiricaly/recruitment/internal/ent/steprun"
	templateModel "github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

type renderContext struct {
	URL string `handlebars:"url"`
}

const workerAdultQualTypeID = "00000000000000000060"

var errStepWithoutParticipants = errors.New("message step without participants")
var errInvalidInitialMessageStep = errors.New("message step cannot be first with mturk player selection")

// StartStep to run step based on stepType
func (s *Session) StartStep(project *ent.Project, run *ent.Run, step *ent.Step, stepRun *ent.StepRun, startTime time.Time) error {
	s.logger.Debug().Msg("Running step")
	defer s.logger.Debug().Msg("Step ran")
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	var err error
	template, err := run.Edges.TemplateOrErr()
	if err != nil {
		template, err = run.QueryTemplate().Only(ctx)
		if err != nil {
			return errors.Wrap(err, "query template for stepRun")
		}
	}

	switch step.Type {
	case stepModel.TypeMTURK_HIT:
		return s.runMTurkHITStep(ctx, project, run, stepRun, template, step, startTime)
	case stepModel.TypeMTURK_MESSAGE:
		return s.runMTurkMessageStep(ctx, project, run, stepRun, template, step, startTime)
	case stepModel.TypeWAIT:
		return nil
	default:
		return errors.Errorf("unknown step type for MTurk: %s", step.Type.String())
	}
}

func (s *Session) runMTurkHITStep(ctx context.Context, project *ent.Project, run *ent.Run, stepRun *ent.StepRun, template *ent.Template, step *ent.Step, startTime time.Time) error {
	// rootURL is already tested from config options, should never fail to parse
	addr, err := url.Parse(s.rootURL)
	if err != nil {
		return errors.Wrap(err, "get root URL")
	}
	addr.Path = "/q/" + stepRun.UrlToken

	question, err := getExternalQuestion(addr.String())
	if err != nil {
		return errors.Wrap(err, "encode HIT question URL")
	}

	isFirstStep := step.Index == 0
	isMturkSelection := template.SelectionType == templateModel.SelectionTypeMTURK_QUALIFICATIONS

	var quals []*mturk.QualificationRequirement
	var assignmentCount int
	if isFirstStep && isMturkSelection {
		assignmentCount = template.ParticipantCount
		for _, q := range template.MturkCriteria.Qualifications {
			ints := make([]*int64, len(q.Values))
			for i, val := range q.Values {
				ints[i] = aws.Int64(int64(val))
			}

			locales := make([]*mturk.Locale, len(q.Locales))
			for i, val := range q.Locales {
				locales[i] = &mturk.Locale{
					Country:     aws.String(val.Country),
					Subdivision: val.Subdivision,
				}
			}

			quals = append(quals, &mturk.QualificationRequirement{
				QualificationTypeId: aws.String(q.ID),
				Comparator:          aws.String(q.Comparator.String()),
				IntegerValues:       ints,
				LocaleValues:        locales,
			})
		}
	} else {
		qualParams := &mturk.CreateQualificationTypeInput{
			AutoGranted:             aws.Bool(true),
			Name:                    aws.String(fmt.Sprintf("empirica_%s", stepRun.ID)),
			Description:             aws.String("internal empirica qual"),
			Keywords:                aws.String("empirica_recruitment_internal"),
			QualificationTypeStatus: aws.String("Active"),
		}
		qualID, err := s.createQualificationType(ctx, qualParams)
		if err != nil {
			return errors.Wrap(err, "create step qual")
		}

		participants, err := stepRun.QueryParticipants().All(ctx)
		if err != nil {
			return errors.Wrap(err, "get step participants")
		}

		for _, participant := range participants {
			params := &mturk.AssociateQualificationWithWorkerInput{
				QualificationTypeId: aws.String(qualID),
				SendNotification:    aws.Bool(false),
				WorkerId:            aws.String(*participant.MturkWorkerID),
			}
			err = s.associateQualificationWithWorker(ctx, params)
			if err != nil {
				return errors.Wrap(err, "associate individual to qual")
			}
		}

		quals = append(quals, &mturk.QualificationRequirement{
			QualificationTypeId: aws.String(qualID),
			Comparator:          aws.String("Exists"),
			ActionsGuarded:      aws.String("DiscoverPreviewAndAccept"),
		})
		assignmentCount = len(participants)
	}

	if template.Adult {
		quals = append(quals, &mturk.QualificationRequirement{
			QualificationTypeId: aws.String(workerAdultQualTypeID),
			Comparator:          aws.String("Exists"),
		})
	}

	params := &mturk.CreateHITInput{
		Question:                    aws.String(question),
		AssignmentDurationInSeconds: aws.Int64(int64(step.HitArgs.Timeout * 60)),
		LifetimeInSeconds:           aws.Int64(int64(step.Duration * 60)),
		MaxAssignments:              aws.Int64(int64(assignmentCount)),
		Title:                       aws.String(step.HitArgs.Title),
		Description:                 aws.String(step.HitArgs.Description),
		Keywords:                    aws.String(step.HitArgs.Keywords),
		Reward:                      aws.String(fmt.Sprintf("%.2f", step.HitArgs.Reward)),
		QualificationRequirements:   quals,
		UniqueRequestToken:          aws.String(stepRun.ID),
		AutoApprovalDelayInSeconds:  aws.Int64(30),
	}

	hitID, err := s.createHit(ctx, params)
	if err != nil {
		return errors.Wrap(err, "create hit")
	}

	_, err = stepRun.Update().SetHitID(hitID).Save(ctx)
	if err != nil {
		return errors.Wrap(err, "save hit ID in StepRun")
	}

	return nil
}

func (s *Session) runMTurkMessageStep(ctx context.Context, project *ent.Project, run *ent.Run, stepRun *ent.StepRun, template *ent.Template, step *ent.Step, startTime time.Time) error {
	if step.Index == 0 && template.SelectionType == templateModel.SelectionTypeMTURK_QUALIFICATIONS {
		return errInvalidInitialMessageStep
	}

	subject := "New Message"
	if step.MsgArgs.Subject == nil {
		subject = *step.MsgArgs.Subject
	}

	participants, err := stepRun.QueryParticipants().All(ctx)
	if err != nil {
		return errors.Wrap(err, "get step participants")
	}

	workerIDs := make([]string, 0, len(participants))
	for _, participant := range participants {
		workerIDs = append(workerIDs, *participant.MturkWorkerID)
	}

	if !s.config.Dev && len(workerIDs) == 0 {
		return errStepWithoutParticipants
	}

	msg := step.MsgArgs.Message
	if step.MsgArgs.URL != nil {
		u, err := url.Parse(*step.MsgArgs.URL)
		if err != nil {
			log.Error().Err(err).Msg("invalid step message URL")
		} else {
			renderCtx := &renderContext{
				URL: u.String(),
			}
			r, err := raymond.Render(msg, renderCtx)
			if err != nil {
				log.Error().Err(err).Msg("failed to render HTML message: URL")
			} else {
				msg = r
			}
		}
	}

	output, err := s.notifyWorkers(ctx, subject, msg, workerIDs)
	if err != nil {
		return errors.Wrap(err, "notify workers")
	}

	failedWorkedIDs := make(map[string]*mturk.NotifyWorkersFailureStatus)
	for _, stat := range output.NotifyWorkersFailureStatuses {
		failedWorkedIDs[*stat.WorkerId] = stat
	}

	for _, p := range participants {
		if p.MturkWorkerID == nil {
			log.Warn().
				Str("participantID", p.ID).
				Msg("notifying participant without workerID?!")
			continue
		}

		stat, ok := failedWorkedIDs[*p.MturkWorkerID]
		if ok {
			log.Error().
				Str("mturkErrCode", *stat.NotifyWorkersFailureCode).
				Str("mturkErr", *stat.NotifyWorkersFailureMessage).
				Msg("could not notify worker")
			continue
		}

		partParams := s.store.Participation.Create().
			SetID(xid.New().String()).
			SetParticipant(p).
			SetStepRun(stepRun).
			SetMturkWorkerID(*p.MturkWorkerID)

		_, err := partParams.Save(ctx)
		if err != nil {
			log.Error().
				Str("workerID", *p.MturkWorkerID).
				Str("stepRunID", stepRun.ID).
				Msg("could not add participant to stepRun")
		}
	}

	return nil
}

// EndStep to end step based on stepType
func (s *Session) EndStep(project *ent.Project, run *ent.Run, step *ent.Step, stepRun *ent.StepRun, nextStep *ent.Step, nextStepRun *ent.StepRun) error {
	s.logger.Debug().Msg("Ending step")
	defer s.logger.Debug().Msg("Step ended")

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	var err error
	template, err := run.Edges.TemplateOrErr()
	if err != nil {
		template, err = run.QueryTemplate().Only(ctx)
		if err != nil {
			return errors.Wrap(err, "query template for stepRun")
		}
	}

	switch step.Type {
	case stepModel.TypeMTURK_HIT:
		return s.endMTurkHITStep(ctx, project, run, template, step, stepRun, nextStep, nextStepRun)
	case stepModel.TypeMTURK_MESSAGE:
		return s.endMTurkMessageStep(ctx, project, run, template, step, stepRun, nextStep, nextStepRun)
	case stepModel.TypeWAIT:
		return nil
	default:
		return errors.Errorf("unknown step type for MTurk: %s", step.Type.String())
	}
}

func (s *Session) endMTurkHITStep(ctx context.Context, project *ent.Project, run *ent.Run, template *ent.Template, step *ent.Step, stepRun *ent.StepRun, nextStep *ent.Step, nextStepRun *ent.StepRun) error {
	// Getting stepRun again, to make sure we have the HIT ID saved previously
	// Maybe should refresh stepRun higher up in the chain...
	s.store.StepRun.Get(ctx, stepRun.ID)

	if stepRun.HitID == nil {
		return errors.New("missing HIT ID to end step")
	}

	err := s.stopHit(ctx, *stepRun.HitID)
	if err != nil {
		log.Error().Err(err).Msg("stop hit failed")
	}

	c, err := s.assignmentsForHit(ctx, *stepRun.HitID)
	if err != nil {
		return errors.New("get assignments for hit step failed")
	}

	participants, err := stepRun.QueryParticipants().
		WithParticipations(func(q *ent.ParticipationQuery) {
			q.Where(participation.HasStepRunWith(stepRunModel.IDEQ(stepRun.ID)))
		}).
		WithProjects().
		All(ctx)
	if err != nil {
		return errors.New("get participants for hit failed")
	}

	m := make(map[string]*ent.Participant)
	for _, p := range participants {
		if p.MturkWorkerID == nil {
			continue
		}
		m[*p.MturkWorkerID] = p
	}

	nextParticipants := make([]*ent.Participant, 0, len(participants))

	for {
		assignment, ok := <-c
		if !ok {
			break
		}

		p := m[*assignment.WorkerId]
		if p == nil {
			p, err = s.store.Participant.Create().
				SetID(xid.New().String()).
				AddProjects(project).
				SetMturkWorkerID(*assignment.WorkerId).
				SetUninitialized(false).
				SetCreatedBy(stepRun).
				AddSteps(stepRun).
				Save(ctx)
			if err != nil {
				log.Error().Msgf("could not add participant with workerID %s for stepRun %s", *assignment.WorkerId, stepRun.ID)
				continue
			}

			m[*p.MturkWorkerID] = p
		} else {
			projects, err := p.Edges.ProjectsOrErr()
			if err != nil {
				return errors.Wrap(err, "get participant projects")
			}
			var found bool
			for _, project := range projects {
				if project.ID == project.ID {
					found = true
					break
				}
			}
			if !found {
				p, err = p.Update().
					AddProjects(project).
					AddSteps(stepRun).
					SetUninitialized(false).
					Save(ctx)
				if err != nil {
					log.Error().Msgf("could not update participant with workerID %s for stepRun %s", *assignment.WorkerId, stepRun.ID)
					continue
				}
			}
		}
		nextParticipants = append(nextParticipants, p)

		participation, err := p.Edges.ParticipationsOrErr()
		if err != nil || participation == nil {
			partParams := s.store.Participation.Create().
				SetID(xid.New().String()).
				SetParticipant(p).
				SetStepRun(stepRun).
				SetMturkWorkerID(*assignment.WorkerId).
				SetMturkAssignmentID(*assignment.AssignmentId).
				SetMturkHitID(*assignment.HITId)

			if assignment.AcceptTime != nil {
				partParams.SetMturkAcceptedAt(*assignment.AcceptTime)
			}
			if assignment.SubmitTime != nil {
				partParams.SetMturkSubmittedAt(*assignment.SubmitTime)
			}

			_, err := partParams.Save(ctx)
			if err != nil {
				log.Error().Msgf("could not add participant with workerID %s for stepRun %s", *assignment.WorkerId, stepRun.ID)
				continue
			}
		}

	}

	if nextStepRun != nil {
		_, err = nextStepRun.Update().
			AddParticipants(nextParticipants...).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "push msg step participants to next run")
		}
	}

	return nil
}

func (s *Session) endMTurkMessageStep(ctx context.Context, project *ent.Project, run *ent.Run, template *ent.Template, step *ent.Step, stepRun *ent.StepRun, nextStep *ent.Step, nextStepRun *ent.StepRun) error {
	if nextStepRun == nil {
		return nil
	}

	participations, err := stepRun.QueryParticipations().WithParticipant().All(ctx)
	if err != nil {
		return errors.New("get participants for msg step failed")
	}

	participants := make([]*ent.Participant, len(participations))

	for i, p := range participations {
		participants[i] = p.Edges.Participant
	}

	_, err = nextStepRun.Update().
		AddParticipants(participants...).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "push msg step participants to next run")
	}

	return nil
}
