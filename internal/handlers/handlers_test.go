package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/1saswata/go-mentorship/internal/store"
)

type mockStore struct {
}

func (m mockStore) CreateTask(name, status string) int {
	return 1
}

func (m mockStore) GetAllTasks() []store.Task {
	return []store.Task{{ID: 1, Name: "Test name", Status: "Test-Status"}}
}

func (m mockStore) UpdateTaskStatus(id int, status string) error {
	return nil
}

func (m mockStore) DeleteTask(id int) error {
	return nil
}

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

func TestListTaskHandler(t *testing.T) {
	t.Run("OneTask", func(t *testing.T) {
		server := TaskServer{Store: mockStore{}}
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		res := httptest.NewRecorder()
		server.ListTaskHandler(res, req)
		wantCode := http.StatusOK
		wantBody := "Test name"
		gotCode := res.Code
		gotBody := res.Body.String()
		if wantCode != gotCode {
			t.Errorf("want: %d , got: %d", wantCode, gotCode)
		}
		if !strings.Contains(gotBody, wantBody) {
			t.Errorf("want %s, got %s", wantBody, gotBody)
		}
	})
}

func TestCreateTaskHandler(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		server := TaskServer{Store: mockStore{}}
		req := httptest.NewRequest(http.MethodPost, "/tasks",
			strings.NewReader(`{"name":"Test","status":"TODO"}`))
		res := httptest.NewRecorder()
		server.CreateTaskHandler(res, req)
		wantCode := http.StatusCreated
		wantBody := "Test"
		gotCode := res.Code
		gotBody := res.Body.String()
		if wantCode != gotCode {
			t.Errorf("want: %d , got: %d", wantCode, gotCode)
		}
		if !strings.Contains(gotBody, wantBody) {
			t.Errorf("want %s, got %s", wantBody, gotBody)
		}
	})

	t.Run("Invalid Json", func(t *testing.T) {
		server := TaskServer{Store: mockStore{}}
		req := httptest.NewRequest(http.MethodPost, "/tasks",
			strings.NewReader(`{"name""Broken Json","status":"TODO"}`))
		res := httptest.NewRecorder()
		server.CreateTaskHandler(res, req)
		wantCode := http.StatusBadRequest
		gotCode := res.Code
		if wantCode != gotCode {
			t.Errorf("want: %d , got: %d", wantCode, gotCode)
		}
	})
}

func TestUpdateTaskHandler(t *testing.T) {
	t.Run("Update Success", func(t *testing.T) {
		server := TaskServer{Store: mockStore{}}
		req := httptest.NewRequest(http.MethodPut, "/tasks",
			strings.NewReader(`{"status":"Complete"}`))
		req.SetPathValue("id", "1")
		res := httptest.NewRecorder()
		server.UpdateTaskHandler(res, req)
		got := res.Code
		want := http.StatusNoContent
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestDeleteTaskHandler(t *testing.T) {
	t.Run("Delete Success", func(t *testing.T) {
		server := TaskServer{Store: mockStore{}}
		req := httptest.NewRequest(http.MethodDelete, "/tasks", nil)
		req.SetPathValue("id", "1")
		res := httptest.NewRecorder()
		server.DeleteTaskHandler(res, req)
		want := http.StatusNoContent
		got := res.Code
		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}
