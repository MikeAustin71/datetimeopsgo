package datetime

import (
	"testing"

	"time"
)

func TestTimeDto_AddTimeDto(t *testing.T) {

	Years := 2018
	Months := 6
	Weeks := 4
	WeekDays := 2
	DateDays := 30
	Hours := 22
	Minutes := 58
	Seconds := 32
	Milliseconds := 515
	Microseconds := 539
	Nanoseconds := 300
	TotSubSecNanoseconds := 515539300
	TotalTimeNanoseconds := int64(82712515539300)

	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
	}

	t2Dto, err := TimeDto{}.New(0, 14, 0, 0, 0, 0, 0, 0, 0, 0)

	err = tDto.AddTimeDto(t2Dto)

	if err != nil {
		t.Errorf("Error returned by t0Dto.AddTimeDto(t2Dto). Error='%v'", err.Error())
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

func TestTimeDto_CopyOut_01(t *testing.T) {

	t0Dto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300).\n" +
			"Error='%v'", err.Error())
		return
	}

	tDto := t0Dto.CopyOut()

	if 2017 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", 2017, tDto.Years)
	}

	if 4 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", 4, tDto.Months)
	}

	if 4 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", 4, tDto.Weeks)
	}

	if 2 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", 2, tDto.WeekDays)
	}

	if 30 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", 30, tDto.DateDays)
	}

	if 22 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", 22, tDto.Hours)
	}

	if 58 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", 58, tDto.Minutes)
	}

	if 32 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", 32, tDto.Seconds)
	}

	if 515 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", 512, tDto.Milliseconds)
	}

	if 539 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", 539, tDto.Microseconds)
	}

	if 300 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", 300, tDto.Nanoseconds)
	}

	if 515539300 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotSubSecNanoseconds)
	}

}

func TestTimeDto_CopyIn_01(t *testing.T) {

	t0Dto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
	}

	tDto, err := TimeDto{}.New(2014, 9, 0, 14, 5, 5, 19, 850, 850, 850)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2014, 9, 0, 14, 5, 5,19,850,850,850).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDto.CopyIn(t0Dto)

	if 2017 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", 2017, tDto.Years)
	}

	if 4 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", 4, tDto.Months)
	}

	if 4 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", 4, tDto.Weeks)
	}

	if 2 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", 2, tDto.WeekDays)
	}

	if 30 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", 30, tDto.DateDays)
	}

	if 22 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", 22, tDto.Hours)
	}

	if 58 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", 58, tDto.Minutes)
	}

	if 32 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", 32, tDto.Seconds)
	}

	if 515 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", 512, tDto.Milliseconds)
	}

	if 539 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", 539, tDto.Microseconds)
	}

	if 300 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", 300, tDto.Nanoseconds)
	}

	if 515539300 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotSubSecNanoseconds)
	}

}

func TestTimeDto_ConvertToNegativeValues(t *testing.T) {

	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDto.ConvertToNegativeValues()

	if -2017 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", -2017, tDto.Years)
	}

	if -4 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", -4, tDto.Months)
	}

	if -4 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", -4, tDto.Weeks)
	}

	if -2 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", -2, tDto.WeekDays)
	}

	if -30 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", -30, tDto.DateDays)
	}

	if -22 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", -22, tDto.Hours)
	}

	if -58 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", -58, tDto.Minutes)
	}

	if -32 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", -32, tDto.Seconds)
	}

	if -515 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", -512, tDto.Milliseconds)
	}

	if -539 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", -539, tDto.Microseconds)
	}

	if -300 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", -300, tDto.Nanoseconds)
	}

	if -515539300 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", -515539300, tDto.TotSubSecNanoseconds)
	}

}

