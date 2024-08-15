package hevy

// A response for fetching a list of workouts
type routineResponse struct {
	paginatedResults
	Routines []Routine `json:"routines"`
}

// Routines gets all routines.
func (c Client) Routines() ([]Routine, error) {
	routines := []Routine{}

	page := 1
	pageCount := 10
	for {
		url := c.constructURL("routines", page, pageCount)
		result := routineResponse{}
		err := c.get(url, &result)
		if err != nil {
			return nil, err
		}

		routines = append(routines, result.Routines...)

		if result.Page == result.PageCount {
			break
		}
		page++
	}

	return routines, nil
}
