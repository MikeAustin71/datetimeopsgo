package common

import (
	"fmt"
	"time"
)

func ExampleSetStartEndTimes() {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := DurationUtility{}

	dur.SetStartEndTimes(t1, t2)

	dDto, err := dur.GetYearMthDays()

	if err != nil {
		panic(err)
	}

	expected := "3-Years 2-months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	fmt.Println("        Expected: ", expected)
	fmt.Println(" YrsMthsDaysTime: ", dDto.DisplayStr)

	expected = "28082h4m2s"
	fmt.Println()
	dDto, err = dur.GetDefaultDuration()
	if err != nil {
		panic(err)
	}
	fmt.Println("Expected Default: ", expected)
	fmt.Println("Default Duration: ", dDto.DisplayStr)

	fmt.Println()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	dDto, err = dur.GetDaysDuration()
	if err != nil {
		panic(err)
	}
	fmt.Println("     Expected Days: ", expected)
	fmt.Println("     Days Duration: ", dDto.DisplayStr)

	fmt.Println()
	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	dDto, err = dur.GetHoursDuration()
	if err != nil {
		panic(err)
	}
	fmt.Println("    Expected Hours: ", expected)
	fmt.Println("    Hours Duration: ", dDto.DisplayStr)

	fmt.Println()
	expected = "3-Years 2-Months 2-Weeks 1-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	dDto, err = dur.GetYrMthWkDayHrMinSecNanosecs()
	if err != nil {
		panic(err)
	}

	fmt.Println("Expected YrMthWeekDayTime : ", expected)
	fmt.Println("          YrMthWeekDayTime: ", dDto.DisplayStr)

	fmt.Println()

	expected = "167-Weeks 1-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	dDto, err = dur.GetWeeksDaysDuration()

	if err != nil {
		panic(err)
	}

	fmt.Println("Expected Weeks Days Time: ", expected)
	fmt.Println("     Weeks Days Duration: ", dDto.DisplayStr)

}
