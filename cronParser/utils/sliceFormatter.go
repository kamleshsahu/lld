package utils

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func FormatSlice(slice []int) string {
	sort.Ints(slice)
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
	sort.Ints(keys)
	strs := make([]string, len(keys))
	for i, val := range keys {
		strs[i] = fmt.Sprintf("%d", val)
	}
	return strings.Join(strs, " ")
}

func FormatMonths(months []time.Month) string {
	sort.Slice(months, func(i, j int) bool {
		return months[i] < months[j]
	})
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
