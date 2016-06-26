package sendbird

// AdminService is an interface for interfacing with the Admin
// endpoints of the Sendbird API
type AdminService interface {
	BroadcastMessage(params *BroadcastMessageRequest) (*Response, error)
	ReadMessages(params *ReadMessagesRequest) ([]AdminMessage, *Response, error)
	DeleteMessage(messageId string) (*DeleteMessage, *Response, error)
	ListMessagingChannels(userId string) ([]AdminMessagingChannel, *Response, error)
	MuteAllChannels(userId string) (*Response, error)
	Mute(params *MuteRequest) ([]string, *Response, error)
	UnMuteAllChannels(userId string) (*Response, error)
	UnMute(params *UnMuteRequest) ([]string, *Response, error)
	MuteList(channelUrls []string) ([]string, *Response, error)
	ConcurrentUserCount() (*ConcurrentUserCount, *Response, error)
	MemberCountInChannel(channelUrl string) (*ChannelMemberCount, *Response, error)
}

// AdminServiceOp handles communication with the AdminService related methods of
// the Sendbird API.
type AdminServiceOp struct {
	client *SendbirdClient
}

var _ AdminService = &AdminServiceOp{}

type BroadcastMessageRequest struct {
	RequestDefaults
	ChannelUrls []string `json:"channel_url,omitempty"` // List of channel urls
	Message     string   `json:"message,omitempty"`     // Message to be sent
	Persistent  bool     `json:"persistent,omitempty"`  // Save the message if True
	Data        string   `json:"data,omitempty"`        // Custom data field
}
type ReadMessagesRequest struct {
	RequestDefaults
	ChannelUrl    string   `json:"channel_url,omitempty"`     // (Optional) Channel url
	TargetUserIds []string `json:"target_user_ids,omitempty"` // (Optional) Target messaging channel with given two users.
	Limit         int      `json:"limit"`                     // A number of messages to read (default: 50)
	MessageId     int64    `json:"message_id"`                // Read a given limit of messages before this id. Set to zero to read messages from the last.
}
type MuteRequest struct {
	RequestDefaults
	Id          string   `json:"id"` // User ID
	ChannelUrls []string `json:"channel_urls"`
	IsSoftMute  bool     `json:"is_soft_mute"` // Channel mute has two modes; hard mute and soft mute. In hard mute mode, messages from muted users are blocked in server. In soft mute mode, messages from muted users are sent to "onMuteMessagesReceived" callback in Client SDK. In most cases, what you need is "hard mute" mode
}
type UnMuteRequest struct {
	RequestDefaults
	Id          string   `json:"id"` // User ID
	ChannelUrls []string `json:"channel_urls"`
}

type AdminMessage struct {
	Id        string           `json:"id"`         // Sender's ID
	Nickname  string           `json:"nickanme"`   // Sender's name
	MessageId int64            `json:"message_id"` // Message ID
	Timestamp int64            `json:"timestamp"`  // Epoch timestamp in milliseconds
	Message   string           `json:"message"`    // Chat message
	File      AdminMessageFile `json:"file"`       // File Info Object... Empty file info object when message is text message
}
type AdminMessageFile struct {
	Url    string `json:"url"`    // File url
	Custom string `json:"custom"` // User custom field
	Type   string `json:"type"`   // File type
	Name   string `json:"name"`   // File name
	Size   int64  `json:"size"`   // File size
}
type DeleteMessage struct {
	AppId int64 `json:"app_id"`
	MsgId int64 `json:"msg_id"`
}
type AdminMessagingChannel struct {
	ChannelUrl         string   `json:"channel_url"`          // Channel URL
	UnreadMessageCount int      `json:"unread_message_count"` // Unread message count
	LastMessage        string   `json:"last_message"`         // Last message
	LastMessageTS      int64    `json:"last_message_ts"`      // Last message timestamp, 0 if last message is empty.
	Members            []Member `json:"members"`
}
type ConcurrentUserCount struct {
	Count int `json:"count"` // Concurrent chatting user count
}
type ChannelMemberCount struct {
	AccumulatedMemberCount int `json:"accumulated_member_count"` // number of members joined the channel at least once
	OnlineMemberCount      int `json:"online_member_count"`      // number of members currently online in the channel
	MemberCount            int `json:"member_count"`             // number of members in the channel
}

