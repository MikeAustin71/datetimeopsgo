package common

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleDurationBreakdown(t *testing.T) {
	t1str := "04/28/2017 19:54:30 -0500 CDT"
	t2str := "04/30/2017 22:58:32 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	durationUtility := DurationUtility{}
	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		t.Error("Time Parse1 Error:", err.Error())
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		t.Error("Time Parse2 Error:", err.Error())
	}

	dur, err := durationUtility.GetDuration(t1, t2)
	if err != nil {
		t.Error("Get Duration Failed: ", err.Error())
	}

	ed := durationUtility.GetDurationBreakDown(dur.StartDateTime, dur.TimeDuration)

	ex1 := "2-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if ed.DurationStr != ex1 {
		t.Error(fmt.Sprintf("Expected duration string of %v, got", ex1), ed.DurationStr)
	}
	// 2-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds

	ex2 := "51h4m2s"

	if ed.DefaultStr != ex2 {
		t.Error(fmt.Sprintf("Expected default druation string: %v, got", ex2), ed.DefaultStr)
	}
}

func TestTimeDurationReturn(t *testing.T) {
	t1str := "04/28/2017 19:54:30 -0500 CDT"
	t2str := "04/30/2017 22:58:32 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	durationUtility := DurationUtility{}
	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		t.Error("Time Parse1 Error:", err.Error())
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		t.Error("Time Parse2 Error:", err.Error())
	}

	dur, err := durationUtility.GetDuration(t1, t2)
	if err != nil {
		t.Error("Get Duration Failed: ", err.Error())
	}

	du := durationUtility.GetDurationBreakDown(dur.StartDateTime, dur.TimeDuration)

	if du.TimeDuration != dur.TimeDuration {
		t.Error(fmt.Sprintf("Expected Time Duration %v, got:", du.TimeDuration), dur)
	}

}

func TestElapsedYearsBreakdown(t *testing.T) {
	t1str := "02/15/2014 19:54:30 -0600 CST"
	t2str := "04/30/2017 22:58:32 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	durationUtility := DurationUtility{}
	t1, err := time.Parse(fmtstr, t1str)
	if err != nil {
		t.Error("Time Parse1 Error:", err.Error())
	}

	t2, err := time.Parse(fmtstr, t2str)
	if err != nil {
		t.Error("Time Parse2 Error:", err.Error())
	}

	dur, err := durationUtility.GetDuration(t1, t2)
	if err != nil {
		t.Error("Get Duration Failed: ", err.Error())
	}

	du := durationUtility.GetDurationBreakDown(dur.StartDateTime, dur.TimeDuration)


	if du.DurationStr != expected {
		t.Error(fmt.Sprintf("Expected: %v, got:", expected), du.DurationStr)
	}

	t3 := t1.AddDate(3, 2, 15).Add(time.Duration(int64((3 * HourNanoSeconds) + (4 * MinuteNanoSeconds) + (2 * SecondNanoseconds))))

	t3EndStr := t3.Format(fmtstr)
	dEndStr := du.EndDateTime.Format(fmtstr)

	if t3EndStr != dEndStr {
		t.Errorf("Expected EndTime: %v Error - Instead got %v", t3EndStr, dEndStr)
	}

}

func TestElapsedTimeBreakdown(t *testing.T) {
	tstr1 := "04/15/2017 19:54:30.123456489 -0500 CDT"
	tstr2 := "04/18/2017 09:21:16.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, err1 := time.Parse(fmtstr, tstr1)

	if err1 != nil {
		t.Error("Error On Time Parse #1: ", err1.Error())
	}

	t2, err2 := time.Parse(fmtstr, tstr2)

	if err2 != nil {
		t.Error("Error On Time Parse #2: ", err2.Error())
	}

	durationUtility := DurationUtility{}

	ed, err4 := durationUtility.GetElapsedTime(t1, t2)
	if err4 != nil {
		t.Error("Error On GetElapsedTime: ", err4.Error())
	}

	ex1 := "2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds"

	if ed.DurationStr != ex1 {
		t.Error(fmt.Sprintf("Expected %v, got", ex1), ed.DurationStr)
	}

	ex2 := "61h26m46.864197832s"

	if ed.DefaultStr != ex2 {
		t.Error(fmt.Sprintf("Expected %v, got", ex2), ed.DefaultStr)
	}

	ex3 := "2-Days 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds"

	if ex3 != ed.NanosecStr {
		t.Error(fmt.Sprintf("Expected %v, got", ex3), ed.NanosecStr)
	}

}

