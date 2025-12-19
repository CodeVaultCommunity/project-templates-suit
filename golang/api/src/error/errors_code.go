// Package errorsapp - This file provides some default utils errors
package errorsapp

// Authentication Errors
const (
	AuthInvalidCredentials = "auth.invalid_credentials"
	AuthTokenExpired       = "auth.token_expired"
	AuthTokenGeneration    = "auth.token_generation_failed"
)

// Database Errors
const (
	DBSQLReadError   = "db.sql_read"
	DBQueryFailed    = "db.query_failed"
	DBInsertFailed   = "db.insert_failed"
	DBConnectionFail = "db.connection_failed"
	DBDeleteFailed   = "db.delete_failed"
)

// Internal Server Errros
const (
	SystemInternal   = "system.internal"
	SystemConfig     = "system.config_error"
	SystemNilPointer = "system.nil_pointer_error"
)

// Internal Server Errros - Collections Module
const (
	SystemCollectionsIndexOutOfBound = "system.collections.IndexOutOfBound"
)

// Bad Requests Errors
const (
	RequestInvalidJSON     = "request.invalid_json"
	RequestInvalidURLParam = "request.invalid_url_param"
	RequestNotFound        = "request.not_found"
	RequestMissingFields   = "request.missing_fields"
	RequestUnauthorized    = "request.unauthorized"
)
