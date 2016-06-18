package sendbird

import ()

type ChatChannelRequest struct {
	RequestDefaults
	ChannelUrl string `json:"channel_url,omitempty"` // Channel Url
	Name       string `json:"name,omitempty"`        // Topic
	CoverUrl   string `json:"cover_url,omitempty"`   // Cover Image Url
	Data       string `json:"data,omitempty"`        // Custom Channel Data
}
type ChatChannelUpdateRequest struct {
	RequestDefaults
	ChannelUrl       string   `json:"channel_url,omitempty"`        // Channel Url
	Name             string   `json:"name,omitempty"`               // Topic
	CoverUrl         string   `json:"cover_url,omitempty"`          // Cover Image Url
	Data             string   `json:"data,omitempty"`               // Custom Channel Data
	TargetChannelUrl string   `json:"target_channel_url,omitempty"` // Target Channel Url
	Ops              []string `json:"ops,omitempty"`                // a list of operators ex) ["user_id_1", "user_id_2"]
}
type ChatChannelMessageRequest struct {
	RequestDefaults
	Id         string `json:"id"`                    // Sender User ID
	ChannelUrl string `json:"channel_url,omitempty"` // (Optional) Channel URL
	Message    string `json:"message"`               // Message to be sent
	Data       string `json:"data"`                  // (Optional) Custom data field
}
type ChatChannelMetadataRequest struct {
	RequestDefaults
	ChannelUrl string   `json:"channel_url"` // Channel URL
	Keys       []string `json:"keys"`        //  Key names
}
type ChatChannelSetMetadataRequest struct {
	RequestDefaults
	ChannelUrl string            `json:"channel_url"` // Channel URL
	Data       map[string]string `json:"data"`
}
type ChatChannelMetacounterRequest struct {
	RequestDefaults
	ChannelUrl string   `json:"channel_url"` // Channel URL
	Keys       []string `json:"keys"`        //  Key names
}
type ChatChannelSetMetacounterRequest struct {
	RequestDefaults
	ChannelUrl string         `json:"channel_url"` // Channel URL
	Data       map[string]int `json:"data"`
}