func TestGetDurationFromElapsedTime(t *testing.T) {
	tstr1 := "04/15/2017 19:54:30.123456489 -0500 CDT"
	tstr2 := "04/18/2017 09:21:16.987654321 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, err1 := time.Parse(fmtstr, tstr1)

	if err1 != nil {
		t.Error("Error On Time Parse #1: ", err1.Error())
	}

	t2, err2 := time.Parse(fmtstr, tstr2)

	if err2 != nil {
		t.Error("Error On Time Parse #2: ", err2.Error())
	}

	du := DurationUtility{}

	dur, err3 := du.GetDuration(t1, t2)

	if err3 != nil {
		t.Error("Error On GetDuration(t1,t2) : ", err3.Error())
	}

	ed, err4 := du.GetElapsedTime(t1, t2)

	if err4 != nil {
		t.Error("Error On GetElapsedTime(t1,t2) : ", err4.Error())
	}

	dur2, err5 := du.GetDurationFromElapsedTime(ed)

	if err5 != nil {
		t.Error("Error on GetDurationFromElapsedTime(ed) :", err5.Error())
	}

	if dur.TimeDuration != dur2 {
		t.Error(fmt.Sprintf("Duration #1 is NOT Equal to Duration #2. Expected %v , got:", dur), dur2)
	}

	if dur.TimeDuration != ed.TimeDuration {
		t.Error(fmt.Sprintf("Duration Utility Time Duration is NOT Equal to Duration #2. Expected %v , got:", ed.TimeDuration), dur2)
	}

}

func TestTimePlusDuration(t *testing.T) {

	tstr1 := "04/15/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, err1 := time.Parse(fmtstr, tstr1)

	if err1 != nil {
		t.Error("Error On Time Parse #1: ", err1.Error())
	}

	secondsInADay := 60 * 60 * 24

	dur := time.Duration(secondsInADay) * time.Second

	du := DurationUtility{}

	t2 := du.GetTimePlusDuration(t1, dur)

	tstr2 := t2.Format(fmtstr)

	expected := "04/16/2017 19:54:30.123456489 -0500 CDT"

	if expected != tstr2 {
		t.Error(fmt.Sprintf("GetTimePlusDuration() gave INVALID Result! Expected %v, got: ", expected), tstr2)
	}

}

func TestDurationUtility_GetTimeMinusDuration(t *testing.T) {
	tstr1 := "04/15/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "04/12/2017 19:54:30.123456489 -0500 CDT"
	t1, err1 := time.Parse(fmtstr, tstr1)

	if err1 != nil {
		t.Error("Error On Time Parse #1: ", err1.Error())
	}

	secondsInADay := 60 * 60 * 24

	duration3Days := secondsInADay * 3

	dur := time.Duration(duration3Days) * time.Second

	du := DurationUtility{}

	subTime := du.GetTimeMinusDuration(t1, dur)

	tstr2 := subTime.Format(fmtstr)

	if tstr2 != expected {
		t.Error(fmt.Sprintf("Expected Time string, '%v', instead got:", expected), tstr2)
	}

}

func TestDurationUtility_GetTimeMinusDuration_02(t *testing.T) {
	tstr1 := "04/15/2017 19:54:30.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "04/10/2017 19:54:30.000000000 -0500 CDT"
	t1, err := time.Parse(fmtstr, tstr1)

	if err != nil {
		t.Error("Error On Time Parse #1: ", err.Error())
	}

	du := DurationUtility{Days: 5}

	dur, err := du.GenerateDuration(du)

	if err != nil {
		t.Errorf("Error returned from du.GenerateDuration(du). Error %v ", err.Error())
	}

	subTime := du.GetTimeMinusDuration(t1, dur)

	tstr2 := subTime.Format(fmtstr)

	if tstr2 != expected {
		t.Error(fmt.Sprintf("Expected Time string 5-days before orignal time. Time Expected, '%v', instead got:", expected), tstr2)
	}

}

func TestAddDurations(t *testing.T) {

	tstr1 := "04/15/2017 19:54:30.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, tstr1)

	secondsInADay := 60 * 60 * 24

	secondsInTwoDays := 60 * 60 * 24 * 2

	// Adding duration of 1-day plus duration of 2-days should
	// equal 3-days.
	dur1 := time.Duration(secondsInADay) * time.Second

	dur2 := time.Duration(secondsInTwoDays) * time.Second

	du := DurationUtility{}

	du2 := du.GetDurationBreakDown(t1, dur1)

	du2.AddDurationToThis(dur2)

	expected := "3-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != du2.DurationStr {
		t.Error(fmt.Sprintf("Expected Total Duration of Three Days, %v - Got: ", expected), du2.DurationStr)
	}

}

