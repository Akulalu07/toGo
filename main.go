package main

import (
	"toGo/cmd"
	"toGo/db"
)

func main() {
	cmd.Init()
	db.Dinit()
	cmd.Execute()
}
