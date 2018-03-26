package datetime

import (
	"testing"
	"time"
	"fmt"
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

func TestDateTzDto_New_02(t *testing.T) {

	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := t1.Location().String()

	dTzDto, err := DateTzDto{}.New(t1, "")

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

func TestDateTzDto_NewNowTz_01(t *testing.T) {

	t0 := time.Now().Local()

	dTz, err := DateTzDto{}.NewNowTz(TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewNowTz(TzIanaUsCentral, FmtDateTimeYrMDayFmtStr). " +
			"Error='%v'", err.Error())
	}

	loc, err := time.LoadLocation(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUsCentral). Error='%v'", err.Error())
	}

	t1 := t0.In(loc)

	actualDur := t1.Sub(dTz.DateTime)

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if actualDur > expectedDur {
		t.Error("Error: Actual Duration exceeded 2-seconds!")
	}

	if TzIanaUsCentral != dTz.TimeZone.LocationName {
		t.Errorf("Error: Expected Time Zone='%v'.  Actual TimeZone='%v' ",
			TzIanaUsCentral, dTz.TimeZone.LocationName)
	}

}


func TestDateTzDto_NewNowLocal_01(t *testing.T) {

	t0 := time.Now().Local()
	dTz, err := DateTzDto{}.NewNowLocal(FmtDateTimeYrMDayFmtStr)


	if err != nil {
		t.Errorf("Error returned from time.NewNowLocal(FmtDateTimeYrMDayFmtStr). " +
			"Error='%v'", err.Error())
	}

	actualDur := t0.Sub(dTz.DateTime)

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if actualDur > expectedDur {
		t.Error("Error: Actual Duration exceeded 2-seconds!")
	}

	if TzGoLocal != dTz.TimeZone.LocationName {
		t.Errorf("Error: Expected Time Zone='%v'.  Actual TimeZone='%v' ",
			TzGoLocal, dTz.TimeZone.LocationName)
	}

}

