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
	t1ExpectedLocationName := TZones.US.Central()

	dTzDto, err := DateTzDto{}.NewDateTimeElements(
		2014,
		2,
		15,
		19,
		54,
		30,
		158712300,
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTimeElements(2014, 2,15,19,54,30,158712300, " +
			"TzUsCentral).\nError='%v'\n", err.Error())
		return
	}

	actualTimeStr := dTzDto.GetDateTimeValue().Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.\nInstead, dTzDto.DateTime='%v'\n",
			t1OutStr, actualTimeStr)
	}

	tz := dTzDto.GetTimeZone()

	if t1ExpectedZone != tz.GetZoneName() {
		t.Errorf("Error: Expected dTzDto.TimeZone='%v'.\n" +
			"Instead, dTzDto.GetTimeZone().ZoneName='%v'\n",
			t1ExpectedZone,tz.GetZoneName())
	}

	if t1ExpectedZoneOffset != tz.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tz.GetTimeZoneOffset()='%v'.\n" +
			"Instead, tz.GetTimeZone().OffsetSeconds='%v'\n",
			t1ExpectedZoneOffset, tz.GetZoneOffsetSeconds())
	}

	if t1ExpectedLocationName != tz.GetLocationName() {
		t.Errorf("Error: Expected tz.GetLocationName()='%v'.\n" +
			"Instead, tz.LocationName='%v'\n",
			t1ExpectedLocationName, tz.GetLocationName())
	}

	if t1.Year() != dTzDto.GetTimeComponents().Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.GetTimeComponents().Years)
	}

	if int(t1.Month()) != dTzDto.GetTimeComponents().Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.GetTimeComponents().Months)
	}

	if t1.Hour() != dTzDto.GetTimeComponents().Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.GetTimeComponents().Hours)
	}

	if t1.Minute() != dTzDto.GetTimeComponents().Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.GetTimeComponents().Minutes)
	}
	if t1.Second() != dTzDto.GetTimeComponents().Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.GetTimeComponents().Seconds)
	}

	if t1.Nanosecond() != dTzDto.GetTimeComponents().TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='%v'.\n" +
			"Instead Total Nanosecond Number='%v'\n", t1.Nanosecond(), dTzDto.GetTimeComponents().TotSubSecNanoseconds)
	}

	r := t1.Nanosecond()

	if r == 0 {
		return
	}

	millisecond := r / int(time.Millisecond)

	if millisecond != dTzDto.GetTimeComponents().Milliseconds {
		t.Errorf("Expected Millisecond='%v'.  Instead, Millisecond='%v'", millisecond, dTzDto.GetTimeComponents().Milliseconds)
	}

	r -= millisecond * int(time.Millisecond)

	if r == 0 {
		return
	}

	microsecond := r / int(time.Microsecond)

	if microsecond != dTzDto.GetTimeComponents().Microseconds {
		t.Errorf("Expected Microsecond='%v'.  Instead, Microsecond='%v'", microsecond, dTzDto.GetTimeComponents().Microseconds)
	}

	r -= microsecond * int(time.Microsecond)

	if r == 0 {
		return
	}

	if r != dTzDto.GetTimeComponents().Nanoseconds {
		t.Errorf("Expected Nanosecond='%v'.  Instead Nanosecond='%v' ", r, dTzDto.GetTimeComponents().Nanoseconds)
	}

}

