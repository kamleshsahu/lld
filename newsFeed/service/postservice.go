package service

import (
	"errors"
	"lld/newsFeed/entity"
)

type IPostService interface {
	IObservable
	CreatePost(post entity.Post) (entity.Post, error)
	DeletePost(postId int) (bool, error)
	GetPostById(postId int) (entity.Post, error)
	GetPostByUserId(userId int) ([]entity.Post, error)
}

type PostService struct {
	Observable
	NextId      int
	Posts       map[int]entity.Post
	UserService IUserService
}

func (p *PostService) DeletePost(postId int) (bool, error) {
	if _, exists := p.Posts[postId]; exists {
		p.Fire("post", entity.PostDTO{Action: "deleted", Id: postId})
		delete(p.Posts, postId)
		return true, nil
	}
	return false, nil
}

func (p *PostService) GetPostById(postId int) (entity.Post, error) {
	if _, exists := p.Posts[postId]; exists {
		return p.Posts[postId], nil
	}
	return entity.Post{}, errors.New("post not found")
}

func (p *PostService) GetPostByUserId(userId int) ([]entity.Post, error) {
	userPosts := make([]entity.Post, 0)
	for _, post := range p.Posts {
		if post.UserId == userId {
			userPosts = append(userPosts, post)
		}
	}
	return userPosts, nil
}

func NewPostService(us IUserService) IPostService {
	return &PostService{UserService: us, Posts: make(map[int]entity.Post)}
}

func (p *PostService) CreatePost(post entity.Post) (entity.Post, error) {
	p.NextId++
	post.Id = p.NextId
	p.Posts[post.Id] = post
	p.Fire("post", entity.PostDTO{Action: "added", Id: post.Id})
	return post, nil
}
