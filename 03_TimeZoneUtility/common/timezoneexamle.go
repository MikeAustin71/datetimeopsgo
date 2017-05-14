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

// Tex003 - Test Example 0003 demonstrates use of
// method TimeZoneUtility.ConvertTz()
func Tex003() {

	tzu := TimeZoneUtility{}

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	tCDT, _ := time.Parse(fmtstr, tstr)
	tzu.ConvertTz(tCDT, TzUsCentral, TzUsPacific)

	fmt.Println("Original Input Time: ", tCDT)
	fmt.Println("tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("tzu.TimeInLocation: ", tzu.TimeInLoc)
	fmt.Println("tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("tzu.TimeOutLocation: ", tzu.TimeOutLoc)
	fmt.Println("tzu.TimeUTC: ", tzu.TimeUTC)

}

// Tex004 - Example of how to parese a time string and
// assign time zone to the 'Location' of the resulting
// time.Time type
func Tex004() {
	tstr := "04/29/2017 19:54:30"
	//fmtstr := "01/02/2006 15:04:05 -0700 MST"
	fmtstr := "01/02/2006 15:04:05"
	tzCDT, _ := time.LoadLocation("America/Chicago")

	testCDT, _ := time.ParseInLocation(fmtstr, tstr, tzCDT)
	fmt.Println("testCDT: ", testCDT)
	fmt.Println("testCDT.Location(): ", testCDT.Location())

	/*
		testCDT:  2017-04-29 19:54:30 -0500 CDT
		testCDT.Location():  America/Chicago
	*/
}

// Tex005 - demonstrates that you cannot change
// the time zone of time string with ParseInLocation
func Tex005() {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	tzPST, _ := time.LoadLocation("America/Los_Angeles")

	tPST, _ := time.ParseInLocation(fmtstr, tstr, tzPST)

	fmt.Println("Original tstr: ", tstr)
	fmt.Println("tPST: ", tPST)

	/*
		Original tstr:  04/29/2017 19:54:30 -0500 CDT
		tPST:  2017-04-29 19:54:30 -0500 CDT
		--------------------------------------------
		NOTE: - You apparently cannot change the
		time zone of a string using ParseInLocation
		when the time zone is part of the original
		string.
	*/

}

// Tex006 - demonstrates how to specify a time zone
// (aka location) using method ParseInLocation
func Tex006() {

	tstr := "04/29/2017 19:54:30"
	fmtstr := "01/02/2006 15:04:05"
	tzPST, _ := time.LoadLocation("America/Los_Angeles")

	tPST, _ := time.ParseInLocation(fmtstr, tstr, tzPST)

	fmt.Println("Original tstr: ", tstr)
	fmt.Println("tPST: ", tPST)

	/*
		Original tstr:  04/29/2017 19:54:30
		tPST:  2017-04-29 19:54:30 -0700 PDT
		--------------------------------------------
		NOTE: - You can change the time zone of a
		time string using ParseInLocation when
		the time zone is NOT specified as part
		of the original time string.
	*/
}

// Tex007 - Demonstrates how to create a timezone
// neutral time string.
func Tex007() {

	tstr := "04/29/2017 19:54:30 -0700 PDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	fmtstr2 := "01/02/2006 15:04:05"

	tzCST, _ := time.LoadLocation("America/Chicago")

	t1, _ := time.Parse(fmtstr, tstr)
	t2str := t1.Format(fmtstr2)

	t3, _ := time.ParseInLocation(fmtstr2, t2str, tzCST)

	fmt.Println("Original tstr: ", tstr)
	fmt.Println("t1: ", t1)
	fmt.Println("t2str: ", t2str)
	fmt.Println("t3: ", t3)
	fmt.Println("t3 Location: ", t3.Location())

	/*
		Original tstr:  04/29/2017 19:54:30 -0700 PDT
		t1:  2017-04-29 19:54:30 -0700 PDT
		t2str:  04/29/2017 19:54:30
		t3:  2017-04-29 19:54:30 -0500 CDT
		t3 Location:  America/Chicago
		--------------------------------------------
		NOTE: - You can change the time zone of a
		time string using ParseInLocation when
		the time zone is NOT specified as part
		of the original time string.
	*/
}

// Tex008 - Demonstrates method ConvertTz
func Tex008() {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu := TimeZoneUtility{}
	tzu.ConvertTz(tIn, ianaCentralTz, ianaPacificTz)

	fmt.Println("Original Time String: ", tstr)
	fmt.Println("tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("tzu.TimeInLocation: ", tzu.TimeInLoc)
	fmt.Println("tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("tzu.TimeOutLocation: ", tzu.TimeOutLoc)
	fmt.Println("tzu.TimeUTC: ", tzu.TimeUTC)
	/*
		Original Time String:  04/29/2017 19:54:30 -0500 CDT
		tzu.TimeIn:  2017-04-29 19:54:30 -0500 CDT
		tzu.TimeInLocation:  America/Chicago
		tzu.TimeOut:  2017-04-29 17:54:30 -0700 PDT
		tzu.TimeOutLocation:  America/Los_Angeles
		tzu.TimeUTC:  2017-04-30 00:54:30 +0000 UTC
	*/
}

func Tex009()  {
	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	tIn := time.Now()
	tzu := TimeZoneUtility{}
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
