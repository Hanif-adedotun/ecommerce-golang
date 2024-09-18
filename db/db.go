package db

import (
	"database/sql"
	"fmt"

	"log"
	"os"

	"github.com/joho/godotenv"

	// Note: If connecting using the App Engine Flex Go runtime, use
	// "github.com/jackc/pgx/stdlib" instead, since v5 requires
	// Go modules which are not supported by App Engine Flex.
	_ "github.com/jackc/pgx/v5/stdlib"
)

// connectTCPSocket initializes a TCP connection pool for a Cloud SQL
// instance of Postgres.
func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_tcp.go: %s environment variable not set.", k)
		}
		return v
	}
	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep secrets safe.
	var (
		dbUser    = mustGetenv("DB_USER")       
		dbPwd     = mustGetenv("DB_PASS")       
		dbTCPHost = mustGetenv("INSTANCE_HOST")
		dbPort    = mustGetenv("DB_PORT")     
		dbName    = mustGetenv("DB_NAME")      
	)

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
		dbTCPHost, dbUser, dbPwd, dbPort, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	fmt.Println("Postgres DB Connected successfully")

	return dbPool, nil
}
