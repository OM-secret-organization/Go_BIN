package learnMysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func RunSql() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "hanbin:hanbin00@tcp(127.0.0.1:3306)/gotestdb")
    if err != nil {
        log.Fatal(err)
    }
    
    // defer db.Close()

    // DB 추가
    result, err := db.Exec("INSERT INTO transaction (depositFrom, depositTo, depositAmount, depositDate, txHash, chainName) value (?, ?, ?, ?, ?, ?)", "0xf57cFb90Dbbc8B51f83f6207EcBa06fbe3f7d528", "0x855697C5e19020E223A48c87391201ACD6057220", "1000000000000", "2022-11-22", "0xc9b0013ac28ba0e34bddf5c6200b7db1a5ec26c976bcc7f99426edd9bde3a6f7", "ETH" )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)


    type Txs struct {
    depositFrom string
    depositTo string
    depositAmount string
    depositDate string
    txHash string
    chainName string
    }

    txInfo := new(Txs)

    err = db.QueryRow("SELECT depositFrom, depositTo, depositAmount, depositDate, txHash, chainName FROM transaction WHERE idx = ? OR ", 1).Scan(&txInfo.depositFrom, &txInfo.depositTo, &txInfo.depositAmount, &txInfo.depositDate, &txInfo.txHash, &txInfo.chainName) // 조회한 결과값을 name 변수에 담음
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("결과 =>",*txInfo)
}