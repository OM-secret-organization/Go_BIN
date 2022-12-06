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

	q2, err := ch.QueueDeclare(
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

	q3, err := ch.QueueDeclare(
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

	fmt.Printf("queue : %+v/n", q)
	fmt.Printf("queue : %+v/n", q2)
	fmt.Printf("queue : %+v/n", q3)


	// err = ch.Publish(
	// 	"",
	// 	"withdrawTxHashQueue",
	// 	false,
	// 	false,
	// 	amqp.Publishing{
	// 		ContentType: "application/json",
	// 		Body: []byte(`[
  //   {
	// 		"type" : "BTC",
	// 		"txHash" : ["aaaaa1", "bbbbb"]
	// 	},
	// 	{
	// 		"type" : "ETH",
	// 		"txHash" : ["ccccc2", "ddddd"]
	// 	},
	// 	{
	// 		"type" : "XRP",
	// 		"txHash" : ["ccccc3", "ddddd"]
	// 	}
	// ]`),
	// 	},
	// )

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("W : Success💙")


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
			"address" : ["9e122bd40b806085a598d897d6fa4e3ff208caee0cc078ee48e3b7c55bdbac6d", "54f40825cf4becbd95cbd7c5dfaf7390613765997fccaea29f416a3c9fe70cbb"]
		},
		{
			"type" : "ETH",
			"address" : ["0xa6c24f17dc3b168acd8e119ded5f17860ec0e5fe2e6c33b56b875b71787659d1", "0x7c5135f7c6c3817deb8cab89f745188e738d07e1376fc58ef1d1abb27b5e1f5c"]
		},
		{
			"type" : "XRP",
			"address" : ["B7ADDEC40BA05CA23DEF50925EA565FAEB267A6DFF221BE65C9C60F52DBEB693", "82D8546F458C482597F3E4661BEF4F62B554713DDD765F532763754FC298923E"]
		}
	]`),
		},
	)


	if err != nil {
		fmt.Println(err)
		panic(err)
	}


	// err = ch.Publish(
	// 	"",
	// 	"withdrawTxHashQueue",
	// 	false,
	// 	false,
	// 	amqp.Publishing{
	// 		ContentType: "application/json",
	// 		Body: []byte(`{
	// 		"type" : "BTC",
	// 		"txHash" : "54f40825cf4becbd95cbd7c5dfaf7390613765997fccaea29f416a3c9fe70cbc"
	// 	}`),
	// 	},
	// )

	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

}