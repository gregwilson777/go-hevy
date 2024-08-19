package hevy_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swrm-io/go-hevy"
)

func TestRoutine(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		switch req.URL.Path {
		case "/v1/routines":
			page := req.URL.Query().Get("page")

			file := fmt.Sprintf("testdata/responses/routine-%s.json", page)
			data, err := os.ReadFile(file)
			assert.NoError(t, err)
			_, err = res.Write(data)
			assert.NoError(t, err)
		}
	}))
	defer srv.Close()

	client := hevy.NewClient("my-fake-api-key")
	client.APIURL = srv.URL

	t.Run("Test Paginated Routines", func(t *testing.T) {
		routines, err := client.Routines()
		assert.NoError(t, err)
		assert.NotEmpty(t, routines)
		assert.Len(t, routines, 3)
	})
}
