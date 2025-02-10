package main

import (
	"toGo/cmd"
	"toGo/db"
)

func main() {
	db.Dinit()
	cmd.Execute()
}
