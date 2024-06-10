package main

import rabbitmq "go-hello-world-rabbitmq/RabbitMQ"

const ConnectionRabbitMq = "amqp://joaoinocen:123456@localhost:5672/"

func main() {
	rabbitmq.Sending(ConnectionRabbitMq)
	rabbitmq.Receiving(ConnectionRabbitMq)
}
