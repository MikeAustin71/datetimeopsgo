package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeDurationDto_DaylightSavings_01(t *testing.T) {

	// This test verifies duration over a daylight savings time threshold.

	locUSCentral, _ := time.LoadLocation(TZones.US.Central())

	t1USCentral := time.Date(2018, time.Month(3), 10, 18, 0, 0, 0, locUSCentral)

	hoursDur := int64(24) * HourNanoSeconds

	t1Dur, err := TimeDurationDto{}.NewStartTimeAddDuration(
		t1USCentral,
		time.Duration(hoursDur),
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dt.TimeDurationDto{}.NewStartEndTimes(t1USCentral, t2USCentral, fmtStr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	outStr := t1Dur.endDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr)

	expectedStr := "2018-03-11 19:00:00.000000000 -0500 CDT"

	if expectedStr != outStr {
		t.Errorf("Error: Expected outStr='%v'.  Instead, outStr='%v'. ", expectedStr, outStr)
	}

	/*

	   -- Gained an hour over Daylight savings threshold
	   Add Date Results - Cumulative Days
	               Start Date Time:  2018-03-10 18:00:00.000000000 -0600 CST
	         -- Duration = 24-Hours --
	          Actual End Date Time:  2018-03-11 19:00:00.000000000 -0500 CDT
	*/
}

/*
*******************************************************************************
 */

func TestTimeDurationDto_GetCumSecondsTimeStr_01(t *testing.T) {

	locUSCentral, err := time.LoadLocation(TZones.US.Central())

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

	tDur, err := TimeDurationDto{}.NewStartEndTimes(
		t1USCentral,
		t2USCentral,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartEndTimesCalcTz()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualStr, _ := tDur.GetCumSecondsTimeStr()

	expectedStr := "62-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expectedStr != actualStr {
		t.Errorf("Error: Expected Duration String='%v'. Actual String='%v'",
			expectedStr, actualStr)
	}

}

func TestTimeDurationDto_GetCumSecondsDto_01(t *testing.T) {

	locUSCentral, err := time.LoadLocation(TZones.US.Central())

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

	tDur, err := TimeDurationDto{}.NewStartEndTimes(
		t1USCentral,
		t2USCentral,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartEndTimesCalcTz()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	d2Dur, _ := tDur.GetCumSecondsCalcDto()

	if d2Dur.Years != 0 {
		t.Error("INVALID Years!")
	}

	if d2Dur.Months != 0 {
		t.Error("INVALID Months!")
	}

	if d2Dur.Weeks != 0 {
		t.Error("INVALID Weeks!")
	}

	if d2Dur.WeekDays != 0 {
		t.Error("INVALID Weeks!")
	}

	if d2Dur.DateDays != 0 {
		t.Error("INVALID DateDays!")
	}

	if d2Dur.Hours != 0 {
		t.Error("INVALID Hours!")
	}

	if d2Dur.Minutes != 0 {
		t.Error("INVALID Minutes!")
	}

	if d2Dur.Seconds != 62 {
		t.Error("INVALID Seconds!")
	}

	if d2Dur.Milliseconds != 0 {
		t.Error("INVALID Milliseconds!")
	}

	if d2Dur.Microseconds != 0 {
		t.Error("INVALID Microseconds!")
	}

	if d2Dur.Nanoseconds != 0 {
		t.Error("INVALID Nanoseconds!")
	}

}

func TestTimeDurationDto_GetYearMthDaysTimeAbbrvStr(t *testing.T) {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto := TimeDurationDto{}

	err :=
		tDto.SetStartEndTimes(
			t2,
			t1,
			TDurCalcType(0).StdYearMth(),
			TZones.US.Central(),
			TCalcMode.LocalTimeZone(),
			FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimes(...). "+
			"Error='%v' ", err.Error())
	}

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := tDto.GetYearMthDaysTimeAbbrvStr()

	if expected != dOut {
		t.Errorf("Expected: %v. Error - got %v", expected, dOut)
	}
}

func TestTimeDurationDto_GetYearsMthsWeeksTimeAbbrvStr(t *testing.T) {

	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:59:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto := TimeDurationDto{}

	err := tDto.SetStartEndTimes(
		t2,
		t1,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimes(...). "+""+
			"Error='%v' ", err.Error())
	}

	expected := "0-Hours 1-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := tDto.GetYearsMthsWeeksTimeAbbrvStr()

	if expected != dOut {
		t.Errorf("Expected: %v. Error - got %v", expected, dOut)
	}

}

