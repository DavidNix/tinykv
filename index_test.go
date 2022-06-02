package tinykv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		idx := make(Index)

		tup := &Tuple{Value: "test"}

		idx.Add(tup)
		idx.Add(&Tuple{Value: "test"})
		idx.Add(&Tuple{Value: "test", Deleted: true})
		idx.Add(&Tuple{Value: "another"})

		require.Equal(t, 2, idx.Count("test"))

		idx.Remove(tup)
		require.Equal(t, 1, idx.Count("test"))

		require.NotPanics(t, func() {
			idx.Remove(&Tuple{Value: "never added"})
		})
	})

	t.Run("invalid input", func(t *testing.T) {
		idx := make(Index)

		require.Panics(t, func() {
			idx.Add(nil)
		})
		require.Panics(t, func() {
			idx.Remove(nil)
		})
	})

}
