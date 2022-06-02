package tinykv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDatabase_BeginRollback(t *testing.T) {
	t.Run("deletes", func(t *testing.T) {
		db := NewDatabase()
		db.Set("test", "1")
		db.Set("delete", "found")

		db.Begin()

		db.Set("test", "2")
		db.Delete("delete")

		require.Equal(t, "2", db.Get("test"))
		require.Equal(t, "NULL", db.Get("delete"))

		db.Rollback()

		require.Equal(t, "1", db.Get("test"))
		require.Equal(t, "found", db.Get("delete"))
	})

	t.Run("counts", func(t *testing.T) {
		db := NewDatabase()
		db.Set("test1", "1")
		db.Set("test2", "1")
		require.Equal(t, 2, db.Count("1"))

		db.Begin()

		db.Set("test1", "2")
		require.Equal(t, 1, db.Count("1"))

		db.Delete("test2")
		require.Zero(t, db.Count("1"))

		require.True(t, db.Rollback())

		require.Equal(t, 2, db.Count("1"))
	})

	t.Run("no transactions", func(t *testing.T) {
		db := NewDatabase()
		require.False(t, db.Rollback())

		db.Begin()
		require.True(t, db.Rollback())
	})
}

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
		db := NewDatabase()

		db.Begin()
		db.Set("a", "foo")
		require.Equal(t, "foo", db.Get("a"))

		db.Begin()
		db.Set("a", "bar")
		require.Equal(t, "bar", db.Get("a"))

		db.Set("a", "baz")
		db.Rollback()
		require.Equal(t, "foo", db.Get("a"))

		db.Rollback()
		require.Equal(t, "NULL", db.Get("a"))
	})

	t.Run("example 4", func(t *testing.T) {
		t.Skip("TODO")
	})
}
