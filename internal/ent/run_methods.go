package ent

import (
	"context"
	"time"

	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/pkg/errors"
)

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
