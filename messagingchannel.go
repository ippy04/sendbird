package sendbird

type MessagingChannelRequest struct {
	RequestDefaults
	Name     string `json:"name,omitempty"`      // (Optional) Topic
	IsGroup  bool   `json:"is_group,omitempty"`  // // weather 1 on 1 or Group chat
	CoverUrl string `json:"cover_url,omitempty"` // (Optional) channel cover image url
	Data     string `json:"data,omitempty"`      // (Optional) custom data
}
type MessagingChannelUpdateRequest struct {
	RequestDefaults
	Name       string `json:"name,omitempty"`        // (Optional) Topic
	IsGroup    bool   `json:"is_group,omitempty"`    // // weather 1 on 1 or Group chat
	CoverUrl   string `json:"cover_url,omitempty"`   // (Optional) channel cover image url
	Data       string `json:"data,omitempty"`        // (Optional) custom data
	ChannelUrl string `json:"channel_url,omitempty"` // Target Channel Url
}
type MessagingChannelInviteRequest struct {
	RequestDefaults
	ChannelUrl string   `json:"channel_url"`
	UserIds    []string `json:"user_ids"`
}
type MessagingChannelHideRequest struct {
	RequestDefaults
	Id         string `json:"ids"`         // User ID
	ChannelUrl string `json:"channel_url"` // Channel URL
}
type MessagingChannelLeaveRequest struct {
	RequestDefaults
	ChannelUrl string   `json:"channel_url"`
	UserIds    []string `json:"user_ids"`
}
type MessagingChannelMetadataRequest struct {
	RequestDefaults
	ChannelUrl string   `json:"channel_url"` // Channel URL
	Keys       []string `json:"keys"`        //  Key names

}
type MessagingChannelSetMetadataRequest struct {
	RequestDefaults
	ChannelUrl string            `json:"channel_url"` // Channel URL
	Data       map[string]string `json:"data"`
}
type MessagingChannelMetacounterRequest struct {
	RequestDefaults
	ChannelUrl string   `json:"channel_url"` // Channel URL
	Keys       []string `json:"keys"`        //  Key names
}
type MessagingChannelSetMetacounterRequest struct {
	RequestDefaults
	ChannelUrl string         `json:"channel_url"` // Channel URL
	Data       map[string]int `json:"data"`
}

type MessagingChannelResponse struct {
	Channel MessagingChannel `json:"channel"`
}
type MessagingChannel struct {
	ChannelUrl string `json:"channel_url"`
	Data       string `json:"data"`
	Name       string `json:"name"`
	IsGroup    bool   `json:"is_group"` // // weather 1 on 1 or Group chat
	CoverUrl   string `json:"cover_url"`
	// CreatedAt int64 `json:"created_at"` // Epoch timestamp in milliseconds
}
type MessagingChannelUrl struct {
	Channel ChannelUrl `json:"channel"`
}
type ChannelUrl struct {
	ChannelUrl string `json:"channel_url"`
}
type MessagingChannelView struct {
	ChannelUrl    string   `json:"channel_url"`
	LastMessage   string   `json:"last_message"`
	LastMessageTS int64    `json:"last_message_ts"`
	CreatedAt     int64    `json:"created_at"`
	Members       []Member `json:"members"`
}

// MessagingChannelService is an interface for interfacing with the Messaging Channel
// endpoints of the Sendbird API
type MessagingChannelService interface {
	Create(*MessagingChannelRequest) (*MessagingChannel, *Response, error)
	Update(params *MessagingChannelUpdateRequest) (*MessagingChannel, *Response, error)
	Delete(channelUrl string) (*MessagingChannelUrl, *Response, error)
	Invite(params *MessagingChannelInviteRequest) (*MessagingChannelUrl, *Response, error)
	Hide(params *MessagingChannelHideRequest) (*MessagingChannelUrl, *Response, error)
	Leave(params *MessagingChannelLeaveRequest) (*MessagingChannelUrl, *Response, error)
	View(channelUrl string) (*MessagingChannelView, *Response, error)
	GetMetadata(params *MessagingChannelMetadataRequest) (map[string]string, *Response, error)
	SetMetadata(params *MessagingChannelSetMetadataRequest) (map[string]string, *Response, error)
	GetMetacounter(params *MessagingChannelMetacounterRequest) (map[string]int, *Response, error)
	SetMetacounter(params *MessagingChannelSetMetacounterRequest) (map[string]int, *Response, error)
	IncreaseMetacounter(params *MessagingChannelSetMetacounterRequest) (map[string]int, *Response, error)
	DecreaseMetacounter(params *MessagingChannelSetMetacounterRequest) (map[string]int, *Response, error)
	MessageCount(channelUrl string) (*MessageCount, *Response, error)
}

// MessagingChannelServiceOp handles communication with the Messaging Channel related methods of
// the Sendbird API.
type MessagingChannelServiceOp struct {
	client *SendbirdClient
}

var _ MessagingChannelService = &MessagingChannelServiceOp{}

// Create creates a group messaging channel using given name
func (s *MessagingChannelServiceOp) Create(params *MessagingChannelRequest) (*MessagingChannel, *Response, error) {

	path := "/messaging/create"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	messagingChannel := new(MessagingChannelResponse)
	resp, err := s.client.Do(req, messagingChannel)

	if err != nil {
		return nil, resp, err
	}

	return &messagingChannel.Channel, resp, nil
}

