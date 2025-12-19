package config

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

/* ----------------------------------------------------
   Helpers
---------------------------------------------------- */

func withEnv(t *testing.T, env map[string]string, f func()) {
	t.Helper()

	backup := os.Environ()
	os.Clearenv()

	for k, v := range env {
		_ = os.Setenv(k, v)
	}

	defer func() {
		os.Clearenv()
		for _, kv := range backup {
			pair := splitEnv(kv)
			_ = os.Setenv(pair[0], pair[1])
		}
	}()

	f()
}

func splitEnv(kv string) [2]string {
	for i := 0; i < len(kv); i++ {
		if kv[i] == '=' {
			return [2]string{kv[:i], kv[i+1:]}
		}
	}
	return [2]string{kv, ""}
}

/* ----------------------------------------------------
   setenv
---------------------------------------------------- */

func Test_setenv(t *testing.T) {
	withEnv(t, map[string]string{}, func() {
		val, err := setenv("A", "default")
		if err != nil {
			t.Fatal(err)
		}
		if val != "default" {
			t.Fatalf("expected default, got %s", val)
		}

		val, err = setenv("A", "other")
		if err != nil {
			t.Fatal(err)
		}
		if val != "default" {
			t.Fatalf("expected preserved value, got %s", val)
		}
	})
}

/* ----------------------------------------------------
   checkNoOptinalAPIKeys
---------------------------------------------------- */

