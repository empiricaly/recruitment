package server

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/empiricaly/recruitment/internal/admin"
	"github.com/empiricaly/recruitment/internal/log"
	"github.com/empiricaly/recruitment/internal/metrics"
	"github.com/empiricaly/recruitment/internal/mturk"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const minPasswordSize = 8

// Config is `tawon agent` command line configuration
type Config struct {
	HTTP        HTTPServerConfig `mapstructure:"http"`
	MTurkConfig *mturk.Config    `mapstructure:"mturk"`
	Admins      []admin.User     `mapstructure:"admins"`
	SecretKey   string           `mapstructure:"secret"`
	DevMode     bool             `mapstructure:"dev"`

	Store   *storage.Config `mapstructure:"store"`
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

	if len(c.SecretKey) != 32 {
		return errors.New("please add a random 32 characters secret key")
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
	storage.ConfigFlags(cmd, "store")
	metrics.ConfigFlags(cmd, "", "recruitment", ":9999", "")
	log.ConfigFlags(cmd, "", "")
	mturk.ConfigFlags(cmd, "")
	HTTPServerConfigFlags(cmd, "")

	flag := "dev"
	bval := false
	cmd.Flags().Bool(flag, bval, "Run in Developer Mode (enables extra tooling ; do not run in production!)")
	viper.SetDefault(flag, bval)

	return nil
}

// HTTPServerConfig is `tawon agent` command line configuration
type HTTPServerConfig struct {
	Addr    string `mapstructure:"addr"`
	RootURL string `mapstructure:"rooturl"`
	HTTPS   bool   `mapstructure:"https"`
	AutoTLS bool   `mapstructure:"autotls"`
	Debug   bool   `mapstructure:"debug"`
}

// Validate configuration is ok
func (c *HTTPServerConfig) Validate() error {
	_, err := url.Parse(c.RootURL)
	if err != nil {
		return errors.Wrap(err, "parsing rooturl")
	}

	return nil
}

// HTTPServerConfigFlags helps configure cobra and viper flags.
func HTTPServerConfigFlags(cmd *cobra.Command, prefix string) error {
	if prefix == "" {
		prefix = "http"
	}

	viper.SetDefault(prefix, &HTTPServerConfig{})

	flag := prefix + ".addr"
	val := ":8880"
	cmd.Flags().String(flag, val, "HTTP server address to listen on")
	viper.SetDefault(flag, val)

	flag = prefix + ".rooturl"
	val = "http://localhost:8880"
	cmd.Flags().String(flag, val, "HTTP server external root URL")
	viper.SetDefault(flag, val)

	flag = "debug"
	bval := false
	cmd.Flags().Bool(flag, bval, "debug http requests")
	viper.SetDefault(flag, bval)

	// flag = "autotls"
	// bval = false
	// cmd.Flags().Bool(flag, bval, "automatically generate TLS certificates")
	// viper.SetDefault(flag, bval)

	return nil
}
