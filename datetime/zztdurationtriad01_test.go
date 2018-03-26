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

	du.SetStartEndTimesTz(t2, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	expected := "0-Hours 0-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearMthDaysTimeAbbrvStr()


	if expected != dOut {
		t.Errorf("Expected: %v. Error - got %v", expected, dOut)
	}
}

func TestDurationTriad_GetYearsMthsWeeksTimeAbbrv(t *testing.T) {

	t1str := "04/30/2017 22:58:31.987654321 -0500 CDT"
	t2str := "04/30/2017 22:59:33.123456789 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	du := DurationTriad{}

	du.SetStartEndTimesTz(t2, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	expected := "0-Hours 1-Minutes 1-Seconds 135-Milliseconds 802-Microseconds 468-Nanoseconds"

	dOut := du.BaseTime.GetYearsMthsWeeksTimeAbbrvStr()

	if expected != dOut {
		t.Errorf("Expected: %v. Error - got %v", expected, dOut)
	}

}

func TestDurationTriad_NewAutoEnd_01(t *testing.T) {

	locCentral, err := time.LoadLocation(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaUsCentral). Error='%v'", err.Error())
	}

	startDateTime := time.Now().In(locCentral)

	durT, err := DurationTriad{}.NewAutoEnd(startDateTime, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewAutoEnd() ")
	}

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if durT.BaseTime.TimeDuration > expectedDur {
		t.Error("Expected duration is greater than 2-seconds. Error!")
	}

}

func TestDurationTriad_NewAutoStart(t *testing.T) {

	durT, err := DurationTriad{}.NewAutoStart(TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewAutoStart() " +
			"Error='%v' ", err.Error())
	}

	durT.SetAutoEnd()

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

	dur, err := DurationTriad{}.NewStartTimeDurationTz(t1, t12Dur, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimeDurationTz(t1, t12Dur). Error='%v'", err.Error())
	}

	// dur.SetStartTimeDurationTz(t1, t12Dur)

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

func TestDurationTriad_NewStartDateTzDuration_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	t1DateTz, err := DateTzDto{}.NewTz(t1, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewTz(t1, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr) " +
			" Error ='%v' ", err.Error())
	}

	dur, err := DurationTriad{}.NewStartDateTzDuration(t1DateTz, t12Dur)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartDateTzDuration(t1DateTz, t12Dur). " +
			"Error='%v'", err.Error())
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

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.TimeZone.LocationName	{
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	t2Local := dur.LocalTime.StartTimeDateTz.DateTime.Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUTC). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.DateTime.Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	if dur.BaseTime.TimeDuration != dur.LocalTime.TimeDuration {
		t.Errorf("Expected Base Time Duration == Local Time Duration. It Did NOT! " +
			"dur.BaseTime.TimeDuration = '%v'  dur.LocalTime.TimeDuration='%v'",
			dur.BaseTime.TimeDuration, dur.LocalTime.TimeDuration)
	}

	if dur.BaseTime.TimeDuration != dur.UTCTime.TimeDuration {
		t.Errorf("Expected Base Time Duration == UTC Time Duration. It Did NOT! " +
			"dur.BaseTime.TimeDuration = '%v'  dur.LocalTime.TimeDuration='%v'",
			dur.BaseTime.TimeDuration, dur.UTCTime.TimeDuration)
	}

}

func TestDurationTriad_NewStartEndDateTzDto_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := DateTzDto{}.New(t1, FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := DateTzDto{}.New(t2, FmtDateTimeYrMDayFmtStr)

	dur, err := DurationTriad{}.NewStartEndDateTzDto(dateTz1, dateTz2)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2). Error='%v'", err.Error())
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

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.TimeZone.LocationName	{
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	t2Local := dur.LocalTime.StartTimeDateTz.DateTime.Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUTC). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.DateTime.Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

}

