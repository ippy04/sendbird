## Golang Sendbird Server API Client

[![Build Status](https://travis-ci.org/ippy04/sendbird.svg?branch=master)](https://travis-ci.org/ippy04/sendbird)
[![Coverage Status](https://coveralls.io/repos/github/ippy04/sendbird/badge.svg?branch=master)](https://coveralls.io/github/ippy04/sendbird?branch=master)

You can view the Sendbird Server API docs here: [https://sendbird.gitbooks.io/sendbird-server-api/content/en/](https://sendbird.gitbooks.io/sendbird-server-api/content/en/)

### Usage
1. Create a Sendbird Application at [https://dashboard.sendbird.com/](https://dashboard.sendbird.com/)
2. Run `go get github.com/ippy04/sendbird`

```go

import "github.com/ippy04/sendbird"

sb := sendbird.NewClient(SENDBIRD_APP_ID, SENDBIRD_API_TOKEN, nil)

```

### Examples
To create a new user:

```go
import "github.com/ippy04/sendbird"

func createUser() {
	sb := sendbird.NewClient(SENDBIRD_APP_ID, SENDBIRD_API_TOKEN, nil)

	params := sendbird.UserRequest{
		Id:               "123456",
		Nickname:         "nickname",
		ImageUrl:         "http://sendbird.com/picture_url",
		IssueAccessToken: true,
	}

	user, _, err := sb.User.Create(params)
	if err != nil {
		log.Fatal("User not created")
	}
}
```

*See tests for more examples*
