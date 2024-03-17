package example

import "fmt"

func main() {
	smsNotification := &SMS
	emailNotification := &Email

	xiaomiCellphone := &Xiaomi
	xiaomiCellphone.SetSendNotification(smsNotification)
	xiaomiCellphone.MobileNotification()
	fmt.Println()

	xiaomiCellphone.SetSendNotification(emailNotification)
	xiaomiCellphone.MobileNotification()
	fmt.Println()

	motorolaCellphone := &Motorola
	motorolaCellphone.SetSendNotification(smsNotification)
	motorolaCellphone.MobileNotification()
	fmt.Println()

	motorolaCellphone.SetSendNotification(emailNotification)
	motorolaCellphone.MobileNotification()
	fmt.Println()
}