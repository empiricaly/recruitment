package log

import (
	"os"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config configures the logger
type Config struct {
	Level    string `mapstructure:"level"`
	ForceTTY bool   `mapstructure:"forcetty"`
}

// ConfigFlags helps configure cobra and viper flags.
func ConfigFlags(cmd *cobra.Command, prefix, level string) error {
	if prefix == "" {
		prefix = "log"
	}

	viper.SetDefault(prefix, &Config{})

	if level == "" {
		level = "info"
	}

	flag := prefix + ".level"
	cmd.Flags().String(flag, level, "Log level: trace, debug, info, warn, error, fatal or panic")
	viper.SetDefault(flag, level)

	flag = prefix + ".forcetty"
	cmd.Flags().Bool(flag, false, "Force behavior of attached TTY (color, human output)")
	viper.SetDefault(flag, false)

	return nil
}

// Validate configuration is ok
func (c *Config) Validate() error {
	_, err := zerolog.ParseLevel(c.Level)
	return err
}

var std *zerolog.Logger

// Init configures the global logger
func Init(config *Config) error {

	level, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(level)

	if config.ForceTTY || isatty.IsTerminal(os.Stderr.Fd()) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return nil
}
