package service

import "lld/newsFeed/entity"

type IFeedService interface {
	IObserver
	GetUserFeed(userId int) ([]entity.Post, error)
	GetUserFeedPaginated(userId int, pageNum int) ([]entity.Post, error)
}

type FeedService struct {
	DefaultFeedSize int
	userService     IUserService
	postService     IPostService
}

func (f *FeedService) Notify(msgtype string, data interface{}) {
	switch msgtype {
	case "post":
		postdto := data.(entity.PostDTO)
		post, _ := f.postService.GetPostById(postdto.Id)
		switch postdto.Action {
		case "added":
			f.addToTimeline(post)
		case "deleted":
			f.removeFromTimeline(post)
		}
	case "relationship":
		relationdto := data.(entity.RelationshipDTO)
		switch relationdto.Action {
		case "followed":
			f.addallPostToTimeline(relationdto.FollowerId, relationdto.FolloweeId)
		case "unfollowed":
			f.removeAllPostFromTimeline(relationdto.FollowerId, relationdto.FolloweeId)
		}
	}
}

func (f *FeedService) addToTimeline(post entity.Post) error {
	user, _ := f.userService.GetUserById(post.UserId)
	for followerId := range user.Followers {
		follower, _ := f.userService.GetUserById(followerId)
		follower.Feed = append(follower.Feed, post.Id)
	}

	return nil
}

func (f *FeedService) addallPostToTimeline(followerId, followeeId int) error {
	follower, _ := f.userService.GetUserById(followerId)

	feed := follower.Feed

	posts, _ := f.postService.GetPostByUserId(followeeId)
	for _, post := range posts {
		feed = append(feed, post.Id)
	}
	follower.Feed = feed
	return nil
}

func (f *FeedService) removeAllPostFromTimeline(followerId, followeeId int) error {
	follower, _ := f.userService.GetUserById(followerId)
	nf := make([]int, 0)
	feed := follower.Feed
	for _, feedId := range feed {
		post, err := f.postService.GetPostById(feedId)
		if err != nil {
			return err
		}
		if post.UserId != followeeId {
			nf = append(nf, post.Id)
		}
	}

	follower.Feed = nf
	return nil
}

func (f *FeedService) removeFromTimeline(post entity.Post) error {
	user, _ := f.userService.GetUserById(post.UserId)
	for followerId := range user.Followers {
		follower, _ := f.userService.GetUserById(followerId)
		for i, postId := range follower.Feed {
			if postId == post.Id {
				follower.Feed = append(follower.Feed[:i], follower.Feed[i+1:]...)
				break
			}
		}
	}
	return nil
}

func (f *FeedService) GetUserFeed(userId int) ([]entity.Post, error) {
	user, err := f.userService.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	var posts []entity.Post
	if len(user.Feed) == 0 {
		return posts, nil
	}

	start := max(len(user.Feed)-f.DefaultFeedSize, 0)
	ids := user.Feed[start:]

	for _, id := range ids {
		post, _ := f.postService.GetPostById(id)
		posts = append(posts, post)
	}

	return posts, nil
}

func (f *FeedService) GetUserFeedPaginated(userId int, pageNum int) ([]entity.Post, error) {
	user, err := f.userService.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	start := len(user.Feed) - f.DefaultFeedSize
	ids := user.Feed[start:]

	var posts []entity.Post

	for _, id := range ids {
		post, _ := f.postService.GetPostById(id)
		posts = append(posts, post)
	}

	return posts, nil
}

func NewFeedService(service IUserService, ps IPostService) IFeedService {
	return &FeedService{userService: service, DefaultFeedSize: 10, postService: ps}
}
