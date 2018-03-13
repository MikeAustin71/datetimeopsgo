package datetime
/*
import (
	"testing"
	"time"
)

func TestDurationUtility_GetYearMthDaysTimeAbbrv(t *testing.T) {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := DurationTriad{}

	du.SetStartEndTimes(t2, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut, err := du.BaseTime.GetYearMthDaysTimeAbbrv()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTimeAbbrv(). Error: %v", err.Error())
	}

	if expected != dOut.DisplayStr {
		t.Errorf("Expected: %v. Error - got %v", expected, dOut.DisplayStr)
	}
}

func TestDurationUtility_GetYearsMthsWeeksTimeAbbrv(t *testing.T) {

	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:59:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := DurationTriad{}

	du.SetStartEndTimes(t2, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	expected := "0-Hours 1-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut, err := du.GetYearsMthsWeeksTimeAbbrv()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTimeAbbrv(). Error: %v", err.Error())
	}

	if expected != dOut.DisplayStr {
		t.Errorf("Expected: %v. Error - got %v", expected, dOut.DisplayStr)
	}

}

func TestDurationUtility_NewStartTimeDuration_01(t *testing.T) {
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

	if t1OutStr != dur.BaseTime.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.", t1OutStr, dur.BaseTime.TimeIn.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.", t2OutStr, dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr))
	}

	if t12Dur != dur.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.TimeDuration)
	}

	dDto, err := dur.GetYearMthDaysTime()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTime. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}

func TestDurationUtility_NewStartEndTimes_01(t *testing.T) {
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

	if t1OutStr != dur.BaseTime.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ", t1OutStr, dur.BaseTime.TimeIn.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ", t1OutStr, dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.TimeDuration {
		t.Errorf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v", tOutDur, dur.TimeDuration)
	}

	dDto, err := dur.GetYearMthDaysTime()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTime. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetYearsMthsWeeksTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearsMthsWeeksTime. Error: %v", err.Error())
	}

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetDefaultDuration()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetDefaultDuration. Error: %v", err.Error())
	}

	expected = "28082h4m2s"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetDaysTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetDaysTime. Error: %v", err.Error())
	}

	expected = "1170-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetHoursTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetHoursTime. Error: %v", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetYrMthWkDayHrMinSecNanosecs()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetYrMthWkDayHrMinSecNanosecs. Error: %v", err.Error())
	}

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetWeeksDaysTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetWeeksDaysTime. Error: %v", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}
}

func TestDurationUtility_NewStartTimeMinusTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)


	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	dur, err := DurationTriad{}.NewStartTimeMinusTime(t2, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimeMinusTime(t2, timeDto). Error='%v'", err.Error())
	}

	if t1OutStr != dur.BaseTime.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.", t1OutStr, dur.BaseTime.TimeIn.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.", t2OutStr, dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr))
	}

	if t12Dur != dur.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.TimeDuration)
	}

	dDto, err := dur.GetYearMthDaysTime()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTime. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}

func TestDurationUtility_NewStartTimePlusTime_01(t *testing.T) {
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

	if t1OutStr != dur.BaseTime.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.", t1OutStr, dur.BaseTime.TimeIn.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.", t2OutStr, dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr))
	}

	if t12Dur != dur.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.TimeDuration)
	}

	dDto, err := dur.GetYearMthDaysTime()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTime. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}

func TestDurationUtility_SetStartEndTimes(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	dur := DurationTriad{}

	dur.SetStartEndTimes(t1, t2, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	if t1OutStr != dur.BaseTime.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v. Instead, got %v ", t1OutStr, dur.BaseTime.TimeIn.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ", t1OutStr, dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.TimeDuration {
		t.Errorf("Error: Expected DurationTriad.TimeDuration of %v. Instead, got %v", tOutDur, dur.TimeDuration)
	}

	dDto, err := dur.GetYearMthDaysTime()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTime. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetYearsMthsWeeksTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearsMthsWeeksTime. Error: %v", err.Error())
	}

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetDefaultDuration()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetDefaultDuration. Error: %v", err.Error())
	}

	expected = "28082h4m2s"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected Default Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetDaysTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetDaysTime. Error: %v", err.Error())
	}

	expected = "1170-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected WeekDays Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetHoursTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetHoursTime. Error: %v", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected Hours Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetYrMthWkDayHrMinSecNanosecs()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetYrMthWkDayHrMinSecNanosecs. Error: %v", err.Error())
	}

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

	dDto, err = dur.GetWeeksDaysTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetWeeksDaysTime. Error: %v", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}

func TestDurationUtility_SetStartEndTimes_02(t *testing.T) {

	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := DurationTriad{}

	dur.SetStartEndTimes(t1, t2, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	dDto, err := dur.GetYrMthWkDayHrMinSecNanosecs()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYrMthWkDayHrMinSecNanosecs. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 864197532-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}

func TestDurationUtility_SetStartEndTimes_03(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := DurationTriad{}

	dur.SetStartEndTimes(t2, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	dDto, err := dur.GetYearsMthsWeeksTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearsMthsWeeksTime. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 864-Milliseconds 197-Microseconds 532-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}

func TestDurationUtility_SetStartEndTimes_04(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	dur := DurationTriad{}

	dur.SetStartEndTimes(t1, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	dDto, err := dur.GetYearsMthsWeeksTime()
	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearsMthsWeeksTime. Error: %v", err.Error())
	}

	expected := "0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YearsMthsWeeksTime Duration: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}

func TestDurationUtility_SetStartTimePlusTime(t *testing.T) {
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

	if t1OutStr != dur.BaseTime.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.", t1OutStr, dur.BaseTime.TimeIn.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.", t2OutStr, dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr))
	}

	if t12Dur != dur.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.TimeDuration)
	}

	dDto, err := dur.GetYearMthDaysTime()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTime. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}

func TestDurationUtility_SetStartTimeMinusTime(t *testing.T) {
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

	dur.SetStartTimeMinusTime(t2, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if t1OutStr != dur.BaseTime.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.", t1OutStr, dur.BaseTime.TimeIn.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.", t2OutStr, dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr))
	}

	if t12Dur != dur.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.TimeDuration)
	}

	dDto, err := dur.GetYearMthDaysTime()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTime. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}

func TestDurationUtility_SetStartTimeDuration(t *testing.T) {
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

	if t1OutStr != dur.BaseTime.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.", t1OutStr, dur.BaseTime.TimeIn.DateTime.Format(fmtstr))
	}

	if t2OutStr != dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.", t2OutStr, dur.EndTimeTzu.TimeIn.DateTime.Format(fmtstr))
	}

	if t12Dur != dur.TimeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.TimeDuration)
	}

	dDto, err := dur.GetYearMthDaysTime()

	if err != nil {
		t.Errorf("Error from DurationTriad.GetYearMthDaysTime. Error: %v", err.Error())
	}

	expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != dDto.DisplayStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, dDto.DisplayStr)
	}

}
*/