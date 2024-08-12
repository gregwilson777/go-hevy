package hevy_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/swrm-io/go-hevy"
)

func TestWorkoutUnmarshal(t *testing.T) {
	workout := hevy.Workout{}

	data, err := os.ReadFile("testdata/base/workout.json")
	assert.NoError(t, err, "error reading testdata/base/workout.json")
	err = json.Unmarshal(data, &workout)
	assert.NoError(t, err, "error unmarshalling testdata/base/workout.json")

	assert.NotEmpty(t, workout)
	assert.IsType(t, uuid.UUID{}, workout.ID)

	assert.NotEmpty(t, workout.StartTime)
	assert.NotEmpty(t, workout.EndTime)
	assert.NotEmpty(t, workout.Exercises)
}

func TestExerciseUnmarshal(t *testing.T) {
	exercise := hevy.Exercise{}

	data, err := os.ReadFile("testdata/base/exercise.json")
	assert.NoError(t, err, "error reading testdata/base/exercise.json")
	err = json.Unmarshal(data, &exercise)
	assert.NoError(t, err, "error unmarshalling testdata/base/exercise.json")

	assert.NotEmpty(t, exercise)
	assert.NotEmpty(t, exercise.Sets)
}

func TestSetUnmarshal(t *testing.T) {
	set := hevy.Set{}

	data, err := os.ReadFile("testdata/base/set.json")
	assert.NoError(t, err, "error reading testdata/base/set.json")
	err = json.Unmarshal(data, &set)
	assert.NoError(t, err, "error unmarshalling testdata/base/exercise.json")

	assert.NotEmpty(t, set)
	assert.Equal(t, hevy.WarmupSet, set.SetType)
}

func TestRoutineUnmarshal(t *testing.T) {
	routine := hevy.Routine{}

	data, err := os.ReadFile("testdata/base/routine.json")
	assert.NoError(t, err, "error reading testdata/base/routine.json")
	err = json.Unmarshal(data, &routine)
	assert.NoError(t, err, "error unmarshalling testdata/base/routine.json")

	assert.NotEmpty(t, routine)
	assert.IsType(t, uuid.UUID{}, routine.ID)
	assert.NotEmpty(t, routine.CreatedAt)
	assert.NotEmpty(t, routine.UpdatedAt)
	assert.NotEmpty(t, routine.Exercises)

}
