package service

import "lld/newsFeed/entity"

type IRelationshipService interface {
	IObservable
	FollowUser(followerId int, followedId int) (bool, error)
	UnfollowUser(followerId int, followedId int) (bool, error)
}

type RelationshipService struct {
	Observable
	UserService IUserService
}

func (r *RelationshipService) FollowUser(followerId int, followeeId int) (bool, error) {
	follower, _ := r.UserService.GetUserById(followerId)
	followee, _ := r.UserService.GetUserById(followeeId)
	follower.Following[followee.Id] = true
	followee.Followers[follower.Id] = true
	msg := entity.RelationshipDTO{follower.Id, followee.Id, "followed"}
	r.Fire("relationship", msg)

	return true, nil
}

func (r *RelationshipService) UnfollowUser(followerId int, followeeId int) (bool, error) {
	follower, _ := r.UserService.GetUserById(followerId)
	followee, _ := r.UserService.GetUserById(followeeId)
	delete(follower.Following, followee.Id)
	delete(followee.Followers, follower.Id)

	msg := entity.RelationshipDTO{FollowerId: follower.Id, FolloweeId: followee.Id, Action: "unfollowed"}
	r.Fire("relationship", msg)

	return true, nil
}

func NewRelationshipService(userService IUserService) IRelationshipService {
	return &RelationshipService{UserService: userService}
}
