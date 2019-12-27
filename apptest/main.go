package main

import (
	"fmt"
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
	ex "github.com/MikeAustin71/datetimeopsgo/datetimeexamples"
	"strings"
	"time"
)

func main() {

	mainTest{}.mainTest039()


}

type mainTest struct {
	input  string
	output string
}

func (mt mainTest) mainTest039() {

	ePrefix := "mainTest039()"

	mt.mainPrintHdr(ePrefix , "-")

	tzName := "Australia/Adelaide"

	tzLocPtr, err := time.LoadLocation(tzName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzName)\n" +
			"tzName='%v'\n" +
			"Error='%v'\n", tzName, err.Error())
		return
	}

	t1 := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		tzLocPtr)

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	tzDef, err := dt.TimeZoneDefDto{}.New(t1)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefDto{}.New(t1)\n" +
			"t1='%v'\n" +
			"Error='%v'\n", t1.Format(fmtStr))
		return
	}

	fmt.Println("t1 Date Time: ", t1.Format(fmtStr))
	ex.PrintOutTimeZoneDefDtoFields(tzDef)

	t2 := time.Date(
		2019,
		time.Month(12),
		30,
		22,
		58,
		32,
		0,
		tzLocPtr)

	tzDef, err = dt.TimeZoneDefDto{}.New(t2)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefDto{}.New(t2)\n" +
			"t2='%v'\n" +
			"Error='%v'\n", t1.Format(fmtStr))
		return
	}

	fmt.Println("t2 Date Time: ", t2.Format(fmtStr))

	ex.PrintOutTimeZoneDefDtoFields(tzDef)
}

func (mt mainTest) mainTest038() {

	ePrefix := "mainTest038()"

	mt.mainPrintHdr(ePrefix , "-")

	t1str := "06/30/2019 22:58:32.000000000 -0400 EDT"
	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	tZoneName := dt.TZones.Other.EST05EDT()

	easternLocPtr, err := time.LoadLocation(tZoneName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.US.Central())\n" +
			"tZoneName='%v'\n" +
			"Error='%v'\n", tZoneName, err.Error())
		return
	}

	t1, err := time.Parse(fmtStr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str)\n" +
			"t1str='%v'\n" +
			"Error='%v'\n", t1str, err.Error())
		return
	}

	t2 := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		easternLocPtr)

	fmt.Println("t1 Date Time: ", t1.Format(fmtStr))
	fmt.Println("t2 Date Time: ", t2.Format(fmtStr))


}

func (mt mainTest) mainTest037() {

	ePrefix := "mainTest037()"

	mt.mainPrintHdr(ePrefix , "-")

	tzName := "EST"

	tzNameLocPtr, err := time.LoadLocation(tzName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzName).\n" +
			"tzName='%v'\n" +
			"Error='%v'\n", tzName, err.Error())
		return
	}

	t1 := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		tzNameLocPtr)

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	fmt.Println("t1 Date Time: ", t1.Format(fmtStr))

	tz2Name := dt.TZones.America.New_York()

	tz2NameLocPtr, err := time.LoadLocation(tz2Name)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzName).\n" +
			"tzName='%v'\n" +
			"Error='%v'\n", tzName, err.Error())
		return
	}

	t2 := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		tz2NameLocPtr)

	fmt.Println("t2 Date Time: ", t2.Format(fmtStr))
}

func (mt mainTest) mainTest036() {

	ePrefix := "mainTest036()"

	mt.mainPrintHdr(ePrefix , "-")

	t1str := "06/30/2019 22:58:32.000000000 -0400 EDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtstr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str)\n" +
			"t1str='%v'\n" +
			"Error='%v'\n", t1str, err.Error())
		return
	}

	tzName := t1.Location().String()

	fmt.Println("t1 date Time String: ", t1str)
	fmt.Println("       t1 date time: ", t1.Format(fmtstr))
	fmt.Println("  t1 Time Zone Name: ", tzName)
	fmt.Println()

/*
	locPtr, err := time.LoadLocation(tzName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzName).\n" +
			"tzName='%v'\n" +
			"Error='%v'\n", tzName, err.Error())
		return
	}
*/

	t2str := "06/30/2019 22:58:32.000000000 -0500 CDT"

	t2, err := time.Parse(fmtstr, t2str)

	if err != nil {
		fmt.Printf("Error returned from time.Parse(fmtstr, t2str)\n" +
			"t2str='%v'\n", t2str, err.Error())
		return
	}

	t3 := t2.In(t1.Location())
	tzName = t3.Location().String()

	fmt.Println("t2 date Time String: ", t2.Format(fmtstr))
	fmt.Println("       t3 date time: ", t3.Format(fmtstr))
	fmt.Println("  t3 Time Zone Name: ", tzName)
	fmt.Println()

}

func (mt mainTest) mainTest035() {
	ePrefix := "mainTest035()"

	mt.mainPrintHdr(ePrefix , "-")

	t1str := "06/30/2019 22:58:32.000000000 -0400 EDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	tZoneName := dt.TZones.US.Eastern()

	easternLocPtr, err := time.LoadLocation(tZoneName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.US.Central())\n" +
			"tZoneName='%v'\n" +
			"Error='%v'\n", tZoneName, err.Error())
		return
	}

	t1, err := time.Parse(fmtstr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str)\n" +
			"t1str='%v'\n" +
			"Error='%v'\n", t1str, err.Error())
		return
	}

	t1Dup := time.Date(
		2019,
		time.Month(6),
		30,
		22,
		58,
		32,
		0,
		easternLocPtr)

	fmt.Printf("   t1 Time: %v\n", t1.Format(fmtstr))
	fmt.Printf("t1Dup Time: %v\n", t1Dup.Format(fmtstr))

	if t1.Equal(t1Dup) {
		fmt.Println("--- t1 is EQUAL to t1Dup ---")
	} else {
		fmt.Println("--- t1 is NOT equal to t1Dup ---")
	}

	fmt.Println()

	t1TzDefDtoDateTime, err :=
		dt.TimeZoneDefDto{}.NewFromTimeZoneName(tZoneName)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefDto{}.NewFromTimeZoneName(tZoneName)\n" +
			"tZoneName='%v'\n" +
			"Error='%v'\n", tZoneName, err.Error())
	}

	mt.mainPrintHdr("       t1TzDefDtoDateTime Data" , "=")
	fmt.Println("          From t1")
	ex.PrintOutTimeZoneDefDtoFields(t1TzDefDtoDateTime)
	fmt.Println()

	t2TzDefDtoDateTime, err :=
		dt.TimeZoneDefDto{}.New(t1)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefDto{}.New(t1)\n" +
			"t1='%v'\n" +
			"Error='%v'\n", t1.Format(fmtstr), err.Error())
		return
	}

	mt.mainPrintHdr("       t2TzDefDtoDateTime Data" , "=")
	fmt.Println("          From t1")
	ex.PrintOutTimeZoneDefDtoFields(t2TzDefDtoDateTime)
	fmt.Println()


	t1TzDefDtoDateTime, err =
		dt.TimeZoneDefDto{}.NewFromTimeZoneName(tZoneName)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefDto{}.NewFromTimeZoneName(tZoneName)\n" +
			"tZoneName='%v'\n" +
			"Error='%v'\n", tZoneName, err.Error())
	}

	t2TzDefDtoDateTime, err =
		dt.TimeZoneDefDto{}.New(t1Dup)

	mt.mainPrintHdr("       t1TzDefDtoDateTime Data" , "*")
	fmt.Println("From: t1Dup")
	ex.PrintOutTimeZoneDefDtoFields(t1TzDefDtoDateTime)
	fmt.Println()

	mt.mainPrintHdr("       t2TzDefDtoDateTime Data" , "*")
	fmt.Println("From: t1Dup")
	ex.PrintOutTimeZoneDefDtoFields(t2TzDefDtoDateTime)
	fmt.Println()

}

