package sendbird

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BotService is an interface for interfacing with the Bot
// endpoints of the Sendbird API
type BotService interface {
	Create(params *BotRequest) (*Bot, *Response, error)
	SendMessage(botUserId string, params *BotMessageRequest) (*BotMessage, *Response, error)
	List() ([]Bot, *Response, error)
	Get(botUserId string) (*Bot, *Response, error)
	Update(botUserId string, params *BotUpdateRequest) (*Bot, *Response, error)
	Delete(botUserId string) (*BotUserId, *Response, error)
	Handler(rw http.ResponseWriter, req *http.Request)
}

type MessageReceivedHandler func(*BotCallback)

// BotServiceOp handles communication with the Bot related methods of
// the Sendbird API.
type BotServiceOp struct {
	client          *SendbirdClient
	MessageReceived MessageReceivedHandler
}

var _ BotService = &BotServiceOp{}

type BotRequest struct {
	RequestDefaultsAPIV2
	BotUserId      string `json:"bot_userid"`
	BotNickname    string `json:"bot_nickname"`
	BotCallbackUrl string `json:"bot_callback_url"`
	IsPrivacyMode  bool   `json:"is_privacy_mode"`
}
type BotUpdateRequest struct {
	RequestDefaultsAPIV2
	BotNickname    string `json:"bot_nickname"`
	BotCallbackUrl string `json:"bot_callback_url"`
	IsPrivacyMode  bool   `json:"is_privacy_mode"`
}
type BotMessageRequest struct {
	RequestDefaultsAPIV2
	Message    string `json:"message"`
	Data       string `json:"data"`
	ChannelUrl string `json:"channel_url"`
}

type BotUserId struct {
	BotUserId string `json:"bot_userid"`
}
type Bot struct {
	BotToken       string `json:"bot_token"`
	BotUserId      string `json:"bot_userid"`
	BotNickname    string `json:"bot_nickname"`
	BotCallbackUrl string `json:"bot_callback_url"`
	IsPrivacyMode  bool   `json:"is_privacy_mode"`
}
type BotMessage struct {
	BotUserId  string `json:"bot_userid"`
	Message    string `json:"message"`
	Data       string `json:"data"`
	ChannelUrl string `json:"channel_url"`
}

type BotCallback struct {
	BotUserId      string   `json:"bot_userid"`               // Recipient bot's username
	Category       string   `json:"bot_message_notification"` // Type of bot notification. For now, there's only one type available. "bot_message_notification"
	Timestamp      int64    `json:"ts"`                       // UNIX timestamp of this message. You could use this field to sort messages sent to your bot.
	BotToken       string   `json:"bot_token"`                // This is a secret token return to you when you create your bot. You can use the bot_token to verify that a callback is sent from SendBird.
	BotNickname    string   `json:"bot_nickname"`             // Recipient bot's nickname
	SenderUsername string   `json:"sender_username"`          // Sender's username. This is equivalent to id field in Server API.
	SenderNickname string   `json:"sender_nickname"`          // Sender's nickname
	Message        string   `json:"message"`                  // Message sent from a sender
	Data           string   `json:"data"`                     // Additional data field sent from a sender
	Mentioned      []string `json:"mentioned"`                // A list of usernames mentioned in this message
	ChannelType    string   `json:"channel_type"`             // Type of channel. Two types are available. 'messaging' for 1on1 channel, 'group_messaging' for group messaging channel.
	ChannelURl     string   `json:"channel_url"`              // Unique url of each channel
}

// Create a bot
func (s *BotServiceOp) Create(params *BotRequest) (*Bot, *Response, error) {

	path := "/v2/bots"
	params.PopulateApiV2Token(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	bot := new(Bot)
	resp, err := s.client.Do(req, bot)

	if err != nil {
		return nil, resp, err
	}

	return bot, resp, nil
}

// Send Message
func (s *BotServiceOp) SendMessage(botUserId string, params *BotMessageRequest) (*BotMessage, *Response, error) {

	path := fmt.Sprintf("v2/bots/%s/send", botUserId)
	params.PopulateApiV2Token(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	message := new(BotMessage)
	resp, err := s.client.Do(req, message)

	if err != nil {
		return nil, resp, err
	}

	return message, resp, nil
}

// Get a list of bots in your application
func (s *BotServiceOp) List() ([]Bot, *Response, error) {

	path := fmt.Sprintf("v2/bots?api_token=%s", s.client.ApiToken)
	req, err := s.client.NewRequest("GET", path, nil)

	bots := []Bot{}
	resp, err := s.client.Do(req, &bots)

	if err != nil {
		return nil, resp, err
	}

	return bots, resp, nil
}

// Retrieve a bot
func (s *BotServiceOp) Get(botUserId string) (*Bot, *Response, error) {

	path := fmt.Sprintf("v2/bots/%s?api_token=%s", botUserId, s.client.ApiToken)
	req, err := s.client.NewRequest("GET", path, nil)

	bot := new(Bot)
	resp, err := s.client.Do(req, bot)

	if err != nil {
		return nil, resp, err
	}

	return bot, resp, nil
}

// Update a bot
func (s *BotServiceOp) Update(botUserId string, params *BotUpdateRequest) (*Bot, *Response, error) {

	path := fmt.Sprintf("v2/bots/%s", botUserId)
	params.PopulateApiV2Token(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	bot := new(Bot)
	resp, err := s.client.Do(req, bot)

	if err != nil {
		return nil, resp, err
	}

	return bot, resp, nil
}

// Delete a bot
func (s *BotServiceOp) Delete(botUserId string) (*BotUserId, *Response, error) {

	params := RequestDefaultsAPIV2{}
	params.PopulateApiV2Token(s.client)
	path := fmt.Sprintf("v2/bots/%s", botUserId)
	req, err := s.client.NewRequest("DELETE", path, params)

	bot := new(BotUserId)
	resp, err := s.client.Do(req, bot)

	if err != nil {
		return nil, resp, err
	}

	return bot, resp, nil
}

func (s *BotServiceOp) Handler(rw http.ResponseWriter, req *http.Request) {
	read, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatal(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO: Message integrity check here

	callbackData := new(BotCallback)
	err = json.Unmarshal(read, callbackData)
	if err != nil {
		log.Fatal("Couldn't parse sendbird json:" + err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if s.MessageReceived != nil {
		go s.MessageReceived(callbackData)
	}

	rw.WriteHeader(http.StatusOK)
}
