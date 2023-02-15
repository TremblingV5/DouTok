package safeMap

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	m := New()
	log.Println(m)
}