// BroadcastMessages broadcasts an admin message to the channels
// Please note that this feature is not available on Free and Sprout plans.
func (s *AdminServiceOp) BroadcastMessage(params *BroadcastMessageRequest) (*Response, error) {

	path := "/admin/broadcast_message"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	var i interface{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ReadMessages reads messages from the target channel
// Either channel_url or target_user_ids must be set
func (s *AdminServiceOp) ReadMessages(params *ReadMessagesRequest) ([]AdminMessage, *Response, error) {

	path := "/admin/read_messages"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	messages := []AdminMessage{}
	resp, err := s.client.Do(req, &messages)

	if err != nil {
		return nil, resp, err
	}

	return messages, resp, nil
}

// DeleteMessage deletes a message from your application
func (s *AdminServiceOp) DeleteMessage(messageId string) (*DeleteMessage, *Response, error) {

	path := "/admin/delete_message"

	params := struct {
		RequestDefaults
		MsgId string `json:"msg_id"`
	}{
		MsgId: messageId,
	}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	message := new(DeleteMessage)
	resp, err := s.client.Do(req, message)

	if err != nil {
		return nil, resp, err
	}

	return message, resp, nil
}

// listMessagingChannels lists messaging channels of the target user
func (s *AdminServiceOp) ListMessagingChannels(userId string) ([]AdminMessagingChannel, *Response, error) {

	path := "/admin/list_messaging_channels"

	params := struct {
		RequestDefaults
		Id string `json:"id"`
	}{
		Id: userId,
	}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	channels := []AdminMessagingChannel{}
	resp, err := s.client.Do(req, &channels)

	if err != nil {
		return nil, resp, err
	}

	return channels, resp, nil
}

// MuteAllChannels mutes a user in all channels in your application. Prohibit a user from sending messages at all.
func (s *AdminServiceOp) MuteAllChannels(userId string) (*Response, error) {

	path := "/admin/mute"

	params := struct {
		RequestDefaults
		Id string `json:"id"`
	}{
		Id: userId,
	}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	var i interface{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Mute mutes a user in some channels in your application
func (s *AdminServiceOp) Mute(params *MuteRequest) ([]string, *Response, error) {

	path := "/admin/mute"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	channels := []string{}
	resp, err := s.client.Do(req, &channels)

	if err != nil {
		return nil, resp, err
	}

	return channels, resp, nil
}

// UnMuteAllChannels unmutes a user to start sending messages in your application. This doesn't unmute channel-wide mute
func (s *AdminServiceOp) UnMuteAllChannels(userId string) (*Response, error) {

	path := "/admin/unmute"

	params := struct {
		RequestDefaults
		Id string `json:"id"`
	}{
		Id: userId,
	}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	var i interface{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UnMute unmutes a user in some channels in your application
func (s *AdminServiceOp) UnMute(params *UnMuteRequest) ([]string, *Response, error) {

	path := "/admin/unmute"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	channels := []string{}
	resp, err := s.client.Do(req, &channels)

	if err != nil {
		return nil, resp, err
	}

	return channels, resp, nil
}

// MuteList gets the list of muted user ids in application-wide mute
func (s *AdminServiceOp) MuteList(channelUrls []string) ([]string, *Response, error) {

	path := "/admin/mute_list"

	params := struct {
		RequestDefaults
		ChannelUrls []string `json:"channel_urls"`
	}{
		ChannelUrls: channelUrls,
	}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	userIds := []string{}
	resp, err := s.client.Do(req, &userIds)

	if err != nil {
		return nil, resp, err
	}

	return userIds, resp, nil
}

// ConcurrentUserCount gets the list of muted user ids in application-wide mute
func (s *AdminServiceOp) ConcurrentUserCount() (*ConcurrentUserCount, *Response, error) {

	path := "/admin/ccu_count"

	params := struct{ RequestDefaults }{}
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	count := new(ConcurrentUserCount)
	resp, err := s.client.Do(req, count)

	if err != nil {
		return nil, resp, err
	}

	return count, resp, nil
}

// MemberCountInChannel gets the list of muted user ids in application-wide mute
func (s *AdminServiceOp) MemberCountInChannel(channelUrl string) (*ChannelMemberCount, *Response, error) {

	path := "/admin/member_count"

	params := struct {
		RequestDefaults
		ChannelUrl string `json:"channel_url"`
	}{
		ChannelUrl: channelUrl,
	}

	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, params)

	count := new(ChannelMemberCount)
	resp, err := s.client.Do(req, count)

	if err != nil {
		return nil, resp, err
	}

	return count, resp, nil
}
