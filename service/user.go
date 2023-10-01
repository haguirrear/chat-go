package service

import (
	"chat/entities"
	"chat/repository"
	"chat/settings"
	"chat/utils"
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{repository: r}
}

func (s *UserService) CreateUser(ctx context.Context, user CreateUserReq) (*CreateUserRes, error) {
	foundUser, err := s.repository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if foundUser != nil {
		return nil, ErrorUserAlreadyCreated
	}

	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	u := entities.User{
		Email:    user.Email,
		Password: hashPass,
	}
	err = s.repository.CreateUser(ctx, &u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		Id:        u.Id,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return res, nil

}

type JWTClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (s *UserService) LoginUser(ctx context.Context, credentials LoginUserReq) (*LoginUserRes, error) {
	foundUser, err := s.repository.GetUserByEmail(ctx, credentials.Email)
	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, ErrorEmailOrPasswordIncorrect
	}
	isCorrect := utils.CheckPasswordIsCorrect(credentials.Password, foundUser.Password)
	if !isCorrect {
		return nil, ErrorEmailOrPasswordIncorrect
	}

	jwtObj := jwt.NewWithClaims(
		jwt.SigningMethodHS256, JWTClaims{
			Id:    foundUser.Id,
			Email: foundUser.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    strconv.Itoa(foundUser.Id),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			},
		},
	)

	token, err := jwtObj.SignedString(settings.GetSettings().SecretKey)
	if err != nil {
		return nil, err
	}

	return &LoginUserRes{
		Id:          foundUser.Id,
		Email:       foundUser.Email,
		AccessToken: token,
	}, nil
}
