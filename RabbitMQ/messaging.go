package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Sending(connectionRabbitMq string) {
	conn, err := amqp.Dial(connectionRabbitMq)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = channel.PublishWithContext(context,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish message")
	log.Printf(" [x] Sent %s\n", body)
}

func Receiving(connectionRabbitMq string) {
	conn, err := amqp.Dial(connectionRabbitMq)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