type ChatChannel struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	ChannelUrl    string `json:"channel_url"`
	MemberCount   int    `json:"member_count"`
	CoverUrl      string `json:"cover_url"`
	CoverImageUrl string `json:"cover_image_url"` // used to handle json returning cover_url or cover_image_url
	Data          string `json:"data"`
	CreatedAt     int64  `json:"created_at"` // Epoch timestamp in milliseconds
}
type ChatChannelUpdate struct {
	ChatChannel
	Ops []string `json:"ops,omitempty"` // a list of operators
}
type Member struct {
	Id    string `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}
type ChatChannelView struct {
	ChatChannel
	Members []Member `json:"members"`
}
type MessageCount struct {
	MessageCount int `json:"message_count"`
}

// ChatChannelService is an interface for interfacing with the Chat Channel
// endpoints of the Sendbird API
type ChatChannelService interface {
	Create(*ChatChannelRequest) (*ChatChannel, *Response, error)
	List() ([]ChatChannel, *Response, error)
	Update(params *ChatChannelUpdateRequest) (*ChatChannelUpdate, *Response, error)
	Delete(channelUrl string) (*Response, error)
	View(channelUrl string) (*ChatChannelView, *Response, error)
	Send(params *ChatChannelMessageRequest) (*Response, error)
	GetMetadata(params *ChatChannelMetadataRequest) (map[string]string, *Response, error)
	SetMetadata(params *ChatChannelSetMetadataRequest) (map[string]string, *Response, error)
	GetMetacounter(params *ChatChannelMetacounterRequest) (map[string]int, *Response, error)
	SetMetacounter(params *ChatChannelSetMetacounterRequest) (map[string]int, *Response, error)
	IncreaseMetacounter(params *ChatChannelSetMetacounterRequest) (map[string]int, *Response, error)
	DecreaseMetacounter(params *ChatChannelSetMetacounterRequest) (map[string]int, *Response, error)
	MessageCount(channelUrl string) (*MessageCount, *Response, error)
}

// ChatChannelServiceOp handles communication with the Chat Channel related methods of
// the Sendbird API.
type ChatChannelServiceOp struct {
	client *SendbirdClient
}

var _ ChatChannelService = &ChatChannelServiceOp{}

func mapCoverImageUrlToCoverUrl(chatChannels []ChatChannel) {

	for i, _ := range chatChannels {
		chatChannels[i].CoverUrl = chatChannels[i].CoverImageUrl
	}

}

// Create a channel using channel URL / name
func (s *ChatChannelServiceOp) Create(params *ChatChannelRequest) (*ChatChannel, *Response, error) {

	path := "/channel/create"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	chatChannel := new(ChatChannel)
	resp, err := s.client.Do(req, chatChannel)

	if err != nil {
		return nil, resp, err
	}

	return chatChannel, resp, nil
}

// Retrieve channel list
func (s *ChatChannelServiceOp) List() ([]ChatChannel, *Response, error) {

	path := "/channel/list"
	params := &RequestDefaults{}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	chatChannels := []ChatChannel{}
	resp, err := s.client.Do(req, &chatChannels)

	mapCoverImageUrlToCoverUrl(chatChannels)

	if err != nil {
		return nil, resp, err
	}

	return chatChannels, resp, nil
}

// Update a channel's information
func (s *ChatChannelServiceOp) Update(params *ChatChannelUpdateRequest) (*ChatChannelUpdate, *Response, error) {

	path := "/channel/update"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	chatChannel := new(ChatChannelUpdate)
	resp, err := s.client.Do(req, chatChannel)

	chatChannel.CoverUrl = chatChannel.CoverImageUrl

	if err != nil {
		return nil, resp, err
	}

	return chatChannel, resp, nil
}

// Delete the channel that matches the URL
func (s *ChatChannelServiceOp) Delete(channelUrl string) (*Response, error) {

	path := "/channel/delete"

	params := ChatChannelRequest{ChannelUrl: channelUrl}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	var i interface{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// View the channel information and its online members.
func (s *ChatChannelServiceOp) View(channelUrl string) (*ChatChannelView, *Response, error) {

	path := "/channel/view"

	params := ChatChannelRequest{ChannelUrl: channelUrl}
	params.PopulateAuthApiToken(s.client)

	req, err := s.client.NewRequest("POST", path, params)

	view := new(ChatChannelView)
	resp, err := s.client.Do(req, view)

	view.CoverUrl = view.CoverImageUrl

	if err != nil {
		return nil, resp, err
	}

	return view, resp, nil
}

// Send a message to the given channel
func (s *ChatChannelServiceOp) Send(params *ChatChannelMessageRequest) (*Response, error) {

	path := "/channel/send"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	var i interface{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Get values of metadata keys
func (s *ChatChannelServiceOp) GetMetadata(params *ChatChannelMetadataRequest) (map[string]string, *Response, error) {

	path := "/channel/get_metadata"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metadata := map[string]string{}
	resp, err := s.client.Do(req, &metadata)

	if err != nil {
		return nil, resp, err
	}

	return metadata, resp, nil
}

// Set values of metadata keys
func (s *ChatChannelServiceOp) SetMetadata(params *ChatChannelSetMetadataRequest) (map[string]string, *Response, error) {

	path := "/channel/set_metadata"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metadata := map[string]string{}
	resp, err := s.client.Do(req, &metadata)

	if err != nil {
		return nil, resp, err
	}

	return metadata, resp, nil
}

// Get values of metacounter keys
func (s *ChatChannelServiceOp) GetMetacounter(params *ChatChannelMetacounterRequest) (map[string]int, *Response, error) {

	path := "/channel/get_metacounter"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metacounters := map[string]int{}
	resp, err := s.client.Do(req, &metacounters)

	if err != nil {
		return nil, resp, err
	}

	return metacounters, resp, nil
}

// Set values of metacounter keys. All values should be integer
func (s *ChatChannelServiceOp) SetMetacounter(params *ChatChannelSetMetacounterRequest) (map[string]int, *Response, error) {

	path := "/channel/set_metacounter"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metacounters := map[string]int{}
	resp, err := s.client.Do(req, &metacounters)

	if err != nil {
		return nil, resp, err
	}

	return metacounters, resp, nil
}

// Increase values of metacounter keys. All deltas should be integer
func (s *ChatChannelServiceOp) IncreaseMetacounter(params *ChatChannelSetMetacounterRequest) (map[string]int, *Response, error) {

	path := "/channel/incr_metacounter"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metacounters := map[string]int{}
	resp, err := s.client.Do(req, &metacounters)

	if err != nil {
		return nil, resp, err
	}

	return metacounters, resp, nil
}

// Decrease values of metacounter keys. All deltas should be integer
func (s *ChatChannelServiceOp) DecreaseMetacounter(params *ChatChannelSetMetacounterRequest) (map[string]int, *Response, error) {

	path := "/channel/decr_metacounter"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metacounters := map[string]int{}
	resp, err := s.client.Do(req, &metacounters)

	if err != nil {
		return nil, resp, err
	}

	return metacounters, resp, nil
}

// Get message count of the channel
func (s *ChatChannelServiceOp) MessageCount(channelUrl string) (*MessageCount, *Response, error) {

	path := "/channel/message_count"

	params := ChatChannelRequest{ChannelUrl: channelUrl}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	count := new(MessageCount)
	resp, err := s.client.Do(req, count)

	if err != nil {
		return nil, resp, err
	}

	return count, resp, nil
}
