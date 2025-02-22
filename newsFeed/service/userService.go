package service

import (
	"errors"
	"lld/newsFeed/entity"
)

type IUserService interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserById(userId int) (*entity.User, error)
}
type UserService struct {
	NextId int
	Users  map[int]*entity.User
}

func (u *UserService) CreateUser(user entity.User) (entity.User, error) {
	user.Following = make(map[int]bool)
	user.Followers = make(map[int]bool)
	user.Feed = make([]int, 0)
	u.NextId++
	user.Id = u.NextId
	u.Users[u.NextId] = &user
	return user, nil
}

func (u *UserService) GetUserById(userId int) (*entity.User, error) {
	if u, ok := u.Users[userId]; ok {
		return u, nil
	}
	return nil, errors.New("user not found")
}

func NewUserService() IUserService {
	return &UserService{Users: make(map[int]*entity.User)}
}
