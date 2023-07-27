package service

import (
	"context"
	"kratos-demo/internal/biz"

	pb "kratos-demo/api/realworld/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{
		uc: uc,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	u, err := s.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	u, err := s.uc.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//func (s *UserService) GetCurrentUser(ctx context.Context, req *pb.EmptyRequest) (*pb.UserReply, error) {
//	return &pb.UserReply{}, nil
//}
//func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
//	return &pb.UserReply{}, nil
//}
//func (s *UserService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.Profile, error) {
//	return &pb.Profile{}, nil
//}
//func (s *UserService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.Profile, error) {
//	return &pb.Profile{}, nil
//}
//func (s *UserService) UnfollowUser(ctx context.Context, req *pb.UnfollowUserRequest) (*pb.Profile, error) {
//	return &pb.Profile{}, nil
//}
