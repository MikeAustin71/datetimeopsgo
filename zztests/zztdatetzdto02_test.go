package zztests

import (
	"github.com/MikeAustin71/datetimeopsgo/datetime"
	"testing"
	"time"
)

func TestDateTzDto_NewDateTimeElements_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.158712300 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := datetime.TZones.US.Central()

	dTzDto, err := datetime.DateTzDto{}.NewDateTimeElements(
		2014,
		2,
		15,
		19,
		54,
		30,
		158712300,
		datetime.TZones.US.Central(),
		datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTimeElements(2014, 2,15,19,54,30,158712300, " +
			"TzUsCentral).\nError='%v'\n", err.Error())
		return
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.\nInstead, dTzDto.DateTime='%v'\n",
			t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone.ZoneName {
		t.Errorf("Error: Expected dTzDto.TimeZone='%v'. Instead, dTzDto.TimeZone.ZoneName='%v'", t1ExpectedZone, dTzDto.TimeZone.ZoneName)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZone.ZoneOffsetSeconds {
		t.Errorf("Error: Expected dTzDto.TimeZoneOffset='%v'.  Instead, dTzDto.TimeZone.OffsetSeconds='%v'", t1ExpectedZoneOffset, dTzDto.TimeZone.ZoneOffsetSeconds)
	}

	if t1ExpectedLocationName != dTzDto.TimeZone.LocationName {
		t.Errorf("Error: Expected dTzDto.TimeLocName='%v'.  Instead, dTzDto.TimeZone.LocationName='%v'", t1ExpectedLocationName, dTzDto.TimeZone.LocationName)
	}

	if t1.Year() != dTzDto.Time.Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != dTzDto.Time.Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Hour() != dTzDto.Time.Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != dTzDto.Time.Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.Time.Minutes)
	}
	if t1.Second() != dTzDto.Time.Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.Time.Seconds)
	}

	if t1.Nanosecond() != dTzDto.Time.TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='%v'.\n" +
			"Instead Total Nanosecond Number='%v'\n", t1.Nanosecond(), dTzDto.Time.TotSubSecNanoseconds)
	}

	r := t1.Nanosecond()

	if r == 0 {
		return
	}

	millisecond := r / int(time.Millisecond)

	if millisecond != dTzDto.Time.Milliseconds {
		t.Errorf("Expected Millisecond='%v'.  Instead, Millisecond='%v'", millisecond, dTzDto.Time.Milliseconds)
	}

	r -= millisecond * int(time.Millisecond)

	if r == 0 {
		return
	}

	microsecond := r / int(time.Microsecond)

	if microsecond != dTzDto.Time.Microseconds {
		t.Errorf("Expected Microsecond='%v'.  Instead, Microsecond='%v'", microsecond, dTzDto.Time.Microseconds)
	}

	r -= microsecond * int(time.Microsecond)

	if r == 0 {
		return
	}

	if r != dTzDto.Time.Nanoseconds {
		t.Errorf("Expected Nanosecond='%v'.  Instead Nanosecond='%v' ", r, dTzDto.Time.Nanoseconds)
	}

}

func TestDateTzDto_NewDateTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := datetime.TZones.US.Central()

	dTzDto, err := datetime.DateTzDto{}.NewDateTime(
		2014,
		2,
		15,
		19,
		54,
		30,
		38,
		175,
		584,
		datetime.TZones.US.Central(),
		datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(2014," +
			"2,15,19,54,30,38, 175, 584, TzUsCentral).\nError='%v'\n", err.Error())
		return
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'", t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone.ZoneName {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneName='%v'. Instead, dTzDto.TimeZone.ZoneName='%v'", t1ExpectedZone, dTzDto.TimeZone.ZoneName)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZone.ZoneOffsetSeconds {
		t.Errorf("Error: Expected dTzDto.TimeZoneOffset.ZoneOffsetSeconds='%v'.  Instead, dTzDto.TimeZoneOffset.ZoneOffsetSeconds='%v'", t1ExpectedZoneOffset, dTzDto.TimeZone.ZoneOffsetSeconds)
	}

	if t1ExpectedLocationName != dTzDto.TimeZone.LocationName {
		t.Errorf("Error: Expected dTzDto.TimeZone.LocationName='%v'.  Instead, dTzDto.TimeZone.LocationName='%v'", t1ExpectedLocationName, dTzDto.TimeZone.LocationName)
	}

	if t1.Year() != dTzDto.Time.Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != dTzDto.Time.Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Hour() != dTzDto.Time.Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != dTzDto.Time.Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.Time.Minutes)
	}

	if t1.Second() != dTzDto.Time.Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.Time.Seconds)
	}

	r := 38 * int(time.Millisecond)

	if 38 != dTzDto.Time.Milliseconds {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'", dTzDto.Time.Milliseconds)
	}

	r += 175 * int(time.Microsecond)

	if 175 != dTzDto.Time.Microseconds {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Time.Microseconds)
	}

	if 584 != dTzDto.Time.Nanoseconds {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Time.Nanoseconds)
	}

	r += 584

	if r != dTzDto.Time.TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='%v'.\n" +
			"Instead Total Nanosecond Number='%v'", r, dTzDto.Time.TotSubSecNanoseconds)
	}

}

