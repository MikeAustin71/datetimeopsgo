package datetime

import (
	"sync"
	"time"
)

// CalendarUtility - This type is used to support calendar conversions.
//
// ------------------------------------------------------------------------
//
// Definition Of Terms
//
//
// Gregorian Calendar
// The Gregorian calendar, which is the calendar used today, was first
// introduced by Pope Gregory XIII via a papal bull in February 1582
// to correct an error in the old Julian calendar.
//
// This error had been accumulating over hundreds of years so that every
// 128 years the calendar was out of sync with the equinoxes and solstices
// by one additional day.
//
// As the centuries passed, the Julian Calendar became more inaccurate.
// Because the calendar was incorrectly determining the date of Easter,
// Pope Gregory XIII reformed the calendar to match the solar year so that
// Easter would once again "fall upon the first Sunday after the first full
// moon on or after the Vernal Equinox.".
//
// Ten days were omitted from the calendar to bring the calendar back in line
// with the solstices, and Pope Gregory XIII decreed that the day following
// Thursday, October 4, 1582 would be Friday, October 15, 1582 and from then
// on the reformed Gregorian calendar would be used.
//
// Reference http://www.searchforancestors.com/utility/gregorian.html
//
//
// Double Dating
// New Year's Day had been celebrated on March 25 under the Julian calendar
// in Great Britain and its colonies, but with the introduction of the
// Gregorian Calendar in 1752, New Year's Day was now observed on January 1.
// When New Year's Day was celebrated on March 25th, March 24 of one year was
// followed by March 25 of the following year. When the Gregorian calendar
// reform changed New Year's Day from March 25 to January 1, the year of George
// Washington's birth, because it took place in February, changed from 1731 to
// 1732. In the Julian Calendar his birthdate is Feb 11, 1731 and in the
// Gregorian Calendar it is Feb 22, 1732. Double dating was used in Great Britain
// and its colonies including America to clarify dates occurring between 1 January
// and 24 March on the years between 1582, the date of the original introduction of
// the Gregorian calendar, and 1752, when Great Britain adopted the calendar.
//
// Double dates were identified with a slash mark (/) representing the Old and New
// Style calendars, e. g., 1731/1732.
//
// Reference http://www.searchforancestors.com/utility/gregorian.html
//
//
// Astronomical Year Numbering
// "Astronomical year numbering is based on AD/CE year numbering, but follows normal
// decimal integer numbering more strictly. Thus, it has a year 0; the years before
// that are designated with negative numbers and the years after that are designated
// with positive numbers."  Wikipedia https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
// The Golang type, 'time.Time' uses Astronomical Year Numbering in that the first year
// before year '1' is year '0'. Thereafter all years before year zero have negative year
// numbers.
//
//
// Proleptic Gregorian Dates
// The Gregorian Calendar was instituted on Friday, October 15, 1582. Prior to
// this date, the Gregorian Calendar was not recognized and therefore did not
// exist. Nevertheless, dates using the Golang type, 'time.Time' can represent
// Gregorian Dates prior to October 15, 1582. Such dates are termed 'Proleptic
// Gregorian Dates'.
//
//
// Proleptic Julian Dates
//

type CalendarUtility struct {
	lock *sync.Mutex
}


