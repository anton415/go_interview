package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSender struct {
	calls int
}

func (s *testSender) Send(recipient, text string) error {
	s.calls++
	return nil
}

func TestNotify(t *testing.T) {
	tests := []struct {
		name          string
		recipient     string
		text          string
		expected      error
		expectedCalls int
	}{
		{"text is empty", "Anton", "", errors.New("text is empty"), 0},
		{"recipient is empty", "", "hello", errors.New("recipient is empty"), 0},
		{"notification", "Anton", "hello", nil, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sender := &testSender{}
			service, err := NewNotificationService(sender)
			if err != nil {
				t.Fatal(err)
			}

			result := service.Notify(tt.recipient, tt.text)

			if tt.expected == nil {
				assert.NoError(t, result)
			} else {
				assert.EqualError(t, result, tt.expected.Error())
			}
			assert.Equal(t, tt.expectedCalls, sender.calls)
		})
	}
}
