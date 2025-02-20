package service

import (
	"context"
	"errors"
	v2 "mall/api/mall/service/v1"
	"mall/app/auth/internal/data"
	"mall/app/auth/internal/util"
	"mall/ent"
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
	return nil, nil
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
	token, err := util.GenerateJWT(usr, 2)
	refreshToken, err := util.GenerateJWT(usr, 24)
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
		return nil, err
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
	}, pwd)
	if err != nil {
		return nil, err
	}
	token, err := util.GenerateJWT(usr, 2)
	refreshToken, err := util.GenerateJWT(usr, 24)
	return &v2.RegisterResponse{
		Token: &v2.Token{
			Token:        token,
			RefreshToken: refreshToken,
		},
	}, nil
}
