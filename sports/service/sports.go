package service

import (
	"context"
	"time"

	"github.com/SylvanSol/Entain_Test/sports/proto/sports"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// sportsService implements the SportsServer interface.
type sportsService struct {
	sports.UnimplementedSportsServer
}

// NewSportsService returns a new instance of sportsService.
func NewSportsService() sports.SportsServer {
	return &sportsService{}
}

// ListEvents returns a list of sports matches. (I don't follow sport so I got ChatGPT to tell me some)
func (s *sportsService) ListEvents(ctx context.Context, req *sports.ListEventsRequest) (*sports.ListEventsResponse, error) {
	now := time.Now()
	return &sports.ListEventsResponse{
		Events: []*sports.Event{
			{
				Id:                  1,
				Name:                "Red Hawks vs Blue Titans",
				Location:            "Thunder Dome",
				AdvertisedStartTime: timestamppb.New(now.Add(2 * time.Hour)),
			},
			{
				Id:                  2,
				Name:                "Iron Bears vs Golden Foxes",
				Location:            "Victory Stadium",
				AdvertisedStartTime: timestamppb.New(now.Add(4 * time.Hour)),
			},
			{
				Id:                  3,
				Name:                "Night Wolves vs Storm Kings",
				Location:            "Arena Eclipse",
				AdvertisedStartTime: timestamppb.New(now.Add(6 * time.Hour)),
			},
		},
	}, nil
}
