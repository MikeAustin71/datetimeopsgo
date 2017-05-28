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
	//tDateTime := "November 12, 2016"
	tDateTime := "2016-11-12 16:26 -0600 CST"
	//tDateTime := "11/12/16 4:26 PM"
	common.TestParseDateTime(tDateTime, "")

}
