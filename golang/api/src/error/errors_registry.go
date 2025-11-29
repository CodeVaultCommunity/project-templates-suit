// Package errorsapp - This file provides some default errors
package errorsapp

import (
	"errors"
	"net/http"
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

func ErrNilPointer() *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    SystemInternal,
		Message: "InternalServerError",
		Err:     errors.New(SystemNilPointer),
	}
}

// System - Collections Module

func ErrIndextOutOfBound() *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    SystemInternal,
		Message: "InternalServerError",
		Err:     errors.New(SystemCollectionsIndexOutOfBound),
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
)

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
