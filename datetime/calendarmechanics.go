package datetime

import (
	"math"
	"sync"
	"time"
)

// calendarMechanics
//
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
// Astronomical Year Numbering
//
// Proleptic Gregorian
//
//
type calendarMechanics struct {
	input  string
	output string
	lock   sync.Mutex
}

// gregorianDateToJulianDayNo - Converts a Gregorian calendar
// date to Julian Day Number (JDN).
//
// This algorithm is valid for all (possibly proleptic) Gregorian
// calendar dates after November 23, −4713. Divisions are integer
// divisions, fractional parts are ignored.
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
func (calMech *calendarMechanics) gregorianDateToJulianDayNo(
	gregorianDateTime time.Time,
	ePrefix string) (
	gregorianDateUtc time.Time,
	julianDayNo int64,
	err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calMech.gregorianDateToJulianDayNo() "

	gregorianDateUtc = gregorianDateTime.UTC()
	julianDayNo = 0
	err =  nil

	// This algorithm is valid for all (possibly proleptic) Gregorian
	// calendar dates after November 23, −4713.
	limitDateTime := time.Date(
		-4713,
		11,
		24,
		0,
		0,
		0,
		0,
		time.UTC)

	if gregorianDateUtc.Before(limitDateTime) {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "gregorianDateTime",
			inputParameterValue: "",
			errMsg:              "Error: 'gregorianDateTime' is invalid. Valid date times must\n" +
				"occur after November 23, −4713 UTC as expressed in Astronomical\n" +
				"year numbering.",
			err:                 nil,
		}

		return gregorianDateUtc, julianDayNo, err
	}

	Year := int64(gregorianDateUtc.Year())
	Month := int64(gregorianDateUtc.Month())
	Day := int64(gregorianDateUtc.Day())

	/*
		JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075
	*/

	julianDayNo =
		(int64(1461) * (Year + int64(4800) +
		(Month - int64(14))/int64(12)))/int64(4) +
		(int64(367) * (Month - int64(2) -
		int64(12) * ((Month - int64(14))/int64(12))))/int64(12) -
		(int64(3) * ((Year + int64(4900) +
		(Month - int64(14))/int64(12))/int64(100)))/int64(4) +
		Day - int64(32075)

	gregorianTimeNanoSecs := int64(gregorianDateUtc.Hour()) * HourNanoSeconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Minute()) * MinuteNanoSeconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Second()) * SecondNanoseconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Nanosecond())

	if gregorianTimeNanoSecs < NoonNanoSeconds {
		julianDayNo -= 1
	}

	return gregorianDateUtc, julianDayNo, err
}

// gregorianDateToJulianDate - Converts a Gregorian Date to a Julian
// Day Number and Time.
//
// Remember that Julian Day Number Times are valid for all dates
// after noon on Monday, January 1, 4713 BC, proleptic Julian calendar
// or November 24, 4714 BC, in the proleptic Gregorian calendar. Therefore,
// using astronomical years encapsulated in the Golang type time.Time,
// this algorithm is valid for all Golang date/times after Gregorian
// calendar (possibly proleptic) values after November 23, −4713.
//
// Reference Wikipedia:
//   https://en.wikipedia.org/wiki/Julian_day
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
//  digitsAfterDecimal  int
//     - The number of digits after the decimal in returned
//       by this method in parameter, 'julianDayNoNoTime'.
//       Effectively, the returned value, 'julianDayNoNoTime',
//       will be rounded to the number of digits to the right
//       of the decimal specified in this parameter.
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
//       the Julian Day Number (JDN)
//
//
//  julianDayNoNoTime   float64
//     - The integer portion of this number (digits to left of
//       the decimal) represents the Julian day number. The fractional
//       digits to the right of the decimal represent elapsed time
//       since noon on the Julian day number. All time values are
//       expressed as Universal Coordinated Time (UTC).
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered this error Type will encapsulate
//       an error message.
//
//
func (calMech *calendarMechanics) gregorianDateToJulianDate(
	gregorianDateTime time.Time,
	digitsAfterDecimal int,
	ePrefix string) (
	gregorianDateUtc time.Time,
	julianDayNoTime float64,
	err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calMech.gregorianDateToJulianDate() "

	gregorianDateUtc = gregorianDateTime.UTC()
	julianDayNoTime = 0.0
	err = nil

	if digitsAfterDecimal < 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "digitsAfterDecimal",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'digitsAfterDecimal' is " +
				"less than ZERO!",
			err:                 nil,
		}

		return gregorianDateUtc, julianDayNoTime, err
	}

	if digitsAfterDecimal > 100 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "digitsAfterDecimal",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'digitsAfterDecimal' is " +
				"greater than 100!",
			err:                 nil,
		}

		return gregorianDateUtc, julianDayNoTime, err
	}

	Year := int64(gregorianDateUtc.Year())
	Month := int64(gregorianDateUtc.Month())
	Day := int64(gregorianDateUtc.Day())

	/*
		JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075
	*/

	julianDayNo :=
		(int64(1461) * (Year + int64(4800) +
			(Month - int64(14))/int64(12)))/int64(4) +
			(int64(367) * (Month - int64(2) -
				int64(12) * ((Month - int64(14))/int64(12))))/int64(12) -
			(int64(3) * ((Year + int64(4900) +
				(Month - int64(14))/int64(12))/int64(100)))/int64(4) +
			Day - int64(32075)

	gregorianTimeNanoSecs := int64(gregorianDateUtc.Hour()) * HourNanoSeconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Minute()) * MinuteNanoSeconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Second()) * SecondNanoseconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Nanosecond())

	if gregorianTimeNanoSecs < NoonNanoSeconds {

		julianDayNo -= 1
		gregorianTimeNanoSecs += NoonNanoSeconds

	} else {

		gregorianTimeNanoSecs -= NoonNanoSeconds
	}

	julianDayNoTime = float64(julianDayNo)

	julianDayNoTime += float64(gregorianTimeNanoSecs) /
									float64(DayNanoSeconds)

	multiplier := math.Pow10(digitsAfterDecimal)

	julianDayNoTime = julianDayNoTime * multiplier

	julianDayNoTime = math.Round(julianDayNoTime)

	julianDayNoTime = julianDayNoTime / multiplier

	return gregorianDateUtc, julianDayNoTime, err
}


