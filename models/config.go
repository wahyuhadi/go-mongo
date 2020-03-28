package models

import (
	"os"
	"framework/db"
)



// Mongo server ip -> localhost -> 127.0.0.1 -> 0.0.0.0
var server = os.Getenv("DATABASE_HOST")

// Database name
var databaseName = os.Getenv("DATABASE_NAME")

// Create a connection
var dbConnect = db.NewConnection(server, databaseName)
