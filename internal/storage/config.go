package storage

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config is store configuration
type Config struct {
	File  string `mapstructure:"file"`
	Debug bool   `mapstructure:"debug"`

	Driver    string `mapstructure:"driver"`
	DriverURI string `mapstructure:"driveruri"`
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

	flag = prefix + ".driver"
	val = "sqlite3"
	cmd.Flags().String(flag, val, "DB Driver")
	viper.SetDefault(flag, val)

	flag = prefix + ".driveruri"
	val = ""
	cmd.Flags().String(flag, val, "Custom DB connection uri (golang sql format)")
	viper.SetDefault(flag, val)

	flag = prefix + ".debug"
	bval := false
	cmd.Flags().Bool(flag, bval, "show sql queries debug logs")
	viper.SetDefault(flag, bval)

	return nil
}
