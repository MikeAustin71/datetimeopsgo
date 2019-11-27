package datetime

import (
	"testing"
	"time"
)

func TestTimeZoneDto_AddTimeDurationDto_01(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := TimeZoneDto{}.New(t1, TZones.US.Pacific(), fmtstr)

	if err != nil {
		t.Errorf("Error returned from TimeZoneDto{}.New(t1, TzUsPacific ). Error='%v'", err.Error())
	}

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t12Dur := t2.Sub(t1)

	tdurDto, err := TimeDurationDto{}.NewStartEndTimesCalcTz(
		t1,
		t2,
		TDurCalcType(0).StdYearMth(),
		TZones.US.Central(),
		fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartEndTimesCalcTz() "+
			"Error='%v' ", err.Error())
	}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddTimeDurationDto(tdurDto)

	if err != nil {
		t.Errorf("Error returned by tzu2.AddTimeDurationDto(tdurDto). "+
			"Error='%v' ", err.Error())
	}

	tzu1OutStr := tzu1.TimeIn.DateTime.Format(fmtstr)

	if t1OutStr != tzu1OutStr {
		t.Errorf("Error: Expected Time1 TimeIn='%v'.  Instead Time1 TimeIn='%v'", t1OutStr, tzu1OutStr)
	}

	tzu2OutStr := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2OutStr {
		t.Errorf("Error: Expected after duration tzu2TimeIn='%v'.  Instead, tzu2TimeIn='%v'", t2OutStr, tzu2OutStr)
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

	actualTimeOutLoc := tzu1.TimeOut.TimeZone.LocationName

	if TZones.US.Pacific() != actualTimeOutLoc {
		t.Errorf("Error: Expected tzu1.TimeOutLoc='%v'.  Instead, tzu1.TimeOutLoc='%v'.", TZones.US.Pacific(), actualTimeOutLoc)
	}

	actualTimeOutLoc = tzu2.TimeOut.TimeZone.LocationName

	if TZones.US.Pacific() != actualTimeOutLoc {
		t.Errorf("Error: Expected tzu2.TimeOutLoc.String()='%v'.  Instead, tzu2.TimeOutLoc='%v'.", TZones.US.Pacific(), actualTimeOutLoc)
	}

}

func TestTimeZoneDto_AddMinusTimeDto(t *testing.T) {
	// expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)

	t12Dur := t2.Sub(t1)

	tzu1, err := TimeZoneDto{}.New(t2, TZones.US.Eastern(), fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.New(t1, TzUsEast). Error='%v'", err.Error())
	}

	tDto := TimeDto{Years: 3, Months: 2, DateDays: 15, Hours: 3, Minutes: 4, Seconds: 2}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddMinusTimeDto(tDto)

	if err != nil {
		t.Errorf("Error returned by tzu2.AddMinusTimeDto(tDto). "+
			"Error='%v'", err.Error())
	}

	tzu2TimeInStr := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t1OutStr != tzu2TimeInStr {
		t.Errorf("Error: Expected tzu2.TimeIn='%v'.  Instead, tzu2.TimeIn='%v'. ", t1OutStr, tzu2TimeInStr)
	}

	tzu2Dur, err := tzu1.Sub(tzu2)

	if err != nil {
		t.Errorf("Error returned by tzu2.Sub(tzu1)")
	}

	if t12Dur != tzu2Dur {
		t.Errorf("Error expected tzu1-tzu2 Duration='%v'.  Instead, Duration='%v'", t12Dur, tzu2Dur)
	}

}

