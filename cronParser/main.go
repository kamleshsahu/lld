package main

import (
	"fmt"
	"time"
)

//
//schudule1, ans1 := CronParser("*/15 * * * trigger")
//fmt.Println(schudule1, ans1)
//
//schudule2, ans2 := CronParser("*/15 2-6 */5 */6 trigger")
//fmt.Println(schudule2, ans2)

//ans3 := CronParser("* * * 4,5,6 trigger")
//fmt.Println(ans3)

//fromDate := time.Date(2028, time.Month(2), 29, 1, 0, 0, 0, time.Now().Location())

func main() {

	exp, _ := CronParser("10 */30 0 29 2 * trigger")
	//fmt.Println(str)
	fmt.Println()
	now := time.Now()
	nextTrigger := exp.Next(now)

	fmt.Println(nextTrigger)

	nextTrigger2 := exp.Next(nextTrigger.Add(time.Second))
	fmt.Println(nextTrigger2)

	nextTrigger3 := exp.Next(nextTrigger2.Add(time.Second))
	fmt.Println(nextTrigger3)
}
