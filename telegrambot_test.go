package telegrambot_test

import (
	"fmt"
	"testing"

	telegrambot "github.com/aikchun/yagotb"
)

func TestNewBot(t *testing.T) {
	_, err := telegrambot.NewBot("abcdefghi")

	if fmt.Sprintf("%v", err) != "This token is invalid." {
		t.Errorf("Suppose to check that the token is invalid.")
	}
}
