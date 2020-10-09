package mturk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/empiricaly/recruitment/internal/ent"
	templateModel "github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

//go:generate rice embed-go

const productionURL = "https://mturk-requester.us-east-1.amazonaws.com"
const sandboxURL = "https://mturk-requester-sandbox.us-east-1.amazonaws.com"
const awsRegion = "us-east-1"

// Session keep the mTurk session state.
type Session struct {
	quals []*model.MTurkQulificationType
	*mturk.MTurk
}

func loadQuals(sandbox bool) ([]*model.MTurkQulificationType, error) {
	file := "prod.json"
	if sandbox {
		file = "sandbox.json"
	}

	box, err := rice.FindBox("quals")
	if err != nil {
		return nil, err
	}

	content, err := box.Bytes(file)
	if err != nil {
		return nil, err
	}

	var jquals struct {
		Qualtypes []struct {
			Type  string
			Name  string
			Items []*model.MTurkQulificationType
		}
	}
	err = json.Unmarshal(content, &jquals)
	if err != nil {
		return nil, err
	}

	quals := []*model.MTurkQulificationType{}
	for _, t := range jquals.Qualtypes {
		for _, item := range t.Items {
			quals = append(quals, item)
		}
	}

	return quals, nil
}

// GetLocales will get JSON Locales
func (s *Session) GetLocales() ([]*model.MTurkLocale, error) {
	file := "locales.json"

	box, err := rice.FindBox("locales")
	if err != nil {
		return nil, err
	}

	content, err := box.Bytes(file)
	if err != nil {
		return nil, err
	}

	var locales []*model.MTurkLocale

	err = json.Unmarshal(content, &locales)
	if err != nil {
		return nil, err
	}

	return locales, nil
}

// GetQuals will get a list of qualification types created by the user
func (s *Session) GetQuals() ([]*model.MTurkQulificationType, error) {
	var quals []*model.MTurkQulificationType
	params := &mturk.ListQualificationTypesInput{MustBeRequestable: aws.Bool(true), MustBeOwnedByCaller: aws.Bool(true)}
	err := s.MTurk.ListQualificationTypesPages(params, func(page *mturk.ListQualificationTypesOutput, lastPage bool) bool {
		for _, qual := range page.QualificationTypes {
			quals = append(quals, &model.MTurkQulificationType{ID: *qual.QualificationTypeId, Name: *qual.Name, Description: *qual.Description, Type: "Custom"})
		}
		return true
	})
	if err != nil {
		log.Error().Err(err).Msg("get custom qualificationTypes")
	}

	for _, qual := range s.quals {
		quals = append(quals, qual)
	}

	return quals, nil
}

// RunStep to run step based on stepType
func (s *Session) RunStep(run *ent.Run, stepRun *ent.StepRun) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	step, err := stepRun.Edges.StepOrErr()
	if err != nil {
		step, err = stepRun.QueryStep().Only(ctx)
		if err != nil {
			return errors.Wrap(err, "query step for stepRun")
		}
	}

	template, err := run.Edges.TemplateOrErr()
	if err != nil {
		template, err = run.QueryTemplate().Only(ctx)
		if err != nil {
			return errors.Wrap(err, "query template for stepRun")
		}
	}

	switch step.Type {
	case "MTURK_HIT":
		return s.runMTurkHITStep(ctx, run, stepRun, template, step)
	case "MTURK_MESSAGE":
		return s.runMTurkMessageStep(ctx, run, stepRun, template, step)
	default:
		return errors.Errorf("unknown step type for MTurk: %s", step.Type.String())
	}
}

func (s *Session) runMTurkMessageStep(ctx context.Context, run *ent.Run, stepRun *ent.StepRun, template *ent.Template, step *ent.Step) error {
	if template.SelectionType == templateModel.SelectionTypeMTURK_QUALIFICATIONS {

	}

	hit := &model.HITStepArgs{}
	err := json.Unmarshal(step.HitArgs, hit)
	if err != nil {
		return errors.Wrap(err, "unmarshal Step Hit args")
	}
	// var notif *model.MessageStepArgsInput = step.MsgArgs
	// params := &mturk.NotifyWorkersInput{MessageText: notif.Message}
	// s.MTurk.NotifyWorkers(params)
	// case "PARTICIPANT_FILTER":
	// 	params := &mturk.CreateHITInput{}

	return nil
}

func (s *Session) runMTurkHITStep(ctx context.Context, run *ent.Run, stepRun *ent.StepRun, template *ent.Template, step *ent.Step) error {
	hitArgs := &model.HITStepArgs{}
	err := json.Unmarshal(step.HitArgs, hitArgs)
	if err != nil {
		return errors.Wrap(err, "unmarshal Step Hit args")
	}

	isFirstStep := step.Index == 0
	isInternalDB := template.SelectionType == templateModel.SelectionTypeINTERNAL_DB

	if isFirstStep && isInternalDB {
		// Go crazzy
	}

	var quals []*mturk.QualificationRequirement
	if !isFirstStep {
		// Go a little less crazy
		// Create a special new qual to allows participants to go to this step
		// and add it to quals
	} else {
		crit := &model.MTurkCriteria{}
		err := json.Unmarshal(template.MturkCriteria, crit)
		if err != nil {
			return errors.Wrap(err, "unmarshal MturkCriteria")
		}
		for _, q := range crit.Qualifications {
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
	}

	params := &mturk.CreateHITInput{
		AssignmentDurationInSeconds: aws.Int64(int64(hitArgs.Duration)),
		LifetimeInSeconds:           aws.Int64(int64(hitArgs.Duration)),
		MaxAssignments:              aws.Int64(int64(template.ParticipantCount)),
		Title:                       aws.String(hitArgs.Title),
		Description:                 aws.String(hitArgs.Description),
		Keywords:                    aws.String(hitArgs.Keywords),
		Reward:                      aws.String(fmt.Sprintf("%f", hitArgs.Reward)),
		QualificationRequirements:   quals,
	}
	hit, err := s.MTurk.CreateHIT(params)
	if err != nil {
		return errors.Wrap(err, "create hit")
	}

	_, err = stepRun.Update().SetHitID(*hit.HIT.HITId).Save(ctx)
	if err != nil {
		return errors.Wrap(err, "save hit ID in StepRun")
	}

	return nil
}

// New create a new session for mTurk
func New(config *Config) (*Session, error) {
	var endpoint string
	if config.Sandbox {
		endpoint = sandboxURL
	} else {
		endpoint = productionURL
	}

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(awsRegion),
		Endpoint: aws.String(endpoint),
	})
	if err != nil {
		return nil, errors.Wrap(err, "create new aws session")
	}

	quals, err := loadQuals(config.Sandbox)
	if err != nil {
		return nil, errors.Wrap(err, "load json quals")
	}

	svc := mturk.New(sess)
	return &Session{
		MTurk: svc,
		quals: quals,
	}, nil
}
