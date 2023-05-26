package services

import (
	"ropc-service/model/entities"
)

var (
	// test user details
	rightUsername       = "right_username"
	rightPassword       = "right_password"
	wrongPassword       = "wrong_password"
	wrongUsername       = "wrong_username"
	hashedRightPassword = "$2a$12$JadjhGXumDBw.8X9o0.EaeNkIDaeGtmHkAmxfgrqApaFT0t.ZVrm."

	rightTestUser = entities.User{
		Username: rightUsername,
		Password: rightPassword,
	}

	wrongTestUser = entities.User{
		Username: wrongUsername,
		Password: wrongPassword,
	}

	// test client details
	rightClientId           = "right_clientId"
	rightClientSecret       = "right_clientSecret"
	wrongClientId           = "wrong_clientId"
	wrongClientSecret       = "wrong_clientSecret"
	hashedRightClientSecret = "$2a$12$0TUqBhM9DdBw980nTxz1EuL3eM/jQQSABDVuO6/lrCsjuUCOCFdxy"

	rightClient = entities.Client{
		ClientId:     rightClientId,
		ClientSecret: rightClientSecret,
	}

	wrongClient = entities.Client{
		ClientId:     wrongClientId,
		ClientSecret: wrongClientSecret,
	}
)
