package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListStudents(t *testing.T) {
	rec := httptest.NewRecorder()
	ListStudents(rec, httptest.NewRequest("GET", "/api/students", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