func TestTimeDurationDto_NewStartTimeAddDate_01(t *testing.T) {

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"


	locationPtr, err := time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	startDateTime := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)

	// Local Tz End Date Time
	// 2019-08-27 19:54:30.038175584 -0500 CDT
	calcEndDateTimeLocal := time.Date(
		2019,
		8,
		27,
		19,
		54,
		30,
		38175584,
		locationPtr)

	var tDurDto TimeDurationDto

	tDurDto, err = TimeDurationDto{}.NewStartTimeAddDate(
		startDateTime,
		5,
		6,
		12,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.LocalTimeZone(),
		fmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartTimeAddDate()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedDateTime := calcEndDateTimeLocal.Format(fmtStr)

	actualDateTime := tDurDto.GetThisEndDateTimeString()

	if expectedDateTime != actualDateTime {
		t.Errorf("Error: Expected End Date Time='%v'\n" +
			"Instead, End Date Time='%v'\n",
			expectedDateTime, actualDateTime)
	}
}

func TestTimeDurationDto_NewStartTimeAddDate_02(t *testing.T) {

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	locationPtr, err := time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	startDateTime := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)

	// UTC Tz End Date Time
	// 2019-08-27 20:54:30.038175584 -0500 CDT
	calcEndDateTimeLocal := time.Date(
		2019,
		8,
		27,
		20,
		54,
		30,
		38175584,
		locationPtr)

	var tDurDto TimeDurationDto

	tDurDto, err = TimeDurationDto{}.NewStartTimeAddDate(
		startDateTime,
		5,
		6,
		12,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.UtcTimeZone(),
		fmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartTimeAddDate()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedDateTime := calcEndDateTimeLocal.Format(fmtStr)

	actualDateTime := tDurDto.GetThisEndDateTimeString()

	if expectedDateTime != actualDateTime {
		t.Errorf("Error: Expected End Date Time='%v'\n" +
			"Instead, End Date Time='%v'\n",
			expectedDateTime, actualDateTime)
	}
}

