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

func (e *Expression) Next(fromTime time.Time) time.Time {
	return e.YearMatch(fromTime)
}

func (e *Expression) YearMatch(fromTime time.Time) time.Time {
	idx := sort.SearchInts(e.Year, fromTime.Year())
	// no matching year, impossible cron, return empty object
	if idx == len(e.Year) {
		return time.Time{}
	}
	// possible answer next year
	if e.Year[idx] > fromTime.Year() {
		// change time to next possible year from expression
		t := time.Date(e.Year[idx], time.January, 1, 0, 0, 0, 0, fromTime.Location())
		//try matching month
		return e.MonthMatch(t)
	}
	// exact year match, now match month
	return e.MonthMatch(fromTime)
}

func (e *Expression) MonthMatch(fromTime time.Time) time.Time {
	idx := sort.Search(len(e.Month), func(i int) bool {
		return e.Month[i] >= fromTime.Month()
	})
	if idx == len(e.Month) {
		// No matching month in the current year, increment to the next year
		nextYear := time.Date(fromTime.Year()+1, time.January, 1, 0, 0, 0, 0, fromTime.Location())
		// Change in year, restart matching with newTime
		return e.YearMatch(nextYear)
	}

	// Matching month found
	month := e.Month[idx]
	if month > fromTime.Month() {
		// Reset time to 1st day of next month
		t := time.Date(fromTime.Year(), month, 1, 0, 0, 0, 0, fromTime.Location())
		// Match day
		return e.DayMatch(t)
	}
	// Exact month found, now match day
	return e.DayMatch(fromTime)
}

func (e *Expression) DayMatch(fromTime time.Time) time.Time {
	// each month can have different no. of daye (bw 1 - 30 or 1-31 or 1-28 in case of feb),
	// so recompute actual no of daye using below function
	days := e.calculateActualDaysOfMonth(fromTime.Year(), int(fromTime.Month()), fromTime.Location())

	idx := sort.Search(len(days), func(i int) bool {
		return days[i] >= fromTime.Day()
	})
	if idx == len(days) {
		// No matching day in the current month, move to the next month
		nextMonth := time.Date(fromTime.Year(), fromTime.Month()+1, 1, 0, 0, 0, 0, fromTime.Location())
		// Increment 1 month can lead to change in year, restart matching with newTime
		return e.YearMatch(nextMonth)
	}

	// Matching day found
	day := days[idx]
	if day > fromTime.Day() {
		t := time.Date(fromTime.Year(), fromTime.Month(), day, 0, 0, 0, 0, fromTime.Location())
		return e.HourMatch(t)
	}
	// Exact day match, now match hour
	return e.HourMatch(fromTime)
}

func (e *Expression) HourMatch(fromTime time.Time) time.Time {
	idx := sort.Search(len(e.Hour), func(i int) bool {
		return e.Hour[i] >= fromTime.Hour()
	})
	if idx == len(e.Hour) {
		// No matching hour in the current day, move to the next day
		nextDay := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day()+1, 0, 0, 0, 0, fromTime.Location())
		// Increment 1 day can lead to change in year, restart matching with newTime
		return e.YearMatch(nextDay)
	}

	// Matching hour found
	hour := e.Hour[idx]
	if hour > fromTime.Hour() {
		t := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), hour, 0, 0, 0, fromTime.Location())
		return e.MinuteMatch(t)
	}
	// Exact Hour Match, now match minutes
	return e.MinuteMatch(fromTime)
}

func (e *Expression) MinuteMatch(fromTime time.Time) time.Time {
	idx := sort.Search(len(e.Minute), func(i int) bool {
		return e.Minute[i] >= fromTime.Minute()
	})
	if idx == len(e.Minute) {
		// No matching minute in the current hour, move to the next hour
		nextHour := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), fromTime.Hour()+1, 0, 0, 0, fromTime.Location())
		// Increment 1 hour can lead to change in year, restart matching with newTime
		return e.YearMatch(nextHour)
	}

	// Matching minute found
	minute := e.Minute[idx]
	if minute > fromTime.Minute() {
		t := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), fromTime.Hour(), minute, 0, 0, fromTime.Location())
		return e.SecondMatch(t)
	}

	// Exact matching minute, now match second
	return e.SecondMatch(fromTime)
}

func (e *Expression) SecondMatch(fromTime time.Time) time.Time {
	idx := sort.Search(len(e.Second), func(i int) bool {
		return e.Second[i] >= fromTime.Second()
	})
	if idx == len(e.Second) {
		// No matching second in the current minute, move to the next minute
		nextMinute := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), fromTime.Hour(), fromTime.Minute()+1, 0, 0, fromTime.Location())
		// Increment 1 minute can lead to change in year, restart matching with newTime
		return e.YearMatch(nextMinute)
	}

	// Matching second found
	second := e.Second[idx]
	t := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), fromTime.Hour(), fromTime.Minute(), second, 0, fromTime.Location())
	return t
}

func (e *Expression) Set(vals []int, p ParserType) {
	switch p {
	case Second:
		e.Second = vals
	case Minute:
		e.Minute = vals
	case Hour:
		e.Hour = vals
	case Day:
		e.Day = vals
	case Month:
		months := make([]time.Month, len(vals))
		for i, val := range vals {
			months[i] = time.Month(val)
		}
		e.Month = months
	case Year:
		e.Year = vals
	}
}

// each month can have different no of day, so need to calculate actual days for given specific month
func (e *Expression) calculateActualDaysOfMonth(year, month int, location *time.Location) []int {
	firstDayOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, location)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	ans := make([]int, 0)
	mymap := make(map[int]bool)
	for _, val := range utils.GenericDefaultList[1 : lastDayOfMonth.Day()+1] {
		mymap[val] = true
	}
	for _, val := range e.Day {
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
