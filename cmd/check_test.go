package cmd

import "testing"

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{"Empty URL", "", false},
		{"Invalid URL", "http//googl", false},
		{"Valid http URL", "http://www.google.com", true},
		{"Valid https URL", "https://www.google.com", true},
		{"URL with no host", "http://", false},
		{"URL with IP", "http://192.168.0.1", true},
		{"URL with port", "http://localhost:8080", true},
		{"URL with special characters", "http://example.com/path?name=val#anchor", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {})
	}
}
