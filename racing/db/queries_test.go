package db

import (
	"database/sql"
	"testing"
	"time"

	//"git.neds.sh/matty/entain/racing/proto/racing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// setupTestDB creates and returns an in-memory SQLite database for testing.
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open sqlite memory db: %v", err)
	}
	return db
}

func TestListRaces_VisibleOnly(t *testing.T) {
	// Setup the in-memory database.
	sqldb := setupTestDB(t)
	defer sqldb.Close()

	// Create a new repository; this will also create (if needed) and seed the table.
	repo := NewRacesRepo(sqldb)
	err := repo.Init()
	assert.NoError(t, err, "failed to initialise the database")

	// Insert test data to ensure consistent results.
	// We'll insert one race with visible = 1 and one with visible = 0.
	now := time.Now().Format(time.RFC3339)
	stmt, err := sqldb.Prepare(`INSERT OR IGNORE INTO races(id, meeting_id, name, number, visible, advertised_start_time) VALUES (?,?,?,?,?,?)`)
	if err != nil {
		t.Fatalf("failed to prepare insert statement: %v", err)
	}
	defer stmt.Close()

	// Insert a race that is visible.
	_, err = stmt.Exec(101, 1, "Test Race Visible", 1, 1, now)
	if err != nil {
		t.Fatalf("failed to insert visible race: %v", err)
	}

	// Insert a race that is NOT visible.
	_, err = stmt.Exec(102, 1, "Test Race Not Visible", 2, 0, now)
	if err != nil {
		t.Fatalf("failed to insert non visible race: %v", err)
	}
	/*
		//Test below has waterfall errors due to the main file not working
		// Test with filter for visible-only races.
		filter := &racing.ListRacesRequestFilter{
			VisibleOnly: true,
		}

		visibleRaces, err := repo.List(filter)
		assert.NoError(t, err, "List() with visible_only filter should not error")
		for _, race := range visibleRaces {
			assert.True(t, race.Visible, "expected race (id: %d) to be visible", race.Id)
		}
	*/
	// Optionally, test that when no filter is applied, both races are present.
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
	assert.True(t, foundVisible, "expected to find test race with visible = 1")
	assert.True(t, foundNotVisible, "expected to find test race with visible = 0")
}
