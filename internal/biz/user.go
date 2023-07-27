package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	pb "kratos-demo/api/realworld/v1"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;unique" json:"username"`
	Email    string `gorm:"column:email;unique" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Bio      string `gorm:"column:bio" json:"bio"`
	Image    string `gorm:"column:image" json:"image"`
}

type UserRepo interface {
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error)
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error)
	//ListUser(ctx context.Context) ([]*User, error)
	//GetUser(ctx context.Context, email string) (*User, error)
	//CreateUser(ctx context.Context, user *User) error
	//UpdateUser(ctx context.Context, email string, user *User) error
	//DeleteUser(ctx context.Context, email string) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	uc.log.WithContext(ctx).Infof("login user: %s", req.User.Email)
	return uc.repo.Login(ctx, req)
}

func (uc *UserUsecase) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	uc.log.WithContext(ctx).Infof("login user: %s", req.User.Email)
	return uc.repo.Register(ctx, req)
}
