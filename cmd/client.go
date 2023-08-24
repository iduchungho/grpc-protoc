package cmd

import (
	"context"
	"log"
	"modules/v2/pkg/generated/v1/greeter"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SayHelloService() error {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
