package mturk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/empiricaly/recruitment/internal/ent"
	templateModel "github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/pkg/errors"
)

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

	// TODO: add correct URL there
	question, err := getExternalQuestion("some url here")
	if err != nil {
		return errors.Wrap(err, "encode HIT question URL")
	}

	params := &mturk.CreateHITInput{
		Question:                    aws.String(question),
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
