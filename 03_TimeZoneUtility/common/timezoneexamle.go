package common

import (
	"fmt"
	"time"
)

// Tex001 - Test Example 1
func Tex001() {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	tzCDT, _ := time.LoadLocation("America/Chicago")

	testCDT, _ := time.ParseInLocation(fmtstr, tstr, tzCDT)
	fmt.Println("testCDT: ", testCDT)
	fmt.Println("testCDT.Location(): ", testCDT.Location())

}

// Tex002 - Test Example 2
func Tex002() {

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	locCDT, _ := time.LoadLocation("America/Chicago")
	eastEDT, err := time.LoadLocation("America/New_York")
	tzUTC, err := time.LoadLocation("Zulu")
	if err != nil {
		panic(err)
	}

	tCDT, _ := time.Parse(fmtstr, tstr)

	fmt.Println("tCDT: ", tCDT)
	fmt.Println("tCDT.Location: ", tCDT.Location())
	fmt.Println("tCDT LoadLocation Result:", locCDT.String())

	testNewYork := tCDT.In(eastEDT)

	fmt.Println("tCDT in Eastern Time Zone:", testNewYork)
	fmt.Println("tCDT as UTC:", tCDT.In(tzUTC))

}
