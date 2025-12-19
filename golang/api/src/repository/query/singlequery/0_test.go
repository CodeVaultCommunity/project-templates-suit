package singlequery

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	errorsapp "mod_name/error"

	"github.com/DATA-DOG/go-sqlmock"
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

func TestNewQuery(t *testing.T) {
	q := New()

	if q == nil {
		t.Fatal("expected non-nil Query")
	}

	if q.Result != nil {
		t.Fatal("expected Result to be nil")
	}
}

func TestNewQueryWithCode(t *testing.T) {
	q := NewWithCode("SELECT 1", 1, "a")

	if q == nil {
		t.Fatal("expected non-nil Query")
	}

	if q.code != "SELECT 1" {
		t.Fatalf("unexpected code: %s", q.code)
	}

	if len(q.args) != 2 {
		t.Fatalf("unexpected args length: %d", len(q.args))
	}
}

func TestNewQueryFromSQLFile(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "query.sql")

	expected := "SELECT * FROM users;"
	if err := os.WriteFile(path, []byte(expected), 0644); err != nil {
		t.Fatal(err)
	}

	q := NewFromSQLFile(path, 42)

	if q.code != expected {
		t.Fatalf("unexpected code: %s", q.code)
	}

	if len(q.args) != 1 || q.args[0] != 42 {
		t.Fatal("unexpected args")
	}
}

/* ----------------------------------------------------
   Setters
---------------------------------------------------- */

func TestSetCode(t *testing.T) {
	q := New()
	q.SetCode("UPDATE table")

	if q.code != "UPDATE table" {
		t.Fatalf("unexpected code: %s", q.code)
	}
}

func TestSetCode_NilReceiver(t *testing.T) {
	var q *Query

	mustPanicNilPointer(t, func() {
		q.SetCode("boom")
	})
}

func TestSetArgs(t *testing.T) {
	q := New()
	q.SetArgs(1, "a", true)

	if len(q.args) != 3 {
		t.Fatalf("unexpected args length: %d", len(q.args))
	}
}

func TestSetCodeFromSQLFile(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "query.sql")

	expected := "SELECT now();"
	if err := os.WriteFile(path, []byte(expected), 0644); err != nil {
		t.Fatal(err)
	}

	q := New()
	q.SetCodeFromSQLFile(path)

	if q.code != expected {
		t.Fatalf("unexpected code: %s", q.code)
	}
}

func TestSetCodeFromSQLFile_NilReceiver(t *testing.T) {
	var q *Query

	mustPanicNilPointer(t, func() {
		q.SetCodeFromSQLFile("any.sql")
	})

	mustPanicNilPointer(t, func() { q.SetArgs() })
}

// Test Exec
func Test_ExecPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Fatal("Expected Panic")
		}
	}()

	q := New()
	tx := &sql.Tx{}
	q.Exec(tx)
}

func Test_NilExecPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Fatal("Expected Panic")
		}
	}()

	q := New()
	q.Exec(nil)
}

func Test_NilPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Fatal("Expected Panic")
		}
	}()

	var q Query
	q.Exec(nil)
}

func Test_Exec(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Fatal("Expected Panic")
		}
	}()

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	q := New()
	tx, _ := db.Begin()
	q.Exec(tx)
}
