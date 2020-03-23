package datetime

import (
	"testing"
	"time"
)

func TestTimeDto_NormalizeTimeElements_01(t *testing.T) {

	t1Dto := TimeDto{}

	t1Dto.Years = 1955
	t1Dto.Months = 15
	t1Dto.DateDays = 32
	t1Dto.Hours = 48
	t1Dto.Minutes = 71
	t1Dto.Seconds = 125
	t1Dto.Milliseconds = 1001
	t1Dto.Microseconds = 1001
	t1Dto.Nanoseconds = 1001

	err := t1Dto.NormalizeTimeElements()

	if err != nil {
		t.Errorf("Error returned by t1Dto.NormalizeTimeElements(). Error='%v'", err.Error())
	}

	t2Dto := TimeDto{}

	t2Dto.Years = 1956
	t2Dto.Months = 3
	t2Dto.Weeks = 4
	t2Dto.WeekDays = 6
	t2Dto.DateDays = 34
	t2Dto.Hours = 1
	t2Dto.Minutes = 13
	t2Dto.Seconds = 6
	t2Dto.Milliseconds = 2
	t2Dto.Microseconds = 2
	t2Dto.Nanoseconds = 1
	t2Dto.TotSubSecNanoseconds = 2002001
	t2Dto.TotTimeNanoseconds = 4386002002001

	if !t1Dto.Equal(t2Dto) {
		t.Error("Expected t1Dto to EQUAL t2Dto. IT DID NOT!")
	}

	/*
	   	After Normalize Time Elements
	   	========================================
	   	TimeDto Printout
	   	========================================
	   Years:  1956
	   Months:  3
	   Weeks:  4
	   WeekDays:  6
	   DateDays:  34
	   Hours:  1
	   Minutes:  13
	   Seconds:  6
	   Milliseconds:  2
	   Microseconds:  2
	   Nanoseconds:  1
	   	Total SubSec Nanoseconds:  2002001
	   	Total Time Nanoseconds:  4386002002001
	   	========================================
	*/

}

