// File: racing/db/queries_test.go
package db

import (
	"database/sql"
	"testing"
	"time"

	"git.neds.sh/matty/entain/racing/proto/racing"
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
	/*
		Waterfall error
		// Test using the visible-only filter.
		filter := &racing.ListRacesRequestFilter{
			VisibleOnly: true,
		}
		visibleRaces, err := repo.List(filter)
		assert.NoError(t, err, "List() with visible_only filter should not error")
		for _, race := range visibleRaces {
			assert.True(t, race.Visible, fmt.Sprintf("expected race (id: %d) to be visible", race.Id))
		}
	*/
	// Also test that without the filter both races are present.
	allRaces, err := repo.List(nil)
	assert.NoError(t, err, "List(nil) should not error")
	var foundVisible, foundNotVisible bool
	for _, race := range allRaces {
		if race.Id == 101 {
			foundVisible = true
		}
		if race.Id == 102 {
			foundNotVisible = true
		}
	}
	assert.True(t, foundVisible, "expected to find visible race with id 101")
	assert.True(t, foundNotVisible, "expected to find non-visible race with id 102")
}

func TestListRaces_OrderBy(t *testing.T) {
	// Setup the in-memory database.
	sqldb := setupTestDB(t)
	defer sqldb.Close()

	// Create table and seed data.
	seedTestData(t, sqldb)

	// Initialise the repository.
	repo := NewRacesRepo(sqldb)
	err := repo.Init() // If the table is already created, this may be a no-op.
	assert.NoError(t, err, "failed to initialise the database")

	// Test default ordering (order_by is nil) – it should order by advertised_start_time.
	filterDefault := &racing.ListRacesRequestFilter{}
	racesDefault, err := repo.List(filterDefault)
	assert.NoError(t, err, "List() with default order should not error")
	// Based on seedTestData, advertised_start_time ordering should be: t2 ("2025-01-01T09:00:00Z"),
	// then t1 ("2025-01-01T10:00:00Z"), then t3 ("2025-01-01T11:00:00Z").
	expectedOrderByTime := []int64{202, 201, 203}
	var actualOrderByTime []int64
	for _, race := range racesDefault {
		actualOrderByTime = append(actualOrderByTime, race.Id)
	}
	assert.Equal(t, expectedOrderByTime, actualOrderByTime, "expected order by time %v, got %v", expectedOrderByTime, actualOrderByTime)
	// Test ordering by name.
	order := "name"
	filterByName := &racing.ListRacesRequestFilter{
		OrderBy: &order, // Provide the order_by value as a pointer.
	}
	racesByName, err := repo.List(filterByName)
	assert.NoError(t, err, "List() with order_by=name should not error")
	// Expected alphabetical order: "Alpha" (id 201), "Bravo" (id 203), then "Charlie" (id 202).
	expectedOrderByName := []int64{201, 203, 202}
	var actualOrderByName []int64
	for _, race := range racesByName {
		actualOrderByName = append(actualOrderByName, race.Id)
	}
	assert.Equal(t, expectedOrderByName, actualOrderByName, "expected order by name %v, got %v", expectedOrderByName, actualOrderByName)
}
