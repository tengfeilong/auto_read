package auto

import (
	"testing"
)

func TestRandInt64(t *testing.T) {
	randInt64 := RandInt64(5, 10)
	t.Log(randInt64)
}
