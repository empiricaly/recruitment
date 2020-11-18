package mturk

import (
	"encoding/json"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/rs/zerolog/log"
)

//go:generate rice embed-go

// GetQuals will get a list of qualification types created by the user
func (s *Session) GetQuals() ([]*model.MTurkQulificationType, error) {
	var quals []*model.MTurkQulificationType
	var validQuals []*model.MTurkQulificationType

	params := &mturk.ListQualificationTypesInput{
		MustBeRequestable:   aws.Bool(true),
		MustBeOwnedByCaller: aws.Bool(true),
	}
	err := s.MTurk.ListQualificationTypesPages(params, func(page *mturk.ListQualificationTypesOutput, _ bool) bool {
		for _, qual := range page.QualificationTypes {
			if qual.QualificationTypeStatus != nil && *qual.QualificationTypeStatus == "Active" {
				if qual.Keywords == nil || !strings.Contains(*qual.Keywords, "empirica_recruitment_internal") {
					quals = append(quals, &model.MTurkQulificationType{ID: *qual.QualificationTypeId, Name: *qual.Name, Description: *qual.Description, Type: "Custom"})
				}
			}
		}
		return true
	})
	if err != nil {
		log.Error().Err(err).Msg("get custom qualificationTypes")
	}

	for _, qual := range s.quals {
		quals = append(quals, qual)
	}

	for _, qual := range quals {
		if strings.Contains(qual.Name, "Inc: [") || strings.Contains(qual.Name, "Exc: [") {
			continue
		}

		validQuals = append(validQuals, qual)
	}

	return validQuals, nil
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
