// cmd/my-cli/cli.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var host string
	app := &cli.App{
		Name:  "Website lookup CLI",
		Usage: "Let's you query IPs, CNAMEs, MX records and Name Servers!",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "host",

				Value:       "http://youtube.com",
				Destination: &host,
			},
		},
		Commands: []cli.Command{
			{
				Name:  "ns",
				Usage: "will retrieve the name servers",
				Action: func(c *cli.Context) error {
					ns, err := net.LookupNS(host)
					if err != nil {
						return err
					}
					for _, host := range ns {
						fmt.Println(host.Host)
					}
					return nil
				},
			},
			{
				Name:  "cname",
				Usage: "will lookup the CNAME for a given host",
				Action: func(c *cli.Context) error {
					cname, err := net.LookupCNAME(host)
					if err != nil {
						return err
					}

					fmt.Println(cname)
					return nil
				},
			},
			{
				Name:  "mx",
				Usage: "will lookup the mail exchange records for a given host",
				Action: func(c *cli.Context) error {
					mxs, err := net.LookupMX(host)
					if err != nil {
						return err
					}

					for _, mx := range mxs {
						fmt.Println(mx.Host)
					}
					return nil
				},
			},
			{
				Name:  "ip",
				Usage: "will lookup the ip address for a given host",
				Action: func(c *cli.Context) error {
					addrs, err := net.LookupIP(host)
					if err != nil {
						return err
					}

					for _, addr := range addrs {
						fmt.Println(addr)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err.Error())
	}
}
