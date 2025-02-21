package main

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	emailConsumer, err := ch.ConsumeWithContext(ctx, "email", "consumer-email", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	for message := range emailConsumer {
		println("routing key: " + message.RoutingKey)
		println("body: " + string(message.Body))
	}
}