func (mt mainTest) mainTest034() {

	ePrefix := "mainTest034()"
	title := fmt.Sprintf("       %v         ", ePrefix)
	ln := strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	// t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t2, _ := time.Parse(fmtstr, t2str)

	tzu1, err := dt.TimeZoneDto{}.New(t2, dt.TZones.US.Eastern(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.New(t1, TzUsEast).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tzu1.IsValid()

	if err != nil {
		fmt.Printf("tzu1 is INVALID!\n" +
			"%v\n", err.Error())
		return
	}

}

func (mt mainTest) mainTest033() {

	ePrefix := "mainTest033()"
	title := fmt.Sprintf("       %v         ", ePrefix)
	ln := strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	locUSCentral, err := time.LoadLocation(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(dt.TZones.Asia.Tokyo())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
		return
	}

	t4USCentral := time.Date(
		2018,
		time.Month(3),
		06,
		20,
		02,
		18,
		792489279, locUSCentral)

	t4AsiaTokyo := t4USCentral.In(locTokyo)

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dTz, err := dt.DateTzDto{}.NewTz(
		t4AsiaTokyo,
		dt.TZones.US.Central(),
		dt.TzConvertType.Relative(),
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewTz(t4AsiaTokyo, TZones.US.Central(), fmtstr).\n" +
			"Error='%v'\n",
			err.Error())
		return
	}

	if !t4USCentral.Equal(dTz.GetDateTimeValue()) {
		fmt.Printf("Error: dTz Toyko to USA-Central\n" +
			"Expected DateTime='%v'.\n" +
			" Instead DateTime='%v'\n",
			t4USCentral.Format(fmtstr), dTz.GetDateTimeValue().Format(fmtstr))
		return
	}

	eTimeZoneDef, err := dt.TimeZoneDefDto{}.New(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefDto{}.New(t4USCentral)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !eTimeZoneDef.Equal(dTz.GetTimeZone()) {
		fmt.Printf("Expected dTz.GetTimeZone().LocationName='%v'.\n"+
			"Instead, dTz.GetTimeZone().LocationName='%v'\n",
			eTimeZoneDef.GetLocationName(), dTz.GetTimeZoneName())
		return
	}

	tDto, err := dt.TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral)\n"+
			"t4USCentral='%v'\nError='%v'\n",
			t4USCentral.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	dTzTimeComponents := dTz.GetTimeComponents()

	fmt.Println("tDto Comparison to dTz ")
	fmt.Printf("               Years: %v  -  %v\n", tDto.Years, dTzTimeComponents.Years )
	fmt.Printf("              Months: %v  -  %v\n", tDto.Months, dTzTimeComponents.Months )
	fmt.Printf("                Days: %v  -  %v\n", tDto.DateDays, dTzTimeComponents.DateDays )
	fmt.Printf("               Hours: %v  -  %v\n", tDto.Hours, dTzTimeComponents.Hours )
	fmt.Printf("             Minutes: %v  -  %v\n", tDto.Minutes, dTzTimeComponents.Minutes )
	fmt.Printf("             Seconds: %v  -  %v\n", tDto.Seconds, dTzTimeComponents.Seconds )
	fmt.Printf("        Milliseconds: %v  -  %v\n", tDto.Milliseconds, dTzTimeComponents.Milliseconds )
	fmt.Printf("        Microseconds: %v  -  %v\n", tDto.Microseconds, dTzTimeComponents.Microseconds )
	fmt.Printf("         Nanoseconds: %v  -  %v\n", tDto.Nanoseconds, dTzTimeComponents.Nanoseconds )
	fmt.Printf("TotSubSecNanoseconds: %v  -  %v\n", tDto.TotSubSecNanoseconds, dTzTimeComponents.TotSubSecNanoseconds )
	fmt.Println()
	fmt.Printf("                dTz Date Time: %v \n", dTz.GetDateTimeValue().Format(fmtstr))

	dTzTimeDtoDateTime, err := dTzTimeComponents.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by dTzTimeComponents.GetDateTime(dt.TZones.US.Central()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("dTz Time Components Date Time: %v\n", dTzTimeDtoDateTime.Format(fmtstr))

	tDtoDateTime, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by tDto.GetDateTime(dt.TZones.US.Central()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("             tDto USA Central: %v \n", tDtoDateTime.Format(fmtstr))
	fmt.Println()

	expectedDt, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from tDto.GetDateTime(TZones.US.Central()). "+
			"Error='%v'", err.Error())
		return
	}

	actualDt, err := dTzTimeComponents.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from dTz.GetTimeComponents().GetDateTime(TZones.US.Central()).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !tDto.Equal(dTz.GetTimeComponents()) {
		fmt.Printf("Expected dTz.Time (TimeDto) == '%v'\n" +
			"Instead, dTz.Time (TimeDto) == '%v'\n",
			expectedDt.Format(dt.FmtDateTimeYrMDayFmtStr), actualDt.Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	if dt.FmtDateTimeYrMDayFmtStr != dTz.GetDateTimeFmt() {
		fmt.Printf("Expected dTz.GetDateTimeFmt()='%v' Instead, dTz.GetDateTimeFmt()='%v' ",
			dt.FmtDateTimeYrMDayFmtStr, dTz.GetDateTimeFmt())
		return
	}

	title = fmt.Sprintf("       %v         ", "!!! Success !!!")
	ln = strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

}


func (mt mainTest) mainTest032() {

	ePrefix := "mainTest032()"
	title := fmt.Sprintf("       %v         ", ePrefix)
	ln := strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	t1Str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	expectedStr := "2014-02-15 19:54:30.038175584 -0500 EST"

	dt1, err := time.Parse(fmtStr, t1Str)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, t1Str).\n" +
			"t1Str='%v'\n", t1Str)
		return
	}

	dtUtil := dt.DTimeUtility{}

	dt2, err := dtUtil.AbsoluteTimeToTimeZoneNameConversion(dt1, dt.TZones.America.New_York())

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by dtUtil.AbsoluteTimeToTimeZoneDtoConversion(dt1, dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dt2Str := dt2.Format(fmtStr)

	var result string

	if dt2Str != expectedStr {
		result = "*** FAILURE ***"
	} else {
		result = "!!! SUCCESS !!!"
	}

	title = fmt.Sprintf("       %v         ", result)
	ln = strings.Repeat("=", len(title))

	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	fmt.Printf("   Start Date Time: %v\n", t1Str)
	fmt.Printf("  Actual Date Time: %v\n", dt2Str)
	fmt.Printf("Expected Date Time: %v\n", expectedStr)
	fmt.Println()
}

func (mt mainTest) mainTest031() {

	ePrefix := "mainTest031()"
	title := fmt.Sprintf("       %v         ", ePrefix)
	ln := strings.Repeat("-", len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	t1Str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	expectedStr := "2014-02-15 19:54:30.038175584 -0500 EST"

	dt1, err := time.Parse(fmtStr, t1Str)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, t1Str).\n" +
			"t1Str='%v'\n", t1Str)
		return
	}

	tzDefDto, err := dt.TimeZoneDefDto{}.NewFromTimeZoneName(dt.TZones.America.New_York())

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by dt.TimeZoneDefDto{}.NewFromTimeZoneName(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dtUtil := dt.DTimeUtility{}

	dt2, err := dtUtil.AbsoluteTimeToTimeZoneDtoConversion(dt1, tzDefDto)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by dtUtil.AbsoluteTimeToTimeZoneDtoConversion(dt1, tzDefDto)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dt2Str := dt2.Format(fmtStr)

	var result string

	if dt2Str != expectedStr {
		result = "*** FAILURE ***"
	} else {
		result = "!!! SUCCESS !!!"
	}

	title = fmt.Sprintf("       %v         ", result)
	ln = strings.Repeat("=", len(title))

	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

	fmt.Printf("   Start Date Time: %v\n", t1Str)
	fmt.Printf("  Actual Date Time: %v\n", dt2Str)
	fmt.Printf("Expected Date Time: %v\n", expectedStr)
	fmt.Println()
}

