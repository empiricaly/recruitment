package server

import (
	"fmt"
	"strings"

	"github.com/empiricaly/recruitment/internal/log"
	"github.com/empiricaly/recruitment/internal/metrics"
	"github.com/empiricaly/recruitment/internal/mturk"
	"github.com/empiricaly/recruitment/internal/store"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const minPasswordSize = 8

type admin struct {
	Name     string
	Username string
	Password string
}

// Config is `tawon agent` command line configuration
type Config struct {
	GQLAddr     string        `mapstructure:"gqladdr"`
	MTurkConfig *mturk.Config `mapstructure:"mturk"`
	Admins      []admin       `mapstructure:"admins"`

	Store   *store.Config   `mapstructure:"store"`
	Metrics *metrics.Config `mapstructure:"metrics"`
	Logger  *log.Config     `mapstructure:"log"`
}

// Validate configuration is ok
func (c *Config) Validate() error {
	for _, admin := range c.Admins {
		if strings.TrimSpace(admin.Name) == "" {
			return errors.New("admin name is required")
		}
		if strings.TrimSpace(admin.Username) == "" {
			return errors.New("admin username is required")
		}
		if strings.TrimSpace(admin.Password) == "" {
			return errors.New("admin password is required")
		}
		if len(strings.TrimSpace(admin.Password)) < minPasswordSize {
			return errors.New(fmt.Sprintf("admin password is too small (%d chars min)", minPasswordSize))
		}
	}

	if len(c.Admins) == 0 {
		return errors.New("please add at least one admin in the configuration")
	}

	if err := c.Store.Validate(); err != nil {
		return errors.Wrap(err, "profiler config error")
	}

	if err := c.Metrics.Validate(); err != nil {
		return errors.Wrap(err, "metrics config error")
	}

	if err := c.Logger.Validate(); err != nil {
		return errors.Wrap(err, "logger config error")
	}

	if err := c.MTurkConfig.Validate(); err != nil {
		return errors.Wrap(err, "mturk config error")
	}

	return nil
}

// ConfigFlags helps configure cobra and viper flags.
func ConfigFlags(cmd *cobra.Command) error {
	store.ConfigFlags(cmd, "store")
	metrics.ConfigFlags(cmd, "", "recruitment", ":9999", "")
	log.ConfigFlags(cmd, "", "")
	mturk.ConfigFlags(cmd, "")

	flag := "gqladdr"
	val := ":8880"
	cmd.Flags().String(flag, val, "GraQL API server address")
	viper.SetDefault(flag, val)

	return nil
}
