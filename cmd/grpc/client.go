package grpc

import (
	"context"
	"fmt"
	"log"
	"modules/pkg/config"
	greeter2 "modules/tools/generated/v1/greeter"
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
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
	_ctx := greeter2.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := _ctx.SayHello(ctx, &greeter2.MessageRequest{
		Message:   "Hello from service provider",
		Timestamp: time.Now().UTC().Unix(),
	})
	if err != nil {
		return err
	}
	log.Printf("Greeting: [%s] at timestamp - [%v]", r.GetMessage(), r.GetTimestamp())
	return nil
}
