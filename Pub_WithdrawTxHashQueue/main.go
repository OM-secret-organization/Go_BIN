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
		"withdrawTxHashQueue",
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
		"withdrawTxHashQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(`{
			"type" : "BTC",
			"txHash" : "df09137495b552f430172ceb2cfec9b3a8eab4752b928c380ea0f6c2e00617e2"
		}`),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Withdraw TxHash Published")

}