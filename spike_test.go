package main

import (
	"net/http"
	"testing"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
)
func TestHandler(t *testing.T) {
	apitest.New().                              // configuration
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"id": "1234", "name": "Andy"}`))
			w.WriteHeader(http.StatusOK)
		}).
		Report(apitest.SequenceDiagram()).
		Get("/user/1234").                      // request
		Expect(t).
		Body(`{"id": "1234", "name": "Andy"}`). // expectations
		Status(http.StatusOK).
		Assert(jsonpath.Contains(`$.b[? @.key=="c"].value`, "result")).
		End()
}