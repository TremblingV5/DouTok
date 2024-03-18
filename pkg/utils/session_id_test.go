package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateSessionIdShouldEqual(t *testing.T) {
	testCase := []struct {
		a, b int64
		c    string
	}{
		{100, 200, "100200"},
		{-1, -1, "-1-1"},
		{1<<63 - 1, 1<<63 - 1, "92233720368547758079223372036854775807"},
		{1<<63 - 1, -1 << 63, "-92233720368547758089223372036854775807"},
	}

	for _, tt := range testCase {
		require.Equal(t, GenerateSessionId(tt.a, tt.b), tt.c)
	}
}

func TestGenerateSessionIdShouldNotEqual(t *testing.T) {
	testCase := []struct {
		a, b int64
		c    string
	}{
		{100, 200, "111222"},
		{-1, -1, "-11"},
		{-1, -1, "1-1"},
		{1<<63 - 1, -1 << 63, "9223372036854775807-9223372036854775808"},
	}

	for _, tt := range testCase {
		require.NotEqual(t, GenerateSessionId(tt.a, tt.b), tt.c)
	}
}
