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

	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	tIn := time.Now()
	tzu := common.TimeZoneUtility{}
	tzu.ConvertTz(tIn, ianaCentralTz, ianaPacificTz)

	fmt.Println("Original Time: ", tIn)
	fmt.Println("Original Time Location: ", tIn.Location())
	fmt.Println("tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("tzu.TimeInLocation: ", tzu.TimeInLoc)
	fmt.Println("tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("tzu.TimeOutLocation: ", tzu.TimeOutLoc)
	fmt.Println("tzu.TimeUTC: ", tzu.TimeUTC)
	/*
		Original Time:  2017-05-13 22:44:08.5476396 -0500 CDT
		Original Time Location:  Local
		tzu.TimeIn:  2017-05-13 22:44:08.5476396 -0500 CDT
		tzu.TimeInLocation:  America/Chicago
		tzu.TimeOut:  2017-05-13 20:44:08.5476396 -0700 PDT
		tzu.TimeOutLocation:  America/Los_Angeles
		tzu.TimeUTC:  2017-05-14 03:44:08.5476396 +0000 UTC
	*/
}
