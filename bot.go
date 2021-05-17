package gotgbot

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: time.Second*30}

func (bot *BotAPI) GetChat(chatId int) (Chat, error) {
	data := map[string]int{"chat_id":chatId}

	marshal, err := json.Marshal(data)
	if err != nil {
		return Chat{}, err
	}

	resp, err := bot.request("getChat", marshal)

	var chat Chat
	if resp.Ok == false {
		return Chat{}, err
	}

	err = json.Unmarshal(resp.Result, &chat)
	if err != nil {
		return Chat{}, err
	}

	return chat, nil
}

func (bot *BotAPI) GetChatMembersCount(chatId int) (int, error) {
	data := map[string]int{"chat_id":chatId}

	marshal, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	resp, err := bot.request("getChatMembersCount", marshal)

	var count int
	err = json.Unmarshal(resp.Result, &count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (bot *BotAPI) request(apiMethodName string, marshal []byte) (APIResponse, error) {
	url := "https://api.telegram.org/bot"+bot.Token+"/"+apiMethodName

	reader := bytes.NewReader(marshal)
	response, err := httpClient.Post(url, "application/json", reader)
	if err != nil {
		return APIResponse{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return APIResponse{}, err
	}

	if bot.Debug {
		log.Println("\n------------\n"+string(body)+"\n------------\n")
	}

	var resp APIResponse
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return APIResponse{}, err
	}

	if !resp.Ok {
		parameters := ResponseParameters{}
		if resp.Parameters != nil {
			parameters = *resp.Parameters
		}

		return resp, Error{Code: resp.ErrorCode, Message: resp.Description, ResponseParameters: parameters}
	}

	return resp, nil
}

// NewBot Make new bot by bot token
func NewBot(token string) BotAPI {
	bot := BotAPI{Token: token}
	return bot
}

// SetDebug Set debug mode
func (bot *BotAPI) SetDebug(debug bool) {
	bot.Debug = true
}
