package db

import (
	"database/sql"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"

	"git.neds.sh/matty/entain/racing/proto/racing"
)

// RacesRepo provides repository access to races.
type RacesRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *racing.ListRacesRequestFilter, sortField, sortDirection string) ([]*racing.Race, error)

	GetByID(id int64) (*racing.Race, error)
}

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewRacesRepo(db *sql.DB) RacesRepo {
	return &racesRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *racesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

func (r *racesRepo) List(filter *racing.ListRacesRequestFilter, sortField, sortDirection string) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getRaceQueries()[racesList]

	query, args = r.applyFilter(query, filter, sortField, sortDirection)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

func (r *racesRepo) applyFilter(query string, filter *racing.ListRacesRequestFilter, sortField, sortDirection string) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	placeholders := strings.Repeat("?,", len(filter.MeetingIds)-1) + "?"
	clauses = append(clauses, "meedting_id IN ("+placeholders+")")
	for _, meetingID := range filter.MeetingIds {
		args = append(args, meetingID)
	}

	if filter.OnlyVisible {
		clauses = append(clauses, "visible = 1")
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	// Allowed fields for sorting
	allowedOrderFields := map[string]bool{
		"advertised_start_time": true,
		"name":                  true,
		"number":                true,
	}

	// Validate and default sortField
	if !allowedOrderFields[sortField] {
		sortField = "advertised_start_time"
	}

	// Validate and default sortDirection
	sortDirection = strings.ToUpper(sortDirection)
	if sortDirection != "DESC" {
		sortDirection = "ASC"
	}

	query += " ORDER BY " + sortField + " " + sortDirection

	return query, args
}

func (m *racesRepo) scanRaces(
	rows *sql.Rows,
) ([]*racing.Race, error) {
	var races []*racing.Race

	for rows.Next() {
		var race racing.Race
		var advertisedStart time.Time

		if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		race.AdvertisedStartTime = ts
		// Derive OPEN/CLOSED status
		if advertisedStart.Before(time.Now()) {
			race.Status = racing.RaceStatus_CLOSED
		} else {
			race.Status = racing.RaceStatus_OPEN
		}
		races = append(races, &race)
	}

	return races, nil
}

// GetByID fetches a single Race by its ID.
func (r *racesRepo) GetByID(id int64) (*racing.Race, error) {
	row := r.db.QueryRow(`SELECT id, meeting_id, name, number, visible, advertised_start_time FROM races WHERE id = ?`, id)
	var (
		race            racing.Race
		advertisedStart time.Time
	)
	if err := row.Scan(
		&race.Id,
		&race.MeetingId,
		&race.Name,
		&race.Number,
		&race.Visible,
		&advertisedStart,
	); err != nil {
		return nil, err
	}

	// Timestamp conversion
	ts, err := ptypes.TimestampProto(advertisedStart)
	if err != nil {
		return nil, err
	}
	race.AdvertisedStartTime = ts

	// Re-use Task 3 logic:
	if advertisedStart.Before(time.Now()) {
		race.Status = racing.RaceStatus_CLOSED
	} else {
		race.Status = racing.RaceStatus_OPEN
	}

	return &race, nil
}