func TestTimeDto_ConvertToAbsoluteValues_01(t *testing.T) {
	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDto.ConvertToNegativeValues()

	if -2017 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", -2017, tDto.Years)
	}

	if -4 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", -4, tDto.Months)
	}

	if -4 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", -4, tDto.Weeks)
	}

	if -2 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", -2, tDto.WeekDays)
	}

	if -30 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", -30, tDto.DateDays)
	}

	if -22 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", -22, tDto.Hours)
	}

	if -58 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", -58, tDto.Minutes)
	}

	if -32 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", -32, tDto.Seconds)
	}

	if -515 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", -512, tDto.Milliseconds)
	}

	if -539 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", -539, tDto.Microseconds)
	}

	if -300 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", -300, tDto.Nanoseconds)
	}

	if -515539300 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", -515539300, tDto.TotSubSecNanoseconds)
	}

	tDto.ConvertToAbsoluteValues()

	if 2017 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", 2017, tDto.Years)
	}

	if 4 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", 4, tDto.Months)
	}

	if 4 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", 4, tDto.Weeks)
	}

	if 2 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", 2, tDto.WeekDays)
	}

	if 30 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", 30, tDto.DateDays)
	}

	if 22 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", 22, tDto.Hours)
	}

	if 58 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", 58, tDto.Minutes)
	}

	if 32 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", 32, tDto.Seconds)
	}

	if 515 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", 512, tDto.Milliseconds)
	}

	if 539 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", 539, tDto.Microseconds)
	}

	if 300 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", 300, tDto.Nanoseconds)
	}

	if 515539300 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotSubSecNanoseconds)
	}

}

func TestTimeDto_Empty(t *testing.T) {

	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if 2017 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", 2017, tDto.Years)
	}

	if 4 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", 4, tDto.Months)
	}

	if 4 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", 4, tDto.Weeks)
	}

	if 2 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", 2, tDto.WeekDays)
	}

	if 30 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", 30, tDto.DateDays)
	}

	if 22 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", 22, tDto.Hours)
	}

	if 58 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", 58, tDto.Minutes)
	}

	if 32 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", 32, tDto.Seconds)
	}

	if 515 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", 512, tDto.Milliseconds)
	}

	if 539 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", 539, tDto.Microseconds)
	}

	if 300 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", 300, tDto.Nanoseconds)
	}

	if 515539300 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotSubSecNanoseconds)
	}

	tDto.Empty()

	if 0 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", 0, tDto.Years)
	}

	if 0 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", 0, tDto.Months)
	}

	if 0 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", 0, tDto.Weeks)
	}

	if 0 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", 0, tDto.WeekDays)
	}

	if 0 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", 0, tDto.DateDays)
	}

	if 0 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", 0, tDto.Hours)
	}

	if 0 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", 0, tDto.Minutes)
	}

	if 0 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", 0, tDto.Seconds)
	}

	if 0 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", 0, tDto.Milliseconds)
	}

	if 0 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", 0, tDto.Microseconds)
	}

	if 0 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", 0, tDto.Nanoseconds)
	}

	if 0 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 0, tDto.TotSubSecNanoseconds)
	}

}

func TestTimeDto_Equal_01(t *testing.T) {

	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if 2017 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", 2017, tDto.Years)
	}

	if 4 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", 4, tDto.Months)
	}

	if 4 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", 4, tDto.Weeks)
	}

	if 2 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", 2, tDto.WeekDays)
	}

	if 30 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", 30, tDto.DateDays)
	}

	if 22 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", 22, tDto.Hours)
	}

	if 58 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", 58, tDto.Minutes)
	}

	if 32 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", 32, tDto.Seconds)
	}

	if 515 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", 512, tDto.Milliseconds)
	}

	if 539 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", 539, tDto.Microseconds)
	}

	if 300 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", 300, tDto.Nanoseconds)
	}

	if 515539300 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotSubSecNanoseconds)
	}

	t2Dto := tDto.CopyOut()

	if !t2Dto.Equal(tDto) {
		t.Error("Error: Expected t2Dto to EQUAL tDto. It did NOT!")
	}
}

