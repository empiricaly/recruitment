package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	participantModel "github.com/empiricaly/recruitment/internal/ent/participant"
	participationModel "github.com/empiricaly/recruitment/internal/ent/participation"
	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	stepRunModel "github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func ginAnswersHandler(s *Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			log.Error().Msg("answers handler: missing URL token")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		workerID := c.Query("workerId")
		if id == "" {
			log.Error().Msg("answers handler: missing workerID")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		assignmentID := c.Query("assignmentId")
		if id == "" {
			log.Error().Msg("answers handler: missing assignmentID")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		hitID := c.Query("hitId")
		if id == "" {
			log.Error().Msg("answers handler: missing hitID")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		// Decode body, which should be key value date to add to Participant
		// e.g. `{ "age": 18, "favoriteColor": "green", "scores": [1,2,3] }`
		dec := json.NewDecoder(c.Request.Body)
		v := map[string]json.RawMessage{}
		err := dec.Decode(&v)
		if err != nil {
			log.Error().Err(err).Msg("answers handler: failed to decode body")
			c.AbortWithStatus(http.StatusNotFound)
		}

		ctx := c.Request.Context()
		var stepRun *ent.StepRun
		err = ent.WithTx(ctx, s.storeConn.Client, func(tx *ent.Tx) error {

			stepRun, err = tx.StepRun.
				Query().
				// WithStep(func(step *ent.StepQuery) {
				// 	step.WithTemplate()
				// }).
				WithStep().
				WithRun(func(run *ent.RunQuery) {
					run.WithProject()
				}).
				WithParticipants().
				Where(stepRunModel.UrlTokenEQ(id)).
				First(ctx)
			if err != nil {
				return errors.Wrap(err, "get stepRun")
			}

			if stepRun.HitID == nil {
				return errors.Errorf("stepRun has nil HIT ID")
			}

			if *stepRun.HitID != hitID {
				return errors.Errorf("stepTun has different HIT ID (%s)", *stepRun.HitID)
			}

			step, err := stepRun.Edges.StepOrErr()
			if err != nil {
				return errors.Wrap(err, "get step")
			}

			if step.Type != stepModel.TypeMTURK_HIT {
				return errors.New("trying to save data on a non-hit step")
			}

			if stepRun.StartedAt == nil {
				return errors.New("stepRun is not running yet, cannot save data")
			}

			run, err := stepRun.Edges.RunOrErr()
			if err != nil {
				return errors.Wrap(err, "get run")
			}

			project, err := run.Edges.ProjectOrErr()
			if err != nil {
				return errors.Wrap(err, "get project")
			}

			participant, err := tx.Participant.
				Query().
				WithParticipations(func(p *ent.ParticipationQuery) {
					p.Where(
						participationModel.And(
							participationModel.MturkWorkerID(workerID),
							participationModel.HasStepRunWith(stepRunModel.IDEQ(stepRun.ID)),
						),
					)
				}).
				WithProjects().
				Where(participantModel.MturkWorkerID(workerID)).
				First(ctx)
			if err != nil && !ent.IsNotFound(err) {
				return errors.Wrap(err, "get participant")
			}

			var createdParticipant bool
			if participant == nil {
				createdParticipant = true
				participant, err = tx.Participant.
					Create().
					SetID(xid.New().String()).
					SetMturkWorkerID(workerID).
					SetUninitialized(false).
					AddProjects(project).
					AddSteps(stepRun).
					SetCreatedBy(stepRun).
					Save(c.Request.Context())
				if err != nil {
					return errors.Wrap(err, "create participant")
				}
			} else {
				projects, err := participant.Edges.ProjectsOrErr()
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
					participant, err = participant.Update().
						AddProjects(project).
						Save(ctx)
					if err != nil {
						return errors.Wrap(err, "update participant project")
					}
				}

				if participant.Uninitialized != nil && *participant.Uninitialized != false {
					participant, err = participant.Update().
						SetUninitialized(false).
						Save(ctx)
					if err != nil {
						return errors.Wrap(err, "Set uninitialized participant")
					}
				}

				stepRunParticipants, err := stepRun.Edges.ParticipantsOrErr()
				if err != nil {
					return errors.Wrap(err, "get stepRun participants")
				}

				var foundStepRunParticipant bool
				for _, p := range stepRunParticipants {
					if p.MturkWorkerID != nil && participant.MturkWorkerID != nil && (*p.MturkWorkerID == *participant.MturkWorkerID) {
						foundStepRunParticipant = true
						break
					}
				}

				if !foundStepRunParticipant {
					participant, err = participant.Update().
						AddSteps(stepRun).
						Save(ctx)
					if err != nil {
						return errors.Wrap(err, "Set stepRun participant")
					}
				}
			}

			participations, err := participant.Edges.ParticipationsOrErr()
			if err != nil && !ent.IsNotLoaded(err) {
				return errors.Wrap(err, "get participations")
			}

			var participation *ent.Participation
			if len(participations) > 0 {
				participation = participations[0]
				if len(participations) > 1 {
					ids := make([]string, len(participations))
					for i, p := range participations {
						ids[i] = p.ID
					}
					log.Warn().Strs("participationIDs", ids).Msg("has multiple participants for one participant/stepRun intersection")
				}
			} else {
				participation, err = tx.Participation.
					Create().
					SetID(xid.New().String()).
					SetMturkWorkerID(workerID).
					SetMturkAssignmentID(assignmentID).
					SetMturkHitID(hitID).
					SetParticipant(participant).
					SetStepRun(stepRun).
					SetAddedParticipant(createdParticipant).
					Save(c.Request.Context())
				if err != nil {
					return errors.Wrap(err, "create participation")
				}
			}

			// Keep for later
			_ = participation

			for key, val := range v {
				_, err = ent.SetDatum(ctx, tx, participant, key, string(val), false)
				if err != nil {
					return errors.Wrap(err, "set datum")
				}
			}

			if stepRun.EndedAt != nil {
				timeExtension := stepRun.EndedAt.Add(time.Minute * time.Duration(step.HitArgs.Timeout))
				remainingTime := timeExtension.Sub(time.Now())

				if remainingTime < 0 {
					log.Error().
						Msg("stepRun no longer running")
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "stepRunEnded"})
					return nil
				}
			}
			return nil
		})
		if err != nil {
			log.Error().Err(err).
				Str("stepRunID", stepRun.ID).
				Str("workerID", workerID).
				Str("hitID", hitID).
				Str("assignmentID", assignmentID).
				Msg("answers handler: commit transaction")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	}
}
