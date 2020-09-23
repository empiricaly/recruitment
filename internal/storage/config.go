package storage

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config is kvstore configuration
type Config struct {
	File  string `mapstructure:"file"`
	Debug bool   `mapstructure:"debug"`
}

// Validate configuration is ok
func (c *Config) Validate() error {

	return nil
}

// ConfigFlags helps configure cobra and viper flags.
func ConfigFlags(cmd *cobra.Command, prefix string) error {
	if prefix == "" {
		prefix = "store"
	}

	viper.SetDefault(prefix, &Config{})

	flag := prefix + ".file"
	val := "recruitment.db"
	cmd.Flags().String(flag, val, "Sqlite3 database file")
	viper.SetDefault(flag, val)

	flag = prefix + ".debug"
	bval := false
	cmd.Flags().Bool(flag, bval, "debug sql queries")
	viper.SetDefault(flag, bval)

	return nil
}
