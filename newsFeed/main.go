package main

import (
	"fmt"
	"lld/newsFeed/entity"
	"lld/newsFeed/service"
)

func main() {

	userService := service.NewUserService()

	postService := service.NewPostService(userService)

	rs := service.NewRelationshipService(userService)

	fs := service.NewFeedService(userService, postService)

	kamlesh := entity.User{Name: "Kamlesh"}
	tikesh := entity.User{Name: "Tikesh"}
	vaibhav := entity.User{Name: "Vaibhav"}

	kamlesh, _ = userService.CreateUser(kamlesh)
	tikesh, _ = userService.CreateUser(tikesh)
	vaibhav, _ = userService.CreateUser(vaibhav)

	rs.FollowUser(kamlesh.Id, tikesh.Id)
	rs.FollowUser(kamlesh.Id, vaibhav.Id)
	rs.FollowUser(vaibhav.Id, tikesh.Id)

	postService.Subscribe(fs)
	rs.Subscribe(fs)

	post := entity.Post{UserId: tikesh.Id, Title: "Post by tikesh"}
	post, _ = postService.CreatePost(post)
	post, _ = postService.CreatePost(entity.Post{UserId: tikesh.Id, Title: "Post by tikesh 2"})
	post, _ = postService.CreatePost(entity.Post{UserId: vaibhav.Id, Title: "Post by vaibhav 3"})

	post, _ = postService.CreatePost(entity.Post{UserId: kamlesh.Id, Title: "Post by kamlesh 4"})

	feed, _ := fs.GetUserFeed(kamlesh.Id)

	for i := 0; i < len(feed); i++ {
		fmt.Println("Kamlesh feed : ", feed[i].Title)
	}
	rs.UnfollowUser(kamlesh.Id, tikesh.Id)

	feed1, _ := fs.GetUserFeed(kamlesh.Id)
	fmt.Println("kamlesh Unfollowed tikesh")

	for i := 0; i < len(feed1); i++ {
		fmt.Println("Kamlesh feed : ", feed1[i].Title)
	}

	vaibhavTimeline, _ := fs.GetUserFeed(vaibhav.Id)
	for i := 0; i < len(vaibhavTimeline); i++ {
		fmt.Println("Vaibhav feed : ", vaibhavTimeline[i].Title)
	}

	rs.UnfollowUser(vaibhav.Id, tikesh.Id)
	fmt.Println("vaibhav Unfollowed tikesh")

	rs.FollowUser(vaibhav.Id, kamlesh.Id)
	fmt.Println("vaibhav Followed kamlesh")

	vaibhavTimeline, _ = fs.GetUserFeed(vaibhav.Id)
	for i := 0; i < len(vaibhavTimeline); i++ {
		fmt.Println("Vaibhav feed : ", vaibhavTimeline[i].Title)
	}

	feed1, _ = fs.GetUserFeed(kamlesh.Id)

	for i := 0; i < len(feed1); i++ {
		fmt.Println("Kamlesh feed : ", feed1[i].Title)
	}

	postService.DeletePost(4)

	fmt.Println("Post 1 deleted")

	vaibhavTimeline, _ = fs.GetUserFeed(vaibhav.Id)
	for i := 0; i < len(vaibhavTimeline); i++ {
		fmt.Println("Vaibhav feed : ", vaibhavTimeline[i].Title)
	}

	feed1, _ = fs.GetUserFeed(kamlesh.Id)

	for i := 0; i < len(feed1); i++ {
		fmt.Println("Kamlesh feed : ", feed1[i].Title)
	}
}
