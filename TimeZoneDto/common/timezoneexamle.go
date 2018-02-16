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
	locCDT, _ := time.LoadLocation(TzUsCentral)
	eastEDT, err := time.LoadLocation(TzUsEast)
	tzUTC, err := time.LoadLocation(TzUTC)
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
// method TimeZoneDto.ConvertTz()
func Tex003() {

	tzu := TimeZoneDto{}

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	tCDT, _ := time.Parse(fmtstr, tstr)
	tzu.ConvertTz(tCDT, TzUsPacific)

	fmt.Println("Original Input Time: ", tCDT)
	fmt.Println("tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("tzu.TimeInLocation: ", tzu.TimeInLoc)
	fmt.Println("tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("tzu.TimeOutLocation: ", tzu.TimeOutLoc)
	fmt.Println("tzu.TimeUTC: ", tzu.TimeUTC)

}

// Tex004 - Example of how to parse a time string and
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
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu := TimeZoneDto{}
	tzu.ConvertTz(tIn, ianaPacificTz)

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

// Tex009 - Demonstrates Time Conversion from
// one time zone to another
func Tex009() {
	ianaPacificTz := "America/Los_Angeles"
	tIn := time.Now()
	tzu := TimeZoneDto{}
	tzu.ConvertTz(tIn, ianaPacificTz)

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

// Tex010 - Demonstrates the Zone() function
func Tex010() {
	ianaPacificTz := "America/Los_Angeles"
	tIn := time.Now()
	tzu := TimeZoneDto{}
	tzu.ConvertTz(tIn, ianaPacificTz)

	fmt.Println("Original Time: ", tIn)
	fmt.Println("Original Time Location: ", tIn.Location())
	fmt.Println("tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("tzu.TimeInLocation: ", tzu.TimeInLoc)
	fmt.Println("tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("tzu.TimeOutLocation: ", tzu.TimeOutLoc)
	zoneName, offsetInt := tzu.TimeOut.Zone()
	// Note: offsetInt is seconds east of UTC
	fmt.Println("tzu.TimeOut Zone Name: ", zoneName)
	fmt.Println("tzu.TimeOut Zone Offset: ", offsetInt)

	fmt.Println("tzu.TimeUTC: ", tzu.TimeUTC)

	/*
		Original Time:  2017-05-14 22:27:42.2495266 -0500 CDT
		Original Time Location:  Local
		tzu.TimeIn:  2017-05-14 22:27:42.2495266 -0500 CDT
		tzu.TimeInLocation:  America/Chicago
		tzu.TimeOut:  2017-05-14 20:27:42.2495266 -0700 PDT
		tzu.TimeOutLocation:  America/Los_Angeles
		tzu.TimeOut Zone Name:  PDT
		tzu.TimeOut Zone Offset:  -25200
		tzu.TimeUTC:  2017-05-15 03:27:42.2495266 +0000 UTC
	*/
}

// ReclassifyTimeAsLocal
func ReclassifyTimeAsLocal() {
	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tz := TimeZoneDto{}
	tIn, err := time.Parse(fmtstr, tPacific)
	if err != nil {
		fmt.Printf("Error returned from time.Parse: %v", err.Error())
		return
	}

	tOut, err := tz.ReclassifyTimeWithNewTz(tIn, "Local")

	tstrOut := tOut.Format(fmtstr)
	fmt.Println("Example Method: ReclassifyTimeWithNewTz()")
	fmt.Println("Input Time : ", tPacific)
	fmt.Println("Output Time: ", tstrOut)
	fmt.Println("Output Time Location: ", tOut.Location())

}

// Tex011
func Tex011() {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	tz := TimeZoneDto{}

	tCDT, _ := time.Parse(fmtstr, tstr)

	tzuOut, err := tz.ConvertTz(tCDT, TzUsEast)

	if err != nil {
		fmt.Println("TimeZoneDto:ConverTz(tCDT,TzUsEast) returned Error: " + err.Error())
		return
	}


	tzuEast := TimeZoneDto{}
	tzuEast.CopyToThis(tzuOut)
	tzuEast.Description = "CDT to Eastern Time Zone Conversion"
	printOutTimeZoneFields(tzuEast)

	tzuLocal, err := tz.ConvertTz(tzuEast.TimeOut, "Local")

	if err != nil {
		fmt.Println("TimeZoneDto:ConverTz(tzuEast.TimeOut,'Local') returned Error: " + err.Error())
		return
	}

	tzuLocal.Description = "Eastern to Local Conversion using 'Local' tz"
	printOutTimeZoneFields(tzuLocal)

}

// Tex021
func Tex021() {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	//tz := TimeZoneDto{}
	tzu, _:= TimeZoneDto{}.New(tIn, ianaPacificTz)

	fmt.Println("Original Time String: ", tstr)
	fmt.Println("          tzu.TimeIn: ", tzu.TimeIn)
	tinLoc, _ := tzu.GetLocationIn()
	fmt.Println("  tzu.TimeInLocation: ", tinLoc)
	tinZone, _ := tzu.GetZoneIn()
	fmt.Println("      tzu.TimeInZone: ", tinZone)
	fmt.Println("         tzu.TimeOut: ", tzu.TimeOut)
	toutLoc, _ := tzu.GetLocationOut()
	fmt.Println("tzu.TimeOutLocation: ", toutLoc)
	toutZone, _ := tzu.GetZoneOut()
	fmt.Println("    tzu.TimeOutZone: ", toutZone)
	fmt.Println("        tzu.TimeUTC: ", tzu.TimeUTC)

	/*
		Original Time String:  04/29/2017 19:54:30 -0500 CDT
							tzu.TimeIn:  2017-04-29 19:54:30 -0500 CDT
			tzu.TimeInLocation:  Local
					tzu.TimeInZone:  CDT
						 tzu.TimeOut:  2017-04-29 17:54:30 -0700 PDT
		tzu.TimeOutLocation:  America/Los_Angeles
				tzu.TimeOutZone:  PDT
						tzu.TimeUTC:  2017-04-30 00:54:30 +0000 UTC
	*/
}


func printOutTimeZoneFields(tz TimeZoneDto) {

	fmt.Println()
	fmt.Println("*********************************")
	fmt.Println("Description: ", tz.Description)
	fmt.Println("TimeIn: ", tz.TimeIn)
	fmt.Println("TimeIn Location: ", tz.TimeInLoc.String())
	fmt.Println("TimeOut: ", tz.TimeOut)
	fmt.Println("TimeOut Location: ", tz.TimeOutLoc.String())
	fmt.Println("Time UTC :", tz.TimeUTC)

}

// ExampleDurationLocalUTCTime
func ExampleDurationLocalUTCTime() {
	t1UTCStr := "2017-07-02 22:00:18.423111300 +0000 UTC"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t2LocalStr := "2017-07-02 17:00:18.423111300 -0500 CDT"
	localTzStr := "America/Chicago"

	t1, _ := time.Parse(fmtstr, t1UTCStr)

	tz := TimeZoneDto{}

	tzLocal, _ := tz.ConvertTz(t1, localTzStr)
	t1OutStr := t1.Format(fmtstr)
	t2OutStr := tzLocal.TimeOut.Format(fmtstr)

	fmt.Println("  t1UTCStr: ", t1UTCStr)
	fmt.Println("  t1OutStr: ", t1OutStr)
	fmt.Println("  t2OutStr: ", t2OutStr)
	fmt.Println("t2LocalStr: ", t2LocalStr)
}