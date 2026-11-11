package main

import (
	"errors"
)

type Sender interface {
	Send(recipient, text string) error
}

type NotificationService struct {
	sender Sender
}

func NewNotificationService(sender Sender) (*NotificationService, error) {
	if sender == nil {
		return nil, errors.New("sender is nil")
	}

	service := &NotificationService{
		sender: sender,
	}
	return service, nil
}

func (s *NotificationService) Notify(recipient, text string) error {
	if recipient == "" {
		return errors.New("recipient is empty")
	}
	if text == "" {
		return errors.New("text is empty")
	}

	return s.sender.Send(recipient, text)
}
