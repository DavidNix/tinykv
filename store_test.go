package tinykv

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore_SetGet(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		store := NewStore()
		store.Set("test", "1")
		got := store.Get("test")

		require.Equal(t, "1", got)

		store.Set("test", "2")
		got = store.Get("test")

		require.Equal(t, "2", got)
	})

	t.Run("key does not exist", func(t *testing.T) {
		store := NewStore()
		got := store.Get("test")

		require.Equal(t, "NULL", got)
	})
}

func TestStore_Delete(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		store := NewStore()
		store.Set("test", "1")
		store.Delete("test")
		got := store.Get("test")

		require.Equal(t, "NULL", got)
	})
}

func TestStore_Count(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		store := NewStore()
		for i := 0; i < 5; i++ {
			store.Set(strconv.Itoa(i), "value")
			store.Set(strconv.Itoa(i+10), "another value")
		}
		got := store.Count("value")

		require.Equal(t, 5, got)
	})

	t.Run("value does not exist", func(t *testing.T) {
		store := NewStore()
		got := store.Count("nope")

		require.Zero(t, got)
	})

	t.Run("deletes", func(t *testing.T) {
		store := NewStore()
		for i := 0; i < 5; i++ {
			store.Set(strconv.Itoa(i), "value")
			store.Set(strconv.Itoa(i), "value")
		}
		store.Delete("1")
		store.Delete("1") // repeat on purpose
		store.Delete("2")

		got := store.Count("value")
		require.Equal(t, 3, got)

		store.Delete("0")
		store.Delete("3")
		store.Delete("4")
		store.Delete("5")
		store.Delete("5") // repeat on purpose

		got = store.Count("value")
		require.Zero(t, got)
	})
}
