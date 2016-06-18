package sendbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAdminBroadcastMessage(t *testing.T) {
	setup()
	defer teardown()

	broadcastRequest := BroadcastMessageRequest{
		ChannelUrls: []string{"chat_channel_url1", "chat_channel_url2"},
		Message:     "hello world",
		Persistent:  true,
		Data:        "extra data",
	}

	mux.HandleFunc("/admin/broadcast_message", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := BroadcastMessageRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, broadcastRequest) {
			t.Errorf("Admin.BroadcastMessage API call received %+v, expected %+v", body, broadcastRequest)
		}

		response := "{}"
		fmt.Fprint(w, response)
	})

	_, err := client.Admin.BroadcastMessage(&broadcastRequest)
	if err != nil {
		t.Errorf("Admin.BroadcastMessage returned error: %v", err)
	}
}

func TestAdminReadMessages(t *testing.T) {
	setup()
	defer teardown()

	readRequest := ReadMessagesRequest{
		ChannelUrl:    "chat_channel_url1",
		TargetUserIds: []string{"123", "456", "789"},
		Limit:         555,
		MessageId:     314516662341,
	}

	mux.HandleFunc("/admin/read_messages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := ReadMessagesRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, readRequest) {
			t.Errorf("Admin.ReadMessages API call received %+v, expected %+v", body, readRequest)
		}

		response := `
		[
		    {
		        "id": "123",
		        "nickanme": "bugs",
		        "message_id": 21315135632,
		        "timestamp": 613451361321,
		        "message": "hola mundial",
		        "file": {}
		    },
		    {
		        "id": "456",       
		        "nickanme": "elmer",
		        "message_id": 32151461234,
		        "timestamp": 1461461463312,
		        "message": "",
		        "file": {
		          "url": "https://path_to_some_file",
		          "custom": "custom file string",
		          "type": "file_type",   
		          "name": "file_name",   
		          "size": 45146124612
		        }
		    }		    
		]`
		fmt.Fprint(w, response)
	})

	expected := []AdminMessage{
		{
			Id:        "123",
			Nickname:  "bugs",
			MessageId: 21315135632,
			Timestamp: 613451361321,
			Message:   "hola mundial",
			File:      AdminMessageFile{},
		},
		{
			Id:        "456",
			Nickname:  "elmer",
			MessageId: 32151461234,
			Timestamp: 1461461463312,
			Message:   "",
			File: AdminMessageFile{
				Url:    "https://path_to_some_file",
				Custom: "custom file string",
				Type:   "file_type",
				Name:   "file_name",
				Size:   45146124612,
			},
		},
	}

	readMessages, _, err := client.Admin.ReadMessages(&readRequest)
	if err != nil {
		t.Errorf("Admin.ReadMessages returned error: %v", err)
	}

	if !reflect.DeepEqual(readMessages, expected) {
		t.Errorf("MessagingChannel.ReadMessages returned %+v, expected %+v", readMessages, expected)
	}
}

func TestAdminDeleteMessage(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/delete_message", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := struct {
			RequestDefaults
			MsgId string `json:"msg_id"`
		}{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedMsgId := "123456"
		if body.MsgId != expectedMsgId {
			t.Errorf("Admin.DeleteMessage API call received msg_id %s, expected %s", body.MsgId, expectedMsgId)
		}

		response := `
		{
		   "app_id": 1354342151,
		   "msg_id": 123456
		}`

		fmt.Fprint(w, response)
	})

	deletedMessage, _, err := client.Admin.DeleteMessage("123456")
	if err != nil {
		t.Errorf("Admin.DeleteMessage returned error: %v", err)
	}

	expected := &DeleteMessage{
		AppId: 1354342151,
		MsgId: 123456,
	}

	if !reflect.DeepEqual(deletedMessage, expected) {
		t.Errorf("Admin.DeleteMessage returned %+v, expected %+v", deletedMessage, expected)
	}
}

func TestAdminListMessagingChannels(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/list_messaging_channels", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := struct {
			RequestDefaults
			Id string `json:"id"`
		}{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedId := "123456"
		if body.Id != expectedId {
			t.Errorf("Admin.ListMessagingChannels API call received id %s, expected %s", body.Id, expectedId)
		}

		response := `
		[
		    {
		        "channel_url": "chat_channel_url",
		        "unread_message_count": 54,
		        "last_message": "Remember me",
		        "last_message_ts": 4524624642362,
		        "members": [
		            {
		                "id": "123",
		                "image": "http://sendbird.com/123.jpg",
		                "name": "bugs"
		            },
		            {
		                "id": "456",
		                "image": "http://sendbird.com/456.jpg",
		                "name": "elmer"
		            }
		        ]
		    }		  
		]`

		fmt.Fprint(w, response)
	})

	messagingChannels, _, err := client.Admin.ListMessagingChannels("123456")
	if err != nil {
		t.Errorf("Admin.ListMessagingChannels returned error: %v", err)
	}

	expected := []AdminMessagingChannel{
		{
			ChannelUrl:         "chat_channel_url",
			UnreadMessageCount: 54,
			LastMessage:        "Remember me",
			LastMessageTS:      4524624642362,
			Members: []Member{
				{
					Id:    "123",
					Image: "http://sendbird.com/123.jpg",
					Name:  "bugs",
				},
				{
					Id:    "456",
					Image: "http://sendbird.com/456.jpg",
					Name:  "elmer",
				},
			},
		},
	}

	if !reflect.DeepEqual(messagingChannels, expected) {
		t.Errorf("Admin.ListMessagingChannels returned %+v, expected %+v", messagingChannels, expected)
	}
}

