package main

import (
	"fmt"

	"github.com/amebalabs/nocreep/app/api"
	"github.com/amebalabs/nocreep/app/store"
)

var revision = "unknown"

func main() {
	fmt.Printf("nocreep %s\n", revision)
	db, err := store.SetupBoltDB()
	if err != nil {
		fmt.Print(err)
	}
	a := api.API{}
	a.Run(db)
}
