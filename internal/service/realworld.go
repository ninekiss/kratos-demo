package service

import (
	"context"

	v1 "kratos-demo/api/realworld/v1"
)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer
}

func NewRealWorldService() *RealWorldService {
	return &RealWorldService{}
}

func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.User, error) {
	return &v1.User{
		User: &v1.User_Data{
			Email:    "aa@niki.com",
			Token:    "123456",
			Username: "niki",
			Bio:      "I work at statefarm",
			Image:    "https://static.productionready.io/images/smiley-cyrus.jpg",
		},
	}, nil
}
