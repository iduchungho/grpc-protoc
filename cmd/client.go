package cmd

import (
	"context"
	"fmt"
	"log"
	"modules/v2/pkg/config"
	"modules/v2/pkg/generated/v1/greeter"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	address string
}

func NewClient() *Client {
	hostAddr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	return &Client{
		address: hostAddr,
	}
}

func (c *Client) SayHelloService() error {
	conn, err := grpc.Dial(c.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	_ctx := greeter.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := _ctx.SayHello(ctx, &greeter.MessageRequest{
		Message:   "Hello from service provider",
		Timestamp: time.Now().UTC().Unix(),
	})
	if err != nil {
		return err
	}
	log.Printf("Greeting: [%s] at timestamp - [%v]", r.GetMessage(), r.GetTimestamp())
	return nil
}
