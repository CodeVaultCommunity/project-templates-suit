package config

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// ---------- setenv() ----------

func Test_setenv(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		predefined   string // value set before test
		want         string
	}{
		{
			name:         "Should use existing environment variable",
			key:          "EXISTING_KEY",
			defaultValue: "default",
			predefined:   "already_set",
			want:         "already_set",
		},
		{
			name:         "Should use default value when key not set",
			key:          "NEW_KEY",
			defaultValue: "default_value",
			predefined:   "",
			want:         "default_value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.predefined != "" {
				_ = os.Setenv(tt.key, tt.predefined)
			} else {
				_ = os.Unsetenv(tt.key)
			}

			got, err := setenv(tt.key, tt.defaultValue)
			if err != nil {
				t.Fatalf("setenv() returned unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("setenv() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------- setAPIConfig() ----------

func Test_setAPIConfig(t *testing.T) {
	tests := []struct {
		name       string
		initialGin string
		wantErr    bool
		wantMode   string
	}{
		{
			name:       "Should set valid gin mode when not defined",
			initialGin: "",
			wantErr:    false,
			wantMode:   gin.DebugMode,
		},
		{
			name:       "Should fix invalid gin mode to default",
			initialGin: "INVALID_MODE",
			wantErr:    false,
			wantMode:   gin.DebugMode,
		},
		{
			name:       "Should keep valid gin mode (release)",
			initialGin: gin.ReleaseMode,
			wantErr:    false,
			wantMode:   gin.ReleaseMode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.initialGin != "" {
				_ = os.Setenv("GIN_MODE", tt.initialGin)
			} else {
				_ = os.Unsetenv("GIN_MODE")
			}

			err := checkAndSetOpitionalAPIKeys()
			if (err != nil) != tt.wantErr {
				t.Fatalf("setAPIConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := os.Getenv("GIN_MODE")
			if got != tt.wantMode {
				t.Errorf("GIN_MODE = %v, want %v", got, tt.wantMode)
			}
		})
	}
}

// ---------- RegisterSwagger() ----------
func TestRegisterSwagger(t *testing.T) {
	t.Run("Should not panic when group is nil", func(_ *testing.T) {
		RegisterSwagger(nil)
	})

	t.Run("Should register Swagger route in debug mode", func(_ *testing.T) {
		gin.SetMode(gin.DebugMode)
		router := gin.New()
		group := router.Group("/api")
		RegisterSwagger(group)
	})
}
