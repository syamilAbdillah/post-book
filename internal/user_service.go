package internal

import (
	"context"

	"github.com/syamilAbdillah/post-book/pkg/webrpc"
)

type UserService struct {}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(ctx context.Context, credentials *webrpc.Credentials) (*webrpc.User, error) {
	return nil, nil
}

func (s *UserService) Register(ctx context.Context, data *webrpc.UserData) (*webrpc.User, map[string]string, error) {
	return nil, nil, nil
}

func (s *UserService) IsAuthenticated(ctx context.Context) (*webrpc.User, error) {
	return nil, nil
}