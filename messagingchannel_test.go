package sendbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMessagingChannelCreate(t *testing.T) {
	setup()
	defer teardown()

	messagingRequest := MessagingChannelRequest{
		Name:     "new messaging channel",
		CoverUrl: "https://sendbird/images/cover_url.jpg",
		Data:     "extra data",
		IsGroup:  true,
	}

	mux.HandleFunc("/messaging/create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, messagingRequest) {
			t.Errorf("MessagingChannel.Create API call received %+v, expected %+v", body, messagingRequest)
		}

		response := `
		{
		    "channel": {
		      "channel_url": "messaging_channel_url",
		      "data": "extra data",
		      "name": "new messaging channel",
		      "is_group": true,
		      "cover_url": "https://sendbird/images/cover_url.jpg"
		    }
		}`

		fmt.Fprint(w, response)
	})

	expected := &MessagingChannel{
		Name:       "new messaging channel",
		ChannelUrl: "messaging_channel_url",
		IsGroup:    true,
		CoverUrl:   "https://sendbird/images/cover_url.jpg",
		Data:       "extra data",
	}

	messagingChannel, _, err := client.Messaging.Create(&messagingRequest)
	if err != nil {
		t.Errorf("MessagingChannel.Create returned error: %v", err)
	}

	if !reflect.DeepEqual(messagingChannel, expected) {
		t.Errorf("MessagingChannel.Create returned %+v, expected %+v", messagingChannel, expected)
	}

}

func TestMessagingChannelUpdate(t *testing.T) {
	setup()
	defer teardown()

	messagingRequest := MessagingChannelUpdateRequest{
		Name:     "new messaging channel",
		CoverUrl: "https://sendbird/images/cover_url.jpg",
		Data:     "extra data",
		IsGroup:  true,
	}

	mux.HandleFunc("/messaging/update", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelUpdateRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, messagingRequest) {
			t.Errorf("MessagingChannel.Update API call received %+v, expected %+v", body, messagingRequest)
		}

		response := `
		{
		    "channel": {
		      "channel_url": "messaging_channel_url",
		      "data": "extra data",
		      "name": "new messaging channel",
		      "is_group": true,
		      "cover_url": "https://sendbird/images/cover_url.jpg"
		    }
		}`

		fmt.Fprint(w, response)
	})

	expected := &MessagingChannel{
		Name:       "new messaging channel",
		ChannelUrl: "messaging_channel_url",
		IsGroup:    true,
		CoverUrl:   "https://sendbird/images/cover_url.jpg",
		Data:       "extra data",
	}

	messagingChannel, _, err := client.Messaging.Update(&messagingRequest)
	if err != nil {
		t.Errorf("MessagingChannel.Update returned error: %v", err)
	}

	if !reflect.DeepEqual(messagingChannel, expected) {
		t.Errorf("MessagingChannel.Update returned %+v, expected %+v", messagingChannel, expected)
	}

}

func TestMessagingChannelDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/messaging/delete", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelUpdateRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedChannelUrl := "messaging_channel_url"
		if body.ChannelUrl != expectedChannelUrl {
			t.Errorf("MessagingChannel.View API call received ChannelUrl %s, expected %s", body.ChannelUrl, expectedChannelUrl)
		}

		response := `
		{
		    "channel": {
		        "channel_url": "messaging_channel_url"
		    }
		}`

		fmt.Fprint(w, response)
	})

	expected := &MessagingChannelUrl{
		Channel: ChannelUrl{ChannelUrl: "messaging_channel_url"},
	}
	messagingChannel, _, err := client.Messaging.Delete("messaging_channel_url")
	if err != nil {
		t.Errorf("MessagingChannel.Delete returned error: %v", err)
	}

	if !reflect.DeepEqual(messagingChannel, expected) {
		t.Errorf("MessagingChannel.Delete returned %+v, expected %+v", messagingChannel, expected)
	}
}

