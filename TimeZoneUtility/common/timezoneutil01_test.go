package common

import (
	"testing"
	"time"
)

func TestTimeZoneUtility_ConvertTz_01(t *testing.T) {
	utcTime := "2017-04-30 00:54:30 +0000 UTC"
	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	mountainTime := "2017-04-29 18:54:30 -0600 MDT"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	ianaMountainTz := "America/Denver"
	tPacificIn, err := time.Parse(fmtstr, pacificTime)

	if err != nil {
		t.Errorf("Received error from time parse tPacificIn: %v", err.Error())
	}

	tzu := TimeZoneUtility{}
	tzuCentral, err := tzu.ConvertTz(tPacificIn, ianaCentralTz)

	if err != nil {
		t.Errorf("Error from TimeZoneUtility.ConvertTz(). Error: %v", err.Error())
	}

	centralTOut := tzuCentral.TimeOut.Format(fmtstr)

	if centralTime != centralTOut {
		t.Errorf("Expected tzuCentral.TimeOut %v, got %v", centralTime, centralTOut)
	}

	tzuMountain, err := tzu.ConvertTz(tzuCentral.TimeOut, ianaMountainTz)

	if err != nil {
		t.Errorf("Error from  tzuMountain TimeZoneUtility.ConvertTz(). Error: %v", err.Error())
	}

	mountainTOut := tzuMountain.TimeOut.Format(fmtstr)

	if mountainTime != mountainTOut {
		t.Errorf("Expected tzuMountain.TimeOut %v, got %v", mountainTime, mountainTOut)
	}

	tzuPacific, err := tzu.ConvertTz(tzuMountain.TimeOut, ianaPacificTz)

	if err != nil {
		t.Errorf("Error from  tzuMountain TimeZoneUtility.ConvertTz(). Error: %v", err.Error())
	}

	pacificTOut := tzuPacific.TimeOut.Format(fmtstr)

	if pacificTime != pacificTOut {

		t.Errorf("Expected tzuPacific.TimeOut %v, got %v", pacificTime, pacificTOut)
	}

	exTOutLoc := "America/Los_Angeles"

	if exTOutLoc != tzuPacific.TimeOutLoc.String() {
		t.Errorf("Expected tzu.TimeOutLoc %v, got %v", exTOutLoc, tzuPacific.TimeOutLoc.String())
	}

	pacificUtcOut := tzuPacific.TimeUTC.Format(fmtstr)

	if utcTime != pacificUtcOut {
		t.Errorf("Expected tzuPacific.TimeUTC %v, got %v", utcTime, pacificUtcOut)
	}

	centralUtcOut := tzuCentral.TimeUTC.Format(fmtstr)

	if utcTime != centralUtcOut {
		t.Errorf("Expected tzuCentral.TimeUTC %v, got %v", utcTime, pacificUtcOut)
	}

	mountainUtcOut := tzuMountain.TimeUTC.Format(fmtstr)

	if utcTime != mountainUtcOut {
		t.Errorf("Expected tzuMountain.TimeUTC %v, got %v", utcTime, pacificUtcOut)
	}

}

func TestTimeZoneUtility_ConvertTz_02(t *testing.T) {
	moscowTz := "Europe/Moscow"
	beijingTz := "Asia/Shanghai"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	moscowTime := "2017-04-30 03:54:30 +0300 MSK"
	utcTime := "2017-04-30 00:54:30 +0000 UTC"
	beijingTime := "2017-04-30 08:54:30 +0800 CST"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"

	tCentral, err := time.Parse(fmtstr, centralTime)

	if err != nil {
		t.Errorf("Error from time.Parse. centralTime = %v. Error= %v", centralTime, err.Error())
	}

	tzuMoscow, err := TimeZoneUtility{}.ConvertTz(tCentral, moscowTz)

	if err != nil {
		t.Errorf("Error from TimeZoneUtility{}.ConvertTz. Central Time = %v. Error= %v", centralTime, err.Error())
	}

	moscowTOut := tzuMoscow.TimeOut.Format(fmtstr)

	if moscowTime != moscowTOut {
		t.Errorf("Error. Moscow Time zone conversion failed! Expected %v. Instead, got %v.", moscowTime, moscowTOut)
	}

	tzuBeijing, err := tzuMoscow.ConvertTz(tCentral, beijingTz)

	if err != nil {
		t.Errorf("Error from tzuMoscow.ConvertTz. Central Time = %v. Error= %v", centralTime, err.Error())
	}

	beijingTOut := tzuBeijing.TimeOut.Format(fmtstr)

	if beijingTime != beijingTOut {
		t.Errorf("Error. Beijing Time zone conversion failed! Expected %v. Instead, got %v.", beijingTime, beijingTOut)
	}

	utcTOut := tzuBeijing.TimeUTC.Format(fmtstr)

	if utcTime != utcTOut {
		t.Errorf("Error. UTC Time from tzuBeijing.TimeUTC failed! Expected %v. Instead, got %v.", utcTime, utcTOut)

	}

}

