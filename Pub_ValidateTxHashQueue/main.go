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
		"validateTxHashQueue",
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
		"validateTxHashQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(`[
    {
			"type" : "BTC",
			"txHash" : ["334a9c3af912a579c8e6ee0931a0c8df1b84f0d802701cae852ebe1ba4880189", "7c78ed008cd0618577661e6fc72bf00fa20dd3a10bdb0777ac369b7874ea7564"]
		},
		{
			"type" : "ETH",
			"txHash" : ["0xa6c24f17dc3b168acd8e119ded5f17860ec0e5fe2e6c33b56b875b71787659d1", "0x7c5135f7c6c3817deb8cab89f745188e738d07e1376fc58ef1d1abb27b5e1f5c"]
		},
		{
			"type" : "XRP",
			"txHash" : ["9C1C3943F6139275B9E4BEBE53AFA829A45F4CE3A536E0A07E9AE3B636EF5A4B", "C64219689E978D767EA428D8CE24DA60AB4ABD0A190CFFDB23944F842A0F06C0"]
		}
	]`),
		},
	)


	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Retract TxHash Published", q)

}