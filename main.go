package main

import (
	"log"
	"modules/v2/cmd"
	"os"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/urfave/cli"
)

var (
	client *cli.App
)

func init() {
	// Initialise a CLI app
	client = cli.NewApp()
	client.Name = "grpc generated"
	client.Usage = "grpc generated client and server"
	client.Version = "0.0.1"
}

func main() {

	client.Commands = []cli.Command{
		{
			Name:  "client",
			Usage: "launch function",
			Action: func(c *cli.Context) error {
				log.Printf("run cli %s", c.App.Name)
				client := cmd.NewClient()
				err := client.SayHelloService()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "server",
			Usage: "start gRPC server",
			Action: func(c *cli.Context) error {
				log.Printf("run cli %s", c.App.Name)
				server := cmd.NewServer()
				return server.GrpcListen()
			},
		},
	}

	// Run the CLI app
	client.Run(os.Args)
}
