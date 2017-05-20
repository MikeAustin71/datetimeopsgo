package common

import (
	"errors"
	"fmt"
	"time"
)

// GetBasicDuration - Returns basic duration as a string
func GetBasicDuration() {
	t1str := "04/29/2017 19:54:30 -0500 CDT"
	t2str := "04/29/2017 20:56:32 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	durationUtility := DurationUtility{}
	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		panic(errors.New("Time Parse1 Error:" + err.Error()))
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		panic(errors.New("Time Parse2 Error:" + err.Error()))
	}

	duration, err := durationUtility.GetDuration(t1, t2)

	fmt.Println("Duration:", duration)
	// Duration: 1h2m2s
}

// GetElapsedTimeDuration - example of
// GetElapsedTime() in DateTime Utility
func GetElapsedTimeDuration() {
	tstr1 := "04/15/2017 19:54:30.123456489 -0500 CDT"
	tstr2 := "04/18/2017 09:21:16.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, tstr1)
	t2, _ := time.Parse(fmtstr, tstr2)
	durationUtility := DurationUtility{}

	ed, _ := durationUtility.GetElapsedTime(t1, t2)
	fmt.Println("Elapsed Time: ", ed.DurationStr)
	// "2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds"

	fmt.Println("")
	fmt.Println("Default Duration: ", ed.DefaultStr)
	// 61h26m46.864197832s

	fmt.Println("")
	fmt.Println("NanosecStr: ", ed.NanosecStr)
	// 2-Days 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds

}

func GetElapsedYears() {
	t1str := "02/15/2014 19:54:30 -0500 CDT"
	t2str := "04/30/2017 22:58:32 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	durationUtility := DurationUtility{}

	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		panic(fmt.Errorf("Time Parse1 Error: %v", err))
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		panic(fmt.Errorf("Time Parse2 Error: %v", err))
	}

	dur, err := durationUtility.GetDuration(t1, t2)
	if err != nil {
		panic(fmt.Errorf("Get Duration Failed: %v", err))
	}

	ed := durationUtility.GetDurationBreakDown(dur)
	fmt.Println("Elapsed Time: ", ed.DurationStr)
	// Elapsed Time:  3-Years 75-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds

	fmt.Println("")
	fmt.Println("Default Duration: ", ed.DefaultStr)
	// Default Duration:  28083h4m2s

	fmt.Println("")
	fmt.Println("NanosecStr: ", ed.NanosecStr)
	// NanosecStr:  3-Years 75-Days 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds

}