func TestDurationEquality(t *testing.T) {

	tstr1 := "04/15/2017 19:54:30.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, tstr1)

	secondsInADay := 60 * 60 * 24

	secondsInTwoDays := 60 * 60 * 24 * 2

	// Adding duration of 1-day plus duration of 2-days should
	// equal 3-days.
	dur1 := time.Duration(secondsInADay) * time.Second

	dur2 := time.Duration(secondsInTwoDays) * time.Second

	du := DurationUtility{}

	du2 := du.GetDurationBreakDown(t1, dur1)

	du2.AddDurationToThis(dur2)

	expected := "3-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != du2.DurationStr {
		t.Error(fmt.Sprintf("Expected Total Duration of Three Days, %v - Got: ", expected), du2.DurationStr)
	}

	secondsInThreeDays := 60 * 60 * 24 * 3
	dur3 := time.Duration(secondsInThreeDays) * time.Second

	du3 := du.GetDurationBreakDown(t1, dur3)

	result := du3.Equal(du2)

	if result == false {
		t.Error("Expected Two Data Utility Structures to be Equal or result = true, Got: ", result)
	}
}

func TestDurationUtility_GetDurationBySeconds(t *testing.T) {
	du := DurationUtility{}

	dur := du.GetDurationBySeconds(10)

	expectedDur := SecondNanoseconds * int64(10)

	if expectedDur != int64(dur) {
		t.Errorf("Expected duration 10 seconds: %v, received duration: %v",int64(expectedDur), dur)
	}

}

func TestDurationUtility_GetDurationByMinutes(t *testing.T) {
	du := DurationUtility{}
	dur := du.GetDurationByMinutes(5)

	expectedDur := MinuteNanoSeconds * int64(5)

	if expectedDur != int64(dur) {
		t.Errorf("Expected duration 10 minutes: %v, received duration: %v",int64(expectedDur), dur)
	}


}

func TestAddDurationStructures(t *testing.T) {
	tstr1 := "04/15/2017 19:54:30.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, tstr1)

	secondsInADay := 60 * 60 * 24

	secondsInTwoDays := 60 * 60 * 24 * 2

	// Adding duration of 1-day plus duration of 2-days should
	// equal 3-days.
	dur1 := time.Duration(secondsInADay) * time.Second

	dur2 := time.Duration(secondsInTwoDays) * time.Second

	du := DurationUtility{}

	du2 := du.GetDurationBreakDown(t1, dur1)

	du3 := du.GetDurationBreakDown(t1, dur2)

	du.CopyToThis(du2)

	du.AddToThis(du3)

	expected := "3-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != du.DurationStr {
		t.Error(fmt.Sprintf("Expected Total Duration of Three Days, %v - Got: ", expected), du.DurationStr)
	}

}

func TestDurationUtility_GenerateDuration(t *testing.T) {
	du := DurationUtility{Days: 3}

	dur, err := du.GenerateDuration(du)

	if err != nil {
		t.Errorf("TestDurationUtility_GenerateDuration error from GenerateDuration - Error: %v", err.Error())
	}

	nanoSecs := int64(dur)

	nanoSecs3Days := DayNanoSeconds * int64(3)

	if nanoSecs != nanoSecs3Days {
		t.Errorf("Expected 3-days equivalent nanoseconds==%v, actually received:%v", nanoSecs3Days, nanoSecs)
	}

}

func TestDurationUtility_GetDurationBreakDown_Zero_Duration(t *testing.T) {

	du := DurationUtility{}

	zeroDuration := time.Duration(0)

	elapsedTime := du.GetDurationBreakDown(time.Time{}, time.Duration(0))

	expectedDurStr := "0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	expectedNanoSecStr := "0-Hours 0-Minutes 0-Seconds 0-Nanoseconds"

	if elapsedTime.TimeDuration != zeroDuration {
		t.Errorf("Expected elapsedTime.TimeDuration == zero, instead got: %v", elapsedTime.TimeDuration)
	}

	if elapsedTime.NanosecStr != expectedNanoSecStr {
		t.Errorf("Expected elapsedTime.NanosecStr to equal '%v', instead got: %v",
			expectedNanoSecStr, elapsedTime.NanosecStr)
	}

	if elapsedTime.DurationStr != expectedDurStr {
		t.Errorf("Expected elapsedTime.DurationStr to equal '%v', instead got: %v",
			expectedDurStr, elapsedTime.DurationStr)

	}
}

func TestDurationUtility_CalculateTargetTimeFromMinusDuration(t *testing.T) {
	tstr1 := "04/15/2017 19:54:30.123456489 +0000 UTC"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	du := DurationUtility{}

	t1, _ := time.Parse(fmtstr, tstr1)

	du.CalculateTargetTimeFromMinusDuration(t1, 1, 0,
		0, 0 , 0, 0, 0, 0, 0)

	tstr2 := du.StartDateTime.Format(fmtstr)

	expected := "04/15/2016 19:54:30.123456489 +0000 UTC"

	if tstr2 != expected {
		t.Errorf("Error: Expected 1-Year duration Start Time of %v. Received %v", expected, tstr2)
	}

}