func Test_checkNoOptinalAPIKeys_DatabaseURL(t *testing.T) {
	withEnv(t, map[string]string{
		"DATABASE_URL": "postgres://x",
		"ENGINE":       "postgre",
	}, func() {
		if err := checkNoOptinalAPIKeys(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func Test_checkNoOptinalAPIKeys_ConstructedDNS(t *testing.T) {
	withEnv(t, map[string]string{
		"ENGINE":           "postgres",
		"DB_USER":          "u",
		"DB_USER_PASSWORD": "p",
		"DB_NAME":          "n",
		"DB_HOST":          "h",
		"DB_PORT":          "5432",
	}, func() {
		if err := checkNoOptinalAPIKeys(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func Test_checkNoOptinalAPIKeys_ConstructedDNS_PostegreFalltrhough(t *testing.T) {
	withEnv(t, map[string]string{
		"ENGINE":           "postgre",
		"DB_USER":          "u",
		"DB_USER_PASSWORD": "p",
		"DB_NAME":          "n",
		"DB_HOST":          "h",
		"DB_PORT":          "5432",
	}, func() {
		if err := checkNoOptinalAPIKeys(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		Engine = "postgre"
		if err := buildDatabaseDNS(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		t.Log(DatabaseDNS)
	})
}

func Test_checkNoOptinalAPIKeys_MissingKey(t *testing.T) {
	withEnv(t, map[string]string{
		"ENGINE": "postgres",
	}, func() {
		if err := checkNoOptinalAPIKeys(); err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

/* ----------------------------------------------------
   checkAndSetOpitionalAPIKeys
---------------------------------------------------- */

func Test_checkAndSetOpitionalAPIKeys_Defaults(t *testing.T) {
	withEnv(t, map[string]string{}, func() {
		if err := checkAndSetOptionalAPIKeys(); err != nil {
			t.Fatal(err)
		}

		if os.Getenv("API_PORT") != ":8080" {
			t.Fatal("API_PORT default not applied")
		}

		if os.Getenv("GIN_MODE") != gin.DebugMode {
			t.Fatal("GIN_MODE default not applied")
		}

		if os.Getenv("SUBPATH") != "/" {
			t.Fatal("SUBPATH default not applied")
		}
	})
}

func Test_checkAndSetOpitionalAPIKeys_InvalidPort(t *testing.T) {
	withEnv(t, map[string]string{
		"API_PORT": "3000",
	}, func() {
		if err := checkAndSetOptionalAPIKeys(); err != nil {
			t.Fatal(err)
		}

		if os.Getenv("API_PORT") != ":3000" {
			t.Fatal("API_PORT not normalized")
		}
	})
}

func Test_checkAndSetOpitionalAPIKeys_EmptyPort(t *testing.T) {
	withEnv(t, map[string]string{
		"API_PORT": "",
	}, func() {
		if err := checkAndSetOptionalAPIKeys(); err != nil {
			t.Fatal(err)
		}

		if os.Getenv("API_PORT") != "" {
			t.Fatal("empty API_PORT not handled safely")
		}
	})
}

func Test_checkAndSetOpitionalAPIKeys_InvalidGinMode(t *testing.T) {
	withEnv(t, map[string]string{
		"GIN_MODE": "INVALID",
	}, func() {
		if err := checkAndSetOptionalAPIKeys(); err != nil {
			t.Fatal(err)
		}

		if os.Getenv("GIN_MODE") != gin.DebugMode {
			t.Fatal("invalid GIN_MODE not fixed")
		}
	})
}

/* ----------------------------------------------------
   buildDatabaseDNS / loadConstraints
---------------------------------------------------- */

func Test_loadConstraints_UsesDatabaseURL(t *testing.T) {
	withEnv(t, map[string]string{
		"ENGINE":       "postgres",
		"DATABASE_URL": "postgres://custom",
	}, func() {
		if err := loadConstraints(); err != nil {
			t.Fatal(err)
		}

		if DatabaseDNS != "postgres://custom" {
			t.Fatal("DATABASE_DNS not respected")
		}
	})
}

func Test_loadConstraints_BuildsPostgresDNS(t *testing.T) {
	withEnv(t, map[string]string{
		"ENGINE":           "postgres",
		"DB_USER":          "u",
		"DB_USER_PASSWORD": "p",
		"DB_NAME":          "n",
		"DB_HOST":          "h",
		"DB_PORT":          "5432",
	}, func() {
		if err := loadConstraints(); err != nil {
			t.Fatal(err)
		}

		if DatabaseDNS == "" {
			t.Fatal("DATABASE_DNS not built")
		}
	})
}

func Test_loadConstraints_InvalidEngine(t *testing.T) {
	withEnv(t, map[string]string{
		"ENGINE":           "oracle",
		"DB_USER":          "u",
		"DB_USER_PASSWORD": "p",
		"DB_NAME":          "n",
		"DB_HOST":          "h",
		"DB_PORT":          "1",
	}, func() {
		if err := loadConstraints(); err == nil {
			t.Fatal("expected error for unsupported engine")
		}
	})
}

/* ----------------------------------------------------
   Load (full flow)
---------------------------------------------------- */

func Test_Load_FullFlow(t *testing.T) {
	withEnv(t, map[string]string{
		"ENGINE":           "postgres",
		"DB_USER":          "u",
		"DB_USER_PASSWORD": "p",
		"DB_NAME":          "n",
		"DB_HOST":          "h",
		"DB_PORT":          "5432",
	}, func() {

		if err := Load(true); err == nil {
			t.Fatal("Load failed: it should force interrupt with `true` arggument")
		}

		if err := Load(false); err != nil {
			t.Fatalf("Load failed: %v", err)
		}

		if gin.Mode() != gin.DebugMode {
			t.Fatal("gin mode not applied")
		}

		if DatabaseDNS == "" {
			t.Fatal("DATABASE_DNS not initialized")
		}
	})
}

/* ----------------------------------------------------
   RegisterSwagger
---------------------------------------------------- */

func TestRegisterSwagger(t *testing.T) {
	t.Run("nil engine", func(t *testing.T) {
		RegisterSwagger(nil)
	})

	t.Run("debug mode", func(t *testing.T) {
		gin.SetMode(gin.DebugMode)
		r := gin.New()
		RegisterSwagger(r)
	})

	t.Run("release mode", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()
		RegisterSwagger(r)
	})
}

/* ----------------------------------------------------
   buildDatabaseDNS — missing env vars
---------------------------------------------------- */

func Test_buildDatabaseDNS_MissingEachRequiredVar(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
	}{
		{
			name: "missing DB_USER",
			env: map[string]string{
				"ENGINE": "postgres",
			},
		},
		{
			name: "missing DB_USER_PASSWORD",
			env: map[string]string{
				"ENGINE":  "postgres",
				"DB_USER": "u",
			},
		},
		{
			name: "missing DB_NAME",
			env: map[string]string{
				"ENGINE":           "postgre",
				"DB_USER":          "u",
				"DB_USER_PASSWORD": "p",
			},
		},
		{
			name: "missing DB_HOST",
			env: map[string]string{
				"ENGINE":           "postgresql",
				"DB_USER":          "u",
				"DB_USER_PASSWORD": "p",
				"DB_NAME":          "n",
			},
		},
		{
			name: "missing DB_PORT",
			env: map[string]string{
				"ENGINE":           "postgres",
				"DB_USER":          "u",
				"DB_USER_PASSWORD": "p",
				"DB_NAME":          "n",
				"DB_HOST":          "h",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withEnv(t, tt.env, func() {
				Engine = "postgres"
				DatabaseDNS = ""

				if err := buildDatabaseDNS(); err == nil {
					t.Fatal("expected error, got nil")
				}
			})
		})
	}
}

/* ----------------------------------------------------
   buildDatabaseDNS — mysql branch
---------------------------------------------------- */

func Test_buildDatabaseDNS_MySQL(t *testing.T) {
	withEnv(t, map[string]string{
		"ENGINE":           "mysql",
		"DB_USER":          "u",
		"DB_USER_PASSWORD": "p",
		"DB_NAME":          "n",
		"DB_HOST":          "h",
		"DB_PORT":          "3306",
	}, func() {
		Engine = "mysql"
		DatabaseDNS = ""

		if err := buildDatabaseDNS(); err != nil {
			t.Fatal(err)
		}

		if DatabaseDNS == "" {
			t.Fatal("mysql DATABASE_DNS not built")
		}
	})
}

/* ----------------------------------------------------
   loadConstraints — ENGINE missing
---------------------------------------------------- */

func Test_loadConstraints_MissingEngine(t *testing.T) {
	withEnv(t, map[string]string{
		"DATABASE_URL": "postgres://x",
	}, func() {
		Engine = ""
		DatabaseDNS = ""

		if err := loadConstraints(); err == nil {
			t.Fatal("expected error when ENGINE is missing")
		}
	})
}

/* ----------------------------------------------------
   checkNoOptinalAPIKeys — ENGINE required
---------------------------------------------------- */

func Test_checkNoOptinalAPIKeys_MissingEngine(t *testing.T) {
	withEnv(t, map[string]string{
		"DATABASE_URL": "postgres://x",
	}, func() {
		if err := checkNoOptinalAPIKeys(); err == nil {
			t.Fatal("expected error for missing ENGINE")
		}
	})
}

/* ----------------------------------------------------
   Load — failure from checkNoOptinalAPIKeys
---------------------------------------------------- */

func Test_Load_FailsOnInvalidRequiredEnv(t *testing.T) {
	withEnv(t, map[string]string{}, func() {
		if err := Load(false); err == nil {
			t.Fatal("expected Load to fail on missing required env")
		}
	})
}

func Test_DummyOperations(t *testing.T) {
	os.Setenv("THIS_KEY_DOES_NOT_EXISTS", "")
	if err := checkNoOptionalKey("THIS_KEY_DOES_NOT_EXISTS", false); err == nil {
		t.Fatal("checkNoOptionalKey should returns an error, but returns nil")
	}
}
