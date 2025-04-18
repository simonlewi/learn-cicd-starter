package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-key-123"},
			},
			expectedKey:   "test-key-123",
			expectedError: nil,
		},
		{
			name:          "missing auth header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header without ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer test-key-123"},
			},
			expectedKey:   "",
			expectedError: nil,
		},
		{
			name: "malformed header without token",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			if tc.expectedError != nil && err != tc.expectedError {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}

			if key != tc.expectedKey {
				t.Errorf("expected key %s, got %s", tc.expectedKey, key)
			}
		})
	}
}
