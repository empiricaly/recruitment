package storage

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v2"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/pkg/errors"
)

type Mapping struct {
	store Store
}

type MappingTxn struct {
	m *Mapping
	t Transaction
}

func NewMapping(store Store) (*Mapping, error) {
	return &Mapping{store: store}, nil
}

func (m *Mapping) Txn(f func(*MappingTxn) error) error {
	return m.store.Txn(func(t Transaction) error {
		txn := &MappingTxn{m, t}
		return f(txn)
	}, true)
}

const collPrefix = "coll_"

func (m *MappingTxn) setKeys(name string, keys [][]byte) error {
	keysJ, err := json.Marshal(keys)
	if err != nil {
		return err
	}

	return m.t.Set([]byte(collPrefix+name), keysJ)
}

func (m *MappingTxn) getKeys(name string) ([][]byte, error) {
	item, err := m.t.Get([]byte(collPrefix + name))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return [][]byte{}, nil
		}
		return nil, err
	}

	keysJ, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	keys := [][]byte{}
	err = json.Unmarshal(keysJ, &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

// Project return a Project by ID.
func (m *MappingTxn) Project(id string) (*model.Project, error) {
	item, err := m.t.Get([]byte(id))
	if err != nil {
		return nil, err
	}

	record, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	project := &model.Project{}
	err = json.Unmarshal(record, project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// ProjectByProjectID returns a Project by projectID.
func (m *MappingTxn) ProjectByProjectID(projectID string) (*model.Project, error) {
	projects, err := m.Projects()
	if err != nil {
		return nil, err
	}
	for _, project := range projects {
		if project.ProjectID == projectID {
			return project, nil
		}
	}

	return nil, errors.New("Project not found")
}

// Projects returns all Projects.
func (m *MappingTxn) Projects() ([]*model.Project, error) {
	keys, err := m.getKeys("projects")
	if err != nil {
		return nil, err
	}

	projects := make([]*model.Project, len(keys))

	for i, key := range keys {
		item, err := m.t.Get([]byte(key))
		if err != nil {
			return nil, err
		}

		keysJ, err := item.ValueCopy(nil)
		if err != nil {
			return nil, err
		}

		project := &model.Project{}
		err = json.Unmarshal(keysJ, project)
		if err != nil {
			return nil, err
		}
		projects[i] = project
	}

	return projects, nil
}

// AddProject saves a new project
func (m *MappingTxn) AddProject(project *model.Project) error {
	keys, err := m.getKeys("projects")
	if err != nil {
		return err
	}

	for _, key := range keys {
		if string(key) == project.ID {
			return errors.New("project ID already exists")
		}
	}

	keys = append(keys, []byte(project.ID))

	out, err := json.Marshal(project)
	if err != nil {
		return err
	}

	err = m.t.Set([]byte(project.ID), out)
	if err != nil {
		return err
	}

	return m.setKeys("projects", keys)
}
