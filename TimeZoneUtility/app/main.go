package main

import (
	"MikeAustin71/datetimeopsgo/TimeZoneUtility/common"

	"time"
	"fmt"
)

/*
import (
	"MikeAustin71/datetimeopsgo/TimeZoneUtility/common"
	"fmt"
	"time"
)

*/

func main() {

	appStartTimeTzu, _ := common.TimeZoneUtility{}.ConvertTz(time.Now().UTC(), "Local")

	dt := common.DateTimeUtility{}
	dateTimeStamp := dt.GetDateTimeStr(appStartTimeTzu.TimeOut)
	logFileNameExt :=  "BaseFileName" + "_" + dateTimeStamp + ".log"
	fmt.Println(" File Name Composite: ", logFileNameExt)
	fmt.Println("     Date Time Stamp: ", dateTimeStamp)
	dateTimeStamp = dt.GetDateTimeNanoSecText(appStartTimeTzu.TimeOut)
	fmt.Println("Full Date Time Stamp: ", dateTimeStamp)

}
