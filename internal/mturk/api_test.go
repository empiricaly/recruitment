package mturk

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/empiricaly/recruitment/internal/log"
)

var sess *Session
var createdHitID string
var title string = "Testing 12"
var description string = "Testing 12 description"

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
		t.Fatal("could not encode external question", err)
	}

	params := &mturk.CreateHITInput{
		Title:                       aws.String(title),
		Description:                 aws.String(description),
		AssignmentDurationInSeconds: aws.Int64(120),
		LifetimeInSeconds:           aws.Int64(300),
		Reward:                      aws.String("1"),
		Question:                    aws.String(question),
	}

	hitID, err := getSession().createHit(ctx, params)
	if err != nil {
		t.Fatal("could not create hit", err)
	}

	// GetHit
	hit, err := getSession().getHit(ctx, hitID)
	if err != nil {
		t.Fatal("could not get created HitID: ", hitID, err)
	}

	// Verify  Hit
	if *hit.Title != title || *hit.Description != description {
		t.Fatal("HIT data doesn't match")
	}

	// StopHit
	err = getSession().stopHit(ctx, hitID)
	if err != nil {
		t.Fatal("could not stop created HitID: ", hitID, err)
	}

	// GetHit
	stoppedHit, err := getSession().getHit(ctx, hitID)
	if err != nil {
		t.Fatal("could not get created HitID: ", hitID, err)
	}

	t.Log("stoppedHit ", *stoppedHit)
	// Check if hit cannot be assigned
	if *stoppedHit.HITStatus == "Assignable" {
		t.Fatal("Failed to stop the HitID: ", hitID, err)
	}
}

var qualName string = "qual test 123"
var qualAnswerKey string = "answer"
var qualDescription string = "qualification description"

func TestCreateQual(t *testing.T) {
	ctx := context.Background()

	params := &mturk.CreateQualificationTypeInput{
		AnswerKey:               aws.String(qualAnswerKey),
		AutoGranted:             aws.Bool(true),
		AutoGrantedValue:        aws.Int64(10),
		Description:             aws.String(qualDescription),
		Keywords:                aws.String("test qualification"),
		Name:                    aws.String(qualName),
		QualificationTypeStatus: aws.String("Active"),
		RetryDelayInSeconds:     aws.Int64(20),
		TestDurationInSeconds:   aws.Int64(60),
	}

	qualID, err := getSession().createQualificationType(ctx, params)
	if err != nil {
		t.Fatal("could not create qualification", err)
	}

	// Get qualification type
	qualificationType, err := getSession().getQualificationType(ctx, qualID)

	// Check if qualifaction type is match
	if *qualificationType.Name != qualName {
		t.Fatal("qualificationType does not match", err)
	}

	// associate worker with qualification
	err = getSession().associateQualificationWithWorker(ctx,
		&mturk.AssociateQualificationWithWorkerInput{
			QualificationTypeId: aws.String(qualID),
			IntegerValue:        aws.Int64(1),
			SendNotification:    aws.Bool(false),
			WorkerId:            aws.String("A1R5P9HWU2CDUT"),
		})
	if err != nil {
		t.Fatal("could not associate qualification with worker", err)
	}

	// Delete Qualification Type
	err = getSession().deleteQualificationType(ctx, qualID)
	if err != nil {
		t.Fatal("could not delete qualificationType", err)
	}
}