func TestTimeDto_Equal_02(t *testing.T) {

	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if 2017 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", 2017, tDto.Years)
	}

	if 4 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", 4, tDto.Months)
	}

	if 4 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", 4, tDto.Weeks)
	}

	if 2 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", 2, tDto.WeekDays)
	}

	if 30 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", 30, tDto.DateDays)
	}

	if 22 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", 22, tDto.Hours)
	}

	if 58 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", 58, tDto.Minutes)
	}

	if 32 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", 32, tDto.Seconds)
	}

	if 515 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", 512, tDto.Milliseconds)
	}

	if 539 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", 539, tDto.Microseconds)
	}

	if 300 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", 300, tDto.Nanoseconds)
	}

	if 515539300 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotSubSecNanoseconds)
	}

	t2Dto := tDto.CopyOut()

	tDto.Nanoseconds = 301

	if t2Dto.Equal(tDto) {
		t.Error("Error: Expected t2Dto NOT EQUAL to tDto. It IS EQUAL!")
	}
}

func TestTimeDto_GetDateTime_01(t *testing.T) {

	year := 2017
	month := 4
	week := 4
	weekDay := 2
	dateDay := 30
	hour := 22
	minute := 58
	second := 32
	millisecond := 0
	microsecond := 0
	totNanoSecs := 515539300

	tDto, err := TimeDto{}.New(year, month, 0, dateDay, hour, minute, second, millisecond, microsecond, totNanoSecs)

	millisecond = 515
	microsecond = 539
	nanosecond := 300

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t1, err := tDto.GetDateTime(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by tDto.GetDateTime(TZones.US.Central()). Error='%v'", err.Error())
	}

	loc, err := time.LoadLocation(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TZones.US.Central()). "+
			"TZones.US.Central()='%v'  Error='%v'", TZones.US.Central(), err.Error())
	}

	t2 := time.Date(year, time.Month(month), dateDay, hour, minute, second, totNanoSecs, loc)

	if year != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", year, tDto.Years)
	}

	if month != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", month, tDto.Months)
	}

	if week != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", week, tDto.Weeks)
	}

	if weekDay != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", weekDay, tDto.WeekDays)
	}

	if dateDay != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", dateDay, tDto.DateDays)
	}

	if hour != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", hour, tDto.Hours)
	}

	if minute != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", minute, tDto.Minutes)
	}

	if second != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", second, tDto.Seconds)
	}

	if millisecond != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", millisecond, tDto.Milliseconds)
	}

	if microsecond != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", microsecond, tDto.Microseconds)
	}

	if nanosecond != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", nanosecond, tDto.Nanoseconds)
	}

	if totNanoSecs != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", totNanoSecs, tDto.TotSubSecNanoseconds)
	}

	if !t1.Equal(t2) {
		t.Errorf("Error: expected t1 to EQUAL t2. They are NOT Equal! t1='%v'   t2='%v'", t1.Format(TZones.US.Central()), t2.Format(TZones.US.Central()))
	}

}

func TestTimeDto_New_01(t *testing.T) {
	/*
		Original t0str:  2017-04-30 22:58:32.515539300 -0500 CDT
		Original t0:  2017-04-30 22:58:32.515539300 -0500 CDT
		========================================
		          TimeDto Printout
		========================================
		            Years:  2017
		           Months:  4
		            Weeks:  4
		         WeekDays:  2
		         DateDays:  30
		            Hours:  22
		          Minutes:  58
		          Seconds:  32
		     Milliseconds:  515
		     Microseconds:  539
		      Nanoseconds:  300
		Total Nanoseconds:  515539300
		========================================
	*/
	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58, 32, 0, 0, 515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if 2017 != tDto.Years {
		t.Errorf("Error: Expected Years='%v'.  Instead, Years='%v'", 2017, tDto.Years)
	}

	if 4 != tDto.Months {
		t.Errorf("Error: Expected Months='%v'.  Instead, Months='%v'", 4, tDto.Months)
	}

	if 4 != tDto.Weeks {
		t.Errorf("Error: Expected Weeks='%v'.  Instead, Weeks='%v'", 4, tDto.Weeks)
	}

	if 2 != tDto.WeekDays {
		t.Errorf("Error: Expected WeekDays='%v'.  Instead, WeekDays='%v'", 2, tDto.WeekDays)
	}

	if 30 != tDto.DateDays {
		t.Errorf("Error: Expected Date Days='%v'.  Instead, Date Days='%v'", 30, tDto.DateDays)
	}

	if 22 != tDto.Hours {
		t.Errorf("Error: Expected Hours='%v'.  Instead, Hours='%v'", 22, tDto.Hours)
	}

	if 58 != tDto.Minutes {
		t.Errorf("Error: Expected Minutes='%v'.  Instead, Minutes='%v'", 58, tDto.Minutes)
	}

	if 32 != tDto.Seconds {
		t.Errorf("Error: Expected Seconds='%v'.  Instead, Seconds='%v'", 32, tDto.Seconds)
	}

	if 515 != tDto.Milliseconds {
		t.Errorf("Error: Expected Milliseconds='%v'.  Instead, Milliseconds='%v'", 512, tDto.Milliseconds)
	}

	if 539 != tDto.Microseconds {
		t.Errorf("Error: Expected Microseconds='%v'.  Instead, Microseconds='%v'", 539, tDto.Microseconds)
	}

	if 300 != tDto.Nanoseconds {
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", 300, tDto.Nanoseconds)
	}

	if 515539300 != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotSubSecNanoseconds)
	}

}

