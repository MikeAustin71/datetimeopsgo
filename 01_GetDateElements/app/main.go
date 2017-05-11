package main

import (
	common "MikeAustin71/datetimeopsgo/01_GetDateElements/common"
	"fmt"
	"time"
)

/*
import (
	common "MikeAustin71/datetimeopsgo/01_GetDateElements/common"
	"fmt"
	"time"
)
*/

func main() {
	tstr1 := "04/15/2017 19:54:30.123456489 -0500 CDT"
	tstr2 := "04/18/2017 09:21:16.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, tstr1)
	t2, _ := time.Parse(fmtstr, tstr2)
	dt := common.DateTimeUtility{}
	dur, _ := dt.GetDuration(t1, t2)
	ed, _ := dt.GetElapsedTime(t1, t2)
	fmt.Println("Elapsed Time: ", ed.DurationStr)
	// "2 Days 13 Hours 26 Minutes 46 Seconds 864 Milliseconds 197 Microseconds 832 Nanoseconds"

	fmt.Println("")
	fmt.Println("Default Duration: ", dur)
	// 61h26m46.864197832s
}