// richardsJulianDayNoTimeToGregorianDateTime - Converts a Julian Day
// Number and Time value to the corresponding date time in the
// Gregorian Calendar. Because the Gregorian Calendar was instituted
// in Friday, October 15, 1582, all Gregorian Calendar dates prior
// to this are extrapolated or proleptic.
//
// "This is an algorithm by E. G. Richards to convert a Julian Day Number,
// J, to a date in the Gregorian calendar (proleptic, when applicable).
// Richards states the algorithm is valid for Julian day numbers greater
// than or equal to 0".
//
// Reference:
//   Richards, E. G. (1998). Mapping Time: The Calendar and its History.
//   Oxford University Press. ISBN 978-0192862051
//
// The Julian Day Number (JDN) is the integer assigned to a whole solar
// day in the Julian day count starting from noon Universal time, with
// Julian day number 0 assigned to the day starting at noon on Monday,
// January 1, 4713 BC, in the proleptic Julian calendar and November 24,
// 4714 BC, in the proleptic Gregorian calendar.
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
//
func (calMech *calendarMechanics) richardsJulianDayNoTimeToGregorianDateTime(
	julianDayNoNoTime float64,
	digitsAfterDecimal int,
	ePrefix string) (
	gregorianDateUtc time.Time,
	err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calMech.JulianDayNoTimeToGregorianDateTime() "

	gregorianDateUtc = time.Time{}
	err = nil

	if julianDayNoNoTime < 0.000 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNoNoTime",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'julianDayNoNoTime' " +
				"is less than Zero!",
			err:                 nil,
		}
		return gregorianDateUtc, err
	}

	if digitsAfterDecimal < 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "digitsAfterDecimal",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'digitsAfterDecimal' is " +
				"less than ZERO!",
			err:                 nil,
		}

		return gregorianDateUtc, err
	}

	if digitsAfterDecimal > 100 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "digitsAfterDecimal",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'digitsAfterDecimal' is " +
				"greater than 100!",
			err:                 nil,
		}

		return gregorianDateUtc, err
	}

	y := int64(4716)
	j := int64(1401)
	m := int64(2)
	n := int64(12)
	r := int64(4)
	p := int64(1461)
	v := int64(3)
	u := int64(5)
	s := int64(153)
	w := int64(2)
	B := int64(274277)
	C := int64(-38)


	julianDayNumInt, julianDayNumFrac := math.Modf(julianDayNoNoTime)

	// Julian Day No as integer
	J := int64(julianDayNumInt)

	f := J + j + ((((4 * J + B) / 146097) * 3) /4) + C

	e := r * f + v // #2

	g := (e % p) / r // #3

	h := u * g + w // #4

	D := ((h % s) / u) + 1  // #5

	M := ((( h / s) + m) % n) + 1  // #6

	Y := (e / p) - y + ((n + m - M)/ n)

	totalNanoSeconds :=
		(int64(julianDayNumFrac * float64(DayNanoSeconds))) +
			NoonNanoSeconds

	hours := 0
	minutes := 0
	seconds := 0
	nanoseconds := 0

	if totalNanoSeconds >= HourNanoSeconds {
		hours = int(totalNanoSeconds / HourNanoSeconds)

		totalNanoSeconds -= int64(hours) * HourNanoSeconds
	}

	if totalNanoSeconds >= MinuteNanoSeconds {
		minutes = int(totalNanoSeconds / MinuteNanoSeconds)

		totalNanoSeconds -= int64(minutes) * MinuteNanoSeconds
	}

	if totalNanoSeconds >= SecondNanoseconds {

		seconds = int(totalNanoSeconds / SecondNanoseconds)

		totalNanoSeconds -= int64(seconds) * SecondNanoseconds
	}

	if totalNanoSeconds >= HalfSecondNanoseconds {
		seconds += 1
	}

	nanoseconds = 0

	gregorianDateUtc = time.Date(
		int(Y),
		time.Month(M),
		int(D),
		hours,
		minutes,
		seconds,
		nanoseconds,
		time.UTC)

	return gregorianDateUtc, err
}