// GregorianDateToJulianDayNo - Converts a Gregorian calendar
// date to Julian Day Number (JDN).
//
// This algorithm is valid for all (possibly proleptic) Gregorian
// calendar dates after November 23, −4713. Divisions are integer
// divisions, fractional parts are ignored.[72]
//
// JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075
//
// Reference Wikipedia:
//   https://en.wikipedia.org/wiki/Julian_day
//
// Note that the input parameter 'gregorianDateTime' will first be
// converted to Universal Coordinated Time (UTC).
//
//  Example: The Julian Date for 00:30:00.0 UT January 1, 2013, is 2 456 293.520 833
//           The Julian Day Number is 2456293
//
// This method calls helper method calendarMechanics.gregorianDateToJulianDayNoTime().
// The algorithm employed by this method is based on the work of E.G. Richards.
//
// Reference:
//   Richards, E. G. (1998). Mapping Time: The Calendar and its History.
//   Oxford University Press. ISBN 978-0192862051
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  gregorianDateTime  time.Time
//     - This date time value will be converted to Universal
//       Coordinated Time (UTC) before conversion to a Julian
//       Day Number (JDN).
//
//
//  ePrefix            string
//     - A string containing the names of the calling functions
//       which invoked this method. The last character in this
//       string should be a blank space.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  gregorianDateUtc   time.Time
//     - The input parameter 'gregorianDateTime' converted to Universal
//       Coordinated Time (UTC). This is the date time used to compute
//       the Julian Day Number (JDN)
//
//
//  julianDayNo        int64
//     - An int64 value representing the Julian Day Number (JDN) equivalent
//       to input parameter 'gregorianDateTime' converted to Universal
//       Coordinated Time (UTC).
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered this error Type will encapsulate
//       an error message.
//
//
//  ke!san Online Julian Day/Date Calculator
//    https://keisan.casio.com/exec/system/1227779487
//
//  How to Calculate the Julian Date
//   https://sciencing.com/calculate-age-lunar-years-5997325.html
//
//  Astronomy
//   https://www.aa.quae.nl/en/reken/juliaansedag.html
//
// How to Calculate Julian Dates
// https://www.howcast.com/videos/how-to-calculate-julian-dates
//
// Astronomical Calculations: The Julian Day
// https://squarewidget.com/julian-day/
//
func (calUtil *CalendarUtility) GregorianDateToJulianDayNo(
	gregorianDateTime time.Time,
	ePrefix string) (
	gregorianDateUtc time.Time,
	julianDayNo int64,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = new(sync.Mutex)
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "DTimeUtility.GregorianDateToJulianDayNo() "

	calendarMech := calendarMechanics{}

	gregorianDateUtc = time.Time{}
	julianDayNo = -1
	err = nil

	var julianDayNoDto JulianDayNoDto

	gregorianDateUtc = gregorianDateTime.UTC()

		julianDayNoDto,
		err =
	calendarMech.gregorianDateToJulianDayNoTime(
		int64(gregorianDateUtc.Year()),
		int(gregorianDateUtc.Month()),
		gregorianDateUtc.Day(),
		gregorianDateUtc.Hour(),
		gregorianDateUtc.Minute(),
		gregorianDateUtc.Second(),
		gregorianDateUtc.Nanosecond(),
		ePrefix)

		if err != nil {
			return gregorianDateUtc, julianDayNo, err
		}

	julianDayNo = julianDayNoDto.GetJulianDayInt64()

	return gregorianDateUtc, julianDayNo, err
}

// GregorianDateToJulianDayNoTime - Converts a Gregorian Date to a Julian
// Date.
//
// Reference Wikipedia:
//   https://en.wikipedia.org/wiki/Julian_day
//
//  Example: The Julian Date for 00:30:00.0 UT January 1, 2013, is
//  2 456 293.520 833
//
// This method calls helper method calendarMechanics.gregorianDateToJulianDayNoTime().
// The algorithm employed by this method is based on the work of E.G. Richards.
//
// Reference:
//   Richards, E. G. (1998). Mapping Time: The Calendar and its History.
//   Oxford University Press. ISBN 978-0192862051
//
// However, the original algorithm has been modified to provide for time fractions
// accurate to nanoseconds.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  gregorianDateTime  time.Time
//     - This date time value will be converted to Universal
//       Coordinated Time (UTC) before conversion to a Julian
//       Day Number (JDN).
//
//
//  ePrefix            string
//     - A string containing the names of the calling functions
//       which invoked this method. The last character in this
//       string should be a blank space.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  gregorianDateUtc   time.Time
//     - The input parameter 'gregorianDateTime' converted to Universal
//       Coordinated Time (UTC). This is the date time used to compute
//       the Julian Date
//
//
//  julianDate          float64
//     - An float64 value representing the Julian Date equivalent to
//       input parameter 'gregorianDateTime'. Note: 'gregorianDateTime'
//       is first converted to Universal Coordinated Time (UTC) before
//       being converted to a Julian Date.
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered this error Type will encapsulate
//       an error message.
//
// Resources:
//  American Association of Variable Star Observers (AAVSO)
//  https://www.aavso.org/jd-calculator
//
func (calUtil *CalendarUtility) GregorianDateToJulianDayNoTime(
	gregorianDateTime time.Time,
	ePrefix string) (
	gregorianDateUtc time.Time,
	julianDayNoDto JulianDayNoDto,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = new(sync.Mutex)
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "DTimeUtility.GregorianDateToJulianDayNoTime() "

	calendarMech := calendarMechanics{}

	gregorianDateUtc = gregorianDateTime.UTC()

	julianDayNoDto,
		err =
		calendarMech.gregorianDateToJulianDayNoTime(
			int64(gregorianDateUtc.Year()),
			int(gregorianDateUtc.Month()),
			gregorianDateUtc.Day(),
			gregorianDateUtc.Hour(),
			gregorianDateUtc.Minute(),
			gregorianDateUtc.Second(),
			gregorianDateUtc.Nanosecond(),
			ePrefix)

	return gregorianDateUtc, julianDayNoDto, err
}



