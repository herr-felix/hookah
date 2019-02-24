package boltstore

import (
	"encoding/json"
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

func getAllFromBucket(s *BoltStore, key string) (model.BuildHistory, error) {
	db, err := s.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var builds model.BuildHistory
	err = db.View(func(tx *bolt.Tx) error {

		err := tx.Bucket([]byte(key)).ForEach(func(k, v []byte) error {

			var build *model.BuildHistoryItem

			err = json.Unmarshal(v, &build)
			if err != nil {
				log.Println(err)
				return err
			}

			builds = append(builds, build)

			return nil
		})
		return err
	})

	if err != nil {
		return nil, err
	}

	return builds, nil
}

// GetAllBuilds get all the builds history of a project
func (s *BoltStore) GetAllBuilds(projectName string) (model.BuildHistory, error) {

	builds, err := getAllFromBucket(s, ":"+projectName+":")

	if err != nil {
		return nil, err
	}

	builds.OrderByStart()

	return builds, nil
}

// GetLatestBuilds get the latest build history for each projects
func (s *BoltStore) GetLatestBuilds() (model.BuildHistory, error) {

	builds, err := getAllFromBucket(s, latestBucketKey)

	if err != nil {
		return nil, err
	}

	builds.OrderByStart()

	return builds, nil
}

// SaveBuild saves a BuildHistory
func (s *BoltStore) SaveBuild(data *model.BuildHistoryItem) error {
	db, err := s.open()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		ID := []byte(data.ID)
		projectName := []byte(":" + data.ProjectName + ":")
		blob, err := json.Marshal(data)
		if err != nil {
			return err
		}
		projectBucket, err := tx.CreateBucketIfNotExists(projectName)
		if err != nil {
			return err
		}
		err = projectBucket.Put(ID, blob)
		if err != nil {
			return err
		}

		latestBucket, err := tx.CreateBucketIfNotExists([]byte(latestBucketKey))
		if err != nil {
			return err
		}
		// TODO: Verify that `data` has a `.start` higher than what's currently in store
		err = latestBucket.Put(projectName, blob)
		if err != nil {
			return err
		}

		return nil
	})
}
