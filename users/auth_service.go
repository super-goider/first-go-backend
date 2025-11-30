package users

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo UserRepo
}

func NewAuthService(repo UserRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req UserCreateRequest) (User, error) {
	// 1. базовые проверки
	if req.Login == "" {
		return User{}, fmt.Errorf("login is empty")
	}

	// 2. проверка на занятость логина
	_, isExist, err := s.repo.GetByLogin(req.Login)
	if err != nil {
		return User{}, err
	}
	if isExist {
		return User{}, fmt.Errorf("login already in use")
	}

	// 3. хэшируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	// 4. создаём User
	u := User{
		Login:        req.Login,
		Email:        req.Email,
		PasswordHash: string(hash),
	}

	// 5. сохраняем в репо
	created, err := s.repo.Add(u)
	if err != nil {
		return User{}, err
	}

	return created, nil
}

func (s *AuthService) Login(req UserLoginRequest) (User, error) {
	if req.Login == "" {
		return User{}, fmt.Errorf("login is empty")
	}

	// 2. проверка на занятость логина
	u, found, err := s.repo.GetByLogin(req.Login)
	if err != nil {
		return User{}, err
	}
	if !found {
		return User{}, fmt.Errorf("invalid login or password")
	}

	newErr := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)) // первый аргумент это хеш, второй пароль
	if newErr != nil {
		return User{}, fmt.Errorf("invalid login or password")
	}

	return u, nil
}
