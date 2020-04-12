package main

import (
	"fmt"
	dt "github.com/MikeAustin71/datetimeopsgo/datetime"
	ex "github.com/MikeAustin71/datetimeopsgo/datetimeexamples"
	"strings"
	"time"
)

func main() {

	mainTest{}.mainTest077()

}

type mainTest struct {
	input  string
	output string
}

func (mt mainTest) mainTest077() {

	lineLen := 70
	titles := []string{"mainTest.mainTest077()",
		"Testing Change To Standard Time",
		"November 2, 2014"}


	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	centralTz, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}


	t1_1 := time.Date(
		2014,
		10,
		31,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t1_1: '%v'\n", t1_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t1_2 := time.Date(
		2014,
		11,
		1,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t1_2: '%v'\n", t1_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t1Dur := t1_2.Sub(t1_1)

	fmt.Printf("t1Dur: '%v'\n", t1Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

	t2_1 := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t2_1: '%v'\n", t2_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t2_2 := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t2_2: '%v'\n", t2_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t2Dur := t2_2.Sub(t2_1)

	fmt.Printf("t2Dur: '%v'\n", t2Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

	t3_1 := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("t3_1: '%v'\n", t3_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t3_2 := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("t3_2: '%v'\n", t3_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t3Dur := t3_2.Sub(t3_1)

	fmt.Printf("t3Dur (UTC): '%v'\n", t3Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

}

func (mt mainTest) mainTest076() {

	lineLen := 70
	titles := []string{"mainTest.mainTest076()",
		"Testing Change to Daylight Savings Time",
		"March 9, 2014"}


	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	centralTz, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1_1 := time.Date(
		2014,
		3,
		7,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t1_1: '%v'\n", t1_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t1_2 := time.Date(
		2014,
		3,
		8,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t1_2: '%v'\n", t1_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t1Dur := t1_2.Sub(t1_1)

	fmt.Printf("t1Dur: '%v'\n", t1Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

	t2_1 := time.Date(
		2014,
		3,
		9,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t2_1: '%v'\n", t2_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t2_2 := time.Date(
		2014,
		3,
		10,
		0,
		0,
		0,
		0,
		centralTz)

	fmt.Printf("t2_2: '%v'\n", t2_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t2Dur := t2_2.Sub(t2_1)

	fmt.Printf("t2Dur: '%v'\n", t2Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

	t3_1 := time.Date(
		2014,
		3,
		9,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("t3_1: '%v'\n", t3_1.Format(dt.DEFAULTDATETIMEFORMAT))

	t3_2 := time.Date(
		2014,
		3,
		10,
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("t3_2: '%v'\n", t3_2.Format(dt.DEFAULTDATETIMEFORMAT))

	t3Dur := t3_2.Sub(t3_1)

	fmt.Printf("t3Dur (UTC): '%v'\n", t3Dur)
	fmt.Println("------------------------------------")
	fmt.Println()
	fmt.Println()

}


func (mt mainTest) mainTest075() {
	// Variation on:
	//   TestDurationTriad_NewStartEndDateTzDto_01
	// datetime\zztdurationtriad01_test.go

	lineLen := 70
	titles := []string{"mainTest.mainTest075()",
		"TestDurationTriad_NewStartEndDateTzDto_01",
		"DurationTriad and Time Duration Dto"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := dt.DateTzDto{}.NewDateTime(t1, dt.FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := dt.DateTzDto{}.NewDateTime(t2, dt.FmtDateTimeYrMDayFmtStr)

	dur, err := dt.DurationTriad{}.NewStartEndDateTzDto(dateTz1, dateTz2)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		fmt.Printf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		fmt.Printf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.BaseTime.TimeDuration {
		fmt.Printf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v", tOutDur, dur.BaseTime.TimeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		fmt.Printf("Error - Expected Default Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, _ = dur.BaseTime.GetCumDaysTimeStr()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumHoursTimeStr()

	if err != nil {
		fmt.Printf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumWeeksDaysTimeStr()

	if err != nil {
		fmt.Printf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}

	loc, err := time.LoadLocation("Local")

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.GetOriginalTzName() {
		fmt.Printf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.GetOriginalTzName())
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.GetDateTimeValue()) {
		fmt.Printf("Expected Local Start Time ='%v'.\n" +
			"Actual Local Start Time ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.GetOriginalTzName())
	}

	t2Local := dur.LocalTime.StartTimeDateTz.GetDateTimeValue().Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.GetDateTimeValue()) {
		fmt.Printf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(dt.FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(dt.TZones.UTC())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(TZones.UTC()). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.GetDateTimeValue()) {
		fmt.Printf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(dt.FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.GetDateTimeValue().Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.GetDateTimeValue()) {
		fmt.Printf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(dt.FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
	}


}

func (mt mainTest) mainTest074() {
	// Variation on:
	// TestDurationTriad_GetYearMthDaysTimeAbbrv
	// datetime\zztdurationtriad01_test.go
	lineLen := 70
	titles := []string{"mainTest.mainTest074()",
		"TestDurationTriad_GetYearMthDaysTimeAbbrv",
		"DurationTriad and Time Duration Dto"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := dt.DurationTriad{}

	err := du.SetStartEndTimesTz(t1, t2, dt.TZones.US.Central(), dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by du.SetStartEndTimesTz(t2, t1, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr). Error='%v' ", err.Error())
	}

	fmt.Printf("Base Time Allocated Nanoseconds= '%v'\n",
		du.BaseTime.Nanoseconds)

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearMthDaysTimeAbbrvStr()

	if expected != dOut {
		fmt.Printf("Expected result= '%v'.\n" +
			                " Instead result= '%v'.\n",
			                expected, dOut)
	}

}

func (mt mainTest) mainTest073() {
	// Variation on:
	// TestDateTzDto_AddDate_01
	// datetime\zztdatetzdto01_test.go
	lineLen := 70
	titles := []string{"mainTest.mainTest073()",
		"Testing AddDateTimeByUtc",
		"Adding \"5-Years, 6-Months 12-Days\""}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")


	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	locationPtr, err := time.LoadLocation(dt.TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	t1 := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)

	t1Utc := t1.In(time.UTC)

	t1Result := t1.AddDate(5, 6, 12)

	t2Utc := t1Utc.AddDate(5, 6, 12)

	t2Result := t2Utc.In(locationPtr)

	titles = []string{"t1 Result",
		"Using Go Date Time Package Addition",
		"Adding 5-Years, 6-Months 12-Days"}

	ex.PrintOutDateTimeTimeZoneFields(
		t1Result,
		titles,
		lineLen,
		fmtStr)

	titles = []string{"t2 Result",
		"Using UTC Conversion",
		"Adding 5-Years, 6-Months 12-Days"}

	ex.PrintOutDateTimeTimeZoneFields(
		t2Result,
		titles,
		lineLen,
		fmtStr)

timeDto := dt.TimeDto{
	Years:                5,
	Months:               6,
	Weeks:                0,
	WeekDays:             0,
	DateDays:             12,
	Hours:                0,
	Minutes:              0,
	Seconds:              0,
	Milliseconds:         0,
	Microseconds:         0,
	Nanoseconds:          0,
	TotSubSecNanoseconds: 0,
	TotTimeNanoseconds:   0,
}

	dur, err := dt.DurationTriad{}.NewEndTimeMinusTimeDtoTz(
		dt.TCalcMode.LocalTimeZone(),
		t2Result,
		timeDto,
		dt.TZones.America.Chicago(),
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("\nCalculated Start Time: %v\n",
		dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtStr))

	fmt.Printf("\n    Actual Start Time: %v\n\n",
		t1.Format(fmtStr))


	fmt.Printf("\nCalculated End Time: %v\n",
		dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtStr))

	fmt.Printf("\n    Actual End Time: %v\n\n",
		t2Result.Format(fmtStr))

	fmt.Printf("\nCalculated UTC Start Time: %v\n",
		dur.UTCTime.StartTimeDateTz.GetDateTimeValue().Format(fmtStr))

	fmt.Printf("\n    Actual UTC Start Time: %v\n\n",
		t1Utc.Format(fmtStr))

	fmt.Printf("\nCalculated UTC End Time: %v\n",
		dur.UTCTime.EndTimeDateTz.GetDateTimeValue().Format(fmtStr))

	fmt.Printf("\n    Actual UTC End Time: %v\n\n",
		t2Utc.Format(fmtStr))

	fmt.Printf("\nCalculated BaseTime Duration: %v\n\n",
		dur.BaseTime.GetYearMthDaysTimeAbbrvStr())

	fmt.Printf("\n     Calculated UTC Duration: %v\n\n",
		dur.UTCTime.GetYearMthDaysTimeAbbrvStr())

	t2DurationDate := t1.Add(dur.UTCTime.TimeDuration)

	fmt.Printf("\n t1 + UTCTime Durtion: %v\n\n",
		t2DurationDate.Format(fmtStr))

	timeDur :=t1Result.Sub(t1)

	timeDurDto, err := dt.TimeDurationDto{}.NewStartTimeDuration(
	t1,
	timeDur,
	fmtStr)

	if err != nil {
		fmt.Printf("Error retured by dt.TimeDurationDto{}.NewStartTimeDuration().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("\n                    t1 time: %v\n",
		t1.Format(fmtStr))

	fmt.Printf("\n             t1 result time: %v\n",
		t1Result.Format(fmtStr))
	
	fmt.Printf("\n     Calculated t1 Duration: %v\n\n",
		timeDurDto.GetYearMthDaysTimeAbbrvStr())


}

func (mt mainTest) mainTest072() {
	// TestDateTzDto_AddDate_01
	// datetime\zztdatetzdto01_test.go


	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	expectedOutDate := t1.Format(fmtstr)

	dtz1, err := dt.DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewStartEndTimes(t1, fmtstr).\n" +
			"Error='%v'", err.Error())
		return
	}

	if expectedOutDate != dtz1.String() {
		fmt.Printf("Error: Expected dtz1.String()='%v'. Instead, dtz1.String()='%v' ", expectedOutDate, dtz1.String())
		return
	}

	t2 := t1.AddDate(5, 6, 12)

	dtz2, err := dtz1.AddDate(
		dt.TCalcMode.LocalTimeZone(),
		5,
		6,
		12,
		fmtstr)

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

}

func (mt mainTest) mainTest071() {
	// TestDurationTriad_NewStartTimeMinusTime_01
	// \datetime\zztdurationtriad01_test.go
	//

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := dt.TimeDto{
		Years: 3, 
		Months: 2, 
		Weeks: 2, 
		WeekDays: 1, 
		Hours: 3, 
		Minutes: 4, 
		Seconds: 2}

	dur, err := dt.DurationTriad{}.NewEndTimeMinusTimeDtoTz(
		 dt.TCalcMode.LocalTimeZone(),
			t2,
			timeDto,
			dt.TZones.US.Central(),
			dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		fmt.Printf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		fmt.Printf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.TimeDuration {
		fmt.Printf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.TimeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		fmt.Printf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}

func (mt mainTest) mainTest070() {
	//t1str := "2020-03-07 03:00:00.000000000 -0500 EST"

	easternTzPtr, err := time.LoadLocation(dt.TZones.America.New_York())

	if err != nil {
		fmt.Printf("Error returned by " +
			"time.LoadLocation(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1 := time.Date(
		2020,
		time.Month(3),
		7,
		3,
		0,
		0,
		0,
		easternTzPtr)

	expectedStr := "2020-03-08 04:00:00.000000000 -0400 EDT"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest070()"

	lineLen := 70
	titles := []string{funcName,
		"Testing DateTimeOps DateTimeMechanics AddDateTimeByUtc",
		"Adding 1-Day To Date Time With IANA Time Zone",
	"Notice That Actual Matches Expected Result"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")


	titles = []string{"Time-1 (t1)",
		"Starting Date 1-day before Daylight Standard Time"}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		titles,
		lineLen,
		fmtStr)

	dtMech := dt.DTimeMechanics{}

	t2 := dtMech.AddDateTimeByUtc(
		t1,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0)

	titles = []string{"Time-2 (t2)",
		"Actual Date 1-day After Daylight Standard Time",
		"Includes added value of 1-Day"}

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		titles,
		lineLen,
		fmtStr)

	lineBreak := strings.Repeat("*", lineLen)
	fmt.Println()
	fmt.Println(lineBreak)
	fmt.Println("        Starting Date Time: ", t1.Format(fmtStr))
	fmt.Println("  Actual After 1-Day Added: ", t2.Format(fmtStr))
	fmt.Println("Expected After 1-Day Added: ", expectedStr)
	fmt.Println(lineBreak)
	fmt.Println()

	titles = []string{funcName, "Successful Completion!"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

}

func (mt mainTest) mainTest069() {
	// t1str := "2020-03-07 03:00:00.000000000 -0500 EST"
	// EST Time Zone Fails

	easternTzPtr, err := time.LoadLocation(dt.TZones.America.New_York())

	if err != nil {
		fmt.Printf("Error returned by " +
			"time.LoadLocation(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1 := time.Date(
		2020,
		time.Month(3),
		7,
		3,
		0,
		0,
		0,
		easternTzPtr)

	expectedStr := "2020-03-08 04:00:00.000000000 -0400 EDT"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest069()"

	lineLen := 70
	titles := []string{funcName,
		"Testing DateTimeOps DateTimeMechanics AddDateTimeByUtc",
		"Adding 1-Day",
		"Notice That Actual Matches Expected Result",
		"When Using IANA Convertible Time Zone"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	titles = []string{"Time-1 (t1)",
		"Starting Date 1-day before Daylight Standard Time"}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		titles,
		lineLen,
		fmtStr)

	dtMech := dt.DTimeMechanics{}

	t2 := dtMech.AddDateTimeByUtc(
		t1,
		0,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		0)

	titles = []string{"Time-2 (t2)",
		"Actual Date 1-day After Daylight Standard Time",
		"Includes added value of 1-Day"}

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		titles,
		lineLen,
		fmtStr)

	lineBreak := strings.Repeat("*", lineLen)
	fmt.Println()
	fmt.Println(lineBreak)
	fmt.Println("        Starting Date Time: ", t1.Format(fmtStr))
	fmt.Println("  Actual After 1-Day Added: ", t2.Format(fmtStr))
	fmt.Println("Expected After 1-Day Added: ", expectedStr)
	fmt.Println(lineBreak)

	var funcResult string

	if t2.Format(fmtStr) != expectedStr {
		titles = []string{"FAILURE!","Expected Date Time Does NOT Match Actual Date Time!"}
		ex.PrintMainHeader(
			titles,
			lineLen,
			"",
			"*")

		funcResult = "@@ FAILURE @@"

	} else {
		titles = []string{"SUCCESS!","Expected Date Time DOES Match Actual Date Time!"}
		ex.PrintMainHeader(
			titles,
			lineLen,
			"",
			"*")

		funcResult = "Successful Completion!"
	}

	fmt.Println()

	titles = []string{funcName, funcResult}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

}

func (mt mainTest) mainTest068() {
	// 8 Mar 2020 - Daylight Saving Time Starts
	// When local standard time is about to reach
	// Sunday, 8 March 2020, 02:00:00 clocks are turned forward 1 hour to
	//Sunday, 8 March 2020, 03:00:00 local daylight time instead.

	// t1str := "2020-03-07 03:01:00.000000000 -0500 EST"


	easternTzPtr, err := time.LoadLocation(dt.TZones.America.New_York())

	if err != nil {
		fmt.Printf("Error returned by " +
			"time.LoadLocation(dt.TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1 := time.Date(
		2020,
		time.Month(3),
		7,
		3,
		0,
		0,
		0,
		easternTzPtr)

	expectedStr := "2020-03-08 04:00:00.000000000 -0400 EDT"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest068()"

	lineLen := 70
	titles := []string{funcName,
		"Testing Golang DateTime Package AddDate",
		"Adding 1-Day",
		"Notice that the Final Date is WRONG!"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	titles = []string{"Time-1 (t1)",
		"Starting Date 1-day before Daylight Standard Time"}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		titles,
		lineLen,
		fmtStr)

	t2 := t1.AddDate(0, 0, 1)

	titles = []string{"Time-2 (t2)",
		"Actual Date 1-day After Daylight Standard Time",
	"Includes added value of 1-Day"}

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		titles,
		lineLen,
		fmtStr)

	lineBreak := strings.Repeat("*", lineLen)
	fmt.Println()
	fmt.Println(lineBreak)
	fmt.Println("        Starting Date Time: ", t1.Format(fmtStr))
	fmt.Println("  Actual After 1-Day Added: ", t2.Format(fmtStr))
	fmt.Println("Expected After 1-Day Added: ", expectedStr)
	fmt.Println(lineBreak)

	var funcResult string

	if t2.Format(fmtStr) != expectedStr {
		titles = []string{"Expected Date Time Does NOT Match Actual Date Time!"}
		ex.PrintMainHeader(
			titles,
			lineLen,
			"",
			"*")

		funcResult = "@@ FAILURE @@"

	} else {
		titles = []string{"Expected Date Time DOES Match Actual Date Time!"}
		ex.PrintMainHeader(
			titles,
			lineLen,
			"",
			"*")

		funcResult = "Successful Completion!"
	}

	fmt.Println()

	titles = []string{funcName, funcResult}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

}

func (mt mainTest) mainTest067() {

// 8 Mar 2020 - Daylight Saving Time Starts
// When local standard time is about to reach
// Sunday, 8 March 2020, 02:00:00 clocks are turned forward 1 hour to
//Sunday, 8 March 2020, 03:00:00 local daylight time instead.

		t1str := "2020-03-08 01:00:00.000000000 -0600 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest067()"

	lineLen := 65
	titles := []string{funcName,
		"Edge Of Daylight Time",
	"Adding 2-Hours"}

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineBreak := strings.Repeat("-", lineLen)
	lineBreak2 := strings.Repeat("*", lineLen)
	t1, _ := time.Parse(fmtStr, t1str)

	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.NewStartEndTimes(t1, TzUsPacific).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var pacificPtr, utcPtr, localPtr *time.Location

	pacificPtr, err = time.LoadLocation(dt.TZones.US.Pacific())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.US.Pacific()).\n" +
			"dt.TZones.US.Pacific()='%v'\n" +
			"Error='%v'\n", dt.TZones.US.Pacific(), err.Error())
		return
	}

	utcPtr, err = time.LoadLocation(dt.TZones.UTC())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.UTC()).\n" +
			"dt.TZones.UTC()='%v'\n" +
			"Error='%v'\n", dt.TZones.UTC(), err.Error())
		return
	}

	localPtr, err = time.LoadLocation(dt.TZones.Local())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.Local()).\n" +
			"dt.TZones.UTC()='%v'\n" +
			"Error='%v'\n", dt.TZones.Local(), err.Error())
		return
	}

	t1In := t1
	t1Out := t1.In(pacificPtr)
	t1Utc := t1.In(utcPtr)
	t1Local := t1.In(localPtr)
	fmt.Println(lineBreak)
	fmt.Println(lineBreak)
	fmt.Println("   Calculated t1 In: ", t1In.Format(fmtStr))
	fmt.Println("  Calculated t1 Out: ", t1Out.Format(fmtStr) )
	fmt.Println("  Calculated t1 UTC: ", t1Utc.Format(fmtStr) )
	fmt.Println("Calculated t1 Local: ", t1Local.Format(fmtStr) )
	fmt.Println(lineBreak)
	fmt.Println()
	ex.PrintOutTimeZoneDtoFields(tzu1, "tzu1 - Before Time Addition")
	fmt.Println(lineBreak2)
	fmt.Println()

	err = tzu1.AddTime(
		2,
		0,
		0,
		0,
		0,
		0)

	if err != nil {
		fmt.Printf("Error returned by tzu1.AddTime(Add 2-hours)\n" +
			"Error='%v'\n", err.Error())
		return
	}
	fmt.Println(lineBreak2)
	ex.PrintOutTimeZoneDtoFields(tzu1, "tzu1 - After 2-Hour Addition")
	fmt.Println(lineBreak2)
	fmt.Println()
}

func (mt mainTest) mainTest066() {
	t1str := "2014-02-15 19:54:30.000000000 -0500 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
	funcName := "mainTest066()"

	lineLen := 65
	titles := []string{funcName,
		"Time Zone Dto Comparison" }

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	lineBreak := strings.Repeat("-", lineLen)
	lineBreak2 := strings.Repeat("*", lineLen)

	t1, _ := time.Parse(fmtStr, t1str)
	t1OutStr := t1.Format(fmtStr)
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.NewStartEndTimes(t1, TzUsPacific).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var pacificPtr, utcPtr, localPtr *time.Location

	pacificPtr, err = time.LoadLocation(dt.TZones.US.Pacific())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.US.Pacific()).\n" +
			"dt.TZones.US.Pacific()='%v'\n" +
			"Error='%v'\n", dt.TZones.US.Pacific(), err.Error())
		return
	}

	utcPtr, err = time.LoadLocation(dt.TZones.UTC())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.UTC()).\n" +
			"dt.TZones.UTC()='%v'\n" +
			"Error='%v'\n", dt.TZones.UTC(), err.Error())
		return
	}

	localPtr, err = time.LoadLocation(dt.TZones.Local())

	if err != nil {
		fmt.Printf("Error returned from " +
			"time.LoadLocation(dt.TZones.Local()).\n" +
			"dt.TZones.UTC()='%v'\n" +
			"Error='%v'\n", dt.TZones.Local(), err.Error())
		return
	}

	t1Out := t1.In(pacificPtr)
	t1Utc := t1.In(utcPtr)
	t1Local := t1.In(localPtr)

	fmt.Println(lineBreak)
	fmt.Println("   Expected t1 In: ", t1.Format(fmtStr))
	fmt.Println("  Expected t1 Out: ", t1Out.Format(fmtStr) )
	fmt.Println("  Expected t1 UTC: ", t1Utc.Format(fmtStr) )
	fmt.Println("Expected t1 Local: ", t1Local.Format(fmtStr) )
	fmt.Println(lineBreak)
	fmt.Println()
	ex.PrintOutTimeZoneDtoFields(tzu1, "tzu1")
	fmt.Println(lineBreak)
	fmt.Println()


	tzu1OutStrTIn := tzu1.TimeIn.GetDateTimeValue().Format(fmtStr)

	if t1OutStr != tzu1OutStrTIn {
		fmt.Printf("Error: Expected tzu1OutStrTIn='%v'.  Instead, tzu1OutStrTIn='%v'", t1OutStr, tzu1OutStrTIn)
	}

	t2 := t1.AddDate(3, 2, 15)
	// t2OutStr := t2.Format(fmtStr)

	tzu2, err := dt.TimeZoneDto{}.NewAddDate(tzu1, 3, 2, 15, fmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDto{}.NewAddDate(tzu1, 3, 2, 15, fmtStr)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t2Out := t2.In(pacificPtr)
	t2Utc := t2.In(utcPtr)
	t2Local := t2.In(localPtr)

	t3Utc := t1Utc.AddDate(3, 2, 15)

	fmt.Println(lineBreak2)
	fmt.Println("   Expected t2 In: ", t2.Format(fmtStr))
	fmt.Println("  Expected t2 Out: ", t2Out.Format(fmtStr) )
	fmt.Println("  Expected t2 UTC: ", t2Utc.Format(fmtStr) )
	fmt.Println("Expected t2 Local: ", t2Local.Format(fmtStr) )
	fmt.Println("      t3Utc Added: ", t3Utc.Format(fmtStr))
	fmt.Println(lineBreak)
	fmt.Println()
	ex.PrintOutTimeZoneDtoFields(tzu2, "tzu2")
	fmt.Println(lineBreak2)


	/*
		tzu2OutStrTIn := tzu2.TimeIn.GetDateTimeValue().Format(fmtStr)

		if t2OutStr != tzu2OutStrTIn {
			fmt.Printf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
			return
		}

		actualDuration, err := tzu2.Sub(tzu1)

		if err != nil {
			fmt.Printf("Error returned by tzu2.Sub(tzu1). Error='%v'", err.Error())
			return
		}

		utcPtr, err := time.LoadLocation(dt.TZones.UTC())

		if err != nil{
			fmt.Printf("Error return from time.LoadLocation(TZones.UTC()).\n" +
				"Error='%v'\n", err.Error())
			return
		}

		t1UTC := t1.In(utcPtr)

		t2UTC := t2.In(utcPtr)

		expectedDuration := t2UTC.Sub(t1UTC)

		if expectedDuration != actualDuration {
			fmt.Printf("Error: Expected Duration='%v'.\n" +
				"Instead, Actual Duration='%v'\n",
				expectedDuration, actualDuration)
		}

		*/

}

func (mt mainTest) mainTest065() {

	funcName := "mainTest065()"
	lineLen := 65
	titles := []string{funcName,
											"UTC Duration Comparison" }

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

	t1str := "2014-02-14 19:54:30.000000000 -0600 CST"
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, err := time.Parse(fmtStr, t1str)

	if err != nil {
		fmt.Printf("Error returned by time.Parse(fmtStr, t1str)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	var utcPtr *time.Location

	utcPtr, err = time.LoadLocation(dt.TZones.UTC())

	if err != nil{
		fmt.Printf("Error return from time.LoadLocation(TZones.UTC()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	equalHeader := strings.Repeat("=", 65)
	fmt.Println(equalHeader)
	fmt.Println("   t1 Start Date Time: ", t1.Format(fmtStr))
	t1Utc := t1.In(utcPtr)
	fmt.Println("     t1 UTC Date Time: ", t1Utc.Format(fmtStr))
	fmt.Println(equalHeader)
	fmt.Println()

	t2 := t1.AddDate(3, 2, 15)
	fmt.Println(equalHeader)
	fmt.Println("   t2 Start Date Time: ", t2.Format(fmtStr))
	t2Utc := t2.In(utcPtr)
	fmt.Println("     t2 UTC Date Time: ", t2Utc.Format(fmtStr))
	fmt.Println(equalHeader)
	fmt.Println()

	t1_2BaseDuration := t2.Sub(t1)
	fmt.Println("  t1-t2 Base Duration: ", t1_2BaseDuration.String())

	t1_2UtcDuration := t2Utc.Sub(t1Utc)
	fmt.Println("   t1-t2 Utc Duration: ", t1_2UtcDuration.String())
	fmt.Println(equalHeader)
	fmt.Println()


	titles = []string{funcName,
		"Successful Completion" }

	ex.PrintMainHeader(
		titles,
		lineLen,
		"=",
		"=")

}

func (mt mainTest) mainTest064() {
	// TestDateTzDto_AddDate_01

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	expectedOutDate := t1.Format(fmtstr)

	dtz1, err := dt.DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		fmt.Printf("Error returned by DateTzDto{}.NewStartEndTimes(t1, fmtstr).\n" +
			"Error='%v'", err.Error())
		return
	}

	if expectedOutDate != dtz1.String() {
		fmt.Printf("Error: Expected dtz1.String()='%v'. Instead, dtz1.String()='%v' ", expectedOutDate, dtz1.String())
	}

	t2 := t1.AddDate(5, 6, 12)

	dtz2, err := dtz1.AddDate(
		dt.TCalcMode.LocalTimeZone(),
		5,
		6,
		12,
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by dtz1.AddDate(5, 6, 12, fmtstr). Error='%v'", err.Error())
		return
	}

	expectedOutDate = t2.Format(fmtstr)

	if expectedOutDate != dtz2.String() {
		fmt.Printf("Error: Expected dtz2.String()='%v'. Instead, dtz2.String()='%v' ", expectedOutDate, dtz2.String())
	}

}

func (mt mainTest) mainTest063() {

	ePrefix := "mainTest063() TestTDurCalcTypeString_034"

	mt.mainPrintHdr(ePrefix , "-")

	expectedStr := "XRAYxxx"

	_, err :=  dt.TDurCalc.XParseString(expectedStr, false)

	if err == nil {
		fmt.Printf("Error: Expected an 'error' return from " +
			"TDurCalc.XParseString(expectedStr, false).\n" +
			"expectedStr='%v'.\n" +
			"NO ERROR WAS RETURNED!!! ",expectedStr)
		return
	}

	mt.mainPrintHdr("Successful Completion" , "=")
}

func (mt mainTest) mainTest062() {

	ePrefix := "mainTest062() TimeZoneDefinition"

	mt.mainPrintHdr(ePrefix , "-")

	lineLen := 65

	utcOffset := "2020-01-19 04:21:18 +0700 +07"
	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	expectedOriginalTz := "Etc/GMT-7"
	expectedConvertibleTz := dt.TZones.Asia.Ho_Chi_Minh()

	utcOffsetTime, err := time.Parse(fmtStr, utcOffset)

	if err != nil {
		fmt.Printf("Received error from time parse utcOffset: %v\n",
			err.Error())
		return
	}

	var tzDef dt.TimeZoneDefinition

	tzDef, err = dt.TimeZoneDefinition{}.NewDateTime(utcOffsetTime)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefinition{}.NewStartEndTimes(utcOffsetTime)\n" +
			"utcOffsetTime= '%v'\n" +
			"Error='%v'\n",
			utcOffsetTime.Format(fmtStr), err.Error())
		return
	}

	originalTzSpec := tzDef.GetOriginalTimeZone()
	
	ex.PrintOutDateTimeTimeZoneFields(
		originalTzSpec.GetReferenceDateTime(),
		[]string{"Original TZ Date Time"},
		lineLen,
		fmtStr)


	ex.PrintOutTimeZoneSpecFields(originalTzSpec, "Original Time Zone")


	convertibleTzSpec := tzDef.GetConvertibleTimeZone()
	
	ex.PrintOutDateTimeTimeZoneFields(
		convertibleTzSpec.GetReferenceDateTime(),
		[]string{"Convertible TZ Date Time"},
		lineLen,
		fmtStr)

	ex.PrintOutTimeZoneSpecFields(convertibleTzSpec, "Convertible Time Zone")

	actualOriginalTz := tzDef.GetOriginalTimeZoneName()

	actualConvertibleTz := tzDef.GetConvertibleTimeZoneName()


	if expectedOriginalTz != actualOriginalTz {
		fmt.Printf("Error: Expected actualOriginalTz='%v'.\n" +
			"Instead, actualOriginalTz='%v'\n",
			expectedOriginalTz, actualOriginalTz)
		return
	}

	if expectedConvertibleTz != actualConvertibleTz {
		fmt.Printf("Error: Expected actualConvertibleTz='%v'.\n" +
			"Instead, actualConvertibleTz='%v'\n",
			expectedConvertibleTz, actualOriginalTz)
		return
	}

	mt.mainPrintHdr("Successful Completion" , "=")

}

func (mt mainTest) mainTest061() {

	ePrefix := "mainTest061() TimeZoneDto.ConvertTz()"

	mt.mainPrintHdr(ePrefix , "-")

	utcTime := "2017-04-30 00:54:30 +0000 UTC"
	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	mountainTime := "2017-04-29 18:54:30 -0600 MDT"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	ianaMountainTz := "America/Denver"
	tPacificIn, err := time.Parse(fmtstr, pacificTime)

	if err != nil {
		fmt.Printf("Received error from time parse tPacificIn: %v\n",
			err.Error())
		return
	}

	tzu := dt.TimeZoneDto{}
	tzuCentral, err := tzu.ConvertTz(
		tPacificIn,
		ianaCentralTz,
		dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("#1 Error from ianaCentralTz TimeZoneDto.ConvertTz().\n" +
			"Error: %v\n", err.Error())
		return
	}

	centralTOut := tzuCentral.TimeOut.GetDateTimeValue().Format(fmtstr)

	if centralTime != centralTOut {
		fmt.Printf("Expected tzuCentral.TimeOut = '%v'.\n" +
			"Instead, tzuCentral.TimeOut = '%v'.\n",
			centralTime, centralTOut)
		return
	}

	tzuMountain, err := tzu.ConvertTz(
		tzuCentral.TimeOut.GetDateTimeValue(),
		ianaMountainTz, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("#2 Error from tzuMountain TimeZoneDto.ConvertTz().\n" +
			"Error: %v\n", err.Error())
		return
	}

	mountainTOut := tzuMountain.TimeOut.GetDateTimeValue().Format(fmtstr)

	if mountainTime != mountainTOut {
		fmt.Printf("Expected tzuMountain.TimeOut= '%v'.\n" +
			"Instead, tzuMountain.TimeOut= '%v'.\n",
			mountainTime, mountainTOut)
		return
	}

	tzuPacific, err := tzu.ConvertTz(
		tzuMountain.TimeOut.GetDateTimeValue(),
		ianaPacificTz, dt.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		fmt.Printf("Error #3 from tzuMountain TimeZoneDto.ConvertTz().\n" +
			"Error: %v\n", err.Error())
		return
	}

	pacificTOut := tzuPacific.TimeOut.GetDateTimeValue().Format(fmtstr)

	if pacificTime != pacificTOut {

		fmt.Printf("Expected tzuPacific.TimeOut= '%v'.\n" +
			"Instead, tzuPacific.TimeOut= '%v'.\n",
			pacificTime, pacificTOut)
		return
	}

	exTOutLoc := "America/Los_Angeles"

	if exTOutLoc != tzuPacific.TimeOut.GetOriginalTzName() {
		fmt.Printf("Expected tzu.TimeOutLoc='%v'.\n" +
			"Instead tzu.TimeOutLoc='%v'\n" +
			"tzuPacific.TimeOut='%v'\n",
			exTOutLoc,
			tzuPacific.TimeOut.GetOriginalTzName(),
			tzuPacific.TimeOut.GetDateTimeValue().Format(dt.FmtDateTimeYrMDayFmtStr))
		return
	}

	pacificUtcOut := tzuPacific.TimeUTC.GetDateTimeValue().Format(fmtstr)

	if utcTime != pacificUtcOut {
		fmt.Printf("Expected tzuPacific.TimeUTC= '%v'\n" +
			"Instead, tzuPacific.TimeUTC= '%v'\n",
			utcTime, pacificUtcOut)
		return
	}

	centralUtcOut := tzuCentral.TimeUTC.GetDateTimeValue().Format(fmtstr)

	if utcTime != centralUtcOut {
		fmt.Printf("Expected tzuCentral.TimeUTC= '%v'\n" +
			"Instead, tzuCentral.TimeUTC= '%v'\n",
			utcTime, pacificUtcOut)
		return
	}

	mountainUtcOut := tzuMountain.TimeUTC.GetDateTimeValue().Format(fmtstr)

	if utcTime != mountainUtcOut {
		fmt.Printf("Expected tzuMountain.TimeUTC= '%v'\n" +
			"Instead, tzuMountain.TimeUTC= '%v'\n",
			utcTime, pacificUtcOut)
		return
	}

	fmt.Println()
	mt.mainPrintHdr("Successful Completion" , "=")
}


func (mt mainTest) mainTest060() {

	ePrefix := "mainTest060()"
	lineLen := 65

	mt.mainPrintHdr(ePrefix , "-")
	mt.mainPrintHdr("TimeZoneDto" , "-")

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtStr, t1str)
	t1OutStr := t1.Format(fmtStr)
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtStr)

	if err != nil {
		fmt.Printf("Error returned from TimeZoneDto{}.NewStartEndTimes(t1, TzUsPacific ). Error='%v'", err.Error())
		return
	}

	fmt.Println("t1OutStr: ", t1OutStr)

	ex.PrintOutDateTimeTimeZoneFields(
		tzu1.TimeIn.GetDateTimeValue(),
		[]string{"Initial tzu1.TimeIn"},
		lineLen,
		fmtStr)

	t2, _ := time.Parse(fmtStr, t2str)
	t2OutStr := t2.Format(fmtStr)

	t12Dur := t2.Sub(t1)

	tdurDto, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(
		t1,
		t2,
		dt.TDurCalcType(0).StdYearMth(),
		dt.TZones.US.Central(),
		fmtStr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}.NewStartEndTimesCalcTz()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddTimeDurationDto(tdurDto)

	if err != nil {
		fmt.Printf("Error returned by tzu2.AddTimeDurationDto(tdurDto).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	tzu1OutStr := tzu1.TimeIn.GetDateTimeValue().Format(fmtStr)

	if t1OutStr != tzu1OutStr {
		fmt.Printf("Error: Expected Time1 TimeIn='%v'.\n" +
			"Instead Time1 TimeIn='%v'\n",
			t1OutStr, tzu1OutStr)
		return
	}

	tzu2OutStr := tzu2.TimeIn.GetDateTimeValue().Format(fmtStr)

	if t2OutStr != tzu2OutStr {
		fmt.Printf("Error: Expected after duration tzu2TimeIn='%v'.\n" +
			"Instead, tzu2TimeIn='%v'\n",
			t2OutStr, tzu2OutStr)
		return
	}

	actualDur := tzu2.TimeIn.Sub(tzu1.TimeIn)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeIn.Sub(tzu1.TimeIn)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeOut.Sub(tzu1.TimeOut)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeOut.Sub(tzu1.TimeOut)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
	}

	actualDur = tzu2.TimeUTC.Sub(tzu1.TimeUTC)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeUTC.Sub(tzu1.TimeUTC)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeLocal.Sub(tzu1.TimeLocal)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeLocal.Sub(tzu1.TimeLocal)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualTimeOutLoc := tzu1.TimeOut.GetOriginalTzName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu1.TimeOutLoc='%v'.\n" +
			"Instead, tzu1.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	actualTimeOutLoc = tzu2.TimeOut.GetOriginalTzName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu2.TimeOutLoc.String()='%v'.\n" +
			"Instead, tzu2.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	mt.mainPrintHdr("Successful Completion" , "=")
}

func (mt mainTest) mainTest059() {
	ePrefix := "mainTest059()"

	mt.mainPrintHdr(ePrefix , "-")
	mt.mainPrintHdr("ConvertTz" , "-")

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	// Invalid Target Iana Time Zone
	invalidTz := "XUZ Time Zone"
	tIn, err := time.Parse(fmtstr, tstr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtstr, tstr)\n" +
			"tstr='%v'\n" +
			"Error='%v'\n", tstr, err.Error())
		return
	}

	tzu := dt.TimeZoneDto{}

	_, err = tzu.ConvertTz(tIn, invalidTz, fmtstr)

	if err == nil {
		fmt.Printf("ConvertTz() failed to detect INVALID Target Time Zone.\n" +
			"err=='nil'\n" +
			"invalidTz='%v'\n", invalidTz)
	}

	return
}

func (mt mainTest) mainTest058() {
	ePrefix := "mainTest058()"
	lineLen := 65

	mt.mainPrintHdr(ePrefix , "-")
	mt.mainPrintHdr("ConvertTzAbbreviationToTimeZone" , "-")

	tzAbbrv := "+10"
	testUtcOffset := "+1000"

	tzMech := dt.TimeZoneMechanics{}

	dateTime := time.Now().UTC()

	var staticTimeZone dt.TimeZoneSpecification
	var err error

	staticTimeZone,
	err = tzMech.ConvertUtcAbbrvToStaticTz(
		dateTime,
		dt.TzConvertType.Relative(),
		"Original Time Zone",
		testUtcOffset,
		ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	fmt.Println("staticTimeZone: ", staticTimeZone.GetLocationName())
	fmt.Println()
	ex.PrintOutDateTimeTimeZoneFields(
		staticTimeZone.GetReferenceDateTime(),
		[]string{"staticTimeZone"},
		lineLen,
		fmtStr)

	ex.PrintOutTimeZoneSpecFields(staticTimeZone, "staticTimeZone")

	var tzSpec dt.TimeZoneSpecification

	tzSpec,
	err =
	tzMech.ConvertTzAbbreviationToTimeZone(
		dateTime,
		dt.TimeZoneConversionType(0).Relative(),
		tzAbbrv+testUtcOffset,
		"",
		ePrefix)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by tzMech.ConvertTzAbbreviationToTimeZone()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(
		tzSpec.GetReferenceDateTime(),
		[]string{"tzSpec"},
		lineLen,
		fmtStr)

	ex.PrintOutTimeZoneSpecFields(tzSpec, "tzSpec")


	mt.mainPrintHdr("SUCCESS" , "!")

}

func (mt mainTest) mainTest057() {

	ePrefix := "mainTest057()"
	lineLen := 65

	mt.mainPrintHdr(ePrefix , "-")

	dtMech := dt.DTimeMechanics{}

	locPtr, err := dtMech.LoadTzLocation(dt.TZones.Asia.Vladivostok(), ePrefix)

	if err != nil {
		fmt.Printf("Error='%v'\n", err.Error())
		return
	}

	dateTime := time.Date(
		2019,
		time.Month(6),
		15,
		11,
		23,
		0,
		0,
		locPtr)

	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	tzDef, err := dt.TimeZoneDefinition{}.NewFromTimeZoneName(
		dateTime, dt.TZones.Asia.Vladivostok(), dt.TzConvertType.Absolute())

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(dateTime, London Time)\n" +
			"dateTime='%v'\nError='%v'\n",
			dateTime.Format(fmtStr), err.Error())
		return
	}

	ex.PrintOutTimeZoneDefFields(tzDef)


	// t1EdtStr :=  "06/20/2019 09:58:32 -0400 EDT"
	// t1EstStr :=  "12/20/2019 09:58:32 -0500 EST"
	// t2CdtStr :=  "06/20/2019 09:58:32 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0600 CST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0600 MDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0700 MST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0700 PDT"
	// t2PstStr :=  "12/20/2019 09:58:32 -0800 PST"

	timeStr := dateTime.Format(fmtStr)

	t2, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		[]string{"t2 Parse Result"},
		lineLen,
		fmtStr)

	tzDef, err = dt.TimeZoneDefinition{}.NewDateTime(t2)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(t1)\n" +
			"t2='%v'\nError='%v'\n", t2.Format(fmtStr), err.Error())
		return
	}

	ex.PrintOutTimeZoneDefFields(tzDef)

	mt.mainPrintHdr("SUCCESS" , "!")
}

