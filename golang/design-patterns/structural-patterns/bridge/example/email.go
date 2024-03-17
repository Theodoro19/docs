package example

import "fmt"

type Email struct {}

func (e *Email) SendNotification() {
	fmt.Println("Notification sent by e-mail")
}