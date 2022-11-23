# Go에서 MySQL 사용하기

## 1. MySQL 설치

```
$ go get github.com/go-sql-driver/mysql
```

## 2. MySQL 임포트

database/sql과 MySQL 드라이버 패키지를 임포트 해준다.
MySQL 드라이버 패키지를 임포트 할 때에는 \_로 alias를 주어 개발자가 직접 사용할 일이 없도록 한다. (database/sql 패키지가 내부적으로 사용!)

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)
```

## 3. db 연결하기

```go
func main() {
    db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb") // Sql.Open(<드라이버명>, <커넥션포트>), testdb는 데이터베이스명
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    //...(db 사용)....
}
```

Sql.Open을 이용하여 DB와 연결을 하는 코드를 썼지만 실제로 db와의 연결은 현 시점에서는 일어나지 않고, (connection 체크조차 하지 않음)
실제 DB connection은 Query 등의 메소드로 실제 DB연결이 필요한 시점에 이루어지게 된다.

## 4. 쿼리 메소드

1. QueryRow() => 하나의 row만 리턴하는 경우
2. Query() => 복수의 row를 리턴하는 경우
3. Scan() => 하나의 row에서 값을 읽어와 로컬 변수에 할당하는 경우
4. Next() => 복수의 row에서 다음 row로 넘어갈 때 사용

5. QueryRow() 사용 - Scan과 함께 사용

```go
func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "hanbin:hanbin00@tcp(127.0.0.1:3306)/gotestdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 하나의 Row를 갖는 SQL 쿼리
    var name string
    err = db.QueryRow("SELECT name FROM test1 WHERE idx = 1").Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)
}
```

2. Query() 사용 - Next, Scan과 함께 사용한다

```go
func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "hanbin:hanbin00@tcp(127.0.0.1:3306)/gotestdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 하나의 Row를 갖는 SQL 쿼리
    var names []string
    var id int
    var name2 string
    rows, err := db.Query("SELECT idx, name FROM test1 WHERE idx >= ?", 2) // ? 자리에 값을 별도의 파라미터로 넣을 수 있음!
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close() // 반드시 닫아주기! (지연하여 닫기) <= 왜?

    for rows.Next() {
        err := rows.Scan(&id, &name2)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name2)
        names = append(names, name2)
    }

    fmt.Println(names)
}
```

3. INSERT 하기

```go
func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "hanbin:hanbin00@tcp(127.0.0.1:3306)/gotestdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // INSERT
    result, err := db.Exec("INSERT INTO test1 (name) value (?)", "hanbin4")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)

    var name string
    err = db.QueryRow("SELECT name FROM test1 WHERE idx = 4").Scan(&name) // 조회한 결과값을 name 변수에 담음
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("결과 =>", name) // INSERT한 값이 정상적으로 저장되고 출력된 것을 확인 할 수 있음
}
```

## 5. BlockChain Tx DB 저장하기

1. 먼저 테이블 스키마 설정해주기

```
CREATE TABLE transaction (
    idx INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    depositFrom VARCHAR(50) NOT NULL,
    depositTo VARCHAR(50) NOT NULL,
    depositAmount VARCHAR(50) NOT NULL,
    depositDate VARCHAR(50) NOT NULL,
    txHash VARCHAR(70) NOT NULL UNIQUE,
    chainName VARCHAR(3) NOT NULL
);
```

2. INSERT

```go
    result, err := db.Exec("INSERT INTO transaction (depositFrom, depositTo, depositAmount, depositDate, txHash, chainName) value (?, ?, ?, ?, ?, ?)", "0xf57cFb90Dbbc8B51f83f6207EcBa06fbe3f7d528", "0x855697C5e19020E223A48c87391201ACD6057220", "1000000000000", "2022-11-22", "0xc9b0013ac28ba0e34bddf5c6200b7db1a5ec26c976bcc7f99426edd9bde3a6f7", "ETH" )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
```

3. SELECT

```go
    type Txs struct {
    depositFrom string
    depositTo string
    depositAmount string
    depositDate string
    txHash string
    chainName string
    }

    txInfo := new(Txs)

    err = db.QueryRow("SELECT depositFrom, depositTo, depositAmount, depositDate, txHash, chainName FROM transaction WHERE idx = ?", 1).Scan(&txInfo.depositFrom, &txInfo.depositTo, &txInfo.depositAmount, &txInfo.depositDate, &txInfo.txHash, &txInfo.chainName)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("결과 =>",*txInfo)
```
