package hevy

import (
	"fmt"

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

// Workouts gets all workouts.
func (c Client) Workouts() ([]Workout, error) {
	workouts := []Workout{}

	page := 1
	pageCount := 10
	for {
		url := c.constructURL("workouts", page, pageCount)
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
	url := c.constructURL("workouts/count", 0, 0)

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
func (c Client) WorkoutsEvents() {
	// ToDO
}

func (c Client) Workout(id uuid.UUID) (Workout, error) {
	path := fmt.Sprintf("workouts/%s", id.String())
	url := c.constructURL(path, 0, 0)

	result := Workout{}

	err := c.get(url, &result)
	if err != nil {
		return Workout{}, err
	}

	return result, nil
}
