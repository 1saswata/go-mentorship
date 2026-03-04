package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/1saswata/go-mentorship/internal/store"
)

type mockStore struct {
	err error
}

func (m mockStore) CreateTask(name, status string) int {
	if m.err == nil {
		return 1
	} else {
		return -1
	}
}

func (m mockStore) GetAllTasks() []store.Task {
	return []store.Task{{ID: 1, Name: "Test name", Status: "Test-Status"}}
}

func (m mockStore) UpdateTaskStatus(id int, status string) error {
	return m.err
}

func (m mockStore) DeleteTask(id int) error {
	return m.err
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
	tests := []struct {
		name     string
		err      error
		reqBody  string
		wantCode int
		wantBody string
	}{
		{name: "Task created", err: nil, reqBody: `{"name":"Test","status":"TODO"}`,
			wantCode: http.StatusCreated, wantBody: "Test"},
		{name: "Invalid Json", err: nil, reqBody: `{"name""Broken Json","status":"TODO"}`,
			wantCode: http.StatusBadRequest, wantBody: ""},
		{name: "Task not created", err: errors.New("db down"),
			reqBody: `{"name":"Test","status":"TODO"}`, wantCode: http.StatusInternalServerError, wantBody: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := TaskServer{Store: mockStore{err: tt.err}}
			req := httptest.NewRequest(http.MethodPost, "/tasks",
				strings.NewReader(tt.reqBody))
			res := httptest.NewRecorder()
			server.CreateTaskHandler(res, req)
			gotCode := res.Code
			gotBody := res.Body.String()
			if tt.wantCode != gotCode {
				t.Errorf("want: %d , got: %d", tt.wantCode, gotCode)
			}
			if !strings.Contains(gotBody, tt.wantBody) {
				t.Errorf("want %s, got %s", tt.wantBody, gotBody)
			}
		})
	}
}

func TestUpdateTaskHandler(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		err     error
		reqBody string
		want    int
	}{
		{name: "Update Success", id: 1, err: nil, reqBody: `{"status":"Complete"}`,
			want: http.StatusNoContent},
		{name: "Update No ID", id: -1, err: nil, reqBody: `{"status":"Complete"}`,
			want: http.StatusBadRequest},
		{name: "Update Fail - Not found", id: 1, err: store.ErrNotFound, reqBody: `{"status":"Complete"}`,
			want: http.StatusNotFound},
		{name: "Update Fail - Bad body", id: 1, err: nil, reqBody: `{"status""Complete"}`,
			want: http.StatusBadRequest},
		{name: "Update Fail - DB error", id: 1, err: errors.New("db down"), reqBody: `{"status":"Complete"}`,
			want: http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := TaskServer{Store: mockStore{err: tt.err}}
			req := httptest.NewRequest(http.MethodPut, "/tasks",
				strings.NewReader(tt.reqBody))
			if tt.id != -1 {
				req.SetPathValue("id", strconv.Itoa(tt.id))
			}
			res := httptest.NewRecorder()
			server.UpdateTaskHandler(res, req)
			got := res.Code
			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestDeleteTaskHandler(t *testing.T) {
	tests := []struct {
		name string
		id   int
		err  error
		want int
	}{
		{name: "Delete Success", id: 1, err: nil,
			want: http.StatusNoContent},
		{name: "Delete No ID", id: -1, err: nil,
			want: http.StatusBadRequest},
		{name: "Update Fail - DB error", id: 1, err: errors.New("DB Error"),
			want: http.StatusInternalServerError},
		{name: "Update Fail - Not found", id: 1, err: store.ErrNotFound,
			want: http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := TaskServer{Store: mockStore{err: tt.err}}
			req := httptest.NewRequest(http.MethodDelete, "/tasks", nil)
			if tt.id != -1 {
				req.SetPathValue("id", strconv.Itoa(tt.id))
			}
			res := httptest.NewRecorder()
			server.DeleteTaskHandler(res, req)
			got := res.Code
			if tt.want != got {
				t.Errorf("want %d got %d", tt.want, got)
			}
		})
	}
}
