package domain

import "rafal-kalinowski.pl/domain/model"

type loginService struct {
	repository LoginRepository
}

// GetLoginChallenge implements LoginService
func (s *loginService) GetLoginChallenge(login *model.LoginChallenge) (*model.LoginChallenge, error) {
	return s.repository.GetLoginChallenge(login)
}

// Login implements LoginService
func (s *loginService) Login(login *model.LoginChallenge) (*model.Login, error) {
	return s.repository.Login(login)
}

// Logout implements LoginService
func (s *loginService) Logout(login *model.Login) error {
	return s.repository.Logout(login)
}

func NewLoginService(repository LoginRepository) LoginService {
	return &loginService{
		repository,
	}
}
