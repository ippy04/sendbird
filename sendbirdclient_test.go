package sendbird

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	client *SendbirdClient
	server *httptest.Server
)

const TestApiToken = "API_TOKEN_1"

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient("SENDBIRD_API_TOKEN", "SENDBIRD_APP_ID", nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient("SENDBIRD_API_TOKEN", "SENDBIRD_APP_ID", nil)

	if c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.BaseURL.String(), defaultBaseURL)
	}

}

func CheckForAuthContentType(t *testing.T, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json, charset=utf8" {
		t.Errorf("Content-Type Request Header should be set to \"application/json, charset=utf8\"")
	}
}
func CheckForAuthParam(t *testing.T, r *http.Request, d interface{}) {
	auth := reflect.ValueOf(d).FieldByName("Auth").String()
	if auth == "" {
		t.Errorf("Required request parameter of Auth not populated")
	}
}
func CheckForV2ApiTokenParam(t *testing.T, r *http.Request, d interface{}) {
	api_token := reflect.ValueOf(d).FieldByName("api_token").String()
	if api_token == "" {
		t.Errorf("Required request parameter of api_token not populated")
	}
}
func CheckForV2ApiTokenQueryString(t *testing.T, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		t.Errorf("Bot.Get - Unable to parse querystring", err)
	}
	if r.Form.Get("api_token") == "" {
		t.Errorf("Required querystring parameter of api_token not populated")
	}
}
