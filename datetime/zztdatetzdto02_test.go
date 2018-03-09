package datetime

import (
	"testing"
	"time"
)

func TestDateTzDto_NewDateTimeElements_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.158712300 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := TzIanaUsCentral

	dTzDto, err := DateTzDto{}.NewDateTimeElements(2014, 2,15,19,54,30,158712300, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTimeElements(2014, 2,15,19,54,30,158712300, TzUsCentral). Error='%v'", err.Error())
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'",t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone.ZoneName {
		t.Errorf("Error: Expected dTzDto.TimeZone='%v'. Instead, dTzDto.TimeZone.ZoneName='%v'",t1ExpectedZone, dTzDto.TimeZone.ZoneName)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZone.ZoneOffsetSeconds {
		t.Errorf("Error: Expected dTzDto.TimeZoneOffset='%v'.  Instead, dTzDto.TimeZone.OffsetSeconds='%v'", t1ExpectedZoneOffset, dTzDto.TimeZone.ZoneOffsetSeconds)
	}

	if t1ExpectedLocationName != dTzDto.TimeZone.LocationName {
		t.Errorf("Error: Expected dTzDto.TimeLocName='%v'.  Instead, dTzDto.TimeZone.LocationName='%v'", t1ExpectedLocationName, dTzDto.TimeZone.LocationName)
	}

	if t1.Year() != int(dTzDto.Time.Years) {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'",t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != int(dTzDto.Time.Months) {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'",int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Hour() != int(dTzDto.Time.Hours) {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'",t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != int(dTzDto.Time.Minutes) {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'",t1.Minute(), dTzDto.Time.Minutes)
	}
	if t1.Second() != int(dTzDto.Time.Seconds) {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'",t1.Second(), dTzDto.Time.Seconds)
	}

	if t1.Nanosecond() != int(dTzDto.Time.TotSubSecNanoseconds) {
		t.Errorf("Expected Total Nanosecond Number='%v'.  Instead Total Nanosecond Number='%v'",t1.Nanosecond(), int(dTzDto.Time.TotSubSecNanoseconds))
	}

	r := t1.Nanosecond()

	if r == 0 {
		return
	}

	millisecond := r / int(time.Millisecond)

	if millisecond != int(dTzDto.Time.Milliseconds) {
		t.Errorf("Expected Millisecond='%v'.  Instead, Millisecond='%v'", millisecond, dTzDto.Time.Milliseconds )
	}

	r -= millisecond * int(time.Millisecond)

	if r == 0 {
		return
	}

	microsecond := r / int(time.Microsecond)

	if microsecond != int(dTzDto.Time.Microseconds) {
		t.Errorf("Expected Microsecond='%v'.  Instead, Microsecond='%v'", microsecond, dTzDto.Time.Microseconds)
	}

	r-= microsecond * int(time.Microsecond)

	if r == 0 {
		return
	}

	if r != int(dTzDto.Time.Nanoseconds) {
		t.Errorf("Expected Nanosecond='%v'.  Instead Nanosecond='%v' ", r, dTzDto.Time.Nanoseconds)
	}

}

func TestDateTzDto_NewDateTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := TzIanaUsCentral

	dTzDto, err := DateTzDto{}.NewDateTime(2014, 2,15,19,54,30,38, 175, 584, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(2014, 2,15,19,54,30,38, 175, 584, TzUsCentral). Error='%v'", err.Error())
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'",t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone.ZoneName {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneName='%v'. Instead, dTzDto.TimeZone.ZoneName='%v'",t1ExpectedZone, dTzDto.TimeZone.ZoneName)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZone.ZoneOffsetSeconds {
		t.Errorf("Error: Expected dTzDto.TimeZoneOffset.ZoneOffsetSeconds='%v'.  Instead, dTzDto.TimeZoneOffset.ZoneOffsetSeconds='%v'", t1ExpectedZoneOffset, dTzDto.TimeZone.ZoneOffsetSeconds)
	}

	if t1ExpectedLocationName != dTzDto.TimeZone.LocationName {
		t.Errorf("Error: Expected dTzDto.TimeZone.LocationName='%v'.  Instead, dTzDto.TimeZone.LocationName='%v'", t1ExpectedLocationName, dTzDto.TimeZone.LocationName)
	}

	if t1.Year() != int(dTzDto.Time.Years) {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'",t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != int(dTzDto.Time.Months) {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'",int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Hour() != int(dTzDto.Time.Hours) {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'",t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != int(dTzDto.Time.Minutes) {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'",t1.Minute(), dTzDto.Time.Minutes)
	}

	if t1.Second() != int(dTzDto.Time.Seconds) {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'",t1.Second(), dTzDto.Time.Seconds)
	}

	r := 38 * int(time.Millisecond)

	if 38 != int(dTzDto.Time.Milliseconds) {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'",  dTzDto.Time.Milliseconds )
	}

	r += 175 * int(time.Microsecond)

	if 175 != int(dTzDto.Time.Microseconds) {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Time.Microseconds)
	}

	if 584 != int(dTzDto.Time.Nanoseconds) {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Time.Nanoseconds)
	}

	r += 584

	if r != int(dTzDto.Time.TotSubSecNanoseconds) {
		t.Errorf("Expected Total Nanosecond Number='%v'.  Instead Total Nanosecond Number='%v'",r, int(dTzDto.Time.TotSubSecNanoseconds))
	}

}

