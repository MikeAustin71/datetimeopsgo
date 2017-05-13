package main

import (
	common "MikeAustin71/datetimeopsgo/03_TimeZoneUtility/common"
	"fmt"
	"time"
)

/*
import (
	common "MikeAustin71/datetimeopsgo/03_TimeZoneUtility/common"
	"fmt"
	"time"
)

*/

func main() {
	tzu := common.TimeZoneUtility{}

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	tCDT, _ := time.Parse(fmtstr, tstr)
	tzu.ConvertTz(tCDT, common.TzUsPacific)

	fmt.Println("Original Input Time: ", tCDT)
	fmt.Println("tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("tzu.TimeInLocation: ", tzu.TimeInLoc)
	fmt.Println("tzu.TimeInZoneStr: ", tzu.TimeInZoneStr)
	fmt.Println("tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("tzu.TimeOutZoneStr: ", tzu.TimeOutZoneStr)
	fmt.Println("tzu.TimeUTC: ", tzu.TimeUTC)

	//common.Tex001()
}
