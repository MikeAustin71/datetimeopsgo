package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDateTimeStr(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	expected := "20170429195430"

	testTime, _ := time.Parse(fmtstr, tstr)

	result := DtMgr{}.GetDateTimeStr(testTime)

	if result != expected {
		t.Error("Expected '20170429195430' got", result)
	}

}

func TestGetDateTimeSecText(t *testing.T) {

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	expected := "2017-04-29 19:54:30"
	testTime, _ := time.Parse(fmtstr, tstr)
	result := DtMgr{}.GetDateTimeSecText(testTime)

	if result != expected {
		t.Error("Expected '", expected, "' got", result)
	}

}

func TestDtMgr_GetMilitaryCompactDateTimeGroup_01(t *testing.T) {

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"

	expected := "291954RAPR17"
	var testTime time.Time
	var err error
	var result string

	testTime, err = time.Parse(fmtstr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, tstr).\n" +
			"fmtstr='%v'  tstr='%v'\n" +
			"Error='%v'\n",
			fmtstr, tstr, err.Error())
		return
	}


	result, err = DtMgr{}.GetMilitaryCompactDateTimeGroup(testTime)

	if err != nil {
		t.Errorf("Error returned by DtMgr{}.GetMilitaryCompactDateTimeGroup(testTime).\n" +
			"testTime='%v'\n" +
			"Error='%v'\n",
			testTime.Format(FmtDateTimeYMDHMSTz), err.Error())
		return
	}

	if result != expected {
		t.Errorf("Error: Expected result='%v'.\n" +
			"Instead, result='%v'\n" +
			"Actual Time='%v'", expected, result, testTime.Format(fmtstr))
	}

}

func TestDtMgr_GetMilitaryCompactDateTimeGroup_02(t *testing.T) {

	tstr := "11/29/2017 19:54:30 -0600 CST"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"

	expected := "291954SNOV17"
	var testTime time.Time
	var err error
	var result string

	testTime, err = time.Parse(fmtstr, tstr)

	if err != nil {
		t.Errorf("Error returned by time.Parse(fmtstr, tstr).\n" +
			"fmtstr='%v'  tstr='%v'\n" +
			"Error='%v'\n",
			fmtstr, tstr, err.Error())
		return
	}


	result, err = DtMgr{}.GetMilitaryCompactDateTimeGroup(testTime)

	if err != nil {
		t.Errorf("Error returned by DtMgr{}.GetMilitaryCompactDateTimeGroup(testTime).\n" +
			"testTime='%v'\n" +
			"Error='%v'\n",
			testTime.Format(FmtDateTimeYMDHMSTz), err.Error())
		return
	}

	if result != expected {
		t.Errorf("Error: Expected result='%v'.\n" +
			"Instead, result='%v'\n" +
			"Actual Time='%v'", expected, result, testTime.Format(fmtstr))
	}

}

func TestGetDateTimeNanoSecText(t *testing.T) {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "2017-04-29 19:54:30.123456489"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DtMgr{}
	result := dt.GetDateTimeNanoSecText(testTime)

	if result != expected {
		t.Error("Expected '", expected, "' got", result)
	}

}

func TestGetDateTimeTzNanoSecText(t *testing.T) {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "04/29/2017 19:54:30.123456489 -0500 CDT"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DtMgr{}
	result := dt.GetDateTimeTzNanoSecText(testTime)

	if result != expected {
		t.Error("Expected '", expected, "' got", result)
	}

}

func TestGetDateTimeTzNanoSecYMDText(t *testing.T) {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "2017-04-29 19:54:30.123456489 -0500 CDT"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DtMgr{}
	result := dt.GetDateTimeTzNanoSecYMDText(testTime)

	if result != expected {
		t.Error("Expected '", expected, "' got", result)
	}

}

func TestGetDateTimeTzNanoSecDowYMDText(t *testing.T) {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "Saturday 2017-04-29 19:54:30.123456489 -0500 CDT"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DtMgr{}
	result := dt.GetDateTimeTzNanoSecDowYMDText(testTime)

	if result != expected {
		t.Error("Expected '", expected, "' got", result)
	}

}

func TestGetDateTimeYMDAbbrvDowNano(t *testing.T) {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "2017-04-29 Sat 19:54:30.123456489 -0500 CDT"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DtMgr{}
	result := dt.GetDateTimeYMDAbbrvDowNano(testTime)

	if result != expected {
		t.Error("Expected '", expected, "' got", result)
	}

}

func TestGetDateTimeTzNanoSecYMDDowText(t *testing.T) {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "2017-04-29 Saturday 19:54:30.123456489 -0500 CDT"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DtMgr{}
	result := dt.GetDateTimeTzNanoSecYMDDowText(testTime)

	if result != expected {
		t.Error("Expected '", expected, "' got", result)
	}

}

func TestGetDateTimeEverything(t *testing.T) {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "Saturday April 29, 2017 19:54:30.123456489 -0500 CDT"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DtMgr{}
	str := dt.GetDateTimeEverything(testTime)

	if str != expected {
		t.Error(fmt.Sprintf("Expected datetime: '%v', got", expected), str)
	}

}

func TestCustomDateTimeFormat(t *testing.T) {
	dt := DtMgr{}
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "Saturday April 29, 2017 19:54:30.123456489 -0500 CDT"
	testTime, _ := time.Parse(fmtstr, tstr)
	result := dt.GetDateTimeCustomFmt(testTime, FmtDateTimeEverything)

	if result != expected {
		t.Error(fmt.Sprintf("Expected: %v, got", expected), result)
	}

}
