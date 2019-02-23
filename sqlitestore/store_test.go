package sqlitestore

import (
	"os"
	"testing"

	"../model"
)

func getStore() *SqliteStore {
	dbPath := "./test.db"
	os.Remove(dbPath)

	store, _ := NewSqliteStore(dbPath)
	return store
}

func TestCreateDB(t *testing.T) {
	dbPath := "./test.db"
	os.Remove(dbPath)

	_, err := NewSqliteStore(dbPath)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveBuild(t *testing.T) {

	store := getStore()

	testBuild := &model.BuildHistory{}

	err := store.SaveBuild(testBuild)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBuild(t *testing.T) {

	store := getStore()

	testBuild := &model.BuildHistory{ID: "test", Output: "Bonjour"}
	store.SaveBuild(testBuild)

	h, err := store.GetBuild("test")

	if err != nil {
		t.Error(err)
	}

	if h.ID != "test" {
		t.Fail()
	}

	if h.Output != "Bonjour" {
		t.Fail()
	}

}

func TestGetLastBuilds(t *testing.T) {
	store := getStore()

	testBuilds := []*model.BuildHistory{
		&model.BuildHistory{ID: "A1", ProjectName: "A", Start: 5},
		&model.BuildHistory{ID: "A2", ProjectName: "A", Start: 4},
		&model.BuildHistory{ID: "B1", ProjectName: "B", Start: 9},
		&model.BuildHistory{ID: "B2", ProjectName: "B", Start: 7},
		&model.BuildHistory{ID: "C1", ProjectName: "C", Start: 3},
		&model.BuildHistory{ID: "C2", ProjectName: "C", Start: 2},
	}

	for _, b := range testBuilds {
		store.SaveBuild(b)
	}

	latests, err := store.GetLatestBuilds()
	if err != nil {
		t.Error(err)
	}
	if len(latests) != 3 {
		t.Error("Rows count should be 3. It is", len(latests))
	}

	if latests[0].ID != "B1" {
		t.Fail()
	}

	if latests[1].ID != "A1" {
		t.Fail()
	}

	if latests[2].ID != "C1" {
		t.Fail()
	}

}

func TestGetAllBuilds(t *testing.T) {

	store := getStore()

	testBuilds := []*model.BuildHistory{
		&model.BuildHistory{ID: "A1", ProjectName: "A", Start: 5},
		&model.BuildHistory{ID: "A2", ProjectName: "A", Start: 4},
		&model.BuildHistory{ID: "B1", ProjectName: "B", Start: 9},
		&model.BuildHistory{ID: "B2", ProjectName: "B", Start: 7},
		&model.BuildHistory{ID: "C1", ProjectName: "C", Start: 3},
		&model.BuildHistory{ID: "C2", ProjectName: "C", Start: 2},
	}

	for _, b := range testBuilds {
		store.SaveBuild(b)
	}

	all, err := store.GetAllBuilds("B")
	if err != nil {
		t.Error(err)
	}
	if len(all) != 2 {
		t.Error("Rows count should be 2. It is", len(all))
	}

	if all[0].ID != "B1" {
		t.Fail()
	}

	if all[1].ID != "B2" {
		t.Fail()
	}
}
