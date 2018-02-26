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

	err = dTz1.AddDateTimeToThis(years, months, days, hours, minutes, seconds, milliseconds, microseconds, nanoseconds, fmtstr)

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

	if t1.Year() != dTzDto.Year {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'",t1.Year(), dTzDto.Year)
	}

	if int(t1.Month()) != dTzDto.Month {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'",int(t1.Month()), dTzDto.Month)
	}

	if t1.Hour() != dTzDto.Hour {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'",t1.Hour(), dTzDto.Hour)
	}

	if t1.Minute() != dTzDto.Minute {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'",t1.Minute(), dTzDto.Minute)
	}

	if t1.Second() != dTzDto.Second {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'",t1.Second(), dTzDto.Second)
	}

	if 38 != dTzDto.Millisecond {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'",  dTzDto.Millisecond )
	}

	if 175 != dTzDto.Microsecond {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Microsecond)
	}

	if 584 != dTzDto.Nanosecond {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Nanosecond)
	}

	if 38175584 != int(dTzDto.TotalNanoSecs) {
		t.Errorf("Expected Total Nanosecond Number='38175584'.  Instead Total Nanosecond Number='%v'",int(dTzDto.TotalNanoSecs))
	}

}