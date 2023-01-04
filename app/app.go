package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "Command Line Application"
	app.Usage = "Find IPs and server names in internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find address IP in Internet",
			Flags:  flags,
			Action: findIPs,
		},
		{
			Name:   "servers",
			Usage:  "Find servers name in Internet",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findIPs(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")
	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
