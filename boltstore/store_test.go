package boltstore

import (
	"os"
	"testing"

	"../model"
)

func getStore() *BoltStore {
	dbPath := "./test.db"
	os.Remove(dbPath)

	store, err := NewBoltStore(dbPath)
	if err != nil {
		panic(err)
	}
	return store
}

func TestCreateDB(t *testing.T) {
	dbPath := "./test.db"
	os.Remove(dbPath)

	_, err := NewBoltStore(dbPath)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveBuild(t *testing.T) {

	store := getStore()

	testBuild := &model.BuildHistoryItem{ID: "TEST", ProjectName: "pako"}

	err := store.SaveBuild(testBuild)
	if err != nil {
		t.Error(err)
	}
}

func TestGetLastBuilds(t *testing.T) {
	store := getStore()

	testBuilds := model.BuildHistory{
		&model.BuildHistoryItem{ID: "A1", ProjectName: "A", Start: 4, Valid: true},
		&model.BuildHistoryItem{ID: "A2", ProjectName: "A", Start: 5, Valid: true},
		&model.BuildHistoryItem{ID: "B1", ProjectName: "B", Start: 7, Valid: true},
		&model.BuildHistoryItem{ID: "B2", ProjectName: "B", Start: 9, Valid: true},
		&model.BuildHistoryItem{ID: "C1", ProjectName: "C", Start: 3, Valid: true},
		&model.BuildHistoryItem{ID: "C2", ProjectName: "C", Start: 2, Valid: true},
	}

	for _, b := range testBuilds {
		store.SaveBuild(b)
	}

	latests, err := store.GetLatestBuilds()
	if err != nil {
		t.Fatal(err)
	}
	if len(latests) != 3 {
		t.Fatal("Rows count should be 3. It is", len(latests))
	}

	if latests[0].ID != "B2" {
		t.Fatal("latests[0].ID should be 'B2', is", latests[0].ID)
	}

	if latests[1].ID != "A2" {
		t.Fatal("latests[1].ID should be 'A2', is", latests[1].ID)
	}

	if latests[2].ID != "C1" {
		t.Fatal("latests[2].ID should be 'C2', is", latests[2].ID)
	}

}

func TestGetAllBuilds(t *testing.T) {

	store := getStore()

	testBuilds := model.BuildHistory{
		&model.BuildHistoryItem{ID: "A1", ProjectName: "A", Start: 5},
		&model.BuildHistoryItem{ID: "A2", ProjectName: "A", Start: 4},
		&model.BuildHistoryItem{ID: "B1", ProjectName: "B", Start: 7},
		&model.BuildHistoryItem{ID: "B2", ProjectName: "B", Start: 9},
		&model.BuildHistoryItem{ID: "C1", ProjectName: "C", Start: 3},
		&model.BuildHistoryItem{ID: "C2", ProjectName: "C", Start: 2},
	}

	for _, b := range testBuilds {
		store.SaveBuild(b)
	}

	all, err := store.GetBuilds("B")
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 2 {
		t.Fatal("Rows count should be 2. It is", len(all))
	}

	if all[0].ID != "B2" {
		t.Fail()
	}

	if all[1].ID != "B1" {
		t.Fail()
	}
}

func TestBuildInvalidation(t *testing.T) {

	store := getStore()

	testBuild1 := &model.BuildHistoryItem{ID: "TEST", ProjectName: "pako", Start: 2, Valid: true}
	testBuild2 := &model.BuildHistoryItem{ID: "TEST2", ProjectName: "pako", Start: 4, Valid: true}
	store.SaveBuild(testBuild1)
	store.SaveBuild(testBuild2)

	err := store.InvalidateBuild("pako", "TEST")
	if err != nil {
		t.Fatal(err)
	}

	all, err := store.GetBuilds("pako")

	// Should not have been valid
	if all[0].Valid {
		t.Fatal("The build should not have been valid")
	}
}