func (mt mainTest) mainTest030() {

	fmt.Println("       mainTest030()         ")
	fmt.Println("-----------------------------")
	fmt.Println()

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtstr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str).\n" +
			"Error='%v'", err.Error())
		return
	}

	// tzLocName := "America/Chicago"
	tzLocName := "Local"

	tzLoc, err := time.LoadLocation(tzLocName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzLocName)\n" +
			"tzLocName='%v'\n" +
			"Error='%v'\n", tzLocName, err.Error())

		return
	}

	t2 := t1.In(tzLoc)

	dtz1, err := dt.DateTzDto{}.New(t2, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.New(t2, fmtstr).\n" +
			"t2='%v'\n" +
			"Error='%v'\n", t2.Format(fmtstr), err.Error())
		return
	}

	tZoneDef := dtz1.GetTimeZone()

	fmt.Println("------ Success!!! ------")
	fmt.Println()
	fmt.Printf("   Time Zone Name: %v\n", tzLocName)
	fmt.Println("  -- tZoneDef Values -- ")
	fmt.Println()
	fmt.Printf("         ZoneName: %v\n", tZoneDef.GetZoneName())
	fmt.Printf("ZoneOffsetSeconds: %v\n", tZoneDef.GetZoneOffsetSeconds())
	fmt.Printf("         ZoneSign: %v\n", tZoneDef.GetZoneSign())
	fmt.Printf("      OffsetHours: %v\n", tZoneDef.GetOffsetHours())
	fmt.Printf("    OffsetMinutes: %v\n", tZoneDef.GetOffsetMinutes())
	fmt.Printf("    OffsetSeconds: %v\n", tZoneDef.GetOffsetSeconds())
	fmt.Printf("       ZoneOffset: %v\n", tZoneDef.GetZoneOffset())
	fmt.Printf("       UTC Offset: %v\n", tZoneDef.GetUtcOffset())
	fmt.Printf("    Location Name: %v\n", tZoneDef.GetLocationName())
	fmt.Printf("        *Location: %v\n", tZoneDef.GetLocationPtr().String())
	fmt.Printf("      Description: %v\n", tZoneDef.GetTagDescription())
	fmt.Println()
	fmt.Println()
}
func (mt mainTest) mainTest029() {

	fmt.Println("       mainTest029()         ")
	fmt.Println("-----------------------------")
	fmt.Println()

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtstr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str).\n" +
			"Error='%v'", err.Error())
		return
	}

	// tzLocName := "America/Chicago"
	tzLocName := "Asia/Kabul"

	tzLoc, err := time.LoadLocation(tzLocName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzLocName)\n" +
			"tzLocName='%v'\n" +
			"Error='%v'\n", tzLocName, err.Error())

		return
	}

	t2 := t1.In(tzLoc)

	dtz1, err := dt.DateTzDto{}.New(t2, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.New(t2, fmtstr).\n" +
			"t2='%v'\n" +
			"Error='%v'\n", t2.Format(fmtstr), err.Error())
		return
	}

	tZoneDef := dtz1.GetTimeZone()

	fmt.Println("------ Success!!! ------")
	fmt.Println()
	fmt.Printf("   Time Zone Name: %v\n", tzLocName)
	fmt.Println("  -- tZoneDef Values -- ")
	fmt.Println()
	fmt.Printf("         ZoneName: %v\n", tZoneDef.GetZoneName())
	fmt.Printf("ZoneOffsetSeconds: %v\n", tZoneDef.GetZoneOffsetSeconds())
	fmt.Printf("         ZoneSign: %v\n", tZoneDef.GetZoneSign())
	fmt.Printf("      OffsetHours: %v\n", tZoneDef.GetOffsetHours())
	fmt.Printf("    OffsetMinutes: %v\n", tZoneDef.GetOffsetMinutes())
	fmt.Printf("    OffsetSeconds: %v\n", tZoneDef.GetOffsetSeconds())
	fmt.Printf("       ZoneOffset: %v\n", tZoneDef.GetZoneOffset())
	fmt.Printf("       UTC Offset: %v\n", tZoneDef.GetUtcOffset())
	fmt.Printf("    Location Name: %v\n", tZoneDef.GetLocationName())
	fmt.Printf("        *Location: %v\n", tZoneDef.GetLocationPtr().String())
	fmt.Printf("      Description: %v\n", tZoneDef.GetTagDescription())
	fmt.Println()
	fmt.Println()
}

func (mt mainTest) mainTest028() {

	fmt.Println("       mainTest028()         ")
	fmt.Println("-----------------------------")
	fmt.Println()

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtstr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtstr, t1str).\n" +
			"Error='%v'", err.Error())
		return
	}

	tzLocName := "Asia/Ho_Chi_Minh"

	tzLoc, err := time.LoadLocation(tzLocName)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(tzLocName)\n" +
			"tzLocName='%v'\n" +
			"Error='%v'\n", tzLocName, err.Error())

		return
	}

	t2 := t1.In(tzLoc)

	dtz1, err := dt.DateTzDto{}.New(t2, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.New(t2, fmtstr).\n" +
			"t2='%v'\n" +
			"Error='%v'\n", t2.Format(fmtstr), err.Error())
		return
	}

	tZoneDef := dtz1.GetTimeZone()

	fmt.Println("------ Success!!! ------")
	fmt.Println()
	fmt.Printf("   Time Zone Name: %v\n", tzLocName)
	fmt.Println("  -- tZoneDef Values -- ")
	fmt.Println()
	fmt.Printf("         ZoneName: %v\n", tZoneDef.GetZoneName())
	fmt.Printf("ZoneOffsetSeconds: %v\n", tZoneDef.GetZoneOffsetSeconds())
	fmt.Printf("         ZoneSign: %v\n", tZoneDef.GetZoneSign())
	fmt.Printf("      OffsetHours: %v\n", tZoneDef.GetOffsetHours())
	fmt.Printf("    OffsetMinutes: %v\n", tZoneDef.GetOffsetMinutes())
	fmt.Printf("    OffsetSeconds: %v\n", tZoneDef.GetOffsetSeconds())
	fmt.Printf("       ZoneOffset: %v\n", tZoneDef.GetZoneOffset())
	fmt.Printf("       UTC Offset: %v\n", tZoneDef.GetUtcOffset())
	fmt.Printf("    Location Name: %v\n", tZoneDef.GetLocationName())
	fmt.Printf("        *Location: %v\n", tZoneDef.GetLocationPtr().String())
	fmt.Printf("      Description: %v\n", tZoneDef.GetTagDescription())
}

