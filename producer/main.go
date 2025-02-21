package main

import (
	"context"
	"strconv"

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
	for i := 1; i <= 100; i++ {
		var value string
		if i%3 == 0 && i%5 == 0 {
			value = "fizzbuzz"
		} else if i%3 == 0 {
			value = "fizz"
		} else if i%5 == 0 {
			value = "buzz"
		} else {
			value = strconv.Itoa(i)
		}
		message := amqp091.Publishing{
			Body: []byte(value),
		}
		err := ch.PublishWithContext(ctx, "notification", "email", false, false, message)
		if err != nil {
			panic(err)
		}
	}
}
