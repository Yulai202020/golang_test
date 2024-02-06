package hello

import (
	"testing"
)

func Test_sum(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		if sum(1, 2) != 3 {
			t.Error("Fail")
		} else {
			t.Logf("OK")
		}
	})
}
