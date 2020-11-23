package model

type RenderContext struct {
	URL         string             `handlebars:"url"`
	Step        *RenderStep        `handlebars:"currentStep"`
	Steps       []*RenderStep      `handlebars:"steps"`
	Template    *RenderTemplate    `handlebars:"template"`
	Run         *RenderRun         `handlebars:"run"`
	Participant *RenderParticipant `handlebars:"participant"`
	WorkerID    string             `handlebars:"workedID"`
}

type RenderTemplate struct {
	Adult            bool
	Sandbox          bool
	Name             string
	ParticipantCount int
	SelectionType    string
}

type RenderRun struct {
	Name      string
	StartedAt string
}

type RenderStep struct {
	Index             int
	Duration          int
	Type              string
	ParticipantsCount int
	StartsAt          string
	StartedAt         string
	EndedAt           string
}

type RenderParticipant struct {
	WorkerID     string
	HITID        string `handlebars:"hitID"`
	AssignmentID string
}
