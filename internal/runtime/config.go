package runtime

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config is runtime configuration
type Config struct {
	Disable bool `mapstructure:"disable"`
	Debug   bool `mapstructure:"debug"`
}

// Validate configuration is ok
func (c *Config) Validate() error {

	return nil
}

// ConfigFlags helps configure cobra and viper flags.
func ConfigFlags(cmd *cobra.Command, prefix string) error {
	if prefix == "" {
		prefix = "runtime"
	}

	viper.SetDefault(prefix, &Config{})

	flag := prefix + ".disable"
	bval := false
	cmd.Flags().Bool(flag, bval, "disable the runtime (don't start runs)")
	viper.SetDefault(flag, bval)

	flag = prefix + ".debug"
	bval = false
	cmd.Flags().Bool(flag, bval, "show runtime debug logs")
	viper.SetDefault(flag, bval)

	return nil
}