func TestDateTzDto_NewNowUTC_01(t *testing.T) {

	t0 := time.Now().Local()

	loc, err := time.LoadLocation(TzIanaUTC)

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TzIanaUsCentral). Error='%v'", err.Error())
	}

	t1 := t0.In(loc)

	dTz, err := DateTzDto{}.NewNowUTC(FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewNowUTC(FmtDateTimeYrMDayFmtStr). " +
			"Error='%v'", err.Error())
	}


	actualDur := t1.Sub(dTz.DateTime)

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if actualDur > expectedDur {
		t.Error("Error: Actual Duration exceeded 2-seconds!")
	}

	if TzIanaUTC != dTz.TimeZone.LocationName {
		t.Errorf("Error: Expected Time Zone='%v'.  Actual TimeZone='%v' ",
			TzIanaUTC, dTz.TimeZone.LocationName)
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

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewTz(t4AsiaTokyo, TzIanaUsCentral, fmtstr). Error='%v'",
			err.Error())
	}

	if !t4USCentral.Equal(dTz.DateTime) {
		t.Errorf("Error: Expected DateTime='%v'. Instead DateTime='%v'",
										t4USCentral.Format(fmtstr), dTz.DateTime.Format(fmtstr))
	}

	eTimeZoneDef, err := TimeZoneDefDto{}.New(t4USCentral)

	if !eTimeZoneDef.Equal(dTz.TimeZone) {
		t.Errorf("Expected dTz.TimeZone.LocationName='%v'. " +
			"Instead, dTz.TimeZone.LocationName='%v'",
			eTimeZoneDef.LocationName, dTz.TimeZone.LocationName)
	}

	tDto, err := TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral) " +
			"t4USCentral='%v' Error='%v'",
			t4USCentral.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	expectedDt, err := tDto.GetDateTime(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned from tDto.GetDateTime(TzIanaUsCentral). " +
			"Error='%v'", err.Error())
	}

	actualDt, err := dTz.Time.GetDateTime(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned from dTz.Time.GetDateTime(TzIanaUsCentral). " +
			"Error='%v'", err.Error())
	}


	if !tDto.Equal(dTz.Time) {
		t.Errorf("Expected dTz.Time (TimeDto) == '%v' Instead, dTz.Time (TimeDto) == '%v'",
			expectedDt.Format(FmtDateTimeYrMDayFmtStr), actualDt.Format(FmtDateTimeYrMDayFmtStr))
	}

	if FmtDateTimeYrMDayFmtStr != dTz.DateTimeFmt {
		t.Errorf("Expected dTz.DateTimeFmt='%v' Instead, dTz.DateTimeFmt='%v' ",
			FmtDateTimeYrMDayFmtStr, dTz.DateTimeFmt)
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

func TestDateTzDto_SetFromTimeTz_01(t *testing.T) {

	locEuropeLondon, err := time.LoadLocation(TzIanaEuropeLondon)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaEuropeLondon). Error='%v'",
				err.Error())
	}

	t1London := time.Date(2018, time.Month(1),15,8,38,29,268154893,
		locEuropeLondon)

	locUSCentral, err := time.LoadLocation(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaUsCentral). Error='%v'", err.Error())
	}

	dTzDto, err := DateTzDto{}.New(t1London, FmtDateTimeEverything)

	if err != nil {
		fmt.Errorf("Error returned by DateTzDto{}.New(t1London, FmtDateTimeYrMDayFmtStr) " +
			" t1London='%v'  Error='%v'",
				t1London.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	locTokyo, err := time.LoadLocation(TzIanaAsiaTokyo)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaAsiaTokyo). Error='%v'", err.Error())
	}

	t4USCentral := time.Date(2018, time.Month(3),06,20,02,18,792489279,locUSCentral)

	t4Tokyo := t4USCentral.In(locTokyo)

	err = dTzDto.SetFromTimeTz(t4Tokyo, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by " +
			"dTzDto.SetFromTimeTz(t4Tokyo, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr) " +
			"t4Tokyo='%v' Error='%v' ",
				t4Tokyo.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	if !t4USCentral.Equal(dTzDto.DateTime) {
		t.Errorf("Expected dTzDto.DateTime='%v'.  Instead dTzDto.DateTime='%v'",
			t4USCentral.Format(FmtDateTimeYrMDayFmtStr), dTzDto.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	eTimeZoneDef, err := TimeZoneDefDto{}.New(t4USCentral)

	if !eTimeZoneDef.Equal(dTzDto.TimeZone) {
		t.Errorf("Expected dTzDto.TimeZone.LocationName='%v'. " +
			"Instead, dTzDto.TimeZone.LocationName='%v'",
				eTimeZoneDef.LocationName, dTzDto.TimeZone.LocationName)
	}

	tDto, err := TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral) " +
			"t4USCentral='%v' Error='%v'",
				t4USCentral.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	expectedDt, err := tDto.GetDateTime(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned from tDto.GetDateTime(TzIanaUsCentral). " +
			"Error='%v'", err.Error())
	}

	actualDt, err := dTzDto.Time.GetDateTime(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned from dTzDto.Time.GetDateTime(TzIanaUsCentral). " +
			"Error='%v'", err.Error())
	}


	if !tDto.Equal(dTzDto.Time) {
		t.Errorf("Expected dTzDto.Time (TimeDto) == '%v' Instead, dTzDto.Time (TimeDto) == '%v'",
			expectedDt.Format(FmtDateTimeYrMDayFmtStr), actualDt.Format(FmtDateTimeYrMDayFmtStr))
	}

	if FmtDateTimeYrMDayFmtStr != dTzDto.DateTimeFmt {
		t.Errorf("Expected dTzDto.DateTimeFmt='%v' Instead, dTzDto.DateTimeFmt='%v' ",
			FmtDateTimeYrMDayFmtStr, dTzDto.DateTimeFmt)
	}

}

