package main

import (
	"apigo/database"
	"apigo/users"
	"fmt"
)

func init() {
	database.InitDatabase("./test.db")
}

func main() {
	users.Migrate()
	fmt.Println("Entry point")
}
