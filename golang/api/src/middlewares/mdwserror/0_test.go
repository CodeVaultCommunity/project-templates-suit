package mdwserror

import (
	"bytes"
	"encoding/json"
	"errors"
	errorsapp "mod_name/error"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestErrorHandlerMiddlewareGetter_Release(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	mw := ErrorHandlerMiddlewareGetter()
	if mw == nil {
		t.Fatal("expected middleware, got nil")
	}
}

func TestErrorHandlerMiddlewareGetter_Debug(t *testing.T) {
	gin.SetMode(gin.DebugMode)

	mw := ErrorHandlerMiddlewareGetter()
	if mw == nil {
		t.Fatal("expected middleware, got nil")
	}
}

func performRequest(r *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/panic", nil)
	r.ServeHTTP(w, req)
	return w
}

func decodeBody(t *testing.T, w *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var body map[string]any
	if err := json.NewDecoder(bytes.NewBuffer(w.Body.Bytes())).Decode(&body); err != nil {
		t.Fatalf("invalid json body: %v", err)
	}
	return body
}

func TestRelease_AppError(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(errorHandlerMiddleware)
	r.GET("/panic", func(c *gin.Context) {
		panic(errorsapp.ErrInvalidJSON)
	})

	w := performRequest(r)
	body := decodeBody(t, w)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}

	if body["code"] != errorsapp.RequestInvalidJSON {
		t.Fatal("expected error code")
	}

	if _, ok := body["debug"]; ok {
		t.Fatal("debug info should not be present in release")
	}
}

func TestRelease_InternalError(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(errorHandlerMiddleware)
	r.GET("/panic", func(c *gin.Context) {
		panic(errors.New("boom"))
	})

	w := performRequest(r)
	body := decodeBody(t, w)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}

	if body["error"] != "Internal server error" {
		t.Fatal("unexpected error message")
	}
}

func TestRelease_Unforeseen(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(errorHandlerMiddleware)
	r.GET("/panic", func(c *gin.Context) {
		panic(123)
	})

	w := performRequest(r)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}

func TestDebug_AppError(t *testing.T) {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(errorHandlerMiddlewareDebug)
	r.GET("/panic", func(c *gin.Context) {
		panic(errorsapp.ErrInvalidJSON)
	})

	w := performRequest(r)
	body := decodeBody(t, w)

	if _, ok := body["debug"]; !ok {
		t.Fatal("expected debug info")
	}
}

func TestDebug_InternalError(t *testing.T) {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(errorHandlerMiddlewareDebug)
	r.GET("/panic", func(c *gin.Context) {
		panic(errors.New("boom"))
	})

	w := performRequest(r)
	body := decodeBody(t, w)

	if _, ok := body["debug"]; !ok {
		t.Fatal("expected debug info")
	}
}

func TestDebug_Unforeseen(t *testing.T) {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(errorHandlerMiddlewareDebug)
	r.GET("/panic", func(c *gin.Context) {
		panic(42)
	})

	w := performRequest(r)
	body := decodeBody(t, w)

	if body["r"] == nil {
		t.Fatal("expected raw panic value in debug")
	}
}
