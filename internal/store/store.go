package store

import "github.com/philippgille/gokv"

// Conn represents a datastore connection.
type Conn struct {
	gokv.Store
}

// Connect creates a connection to a messaging service with the given config.
func Connect(config *Config) (c *Conn, err error) {
	if config.Badger.Enabled {
		c, err = ConnectBadger(config)
	}

	if config.Redis.Enabled {
		c, err = ConnectRedis(config)
	}

	return c, err
}