func TestDateTzDto_NewDateTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := TZones.US.Central()

	dTzDto, err := DateTzDto{}.NewDateTimeComponents(
		2014,
		2,
		15,
		19,
		54,
		30,
		38,
		175,
		584,
		TZones.US.Central(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTimeComponents(2014," +
			"2,15,19,54,30,38, 175, 584, TzUsCentral).\nError='%v'\n", err.Error())
		return
	}

	actualTimeStr := dTzDto.GetDateTimeValue().Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'", t1OutStr, actualTimeStr)
	}

	tz := dTzDto.GetTimeZone()

	if t1ExpectedZone != tz.GetZoneName() {
		t.Errorf("Error: Expected tz.ZoneName='%v'.\n" +
			"Instead, tz.ZoneName='%v'\n",
			t1ExpectedZone, tz.GetZoneName())
	}

	if t1ExpectedZoneOffset != tz.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tz.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tz.GetZoneOffsetSeconds()='%v'",
			t1ExpectedZoneOffset, tz.GetZoneOffsetSeconds())
	}

	if t1ExpectedLocationName != tz.GetLocationName() {
		t.Errorf("Error: Expected tz.GetLocationName()='%v'.\n" +
			"Instead, tz.GetLocationName()='%v'\n",
			t1ExpectedLocationName, tz.GetLocationName())
	}

	if t1.Year() != dTzDto.GetTimeComponents().Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.GetTimeComponents().Years)
	}

	if int(t1.Month()) != dTzDto.GetTimeComponents().Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.GetTimeComponents().Months)
	}

	if t1.Hour() != dTzDto.GetTimeComponents().Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.GetTimeComponents().Hours)
	}

	if t1.Minute() != dTzDto.GetTimeComponents().Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.GetTimeComponents().Minutes)
	}

	if t1.Second() != dTzDto.GetTimeComponents().Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.GetTimeComponents().Seconds)
	}

	r := 38 * int(time.Millisecond)

	if 38 != dTzDto.GetTimeComponents().Milliseconds {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'", dTzDto.GetTimeComponents().Milliseconds)
	}

	r += 175 * int(time.Microsecond)

	if 175 != dTzDto.GetTimeComponents().Microseconds {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.GetTimeComponents().Microseconds)
	}

	if 584 != dTzDto.GetTimeComponents().Nanoseconds {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.GetTimeComponents().Nanoseconds)
	}

	r += 584

	if r != dTzDto.GetTimeComponents().TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='%v'.\n" +
			"Instead Total Nanosecond Number='%v'", r, dTzDto.GetTimeComponents().TotSubSecNanoseconds)
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
		t.Errorf("Error returned by DateTzDto{}.New(t1).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualTimeStr := dTzDto.GetDateTimeValue().Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'", t1OutStr, actualTimeStr)
	}

	tz := dTzDto.GetTimeZone()

	if t1ExpectedZone != tz.GetZoneName() {
		t.Errorf("Error: Expected tz.GetZoneName()='%v'.\n" +
			"Instead, tz.GetZoneName()='%v'",
			t1ExpectedZone, tz.GetZoneName())
	}

	if t1ExpectedZoneOffset != tz.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tz.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tz.ZoneOffsetSeconds()='%v'",
			t1ExpectedZoneOffset, tz.GetZoneOffsetSeconds())
	}

	if t1ExpectedLocationName != tz.GetLocationName() {
		t.Errorf("Error: Expected tz.GetLocationName()='%v'.\n" +
			"Instead, tz.LocationName()='%v'",
			t1ExpectedLocationName, tz.GetLocationName())
	}

	if t1.Year() != dTzDto.GetTimeComponents().Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.GetTimeComponents().Years)
	}

	if int(t1.Month()) != dTzDto.GetTimeComponents().Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.GetTimeComponents().Months)
	}

	if t1.Hour() != dTzDto.GetTimeComponents().Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.GetTimeComponents().Hours)
	}

	if t1.Minute() != dTzDto.GetTimeComponents().Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.GetTimeComponents().Minutes)
	}

	if t1.Second() != dTzDto.GetTimeComponents().Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.GetTimeComponents().Seconds)
	}

	if 38 != dTzDto.GetTimeComponents().Milliseconds {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'", dTzDto.GetTimeComponents().Milliseconds)
	}

	if 175 != dTzDto.GetTimeComponents().Microseconds {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.GetTimeComponents().Microseconds)
	}

	if 584 != dTzDto.GetTimeComponents().Nanoseconds {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.GetTimeComponents().Nanoseconds)
	}

	if 38175584 != dTzDto.GetTimeComponents().TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='38175584'.\n" +
			"Instead Total Nanosecond Number='%v'\n", dTzDto.GetTimeComponents().TotSubSecNanoseconds)
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
		t.Errorf("Error returned by DateTzDto{}.New(t1).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualTimeStr := dTzDto.GetDateTimeValue().Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'", t1OutStr, actualTimeStr)
	}

	tz := dTzDto.GetTimeZone()

	if t1ExpectedZone != tz.GetZoneName() {
		t.Errorf("Error: Expected tz.GetZoneName()='%v'.\n" +
			"Instead, tz.GetZoneName()='%v'\n",
			t1ExpectedZone, tz.GetZoneName())
	}

	if t1ExpectedZoneOffset != tz.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tz.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tz.GetZoneOffsetSeconds()='%v'\n",
			t1ExpectedZoneOffset, tz.GetZoneOffsetSeconds())
	}

	if t1ExpectedLocationName != tz.GetLocationName() {
		t.Errorf("Error: Expected tz.GetLocationName()='%v'.\n" +
			"Instead, tz.GetLocationName()='%v'\n",
			t1ExpectedLocationName, tz.GetLocationName())
	}

	if t1.Year() != dTzDto.GetTimeComponents().Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.GetTimeComponents().Years)
	}

	if int(t1.Month()) != dTzDto.GetTimeComponents().Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.GetTimeComponents().Months)
	}

	if t1.Hour() != dTzDto.GetTimeComponents().Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.GetTimeComponents().Hours)
	}

	if t1.Minute() != dTzDto.GetTimeComponents().Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.GetTimeComponents().Minutes)
	}

	if t1.Second() != dTzDto.GetTimeComponents().Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.GetTimeComponents().Seconds)
	}

	if 38 != dTzDto.GetTimeComponents().Milliseconds {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'", dTzDto.GetTimeComponents().Milliseconds)
	}

	if 175 != dTzDto.GetTimeComponents().Microseconds {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.GetTimeComponents().Microseconds)
	}

	if 584 != dTzDto.GetTimeComponents().Nanoseconds {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.GetTimeComponents().Nanoseconds)
	}

	if 38175584 != dTzDto.GetTimeComponents().TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='38175584'.\n" +
			"Instead Total Nanosecond Number='%v'\n", dTzDto.GetTimeComponents().TotSubSecNanoseconds)
	}

}

