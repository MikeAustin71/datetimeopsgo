package datetime

import (
	"testing"
	"time"
)

func TestTimeDurationDto_DaylightSavings_01(t *testing.T) {

	// This test verifies duration over a daylight savings time threshold.

	locUSCentral, _ := time.LoadLocation(TzIanaUsCentral)

	t1USCentral := time.Date(2018, time.Month(3),10,18,0,0,0, locUSCentral)

	hoursDur := int64(24) * HourNanoSeconds

	t1Dur, err := TimeDurationDto{}.NewStartTimeDurationTzCalc(t1USCentral, time.Duration(hoursDur),
										TzIanaUsCentral,	TDurCalcTypeSTDYEARMTH, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dt.TimeDurationDto{}.New(t1USCentral, t2USCentral, fmtStr). " +
			"Error='%v'\n", err.Error())
	}

	outStr := t1Dur.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr)

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

func TestTimeDurationDto_GetYearMthDaysTimeAbbrv(t *testing.T) {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto := TimeDurationDto{}

	tDto.SetStartEndTimesTzCalc(t2, t1, TDurCalcTypeSTDYEARMTH, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := tDto.GetYearMthDaysTimeAbbrvStr()


	if expected != dOut {
		t.Errorf("Expected: %v. Error - got %v", expected, dOut)
	}
}

// TODO - Fix these tests
func TestTimeDurationDto_GetYearsMthsWeeksTimeAbbrv(t *testing.T) {

	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:59:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := DurationTriad{}

	du.SetStartEndTimes(t2, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	expected := "0-Hours 1-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearsMthsWeeksTimeAbbrvStr()

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

	dur, err := DurationTriad{}.NewStartTimeDuration(t1, t12Dur, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimeDuration(t1, t12Dur). Error='%v'", err.Error())
	}

	// dur.SetStartTimeDuration(t1, t12Dur)

	if t1OutStr != dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.BaseTime.TimeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

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

	dur, err := DurationTriad{}.NewStartEndTimes(t1, t2, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimes(t1, t2). Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.BaseTime.TimeDuration {
		t.Errorf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v", tOutDur, dur.BaseTime.TimeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()


	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, _ = dur.BaseTime.GetCumDaysTimeStr()

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). " +
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumWeeksDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). " +
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

	dur, err := DurationTriad{}.NewEndTimeMinusTimeDto(t2, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDto(t2, timeDto). Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr))
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
	dur, err := DurationTriad{}.NewStartTimePlusTime(t1, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimePlusTime(t1, timeDto). Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr))
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

func TestTimeDurationDto_SetStartEndTimes(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	dur := DurationTriad{}

	dur.SetStartEndTimes(t1, t2, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	if t1OutStr != dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.BaseTime.TimeDuration {
		t.Errorf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v",
			tOutDur, dur.BaseTime.TimeDuration)

	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr = dur.BaseTime.GetDefaultDurationStr()

	expected = "28082h4m2s"

	if expected != outStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err := dur.BaseTime.GetCumDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by dur.BaseTime.GetCumDaysTimeStr(). " +
			"Error='%v'", err.Error())
	}

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). " +
			"Error='%v'", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, outStr)
	}

	outStr = dur.BaseTime.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumWeeksDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by dur.BaseTime.GetCumWeeksDaysTimeStr(). " +
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

	dur := DurationTriad{}

	dur.SetStartEndTimes(t1, t2, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	outStr := dur.BaseTime.GetYrMthWkDayHrMinSecNanosecsStr()

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

	dur := DurationTriad{}

	dur.SetStartEndTimes(t2, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	outStr := dur.BaseTime.GetYearsMthsWeeksTimeStr()

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

	dur.SetStartEndTimes(t1, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

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
	dur.SetStartTimePlusTime(t1, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if t1OutStr != dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr))
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

	dur.SetEndTimeMinusTimeDto(t2, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if t1OutStr != dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr))
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

	dur.SetStartTimeDuration(t1, t12Dur, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if t1OutStr != dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.DateTime.Format(fmtstr))
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