func TestMessagingChannelInvite(t *testing.T) {
	setup()
	defer teardown()

	inviteRequest := MessagingChannelInviteRequest{
		ChannelUrl: "messaging_channel_url",
		UserIds:    []string{"123", "456", "789"},
	}

	mux.HandleFunc("/messaging/invite", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelInviteRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, inviteRequest) {
			t.Errorf("MessagingChannel.Invite API call received %+v, expected %+v", body, inviteRequest)
		}

		response := `
		{
		    "channel": {
		        "channel_url": "messaging_channel_url"
		    }
		}`

		fmt.Fprint(w, response)
	})

	expected := &MessagingChannelUrl{
		Channel: ChannelUrl{ChannelUrl: "messaging_channel_url"},
	}
	messagingChannel, _, err := client.Messaging.Invite(&inviteRequest)
	if err != nil {
		t.Errorf("MessagingChannel.Invite returned error: %v", err)
	}

	if !reflect.DeepEqual(messagingChannel, expected) {
		t.Errorf("MessagingChannel.Invite returned %+v, expected %+v", messagingChannel, expected)
	}
}

func TestMessagingChannelHide(t *testing.T) {
	setup()
	defer teardown()

	hideRequest := MessagingChannelHideRequest{
		ChannelUrl: "messaging_channel_url",
		Id:         "456",
	}

	mux.HandleFunc("/messaging/hide", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelHideRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, hideRequest) {
			t.Errorf("MessagingChannel.Hide API call received %+v, expected %+v", body, hideRequest)
		}

		response := `
		{
		    "channel": {
		        "channel_url": "messaging_channel_url"
		    }
		}`

		fmt.Fprint(w, response)
	})

	expected := &MessagingChannelUrl{
		Channel: ChannelUrl{ChannelUrl: "messaging_channel_url"},
	}
	messagingChannel, _, err := client.Messaging.Hide(&hideRequest)
	if err != nil {
		t.Errorf("MessagingChannel.Hide returned error: %v", err)
	}

	if !reflect.DeepEqual(messagingChannel, expected) {
		t.Errorf("MessagingChannel.Hide returned %+v, expected %+v", messagingChannel, expected)
	}
}

func TestMessagingChannelLeave(t *testing.T) {
	setup()
	defer teardown()

	leaveRequest := MessagingChannelLeaveRequest{
		ChannelUrl: "messaging_channel_url",
		UserIds:    []string{"123", "456", "789"},
	}

	mux.HandleFunc("/messaging/leave", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelLeaveRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, leaveRequest) {
			t.Errorf("MessagingChannel.Leave API call received %+v, expected %+v", body, leaveRequest)
		}

		response := `
		{
		    "channel": {
		        "channel_url": "messaging_channel_url"
		    }
		}`

		fmt.Fprint(w, response)
	})

	expected := &MessagingChannelUrl{
		Channel: ChannelUrl{ChannelUrl: "messaging_channel_url"},
	}
	messagingChannel, _, err := client.Messaging.Leave(&leaveRequest)
	if err != nil {
		t.Errorf("MessagingChannel.Leave returned error: %v", err)
	}

	if !reflect.DeepEqual(messagingChannel, expected) {
		t.Errorf("MessagingChannel.Leave returned %+v, expected %+v", messagingChannel, expected)
	}
}

func TestMessagingChannelView(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/messaging/view", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelUpdateRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedChannelUrl := "messaging_channel_url"
		if body.ChannelUrl != expectedChannelUrl {
			t.Errorf("MessagingChannel.View API call received ChannelUrl %s, expected %s", body.ChannelUrl, expectedChannelUrl)
		}

		response := `
	    {
	      "channel_url": "messaging_channel_url",
	      "last_message": "I am the last hope",
	      "last_message_ts": 1234567890,   
	      "created_at": 12421521552,
	      "members": [
          	{	
	          "id": "123",               
	          "image": "http://sendbird/user123.jpg",
	          "name": "bugs"          
	        },
          	{	
	          "id": "456",               
	          "image": "http://sendbird/user456.jpg",         
	          "name": "elmer"
	        }
	      ]
		}`
		fmt.Fprint(w, response)
	})

	MessagingChannel, _, err := client.Messaging.View("messaging_channel_url")

	if err != nil {
		t.Errorf("MessagingChannel.View returned error: %v", err)
	}

	expected := &MessagingChannelView{
		ChannelUrl:    "messaging_channel_url",
		LastMessage:   "I am the last hope",
		LastMessageTS: 1234567890,
		CreatedAt:     12421521552,
		Members: []Member{
			{
				Id:    "123",
				Image: "http://sendbird/user123.jpg",
				Name:  "bugs",
			},
			{
				Id:    "456",
				Image: "http://sendbird/user456.jpg",
				Name:  "elmer",
			},
		},
	}

	if !reflect.DeepEqual(MessagingChannel, expected) {
		t.Errorf("MessagingChannel.View returned %+v, expected %+v", MessagingChannel, expected)
	}
}

