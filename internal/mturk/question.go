package mturk

import "encoding/xml"

func getExternalQuestion(url string) (string, error) {
	return (&externalQuestion{
		ExternalURL: url,
	}).xmlEncode()
}

type externalQuestion struct {
	XMLName xml.Name `xml:"http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2006-07-14/ExternalQuestion.xsd ExternalQuestion"`

	// The URL of your web form, to be displayed in a frame in the Worker's web
	// browser. This URL must use the HTTPS protocol.
	// Amazon Mechanical Turk appends the following parameters to this URL:
	// assignmentId, hitId, turkSubmitTo, and workerId. For more information
	// about these appended parameters, see the sections following this table.
	ExternalURL string

	// The height of the frame, in pixels.
	// If you set the value to 0, your HIT will automatically resize to fit
	// within the Worker's browser window.
	// This defaults to 0 as the default Go value, and that's probably what we
	// want.
	FrameHeight int
}

func (e *externalQuestion) xmlEncode() (s string, err error) {
	var buf []byte
	buf, err = xml.Marshal(e)
	return string(buf), err
}
