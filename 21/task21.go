package main

import "fmt"

type poluchatel interface {
	poluchenie()
}

type otprovitel struct{}

func (o otprovitel) otdat() {
	fmt.Println("peredano")
}

type pochta struct {
	otpr otprovitel
}

func (p pochta) poluchenie() {
	p.otpr.otdat()
}

func main() {
	otprovit := otprovitel{}
	pocht := pochta{otpr: otprovit}
	pocht.poluchenie()
}
