package main

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Send(message string) error {
	if e.Address == "" {
		return errors.New("adresa net")
	}
	fmt.Println(e.Address, "otpravil", message)
	return nil
}

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Send(message string) error {
	if s.Phone == "" {
		return errors.New("nomera net")
	}
	fmt.Println(s.Phone, "poluchil", message)
	return nil
}

func notify(n Notifier, message string) {
	err := n.Send(message)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
}

func main() {
	email := EmailNotifier{Address: "pochta.com"}
	sms := SMSNotifier{Phone: "+7424242424242"}

	notify(email, "soo na pochtu")

	notify(sms, "soo na sms")

	emptyEmail := EmailNotifier{Address: ""}
	notify(emptyEmail, "eto ne budet vidno")
}