func TestMessagingChannelGetMetadata(t *testing.T) {
	setup()
	defer teardown()

	metadataRequest := MessagingChannelMetadataRequest{
		ChannelUrl: "channel_url",
		Keys:       []string{"key1", "key2"},
	}

	mux.HandleFunc("/messaging/get_metadata", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelMetadataRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metadataRequest) {
			t.Errorf("MessagingChannel.GetMetadata API call received %+v, expected %+v", body, metadataRequest)
		}

		response := `
	    {
	      "key1": "foo",
	      "key2": "bar"
		}`
		fmt.Fprint(w, response)
	})

	metadata, _, err := client.Messaging.GetMetadata(&metadataRequest)

	if err != nil {
		t.Errorf("MessagingChannel.GetMetadata returned error: %v", err)
	}

	expected := map[string]string{
		"key1": "foo",
		"key2": "bar",
	}

	if !reflect.DeepEqual(metadata, expected) {
		t.Errorf("MessagingChannel.GetMetadata returned %+v, expected %+v", metadata, expected)
	}
}

func TestMessagingChannelSetMetadata(t *testing.T) {
	setup()
	defer teardown()

	metadataRequest := MessagingChannelSetMetadataRequest{
		ChannelUrl: "channel_url",
		Data:       map[string]string{"key1": "foo", "key2": "bar"},
	}

	mux.HandleFunc("/messaging/set_metadata", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelSetMetadataRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metadataRequest) {
			t.Errorf("MessagingChannel.SetMetadata API call received %+v, expected %+v", body, metadataRequest)
		}

		response := `
	    {
	      "key1": "foo",
	      "key2": "bar"
		}`
		fmt.Fprint(w, response)
	})

	metadata, _, err := client.Messaging.SetMetadata(&metadataRequest)

	if err != nil {
		t.Errorf("MessagingChannel.SetMetadata returned error: %v", err)
	}

	expected := map[string]string{
		"key1": "foo",
		"key2": "bar",
	}

	if !reflect.DeepEqual(metadata, expected) {
		t.Errorf("MessagingChannel.SetMetadata returned %+v, expected %+v", metadata, expected)
	}
}

func TestMessagingChannelGetMetacounter(t *testing.T) {
	setup()
	defer teardown()

	metacounterRequest := MessagingChannelMetacounterRequest{
		ChannelUrl: "channel_url",
		Keys:       []string{"key1", "key2"},
	}

	mux.HandleFunc("/messaging/get_metacounter", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelMetacounterRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metacounterRequest) {
			t.Errorf("MessagingChannel.GetMetacounter API call received %+v, expected %+v", body, metacounterRequest)
		}

		response := `
	    {
	      "key1": 34,
	      "key2": 56
		}`
		fmt.Fprint(w, response)
	})

	metacounters, _, err := client.Messaging.GetMetacounter(&metacounterRequest)

	if err != nil {
		t.Errorf("MessagingChannel.GetMetacounter returned error: %v", err)
	}

	expected := map[string]int{
		"key1": 34,
		"key2": 56,
	}

	if !reflect.DeepEqual(metacounters, expected) {
		t.Errorf("MessagingChannel.GetMetacounter returned %+v, expected %+v", metacounters, expected)
	}
}

