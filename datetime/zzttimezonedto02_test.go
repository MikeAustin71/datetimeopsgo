package datetime

import (
	"testing"
	"time"
)

func TestTimeZoneUtility_NewTimeAddDate_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	var utcPtr, centraTzPtr *time.Location
		var t1Utc, t2Utc, t2Central time.Time
	var tzu2 TimeZoneDto
	var err error

	utcPtr, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned by utcPtr, err = time.LoadLocation(TZones.UTC())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	centraTzPtr, err = time.LoadLocation(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by centraTzPtr, err = time.LoadLocation(TZones.US.Central())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1Utc = t1.In(utcPtr)

	t2Utc = t1Utc.AddDate(3, 2, 15)

	t2Central = t2Utc.In(centraTzPtr)

	t2OutStr := t2Central.Format(fmtstr)

	tzu2, err = TimeZoneDto{}.NewTimeAddDate(t1, TZones.US.Pacific(), 3, 2, 15, fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewTimeAddDate(t1,TzUsPacific , 3, 2, 15).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzu2OutStrTIn := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration := tzu2.TimeUTC.SubDateTime(t1Utc)

	expectedDuration := t2Utc.Sub(t1Utc)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration)
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

	tzu2, err := TimeZoneDto{}.NewTimeAddTime(t1, TZones.US.Pacific(), 3, 32, 18, 122, 58, 615, fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewTimeAddTime(t1, TzUsPacific, 3, 32, 18,122,58,615 ).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzu2OutStrTIn := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration := tzu2.TimeIn.SubDateTime(t1)

	expectedDuration := t2.Sub(t1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration)
	}

}

func TestTimeZoneUtility_NewTimeAddDateTime_01(t *testing.T) {
	// expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	t1str := "2014-02-15 19:54:30.000000000 -0600 CST"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	var err error
	var utcPtr, centralTzPtr, easternTzPtr *time.Location
	var t2Central, t2Eastern, t2Utc, t1Utc time.Time

	utcPtr, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.UTC())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	centralTzPtr, err = time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(" +
			"TZones.America.Chicago())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	easternTzPtr, err = time.LoadLocation(TZones.America.New_York())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(" +
			"TZones.America.New_York())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1Utc = t1.In(utcPtr)

	dtMech := DTimeMechanics{}

	t2Utc = dtMech.AddDateTime(
		t1Utc,
		3,
		2,
		15,
		3,
		4,
		2,
		0,
		0,
		0)

	t2Central = t2Utc.In(centralTzPtr)
	t2CentralOutStr := t2Central.Format(fmtstr)

	t2Eastern = t2Utc.In(easternTzPtr)
	t2EasternOutStr := t2Eastern.Format(fmtstr)

	t12Dur := t2Utc.Sub(t1Utc)

	tzu2, err := TimeZoneDto{}.NewTimeAddDateTime(
		t1,
		TZones.US.Eastern(),
		3,
		2,
		15,
		3,
		4,
		2,
		0,
		0,
		0,
		fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewAddDateTime(tzu1, 3,2, 15, 3, 4, 2,0, 0, 0).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzu2TimeInStr := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t2CentralOutStr != tzu2TimeInStr {
		t.Errorf("Error: Expected tzu2.TimeIn='%v'.  Instead, tzu2.TimeIn='%v'. ", t2CentralOutStr, tzu2TimeInStr)
		return
	}

	tzu2TimeOutStr := tzu2.TimeOut.GetDateTimeValue().Format(fmtstr)

	if t2EasternOutStr != tzu2TimeOutStr {
		t.Errorf("Error: Expected tzu2.TimeOut='%v'\n" +
			"Instead, tzu2.TimeOut='%v'\n",
			t2EasternOutStr, tzu2TimeOutStr)
		return
	}

	tzu2Dur := tzu2.TimeUTC.SubDateTime(t1Utc)

	if t12Dur != tzu2Dur {
		t.Errorf("Error: Expected tzu1-tzu2 Duration='%v'.  Instead, Duration='%v'", t12Dur, tzu2Dur)
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

	tzu2, err := TimeZoneDto{}.NewTimeAddDuration(t1, TZones.US.Pacific(), time.Duration(dNanSecs), fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewTimeAddTime(t1, TzUsPacific, 3, 32, 18,122,58,615 ).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tzu2OutStrTIn := tzu2.TimeIn.GetDateTimeValue().Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration := tzu2.TimeIn.SubDateTime(t1)

	expectedDuration := time.Duration(dNanSecs)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration)
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

	tOut, err := tz.ReclassifyTimeWithNewTz(
		tIn,
		TzConvertType.Relative(),
		TZones.Local())

	if err != nil {
		t.Errorf("Error returned by tz.ReclassifyTimeWithNewTz(tIn, TZones.Local())\n" +
			"Error='%v'\n", err.Error())
		return
	}

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

	tOut, err := tz.ReclassifyTimeWithNewTz(
		tIn,
		TzConvertType.Relative(),
		TZones.US.Hawaii())

	if err != nil {
		t.Errorf("Error returned by tz.ReclassifyTimeWithNewTz(tIn, TZones.US.Hawaii())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tOutLoc := tOut.Location()

	if tOutLoc.String() != TZones.US.Hawaii() {
		t.Errorf("Expected tOutLocation == '%v', instead tOutLocation == '%v'", TZones.US.Hawaii(), tOutLoc.String())
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

	tOut, err := tz.ReclassifyTimeWithNewTz(
		tIn,
		TzConvertType.Relative(),
		TZones.US.Mountain())

	if err != nil {
		t.Errorf("Error returned by tz.ReclassifyTimeWithNewTz(tIn, TZones.US.Mountain())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tOutLoc := tOut.Location()

	if tOutLoc.String() != TZones.US.Mountain() {
		t.Errorf("Expected tOutLocation == '%v', instead tOutLocation == '%v'", TZones.US.Hawaii(), tOutLoc.String())
	}

}
