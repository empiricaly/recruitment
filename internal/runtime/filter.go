package runtime

import (
	"bufio"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/davecgh/go-spew/spew"
	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func jsfilter(participants []*ent.Participant, js string) ([]*ent.Participant, error) {
	subProcess := exec.Command("node", "../../scripts/participant_filter/filter.js")
	subProcess.Stderr = os.Stderr

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		return nil, errors.Wrap(err, "opening stdin on node process")
	}
	input := bufio.NewWriter(stdin)

	stdout, err := subProcess.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "opening stdout on node process")
	}
	output := bufio.NewScanner(stdout)

	if err = subProcess.Start(); err != nil {
		return nil, errors.Wrap(err, "running node process")
	}

	_, err = input.WriteString(js)
	if err != nil {
		return nil, errors.Wrap(err, "writing js function to node")
	}

	_, err = input.WriteString("\nEND_OF_JS\n")
	if err != nil {
		return nil, errors.Wrap(err, "writing end of js to node")
	}

	_, err = input.WriteString("\nEND_OF_ARGS\n")
	if err != nil {
		return nil, errors.Wrap(err, "writing end of args to node")
	}

	err = input.Flush()
	if err != nil {
		return nil, errors.Wrap(err, "flush args to node")
	}

	for _, p := range participants {
		b, err := json.Marshal(p)
		if err != nil {
			return nil, errors.Wrap(err, "JSON encode participant")
		}

		_, err = input.Write(b)
		if err != nil {
			return nil, errors.Wrap(err, "writing participant to node")
		}

		_, err = input.WriteString("\n")
		if err != nil {
			return nil, errors.Wrap(err, "writing new line to node")
		}

		err = input.Flush()
		if err != nil {
			return nil, errors.Wrap(err, "flush participant to node")
		}

	}

	err = stdin.Close()
	if err != nil {
		return nil, errors.Wrap(err, "finished sending input to node")
	}

	for output.Scan() {
		line := output.Bytes()
		p := &participantChanges{}
		err = json.Unmarshal(line, p)
		if err != nil {
			log.Error().Err(err).Msg("unmarshal json from node")
			continue
		}
		spew.Dump(p)
	}
	if err := output.Err(); err != nil {
		return nil, errors.Wrap(err, "reading output from node process")
	}

	subProcess.Wait()

	return nil, nil
}

type participantChanges struct {
	ID      string
	Changes map[string]interface{}
	Keep    bool
}
