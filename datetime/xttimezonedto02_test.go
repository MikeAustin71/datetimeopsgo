package datetime

import (
	"testing"
	"time"
)

func TestTimeZoneUtility_NewTimeAddDate_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2 := t1.AddDate(3,2, 15)
	t2OutStr := t2.Format(fmtstr)

	tzu2, err := TimeZoneDto{}.NewTimeAddDate(t1, TzIanaUsPacific , 3, 2, 15, fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewTimeAddDate(t1,TzUsPacific , 3, 2, 15). Error='%v'", err.Error())
	}

	tzu2OutStrTIn := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration := tzu2.TimeIn.SubDateTime(t1)

	expectedDuration := t2.Sub(t1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration )
	}

}

func TestTimeZoneUtility_NewTimeAddTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	dNanSecs := int64(0)

	dNanSecs = int64(time.Hour) * int64(3)
	dNanSecs += int64(time.Minute) * int64(32)
	dNanSecs += int64(time.Second) * int64(18)
	dNanSecs += int64(time.Millisecond) * int64(122)
	dNanSecs += int64(time.Microsecond) * int64(58)
	dNanSecs += int64(615) // Nanoseconds

	t2 := t1.Add(time.Duration(dNanSecs))
	t2OutStr := t2.Format(fmtstr)

	tzu2, err := TimeZoneDto{}.NewTimeAddTime(t1, TzIanaUsPacific, 3, 32, 18,122,58,615 , fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewTimeAddTime(t1, TzUsPacific, 3, 32, 18,122,58,615 ). Error='%v'", err.Error())
	}

	tzu2OutStrTIn := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration := tzu2.TimeIn.SubDateTime(t1)

	expectedDuration := t2.Sub(t1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration )
	}

}

func TestTimeZoneUtility_NewTimeAddDateTime_01(t *testing.T) {
	// expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	tzu2, err := TimeZoneDto{}.NewTimeAddDateTime(t1, TzIanaUsEast, 3,2, 15, 3, 4, 2,0, 0, 0, fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewAddDateTime(tzu1, 3,2, 15, 3, 4, 2,0, 0, 0). Error='%v'", err.Error())
	}

	tzu2TimeInStr := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2TimeInStr {
		t.Errorf("Error: Expected tzu2.TimeIn='%v'.  Instead, tzu2.TimeIn='%v'. ",t2OutStr, tzu2TimeInStr)
	}

	tzu2Dur:= tzu2.TimeIn.SubDateTime(t1)

	if t12Dur != tzu2Dur {
		t.Errorf("Error expected tzu1-tzu2 Duration='%v'.  Instead, Duration='%v'",t12Dur, tzu2Dur)
	}

}

func TestTimeZoneUtility_NewTimeAddDuration_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	dNanSecs := int64(0)

	dNanSecs = int64(time.Hour) * int64(3)
	dNanSecs += int64(time.Minute) * int64(32)
	dNanSecs += int64(time.Second) * int64(18)
	dNanSecs += int64(time.Millisecond) * int64(122)
	dNanSecs += int64(time.Microsecond) * int64(58)
	dNanSecs += int64(615) // Nanoseconds

	t2 := t1.Add(time.Duration(dNanSecs))
	t2OutStr := t2.Format(fmtstr)

	tzu2, err := TimeZoneDto{}.NewTimeAddDuration(t1, TzIanaUsPacific, time.Duration(dNanSecs), fmtstr )

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewTimeAddTime(t1, TzUsPacific, 3, 32, 18,122,58,615 ). Error='%v'", err.Error())
	}

	tzu2OutStrTIn := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration := tzu2.TimeIn.SubDateTime(t1)

	expectedDuration := time.Duration(dNanSecs)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration )
	}

}

func TestTimeZoneUtility_ReclassifyTimeWithTzLocal(t *testing.T) {
	/*
			Example Method: ReclassifyTimeWithNewTz()
		Input Time :  2017-04-29 17:54:30 -0700 PDT
		Output Time:  2017-04-29 17:54:30 -0500 CDT
	*/

	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tz := TimeZoneDto{}
	tIn, err := time.Parse(fmtstr, tPacific)
	if err != nil {
		t.Errorf("Error returned from time.Parse: %v", err.Error())
	}

	tOut, err := tz.ReclassifyTimeWithNewTz(tIn, "Local")

	tOutLoc := tOut.Location()

	if tOutLoc.String() != "Local" {
		t.Errorf("Expected tOutLocation == 'Local', instead go Location: '%v'", tOutLoc.String())
	}

}

func TestTimeZoneUtility_ReclassifyTimeWithNewTz(t *testing.T) {

	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tz := TimeZoneDto{}
	tIn, err := time.Parse(fmtstr, tPacific)
	if err != nil {
		t.Errorf("Error returned from time.Parse: %v", err.Error())
	}

	tOut, err := tz.ReclassifyTimeWithNewTz(tIn, TzIanaUsHawaii)

	tOutLoc := tOut.Location()

	if tOutLoc.String() != TzIanaUsHawaii {
		t.Errorf("Expected tOutLocation == '%v', instead tOutLocation == '%v'", TzIanaUsHawaii, tOutLoc.String())
	}

}

func TestTimeZoneUtility_ReclassifyTimeAsMountain(t *testing.T) {

	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tz := TimeZoneDto{}
	tIn, err := time.Parse(fmtstr, tPacific)
	if err != nil {
		t.Errorf("Error returned from time.Parse: %v", err.Error())
	}

	tOut, err := tz.ReclassifyTimeWithNewTz(tIn, TzIanaUsMountain)

	tOutLoc := tOut.Location()

	if tOutLoc.String() != TzIanaUsMountain {
		t.Errorf("Expected tOutLocation == '%v', instead tOutLocation == '%v'", TzIanaUsHawaii, tOutLoc.String())
	}

}
