package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
	frontend(resp, req)

	if resp.Header().Get("content-type") != "text/html" {
		t.Errorf("content-type expected %s got %s", "text/html; charset=utf-8", resp.Header().Get("content-type"))
	}

	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("status code expected %d got %d", http.StatusOK, resp.Result().StatusCode)
	}
}

func TestMain3(t *testing.T) {
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
	frontend(resp, req)

	if resp.Header().Get("content-type") != "text/html" {
		t.Errorf("content-type expected %s got %s", "text/html; charset=utf-8", resp.Header().Get("content-type"))
	}

	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("status code expected %d got %d", http.StatusOK, resp.Result().StatusCode)
	}
}

func TestMain2(t *testing.T) {
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
	frontend(resp, req)

	if resp.Header().Get("content-type") != "text/html" {
		t.Errorf("content-type expected %s got %s", "text/html; charset=utf-8", resp.Header().Get("content-type"))
	}

	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("status code expected %d got %d", http.StatusOK, resp.Result().StatusCode)
	}
}

func TestMain4(t *testing.T) {
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
	frontend(resp, req)

	if resp.Header().Get("content-type") != "text/html" {
		t.Errorf("content-type expected %s got %s", "text/html; charset=utf-8", resp.Header().Get("content-type"))
	}

	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("status code expected %d got %d", http.StatusOK, resp.Result().StatusCode)
	}
}

func TestMain5(t *testing.T) {
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
	frontend(resp, req)

	if resp.Header().Get("content-type") != "text/html" {
		t.Errorf("content-type expected %s got %s", "text/html; charset=utf-8", resp.Header().Get("content-type"))
	}

	if resp.Result().StatusCode != http.StatusOK {
		t.Errorf("status code expected %d got %d", http.StatusOK, resp.Result().StatusCode)
	}
}
