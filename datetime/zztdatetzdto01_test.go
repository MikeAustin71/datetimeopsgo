package datetime

import (
	"testing"
	"time"
)

func TestDateTzDto_AddDate_01(t *testing.T) {

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	expectedOutDate := t1.Format(fmtstr)

	dtz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1, fmtstr). Error='%v'", err.Error())
	}

	if expectedOutDate != dtz1.String() {
		t.Errorf("Error: Expected dtz1.String()='%v'. Instead, dtz1.String()='%v' ",expectedOutDate, dtz1.String())
	}

	t2 := t1.AddDate(5, 6, 12)

	dtz2, err := dtz1.AddDate(5, 6, 12, fmtstr)

	if err != nil {
		t.Errorf("Error returned by dtz1.AddDate(5, 6, 12, fmtstr). Error='%v'", err.Error())
	}

	expectedOutDate = t2.Format(fmtstr)

	if expectedOutDate != dtz2.String() {
		t.Errorf("Error: Expected dtz2.String()='%v'. Instead, dtz2.String()='%v' ",expectedOutDate, dtz2.String())
	}

}

func TestDateTzDto_AddDateToThis_01(t *testing.T) {
	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	expectedOutDate := t1.Format(fmtstr)

	dtz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1, fmtstr). Error='%v'", err.Error())
	}

	if expectedOutDate != dtz1.String() {
		t.Errorf("Error: Expected dtz1.String()='%v'. Instead, dtz1.String()='%v' ",expectedOutDate, dtz1.String())
	}

	t2 := t1.AddDate(5, 6, 12)

	err = dtz1.AddDateToThis(5, 6, 12)

	if err != nil {
		t.Errorf("Error returned by AddDateToThis(5, 6, 12). Error='%v'", err.Error())
	}

	expectedOutDate = t2.Format(fmtstr)

	if expectedOutDate != dtz1.String() {
		t.Errorf("Error: Expected updated dtz1.String()='%v'. Instead, dtz1.String()='%v' ",expectedOutDate, dtz1.String())
	}

}

func TestDateTzDto_AddDuration_01(t *testing.T) {

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t2str := "2018-02-15 20:54:30.038175584 -0600 CST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	duration := t2.Sub(t1)

	dTz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1, fmtstr). Error='%v'", err.Error())
	}

	expectedOutStr := t1.Format(fmtstr)

	if expectedOutStr != dTz1.String() {
		t.Errorf("Error: Expected dTz1.String()='%v'. Instead, dTz1.String()='%v'",expectedOutStr, dTz1.String())
	}

	dTz2, err := dTz1.AddDuration(duration, fmtstr)

	if err != nil {
		t.Errorf("Error returned by dTz1.AddDuration(duration, fmtstr). Error='%v'", err.Error())
	}

	expectedOutStr = t2.Format(fmtstr)

	if expectedOutStr != dTz2.String() {
		t.Errorf("Error: Expected dTz2.String()='%v'. Instead, dTz2.String()='%v'",expectedOutStr, dTz2.String())
	}

}

func TestDateTzDto_AddDurationToThis_01(t *testing.T) {
	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t2str := "2018-02-15 20:54:30.038175584 -0600 CST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	duration := t2.Sub(t1)

	dTz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1, fmtstr). Error='%v'", err.Error())
	}

	expectedOutStr := t1.Format(fmtstr)

	if expectedOutStr != dTz1.String() {
		t.Errorf("Error: Expected dTz1.String()='%v'. Instead, dTz1.String()='%v'",expectedOutStr, dTz1.String())
	}

	err = dTz1.AddDurationToThis(duration)

	if err != nil {
		t.Errorf("Error returned by dTz1.AddDurationToThis(duration). Error='%v'", err.Error())
	}

	expectedOutStr = t2.Format(fmtstr)

	if expectedOutStr != dTz1.String() {
		t.Errorf("Error: Expected updated dTz1.String()='%v'. Instead, updated dTz1.String()='%v'",expectedOutStr, dTz1.String())
	}

}

