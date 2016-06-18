package sendbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestChatChannelCreate(t *testing.T) {
	setup()
	defer teardown()

	chatRequest := ChatChannelRequest{
		ChannelUrl: "chat_channel_url",
		Name:       "new chat channel",
		CoverUrl:   "https://sendbird/images/cover_url.jpg",
		Data:       "extra data",
	}

	mux.HandleFunc("/channel/create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, chatRequest) {
			t.Errorf("ChatChannel.Create API call received %+v, expected %+v", body, chatRequest)
		}

		response := `
		{
		    "id": 123456, 
		    "name": "new chat channel", 
		    "channel_url": "chat_channel_url",
		    "member_count": 5,
		    "cover_url": "https://sendbird/images/cover_url.jpg",
		    "data": "extra data",
		    "created_at": 1234567890123
		}`
		fmt.Fprint(w, response)
	})

	expected := &ChatChannel{
		Id:          123456,
		Name:        "new chat channel",
		ChannelUrl:  "chat_channel_url",
		MemberCount: 5,
		CoverUrl:    "https://sendbird/images/cover_url.jpg",
		Data:        "extra data",
		CreatedAt:   1234567890123,
	}

	chatChannel, _, err := client.Chat.Create(&chatRequest)
	if err != nil {
		t.Errorf("ChatChannel.Create returned error: %v", err)
	}

	if !reflect.DeepEqual(chatChannel, expected) {
		t.Errorf("ChatChannel.Create returned %+v, expected %+v", chatChannel, expected)
	}

}

