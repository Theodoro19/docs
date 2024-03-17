package example

import (
	"strings"
	"testing"
)

func TestGetTelegramNotifier(t *testing.T)  {
	notification := Notification{"5552233", "Hello world!"}
	notifier, err := GetNotifier(Telegram)

	if err != nil {
		t.Fatal("Notifier type 'Telegram' must be returned")
	}

	resultStr := notifier.SendNotification(notification)
	if !strings.Contains(resultStr, "was sent to number") {
		t.Error("Telegram notifier result message was not correct.")
	}
}

func TestGetWhatsappNotifier(t *testing.T)  {
	notification := Notification{"5552233", "Hello world!"}
	notifier, err := GetNotifier(Whatsapp)

	if err != nil {
		t.Fatal("Notifier type 'Email' must be returned")
	}

	resultStr := notifier.SendNotification(notification)
	if !strings.Contains(resultStr, "was sent to number") {
		t.Error("Whatsapp notifier result message was not correct.")
	}
}

func TestGetNonExistanceNotifier(t *testing.T)  {
	_, err := GetNotifier(3)

	if err == nil {
		t.Fatal("Error must be returned for unrecognized notifier type.")
	}
}