func TestTimeZoneUtility_ConvertTz_03(t *testing.T) {
	t1UTCStr := "2017-07-02 22:00:18.423111300 +0000 UTC"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t2LocalStr := "2017-07-02 17:00:18.423111300 -0500 CDT"
	localTzStr := "America/Chicago"

	t1, _ := time.Parse(fmtstr, t1UTCStr)

	tz := TimeZoneUtility{}

	tzLocal, _ := tz.ConvertTz(t1, localTzStr)
	t1OutStr := tzLocal.TimeIn.Format(fmtstr)
	t2OutStr := tzLocal.TimeOut.Format(fmtstr)

	if t1UTCStr != t1OutStr {
		t.Errorf("Expected Input Time: %v. Error - Instead, got %v", t1UTCStr, t1OutStr)
	}

	if t2LocalStr != t2OutStr {
		t.Errorf("Expected Output Local Time: %v. Error - Instead, got %v", t2LocalStr, t2OutStr)
	}

}

func TestTimeZoneUtility_ConvertTz_04(t *testing.T) {
	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	ianaCentralTz := "America/Chicago"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"

	tPacific, err := time.Parse(fmtstr, pacificTime)

	if err != nil {
		t.Errorf("Error from time.Parse. pacificTime = %v. Error= %v", pacificTime, err.Error())
	}

	tzuCentral, err := TimeZoneUtility{}.ConvertTz(tPacific, ianaCentralTz)

	tOutStr := tzuCentral.TimeOut.Format(fmtstr)

	if centralTime != tOutStr {
		t.Errorf("Error. Central Time zone conversion failed! Expected %v. Instead, got %v.", centralTime, tOutStr)
	}

	tzLocal, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error received from time.LoadLocation(\"Local\") Error='%v'", err.Error())
	}

	expectedLocalTime :=  tPacific.In(tzLocal).Format(fmtstr)

	actualLocalTime := tzuCentral.TimeLocal.Format(fmtstr)

	if expectedLocalTime != actualLocalTime {
		t.Errorf("Error: Expected Local Time='%v'.  Actual Local Time='%v'", expectedLocalTime, actualLocalTime)
	}

}


func TestTimeZoneUtility_ConvertTz_05(t *testing.T) {

	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	ianaCentralTz := "America/Chicago"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"

	tPacific, err := time.Parse(fmtstr, pacificTime)

	if err != nil {
		t.Errorf("Error from time.Parse. pacificTime = %v. Error= %v", pacificTime, err.Error())
	}

	tzuCentral, err := TimeZoneUtility{}.ConvertTz(tPacific, ianaCentralTz)

	tOutStr := tzuCentral.TimeOut.Format(fmtstr)

	if centralTime != tOutStr {
		t.Errorf("Error. Central Time zone conversion failed! Expected %v. Instead, got %v.", centralTime, tOutStr)
	}

}

func TestInvalidTargetTzInConversion(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	// Invalid Target Iana Time Zone
	invalidTz := "XUZ Time Zone"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu := TimeZoneUtility{}
	_, err := tzu.ConvertTz(tIn, invalidTz)

	if err == nil {
		t.Error("ConverTz() failed to detect INVALID Tartet Time Zone. Got: ", "err==nil")
	}

}

func TestTimeZoneUtility_GetLocationIn01(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _:= TimeZoneUtility{}.New(tIn, ianaPacificTz)

	expectedLocation:="Local"

	actualLocation, err := tzu.GetLocationIn()

	if err != nil {
		t.Errorf("Error returned from tzu.GetLocationIn(). Error='%v'", err.Error())
	}

	if expectedLocation != actualLocation {
		t.Errorf("Expected Location='%v'. Instead, actual location='%v'",expectedLocation, actualLocation)
	}

}