func TestDateTzDto_New_01(t *testing.T) {

	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := t1.Location().String()

	dTzDto, err := DateTzDto{}.New(t1, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1). Error='%v'", err.Error())
	}

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

	if t1.Year() != int(dTzDto.Time.Years) {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'",t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != int(dTzDto.Time.Months) {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'",int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Hour() != int(dTzDto.Time.Hours) {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'",t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != int(dTzDto.Time.Minutes) {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'",t1.Minute(), dTzDto.Time.Minutes)
	}

	if t1.Second() != int(dTzDto.Time.Seconds) {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'",t1.Second(), dTzDto.Time.Seconds)
	}

	if 38 != int(dTzDto.Time.Milliseconds) {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'",  dTzDto.Time.Milliseconds )
	}

	if 175 != int(dTzDto.Time.Microseconds) {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Time.Microseconds)
	}

	if 584 != int(dTzDto.Time.Nanoseconds) {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Time.Nanoseconds)
	}

	if 38175584 != int(dTzDto.Time.TotSubSecNanoseconds) {
		t.Errorf("Expected Total Nanosecond Number='38175584'.  Instead Total Nanosecond Number='%v'",int(dTzDto.Time.TotSubSecNanoseconds))
	}

}

func TestDateTzDto_NewTimeDto_01(t *testing.T) {

	t0str := "2017-04-30 22:58:32.515539300 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t0, err := time.Parse(fmtstr, t0str)

	if err != nil {
		t.Errorf("Error retruned from time.Parse(fmtstr, t0str). t0str='%v'  Error='%v'", t0str, err.Error())
	}

	tDto, err := TimeDto{}.New(2017, 04, 0, 30, 22, 58,32,0,0, 515539300)

	if err != nil {
		t.Errorf("Error returned from TimeDto{}.New(...)  Error='%v'", err.Error())
	}

	dTzDto, err := DateTzDto{}.NewTimeDto(tDto, TzIanaUsCentral, fmtstr)

	if !dTzDto.DateTime.Equal(t0) {
		t.Errorf("Error returned from DateTzDto{}.NewTimeDto(tDto, TzIanaUsCentral, fmtstr). Error='%v'", err.Error())
	}

	if t0str != dTzDto.String() {
		t.Errorf("Error on formats: Expected date time string='%v'. Instead, date time string='%v'", t0str, dTzDto.String())
	}

}

func TestDateTzDto_NewTz_01(t *testing.T) {

	locUSCentral, err := time.LoadLocation(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaUsCentral). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(TzIanaAsiaTokyo)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaAsiaTokyo). Error='%v'", err.Error())
	}

	t4USCentral := time.Date(2018, time.Month(3),06,20,02,18,792489279,locUSCentral)

	t4AsiaTokyo := t4USCentral.In(locTokyo)

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dTz, err := DateTzDto{}.NewTz(t4AsiaTokyo, TzIanaUsCentral, fmtstr)

	if !t4USCentral.Equal(dTz.DateTime) {
		t.Errorf("Error: Expected DateTime='%v'. Instead DateTime='%v'",
										t4USCentral.Format(fmtstr), dTz.DateTime.Format(fmtstr))
	}

}

