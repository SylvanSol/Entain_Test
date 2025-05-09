package service

import (
	"database/sql"
	"fmt"
	"strings"

	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// racingService implements the Racing interface.
type racingService struct {
	racing.UnimplementedRacingServer // Embedding due to later version of Go
	racesRepo                        db.RacesRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewRacingService(racesRepo db.RacesRepo) racing.RacingServer {
	return &racingService{racesRepo: racesRepo}
}

func (s *racingService) ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error) {
	// fallback values in case the Sort field is nil
	field := "advertised_start_time"
	direction := "ASC"
	if in.Sort != nil {
		if in.Sort.Field != "" {
			field = in.Sort.Field
		}
		if strings.ToUpper(in.Sort.Direction) == "DESC" {
			direction = "DESC"
		}
	}

	races, err := s.racesRepo.List(in.Filter, field, direction)
	if err != nil {
		return nil, fmt.Errorf("failed to list races: %v", err)
	}

	return &racing.ListRacesResponse{Races: races}, nil
}

func (s *racingService) GetRace(ctx context.Context, req *racing.GetRaceRequest) (*racing.GetRaceResponse, error) {
	race, err := s.racesRepo.GetByID(req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "race %d not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "error fetching race: %v", err)
	}
	return &racing.GetRaceResponse{Race: race}, nil
}