func TestDateTzDto_New_01(t *testing.T) {

	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := t1.Location().String()

	dTzDto, err := datetime.DateTzDto{}.New(t1, datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'", t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone.ZoneName {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneName='%v'. Instead, dTzDto.TimeZone.ZoneName='%v'", t1ExpectedZone, dTzDto.TimeZone.ZoneName)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZone.ZoneOffsetSeconds {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneOffsetSeconds='%v'.  Instead, dTzDto.TimeZone.ZoneOffsetSeconds='%v'", t1ExpectedZoneOffset, dTzDto.TimeZone.ZoneOffsetSeconds)
	}

	if t1ExpectedLocationName != dTzDto.TimeZone.LocationName {
		t.Errorf("Error: Expected dTzDto.TimeZone.LocationName='%v'.  Instead, dTzDto.TimeZone.LocationName='%v'", t1ExpectedLocationName, dTzDto.TimeZone.LocationName)
	}

	if t1.Year() != dTzDto.Time.Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != dTzDto.Time.Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Hour() != dTzDto.Time.Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != dTzDto.Time.Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.Time.Minutes)
	}

	if t1.Second() != dTzDto.Time.Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.Time.Seconds)
	}

	if 38 != dTzDto.Time.Milliseconds {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'", dTzDto.Time.Milliseconds)
	}

	if 175 != dTzDto.Time.Microseconds {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Time.Microseconds)
	}

	if 584 != dTzDto.Time.Nanoseconds {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Time.Nanoseconds)
	}

	if 38175584 != dTzDto.Time.TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='38175584'.\n" +
			"Instead Total Nanosecond Number='%v'\n", dTzDto.Time.TotSubSecNanoseconds)
	}

}

func TestDateTzDto_New_02(t *testing.T) {

	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := t1.Location().String()

	dTzDto, err := datetime.DateTzDto{}.New(t1, "")

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'", t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone.ZoneName {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneName='%v'. Instead, dTzDto.TimeZone.ZoneName='%v'", t1ExpectedZone, dTzDto.TimeZone.ZoneName)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZone.ZoneOffsetSeconds {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneOffsetSeconds='%v'.  Instead, dTzDto.TimeZone.ZoneOffsetSeconds='%v'", t1ExpectedZoneOffset, dTzDto.TimeZone.ZoneOffsetSeconds)
	}

	if t1ExpectedLocationName != dTzDto.TimeZone.LocationName {
		t.Errorf("Error: Expected dTzDto.TimeZone.LocationName='%v'.  Instead, dTzDto.TimeZone.LocationName='%v'", t1ExpectedLocationName, dTzDto.TimeZone.LocationName)
	}

	if t1.Year() != dTzDto.Time.Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != dTzDto.Time.Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Hour() != dTzDto.Time.Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != dTzDto.Time.Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.Time.Minutes)
	}

	if t1.Second() != dTzDto.Time.Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.Time.Seconds)
	}

	if 38 != dTzDto.Time.Milliseconds {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'", dTzDto.Time.Milliseconds)
	}

	if 175 != dTzDto.Time.Microseconds {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Time.Microseconds)
	}

	if 584 != dTzDto.Time.Nanoseconds {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Time.Nanoseconds)
	}

	if 38175584 != dTzDto.Time.TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='38175584'.\n" +
			"Instead Total Nanosecond Number='%v'\n", dTzDto.Time.TotSubSecNanoseconds)
	}

}

