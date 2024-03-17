package example

import "fmt"

type Xiaomi struct {
	notification Notification
}

func (x *Xiaomi) MobileNotification(){
	fmt.Println("Send notification from Xiaomi")
	x.notification.SendNotification()
}

func (x *Xiaomi) SetSendNotification(n notification) {
	x.notification = n
}