package sendbird

type UserRequest struct {
	RequestDefaults
	Id               string `json:"id,omitempty"`
	Nickname         string `json:"nickname,omitempty"`
	ImageUrl         string `json:"image_url,omitempty"`
	IssueAccessToken bool   `json:"issue_access_token,omitempty"`
}
type BlockRequest struct {
	RequestDefaults
	Id       string `json:"id"`
	TargetId string `json:"target_id,omitempty"`
}
type DeactivateRequest struct {
	RequestDefaults
	Id string `json:"id"`
}

type User struct {
	Id          string `json:"id,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	Nickname    string `json:"nickname"`
	Picture     string `json:"picture"`
	AccessToken string `json:"access_token,omitempty"`
	IsActive    bool   `json:"is_active,omitempty"`
}

// UserService is an interface for interfacing with the User
// endpoints of the Sendbird API
type UserService interface {
	Create(*UserRequest) (*User, *Response, error)
	Update(params *UserRequest) (*User, *Response, error)
	Auth(params *UserRequest) (*User, *Response, error)
	Block(params *BlockRequest) (*Response, error)
	UnBlock(params *BlockRequest) (*Response, error)
	Deactivate(params *DeactivateRequest) (*Response, error)
}

// UserServiceOp handles communication with the User related methods of
// the Sendbird API.
type UserServiceOp struct {
	client *SendbirdClient
}

var _ UserService = &UserServiceOp{}

// Create creates a new user account using id / nickname / profile image combination
func (s *UserServiceOp) Create(params *UserRequest) (*User, *Response, error) {

	path := "/user/create"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	createdUser := new(User)
	resp, err := s.client.Do(req, createdUser)

	// log.Info(resp)

	if err != nil {
		return nil, resp, err
	}

	return createdUser, resp, nil
}

// Update updates a user
func (s *UserServiceOp) Update(params *UserRequest) (*User, *Response, error) {

	path := "/user/update"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	updatedUser := new(User)
	resp, err := s.client.Do(req, updatedUser)

	if err != nil {
		return nil, resp, err
	}

	return updatedUser, resp, nil
}

// Auth retrieve a user's information
func (s *UserServiceOp) Auth(params *UserRequest) (*User, *Response, error) {

	path := "/user/auth"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	user := new(User)
	resp, err := s.client.Do(req, user)

	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// Block blocks a target user
func (s *UserServiceOp) Block(params *BlockRequest) (*Response, error) {

	path := "/user/block"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	var i interface{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Unblock unblocks a target user.
func (s *UserServiceOp) UnBlock(params *BlockRequest) (*Response, error) {

	path := "/user/unblock"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	var i interface{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivate deactivates a user.
func (s *UserServiceOp) Deactivate(params *DeactivateRequest) (*Response, error) {

	path := "/user/deactivate"
	params.PopulateAuthApiToken(s.client)
	req, err := s.client.NewRequest("POST", path, *params)

	var i interface{}
	resp, err := s.client.Do(req, i)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
