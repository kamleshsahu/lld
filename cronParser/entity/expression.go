package entity

import (
	"lld/cronParser/utils"
	"sort"
	"time"
)

type Expression struct {
	Second []int
	Minute []int
	Hour   []int
	Day    []int
	Month  []time.Month
	Year   []int
}

type ParserType string

const (
	Second      ParserType = "second"
	Minute      ParserType = "minute"
	Hour        ParserType = "hour"
	Day         ParserType = "day"
	Month       ParserType = "month"
	Year        ParserType = "year"
	Description ParserType = "description"
)

func (p ParserType) String() string {
	return string(p)
}

func (s *Expression) Next(fromTime time.Time) time.Time {
	return s.YearMatch(fromTime)
}

func (s *Expression) YearMatch(fromTime time.Time) time.Time {
	idx := sort.SearchInts(s.Year, fromTime.Year())
	// no matching year, impossible cron, return empty object
	if idx == len(s.Year) {
		return time.Time{}
	}
	// possible answer next year
	if s.Year[idx] > fromTime.Year() {
		// change time to next possible year from expression
		t := time.Date(s.Year[idx], time.January, 1, 0, 0, 0, 0, fromTime.Location())
		//try matching month
		return s.MonthMatch(t)
	}
	// exact year match, now match month
	return s.MonthMatch(fromTime)
}

func (s *Expression) MonthMatch(fromTime time.Time) time.Time {
	idx := sort.Search(len(s.Month), func(i int) bool {
		return s.Month[i] >= fromTime.Month()
	})
	if idx == len(s.Month) {
		// No matching month in the current year, increment to the next year
		nextYear := time.Date(fromTime.Year()+1, time.January, 1, 0, 0, 0, 0, fromTime.Location())
		// Change in year, restart matching with newTime
		return s.YearMatch(nextYear)
	}

	// Matching month found
	month := s.Month[idx]
	if month > fromTime.Month() {
		// Reset time to 1st day of next month
		t := time.Date(fromTime.Year(), month, 1, 0, 0, 0, 0, fromTime.Location())
		// Match day
		return s.DayMatch(t)
	}
	// Exact month found, now match day
	return s.DayMatch(fromTime)
}

func (s *Expression) DayMatch(fromTime time.Time) time.Time {
	// each month can have different no. of days (bw 1 - 30 or 1-31 or 1-28 in case of feb),
	// so recompute actual no of days using below function
	days := s.calculateActualDaysOfMonth(fromTime.Year(), int(fromTime.Month()), fromTime.Location())

	idx := sort.Search(len(days), func(i int) bool {
		return days[i] >= fromTime.Day()
	})
	if idx == len(days) {
		// No matching day in the current month, move to the next month
		nextMonth := time.Date(fromTime.Year(), fromTime.Month()+1, 1, 0, 0, 0, 0, fromTime.Location())
		// Increment 1 month can lead to change in year, restart matching with newTime
		return s.YearMatch(nextMonth)
	}

	// Matching day found
	day := days[idx]
	if day > fromTime.Day() {
		t := time.Date(fromTime.Year(), fromTime.Month(), day, 0, 0, 0, 0, fromTime.Location())
		return s.HourMatch(t)
	}
	// Exact day match, now match hour
	return s.HourMatch(fromTime)
}

func (s *Expression) HourMatch(fromTime time.Time) time.Time {
	idx := sort.Search(len(s.Hour), func(i int) bool {
		return s.Hour[i] >= fromTime.Hour()
	})
	if idx == len(s.Hour) {
		// No matching hour in the current day, move to the next day
		nextDay := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day()+1, 0, 0, 0, 0, fromTime.Location())
		// Increment 1 day can lead to change in year, restart matching with newTime
		return s.YearMatch(nextDay)
	}

	// Matching hour found
	hour := s.Hour[idx]
	if hour > fromTime.Hour() {
		t := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), hour, 0, 0, 0, fromTime.Location())
		return s.MinuteMatch(t)
	}
	// Exact Hour Match, now match minutes
	return s.MinuteMatch(fromTime)
}

func (s *Expression) MinuteMatch(fromTime time.Time) time.Time {
	idx := sort.Search(len(s.Minute), func(i int) bool {
		return s.Minute[i] >= fromTime.Minute()
	})
	if idx == len(s.Minute) {
		// No matching minute in the current hour, move to the next hour
		nextHour := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), fromTime.Hour()+1, 0, 0, 0, fromTime.Location())
		// Increment 1 hour can lead to change in year, restart matching with newTime
		return s.YearMatch(nextHour)
	}

	// Matching minute found
	minute := s.Minute[idx]
	if minute > fromTime.Minute() {
		t := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), fromTime.Hour(), minute, 0, 0, fromTime.Location())
		return s.SecondMatch(t)
	}

	// Exact matching minute, now match second
	return s.SecondMatch(fromTime)
}

func (s *Expression) SecondMatch(fromTime time.Time) time.Time {
	idx := sort.Search(len(s.Second), func(i int) bool {
		return s.Second[i] >= fromTime.Second()
	})
	if idx == len(s.Second) {
		// No matching second in the current minute, move to the next minute
		nextMinute := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), fromTime.Hour(), fromTime.Minute()+1, 0, 0, fromTime.Location())
		// Increment 1 minute can lead to change in year, restart matching with newTime
		return s.YearMatch(nextMinute)
	}

	// Matching second found
	second := s.Second[idx]
	t := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), fromTime.Hour(), fromTime.Minute(), second, 0, fromTime.Location())
	return t
}

func (s *Expression) Set(vals []int, p ParserType) {
	switch p {
	case Second:
		s.Second = vals
	case Minute:
		s.Minute = vals
	case Hour:
		s.Hour = vals
	case Day:
		s.Day = vals
	case Month:
		months := make([]time.Month, len(vals))
		for i, val := range vals {
			months[i] = time.Month(val)
		}
		s.Month = months
	case Year:
		s.Year = vals
	}
}

// each month can have different no of day, so need to calculate actual days for given specific month
func (s *Expression) calculateActualDaysOfMonth(year, month int, location *time.Location) []int {
	firstDayOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, location)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	ans := make([]int, 0)
	mymap := make(map[int]bool)
	for _, val := range utils.GenericDefaultList[1 : lastDayOfMonth.Day()+1] {
		mymap[val] = true
	}
	for _, val := range s.Day {
		if mymap[val] {
			ans = append(ans, val)
		}
	}
	return ans
}

func NewExpression() *Expression {
	exp := &Expression{
		Second: make([]int, 0),
		Minute: make([]int, 0),
		Hour:   make([]int, 0),
		Day:    make([]int, 0),
		Month:  make([]time.Month, 0),
		Year:   make([]int, 0),
	}
	return exp
}
