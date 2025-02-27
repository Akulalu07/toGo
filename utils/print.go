package utils

import (
	"os"

	"github.com/fatih/color"
)

func Fatal(a ...interface{}) {
	color.RGB(230, 42, 42).Println(a)
	os.Exit(1)
}

func Good(some string) {
	color.Green(some)
}
