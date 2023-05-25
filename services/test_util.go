package services

import "ropc-service/model"

var (
	// test user details
	rightUsername       = "right_username"
	rightPassword       = "right_password"
	wrongPassword       = "wrong_password"
	wrongUsername       = "wrong_username"
	hashedRightPassword = "$2a$12$JadjhGXumDBw.8X9o0.EaeNkIDaeGtmHkAmxfgrqApaFT0t.ZVrm."

	rightTestUser = model.User{
		Username: rightUsername,
		Password: rightPassword,
	}

	wrongTestUser = model.User{
		Username: wrongUsername,
		Password: wrongPassword,
	}

	// test client details
	rightClientId           = "right_clientId"
	rightClientSecret       = "right_clientSecret"
	wrongClientId           = "wrong_clientId"
	wrongClientSecret       = "wrong_clientSecret"
	hashedRightClientSecret = "$2a$12$0TUqBhM9DdBw980nTxz1EuL3eM/jQQSABDVuO6/lrCsjuUCOCFdxy"

	rightClient = model.Client{
		ClientId:     rightClientId,
		ClientSecret: rightClientSecret,
	}

	wrongClient = model.Client{
		ClientId:     wrongClientId,
		ClientSecret: wrongClientSecret,
	}
)
