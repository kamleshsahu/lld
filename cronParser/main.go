package main

import (
	"cronParser/cronParser"
	"fmt"
	"time"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//var cronExpression string
	//
	//cronExpression, err := reader.ReadString('\n')
	//
	//if err != nil {
	//	fmt.Println("Error reading input:", err)
	//	return
	//}

	//cronExpression = cronExpression[:len(cronExpression)-1]

	cp := cronParser.NewDefaultCronParser(nil)
	exp, err := cp.Parse("*/15 2-4,4-8 2,1 1 2 /usr/bin/find")

	next := exp.Next(time.Now())
	fmt.Println(next.String(), next.Weekday())

	if err != nil {
		fmt.Println("Error parsing cron expression:", err)
		return
	}
	fmt.Println()
	fmt.Println(exp.ToString())
}