func (mt mainTest) mainTest056() {

	ePrefix := "mainTest056()"

	mt.mainPrintHdr(ePrefix , "-")

	dtMech := dt.DTimeMechanics{}
	locUSCentral, err :=
		dtMech.LoadTzLocation(dt.TZones.US.Central(), ePrefix)

	if err != nil {
		fmt.Printf("Error='%v'\n", err.Error())
		return
	}

	locTokyo, err := dtMech.LoadTzLocation(dt.TZones.Asia.Tokyo(), ePrefix)

	if err != nil {
		fmt.Printf("Error='%v'", err.Error())
		return
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

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
		fmt.Printf("Error: Expected DateTime='%v'. Instead DateTime='%v'",
			t4USCentral.Format(fmtstr), dTz.GetDateTimeValue().Format(fmtstr))
	}

	eTimeZoneDef, err := dt.TimeZoneDefinition{}.NewDateTime(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeZoneDefinition{}.NewStartEndTimes(t4USCentral)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	areEqual := eTimeZoneDef.Equal(dTz.GetTimeZoneDef())

	if ! areEqual {
		fmt.Printf("Expected dTz.GetTimeZoneDef().LocationName='%v'.\n"+
			"Instead, dTz.GetTimeZoneDef().LocationName='%v'\n",
			eTimeZoneDef.GetOriginalLocationName(), dTz.GetOriginalTzName())
	}

	tDto, err := dt.TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		fmt.Printf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral)\n"+
			"t4USCentral='%v'\nError='%v'\n",
			t4USCentral.Format(dt.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	expectedDt, err := tDto.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from tDto.GetDateTime(TZones.US.Central()). "+
			"Error='%v'", err.Error())
	}

	timeComponents := dTz.GetTimeComponents()

	actualDt, err := timeComponents.GetDateTime(dt.TZones.US.Central())

	if err != nil {
		fmt.Printf("Error returned from dTz.GetTimeComponents().GetDateTime(TZones.US.Central()).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !tDto.Equal(dTz.GetTimeComponents()) {
		fmt.Printf("Expected dTz.Time (TimeDto) == '%v' Instead, dTz.Time (TimeDto) == '%v'",
			expectedDt.Format(dt.FmtDateTimeYrMDayFmtStr),
			actualDt.Format(dt.FmtDateTimeYrMDayFmtStr))
	}

	if dt.FmtDateTimeYrMDayFmtStr != dTz.GetDateTimeFmt() {
		fmt.Printf("Expected dTz.GetDateTimeFmt()='%v' Instead, dTz.GetDateTimeFmt()='%v' ",
			dt.FmtDateTimeYrMDayFmtStr, dTz.GetDateTimeFmt())
	}

}

