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
	CREATE INDEX IF NOT EXISTS builds_project_name ON builds(projectName);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("%q: %s", err, sqlStmt)
	}

	return store, nil
}

func (s *SqliteStore) open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", s.dbPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %s", err)
	}

	return db, nil
}

// GetAllBuilds get all the builds history of a project
func (s *SqliteStore) GetAllBuilds(projectName string) ([]*model.BuildHistory, error) {
	db, err := s.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		SELECT
			id,
			name,
			projectName,
			start,
			duration,
			status,
			output
		FROM builds 
		WHERE projectName = ?
		ORDER BY start DESC;`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(projectName)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var builds []*model.BuildHistory
	for rows.Next() {
		b := &model.BuildHistory{}
		err = rows.Scan(
			&b.ID,
			&b.Name,
			&b.ProjectName,
			&b.Start,
			&b.Duration,
			&b.Status,
			&b.Output,
		)
		if err != nil {
			return nil, err
		}
		builds = append(builds, b)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return builds, nil
}

// GetLatestBuilds get the latest build history for each projects
func (s *SqliteStore) GetLatestBuilds() ([]*model.BuildHistory, error) {
	db, err := s.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT
			id,
			name,
			projectName,
			start,
			duration,
			status
		FROM builds 
		WHERE (projectName, start) IN (
			SELECT projectName, MAX(start)
			FROM builds
			GROUP BY projectName
		) ORDER BY start DESC;`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var builds []*model.BuildHistory
	for rows.Next() {
		b := &model.BuildHistory{}
		err = rows.Scan(
			&b.ID,
			&b.Name,
			&b.ProjectName,
			&b.Start,
			&b.Duration,
			&b.Status,
		)
		if err != nil {
			return nil, err
		}
		builds = append(builds, b)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return builds, nil
}

// GetBuild retrieve a build by its ID
func (s *SqliteStore) GetBuild(ID string) (*model.BuildHistory, error) {
	db, err := s.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		SELECT
			id,
			name,
			projectName,
			start,
			duration,
			status,
			output
		FROM builds 
		WHERE id = ?;`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	buildHistory := &model.BuildHistory{}
	err = stmt.QueryRow(ID).Scan(
		&buildHistory.ID,
		&buildHistory.Name,
		&buildHistory.ProjectName,
		&buildHistory.Start,
		&buildHistory.Duration,
		&buildHistory.Status,
		&buildHistory.Output,
	)
	if err != nil {
		return nil, err
	}

	return buildHistory, nil
}

// SaveBuild saves a BuildHistory
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