func (mt mainTest) mainTest027() {

	fmt.Println("       mainTest027()         ")
	fmt.Println("-----------------------------")
	fmt.Println()

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	expectedOutDate := t1.Format(fmtstr)

	dtz1, err := dt.DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.New(t1, fmtstr).\n" +
			"Error='%v'", err.Error())
		return
	}

	if expectedOutDate != dtz1.String() {
		fmt.Printf("Error: Expected dtz1.String()='%v'.\n" +
			"Instead, dtz1.String()='%v'\n", expectedOutDate, dtz1.String())
	}

	t2 := t1.AddDate(5, 6, 12)

	dtz2, err := dtz1.AddDate(5, 6, 12, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dtz1.AddDate(5, 6, 12, fmtstr).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedOutDate = t2.Format(fmtstr)

	if expectedOutDate != dtz2.String() {
		fmt.Printf("Error: Expected dtz2.String()='%v'.\n" +
			"Instead, dtz2.String()='%v'\n", expectedOutDate, dtz2.String())
		return
	}

	fmt.Println("------ Success!!! ------")

}

/*
	func (mt mainTest) mainTest026() {

	fmt.Println("       mainTest026()         ")
	fmt.Println("-----------------------------")
	fmt.Println()
	
	tstr := "12/06/2019 03:12:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	testTime, err := time.Parse(fmtStr, tstr)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtStr, tstr)\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
		return
	}

	var milTzDto dt.MilitaryDateTzDto
	var dateTzDto dt.DateTzDto

	milTzDto, err = dt.MilitaryDateTzDto{}.New(testTime, "Sierra")

	if err != nil {
		fmt.Printf("Error returned by MilitaryDateTzDto{}.New(testTime, \"Sierra\")\n" +
			"testTime='%v'\n" +
			"Error='%v'\n", testTime.Format(fmtStr), err.Error())
		return
	}

	dateTzDto, err = dt.DateTzDto{}.NewFromMilitaryDateTz(milTzDto, fmtStr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewFromMilitaryDateTz(milTzDto, fmtStr)\n" +
			"milTzDto.DateTime='%v'\n" +
			"Error='%v'\n", milTzDto.DateTime.Format(fmtStr), err.Error())
		return
	}

	fmt.Println("         Success!!           ")
	fmt.Println("-----------------------------")
	fmt.Println()
	fmt.Printf("Military DateTime: %v", dateTzDto.GetDateTimeValue().Format(fmtStr))

}
*/

func (mt mainTest) mainTest025() {

	// expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	fmt.Println("       mainTest025()     ")

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Eastern(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.New(t1, TzUsEast).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDto := dt.TimeDto{Years: 3, Months: 2, DateDays: 15, Hours: 3, Minutes: 4, Seconds: 2}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddPlusTimeDto(tDto)

	if err != nil {
		fmt.Printf("Error returned by tzu2.AddPlusTimeDto(tDto). " +
			"Error='%v' ", err.Error())
		return
	}

	tzu2TimeInStr := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t2OutStr != tzu2TimeInStr {
		fmt.Printf("Error: Expected tzu2.TimeIn='%v'.  Instead, tzu2.TimeIn='%v'. ", t2OutStr, tzu2TimeInStr)
		return
	}

	fmt.Println("       Successful Completion!     ")
}

func (mt mainTest) mainTest024() {
	ePrefix := "mainTest.mainTest024() "
	fmt.Println(ePrefix)
	tstr := "12/02/2019 22:05:00 -0600 CST"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	var testTime, expectedMilTime time.Time
	var err error
	var actualMilDateTimeGroup, expectedMilDateTimeGroup string
	var milDatTzDto dt.MilitaryDateTzDto
	var expectedMilTimeLoc *time.Location

	testTime, err = time.Parse(fmtstr, tstr)

	expectedMilTimeLoc, err = time.LoadLocation(dt.TZones.Military.Quebec())

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(dt.TZones.Military.Quebec())\n" +
			"dt.TZones.Military.Quebec()='%v'\n" +
			"Error='%v'\n", dt.TZones.Military.Quebec(), err.Error())
	}

	expectedMilTime = testTime.In(expectedMilTimeLoc)

	milDatTzDto, err = dt.MilitaryDateTzDto{}.New(testTime, "Q")

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by MilitaryDateTzDto{}.New(testTime, \"Q\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedMilDateTimeGroup, err = dt.DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by DtMgr{}.GetMilitaryOpenDateTimeGroup(expectedMilTime)\n" +
			"expectedMilTime='%v'\n" +
			"Error='%v'\n",
			expectedMilTime.Format(fmtstr) ,err.Error())
		return
	}

	actualMilDateTimeGroup, err = milDatTzDto.GetOpenDateTimeGroup()

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by milDatTzDto.GetOpenDateTimeGroup()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if expectedMilDateTimeGroup != actualMilDateTimeGroup {
		fmt.Printf(ePrefix +
			"\nError: Expected Military Date Time Group='%v'.\n" +
			"Actual Military Date Time Group='%v'\n" +
			"Military Time='%v'",
			expectedMilDateTimeGroup, actualMilDateTimeGroup, expectedMilTime.Format(fmtstr))
	}

	fmt.Println("***** Success *****")
	fmt.Printf("Expected Military Date Time Group: %v\n", expectedMilDateTimeGroup)
	fmt.Printf("  Actual Military Date Time Group: %v\n", actualMilDateTimeGroup)
	fmt.Printf("              Original Start Time: %v\n", testTime.Format(fmtstr))
	fmt.Printf("      Military Time in Quebect Tz: %v\n", milDatTzDto.DateTime.Format(fmtstr))
}

func (mt mainTest) mainTest023() {
	ePrefix := "mainTest.mainTest023() "

	fmt.Println(ePrefix)
	tstr := "11/29/2017 19:54:30 -0600 CST"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"

	expected := "291954SNOV17"
	var testTime time.Time
	var err error
	var result string
	var milDatTzDto dt.MilitaryDateTzDto

	testTime, err = time.Parse(fmtstr, tstr)

	milDatTzDto, err = dt.MilitaryDateTzDto{}.New(testTime, "S")

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by MilitaryDateTzDto{}.New(testTime, \"S\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	result, err = milDatTzDto.GetCompactDateTimeGroup()

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by milDatTzDto.GetCompactDateTimeGroup()\n")
		return
	}

	if result != expected {
		fmt.Printf(ePrefix +
			"\nError: Expected result='%v'.\n" +
			"Instead, result='%v'.\n", expected, result)
		return
	}

	fmt.Println("  Result= ", result)
	fmt.Println("Expected= ", expected)

}

func (mt mainTest) mainTest022() {

	ePrefix := "mainTest.mainTest022() "

	fmt.Println(ePrefix)

	t := time.Now().Local()

	militaryDateTime, err := dt.DtMgr{}.GetMilitaryCompactDateTimeGroup(t)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError retunred by dt.DtMgr{}.GetMilitaryCompactDateTimeGroup(t).\n" +
			"t='%v'\nError='%v'\n", t.Format(dt.FmtDateTimeYMDHMSTz),
			err.Error())
		return
	}

	fmt.Println("Standard Format: ", t.Format(dt.FmtDateTimeYMDHMSTz))
	fmt.Println("Military Format: ", militaryDateTime)

}

