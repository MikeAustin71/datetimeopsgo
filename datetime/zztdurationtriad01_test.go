package datetime

import (
	"testing"
	"time"
)

func TestDurationTriad_GetYearMthDaysTimeAbbrv(t *testing.T) {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := DurationTriad{}

	err := du.SetStartEndTimesTz(
		t2,
		t1,
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by du.SetStartEndTimesTz(t2, t1, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr).\nError='%v'\n", err.Error())
	}

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearMthDaysTimeAbbrvStr()

	if expected != dOut {
		t.Errorf("Expected: %v.\nInstead, " +
			"result= %v.\n", expected, dOut)
	}
}

func TestDurationTriad_GetYearsMthsWeeksTimeAbbrv(t *testing.T) {

	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:59:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := DurationTriad{}

	err := du.SetStartEndTimesTz(t2, t1, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by du.SetStartEndTimesTz(t2, t1, TZones.US.Central(), "+
			"FmtDateTimeYrMDayFmtStr). Error='%v' ", err.Error())
	}

	expected := "0-Hours 1-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearsMthsWeeksTimeAbbrvStr()

	if expected != dOut {
		t.Errorf("Expected: %v. Error - got %v", expected, dOut)
	}

}

func TestDurationTriad_NewAutoEnd_01(t *testing.T) {

	locCentral, err := time.LoadLocation(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	startDateTime := time.Now().In(locCentral)

	durT, err := DurationTriad{}.NewAutoEnd(startDateTime, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewAutoEnd().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if durT.BaseTime.TimeDuration > expectedDur {
		t.Error("Expected duration is greater than 2-seconds. Error!")
	}

}

func TestDurationTriad_NewAutoStart(t *testing.T) {

	durT, err := DurationTriad{}.NewAutoStart(TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewAutoStart().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	err = durT.SetAutoEnd()

	if err != nil {
		t.Errorf("Error returned by durT.SetAutoEnd() "+
			"Error='%v' ", err.Error())
	}

	expectedDur := time.Duration(int64(3) * int64(time.Second))

	if durT.BaseTime.TimeDuration > expectedDur {
		t.Error("Expected duration is greater than 3-seconds. Error!")
	}

}

func TestDurationTriad_NewStartTimeDuration_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	dur, err := DurationTriad{}.NewStartTimeDurationTz(t1, t12Dur, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimeDurationTz(t1, t12Dur).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// dur.SetStartTimeDurationTz(t1, t12Dur)

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
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

func TestDurationTriad_NewStartDateTzDuration_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	t1DateTz, err := DateTzDto{}.NewTz(
		t1,
		TZones.US.Central(),
		TzConvertType.Relative(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewTz(t1, TZones.US.Central(), FmtDateTimeYrMDayFmtStr) "+
			" Error ='%v' ", err.Error())
	}

	dur, err := DurationTriad{}.NewStartDateTzDuration(t1DateTz, t12Dur)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartDateTzDuration(t1DateTz, t12Dur).\n"+
			"Error='%v'\n", err.Error())
		return
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
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.BaseTime.TimeDuration)
	}

	outStr := dur.BaseTime.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.GetOriginalTzName())
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.GetDateTimeValue()) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
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

func TestDurationTriad_NewStartEndDateTzDto_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := DateTzDto{}.NewDateTime(t1, FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := DateTzDto{}.NewDateTime(t2, FmtDateTimeYrMDayFmtStr)

	dur, err := DurationTriad{}.NewStartEndDateTzDto(dateTz1, dateTz2)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
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
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
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

func TestDurationTriad_NewStartEndDateTzDtoTz(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := DateTzDto{}.NewDateTime(t1, FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := DateTzDto{}.NewDateTime(t2, FmtDateTimeYrMDayFmtStr)

	dur, err := DurationTriad{}.NewStartEndDateTzDtoTz(dateTz1, dateTz2, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v.\n" +
			"Instead, got %v ",
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
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
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
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

func TestDurationTriad_NewStartEndDateTzDtoCalcTz(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := DateTzDto{}.NewDateTime(t1, FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := DateTzDto{}.NewDateTime(t2, FmtDateTimeYrMDayFmtStr)

	dur, err := DurationTriad{}.NewStartEndDateTzDtoCalcTz(
		dateTz1,
		dateTz2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.StartTimeDateTz of %v.\n" +
			"Instead, got %v\n" ,
			t1OutStr, dur.BaseTime.StartTimeDateTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.EndTimeDateTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.EndTimeDateTz.GetDateTimeValue().Format(fmtstr))
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
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
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
		t.Errorf("Expected Local Start Time ='%v'.\nActual Local Start Time ='%v'.\n",
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

func TestDurationTriad_NewStartEndTimes_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	dur, err := DurationTriad{}.NewStartEndTimesTz(t1, t2, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
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
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr(). "+
			"Error='%v'", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v. Instead, got %v",
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

func TestDurationTriad_NewStartTimeMinusTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	dur, err := DurationTriad{}.NewEndTimeMinusTimeDtoTz(
		TCalcMode.LocalTimeZone(),
		t2,
		timeDto,
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
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

func TestDurationTriad_NewStartTimePlusTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}
	dur, err := DurationTriad{}.NewStartTimePlusTimeDtoTz(
		TCalcMode.LocalTimeZone(),
		t1,
		timeDto,
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimePlusTimeDtoTz(t1, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
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

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'. ",
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
