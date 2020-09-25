package server

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/empiricaly/recruitment/internal/admin"
	logger "github.com/empiricaly/recruitment/internal/log"
	"github.com/empiricaly/recruitment/internal/metrics"
	"github.com/empiricaly/recruitment/internal/mturk"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Server encapsulates the state of running the server
type Server struct {
	ctx       context.Context
	config    *Config
	storeConn *storage.Conn

	mturk   *mturk.Session
	metrics *metrics.Metrics
	done    chan struct{}
	wg      sync.WaitGroup
}

// Run starts the server with the given configuration
func Run(ctx context.Context, config *Config) (err error) {
	s := Server{
		ctx:    ctx,
		config: config,
		done:   make(chan struct{}),
	}

	err = logger.Init(config.Logger)
	if err != nil {
		return errors.Wrap(err, "could not initialize logger")
	}

	if e := log.Debug(); e.Enabled() {
		log.Debug().Interface("config", config).Msg("")
	}

	s.metrics, err = metrics.New(config.Metrics)
	if err != nil {
		return errors.Wrap(err, "could not initialize metrics")
	}

	s.storeConn, err = storage.Connect(config.Store)
	if err != nil {
		return errors.Wrap(err, "store err")
	}

	if config.DevMode {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP)

		// go func() {
		// 	for range sig {
		// 		log.Warn().Msg("Clearing DB")
		// 		dropStart := time.Now()
		// 		err = s.storeConn.DropAll()
		// 		if err != nil {
		// 			log.Error().Err(err).Msg("Failed to clear DB\n")
		// 		} else {
		// 			log.Warn().Fields(map[string]interface{}{"took": time.Since(dropStart).String()}).Msg("Cleared DB")
		// 		}
		// 	}
		// }()
	}

	err = admin.Init(ctx, s.config.Admins, s.storeConn)
	if err != nil {
		return errors.Wrap(err, "init admins")
	}

	s.mturk, err = mturk.New(config.MTurkConfig)
	if err != nil {
		return errors.Wrap(err, "init mturk")
	}

	s.startGraphqlServer()

	err = s.start()
	if err != nil {
		return errors.Wrap(err, "could not subscribe to configuration topic")
	}

	log.Info().Msg("Server started")

	<-s.ctx.Done()

	close(s.done)
	s.wg.Wait()

	return s.storeConn.Close()
}

func (s *Server) start() error {

	return nil
}
