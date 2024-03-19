package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateSessionIdShouldEqual(t *testing.T) {
	cases := []struct {
		name string
		a, b int64
		c    string
	}{
		{"GenerateSessionId with two positive numbers",
			100, 200, "100200"},
		{"GenerateSessionId with two negative numbers",
			-1, -1, "-1-1"},
		{"GenerateSessionId with two MaxInt64",
			1<<63 - 1, 1<<63 - 1, "92233720368547758079223372036854775807"},
		{"GenerateSessionId with MaxInt64 and MinInt64",
			1<<63 - 1, -1 << 63, "-92233720368547758089223372036854775807"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, GenerateSessionId(c.a, c.b), c.c)
		})
	}
}

func TestGenerateSessionIdShouldNotEqual(t *testing.T) {
	cases := []struct {
		name string
		a, b int64
		c    string
	}{
		{"GenerateSessionId with two positive numbers",
			100, 200, "111222"},
		{"GenerateSessionId with two negative numbers, case 1",
			-1, -1, "-11"},
		{"GenerateSessionId with two negative numbers, case 2",
			-1, -1, "1-1"},
		{"GenerateSessionId with MaxInt64 and MinInt64",
			1<<63 - 1, -1 << 63, "9223372036854775807-9223372036854775808"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.NotEqual(t, GenerateSessionId(c.a, c.b), c.c)
		})
	}
}
