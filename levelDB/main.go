package levelDB

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// bytecode => string 변환 함수
func decode (b []byte) string {
	return string(b[:len(b)])
}

func RunLevelDb() {
	// DB creation
	filePath := "levelDB/test_path/test_file"
	db, err := leveldb.OpenFile(filePath, nil)
	if err != nil {
		panic(err)
	}

	// DB Create & Read
	err = db.Put([]byte("Key"), []byte("value"), nil) // key와 value를 byte로 바꿔 db에 저장
	err = db.Put([]byte("han"), []byte("bin"), nil) // key와 value를 byte로 바꿔 db에 저장

	data, err := db.Get([]byte("han"), nil) // byte값으로 db 조회
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DB create : ", decode(data)) // 조회결과인 data는 value의 byte값. decode 함수를 이용해 string으로 변환

	// ----- 여기까지 실행 시 test_path라는 디렉토리 생성 - 내부에 test_file 디렉토리.. 거기에 저장됨 알아볼수없는형태로
	// db.Get으로 Put 하지 않은 Key값 입력시 `leveldb: not found` err 출력됨

	// DB Delete
	err = db.Delete([]byte("han"), nil)
	data, err = db.Get([]byte("han"), nil) // byte값으로 db 조회
	if err != nil {
		fmt.Println("DB delete", err)
	}
	fmt.Println("DB delete", decode(data)) 



	// Batch : 여러 항목 일괄 처리
	// batch를 만든 뒤 필요한 동작들을 적어두고 db.Write로 batch 적용
	batch := new(leveldb.Batch)
	batch.Put([]byte("key1"), []byte("value1"))
	batch.Put([]byte("key2"), []byte("value2"))
	batch.Put([]byte("key3"), []byte("value3"))
	batch.Delete([]byte("key2"))
	err = db.Write(batch, nil)

	



	// Iterator : 반복
	iter := db.NewIterator(nil, nil) // 전체항목 순환조회
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println("iter : ", decode(key), decode(value))
	}
	iter.Release() // 반복해제
	err = iter.Error()
}
