package datetime

import (
	"testing"
	"time"
)

func TestDurationTriad_CopyIn_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	var durX, dur DurationTriad
	var err error

	durX, err = DurationTriad{}.NewStartEndTimes(
		t1,
		t2,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by durX=DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dur.CopyIn(durX)

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
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, dur.BaseTime.timeDuration)
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

func TestDurationTriad_GetYearMthDaysTimeAbbrv(t *testing.T) {
	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:58:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := DurationTriad{}

	err := du.SetStartEndTimes(
		t1,
		t2,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
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

	err := du.SetStartEndTimes(
		t2,
		t1,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

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

	durT, err := DurationTriad{}.NewAutoEnd(
		startDateTime,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewAutoEnd().\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if durT.BaseTime.timeDuration > expectedDur {
		t.Error("Expected duration is greater than 2-seconds. Error!")
	}

}

func TestDurationTriad_NewAutoStart(t *testing.T) {

	durT, err := DurationTriad{}.NewAutoStart(
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

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

	if durT.BaseTime.timeDuration > expectedDur {
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

	dur, err := DurationTriad{}.NewStartTimeDuration(
		t1,
		t12Dur,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimeTzDuration(t1, t12Dur).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	// dur.SetStartTimeDurationTz(t1, t12Dur)

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != dur.BaseTime.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.BaseTime.timeDuration)
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
		return
	}

	dur, err := DurationTriad{}.NewStartTimeTzDuration(
		t1DateTz,
		t12Dur,
		TDurCalc.StdYearMth(),
		t1DateTz.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		t1DateTz.GetDateTimeFmt())

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartDateTzDuration(t1DateTz, t12Dur).\n"+
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
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v", t12Dur, dur.BaseTime.timeDuration)
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

	if t1Local.Location().String() != dur.LocalTime.startDateTimeTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.startDateTimeTz.GetOriginalTzName())
	}

	if !t1Local.Equal(dur.LocalTime.startDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
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

func TestDurationTriad_NewStartEndTimesTz_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := DateTzDto{}.NewDateTime(t1, FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := DateTzDto{}.NewDateTime(t2, FmtDateTimeYrMDayFmtStr)

	dur, err := DurationTriad{}.NewStartEndTimesTz(
		dateTz1,
		dateTz2,
		TDurCalc.StdYearMth(),
		dateTz1.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		dateTz1.GetDateTimeFmt())

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
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
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, dur.BaseTime.timeDuration)
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
		t.Errorf("Error - Expected WeekDays Duration: %v.\n" +
			"Instead, got %v", expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumHoursTimeStr()

	if err != nil {
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr().\n"+
			"Error='%v'\n", err.Error())
	}

	expected = "28082-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Hours Duration: %v.\n" +
			"Instead, got %v\n", expected, outStr)
	}

	outStr = dur.BaseTime.GetYrMthWkDayHrMinSecNanosecsStr()

	expected = "3-Years 2-Months 2-Weeks 1-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthWkDayHourSecNanosec Duration: %v.\n" +
			"Instead, got %v",
			expected, outStr)
	}

	outStr, err = dur.BaseTime.GetCumWeeksDaysTimeStr()

	if err != nil {
		t.Errorf("Error returned by dur.BaseTime.GetCumHoursTimeStr().\n"+
			"Error='%v'\n", err.Error())
	}

	expected = "167-Weeks 1-WeekDays 2-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected Weeks WeekDays Duration: %v.\n" +
			"Instead, got %v\n",
			expected, outStr)
	}

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\").\n" +
			"Error='%v'\n", err.Error())
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

	t2Local := dur.LocalTime.startDateTimeTz.
		GetDateTimeValue().Add(dur.LocalTime.timeDuration)

	if !t2Local.Equal(dur.LocalTime.endDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected Local End Time='%v'.\n" +
			"Actual Local End Time='%v'.\n",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.endDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.UTC()).\n" +
			"Error='%v'\n", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.startDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC Start Time='%v'.\n" +
			"Actual UTC Start Time='%v'.\n",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.startDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.startDateTimeTz.GetDateTimeValue().Add(dur.UTCTime.timeDuration)

	if !t2UTC.Equal(dur.UTCTime.endDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected UTC End Time='%v'.\n" +
			"Actual UTC End Time='%v'.\n",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.endDateTimeTz.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

}

func TestDurationTriad_NewStartEndTimesTz_02(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := DateTzDto{}.NewDateTime(t1, FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := DateTzDto{}.NewDateTime(t2, FmtDateTimeYrMDayFmtStr)

	dur, err := DurationTriad{}.NewStartEndTimesTz(
		dateTz1,
		dateTz2,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v.\n" +
			"Instead, got %v ",
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.BaseTime.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, dur.BaseTime.timeDuration)
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

func TestDurationTriad_NewStartEndTimesTz_03(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := DateTzDto{}.NewDateTime(t1, FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := DateTzDto{}.NewDateTime(t2, FmtDateTimeYrMDayFmtStr)

	dur, err := DurationTriad{}.NewStartEndTimesTz(
		dateTz1,
		dateTz2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v.\n" +
			"Instead, got %v\n" ,
			t1OutStr, dur.BaseTime.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, dur.BaseTime.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != dur.BaseTime.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, dur.BaseTime.timeDuration)
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

	if t1Local.Location().String() != dur.LocalTime.startDateTimeTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'.\n",
			t1Local.Location().String(),
			dur.LocalTime.startDateTimeTz.GetOriginalTzName())
	}

	if !t1Local.Equal(dur.LocalTime.startDateTimeTz.GetDateTimeValue()) {
		t.Errorf("Expected Local Start Time ='%v'.\nActual Local Start Time ='%v'.\n",
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

func TestDurationTriad_NewStartEndTimes_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	dur, err := DurationTriad{}.NewStartEndTimes(
		t1,
		t2,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2).\n" +
			"Error='%v'\n", err.Error())
		return
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
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, dur.BaseTime.timeDuration)
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

func TestDurationTriad_NewEndTimeMinusTimeDto_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	dur, err := DurationTriad{}.NewEndTimeMinusTimeDto(
		t2,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto).\n" +
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

func TestDurationTriad_NewEndTimeTzMinusTimeDto_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var dur DurationTriad
	var err error
	var dTz2 DateTzDto

	dTz2, err = DateTzDto{}.NewDateTime(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t2, fmtstr).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dur, err = DurationTriad{}.NewEndTimeTzMinusTimeDto(
		dTz2,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewEndTimeTzMinusTimeDto(dTz2, timeDto).\n" +
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

func TestDurationTriad_NewStartTimePlusTimeDto_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	dur, err := DurationTriad{}.NewStartTimePlusTimeDto(
		t1,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimePlusTimeDtoTz(t1, timeDto).\n" +
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
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.startDateTimeTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'. ",
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

func TestDurationTriad_NewStartTimeTzPlusTimeDto_01(t *testing.T) {
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
	var dur DurationTriad
	var dTz1 DateTzDto

	dTz1, err = DateTzDto{}.NewDateTime(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(t1, fmtstr).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dur, err = DurationTriad{}.NewStartTimeTzPlusTimeDto(
		dTz1,
		timeDto,
		TDurCalc.StdYearMth(),
		TZones.US.Central(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimePlusTimeDtoTz(t1, timeDto).\n" +
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
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.startDateTimeTz.GetOriginalTzName() {
		t.Errorf("Expected Local Time Zone Location ='%v'.\n" +
			"Actual Time Zone Location ='%v'. ",
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
