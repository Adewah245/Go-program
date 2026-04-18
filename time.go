package main

import (
	"fmt"
	"time"
)

func timeset(now string) string {
	now = time.Now().Format(now)
	hour := time.Now().Local().Weekday()
	if hour >= 0 && hour < 12 {
		fmt.Println("Good Morning; It's: ")
	} else if hour >= 12 && hour < 16 {
		fmt.Println("Good Afternoon; It's: ")
	} else if hour >= 16 && hour < 24 {
		fmt.Println("Good Night; It's: ")
	}
	return hour.String()
}
