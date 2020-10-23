package storage

import (
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

// runMigrations is for running custom migrations.
func (c *Conn) runMigrations() error {
	c.logger.Debug().Msg("Running migrations")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := c.runMigrationParticipantsToProjects(ctx)
	if err != nil {
		return errors.Wrap(err, "runMigrationParticipantsToProjects")
	}

	return nil
}

func (c *Conn) runMigrationParticipantsToProjects(ctx context.Context) error {
	c.logger.Debug().Msg("Running migrations – Participants should have a project")

	participants, err := c.Participant.Query().WithProjects().WithSteps().All(ctx)
	if err != nil {
		return errors.Wrap(err, "getting all participants")
	}

	for _, p := range participants {
		projs, err := p.Edges.ProjectsOrErr()
		if err != nil {
			continue
		}
		if len(projs) == 0 {
			steps, err := p.Edges.StepsOrErr()
			if err != nil {
				continue
			}
			run, err := steps[0].QueryRun().Only(ctx)
			if err != nil {
				continue
			}
			project, err := run.QueryProject().Only(ctx)
			if err != nil {
				continue
			}

			log.Debug().Msg("adding project")
			p.Update().AddProjects(project).Save(ctx)
		}
	}

	return nil
}
