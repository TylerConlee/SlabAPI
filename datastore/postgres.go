package datastore

import (
	"database/sql"
	"fmt"
	"os"

	// postgres driver
	_ "github.com/lib/pq"
)

// Db is our database struct used for interacting with the database
type Db struct {
	*sql.DB
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

// ConnString returns a connection string based on the parameters it's given
func ConnString() string {
	host := getEnv("SLAB_PG_HOST", "localhost")
	port := getEnv("SLAB_PG_PORT", 5432)
	user := getEnv("SLAB_PG_USER", "postgres")
	pass := getEnv("SLAB_PG_PASS", "")
	dbName := getEnv("SLAB_PG_DBNAME", "slabapi")
	return fmt.Sprintf(
		"host=%s port=%d user=%s pass=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbName,
	)
}

func getEnv(key string, fallback interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
