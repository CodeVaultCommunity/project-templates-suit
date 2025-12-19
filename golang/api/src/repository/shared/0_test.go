package reposhared

import (
	"os"
	"path/filepath"
	"testing"

	errorsapp "mod_name/error"
)

/* ----------------------------------------------------
   Helpers
---------------------------------------------------- */

func mustPanicAppError(t *testing.T, f func()) {
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
   ReadSQLFileAsBytes
---------------------------------------------------- */

func TestReadSQLFileAsBytes_FileNotFound(t *testing.T) {
	mustPanicAppError(t, func() {
		ReadSQLFileAsBytes("definitely-not-exists.sql")
	})
}

func TestReadSQLFileAsBytes_EmptyFile(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "empty.sql")

	if err := os.WriteFile(path, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	mustPanicAppError(t, func() {
		ReadSQLFileAsBytes(path)
	})
}

func TestReadSQLFileAsBytes_Success(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "query.sql")

	content := []byte("SELECT 1;")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}

	b := ReadSQLFileAsBytes(path)

	if b == nil {
		t.Fatal("expected non-nil byte pointer")
	}

	if b[0] != content[0] {
		t.Fatalf("expected first byte %v, got %v", content[0], b[0])
	}
}

/* ----------------------------------------------------
   ReadSQLFile
---------------------------------------------------- */

func TestReadSQLFile_Success(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "query.sql")

	expected := "SELECT * FROM users;"
	if err := os.WriteFile(path, []byte(expected), 0644); err != nil {
		t.Fatal(err)
	}

	result := ReadSQLFile(path)

	if result != expected {
		t.Fatalf("expected %q, got %q", expected, result)
	}
}
