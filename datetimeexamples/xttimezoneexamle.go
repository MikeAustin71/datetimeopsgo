package datetimeexamples


import (
	dt "../datetime"
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
	locCDT, _ := time.LoadLocation(dt.TzIanaUsCentral)
	eastEDT, err := time.LoadLocation(dt.TzIanaUsEast)
	tzUTC, err := time.LoadLocation(dt.TzIanaZulu)
	if err != nil {
		panic(err)
	}

	tCDT, _ := time.Parse(fmtstr, tstr)

	fmt.Println("                    tCDT: ", tCDT.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("           tCDT.Location: ", tCDT.Location())
	fmt.Println("tCDT LoadLocation Result: ", locCDT.String())

	testNewYork := tCDT.In(eastEDT)

	fmt.Println("tCDT in Eastern Time Zone:", testNewYork)
	fmt.Println("tCDT as UTC:", tCDT.In(tzUTC))

}

// Tex003 - Test Example 0003 demonstrates use of
// method TimeZoneDto.ConvertTz()
func Tex003() {



	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	tCDT, _ := time.Parse(fmtstr, tstr)
	tzu, err := dt.TimeZoneDto{}.New(tCDT, dt.TzIanaUsPacific, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.New(tCDT, TzIanaUsPacific) tCDT='%v'  Error='%v'\n", tCDT.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	fmt.Println("  Original Input Time: ", tCDT)
	fmt.Println("           tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("   tzu.TimeInLocation: ", tzu.TimeIn.TimeZone.LocationName)
	fmt.Println("          tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("  tzu.TimeOutLocation: ", tzu.TimeOut.TimeZone.LocationName)
	fmt.Println("          tzu.TimeUTC: ", tzu.TimeUTC)
	fmt.Println("        tzu.TimeLocal: ", tzu.TimeLocal)
	fmt.Println("tzu.TimeLocalLocation: ", tzu.TimeLocal.TimeZone.LocationName)

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
	tzu, err :=	dt.TimeZoneDto{}.ConvertTz(tIn, ianaPacificTz, fmtstr)

	if err != nil {
		fmt.Printf("Error returned from TimeZoneDto{}.ConvertTz(tIn, ianaPacificTz) tIn='%v'  Error='%v' \n", tIn.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	fmt.Println(" Original Time String: ", tstr)
	fmt.Println("           tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("   tzu.TimeInLocation: ", tzu.TimeIn.TimeZone.LocationName)
	fmt.Println("          tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("  tzu.TimeOutLocation: ", tzu.TimeOut.TimeZone.LocationName)
	fmt.Println("          tzu.TimeUTC: ", tzu.TimeUTC)
	fmt.Println("        tzu.TimeLocal: ", tzu.TimeLocal)
	fmt.Println("tzu.TimeLocalLocation: ", tzu.TimeLocal.TimeZone.LocationName)
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
	tzu, err := dt.TimeZoneDto{}.ConvertTz(tIn, ianaPacificTz, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by ConvertTz(tIn, ianaPacificTz). tIn='%v'  Error='%v' \n", tIn.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	fmt.Println("         Original Time: ", tIn.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("Original Time Location: ", tIn.Location())
	fmt.Println("            tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("    tzu.TimeInLocation: ", tzu.TimeIn.TimeZone.LocationName)
	fmt.Println("           tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("   tzu.TimeOutLocation: ", tzu.TimeOut.TimeZone.LocationName)
	fmt.Println("           tzu.TimeUTC: ", tzu.TimeUTC)
	fmt.Println("         tzu.TimeLocal: ", tzu.TimeLocal)
	fmt.Println(" tzu.TimeLocalLocation: ", tzu.TimeLocal.TimeZone.LocationName)

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
	tzu, err := dt.TimeZoneDto{}.ConvertTz(tIn, ianaPacificTz, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by ConvertTz(tIn, ianaPacificTz). tIn='%v'  Error='%v' \n", tIn.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	fmt.Println("          Original Time: ", tIn.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println(" Original Time Location: ", tIn.Location())
	fmt.Println("             tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("     tzu.TimeInLocation: ", tzu.TimeIn.TimeZone.LocationName)
	fmt.Println("            tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("    tzu.TimeOutLocation: ", tzu.TimeOut.TimeZone.LocationName)
	fmt.Println("  tzu.TimeOut Zone Name: ", tzu.TimeOut.TimeZone.ZoneName)
	fmt.Println("tzu.TimeOut Zone Offset: ", tzu.TimeOut.TimeZone.ZoneOffsetSeconds)
	fmt.Println("            tzu.TimeUTC: ", tzu.TimeUTC)
	fmt.Println("          tzu.TimeLocal: ", tzu.TimeLocal)
	fmt.Println(" tzu.TimeLocal Location: ", tzu.TimeLocal.TimeZone.LocationName)
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
	tz := dt.TimeZoneDto{}
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
	tz := dt.TimeZoneDto{}

	tCDT, _ := time.Parse(fmtstr, tstr)

	tzuOut, err := tz.ConvertTz(tCDT, dt.TzIanaUsEast, fmtstr)

	if err != nil {
		fmt.Println("TimeZoneDto:ConverTz(tCDT,TzUsEast) returned Error: " + err.Error())
		return
	}


	tzuEast := dt.TimeZoneDto{}
	tzuEast.CopyIn(tzuOut)
	tzuEast.Description = "CDT to Eastern Time Zone Conversion"
	PrintOutTimeZoneFields(tzuEast)

	tzuLocal, err := tz.ConvertTz(tzuEast.TimeOut.DateTime, "Local", fmtstr)

	if err != nil {
		fmt.Println("TimeZoneDto:ConverTz(tzuEast.TimeOut,'Local') returned Error: " + err.Error())
		return
	}

	tzuLocal.Description = "Eastern to Local Conversion using 'Local' tz"
	PrintOutTimeZoneFields(tzuLocal)

}

// TzExampleParseIn0012 - Gives example of two methods for setting
// time in a specific time zone.
func TzExampleParseIn0012() {
	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	t2str := "2014-02-15 19:54:30.038175584"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	neutralFmtStr := "2006-01-02 15:04:05.000000000"
	pacificLoc, err := time.LoadLocation(dt.TzIanaUsPacific)

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(common.TzIanaUsPacific). Error='%v'", err.Error())
		return
	}

	tCentral, _ := time.Parse(fmtstr, t1str)

	tPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, pacificLoc)

	// "01/02/2006 15:04:05.000000000"
	// "02/15/2014 19:54:30.038175584"
	tPacific2, err := time.ParseInLocation(neutralFmtStr, t2str , pacificLoc)

	if err!= nil {
		fmt.Printf("Error returned from time.ParseInLocation(\"01/02/2006 15:04:05.000000000\", \"02/15/2014 19:54:30.038175584\", pacificLoc). Error='%v'", err.Error())
		return

	}

	fmt.Println(" tCentral: ", tCentral.Format(fmtstr))
	fmt.Println(" tPacific: ", tPacific.Format(fmtstr))
	fmt.Println("tPacific2: ", tPacific2.Format(fmtstr))
	fmt.Println()

}

// TimeZoneDefExample014 - Provides example for creating a TimeZoneDefDto
// object.
func TimeZoneDefExample014() {
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	//neutralFmtStr := "2006-01-02 15:04:05.000000000"
	hongKongLoc, _ := time.LoadLocation(dt.TzIanaAsiaHongKong)
	beijingLoc, _ :=time.LoadLocation(dt.TzIanaAsiaShanghai)
	usPacificLoc, _ :=time.LoadLocation(dt.TzIanaUsPacific)

	tHongKong := time.Date(2014, 2, 15, 19, 54, 30, 38175584, hongKongLoc)
	tBeijing := time.Date(2014, 2, 15, 19, 54, 30, 38175584, beijingLoc)
	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)

	tzDef, err := dt.TimeZoneDefDto{}.New(tUsPacific)

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

// Tex021
func Tex021() {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	//tz := TimeZoneDto{}
	tzu, _:= dt.TimeZoneDto{}.New(tIn, ianaPacificTz, fmtstr)

	fmt.Println("Original Time String: ", tstr)
	fmt.Println("          tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("  tzu.TimeInLocation: ", tzu.TimeIn.TimeZone.LocationName)
	fmt.Println("      tzu.TimeInZone: ", tzu.TimeIn.TimeZone.ZoneName)
	fmt.Println("         tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println(" tzu.TimeOutLocation: ", tzu.TimeOut.TimeZone.LocationName)
	fmt.Println("    tzu.TimeOut Zone: ", tzu.TimeOut.TimeZone.ZoneName)
	fmt.Println("         tzu.TimeUTC: ", tzu.TimeUTC)
	fmt.Println("       tzu.TimeLocal: ", tzu.TimeLocal)

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

// TestExample022
func TestExample022() {
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

	tzu := dt.TimeZoneDto{}
	tzuCentral, err := tzu.ConvertTz(tPacificIn, ianaCentralTz, fmtstr)

	if err != nil {
		fmt.Printf("Error from TimeZoneDto.ConvertTz(). Error: %v\n", err.Error())
		return
	}

	centralTOut := tzuCentral.TimeOut.DateTime.Format(fmtstr)

	if centralTime != centralTOut {
		fmt.Printf("Expected tzuCentral.TimeOut %v, got %v\n", centralTime, centralTOut)
		return
	}

	tzuMountain, err := tzu.ConvertTz(tzuCentral.TimeOut.DateTime, ianaMountainTz, fmtstr)

	if err != nil {
		fmt.Printf("Error from  tzuMountain TimeZoneDto.ConvertTz(). Error: %v\n", err.Error())
		return
	}

	mountainTOut := tzuMountain.TimeOut.DateTime.Format(fmtstr)

	if mountainTime != mountainTOut {
		fmt.Printf("Expected tzuMountain.TimeOut %v, got %v\n", mountainTime, mountainTOut)
		return
	}

	tzuPacific, err := tzu.ConvertTz(tzuMountain.TimeOut.DateTime, ianaPacificTz, fmtstr)

	if err != nil {
		fmt.Printf("Error from  tzuMountain TimeZoneDto.ConvertTz(). Error: %v\n", err.Error())
		return
	}

	pacificTOut := tzuPacific.TimeOut.DateTime.Format(fmtstr)

	if pacificTime != pacificTOut {

		fmt.Printf("Expected tzuPacific.TimeOut %v, got %v\n", pacificTime, pacificTOut)
		return
	}

	exTOutLoc := "America/Los_Angeles"

	if exTOutLoc != tzuPacific.TimeOut.TimeZone.LocationName {
		fmt.Printf("Expected tzu.TimeOutLoc %v, got %v.  tzuPacific.TimeOut='%v'\n", exTOutLoc, tzuPacific.TimeOut.TimeZone.LocationName, tzuPacific.TimeOut.DateTime.Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	fmt.Println("Successful Completion!")
}

// TestExampleNewAddDate023
func TestExampleNewAddDate023() {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TzIanaUsPacific, fmtstr)
	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.New(t1, TzUsPacific). Error='%v'", err.Error())
		return
	}

	tzu1OutStrTIn := tzu1.TimeIn.DateTime.Format(fmtstr)

	if t1OutStr != tzu1OutStrTIn {
		fmt.Printf("Error: Expected tzu1OutStrTIn='%v'.  Instead, tzu1OutStrTIn='%v'", t1OutStr, tzu1OutStrTIn )
		return
	}

	t2 := t1.AddDate(3,2, 15)
	t2OutStr := t2.Format(fmtstr)

	tzu2, err := dt.TimeZoneDto{}.NewAddDate(tzu1, 3, 2, 15, fmtstr)

	tzu2OutStrTIn := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		fmt.Printf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
		return
	}

	actualDuration, err := tzu2.Sub(tzu1)

	if err != nil {
		fmt.Printf("Error returned by tzu2.Sub(tzu1). Error='%v'", err.Error())
		return
	}

	expectedDuration := t2.Sub(t1)

	if expectedDuration != actualDuration {
		fmt.Printf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration )
		return
	}

	fmt.Println("Successful Completion")
}

// PrintOut DateTzDto Fields
func PrintOutDateTzDtoFields(dtz dt.DateTzDto) {
	fmt.Println("----------------------------------")
	fmt.Println("           DateTzDto")
	fmt.Println("        Date Time Info")
	fmt.Println("----------------------------------")
	fmt.Println("  Description: ", dtz.Description)
	fmt.Println("         Year: ", dtz.Year)
	fmt.Println("        Month: ", dtz.Month)
	fmt.Println("          Day: ", dtz.Day)
	fmt.Println("         Hour: ", dtz.Hour)
	fmt.Println("       Minute: ", dtz.Minute)
	fmt.Println("       Second: ", dtz.Second)
	fmt.Println("  Millisecond: ", dtz.Millisecond)
	fmt.Println("  Microsecond: ", dtz.Microsecond)
	fmt.Println("   Nanosecond: ", dtz.Nanosecond)
	fmt.Println("TotalNanoSecs: ", dtz.Nanosecond)
	fmt.Println("     DateTime: ", dtz.DateTime.Format(dtz.DateTimeFmt))
	fmt.Println("  DateTimeFmt: ", dtz.DateTimeFmt)
	fmt.Println("----------------------------------")
	fmt.Println("         Time Zone Info")
	fmt.Println("----------------------------------")
	PrintOutTimeZoneDefDtoFields(dtz.TimeZone)

}

// PrintOutTimeZoneDefDtoFields
func PrintOutTimeZoneDefDtoFields(tzDef dt.TimeZoneDefDto) {

	fmt.Println("           ZoneName: ", tzDef.ZoneName)
	fmt.Println("         ZoneOffset: ",tzDef.ZoneOffset)
	fmt.Println("  ZoneOffsetSeconds:", tzDef.ZoneOffsetSeconds)
	fmt.Println("           ZoneSign: ",tzDef.ZoneSign)
	fmt.Println("        OffsetHours: ", tzDef.OffsetHours)
	fmt.Println("      OffsetMinutes: ", tzDef.OffsetMinutes)
	fmt.Println("      OffsetSeconds: ", tzDef.OffsetSeconds)
	fmt.Println("    Location String: ", tzDef.Location.String())
	fmt.Println("      Location Name: ", tzDef.LocationName)
}

// PrintOutTimeZoneFields
func PrintOutTimeZoneFields(tz dt.TimeZoneDto) {
	tz.SetDateTimeFormatStr(dt.FmtDateTimeYrMDayFmtStr)
	fmt.Println()
	fmt.Println("*********************************")
	fmt.Println("     Description: ", tz.Description)
	fmt.Println("          TimeIn: ", tz.TimeIn)
	fmt.Println(" TimeIn Location: ", tz.TimeIn.TimeZone.LocationName)
	fmt.Println("         TimeOut: ", tz.TimeOut)
	fmt.Println("TimeOut Location: ", tz.TimeOut.TimeZone.LocationName)
	fmt.Println("       Time UTC : ", tz.TimeUTC)
	fmt.Println("       TimeLocal: ", tz.TimeLocal)

}

// ExampleDurationLocalUTCTime
func ExampleDurationLocalUTCTime() {
	t1UTCStr := "2017-07-02 22:00:18.423111300 +0000 UTC"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t2LocalStr := "2017-07-02 17:00:18.423111300 -0500 CDT"
	localTzStr := "America/Chicago"

	t1, _ := time.Parse(fmtstr, t1UTCStr)

	tz := dt.TimeZoneDto{}

	tzLocal, _ := tz.ConvertTz(t1, localTzStr, fmtstr)
	t1OutStr := t1.Format(fmtstr)
	t2OutStr := tzLocal.TimeOut.DateTime.Format(fmtstr)

	fmt.Println("  t1UTCStr: ", t1UTCStr)
	fmt.Println("  t1OutStr: ", t1OutStr)
	fmt.Println("  t2OutStr: ", t2OutStr)
	fmt.Println("t2LocalStr: ", t2LocalStr)
}