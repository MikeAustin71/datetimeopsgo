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

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationCalcTz(t1USCentral, time.Duration(hoursDur),
		TDurCalcType(0).StdYearMth(), TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dt.TimeDurationDto{}.NewStartEndTimes(t1USCentral, t2USCentral, fmtStr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	outStr := t1Dur.EndTimeDateTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr)

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

	tDur, err := TimeDurationDto{}.NewStartEndTimesCalcTz(
		t1USCentral,
		t2USCentral,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
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

	tDur, err := TimeDurationDto{}.NewStartEndTimesCalcTz(
		t1USCentral,
		t2USCentral,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
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
		tDto.SetStartEndTimesCalcTz(
			t2,
			t1,
			TDurCalcType(0).StdYearMth(),
			TZones.US.Central(),
			FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimesCalcTz(...). "+
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

	err := tDto.SetStartEndTimesCalcTz(
		t2,
		t1,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimesCalcTz(...). "+""+
			"Error='%v' ", err.Error())
	}

	expected := "0-Hours 1-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := tDto.GetYearsMthsWeeksTimeAbbrvStr()

	if expected != dOut {
		t.Errorf("Expected: %v. Error - got %v", expected, dOut)
	}

}

func TestTimeDurationDto_NewStartTimeDuration_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	tDto, err := TimeDurationDto{}.NewStartTimeDurationCalcTz(t1, t12Dur,
		TDurCalcType(0).StdYearMth(), TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartTimeDurationCalcTz(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	// tDto.SetStartTimeDurationTz(t1, t12Dur)

	if t1OutStr != tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, tDto.TimeDuration)
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

	tDto, err := TimeDurationDto{}.NewStartEndTimesCalcTz(t1, t2, TDurCalcType(0).StdYearMth(),
		TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesCalcTz(...).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.TimeDuration {
		t.Errorf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v", tOutDur, tDto.TimeDuration)
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

func TestTimeDurationDto_NewStartEndDateTzDtoCalcTz_01(t *testing.T) {
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

	tDto, err := TimeDurationDto{}.NewStartEndTimesDateTzDtoCalcTz(t1Dtz, t2Dtz, TDurCalcType(0).StdYearMth(),
		TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesDateTzDtoCalcTz(...).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.TimeDuration {
		t.Errorf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v", tOutDur, tDto.TimeDuration)
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

func TestTimeDurationDto_NewStartEndDateTzDtoTz_01(t *testing.T) {
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

	tDto, err := TimeDurationDto{}.NewStartEndDateTzDtoTz(t1Dtz, t2Dtz,
		TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesDateTzDtoCalcTz(...).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.TimeDuration {
		t.Errorf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v", tOutDur, tDto.TimeDuration)
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

func TestTimeDurationDto_NewStartEndDateTzDto_01(t *testing.T) {
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

	tDto, err := TimeDurationDto{}.NewStartEndDateTzDto(
		t1Dtz,
		t2Dtz,
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndDateTzDto(...).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.TimeDuration {
		t.Errorf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v", tOutDur, tDto.TimeDuration)
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

func TestTimeDurationDto_NewStartTimeMinusTime_01(t *testing.T) {
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
		TCalcMode.LocalTimeZone(),
		t2,
		timeDto,
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, tDto.TimeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}

func TestTimeDurationDto_NewStartTimePlusTime_01(t *testing.T) {
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
		TCalcMode.LocalTimeZone(),
		t1,
		timeDto,
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimePlusTimeDtoTz(t1, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, tDto.TimeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestTimeDurationDto_SetStartEndTimes(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tDto := TimeDurationDto{}

	err := tDto.SetStartEndTimesCalcTz(
		t1,
		t2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimesCalcTz(...). "+""+
			"Error='%v' ", err.Error())
	}

	if t1OutStr != tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, tDto.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.TimeDuration {
		t.Errorf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v",
			tOutDur, tDto.TimeDuration)

	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = tDto.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr = tDto.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = tDto.GetCumDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by tDto.BaseTime.GetCumDaysTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
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
		t.Errorf("Error returned by tDto.BaseTime.GetCumWeeksDaysTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

}

func TestTimeDurationDto_SetStartEndTimes_02(t *testing.T) {

	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto := TimeDurationDto{}

	err := tDto.SetStartEndTimesCalcTz(
		t1,
		t2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimesCalcTz(...). "+""+
			"Error='%v' ", err.Error())
	}

	outStr := tDto.GetYrMthWkDayHrMinSecNanosecsStr()

	expected := "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 864197532-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestTimeDurationDto_SetStartEndTimes_03(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto := TimeDurationDto{}

	err := tDto.SetStartEndTimesCalcTz(
		t2,
		t1,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimesCalcTz(...). "+
			"Error='%v'", err.Error())
	}

	outStr := tDto.GetYearsMthsWeeksTimeStr()

	expected := "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 864-Milliseconds 197-Microseconds 532-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

}

func TestTimeDurationDto_SetStartEndTimes_04(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	dur := DurationTriad{}

	err := dur.SetStartEndTimesTz(t1, t1, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartEndTimesTz(...). "+""+
			"Error='%v' ", err.Error())
	}

	outStr := dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected := "0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestTimeDurationDto_SetStartTimePlusTime(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	dur := DurationTriad{}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err := dur.SetStartTimePlusTimeDtoTz(
		TCalcMode.LocalTimeZone(),
		t1,
		timeDto,
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimePlusTimeDtoTz(...). "+
			"Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.TimeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}

func TestTimeDurationDto_SetStartTimeMinusTime(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	dur := DurationTriad{}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err := dur.SetEndTimeMinusTimeDtoTz(
		TCalcMode.LocalTimeZone(),
		t2,
		timeDto,
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetEndTimeMinusTimeDtoTz(...). "+
			"Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.TimeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestTimeDurationDto_SetStartTimeDuration(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	dur := DurationTriad{}

	err := dur.SetStartTimeDurationTz(
		t1,
		t12Dur,
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimeDurationTz(...). "+
			"Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.TimeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}
