package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	HealthHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.StatusCode)
	}
	if got := res.Header.Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected content-type %q, got %q", "application/json", got)
	}

	var payload map[string]string
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload["status"] != "ok" {
		t.Fatalf("expected status %q, got %q", "ok", payload["status"])
	}
}

func TestHealthHandlerRejectsNonGET(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rec := httptest.NewRecorder()

	HealthHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected %d, got %d", http.StatusMethodNotAllowed, res.StatusCode)
	}
	if got := res.Header.Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected content-type %q, got %q", "application/json", got)
	}
	if got := res.Header.Get("Allow"); got != http.MethodGet {
		t.Fatalf("expected allow %q, got %q", http.MethodGet, got)
	}

	var payload errorResponse
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload.Error != "method_not_allowed" {
		t.Fatalf("expected error %q, got %q", "method_not_allowed", payload.Error)
	}
	if payload.Message != "health endpoint accepts only GET" {
		t.Fatalf("expected message %q, got %q", "health endpoint accepts only GET", payload.Message)
	}
}

func TestWriteJSONError(t *testing.T) {
	rec := httptest.NewRecorder()

	err := writeJSONError(rec, http.StatusBadRequest, "bad_request", "invalid payload")
	if err != nil {
		t.Fatalf("writeJSONError returned error: %v", err)
	}

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
	if got := res.Header.Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected content-type %q, got %q", "application/json", got)
	}

	var payload errorResponse
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload.Error != "bad_request" {
		t.Fatalf("expected error %q, got %q", "bad_request", payload.Error)
	}
	if payload.Message != "invalid payload" {
		t.Fatalf("expected message %q, got %q", "invalid payload", payload.Message)
	}
}
