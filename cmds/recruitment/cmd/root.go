/*
Copyright © 2020 Nicolas Paton <nicolas@paton.dev>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/empiricaly/recruitment/internal/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var cfgRead bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "recruitment",
	Short: "Empirica recruitment facilates MTurk based experiment recruitment",
	Run: func(cmd *cobra.Command, args []string) {
		conf := new(server.Config)
		err := viper.Unmarshal(conf)
		// spew.Dump(cmd)
		// spew.Dump(conf)
		if err != nil {
			log.Fatal().Err(err).Msg("could not parse configuration")
		}

		err = conf.Validate()
		if err != nil {
			log.Fatal().Err(err).Msg("invalid config")
		}

		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			s := make(chan os.Signal, 1)
			signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			<-s
			cancel()
		}()

		err = server.Run(ctx, conf)
		if err != nil {
			log.Fatal().Err(err).Msg("failed starting server")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	server.ConfigFlags(rootCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")
	viper.BindPFlags(rootCmd.Flags())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("recruitment")
	}

	viper.SetEnvPrefix("recruitment")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		cfgRead = true
	}
}
