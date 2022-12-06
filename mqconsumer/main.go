package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func main () {
	fmt.Println("Consumer run")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // 여기 url 설정을 맞춰주면 통신 가능!
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
		"retractTxDataQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)


	forever := make(chan bool)
	go func() {
		fmt.Println("?????")
		d := <- msgs 
			var tx Transaction
			err := json.Unmarshal(d.Body, &tx)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			fmt.Printf("트랜잭션이 들어왔다 : %+v\n", tx)

		// 	for _, hash := range list {
		// 		switch hash.Type {
		// 		case "BTC" :
		// 			fmt.Println("BTC hash list came")
				
		// 		case "ETH" :
		// 			fmt.Println("ETH hash list came")

		// 		case "XRP" : 
		// 			fmt.Println("XRP hash list came")
		// 	}
		// }
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
	Quantity interface{} `json:"quantity"`
	TxHash string `json:"txHash"`
}