package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func (s *Session) stopHit(hitID string) error {
	// UpdateExpirationForHITWithContext(0)
	return nil
}

func (s *Session) getHit(hitID string) (*mturk.HIT, error) {
	// GetHITWithContext
	// devHITMap[hitID]
	return nil, nil
}

var devHITMap map[string]*mturk.HIT

func init() {
	devHITMap = make(map[string]*mturk.HIT)
}

func (s *Session) createHit(ctx context.Context, params *mturk.CreateHITInput) (string, error) {
	var hitID string
	if s.config.Dev {
		log.Debug().Interface("input", params).Msg("Creating HIT")
		hitID = xid.New().String()
		devHITMap[hitID] = &mturk.HIT{
			AssignmentDurationInSeconds: params.AssignmentDurationInSeconds,
		}
	} else {
		hit, err := s.MTurk.CreateHITWithContext(ctx, params)
		if err != nil {
			return "", err
		}
		hitID = *hit.HIT.HITId
	}

	return hitID, nil
}

func (s *Session) getQual(qualID string) (*mturk.Qualification, error) {
	return nil, nil
}

func (s *Session) deleteQual(qualID string) error {
	return nil
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
