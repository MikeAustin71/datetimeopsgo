package common

import (
	"fmt"
	"time"
)

// GetCurrentTimeAsInts - Breaks down time
// to constituent elements as integers.
func GetCurrentTimeAsInts() {
	// Get current time
	t := time.Now().Local()
	var i int64
	i = int64(t.Month())
	fmt.Println("The integer month is: ", i)
	i = int64(t.Day())
	fmt.Println("The integer day is:", i)
	i = int64(t.Year())
	fmt.Println("The integer year is:", i)
	i = int64(t.Hour())
	fmt.Println("The integer hour is:", i)
	i = int64(t.Minute())
	fmt.Println("The integer minute is:", i)
	i = int64(t.Second())
	fmt.Println("The integer second is:", i)
	i = int64(t.Nanosecond())
	fmt.Println("The integer nanosecond is", i)
}
