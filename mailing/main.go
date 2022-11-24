package mailing

import (
	"gopkg.in/gomail.v2"
)

func SendMail(mailType string, receiver ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "hnbn707@naver.com")
	m.SetHeader("To", receiver...)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "메일전송")
	m.SetBody("text/html", "테스트 중입니다")
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.naver.com", 587, "hnbn707", "roalroal123")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}	
	return nil
}