package repository

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"mod_name/config"

	"github.com/DATA-DOG/go-sqlmock"
)

func mustPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic, got none")
		}
	}()
	f()
}

/* ----------------------------------------------------
   DatabaseInit
---------------------------------------------------- */

func TestDatabaseInit_OpenFails(t *testing.T) {
	config.Engine = "invalid-driver"
	config.DatabaseDNS = "dsn"

	if err := DatabaseInit(); err == nil {
		t.Fatal("expected error from sql.Open")
	}
}

/* ----------------------------------------------------
   DatabaseClose
---------------------------------------------------- */

func TestDatabaseClose_NilDB(t *testing.T) {
	SQLDB = nil

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic when DB is nil")
		}
	}()

	DatabaseClose()
}

// utils.go
func TestFastTransaction_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	SQLDB = db

	mock.ExpectBegin()
	mock.ExpectExec("SELECT 1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	tmp := t.TempDir()
	sqlfile := filepath.Join(tmp, "q.sql")

	if err := os.WriteFile(sqlfile, []byte("SELECT 1"), 0644); err != nil {
		t.Fatal(err)
	}

	result := FastAtomicTransaction(sqlfile)

	if result == nil {
		t.Fatal("expected sql.Result")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestFastTransaction_ExecFails_Rollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	SQLDB = db

	mock.ExpectBegin()
	mock.ExpectExec("SELECT 1").
		WillReturnError(errors.New("exec failed"))
	mock.ExpectRollback()

	tmp := t.TempDir()
	sqlfile := filepath.Join(tmp, "q.sql")

	if err := os.WriteFile(sqlfile, []byte("SELECT 1"), 0644); err != nil {
		t.Fatal(err)
	}

	mustPanic(t, func() {
		FastAtomicTransaction(sqlfile)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestFastTransactions_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	SQLDB = db

	mock.ExpectBegin()
	mock.ExpectExec("SELECT 1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("SELECT 2").
		WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()

	tmp := t.TempDir()

	files := []string{
		filepath.Join(tmp, "q1.sql"),
		filepath.Join(tmp, "q2.sql"),
	}

	if err := os.WriteFile(files[0], []byte("SELECT 1"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(files[1], []byte("SELECT 2"), 0644); err != nil {
		t.Fatal(err)
	}

	results := FastAtomicTransactions(
		files,
		[]any{},
		[]any{},
	)

	if results == nil || results.IsEmpty() {
		t.Fatal("expected results")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestFastTransactions_PartialFailure_Rollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	SQLDB = db

	mock.ExpectBegin()
	mock.ExpectExec("SELECT 1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("SELECT 2").
		WillReturnError(errors.New("boom"))
	mock.ExpectRollback()

	tmp := t.TempDir()

	files := []string{
		filepath.Join(tmp, "q1.sql"),
		filepath.Join(tmp, "q2.sql"),
	}

	if err := os.WriteFile(files[0], []byte("SELECT 1"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(files[1], []byte("SELECT 2"), 0644); err != nil {
		t.Fatal(err)
	}

	mustPanic(t, func() {
		FastAtomicTransactions(
			files,
			[]any{},
			[]any{},
		)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