func TestDateTzDto_SetFromTimeDto(t *testing.T) {
	
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

	t4Dto, err := TimeDto{}.New(year, month, 0, day, hour, minute,
		second, 0, 0, nSecs)

	if err != nil {
		t.Errorf("Error returned by t4USCentral TimeDto{}.New(). Error='%v'", err.Error())
	}
	
	t4TZoneDef, err := TimeZoneDefDto{}.New(t4USCentral)
	
	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(t4USCentral). Error='%v'", err.Error())
	}
	
	locTokyo, err := time.LoadLocation(TzIanaAsiaTokyo)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaAsiaTokyo). Error='%v'", err.Error())
	}

	t5Tokyo := time.Date(2012, 9, 30, 11, 58, 48, 123456789, locTokyo)
		

	t5Dto, err := TimeDto{}.New(2012, 9, 0, 30, 11,
							58, 48,  0, 0, 123456789)

	if err != nil {
		t.Errorf("Error returned by t5Tokyo TimeDto{}.New(). Error='%v'", err.Error())
	}

	t5TZoneDef, err := TimeZoneDefDto{}.New(t5Tokyo)
	
	dTz1, err := DateTzDto{}.New(t5Tokyo, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t4USCentral, FmtDateTimeYrMDayFmtStr)")
	}
	
	if !t5Dto.Equal(dTz1.Time) {
		t.Error("Expected t5Dto == dTz1.Time. It DID NOT!")
	}
	
	if !t5TZoneDef.Equal(dTz1.TimeZone) {
		t.Error("Expected t5TZoneDef == dTz1.TimeZone. It DID NOT!")
	}

	err = dTz1.SetFromTimeDto(t4Dto, TzIanaUsCentral)
	
	if err != nil {
		t.Errorf("Error returned from dTz1.SetFromTimeDto(t4Dto, TzIanaUsCentral). " +
			"Error='%v'", err.Error())
	}
	
	if !t4USCentral.Equal(dTz1.DateTime) {
		t.Errorf("Expected dTz1.DateTime='%v'.  Instead, dTz1.DateTime='%v'.",
			t4USCentral.Format(FmtDateTimeYrMDayFmtStr),
			dTz1.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}
	
	if !t4Dto.Equal(dTz1.Time) {
		t.Error("Expected t4Dto TimeDto == dTz1.Time Time Dto. THEY ARE NOT EQUAL!")
	}
	
	if !t4TZoneDef.Equal(dTz1.TimeZone) {
		t.Error("Expected t4TZoneDef TimeZoneDef == dTz1.TimeZone TimeZoneDef. " +
			"THEY ARE NOT EQUAL!")
	}
	
	if year != dTz1.Time.Years {
		t.Errorf("Error: Expected Years='%v'. Instead, Years='%v'",year, dTz1.Time.Years)
	}

	if month != dTz1.Time.Months {
		t.Errorf("Error: Expected Months='%v'. Instead, Months='%v'",month, dTz1.Time.Months)
	}

	if day != dTz1.Time.DateDays {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",day, dTz1.Time.DateDays)
	}

	if hour != dTz1.Time.Hours {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",hour, dTz1.Time.Hours)
	}

	if minute != dTz1.Time.Minutes {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",minute, dTz1.Time.Minutes)
	}

	if second != dTz1.Time.Seconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",second, dTz1.Time.Seconds)
	}

	if 792 != dTz1.Time.Milliseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",792, dTz1.Time.Milliseconds)
	}

	if 489 != dTz1.Time.Microseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",489, dTz1.Time.Microseconds)
	}

	if 279 != dTz1.Time.Nanoseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'",279, dTz1.Time.Nanoseconds)
	}

	if nSecs != dTz1.Time.TotSubSecNanoseconds {
		t.Errorf("Error: Expected dTz1.Time.TotSubSecNanoseconds='%v'. "+
			"Instead, dTz1.Time.TotSubSecNanoseconds='%v'", nSecs, dTz1.Time.TotSubSecNanoseconds)
	}

	totTime := int64(hour) * int64(time.Hour)
	totTime += int64(minute) * int64(time.Minute)
	totTime += int64(second) * int64(time.Second)
	totTime += int64(nSecs)

	if totTime != dTz1.Time.TotTimeNanoseconds {
		t.Errorf("Error: Expected tDto.TotTimeNanoseconds='%v'. "+
			"Instead, tDto.TotTimeNanoseconds='%v'", totTime, dTz1.Time.TotTimeNanoseconds)
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
