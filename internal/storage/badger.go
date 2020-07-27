package storage

import (
	"os"

	"github.com/dgraph-io/badger/v2"
	"github.com/mattn/go-isatty"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// logger is a wrapper of zerolog for badger internal logging
type logger struct {
	log *zerolog.Logger
}

func (l *logger) Errorf(msg string, vars ...interface{}) {
	l.log.Error().Msgf(msg, vars...)
}
func (l *logger) Warningf(msg string, vars ...interface{}) {
	l.log.Warn().Msgf(msg, vars...)
}
func (l *logger) Infof(msg string, vars ...interface{}) {
	l.log.Info().Msgf(msg, vars...)
}
func (l *logger) Debugf(msg string, vars ...interface{}) {
	l.log.Debug().Msgf(msg, vars...)
}

// ConnectBadger creates a connection to Badger with the given config.
func ConnectBadger(config *Config) (*Conn, error) {
	if err := config.Badger.Validate(); err != nil {
		return nil, errors.Wrap(err, "badger config error")
	}
	conf := config.Badger

	level, err := zerolog.ParseLevel(config.Badger.LogLevel)
	if err != nil {
		return nil, errors.Wrap(err, "invalid badger loglevel")
	}
	l := zerolog.New(os.Stderr)

	l = l.Level(level).With().Logger()

	if isatty.IsTerminal(os.Stderr.Fd()) {
		l = l.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Logger()
	}

	// Set default values
	// Open the Badger database located in the options.Dir directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions(conf.Dir)
	opts.Logger = &logger{&l}
	db, err := badger.Open(opts)
	if err != nil {
		return nil, errors.Wrap(err, "badger failed conn")
	}

	return &Conn{Store: &badgerDB{DB: db}}, nil
}

type badgerDB struct {
	*badger.DB
}

type badgerTransaction struct {
	*badger.Txn
}

func (t *badgerTransaction) Get(key []byte) (Item, error) {
	return t.Txn.Get(key)
}

func (db *badgerDB) Txn(t func(Transaction) error, writes bool) error {
	txn := db.NewTransaction(writes)
	defer txn.Discard()
	err := t(&badgerTransaction{Txn: txn})
	if err != nil {
		return err
	}
	return txn.Commit()
}

// BadgerConfig contains values need to configure BadgerDB.
type BadgerConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	LogLevel string `mapstructure:"loglevel"`
	Dir      string `mapstructure:"dir"`
}

func badgerConfigFlags(cmd *cobra.Command, prefix string) error {
	if prefix == "" {
		prefix = "badger"
	}

	viper.SetDefault(prefix, &BadgerConfig{})

	flag := prefix + ".dir"
	val := "./data"
	cmd.Flags().String(flag, val, "BadgerDB storage directory")
	viper.SetDefault(flag, val)

	flag = prefix + ".loglevel"
	val = "warn"
	cmd.Flags().String(flag, val, "BadgerDB specific log level (logs too much)")
	viper.SetDefault(flag, val)

	flag = prefix + ".enabled"
	cmd.Flags().Bool(flag, true, "Use Badger")
	viper.SetDefault(flag, true)

	return nil
}

// Validate configuration is ok
func (h *BadgerConfig) Validate() error {

	return nil
}
