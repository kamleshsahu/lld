package entity

type TimeUnit string

const (
	Minute    TimeUnit = "minute"
	Hour      TimeUnit = "hour"
	Day       TimeUnit = "day"
	Month     TimeUnit = "month"
	DayOfWeek TimeUnit = "dayOfWeek"
)

func (p *TimeUnit) String() string {
	return string(*p)
}

