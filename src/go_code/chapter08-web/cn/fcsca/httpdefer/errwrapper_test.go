package httpdefer

import "testing"

func TestErrwrapper(t *testing.T) {
	tests := []struct {
		code    int
		message string
	}{
		{500, ""},
	}
}