func (mt mainTest) mainTest055() {

	ePrefix := "mainTest055()"

	mt.mainPrintHdr(ePrefix , "-")

	// t1EdtStr :=  "06/20/2019 09:58:32 -0400 EDT"
	// t1EstStr :=  "12/20/2019 09:58:32 -0500 EST"
	// t2CdtStr :=  "06/20/2019 09:58:32 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0600 CST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0600 MDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0700 MST"
	t2PdtStr :=  "06/20/2019 09:58:32 -0700 PDT"
	// t2PstStr :=  "12/20/2019 09:58:32 -0800 PST"

	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	timeStr := t2PdtStr
	t1, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	tzDef, err := dt.TimeZoneDefinition{}.NewDateTime(t1)

	if err != nil {
		fmt.Printf("Error returned by dt.TimeZoneDefinition{}.NewStartEndTimes(t1)\n" +
			"t1='%v'\nError='%v'\n", t1.Format(fmtStr), err.Error())
		return
	}

	ex.PrintOutTimeZoneDefFields(tzDef)

}



func (mt mainTest) mainTest054() {

	ePrefix := "mainTest054()"
	lineLen := 65
	mt.mainPrintHdr(ePrefix , "-")

	// t1EdtStr :=  "06/20/2019 09:58:32 -0400 EDT"
	// t1EstStr :=  "12/20/2019 09:58:32 -0500 EST"
	// t2CdtStr :=  "06/20/2019 09:58:32 -0500 CDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0600 CST"
	// t2PdtStr :=  "06/20/2019 09:58:32 -0600 MDT"
	// t2PdtStr :=  "12/20/2019 09:58:32 -0700 MST"
	t2PdtStr :=  "06/20/2019 09:58:32 -0700 PDT"
	// t2PstStr :=  "12/20/2019 09:58:32 -0800 PST"

	fmtStr := "01/02/2006 15:04:05 -0700 MST"
	timeStr := t2PdtStr
	t1, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Initial Date Time"},
		lineLen,
		fmtStr)

	t2 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		t1.Location())

	ex.PrintOutDateTimeTimeZoneFields(
		t2,
		[]string{"t2 Date Time"},
		lineLen,
		fmtStr)

	pacificTz := dt.TZones.America.Los_Angeles()

	pacificTzPtr, err := time.LoadLocation(pacificTz)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(pacificTz)\n" +
			"pacificTz='%v'\nError='%v'\n", pacificTz, err.Error())
		return
	}

	t3 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		pacificTzPtr)

	ex.PrintOutDateTimeTimeZoneFields(
		t3,
		[]string{"t3 Date Time"},
		lineLen,
		fmtStr)

}