func TestDateTzDto_NewNowTz_01(t *testing.T) {

	t0 := time.Now().Local()

	dTz, err := datetime.DateTzDto{}.NewNowTz(datetime.TZones.US.Central(), datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewNowTz(TZones.US.Central(), FmtDateTimeYrMDayFmtStr). "+
			"Error='%v'", err.Error())
	}

	loc, err := time.LoadLocation(datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	t1 := t0.In(loc)

	actualDur := t1.Sub(dTz.DateTime)

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if actualDur > expectedDur {
		t.Error("Error: Actual Duration exceeded 2-seconds!")
	}

	if datetime.TZones.US.Central() != dTz.TimeZone.LocationName {
		t.Errorf("Error: Expected Time Zone='%v'.  Actual TimeZone='%v' ",
			datetime.TZones.US.Central(), dTz.TimeZone.LocationName)
	}

}

func TestDateTzDto_NewNowLocal_01(t *testing.T) {

	t0 := time.Now().Local()
	dTz, err := datetime.DateTzDto{}.NewNowLocal(datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from time.NewNowLocal(FmtDateTimeYrMDayFmtStr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	actualDur := t0.Sub(dTz.DateTime)

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if actualDur > expectedDur {
		t.Error("Error: Actual Duration exceeded 2-seconds!")
	}

	if datetime.TZones.Local() != dTz.TimeZone.LocationName {
		t.Errorf("Error: Expected Time Zone='%v'.  Actual TimeZone='%v' ",
			datetime.TZones.Local(), dTz.TimeZone.LocationName)
	}

}

func TestDateTzDto_NewNowUTC_01(t *testing.T) {

	t0 := time.Now().Local()

	loc, err := time.LoadLocation(datetime.TZones.UTC())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	t1 := t0.In(loc)

	dTz, err := datetime.DateTzDto{}.NewNowUTC(datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewNowUTC(FmtDateTimeYrMDayFmtStr). "+
			"Error='%v'", err.Error())
		return
	}

	actualDur := t1.Sub(dTz.DateTime)

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if actualDur > expectedDur {
		t.Error("Error: Actual Duration exceeded 2-seconds!")
	}

	if datetime.TZones.UTC() != dTz.TimeZone.LocationName {
		t.Errorf("Error: Expected Time Zone='%v'.  Actual TimeZone='%v' ",
			datetime.TZones.UTC(), dTz.TimeZone.LocationName)
	}

}

func TestDateTzDto_NewTimeDto_01(t *testing.T) {

	t0str := "2017-04-30 22:58:32.515539300 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t0, err := time.Parse(fmtstr, t0str)

	if err != nil {
		t.Errorf("Error retruned from time.Parse(fmtstr, t0str). t0str='%v'  Error='%v'", t0str, err.Error())
	}

	tDto, err := datetime.TimeDto{}.New(2017, 04, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned from TimeDto{}.New(...)  Error='%v'", err.Error())
	}

	dTzDto, err := datetime.DateTzDto{}.NewTimeDto(tDto, datetime.TZones.US.Central(), fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewTimeDto(tDto, TZones.US.Central(), fmtstr)\n" +
			"Error='%v\n", err.Error())
		return
	}

	if !dTzDto.DateTime.Equal(t0) {
		t.Error("Error: Expected dTzDto.DateTime==t0.\n" +
			"Instead, they are NOT Equal!\n")
	}

	if t0str != dTzDto.String() {
		t.Errorf("Error on formats: Expected date time string='%v'. Instead, date time string='%v'", t0str, dTzDto.String())
	}

}

