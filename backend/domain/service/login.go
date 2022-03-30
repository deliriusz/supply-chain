package domain

import "rafal-kalinowski.pl/domain/model"

type LoginService interface {
	GetLoginChallenge(*model.LoginChallenge) (*model.LoginChallenge, error)
	Login(*model.LoginChallenge) (*model.Login, error)
	Logout(*model.Login) error
	GetSessionById(string) (*model.Login, error)
}

type LoginRepository interface {
	GetLoginChallenge(*model.LoginChallenge) (*model.LoginChallenge, error)
	Login(*model.LoginChallenge) (*model.Login, error)
	Logout(*model.Login) error
	GetSessionById(string) (*model.Login, error)
}
