package tinykv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore_SetGet(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		store := NewStore()
		store.Set("test", "1")
		got := store.Get("test")

		require.Equal(t, "1", got)
	})

	t.Run("key does not exist", func(t *testing.T) {
		store := NewStore()
		got := store.Get("test")

		require.Equal(t, "NULL", got)
	})
}
