package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")

	conn, err := amqp.Dial("amqp://guest:guest@hanbin.shop:5672/")
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
			"address" : ["0xeBec795c9c8bBD61FFc14A6662944748F299cAcf"]
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