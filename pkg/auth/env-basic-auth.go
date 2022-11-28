package auth

import "errors"

type (
	Config struct {
		Login    string
		Password string
	}

	EnvBasicAuth struct {
		config Config
	}
)

func NewEnvBasicAuth(cfg Config) *EnvBasicAuth {
	return &EnvBasicAuth{
		config: cfg,
	}
}

func (a *EnvBasicAuth) Authentication(username, password string, hasAuth bool) error {
	if hasAuth && username == a.config.Login && password == a.config.Password {
		return nil
	}
	return errors.New("unauthorized user")
}