func TestDateTzDto_NewTz_01(t *testing.T) {

	locUSCentral, err := time.LoadLocation(datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(datetime.TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

	t4AsiaTokyo := t4USCentral.In(locTokyo)

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dTz, err := datetime.DateTzDto{}.NewTz(t4AsiaTokyo, datetime.TZones.US.Central(), fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewTz(t4AsiaTokyo, TZones.US.Central(), fmtstr).\n" +
			"Error='%v'\n",
			err.Error())
		return
	}

	if !t4USCentral.Equal(dTz.DateTime) {
		t.Errorf("Error: Expected DateTime='%v'. Instead DateTime='%v'",
			t4USCentral.Format(fmtstr), dTz.DateTime.Format(fmtstr))
	}

	eTimeZoneDef, err := datetime.TimeZoneDefDto{}.New(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(t4USCentral)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !eTimeZoneDef.Equal(dTz.TimeZone) {
		t.Errorf("Expected dTz.TimeZone.LocationName='%v'. "+
			"Instead, dTz.TimeZone.LocationName='%v'",
			eTimeZoneDef.LocationName, dTz.TimeZone.LocationName)
	}

	tDto, err := datetime.TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral)\n"+
			"t4USCentral='%v'\nError='%v'\n",
			t4USCentral.Format(datetime.FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	expectedDt, err := tDto.GetDateTime(datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from tDto.GetDateTime(TZones.US.Central()). "+
			"Error='%v'", err.Error())
	}

	actualDt, err := dTz.Time.GetDateTime(datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from dTz.Time.GetDateTime(TZones.US.Central()).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !tDto.Equal(dTz.Time) {
		t.Errorf("Expected dTz.Time (TimeDto) == '%v' Instead, dTz.Time (TimeDto) == '%v'",
			expectedDt.Format(datetime.FmtDateTimeYrMDayFmtStr), actualDt.Format(datetime.FmtDateTimeYrMDayFmtStr))
	}

	if datetime.FmtDateTimeYrMDayFmtStr != dTz.DateTimeFmt {
		t.Errorf("Expected dTz.DateTimeFmt='%v' Instead, dTz.DateTimeFmt='%v' ",
			datetime.FmtDateTimeYrMDayFmtStr, dTz.DateTimeFmt)
	}

}

func TestDateTzDto_SetFromDateTime_01(t *testing.T) {
	t0str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t0, _ := time.Parse(fmtstr, t0str)
	t1, _ := time.Parse(fmtstr, t1str)

	dTzDto, err := datetime.DateTzDto{}.New(t0, datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t0, FmtDateTimeYrMDayFmtStr)\n" +
			"Error='%v'\n", err.Error())

		return
	}

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := datetime.TZones.US.Central()

	err = dTzDto.SetFromDateTime(2014, 2, 15, 19, 54, 30, 38, 175, 584, datetime.TZones.US.Central(), datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(2014, 2,15,19,54,30,38, 175, 584, TzUsCentral). Error='%v'", err.Error())
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'", t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone.ZoneName {
		t.Errorf("Error: Expected dTzDto.TimeZone='%v'. Instead, dTzDto.TimeZone='%v'", t1ExpectedZone, dTzDto.TimeZone)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZone.ZoneOffsetSeconds {
		t.Errorf("Error: Expected dTzDto.TimeZone.ZoneOffsetSeconds='%v'.  Instead, dTzDto.TimeZone.ZoneOffsetSeconds='%v'", t1ExpectedZoneOffset, dTzDto.TimeZone.ZoneOffsetSeconds)
	}

	if t1ExpectedLocationName != dTzDto.TimeZone.LocationName {
		t.Errorf("Error: Expected dTzDto.TimeZone.LocationName='%v'.  Instead, dTzDto.TimeZone.LocationName='%v'", t1ExpectedLocationName, dTzDto.TimeZone.LocationName)
	}

	if t1.Year() != dTzDto.Time.Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.Time.Years)
	}

	if int(t1.Month()) != dTzDto.Time.Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.Time.Months)
	}

	if t1.Hour() != dTzDto.Time.Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.Time.Hours)
	}

	if t1.Minute() != dTzDto.Time.Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.Time.Minutes)
	}

	if t1.Second() != dTzDto.Time.Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.Time.Seconds)
	}

	r := 38 * int(time.Millisecond)

	if 38 != dTzDto.Time.Milliseconds {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'", dTzDto.Time.Milliseconds)
	}

	r += 175 * int(time.Microsecond)

	if 175 != dTzDto.Time.Microseconds {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Time.Microseconds)
	}

	if 584 != dTzDto.Time.Nanoseconds {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Time.Nanoseconds)
	}

	r += 584

	if r != dTzDto.Time.TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='%v'.\n" +
			"Instead Total Nanosecond Number='%v'", r, dTzDto.Time.TotSubSecNanoseconds)
	}

}

