package telegrambot

type SendMessagePayload struct {
	ChatId           int64                `json:"chat_id"`
	Text             string               `json:"text"`
	ReplyToMessageID int64                `json:"reply_to_message_id", omitempty`
	ParseMode        string               `json:"parse_mode", omitempty`
	reply_markup     InlineKeyboardMarkup `json:"reply_markup", omitempty`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string `json:"text"`
	Url                          string `json:"url", omitempty`
	CallbackData                 string `json:"callback_data", omitempty`
	SwitchInlineQuery            string `json:"switch_inline_query", omitempty`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat", omitempty`
}
