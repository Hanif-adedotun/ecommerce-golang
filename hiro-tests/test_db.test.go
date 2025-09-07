package db

import (
	"database/sql"
	"errors"
	"os"
	"testing"
)

func TestConnectDB(t *testing.T) {
	// Test case 1: Successful database connection
	t.Run("Successful database connection", func(t *testing.T) {
		envVars := map[string]string{
			"DB_USER": "test_user",
			"DB_PASS": "test_pass",
			"INSTANCE_HOST": "localhost",
			"DB_PORT": "5432",
			"DB_NAME": "test_database",
		}
		for k, v := range envVars {
			os.Setenv(k, v)
		}
		defer func() {
			for k := range envVars {
				os.Unsetenv(k)
			}
		}()
		_, err := ConnectDB()
		if err != nil {
			t.Errorf("ConnectDB() error = %v", err)
		}
	})

	// Test case 2: Missing environment variables
	t.Run("Missing environment variables", func(t *testing.T) {
		os.Unsetenv("DB_USER")
		defer os.Setenv("DB_USER", "test_user")
		_, err := ConnectDB()
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("ConnectDB() error = %v, want %v", err, os.ErrNotExist)
		}
	})

	// Test case 3: Invalid database URI
	t.Run("Invalid database URI", func(t *testing.T) {
		envVars := map[string]string{
			"DB_USER": "test_user",
			"DB_PASS": "test_pass",
			"INSTANCE_HOST": "invalid_host",
			"DB_PORT": "5432",
			"DB_NAME": "test_database",
		}
		for k, v := range envVars {
			os.Setenv(k, v)
		}
		defer func() {
			for k := range envVars {
				os.Unsetenv(k)
			}
		}()
		_, err := ConnectDB()
		if err == nil {
			t.Errorf("ConnectDB() error = %v, want non-nil error", err)
		}
	})
}