package safeMap

import "testing"

func TestSet(t *testing.T) {
	m := New()

	m.Set("1", "1")
	m.Set("1", "3")
	m.Set("1", "2")
}
