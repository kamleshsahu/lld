package entity

import (
	"cronParser/utils"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Expression struct {
	Minute    []int
	Hour      []int
	Day       map[int]bool
	Month     []time.Month
	DayOfWeek map[int]bool
	Command   string
}

func (e *Expression) Set(vals []int, timeUnit TimeUnit) {
	switch timeUnit {
	case Minute:
		e.Minute = vals
	case Hour:
		e.Hour = vals
	case Day:
		e.Day = utils.ToMap(vals)
	case Month:
		months := make([]time.Month, len(vals))
		for i, val := range vals {
			months[i] = time.Month(val)
		}
		e.Month = months
	case DayOfWeek:
		e.DayOfWeek = utils.ToMap(vals)
	}
}

func (e *Expression) ToString() string {
	fields := []struct {
		description string
		value       string
	}{
		{"minute", utils.FormatSlice(e.Minute)},
		{"hour", utils.FormatSlice(e.Hour)},
		{"day of month", utils.FormatMap(e.Day)},
		{"month", utils.FormatMonths(e.Month)},
		{"day of week", utils.FormatMap(e.DayOfWeek)},
		{"command", e.Command},
	}

	var result strings.Builder
	for _, field := range fields {
		result.WriteString(fmt.Sprintf("%-*s %s\n", 20, field.description, field.value))
	}

	return result.String()
}

func (e *Expression) Next(fromTime time.Time) time.Time {
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
		return e.MonthMatch(nextYear)
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
		return e.MonthMatch(nextMonth)
	}

	// Matching day found
	day := days[idx]
	if day > fromTime.Day() {
		t := time.Date(fromTime.Year(), fromTime.Month(), day, 0, 0, 0, 0, fromTime.Location())
		return t
	}
	// Exact day match, now match hour
	return fromTime
}

func (e *Expression) calculateActualDaysOfMonth(year, month int, location *time.Location) []int {
	firstDayOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, location)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	ans := make([]int, 0)
	mymap := make(map[int]int)
	for _, val := range utils.GenericDefaultList[1 : lastDayOfMonth.Day()+1] {
		mymap[val] = int(time.Date(year, time.Month(month), val, 0, 0, 0, 0, location).Weekday())
	}
	for val, exist := range e.Day {
		if !exist {
			continue
		}
		if _, exist = mymap[val]; !exist {
			continue
		}
		if !e.DayOfWeek[mymap[val]] {
			continue
		}
		ans = append(ans, val)
	}
	sort.Ints(ans)
	return ans
}

func NewExpression() *Expression {
	exp := &Expression{
		Minute:    make([]int, 0),
		Hour:      make([]int, 0),
		Day:       make(map[int]bool),
		Month:     make([]time.Month, 0),
		DayOfWeek: make(map[int]bool),
	}
	return exp
}