// Update updates a group messaging channel using given channel_url
func (s *MessagingChannelServiceOp) Update(params *MessagingChannelUpdateRequest) (*MessagingChannel, *Response, error) {

	path := "/messaging/update"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	messagingChannel := new(MessagingChannelResponse)
	resp, err := s.client.Do(req, messagingChannel)

	if err != nil {
		return nil, resp, err
	}

	return &messagingChannel.Channel, resp, nil
}

// Delete deletes a group messaging channel using given channel_url
func (s *MessagingChannelServiceOp) Delete(channelUrl string) (*MessagingChannelUrl, *Response, error) {

	path := "/messaging/delete"

	params := MessagingChannelUpdateRequest{ChannelUrl: channelUrl}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	deleted := new(MessagingChannelUrl)
	resp, err := s.client.Do(req, deleted)

	if err != nil {
		return nil, resp, err
	}

	return deleted, resp, nil
}

// Invite invites users to the channel
func (s *MessagingChannelServiceOp) Invite(params *MessagingChannelInviteRequest) (*MessagingChannelUrl, *Response, error) {

	path := "/messaging/invite"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	messagingChannelUrl := new(MessagingChannelUrl)
	resp, err := s.client.Do(req, messagingChannelUrl)

	if err != nil {
		return nil, resp, err
	}

	return messagingChannelUrl, resp, nil
}

// Hide hides a messaging channel from the user messaging channel list
func (s *MessagingChannelServiceOp) Hide(params *MessagingChannelHideRequest) (*MessagingChannelUrl, *Response, error) {

	path := "/messaging/hide"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	messagingChannelUrl := new(MessagingChannelUrl)
	resp, err := s.client.Do(req, messagingChannelUrl)

	if err != nil {
		return nil, resp, err
	}

	return messagingChannelUrl, resp, nil
}

// Leave leaves a messaging channel
func (s *MessagingChannelServiceOp) Leave(params *MessagingChannelLeaveRequest) (*MessagingChannelUrl, *Response, error) {

	path := "/messaging/leave"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	messagingChannelUrl := new(MessagingChannelUrl)
	resp, err := s.client.Do(req, messagingChannelUrl)

	if err != nil {
		return nil, resp, err
	}

	return messagingChannelUrl, resp, nil

}

// View views information of the channel
func (s *MessagingChannelServiceOp) View(channelUrl string) (*MessagingChannelView, *Response, error) {

	path := "/messaging/view"

	params := MessagingChannelUpdateRequest{ChannelUrl: channelUrl}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	view := new(MessagingChannelView)
	resp, err := s.client.Do(req, view)

	if err != nil {
		return nil, resp, err
	}

	return view, resp, nil
}

// GetMetadata gets values of metadata keys
func (s *MessagingChannelServiceOp) GetMetadata(params *MessagingChannelMetadataRequest) (map[string]string, *Response, error) {

	path := "/messaging/get_metadata"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metadata := map[string]string{}
	resp, err := s.client.Do(req, &metadata)

	if err != nil {
		return nil, resp, err
	}

	return metadata, resp, nil
}

// SetMetadata sets values of metadata keys
func (s *MessagingChannelServiceOp) SetMetadata(params *MessagingChannelSetMetadataRequest) (map[string]string, *Response, error) {

	path := "/messaging/set_metadata"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metadata := map[string]string{}
	resp, err := s.client.Do(req, &metadata)

	if err != nil {
		return nil, resp, err
	}

	return metadata, resp, nil
}

// GetMetacounter gets values of metacounter keys
func (s *MessagingChannelServiceOp) GetMetacounter(params *MessagingChannelMetacounterRequest) (map[string]int, *Response, error) {

	path := "/messaging/get_metacounter"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metacounter := map[string]int{}
	resp, err := s.client.Do(req, &metacounter)

	if err != nil {
		return nil, resp, err
	}

	return metacounter, resp, nil
}

// SetMetacounter sets values of metacounter keys. All values should be integer
func (s *MessagingChannelServiceOp) SetMetacounter(params *MessagingChannelSetMetacounterRequest) (map[string]int, *Response, error) {

	path := "/messaging/set_metacounter"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metacounter := map[string]int{}
	resp, err := s.client.Do(req, &metacounter)

	if err != nil {
		return nil, resp, err
	}

	return metacounter, resp, nil
}

// IncreaseMetacounter increases values of metacounter keys. All deltas should be integer
func (s *MessagingChannelServiceOp) IncreaseMetacounter(params *MessagingChannelSetMetacounterRequest) (map[string]int, *Response, error) {

	path := "/messaging/incr_metacounter"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metacounter := map[string]int{}
	resp, err := s.client.Do(req, &metacounter)

	if err != nil {
		return nil, resp, err
	}

	return metacounter, resp, nil
}

// DecreaseMetacounter decreases values of metacounter keys. All deltas should be integer
func (s *MessagingChannelServiceOp) DecreaseMetacounter(params *MessagingChannelSetMetacounterRequest) (map[string]int, *Response, error) {

	path := "/messaging/decr_metacounter"

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	metacounter := map[string]int{}
	resp, err := s.client.Do(req, &metacounter)

	if err != nil {
		return nil, resp, err
	}

	return metacounter, resp, nil
}

// MessageCount gets message count of the channel
func (s *MessagingChannelServiceOp) MessageCount(channelUrl string) (*MessageCount, *Response, error) {

	path := "/messaging/message_count"

	params := MessagingChannelUpdateRequest{ChannelUrl: channelUrl}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	count := new(MessageCount)
	resp, err := s.client.Do(req, count)

	if err != nil {
		return nil, resp, err
	}

	return count, resp, nil
}