func TestDateTzDto_NewFromMilitaryDateTz_01(t *testing.T) {

	tstr := "12/06/2019 03:12:00 -0600 CST"
	fmtStr := "01/02/2006 15:04:05 -0700 MST"

	testTime, err := time.Parse(fmtStr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtStr, tstr)\n" +
			"fmtStr='%v'\n" +
			"tstr='%v'\n" +
			"Error='%v'\n",fmtStr, tstr, err.Error())
	}

	var milTzDto MilitaryDateTzDto
	var dateTzDto DateTzDto

	milTzDto, err = MilitaryDateTzDto{}.New(testTime, "Sierra")

	if err != nil {
		t.Errorf("Error returned by MilitaryDateTzDto{}.New(testTime, \"Sierra\")\n" +
			"testTime='%v'\n" +
			"Error='%v'\n", testTime.Format(fmtStr), err.Error())
		return
	}

	dateTzDto, err = DateTzDto{}.NewFromMilitaryDateTz(milTzDto, fmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewFromMilitaryDateTz(milTzDto, fmtStr)\n" +
			"milTzDto.DateTime='%v'\n" +
			"Error='%v'\n", milTzDto.DateTime.Format(fmtStr), err.Error())
		return
	}

	if !testTime.Equal(dateTzDto.GetDateTimeValue()) {
		t.Errorf("Error: Expected dateTzDto.DateTime='%v'\n" +
			"Instead, dateTzDto.DateTime='%v'\n",
			testTime.Format(FmtDateTimeTzNanoYMDDow),
			dateTzDto.GetDateTimeValue().Format(FmtDateTimeTzNanoYMDDow))
	}
}