func TestTimeDto_New_02(t *testing.T) {

	tDto, err := TimeDto{}.New(0, 0, -8, 0, 0, 0, 0, 0, 0, 0)

	if err != nil {
		t.Errorf("Error returned from TimeDto{}.New()\nWeeks=-8.\nError='%v'\n", err.Error())
		return
	}

	Years := 0
	Months := 0
	Weeks := -8
	WeekDays := 0
	DateDays := -56
	Hours := 0
	Minutes := 0
	Seconds := 0
	Milliseconds := 0
	Microseconds := 0
	Nanoseconds := 0

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

	/*
			====================================
		         Original TimeDto
		====================================
		========================================
		          TimeDto Printout
		========================================
		                   Years:  0
		                  Months:  0
		                   Weeks:  -8
		                WeekDays:  0
		                DateDays:  0
		                   Hours:  0
		                 Minutes:  0
		                 Seconds:  0
		            Milliseconds:  0
		            Microseconds:  0
		             Nanoseconds:  0
		Total SubSec Nanoseconds:  0
		  Total Time Nanoseconds:  0
		========================================
		------------------------------------
		        Normalized TimeDto
		------------------------------------
		========================================
		          TimeDto Printout
		========================================
		                   Years:  -1
		                  Months:  11
		                   Weeks:  4
		                WeekDays:  2
		                DateDays:  30
		                   Hours:  0
		                 Minutes:  0
		                 Seconds:  0
		            Milliseconds:  0
		            Microseconds:  0
		             Nanoseconds:  0
		Total SubSec Nanoseconds:  0
		  Total Time Nanoseconds:  0
		========================================
		             Start Date:  -0001-11-30 00:00:00.000000000 +0000 UCT
		Start Date Plus 8-Weeks:  0000-01-25 00:00:00.000000000 +0000 UCT
	*/

}

