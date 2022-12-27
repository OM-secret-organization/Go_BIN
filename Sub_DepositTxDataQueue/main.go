package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func main () {
	fmt.Println("Consumer run")

	conn, err := amqp.Dial("amqp://guest:guest@hanbin.shop:5672/") // 여기 url 설정을 맞춰주면 통신 가능!
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

	msgs, err := ch.Consume(
		"depositTxDataQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println("channel consume err", err)
	}


	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var tx Transaction
			err := json.Unmarshal(d.Body, &tx)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			fmt.Printf("트랜잭션 : %+v\n", tx)
	}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - waiting for messages")

	<- forever
}

type HashList struct {
	Type string `json:"type"`
	TxHash []string `json:"txHash"`
}

type Transaction struct {
	To string `json:"to"`
	From string `json:"from"`
	Quantity string `json:"quantity"`
	TxHash string `json:"txHash"`
	Error bool `json:"error"`
}