package main

import (
	"apigo/database"
	"fmt"
)

func init() {
	database.InitDatabase("./test.db")
}

func main() {
	fmt.Println("Entry point")
}
