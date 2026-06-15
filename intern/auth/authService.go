package auth

import (
	"context"
	"mini-ozon/intern/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repositories.UserRepository
}

func CreateAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
func (a AuthService) Login(email, password string, ctx context.Context) (string, error) {
	user, err := a.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}
	token, err := GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}
