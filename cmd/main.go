package main

import (
	"log"
	"modules/cmd/grpc"
	"os"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/urfave/cli"
)

var (
	client *cli.App
)

func init() {
	// Initialise a CLI internal
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
				client := grpc.NewClient()
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
				server := grpc.NewServer()
				return server.GrpcListen()
			},
		},
	}

	// Run the CLI internal
	client.Run(os.Args)
}