func TestTimeZoneUtility_GetLocationOut_01(t *testing.T) {

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _:= TimeZoneUtility{}.New(tIn, ianaPacificTz)

	expectedLocation:="America/Los_Angeles"

	actualLocation, err := tzu.GetLocationOut()

	if err != nil {
		t.Errorf("Error returned from tzu.GetLocationOut(). Error='%v'", err.Error())
	}

	if expectedLocation != actualLocation {
		t.Errorf("Expected Location Out='%v'. Instead, actual location out='%v'",expectedLocation, actualLocation)
	}
}

func TestTimeZoneUtility_GetZoneIn_01(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _:= TimeZoneUtility{}.New(tIn, ianaPacificTz)

	expectedZone := "CDT"
	actualZone, err := tzu.GetZoneIn()

	if err != nil {
		t.Errorf("Error returned from tzu.GetZoneIn(). Error='%v'", err.Error())
	}

	if expectedZone != actualZone {
		t.Errorf("Expected Zone In='%v'. Instead, actual Zone In='%v'",expectedZone, actualZone)
	}

}

func TestTimeZoneUtility_GetZoneOut_01(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _:= TimeZoneUtility{}.New(tIn, ianaPacificTz)

	expectedZone := "PDT"
	actualZone, err := tzu.GetZoneOut()

	if err != nil {
		t.Errorf("Error returned from tzu.GetZoneOut(). Error='%v'", err.Error())
	}

	if expectedZone != actualZone {
		t.Errorf("Expected Zone Out='%v'. Instead, actual Zone Out='%v'",expectedZone, actualZone)
	}

}

func TestTimeZoneUtility_IsValidTimeZone_01(t *testing.T) {
	tIn := time.Now()

	tzu := TimeZoneUtility{}

	isValidTz, isValidIanaTz, isValidLocalTz := tzu.IsValidTimeZone(tIn.Location().String())

	if isValidTz == false {
		t.Error("Expected Now() Location to yield 'Local' Time Zone isValidTz == VALID ('true'), instead got: ", isValidTz)
	}

	if isValidIanaTz == true {
		t.Error("Passed Time Zone was 'Local' Time Zone. Expected isValidIanaTz == false, got: ", isValidIanaTz)
	}

	if isValidLocalTz == false {
		t.Error("Passed Time Zone was 'Local' Time Zone. Expected isValidLocalTz == true, got: ", isValidIanaTz)
	}

}

func TestTimeZoneUtility_IsValidTimeZone02(t *testing.T) {

	tzu := TimeZoneUtility{}

	isValidTz, isValidIanaTz, isValidLocalTz := tzu.IsValidTimeZone("America/Chicago")

	if isValidTz == false {
		t.Error("Expected 'America/Chicago' to yield isValidTz = 'true', instead got", isValidTz)
	}

	if isValidIanaTz == false {
		t.Error("Expected 'America/Chicago' to yield isValidIanaTz = 'true', instead got", isValidIanaTz)
	}

	if isValidLocalTz == true {
		t.Error("Expected 'America/Chicago' to yield isValidLocalTz = 'false', instead got", isValidLocalTz)
	}

}

func TestTimeZoneUtility_Location_01(t *testing.T) {
	utcTime := "2017-04-30 00:54:30 +0000 UTC"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"

	tUtc, _ := time.Parse(fmtstr, utcTime)

	tzu := TimeZoneUtility{}

	tzuPacific, err := tzu.ConvertTz(tUtc, ianaPacificTz)

	if err != nil {
		t.Errorf("Error from TimeZoneUtility{}.ConvertTz. Utc Time = %v. Error= %v", utcTime, err.Error())
	}

	tzOutPacific := tzuPacific.TimeOutLoc.String()

	if tzOutPacific != ianaPacificTz {
		t.Errorf("Error: Expected tzOutPacific %v. Instead, got %v", ianaPacificTz, tzOutPacific)
	}

}

func TestTimeZoneUtility_Location_02(t *testing.T) {

	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tPacific, _ := time.Parse(fmtstr, pacificTime)

	tzu := TimeZoneUtility{}

	tzuLocal, err := tzu.ConvertTz(tPacific, "Local")

	if err != nil {
		t.Errorf("Error from TimeZoneUtility{}.ConvertTz. Pacific Time = %v. Error= %v", pacificTime, err.Error())
	}

	tzOutLocal := tzuLocal.TimeOutLoc.String()

	if "Local" != tzOutLocal {
		t.Errorf("Error: Expected tzOutLocal 'Local'. Instead, got %v", tzOutLocal)
	}

}

