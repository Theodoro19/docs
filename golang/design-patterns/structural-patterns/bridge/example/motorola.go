package example

import "fmt"

type Motorola struct {
	notification Notification
}

func (m *Motorola) MobileNotification(){
	fmt.println("Send notification from Motorola")
	m.notification.SendNotification()
}

func (m *Motorola) SetSendNotification(n notification) {
	m.notification = n
}