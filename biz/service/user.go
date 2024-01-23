package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"ulog-backend/biz/dal"
	"ulog-backend/biz/model"
	"ulog-backend/pkg/encrypt"
)

var (
	u = dal.Q.User
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{
		ctx: ctx,
		c:   c,
	}
}

func (userService UserService) GetUserInfoByName(name string) (*model.User, error) {
	return u.Where(u.Name.Eq(name)).First()
}

func (userService UserService) PasswordRegisterUser(name, password string) (*model.User, error) {
	password, err := encrypt.PasswordHash(password)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Name:     name,
		Password: &password,
	}
	err = u.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService UserService) PasswordAuthenticateUser(name, password string) (*model.User, error) {
	user, err := u.Where(u.Name.Eq(name)).First()
	if err != nil {
		return nil, err
	}
	err = encrypt.PasswordVerify(*user.Password, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