func TestTimeZoneUtility_MakeDateTz(t *testing.T) {
	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"

	dtTzDto := DateTzDto{Year: 2017, Month: 4, Day: 29, Hour: 17, Minute: 54, Second: 30, TimeZone: "America/Los_Angeles"}

	tzu := TimeZoneUtility{}

	tOut, err := tzu.MakeDateTz(dtTzDto)

	if err != nil {
		t.Errorf("Error returned from TimeZoneUtility.MakeDateTz(). Error: %v", err.Error())
	}

	tOutStr := tOut.Format(fmtstr)

	if tPacific != tOutStr {
		t.Errorf("Error - Expected output time string: %v. Instead, got %v.", tPacific, tOutStr)
	}

}

func TestTimeZoneUtility_New_01(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	//tz := TimeZoneUtility{}
	tzu, _:= TimeZoneUtility{}.New(tIn, ianaPacificTz)
	expectedTimeIn := "2017-04-29 19:54:30 -0500 CDT"
	if expectedTimeIn != tzu.TimeIn.String() {
		t.Errorf("Expected Time In='%v'. Instead, Time In='%v'", expectedTimeIn, tzu.TimeIn.String() )
	}

	expectedTimeOut := "2017-04-29 17:54:30 -0700 PDT"

	if expectedTimeOut != tzu.TimeOut.String() {
		t.Errorf("Expected Time Out='%v'. Instead, Time Out='%v'", expectedTimeOut, tzu.TimeOut.String())
	}

}

func TestTimeZoneUtility_NewAddDuration_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := TimeZoneUtility{}.New(t1, TzUsPacific )

	if err != nil {
		t.Errorf("Error returned from TimeZoneUtility{}.New(t1, TzUsPacific ). Error='%v'",err.Error())
	}

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t12Dur := t2.Sub(t1)

	tzu2, err := TimeZoneUtility{}.NewAddDuration(tzu1, t12Dur)

	if err != nil {
		t.Errorf("Error returned by TimeZoneUtility{}.NewAddDuration(tzu1, t12Dur). Error='%v'", err.Error())
	}

	tzu1OutStr := tzu1.TimeIn.Format(fmtstr)

	if t1OutStr != tzu1OutStr {
		t.Errorf("Error: Expected Time1 TimeIn='%v'.  Instead Time1 TimeIn='%v'",t1OutStr, tzu1OutStr)
	}

	tzu2OutStr := tzu2.TimeIn.Format(fmtstr)

	if t2OutStr != tzu2OutStr {
		t.Errorf("Error: Expected after duration tzu2TimeIn='%v'.  Instead, tzu2TimeIn='%v'",t2OutStr, tzu2OutStr)
	}

	actualDur := tzu2.TimeIn.Sub(tzu1.TimeIn)

	if t12Dur != actualDur {
		t.Errorf("Error: Expected tzu2.TimeIn.Sub(tzu1.TimeIn)='%v'.  Instead, duration='%v'", t12Dur, actualDur)
	}

	actualDur = tzu2.TimeOut.Sub(tzu1.TimeOut)

	if t12Dur != actualDur {
		t.Errorf("Error: Expected tzu2.TimeOut.Sub(tzu1.TimeOut)='%v'.  Instead, duration='%v'", t12Dur, actualDur)
	}

	actualDur = tzu2.TimeUTC.Sub(tzu1.TimeUTC)

	if t12Dur != actualDur {
		t.Errorf("Error: Expected tzu2.TimeUTC.Sub(tzu1.TimeUTC)='%v'.  Instead, duration='%v'", t12Dur, actualDur)
	}

	actualDur = tzu2.TimeLocal.Sub(tzu1.TimeLocal)

	if t12Dur != actualDur {
		t.Errorf("Error: Expected tzu2.TimeLocal.Sub(tzu1.TimeLocal)='%v'.  Instead, duration='%v'", t12Dur, actualDur)
	}

	actualTimeOutLoc := tzu1.TimeOutLoc.String()

	if TzUsPacific != actualTimeOutLoc {
		t.Errorf("Error: Expected tzu1.TimeOutLoc='%v'.  Instead, tzu1.TimeOutLoc='%v'.",TzUsPacific, actualTimeOutLoc)
	}

	actualTimeOutLoc = tzu2.TimeOutLoc.String()

	if TzUsPacific != actualTimeOutLoc {
		t.Errorf("Error: Expected tzu2.TimeOutLoc.String()='%v'.  Instead, tzu2.TimeOutLoc='%v'.",TzUsPacific, actualTimeOutLoc)
	}

}

