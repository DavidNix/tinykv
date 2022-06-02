package tinykv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDatabase_Feature(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		db := NewDatabase()

		require.Equal(t, "NULL", db.Get("a"))

		db.Set("a", "foo")
		db.Set("b", "foo")

		require.Equal(t, 2, db.Count("foo"))
		require.Zero(t, db.Count("bar"))

		db.Delete("a")
		require.Equal(t, 1, db.Count("foo"))

		db.Set("b", "baz")
		require.Equal(t, 0, db.Count("foo"))

		require.Equal(t, "baz", db.Get("b"))
		require.Equal(t, "NULL", db.Get("B"))
	})

	t.Run("example 2", func(t *testing.T) {
		db := NewDatabase()
		db.Set("a", "foo")
		db.Set("a", "foo")

		require.Equal(t, 1, db.Count("foo"))
		require.Equal(t, "foo", db.Get("a"))

		db.Delete("a")
		require.Equal(t, "NULL", db.Get("a"))
		require.Zero(t, db.Count("foo"))
	})

	t.Run("example 3", func(t *testing.T) {
		t.Skip("TODO")
	})
}
