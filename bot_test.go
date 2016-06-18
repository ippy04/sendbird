package sendbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBotCreate(t *testing.T) {
	setup()
	defer teardown()

	botRequest := BotRequest{
		BotUserId:      "helper_bot",
		BotNickname:    "Helper Bot",
		BotCallbackUrl: "https://yourdomain.com/sendbird_bot",
		IsPrivacyMode:  true,
	}

	mux.HandleFunc("/v2/bots", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := BotRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForV2ApiTokenParam(t, r, body)

		if !reflect.DeepEqual(body, botRequest) {
			t.Errorf("Bot.Create API call received %+v, expected %+v", body, botRequest)
		}

		response := `
		{
		    "bot_token": "<BOT_TOKEN>", 
		    "bot_userid": "helper_bot", 
		    "bot_nickname": "Helper Bot", 
		    "bot_callback_url": "https://yourdomain.com/sendbird_bot",
		    "is_privacy_mode": true
		}`
		fmt.Fprint(w, response)
	})

	expected := &Bot{
		BotToken:       "<BOT_TOKEN>",
		BotUserId:      "helper_bot",
		BotNickname:    "Helper Bot",
		BotCallbackUrl: "https://yourdomain.com/sendbird_bot",
		IsPrivacyMode:  true,
	}

	bot, _, err := client.Bot.Create(&botRequest)
	if err != nil {
		t.Errorf("Bot.Create returned error: %v", err)
	}

	if !reflect.DeepEqual(bot, expected) {
		t.Errorf("Bot.Create returned %+v, expected %+v", bot, expected)
	}
}

func TestBotSendMessage(t *testing.T) {
	setup()
	defer teardown()

	botMsgRequest := BotMessageRequest{
		Message:    "Hello John, here's an answer for your question",
		Data:       "extra data",
		ChannelUrl: "<CHANNER_URL>",
	}

	mux.HandleFunc("/v2/bots/123/send", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := BotMessageRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForV2ApiTokenParam(t, r, body)

		if !reflect.DeepEqual(body, botMsgRequest) {
			t.Errorf("Bot.SendMessage API call received %+v, expected %+v", body, botMsgRequest)
		}

		response := `
		{
		    "bot_userid": "hello_bot",
		    "channel_url": "<CHANNEL_URL>",
		    "message" : "Hello John, here's an answer for your question",
		    "data" : "extra data"
		}`
		fmt.Fprint(w, response)
	})

	expected := &BotMessage{

		BotUserId:  "hello_bot",
		ChannelUrl: "<CHANNEL_URL>",
		Message:    "Hello John, here's an answer for your question",
		Data:       "extra data",
	}

	msg, _, err := client.Bot.SendMessage("123", &botMsgRequest)
	if err != nil {
		t.Errorf("Bot.SendMessage returned error: %v", err)
	}

	if !reflect.DeepEqual(msg, expected) {
		t.Errorf("Bot.SendMessage returned %+v, expected %+v", msg, expected)
	}
}

func TestBotList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/bots", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		CheckForAuthContentType(t, r)
		CheckForV2ApiTokenQueryString(t, r)

		response := `
		[
		    {
		        "bot_token": "<BOT_TOKEN1>", 
		        "bot_userid": "helper_bot", 
		        "bot_nickname": "Helper Bot", 
		        "bot_callback_url": "https://yourdomain.com/sendbird_bot",
		        "is_privacy_mode": true
		    },
		    {
		        "bot_token": "<BOT_TOKEN2>", 
		        "bot_userid": "helper_bot2", 
		        "bot_nickname": "Helper Bot2", 
		        "bot_callback_url": "https://yourdomain.com/sendbird_bot",
		        "is_privacy_mode": true
		    }
		]`
		fmt.Fprint(w, response)
	})

	expected := []Bot{
		Bot{
			BotToken:       "<BOT_TOKEN1>",
			BotUserId:      "helper_bot",
			BotNickname:    "Helper Bot",
			BotCallbackUrl: "https://yourdomain.com/sendbird_bot",
			IsPrivacyMode:  true,
		},
		Bot{
			BotToken:       "<BOT_TOKEN2>",
			BotUserId:      "helper_bot2",
			BotNickname:    "Helper Bot2",
			BotCallbackUrl: "https://yourdomain.com/sendbird_bot",
			IsPrivacyMode:  true,
		},
	}

	bots, _, err := client.Bot.List()
	if err != nil {
		t.Errorf("Bot.List returned error: %v", err)
	}

	if !reflect.DeepEqual(bots, expected) {
		t.Errorf("Bot.List returned %+v, expected %+v", bots, expected)
	}
}

func TestBotGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/bots/helper_bot", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		CheckForAuthContentType(t, r)
		CheckForV2ApiTokenQueryString(t, r)

		response := `
		{
		    "bot_token": "<BOT_TOKEN1>", 
		    "bot_userid": "helper_bot", 
		    "bot_nickname": "Helper Bot", 
		    "bot_callback_url": "https://yourdomain.com/sendbird_bot",
		    "is_privacy_mode": true
		}`
		fmt.Fprint(w, response)
	})

	expected := &Bot{
		BotToken:       "<BOT_TOKEN1>",
		BotUserId:      "helper_bot",
		BotNickname:    "Helper Bot",
		BotCallbackUrl: "https://yourdomain.com/sendbird_bot",
		IsPrivacyMode:  true,
	}

	bot, _, err := client.Bot.Get("helper_bot")
	if err != nil {
		t.Errorf("Bot.Get returned error: %v", err)
	}

	if !reflect.DeepEqual(bot, expected) {
		t.Errorf("Bot.Get returned %+v, expected %+v", bot, expected)
	}
}

func TestBotUpdate(t *testing.T) {
	setup()
	defer teardown()

	botRequest := BotUpdateRequest{
		BotNickname:    "Helper Bot",
		BotCallbackUrl: "https://yourdomain.com/sendbird_bot",
		IsPrivacyMode:  true,
	}

	mux.HandleFunc("/v2/bots/helper_bot", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := BotUpdateRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForV2ApiTokenParam(t, r, body)

		if !reflect.DeepEqual(body, botRequest) {
			t.Errorf("Bot.Update API call received %+v, expected %+v", body, botRequest)
		}

		response := `
		{
		    "bot_token": "<BOT_TOKEN>", 
		    "bot_userid": "helper_bot", 
		    "bot_nickname": "Helper Bot", 
		    "bot_callback_url": "https://yourdomain.com/sendbird_bot",
		    "is_privacy_mode": true
		}`
		fmt.Fprint(w, response)
	})

	expected := &Bot{
		BotToken:       "<BOT_TOKEN>",
		BotUserId:      "helper_bot",
		BotNickname:    "Helper Bot",
		BotCallbackUrl: "https://yourdomain.com/sendbird_bot",
		IsPrivacyMode:  true,
	}

	bot, _, err := client.Bot.Update("helper_bot", &botRequest)
	if err != nil {
		t.Errorf("Bot.Update returned error: %v", err)
	}

	if !reflect.DeepEqual(bot, expected) {
		t.Errorf("Bot.Update returned %+v, expected %+v", bot, expected)
	}
}

func TestBotDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/bots/helper_bot", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		CheckForAuthContentType(t, r)

		body := BotUpdateRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForV2ApiTokenParam(t, r, body)

		response := `
		{
			"bot_userid": "helper_bot"
		}`
		fmt.Fprint(w, response)
	})

	expected := &BotUserId{
		BotUserId: "helper_bot",
	}

	bot, _, err := client.Bot.Delete("helper_bot")
	if err != nil {
		t.Errorf("Bot.Delete returned error: %v", err)
	}

	if !reflect.DeepEqual(bot, expected) {
		t.Errorf("Bot.Delete returned %+v, expected %+v", bot, expected)
	}
}
