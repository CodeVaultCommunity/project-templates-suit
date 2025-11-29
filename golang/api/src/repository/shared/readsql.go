package shared

import (
	errorsapp "mod_name/error"
	"net/http"
	"os"
)

func ReadSQLFileAsBytes(filepath string) *byte {
	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(errorsapp.New(http.StatusInternalServerError, errorsapp.DBSQLReadError, "sql file is empty", err))
	}

	if len(content) == 0 {
		panic(errorsapp.New(http.StatusInternalServerError, errorsapp.DBSQLReadError, "sql file is empty", nil))
	}

	return &content[0]
}

func ReadSQLFile(filepath string) string {
	return string(*ReadSQLFileAsBytes(filepath))
}
