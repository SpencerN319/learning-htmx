//go:build unit

package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	tc := []struct {
		description string
		req         *http.Request
		expStatus   int
	}{
		{
			description: "GET request should return 200",
			req:         httptest.NewRequest(http.MethodGet, "/health", nil),
			expStatus:   http.StatusOK,
		},
		{
			description: "GET request should return service unavailable",
			req:         httptest.NewRequest(http.MethodPost, "/health", nil),
			expStatus:   http.StatusMethodNotAllowed,
		},
	}

	svr := newServer(context.TODO())
	svr.status.Store(http.StatusOK)
	defer svr.Close()

	for i, c := range tc {
		w := httptest.NewRecorder()
		svr.ServeHTTP(w, c.req)
		res := w.Result()

		if res.StatusCode != c.expStatus {
			t.Errorf("[%d] %s: expected %d; received %d", i, c.description, c.expStatus, res.StatusCode)
		}
	}
}