func TestTimeDurationDto_NewStartTimeAddDate_03(t *testing.T) {
/*
		Daylight Savings Time Changed To Standard Time
	 on November 2, 2014

*/
	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	locationPtr, err := time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-11-03 00:00:00.000000000 -0600 CST"
	endDateTimeLocal := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		locationPtr)

	// Local Start Date Time
	// 2014-11-02 00:00:00.000000000 -0500 CDT
	calcStartDateTimeLocal := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		locationPtr)

	// UTC Start Date Time
	// 2014-11-02 01:00:00.000000000 -0500 CDT
	calcStartDateTimeUTC := time.Date(
		2014,
		11,
		2,
		1,
		0,
		0,
		0,
		locationPtr)

	var tDurDto, tDurDto2 TimeDurationDto

	tDurDto, err = TimeDurationDto{}.NewStartTimeAddDate(
		endDateTimeLocal,
		0,
		0,
		-1,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.LocalTimeZone(),
		fmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartTimeAddDate()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedDateTimeLocal := calcStartDateTimeLocal.Format(fmtStr)

	actualDateTimeLocal := tDurDto.GetThisStartDateTimeString()

	if expectedDateTimeLocal != actualDateTimeLocal {
		t.Errorf("\nError: Expected Start Date Time Local='%v'\n" +
			"Instead, Start Date Time Local='%v'\n",
			expectedDateTimeLocal, actualDateTimeLocal)
	}

	tDurDto2, err = TimeDurationDto{}.NewStartTimeAddDate(
		endDateTimeLocal,
		0,
		0,
		-1,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.UtcTimeZone(),
		fmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartTimeAddDate()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedDateTimeUTC := calcStartDateTimeUTC.Format(fmtStr)

	actualDateTimeUTC := tDurDto2.GetThisStartDateTimeString()

	if expectedDateTimeUTC != actualDateTimeUTC {
		t.Errorf("Error: Expected Start Date Time UTC='%v'\n" +
			"Instead, Start Date Time UTC='%v'\n",
			expectedDateTimeUTC, actualDateTimeUTC)
	}

}

func TestTimeDurationDto_NewStartTimeAddDate_04(t *testing.T) {

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	locationPtr, err := time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	startDateTime := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)

	_, err = TimeDurationDto{}.NewStartTimeAddDate(
		startDateTime,
		5,
		6,
		12,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.None(),
		fmtStr)

	if err == nil {
		t.Error("Expected error returned from " +
			"TimeDurationDto{}.NewStartTimeAddDate()\n" +
			"because parameter TCalcMode.None() is invalid.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}
}

func TestTimeDurationDto_NewStartTimeAddDate_05(t *testing.T) {

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	locationPtr, err := time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	startDateTime := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)

	_, err = TimeDurationDto{}.NewStartTimeAddDate(
		startDateTime,
		5,
		6,
		12,
		TDurCalc.StdYearMth(),
		"Invalid Time Zone Location",
		TCalcMode.LocalTimeZone(),
		fmtStr)

	if err == nil {
		t.Error("Expected error returned from " +
			"TimeDurationDto{}.NewStartTimeAddDate()\n" +
			"because parameter Time Zone Location is invalid.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}
}

func TestTimeDurationDto_NewStartTimeAddDate_06(t *testing.T) {

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	locationPtr, err := time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	startDateTime := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)

	_, err = TimeDurationDto{}.NewStartTimeAddDate(
		startDateTime,
		5,
		6,
		12,
		TDurCalc.None(),
		TZones.America.Chicago(),
		TCalcMode.LocalTimeZone(),
		fmtStr)

	if err == nil {
		t.Error("Expected error returned from " +
			"TimeDurationDto{}.NewStartTimeAddDate()\n" +
			"because parameter TDurCalc.None() is invalid.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}
}

func TestTimeDurationDto_NewStartTimeAddDateTime_01(t *testing.T) {

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"


	locationPtr, err := time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	startDateTime := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)


	// Local Tz End Date Time
	// 2019-08-27 20:55:31.038175594 -0500 CDT
	calcEndDateTimeLocal := time.Date(
		2019,
		8,
		27,
		20,
		55,
		31,
		38175594,
		locationPtr)

	var tDurDto TimeDurationDto
	// 5-years, 6-months, 12-days,
	// 1-hour, 1-minute, 1-second, 0-milliseconds,
	// 0-microseconds, 10-nanoseconds

	tDurDto, err = TimeDurationDto{}.NewStartTimeAddDateTime(
		startDateTime,
		5,
		6,
		12,
		1,
		1,
		1,
		0,
		0,
		10,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.LocalTimeZone(),
		fmtStr)

	if err != nil {
		t.Errorf("Error returned by " +
			"TimeDurationDto{}.NewStartTimeAddDateTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedDateTime := calcEndDateTimeLocal.Format(fmtStr)

	actualDateTime := tDurDto.GetThisEndDateTimeString()

	if expectedDateTime != actualDateTime {
		t.Errorf("Error: Expected End Date Time='%v'\n" +
			"Instead, End Date Time='%v'\n",
			expectedDateTime, actualDateTime)
	}
}

func TestTimeDurationDto_NewStartTimeAddDateTime_02(t *testing.T) {

	fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"


	locationPtr, err := time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(dt.TZones.America.Chicago()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// "2014-02-15 19:54:30.038175584 -0600 CST"
	calcStartDateTimeLocal := time.Date(
		2014,
		2,
		15,
		19,
		54,
		30,
		38175584,
		locationPtr)


	// Local Tz End Date Time
	// 2019-08-27 20:55:31.038175594 -0500 CDT
	endDateTime := time.Date(
		2019,
		8,
		27,
		20,
		55,
		31,
		38175594,
		locationPtr)

	var tDurDto TimeDurationDto
	// -5-years, -6-months, -12-days,
	// -1-hour, -1-minute, -1-second, 0-milliseconds,
	// 0-microseconds, -10-nanoseconds

	tDurDto, err = TimeDurationDto{}.NewStartTimeAddDateTime(
		endDateTime,
		-5,
		-6,
		-12,
		-1,
		-1,
		-1,
		0,
		0,
		-10,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.LocalTimeZone(),
		fmtStr)

	if err != nil {
		t.Errorf("Error returned by " +
			"TimeDurationDto{}.NewStartTimeAddDateTime()\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedDateTime := calcStartDateTimeLocal.Format(fmtStr)

	actualDateTime := tDurDto.GetThisStartDateTimeString()

	if expectedDateTime != actualDateTime {
		t.Errorf("\nError: Expected Start Date Time='%v'\n" +
			"Instead, Start Date Time='%v'\n",
			expectedDateTime, actualDateTime)
	}
}

func TestTimeDurationDto_NewDefaultStartEndTimes_001(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tDto, err := TimeDurationDto{}.NewDefaultStartEndTimes(
		t1,
		t2)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewDefaultStartEndTimes(...).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, _ = tDto.GetCumDaysTimeStr()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, err = tDto.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = tDto.GetCumWeeksDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestTimeDurationDto_NewDefaultStartEndTimesTz_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t1Dtz, err := DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewStartEndTimes(t1, fmtstr). Error='%v'", err.Error())
	}

	t2Dtz, err := DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewStartEndTimes(t2, fmtstr). Error='%v'", err.Error())
	}

	tDto, err := TimeDurationDto{}.NewDefaultStartEndTimesTz(
		t1Dtz,
		t2Dtz)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewDefaultStartEndTimesTz(...).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, _ = tDto.GetCumDaysTimeStr()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, err = tDto.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = tDto.GetCumWeeksDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}
}

func TestTimeDurationDto_NewDefaultStartTimeDuration_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	tDto, err := TimeDurationDto{}.NewDefaultStartTimeDuration(
		t1,
		t12Dur)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartTimeDurationCalcTz(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}
}

func TestTimeDurationDto_NewDefaultStartTimeTzDuration_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var tDto TimeDurationDto
	var err error
	var dTzT1 DateTzDto

	dTzT1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	tDto, err = TimeDurationDto{}.NewDefaultStartTimeTzDuration(
		dTzT1,
		t12Dur)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartTimeDurationCalcTz(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}
}

func TestTimeDurationDto_NewDefaultStartTimePlusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	tDto, err := TimeDurationDto{}.NewDefaultStartTimePlusTimeDto(
		t1,
		timeDto)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewDefaultStartTimePlusTimeDto(t1, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}
}

func TestTimeDurationDto_NewDefaultStartTimeTzPlusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var tDto TimeDurationDto
	var err error
	var dTzT1 DateTzDto

	dTzT1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDto, err = TimeDurationDto{}.NewDefaultStartTimeTzPlusTimeDto(
		dTzT1,
		timeDto)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewDefaultStartTimePlusTimeDto(t1, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}
}

func TestTimeDurationDto_NewDefaultEndTimeMinusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	tDto, err := TimeDurationDto{}.NewDefaultEndTimeMinusTimeDto(
		t2,
		timeDto)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewDefaultEndTimeMinusTimeDto(t2, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}
}

