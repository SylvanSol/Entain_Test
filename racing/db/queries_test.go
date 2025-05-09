// File: racing/db/queries_test.go
package db

import (
	"database/sql"
	//"fmt"
	"testing"
	"time"

	"git.neds.sh/matty/entain/racing/proto/racing"
	//tspb "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// setupTestDB creates an in-memory SQLite database for testing.
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open sqlite memory db: %v", err)
	}
	return db
}

// seedTestData creates the races table and inserts controlled data for ordering tests.
func seedTestData(t *testing.T, db *sql.DB) {
	// Create the races table.
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS races (
			id INTEGER PRIMARY KEY,
			meeting_id INTEGER,
			name TEXT,
			number INTEGER,
			visible INTEGER,
			advertised_start_time DATETIME
		)
	`)
	if err != nil {
		t.Fatalf("failed to create races table: %v", err)
	}

	// Prepare an insert statement.
	stmt, err := db.Prepare(`
		INSERT OR IGNORE INTO races(id, meeting_id, name, number, visible, advertised_start_time)
		VALUES (?,?,?,?,?,?)
	`)
	if err != nil {
		t.Fatalf("failed to prepare insert statement: %v", err)
	}
	defer stmt.Close()

	// Insert test data.
	// Times are set intentionally out of order.
	t1 := "2025-01-01T10:00:00Z"
	t2 := "2025-01-01T09:00:00Z"
	t3 := "2025-01-01T11:00:00Z"

	// Insert visible races.
	_, err = stmt.Exec(201, 1, "Alpha", 1, 1, t1)
	assert.NoError(t, err, "failed to insert race Alpha")
	_, err = stmt.Exec(202, 1, "Charlie", 2, 1, t2)
	assert.NoError(t, err, "failed to insert race Charlie")
	_, err = stmt.Exec(203, 1, "Bravo", 3, 1, t3)
	assert.NoError(t, err, "failed to insert race Bravo")
	// Insert a non-visible race (should be excluded when filtering by visible only).
	_, err = stmt.Exec(204, 1, "Delta", 4, 0, t1)
	assert.NoError(t, err, "failed to insert race Delta (non-visible)")
}

// TASK 1
func TestListRaces_VisibleOnly(t *testing.T) {
	// Setup the in-memory database.
	sqldb := setupTestDB(t)
	defer sqldb.Close()

	// Create the races table.
	_, err := sqldb.Exec(`
		CREATE TABLE IF NOT EXISTS races (
			id INTEGER PRIMARY KEY,
			meeting_id INTEGER,
			name TEXT,
			number INTEGER,
			visible INTEGER,
			advertised_start_time DATETIME
		)
	`)
	assert.NoError(t, err, "failed to create races table")

	// Initialise the repository.
	repo := NewRacesRepo(sqldb)
	err = repo.Init()
	assert.NoError(t, err, "failed to initialise the database")

	// Insert test data: one visible race and one non-visible race.
	now := time.Now().Format(time.RFC3339)
	stmt, err := sqldb.Prepare(`
		INSERT OR IGNORE INTO races(id, meeting_id, name, number, visible, advertised_start_time)
		VALUES (?,?,?,?,?,?)
	`)
	assert.NoError(t, err, "failed to prepare insert statement")
	defer stmt.Close()

	// Visible race.
	_, err = stmt.Exec(101, 1, "Test Race Visible", 1, 1, now)
	assert.NoError(t, err, "failed to insert visible race")
	// Non-visible race.
	_, err = stmt.Exec(102, 1, "Test Race Not Visible", 2, 0, now)
	assert.NoError(t, err, "failed to insert non-visible race")

}

// TASK 2
func TestListRaces_OrderBy(t *testing.T) {
	sqldb := setupTestDB(t)
	defer sqldb.Close()
	seedTestData(t, sqldb)
	repo := NewRacesRepo(sqldb)

	filter := &racing.ListRacesRequestFilter{}
	races, err := repo.List(filter, "advertised_start_time", "asc")
	assert.NoError(t, err)
	expected := []int64{202, 201, 203}
	var actual []int64
	for _, r := range races {
		actual = append(actual, r.Id)
	}
	assert.Equal(t, expected, actual)

	races, err = repo.List(filter, "name", "asc")
	assert.NoError(t, err)
	expected = []int64{201, 203, 202}
	actual = nil
	for _, r := range races {
		actual = append(actual, r.Id)
	}
	assert.Equal(t, expected, actual)
}
