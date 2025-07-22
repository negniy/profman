package main

import (
	"context"
	"log"
	"os"
	"profman/pkg/cli"
)

func main() {
	app := cli.RegCommands()
	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
