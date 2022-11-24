package main

import "study/learngo/httptest"

func main() {
	// levelDB.RunLevelDb()
	// rpc.Run()
	// learnMysql.RunSql()
	// receivers := []string{"hnbn707@gmail.com", "hnbn9414@gmail.com", "gksqls29@naver.com"}
	// err := mailing.SendMail("test", receivers...)
	// if err != nil {
	// 	fmt.Println("메일전송실패")
	// }
	// fmt.Println("메일전송성공")

	httptest.Get("http://localhost:3000/test")
}