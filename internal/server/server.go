package server

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/empiricaly/recruitment/internal/admin"
	logger "github.com/empiricaly/recruitment/internal/log"
	"github.com/empiricaly/recruitment/internal/metrics"
	"github.com/empiricaly/recruitment/internal/mturk"
	"github.com/empiricaly/recruitment/internal/runtime"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Server encapsulates the state of running the server
type Server struct {
	ctx       context.Context
	config    *Config
	storeConn *storage.Conn

	mturk        *mturk.Session
	mturkSandbox *mturk.Session
	runtime      *runtime.Runtime
	metrics      *metrics.Metrics
	done         chan struct{}
	wg           sync.WaitGroup
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

	if config.DevMode {
		log.Debug().Interface("config", config).Msg("Configuration")
	}

	s.metrics, err = metrics.New(config.Metrics)
	if err != nil {
		return errors.Wrap(err, "could not initialize metrics")
	}

	s.storeConn, err = storage.Connect(ctx, config.Store)
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

	s.mturk, err = mturk.New(config.MTurkConfig, false, config.HTTP.RootURL, s.storeConn)
	if err != nil {
		return errors.Wrap(err, "init mturk")
	}

	s.mturkSandbox, err = mturk.New(config.MTurkConfig, true, config.HTTP.RootURL, s.storeConn)
	if err != nil {
		return errors.Wrap(err, "init mturk")
	}

	s.runtime, err = runtime.Start(config.Runtime, s.storeConn, s.mturk, s.mturkSandbox)
	if err != nil {
		return errors.Wrap(err, "init runtime")
	}

	s.startHTTPServer()

	err = s.start()
	if err != nil {
		return errors.Wrap(err, "could not subscribe to configuration topic")
	}

	addr := config.HTTP.Addr
	if strings.HasPrefix(addr, ":") {
		addr = "https://localhost" + addr
	}
	log.Info().Msgf("Recruitment started at %s", addr)

	<-s.ctx.Done()

	s.runtime.Stop()
	close(s.done)
	s.wg.Wait()

	return s.storeConn.Close()
}

func (s *Server) start() error {

	return nil
}
