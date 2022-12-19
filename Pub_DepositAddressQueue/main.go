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

	fmt.Println("Successfully Connected to our RabbitMQ INstance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()


	q, err := ch.QueueDeclare(
		"depositAddressQueue",
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


	err = ch.Publish(
		"",
		"depositAddressQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(`[
    {
			"type" : "BTC",
			"address" : ["16WCKxh6J8D8rs9xCSRjoBMAGV3QVgRDk8", "bc1qwqdg6squsna38e46795at95yu9atm8azzmyvckulcc7kytlcckxswvvzej"]
		},
		{
			"type" : "ETH",
			"address" : ["0xF1fD681a123b19c64110E03AB3F00fD16e9ec185", "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"]
		},
		{
			"type" : "XRP",
			"address" : ["rEb8TK3gBgk5auZkwc6sHnwrGVJH8DuaLh:102171667", "rEb8TK3gBgk5auZkwc6sHnwrGVJH8DuaLh:458813248"]
		}
	]`),
		},
	)


	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Deposit Address Published", q)

}