func (mt mainTest) mainTest021() {
	tz := dt.TZones.Other.Factory()
	dtz, err := dt.DateTzDto{}.NewNowTz(tz, dt.FmtDateTimeYrMDayFmtStr )

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewNowTz(tz, "	+
			"dt.FmtDateTimeYrMDayFmtStr ).\n Error='%v' \n", err.Error())
		return
	}

	tzDef2, err := dt.TimeZoneDefDto{}.New(dtz.GetDateTimeValue())
	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	fmt.Println()
	fmt.Println("Testing tz", tz)
	fmt.Println("-------------------------------------")
	fmt.Println("    Zone Name: ", tzDef2.GetZoneName())
	fmt.Println("  Zone Offset: ", tzDef2.GetZoneOffset())
	fmt.Println("   UTC Offset: ", tzDef2.GetUtcOffset())
	fmt.Println("Location Name: ", tzDef2.GetLocationName())
	fmt.Println("    *Location: ", tzDef2.GetLocationPtr().String())
}


func (mt mainTest) mainTest020() {

	tz1 := "Cuba"
	tz2 := "America/Havana"

	dtz, err := dt.DateTzDto{}.NewNowTz(tz1, dt.FmtDateTimeYrMDayFmtStr )

	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	tzDef, err := dt.TimeZoneDefDto{}.New(dtz.GetDateTimeValue())

	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	dtz2, err := dt.DateTzDto{}.NewNowTz(tz2, dt.FmtDateTimeYrMDayFmtStr )

	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	tzDef2, err := dt.TimeZoneDefDto{}.New(dtz2.GetDateTimeValue())
	if err != nil {
		fmt.Printf("%v \n", err.Error())
		return
	}

	fmt.Println("Testing tz1", tz1)
	fmt.Println("-------------------------------------")
	fmt.Println("    Zone Name: ", tzDef.GetZoneName())
	fmt.Println("       Offset: ", tzDef.GetZoneOffset())
	fmt.Println("Location Name: ", tzDef.GetLocationName())
	fmt.Println("    *Location: ", tzDef.GetLocationPtr().String())

	fmt.Println()
	fmt.Println("Testing tz2", tz2)
	fmt.Println("-------------------------------------")
	fmt.Println("    Zone Name: ", tzDef2.GetZoneName())
	fmt.Println("       Offset: ", tzDef2.GetZoneOffset())
	fmt.Println("Location Name: ", tzDef2.GetLocationName())
	fmt.Println("    *Location: ", tzDef2.GetLocationPtr().String())
}

