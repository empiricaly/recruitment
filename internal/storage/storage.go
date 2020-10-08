package storage

import (
	"context"

	"github.com/empiricaly/recruitment/internal/ent"
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
	for _, m := range msg {
		log.Debug().Msgf("%v", m)
	}
}

// Connect creates a connection to a messaging service with the given config.
func Connect(config *Config) (c *Conn, err error) {
	options := []ent.Option{}
	if config.Debug {
		options = append(options, ent.Log(dbLog))
	}

	client, err := ent.Open("sqlite3", "file:"+config.File+"?mode=rwc&_fk=1", options...)
	if err != nil {
		return nil, errors.Wrap(err, "open sqlite conn")
	}

	if config.Debug {
		client = client.Debug()
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, errors.Wrap(err, "write sqlite schema")
	}

	// return c, err
	return &Conn{Client: client}, nil
}

// Close cleanlu close the database connection.
func (c *Conn) Close() error {
	return c.Client.Close()
}
