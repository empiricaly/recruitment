package runtime

import (
	"bufio"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"

	rice "github.com/GeertJohan/go.rice"
	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var participantFilterJSFile = ""

func jsfilter(ctx context.Context, conn *storage.Conn, participants []*ent.Participant, js string) ([]*ent.Participant, error) {
	if participantFilterJSFile == "" {
		box, err := rice.FindBox("scripts")
		if err != nil {
			return nil, errors.Wrap(err, "opening directory rice for filter.js file")
		}

		content, err := box.Bytes("filter.js")
		f, err := ioutil.TempFile("", "empirica_recruitment_js_*.js")
		if err != nil {
			return nil, errors.Wrap(err, "creating temporary filter.js file")
		}

		_, err = f.Write(content)
		if err != nil {
			return nil, errors.Wrap(err, "writing temporary filter.js file")
		}

		err = f.Close()
		if err != nil {
			return nil, errors.Wrap(err, "closing temporary filter.js file")
		}

		participantFilterJSFile = f.Name()
	}

	subProcess := exec.Command("node", participantFilterJSFile)
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

	if _, err = input.WriteString(js); err != nil {
		return nil, errors.Wrap(err, "writing js function to node")
	}

	if _, err = input.WriteString("\nEND_OF_JS\n"); err != nil {
		return nil, errors.Wrap(err, "writing end of js to node")
	}

	if _, err = input.WriteString("\nEND_OF_ARGS\n"); err != nil {
		return nil, errors.Wrap(err, "writing end of args to node")
	}

	if err = input.Flush(); err != nil {
		return nil, errors.Wrap(err, "flush args to node")
	}

	for _, p := range participants {
		b, err := json.Marshal(p)
		if err != nil {
			return nil, errors.Wrap(err, "JSON encode participant")
		}

		if _, err = input.Write(b); err != nil {
			return nil, errors.Wrap(err, "writing participant to node")
		}

		if _, err = input.WriteString("\n"); err != nil {
			return nil, errors.Wrap(err, "writing new line to node")
		}

		if err = input.Flush(); err != nil {
			return nil, errors.Wrap(err, "flush participant to node")
		}
	}

	if err = stdin.Close(); err != nil {
		return nil, errors.Wrap(err, "finished sending input to node")
	}

	changes := make([]*participantChanges, 0, len(participants))
	for output.Scan() {
		line := output.Bytes()
		change := &participantChanges{}
		err = json.Unmarshal(line, change)
		if err != nil {
			log.Error().Err(err).Msg("unmarshal json from node")
			continue
		}
		changes = append(changes, change)
	}
	if err := output.Err(); err != nil {
		return nil, errors.Wrap(err, "reading output from node process")
	}

	subProcess.Wait()

	participantsMap := make(map[string]*ent.Participant)
	for _, p := range participants {
		participantsMap[p.ID] = p
	}

	outputParticipants := make([]*ent.Participant, 0)

	if err := ent.WithTx(ctx, conn.Client, func(tx *ent.Tx) error {
		for _, change := range changes {
			if change.Keep {
				p, ok := participantsMap[change.ID]
				if !ok {
					log.Warn().Str("ID", change.ID).Msg("returned unknown filtered participant")
				} else {
					outputParticipants = append(outputParticipants, p)
				}
			}

			if len(change.Changes) == 0 {
				continue
			}

			for key, val := range change.Changes {
				_, err := ent.SetDatum(ctx, tx, participantsMap[change.ID], key, val, false)
				if err != nil {
					return errors.Wrap(err, "set datum")
				}
			}
		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "set data on fitered participant: commit transaction")
	}

	return outputParticipants, nil
}

type participantChanges struct {
	ID      string
	Changes map[string]string
	Keep    bool
}
