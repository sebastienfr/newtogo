package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var port = 8020

func main() {
	// new app
	app := cli.NewApp()
	app.Name = "mytodolist"
	app.Usage = "mytodolist service launcher"
	// ...
	// command line flags
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Value: port,
			Name:  "port",
			Usage: "Set the listening port of the webserver",
		},
	}

	// main action
	// sub action are possible also
	app.Action = func(c *cli.Context) error {
		// parse parameters
		port = c.Int("port")
		fmt.Printf("port : %d\n", port)
		return nil
	}

	// run the app
	err := app.Run(os.Args)
	if err != nil {
		fmt.Errorf("Run error %q\n", err)
	}
}
