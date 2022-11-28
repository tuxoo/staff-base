package auth

type BasicAuth interface {
	Authentication(username, password string, hasAuth bool) error
}
