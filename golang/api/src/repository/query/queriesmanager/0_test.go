package queriesmanager

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	errorsapp "mod_name/error"
)

/* ----------------------------------------------------
   Helpers
---------------------------------------------------- */

func mustPanicNilPointer(t *testing.T, f func()) {
	t.Helper()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic, got nil")
		}

		if _, ok := r.(*errorsapp.AppError); !ok {
			t.Fatalf("expected *AppError panic, got %T", r)
		}
	}()

	f()
}

/* ----------------------------------------------------
   Constructors
---------------------------------------------------- */

func TestNewQueriesManager(t *testing.T) {
	m := New()

	if m == nil {
		t.Fatal("expected non-nil manager")
	}

	if m.queries == nil {
		t.Fatal("expected queries list to be initialized")
	}

	if m.queries.IsEmpty() {
		// empty is expected, but list must exist
		return
	}
}

func TestNewQueryFromMultipleSQLFile(t *testing.T) {
	tmp := t.TempDir()

	files := []string{
		filepath.Join(tmp, "q1.sql"),
		filepath.Join(tmp, "q2.sql"),
	}

	if err := os.WriteFile(files[0], []byte("SELECT 1;"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(files[1], []byte("SELECT 2;"), 0644); err != nil {
		t.Fatal(err)
	}

	m := NewFromManyFiles(files, []any{}, []any{})

	if m == nil {
		t.Fatal("expected manager")
	}

	func() {
		defer func() {
			if p := recover(); p == nil {
				panic("Expect error, gets nil")
			}
		}()

		NewFromManyFiles(files)
	}()

	if m.queries.IsEmpty() {
		t.Fatal("expected queries to be added")
	}
}

/* ----------------------------------------------------
   Add methods
---------------------------------------------------- */

func TestAddQuery(t *testing.T) {
	m := New()

	m.AddQuery("SELECT 1", 1)

	if m.queries.IsEmpty() {
		t.Fatal("expected query to be added")
	}
}

func TestAddQuery_NilManager(t *testing.T) {
	var m *QueriesManager

	mustPanicNilPointer(t, func() {
		m.AddQuery("boom")
	})
}

func TestAddCodeFromSQLFile(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "q.sql")

	if err := os.WriteFile(path, []byte("SELECT now();"), 0644); err != nil {
		t.Fatal(err)
	}

	m := New()
	m.AddCodeFromSQLFile(path)

	if m.queries.IsEmpty() {
		t.Fatal("expected query to be added from sql file")
	}
}

func TestAddCodeFromSQLFiles(t *testing.T) {
	tmp := t.TempDir()

	files := []string{
		filepath.Join(tmp, "q1.sql"),
		filepath.Join(tmp, "q2.sql"),
	}

	if err := os.WriteFile(files[0], []byte("SELECT 1;"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(files[1], []byte("SELECT 2;"), 0644); err != nil {
		t.Fatal(err)
	}

	args := [][]any{
		{1},
		{"a"},
	}

	m := New()
	m.AddCodeFromSQLFiles(files, args...)

	if m.queries.IsEmpty() {
		t.Fatal("expected queries to be added")
	}
}

func TestAddCodeFromSQLFiles_NilManager(t *testing.T) {
	var m *QueriesManager

	mustPanicNilPointer(t, func() {
		m.AddCodeFromSQLFiles([]string{})
	})
}

/* ----------------------------------------------------
   GetQueries
---------------------------------------------------- */

func TestGetQueries(t *testing.T) {
	m := New()

	q := m.GetQueries()
	if q == nil {
		t.Fatal("expected queries list")
	}
}

func TestGetQueries_NilManager(t *testing.T) {
	var m *QueriesManager

	mustPanicNilPointer(t, func() {
		_ = m.GetQueries()
	})

	mustPanicNilPointer(t, func() { m.AddCodeFromSQLFile("") })
}

func TestAddCodeFromSQLFiles_LengthMismatch(t *testing.T) {
	m := New()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic")
		}

		if _, ok := r.(*errorsapp.AppError); !ok {
			t.Fatalf("expected *AppError, got %T", r)
		}
	}()

	m.AddCodeFromSQLFiles(
		[]string{"a.sql", "b.sql"},
	)
}

// Test Exec

func TestNilExec(t *testing.T) {
	var manager *QueriesManager

	mustPanicNilPointer(t, func() { manager.Exec(nil) })

	manager = New()
	mustPanicNilPointer(t, func() { manager.Exec(nil) })
}

func TestExec(t *testing.T) {
	manager := New()
	manager.AddQuery("")
	tx := &sql.Tx{}

	defer func() {
		if p := recover(); p == nil {
			t.Fatal("Expect panic")
		}
	}()
	manager.Exec(tx)
}
