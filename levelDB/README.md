# Go levelDB 사용하기

1.  goleveldb package 설치

    ```
    $ go get github.com/syndtr/goleveldb/leveldb
    ```

2.  설치한 package import하기

    ```go
    import (
      "fmt"

      "github.com/syndtr/goleveldb/leveldb"
      "github.com/syndtr/goleveldb/leveldb/filter"
      "github.com/syndtr/goleveldb/leveldb/opt"
    )
    ```

3.  DB 조작에 필요한 함수 생성

    ```go
    // bytecode => string 변환 함수
    func decode (b []byte) string {
      return string(b[:len(b)])
    }
    ```

4.  levelDB 파일 생성

    ```go
    func main() {
      db, err := leveldb.OpenFile(filePath, nil)
      if err != nil {
        panic(err)
      }
      err = db.Put([]byte("Key"), []byte("value"), nil)
      data, err := db.Get([]byte("Key"), nil)
      fmt.Println(decode(data))
    }
    ```

5.  levelDB를 이용한 batch, iterator 처리
    - 코드 참조!

# 정리

한마디로 levelDB는 간단하게 사용하는 memory DB이다.
DB로 사용할 경로를 지정해주면 내부에 데이터가 저장되며
필요하다면 로컬의 텍스트파일 등을 읽어와 값을 넣어줄 수 있다.

DB에 address값을 넣어두고
그때그때 이 address값이 존재하는지만 확인하면 되기때문에
필요한 메소드들 (HasValue, SetValue)만 있으면 이거만 가지고 조작을 하면 된다.

하지만 계속해서 추가되는 address값을 담고 서버가 꺼지더라도 영구히 저장하기 위해서는 다른 방법이 필요할 듯 하다.