func TestTimeZoneDto_AddPlusTimeDto(t *testing.T) {

	// expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	tzu1, err := TimeZoneDto{}.New(t1, TZones.US.Eastern(), fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.New(t1, TzUsEast). Error='%v'", err.Error())
	}

	tDto := TimeDto{Years: 3, Months: 2, DateDays: 15, Hours: 3, Minutes: 4, Seconds: 2}

	tzu2 := tzu1.CopyOut()

	err = tzu2.AddPlusTimeDto(tDto)

	if err != nil {
		t.Errorf("Error returned by tzu2.AddPlusTimeDto(tDto). "+
			"Error='%v' ", err.Error())
	}

	tzu2TimeInStr := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2TimeInStr {
		t.Errorf("Error: Expected tzu2.TimeIn='%v'.  Instead, tzu2.TimeIn='%v'. ", t2OutStr, tzu2TimeInStr)
	}

	tzu2Dur, err := tzu2.Sub(tzu1)

	if err != nil {
		t.Errorf("Error returned by tzu2.Sub(tzu1)")
	}

	if t12Dur != tzu2Dur {
		t.Errorf("Error expected tzu1-tzu2 Duration='%v'.  Instead, Duration='%v'", t12Dur, tzu2Dur)
	}

}

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

	tzu := TimeZoneDto{}
	tzuCentral, err := tzu.ConvertTz(tPacificIn, ianaCentralTz, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error from TimeZoneDto.ConvertTz(). Error: %v", err.Error())
	}

	centralTOut := tzuCentral.TimeOut.DateTime.Format(fmtstr)

	if centralTime != centralTOut {
		t.Errorf("Expected tzuCentral.TimeOut %v, got %v", centralTime, centralTOut)
	}

	tzuMountain, err := tzu.ConvertTz(tzuCentral.TimeOut.DateTime, ianaMountainTz, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error from  tzuMountain TimeZoneDto.ConvertTz(). Error: %v", err.Error())
	}

	mountainTOut := tzuMountain.TimeOut.DateTime.Format(fmtstr)

	if mountainTime != mountainTOut {
		t.Errorf("Expected tzuMountain.TimeOut %v, got %v", mountainTime, mountainTOut)
	}

	tzuPacific, err := tzu.ConvertTz(tzuMountain.TimeOut.DateTime, ianaPacificTz, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error from  tzuMountain TimeZoneDto.ConvertTz(). Error: %v", err.Error())
	}

	pacificTOut := tzuPacific.TimeOut.DateTime.Format(fmtstr)

	if pacificTime != pacificTOut {

		t.Errorf("Expected tzuPacific.TimeOut %v, got %v", pacificTime, pacificTOut)
	}

	exTOutLoc := "America/Los_Angeles"

	if exTOutLoc != tzuPacific.TimeOut.TimeZone.LocationName {
		t.Errorf("Expected tzu.TimeOutLoc %v, got %v.  tzuPacific.TimeOut='%v'", exTOutLoc, tzuPacific.TimeOut.TimeZone.LocationName, tzuPacific.TimeOut.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	pacificUtcOut := tzuPacific.TimeUTC.DateTime.Format(fmtstr)

	if utcTime != pacificUtcOut {
		t.Errorf("Expected tzuPacific.TimeUTC %v, got %v", utcTime, pacificUtcOut)
	}

	centralUtcOut := tzuCentral.TimeUTC.DateTime.Format(fmtstr)

	if utcTime != centralUtcOut {
		t.Errorf("Expected tzuCentral.TimeUTC %v, got %v", utcTime, pacificUtcOut)
	}

	mountainUtcOut := tzuMountain.TimeUTC.DateTime.Format(fmtstr)

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

	tzuMoscow, err := TimeZoneDto{}.ConvertTz(tCentral, moscowTz, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error from TimeZoneDto{}.ConvertTz. Central Time = %v. Error= %v", centralTime, err.Error())
	}

	moscowTOut := tzuMoscow.TimeOut.DateTime.Format(fmtstr)

	if moscowTime != moscowTOut {
		t.Errorf("Error. Moscow Time zone conversion failed! Expected %v. Instead, got %v.", moscowTime, moscowTOut)
	}

	tzuBeijing, err := tzuMoscow.ConvertTz(tCentral, beijingTz, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error from tzuMoscow.ConvertTz. Central Time = %v. Error= %v", centralTime, err.Error())
	}

	beijingTOut := tzuBeijing.TimeOut.DateTime.Format(fmtstr)

	if beijingTime != beijingTOut {
		t.Errorf("Error. Beijing Time zone conversion failed! Expected %v. Instead, got %v.", beijingTime, beijingTOut)
	}

	utcTOut := tzuBeijing.TimeUTC.DateTime.Format(fmtstr)

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

	tz := TimeZoneDto{}

	tzLocal, _ := tz.ConvertTz(t1, localTzStr, fmtstr)
	t1OutStr := tzLocal.TimeIn.DateTime.Format(fmtstr)
	t2OutStr := tzLocal.TimeOut.DateTime.Format(fmtstr)

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

	tzuCentral, err := TimeZoneDto{}.ConvertTz(tPacific, ianaCentralTz, fmtstr)

	tOutStr := tzuCentral.TimeOut.DateTime.Format(fmtstr)

	if centralTime != tOutStr {
		t.Errorf("Error. Central Time zone conversion failed! Expected %v. Instead, got %v.", centralTime, tOutStr)
	}

	tzLocal, err := time.LoadLocation("Local")

	if err != nil {
		t.Errorf("Error received from time.LoadLocation(\"Local\") Error='%v'", err.Error())
	}

	expectedLocalTime := tPacific.In(tzLocal).Format(fmtstr)

	actualLocalTime := tzuCentral.TimeLocal.DateTime.Format(fmtstr)

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

	tzuCentral, err := TimeZoneDto{}.ConvertTz(tPacific, ianaCentralTz, fmtstr)

	tOutStr := tzuCentral.TimeOut.DateTime.Format(fmtstr)

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
	tzu := TimeZoneDto{}
	_, err := tzu.ConvertTz(tIn, invalidTz, fmtstr)

	if err == nil {
		t.Error("ConverTz() failed to detect INVALID Tartet Time Zone. Got: ", "err==nil")
	}

}

func TestTimeZoneUtility_GetLocationIn01(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _ := TimeZoneDto{}.New(tIn, ianaPacificTz, fmtstr)

	expectedLocation := "Local"

	actualLocation := tzu.TimeIn.TimeZone.LocationName

	if expectedLocation != actualLocation {
		t.Errorf("Expected Location='%v'. Instead, actual location='%v'", expectedLocation, actualLocation)
	}

}

func TestTimeZoneUtility_GetLocationOut_01(t *testing.T) {

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _ := TimeZoneDto{}.New(tIn, ianaPacificTz, fmtstr)

	expectedLocation := "America/Los_Angeles"

	actualLocation := tzu.TimeOut.TimeZone.LocationName

	if expectedLocation != actualLocation {
		t.Errorf("Expected Location Out='%v'. Instead, actual location out='%v'", expectedLocation, actualLocation)
	}
}

func TestTimeZoneUtility_GetZoneIn_01(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _ := TimeZoneDto{}.New(tIn, ianaPacificTz, fmtstr)

	expectedZone := "CDT"
	actualZone := tzu.TimeIn.TimeZone.ZoneName

	if expectedZone != actualZone {
		t.Errorf("Expected Zone In='%v'. Instead, actual Zone In='%v'", expectedZone, actualZone)
	}

}

func TestTimeZoneUtility_GetZoneOut_01(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu, _ := TimeZoneDto{}.New(tIn, ianaPacificTz, fmtstr)

	expectedZone := "PDT"
	actualZone := tzu.TimeOut.TimeZone.ZoneName

	if expectedZone != actualZone {
		t.Errorf("Expected Zone Out='%v'. Instead, actual Zone Out='%v'", expectedZone, actualZone)
	}

}

func TestTimeZoneUtility_IsValidTimeZone_01(t *testing.T) {
	tIn := time.Now()

	tzu := TimeZoneDto{}

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

	if tzu.TimeIn.TimeZone.IsValid() {
		t.Error("Expected tzu.TimeInZone.IsValidDateTime()=='false'.  Instead, the result was 'true'!")
	}

}

func TestTimeZoneUtility_IsValidTimeZone02(t *testing.T) {

	tzu := TimeZoneDto{}

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

	tzu := TimeZoneDto{}

	tzuPacific, err := tzu.ConvertTz(tUtc, ianaPacificTz, fmtstr)

	if err != nil {
		t.Errorf("Error from TimeZoneDto{}.ConvertTz. Utc Time = %v. Error= %v", utcTime, err.Error())
	}

	tzOutPacific := tzuPacific.TimeOut.TimeZone.LocationName

	if tzOutPacific != ianaPacificTz {
		t.Errorf("Error: Expected tzOutPacific %v. Instead, got %v", ianaPacificTz, tzOutPacific)
	}

}

func TestTimeZoneUtility_Location_02(t *testing.T) {

	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tPacific, _ := time.Parse(fmtstr, pacificTime)

	tzu := TimeZoneDto{}

	tzuLocal, err := tzu.ConvertTz(tPacific, "Local", fmtstr)

	if err != nil {
		t.Errorf("Error from TimeZoneDto{}.ConvertTz. Pacific Time = %v. Error= %v", pacificTime, err.Error())
	}

	tzOutLocal := tzuLocal.TimeOut.TimeZone.LocationName

	if "Local" != tzOutLocal {
		t.Errorf("Error: Expected tzOutLocal 'Local'. Instead, got %v", tzOutLocal)
	}

}

func TestTimeZoneUtility_New_01(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	tIn, _ := time.Parse(fmtstr, tstr)

	tzu, _ := TimeZoneDto{}.New(tIn, ianaPacificTz, FmtDateTimeYrMDayFmtStr)
	expectedTimeIn := "2017-04-29 19:54:30.000000000 -0500 CDT"
	if expectedTimeIn != tzu.TimeIn.String() {
		t.Errorf("Expected Time In='%v'. Instead, Time In='%v'", expectedTimeIn, tzu.TimeIn.String())
	}

	expectedTimeOut := "2017-04-29 17:54:30.000000000 -0700 PDT"

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
	tzu1, err := TimeZoneDto{}.New(t1, TZones.US.Pacific(), fmtstr)

	if err != nil {
		t.Errorf("Error returned from TimeZoneDto{}.New(t1, TzUsPacific ). Error='%v'", err.Error())
	}

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	t12Dur := t2.Sub(t1)

	tzu2, err := TimeZoneDto{}.NewAddDuration(tzu1, t12Dur, fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewAddDuration(tzu1, t12Dur). Error='%v'", err.Error())
	}

	tzu1OutStr := tzu1.TimeIn.DateTime.Format(fmtstr)

	if t1OutStr != tzu1OutStr {
		t.Errorf("Error: Expected Time1 TimeIn='%v'.  Instead Time1 TimeIn='%v'", t1OutStr, tzu1OutStr)
	}

	tzu2OutStr := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2OutStr {
		t.Errorf("Error: Expected after duration tzu2TimeIn='%v'.  Instead, tzu2TimeIn='%v'", t2OutStr, tzu2OutStr)
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

	actualTimeOutLoc := tzu1.TimeOut.TimeZone.LocationName

	if TZones.US.Pacific() != actualTimeOutLoc {
		t.Errorf("Error: Expected tzu1.TimeOutLoc='%v'.  Instead, tzu1.TimeOutLoc='%v'.", TZones.US.Pacific(), actualTimeOutLoc)
	}

	actualTimeOutLoc = tzu2.TimeOut.TimeZone.LocationName

	if TZones.US.Pacific() != actualTimeOutLoc {
		t.Errorf("Error: Expected tzu2.TimeOutLoc.String()='%v'.  Instead, tzu2.TimeOutLoc='%v'.", TZones.US.Pacific(), actualTimeOutLoc)
	}

}

func TestTimeZoneUtility_NewAddDate_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := TimeZoneDto{}.New(t1, TZones.US.Pacific(), fmtstr)
	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.New(t1, TzUsPacific). Error='%v'", err.Error())
	}

	tzu1OutStrTIn := tzu1.TimeIn.DateTime.Format(fmtstr)

	if t1OutStr != tzu1OutStrTIn {
		t.Errorf("Error: Expected tzu1OutStrTIn='%v'.  Instead, tzu1OutStrTIn='%v'", t1OutStr, tzu1OutStrTIn)
	}

	t2 := t1.AddDate(3, 2, 15)
	t2OutStr := t2.Format(fmtstr)

	tzu2, err := TimeZoneDto{}.NewAddDate(tzu1, 3, 2, 15, fmtstr)

	tzu2OutStrTIn := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration, err := tzu2.Sub(tzu1)

	if err != nil {
		t.Errorf("Error returned by tzu2.Sub(tzu1). Error='%v'", err.Error())
	}

	expectedDuration := t2.Sub(t1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration)
	}
}

