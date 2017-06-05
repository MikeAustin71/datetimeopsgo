package main

import (
	"MikeAustin71/datetimeopsgo/05_DateTimeFormatsUtility/common"
)

/*

import (
	"MikeAustin71/datetimeopsgo/05_DateTimeFormatsUtility/common"
	"errors"
	"fmt"
	"time"
)

*/

func main() {

	//fmtDateTime := "2006-01-02 15:04:05 -0700 MST"
	// tDateTime := "Monday 11/12/2016 4:26 PM"
	// tDateTime := "7-6-16 9:30AM"
	// tDateTime := "November 12, 2016"
	//tDateTime := "November 12, 11:26pm -0600 CST 2016"
	// tDateTime := "November 12, 2016 11:6pm +0000 UTC"
	 tDateTime := "November 12, 2016 1:6pm +0000 UTC"
	//tDateTime := "2016-11-12 23:26:00 +0000 UTC"
	//tDateTime := "2016-11-26 16:26 -0600"
	//tDateTime := "11/12/16 4:26 PM"
	// tDateTime := "5/27/2017 11:42PM CDT"
	//tDateTime := "2016-11-26 16:26 CDT -0600"
	//tDateTime := "June 12th, 2016 4:26 PM"
	//tDateTime := "05.03.2017"
	//tDateTime := "Mon"

	common.TestParseDateTimeCreateFormatsInMemory(tDateTime, "")

	//common.TestParseSampleDateTimes()
	//common.WriteAllFormatsToFile()
	//common.WriteFormatStatsToFile()

	// 2006-01-02 15:04 -0700 MST
	// 0000024 2006-1-2 15:04 MST -0700

	/*
		t, err := time.Parse("2006-01-02 15:04 MST -0700", tDateTime)

		if err != nil {
			panic(err)
		}

		fmt.Println("Success!")
		fmt.Println("Original Time: ", tDateTime)
		fmt.Println("  Parsed Time: ", t)
	*/

}
