package db

import (
	"fmt"
	"strings"
)

const (
	racesList = "list"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races
		`,
	}
}

func buildListracesQuery(orderBy, sortDirection string) (string, error) {
	baseQuery := `
		SELECT 
			id, 
			meeting_id, 
			name, 
			number, 
			visible, 
			advertised_start_time 
		FROM races
	`
	allowedOrderFields := map[string]bool{
		"advertised_start_time": true,
		"name":                  true,
		"number":                true,
	}

	if orderBy == "" {
		orderBy = "advertised_start_time"
	}
	if !allowedOrderFields[orderBy] {
		return "", fmt.Errorf("invalid order_by field: %s", orderBy)
	}

	if sortDirection != "" && strings.ToLower(sortDirection) != "asc" && strings.ToLower(sortDirection) != "desc" {
		return "", fmt.Errorf("invalid sort_direction: %s", sortDirection)
	}

	if sortDirection == "" {
		sortDirection = "ASC"
	}
	query := fmt.Sprintf("%s ORDER BY %s %s", baseQuery, orderBy, sortDirection)
	return query, nil
}