func TestDateTzDto_AddTime_01(t *testing.T) {

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t1str := "2018-02-22 19:21:30.000000000 -0600 CST"

	hours := 50
	minutes := 24
	seconds := 38
	milliseconds := 600
	microseconds := 1500
	nanoseconds := 68473

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	t1, _ := time.Parse(fmtstr, t1str)

	dTz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1, fmtstr). Error='%v' ", err.Error())
	}

	expectedOutputStr := t1.Format(fmtstr)

	if expectedOutputStr != dTz1.String() {
		t.Errorf("Error: Expected dTz1.String()='%v'. Instead, dTz1.String()='%v'.",expectedOutputStr, dTz1.String())
	}

	t2 := t1.Add(time.Duration(totNanoSecs))

	dTz2, err := dTz1.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, fmtstr)

	if err != nil {
		t.Errorf("Error returned by dTz1.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds, fmtstr). Error='%v'", err.Error())
	}

	expectedOutputStr = t2.Format(fmtstr)

	if expectedOutputStr != dTz2.String() {
		t.Errorf("Error: Expected updated dTz2.String()='%v'. Instead, dTz2.String()='%v'. ",expectedOutputStr, dTz2.String())
	}

}

func TestDateTzDto_AddTimeToThis_01(t *testing.T) {

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t1str := "2018-02-22 19:21:30.000000000 -0600 CST"

	hours := 50
	minutes := 24
	seconds := 38
	milliseconds := 600
	microseconds := 1500
	nanoseconds := 68473

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	t1, _ := time.Parse(fmtstr, t1str)

	dTz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1, fmtstr). Error='%v' ", err.Error())
	}

	expectedOutputStr := t1.Format(fmtstr)

	if expectedOutputStr != dTz1.String() {
		t.Errorf("Error: Expected dTz1.String()='%v'. Instead, dTz1.String()='%v'.",expectedOutputStr, dTz1.String())
	}

	t2 := t1.Add(time.Duration(totNanoSecs))

	err = dTz1.AddTimeToThis(hours, minutes, seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		t.Errorf("Error returned by dTz1.AddTimeToThis(hours, minutes, seconds, milliseconds, microseconds, nanoseconds). Error='%v'", err.Error())
	}

	expectedOutputStr = t2.Format(fmtstr)

	if expectedOutputStr != dTz1.String() {
		t.Errorf("Error: Expected updated dTz1.String()='%v'. Instead, dTz1.String()='%v'. ",expectedOutputStr, dTz1.String())
	}

}

func TestDateTzDto_AddDateTime(t *testing.T) {

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t1str := "2014-02-18 19:21:30.000000000 -0600 CST"

	years := 5
	months := 6
	days := 18
	hours := 50
	minutes := 24
	seconds := 38
	milliseconds := 600
	microseconds := 1500
	nanoseconds := 68473

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	t1, _ := time.Parse(fmtstr, t1str)

	dTz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1, fmtstr). Error='%v' ", err.Error())
	}

	expectedOutputStr := t1.Format(fmtstr)

	if expectedOutputStr != dTz1.String() {
		t.Errorf("Error: Expected dTz1.String()='%v'. Instead, dTz1.String()='%v'.",expectedOutputStr, dTz1.String())
	}

	t2 := t1.AddDate(years, months, days)

	t3 := t2.Add(time.Duration(totNanoSecs))

	dTz2, err := dTz1.AddDateTime(years, months, days, hours, minutes, seconds, milliseconds, microseconds, nanoseconds, fmtstr)

	if err != nil {
		t.Errorf("Error returned by dTz1.AddDateTime(years, months, days, hours, minutes, ...) Error='%v'.", err.Error() )
	}

	expectedOutputStr = t3.Format(fmtstr)

	if expectedOutputStr != dTz2.String() {
		t.Errorf("Error: Expected updated dTz2.String()='%v'. Instead, dTz2.String()='%v'.",expectedOutputStr, dTz2.String())
	}

}

