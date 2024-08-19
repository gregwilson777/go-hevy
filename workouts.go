package hevy

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// A response for fetching a list of workouts
type workoutResponse struct {
	paginatedResults
	Workouts []Workout `json:"workouts"`
}

type workoutCountResponse struct {
	Count int `json:"workout_count"`
}

type workoutEventResponse struct {
	paginatedResults
	Events []Event
}

// Workouts gets all workouts.
func (c Client) Workouts() ([]Workout, error) {
	workouts := []Workout{}

	page := 1
	pageCount := 10

	for {
		q := map[string]string{
			"page":      fmt.Sprintf("%d", page),
			"pageCount": fmt.Sprintf("%d", pageCount),
		}
		url := c.constructURL("workouts", q)
		result := workoutResponse{}
		err := c.get(url, &result)
		if err != nil {
			return nil, err
		}

		workouts = append(workouts, result.Workouts...)

		if result.Page == result.PageCount {
			break
		}
		page++
	}

	return workouts, nil
}

// WorkoutCount returns a count of workouts
func (c Client) WorkoutCount() (int, error) {
	url := c.constructURL("workouts/count", map[string]string{})

	result := workoutCountResponse{}

	err := c.get(url, &result)
	if err != nil {
		return 0, err
	}

	return result.Count, nil
}

// WorkoutEvents retrieves a paged list of workout events (updates or deletes) since a given date.
// Events are ordered from newest to oldest. The intention is to allow clients to keep their local
// cache of workouts up to date without having to fetch the entire list of workouts.
func (c Client) WorkoutEvents(since time.Time) ([]Event, error) {
	events := []Event{}

	page := 1
	pageCount := 10

	for {
		q := map[string]string{
			"page":      fmt.Sprintf("%d", page),
			"pageCount": fmt.Sprintf("%d", pageCount),
			"since":     since.Format("RFC3339Nano"),
		}
		url := c.constructURL("workouts/events", q)
		result := workoutEventResponse{}
		err := c.get(url, &result)
		if err != nil {
			return nil, err
		}

		events = append(events, result.Events...)

		if result.Page == result.PageCount {
			break
		}
		page++
	}

	return events, nil
}

func (c Client) Workout(id uuid.UUID) (Workout, error) {
	path := fmt.Sprintf("workouts/%s", id.String())
	url := c.constructURL(path, map[string]string{})

	result := Workout{}

	err := c.get(url, &result)
	if err != nil {
		return Workout{}, err
	}

	return result, nil
}
