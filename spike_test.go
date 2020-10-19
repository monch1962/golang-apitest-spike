package main

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
)

var host = "http://localhost:8000/"
func TestHandler(t *testing.T) {
	var getUserMock = apitest.NewMock().
    	Get("/user/1234").
    	RespondWith().
    	Body(`{"name": "jon"}`).
    	Status(http.StatusOK).
		End()

	apitest.New().                              // configuration
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"id": "1234", "name": "Andy"}`))
			w.WriteHeader(http.StatusOK)
		}).
		Mocks(getUserMock).
		Report(apitest.SequenceDiagram()).
		Get("/user/1234").                      // request
		Expect(t).
		Body(`{"id": "1234", "name": "Andy"}`). // expectations
		Status(http.StatusOK).
		Assert(jsonpath.Contains(`$.b[? @.key=="c"].value`, "result")).
		End()
}

func TestMocks_Standalone(t *testing.T) {
	cli := http.Client{Timeout: 5}
	if os.Getenv("MOCKS") != "" {
		defer apitest.NewMock().
			Post(host + "/path").
			Body(`{"a", 12345}`).
			RespondWith().
			Status(http.StatusCreated).
			EndStandalone()()
	}

	resp, err := cli.Post(host + "/path",
		"application/json",
		strings.NewReader(`{"a", 12345}`))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
