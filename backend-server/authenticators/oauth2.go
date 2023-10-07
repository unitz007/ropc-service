package authenticators

type Oauth2 interface {
	Authenticate(grantType string)
}

type GrantType interface {
}