func TestDateTzDto_SetNewTimeZone_01(t *testing.T) {

	locUSCentral, err := time.LoadLocation(datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(datetime.TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

	t4AsiaTokyo := t4USCentral.In(locTokyo)

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dTz1, err := datetime.DateTzDto{}.New(t4USCentral, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t4USCentral, fmtstr)\n" +
			"Error='%v\n", err.Error())
		return
	}

	err = dTz1.SetNewTimeZone(datetime.TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by dTz1.SetNewTimeZone(TZones.Asia.Tokyo()). "+
			"TZones.Asia.Tokyo()='%v' Error='%v' ", datetime.TZones.Asia.Tokyo(), err.Error())
	}

	if !t4AsiaTokyo.Equal(dTz1.DateTime) {
		t.Errorf("Error: Expected converted dTz1 date time = '%v'.  Instead, dTz1 date time='%v'",
			t4AsiaTokyo.Format(datetime.FmtDateTimeYrMDayFmtStr), dTz1.DateTime.Format(datetime.FmtDateTimeYrMDayFmtStr))
	}

	if datetime.TZones.Asia.Tokyo() != dTz1.TimeZone.LocationName {
		t.Errorf("Error: Expected dTz1 Time Zone Location Name ='%v'. "+
			"Instead, Time Zone Location Name='%v'", datetime.TZones.Asia.Tokyo(), dTz1.TimeZone.LocationName)
	}

}

func TestDateTzDto_SetFromTimeTz_01(t *testing.T) {

	locEuropeLondon, err := time.LoadLocation(datetime.TZones.Europe.London())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Europe.London()). Error='%v'",
			err.Error())
	}

	t1London := time.Date(2018, time.Month(1), 15, 8, 38, 29, 268154893,
		locEuropeLondon)

	locUSCentral, err := time.LoadLocation(datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	dTzDto, err := datetime.DateTzDto{}.New(t1London, datetime.FmtDateTimeEverything)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1London, FmtDateTimeYrMDayFmtStr) "+
			" t1London='%v'  Error='%v'",
			t1London.Format(datetime.FmtDateTimeYrMDayFmtStr), err.Error())
	}

	locTokyo, err := time.LoadLocation(datetime.TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

	t4Tokyo := t4USCentral.In(locTokyo)

	err = dTzDto.SetFromTimeTz(t4Tokyo, datetime.TZones.US.Central(), datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by "+
			"dTzDto.SetFromTimeTz(t4Tokyo, TZones.US.Central(), FmtDateTimeYrMDayFmtStr) "+
			"t4Tokyo='%v' Error='%v' ",
			t4Tokyo.Format(datetime.FmtDateTimeYrMDayFmtStr), err.Error())
	}

	if !t4USCentral.Equal(dTzDto.DateTime) {
		t.Errorf("Expected dTzDto.DateTime='%v'.  Instead dTzDto.DateTime='%v'",
			t4USCentral.Format(datetime.FmtDateTimeYrMDayFmtStr), dTzDto.DateTime.Format(datetime.FmtDateTimeYrMDayFmtStr))
	}

	eTimeZoneDef, err := datetime.TimeZoneDefDto{}.New(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(t4USCentral)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !eTimeZoneDef.Equal(dTzDto.TimeZone) {
		t.Errorf("Expected dTzDto.TimeZone.LocationName='%v'.\n"+
			"Instead, dTzDto.TimeZone.LocationName='%v'\n",
			eTimeZoneDef.LocationName, dTzDto.TimeZone.LocationName)
	}

	tDto, err := datetime.TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral) "+
			"t4USCentral='%v' Error='%v'",
			t4USCentral.Format(datetime.FmtDateTimeYrMDayFmtStr), err.Error())

		return
	}

	expectedDt, err := tDto.GetDateTime(datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from tDto.GetDateTime(TZones.US.Central()). "+
			"Error='%v'", err.Error())

		return
	}

	actualDt, err := dTzDto.Time.GetDateTime(datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from dTzDto.Time.GetDateTime(TZones.US.Central()).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !tDto.Equal(dTzDto.Time) {
		t.Errorf("Expected dTzDto.Time (TimeDto) == '%v' Instead, dTzDto.Time (TimeDto) == '%v'",
			expectedDt.Format(datetime.FmtDateTimeYrMDayFmtStr), actualDt.Format(datetime.FmtDateTimeYrMDayFmtStr))
	}

	if datetime.FmtDateTimeYrMDayFmtStr != dTzDto.DateTimeFmt {
		t.Errorf("Expected dTzDto.DateTimeFmt='%v' Instead, dTzDto.DateTimeFmt='%v' ",
			datetime.FmtDateTimeYrMDayFmtStr, dTzDto.DateTimeFmt)
	}

}

