package service

import (
	"context"
	"log"
	"modules/v2/pkg/plugin/v1/greeter"
	"time"
)

type Greeter struct {
	greeter.UnimplementedGreeterServer
}

func (g *Greeter) SayHello(context context.Context, msg *greeter.MessageRequest) (*greeter.MessageReply, error) {
	log.Printf("received message: %s at time: %v\n", msg.GetMessage(), msg.GetTimestamp())
	return &greeter.MessageReply{
		Message:   "Hello " + msg.GetMessage(),
		Timestamp: time.Now().UTC().Unix(),
	}, nil
}
