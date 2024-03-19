package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateSessionIdShouldEqual(t *testing.T) {
	testCase := []struct {
		a, b int64
		c    string
		info string
	}{
		{100, 200, "100200",
			"GenerateSessionId with two positive numbers"},
		{-1, -1, "-1-1",
			"GenerateSessionId with two negative numbers"},
		{1<<63 - 1, 1<<63 - 1, "92233720368547758079223372036854775807",
			"GenerateSessionId with two MaxInt64"},
		{1<<63 - 1, -1 << 63, "-92233720368547758089223372036854775807",
			"GenerateSessionId with MaxInt64 and MinInt64"},
	}
	for _, tt := range testCase {
		require.Equal(t, GenerateSessionId(tt.a, tt.b), tt.c, tt.info)
	}
}

func TestGenerateSessionIdShouldNotEqual(t *testing.T) {
	testCase := []struct {
		a, b int64
		c    string
		info string
	}{
		{100, 200, "111222",
			"GenerateSessionId with two positive numbers"},
		{-1, -1, "-11",
			"GenerateSessionId with two negative numbers, case 1"},
		{-1, -1, "1-1",
			"GenerateSessionId with two negative numbers, case 2"},
		{1<<63 - 1, -1 << 63, "9223372036854775807-9223372036854775808",
			"GenerateSessionId with MaxInt64 and MinInt64"},
	}

	for _, tt := range testCase {
		require.NotEqual(t, GenerateSessionId(tt.a, tt.b), tt.c, tt.info)
	}
}
