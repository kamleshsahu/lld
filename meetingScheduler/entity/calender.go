package entity

import "time"

type Calender struct {
	IsBooked map[time.Time]time.Time
}

//12 - 1
//12:30 - 1:30

//12 -
