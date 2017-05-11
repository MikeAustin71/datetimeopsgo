package common

import (
	"errors"
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

// GetEverythingFormat - provides a sample of the
// 'GetEverything Date Time Format!
func GetEverythingFormat() {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DateTimeUtility{}
	str := dt.GetDateTimeEverything(testTime)
	fmt.Println("Everything Format: ", str)
	// Saturday April 29, 2017 19:54:30.123456489 -0500 CDT
}

// GetCurrentTimeAsString - Get current time in the form of a string
func GetCurrentTimeAsString() {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	dt := DateTimeUtility{}
	t, err := time.Parse(fmtstr, tstr)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := dt.GetDateTimeStr(t)

	fmt.Println(tstr, "=", t)
	fmt.Println("result=", result)
}

// GetBasicDuration - Returns basic duration as a string
func GetBasicDuration() {
	t1str := "04/29/2017 19:54:30 -0500 CDT"
	t2str := "04/29/2017 20:56:32 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	dt := DateTimeUtility{}
	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		panic(errors.New("Time Parse1 Error:" + err.Error()))
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		panic(errors.New("Time Parse2 Error:" + err.Error()))
	}

	duration, err := dt.GetDuration(t1, t2)

	fmt.Println("Duration:", duration)
}
