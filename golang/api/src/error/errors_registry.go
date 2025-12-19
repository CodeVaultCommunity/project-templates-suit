// Package errorsapp - This file provides some default errors
package errorsapp

import (
	"errors"
	"net/http"
	"strings"
)

// Authentication Errors
var (
	ErrInvalidCredentials = &AppError{
		Status:  http.StatusUnauthorized,
		Code:    AuthInvalidCredentials,
		Message: "Invalid login or password",
	}
	ErrTokenExpired = &AppError{
		Status:  http.StatusUnauthorized,
		Code:    AuthTokenExpired,
		Message: "Expired token",
	}
	ErrTokenGeneration = &AppError{
		Status:  http.StatusInternalServerError,
		Code:    AuthTokenGeneration,
		Message: "Error generating JWT token",
	}
)

// System
var (
	ErrInternal = &AppError{
		Status:  http.StatusInternalServerError,
		Code:    SystemInternal,
		Message: "InternalServerError",
	}
)

// ErrNilPointer creates a new instance of `ErrNilPointer`
// to capture the stack trace
func ErrNilPointer() *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    SystemInternal,
		Message: "InternalServerError",
		Err:     errors.New(SystemNilPointer),
	}
}

// ErrListIsEmpty creates a new instance of `ErrListIsEmpty`
// to capture the stack trace
func ErrListIsEmpty() *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    SystemInternal,
		Message: "InternalServerError",
		Err:     errors.New("List is Empty"),
	}
}

// System - Collections Module

// ErrIndexOutOfBound creates a new instance of `ErrIndexOutOfBound`
// to capture the stack trace
func ErrIndexOutOfBound() *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    SystemInternal,
		Message: "InternalServerError",
		Err:     errors.New(SystemCollectionsIndexOutOfBound),
	}
}

// System - Repository Module

// ErrMismatchedArgs creates a new instance of `ErrMismatchedArgs`
// to capture the stack trace
func ErrMismatchedArgs() *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    SystemInternal,
		Message: "filespath and args length mismatch",
		Err:     errors.New("len(filespath) != len(args)"),
	}
}

// Requests Errors
var (
	ErrInvalidJSON = &AppError{
		Status:  http.StatusBadRequest,
		Code:    RequestInvalidJSON,
		Message: "Invalid JSON",
	}
	ErrUnauthorized = &AppError{
		Status:  http.StatusUnauthorized,
		Code:    RequestUnauthorized,
		Message: "Unauthorized access",
	}
	ErrNotFound = &AppError{
		Status:  http.StatusNotFound,
		Code:    RequestNotFound,
		Message: "Object not found",
	}
)

// ErrNotOptionalURLParam creates a new instance of `ErrNotOptionalURLParam`
// to capture the stack trace
func ErrNotOptionalURLParam(paramName string) *AppError {
	builder := strings.Builder{}
	builder.WriteString("arggument ")
	builder.WriteString(paramName)
	builder.WriteString(" is not optional")

	return &AppError{
		Status:  http.StatusBadRequest,
		Code:    RequestInvalidURLParam,
		Message: builder.String(),
	}
}

// ErrURLParamInvalidType creates a new instance of `ErrURLParamInvalidType`
// to capture the stack trace
func ErrURLParamInvalidType(paramName string, expectedType string) *AppError {
	builder := strings.Builder{}
	builder.WriteString("arggument ")
	builder.WriteString(paramName)
	builder.WriteString(" should be ")
	builder.WriteString(expectedType)

	return &AppError{
		Status:  http.StatusBadRequest,
		Code:    RequestInvalidURLParam,
		Message: builder.String(),
	}
}

// Database
var (
	ErrSQLRead = &AppError{
		Status:  http.StatusInternalServerError,
		Code:    DBSQLReadError,
		Message: "Error while reading SQL file",
	}
	ErrSQLQuery = &AppError{
		Status:  http.StatusInternalServerError,
		Code:    DBQueryFailed,
		Message: "Error while execute query",
	}
	ErrSQLInsert = &AppError{
		Status:  http.StatusInternalServerError,
		Code:    DBInsertFailed,
		Message: "Error while inserting data into the database",
	}
	ErrDBConnection = &AppError{
		Status:  http.StatusInternalServerError,
		Code:    DBConnectionFail,
		Message: "Database connection failed",
	}
	ErrSQLDelete = &AppError{
		Status:  http.StatusInternalServerError,
		Code:    DBDeleteFailed,
		Message: "Error while deleting data from the database",
	}
)
