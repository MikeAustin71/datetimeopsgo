package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeDurationDto_CopyIn_01(t *testing.T) {

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)
	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)
	t12Dur := t2.Sub(t1)

	timeDto := TimeDto{Years: 3, Months: 2, Weeks: 2, WeekDays: 1, Hours: 3, Minutes: 4, Seconds: 2}

	var tDtoX, tDto TimeDurationDto

	var err error

	tDtoX, err = TimeDurationDto{}.NewDefaultStartTimePlusTimeDto(
		t1,
		timeDto)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewDefaultStartTimePlusTimeDto(t1, timeDto).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tX1 := time.Date(
		2005,
		10,
		30,
		20,
		0,
		0,
		0,
		time.UTC)

	tX2 := time.Date(
		2007,
		5,
		15,
		2,
		0,
		0,
		0,
		time.UTC)

	tDto, err = TimeDurationDto{}.NewDefaultStartEndTimes(
		tX1,
		tX2)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}." +
			"NewDefaultStartEndTimes(tX1, tX2).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDto.CopyIn(tDtoX)

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected Start Time %v. Instead, got %v.",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error- Expected End Time %v. Instead, got %v.",
			t2OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t12Dur != tDto.timeDuration {
		t.Errorf("Error- Expected Time Duration %v. Instead, got %v",
			t12Dur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v",
			expected, outStr)
	}

	if !tDto.Equal(tDtoX) {
		t.Error("Error - Expected tDto == tDtoX.\n" +
			"However, they are NOT Equal!\n")
	}

}

