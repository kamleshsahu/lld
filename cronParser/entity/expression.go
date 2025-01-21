package entity

import (
	"cronParser/utils"
	"fmt"
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
