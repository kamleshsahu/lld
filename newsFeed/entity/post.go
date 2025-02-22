package entity

type Post struct {
	Id     int
	Time   int
	UserId int
	Title  string
}

type User struct {
	Id        int
	Name      string
	Followers map[int]bool
	Following map[int]bool
	Feed      []int
}

type PostDTO struct {
	Id     int
	Action string
}

type RelationshipDTO struct {
	FollowerId int
	FolloweeId int
	Action     string
}
