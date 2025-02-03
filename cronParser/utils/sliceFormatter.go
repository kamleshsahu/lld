package utils

import (
	"fmt"
	"strings"
	"time"
)

func FormatSlice(slice []int) string {
	strs := make([]string, len(slice))
	for i, val := range slice {
		strs[i] = fmt.Sprintf("%d", val)
	}
	return strings.Join(strs, " ")
}

func FormatMap(m map[int]bool) string {
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	strs := make([]string, len(keys))
	for i, val := range keys {
		strs[i] = fmt.Sprintf("%d", val)
	}
	return strings.Join(strs, " ")
}

func FormatMonths(months []time.Month) string {
	//sort.Slice(months, func(i, j int) bool {
	//	return months[i] < months[j]
	//})

	strs := make([]string, len(months))
	for i, month := range months {
		strs[i] = fmt.Sprintf("%d", int(month))
	}
	return strings.Join(strs, " ")
}

func ToMap(arr []int) map[int]bool {
	m := make(map[int]bool)
	for _, v := range arr {
		m[v] = true
	}
	return m
}

var (
	GenericDefaultList = []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	}
)
