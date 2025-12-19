// Package reposhared provides shared functions that could be used for all repository package and subpackage
package reposhared

import (
	errorsapp "mod_name/error"
	"net/http"
	"os"
)

// ReadSQLFileAsBytes load the content at sql file `filepath` as a byte array
func ReadSQLFileAsBytes(filepath string) []byte {
	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(errorsapp.New(http.StatusInternalServerError, errorsapp.DBSQLReadError, "sql file is empty", err))
	}

	if len(content) == 0 {
		panic(errorsapp.New(http.StatusInternalServerError, errorsapp.DBSQLReadError, "sql file is empty", nil))
	}

	return content
}

// ReadSQLFile load the content at sql file `filepath` as a string
func ReadSQLFile(filepath string) string {
	return string(ReadSQLFileAsBytes(filepath))
}
