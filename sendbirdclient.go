package sendbird

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://api.sendbird.com"
	contentType    = "application/json, charset=utf8"
)

// -----------------------
// Response is a Sendbird response. This wraps the standard http.Response returned from Sendbird.
type Response struct {
	*http.Response
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	Message string
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// Client manages communication with Sendbird
type SendbirdClient struct {
	// HTTP client used to communicate with the DO API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	AppId       string
	ApiToken    string
	ContentType string

	// Services used for communicating with the API
	Users     UserService
	Chat      ChatChannelService
	Messaging MessagingChannelService
	Admin     AdminService
	Bot       BotService

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback
}

// NewClient returns a new Sendbird API client.
func NewClient(appId string, apiToken string, httpClient *http.Client) *SendbirdClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &SendbirdClient{
		client:      httpClient,
		BaseURL:     baseURL,
		AppId:       appId,
		ApiToken:    apiToken,
		ContentType: contentType,
	}

	c.Users = &UserServiceOp{client: c}
	c.Chat = &ChatChannelServiceOp{client: c}
	c.Messaging = &MessagingChannelServiceOp{client: c}
	c.Admin = &AdminServiceOp{client: c}
	c.Bot = &BotServiceOp{client: c}

	return c
}

func (c *SendbirdClient) NormalizeId(Id string) string {
	lastSlash := strings.LastIndex(Id, "/")
	if lastSlash >= 0 {
		return Id[lastSlash+1:]
	} else {
		return Id
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *SendbirdClient) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	url := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", c.ContentType)

	return req, nil
}

// OnRequestCompleted sets the DO API request completion callback
func (c *SendbirdClient) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

// newResponse creates a new Response for the provided http.Response
func newResponse(r *http.Response) *Response {
	response := Response{Response: r}
	// response.populateRate()

	return &response
}

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	if r.StatusCode == 404 {
		return &ErrorResponse{Response: r, Message: "Resouce Not Found - please check the URL"}
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}

	return errorResponse
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *SendbirdClient) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {

		if w, ok := v.(io.Writer); ok {
			_, err := io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err := json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err
}
