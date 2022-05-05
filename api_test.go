package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"api-webservice/controllers"
	"github.com/gorilla/mux"
)

func TestBlockedRequest(t *testing.T) {
	t.Run("returns url blocked policy", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/urlinfo/1/google.com:80/search?query", nil)

		response := httptest.NewRecorder()

		vars := map[string]string {
			"url": "google.com:80",
		}

		request = mux.SetURLVars(request, vars)

		handler := controllers.GetPolicy()
		handler.ServeHTTP(response, request)
		
		if response.Code != http.StatusOK {
			t.Fatalf("expected code %d, got %d", http.StatusOK, response.Code)
		}
		bodyStr := string(response.Body.Bytes())
		if len(bodyStr) <= 0 {
			t.Fatalf("expected non-empty response body")
		}

	})
}

func TestUnblockedRequest(t *testing.T) {
	t.Run("returns url unblcocked policy", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/urlinfo/1/cisco.com:80/search?query", nil)

		response := httptest.NewRecorder()

		vars := map[string]string {
			"url": "cisco.com:80",
		}

		request = mux.SetURLVars(request, vars)

		handler := controllers.GetPolicy()
		handler.ServeHTTP(response, request)
		
		if response.Code != http.StatusOK {
			t.Fatalf("expected code %d, got %d", http.StatusOK, response.Code)
		}
		bodyStr := string(response.Body.Bytes())
		if len(bodyStr) <= 0 {
			t.Fatalf("expected non-empty response body")
		}

	})
}

func TestFalseUrl(t *testing.T) {
	t.Run("return 204 No Content status", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/urlinfo/1/youtube.com:80/search?query", nil)

		response := httptest.NewRecorder()

		vars := map[string]string {
			"url": "youtube.com:80",
		}

		request = mux.SetURLVars(request, vars)

		handler := controllers.GetPolicy()
		handler.ServeHTTP(response, request)
		
		if response.Code != http.StatusNoContent {
			t.Fatalf("expected code %d, got %d", http.StatusNoContent, response.Code)
		}
	})
}
