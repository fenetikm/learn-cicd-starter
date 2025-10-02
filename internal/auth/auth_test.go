package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input map[string]string
		err   bool
		want  string
	}

	tests := []test{
		{input: map[string]string{"bob": "sam"}, err: true, want: ""},
		{input: map[string]string{}, err: true, want: ""},
		{input: map[string]string{"Authorization": ""}, err: true, want: ""},
		{input: map[string]string{"Authorization": "ApiKey"}, err: true, want: ""},
		{input: map[string]string{"Authorization": "ApiKey 123"}, err: false, want: "123"},
	}

	for _, tc := range tests {
		headers := http.Header{}
		for key := range tc.input {
			headers.Set(key, tc.input[key])
		}

		got, err := GetAPIKey(headers)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
		if err == nil && tc.err == true {
			t.Fatalf("expected err")
		}
	}

}