func (mt mainTest) mainTest019() {
	locUSCentral, err := time.LoadLocation(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	year := 2018
	month := 3
	day := 6
	hour := 20
	minute := 2
	second := 18
	nSecs := 792489279

	t1USCentral := time.Date(year, time.Month(month), day, hour, minute, second, nSecs, locUSCentral)

	minute = 3
	second = 20
	t2USCentral := time.Date(year, time.Month(month), day, hour, minute, second, nSecs, locUSCentral)

	tDur, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(
		t1USCentral,
		t2USCentral,
		dt.TDurCalcType(0).StdYearMth(),
		dt.TZones.US.Central(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesCalcTz()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	str, _ := tDur.GetCumSecondsTimeStr()

	fmt.Println("Cumulative Seconds")
	fmt.Println(str)
	// 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds
}

func (mt mainTest) mainTest018() {
	t1Dto := dt.TimeDto{}

	/*
		t1Dto.Years = 1955
		t1Dto.Months = 15
		t1Dto.DateDays = 32
		t1Dto.Hours = 48
		t1Dto.Minutes = 71
		t1Dto.Seconds = 125
		t1Dto.Milliseconds = 1001
		t1Dto.Microseconds = 1001
		t1Dto.Nanoseconds = 1001
	*/

	t1Dto.Years = 0
	t1Dto.Months = 0
	t1Dto.DateDays = 32
	t1Dto.Hours = 48
	t1Dto.Minutes = 71
	t1Dto.Seconds = 125
	t1Dto.Milliseconds = 0
	t1Dto.Microseconds = 0
	t1Dto.Nanoseconds = 0

	err := t1Dto.NormalizeTimeElements()
	if err != nil {
		fmt.Printf("Error returned by t1Dto.NormalizeTimeElements(). "+
			"Error='%v' ", err.Error())
		return
	}

	fmt.Println("After Normalize Time Elements ")
	ex.PrintOutTimeDtoFields(t1Dto)

	_, err = t1Dto.NormalizeDays()

	if err != nil {
		fmt.Printf("Error returned by t1Dto.NormalizeDays(). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("After ")
	ex.PrintOutTimeDtoFields(t1Dto)
	dateTime, err := t1Dto.GetDateTime(dt.TZones.Other.UTC())

	if err != nil {
		fmt.Printf("Error returned by t1Dto.GetDateTime(dt.TZones.UTC()). Error='%v'\n",
			err.Error())
		return
	}

	fmt.Println("t1Dto.GetDateTime(): ", dateTime.Format(dt.FmtDateTimeYrMDayFmtStr))

}

func (mt mainTest) mainTest017() {
	t1Dto := dt.TimeDto{}
	/*
		t1Dto.Years = 1955
		t1Dto.Months = 15
		t1Dto.DateDays = 32
		t1Dto.Hours = 48
		t1Dto.Minutes = 71
		t1Dto.Seconds = 125
		t1Dto.Milliseconds = 1001
		t1Dto.Microseconds = 1001
		t1Dto.Nanoseconds = 1001
	*/

	t1Dto.Years = 1955
	t1Dto.Months = 15
	t1Dto.DateDays = 32
	t1Dto.Hours = 48
	t1Dto.Minutes = 71
	t1Dto.Seconds = 125
	t1Dto.Milliseconds = 0
	t1Dto.Microseconds = 0
	t1Dto.Nanoseconds = 123456789

	fmt.Println("Original TimeDto Values")
	ex.PrintOutTimeDtoFields(t1Dto)

	err := t1Dto.NormalizeTimeElements()

	if err != nil {
		fmt.Printf("Error returned by t1Dto.NormalizeTimeElements(). Error='%v' \n",
			err.Error())
		return
	}

	fmt.Println("After Normalize Time Elements")
	ex.PrintOutTimeDtoFields(t1Dto)

	_, err = t1Dto.NormalizeDays()

	if err != nil {
		fmt.Printf("Error returned by t1Dto.NormalizeDays(). Error='%v' \n",
			err.Error())
		return
	}

	fmt.Println("After Normalize Days")
	ex.PrintOutTimeDtoFields(t1Dto)

	utcLoc, _ := time.LoadLocation(dt.TZones.Other.UTC())

	tDate := time.Date(1956, 3, 34, 1, 13, 6, 2002001, utcLoc)

	fmt.Println("tDate: ", tDate.Format(dt.FmtDateTimeYrMDayFmtStr))

	fmt.Println("SUCCESSFUL COMPLETION!")
	/*

	   Original TimeDto Values
	   ========================================
	             TimeDto Printout
	   ========================================
	                      Years:  1955
	                     Months:  15
	                      Weeks:  0
	                   WeekDays:  0
	                   DateDays:  32
	                      Hours:  48
	                    Minutes:  71
	                    Seconds:  125
	               Milliseconds:  1001
	               Microseconds:  1001
	                Nanoseconds:  1001
	   Total SubSec Nanoseconds:  0
	     Total Time Nanoseconds:  0
	   ========================================
	   After Normalize Time Elements
	   ========================================
	             TimeDto Printout
	   ========================================
	                      Years:  1956
	                     Months:  3
	                      Weeks:  4
	                   WeekDays:  6
	                   DateDays:  34
	                      Hours:  1
	                    Minutes:  13
	                    Seconds:  6
	               Milliseconds:  2
	               Microseconds:  2
	                Nanoseconds:  1
	   Total SubSec Nanoseconds:  2002001
	     Total Time Nanoseconds:  4386002002001
	   ========================================
	   After Normalize Days
	   ========================================
	             TimeDto Printout
	   ========================================
	                      Years:  1956
	                     Months:  4
	                      Weeks:  0
	                   WeekDays:  3
	                   DateDays:  3
	                      Hours:  1
	                    Minutes:  13
	                    Seconds:  6
	               Milliseconds:  2
	               Microseconds:  2
	                Nanoseconds:  1
	   Total SubSec Nanoseconds:  2002001
	     Total Time Nanoseconds:  4386002002001
	   ========================================
	   tDate:  1956-04-03 01:13:06.002002001 +0000 UCT

	*/

}

func (mt mainTest) mainTest016() {
	locUSCentral, err := time.LoadLocation(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	year := 2018
	month := 3
	day := 6
	hour := 20
	minute := 2
	second := 18
	nSecs := 792489279

	t4USCentral := time.Date(year, time.Month(month), day, hour, minute, second, nSecs, locUSCentral)

	t4Dto, err := dt.TimeDto{}.New(year, month, 0, day, hour, minute,
		second, 0, 0, nSecs)

	if err != nil {
		fmt.Printf("Error returned by t4USCentral TimeDto{}.New(). Error='%v'\n", err.Error())
		return
	}

	t4TZoneDef, err := dt.TimeZoneDefDto{}.New(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefDto{}.New(t4USCentral). Error='%v'", err.Error())
		return
	}

	locTokyo, err := time.LoadLocation(dt.TZones.Asia.Tokyo())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
		return
	}

	t5Tokyo := time.Date(2012, 9, 30, 11, 58, 48, 123456789, locTokyo)

	t5Dto, err := dt.TimeDto{}.New(2012, 9, 0, 30, 11,
		58, 48, 0, 0, 123456789)

	if err != nil {
		fmt.Printf("Error returned by t5Tokyo TimeDto{}.New(). Error='%v'", err.Error())
		return
	}

	t5TZoneDef, err := dt.TimeZoneDefDto{}.New(t5Tokyo)

	dTz1, err := dt.DateTzDto{}.New(t5Tokyo, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.New(t4USCentral, FmtDateTimeYrMDayFmtStr)\n")
		return
	}

	timeComponents := dTz1.GetTimeComponents()
	
	if !t5Dto.Equal(timeComponents) {
		fmt.Print("Expected t5Dto == dTz1.timeComponents. It DID NOT!\n")

		fmt.Println("t5Dto")
		ex.PrintOutTimeDtoFields(t5Dto)
		fmt.Println("\n\ndTz1.Time")
		ex.PrintOutTimeDtoFields(timeComponents)
		return
	}

	if !t5TZoneDef.Equal(dTz1.GetTimeZone()) {
		fmt.Print("Expected t5TZoneDef == dTz1.GetTimeZone(). It DID NOT!")
	}

	err = dTz1.SetFromTimeDto(t4Dto, dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from dTz1.SetFromTimeDto(t4Dto, TZones.US.Central()). "+
			"Error='%v'\n", err.Error())
		return
	}

	if !t4USCentral.Equal(dTz1.GetDateTimeValue()) {
		fmt.Printf("Expected dTz1.DateTime='%v'.\n" +
			"Instead, dTz1.DateTime='%v'.\n",
			t4USCentral.Format(dt.FmtDateTimeYrMDayFmtStr),
			dTz1.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}
	
	if !t4Dto.Equal(timeComponents) {
		fmt.Print("Expected t4Dto TimeDto == dTz1.Time Time Dto. THEY ARE NOT EQUAL!\n")
		fmt.Println("t4Dto")

		ex.PrintOutTimeDtoFields(t5Dto)
		fmt.Println("\n\ndTz1.Time")
		ex.PrintOutTimeDtoFields(timeComponents)
		return
	}

	if !t4TZoneDef.Equal(dTz1.GetTimeZone()) {
		fmt.Print("Expected t4TZoneDef TimeZoneDef == dTz1.GetTimeZone() TimeZoneDef. " +
			"THEY ARE NOT EQUAL!\n")

		fmt.Println("t4TZoneDef")
		ex.PrintOutTimeZoneDefDtoFields(t4TZoneDef)
		fmt.Println("\n\ndTz1.GetTimeZone()")
		ex.PrintOutTimeZoneDefDtoFields(dTz1.GetTimeZone())

		return
	}

	if year != dTz1.GetTimeComponents().Years {
		fmt.Printf("Error: Expected Years='%v'.\n" +
			"Instead, Years='%v'\n", year, dTz1.GetTimeComponents().Years)
		return
	}

	if month != dTz1.GetTimeComponents().Months {
		fmt.Printf("Error: Expected Months='%v'.\n" +
			"Instead, Months='%v'\n", month, dTz1.GetTimeComponents().Months)
		return
	}

	if day != dTz1.GetTimeComponents().DateDays {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			day, dTz1.GetTimeComponents().DateDays)
		return
	}

	if hour != dTz1.GetTimeComponents().Hours {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n", 
			hour, dTz1.GetTimeComponents().Hours)
		return
	}

	if minute != dTz1.GetTimeComponents().Minutes {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n", 
			minute, dTz1.GetTimeComponents().Minutes)
		return
	}

	if second != dTz1.GetTimeComponents().Seconds {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n", 
			second, dTz1.GetTimeComponents().Seconds)
		return
	}

	if 792 != dTz1.GetTimeComponents().Milliseconds {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			792, dTz1.GetTimeComponents().Milliseconds)
		return
	}

	if 489 != dTz1.GetTimeComponents().Microseconds {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			489, dTz1.GetTimeComponents().Microseconds)
		return
	}

	if 279 != dTz1.GetTimeComponents().Nanoseconds {
		fmt.Printf("Error: Expected Days='%v'. Instead, Days='%v'\n",
			279, dTz1.GetTimeComponents().Nanoseconds)
		return
	}

	if nSecs != dTz1.GetTimeComponents().TotSubSecNanoseconds {
		fmt.Printf("Error: Expected dTz1.GetTimeComponents().TotSubSecNanoseconds='%v'. "+
			"Instead, dTz1.GetTimeComponents().TotSubSecNanoseconds='%v'\n", nSecs, dTz1.GetTimeComponents().TotSubSecNanoseconds)
		return
	}

	totTime := int64(hour) * int64(time.Hour)
	totTime += int64(minute) * int64(time.Minute)
	totTime += int64(second) * int64(time.Second)
	totTime += int64(nSecs)

	if totTime != dTz1.GetTimeComponents().TotTimeNanoseconds {
		fmt.Printf("Error: Expected tDto.TotTimeNanoseconds='%v'. "+
			"Instead, tDto.TotTimeNanoseconds='%v'\n", totTime, dTz1.GetTimeComponents().TotTimeNanoseconds)
		return
	}

	fmt.Println("SUCCESSFUL COMPLETION!!!")
}