func TestDateTzDto_SetFromTimeDto(t *testing.T) {

	locUSCentral, err := time.LoadLocation(datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	year := 2018
	month := 3
	day := 6
	hour := 20
	minute := 2
	second := 18
	nSecs := 792489279

	t4USCentral := time.Date(year, time.Month(month), day, hour, minute, second, nSecs, locUSCentral)

	t4Dto, err := datetime.TimeDto{}.New(year, month, 0, day, hour, minute,
		second, 0, 0, nSecs)

	if err != nil {
		t.Errorf("Error returned by t4USCentral TimeDto{}.New(). Error='%v'", err.Error())
	}

	t4TZoneDef, err := datetime.TimeZoneDefDto{}.New(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(t4USCentral). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(datetime.TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
	}

	t5Tokyo := time.Date(2012, 9, 30, 11, 58, 48, 123456789, locTokyo)

	t5Dto, err := datetime.TimeDto{}.New(2012, 9, 0, 30, 11,
		58, 48, 0, 0, 123456789)

	if err != nil {
		t.Errorf("Error returned by t5Tokyo TimeDto{}.New(). Error='%v'", err.Error())
	}

	t5TZoneDef, err := datetime.TimeZoneDefDto{}.New(t5Tokyo)

	dTz1, err := datetime.DateTzDto{}.New(t5Tokyo, datetime.FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t4USCentral, FmtDateTimeYrMDayFmtStr)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !t5Dto.Equal(dTz1.Time) {
		t.Error("Expected t5Dto == dTz1.Time. It DID NOT!")
	}

	if !t5TZoneDef.Equal(dTz1.TimeZone) {
		t.Error("Expected t5TZoneDef == dTz1.TimeZone. It DID NOT!")
	}

	err = dTz1.SetFromTimeDto(t4Dto, datetime.TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from dTz1.SetFromTimeDto(t4Dto, TZones.US.Central()). "+
			"Error='%v'", err.Error())
	}

	if !t4USCentral.Equal(dTz1.DateTime) {
		t.Errorf("Expected dTz1.DateTime='%v'.  Instead, dTz1.DateTime='%v'.",
			t4USCentral.Format(datetime.FmtDateTimeYrMDayFmtStr),
			dTz1.DateTime.Format(datetime.FmtDateTimeYrMDayFmtStr))
	}

	if !t4Dto.Equal(dTz1.Time) {
		t.Error("Expected t4Dto TimeDto == dTz1.Time Time Dto. THEY ARE NOT EQUAL!")
	}

	if !t4TZoneDef.Equal(dTz1.TimeZone) {
		t.Error("Expected t4TZoneDef TimeZoneDef == dTz1.TimeZone TimeZoneDef. " +
			"THEY ARE NOT EQUAL!")
	}

	if year != dTz1.Time.Years {
		t.Errorf("Error: Expected Years='%v'. Instead, Years='%v'", year, dTz1.Time.Years)
	}

	if month != dTz1.Time.Months {
		t.Errorf("Error: Expected Months='%v'. Instead, Months='%v'", month, dTz1.Time.Months)
	}

	if day != dTz1.Time.DateDays {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", day, dTz1.Time.DateDays)
	}

	if hour != dTz1.Time.Hours {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", hour, dTz1.Time.Hours)
	}

	if minute != dTz1.Time.Minutes {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", minute, dTz1.Time.Minutes)
	}

	if second != dTz1.Time.Seconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", second, dTz1.Time.Seconds)
	}

	if 792 != dTz1.Time.Milliseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", 792, dTz1.Time.Milliseconds)
	}

	if 489 != dTz1.Time.Microseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", 489, dTz1.Time.Microseconds)
	}

	if 279 != dTz1.Time.Nanoseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", 279, dTz1.Time.Nanoseconds)
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

	dTz1, err := datetime.DateTzDto{}.New(t1, fmtstr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.New(t1, fmtstr). Error='%v'", err.Error())
	}

	dTz2, err := datetime.DateTzDto{}.New(t2, fmtstr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.New(t2, fmtstr).\nError='%v'\n", err.Error())
		return
	}

	actualDuration := dTz2.Sub(dTz1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: expected duration='%v'.  Instead, duration='%v' ", expectedDuration, actualDuration)
	}

}
