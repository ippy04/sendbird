package sendbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUserCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := UserRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		response := `
    	{ 
    		"id": "123456",
   			"user_id": "7890",
		    "nickname": "nickname",
		    "picture": "http://sendbird.com/picture_url",
		    "access_token": "ABCDEFG"
		}`
		fmt.Fprint(w, response)
	})

	expected := &User{
		Id:          "123456",
		UserId:      "7890",
		Nickname:    "nickname",
		Picture:     "http://sendbird.com/picture_url",
		AccessToken: "ABCDEFG",
	}

	userRequest := UserRequest{
		Id:               "123456",
		Nickname:         "nickname",
		ImageUrl:         "http://sendbird.com/picture_url",
		IssueAccessToken: true,
	}

	user, _, err := client.Users.Create(&userRequest)
	if err != nil {
		t.Errorf("User.Create returned error: %v", err)
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("User.Create returned %+v, expected %+v", user, expected)
	}

}

func TestUserUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := UserRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		response := `
    	{ 
    		"id": "123456",
   			"user_id": "7890",
		    "nickname": "nickname",
		    "picture": "http://sendbird.com/picture_url",
		    "access_token": "ABCDEFG"
		}`
		fmt.Fprint(w, response)
	})

	expected := &User{
		Id:          "123456",
		UserId:      "7890",
		Nickname:    "nickname",
		Picture:     "http://sendbird.com/picture_url",
		AccessToken: "ABCDEFG",
	}

	userRequest := UserRequest{
		Id:               "123456",
		Nickname:         "nickname",
		ImageUrl:         "http://sendbird.com/picture_url",
		IssueAccessToken: true,
	}

	user, _, err := client.Users.Update(&userRequest)
	if err != nil {
		t.Errorf("User.Update returned error: %v", err)
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("User.Update returned %+v, expected %+v", user, expected)
	}

}

func TestUserAuth(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/auth", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := UserRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		response := `
    	{ 
    		"id": "123456",
   			"user_id": "7890",
		    "nickname": "nickname",
		    "picture": "http://sendbird.com/picture_url",
		    "access_token": "ABCDEFG",
		    "is_active": true
		}`
		fmt.Fprint(w, response)
	})

	expected := &User{
		Id:          "123456",
		UserId:      "7890",
		Nickname:    "nickname",
		Picture:     "http://sendbird.com/picture_url",
		AccessToken: "ABCDEFG",
		IsActive:    true,
	}

	userRequest := UserRequest{
		Id:               "123456",
		IssueAccessToken: true,
	}

	user, _, err := client.Users.Auth(&userRequest)
	if err != nil {
		t.Errorf("User.Auth returned error: %v", err)
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("User.Auth returned %+v, expected %+v", user, expected)
	}

}

func TestUserBlock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/block", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := BlockRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		response := "{}"
		fmt.Fprint(w, response)
	})

	blockRequest := BlockRequest{
		Id:       "12345",
		TargetId: "67890",
	}

	_, err := client.Users.Block(&blockRequest)
	if err != nil {
		t.Errorf("User.Block returned error: %v", err)
	}
}

func TestUserUnBlock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/unblock", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := BlockRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		response := "{}"
		fmt.Fprint(w, response)
	})

	unblockRequest := BlockRequest{
		Id:       "12345",
		TargetId: "67890",
	}

	_, err := client.Users.UnBlock(&unblockRequest)
	if err != nil {
		t.Errorf("User.UnBlock returned error: %v", err)
	}
}

func TestUserDeactivate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/deactivate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		CheckForAuthContentType(t, r)

		body := DeactivateRequest{}
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&body)
		if err != nil {
			t.Errorf("error decoding request json", err)
		}

		CheckForAuthParam(t, r, body)

		response := "{}"
		fmt.Fprint(w, response)
	})

	deactivateRequest := DeactivateRequest{
		Id: "12345",
	}

	_, err := client.Users.Deactivate(&deactivateRequest)
	if err != nil {
		t.Errorf("User.Deactivate returned error: %v", err)
	}
}
