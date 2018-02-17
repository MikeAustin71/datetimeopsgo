package main

import (
	// "MikeAustin71/datetimeopsgo/TimeZoneDto/common"

	"time"
	"fmt"
	"MikeAustin71/datetimeopsgo/TimeZoneDto/common"
)

/*
import (
	"MikeAustin71/datetimeopsgo/TimeZoneDto/common"
	"fmt"
	"time"
)

*/

func main() {

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	//neutralFmtStr := "2006-01-02 15:04:05.000000000"
	hongKongLoc, _ := time.LoadLocation(common.TzIanaAsiaHongKong)
	beijingLoc, _ :=time.LoadLocation(common.TzIanaAsiaShanghai)
	usPacificLoc, _ :=time.LoadLocation(common.TzIanaUsPacific)

	tHongKong := time.Date(2014, 2, 15, 19, 54, 30, 38175584, hongKongLoc)
	tBeijing := time.Date(2014, 2, 15, 19, 54, 30, 38175584, beijingLoc)
	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)

	tzDef, err := common.TimeZoneDefDto{}.New(tUsPacific)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefDto{}.New(tUsPacific). Error='%v'", err.Error())
		return
	}


	fmt.Println(" tHongKong: ", tHongKong.Format(fmtstr))
	fmt.Println("  tBeijing: ", tBeijing.Format(fmtstr))
	fmt.Println("----------------------------------------")
	fmt.Println("       tUsPacific: ", tUsPacific.Format(fmtstr))
	fmt.Println("         ZoneName: ", tzDef.ZoneName)
	fmt.Println("       ZoneOffset: ",tzDef.ZoneOffset)
	fmt.Println("ZoneOffsetSeconds:", tzDef.ZoneOffsetSeconds)
	fmt.Println("         ZoneSign: ",tzDef.ZoneSign)
	fmt.Println("      OffsetHours: ", tzDef.OffsetHours)
	fmt.Println("    OffsetMinutes: ", tzDef.OffsetMinutes)
	fmt.Println("    OffsetSeconds: ", tzDef.OffsetSeconds)
	fmt.Println("  Location String: ", tzDef.Location.String())
	fmt.Println("    Location Name: ", tzDef.LocationName)




}