func TestMessagingChannelSetMetacounter(t *testing.T) {
	setup()
	defer teardown()

	metacounterRequest := MessagingChannelSetMetacounterRequest{
		ChannelUrl: "channel_url",
		Data:       map[string]int{"key1": 45, "key2": 67},
	}

	mux.HandleFunc("/messaging/set_metacounter", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelSetMetacounterRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metacounterRequest) {
			t.Errorf("MessagingChannel.SetMetacounter API call received %+v, expected %+v", body, metacounterRequest)
		}

		response := `
	    {
	      "key1": 45,
	      "key2": 67
		}`
		fmt.Fprint(w, response)
	})

	metacounters, _, err := client.Messaging.SetMetacounter(&metacounterRequest)

	if err != nil {
		t.Errorf("MessagingChannel.SetMetacounter returned error: %v", err)
	}

	expected := map[string]int{
		"key1": 45,
		"key2": 67,
	}

	if !reflect.DeepEqual(metacounters, expected) {
		t.Errorf("MessagingChannel.SetMetacounter returned %+v, expected %+v", metacounters, expected)
	}
}

func TestMessagingChannelIncreaseMetacounter(t *testing.T) {
	setup()
	defer teardown()

	metacounterRequest := MessagingChannelSetMetacounterRequest{
		ChannelUrl: "channel_url",
		Data:       map[string]int{"key1": 2, "key2": 5},
	}

	mux.HandleFunc("/messaging/incr_metacounter", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelSetMetacounterRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metacounterRequest) {
			t.Errorf("MessagingChannel.IncreaseMetacounter API call received %+v, expected %+v", body, metacounterRequest)
		}

		response := `
	    {
	      "key1": 40,
	      "key2": 63
		}`
		fmt.Fprint(w, response)
	})

	metacounters, _, err := client.Messaging.IncreaseMetacounter(&metacounterRequest)

	if err != nil {
		t.Errorf("MessagingChannel.IncreaseMetacounter returned error: %v", err)
	}

	expected := map[string]int{
		"key1": 40,
		"key2": 63,
	}

	if !reflect.DeepEqual(metacounters, expected) {
		t.Errorf("MessagingChannel.IncreaseMetacounter returned %+v, expected %+v", metacounters, expected)
	}
}

func TestMessagingChannelDecreaseMetacounter(t *testing.T) {
	setup()
	defer teardown()

	metacounterRequest := MessagingChannelSetMetacounterRequest{
		ChannelUrl: "channel_url",
		Data:       map[string]int{"key1": 2, "key2": 5},
	}

	mux.HandleFunc("/messaging/decr_metacounter", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelSetMetacounterRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, metacounterRequest) {
			t.Errorf("MessagingChannel.DecreaseMetacounter API call received %+v, expected %+v", body, metacounterRequest)
		}

		response := `
	    {
	      "key1": 4,
	      "key2": 6
		}`
		fmt.Fprint(w, response)
	})

	metacounters, _, err := client.Messaging.DecreaseMetacounter(&metacounterRequest)

	if err != nil {
		t.Errorf("MessagingChannel.DecreaseMetacounter returned error: %v", err)
	}

	expected := map[string]int{
		"key1": 4,
		"key2": 6,
	}

	if !reflect.DeepEqual(metacounters, expected) {
		t.Errorf("MessagingChannel.DecreaseMetacounter returned %+v, expected %+v", metacounters, expected)
	}
}

func TestMessagingChannelMessageCount(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/messaging/message_count", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MessagingChannelUpdateRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedChannelUrl := "messaging_channel_url"
		if body.ChannelUrl != expectedChannelUrl {
			t.Errorf("MessagingChannel.MessageCount API call received ChannelUrl %s, expected %s", body.ChannelUrl, expectedChannelUrl)
		}

		response := `
	    {
	       "message_count": 415
		}`
		fmt.Fprint(w, response)
	})

	messageCount, _, err := client.Messaging.MessageCount("messaging_channel_url")

	if err != nil {
		t.Errorf("MessagingChannel.MessageCount returned error: %v", err)
	}

	expected := &MessageCount{
		MessageCount: 415,
	}

	if !reflect.DeepEqual(messageCount, expected) {
		t.Errorf("MessagingChannel.MessageCount returned %+v, expected %+v", messageCount, expected)
	}
}