func (mt mainTest) mainTest053() {

	ePrefix := "mainTest053()"

	mt.mainPrintHdr(ePrefix , "=")

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _ := dt.TimeZoneDto{}.New(tIn, ianaPacificTz, fmtstr)

	fmt.Println(" Time In: ", tzu.TimeIn.GetDateTimeValue().Format(fmtstr))
	fmt.Println("Time Out: ", tzu.TimeOut.GetDateTimeValue().Format(fmtstr))

	expectedZone := "PDT"

	actualZone := tzu.TimeOut.GetOriginalTzAbbreviation()

	if expectedZone != actualZone {
		fmt.Printf("Expected Zone Out='%v'.\n" +
			"Instead, actual Zone Out='%v'\n", expectedZone, actualZone)
		return
	}

	mt.mainPrintHdr("SUCCESS" , "!!!")


}


func (mt mainTest) mainTest052() {

	ePrefix := "mainTest052()"

	mt.mainPrintHdr(ePrefix , "=")

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := dt.TimeZoneDto{}.New(t1, dt.TZones.US.Pacific(), fmtstr)

	if err != nil {
		fmt.Printf("Error returned from TimeZoneDto{}.NewStartEndTimes(t1, TzUsPacific ).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tzu1.IsValid()

	if err != nil {
		fmt.Printf("'tzu1' is INVALID!\n" +
			"Error='%v'\n", err.Error())
		return
	}

	fmt.Printf("\nPassed IsValid() Check #1!\n")

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t12Dur := t2.Sub(t1)

	tdurDto, err := dt.TimeDurationDto{}.NewStartEndTimesCalcTz(
		t1,
		t2,
		dt.TDurCalcType(0).StdYearMth(),
		dt.TZones.US.Central(),
		fmtstr)

	if err != nil {
		fmt.Printf("Error returned by TimeDurationDto{}.NewStartEndTimesCalcTz()\n"+
			"Error='%v' ", err.Error())
		return
	}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddTimeDurationDto(tdurDto)
// Error expected here!
	if err != nil {
		fmt.Printf("Error returned by tzu2.AddTimeDurationDto(tdurDto).\n"+
			"Error='%v' ", err.Error())
		return
	}

	tzu1OutStr := tzu1.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t1OutStr != tzu1OutStr {
		fmt.Printf("Error: Expected Time1 TimeIn='%v'.\n" +
			"Instead Time1 TimeIn='%v'\n",
			t1OutStr, tzu1OutStr)
		return
	}

	tzu2OutStr := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t2OutStr != tzu2OutStr {
		fmt.Printf("Error: Expected after duration tzu2TimeIn='%v'.\n" +
			"Instead, tzu2TimeIn='%v'\n",
			t2OutStr, tzu2OutStr)
		return
	}

	actualDur := tzu2.TimeIn.Sub(tzu1.TimeIn)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeIn.Sub(tzu1.TimeIn)='%v'.\n" +
			"Instead, duration='%v'\n",
			t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeOut.Sub(tzu1.TimeOut)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeOut.Sub(tzu1.TimeOut)='%v'.\n" +
			"Instead, duration='%v'\n", t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeUTC.Sub(tzu1.TimeUTC)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeUTC.Sub(tzu1.TimeUTC)='%v'.\n" +
			"Instead, duration='%v'\n", t12Dur, actualDur)
		return
	}

	actualDur = tzu2.TimeLocal.Sub(tzu1.TimeLocal)

	if t12Dur != actualDur {
		fmt.Printf("Error: Expected tzu2.TimeLocal.Sub(tzu1.TimeLocal)='%v'.\n" +
			"Instead, duration='%v'\n", t12Dur, actualDur)
		return
	}

	actualTimeOutLoc := tzu1.TimeOut.GetOriginalTzName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu1.TimeOutLoc='%v'.\n" +
			"Instead, tzu1.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	actualTimeOutLoc = tzu2.TimeOut.GetOriginalTzName()

	if dt.TZones.US.Pacific() != actualTimeOutLoc {
		fmt.Printf("Error: Expected tzu2.TimeOutLoc.String()='%v'.\n" +
			"Instead, tzu2.TimeOutLoc='%v'.\n",
			dt.TZones.US.Pacific(), actualTimeOutLoc)
		return
	}

	mt.mainPrintHdr("SUCCESS" , "!")

}

