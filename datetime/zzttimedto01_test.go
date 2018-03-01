package datetime

import (
	"testing"

	"time"
)

func TestTimeDto_CopyOut_01(t *testing.T) {

	t0Dto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
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

	if 515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotNanoseconds)
	}


}

func TestTimeDto_CopyIn_01(t *testing.T) {

	t0Dto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
	}

	tDto, err := TimeDto{}.New(2014, 9, 0, 14, 5, 5,19,850,850,850)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2014, 9, 0, 14, 5, 5,19,850,850,850). Error='%v'", err.Error())
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

	if 515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotNanoseconds)
	}

}

func TestTimeDto_ConvertToNegativeValues(t *testing.T) {

	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
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

	if -515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", -515539300, tDto.TotNanoseconds)
	}

}

func TestTimeDto_ConvertToAbsoluteValues_01(t *testing.T) {
	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
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

	if -515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", -515539300, tDto.TotNanoseconds)
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

	if 515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotNanoseconds)
	}

}

func TestTimeDto_Empty(t *testing.T) {

	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
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

	if 515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotNanoseconds)
	}


	tDto.Empty()

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
	}

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

	if 0 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 0, tDto.TotNanoseconds)
	}

}

func TestTimeDto_Equal_01(t *testing.T) {

	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
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

	if 515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotNanoseconds)
	}

	t2Dto := tDto.CopyOut()

	if !t2Dto.Equal(tDto) {
		t.Error("Error: Expected t2Dto to EQUAL tDto. It did NOT!")
	}
}

func TestTimeDto_Equal_02(t *testing.T) {

	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
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

	if 515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotNanoseconds)
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

	tDto, err := TimeDto{}.New(year, month, 0, dateDay, hour, minute,second,millisecond,microsecond,totNanoSecs)

	millisecond = 515
	microsecond = 539
	nanosecond := 300

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
	}

	t1, err := tDto.GetDateTime(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned by tDto.GetDateTime(TzIanaUsCentral). Error='%v'", err.Error())
	}

	loc, err := time.LoadLocation(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned by time.LoadLocation(TzIanaUsCentral). " +
							"TzIanaUsCentral='%v'  Error='%v'", TzIanaUsCentral, err.Error())
	}

	t2 := time.Date(year, time.Month(month), dateDay, hour, minute, second, totNanoSecs, loc )

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

	if totNanoSecs != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", totNanoSecs, tDto.TotNanoseconds)
	}

	if !t1.Equal(t2) {
		t.Errorf("Error: expected t1 to EQUAL t2. They are NOT Equal! t1='%v'   t2='%v'", t1.Format(TzIanaUsCentral), t2.Format(TzIanaUsCentral))
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
	tDto, err := TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300)

	if err != nil {
		t.Errorf("Error returned by TimeDto{}.New(2017, 4, 0, 30, 22, 58,32,0,0,515539300). Error='%v'", err.Error())
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

	if 515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotNanoseconds)
	}

}

func TestTimeDto_New_02(t *testing.T) {

	tDto, err := TimeDto{}.New(0, 0, -8, 0, 0, 0, 0, 0, 0, 0 )

	if err != nil {
		t.Errorf("Error returned from TimeDto{}.New() Weeks=-8. Error='%v'", err.Error())
	}

	if -8 != tDto.Weeks {
		t.Errorf("Error: Expected tDto.Weeks= -8. Instead, tDto.Weeks='%v'", tDto.Weeks)
	}

	if -56 != tDto.DateDays {
		t.Errorf("Error: Expected tDto.DateDays= -56. Instead, tDto.DateDays='%v'", tDto.DateDays)
	}

}

func TestTimeDto_NewF1romDateTime_01(t *testing.T) {

	// t1str :="2017-04-30 22:58:32.515539300 -0500 CDT"
	// t1, err := time.Parse(FmtDateTimeYrMDayFmtStr, t1str)

	dTzDto, err := DateTzDto{}.NewDateTimeElements(2017, 04, 30, 22, 58, 32,515539300, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewDateTimeElements(year, month, day,...). Error='%v'", err.Error())
	}

	tDto, err := TimeDto{}.NewFromDateTime(dTzDto.DateTime)

	dt2, err := tDto.GetDateTime(TzIanaUsCentral)

	if err != nil {
		t.Errorf("Error returned by tDto.GetDateTime(TzIanaUsCentral). Error='%v'", err.Error())
	}

	if !dt2.Equal(dTzDto.DateTime) {
		t.Error("Error: Expected dt2 datetime to EQUAL dTzDto.DateTime. It did NOT!")
	}

}

func TestTimeDto_NewFromDateTzDto_01(t *testing.T) {

	// t1str :="2017-04-30 22:58:32.515539300 -0500 CDT"
	// t1, err := time.Parse(FmtDateTimeYrMDayFmtStr, t1str)

	dTzDto, err := DateTzDto{}.NewDateTimeElements(2017, 04, 30, 22, 58, 32,515539300, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)

	if err != nil {
		t.Errorf("Error returned from DateTzDto{}.NewDateTimeElements(year, month, day,...). Error='%v'", err.Error())
	}

	tDto, err := TimeDto{}.NewFromDateTzDto(dTzDto)

	t2, err := tDto.GetDateTime(TzIanaUsCentral)

	if !dTzDto.DateTime.Equal(t2) {
		t.Error("Error: Expected t2 to EQUAL dTzDto. It did NOT!")
	}

}