func TestTimeZoneUtility_NewAddDateTime_01(t *testing.T) {
	// expected := "3-Years 2-Months 15-WeekDays 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	tzu1, err := TimeZoneDto{}.New(t1, TZones.US.Eastern(), fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.New(t1, TzUsEast). Error='%v'", err.Error())
	}

	tzu2, err := TimeZoneDto{}.NewAddDateTime(tzu1, 3, 2, 15, 3,
		4, 2, 0, 0, 0, fmtstr)

	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.NewAddDateTime(tzu1, 3,2, 15, 3, 4, 2,0, 0, 0). Error='%v'", err.Error())
	}

	tzu2TimeInStr := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2TimeInStr {
		t.Errorf("Error: Expected tzu2.TimeIn='%v'.  Instead, tzu2.TimeIn='%v'. ", t2OutStr, tzu2TimeInStr)
	}

	tzu2Dur, err := tzu2.Sub(tzu1)

	if err != nil {
		t.Errorf("Error returned by tzu2.Sub(tzu1)")
	}

	if t12Dur != tzu2Dur {
		t.Errorf("Error expected tzu1-tzu2 Duration='%v'.  Instead, Duration='%v'", t12Dur, tzu2Dur)
	}

}

func TestTimeZoneUtility_NewAddTime_01(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	tzu1, err := TimeZoneDto{}.New(t1, TZones.US.Pacific(), fmtstr)
	if err != nil {
		t.Errorf("Error returned by TimeZoneDto{}.New(t1, TzUsPacific). Error='%v'", err.Error())
	}

	tzu1OutStrTIn := tzu1.TimeIn.DateTime.Format(fmtstr)

	if t1OutStr != tzu1OutStrTIn {
		t.Errorf("Error: Expected tzu1OutStrTIn='%v'.  Instead, tzu1OutStrTIn='%v'", t1OutStr, tzu1OutStrTIn)
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

	tzu2, err := TimeZoneDto{}.NewAddTime(tzu1, 3, 32, 18,
		122, 58, 615, fmtstr)

	tzu2OutStrTIn := tzu2.TimeIn.DateTime.Format(fmtstr)

	if t2OutStr != tzu2OutStrTIn {
		t.Errorf("Error: Expected tzu2OutStrTIn='%v'.  Instead, tzu2OutStrTIn='%v'", t2OutStr, tzu2OutStrTIn)
	}

	actualDuration, err := tzu2.Sub(tzu1)

	if err != nil {
		t.Errorf("Error returned by tzu2.Sub(tzu1). Error='%v'", err.Error())
	}

	expectedDuration := t2.Sub(t1)

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Duration='%v'. Instead, Actual Duration='%v'", expectedDuration, actualDuration)
	}
}
