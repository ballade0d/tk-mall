package service

import (
	"context"
	"errors"
	v2 "mall/api/mall/service/v1"
	"mall/app/gateway/internal/data"
	"mall/app/gateway/internal/util"
	"mall/ent"
	"time"
)

type UserService struct {
	v2.UnimplementedUserServiceServer
	userRepo     data.UserRepo
	passwordRepo data.PasswordRepo
}

func NewUserService(userRepo data.UserRepo, passwordRepo data.PasswordRepo) *UserService {
	return &UserService{
		userRepo:     userRepo,
		passwordRepo: passwordRepo,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *v2.GetUserRequest) (*v2.GetUserResponse, error) {
	id := ctx.Value("claims").(*util.Claims).UserId
	usr, err := s.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &v2.GetUserResponse{
		User: &v2.User{
			Id:    int64(usr.ID),
			Name:  usr.Name,
			Email: usr.Email,
		},
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *v2.LoginRequest) (*v2.LoginResponse, error) {
	usr, err := s.userRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	pwd, err := usr.QueryPassword().Only(ctx)
	if err != nil {
		return nil, err
	}
	token, err := util.GenToken(usr, time.Duration(2)*time.Hour)
	refreshToken, err := util.GenRefreshToken(usr, time.Duration(24)*time.Hour)
	if err != nil {
		return nil, err
	}
	if util.Verify(pwd.Password, req.Password) {
		return &v2.LoginResponse{
			Token: &v2.Token{
				Token:        token,
				RefreshToken: refreshToken,
			},
		}, nil
	}
	return nil, nil
}

func (s *UserService) Register(ctx context.Context, req *v2.RegisterRequest) (*v2.RegisterResponse, error) {
	usr, err := s.userRepo.FindUserByEmail(ctx, req.Email)
	var entErr *ent.NotFoundError
	if !errors.As(err, &entErr) {
		return nil, errors.New("user already exists")
	}
	encrypted, err := util.Encrypt(req.Password)
	if err != nil {
		return nil, err
	}
	pwd, err := s.passwordRepo.CreatePassword(ctx, encrypted)
	if err != nil {
		return nil, err
	}
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("name, email and password are required")
	}
	usr, err = s.userRepo.CreateUser(ctx, &ent.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  "user",
	}, pwd)
	if err != nil {
		return nil, err
	}
	token, err := util.GenToken(usr, time.Duration(2)*time.Hour)
	refreshToken, err := util.GenRefreshToken(usr, time.Duration(24)*time.Hour)
	return &v2.RegisterResponse{
		Token: &v2.Token{
			Token:        token,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (s *UserService) RefreshToken(ctx context.Context, req *v2.RefreshTokenRequest) (*v2.RefreshTokenResponse, error) {
	claims, err := util.VerifyJWT(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	if claims.GrantType != "refresh_token" {
		return nil, errors.New("invalid refresh token")
	}
	usr, err := s.userRepo.FindUserByID(ctx, claims.UserId)
	if err != nil {
		return nil, err
	}
	token, err := util.GenToken(usr, time.Duration(2)*time.Hour)
	refreshToken, err := util.GenRefreshToken(usr, time.Duration(24)*time.Hour)
	return &v2.RefreshTokenResponse{
		Token: &v2.Token{
			Token:        token,
			RefreshToken: refreshToken,
		},
	}, nil
}
