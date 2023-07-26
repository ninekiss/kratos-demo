package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-demo/api/realworld/v1"
	"kratos-demo/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	fmt.Println("login user: ", req.User.Email, req)
	return &pb.UserReply{
		User: &pb.UserReply_Data{
			Username: "niki",
			Email:    req.User.Email,
			Bio:      "I work at statefarm",
			Image:    "https://static.productionready.io/images/smiley-cyrus.jpg",
		},
	}, nil
}
