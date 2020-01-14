package datetimeexamples

import (
	"fmt"
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
	"strings"
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
	locCDT, _ := time.LoadLocation(dt.TZones.US.Central())
	eastEDT, err := time.LoadLocation(dt.TZones.US.Eastern())
	tzUTC, err := time.LoadLocation(dt.TZones.Zulu())
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
	tzu, err := dt.TimeZoneDto{}.New(tCDT, dt.TZones.US.Pacific(), dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.New(tCDT, TZones.US.Pacific()) tCDT='%v'  Error='%v'\n", tCDT.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	fmt.Println("  Original Input Time: ", tCDT)
	fmt.Println("           tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("   tzu.TimeInLocation: ", tzu.TimeIn.GetTimeZoneName())
	fmt.Println("          tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("  tzu.TimeOutLocation: ", tzu.TimeOut.GetTimeZoneName())
	fmt.Println("          tzu.TimeUTC: ", tzu.TimeUTC)
	fmt.Println("        tzu.TimeLocal: ", tzu.TimeLocal)
	fmt.Println("tzu.TimeLocalLocation: ", tzu.TimeLocal.GetTimeZoneName())

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
	tzu, err := dt.TimeZoneDto{}.ConvertTz(tIn, ianaPacificTz, fmtstr)

	if err != nil {
		fmt.Printf("Error returned from TimeZoneDto{}.ConvertTz(tIn, ianaPacificTz) tIn='%v'  Error='%v' \n", tIn.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	fmt.Println(" Original Time String: ", tstr)
	fmt.Println("           tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("   tzu.TimeInLocation: ", tzu.TimeIn.GetTimeZoneName())
	fmt.Println("          tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("  tzu.TimeOutLocation: ", tzu.TimeOut.GetTimeZoneName())
	fmt.Println("          tzu.TimeUTC: ", tzu.TimeUTC)
	fmt.Println("        tzu.TimeLocal: ", tzu.TimeLocal)
	fmt.Println("tzu.TimeLocalLocation: ", tzu.TimeLocal.GetTimeZoneName())
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
	fmt.Println("    tzu.TimeInLocation: ", tzu.TimeIn.GetTimeZoneName())
	fmt.Println("           tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("   tzu.TimeOutLocation: ", tzu.TimeOut.GetTimeZoneName())
	fmt.Println("           tzu.TimeUTC: ", tzu.TimeUTC)
	fmt.Println("         tzu.TimeLocal: ", tzu.TimeLocal)
	fmt.Println(" tzu.TimeLocalLocation: ", tzu.TimeLocal.GetTimeZoneName())

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

	tzIn := tzu.TimeIn.GetTimeZone()
	tzOut := tzu.TimeOut.GetTimeZone()
	tzLocal := tzu.TimeLocal.GetTimeZone()

	fmt.Println("          Original Time: ", tIn.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println(" Original Time Location: ", tIn.Location())
	fmt.Println("             tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("       Time In Location: ", tzIn.GetOriginalLocationName())
	fmt.Println("            tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println("      Time Out Location: ", tzOut.GetOriginalLocationName())
	fmt.Println("     Time Out Zone Name: ", tzOut.GetZoneName())
	fmt.Println("   Time Out Zone Offset: ", tzOut.GetZoneOffset())
	fmt.Println("            tzu.TimeUTC: ", tzu.TimeUTC)
	fmt.Println("          tzu.TimeLocal: ", tzu.TimeLocal)
	fmt.Println("    Time Local Location: ", tzLocal.GetOriginalLocationName())
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

	if err != nil {
		fmt.Printf("Error returned by tz.ReclassifyTimeWithNewTz(tIn, \"Local\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

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

	tzuOut, err := tz.ConvertTz(tCDT, dt.TZones.US.Eastern(), fmtstr)

	if err != nil {
		fmt.Println("TimeZoneDto:ConvertTz(tCDT,TzUsEast) returned Error: " + err.Error())
		return
	}

	tzuEast := dt.TimeZoneDto{}
	tzuEast.CopyIn(tzuOut)
	tzuEast.Description = "CDT to Eastern Time Zone Conversion"
	PrintOutTimeZoneFields(tzuEast)

	tzuLocal, err := tz.ConvertTz(tzuEast.TimeOut.GetDateTimeValue(), "Local", fmtstr)

	if err != nil {
		fmt.Println("TimeZoneDto:ConvertTz(tzuEast.TimeOut,'Local') returned Error: " + err.Error())
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
	pacificLoc, err := time.LoadLocation(dt.TZones.US.Pacific())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(common.TZones.US.Pacific()). Error='%v'", err.Error())
		return
	}

	tCentral, _ := time.Parse(fmtstr, t1str)

	tPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, pacificLoc)

	// "01/02/2006 15:04:05.000000000"
	// "02/15/2014 19:54:30.038175584"
	tPacific2, err := time.ParseInLocation(neutralFmtStr, t2str, pacificLoc)

	if err != nil {
		fmt.Printf("Error returned from time.ParseInLocation(\"01/02/2006 15:04:05.000000000\", \"02/15/2014 19:54:30.038175584\", pacificLoc). Error='%v'", err.Error())
		return

	}

	fmt.Println(" tCentral: ", tCentral.Format(fmtstr))
	fmt.Println(" tPacific: ", tPacific.Format(fmtstr))
	fmt.Println("tPacific2: ", tPacific2.Format(fmtstr))
	fmt.Println()

}

// TimeZoneDefExample014 - Provides example for creating a TimeZoneDefinition
// object.
func TimeZoneDefExample014() {
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	//neutralFmtStr := "2006-01-02 15:04:05.000000000"
	hongKongLoc, _ := time.LoadLocation(dt.TZones.Asia.Hong_Kong())
	beijingLoc, _ := time.LoadLocation(dt.TZones.Asia.Shanghai())
	usPacificLoc, _ := time.LoadLocation(dt.TZones.US.Pacific())

	tHongKong := time.Date(2014, 2, 15, 19, 54, 30, 38175584, hongKongLoc)
	tBeijing := time.Date(2014, 2, 15, 19, 54, 30, 38175584, beijingLoc)
	tUsPacific := time.Date(2014, 2, 15, 19, 54, 30, 38175584, usPacificLoc)

	tzDef, err := dt.TimeZoneDefinition{}.New(tUsPacific)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefinition{}.New(tUsPacific). Error='%v'", err.Error())
		return
	}

	offsetSignChar,
	_,
	offsetHours,
	offsetMinutes,
	offsetSeconds := tzDef.GetOriginalOffsetElements()

	fmt.Println(" tHongKong: ", tHongKong.Format(fmtstr))
	fmt.Println("  tBeijing: ", tBeijing.Format(fmtstr))
	fmt.Println("----------------------------------------")
	fmt.Println("       tUsPacific: ", tUsPacific.Format(fmtstr))
	fmt.Println("         ZoneName: ", tzDef.GetZoneName())
	fmt.Println("       ZoneOffset: ", tzDef.GetZoneOffset())
	fmt.Println("       UTC Offset: ", tzDef.GetOriginalUtcOffset())
	fmt.Println("ZoneOffsetSeconds:", tzDef.GetZoneOffsetSeconds())
	fmt.Println("         ZoneSign: ", offsetSignChar)
	fmt.Println("      OffsetHours: ", offsetHours)
	fmt.Println("    OffsetMinutes: ", offsetMinutes)
	fmt.Println("    OffsetSeconds: ", offsetSeconds)
	fmt.Println("  Location String: ", tzDef.GetOriginalLocationPtr().String())
	fmt.Println("    Location Name: ", tzDef.GetOriginalLocationName())

}

// Tex021
func Tex021() {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	//tz := TimeZoneDto{}
	tzu, _ := dt.TimeZoneDto{}.New(tIn, ianaPacificTz, fmtstr)

	tzIn := tzu.TimeIn.GetTimeZone()
	tzOut := tzu.TimeOut.GetTimeZone()

	fmt.Println("Original Time String: ", tstr)
	fmt.Println("          tzu.TimeIn: ", tzu.TimeIn)
	fmt.Println("  tzu.TimeInLocation: ", tzIn.GetOriginalLocationName())
	fmt.Println("      tzu.TimeInZone: ", tzIn.GetZoneName())
	fmt.Println("         tzu.TimeOut: ", tzu.TimeOut)
	fmt.Println(" tzu.TimeOutLocation: ", tzOut.GetOriginalLocationName())
	fmt.Println("    tzu.TimeOut Zone: ", tzOut.GetZoneName())
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

	centralTOut := tzuCentral.TimeOut.GetDateTimeValue().Format(fmtstr)

	if centralTime != centralTOut {
		fmt.Printf("Expected tzuCentral.TimeOut %v, got %v\n", centralTime, centralTOut)
		return
	}

	tzuMountain, err := tzu.ConvertTz(tzuCentral.TimeOut.GetDateTimeValue(), ianaMountainTz, fmtstr)

	if err != nil {
		fmt.Printf("Error from  tzuMountain TimeZoneDto.ConvertTz(). Error: %v\n", err.Error())
		return
	}

	mountainTOut := tzuMountain.TimeOut.GetDateTimeValue().Format(fmtstr)

	if mountainTime != mountainTOut {
		fmt.Printf("Expected tzuMountain.TimeOut %v, got %v\n", mountainTime, mountainTOut)
		return
	}

	tzuPacific, err := tzu.ConvertTz(tzuMountain.TimeOut.GetDateTimeValue(), ianaPacificTz, fmtstr)

	if err != nil {
		fmt.Printf("Error from  tzuMountain TimeZoneDto.ConvertTz(). Error: %v\n", err.Error())
		return
	}

	pacificTOut := tzuPacific.TimeOut.GetDateTimeValue().Format(fmtstr)

	if pacificTime != pacificTOut {

		fmt.Printf("Expected tzuPacific.TimeOut %v, got %v\n", pacificTime, pacificTOut)
		return
	}

	exTOutLoc := "America/Los_Angeles"

	if exTOutLoc != tzuPacific.TimeOut.GetTimeZoneName() {
		fmt.Printf("Expected tzu.TimeOutLoc='%v'.\n" +
			"Instead tzu.TimeOutLoc='%v'.\n" +
			"tzuPacific.TimeOut='%v'\n",
			exTOutLoc, tzuPacific.TimeOut.GetTimeZoneName(),
			tzuPacific.TimeOut.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
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
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtstr)
	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.New(t1, TzUsPacific). Error='%v'", err.Error())
		return
	}

	tzu1OutStrTIn := tzu1.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t1OutStr != tzu1OutStrTIn {
		fmt.Printf("Error: Expected tzu1OutStrTIn='%v'.  Instead, tzu1OutStrTIn='%v'", t1OutStr, tzu1OutStrTIn)
		return
	}

	t2 := t1.AddDate(3, 2, 15)
	t2OutStr := t2.Format(fmtstr)

	tzu2, err := dt.TimeZoneDto{}.NewAddDate(tzu1, 3, 2, 15, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDto{}.NewAddDate(tzu1, 3, 2, 15, fmtstr)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzu2OutStrTIn := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

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
		fmt.Printf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration)
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
	fmt.Println("  Description: ", dtz.GetDescription())
	fmt.Println("         Year: ", dtz.GetTimeComponents().Years)
	fmt.Println("        Month: ", dtz.GetTimeComponents().Months)
	fmt.Println("        Weeks: ", dtz.GetTimeComponents().Weeks)
	fmt.Println("      WeekDay: ", dtz.GetTimeComponents().WeekDays)
	fmt.Println("      DateDay: ", dtz.GetTimeComponents().DateDays)
	fmt.Println("         Hour: ", dtz.GetTimeComponents().Hours)
	fmt.Println("       Minute: ", dtz.GetTimeComponents().Minutes)
	fmt.Println("       Second: ", dtz.GetTimeComponents().Seconds)
	fmt.Println("  Millisecond: ", dtz.GetTimeComponents().Milliseconds)
	fmt.Println("  Microsecond: ", dtz.GetTimeComponents().Microseconds)
	fmt.Println("   Nanosecond: ", dtz.GetTimeComponents().Nanoseconds)
	fmt.Println("TotalNanoSecs: ", dtz.GetTimeComponents().Nanoseconds)
	fmt.Println("     DateTime: ", dtz.GetDateTimeValue().Format(dtz.GetDateTimeFmt()))
	fmt.Println("  DateTimeFmt: ", dtz.GetDateTimeFmt())
	fmt.Println("----------------------------------")
	fmt.Println("         Time Zone Info")
	fmt.Println("----------------------------------")
	PrintOutTimeZoneDefFields(dtz.GetTimeZone())

}

// PrintOutTimeZoneDefFields - Prints elements included in
// TimeZoneDefinition input parameter.
func PrintOutTimeZoneDefFields(tzDef dt.TimeZoneDefinition) {

	var milTzLetter, milTzName string
	var err error

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	milTzLetter, err = tzDef.GetMilitaryTimeZoneLetter()

	if err != nil {
		milTzLetter = "Empty String"
	}

	milTzName, err = tzDef.GetMilitaryTimeZoneName()

	if err != nil {
		milTzName = "Empty String"
	}

	offsetSignChar,
	_,
	offsetHours,
	offsetMinutes,
	offsetSeconds := tzDef.GetOriginalOffsetElements()

	fmt.Println("-----------------------------------------------")
	fmt.Println("              Original Time Zone               ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("      Reference Date Time: ", tzDef.GetOriginalDateTime().Format(fmtStr))
	fmt.Println("                 ZoneName: ", tzDef.GetZoneName())
	fmt.Println("               ZoneOffset: ", tzDef.GetZoneOffset())
	fmt.Println("               UTC Offset: ", tzDef.GetOriginalUtcOffset())
	fmt.Println("        ZoneOffsetSeconds:", tzDef.GetZoneOffsetSeconds())
	fmt.Println("                 ZoneSign: ", offsetSignChar)
	fmt.Println("              OffsetHours: ", offsetHours)
	fmt.Println("            OffsetMinutes: ", offsetMinutes)
	fmt.Println("            OffsetSeconds: ", offsetSeconds)
	fmt.Println("          Location String: ", tzDef.GetOriginalLocationPtr().String())
	fmt.Println("            Location Name: ", tzDef.GetOriginalLocationName())
	fmt.Println("       Location Name Type: ", tzDef.GetOriginalLocationNameType().String())
	fmt.Println("Military Time Zone Letter: ", milTzLetter)
	fmt.Println("  Military Time Zone Name: ", milTzName)
	fmt.Println("          Tag Description: ", tzDef.GetOriginalTagDescription())
	fmt.Println( "          Time Zone Type: ", tzDef.GetTimeZoneType().String())
	fmt.Println("-----------------------------------------------")
	fmt.Println("                Convertible Time Zone          ")
	fmt.Println("-----------------------------------------------")

}

// PrintOutTimeZoneFields
func PrintOutTimeZoneFields(tz dt.TimeZoneDto) {
	tz.SetDateTimeFormatStr(dt.FmtDateTimeYrMDayFmtStr)
	fmt.Println()
	fmt.Println("*********************************")
	fmt.Println("     Description: ", tz.Description)
	fmt.Println("          TimeIn: ", tz.TimeIn.String())
	fmt.Println("      TimeIn Time Zone")
	PrintOutTimeZoneDefFields(tz.TimeIn.GetTimeZone())
	fmt.Println()
	fmt.Println("         TimeOut: ", tz.TimeOut.String())
	fmt.Println("      TimeOut Time Zone")
	PrintOutTimeZoneDefFields(tz.TimeOut.GetTimeZone())
	fmt.Println()
	fmt.Println("       Time UTC : ", tz.TimeUTC.String())
	fmt.Println("      TimeUTC Time Zone")
	PrintOutTimeZoneDefFields(tz.TimeUTC.GetTimeZone())
	fmt.Println()
	fmt.Println("       TimeLocal: ", tz.TimeLocal.String())
	fmt.Println("      TimeLocal Time Zone")
	PrintOutTimeZoneDefFields(tz.TimeLocal.GetTimeZone())
	fmt.Println()

}

// PrintOutTimeDtoFields
func PrintOutTimeDtoFields(tDto dt.TimeDto) {
	fmt.Println("========================================")
	fmt.Println("          TimeDto Printout")
	fmt.Println("========================================")
	fmt.Println("                   Years: ", tDto.Years)
	fmt.Println("                  Months: ", tDto.Months)
	fmt.Println("                   Weeks: ", tDto.Weeks)
	fmt.Println("                WeekDays: ", tDto.WeekDays)
	fmt.Println("                DateDays: ", tDto.DateDays)
	fmt.Println("                   Hours: ", tDto.Hours)
	fmt.Println("                 Minutes: ", tDto.Minutes)
	fmt.Println("                 Seconds: ", tDto.Seconds)
	fmt.Println("            Milliseconds: ", tDto.Milliseconds)
	fmt.Println("            Microseconds: ", tDto.Microseconds)
	fmt.Println("             Nanoseconds: ", tDto.Nanoseconds)
	fmt.Println("Total SubSec Nanoseconds: ", tDto.TotSubSecNanoseconds)
	fmt.Println("  Total Time Nanoseconds: ", tDto.TotTimeNanoseconds)
	fmt.Println("========================================")

}

func PrintOutDateTimeTimeZoneFields(dt time.Time, dtLabel string) {
	fmt.Println("=========================================")
	fmt.Println("     Time Zone Fields From time.Time     ")
	fmt.Println("=========================================")

	zoneName, zoneOffsetSeconds := dt.Zone()

	locationPtr := dt.Location()

	var locationName, locationPtrStr,
	  offsetSignStr string

	if locationPtr == nil {
		locationName = ""
		locationPtrStr = ""
	} else {
		locationName = dt.Location().String()
		locationPtrStr = locationName
	}

	var zoneSign, offsetHours, offsetMinutes,
	offsetSeconds, totalOffsetSeconds int


	totalOffsetSeconds = zoneOffsetSeconds

	if zoneOffsetSeconds < 0 {
		zoneSign = -1
		offsetSignStr = "-"
	} else {
		zoneSign = 1
		offsetSignStr = "+"
	}

	unSignedZoneOffsetSeconds := zoneOffsetSeconds

	unSignedZoneOffsetSeconds *= zoneSign

	if unSignedZoneOffsetSeconds != 0 {
		offsetHours = unSignedZoneOffsetSeconds / 3600 // compute hours
		unSignedZoneOffsetSeconds -= offsetHours * 3600

		if unSignedZoneOffsetSeconds != 0 {
			offsetMinutes = unSignedZoneOffsetSeconds / 60 // compute minutes
			unSignedZoneOffsetSeconds -= offsetMinutes * 60
		}

		offsetSeconds = unSignedZoneOffsetSeconds
	}
	utcOffset := fmt.Sprintf("UTC" + offsetSignStr +
		"%02d%02d", offsetHours, offsetMinutes)

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	/*
	abbreviationLookUp := fmt.Sprintf(
		zoneName + offsetSignStr +
			"%02d%02d", offsetHours, offsetMinutes)
*/

	abbreviationLookUp := zoneName + utcOffset[3:]

	fieldLen := len("    Total Offset Seconds:")

	lenDtLabel := len(dtLabel) + 2

	if lenDtLabel == 0 {
		dtLabel = "        Input Date Time: "
	} else if lenDtLabel >= fieldLen {
		dtLabel += ": "
	} else {
		xSpacer := strings.Repeat(" ", fieldLen - lenDtLabel)
		dtLabel = xSpacer + dtLabel + ": "
	}

	fmt.Print()
	fmt.Println(dtLabel, dt.Format(fmtStr))
	fmt.Println("              Zone Name: ", zoneName)
	fmt.Println("    Abbreviation Lookup: ", abbreviationLookUp)
	fmt.Println("              Zone Sign: ", offsetSignStr)
	fmt.Println("      Zone Offset Hours: ", offsetHours)
	fmt.Println("    Zone Offset Minutes: ", offsetMinutes)
	fmt.Println("    Zone Offset Seconds: ", offsetSeconds)
	fmt.Println("   Total Offset Seconds: ", totalOffsetSeconds)
	fmt.Println("             UTC Offset: ", utcOffset)
	fmt.Println("       Location Pointer: ", locationPtrStr)
	fmt.Println("          Location Name: ", locationName)
	fmt.Println("=========================================")
	fmt.Println()

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
	t2OutStr := tzLocal.TimeOut.GetDateTimeValue().Format(fmtstr)

	fmt.Println("  t1UTCStr: ", t1UTCStr)
	fmt.Println("  t1OutStr: ", t1OutStr)
	fmt.Println("  t2OutStr: ", t2OutStr)
	fmt.Println("t2LocalStr: ", t2LocalStr)
}
