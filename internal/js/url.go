package js

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"

	rice "github.com/GeertJohan/go.rice"
	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/pkg/errors"
)

var urlJSFile = ""

func UrlJS(participant *ent.Participant, currentStep *ent.StepRun, steps []*ent.Step, run *ent.Run, js string) (string, error) {
	if urlJSFile == "" {
		box, err := rice.FindBox("scripts")
		if err != nil {
			return "", errors.Wrap(err, "opening directory rice for url.js file")
		}

		content, err := box.Bytes("url.js")
		f, err := ioutil.TempFile("", "empirica_recruitment_url_js_*.js")
		if err != nil {
			return "", errors.Wrap(err, "creating temporary url.js file")
		}

		_, err = f.Write(content)
		if err != nil {
			return "", errors.Wrap(err, "writing temporary url.js file")
		}

		err = f.Close()
		if err != nil {
			return "", errors.Wrap(err, "closing temporary url.js file")
		}

		urlJSFile = f.Name()
	}

	subProcess := exec.Command("node", urlJSFile)
	subProcess.Stderr = os.Stderr

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		return "", errors.Wrap(err, "opening stdin on node process")
	}
	input := bufio.NewWriter(stdin)

	stdout, err := subProcess.StdoutPipe()
	if err != nil {
		return "", errors.Wrap(err, "opening stdout on node process")
	}
	output := bufio.NewScanner(stdout)

	if err = subProcess.Start(); err != nil {
		return "", errors.Wrap(err, "running node process")
	}

	if _, err = input.WriteString(js); err != nil {
		return "", errors.Wrap(err, "writing js function to node")
	}

	if _, err = input.WriteString("\nEND_OF_JS\n"); err != nil {
		return "", errors.Wrap(err, "writing end of js to node")
	}

	// start insert participant
	p, err := json.Marshal(participant)
	if err != nil {
		return "", errors.Wrap(err, "JSON encode participant")
	}

	if _, err = input.Write(p); err != nil {
		return "", errors.Wrap(err, "writing participant to node")
	}

	if _, err = input.WriteString("\n"); err != nil {
		return "", errors.Wrap(err, "writing new line to node")
	}

	if err = input.Flush(); err != nil {
		return "", errors.Wrap(err, "flush participant to node")
	}
	// End insert participant

	// start insert run
	r, err := json.Marshal(run)
	if err != nil {
		return "", errors.Wrap(err, "JSON encode run")
	}

	if _, err = input.Write(r); err != nil {
		return "", errors.Wrap(err, "writing run to node")
	}

	if _, err = input.WriteString("\n"); err != nil {
		return "", errors.Wrap(err, "writing new line to node")
	}

	if err = input.Flush(); err != nil {
		return "", errors.Wrap(err, "flush run to node")
	}
	// End insert run

	// start insert currentStep
	s, err := json.Marshal(currentStep)
	if err != nil {
		return "", errors.Wrap(err, "JSON encode currentStep")
	}

	if _, err = input.Write(s); err != nil {
		return "", errors.Wrap(err, "writing currentStep to node")
	}

	if _, err = input.WriteString("\n"); err != nil {
		return "", errors.Wrap(err, "writing new line to node")
	}

	if err = input.Flush(); err != nil {
		return "", errors.Wrap(err, "flush currentStep to node")
	}
	// End insert currentStep

	if _, err = input.WriteString("\nEND_OF_ARGS\n"); err != nil {
		return "", errors.Wrap(err, "writing end of args to node")
	}

	if err = input.Flush(); err != nil {
		return "", errors.Wrap(err, "flush args to node")
	}

	for _, s := range steps {
		b, err := json.Marshal(s)
		if err != nil {
			return "", errors.Wrap(err, "JSON encode step")
		}

		if _, err = input.Write(b); err != nil {
			return "", errors.Wrap(err, "writing step to node")
		}

		if _, err = input.WriteString("\n"); err != nil {
			return "", errors.Wrap(err, "writing new line to node")
		}

		if err = input.Flush(); err != nil {
			return "", errors.Wrap(err, "flush step to node")
		}
	}

	if err = stdin.Close(); err != nil {
		return "", errors.Wrap(err, "finished sending input to node")
	}

	var url string
	for output.Scan() {
		line := output.Bytes()
		url = string(line)
	}
	if err := output.Err(); err != nil {
		return "", errors.Wrap(err, "reading output from node process")
	}

	subProcess.Wait()
	return url, nil
}
