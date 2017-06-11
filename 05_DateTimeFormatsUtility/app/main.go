package main

import (
	"MikeAustin71/datetimeopsgo/05_DateTimeFormatsUtility/common"
	//"time"
	//"fmt"
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
	// tDateTime := "Saturday 11/12/2016 4:26 PM"
	//tDateTime := "7-6-16 9:30AM"
	// tDateTime := "November 12, 2016"
	//tDateTime := "November 12, 11:26pm -0600 CST 2016"
	//tDateTime := "November 12, 2016 11:6pm +0000 UTC"
	//tDateTime := "November 12, 2016 1:6pm +0000 UTC"
	//tDateTime := "November 12, 2016 1:06pm -0500 EST"
	//tDateTime := "5/31/2017 23:2:17 -0700 PDT"
	//tDateTime := "2016-11-12 23:26:00 +0000 UTC"
	//tDateTime := "2016-11-12 23:26:00Z"
	//tDateTime := "2017-6-12 11:26 p.m. Z"
	//tDateTime := "2017-11-26 16:26 -0600"
	//tDateTime := "2017-6-5 17:16 +0100 BST"
	//tDateTime := "11/12/16 4:26 PM"
	//tDateTime := "11/12/16 4:4 P.M."
	//tDateTime := "11/12/16 4:4:4.012 AM"
	//tDateTime := "11/1/16 4:4:04.012 A.M."
	//tDateTime:= "Monday June 5, 2017 17:24:46.064223400 -0500 CDT"
	//tDateTime := "6-5-2017 17:30:17 -0700 PDT"
	//tDateTime := "11/12/16 4:04:0.012 PM"
	//tDateTime := "11/2/16 04:04:0.012 PM"
	//tDateTime := "11/12/16 4:04:00.012 PM"
	//tDateTime := "11/12/16 04:4:0.012 PM"
	//tDateTime := "11/12/16 04:04:00.012 AM"
	//tDateTime := "11/12/16 04:04:00.012 A.M."
	//tDateTime := "11/12/16 04:4:0.012 P.M."

	// tDateTime := "5/27/2017 11:42PM CDT"
	//tDateTime := "06/1/2017 11:42 -0700 PDT"
	//tDateTime := "2016-11-26 16:26 CDT -0600"
	//tDateTime := "2016/11/26 16:2:3 PDT -0700"
	//tDateTime := "June 12th, 2016 4:26 PM"
	//tDateTime := "05.03.2017"
	//tDateTime := "2017.3.5"
	//tDateTime := "6/27/2017 23:26:01 -0500 CDT"
	//tDateTime := "xxxxxxxxxxxxxxxxxxxxxxxxxxx"
	//tDateTime:= "7-6-16 9:30AM"
	//tDateTime := "12 Nov 2016"
	//tDateTime := "23:26:01 -0500 CDT"
	//tDateTime := "11-26-2016 16:26 -0600 CST"
	//tDateTime:= "Monday June 5th2017 17:24:46.064223400 -0500 CDT"
	//tDateTime := "5/27/2017 11:42PMCDT"
	//tDateTime := "06/1/2017 11:42 PM-0700 PDT"
	//tDateTime:= "June 3rd, 2017 11:42:00PM -0700 PDT"
	//tDateTime:= "2017-6-12 11:26 p.m. Z"
	//tDateTime := "November 12, 11:26pm -0600 CST 2016"
	//tDateTime := "Saturday 11/12/2016 4:26 PM"

	 // common.TestParseDateTimeCreateFormatsInMemory( tDateTime, "")
	common.HammerSampleDateTimes()

// 11/12/16 4:04:0.012 PM
	//FmtDateTimeEverything := "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
	//fmt.Println("Time Now: ", time.Now().Format(FmtDateTimeEverything))
	//common.TestParseSampleDateTimes()
	//common.WriteAllFormatsToFile()
	//common.WriteFormatStatsToFile()

	// 2006-01-02 15:04 -0700 MST
	// 0000024 2006-1-2 15:04 MST -0700

/*
		t, err := time.Parse("15:04:05 -0700 MST", tDateTime)

		if err != nil {
			panic(err)
		}

		fmt.Println("Success!")
		fmt.Println("Original Time: ", tDateTime)
		fmt.Println("  Parsed Time: ", t)
*/

}