func TestTimeDurationDto_IsEmpty_01(t *testing.T){

	t1str := "02/15/2014 19:54:30.000000000 -0600 CST"
	t2str := "04/30/2017 22:58:32.000000000 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"

	t1, _ := time.Parse(fmtstr, t1str)
	t1OutStr := t1.Format(fmtstr)

	t2, _ := time.Parse(fmtstr, t2str)
	t2OutStr := t2.Format(fmtstr)

	tDto, err := TimeDurationDto{}.NewDefaultStartEndTimes(
		t1,
		t2)

	if err != nil {
		t.Errorf("Error returned by DurationTriad{}.NewDefaultStartEndTimes(...).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if t1OutStr != tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.startDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.startDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	if t2OutStr != tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr) {
		t.Errorf("Error: Expected DurationTriad.endDateTimeTz of %v. Instead, got %v ",
			t1OutStr, tDto.endDateTimeTz.GetDateTimeValue().Format(fmtstr))
	}

	tOutDur := t2.Sub(t1)

	if tOutDur != tDto.timeDuration {
		t.Errorf("Error: Expected DurationTriad.timeDuration of %v. Instead, got %v", tOutDur, tDto.timeDuration)
	}

	outStr := tDto.GetYearMthDaysTimeStr()

	expected := "3-Years 2-Months 15-Days 3-Hours 4-Minutes 2-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"

	if expected != outStr {
		t.Errorf("Error - Expected YrMthDay: %v. Instead, got %v", expected, outStr)
	}

	if tDto.IsEmpty() {
		t.Error("Error - Expected IsEmpty() == 'false'.\n" +
			"Instead, IsEmpty()=='true'")
	}
}

func TestTimeDurationDto_IsEmpty_02(t *testing.T){

	tDto := TimeDurationDto{}

	if !tDto.IsEmpty() {
		t.Error("Error - Expected IsEmpty() == 'true'.\n" +
			"Instead, IsEmpty() == 'false'.\n")
	}

}

func TestTimeDurationDto_DurCompare_01(t *testing.T) {
/*
		Testing Change To Standard Time
		November 2, 2014

		TCalcMode.LocalTimeZone() and TCalcMode.UtcTimeZone()
		yield the same result because duration is NOT added to
		starting date time.

*/
	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t2_1 := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		centralTz)

	t2_2 := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		centralTz)

	var tDur1, tDur2 TimeDurationDto

	tDur1, err = TimeDurationDto{}.NewStartEndTimes(
		t2_1,
		t2_2,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartEndTimes(TCalcMode.LocalTimeZone()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	expectedHours := time.Duration( HourNanoSeconds * 25)

	actualDuration := tDur1.GetThisTimeDuration()

	if expectedHours != actualDuration {
		t.Errorf("Error: Expected Local Time Duration = 25-hours.\n" +
			"Instead, actual duration = '%v'.", actualDuration.String())
	}

	tDur2, err = TimeDurationDto{}.NewStartEndTimes(
		t2_1,
		t2_2,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.UtcTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartEndTimes(TCalcMode.UtcTimeZone()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualDuration = tDur2.GetThisTimeDuration()

	if expectedHours != actualDuration {
		t.Errorf("Error: Expected UTC Time Duration = 25-hours.\n" +
			"Instead, actual duration = '%v'.", actualDuration.String())
	}
}

func TestTimeDurationDto_DurCompare_02(t *testing.T) {
	/*
		Testing Change To Standard Time
		November 2, 2014

		TCalcMode.LocalTimeZone() and TCalcMode.UtcTimeZone()
		yield different results because duration is added to
		starting date time.

	*/
	var err error
	var centralTz *time.Location

	centralTz, err = time.LoadLocation(TZones.America.Chicago())

	if err != nil {
		fmt.Printf("Error returned from time.LoadLocation(\"dt.TZones.America.Chicago()\")\n" +
			"Error='%v'\n", err.Error())
		return
	}

	initialStartDate := time.Date(
		2014,
		11,
		2,
		0,
		0,
		0,
		0,
		centralTz)

	localTzEndDate := time.Date(
		2014,
		11,
		3,
		0,
		0,
		0,
		0,
		centralTz)

	uTCTzEndDate := localTzEndDate.In(time.UTC)

	var tDur1, tDur2 TimeDurationDto

	expectedDuration := time.Duration( HourNanoSeconds * 25)

	tDur1, err = TimeDurationDto{}.NewStartTimeDuration(
		initialStartDate,
		expectedDuration,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartEndTimes(TCalcMode.LocalTimeZone()).\n" +
			"Error='%v'\n", err.Error())
		return
	}


	actualDuration := tDur1.GetThisTimeDuration()

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected Local Time Duration = 25-hours.\n" +
			"Instead, actual duration = '%v'.", actualDuration.String())
		return
	}

	startDateTime := tDur1.GetThisStartDateTimeString()

	if startDateTime != initialStartDate.Format(FmtDateTimeYrMDayFmtStr) {
		t.Errorf("Error: Expected starting date time= '%v'.\n" +
			"Instead, actual starting date time= '%v'\n",
			initialStartDate.Format(FmtDateTimeYrMDayFmtStr), startDateTime)
		return
	}

	endDateTimeStr := tDur1.GetThisEndDateTimeString()

	if endDateTimeStr != localTzEndDate.Format(FmtDateTimeYrMDayFmtStr) {
		t.Errorf("Error: Expected ending date time= '%v'.\n" +
			"Instead, actual starting date time= '%v'\n",
			localTzEndDate.Format(FmtDateTimeYrMDayFmtStr), endDateTimeStr)
		return
	}

	tDur2, err = TimeDurationDto{}.NewStartTimeDuration(
		initialStartDate,
		expectedDuration,
		TDurCalc.StdYearMth(),
		TZones.America.Chicago(),
		TCalcMode.UtcTimeZone(),
		FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned by TimeDurationDto{}.NewStartEndTimes(TCalcMode.UtcTimeZone()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	actualDuration = tDur2.GetThisTimeDuration()

	if expectedDuration != actualDuration {
		t.Errorf("Error: Expected UTC Time Duration = 25-hours.\n" +
			"Instead, actual duration = '%v'.", actualDuration.String())
		return
	}

	endDateTime := tDur2.GetThisEndDateTime().In(time.UTC)
	endDateTimeStr = endDateTime.Format(FmtDateTimeYrMDayFmtStr)

	if endDateTimeStr != uTCTzEndDate.Format(FmtDateTimeYrMDayFmtStr) {
		t.Errorf("Error: Expected ending date time UTC = '%v'.\n" +
			"Instead, actual ending date time UTC = '%v'\n",
			uTCTzEndDate.Format(FmtDateTimeYrMDayFmtStr), endDateTimeStr)
	}

}