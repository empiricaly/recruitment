package mturk

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const productionURL = "https://mturk-requester.us-east-1.amazonaws.com"
const sandboxURL = "https://mturk-requester-sandbox.us-east-1.amazonaws.com"
const awsRegion = "us-east-1"

// Session keep the mTurk session state.
type Session struct {
	quals []*model.MTurkQulificationType
	*mturk.MTurk
	rootURL string
	config  *Config
	store   *storage.Conn
	sandbox bool

	logger zerolog.Logger
}

// New create a new session for mTurk
func New(config *Config, sandbox bool, rootURL string, store *storage.Conn) (*Session, error) {
	var endpoint string
	if sandbox {
		endpoint = sandboxURL
	} else {
		endpoint = productionURL
	}

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(awsRegion),
		Endpoint: aws.String(endpoint),
	})
	if err != nil {
		return nil, errors.Wrap(err, "create new aws session")
	}

	quals, err := loadQuals(sandbox)
	if err != nil {
		return nil, errors.Wrap(err, "load json quals")
	}

	svc := mturk.New(sess)
	return &Session{
		config:  config,
		rootURL: rootURL,
		sandbox: sandbox,
		MTurk:   svc,
		quals:   quals,
		store:   store,
		logger:  log.With().Str("pkg", "mturk").Logger(),
	}, nil
}
