package main

import (
	"github.com/urfave/cli/v2"
)

func main(){
	app := &cli.App{
		Name: "healtherchecker",
		Usage: "tool to check website status",
		Flags:[]cli.Flag{
			&cli.Flag{
				Name: "domain",
				Aleases:{}string{'d'},
				Usage: "domain name",
				Required: true,
			}
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check.",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}