// This is a stub, need to finish

package storage_testing

import (
	"testing"

	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/enttest"
	"github.com/empiricaly/recruitment/internal/ent/migrate"

	_ "github.com/mattn/go-sqlite3"
)

func Setup(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
}

func Teardown(t *testing.T) {
	client.Close()
}
