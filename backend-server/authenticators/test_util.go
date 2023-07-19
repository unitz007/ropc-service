package authenticators

import (
	"ropc-service/model/entities"
)

var (
	// test user details
	RightUsername       = "right_username"
	RightPassword       = "right_password"
	WrongPassword       = "wrong_password"
	WrongUsername       = "wrong_username"
	hashedRightPassword = "$2a$12$JadjhGXumDBw.8X9o0.EaeNkIDaeGtmHkAmxfgrqApaFT0t.ZVrm."

	RightTestUser = entities.User{
		Username: RightUsername,
		Password: RightPassword,
	}

	WrongTestUser = entities.User{
		Username: WrongUsername,
		Password: WrongPassword,
	}

	// test client details
	rightClientId           = "right_clientId"
	rightClientSecret       = "right_clientSecret"
	WrongClientId           = "wrong_clientId"
	WrongClientSecret       = "wrong_clientSecret"
	hashedRightClientSecret = "$2a$12$0TUqBhM9DdBw980nTxz1EuL3eM/jQQSABDVuO6/lrCsjuUCOCFdxy"

	_ = entities.Client{
		ClientId:     rightClientId,
		ClientSecret: rightClientSecret,
	}

	WrongClient = entities.Client{
		ClientId:     WrongClientId,
		ClientSecret: WrongClientSecret,
	}
)
