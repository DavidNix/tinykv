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
		idx.Add(&Tuple{Value: "test", Key: "1"})
		idx.Add(&Tuple{Value: "test", Key: "1"})
		idx.Add(&Tuple{Value: "test", Key: "2", Deleted: true})
		idx.Add(&Tuple{Value: "another"})

		require.Equal(t, 2, idx.Count("test"))
	})

	t.Run("invalid input", func(t *testing.T) {
		idx := make(Index)

		require.Panics(t, func() {
			idx.Add(nil)
		})
	})
}
