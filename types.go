package telegrambot

type Bot struct {
	Token    string
	Username string
	Handlers map[string]func(*Bot, *Update, []string)
}

type Update struct {
	UpdateID          int64          `json:"update_id"`
	Message           Message        `json:"message,omitempty"`
	EditedMessage     Message        `json:"edited_message,omitempty"`
	ChannelPost       Message        `json:"channel_post,omitempty"`
	EditedChannelPost Message        `json:"edited_channel_post,omitempty"`
	CallbackQuery     *CallbackQuery `json:"callback_query,omitempty"`
}

type User struct {
	ID                      int64  `json:"id"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

type Message struct {
	MessageID      int64    `json:"message_id"`
	Text           string   `json:"text"`
	Chat           Chat     `json:"chat"`
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	From           User     `json:"from"`
	Date           int64    `json:"date"`
}

type Chat struct {
	Id int64 `json:"id"`
}

type Response struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type GetMeResponse struct {
	User User `json:"result"`
}

type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

type AnswerCallbackQueryPayload struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}
