package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the CLI app ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Command line application"
	app.Usage = "Find IPs and server names over the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IPs on internet",
			Flags:  flags,
			Action: findIps,
		}
	}

	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")
	fmt.Println("Searching ip for host ", host)

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