func TestDateTzDto_NewNowTz_01(t *testing.T) {

	t0 := time.Now().Local()

	dTz, err := DateTzDto{}.NewNowTz(TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewNowTz(TZones.US.Central(), FmtDateTimeYrMDayFmtStr). "+
			"Error='%v'", err.Error())
	}

	loc, err := time.LoadLocation(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	t1 := t0.In(loc)

	actualDur := t1.Sub(dTz.GetDateTimeValue())

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if actualDur > expectedDur {
		t.Error("Error: Actual Duration exceeded 2-seconds!")
	}

	if TZones.US.Central() != dTz.GetTimeZoneName() {
		t.Errorf("Error: Expected Time Zone='%v'.  Actual TimeZone='%v' ",
			TZones.US.Central(), dTz.GetTimeZoneName())
	}

}

func TestDateTzDto_NewNowLocal_01(t *testing.T) {

	t0 := time.Now().Local()
	dTz, err := DateTzDto{}.NewNowLocal(FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from time.NewNowLocal(FmtDateTimeYrMDayFmtStr).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	actualDur := t0.Sub(dTz.GetDateTimeValue())

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if actualDur > expectedDur {
		t.Error("Error: Actual Duration exceeded 2-seconds!")
	}

	if TZones.Local() != dTz.GetTimeZoneName() {
		t.Errorf("Error: Expected Time Zone='%v'.  Actual TimeZone='%v' ",
			TZones.Local(), dTz.GetTimeZoneName())
	}

}

func TestDateTzDto_NewNowUTC_01(t *testing.T) {

	t0 := time.Now().Local()

	loc, err := time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned from time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	t1 := t0.In(loc)

	dTz, err := DateTzDto{}.NewNowUTC(FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewNowUTC(FmtDateTimeYrMDayFmtStr). "+
			"Error='%v'", err.Error())
		return
	}

	actualDur := t1.Sub(dTz.GetDateTimeValue())

	expectedDur := time.Duration(int64(2) * int64(time.Second))

	if actualDur > expectedDur {
		t.Error("Error: Actual Duration exceeded 2-seconds!")
	}

	if TZones.UTC() != dTz.GetTimeZoneName() {
		t.Errorf("Error: Expected Time Zone='%v'.  Actual TimeZone='%v' ",
			TZones.UTC(), dTz.GetTimeZoneName())
	}

}

func TestDateTzDto_NewTimeDto_01(t *testing.T) {

	t0str := "2017-04-30 22:58:32.515539300 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t0, err := time.Parse(fmtstr, t0str)

	if err != nil {
		t.Errorf("Error retruned from time.Parse(fmtstr, t0str). t0str='%v'  Error='%v'", t0str, err.Error())
	}

	tDto, err := TimeDto{}.New(2017, 04, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned from TimeDto{}.New(...)  Error='%v'", err.Error())
	}

	dTzDto, err := DateTzDto{}.NewTimeDto(tDto, TZones.US.Central(), fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewTimeDto(tDto, TZones.US.Central(), fmtstr)\n" +
			"Error='%v\n", err.Error())
		return
	}

	if !dTzDto.GetDateTimeValue().Equal(t0) {
		t.Error("Error: Expected dTzDto.DateTime==t0.\n" +
			"Instead, they are NOT Equal!\n")
	}

	if t0str != dTzDto.String() {
		t.Errorf("Error on formats: Expected date time string='%v'. Instead, date time string='%v'", t0str, dTzDto.String())
	}

}