// JulianDayNoTimeToGregorianCalendar - Converts a Julian Day
// Number and Time value to the corresponding date time in the
// Gregorian Calendar. Because the Gregorian Calendar was instituted
// in Friday, October 15, 1582, all Gregorian Calendar dates prior
// to this are extrapolated or proleptic. This method uses the
// 'Richards' algorithm.
//
// "This is an algorithm by E. G. Richards to convert a Julian Day Number,
// J, to a date in the Gregorian calendar (proleptic, when applicable).
// Richards states the algorithm is valid for Julian day numbers greater
// than or equal to 0".  https://en.wikipedia.org/wiki/Julian_day
//
//   Richards, E. G. (1998). Mapping Time: The Calendar and its History.
//   Oxford University Press. ISBN 978-0192862051
//
// Julian Day numbers start on day zero at noon. This means that Julian
// Day Number Times are valid for all dates on or after noon on Monday,
// January 1, 4713 BCE, in the proleptic Julian calendar or November 24,
// 4714 BCE, in the proleptic Gregorian calendar. Remember that the Golang
// 'time.Time' type uses Astronomical Year numbering with the Gregorian
// Calendar. In other words, the 'time.Time' type recognizes the year
// zero. Dates expressed in the 'Common Era' ('BCE' Before Common Era
// or 'CE' Common Era). Therefore a 'time.Time' year of '-4713' is equal
// to the year '4714 BCE'
//
// This means that the 'Richards' algorithm employed by this
// method is valid for all 'time.Time' (possibly proleptic) Gregorian
// dates on or after noon November 24, −4713.
//
// For information on the Julian Day Number/Time see:
//   https://en.wikipedia.org/wiki/Julian_day
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  julianDayNoNoTime   float64
//     - The integer portion of this number (digits to left of
//       the decimal) represents the Julian day number. The fractional
//       digits to the right of the decimal represent elapsed time
//       since noon on the Julian day number. All time values are
//       expressed as Universal Coordinated Time (UTC).
//
//  digitsAfterDecimal  int
//     - The number of digits after the decimal in input parameter
//       'julianDayNoNoTime' which will be used in the conversion
//       algorithm. Effectively, 'julianDayNoNoTime' will be rounded
//       to the number of digits to the right of the decimal specified
//       in this parameter.
//
//  ePrefix             string
//     - A string containing the names of the calling functions
//       which invoked this method. The last character in this
//       string should be a blank space.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  gregorianDateUtc    time.Time
//     - The returned parameter 'gregorianDateTime' represents the input
//       'julianDayNoNoTime' converted to the Gregorian calendar. This
//       returned 'time.Time' type is always configured as Universal
//       Coordinated Time (UTC). In addition, as a Golan 'time.Time'
//       type, the date is expressed using astronomical years. Astronomical
//       year numbering includes a Zero Year. Therefore, 1BCE is stored
//       as year zero in this return value.
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered this error Type will encapsulate
//       an error message.
//
//
// ------------------------------------------------------------------------
//
// Resources
//
//  PHP Julian date converter algorithms (Stack Overflow)
//   https://stackoverflow.com/questions/45586444/php-julian-date-converter-algorithms
//
//  Richards Algorithm
//   https://en.wikipedia.org/wiki/Julian_day
//
func (calUtil *CalendarUtility) JulianDayNoTimeToGregorianCalendar(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	gregorianDateUtc time.Time,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = new(sync.Mutex)
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "DTimeUtility.JulianDayNoTimeToGregorianCalendar() "

	calendarMech := calendarMechanics{}

	gregorianDateUtc,
		err =
		calendarMech.richardsJulianDayNoTimeToGregorianCalendar(
			julianDayNoDto,
			ePrefix)

	return gregorianDateUtc, err
}