func TestDurationTriad_NewStartEndDateTzDtoTz(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := DateTzDto{}.New(t1, FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := DateTzDto{}.New(t2, FmtDateTimeYrMDayFmtStr)

	dur, err := DurationTriad{}.NewStartEndDateTzDtoTz(dateTz1, dateTz2, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2). Error='%v'", err.Error())
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

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.TimeZone.LocationName	{
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	t2Local := dur.LocalTime.StartTimeDateTz.DateTime.Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUTC). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.DateTime.Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

}


func TestDurationTriad_NewStartEndDateTzDtoCalcTz(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	dateTz1, err := DateTzDto{}.New(t1, FmtDateTimeYrMDayFmtStr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	dateTz2, err := DateTzDto{}.New(t2, FmtDateTimeYrMDayFmtStr)

	dur, err := DurationTriad{}.NewStartEndDateTzDtoCalcTz(
																	dateTz1,
																	dateTz2,
																	TDurCalcTypeSTDYEARMTH,
																	TzIanaUsCentral,
																	FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2). Error='%v'", err.Error())
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

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.TimeZone.LocationName	{
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	t2Local := dur.LocalTime.StartTimeDateTz.DateTime.Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUTC). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.DateTime.Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
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

	dur, err := DurationTriad{}.NewStartEndTimesTz(t1, t2, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartEndTimesTz(t1, t2). Error='%v'", err.Error())
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

	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.TimeZone.LocationName	{
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	t2Local := dur.LocalTime.StartTimeDateTz.DateTime.Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUTC). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.DateTime.Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
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

	dur, err := DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewEndTimeMinusTimeDtoTz(t2, timeDto). Error='%v'", err.Error())
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
	dur, err := DurationTriad{}.NewStartTimePlusTimeDtoTz(t1, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewStartTimePlusTimeDtoTz(t1, timeDto). Error='%v'", err.Error())
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


	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.TimeZone.LocationName	{
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	t2Local := dur.LocalTime.StartTimeDateTz.DateTime.Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUTC). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.DateTime.Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}


}

func TestDurationTriad_SetStartEndTimes(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	dur := DurationTriad{}

	dur.SetStartEndTimesTz(t1, t2, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

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


	loc, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(\"Local\"). Error='%v'", err.Error())
	}

	t1Local := t1.In(loc)

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.TimeZone.LocationName	{
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	t2Local := dur.LocalTime.StartTimeDateTz.DateTime.Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUTC). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.DateTime.Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

}

func TestDurationTriad_SetStartEndTimes_02(t *testing.T) {

	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := DurationTriad{}

	dur.SetStartEndTimesTz(t1, t2, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

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

	if t1Local.Location().String() != dur.LocalTime.StartTimeDateTz.TimeZone.LocationName	{
		t.Errorf("Expected Local Time Zone Location ='%v'. Actual Time Zone Location ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	if !t1Local.Equal(dur.LocalTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected Local Start Time ='%v'. Actual Local Start Time ='%v'. ",
			t1Local.Location().String(),
			dur.LocalTime.StartTimeDateTz.TimeZone.LocationName)
	}

	t2Local := dur.LocalTime.StartTimeDateTz.DateTime.Add(dur.LocalTime.TimeDuration)

	if !t2Local.Equal(dur.LocalTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected Local End Time='%v'. Actual Local End Time='%v'. ",
			t2Local.Format(FmtDateTimeYrMDayFmtStr),
			dur.LocalTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	loc, err = time.LoadLocation(TzIanaUTC)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUTC). Error='%v'", err.Error())
	}

	t1UTC := t1.In(loc)

	if !t1UTC.Equal(dur.UTCTime.StartTimeDateTz.DateTime) {
		t.Errorf("Expected UTC Start Time='%v'. Actual UTC Start Time='%v'. ",
			t1UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.StartTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	t2UTC := dur.UTCTime.StartTimeDateTz.DateTime.Add(dur.UTCTime.TimeDuration)

	if !t2UTC.Equal(dur.UTCTime.EndTimeDateTz.DateTime) {
		t.Errorf("Expected UTC End Time='%v'. Actual UTC End Time='%v'. ",
			t2UTC.Format(FmtDateTimeYrMDayFmtStr),
			dur.UTCTime.EndTimeDateTz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

}

func TestDurationTriad_SetStartEndTimes_03(t *testing.T) {
	t1str := "02/15/2014 19:54:30.123456789 -0600 CST"
	t2str := "04/30/2017 22:58:32.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	dur := DurationTriad{}

	dur.SetStartEndTimesTz(t2, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

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

	dur.SetStartEndTimesTz(t1, t1, TzIanaUsCentral,FmtDateTimeYrMDayFmtStr)

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
	dur.SetStartTimePlusTimeDtoTz(t1, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

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

	dur.SetEndTimeMinusTimeDtoTz(t2, timeDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

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

	dur.SetStartTimeDurationTz(t1, t12Dur, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

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
