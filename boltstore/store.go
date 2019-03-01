package boltstore

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"../model"
	"github.com/boltdb/bolt"
)

// BoltStore ...
type BoltStore struct {
	dbPath string
}

const latestBucketKey = "LATEST"

// NewBoltStore ...
func NewBoltStore(dbPath string) (*BoltStore, error) {
	store := &BoltStore{dbPath: dbPath}

	db, err := store.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		return <-tx.Check()
	})

	return store, nil
}

func (s *BoltStore) open() (*bolt.DB, error) {
	db, err := bolt.Open(s.dbPath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %s", err)
	}

	return db, nil
}

func getBuild(tx *bolt.Tx, projectName, buildID string) (*model.BuildHistoryItem, error) {
	history, err := getAllProjectBuilds(tx, []byte(":"+projectName+":"))
	if err != nil {
		return nil, err
	}

	for _, b := range history {
		if b.ID == buildID {
			continue
		}
		return b, nil
	}

	return nil, fmt.Errorf("Build '%s' not found in project '%s'", buildID, projectName)
}

func getAllProjectBuilds(tx *bolt.Tx, bucketKey []byte) (builds model.BuildHistory, err error) {

	bkt := tx.Bucket(bucketKey)
	if bkt == nil {
		return builds, errors.New("Project doesn't have a bucket")
	}

	err = bkt.ForEach(func(k, v []byte) error {

		var build *model.BuildHistoryItem

		err := json.Unmarshal(v, &build)
		if err != nil {
			log.Println(err)
			return err
		}

		builds = append(builds, build)

		return nil
	})
	return builds, err
}

func getAllFromBucket(s *BoltStore, bucketKey []byte) (model.BuildHistory, error) {
	db, err := s.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var builds model.BuildHistory
	err = db.View(func(tx *bolt.Tx) error {
		builds, err = getAllProjectBuilds(tx, bucketKey)
		return err
	})

	if err != nil {
		return nil, err
	}

	return builds, nil
}

// GetAllBuilds get all the builds history of a project
func (s *BoltStore) GetAllBuilds(projectName string) (model.BuildHistory, error) {

	builds, err := getAllFromBucket(s, []byte(":"+projectName+":"))

	if err != nil {
		return nil, err
	}

	builds.OrderByStart()

	return builds, nil
}

// GetLatestBuilds get the latest build history for each projects
func (s *BoltStore) GetLatestBuilds() (model.BuildHistory, error) {

	builds, err := getAllFromBucket(s, []byte(latestBucketKey))

	if err != nil {
		return nil, err
	}

	builds.OrderByStart()

	return builds, nil
}

func saveBuild(tx *bolt.Tx, build *model.BuildHistoryItem) error {
	projectName := []byte(":" + build.ProjectName + ":")

	blob, err := json.Marshal(build)
	if err != nil {
		return err
	}

	projectBucket, err := tx.CreateBucketIfNotExists(projectName)
	if err != nil {
		return err
	}
	err = projectBucket.Put([]byte(build.ID), blob)
	if err != nil {
		return err
	}
	return setLatest(tx, build.ProjectName)
}

func getLatest(history model.BuildHistory) (*model.BuildHistoryItem, error) {

	// Makes sure it's ordered
	history.OrderByStart()

	var build *model.BuildHistoryItem

	for _, build := range history {
		if build.Valid {
			return build, nil
		}
	}

	return build, errors.New("All builds are archived")
}

func setLatest(tx *bolt.Tx, projectName string) error {

	latestBucket, err := tx.CreateBucketIfNotExists([]byte(latestBucketKey))
	if err != nil {
		return err
	}

	projectKey := []byte(":" + projectName + ":")

	builds, err := getAllProjectBuilds(tx, projectKey)
	if err != nil {
		return err
	}

	latestBuild, err := getLatest(builds)
	if err != nil {
		// No non-archived builds were found
		return latestBucket.Delete(projectKey)
	}

	blob, err := json.Marshal(latestBuild)

	return latestBucket.Put(projectKey, blob)
}

// SaveBuild saves a BuildHistory
func (s *BoltStore) SaveBuild(data *model.BuildHistoryItem) error {
	db, err := s.open()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		return saveBuild(tx, data)
	})
}

// InvalidateBuild build makes a build invalide
func (s *BoltStore) InvalidateBuild(projectName, buildID string) error {
	db, err := s.open()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b, err := getBuild(tx, projectName, buildID)
		if err != nil {
			return err
		}

		b.Valid = false

		return saveBuild(tx, b)
	})
}