// JulianDayNoTimeToJulianCalendar - Converts a Julian Day Number and
// Time value to the corresponding date time in the Julian Calendar.
//
// Note that Augustus corrected errors in the observance of leap years
// by omitting leap days until AD 8. Julian calendar dates before March
// AD 4 are proleptic, and do not necessarily match the dates actually
// observed in the Roman Empire.
//
// Background:
//
// "The Julian calendar, proposed by Julius Caesar in 708 Ab urbe condita
// (AUC) (46 BC), was a reform of the Roman calendar. It took effect on
// 1 January 709 AUC (45 BC), by edict. It was designed with the aid of
// Greek mathematicians and Greek astronomers such as Sosigenes of Alexandria.
//
// The [Julian] calendar was the predominant calendar in the Roman world,
// most of Europe, and in European settlements in the Americas and elsewhere,
// until it was gradually replaced by the Gregorian calendar, promulgated in
// 1582 by Pope Gregory XIII. The Julian calendar is still used in parts of
// the Eastern Orthodox Church and in parts of Oriental Orthodoxy as well as
// by the Berbers.
//
// The Julian calendar has two types of year: a normal year of 365 days and
// a leap year of 366 days. They follow a simple cycle of three normal years
// and one leap year, giving an average year that is 365.25 days long. That
// is more than the actual solar year value of 365.24219 days, which means
// the Julian calendar gains a day every 128 years.
//
// During the 20th and 21st centuries, a date according to the Julian calendar
// is 13 days earlier than its corresponding Gregorian date."
//
// Wikipedia https://en.wikipedia.org/wiki/Julian_calendar
//
//
// "Augustus corrected errors in the observance of leap years by omitting leap
// days until AD 8. Julian calendar dates before March AD 4 are proleptic, and
// do not necessarily match the dates actually observed in the Roman Empire."
//
// Nautical almanac offices of the United Kingdom and United States, 1961, p. 411"
//
// Conversion between Julian and Gregorian calendars:
//  https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars
//
// This method uses the 'Richards' algorithm to convert Julian Day Number and
// Times to the Julian Calendar.
//
// Reference:
//   Richards, E. G. (1998). Mapping Time: The Calendar and its History.
//   Oxford University Press. ISBN 978-0192862051
//
//   Wikipedia - Julian Day
//   https://en.wikipedia.org/wiki/Julian_day
//
// The Julian Calendar date time returned by this method is generated from
// the Julian Day Number. The Julian Day Number (JDN) is the integer assigned
// to a whole solar day in the Julian day count starting from noon Universal
// time, with Julian day number 0 assigned to the day starting at noon on
// Monday, January 1, 4713 BC, in the proleptic Julian calendar and November
// 24, 4714 BC, in the proleptic Gregorian calendar.
//
// The Julian Day Number Time is a floating point number with an integer
// to the left of the decimal point representing the Julian Day Number
// and the fraction to the right of the decimal point representing time
// in hours minutes and seconds.
//
// Julian Day numbers start on day zero at noon. This means that Julian
// Day Number Times are valid for all dates on or after noon on Monday,
// January 1, 4713 BCE, in the proleptic Julian calendar or November 24,
// 4714 BCE, in the proleptic Gregorian calendar. Remember that the Golang
// 'time.Time' type uses Astronomical Year numbering with the Gregorian
// Calendar. In other words, the 'time.Time' type recognizes the year
// zero. Dates expressed in the 'Common Era' ('BCE' Before Common Era
// or 'CE' Common Era). Therefore a 'time.Time' year of '-4713' is equal
// to the year '4714 BCE'
//
// This means that the 'Richards' algorithm employed by this method is valid
// for all 'time.Time' (possibly proleptic) Julian dates on or after noon
// November 24, −4713 (Gregorian Calendar proleptic).
//
// For information on the Julian Day Number/Time see:
//   https://en.wikipedia.org/wiki/Julian_day
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  julianDayNoNoTime   float64
//     - The integer portion of this number (digits to left of
//       the decimal) represents the Julian day number. The fractional
//       digits to the right of the decimal represent elapsed time
//       since noon on the Julian day number. All time values are
//       expressed as Universal Coordinated Time (UTC).
//
//
//  digitsAfterDecimal  int
//     - The number of digits after the decimal in input parameter
//       'julianDayNoNoTime' which will be used in the conversion
//       algorithm. Effectively, 'julianDayNoNoTime' will be rounded
//       to the number of digits to the right of the decimal specified
//       in this parameter.
//
//
//  ePrefix             string
//     - A string containing the names of the calling functions
//       which invoked this method. The last character in this
//       string should be a blank space.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  julianDateUtc    time.Time
//     - The returned parameter 'gregorianDateTime' represents the input
//       'julianDayNoNoTime' converted to the Gregorian calendar. This
//       returned 'time.Time' type is always configured as Universal
//       Coordinated Time (UTC). In addition, as a Golang 'time.Time'
//       type, the date is expressed using astronomical years. Astronomical
//       year numbering includes a zero year. Therefore, 1BCE is stored
//       as year zero in this return value.
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered this error Type will encapsulate
//       an error message.
//
//
// ------------------------------------------------------------------------
//
// Resources
//
//  Julian Day Wikipedia
//  https://en.wikipedia.org/wiki/Julian_day
//
//  PHP Julian date converter algorithms (Stack Overflow)
//   https://stackoverflow.com/questions/45586444/php-julian-date-converter-algorithms
//
//
func (calUtil *CalendarUtility) JulianDayNoTimeToJulianCalendar(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	julianDateTimeUtc time.Time,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = new(sync.Mutex)
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "DTimeUtility.JulianDayNoTimeToJulianCalendar() "

	calendarMech := calendarMechanics{}

	julianDateTimeUtc,
		err =
		calendarMech.richardsJulianDayNoTimeToJulianCalendar(
			julianDayNoDto,
			ePrefix)

	return julianDateTimeUtc, err
}