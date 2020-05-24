package store

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config is kvstore configuration
type Config struct {
	Badger BadgerConfig `mapstructure:"badger"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

// Validate configuration is ok
func (c *Config) Validate() error {
	if err := c.Badger.Validate(); err != nil {
		return errors.Wrap(err, "badger config error")
	}

	if err := c.Redis.Validate(); err != nil {
		return errors.Wrap(err, "redis config error")
	}

	if c.Redis.Enabled {
		c.Badger.Enabled = false
	}

	if !c.Redis.Enabled && !c.Badger.Enabled {
		return errors.New("at least one store must be enabled")
	}

	return nil
}

// ConfigFlags helps configure cobra and viper flags.
func ConfigFlags(cmd *cobra.Command, prefix string) error {
	if prefix == "" {
		prefix = "store"
	}

	viper.SetDefault(prefix, &Config{})

	badgerConfigFlags(cmd, prefix+".badger")
	redisConfigFlags(cmd, prefix+".redis")

	return nil
}