func TestDateTzDto_SetFromDateTime_01(t *testing.T) {
	t0str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t0, _ := time.Parse(fmtstr, t0str)
	t1, _ := time.Parse(fmtstr, t1str)

	dTzDto, err := DateTzDto{}.New(t0, FmtDateTimeYrMDayFmtStr)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := TzIanaUsCentral

	err = dTzDto.SetFromDateTime(2014, 2,15,19,54,30,38, 175, 584, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(2014, 2,15,19,54,30,38, 175, 584, TzUsCentral). Error='%v'", err.Error())
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'",t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone.ZoneName {
		t.Errorf("Error: Expected dTzDto.TimeZone='%v'. Instead, dTzDto.TimeZone='%v'",t1ExpectedZone, dTzDto.TimeZone)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZone.ZoneOffsetSeconds {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneOffsetSeconds='%v'.  Instead, dTzDto.TimeZone.ZoneOffsetSeconds='%v'", t1ExpectedZoneOffset, dTzDto.TimeZone.ZoneOffsetSeconds)
	}

	if t1ExpectedLocationName != dTzDto.TimeZone.LocationName {
		t.Errorf("Error: Expected dTzDto.TimeZone.LocationName='%v'.  Instead, dTzDto.TimeZone.LocationName='%v'", t1ExpectedLocationName, dTzDto.TimeZone.LocationName)
	}

	if t1.Year() != int(dTzDto.Time.Years) {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'",t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != int(dTzDto.Time.Months) {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'",int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Hour() != int(dTzDto.Time.Hours) {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'",t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != int(dTzDto.Time.Minutes) {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'",t1.Minute(), dTzDto.Time.Minutes)
	}

	if t1.Second() != int(dTzDto.Time.Seconds) {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'",t1.Second(), dTzDto.Time.Seconds)
	}

	r := 38 * int(time.Millisecond)

	if 38 != int(dTzDto.Time.Milliseconds) {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'",  dTzDto.Time.Milliseconds )
	}

	r += 175 * int(time.Microsecond)

	if 175 != int(dTzDto.Time.Microseconds) {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Time.Microseconds)
	}

	if 584 != int(dTzDto.Time.Nanoseconds) {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Time.Nanoseconds)
	}

	r += 584

	if r != int(dTzDto.Time.TotSubSecNanoseconds) {
		t.Errorf("Expected Total Nanosecond Number='%v'.  Instead Total Nanosecond Number='%v'",r, int(dTzDto.Time.TotSubSecNanoseconds))
	}

}

func TestDateTzDto_SetNewTimeZone_01(t *testing.T) {

	locUSCentral, err := time.LoadLocation(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaUsCentral). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(TzIanaAsiaTokyo)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaAsiaTokyo). Error='%v'", err.Error())
	}

	t4USCentral := time.Date(2018, time.Month(3),06,20,02,18,792489279,locUSCentral)

	t4AsiaTokyo := t4USCentral.In(locTokyo)

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dTz1, err := DateTzDto{}.New(t4USCentral, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t4USCentral, fmtstr) Error='%v", err.Error())
	}

	dTz1.SetNewTimeZone(TzIanaAsiaTokyo)

	if !t4AsiaTokyo.Equal(dTz1.DateTime) {
		t.Errorf("Error: Expected converted dTz1 date time = '%v'.  Instead, dTz1 date time='%v'",
			t4AsiaTokyo.Format(FmtDateTimeYrMDayFmtStr), dTz1.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	if TzIanaAsiaTokyo != dTz1.TimeZone.LocationName {
		t.Errorf("Error: Expected dTz1 Time Zone Location Name ='%v'. "+
			"Instead, Time Zone Location Name='%v'",TzIanaAsiaTokyo, dTz1.TimeZone.LocationName)
	}

}

func TestDateTzDto_Sub_01(t *testing.T) {
	t1str := "2014-02-15 19:54:30.038175584 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t2str := "2018-02-15 20:54:30.038175584 -0600 CST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)

	expectedDuration := t2.Sub(t1)

	dTz1, err := DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.New(t1, fmtstr). Error='%v'", err.Error())
	}

	dTz2, err := DateTzDto{}.New(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.New(t2, fmtstr). Error='%v'", err.Error())
	}

	actualDuration := dTz2.Sub(dTz1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: expected duration='%v'.  Instead, duration='%v' ",expectedDuration, actualDuration)
	}


}