func (mt mainTest) mainTest015() {

	// t1str :="2017-04-30 22:58:32.515539300 -0500 CDT"
	// t1, err := time.Parse(FmtDateTimeYrMDayFmtStr, t1str)

	dTzDto, err := dt.DateTzDto{}.NewDateTimeElements(2017, 04, 30, 22, 58, 32, 515539300, dt.TZones.US.Central(), dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned from DateTzDto{}.NewDateTimeElements(year, month, day,...). "+
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Println("dTzDto.DateTime: ", dTzDto.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))

	tDto, err := dt.TimeDto{}.NewFromDateTime(dTzDto.GetDateTimeValue())

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDto{}.NewFromDateTime(dTzDto.DateTime)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dt2, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by tDto.GetDateTime(TZones.US.Central()). Error='%v'\n", err.Error())
		return
	}

	if !dt2.Equal(dTzDto.GetDateTimeValue()) {
		fmt.Printf("Error: Expected dTzDto.DateTime='%v'. It did NOT! dTzDto.DateTime='%v'\n",
			dt2.Format(dt.FmtDateTimeYrMDayFmtStr),
			dTzDto.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	fmt.Println("Success!")
}

func (mt mainTest) mainTest014() {
	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	t2str := "2017-04-30 22:58:32.000000000 -0500 CDT"
	t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)

	//t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t1OutStr := t1.Format(dt.FmtDateTimeYrMDayFmtStr)
	//t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)
	t2OutStr := t2.Format(dt.FmtDateTimeYrMDayFmtStr)
	t12Dur := t2.Sub(t1)

	timeDto := dt.TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err := timeDto.NormalizeTimeElements()

	if err != nil {
		fmt.Printf("Error returned by timeDto.NormalizeTimeElements(). "+
			"Error='%v' ", err.Error())
		return
	}

	dur, err := dt.DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto, dt.TZones.US.Central(), dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto). "+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error- Expected Start Time %v. Instead, got %v.\n",
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr) {
		fmt.Printf("Error- Expected End Time %v. Instead, got %v.\n",
			t2OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	if t12Dur != dur.BaseTime.TimeDuration {
		fmt.Printf("Error- Expected Time Duration %v. Instead, got %v\n",
			t12Dur, dur.BaseTime.TimeDuration)
		return
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthDay: %v. Instead, got %v\n", expected, outStr)
		return
	}

	fmt.Println("Successful Completion!")
}

func (mt mainTest) mainTest012() {
	// 101095442000000000
	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	t2str := "2017-04-30 22:58:32.000000000 -0500 CDT"
	t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)
	eDur := t2.Sub(t1)

	timeDto := dt.TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	timeDto.ConvertToNegativeValues()

	dx1 := t2.AddDate(timeDto.Years, timeDto.Months, 0)

	dur := (int64(timeDto.Weeks*7) + int64(timeDto.WeekDays)) * dt.DayNanoSeconds
	dur += int64(timeDto.Hours) * dt.HourNanoSeconds
	dur += int64(timeDto.Minutes) * dt.MinuteNanoSeconds
	dur += int64(timeDto.Seconds) * dt.SecondNanoseconds

	dx2 := dx1.Add(time.Duration(dur))

	fmt.Println("Expected Start Date Time: ", t1.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("  Actual Start Date Time: ", dx2.Format(dt.FmtDateTimeYrMDayFmtStr))
	fmt.Println("       Expected Duration: ", int64(eDur))
	fmt.Println("          ActualDuration:  101095442000000000")

}

func (mt mainTest) mainTest011() {

	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	t2str := "2017-04-30 22:58:32.000000000 -0500 CDT"
	t1, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t1str)
	t2, _ := time.Parse(dt.FmtDateTimeYrMDayFmtStr, t2str)

	tDto, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(t1, t2,
		dt.TDurCalcType(0).StdYearMth(), dt.TZones.US.Central(), dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesCalcTz() "+
			"Error='%v'", err.Error())
		return
	}

	ex.PrintTimeDurationDto(tDto)

}

func (mt mainTest) mainTest010() {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := dt.TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err := timeDto.NormalizeTimeElements()

	if err != nil {
		fmt.Printf("Error returned by timeDto.NormalizeTimeElements(). "+
			"Error='%v' ", err.Error())
		return
	}

	dur, err := dt.TimeDurationDto{}.NewEndTimeMinusTimeDto(t2, timeDto, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto). Error='%v'\n", err.Error())
		return
	}

	fmt.Println("Expected Start Date Time: ", t1OutStr)
	fmt.Println("  Actual Start Date Time: ", dur.StartTimeDateTz.String())
	fmt.Println("-----------------------------------------")
	fmt.Println("  Expected End Date Time: ", t2OutStr)
	fmt.Println("    Actual End Date Time: ", dur.EndTimeDateTz.String())
	fmt.Println("-----------------------------------------")
	fmt.Println("       Expected Duration: ", t12Dur)
	fmt.Println("         Actual Duration: ", dur.TimeDuration.String())
}

func (mt mainTest) mainTest009() {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto, err :=
		dt.TimeDurationDto{}.NewStartEndTimesCalcTz(t2, t1,
			dt.TDurCalcType(0).StdYearMth(), dt.TZones.US.Central(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesCalcTz(). "+
			" Error='%v'\n", err.Error())
		return
	}

	fmt.Println("TimeDurationDto")
	ex.PrintTimeDurationDto(tDto)

	durT, err := dt.DurationTriad{}.NewStartEndTimesTz(t2, t1, dt.TZones.US.Central(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDurationDto{}.NewStartEndTimesCalcTz(). "+
			" Error='%v'\n", err.Error())
		return
	}

	fmt.Println("DurationTriad BaseTimeDto")
	ex.PrintTimeDurationDto(durT.BaseTime)
}

func (mt mainTest) mainTest008() {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := dt.DurationTriad{}

	err := du.SetStartEndTimesTz(t2, t1, dt.TZones.US.Central(), dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by du.SetStartEndTimesTz(t2, t1, dt.TZones.US.Central(), "+
			"dt.FmtDateTimeYrMDayFmtStr). Error='%v' ", err.Error())
		return
	}

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearMthDaysTimeAbbrvStr()

	fmt.Println("Expected: ", expected)
	fmt.Println("  Actual: ", dOut)
	fmt.Println("Start Time: ", du.BaseTime.StartTimeDateTz.String())
	fmt.Println("  End Time: ", du.BaseTime.EndTimeDateTz.String())

}

func (mt mainTest) mainTest007() {

	mthTest := int(time.Month(0))

	fmt.Println("===============================")
	fmt.Println("       Month Zero Test")
	fmt.Println("===============================")
	fmt.Println("int (time.Month(0))= ", mthTest)

	/* Result = Sending zero to time.Month(0) yields a
	a zero value. NOT GOOD! Best to use '1' in
	place of zero month number.

		===============================
					 Month Zero Test
		===============================
		int (time.Month(0))=  0

	*/

}

