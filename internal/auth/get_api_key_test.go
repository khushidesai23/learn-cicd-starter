// Add a couple unit tests for GetAPIKey.
package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("valid API key", func(t *testing.T) {
		expectedKey := "valid_api_key"
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey "+expectedKey)

		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if apiKey != expectedKey {
			t.Errorf("expected API key %q, got %q", expectedKey, apiKey)
		}
	})

	t.Run("missing authorization header", func(t *testing.T) {
		headers := http.Header{}

		apiKey, err := GetAPIKey(headers)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
		if err != ErrNoAuthHeaderIncluded {
			t.Fatalf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
		}
		if apiKey != "" {
			t.Errorf("expected empty API key, got %q", apiKey)
		}
	})
}
