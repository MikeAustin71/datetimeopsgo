package datetime

import (
	"testing"

)

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
		t.Errorf("Error: Expected Nanoseconds='%v'.  Instead, Nanoseconds='%v'", 539, tDto.Nanoseconds)
	}

	if 515539300 != tDto.TotNanoseconds {
		t.Errorf("Error: Expected Total Nanoseconds='%v'.  Instead, Total Nanoseconds='%v'", 515539300, tDto.TotNanoseconds)
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
