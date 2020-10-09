package mturk

import (
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func (s *Session) createHit(params *mturk.CreateHITInput) (string, error) {
	var hitID string
	if s.config.Dev {
		log.Debug().Interface("input", params).Msg("Creating HIT")
		hitID = xid.New().String()
	} else {
		hit, err := s.MTurk.CreateHIT(params)
		if err != nil {
			return "", err
		}
		hitID = *hit.HIT.HITId
	}

	return hitID, nil
}

func (s *Session) createQual(params *mturk.CreateQualificationTypeInput) (string, error) {
	var qualID string
	if s.config.Dev {
		log.Debug().Interface("input", params).Msg("Creating Qualification")
		qualID = xid.New().String()
	} else {
		qual, err := s.MTurk.CreateQualificationType(params)
		if err != nil {
			return "", err
		}
		qualID = *qual.QualificationType.QualificationTypeId
	}

	return qualID, nil
}
