package example

import (
	"errors"
	"fmt"
)

const (
	Telegram   = 1
	Whatsapp = 2
)

type Notification struct {
	To      string
	Message string
}

type Notifier interface {
	SendNotification(Notification) string
}

type TelegramNotification struct{}
func (n TelegramNotification) SendNotification(notification Notification) string {
	return fmt.Sprintf("`%s` was sent to number `%s` successfully!", notification.Message, notification.To)
}

type WhatsappNotification struct{}
func (n WhatsappNotification) SendNotification(notification Notification) string {
	return fmt.Sprintf("`%s` was sent to email `%s` successfully!", notification.Message, notification.To)
}

// Factory .....
func GetNotifier(notifierType int) (Notifier, error) {
	switch notifierType {
		case Telegram:
			return new(TelegramNotification), nil
		case Whatsapp:
			return new(WhatsappNotification), nil
		default:
		return nil, errors.New(fmt.Sprintf("Notifier type %d not recognized.", notifierType))
	}
}