func TestDurationUtility_CalculateDurationElements_Years(t *testing.T) {
	t1Str := "04/30/2017 22:58:32.000000000 -0500 CDT"

	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1Str)
	t2 := t1.AddDate(4, 0, 2 )

	td := t2.Sub(t1)

	du := DurationUtility{StartDateTime:t1, TimeDuration:td}
	du.CalculateDurationElements()
	du.CalculateDurationStrings()
	expected := "4-Years 0-Months 2-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if du.DurationStr != expected {
		t.Errorf("Expected Duration %v. Error - Received %v", expected, du.DurationStr)
	}

}

func TestDurationUtility_CalculateDurationElements_Months(t *testing.T) {
	t1Str := "04/30/2017 22:58:32.000000000 -0500 CDT"

	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1Str)
	t2 := t1.AddDate(4, 3, 2 )

	td := t2.Sub(t1)

	du := DurationUtility{StartDateTime:t1, TimeDuration:td}
	du.CalculateDurationElements()
	du.CalculateDurationStrings()
	expected := "4-Years 3-Months 2-Days 0-Hours 0-Minutes 0-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if du.DurationStr != expected {
		t.Errorf("Expected Duration %v. Error - Received %v", expected, du.DurationStr)
	}

}

func TestDurationUtility_SetStartTimeDuration(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 23:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "3-Years 2-Months 15-Days 4-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	du := DurationUtility{}
	t1, _ := time.Parse(fmtstr, t1str)
	t1fmt := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2fmt := t2.Format(fmtstr)
	dur := time.Duration(int64(101099042000000000))

	du.SetStartTimeDuration(t1, dur)

	if du.DurationStr != expected {
		t.Errorf("Expected Duration %v. Error - Received %v", expected, du.DurationStr)
	}

	if t1fmt != du.StartDateTime.Format(fmtstr) {
		t.Errorf("Expected Start Time: %v. Error - Received %v", t1fmt, du.StartDateTime.Format(fmtstr))
	}

	if t2fmt != du.EndDateTime.Format(fmtstr) {
		t.Errorf("Expected End Time: %v. Error - Received %v", t2fmt, du.EndDateTime.Format(fmtstr))
	}

	expected2 := fmt.Sprintf("%v",dur)
	actual := fmt.Sprintf("%v", du.TimeDuration)

	if expected2 != actual {
		t.Errorf("Expected duration, %v. Error - Received duration %v.", expected2, actual)
	}

}

func TestDurationUtility_CalculateDurationFromElements(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 23:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	expected := "3-Years 2-Months 15-Days 4-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	t1fmt := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2fmt := t2.Format(fmtstr)
	dur := int64(101099042000000000)


	du := DurationUtility{StartDateTime:t1, Years:3, Months:2, Days:15, Hours:4, Minutes:4, Seconds:2}

	du.CalculateDurationFromElements()

	if du.DurationStr != expected {
		t.Errorf("Expected Duration %v. Error - Received %v", expected, du.DurationStr)
	}

	if t1fmt != du.StartDateTime.Format(fmtstr) {
		t.Errorf("Expected Start Time: %v. Error - Received %v", t1fmt, du.StartDateTime.Format(fmtstr))
	}

	if t2fmt != du.EndDateTime.Format(fmtstr) {
		t.Errorf("Expected End Time: %v. Error - Received %v", t2fmt, du.EndDateTime.Format(fmtstr))
	}

	actual := int64(du.TimeDuration)

	if dur != actual {
		t.Errorf("Expected duration, %v. Error - Received duration %v.", dur, actual)
	}

}

func TestDurationUtility_CalculateDurationFromMinusElements(t *testing.T) {
	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
	t1fmt := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2fmt := t2.Format(fmtstr)
	//dur := time.Duration(int64(101099042000000000))


	du := DurationUtility{StartDateTime:t2, Years:-3, Months:-2, Days:-15, Hours:-3, Minutes:-4, Seconds:-2}

	du.CalculateDurationFromElements()

	if du.DurationStr != expected {
		t.Errorf("Expected Duration %v. Error - Received %v", expected, du.DurationStr)
	}

	if t1fmt != du.StartDateTime.Format(fmtstr) {
		t.Errorf("Expected Start Time: %v. Error - Received %v", t1fmt, du.StartDateTime.Format(fmtstr))
	}

	if t2fmt != du.EndDateTime.Format(fmtstr) {
		t.Errorf("Expected End Time: %v. Error - Received %v", t2fmt, du.EndDateTime.Format(fmtstr))
	}

	expected2 := fmt.Sprintf("28082h4m2s")
	actual := fmt.Sprintf("%v", du.TimeDuration)

	if expected2 != actual {
		t.Errorf("Expected duration, %v. Error - Received duration %v.", expected2, actual)
	}

}

