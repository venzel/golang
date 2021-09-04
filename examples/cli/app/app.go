package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "CLI venzel"
	
	app.Usage = "Busca ips e dns"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "venzel.com.br", // default
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de enderecos",
			Flags: flags,
			Action: find_ips,
		},
		{
			Name: "dns",
			Usage: "Busca servidores",
			Flags: flags,
			Action: find_dns,
		},
	}

	return app
}

func find_ips(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func find_dns(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}