func TestDateTzDto_AddDateTimeToThis_01(t *testing.T) {

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t1str := "2014-02-18 19:21:30.000000000 -0600 CST"

	years := 5
	months := 6
	days := 18
	hours := 50
	minutes := 24
	seconds := 38
	milliseconds := 600
	microseconds := 1500
	nanoseconds := 68473

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	t1, _ := time.Parse(fmtstr, t1str)

	dTz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1, fmtstr). Error='%v' ", err.Error())
	}

	expectedOutputStr := t1.Format(fmtstr)

	if expectedOutputStr != dTz1.String() {
		t.Errorf("Error: Expected dTz1.String()='%v'. Instead, dTz1.String()='%v'.",expectedOutputStr, dTz1.String())
	}

	t2 := t1.AddDate(years, months, days)

	t3 := t2.Add(time.Duration(totNanoSecs))

	err = dTz1.AddDateTimeToThis(years, months, days, hours, minutes, seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		t.Errorf("Error returned by dTz1.AddDateTimeToThis(years, months, days, hours, minutes, seconds, ...). Error='%v'", err.Error())
	}

	expectedOutputStr = t3.Format(fmtstr)

	if expectedOutputStr != dTz1.String() {
		t.Errorf("Error: Expected updated dTz1.String()='%v'. Instead, dTz1.String()='%v'.",expectedOutputStr, dTz1.String())
	}

}

func TestDateTzDto_CopyIn_01(t *testing.T) {

	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t2str := "2018-02-22 19:21:30.000000000 -0600 CST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	dtz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1, fmtstr). Error='%v' ", err.Error())
	}

	actualOutStr := dtz1.String()

	if actualOutStr != t1OutStr {
		t.Errorf("Expected dtz1.String()='%v'. Instead, dtz1.String()='%v'", actualOutStr, t1OutStr)
	}

	dtz2, err := DateTzDto{}.New(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t2, fmtstr). Error='%v'", err.Error())
	}

	actualOutStr = dtz2.String()

	if actualOutStr != t2OutStr {
		t.Errorf("Expected dtz2.String()='%v'. Instead, dtz2.String()='%v'", actualOutStr, t2OutStr)
	}

	dtz2.CopyIn(dtz1)

}

func TestDateTzDto_CopyOut(t *testing.T) {

	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := t1.Location().String()

	dTzDtoInitial, err := DateTzDto{}.New(t1, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1). Error='%v'", err.Error())
	}

	dTzDto := dTzDtoInitial.CopyOut()

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'",t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone.ZoneName {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneName='%v'. Instead, dTzDto.TimeZone.ZoneName='%v'",t1ExpectedZone, dTzDto.TimeZone.ZoneName)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZone.ZoneOffsetSeconds {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneOffsetSeconds='%v'.  Instead, dTzDto.TimeZone.ZoneOffsetSeconds='%v'", t1ExpectedZoneOffset, dTzDto.TimeZone.ZoneOffsetSeconds)
	}

	if t1ExpectedLocationName != dTzDto.TimeZone.LocationName {
		t.Errorf("Error: Expected dTzDto.TimeZone.LocationName='%v'.  Instead, dTzDto.TimeZone.LocationName='%v'", t1ExpectedLocationName, dTzDto.TimeZone.LocationName)
	}

	if t1.Year() != dTzDto.Time.Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'",t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != dTzDto.Time.Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'",int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Day() != dTzDto.Time.DateDays {
		t.Errorf("Expected Date Day Number='%v'.  Instead Date Day Number='%v'",t1.Day(), dTzDto.Time.DateDays)
	}

	if t1.Hour() != dTzDto.Time.Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'",t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != dTzDto.Time.Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'",t1.Minute(), dTzDto.Time.Minutes)
	}

	if t1.Second() != dTzDto.Time.Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'",t1.Second(), dTzDto.Time.Seconds)
	}

	if 38 != dTzDto.Time.Milliseconds {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'",  dTzDto.Time.Milliseconds )
	}

	if 175 != dTzDto.Time.Microseconds {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Time.Microseconds)
	}

	if 584 != dTzDto.Time.Nanoseconds {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Time.Nanoseconds)
	}

	if 38175584 != dTzDto.Time.TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanoseconds Number='38175584'.  Instead Total Nanoseconds Number='%v'", dTzDto.Time.TotSubSecNanoseconds)
	}

}

