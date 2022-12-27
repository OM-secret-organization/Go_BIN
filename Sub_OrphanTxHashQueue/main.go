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
		"validateTxDataQueue",
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
			var hashlist HashList
			err := json.Unmarshal(d.Body, &hashlist)

			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			for _, hash := range hashlist.TxHash {
				fmt.Println("잘못된 해시 : ", hash)
			}
			fmt.Printf("orphan transaction 발견! : %+v\n", hashlist)
	}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - waiting for messages")

	<- forever
}

type HashList struct {
	TxHash []string `json:"txHash"`
}

type Transaction struct {
	To string `json:"to"`
	From string `json:"from"`
	Quantity string `json:"quantity"`
	TxHash string `json:"txHash"`
	Error bool `json:"error"`
}



// 입금 -> 트랜잭션정보
// 해시 -> 출금 -> 트랜잭션정보
// 해시 -> 철회 -> 트랜잭션정보
// ---
// []해시 -> 검증 -> 검증실패한 트랜잭션 해시