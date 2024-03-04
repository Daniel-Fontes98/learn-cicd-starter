package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		headers    http.Header
		expected   string
		expectErr  bool
		errMessage string
	}{
		// Test case 1: Valid authorization header
		{
			headers:    http.Header{"Authorization": []string{"ApiKey my-api-key"}},
			expected:   "my-api-key",
			expectErr:  false,
			errMessage: "",
		},
		// Test case 2: Missing authorization header
		{
			headers:    http.Header{},
			expected:   "",
			expectErr:  true,
			errMessage: "no authorization header included",
		},
		// Test case 3: Malformed authorization header
		{
			headers:    http.Header{"Authorization": []string{"Bearer my-api-key"}},
			expected:   "",
			expectErr:  true,
			errMessage: "malformed authorization header",
		},
	}

	for _, tc := range testCases {
		apiKey, err := GetAPIKey(tc.headers)

		if tc.expectErr {
			if err == nil {
				t.Errorf("Expected error, but got nil")
			} else if err.Error() != tc.errMessage {
				t.Errorf("Expected error message '%s', but got '%s'", tc.errMessage, err.Error())
			}
		} else {
			if err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
			if apiKey != tc.expected {
				t.Errorf("Expected API key '%s', but got '%s'", tc.expected, apiKey)
			}
		}
	}
}
