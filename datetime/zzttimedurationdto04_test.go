package datetime

import (
	"testing"
	"time"
)

func TestTimeDurationDto_SetStartEndTimes_01(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tDto := TimeDurationDto{}

	err := tDto.SetStartEndTimes(
		t1,
		t2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimes(...). "+""+
			"Error='%v' ", err.Error())
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
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v",
			tOutDur, tDto.timeDuration)

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

	err := tDto.SetStartEndTimes(
		t1,
		t2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimes(...). "+""+
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

	err := tDto.SetStartEndTimes(
		t2,
		t1,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimes(...). "+
			"Error='%v'", err.Error())
	}

	outStr := tDto.GetYearsMthsWeeksTimeStr()

	expected := "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 864-Milliseconds 197-Microseconds 532-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, outStr)
	}

}

func TestTimeDurationDto_SetStartEndTimes_04(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tDto := TimeDurationDto{}

	err := tDto.SetStartEndTimes(
		t1,
		t2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		"")

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimes(...). "+""+
			"Error='%v' ", err.Error())
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
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v",
			tOutDur, tDto.timeDuration)

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

func TestTimeDurationDto_SetStartEndTimes_05(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto := TimeDurationDto{}

	err := tDto.SetStartEndTimes(
		t1,
		t2,
		TDurCalc.None(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetStartEndTimes(...)\n"+
			"because input parameter 'TDurCalc.None()' is invalid!\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeDurationDto_SetStartEndTimes_06(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto := TimeDurationDto{}

	err := tDto.SetStartEndTimes(
		t1,
		t2,
		TDurCalc.StdYearMth(),
		"Invalid time zone location name",
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetStartEndTimes(...)\n"+
			"because input parameter time zone location name is invalid!\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeDurationDto_SetStartEndTimes_07(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	tDto := TimeDurationDto{}

	err := tDto.SetStartEndTimes(
		t1,
		t2,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.None(),
		FmtDateTimeYrMDayFmtStr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetStartEndTimes(...)\n"+
			"because input parameter 'TCalcMode.None()' is invalid!\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeDurationDto_SetDefaultStartEndTimes_01(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tDto := TimeDurationDto{}

	err := tDto.SetDefaultStartEndTimes(
		t1,
		t2)

	if err != nil {
		t.Errorf("Error returned by tDto.SetDefaultStartEndTimes(...). "+""+
			"Error='%v' ", err.Error())
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
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v",
			tOutDur, tDto.timeDuration)

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

func TestTimeDurationDto_SetStartEndTimesTz_01(t *testing.T) {
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

	err = dur.SetStartEndTimesTz(
		datTz1,
		datTz1,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by dur.SetStartEndTimesTz(...). "+""+
			"Error='%v' ", err.Error())
		return
	}

	outStr := dur.BaseTime.GetYearsMthsWeeksTimeStr()

	expected := "0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v",
			expected, outStr)
	}

}


func TestTimeDurationDto_SetStartEndTimesTz_02(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tDto := TimeDurationDto{}

	var err error

	var dTz1, dTz2 DateTzDto

	dTz1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	dTz2, err = DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t2, fmtstr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetStartEndTimesTz(
		dTz1,
		dTz2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartEndTimes(...).\n"+
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
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v",
			tOutDur, tDto.timeDuration)

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

func TestTimeDurationDto_SetDefaultStartEndTimesTz_01(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tDto := TimeDurationDto{}

	var err error

	var dTz1, dTz2 DateTzDto

	dTz1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	dTz2, err = DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t2, fmtstr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetDefaultStartEndTimesTz(
		dTz1,
		dTz2)

	if err != nil {
		t.Errorf("Error returned by tDto.SetDefaultStartEndTimesTz(...).\n"+
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
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v",
			tOutDur, tDto.timeDuration)

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

func TestTimeDurationDto_SetEndTimeMinusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	var err error
	var africaCairoPtr *time.Location

	africaCairoPtr, err = time.LoadLocation(TZones.Africa.Cairo())

	if err != nil {
		t.Errorf("Error returned by time." +
			"LoadLocation(TZones.Africa.Cairo()).\n" +
			"Error='%v'\n", err.Error())
		return
	}


	t01 := time.Date(
		2010,
		time.Month(1),
		15,
		13,
		0,
		0,
		0,
		africaCairoPtr)


	t02 := time.Date(
		2012,
		time.Month(3),
		9,
		2,
		0,
		0,
		0,
		africaCairoPtr)

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var tDto TimeDurationDto

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		t01,
		t02)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}." +
			"NewDefaultStartEndTimes(t01, t02).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetEndTimeMinusTimeDto(
		t2,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		fmtstr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetEndTimeMinusTimeDto().\n" +
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

func TestTimeDurationDto_SetEndTimeMinusTimeDto_002(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	var err error
	var africaCairoPtr *time.Location

	africaCairoPtr, err = time.LoadLocation(TZones.Africa.Cairo())

	if err != nil {
		t.Errorf("Error returned by time." +
			"LoadLocation(TZones.Africa.Cairo()).\n" +
			"Error='%v'\n", err.Error())
		return
	}


	t01 := time.Date(
		2010,
		time.Month(1),
		15,
		13,
		0,
		0,
		0,
		africaCairoPtr)


	t02 := time.Date(
		2012,
		time.Month(3),
		9,
		2,
		0,
		0,
		0,
		africaCairoPtr)

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var tDto TimeDurationDto

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		t01,
		t02)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}." +
			"NewDefaultStartEndTimes(t01, t02).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetEndTimeMinusTimeDto(
		t2,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		"")

	if err != nil {
		t.Errorf("Error returned by tDto.SetEndTimeMinusTimeDto().\n" +
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

func TestTimeDurationDto_SetEndTimeMinusTimeDto_003(t *testing.T) {
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	var err error
	var africaCairoPtr *time.Location

	africaCairoPtr, err = time.LoadLocation(TZones.Africa.Cairo())

	if err != nil {
		t.Errorf("Error returned by time." +
			"LoadLocation(TZones.Africa.Cairo()).\n" +
			"Error='%v'\n", err.Error())
		return
	}


	t01 := time.Date(
		2010,
		time.Month(1),
		15,
		13,
		0,
		0,
		0,
		africaCairoPtr)


	t02 := time.Date(
		2012,
		time.Month(3),
		9,
		2,
		0,
		0,
		0,
		africaCairoPtr)

	t2, _ := time.Parse(fmtstr, t2str)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var tDto TimeDurationDto

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		t01,
		t02)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}." +
			"NewDefaultStartEndTimes(t01, t02).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetEndTimeMinusTimeDto(
		t2,
		timeDto,
		TDurCalc.None(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		fmtstr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetEndTimeMinusTimeDto()\n" +
			"because input parameter 'TDurCalc.None()' is invalid!\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestTimeDurationDto_SetEndTimeMinusTimeDto_004(t *testing.T) {
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	var err error
	var africaCairoPtr *time.Location

	africaCairoPtr, err = time.LoadLocation(TZones.Africa.Cairo())

	if err != nil {
		t.Errorf("Error returned by time." +
			"LoadLocation(TZones.Africa.Cairo()).\n" +
			"Error='%v'\n", err.Error())
		return
	}


	t01 := time.Date(
		2010,
		time.Month(1),
		15,
		13,
		0,
		0,
		0,
		africaCairoPtr)


	t02 := time.Date(
		2012,
		time.Month(3),
		9,
		2,
		0,
		0,
		0,
		africaCairoPtr)

	t2, _ := time.Parse(fmtstr, t2str)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var tDto TimeDurationDto

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		t01,
		t02)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}." +
			"NewDefaultStartEndTimes(t01, t02).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetEndTimeMinusTimeDto(
		t2,
		timeDto,
		TDurCalc.StdYearMth(),
		"Invalid Time Zone Location Name",
		TCalcMode.LocalTimeZone(),
		fmtstr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetEndTimeMinusTimeDto()\n" +
			"because input parameter 'Time Zone Location Name' is invalid!\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestTimeDurationDto_SetEndTimeMinusTimeDto_005(t *testing.T) {
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	var err error
	var africaCairoPtr *time.Location

	africaCairoPtr, err = time.LoadLocation(TZones.Africa.Cairo())

	if err != nil {
		t.Errorf("Error returned by time." +
			"LoadLocation(TZones.Africa.Cairo()).\n" +
			"Error='%v'\n", err.Error())
		return
	}


	t01 := time.Date(
		2010,
		time.Month(1),
		15,
		13,
		0,
		0,
		0,
		africaCairoPtr)


	t02 := time.Date(
		2012,
		time.Month(3),
		9,
		2,
		0,
		0,
		0,
		africaCairoPtr)

	t2, _ := time.Parse(fmtstr, t2str)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var tDto TimeDurationDto

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		t01,
		t02)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}." +
			"NewDefaultStartEndTimes(t01, t02).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetEndTimeMinusTimeDto(
		t2,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.None(),
		fmtstr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetEndTimeMinusTimeDto()\n" +
			"because input parameter 'TCalcMode.None()' is invalid!\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestTimeDurationDto_SetDefaultEndTimeMinusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	var err error
	var africaCairoPtr *time.Location

	africaCairoPtr, err = time.LoadLocation(TZones.Africa.Cairo())

	if err != nil {
		t.Errorf("Error returned by time." +
			"LoadLocation(TZones.Africa.Cairo()).\n" +
			"Error='%v'\n", err.Error())
		return
	}


	t01 := time.Date(
		2010,
		time.Month(1),
		15,
		13,
		0,
		0,
		0,
		africaCairoPtr)


	t02 := time.Date(
		2012,
		time.Month(3),
		9,
		2,
		0,
		0,
		0,
		africaCairoPtr)

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var tDto TimeDurationDto

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		t01,
		t02)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}." +
			"NewDefaultStartEndTimes(t01, t02).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetDefaultEndTimeMinusTimeDto(
		t2,
		timeDto)

	if err != nil {
		t.Errorf("Error returned by tDto.SetDefaultEndTimeMinusTimeDto().\n" +
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

func TestTimeDurationDto_SetStartTimeDuration_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2007,
		5,
		7,
		14,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2008,
		9,
		7,
		6,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetStartTimeDuration(
		t1,
		t12Dur,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartTimeDuration(t1, t12Dur).\n"+
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

func TestTimeDurationDto_SetStartTimeDuration_002(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2007,
		5,
		7,
		14,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2008,
		9,
		7,
		6,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetStartTimeDuration(
		t1,
		t12Dur,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		"")

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartTimeDuration(t1, t12Dur).\n"+
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

func TestTimeDurationDto_SetStartTimeDuration_003(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t2, _ := time.Parse(fmtstr, t2str)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2007,
		5,
		7,
		14,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2008,
		9,
		7,
		6,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetStartTimeDuration(
		t1,
		t12Dur,
		TDurCalc.None(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetStartTimeDuration(t1, t12Dur)\n" +
			"because input parameter 'TDurCalc.None()' is invalid.\n"+
			"However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeDurationDto_SetStartTimeDuration_004(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t2, _ := time.Parse(fmtstr, t2str)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2007,
		5,
		7,
		14,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2008,
		9,
		7,
		6,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetStartTimeDuration(
		t1,
		t12Dur,
		TDurCalc.StdYearMth(),
		"Invalid Time Zone Location Name",
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetStartTimeDuration(t1, t12Dur)\n" +
			"because input parameter 'Time Zone Location Name' is invalid.\n"+
			"However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeDurationDto_SetStartTimeDuration_005(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t2, _ := time.Parse(fmtstr, t2str)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2007,
		5,
		7,
		14,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2008,
		9,
		7,
		6,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetStartTimeDuration(
		t1,
		t12Dur,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.None(),
		FmtDateTimeYrMDayFmtStr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetStartTimeDuration(t1, t12Dur)\n" +
			"because input parameter 'TCalcMode.None()' is invalid.\n"+
			"However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeDurationDto_SetDefaultStartTimeDuration_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2007,
		5,
		7,
		14,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2008,
		9,
		7,
		6,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetDefaultStartTimeDuration(
		t1,
		t12Dur)

	if err != nil {
		t.Errorf("Error returned by tDto.SetDefaultStartTimeDuration(t1, t12Dur).\n"+
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

func TestTimeDurationDto_SetStartTimeTzDuration_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto
	var dTz1 DateTzDto

	tX1 := time.Date(
		2007,
		5,
		7,
		14,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2008,
		9,
		7,
		6,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	dTz1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetStartTimeTzDuration(
		dTz1,
		t12Dur,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartTimeTzDuration(dTz1, t12Dur).\n"+
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

func TestTimeDurationDto_SetDefaultStartTimeTzDuration_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto
	var dTz1 DateTzDto

	tX1 := time.Date(
		2007,
		5,
		7,
		14,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2008,
		9,
		7,
		6,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(t1, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	dTz1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetDefaultStartTimeTzDuration(
		dTz1,
		t12Dur)

	if err != nil {
		t.Errorf("Error returned by tDto.SetDefaultStartTimeTzDuration(dTz1, t12Dur).\n"+
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

func TestTimeDurationDto_SetStartTimePlusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2005,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2007,
		5,
		3,
		18,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		tX1,
		tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err = tDto.SetStartTimePlusTimeDto(
		t1,
		timeDto,
		TDurCalc.StdYearMth(),
		t1.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartTimePlusTimeDto(t1, timeDto).\n" +
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

func TestTimeDurationDto_SetStartTimePlusTimeDto_002(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2005,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2007,
		5,
		3,
		18,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		tX1,
		tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err = tDto.SetStartTimePlusTimeDto(
		t1,
		timeDto,
		TDurCalc.StdYearMth(),
		t1.Location().String(),
		TCalcMode.LocalTimeZone(),
		"")

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartTimePlusTimeDto(t1, timeDto).\n" +
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

func TestTimeDurationDto_SetStartTimePlusTimeDto_003(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2005,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2007,
		5,
		3,
		18,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		tX1,
		tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err = tDto.SetStartTimePlusTimeDto(
		t1,
		timeDto,
		TDurCalc.None(),
		t1.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetStartTimePlusTimeDto(t1, timeDto)\n" +
			"because input parameter 'TDurCalc.None()' is invalid.\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestTimeDurationDto_SetStartTimePlusTimeDto_004(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2005,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2007,
		5,
		3,
		18,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		tX1,
		tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err = tDto.SetStartTimePlusTimeDto(
		t1,
		timeDto,
		TDurCalc.StdYearMth(),
		"Invalid Time Zone Location Name",
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetStartTimePlusTimeDto(t1, timeDto)\n" +
			"because input parameter 'Time Zone Location Name' is invalid.\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeDurationDto_SetStartTimePlusTimeDto_005(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2005,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2007,
		5,
		3,
		18,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		tX1,
		tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err = tDto.SetStartTimePlusTimeDto(
		t1,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.None(),
		FmtDateTimeYrMDayFmtStr)

	if err == nil {
		t.Error("Expected an error return from tDto.SetStartTimePlusTimeDto(t1, timeDto)\n" +
			"because input parameter 'TCalcMode.None()' is invalid.\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}

}

func TestTimeDurationDto_SetDefaultStartTimePlusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2005,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2007,
		5,
		3,
		18,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		tX1,
		tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	err = tDto.SetDefaultStartTimePlusTimeDto(
		t1,
		timeDto)

	if err != nil {
		t.Errorf("Error returned by tDto.SetDefaultStartTimePlusTimeDto(t1, timeDto).\n" +
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

func TestTimeDurationDto_SetStartTimeTzPlusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2005,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2007,
		5,
		3,
		18,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		tX1,
		tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var dTz1 DateTzDto

	dTz1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetStartTimeTzPlusTimeDto(
		dTz1,
		timeDto,
		TDurCalc.StdYearMth(),
		t1.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by tDto.SetStartTimeTzPlusTimeDto(t1, timeDto).\n" +
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

func TestTimeDurationDto_SetDefaultStartTimeTzPlusTimeDto_001(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	var err error
	var tDto TimeDurationDto

	tX1 := time.Date(
		2005,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2007,
		5,
		3,
		18,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		tX1,
		tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewDefaultStartEndTimes(tX1, tX2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var dTz1 DateTzDto

	dTz1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	err = tDto.SetDefaultStartTimeTzPlusTimeDto(
		dTz1,
		timeDto)

	if err != nil {
		t.Errorf("Error returned by tDto.SetDefaultStartTimeTzPlusTimeDto(t1, timeDto).\n" +
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
