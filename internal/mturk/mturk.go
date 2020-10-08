package mturk

import (
	"encoding/json"
	"fmt"

	rice "github.com/GeertJohan/go.rice"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/empiricaly/recruitment/internal/ent"
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
func (s *Session) RunStep(run *ent.Run, step *ent.Step) error {

	switch step.Type {
	case "MTURK_HIT":
		hit := &model.HITStepArgs{}
		err := json.Unmarshal(step.HitArgs, hit)
		if err != nil {
			return errors.Wrap(err, "unmarshal Step Hit args")
		}
		params := &mturk.CreateHITInput{
			AssignmentDurationInSeconds: aws.Int64(int64(hit.Duration)),
			LifetimeInSeconds:           aws.Int64(int64(hit.Duration)),
			MaxAssignments:              aws.Int64(int64(run.Edges.Procedure.ParticipantCount)),
			Title:                       aws.String(hit.Title),
			Description:                 aws.String(hit.Description),
			Keywords:                    aws.String(hit.Keywords),
			Reward:                      aws.String(fmt.Sprintf("%f", hit.Reward)),
		}
		s.MTurk.CreateHIT(params)

	case "MTURK_MESSAGE":
		// var notif *model.MessageStepArgsInput = step.MsgArgs
		// params := &mturk.NotifyWorkersInput{MessageText: notif.Message}
		// s.MTurk.NotifyWorkers(params)
	// case "PARTICIPANT_FILTER":
	// 	params := &mturk.CreateHITInput{}
	default:
		return errors.Errorf("unknown step type for MTurk: %s", step.Type.String())
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
