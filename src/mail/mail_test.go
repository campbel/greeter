package mail

import (
	"errors"
	"testing"
)

type mockProvider struct {
	send func() error
}

func (m mockProvider) Send(to, from, subject, message string) error {
	return m.send()
}

func TestSendFailed(t *testing.T) {
	called := false
	provider := mockProvider{
		send: func() error {
			called = true
			return errors.New("foo error")
		},
	}

	sender := NewSender(provider)

	err := sender.Send("", "")
	if err == nil {
		t.Error("Expected an error to occur since no providers returned nil errors")
	}

	if !called {
		t.Error("Expected the senders send function to be called")
	}
}

func TestSendSucceededOnFirst(t *testing.T) {
	called1 := false
	provider1 := mockProvider{
		send: func() error {
			called1 = true
			return nil
		},
	}

	called2 := false
	provider2 := mockProvider{
		send: func() error {
			called2 = true
			return errors.New("foo error")
		},
	}

	sender := NewSender(provider1, provider2)

	err := sender.Send("", "")
	if err != nil {
		t.Error("Expected no error to occur since the second provider returned nil error")
	}

	if !called1 {
		t.Error("Expected the first provider to be called")
	}

	if called2 {
		t.Error("Expected the second provider to NOT be called")
	}
}

func TestSendSucceededOnSecond(t *testing.T) {
	called1 := false
	provider1 := mockProvider{
		send: func() error {
			called1 = true
			return errors.New("foo error")
		},
	}

	called2 := false
	provider2 := mockProvider{
		send: func() error {
			called2 = true
			return nil
		},
	}

	sender := NewSender(provider1, provider2)

	err := sender.Send("", "")
	if err != nil {
		t.Error("Expected no error to occur since the second provider returned nil error")
	}

	if !called1 {
		t.Error("Expected the first provider to be called")
	}

	if !called2 {
		t.Error("Expected the second provider to be called")
	}
}
