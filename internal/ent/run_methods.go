package ent

import (
	"context"
	"fmt"
	"time"

	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	errs "github.com/pkg/errors"
)

// Validate will validate all fields on Run
func (run *Run) Validate() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var multiError error

	template, err := run.QueryTemplate().WithSteps().Only(ctx)
	if err != nil {
		err = errs.Wrap(err, "query validate: template")
		return err
	}

	// Check participantCount
	if template.ParticipantCount < 1 {
		multiError = multierror.Append(multiError, errors.New("template: participantCount cannot be less than one"))
	}

	// Check steps
	for i, step := range template.Edges.Steps {
		switch step.Type {
		case stepModel.TypeMTURK_HIT:
			if step.Duration < 5 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_hit: duration cannot be less that 5 minutes, index: %v", i+1))
			}

			if len(step.HitArgs.Title) == 0 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_hit: tittle cannot be empty, index: %v", i+1))
			}

			if len(step.HitArgs.Description) == 0 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_hit: description cannot be empty, index: %v", i+1))
			}

			if len(step.HitArgs.Keywords) == 0 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_hit: keywords cannot be empty, index: %v", i+1))
			}

			if step.HitArgs.Reward <= 0 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_hit: reward cannot be less than or equal to 0, index: %v", i+1))
			}

			if step.HitArgs.Timeout < 1 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_hit: HIT Timeout cannot be less than to 1, index: %v", i+1))
			}

			if len(step.MsgArgs.Message) == 0 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_hit: message cannot be empty, index: %v", i+1))
			}

			break
		case stepModel.TypeMTURK_MESSAGE:
			if step.Duration < 0 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_message: duration cannot be less that 0, index: %v", i+1))
			}

			if len(*step.MsgArgs.Subject) == 0 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_message: subject cannot be empty, index: %v", i+1))
			}

			if len(step.MsgArgs.Message) == 0 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_mturk_message: message cannot be empty, index: %v", i+1))
			}

			break
		case stepModel.TypePARTICIPANT_FILTER:
			if step.Duration < 0 {
				multiError = multierror.Append(multiError, fmt.Errorf("step_participant_filter: duration cannot be less that 0, index: %v", i+1))
			}
			break
		}

	}

	return multiError
}

// Relations returns all the Run close relations (steps, stepRuns, and template).
func (run *Run) Relations() (currentStepRun *StepRun, currentStep *Step, stepRuns []*StepRun, template *Template, steps []*Step, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	currentStepRun, err = run.QueryCurrentStep().Only(ctx)
	if err != nil && !IsNotFound(err) {
		err = errors.Wrap(err, "query relations: currentStepRun")
		return
	}

	template, err = run.QueryTemplate().Only(ctx)
	if err != nil {
		err = errors.Wrap(err, "query relations: template")
		return
	}

	steps, err = template.QuerySteps().Order(Asc(stepModel.FieldIndex)).All(ctx)
	if err != nil {
		err = errors.Wrap(err, "query relations: steps")
		return
	}

	stepRuns, err = run.QuerySteps().All(ctx)
	if err != nil {
		err = errors.Wrap(err, "query relations: runsteps")
		return
	}
	if currentStepRun == nil {
		// run this step and go to next
		return
	}

	currentStep, err = currentStepRun.QueryStep().Only(ctx)
	if err != nil {
		err = errors.Wrap(err, "query relations: currentStep")
		return
	}

	return
}
