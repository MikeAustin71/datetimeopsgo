package datetime

import (
	"sync"
)

type CalendarDateTime struct {
	year                 int64  // Number of Years
	month                int    // Number of Months
	dateDays             int    // Total Number of Days. Weeks x 7 plus WeekDays
	hours                int    // Number of Hours.
	minutes              int    // Number of Minutes
	seconds              int    // Number of Seconds
	milliseconds         int    // Number of Milliseconds
	microseconds         int    // Number of Microseconds
	nanoseconds          int    // Remaining Nanoseconds after Milliseconds & Microseconds
	totSubSecNanoseconds int    // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
	//  plus remaining Nanoseconds
	totTimeNanoseconds   int64  // Total Number of equivalent Nanoseconds for Hours + Minutes
	//  + Seconds + Milliseconds + Nanoseconds
	julianDayNumber JulianDayNoDto // Encapsulates Julian Day Number/Time
	timeZone TimeZoneDefinition // Contains a detailed definition and descriptions of the Time
	//                             Zone and Time Zone Location associated with this date time.
	calendar CalendarSpec       // Designates calendar associated with this date/time
	dateTimeFmt    string       // Date Time Format String. Empty string or default is
	//                             "2006-01-02 15:04:05.000000000 -0700 MST"
	lock          *sync.Mutex   // Used for coordinating thread safe operations.
}

// NewInt - Creates and returns a
func(calDTime CalendarDateTime) NewInt(
	year,
	month,
	day,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	timeZoneLocation string,
	calendar CalendarSpec,
	dateTimeFmt string) (calDateTime CalendarDateTime, err error) {


	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix := "CalendarDateTime.New() "

	calDTimeUtil := calendarDateTimeUtility{}

	calDateTime = CalendarDateTime{}

	err = calDTimeUtil.setCalDateTime(
		&calDateTime,
		int64(year),
		month,
		day,
		hours,
		minutes,
		seconds,
		nanoseconds,
		timeZoneLocation,
		calendar,
		dateTimeFmt,
		ePrefix)

	return calDateTime, err
}