func (mt mainTest) mainTest006() {

	locUTC, _ := time.LoadLocation(dt.TZones.Other.UTC())

	fmt.Println()
	fmt.Println("2018-00")
	fmt.Println()

	tDateTime := time.Date(2018, 0, 0, 0, 0, 0, 0, locUTC)

	ex.PrintDateTime(tDateTime, dt.FmtDateTimeYrMDayFmtStr)
	/*		Result - Don't Use a Zero Month
			2018-00
			----------------------------------
								 Date Time
			----------------------------------
			Date Time:  2017-11-30 00:00:00.000000000 +0000 UCT
			The integer month is:  11
			The integer day is: 30
			The integer year is: 2017
			The integer hour is: 0
			The integer minute is: 0
			The integer second is: 0
			The integer nanosecond is 0
	*/

	fmt.Println()
	fmt.Println("2018-01")
	fmt.Println()

	t2 := time.Date(2018, 1, 0, 0, 0, 0, 0, locUTC)

	ex.PrintDateTime(t2, dt.FmtDateTimeYrMDayFmtStr)
	/* Result - Best Approach - Use 1 as month number instead of zero month number.
	Also - Use Zero Days. Convert days to duration and add the duration.
	2018-01

	----------------------------------
						 Date Time
	----------------------------------
	Date Time:  2017-12-31 00:00:00.000000000 +0000 UCT
	The integer month is:  12
	The integer day is: 31
	The integer year is: 2017
	The integer hour is: 0
	The integer minute is: 0
	The integer second is: 0
	The integer nanosecond is 0
	*/

	fmt.Println()
	fmt.Println("Add 1 Day")
	fmt.Println()

	dur := int64(24) * dt.HourNanoSeconds

	t3 := t2.Add(time.Duration(dur))

	ex.PrintDateTime(t3, dt.FmtDateTimeYrMDayFmtStr)
	/*
		Add 1 Day to	2017-12-31 00:00:00.000000000 +0000 UCT
		Gives desired result

		----------------------------------
							 Date Time
		----------------------------------
		Date Time:  2018-01-01 00:00:00.000000000 +0000 UCT
		The integer month is:  1
		The integer day is: 1
		The integer year is: 2018
		The integer hour is: 0
		The integer minute is: 0
		The integer second is: 0
		The integer nanosecond is 0
	*/

}

func (mt mainTest) mainTest005() {

	locUTC, _ := time.LoadLocation(dt.TZones.Other.UTC())

	tDateTime := time.Date(2018, 2, 0, 0, 0, 0, 0, locUTC)

	ex.PrintDateTime(tDateTime, dt.FmtDateTimeYrMDayFmtStr)

	fmt.Println()
	fmt.Println("Adding 3-days")
	fmt.Println()

	dur := int64(3) * dt.DayNanoSeconds
	t2 := tDateTime.Add(time.Duration(dur))

	ex.PrintDateTime(t2, dt.FmtDateTimeYrMDayFmtStr)

	expectedDt := time.Date(2018, 2, 3, 0, 0, 0, 0, locUTC)

	fmt.Println()
	fmt.Println("Complete Date 2018-02-03")
	fmt.Println()

	ex.PrintDateTime(expectedDt, dt.FmtDateTimeYrMDayFmtStr)

}

func (mt mainTest) mainTest004() {
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	locUSCentral, _ := time.LoadLocation(dt.TZones.US.Central())
	locUSPacific, _ := time.LoadLocation(dt.TZones.US.Pacific())
	locParis, _ := time.LoadLocation(dt.TZones.Europe.Paris())
	locCairo, _ := time.LoadLocation(dt.TZones.Africa.Cairo())
	locMoscow, _ := time.LoadLocation(dt.TZones.Europe.Moscow())
	locTokyo, _ := time.LoadLocation(dt.TZones.Asia.Tokyo())

	t1USCentral := time.Date(1948, time.Month(9), 7, 4, 32, 16, 8185431, locUSCentral)
	t1USPacific := t1USCentral.In(locUSPacific)
	t1EuropeParis := t1USPacific.In(locParis)
	t1AfricaCairo := t1EuropeParis.In(locCairo)
	t1EuropeMoscow := t1AfricaCairo.In(locMoscow)
	t1AsiaTokyo := t1EuropeMoscow.In(locTokyo)
	t1bUSCentral := t1AsiaTokyo.In(locUSCentral)

	fmt.Println("t1USCentral: ", t1USCentral.Format(fmtStr))
	fmt.Println("t1USPacific: ", t1USPacific.Format(fmtStr))
	fmt.Println("t1EuropeParis: ", t1EuropeParis.Format(fmtStr))
	fmt.Println("t1AfricaCairo: ", t1AfricaCairo.Format(fmtStr))
	fmt.Println("t1EuropeMoscow: ", t1EuropeMoscow.Format(fmtStr))
	fmt.Println("t1AsiaTokyo: ", t1AsiaTokyo.Format(fmtStr))
	fmt.Println("t1bUSCentral: ", t1bUSCentral.Format(fmtStr))

}

func (mt mainTest) mainTest003() {
	loc, _ := time.LoadLocation(dt.TZones.US.Central())

	t1 := time.Date(2014, time.Month(15), 67, 19, 54, 30, 158712300, loc)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	addYear := 0
	addMonth := 15
	addDay := 64
	addHours := 0
	addMinutes := 0
	addSeconds := 0
	addMilliSeconds := 0
	addMicroSeconds := 0
	addNanoSeconds := 0

	var totDuration int64

	t2 := t1.AddDate(addYear, addMonth, addDay)

	totDuration = int64(addHours) * int64(time.Hour)
	totDuration += int64(addMinutes) * int64(time.Minute)
	totDuration += int64(addSeconds) * int64(time.Second)
	totDuration += int64(addMilliSeconds) * int64(time.Millisecond)
	totDuration += int64(addMicroSeconds) * int64(time.Microsecond)
	totDuration += int64(addNanoSeconds)

	t3 := t2.Add(time.Duration(totDuration))

	fmt.Println("t1: ", t1.Format(fmtstr))
	fmt.Println("t2: ", t2.Format(fmtstr))
	fmt.Println("t2: ", t3.Format(fmtstr))

}

func (mt mainTest) mainTest002() {

	tDto, err := dt.TimeDto{}.New(0, 0, -8, 0, 0, 0, 0, 0, 0, 0)

	if err != nil {
		fmt.Printf("Error returned from TimeDto{}.New(0, 0, -8, 0, 0, 0, 0, 0, 0, 0 ) Error='%v' \n", err.Error())
	}

	ex.PrintOutTimeDtoFields(tDto)

}

func (mt mainTest) mainTest001() {

	loc, _ := time.LoadLocation(dt.TZones.US.Central())
	t1 := time.Date(2014, time.Month(2), 15, 19, 54, 30, 158712300, loc)
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	tDto, err := dt.TimeDto{}.New(2014, 2, 0, 15, 19, 54, 30, 0, 0, 158712300)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeDto{}.New(year, month, ...). Error=%v \n", err.Error())
		return
	}

	t2, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned by tDto.GetDateTime(dt.TZones.US.Central()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Println("t1: ", t1.Format(fmtstr))
	fmt.Println("t2: ", t2.Format(fmtstr))

}

func (mt mainTest) mainPrintHdr(textToPrint string, repeatStr string) {
	title := fmt.Sprintf("       %v         ", textToPrint)
	ln := strings.Repeat(repeatStr, len(title))
	fmt.Println(ln)
	fmt.Println(title)
	fmt.Println(ln)
	fmt.Println()

}
