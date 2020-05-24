package server

import (
	"github.com/empiricaly/recruitment/internal/log"
	"github.com/empiricaly/recruitment/internal/metrics"
	"github.com/empiricaly/recruitment/internal/store"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config is `tawon agent` command line configuration
type Config struct {
	Store   *store.Config   `mapstructure:"store"`
	Metrics *metrics.Config `mapstructure:"metrics"`
	Logger  *log.Config     `mapstructure:"log"`

	GQLAddr string `mapstructure:"gqladdr"`
}

// Validate configuration is ok
func (c *Config) Validate() error {
	if err := c.Store.Validate(); err != nil {
		return errors.Wrap(err, "profiler config error")
	}

	if err := c.Metrics.Validate(); err != nil {
		return errors.Wrap(err, "metrics config error")
	}

	if err := c.Logger.Validate(); err != nil {
		return errors.Wrap(err, "logger config error")
	}

	return nil
}

// ConfigFlags helps configure cobra and viper flags.
func ConfigFlags(cmd *cobra.Command) error {
	store.ConfigFlags(cmd, "store")
	metrics.ConfigFlags(cmd, "", "recruitment", ":9999", "")
	log.ConfigFlags(cmd, "", "")

	flag := "gqladdr"
	val := ":8880"
	cmd.Flags().String(flag, val, "GraQL API server address")
	viper.SetDefault(flag, val)

	return nil
}
