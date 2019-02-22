package main

import (
	"fmt"

	"github.com/amebalabs/nocreep/app/api"
)

var revision = "unknown"

func main() {
	fmt.Printf("nocreep %s\n", revision)
	api.Run()
}
