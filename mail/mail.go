package mail

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	gomail "gopkg.in/mail.v2"
)

func SendMail(mailData string) {

	m := gomail.NewMessage()

	m.SetHeader("From", "your@example.com")

	m.SetHeader("To", "your@gmail.com",)

	m.SetHeader("Subject", "Yeni Order")

	m.SetBody("text/plain", mailData)

	d := gomail.NewDialer("mail.exampleSmptServer.com", 465, "your@example.com", "password")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func SendVerifyCode(dataUser string) string {
	token := CreateRandomKey(6)
	m := gomail.NewMessage()

	m.SetHeader("From", "your@example.com")

	m.SetHeader("To", dataUser)

	m.SetHeader("Subject", "Your Code")

	m.SetBody("text/plain", "Your verification code: : "+token)

	d := gomail.NewDialer("mail.exampleSmptServer.com", 465, "your@example.com", "password")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return token
}

func SendForgetCode(dataUser string) string {
	token := CreateRandomKey(8)
	m := gomail.NewMessage()

	m.SetHeader("From", "your@example.com")

	m.SetHeader("To", dataUser)

	m.SetHeader("Subject", "Your new password link")

	m.SetBody("text/plain", "Your new link : "+"http://example.com/?token="+token)

	d := gomail.NewDialer("mail.exampleSmptServer.com", 465, "your@example.com", "password")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return token
}

func CreateRandomKey(size int) string {
	s := ""
	for i := 1; i <= size; i++ {
		s += GetSingleRandomCharacter()
	}
	return s
}

func GetSingleRandomCharacter() string {
	rand.Seed(time.Now().UnixNano())
	alphabet := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	k := string(alphabet[rand.Intn(len(alphabet))])
	return k
}
