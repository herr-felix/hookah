package sqlitestore

import (
	"database/sql"
	"fmt"

	"../model"
	_ "github.com/mattn/go-sqlite3"
)

// SqliteStore ...
type SqliteStore struct {
	dbPath string
}

// NewSqliteStore ...
func NewSqliteStore(dbPath string) (*SqliteStore, error) {
	store := &SqliteStore{dbPath: dbPath}

	db, err := store.open()
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping database: %s", err)
	}

	return store, nil
}

func (s *SqliteStore) open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", s.dbPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %s", err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS builds (
		id          TEXT    NOT NULL PRIMARY KEY, 
		name        TEXT    NOT NULL,
		projectName TEXT    NOT NULL,
		start       INTEGER NOT NULL,
		duration    INTEGER NOT NULL,
		status      TEXT    NOT NULL,
		output      TEXT    NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("%q: %s", err, sqlStmt)
	}

	return db, nil
}

// func (s *SqliteStore) GetSummaries() ([]model.ProjectSummary, error) {

// }

// func (s *SqliteStore) GetBuild(ID string) (*model.BuildHistory, error) {

// }

// SaveBuild saves a build_history
func (s *SqliteStore) SaveBuild(data *model.BuildHistory) error {
	db, err := s.open()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
		INSERT INTO builds (
			id,
			name,
			projectName,
			start,
			duration,
			status,
			output ) 
			VALUES (?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		data.ID,
		data.Name,
		data.ProjectName,
		data.Start,
		data.Duration,
		data.Status,
		data.Output,
	)
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}
