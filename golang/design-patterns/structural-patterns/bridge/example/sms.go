package example

import "fmt"

type SMS struct {}

func (s *SMS) SendNotification() {
	fmt.Println("Notification sent by sms")
}