package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	t.Run("OK HealthCheck", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		res := httptest.NewRecorder()
		HealthCheckHandler(res, req)
		wantCode := http.StatusOK
		wantBody := "OK\n"
		gotCode := res.Code
		gotBody := res.Body.String()
		if wantCode != gotCode {
			t.Errorf("want: %d got: %d", wantCode, gotCode)
		}
		if wantBody != gotBody {
			t.Errorf("want: %s got: %s", wantBody, gotBody)
		}
	})
}
