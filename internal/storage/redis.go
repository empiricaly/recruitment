package storage

import (
	"github.com/philippgille/gokv/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// // ConnectRedis creates a connection to Redis with the given config.
// func ConnectRedis(config *Config) (*Conn, error) {
// 	if err := config.Redis.Validate(); err != nil {
// 		return nil, errors.Wrap(err, "redis config error")
// 	}
// 	conf := config.Redis

// 	client, err := redis.NewClient(redis.Options{
// 		Address:  conf.Address,
// 		Password: conf.Password,
// 		DB:       conf.DB,
// 	})
// 	if err != nil {
// 		return nil, errors.Wrap(err, "badger failed conn")
// 	}

// 	return &Conn{Store: client}, nil
// }

// RedisConfig contains values need to configure RedisDB.
type RedisConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Address of the Redis server, including the port.
	// Optional ("localhost:6379" by default).
	Address string `mapstructure:"addr"`
	// Password for the Redis server.
	// Optional ("" by default).
	Password string `mapstructure:"password"`
	// DB to use.
	// Optional (0 by default).
	DB int `mapstructure:"db"`
}

func redisConfigFlags(cmd *cobra.Command, prefix string) error {
	if prefix == "" {
		prefix = "redis"
	}

	viper.SetDefault(prefix, &RedisConfig{})

	flag := prefix + ".addr"
	val := redis.DefaultOptions.Address
	cmd.Flags().String(flag, val, "Redis server address")
	viper.SetDefault(flag, val)

	flag = prefix + ".password"
	val = ""
	cmd.Flags().String(flag, val, "Redis server password")
	viper.SetDefault(flag, val)

	flag = prefix + ".db"
	valI := 0
	cmd.Flags().Int(flag, valI, "Redis DB to use")
	viper.SetDefault(flag, valI)

	flag = prefix + ".enabled"
	cmd.Flags().Bool(flag, false, "Use Redis")
	viper.SetDefault(flag, false)

	return nil
}

// Validate configuration is ok
func (h *RedisConfig) Validate() error {

	return nil
}
