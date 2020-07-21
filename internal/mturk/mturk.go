package mturk

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/pkg/errors"
)

const productionURL = "https://mturk-requester.us-east-1.amazonaws.com"
const sandboxURL = "https://mturk-requester-sandbox.us-east-1.amazonaws.com"
const awsRegion = "us-east-1"

// Session keep the mTurk session state.
type Session struct {
	*mturk.MTurk
}

// New create a new session for mTurk
func New(config *Config) (*Session, error) {
	var endpoint string
	if config.Sandbox {
		endpoint = sandboxURL
	} else {
		endpoint = productionURL
	}

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(awsRegion),
		Endpoint: aws.String(endpoint),
	})

	if err != nil {
		return nil, errors.Wrap(err, "create new aws sessoin")
	}

	svc := mturk.New(sess)
	return &Session{
		MTurk: svc,
	}, nil
}
