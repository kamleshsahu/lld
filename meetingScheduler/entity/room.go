package entity

type Room struct {
	Calender Calender
	Meetings []Meeting
}

type Meeting struct {
	UsersEmails []string
}
