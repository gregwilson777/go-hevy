package hevy_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/swrm-io/go-hevy"
)

func TestWorkouts(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		page := req.URL.Query().Get("page")

		file := fmt.Sprintf("testdata/responses/workout-%s.json", page)
		data, err := os.ReadFile(file)
		assert.NoError(t, err)
		res.Write(data)
	}))
	defer testServer.Close()

	client := hevy.NewClient("my-fake-api-key")
	client.ApiURL = testServer.URL

	workouts, err := client.Workouts()
	assert.NoError(t, err)
	assert.NotEmpty(t, workouts)
	assert.Len(t, workouts, 6)
}

func TestWorkoutCount(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		data, err := os.ReadFile("testdata/responses/workout-count.json")
		assert.NoError(t, err)
		res.Write(data)
	}))
	defer testServer.Close()

	client := hevy.NewClient("my-fake-api-key")
	client.ApiURL = testServer.URL

	count, err := client.WorkoutCount()
	assert.NoError(t, err)
	assert.Equal(t, 21, count)
}

func TestWorkout(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		data, err := os.ReadFile("testdata/responses/single-workout.json")
		assert.NoError(t, err)
		res.Write(data)
	}))
	defer testServer.Close()

	client := hevy.NewClient("my-fake-api-key")
	client.ApiURL = testServer.URL

	workoutID, err := uuid.Parse("b459cba5-cd6d-463c-abd6-54f8eafcadcb")
	assert.NoError(t, err)
	workout, err := client.Workout(workoutID)
	assert.NoError(t, err)
	assert.NotEmpty(t, workout)
	assert.Equal(t, workoutID, workout.ID)
	assert.Equal(t, "Morning Workout 💪", workout.Title)
}
