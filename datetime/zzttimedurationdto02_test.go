package datetime

import (
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
