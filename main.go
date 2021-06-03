package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arpesam/go-cmd-line/app"
)

func main() {
	fmt.Println("Starting execution...")

	app := app.Generate()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// if err := app.Run(os.Args); err != nil {
	// 	log.Fatal(err)
	// }
}
