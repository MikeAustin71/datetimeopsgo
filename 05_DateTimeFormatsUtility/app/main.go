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
	//tDateTime := "November 12, 2016"
	tDateTime := "2016-11-26 16:26 -0600"
	//tDateTime := "11/12/16 4:26 PM"
	// tDateTime := "5/27/2017 11:42PM CDT"
	common.TestParseDateTime(tDateTime, "")

}
