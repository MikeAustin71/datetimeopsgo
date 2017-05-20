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

	ed := durationUtility.GetDurationBreakDown(dur)

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

	du := durationUtility.GetDurationBreakDown(dur)

	if du.TimeDuration != dur {
		t.Error(fmt.Sprintf("Expected Time Duration %v, got:", du.TimeDuration), dur)
	}

}

func TestElapsedYearsBreakdown(t *testing.T) {
	t1str := "02/15/2014 19:54:30 -0500 CDT"
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

	du := durationUtility.GetDurationBreakDown(dur)

	expected := "3-Years 75-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if du.DurationStr != expected {
		t.Error(fmt.Sprintf("Expected: %v, got:", expected), du.DurationStr)
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

	if dur != dur2 {
		t.Error(fmt.Sprintf("Duration #1 is NOT Equal to Duration #2. Expected %v , got:", dur), dur2)
	}

	if dur != ed.TimeDuration {
		t.Error(fmt.Sprintf("Duration Utility Time Duration is NOT Equal to Duration #2. Expected %v , got:", ed.TimeDuration), dur2)
	}

}