func TestTimeDto_NormalizeTimeElements_02(t *testing.T) {

	t1Dto := TimeDto{}

	t1Dto.Years = 1955
	t1Dto.Months = 15
	t1Dto.DateDays = 32
	t1Dto.Hours = 48
	t1Dto.Minutes = 71
	t1Dto.Seconds = 125
	t1Dto.Milliseconds = 0
	t1Dto.Microseconds = 0
	t1Dto.Nanoseconds = 123456789

	err := t1Dto.NormalizeTimeElements()

	if err != nil {
		t.Errorf("Error returned by t1Dto.NormalizeTimeElements(). Error='%v'", err.Error())
	}

	t2Dto := TimeDto{}

	t2Dto.Years = 1956
	t2Dto.Months = 3
	t2Dto.Weeks = 4
	t2Dto.WeekDays = 6
	t2Dto.DateDays = 34
	t2Dto.Hours = 1
	t2Dto.Minutes = 13
	t2Dto.Seconds = 5
	t2Dto.Milliseconds = 123
	t2Dto.Microseconds = 456
	t2Dto.Nanoseconds = 789
	t2Dto.TotSubSecNanoseconds = 123456789
	t2Dto.TotTimeNanoseconds = 4385123456789

	if !t1Dto.Equal(t2Dto) {
		t.Error("Expected t1Dto to EQUAL t2Dto. IT DID NOT!")
	}
	/*
	   After Normalize Time Elements
	   ========================================
	             TimeDto Printout
	   ========================================
	                      Years:  1956
	                     Months:  3
	                      Weeks:  4
	                   WeekDays:  6
	                   DateDays:  34
	                      Hours:  1
	                    Minutes:  13
	                    Seconds:  5
	               Milliseconds:  123
	               Microseconds:  456
	                Nanoseconds:  789
	   Total SubSec Nanoseconds:  123456789
	     Total Time Nanoseconds:  4385123456789
	*/

}
func TestTimeDto_NormalizeDays_01(t *testing.T) {

	t1Dto := TimeDto{}

	t1Dto.Years = 1955
	t1Dto.Months = 15
	t1Dto.Weeks = 0
	t1Dto.WeekDays = 0
	t1Dto.DateDays = 32
	t1Dto.Hours = 48
	t1Dto.Minutes = 71
	t1Dto.Seconds = 125
	t1Dto.Milliseconds = 1001
	t1Dto.Microseconds = 1001
	t1Dto.Nanoseconds = 1001

	err := t1Dto.NormalizeTimeElements()

	if err != nil {
		t.Errorf("Error returned by t1Dto.NormalizeTimeElements(). Error='%v'", err.Error())
	}

	_, err = t1Dto.NormalizeDays()

	if err != nil {
		t.Errorf("Error returned by t1Dto.NormalizeDays(). Error='%v'", err.Error())
	}

	t2Dto := TimeDto{}

	t2Dto.Years = 1956
	t2Dto.Months = 4
	t2Dto.Weeks = 0
	t2Dto.WeekDays = 3
	t2Dto.DateDays = 3
	t2Dto.Hours = 1
	t2Dto.Minutes = 13
	t2Dto.Seconds = 6
	t2Dto.Milliseconds = 2
	t2Dto.Microseconds = 2
	t2Dto.Nanoseconds = 1
	t2Dto.TotSubSecNanoseconds = 2002001
	t2Dto.TotTimeNanoseconds = 4386002002001

	if !t1Dto.Equal(t2Dto) {
		t.Error("Expected t1Dto to EQUAL t2Dto. IT DID NOT!")
	}

	expectedDateTime := "1956-04-03 01:13:06.002002001 +0000 UTC"

	actualDateTime, err := t1Dto.GetDateTime(TZones.UTC())

	if err != nil {
		t.Errorf("Error returned by t1Dto.GetDateTime(TZones.UTC()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if expectedDateTime != actualDateTime.Format(FmtDateTimeYrMDayFmtStr) {
		t.Errorf("Error: Expected t1Dto.GetDateTime(TZones.UTC())='%v'.  Instead datetime='%v'",
			expectedDateTime, actualDateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

}


func TestTimeDto_SubTimeDto_01(t *testing.T) {

	Years := 2019
	Months := 6
	Weeks := 4
	WeekDays := 2
	DateDays := 30
	Hours := 0
	Minutes := 0
	Seconds := 1
	Milliseconds := 0
	Microseconds := 0
	Nanoseconds := 0
	TotSubSecNanoseconds := 0
	TotalTimeNanoseconds := int64(time.Second)

	tDto, err :=
		TimeDto{}.NewTimeComponents(
			2019,
			8,
			0,
			30,
			0,
			0,
			1,
			0,
			0,
			0)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewStartEndTimes(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
	}

	t2Dto, err := TimeDto{}.NewTimeComponents(
		0,
		2,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0)

	err = tDto.SubTimeDto(t2Dto)

	if err != nil {
		t.Errorf("Error returned by t0Dto.SubTimeDto(t2Dto). Error='%v'", err.Error())
	}

	if Years != tDto.Years {
		t.Errorf("Error: Expected tDto.Years='%v'. Instead, tDto.Years='%v'",
			Years, tDto.Years)
	}

	if Months != tDto.Months {
		t.Errorf("Error: Expected tDto.Months='%v'. Instead, tDto.Months='%v'",
			Months, tDto.Months)
	}

	if Weeks != tDto.Weeks {
		t.Errorf("Error: Expected tDto.Weeks='%v'. Instead, tDto.Weeks='%v'",
			Weeks, tDto.Weeks)
	}

	if WeekDays != tDto.WeekDays {
		t.Errorf("Error: Expected tDto.WeekDays='%v'. Instead, tDto.WeekDays='%v'",
			WeekDays, tDto.WeekDays)
	}

	if DateDays != tDto.DateDays {
		t.Errorf("Error: Expected tDto.DateDays='%v'. Instead, tDto.DateDays='%v'",
			DateDays, tDto.DateDays)
	}

	if Hours != tDto.Hours {
		t.Errorf("Error: Expected tDto.Hours='%v'. Instead, tDto.Hours='%v'",
			Hours, tDto.Hours)
	}

	if Minutes != tDto.Minutes {
		t.Errorf("Error: Expected tDto.Minutes='%v'. Instead, tDto.Minutes='%v'",
			Minutes, tDto.Minutes)
	}

	if Seconds != tDto.Seconds {
		t.Errorf("Error: Expected tDto.Seconds='%v'. Instead, tDto.Seconds='%v'",
			Seconds, tDto.Seconds)
	}

	if Milliseconds != tDto.Milliseconds {
		t.Errorf("Error: Expected tDto.Milliseconds='%v'. Instead, tDto.Milliseconds='%v'",
			Milliseconds, tDto.Milliseconds)
	}

	if Microseconds != tDto.Microseconds {
		t.Errorf("Error: Expected tDto.Microseconds='%v'. Instead, tDto.Microseconds='%v'",
			Microseconds, tDto.Microseconds)
	}

	if Nanoseconds != tDto.Nanoseconds {
		t.Errorf("Error: Expected tDto.Nanoseconds='%v'. Instead, tDto.Nanoseconds='%v'",
			Nanoseconds, tDto.Nanoseconds)
	}

	if TotSubSecNanoseconds != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected tDto.TotSubSecNanoseconds='%v'. Instead, tDto.TotSubSecNanoseconds='%v'",
			TotSubSecNanoseconds, tDto.TotSubSecNanoseconds)
	}

	if TotalTimeNanoseconds != tDto.TotTimeNanoseconds {
		t.Errorf("Error: Expected tDto.TotTimeNanoseconds='%v'. Instead, tDto.TotTimeNanoseconds='%v'",
			TotalTimeNanoseconds, tDto.TotTimeNanoseconds)
	}
}
