package server

import (
	"github.com/gin-gonic/gin"
)

func ginAnswersHandler(s *Server) func(c *gin.Context) {
	return func(c *gin.Context) {
		// id := c.Param("id")
		// if id == "" {
		// 	log.Error().Msg("answers handler: missing URL token")
		// 	c.AbortWithStatus(http.StatusNotFound)
		// 	return
		// }

		// workerID := c.Query("workerID")
		// if id == "" {
		// 	log.Error().Msg("answers handler: missing workerID")
		// 	c.AbortWithStatus(http.StatusNotFound)
		// 	return
		// }

		// assignmentID := c.Query("assignmentID")
		// if id == "" {
		// 	log.Error().Msg("answers handler: missing assignmentID")
		// 	c.AbortWithStatus(http.StatusNotFound)
		// 	return
		// }

		// hitID := c.Query("hitID")
		// if id == "" {
		// 	log.Error().Msg("answers handler: missing hitID")
		// 	c.AbortWithStatus(http.StatusNotFound)
		// 	return
		// }

		// stepRun, err := s.storeConn.StepRun.
		// 	Query().
		// 	// WithStep(func(step *ent.StepQuery) {
		// 	// 	step.WithTemplate()
		// 	// }).
		// 	WithRun().
		// 	Where(stepRunModel.UrlTokenEQ(id)).
		// 	First(c.Request.Context())
		// if err != nil {
		// 	log.Error().Err(err).Msg("answers handler: get stepRun")
		// 	c.AbortWithStatus(http.StatusNotFound)
		// 	return
		// }
		// // step, err := stepRun.Edges.StepOrErr()
		// // if err != nil {
		// // 	log.Error().Err(err).Msg("answers handler: get step")
		// // 	c.AbortWithStatus(http.StatusNotFound)
		// // 	return
		// // }

		// participant, err := s.storeConn.Participant.
		// 	Query().
		// 	WithParticipations(func(p *ent.ParticipationQuery) {
		// 		p.Where(participationModel.MturkWorkerID(workerID))
		// 	}).
		// 	Where(participantModel.MturkWorkerID(workerID)).
		// 	First(c.Request.Context())
		// if err != nil && !ent.IsNotFound(err) {
		// 	log.Error().Err(err).Msg("answers handler: get participant")
		// 	c.AbortWithStatus(http.StatusNotFound)
		// 	return
		// }
	}
}