func TestChatChannelList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channel/list", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		response := `
		[
		    {
				    "id": 123456, 
				    "name": "new chat channel", 
				    "channel_url": "chat_channel_url",
				    "member_count": 5,
				    "cover_image_url": "https://sendbird/images/cover_url.jpg",
				    "data": "extra data",
				    "created_at": 1234567890123
		    },
		    {
				    "id": 789, 
				    "name": "new chat channel 2", 
				    "channel_url": "chat_channel_url_2",
				    "member_count": 10,
				    "cover_image_url": "https://sendbird/images/cover_url_2.jpg",
				    "data": "extra data 2",
				    "created_at": 3421353123553212
		    }
		]`

		fmt.Fprint(w, response)
	})

	expected := []ChatChannel{
		{
			Id:            123456,
			Name:          "new chat channel",
			ChannelUrl:    "chat_channel_url",
			MemberCount:   5,
			CoverUrl:      "https://sendbird/images/cover_url.jpg",
			CoverImageUrl: "https://sendbird/images/cover_url.jpg",
			Data:          "extra data",
			CreatedAt:     1234567890123,
		},
		{
			Id:            789,
			Name:          "new chat channel 2",
			ChannelUrl:    "chat_channel_url_2",
			MemberCount:   10,
			CoverUrl:      "https://sendbird/images/cover_url_2.jpg",
			CoverImageUrl: "https://sendbird/images/cover_url_2.jpg",
			Data:          "extra data 2",
			CreatedAt:     3421353123553212,
		},
	}

	chatChannel, _, err := client.Chat.List()
	if err != nil {
		t.Errorf("ChatChannel.List returned error: %v", err)
	}

	if !reflect.DeepEqual(chatChannel, expected) {
		t.Errorf("ChatChannel.List returned %+v, expected %+v", chatChannel, expected)
	}

}

func TestChatChannelUpdate(t *testing.T) {
	setup()
	defer teardown()

	chatRequest := ChatChannelUpdateRequest{
		ChannelUrl: "chat_channel_url",
		Name:       "new chat channel",
		CoverUrl:   "https://sendbird/images/cover_url.jpg",
		Data:       "extra data",
		Ops:        []string{"user_id_1", "user_id_2"},
	}

	mux.HandleFunc("/channel/update", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelUpdateRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, chatRequest) {
			t.Errorf("ChatChannel.Update API call received %+v, expected %+v", body, chatRequest)
		}

		response := `
		{
		    "id": 123456, 
		    "name": "new chat channel", 
		    "channel_url": "chat_channel_url",
		    "member_count": 5,
		    "cover_image_url": "https://sendbird/images/cover_url.jpg",
		    "data": "extra data",
		    "created_at": 1234567890123,
		    "ops": ["user_id_1", "user_id_2"]
		}`

		fmt.Fprint(w, response)
	})

	expected := &ChatChannelUpdate{
		ChatChannel: ChatChannel{
			Id:            123456,
			Name:          "new chat channel",
			ChannelUrl:    "chat_channel_url",
			MemberCount:   5,
			CoverUrl:      "https://sendbird/images/cover_url.jpg",
			CoverImageUrl: "https://sendbird/images/cover_url.jpg",
			Data:          "extra data",
			CreatedAt:     1234567890123,
		},
		Ops: []string{"user_id_1", "user_id_2"},
	}

	chatChannel, _, err := client.Chat.Update(&chatRequest)
	if err != nil {
		t.Errorf("ChatChannel.Update returned error: %v", err)
	}

	if !reflect.DeepEqual(chatChannel, expected) {
		t.Errorf("ChatChannel.Update returned %+v, expected %+v", chatChannel, expected)
	}

}

func TestChatChannelDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channel/delete", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedChannelUrl := "chat_channel_url"
		if body.ChannelUrl != expectedChannelUrl {
			t.Errorf("ChatChannel.View API call received ChannelUrl %s, expected %s", body.ChannelUrl, expectedChannelUrl)
		}

		response := "{}"

		fmt.Fprint(w, response)
	})

	_, err := client.Chat.Delete("chat_channel_url")
	if err != nil {
		t.Errorf("ChatChannel.Delete returned error: %v", err)
	}
}

func TestChatChannelView(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channel/view", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedChannelUrl := "chat_channel_url"
		if body.ChannelUrl != expectedChannelUrl {
			t.Errorf("ChatChannel.View API call received ChannelUrl %s, expected %s", body.ChannelUrl, expectedChannelUrl)
		}

		response := `
	    {
	      "channel_url": "chat_channel_url",        
	      "name": "channel name",               
	      "created_at": 12421521552,            
	      "cover_image_url": "http://sendbird.com/cover_image_url.jpg",
	      "data": "custom data",               
	      "members": [
	        {
	          "id": "1",
	          "image": "http://sendbird.com/user_profile1.jpg",
	          "name": "Bugs"
	        }
	      ]
		}`
		fmt.Fprint(w, response)
	})

	chatChannel, _, err := client.Chat.View("chat_channel_url")

	if err != nil {
		t.Errorf("ChatChannel.View returned error: %v", err)
	}

	expected := &ChatChannelView{
		ChatChannel: ChatChannel{
			Name:          "channel name",
			ChannelUrl:    "chat_channel_url",
			CoverUrl:      "http://sendbird.com/cover_image_url.jpg",
			CoverImageUrl: "http://sendbird.com/cover_image_url.jpg",
			Data:          "custom data",
			CreatedAt:     12421521552,
		},
		Members: []Member{
			{
				Id:    "1",
				Image: "http://sendbird.com/user_profile1.jpg",
				Name:  "Bugs",
			},
		},
	}

	if !reflect.DeepEqual(chatChannel, expected) {
		t.Errorf("ChatChannel.View returned %+v, expected %+v", chatChannel, expected)
	}
}

func TestChatChannelSend(t *testing.T) {
	setup()
	defer teardown()

	messageRequest := ChatChannelMessageRequest{
		Id:         "1234",
		ChannelUrl: "channel_url",
		Message:    "The medium is the message",
		Data:       "message data",
	}

	mux.HandleFunc("/channel/send", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelMessageRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		if !reflect.DeepEqual(body, messageRequest) {
			t.Errorf("ChatChannel.Send API call received %+v, expected %+v", body, messageRequest)
		}

		CheckForAuthParam(t, r, body)

		response := "{}"

		fmt.Fprint(w, response)
	})

	_, err := client.Chat.Send(&messageRequest)

	if err != nil {
		t.Errorf("ChatChannel.Send returned error: %v", err)
	}
}

func TestChatChannelGetMetadata(t *testing.T) {
	setup()
	defer teardown()

	metadataRequest := ChatChannelMetadataRequest{
		ChannelUrl: "channel_url",
		Keys:       []string{"key1", "key2"},
	}

	mux.HandleFunc("/channel/get_metadata", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelMetadataRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metadataRequest) {
			t.Errorf("ChatChannel.GetMetadata API call received %+v, expected %+v", body, metadataRequest)
		}

		response := `
	    {
	      "key1": "foo",
	      "key2": "bar"
		}`
		fmt.Fprint(w, response)
	})

	metadata, _, err := client.Chat.GetMetadata(&metadataRequest)

	if err != nil {
		t.Errorf("ChatChannel.GetMetadata returned error: %v", err)
	}

	expected := map[string]string{
		"key1": "foo",
		"key2": "bar",
	}

	if !reflect.DeepEqual(metadata, expected) {
		t.Errorf("ChatChannel.GetMetadata returned %+v, expected %+v", metadata, expected)
	}
}

func TestChatChannelSetMetadata(t *testing.T) {
	setup()
	defer teardown()

	metadataRequest := ChatChannelSetMetadataRequest{
		ChannelUrl: "channel_url",
		Data:       map[string]string{"key1": "foo", "key2": "bar"},
	}

	mux.HandleFunc("/channel/set_metadata", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelSetMetadataRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metadataRequest) {
			t.Errorf("ChatChannel.SetMetadata API call received %+v, expected %+v", body, metadataRequest)
		}

		response := `
	    {
	      "key1": "foo",
	      "key2": "bar"
		}`
		fmt.Fprint(w, response)
	})

	metadata, _, err := client.Chat.SetMetadata(&metadataRequest)

	if err != nil {
		t.Errorf("ChatChannel.SetMetadata returned error: %v", err)
	}

	expected := map[string]string{
		"key1": "foo",
		"key2": "bar",
	}

	if !reflect.DeepEqual(metadata, expected) {
		t.Errorf("ChatChannel.SetMetadata returned %+v, expected %+v", metadata, expected)
	}
}

func TestChatChannelGetMetacounter(t *testing.T) {
	setup()
	defer teardown()

	metacounterRequest := ChatChannelMetacounterRequest{
		ChannelUrl: "channel_url",
		Keys:       []string{"key1", "key2"},
	}

	mux.HandleFunc("/channel/get_metacounter", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelMetacounterRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metacounterRequest) {
			t.Errorf("ChatChannel.GetMetacounter API call received %+v, expected %+v", body, metacounterRequest)
		}

		response := `
	    {
	      "key1": 34,
	      "key2": 56
		}`
		fmt.Fprint(w, response)
	})

	metacounters, _, err := client.Chat.GetMetacounter(&metacounterRequest)

	if err != nil {
		t.Errorf("ChatChannel.GetMetacounter returned error: %v", err)
	}

	expected := map[string]int{
		"key1": 34,
		"key2": 56,
	}

	if !reflect.DeepEqual(metacounters, expected) {
		t.Errorf("ChatChannel.GetMetacounter returned %+v, expected %+v", metacounters, expected)
	}
}

func TestChatChannelSetMetacounter(t *testing.T) {
	setup()
	defer teardown()

	metacounterRequest := ChatChannelSetMetacounterRequest{
		ChannelUrl: "channel_url",
		Data:       map[string]int{"key1": 45, "key2": 67},
	}

	mux.HandleFunc("/channel/set_metacounter", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelSetMetacounterRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metacounterRequest) {
			t.Errorf("ChatChannel.SetMetacounter API call received %+v, expected %+v", body, metacounterRequest)
		}

		response := `
	    {
	      "key1": 45,
	      "key2": 67
		}`
		fmt.Fprint(w, response)
	})

	metacounters, _, err := client.Chat.SetMetacounter(&metacounterRequest)

	if err != nil {
		t.Errorf("ChatChannel.SetMetacounter returned error: %v", err)
	}

	expected := map[string]int{
		"key1": 45,
		"key2": 67,
	}

	if !reflect.DeepEqual(metacounters, expected) {
		t.Errorf("ChatChannel.SetMetacounter returned %+v, expected %+v", metacounters, expected)
	}
}

func TestChatChannelIncreaseMetacounter(t *testing.T) {
	setup()
	defer teardown()

	metacounterRequest := ChatChannelSetMetacounterRequest{
		ChannelUrl: "channel_url",
		Data:       map[string]int{"key1": 2, "key2": 5},
	}

	mux.HandleFunc("/channel/incr_metacounter", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelSetMetacounterRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metacounterRequest) {
			t.Errorf("ChatChannel.IncreaseMetacounter API call received %+v, expected %+v", body, metacounterRequest)
		}

		response := `
	    {
	      "key1": 40,
	      "key2": 63
		}`
		fmt.Fprint(w, response)
	})

	metacounters, _, err := client.Chat.IncreaseMetacounter(&metacounterRequest)

	if err != nil {
		t.Errorf("ChatChannel.IncreaseMetacounter returned error: %v", err)
	}

	expected := map[string]int{
		"key1": 40,
		"key2": 63,
	}

	if !reflect.DeepEqual(metacounters, expected) {
		t.Errorf("ChatChannel.IncreaseMetacounter returned %+v, expected %+v", metacounters, expected)
	}
}

func TestChatChannelDecreaseMetacounter(t *testing.T) {
	setup()
	defer teardown()

	metacounterRequest := ChatChannelSetMetacounterRequest{
		ChannelUrl: "channel_url",
		Data:       map[string]int{"key1": 2, "key2": 5},
	}

	mux.HandleFunc("/channel/decr_metacounter", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelSetMetacounterRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metacounterRequest) {
			t.Errorf("ChatChannel.DecreaseMetacounter API call received %+v, expected %+v", body, metacounterRequest)
		}

		response := `
	    {
	      "key1": 4,
	      "key2": 6
		}`
		fmt.Fprint(w, response)
	})

	metacounters, _, err := client.Chat.DecreaseMetacounter(&metacounterRequest)

	if err != nil {
		t.Errorf("ChatChannel.DecreaseMetacounter returned error: %v", err)
	}

	expected := map[string]int{
		"key1": 4,
		"key2": 6,
	}

	if !reflect.DeepEqual(metacounters, expected) {
		t.Errorf("ChatChannel.DecreaseMetacounter returned %+v, expected %+v", metacounters, expected)
	}
}

func TestChatChannelMessageCount(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channel/message_count", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ChatChannelRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedChannelUrl := "chat_channel_url"
		if body.ChannelUrl != expectedChannelUrl {
			t.Errorf("ChatChannel.MessageCount API call received ChannelUrl %s, expected %s", body.ChannelUrl, expectedChannelUrl)
		}

		response := `
	    {
	       "message_count": 415
		}`
		fmt.Fprint(w, response)
	})

	messageCount, _, err := client.Chat.MessageCount("chat_channel_url")

	if err != nil {
		t.Errorf("ChatChannel.MessageCount returned error: %v", err)
	}

	expected := &MessageCount{
		MessageCount: 415,
	}

	if !reflect.DeepEqual(messageCount, expected) {
		t.Errorf("ChatChannel.MessageCount returned %+v, expected %+v", messageCount, expected)
	}
}