func TestDateTzDto_NewTz_01(t *testing.T) {

	locUSCentral, err := time.LoadLocation(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

	t4AsiaTokyo := t4USCentral.In(locTokyo)

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dTz, err := DateTzDto{}.NewTz(t4AsiaTokyo, TZones.US.Central(), fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewTz(t4AsiaTokyo, TZones.US.Central(), fmtstr).\n" +
			"Error='%v'\n",
			err.Error())
		return
	}

	if !t4USCentral.Equal(dTz.GetDateTimeValue()) {
		t.Errorf("Error: Expected DateTime='%v'. Instead DateTime='%v'",
			t4USCentral.Format(fmtstr), dTz.GetDateTimeValue().Format(fmtstr))
	}

	eTimeZoneDef, err := TimeZoneDefDto{}.New(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(t4USCentral)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !eTimeZoneDef.Equal(dTz.GetTimeZone()) {
		t.Errorf("Expected dTz.GetTimeZone().LocationName='%v'. "+
			"Instead, dTz.GetTimeZone().LocationName='%v'",
			eTimeZoneDef.GetLocationName(), dTz.GetTimeZoneName())
	}

	tDto, err := TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral)\n"+
			"t4USCentral='%v'\nError='%v'\n",
			t4USCentral.Format(FmtDateTimeYrMDayFmtStr), err.Error())
		return
	}

	expectedDt, err := tDto.GetDateTime(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from tDto.GetDateTime(TZones.US.Central()). "+
			"Error='%v'", err.Error())
	}

	timeComponents := dTz.GetTimeComponents()

	actualDt, err := timeComponents.GetDateTime(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from dTz.GetTimeComponents().GetDateTime(TZones.US.Central()).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !tDto.Equal(dTz.GetTimeComponents()) {
		t.Errorf("Expected dTz.Time (TimeDto) == '%v' Instead, dTz.Time (TimeDto) == '%v'",
			expectedDt.Format(FmtDateTimeYrMDayFmtStr), actualDt.Format(FmtDateTimeYrMDayFmtStr))
	}

	if FmtDateTimeYrMDayFmtStr != dTz.GetDateTimeFmt() {
		t.Errorf("Expected dTz.GetDateTimeFmt()='%v' Instead, dTz.GetDateTimeFmt()='%v' ",
			FmtDateTimeYrMDayFmtStr, dTz.GetDateTimeFmt())
	}

}

func TestDateTzDto_SetFromDateTimeComponents_01(t *testing.T) {
	t0str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t0, _ := time.Parse(fmtstr, t0str)
	t1, _ := time.Parse(fmtstr, t1str)

	dTzDto, err := DateTzDto{}.New(t0, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t0, FmtDateTimeYrMDayFmtStr)\n" +
			"Error='%v'\n", err.Error())

		return
	}

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := TZones.US.Central()

	err = dTzDto.SetFromDateTimeComponents(2014, 2, 15, 19, 54, 30, 38, 175, 584, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTimeComponents(2014, 2,15,19,54,30,38, 175, 584, TzUsCentral). Error='%v'", err.Error())
	}

	actualTimeStr := dTzDto.GetDateTimeValue().Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'", t1OutStr, actualTimeStr)
	}

	tz := dTzDto.GetTimeZone()

	if t1ExpectedZone != tz.GetZoneName() {
		t.Errorf("Error: Expected tz.GetZoneName()='%v'.\n" +
			"Instead, tz.GetZoneName()='%v'\n",
			t1ExpectedZone,  tz.GetZoneName())
	}

	if t1ExpectedZoneOffset != tz.GetZoneOffsetSeconds() {
		t.Errorf("Error: Expected tz.GetZoneOffsetSeconds()='%v'.\n" +
			"Instead, tz.GetZoneOffsetSeconds()='%v'\n",
			t1ExpectedZoneOffset, tz.GetZoneOffsetSeconds())
	}

	if t1ExpectedLocationName != dTzDto.GetTimeZoneName() {
		t.Errorf("Error: Expected dTzDto.GetTimeZoneName()='%v'.\n" +
			"Instead, dTzDto.GetTimeZoneName()='%v'",
			t1ExpectedLocationName, dTzDto.GetTimeZoneName())
	}

	if t1.Year() != dTzDto.GetTimeComponents().Years {
		t.Errorf("Expected Year='%v'.  Instead Year='%v'", t1.Year(), dTzDto.GetTimeComponents().Years)
	}

	if int(t1.Month()) != dTzDto.GetTimeComponents().Months {
		t.Errorf("Expected Month Number='%v'.  Instead Month Number='%v'", int(t1.Month()), dTzDto.GetTimeComponents().Months)
	}

	if t1.Hour() != dTzDto.GetTimeComponents().Hours {
		t.Errorf("Expected Hour Number='%v'.  Instead Hour Number='%v'", t1.Hour(), dTzDto.GetTimeComponents().Hours)
	}

	if t1.Minute() != dTzDto.GetTimeComponents().Minutes {
		t.Errorf("Expected Minute Number='%v'.  Instead Minute Number='%v'", t1.Minute(), dTzDto.GetTimeComponents().Minutes)
	}

	if t1.Second() != dTzDto.GetTimeComponents().Seconds {
		t.Errorf("Expected Second Number='%v'.  Instead Second Number='%v'", t1.Second(), dTzDto.GetTimeComponents().Seconds)
	}

	r := 38 * int(time.Millisecond)

	if 38 != dTzDto.GetTimeComponents().Milliseconds {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'", dTzDto.GetTimeComponents().Milliseconds)
	}

	r += 175 * int(time.Microsecond)

	if 175 != dTzDto.GetTimeComponents().Microseconds {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.GetTimeComponents().Microseconds)
	}

	if 584 != dTzDto.GetTimeComponents().Nanoseconds {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.GetTimeComponents().Nanoseconds)
	}

	r += 584

	if r != dTzDto.GetTimeComponents().TotSubSecNanoseconds {
		t.Errorf("Expected Total Nanosecond Number='%v'.\n" +
			"Instead Total Nanosecond Number='%v'", r, dTzDto.GetTimeComponents().TotSubSecNanoseconds)
	}

}

