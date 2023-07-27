package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	pb "kratos-demo/api/realworld/v1"
	"kratos-demo/internal/biz"
	"kratos-demo/utils"
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

func setToken(u *pb.UserReply) error {
	token, err := utils.GenToken(u.User.Username)
	if err != nil {
		return err
	}
	u.User.Token = token
	return nil
}

func setDefault(u *pb.UserReply) {
	if u.User.Image == "" {
		u.User.Image = "https://static.productionready.io/images/smiley-cyrus.jpg"
	}
	if u.User.Bio == "" {
		u.User.Bio = "I work at statefarm"
	}
}

func (r *userRepo) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	var username string
	if md, ok := metadata.FromServerContext(ctx); ok {
		username = md.Get("username")
		fmt.Println("username: ", username)
	}

	ur := &pb.UserReply{}
	err := r.data.db.Model(&biz.User{}).Where("email = ? AND password = ?", req.User.Email, req.User.Password).Scan(&ur.User).Error
	if err != nil {
		return nil, err
	}

	setDefault(ur)

	err = setToken(ur)
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func (r *userRepo) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	ur := &pb.UserReply{}
	err := r.data.db.Create(
		&biz.User{
			Username: req.User.Username,
			Email:    req.User.Email,
			Password: req.User.Password,
		}).Scan(&ur.User).Error
	if err != nil {
		return nil, err
	}

	setDefault(ur)

	err = setToken(ur)
	if err != nil {
		return nil, err
	}

	return ur, nil
}
