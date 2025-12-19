package transaction

import (
	"database/sql"
	"errors"
	"testing"

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

func TestNewTransaction(t *testing.T) {
	tx := New()
	if tx == nil {
		t.Fatal("expected transaction instance")
	}
	if tx.tx != nil {
		t.Fatal("expected tx.tx to be nil")
	}
}

func TestOpenTransaction_NilReceiver(t *testing.T) {
	var tx *Transaction

	mustPanic(t, func() {
		tx.OpenTransaction(&sql.DB{})
	})
}

func TestOpenTransaction_AlreadyOpen(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	tr := New()
	tr.tx = &sql.Tx{}

	mustPanic(t, func() {
		tr.OpenTransaction(db)
	})
}

func TestOpenTransaction_BeginFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin().WillReturnError(errors.New("begin failed"))

	tr := New()

	mustPanic(t, func() {
		tr.OpenTransaction(db)
	})
}

func TestOpenTransaction_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()

	tr := New()
	tr.OpenTransaction(db)

	if tr.tx == nil {
		t.Fatal("expected tx to be set")
	}
}

func TestCommit_NilReceiver(t *testing.T) {
	var tr *Transaction

	mustPanic(t, func() {
		tr.Commit()
	})
}

func TestCommit_Fails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectCommit().WillReturnError(errors.New("commit failed"))

	tr := New()
	tr.OpenTransaction(db)

	mustPanic(t, func() {
		tr.Commit()
	})
}

func TestCommit_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectCommit()

	tr := New()
	tr.OpenTransaction(db)
	tr.Commit()

	if tr.tx != nil {
		t.Fatal("expected tx to be nil after commit")
	}
}

func TestRollback_RepropagatesPanic(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectRollback()

	tr := New()
	tr.OpenTransaction(db)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic")
		}
		if tr.tx != nil {
			t.Fatal("expected tx to be nil after rollback")
		}
	}()

	func() {
		defer tr.Rollback()
		panic("boom")
	}()
}

func TestRollback_NoTx_NoPanic(t *testing.T) {
	tr := New()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectRollback()

	tr.OpenTransaction(db)

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic")
		}
	}()

	func() {
		defer tr.Rollback()
		panic("boom")
	}()
}

type fakeQuery struct {
	exec func(*sql.Tx)
}

func (f *fakeQuery) Exec(tx *sql.Tx) {
	f.exec(tx)
}

func TestAtomic_NilQuery(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	tr := New()

	mustPanic(t, func() {
		tr.Atomic(db, nil)
	})
}

func TestAtomic_RollbackOnPanic(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectRollback()

	tr := New()

	q := &fakeQuery{
		exec: func(tx *sql.Tx) {
			panic("query failed")
		},
	}

	mustPanic(t, func() {
		tr.Atomic(db, q)
	})
}

func TestAtomic_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectCommit()

	tr := New()

	q := &fakeQuery{
		exec: func(tx *sql.Tx) {},
	}

	tr.Atomic(db, q)

	if tr.tx != nil {
		t.Fatal("expected tx to be nil after atomic")
	}
}

func Test_Dummy(t *testing.T) {
	var tx *Transaction
	mustPanic(t, tx.Rollback)
}
