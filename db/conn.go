package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

 const (
	host		= "localhost"
	port		= 5432
	user		= "postgres"
	password 	= "1234"
	dbname  	= "postgres"
 )

 func ConnectDB () (*sql.DB, error) {
	// Set up the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		// return nil, fmt.Errorf("failed to open database: %v", err)
		panic(err)
	}

	// Test the connection
	err = db.Ping()

	if err != nil {
		// return nil, fmt.Errorf("failed to ping database: %v", err)
		panic(err)
	}

	fmt.Println("Connected to " + dbname + " database!")
	
	// Return the database connection
	return db, nil
}