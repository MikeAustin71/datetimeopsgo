package common

import (
	"testing"
	"time"
)

func TestTimeZoneUtility_ReclassifyTimeWithTzLocal(t *testing.T) {
	/*
			Example Method: ReclassifyTimeWithNewTz()
		Input Time :  2017-04-29 17:54:30 -0700 PDT
		Output Time:  2017-04-29 17:54:30 -0500 CDT
	*/

	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tz := TimeZoneUtility{}
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
	tz := TimeZoneUtility{}
	tIn, err := time.Parse(fmtstr, tPacific)
	if err != nil {
		t.Errorf("Error returned from time.Parse: %v", err.Error())
	}

	tOut, err := tz.ReclassifyTimeWithNewTz(tIn, TzUsHawaii)

	tOutLoc := tOut.Location()

	if tOutLoc.String() != TzUsHawaii {
		t.Errorf("Expected tOutLocation == '%v', instead tOutLocation == '%v'", TzUsHawaii, tOutLoc.String())
	}

}

func TestTimeZoneUtility_ReclassifyTimeAsMountain(t *testing.T) {

	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tz := TimeZoneUtility{}
	tIn, err := time.Parse(fmtstr, tPacific)
	if err != nil {
		t.Errorf("Error returned from time.Parse: %v", err.Error())
	}

	tOut, err := tz.ReclassifyTimeWithNewTz(tIn, TzUsMountain)

	tOutLoc := tOut.Location()

	if tOutLoc.String() != TzUsMountain {
		t.Errorf("Expected tOutLocation == '%v', instead tOutLocation == '%v'", TzUsHawaii, tOutLoc.String())
	}

}
