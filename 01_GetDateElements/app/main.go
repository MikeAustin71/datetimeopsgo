package main

import (
	common "MikeAustin71/datetimeopsgo/01_GetDateElements/common"
	"fmt"
	"time"
)

func main() {

	dt3()

}

/*
func dt1()  {
// Get current time in the form of a string
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"

	t, err := time.Parse(fmtstr, tstr)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := common.GetDateTimeStr(t)

	fmt.Println(tstr, "=", t)
	fmt.Println("result=", result)

}
*/

/*
func dt2() {

	dt := common.DateTimeUtility{}
	now := time.Now()

	s := dt.GetDateTimeEverything(now)

	fmt.Println("Time String: ", s)

}
*/

func dt3() {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := common.DateTimeUtility{}
	str := dt.GetDateTimeEverything(testTime)
	fmt.Println("Everything Format: ", str)
	// Saturday April 29, 2017 19:54:30.123456489 -0500 CDT
}
