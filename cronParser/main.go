package main

import (
	"bufio"
	"cronParser/cronParser"
	"fmt"
	"os"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		var cronExpression string

		cronExpression, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		cronExpression = cronExpression[:len(cronExpression)-1]

		cp := cronParser.NewDefaultCronParser()
		exp, err := cp.Parse(cronExpression)

		if err != nil {
			fmt.Println("Error parsing cron expression:", err)
			return
		}

		fmt.Println(exp.ToString())
	}
}
