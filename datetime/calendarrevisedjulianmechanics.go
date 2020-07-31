package datetime

import "sync"

type calendarRevisedJulianMechanics struct {

	lock *sync.Mutex
}


// revisedJulianCalendarToJulianDayNo - Converts a date time under the revised
// Julian Calendar to a Julian Day Number/Time.
//
// The Revised Julian calendar, also known as the Milanković calendar, or,
// less formally, new calendar, is a calendar proposed by the Serbian scientist
// Milutin Milanković in 1923, which effectively discontinued the 340 years
// of divergence between the naming of dates sanctioned by those Eastern
// Orthodox churches adopting it and the Gregorian calendar that has come
// to predominate worldwide. This calendar was intended to replace the
// ecclesiastical calendar based on the Julian calendar hitherto in use
// by all of the Eastern Orthodox Church. From 1 March 1600 through 28 February 2800,
// the Revised Julian calendar aligns its dates with the Gregorian calendar,
// which was proclaimed in 1582 by Pope Gregory XIII for adoption by the
// Christian world. The calendar has been adopted by the Orthodox churches
// of Constantinople, Albania, Alexandria, Antioch, Bulgaria, Cyprus,
// Greece, and Romania.
//
// The Revised Julian calendar has the same months and month lengths as the
// Julian calendar, but, in the Revised Julian calendar, years evenly
// divisible by 100 are not leap years, except that years with remainders
// of 200 or 600 when divided by 900 remain leap years, e.g. 2000 and 2400
// as in the Gregorian Calendar.
//
// For additional information, reference:
//    https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
//
// Summary
//
// 1. Years evenly divisible by 4 are leap years unless they are
//    century years.
//
// 2. Years evenly divisible by 100 are not leap years unless when
//    divided by 900 those years have remainders of 200 or 600 in
//    which case they are leap years.
//
func (calRJulMech *calendarRevisedJulianMechanics) revisedJulianCalendarToJulianDayNo(
	years int64,
	months,
	days,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	ePrefix string) (
	julianDayNoDto JulianDayNoDto,
	err error) {

	if calRJulMech.lock == nil {
		calRJulMech.lock = new(sync.Mutex)
	}

	calRJulMech.lock.Lock()

	defer calRJulMech.lock.Unlock()

	ePrefix += "calendarMechanics.revisedJulianCalendarToJulianDayNo() "

	err = nil

	julianDayNoDto = JulianDayNoDto{}

	// Julian Day numbers start on day zero at noon. For the Julian Calendar,
	// this means that Julian Day Number Zero begins at noon on Monday,
	// January 1, 4713 BCE, in the proleptic Julian calendar. Using
	// astronomical year numbering this is Monday, January 1, -4712
	//
	// Julian Day Calculation Summary
	//
	// 1. Julian Day 0 is Julian Calendar Day: Monday, January 1, -4712
	// using Astronomical year numbering.
	//
	// 2. Years evenly divisible by 4 are leap years unless they are
	//    century years.
	//
	// 3. Years evenly divisible by 100 are not leap years unless when
	//    divided by 900 those years have remainders of 200 or 600 in
	//    which case they are leap years.

	// baseYear is year prior to years - 4712

	return julianDayNoDto, err
}