func (mt mainTest) mainTest051() {
	ePrefix := "mainTest051()"
	lineLen := 65
	mt.mainPrintHdr(ePrefix , "-")

	timeZone := dt.TZones.Local()
	// t2PdtStr :=  "06/20/2019 09:58:32.000000000 -0700 PDT"
	utcLoc, err := time.LoadLocation(timeZone)

	if err != nil {
		fmt.Printf("Error returned by time.LoadLocation(dt.TZones.UTC())\n" +
			"Error='%v'\n", err.Error())
	}

	txUtc := time.Date(
		2019,
		time.Month(12),
		15,
		11,
		0,
		0,
		0,
		utcLoc)

	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"
	timeStr := txUtc.Format(fmtStr)
	fmt.Println()
	fmt.Println("#1   timeStr: ", timeStr)
	fmt.Println("#1 time zone: ", timeZone)
	fmt.Println()

		t1, err := time.Parse(fmtStr, timeStr)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.Parse(fmtStr, timeStr)\n" +
			"timeStr='%v'\nError='%v'\n", timeStr, err.Error())
		return
	}

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Initial Date Time"},
		lineLen,
		fmtStr)

	_, err = time.LoadLocation(t1.Location().String())

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by time.LoadLocation(t1.Location().String())\n" +
			"t1.Location().String()='%v'\n" +
			"Error='%v'\n", t1.Location().String(), err.Error())
	}

	tzMech := dt.TimeZoneMechanics{}

	var tzSpec dt.TimeZoneSpecification

	tzSpec,
	err = tzMech.GetConvertibleTimeZoneFromDateTime(
		t1,
		dt.TzConvertType.Absolute(),
		ePrefix,
		ePrefix)

	if err != nil {
		fmt.Printf(ePrefix +
			"\nError returned by tzMech.GetConvertibleTimeZoneFromDateTime(t1, ePrefix)\n" +
			"t1='%v'\n" +
			"Error='%v'\n", t1.Format(fmtStr), err.Error())
		return
	}

	fmt.Println()
	fmt.Printf("        ianaTimeZoneName: %v\n", tzSpec.GetLocationName())
	fmt.Printf("         ianaLocationPtr: %v\n", tzSpec.GetLocationPointer().String())
	fmt.Printf("                 tzClass: %v\n\n", tzSpec.GetTimeZoneClass().String())

	fmt.Printf(ePrefix +
		"\nSuccess!\n" +
		"time zone pointer name = '%v'\n",t1.Location().String())

}