func TestDateTzDto_SetNewTimeZone_01(t *testing.T) {

	locUSCentral, err := time.LoadLocation(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

	t4AsiaTokyo := t4USCentral.In(locTokyo)

	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	dTz1, err := DateTzDto{}.New(t4USCentral, fmtstr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t4USCentral, fmtstr)\n" +
			"Error='%v\n", err.Error())
		return
	}

	err = dTz1.SetNewTimeZone(TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by dTz1.SetNewTimeZone(TZones.Asia.Tokyo()). "+
			"TZones.Asia.Tokyo()='%v' Error='%v' ", TZones.Asia.Tokyo(), err.Error())
	}

	if !t4AsiaTokyo.Equal(dTz1.GetDateTimeValue()) {
		t.Errorf("Error: Expected converted dTz1 date time = '%v'.  Instead, dTz1 date time='%v'",
			t4AsiaTokyo.Format(FmtDateTimeYrMDayFmtStr),
			dTz1.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	if TZones.Asia.Tokyo() != dTz1.GetTimeZoneName() {
		t.Errorf("Error: Expected dTz1 Time Zone Location Name ='%v'. "+
			"Instead, Time Zone Location Name='%v'", TZones.Asia.Tokyo(), dTz1.GetTimeZoneName())
	}

}

func TestDateTzDto_SetFromTimeTz_01(t *testing.T) {

	locEuropeLondon, err := time.LoadLocation(TZones.Europe.London())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Europe.London()). Error='%v'",
			err.Error())
	}

	t1London := time.Date(2018, time.Month(1), 15, 8, 38, 29, 268154893,
		locEuropeLondon)

	locUSCentral, err := time.LoadLocation(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.US.Central()). Error='%v'", err.Error())
	}

	dTzDto, err := DateTzDto{}.New(t1London, FmtDateTimeEverything)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t1London, FmtDateTimeYrMDayFmtStr) "+
			" t1London='%v'  Error='%v'",
			t1London.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	locTokyo, err := time.LoadLocation(TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
	}

	t4USCentral := time.Date(2018, time.Month(3), 06, 20, 02, 18, 792489279, locUSCentral)

	t4Tokyo := t4USCentral.In(locTokyo)

	err = dTzDto.SetFromTimeTz(t4Tokyo, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by "+
			"dTzDto.SetFromTimeTz(t4Tokyo, TZones.US.Central(), FmtDateTimeYrMDayFmtStr) "+
			"t4Tokyo='%v' Error='%v' ",
			t4Tokyo.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	if !t4USCentral.Equal(dTzDto.GetDateTimeValue()) {
		t.Errorf("Expected dTzDto.DateTime='%v'.  Instead dTzDto.DateTime='%v'",
			t4USCentral.Format(FmtDateTimeYrMDayFmtStr),
			dTzDto.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	eTimeZoneDef, err := TimeZoneDefDto{}.New(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(t4USCentral)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !eTimeZoneDef.Equal(dTzDto.GetTimeZone()) {
		t.Errorf("Expected dTzDto.GetTimeZone().LocationName='%v'.\n"+
			"Instead, dTzDto.GetTimeZone().LocationName='%v'\n",
			eTimeZoneDef.GetLocationName(), dTzDto.GetTimeZoneName())
	}

	tDto, err := TimeDto{}.NewFromDateTime(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewFromDateTime(t4USCentral) "+
			"t4USCentral='%v' Error='%v'",
			t4USCentral.Format(FmtDateTimeYrMDayFmtStr), err.Error())

		return
	}

	expectedDt, err := tDto.GetDateTime(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from tDto.GetDateTime(TZones.US.Central()). "+
			"Error='%v'", err.Error())

		return
	}

	timeComponents := dTzDto.GetTimeComponents()

	actualDt, err := timeComponents.GetDateTime(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from dTzDto.GetTimeComponents().GetDateTime(TZones.US.Central()).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if !tDto.Equal(dTzDto.GetTimeComponents()) {
		t.Errorf("Expected dTzDto.Time (TimeDto) == '%v' Instead, dTzDto.Time (TimeDto) == '%v'",
			expectedDt.Format(FmtDateTimeYrMDayFmtStr), actualDt.Format(FmtDateTimeYrMDayFmtStr))
	}

	if FmtDateTimeYrMDayFmtStr != dTzDto.GetDateTimeFmt() {
		t.Errorf("Expected dTzDto.GetDateTimeFmt()='%v' Instead, dTzDto.GetDateTimeFmt()='%v' ",
			FmtDateTimeYrMDayFmtStr, dTzDto.GetDateTimeFmt())
	}

}

func TestDateTzDto_SetFromTimeDto(t *testing.T) {

	locUSCentral, err := time.LoadLocation(TZones.US.Central())

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

	t4Dto, err := TimeDto{}.New(year, month, 0, day, hour, minute,
		second, 0, 0, nSecs)

	if err != nil {
		t.Errorf("Error returned by t4USCentral TimeDto{}.New(). Error='%v'", err.Error())
	}

	t4TZoneDef, err := TimeZoneDefDto{}.New(t4USCentral)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDefDto{}.New(t4USCentral). Error='%v'", err.Error())
	}

	locTokyo, err := time.LoadLocation(TZones.Asia.Tokyo())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.Asia.Tokyo()). Error='%v'", err.Error())
	}

	t5Tokyo := time.Date(2012, 9, 30, 11, 58, 48, 123456789, locTokyo)

	t5Dto, err := TimeDto{}.New(2012, 9, 0, 30, 11,
		58, 48, 0, 0, 123456789)

	if err != nil {
		t.Errorf("Error returned by t5Tokyo TimeDto{}.New(). Error='%v'", err.Error())
	}

	t5TZoneDef, err := TimeZoneDefDto{}.New(t5Tokyo)

	dTz1, err := DateTzDto{}.New(t5Tokyo, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.New(t4USCentral, FmtDateTimeYrMDayFmtStr)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !t5Dto.Equal(dTz1.GetTimeComponents()) {
		t.Error("Expected t5Dto == dTz1.GetTimeComponents(). It DID NOT!")
	}

	if !t5TZoneDef.Equal(dTz1.GetTimeZone()) {
		t.Error("Expected t5TZoneDef == dTz1.GetTimeZone(). It DID NOT!")
	}

	err = dTz1.SetFromTimeDto(t4Dto, TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned from dTz1.SetFromTimeDto(t4Dto, TZones.US.Central()). "+
			"Error='%v'", err.Error())
	}

	if !t4USCentral.Equal(dTz1.GetDateTimeValue()) {
		t.Errorf("Expected dTz1.DateTime='%v'.  Instead, dTz1.DateTime='%v'.",
			t4USCentral.Format(FmtDateTimeYrMDayFmtStr),
			dTz1.GetDateTimeValue().Format(FmtDateTimeYrMDayFmtStr))
	}

	if !t4Dto.Equal(dTz1.GetTimeComponents()) {
		t.Error("Expected t4Dto TimeDto == dTz1.Time Time Dto. THEY ARE NOT EQUAL!")
	}

	if !t4TZoneDef.Equal(dTz1.GetTimeZone()) {
		t.Error("Expected t4TZoneDef TimeZoneDef == dTz1.TimeZone TimeZoneDef. " +
			"THEY ARE NOT EQUAL!")
	}

	if year != dTz1.GetTimeComponents().Years {
		t.Errorf("Error: Expected Years='%v'. Instead, Years='%v'", year, dTz1.GetTimeComponents().Years)
	}

	if month != dTz1.GetTimeComponents().Months {
		t.Errorf("Error: Expected Months='%v'. Instead, Months='%v'", month, dTz1.GetTimeComponents().Months)
	}

	if day != dTz1.GetTimeComponents().DateDays {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", day, dTz1.GetTimeComponents().DateDays)
	}

	if hour != dTz1.GetTimeComponents().Hours {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", hour, dTz1.GetTimeComponents().Hours)
	}

	if minute != dTz1.GetTimeComponents().Minutes {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", minute, dTz1.GetTimeComponents().Minutes)
	}

	if second != dTz1.GetTimeComponents().Seconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", second, dTz1.GetTimeComponents().Seconds)
	}

	if 792 != dTz1.GetTimeComponents().Milliseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", 792, dTz1.GetTimeComponents().Milliseconds)
	}

	if 489 != dTz1.GetTimeComponents().Microseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", 489, dTz1.GetTimeComponents().Microseconds)
	}

	if 279 != dTz1.GetTimeComponents().Nanoseconds {
		t.Errorf("Error: Expected Days='%v'. Instead, Days='%v'", 279, dTz1.GetTimeComponents().Nanoseconds)
	}

	if nSecs != dTz1.GetTimeComponents().TotSubSecNanoseconds {
		t.Errorf("Error: Expected dTz1.GetTimeComponents().TotSubSecNanoseconds='%v'. "+
			"Instead, dTz1.GetTimeComponents().TotSubSecNanoseconds='%v'", nSecs, dTz1.GetTimeComponents().TotSubSecNanoseconds)
	}

	totTime := int64(hour) * int64(time.Hour)
	totTime += int64(minute) * int64(time.Minute)
	totTime += int64(second) * int64(time.Second)
	totTime += int64(nSecs)

	if totTime != dTz1.GetTimeComponents().TotTimeNanoseconds {
		t.Errorf("Error: Expected tDto.TotTimeNanoseconds='%v'. "+
			"Instead, tDto.TotTimeNanoseconds='%v'", totTime, dTz1.GetTimeComponents().TotTimeNanoseconds)
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
		t.Errorf("Error returned from DateTzDto{}.New(t2, fmtstr).\nError='%v'\n", err.Error())
		return
	}

	actualDuration := dTz2.Sub(dTz1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: expected duration='%v'.  Instead, duration='%v' ", expectedDuration, actualDuration)
	}

}
