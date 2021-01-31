package telegrambot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func NewBot(token string) (*Bot, error) {
	var u User
	b := Bot{
		Token:    token,
		Handlers: make(map[string]func(*Bot, *Update, []string)),
	}

	r, err := b.GetMe()

	if err != nil {
		return nil, error
	}

	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&u)

	if err != nil {

		return nil, err
	}

	b.Username = u.Username

	return &b, err
}

func callAPI(bot *Bot, method string, body io.Reader) (resp *http.Response, err error) {
	URL_PATTERN := fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.Token, method)
	return http.Post(URL_PATTERN, "application/json", body)
}

func (bot *Bot) SendMessage(body io.Reader) {
	resp, err := callAPI(bot, "sendMessage", body)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func GetMe(bot *Bot) (resp *http.Response, err error) {
	URL_PATTERN := fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.Token, "getMe")
	return http.Get(URL_PATTERN)
}

func (bot *Bot) AddHandler(s string, f func(*Bot, *Update, []string)) {
	bot.Handlers[s] = f
	ss := fmt.Sprintf("%s%s", s, bot.Username)
	bot.Handlers[ss] = f
}

func (bot *Bot) HandleUpdate(u *Update) {
	s := u.Message.Text

	trimmed := strings.Trim(s, " ")
	tokens := strings.Split(trimmed, " ")
	funcName := tokens[0]
	args := tokens[1:]

	if f, ok := bot.Handlers[funcName]; ok {
		f(bot, u, args)
	}

}
