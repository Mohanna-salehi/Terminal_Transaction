package api

import (
	"Terminal_Transaction/internal/database"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_GetTerminals(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	h := NewHandler(db)
	req := httptest.NewRequest("GET", "/api/reports/terminal", nil)
	resp := httptest.NewRecorder()
	start := time.Now()
	h.GetTerminals(resp, req)
	elapsed := time.Since(start)
	t.Log("elapsed:", elapsed)

	if resp.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			resp.Code,
		)
	}
}

func TestGetDailyTransaction(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	h := NewHandler(db)
	req := httptest.NewRequest("GET", "/api/reports/terminals/daily", nil)
	resp := httptest.NewRecorder()
	start := time.Now()
	h.GetDailyTransaction(resp, req)
	elapsed := time.Since(start)
	t.Log("elapsed:", elapsed)

	if resp.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			resp.Code,
		)
	}
}

func TestGetLastTransactionResponse(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	h := NewHandler(db)
	req := httptest.NewRequest("GET", "/api/reports/terminals/last-transactions", nil)
	resp := httptest.NewRecorder()
	start := time.Now()
	h.GetLastTransaction(resp, req)
	elapsed := time.Since(start)
	t.Log("elapsed:", elapsed)

	if resp.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			resp.Code,
		)
	}
}
