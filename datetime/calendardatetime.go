package datetime

import (
	"sync"
)

type CalendarDateTime struct {
	year                 int64 // Number of Years
	month                int   // Number of Months
	dateDays             int   // Total Number of Days. Weeks x 7 plus WeekDays
	usDayOfWeekNumber    int   // Day of week beginning with Sunday=0
	hours                int   // Number of Hours.
	minutes              int   // Number of Minutes
	seconds              int   // Number of Seconds
	milliseconds         int   // Number of Milliseconds
	microseconds         int   // Number of Microseconds
	nanoseconds          int   // Remaining Nanoseconds after Milliseconds & Microseconds
	totSubSecNanoseconds int   // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
	//  plus remaining Nanoseconds
	totTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
	//                              + Seconds + Milliseconds + Nanoseconds
	julianDayNumber JulianDayNoDto     // Encapsulates Julian Day Number/Time
	timeZone        TimeZoneDefinition // Contains a detailed definition and descriptions of the Time
	//                                        Zone and Time Zone Location associated with this date time.
	calendar          CalendarSpec        // Designates the calendar associated with this date/time.
	yearNumberingMode CalendarYearNumMode // Designates the year numbering system associated
	//                                         with this date/time.
	dateTimeFmt string // Date Time Format String. Empty string or default is
	//                     "2006-01-02 15:04:05.000000000 -0700 MST"
	lock *sync.Mutex // Used for coordinating thread safe operations.
}

// NewGregorianDate - Creates a new instance of 'CalendarDateTime' formatted
// for a Gregorian Date Time.
//
func (calDTime CalendarDateTime) NewGregorianDate(
	year,
	month,
	day,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	timeZoneLocation string,
	dateTimeFmt string) (calDateTime CalendarDateTime, err error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix := "CalendarDateTime.NewGregorianDate() "

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
		CalendarSpec(0).Gregorian(),
		CalendarYearNumMode(0).Astronomical(),
		dateTimeFmt,
		ePrefix)

	return calDateTime, err
}

// NewJulianDate - Creates a new instance of 'CalendarDateTime' formatted
// for a Julian Date Time.
//
func (calDTime CalendarDateTime) NewJulianDate(
	year,
	month,
	day,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	timeZoneLocation string,
	dateTimeFmt string) (calDateTime CalendarDateTime, err error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix := "CalendarDateTime.NewJulianDate() "

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
		CalendarSpec(0).Julian(),
		CalendarYearNumMode(0).Astronomical(),
		dateTimeFmt,
		ePrefix)

	return calDateTime, err
}

// NewInt - Creates and returns a populated 'CalendarDateTime' instance.
//
func (calDTime CalendarDateTime) NewInt(
	year,
	month,
	day,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	timeZoneLocation string,
	calendar CalendarSpec,
	yearNumberingSystem CalendarYearNumMode,
	dateTimeFmt string) (calDateTime CalendarDateTime, err error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix := "CalendarDateTime.NewInt() "

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
		yearNumberingSystem,
		dateTimeFmt,
		ePrefix)

	return calDateTime, err
}

// GetDateTimeStr - Returns the equivalent date time string
// reflecting the date time value of the current 'CalendarDateTime'
// instance.
//
func (calDTime *CalendarDateTime) GetDateTimeStr() (string, error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix := "CalendarDateTime.GetDateTimeStr() "

	calDTimeUtil := calendarDateTimeUtility{}

	return calDTimeUtil.generateDateTimeStr(
		calDTime.year,
		calDTime.month,
		calDTime.dateDays,
		calDTime.usDayOfWeekNumber,
		calDTime.hours,
		calDTime.minutes,
		calDTime.seconds,
		calDTime.nanoseconds,
		calDTime.dateTimeFmt,
		ePrefix)
}