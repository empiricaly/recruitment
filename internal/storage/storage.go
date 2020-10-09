package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/migrate"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	// SQLLite3 package required by ent to init the SQLLite Storage
	_ "github.com/mattn/go-sqlite3"
)

// Conn represents a datastore connection.
type Conn struct {
	*ent.Client
}

// Connect creates a connection to a messaging service with the given config.
func dbLog(msg ...interface{}) {
	if len(msg) == 1 {
		log.Debug().Interface("query", msg[0]).Msg("db")
	} else {
		log.Debug().Interface("query", msg).Msg("db log")
	}
	// spew.Dump(msg)
	// for _, m := range msg {
	// 	log.Debug().Msgf("%v", m)
	// }
}

// Connect creates a connection to a messaging service with the given config.
func Connect(ctx context.Context, config *Config) (c *Conn, err error) {
	connString := fmt.Sprintf("file:%s?mode=rwc&_fk=1", config.File)
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(0)
	db.SetConnMaxLifetime(0)

	drv := entsql.OpenDB("sqlite3", db)

	options := []ent.Option{ent.Driver(drv)}
	if config.Debug {
		options = append(options, ent.Log(dbLog))
		options = append(options, ent.Debug())
	}
	client := ent.NewClient(options...)

	// client, err := ent.Open("sqlite3", connString, options...)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "open sqlite conn")
	// }

	if config.Debug {
		client = client.Debug()
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		return nil, errors.Wrap(err, "write sqlite schema")
	}

	return &Conn{Client: client}, nil
}

// Close cleanlu close the database connection.
func (c *Conn) Close() error {
	return c.Client.Close()
}
