package errorsapp

import (
	"errors"
	"net/http"
	"testing"
)

func TestAppError_Error_WithUnderlyingError(t *testing.T) {
	baseErr := errors.New("base error")

	appErr := &AppError{
		Message: "wrapper message",
		Err:     baseErr,
	}

	if appErr.Error() != "base error" {
		t.Fatalf("expected underlying error message, got %q", appErr.Error())
	}
}

func TestAppError_Error_WithoutUnderlyingError(t *testing.T) {
	appErr := &AppError{
		Message: "only message",
	}

	if appErr.Error() != "only message" {
		t.Fatalf("expected message, got %q", appErr.Error())
	}
}

func TestNewAppError(t *testing.T) {
	err := errors.New("inner")

	appErr := New(
		http.StatusBadRequest,
		"TEST_CODE",
		"test message",
		err,
	)

	if appErr.Status != http.StatusBadRequest {
		t.Fatalf("unexpected status: %d", appErr.Status)
	}
	if appErr.Code != "TEST_CODE" {
		t.Fatalf("unexpected code: %s", appErr.Code)
	}
	if appErr.Message != "test message" {
		t.Fatalf("unexpected message: %s", appErr.Message)
	}
	if appErr.Err != err {
		t.Fatal("underlying error not set correctly")
	}
}

/* ----------------------------------------------------
   Static errors (var)
---------------------------------------------------- */

func TestStaticErrors_NotNil(t *testing.T) {
	tests := []struct {
		name string
		err  *AppError
	}{
		{"ErrInvalidCredentials", ErrInvalidCredentials},
		{"ErrTokenExpired", ErrTokenExpired},
		{"ErrTokenGeneration", ErrTokenGeneration},
		{"ErrInternal", ErrInternal},
		{"ErrInvalidJSON", ErrInvalidJSON},
		{"ErrUnauthorized", ErrUnauthorized},
		{"ErrSQLRead", ErrSQLRead},
		{"ErrSQLQuery", ErrSQLQuery},
		{"ErrSQLInsert", ErrSQLInsert},
		{"ErrDBConnection", ErrDBConnection},
		{"ErrSQLDelete", ErrSQLDelete},
		{"ErrNotFound", ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Fatal("expected error, got nil")
			}
			if tt.err.Message == "" {
				t.Fatal("expected message to be set")
			}
			if tt.err.Code == "" {
				t.Fatal("expected code to be set")
			}
		})
	}
}

/* ----------------------------------------------------
   Factory errors (func)
---------------------------------------------------- */

func TestErrNilPointer(t *testing.T) {
	e1 := ErrNilPointer()
	e2 := ErrNilPointer()

	if e1 == e2 {
		t.Fatal("expected different instances")
	}

	if e1.Err == nil {
		t.Fatal("expected underlying error")
	}

	if e1.Error() != SystemNilPointer {
		t.Fatalf("unexpected error message: %s", e1.Error())
	}
}

func TestErrMismatchedArgs(t *testing.T) {
	e1 := ErrMismatchedArgs()
	e2 := ErrMismatchedArgs()

	if e1 == e2 {
		t.Fatal("expected different instances")
	}

	if e1.Err == nil {
		t.Fatal("expected underlying error")
	}

	if e1.Error() != "len(filespath) != len(args)" {
		t.Fatalf("unexpected error message: %s", e1.Error())
	}
}

func TestErrNotOptionalQueryParam(t *testing.T) {
	e1 := ErrNotOptionalURLParam("test")
	e2 := ErrNotOptionalURLParam("test")

	if e1 == e2 {
		t.Fatal("expected different instances")
	}

	if e1.Err == nil {
		t.Fatal("expected underlying error")
	}

	if e1.Error() != "arggument test is not optional" {
		t.Fatalf("unexpected error message: %s", e1.Error())
	}
}

func TestErrQueryParamInvalidType(t *testing.T) {
	e1 := ErrURLParamInvalidType("test", "string")
	e2 := ErrURLParamInvalidType("test", "string")

	if e1 == e2 {
		t.Fatal("expected different instances")
	}

	if e1.Err == nil {
		t.Fatal("expected underlying error")
	}

	if e1.Error() != "arggument test should be string" {
		t.Fatalf("unexpected error message: %s", e1.Error())
	}
}

func TestErrListIsEmpty(t *testing.T) {
	err := ErrListIsEmpty()

	if err.Err == nil {
		t.Fatal("expected underlying error")
	}

	if err.Error() != "List is Empty" {
		t.Fatalf("unexpected message: %s", err.Error())
	}
}

func TestErrIndexOutOfBound(t *testing.T) {
	err := ErrIndexOutOfBound()

	if err.Err == nil {
		t.Fatal("expected underlying error")
	}

	if err.Error() != SystemCollectionsIndexOutOfBound {
		t.Fatalf("unexpected message: %s", err.Error())
	}
}