func TestAdminMuteAllChannels(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/mute", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := struct {
			RequestDefaults
			Id string `json:"id"`
		}{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedId := "123456"
		if body.Id != expectedId {
			t.Errorf("Admin.MuteAllChannels API call received id %s, expected %s", body.Id, expectedId)
		}

		response := "{}"

		fmt.Fprint(w, response)
	})

	_, err := client.Admin.MuteAllChannels("123456")
	if err != nil {
		t.Errorf("Admin.MuteAllChannels returned error: %v", err)
	}
}

func TestAdminMute(t *testing.T) {
	setup()
	defer teardown()

	muteRequest := MuteRequest{
		Id:          "123",
		ChannelUrls: []string{"channel_1", "channel_2"},
		IsSoftMute:  true,
	}
	mux.HandleFunc("/admin/mute", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := MuteRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, muteRequest) {
			t.Errorf("Admin.Mute API call received %+v, expected %+v", body, muteRequest)
		}

		response := `["channel_1", "channel_2"]`

		fmt.Fprint(w, response)
	})

	expected := []string{"channel_1", "channel_2"}

	muteChannels, _, err := client.Admin.Mute(&muteRequest)
	if err != nil {
		t.Errorf("Admin.Mute returned error: %v", err)
	}

	if !reflect.DeepEqual(muteChannels, expected) {
		t.Errorf("Admin.Mute returned %+v, expected %+v", muteChannels, expected)
	}
}

func TestAdminUnMuteAllChannels(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/unmute", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := struct {
			RequestDefaults
			Id string `json:"id"`
		}{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedId := "123456"
		if body.Id != expectedId {
			t.Errorf("Admin.UnMuteAllChannels API call received id %s, expected %s", body.Id, expectedId)
		}

		response := "{}"

		fmt.Fprint(w, response)
	})

	_, err := client.Admin.UnMuteAllChannels("123456")
	if err != nil {
		t.Errorf("Admin.UnMuteAllChannels returned error: %v", err)
	}
}

func TestAdminUnMute(t *testing.T) {
	setup()
	defer teardown()

	unmuteRequest := UnMuteRequest{
		Id:          "123",
		ChannelUrls: []string{"channel_1", "channel_2"},
	}
	mux.HandleFunc("/admin/unmute", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := UnMuteRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		if !reflect.DeepEqual(body, unmuteRequest) {
			t.Errorf("Admin.UnMute API call received %+v, expected %+v", body, unmuteRequest)
		}

		response := `["channel_1", "channel_2"]`

		fmt.Fprint(w, response)
	})

	expected := []string{"channel_1", "channel_2"}

	unmuteChannels, _, err := client.Admin.UnMute(&unmuteRequest)
	if err != nil {
		t.Errorf("Admin.UnMute returned error: %v", err)
	}

	if !reflect.DeepEqual(unmuteChannels, expected) {
		t.Errorf("Admin.UnMute returned %+v, expected %+v", unmuteChannels, expected)
	}
}

func TestAdminMuteList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/mute_list", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := struct {
			RequestDefaults
			ChannelUrls []string `json:"channel_urls"`
		}{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedChannelRequest := []string{"channel_1", "channel_2"}
		if !reflect.DeepEqual(body.ChannelUrls, expectedChannelRequest) {
			t.Errorf("Admin.MuteList API call received id %s, expected %s", body.ChannelUrls, expectedChannelRequest)
		}

		response := `["123", "456", "789"]`

		fmt.Fprint(w, response)
	})

	expected := []string{"123", "456", "789"}

	userIds, _, err := client.Admin.MuteList([]string{"channel_1", "channel_2"})
	if err != nil {
		t.Errorf("Admin.MuteList returned error: %v", err)
	}

	if !reflect.DeepEqual(userIds, expected) {
		t.Errorf("Admin.MuteList returned %+v, expected %+v", userIds, expected)
	}
}

func TestConcurrentUserCount(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/ccu_count", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := RequestDefaults{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		response := `
		{
			"count": 1314
		}`

		fmt.Fprint(w, response)
	})

	expected := &ConcurrentUserCount{Count: 1314}

	userCount, _, err := client.Admin.ConcurrentUserCount()
	if err != nil {
		t.Errorf("Admin.ConcurrentUserCount returned error: %v", err)
	}

	if !reflect.DeepEqual(userCount, expected) {
		t.Errorf("Admin.ConcurrentUserCount returned %+v, expected %+v", userCount, expected)
	}
}

func TestMemberCountInChannel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/admin/member_count", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := struct {
			RequestDefaults
			ChannelUrl string `json:"channel_url"`
		}{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		expectedChannelUrl := "channel_43"
		if body.ChannelUrl != expectedChannelUrl {
			t.Errorf("Admin.MemberCountInChannel API call received channelUrl %s, expected %s", body.ChannelUrl, expectedChannelUrl)
		}

		response := `
		{
		  "accumulated_member_count": 12,
		  "online_member_count": 31,
		  "member_count": 61
		}`

		fmt.Fprint(w, response)
	})

	expected := &ChannelMemberCount{
		AccumulatedMemberCount: 12,
		OnlineMemberCount:      31,
		MemberCount:            61,
	}

	userCount, _, err := client.Admin.MemberCountInChannel("channel_43")
	if err != nil {
		t.Errorf("Admin.MemberCountInChannel returned error: %v", err)
	}

	if !reflect.DeepEqual(userCount, expected) {
		t.Errorf("Admin.MemberCountInChannel returned %+v, expected %+v", userCount, expected)
	}
}
