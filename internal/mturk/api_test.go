package mturk

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/empiricaly/recruitment/internal/log"
)

var sess *Session

func init() {
	log.Init(&log.Config{
		Level:    "debug",
		ForceTTY: true,
	})
}

func getSession() *Session {
	if sess == nil {
		var err error
		config := &Config{
			Sandbox: true,
		}
		sess, err = New(config)
		if err != nil {
			panic(err)
		}
	}

	return sess
}

func TestCreateHIT(t *testing.T) {
	ctx := context.Background()
	question, err := getExternalQuestion("https://www.example.com")
	if err != nil {
		t.Error("could not encode external question", err)
		return
	}

	params := &mturk.CreateHITInput{
		Title:                       aws.String("Testing 12"),
		Description:                 aws.String("Testing 12"),
		AssignmentDurationInSeconds: aws.Int64(120),
		LifetimeInSeconds:           aws.Int64(300),
		Reward:                      aws.String("1"),
		Question:                    aws.String(question),
		QualificationRequirements: []*mturk.QualificationRequirement{
			{
				ActionsGuarded: aws.String("DiscoverPreviewAndAccept"),
			},
		},
	}
	_, err = getSession().createHit(ctx, params)
	if err != nil {
		t.Error("could not create hit", err)
		return
	}

	//
	// add getHIT by ID,verify it was create correctly
	//
}