func TestTimeDurationDto_NewDefaultEndTimeTzMinusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var err error
	var tDto TimeDurationDto
	var dTz2 DateTzDto

	dTz2, err = DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}" +
			".NewDateTime(t2, fmtstr).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDto, err = TimeDurationDto{}.NewDefaultEndTimeTzMinusTimeDto(
		dTz2,
		timeDto)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}." +
			"NewDefaultEndTimeTzMinusTimeDto(NewDefaultEndTimeTzMinusTimeDto, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}
}

func TestTimeDurationDto_NewStartTimeAddDuration_S2_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	tDto, err := TimeDurationDto{}.NewStartTimeAddDuration(
		t1,
		t12Dur,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartTimeDurationCalcTz(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}

func TestTimeDurationDto_NewStartEndTimes_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tDto, err := TimeDurationDto{}.NewStartEndTimes(
		t1,
		t2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesCalcTz(...).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, _ = tDto.GetCumDaysTimeStr()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, err = tDto.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = tDto.GetCumWeeksDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}
}

func TestTimeDurationDto_NewStartEndTimesTz_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t1Dtz, err := DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewStartEndTimes(t1, fmtstr). Error='%v'", err.Error())
	}

	t2Dtz, err := DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewStartEndTimes(t2, fmtstr). Error='%v'", err.Error())
	}

	tDto, err := TimeDurationDto{}.NewStartEndTimesTz(
		t1Dtz,
		t2Dtz,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesDateTzDtoCalcTz(...).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, _ = tDto.GetCumDaysTimeStr()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, err = tDto.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = tDto.GetCumWeeksDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestTimeDurationDto_NewStartEndTimesTz_02(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t1Dtz, err := DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewStartEndTimes(t1, fmtstr). Error='%v'", err.Error())
	}

	t2Dtz, err := DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewStartEndTimes(t2, fmtstr). Error='%v'", err.Error())
	}

	tDto, err := TimeDurationDto{}.NewStartEndTimesTz(
		t1Dtz,
		t2Dtz,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesDateTzDtoCalcTz(...).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, _ = tDto.GetCumDaysTimeStr()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, err = tDto.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = tDto.GetCumWeeksDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestTimeDurationDto_NewStartEndTimesTz_03(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t1Dtz, err := DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewStartEndTimes(t1, fmtstr). Error='%v'", err.Error())
	}

	t2Dtz, err := DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewStartEndTimes(t2, fmtstr). Error='%v'", err.Error())
	}

	tDto, err := TimeDurationDto{}.NewStartEndTimesTz(
		t1Dtz,
		t2Dtz,
		TDurCalc.StdYearMth(),
		t1Dtz.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndDateTzDto(...).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, _ = tDto.GetCumDaysTimeStr()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, err = tDto.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = tDto.GetCumWeeksDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestTimeDurationDto_NewStartTimeMinusTimeDto_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	tDto, err := TimeDurationDto{}.NewEndTimeMinusTimeDto(
		t2,
		timeDto,
		TDurCalc.StdYearMth(),
		t2.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}

func TestTimeDurationDto_NewStartTimePlusTimeDto_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	tDto, err := TimeDurationDto{}.NewStartTimePlusTimeDto(
		t1,
		timeDto,
		TDurCalc.StdYearMth(),
		t1.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimePlusTimeDtoTz(t1, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}
}

