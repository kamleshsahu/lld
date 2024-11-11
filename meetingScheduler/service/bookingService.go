package service

import (
	"awesomeProject/meeting/entity"
	"fmt"
	"time"
)

type bookingService struct {
	meetingId           int
	RoomsByDay          map[int][]entity.Room
	MeetingsById        map[int]entity.Meeting
	notificationService INotificationService
}

func (b *bookingService) Book(start, end time.Time) (meetingId int, roomId int, err error) {

	day := start.Day()
	rooms := b.RoomsByDay[day]

	for roomId, room := range rooms {
		isConflict := false
		for startTime, endTime := range room.Calender.IsBooked {
			if start.After(startTime) && start.Before(endTime) {
				isConflict = true
				break
			}
			if end.After(startTime) && end.Before(endTime) {
				isConflict = true
				break
			}
		}
		if !isConflict {
			meetingId++
			meeting := entity.Meeting{UsersEmails: make([]string, 0)}
			b.MeetingsById[meetingId] = meeting
			room.Calender.IsBooked[start] = end
			return meetingId, roomId, nil
		}
	}
	return -1, -1, fmt.Errorf("no meeting room found")
}

func (b *bookingService) InviteUser(meetingId int, usersEmails []string) (err error) {
	meeting := b.MeetingsById[meetingId]
	for _, email := range usersEmails {
		meeting.UsersEmails = append(b.MeetingsById[meetingId].UsersEmails, email)
		b.notificationService.Notify(fmt.Sprintf("You are invited to meeting %d", meetingId), email)
	}

	return nil
}

func NewBookingService(rooms int, notifyService INotificationService) IBookingService {

	roomsByDay := make(map[int][]entity.Room)
	for i := 0; i < 31; i++ {
		roomsByDay[i] = make([]entity.Room, rooms)
	}

	return &bookingService{
		RoomsByDay:          roomsByDay,
		MeetingsById:        make(map[int]entity.Meeting),
		notificationService: notifyService,
	}
}