func TestDateTzDto_GetTimeDto_01(t *testing.T) {

	locUSCentral, err := time.LoadLocation(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaUsCentral). Error='%v'", err.Error())
	}

	year := 2018
	month := 3
	day := 6
	hour := 20
	minute := 2
	second := 18
	nSecs := 792489279

	t4USCentral := time.Date(year, time.Month(month),day,hour,minute,second,nSecs,locUSCentral)

	dTz1, err := DateTzDto{}.New(t4USCentral, FmtDateTimeYrMDayFmtStr)
	
	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t4USCentral, FmtDateTimeYrMDayFmtStr)")
	}
	
	tDto, err := dTz1.GetTimeDto()
	
	if err != nil {
		t.Errorf("Error returned by dTz1.GetTimeDto(). Error=%v'", err.Error())
	}

	if year != tDto.Years {
		t.Errorf("Error: Expected Years='%v'. Instead, Years='%v'",year, tDto.Years)
	}

	if month != tDto.Months {
		t.Errorf("Error: Expected Months='%v'. Instead, Months='%v'",month, tDto.Months)
	}

	if day != tDto.DateDays {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",day, tDto.DateDays)
	}

	if hour != tDto.Hours {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",hour, tDto.Hours)
	}

	if minute != tDto.Minutes {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",minute, tDto.Minutes)
	}

	if second != tDto.Seconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",second, tDto.Seconds)
	}

	if 792 != tDto.Milliseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",792, tDto.Milliseconds)
	}
	
	if 489 != tDto.Microseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",489, tDto.Microseconds)
	}
	
	if 279 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",279, tDto.Nanoseconds)
	}
	
	if nSecs != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected tDto.TotSubSecNanoseconds='%v'. "+
			"Instead, tDto.TotSubSecNanoseconds='%v'", nSecs, tDto.TotSubSecNanoseconds)
	}

	totTime := int64(hour) * int64(time.Hour)
	totTime += int64(minute) * int64(time.Minute)
	totTime += int64(second) * int64(time.Second)
	totTime += int64(nSecs)

	if totTime != tDto.TotTimeNanoseconds {
		t.Errorf("Error: Expected tDto.TotTimeNanoseconds='%v'. "+
			"Instead, tDto.TotTimeNanoseconds='%v'", totTime, tDto.TotTimeNanoseconds)
	}

}

func TestDateTzDto_GetDateTimeTzNanoSecYMDDowText(t *testing.T) {

	tDto := TimeDto {
		Years: 2018,
		Months: 2,
		DateDays: 6,
		Hours: 20,
		Minutes: 2,
		Seconds: 18,
		Nanoseconds: 792489279,
	}

	tDto.NormalizeTimeElements()

	dTz, err := DateTzDto{}.NewTimeDto(tDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by " +
			"DateTzDto{}.NewTimeDto(tDto, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr) " +
			"Error='%v'", err.Error())
	}

	expectedStr := "2018-02-06 20:02:18.792489279 -0600 CST"

	actualStr := dTz.GetDateTimeYrMDayTzFmtStr()

	if expectedStr != actualStr {
		t.Errorf("Expected date time string='%v'.  Instead date time string='%v' ",
			expectedStr, actualStr)
	}
}