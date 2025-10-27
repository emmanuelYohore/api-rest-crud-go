package main

import (
	"github.com/emmanuelYohore/api-rest-crud-go/database"
	"github.com/emmanuelYohore/api-rest-crud-go/routes"
)

func main() {
	database.ConnectDatabase()
	routes.Routes()
}
