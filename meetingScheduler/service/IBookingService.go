package service

import "time"

type IBookingService interface {
	Book(start, end time.Time) (meetingId int, roomId int, err error)
	InviteUser(meetingId int, usersEmails []string) (err error)
}
