package datetime

import (
	"testing"
	"time"
)

func TestDurationTriad_SetStartEndTimes(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	dur := DurationTriad{}

	err := dur.SetStartEndTimesTz(t1, t2, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartEndTimesTz(t1, t2, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
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

	outStr, err = dur.BaseTime.GetCumDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by dur.BaseTime.GetCumDaysTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "1170-Days 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v",
			expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
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
		t.Errorf("Error returned by dur.BaseTime.GetCumWeeksDaysTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v", expected, outStr)
	}

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.GetOriginalTzName())
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.GetDateTimeValue()) {
		t.Errorf("Expected Local Start Time ='%v'.\n" +
			"Actual Local Start Time ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.GetOriginalTzName())
	}

	t2Local := dur.LocalTime.StartTimeDateTz.GetDateTimeValue().Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.GetDateTimeValue()) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.UTC()). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.GetDateTimeValue().Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

}

func TestDurationTriad_SetStartEndTimes_02(t *testing.T) {

	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := DurationTriad{}

	err := dur.SetStartEndTimesTz(t1, t2, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartEndTimesTz(t1, t2, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
	}

	outStr := dur.BaseTime.GetYrMthWkDayHrMinSecNanosecsStr()

	expected := "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 864197532-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v",
			expected, outStr)
	}

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.GetOriginalTzName())
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.GetDateTimeValue()) {
		t.Errorf("Expected Local Start Time ='%v'.\n" +
			"Actual Local Start Time ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.GetOriginalTzName())
	}

	t2Local := dur.LocalTime.StartTimeDateTz.GetDateTimeValue().Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.GetDateTimeValue()) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.UTC()). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.GetDateTimeValue().Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

}

func TestDurationTriad_SetStartEndTimes_03(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := DurationTriad{}

	err := dur.SetStartEndTimesTz(t2, t1, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartEndTimesTz(t2, t1, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
	}

	outStr := dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected := "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 864-Milliseconds 197-Microseconds 532-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

}

func TestDurationTriad_SetStartEndTimes_04(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	dur := DurationTriad{}

	err := dur.SetStartEndTimesTz(t1, t1, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartEndTimesTz(t1, t1, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
	}

	outStr := dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected := "0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestDurationTriad_SetStartTimePlusTime(t *testing.T) {
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
		t.Errorf("Error returned by dur.SetStartTimePlusTimeDtoTz(t1, timeDto, "+
			"TZones.US.Central(), FmtDateTimeYrMDayFmtStr). Error='%v' ", err.Error())
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

func TestDurationTriad_SetStartTimeMinusTime(t *testing.T) {
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
		t.Errorf("Error returned by dur.SetEndTimeMinusTimeDtoTz(t2, timeDto, "+
			"TZones.US.Central(), FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
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

func TestDurationTriad_SetStartTimeDuration(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	dur := DurationTriad{}

	err := dur.SetStartTimeDurationTz(t1, t12Dur, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimeDurationTz(t1, t12Dur, "+
			"TZones.US.Central(), FmtDateTimeYrMDayFmtStr). Error='%v' ", err.Error())
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

func TestDurationTriad_SetStartTimePlusTimeDto_01(t *testing.T) {
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

	err := dur.SetStartTimePlusTimeDto(
		TCalcMode.LocalTimeZone(),
		t1,
		timeDto,
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimePlusTimeDto(t1, timeDto, "+
			"FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
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