func TestTimeZoneUtility_NewAddDate_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := TimeZoneUtility{}.New(t1, TzUsPacific)
	if err != nil {
		t.Errorf("Error returned by TimeZoneUtility{}.New(t1, TzUsPacific). Error='%v'", err.Error())
	}

	tzu1OutStrTIn := tzu1.TimeIn.Format(fmtstr)

	if t1OutStr != tzu1OutStrTIn {
		t.Errorf("Error: Expected tzu1OutStrTIn='%v'.  Instead, tzu1OutStrTIn='%v'", t1OutStr, tzu1OutStrTIn )
	}

	t2 := t1.AddDate(3,2, 15)
	t2OutStr := t2.Format(fmtstr)

	tzu2, err := TimeZoneUtility{}.NewAddDate(tzu1, 3, 2, 15)

	tzu2OutStrTIn := tzu2.TimeIn.Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration, err := tzu2.Sub(tzu1)

	if err != nil {
		t.Errorf("Error returned by tzu2.Sub(tzu1). Error='%v'", err.Error())
	}

	expectedDuration := t2.Sub(t1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration )
	}
}


func TestTimeZoneUtility_NewAddDateTime_01(t *testing.T) {
	// expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	tzu1, err := TimeZoneUtility{}.New(t1, TzUsEast)

	if err != nil {
		t.Errorf("Error returned by TimeZoneUtility{}.New(t1, TzUsEast). Error='%v'", err.Error())
	}

	tzu2, err := TimeZoneUtility{}.NewAddDateTime(tzu1, 3,2, 15, 3, 4, 2,0, 0, 0)

	if err != nil {
		t.Errorf("Error returned by TimeZoneUtility{}.NewAddDateTime(tzu1, 3,2, 15, 3, 4, 2,0, 0, 0). Error='%v'", err.Error())
	}

	tzu2TimeInStr := tzu2.TimeIn.Format(fmtstr)

	if t2OutStr != tzu2TimeInStr {
		t.Errorf("Error: Expected tzu2.TimeIn='%v'.  Instead, tzu2.TimeIn='%v'. ",t2OutStr, tzu2TimeInStr)
	}

	tzu2Dur, err := tzu2.Sub(tzu1)

	if err != nil {
		t.Errorf("Error returned by tzu2.Sub(tzu1)")
	}

	if t12Dur != tzu2Dur {
		t.Errorf("Error expected tzu1-tzu2 Duration='%v'.  Instead, Duration='%v'",t12Dur, tzu2Dur)
	}

}

func TestTimeZoneUtility_NewAddTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := TimeZoneUtility{}.New(t1, TzUsPacific)
	if err != nil {
		t.Errorf("Error returned by TimeZoneUtility{}.New(t1, TzUsPacific). Error='%v'", err.Error())
	}

	tzu1OutStrTIn := tzu1.TimeIn.Format(fmtstr)

	if t1OutStr != tzu1OutStrTIn {
		t.Errorf("Error: Expected tzu1OutStrTIn='%v'.  Instead, tzu1OutStrTIn='%v'", t1OutStr, tzu1OutStrTIn )
	}

	dNanSecs := int64(0)

	dNanSecs = int64(time.Hour) * int64(3)
	dNanSecs += int64(time.Minute) * int64(32)
	dNanSecs += int64(time.Second) * int64(18)
	dNanSecs += int64(time.Millisecond) * int64(122)
	dNanSecs += int64(time.Microsecond) * int64(58)
	dNanSecs += int64(615) // Nanoseconds

	t2 := t1.Add(time.Duration(dNanSecs))
	t2OutStr := t2.Format(fmtstr)

	tzu2, err := TimeZoneUtility{}.NewAddTime(tzu1, 3, 32, 18,122,58,615 )

	tzu2OutStrTIn := tzu2.TimeIn.Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration, err := tzu2.Sub(tzu1)

	if err != nil {
		t.Errorf("Error returned by tzu2.Sub(tzu1). Error='%v'", err.Error())
	}

	expectedDuration := t2.Sub(t1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration )
	}
}


