package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"retractTxHashQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("W : Success💙", q)


	err = ch.Publish(
		"",
		"retractTxHashQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(`{
			"type" : "XRP",
			"txHash" : "30C4A034F43243683B22D83AFFA6F9BC6BBCC44A9AD3B4C879FAC0CF2F6430F1"
		}`),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("retract TxHash Published")

}