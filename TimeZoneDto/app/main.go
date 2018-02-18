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

	TestExample031()

}

// TestExample031
func TestExample031() {
	// utcTime := "2017-04-30 00:54:30 +0000 UTC"
	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	mountainTime := "2017-04-29 18:54:30 -0600 MDT"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	ianaMountainTz := "America/Denver"
	tPacificIn, err := time.Parse(fmtstr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n", err.Error())
		return
	}

	tzu := common.TimeZoneDto{}
	tzuCentral, err := tzu.ConvertTz(tPacificIn, ianaCentralTz)

	if err != nil {
		fmt.Printf("Error from TimeZoneDto.ConvertTz(). Error: %v\n", err.Error())
		return
	}

	centralTOut := tzuCentral.TimeOut.Format(fmtstr)

	if centralTime != centralTOut {
		fmt.Printf("Expected tzuCentral.TimeOut %v, got %v\n", centralTime, centralTOut)
		return
	}

	tzuMountain, err := tzu.ConvertTz(tzuCentral.TimeOut, ianaMountainTz)

	if err != nil {
		fmt.Printf("Error from  tzuMountain TimeZoneDto.ConvertTz(). Error: %v\n", err.Error())
		return
	}

	mountainTOut := tzuMountain.TimeOut.Format(fmtstr)

	if mountainTime != mountainTOut {
		fmt.Printf("Expected tzuMountain.TimeOut %v, got %v\n", mountainTime, mountainTOut)
		return
	}

	tzuPacific, err := tzu.ConvertTz(tzuMountain.TimeOut, ianaPacificTz)

	if err != nil {
		fmt.Printf("Error from  tzuMountain TimeZoneDto.ConvertTz(). Error: %v\n", err.Error())
		return
	}

	pacificTOut := tzuPacific.TimeOut.Format(fmtstr)

	if pacificTime != pacificTOut {

		fmt.Printf("Expected tzuPacific.TimeOut %v, got %v\n", pacificTime, pacificTOut)
		return
	}

	exTOutLoc := "America/Los_Angeles"

	if exTOutLoc != tzuPacific.TimeOutZone.LocationName {
		fmt.Printf("Expected tzu.TimeOutLoc %v, got %v.  tzuPacific.TimeOut='%v'\n", exTOutLoc, tzuPacific.TimeOutZone.LocationName, tzuPacific.TimeOut.Format(common.TzDtoYrMDayFmtStr))
		return
	}

	fmt.Println("Successful Completion!")
}