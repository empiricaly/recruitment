package mturk

import (
	"context"
	"fmt"
	"time"

	errs "github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

var devHITMap map[string]*mturk.HIT

func init() {
	devHITMap = make(map[string]*mturk.HIT)
}

func (s *Session) stopHit(ctx context.Context, hitID string) error {
	if s.config.Dev {
		log.Debug().Interface("hitID", hitID).Msg("Stopping HIT")
		selectedHit := devHITMap[hitID]
		if selectedHit == nil {
			log.Warn().Msg("failed to stop Hit (dev mode)")
			return nil
		}
		delete(devHITMap, hitID)
	} else {
		date := "1999-12-31"
		t, _ := time.Parse("layoutISO", date)
		_, err := s.MTurk.UpdateExpirationForHITWithContext(ctx, &mturk.UpdateExpirationForHITInput{
			HITId:    aws.String(hitID),
			ExpireAt: aws.Time(t),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) getHit(ctx context.Context, hitID string) (*mturk.HIT, error) {
	var selectedHit *mturk.HIT
	if s.config.Dev {
		log.Debug().Interface("hitID", hitID).Msg("Getting HIT")
		selectedHit = devHITMap[hitID]
		if selectedHit == nil {
			log.Error().Msg("failed to get HIT")
			return nil, errs.Wrap(fmt.Errorf("getHit Dev"), "get Hit")
		}
	} else {
		hit, err := s.MTurk.GetHITWithContext(ctx, &mturk.GetHITInput{HITId: &hitID})
		if err != nil {
			return nil, err
		}
		selectedHit = hit.HIT
	}

	return selectedHit, nil
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
			log.Error().Err(err).Msg("failed to createHit HIT")
			return "", err
		}
		hitID = *hit.HIT.HITId
	}

	return hitID, nil
}

func (s *Session) getQualificationType(ctx context.Context, qualID string) (*mturk.QualificationType, error) {
	var selectedQual *mturk.QualificationType
	if s.config.Dev {
		log.Debug().Interface("qualID", qualID).Msg("Getting Qualification")
	} else {
		qual, err := s.MTurk.GetQualificationTypeWithContext(ctx, &mturk.GetQualificationTypeInput{QualificationTypeId: aws.String(qualID)})
		if err != nil {
			return nil, err
		}
		selectedQual = qual.QualificationType
	}

	return selectedQual, nil
}

func (s *Session) deleteQualificationType(ctx context.Context, qualID string) error {
	if s.config.Dev {
		log.Debug().Interface("qualID", qualID).Msg("Deleting Qualification")
	} else {
		_, err := s.MTurk.DeleteQualificationTypeWithContext(ctx, &mturk.DeleteQualificationTypeInput{QualificationTypeId: aws.String(qualID)})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) associateQualificationWithWorker(ctx context.Context, params *mturk.AssociateQualificationWithWorkerInput) error {
	if s.config.Dev {
		log.Debug().Interface("input", params).Msg("Associating qualification with worker")
	} else {
		_, err := s.MTurk.AssociateQualificationWithWorkerWithContext(ctx, params)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) createQualificationType(ctx context.Context, params *mturk.CreateQualificationTypeInput) (string, error) {
	var qualID string
	if s.config.Dev {
		log.Debug().Interface("input", params).Msg("Creating Qualification")
		qualID = xid.New().String()
	} else {
		qual, err := s.MTurk.CreateQualificationTypeWithContext(ctx, params)
		if err != nil {
			return "", err
		}
		qualID = *qual.QualificationType.QualificationTypeId
	}

	return qualID, nil
}

func (s *Session) notifyWorkers(ctx context.Context, subject, text string, workerIDs []string) (*mturk.NotifyWorkersOutput, error) {
	if s.config.Dev {
		log.Debug().Strs("players", workerIDs).Msg("Notify Players")
	} else {
		params := &mturk.NotifyWorkersInput{
			WorkerIds:   aws.StringSlice(workerIDs),
			Subject:     aws.String(subject),
			MessageText: aws.String(text),
		}
		res, err := s.MTurk.NotifyWorkersWithContext(ctx, params)
		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return &mturk.NotifyWorkersOutput{}, nil
}

func (s *Session) assignmentsForHit(ctx context.Context, hitID string) (chan *mturk.Assignment, error) {
	c := make(chan *mturk.Assignment)

	if s.config.Dev {
		log.Debug().Msg("Getting assignments")
		close(c)
	} else {
		go func() {
			defer close(c)

			for {
				var nextToken *string
				params := &mturk.ListAssignmentsForHITInput{
					AssignmentStatuses: aws.StringSlice([]string{"Approved"}),
					HITId:              aws.String(hitID),
					MaxResults:         aws.Int64(100),
				}

				if nextToken != nil {
					params.NextToken = nextToken
				}

				assignments, err := s.MTurk.ListAssignmentsForHITWithContext(ctx, params)
				if err != nil {
					log.Error().Err(err).Msg("list assignments for hit")
				}

				for _, assignment := range assignments.Assignments {
					c <- assignment
				}

				if assignments.NextToken != nil && *assignments.NextToken != "" {
					nextToken = assignments.NextToken
				} else {
					return
				}
			}
		}()
	}

	return c, nil
}
