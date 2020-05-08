package datetime

import (
	"testing"
	"time"
)

func TestDurationTriad_SetStartEndTimes_01(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	dur := DurationTriad{}

	err := dur.SetStartEndTimes(
		t1,
		t2,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartEndTimesTz(t1, t2, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.BaseTime.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v",
			tOutDur, dur.BaseTime.timeDuration)

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

	if t1Local.Location().String() != dur.LocalTime.startDateTimeTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.startDateTimeTz.GetOriginalTzName())
	}

	if !t1Local.Equal(dur.LocalTime.startDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected Local Start Time ='%v'.\n" +
			"Actual Local Start Time ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.startDateTimeTz.GetOriginalTzName())
	}

	t2Local := dur.LocalTime.startDateTimeTz.GetDateTimeValue().Add(dur.LocalTime.timeDuration)

	if !t2Local.Equal(dur.LocalTime.endDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.endDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.UTC()). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.startDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.startDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.startDateTimeTz.GetDateTimeValue().Add(dur.UTCTime.timeDuration)

	if !t2UTC.Equal(dur.UTCTime.endDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.endDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

}

func TestDurationTriad_SetStartEndTimes_02(t *testing.T) {

	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := DurationTriad{}

	err := dur.SetStartEndTimes(
		t1,
		t2,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

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

	if t1Local.Location().String() != dur.LocalTime.startDateTimeTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.startDateTimeTz.GetOriginalTzName())
	}

	if !t1Local.Equal(dur.LocalTime.startDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected Local Start Time ='%v'.\n" +
			"Actual Local Start Time ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.startDateTimeTz.GetOriginalTzName())
	}

	t2Local := dur.LocalTime.startDateTimeTz.GetDateTimeValue().Add(dur.LocalTime.timeDuration)

	if !t2Local.Equal(dur.LocalTime.endDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.endDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.UTC()). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.startDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.startDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.startDateTimeTz.GetDateTimeValue().Add(dur.UTCTime.timeDuration)

	if !t2UTC.Equal(dur.UTCTime.endDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.endDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

}

func TestDurationTriad_SetStartEndTimes_03(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := DurationTriad{}

	err := dur.SetStartEndTimes(
		t1,
		t2,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartEndTimes(t1, t2, TZones.US.Central(), "+
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

	err := dur.SetStartEndTimes(
		t1,
		t1,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartEndTimes(t1, t1, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
	}

	outStr := dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected := "0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestTimeDuration_SetDefaultStartEndTimesTz_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	var datTz1 DateTzDto

	var err error

	datTz1, err = DateTzDto{}.NewDateTime(t1, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1). "+""+
			"Error='%v' ", err.Error())
		return
	}

	dur := DurationTriad{}

	err = dur.SetDefaultStartEndTimesTz(
		datTz1,
		datTz1)

	if err != nil {
		t.Errorf("Error returned by dur.SetDefaultStartEndTimesTz(...).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	outStr := dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected := "0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestDurationTriad_SetDefaultStartEndTimes_01(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	dur := DurationTriad{}

	err := dur.SetDefaultStartEndTimes(
		t1,
		t2)

	if err != nil {
		t.Errorf("Error returned by dur.SetDefaultStartEndTimes(t1, t2, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.BaseTime.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v",
			tOutDur, dur.BaseTime.timeDuration)

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

	if t1Local.Location().String() != dur.LocalTime.startDateTimeTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.startDateTimeTz.GetOriginalTzName())
	}

	if !t1Local.Equal(dur.LocalTime.startDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected Local Start Time ='%v'.\n" +
			"Actual Local Start Time ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.startDateTimeTz.GetOriginalTzName())
	}

	t2Local := dur.LocalTime.startDateTimeTz.GetDateTimeValue().Add(dur.LocalTime.timeDuration)

	if !t2Local.Equal(dur.LocalTime.endDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.endDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.UTC()). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.startDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.startDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.startDateTimeTz.GetDateTimeValue().Add(dur.UTCTime.timeDuration)

	if !t2UTC.Equal(dur.UTCTime.endDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.endDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

}

func TestDurationTriad_SetEndTimeMinusTimeDto_001(t *testing.T) {
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

	err := dur.SetEndTimeMinusTimeDto(
		t2,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetEndTimeMinusTimeDto(...). "+
			"Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestDurationTriad_SetEndTimeMinusTimeDto_002(t *testing.T) {
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

	err := dur.SetEndTimeMinusTimeDto(
		t2,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetEndTimeMinusTimeDto(t2, timeDto, "+
			"TZones.US.Central(), FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestDurationTriad_SetDefaultEndTimeMinusTimeDto_001(t *testing.T) {
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

	err := dur.SetDefaultEndTimeMinusTimeDto(
		t2,
		timeDto,
		TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by dur.SetDefaultEndTimeMinusTimeDto(...). "+
			"Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestDurationTriad_SetDefaultEndTimeTzMinusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var dTzDto DateTzDto

	var err error

	dTzDto, err = DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}." +
			"NewDateTime(t2, fmtstr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	dur := DurationTriad{}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err = dur.SetDefaultEndTimeTzMinusTimeDto(
		dTzDto,
		timeDto)

	if err != nil {
		t.Errorf("Error returned by dur.SetEndTimeMinusTimeDto(...). "+
			"Error='%v'", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestDurationTriad_SetStartTimeTzMinusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)


	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	dur := DurationTriad{}
	var err error
	var dTz2 DateTzDto

	dTz2, err = DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t2, fmtstr). "+
			"Error='%v'", err.Error())
	}

	err = dur.SetEndTimeTzMinusTimeDto(
		dTz2,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetEndTimeTzMinusTimeDto(...). "+
			"Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

}

func TestDurationTriad_SetStartTimeDuration_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	dur := DurationTriad{}

	err := dur.SetStartTimeDuration(
		t1,
		t12Dur,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimeDurationTz(...). "+
			"Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}

func TestDurationTriad_SetStartTimeDuration_02(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	dur := DurationTriad{}

	err := dur.SetStartTimeDuration(
		t1,
		t12Dur,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimeDuration(t1, t12Dur, "+
			"TZones.US.Central(), FmtDateTimeYrMDayFmtStr). Error='%v' ", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}

func TestDurationTriad_SetStartTimeTzDuration_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	dur := DurationTriad{}

	var dTz1 DateTzDto
	var err error

	dTz1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr). "+
			"Error='%v'", err.Error())
	}

	err = dur.SetStartTimeTzDuration(
		dTz1,
		t12Dur,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimeDurationTz(...). "+
			"Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
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
		t1,
		timeDto,
		TDurCalc.StdYearMth(),
		t1.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimePlusTimeDto(t1, timeDto, "+
			"FmtDateTimeYrMDayFmtStr) Error='%v' ", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}
}

func TestDurationTriad_SetStartTimePlusTimeDto_02(t *testing.T) {
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
		t1,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimePlusTimeDtoTz(...). "+
			"Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}
}

func TestDurationTriad_SetStartTimePlusTimeDto_03(t *testing.T) {
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
		t1,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimePlusTimeDto(t1, timeDto, "+
			"TZones.US.Central(), FmtDateTimeYrMDayFmtStr). Error='%v' ", err.Error())
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}

func TestDurationTriad_SetStartTimeTzPlusTimeDto_01(t *testing.T) {
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

	var err error
	var dTz1 DateTzDto

	dTz1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = dur.SetStartTimeTzPlusTimeDto(
		dTz1,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartTimeTzPlusTimeDto(dTz1, timeDto)\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, dur.BaseTime.timeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

}
