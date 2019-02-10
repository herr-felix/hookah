package sqlitestore

import (
	"os"
	"testing"

	"../model"
)

func TestSaveBuild(t *testing.T) {
	dbPath := "./test.db"
	os.Remove(dbPath)

	store, err := NewSqliteStore(dbPath)
	if err != nil {
		t.Error(err)
	}

	testBuild := &model.BuildHistory{}

	store.SaveBuild(testBuild)
	if err != nil {
		t.Error(err)
	}
}