func TestTimeDto_NewFromDateTime_01(t *testing.T) {

	// t1str :="2017-04-30 22:58:32.515539300 -0500 CDT"
	// t1, err := time.Parse(FmtDateTimeYrMDayFmtStr, t1str)

	dTzDto, err := DateTzDto{}.NewDateTimeElements(2017, 04, 30, 22, 58, 32, 515539300, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewDateTimeElements(year, month, day,...).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	tDto, err := TimeDto{}.NewFromDateTime(dTzDto.GetDateTimeValue())

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewFromDateTime(dTzDto.DateTime)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	dt2, err := tDto.GetDateTime(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by tDto.GetDateTime(TZones.US.Central()).\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !dt2.Equal(dTzDto.GetDateTimeValue()) {
		t.Error("Error: Expected dt2 datetime to EQUAL dTzDto.DateTime. It did NOT!")
	}

}

func TestTimeDto_NewFromDateTzDto_01(t *testing.T) {

	// t1str :="2017-04-30 22:58:32.515539300 -0500 CDT"
	// t1, err := time.Parse(FmtDateTimeYrMDayFmtStr, t1str)

	dTzDto, err := DateTzDto{}.NewDateTimeElements(2017, 04, 30, 22, 58, 32, 515539300, TZones.US.Central(), FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewDateTimeElements(year, month, day,...). Error='%v'", err.Error())
	}

	tDto, err := TimeDto{}.NewFromDateTzDto(dTzDto)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewFromDateTzDto(dTzDto)\n" +
			"Error='%v'\n", err.Error())
		return
	}

	t2, err := tDto.GetDateTime(TZones.US.Central())

	if err != nil {
		t.Errorf("Error returned by tDto.GetDateTime(TZones.US.Central())\n" +
			"Error='%v'\n", err.Error())
		return
	}

	if !dTzDto.GetDateTimeValue().Equal(t2) {
		t.Error("Error: Expected t2 to EQUAL dTzDto. It did NOT!")
	}

}

func TestTimeDto_NewTimeElements_01(t *testing.T) {
	year := 69
	month := 5
	day := 27
	hour := 15
	minute := 30
	second := 2
	millisecond := 784
	microsecond := 303
	nanosecond := 848
	totNanosecs := 784303848

	tDto, err := TimeDto{}.NewTimeElements(year, month, day, hour, minute, second, totNanosecs)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.NewTimeElements(...).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if year != tDto.Years {
		t.Errorf("Error: Expected tDto.Years='%v'. Instead, tDto.Years='%v'",
			year, tDto.Years)
	}

	if month != tDto.Months {
		t.Errorf("Error: Expected tDto.Months='%v'. Instead, tDto.Months='%v'",
			month, tDto.Months)
	}

	if day != tDto.DateDays {
		t.Errorf("Error: Expected tDto.DateDays='%v'. Instead, tDto.DateDays='%v'",
			day, tDto.DateDays)
	}

	if hour != tDto.Hours {
		t.Errorf("Error: Expected tDto.Hours='%v'. Instead, tDto.Hours='%v'",
			hour, tDto.Hours)
	}

	if minute != tDto.Minutes {
		t.Errorf("Error: Expected tDto.Minutes='%v'. Instead, tDto.Minutes='%v'",
			minute, tDto.Minutes)
	}

	if second != tDto.Seconds {
		t.Errorf("Error: Expected tDto.Seconds='%v'. Instead, tDto.Seconds='%v'",
			second, tDto.Seconds)
	}

	if millisecond != tDto.Milliseconds {
		t.Errorf("Error: Expected tDto.Milliseconds='%v'. Instead, tDto.Milliseconds='%v'",
			millisecond, tDto.Milliseconds)
	}

	if microsecond != tDto.Microseconds {
		t.Errorf("Error: Expected tDto.Microseconds='%v'. Instead, tDto.Microseconds='%v'",
			microsecond, tDto.Microseconds)
	}

	if nanosecond != tDto.Nanoseconds {
		t.Errorf("Error: Expected tDto.Nanoseconds='%v'. Instead, tDto.Nanoseconds='%v'",
			nanosecond, tDto.Nanoseconds)
	}

	dur := int64(hour) * int64(time.Hour)
	dur += int64(minute) * int64(time.Minute)
	dur += int64(second) * int64(time.Second)
	dur += int64(totNanosecs)

	if dur != tDto.TotTimeNanoseconds {
		t.Errorf("Error: Expected tDto.TotTimeNanoseconds='%v'.  "+
			"Instead, tDto.TotTimeNanoseconds='%v'",
			dur, tDto.TotTimeNanoseconds)
	}

	if totNanosecs != tDto.TotSubSecNanoseconds {
		t.Errorf("Error: Expected tDto.TotSubSecNanoseconds='%v'.  "+
			"Instead, tDto.TotSubSecNanoseconds='%v'",
			totNanosecs, tDto.TotSubSecNanoseconds)
	}

}