func (mt mainTest) mainTest050() {
/*
	 =========================================
	      Time Zone Fields From time.Time
	 =========================================
	    t1 Initial Date Time:  12/30/2019 09:00:00.000000000 +0700 +07
	               Zone Name:  +07
	     Abbreviation Lookup:  +07+0700
	               Zone Sign:  +
	       Zone Offset Hours:  7
	     Zone Offset Minutes:  0
	     Zone Offset Seconds:  0
	    Total Offset Seconds:  25200
	              UTC Offset:  UTC+0700
	        Location Pointer:  Asia/Ho_Chi_Minh
	           Location Name:  Asia/Ho_Chi_Minh
	 =========================================
 */
	ePrefix := "mainTest050()"
	lineLen := 65
	mt.mainPrintHdr(ePrefix , "-")

	// "Asia/Ho_Chi_Minh"
// timeZoneName := "Asia/Vladivostok"
timeZoneName := "Asia/Ho_Chi_Minh"

tzLocPtr, err := time.LoadLocation(timeZoneName)

if err != nil {
	fmt.Printf(ePrefix +
		"Error returned by time.LoadLocation(timeZoneName)\n" +
		"timeZoneName='%v'\n" +
		"Error='%v'\n", timeZoneName, err.Error())

	return

}

	t1 := time.Date(
		2019,
		time.Month(12),
		30,
		9,
		0,
		0,
		0,
		tzLocPtr)

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	ex.PrintOutDateTimeTimeZoneFields(
		t1,
		[]string{"t1 Initial Date Time"},
		lineLen,
		fmtStr)

}

// Prints Text Title Lines to the Console
func (mt mainTest) mainPrintHdr(textToPrint string, repeatStr string) {


	lenTextToPrint := len(textToPrint)

	lenExtra := 5

	lenBar := (lenExtra * 2) + lenTextToPrint

	blankMargin := strings.Repeat(" ", lenExtra)

	title := blankMargin + textToPrint + blankMargin
	bar := strings.Repeat(repeatStr, lenBar)

	fmt.Println(bar)
	fmt.Println(title)
	fmt.Println(bar)
	fmt.Println()

}
