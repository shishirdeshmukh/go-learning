package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello Time")

	// presetime := time.Now()
	//	fmt.Println(presetime)

	//fmt.Println(presetime.Format("01-02-2006 15:04:05 Monday"))

	// createdDate := time.Date(2025, time.January, 10, 21, 21, 0, 0, time.UTC)
	// fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

	// 4. Comparing Times
	// t2 := presetime.Add((10 + time.Hour))

	// fmt.Println((presetime.Before(t2)))
	// fmt.Println((presetime.Before(t2)))
	// fmt.Println((presetime.Equal(t2)))

	//Sleep
	// time.Sleep((3 * time.Second))
	// fmt.Println("After Sleep")

	//time zone

	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Kolkata")
	ny := now.In(loc)

	fmt.Println(ny)
	fmt.Println(loc)

}
