package main

import (
	"fmt"
	"strconv"

	"github.com/amebalabs/nocreep/app/model"

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
	for i := 1; i <= 10; i++ {
		db.AddDevice(model.Device{ID: strconv.Itoa(i), OS: "iOS 11.2 " + strconv.Itoa(i)})
	}
	api.Run(db)
}
