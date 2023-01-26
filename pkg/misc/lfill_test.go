package misc

import (
	"testing"
)

func TestLFill(t *testing.T) {
	src := "23"

	res := LFill(src, 6)

	if res != "000023" {
		panic("LFill defeat")
	}
}
