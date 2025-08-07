package hevy

import "fmt"

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
		q := map[string]string{
			"page":      fmt.Sprintf("%d", page),
			"pageCount": fmt.Sprintf("%d", pageCount),
		}
		url := c.constructURL("routines", q)
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

	// Fetch routine folders to get the folder information for each routine
	folders, err := c.RoutineFolders()
	if err != nil {
		return nil, err
	}

	for _, r := range routines {
		// Find the folder for each routine
		for _, f := range folders {
			if r.FolderID == f.ID {
				r.FolderName = f.Title
				break
			}
		}
	}

	return routines, nil
}

type routineFolderResponse struct {
	paginatedResults
	RoutineFolders []RoutineFolder `json:"routine_folders"`
}

// Routines gets all routine folders.
func (c Client) RoutineFolders() ([]RoutineFolder, error) {
	routines := []RoutineFolder{}

	page := 1
	pageCount := 10
	for {
		q := map[string]string{
			"page":      fmt.Sprintf("%d", page),
			"pageCount": fmt.Sprintf("%d", pageCount),
		}
		url := c.constructURL("routine_folders", q)
		result := routineFolderResponse{}
		err := c.get(url, &result)
		if err != nil {
			return nil, err
		}

		routines = append(routines, result.RoutineFolders...)

		if result.Page == result.PageCount {
			break
		}
		page++
	}

	return routines, nil
}
