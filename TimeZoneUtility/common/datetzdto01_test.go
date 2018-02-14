package common

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
	t1ExpectedLocationName := TzUsCentral

	dTzDto, err := DateTzDto{}.NewDateTimeElements(2014, 2,15,19,54,30,158712300, TzUsCentral)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTimeElements(2014, 2,15,19,54,30,158712300, TzUsCentral). Error='%v'", err.Error())
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'",t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone {
		t.Errorf("Error: Expected dTzDto.TimeZone='%v'. Instead, dTzDto.TimeZone='%v'",t1ExpectedZone, dTzDto.TimeZone)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZoneOffset {
		t.Errorf("Error: Expected dTzDto.TimeZoneOffset='%v'.  Instead, dTzDto.TimeZoneOffset='%v'", t1ExpectedZoneOffset, dTzDto.TimeZoneOffset)
	}

	if t1ExpectedLocationName != dTzDto.TimeLocName {
		t.Errorf("Error: Expected dTzDto.TimeLocName='%v'.  Instead, dTzDto.TimeLocName='%v'", t1ExpectedLocationName, dTzDto.TimeLocName)
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

	if t1.Nanosecond() != int(dTzDto.TotalNanoSecs) {
		t.Errorf("Expected Total Nanosecond Number='%v'.  Instead Total Nanosecond Number='%v'",t1.Nanosecond(), int(dTzDto.TotalNanoSecs))
	}

	r := t1.Nanosecond()

	if r == 0 {
		return
	}

	millisecond := r / int(time.Millisecond)

	if millisecond != dTzDto.Millisecond {
		t.Errorf("Expected Millisecond='%v'.  Instead, Millisecond='%v'", millisecond, dTzDto.Millisecond )
	}

	r -= millisecond * int(time.Millisecond)

	if r == 0 {
		return
	}

	microsecond := r / int(time.Microsecond)

	if microsecond != dTzDto.Microsecond {
		t.Errorf("Expected Microsecond='%v'.  Instead, Microsecond='%v'", microsecond, dTzDto.Microsecond)
	}

	r-= microsecond * int(time.Microsecond)

	if r == 0 {
		return
	}

	if r != dTzDto.Nanosecond {
		t.Errorf("Expected Nanosecond='%v'.  Instead Nanosecond='%v' ", r, dTzDto.Nanosecond)
	}

}

func TestDateTzDto_NewDateTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.038175584 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t1OutStr := t1.Format(fmtstr)

	t1ExpectedZone, t1ExpectedZoneOffset := t1.Zone()
	t1ExpectedLocationName := TzUsCentral

	dTzDto, err := DateTzDto{}.NewDateTime(2014, 2,15,19,54,30,38, 175, 584, TzUsCentral)

	if err != nil {
		t.Errorf("Error returned by DateTzDto{}.NewDateTime(2014, 2,15,19,54,30,38, 175, 584, TzUsCentral). Error='%v'", err.Error())
	}

	actualTimeStr := dTzDto.DateTime.Format(fmtstr)

	if t1OutStr != actualTimeStr {
		t.Errorf("Error: expected dTzDto.DateTime='%v'.  Instead, dTzDto.DateTime='%v'",t1OutStr, actualTimeStr)
	}

	if t1ExpectedZone != dTzDto.TimeZone {
		t.Errorf("Error: Expected dTzDto.TimeZone='%v'. Instead, dTzDto.TimeZone='%v'",t1ExpectedZone, dTzDto.TimeZone)
	}

	if t1ExpectedZoneOffset != dTzDto.TimeZoneOffset {
		t.Errorf("Error: Expected dTzDto.TimeZoneOffset='%v'.  Instead, dTzDto.TimeZoneOffset='%v'", t1ExpectedZoneOffset, dTzDto.TimeZoneOffset)
	}

	if t1ExpectedLocationName != dTzDto.TimeLocName {
		t.Errorf("Error: Expected dTzDto.TimeLocName='%v'.  Instead, dTzDto.TimeLocName='%v'", t1ExpectedLocationName, dTzDto.TimeLocName)
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

	r := 38 * int(time.Millisecond)

	if 38 != dTzDto.Millisecond {
		t.Errorf("Expected Millisecond='38'.  Instead, Millisecond='%v'",  dTzDto.Millisecond )
	}

	r += 175 * int(time.Microsecond)

	if 175 != dTzDto.Microsecond {
		t.Errorf("Expected Microsecond='175'.  Instead, Microsecond='%v'", dTzDto.Microsecond)
	}

	if 584 != dTzDto.Nanosecond {
		t.Errorf("Expected Nanosecond='584'.  Instead Nanosecond='%v' ", dTzDto.Nanosecond)
	}

	r += 584

	if r != int(dTzDto.TotalNanoSecs) {
		t.Errorf("Expected Total Nanosecond Number='%v'.  Instead Total Nanosecond Number='%v'",r, int(dTzDto.TotalNanoSecs))
	}

}