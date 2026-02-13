package console

import "testing"

func TestResolveTitle(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty falls back to default", "", AppTitle},
		{"custom title preserved", "My Dashboard", "My Dashboard"},
	}

	for _, tt := range tests {
		if got := ResolveTitle(tt.input); got != tt.want {
			t.Errorf("ResolveTitle(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
