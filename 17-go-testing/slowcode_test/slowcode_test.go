package slowcode_test

import (
	"cbtnuggets/slowcode"
	"testing"
)

func TestRunSlowly(t *testing.T) {
	t.Run("First", func(t *testing.T) {
		t.Parallel()
		slowcode.RunSlowly(2)
	})
	t.Run("Second", func(t *testing.T) {
		t.Parallel()
		slowcode.RunSlowly(2)
	})
	t.Run("Third", func(t *testing.T) {
		t.Parallel()
		slowcode.RunSlowly(2)
	})

}
