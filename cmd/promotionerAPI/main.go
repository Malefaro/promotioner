package main

import (
	"flag"
	"fmt"
)

var port = flag.String("port", ":8000", "server port")
var authAddress = flag.String("authAddress", "localhost:8001", "auth server address")
var databaseAddress = flag.String("databaseAddress", "localhost:5432", "database address")
var databaseUsername = flag.String("databaseUsername", "test_user", "database user")
var databaseName = flag.String("databaseName", "promotioner_test_database", "database name")
var databasePassword = flag.String("databasePassword", "test_user_password", "database password")

func main() {
	flag.Parse()
	fmt.Println(*databaseAddress)
	// todo: connection to authService
	// todo: connection to database
	// todo: run API server
}
