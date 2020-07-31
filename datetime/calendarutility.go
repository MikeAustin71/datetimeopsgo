package datetime

import (
	"errors"
	"fmt"
	"math"
	"math/big"
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
// accurate to subMicrosecondNanoseconds.
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
		calendarMech.julianDayNoTimeToGregorianCalendar(
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


// GetCompleteYearInterval - Computes the interval of whole years between a base date/time
// and a target date/time. To guarantee complete accuracy, the input DateTransferDto
// objects must contain the correct setting for the internal member variable 'isLeapYear'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// baseDateTimeDto          DateTransferDto
//     - Specifies the base date time. This value will be compared
//       with parameter 'targetDateDto' to compute the number of
//       whole or complete years between the two date times.
//
//
// targetDateTimeDto        DateTransferDto
//     - Specifies the target date time. This value will be compared
//       with parameter 'baseDateTimeDto' to compute the number of
//       whole or complete years between the two date times.
//
//
//  ePrefix                 string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  completedYearsInterval  int64
//     - If this method completes successfully, this value will contain
//       the total number of consecutive whole years between the date
//       times specified by input parameters, 'baseDateTimeDto' and
//       'targetDateTimeDto'.
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calUtil *CalendarUtility) GetCompleteYearInterval(
	baseDateTimeDto DateTransferDto,
	targetDateTimeDto DateTransferDto,
	ePrefix string) (completedYearsInterval int64, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	completedYearsInterval = -999999999

	err = nil

	ePrefix += "CalendarUtility.NewGetCompleteYearInterval() "

	if !baseDateTimeDto.IsValidInitialize() {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'baseDateTimeDto' is INVALID!\n")
		return completedYearsInterval, err
	}

	if !targetDateTimeDto.IsValidInitialize() {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'targetDateTimeDto' is INVALID!\n")
		return completedYearsInterval, err
	}

	if baseDateTimeDto.GetYear() == targetDateTimeDto.GetYear() {
		// Years are equal. Nothing to do.
		completedYearsInterval = 0
		return completedYearsInterval, err
	}

	var baseTargetComparisonResult int

	var tempBaseDateTimeDto, tempTargetDateTimeDto DateTransferDto

	tempBaseDateTimeDto = baseDateTimeDto.CopyOut()
	tempTargetDateTimeDto = targetDateTimeDto.CopyOut()

	baseTargetComparisonResult, err = tempBaseDateTimeDto.Compare(&tempTargetDateTimeDto, ePrefix)

	if err != nil {
		return completedYearsInterval, err
	}

	if baseTargetComparisonResult == 1 {
		tempBaseDateTimeDto.ExchangeValues(&tempTargetDateTimeDto)
	}

	fmt.Println()
	fmt.Printf("  Temp Base Year: %v\n", tempBaseDateTimeDto.GetYear())
	fmt.Printf("Temp Target Year: %v\n", tempTargetDateTimeDto.GetYear())

	// baseYear is now less than targetYear

	// Now determine whether base year is a full year
	// or partial year and make necessary adjustments.
	//

	completedYearsInterval =
		tempTargetDateTimeDto.GetYear() -
			tempBaseDateTimeDto.GetYear() - 1


	return completedYearsInterval, err
}

// GetElapsedWholeOrdinalDaysInYear - Returns the number of elapsed
// 24-hour days in a year based on specified month, day, hour, minute,
// second and nanosecond.
//
func (calUtil *CalendarUtility) GetElapsedWholeOrdinalDaysInYear(
	isLeapYear bool,
	month,
	day int,
	hour int,
	minute int,
	second int,
	nanosecond int,
	ePrefix string) (elapsedCompleteDays int, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	elapsedCompleteDays = math.MinInt32
	err = nil

	ePrefix += "CalendarUtility.GetElapsedWholeOrdinalDaysInYear() "

	calMech2 := CalendarUtility{}

	var ordinalDayNumber int


	ordinalDayNumber, err = calMech2.GetOrdinalDayNumber(
		isLeapYear,
		month,
		day,
		ePrefix)

	if err != nil {
		return elapsedCompleteDays, err
	}

	isFullDay := false

	if hour == 23 &&
		minute == 59 &&
		(second == 59 || second == 60) &&
		nanosecond== 999999999 {
		isFullDay = true
	}

	if ordinalDayNumber > 0 {
		if isFullDay {
			elapsedCompleteDays = ordinalDayNumber
		} else {
			// This is a partial day.
			elapsedCompleteDays = ordinalDayNumber - 1
		}

	} else {
		elapsedCompleteDays = 0
	}


	return elapsedCompleteDays, err
}

// GetJulianDayNoTimeFraction - Computes the Julian Day Number time
// fraction. The Julian Day starts a 12:00:00-hours or Noon.
//
// For more information on the Julian Day Number, reference:
//   https://en.wikipedia.org/wiki/Julian_day
//
//
// Note: If the 'second' value is set to '60' a leap second is assumed
// and the calculation is adjusted accordingly.
//
// For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second              int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of 'leap seconds'.
//       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond          int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  julianDayNoTimeFraction   *big.Float
//     - If successful this method will return the Julian Day Number
//       time fraction as a floating point value. Obviously, the value
//       only contains the time fraction and therefore does NOT contain
//       the integer portion of the Julian Day Number.
//
//  julianDayNoAdjustment     int64
//     - If successful this method returns an integer value signaling
//       whether an adjustment is the integer Julian Day Number is
//       required. If the time value represented by the input parameters
//       occurs before Noon (12:00:00.000000000), it signals that the
//       moment occurred on the previous Julian Day Number. Therefore,
//       if the time value occurs before Noon (12:00:00.000000000), this
//       return value is set to -1 (minus one). Otherwise, the return
//       value is set to zero.
//
//  err                       error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calUtil *CalendarUtility) GetJulianDayNoTimeFraction(
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	julianDayNoTimeFraction *big.Float,
	julianDayNoAdjustment int64,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "CalendarUtility.GetJulianDayNoTimeFraction() "

	julianDayNoTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	julianDayNoAdjustment = 0

	err = nil

	calMech2 := CalendarUtility{}

	var totalDateTimeNanoSecs, noonNanoSecs int64

	noonNanoSecs = int64(time.Hour) * 12

	totalDateTimeNanoSecs, err =
		calMech2.GetTimeTotalNanoseconds(
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

	if totalDateTimeNanoSecs == noonNanoSecs {
		// Fraction is zero.
		return julianDayNoTimeFraction, julianDayNoAdjustment, err
	}

	if totalDateTimeNanoSecs  < noonNanoSecs {
		julianDayNoAdjustment = -1

		totalDateTimeNanoSecs += noonNanoSecs

	} else {
		// totalDateTimeNanoSecs  > noonNanoSecs
		julianDayNoAdjustment = 0
		totalDateTimeNanoSecs -= noonNanoSecs
	}

	var twentyFourHourNanoseconds *big.Float

	if second == 60 {

		twentyFourHourNanoseconds =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64((int64(time.Hour) * 24) + 1)

	} else {

		twentyFourHourNanoseconds =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(int64(time.Hour) * 24)
	}

	actualNanoSec :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(totalDateTimeNanoSecs)

	julianDayNoTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(actualNanoSec, twentyFourHourNanoseconds)

	return julianDayNoTimeFraction, julianDayNoAdjustment, err
}

// GetStandardDayTimeFraction - Computes the time for a standard
// day as a fractional value.
//
// The time value represented by the input parameters is converted
// to total nanoseconds and divided by the number of nanoseconds
// in a standard 24-hour day.
//
// The month and day input parameters are used to validate 'leap seconds'.
//
// Important: If the 'second' value is set to '60' a leap second is assumed
// and the total number of nanoseconds is divided by 24-hours plus one
// second.
//
// For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  month              int
//     - The month time component for this time value.
//       The valid range is 1 - 12 inclusive. This parameter
//       is used to validate 'leap seconds'.
//
//
//  day                int
//     - The day time component for this time value.
//       The valid range is 1 - 31 inclusive. This parameter
//       is used to validate 'leap seconds'.
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of leap seconds.
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  julianDayNoTimeFraction   *big.Float
//     - If successful this method will return the Julian Day Number
//       time fraction as a floating point value. Obviously, the value
//       only contains the time fraction and therefore does NOT contain
//       the integer portion of the Julian Day Number.
//
//  err                       error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calUtil *CalendarUtility) GetStandardDayTimeFraction(
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	standardDayTimeFraction *big.Float,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	standardDayTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	ePrefix += "CalendarUtility.GetStandardDayTimeFraction() "

	if month < 0 || month > 12 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'month' is INVALID!\n" +
			"month='%v'\n", month)

		return standardDayTimeFraction, err
	}

	if day < 1 || day > 31 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'day' is INVALID!\n" +
			"day='%v'\n", day)

		return standardDayTimeFraction, err
	}

	if hour < 0 || hour > 23 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'hour' is INVALID!\n" +
			"hour='%v'\n", hour)

		return standardDayTimeFraction, err
	}

	if minute < 0 || minute > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'minute' is INVALID!\n" +
			"minute='%v'\n", minute)

		return standardDayTimeFraction, err
	}

	if nanosecond < 0 || nanosecond > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'nanosecond' is INVALID!\n" +
			"nanosecond='%v'\n", nanosecond)

		return standardDayTimeFraction, err
	}

	calUtil2 := CalendarUtility{}

	isValidSecond := calUtil2.IsValidSecond(
		month,
		day,
		hour,
		minute,
		second)

	if !isValidSecond {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'nanosecond' is INVALID!\n" +
			"nanosecond='%v'\n", nanosecond)

		return standardDayTimeFraction, err
	}

	var totalDateTimeNanoSecs int64

	totalDateTimeNanoSecs, err =
		calUtil2.GetTimeTotalNanoseconds(
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

	var twentyFourHourNanoseconds *big.Float

	if second == 60 {

		twentyFourHourNanoseconds =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64((int64(time.Hour) * 24) + 1)

	} else {

		twentyFourHourNanoseconds =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(int64(time.Hour) * 24)
	}

	actualNanoSec :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(totalDateTimeNanoSecs)

	standardDayTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(actualNanoSec, twentyFourHourNanoseconds)

	return standardDayTimeFraction, err
}

// GetMonthDayFromOrdinalDayNo - Receives an Ordinal Day Number and returns
// the associated month and day number. The input parameter 'isLeapYear'
// specifies whether the Ordinal Day Number is included in a standard year
// (365-Days) or a Leap Year (366-Days).
//
// The return value 'yearAdjustment' this value will be populated with one
// of two values: Zero (0) or Minus One (-1). A value of Zero signals that
// no prior year adjustment is required. A value of Minus One indicates that
// the ordinal date, passed as an input parameter, represents December 31st
// of the prior year. In this event, an adjustment to the original year value
// is necessary.
//
// For more information on Ordinal Date, reference:
//    https://en.wikipedia.org/wiki/Ordinal_date
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  ordinalDate        int64
//     - A value with a valid range of 1-366 inclusive which specifies
//       the day of the year expressed as an ordinal day number or ordinal
//       date. For more information on 'ordinal dates', reference:
//         https://en.wikipedia.org/wiki/Ordinal_date
//
//
//  isLeapYear         bool
//     - If 'true' it signals that the input parameter 'ordinalDate' represents
//       an ordinal date within a leap year.
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// yearAdjustment     int
//     - If the method completes successfully, this value will
//       be populated with one of three values:
//
//           One (+1)       -
//             A value of plus one (+1) signals that the ordinal day
//             day number represents the first day of the next year
//             (January 1st of the next year). Effectively the year
//             value is equal to year + 1, month is equal to '1' and
//             the day value is equal to '1'.
//
//           Zero (0)       -
//             A value of Zero signals that no year adjustment is
//             required. The ordinal day number was converted to
//             month and day number in the current year.
//
//           Minus One (-1) -
//             A value of Minus One indicates that the ordinal date
//             represents December 31st of the prior year. 'year'
//             is therefore equal to year - 1, month = 12 and day = 31.
//
//  month             int
//     - If the method completes successfully, this value
//       will contain the month number associated with the
//       input parameter 'ordinalDate'.
//
//
//  day                int
//     - If successful this value will contain the day number
//       associated with the input parameter, 'ordinalDate'.
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calUtil *CalendarUtility) GetMonthDayFromOrdinalDayNo(
	ordinalDate int64,
	isLeapYear bool,
	ePrefix string)(
	yearAdjustment int,
	month int,
	day int,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "CalendarUtility.GetMonthDayFromOrdinalDayNo() "

	yearAdjustment = math.MinInt32
	month = -1
	day = -1
	err = nil

	if isLeapYear {

		if ordinalDate == 367 {
			month = 1
			day = 1
			yearAdjustment = 1
			return yearAdjustment, month, day, err
		}

		if ordinalDate < 0 ||
			ordinalDate > 366 {
			err = fmt.Errorf("\n" +ePrefix + "Error:\n" +
				"Input Parameter 'ordinalDate' is INVALID!\n" +
				"ordinalDate='%v'\n",
				ordinalDate)
			return yearAdjustment, month, day, err
		}

	} else {
		// This is NOT a leap year

		if ordinalDate == 366 {
			month = 1
			day = 1
			yearAdjustment = 1
			return yearAdjustment, month, day, err
		}

		if ordinalDate < 0 ||
			ordinalDate > 365 {
			err = fmt.Errorf("\n" +ePrefix + "Error:\n" +
				"Input Parameter 'ordinalDate' is INVALID!\n" +
				"ordinalDate='%v'\n",
				ordinalDate)
			return yearAdjustment, month, day, err
		}

	}


	if ordinalDate == 0 {
		yearAdjustment = -1
		month = 12
		day = 31
		return yearAdjustment, month, day, err
	}

	ordDays := []int64 {
		0,
		31,
		59,
		90,
		120,
		151,
		181,
		212,
		243,
		273,
		304,
		334,
	}

	mthDays := []int {
		-1,
		31,
		28,
		31,
		30,
		31,
		30,
		31,
		31,
		30,
		31,
		30,
		31,
	}

	for i:=11; i > -1; i-- {

		if ordDays[i] <= ordinalDate {

			ordinalDate -= ordDays[i]

			month = i + 1

			if month > 2 && isLeapYear {
				ordinalDate--
			}

			if ordinalDate == 0 &&
				month==1 {
				yearAdjustment = -1
				month = 12
				day = 31

				break

			} else if ordinalDate == 0 {
				month--
				day = mthDays[month]
				break
			} else {
				day = int(ordinalDate)
				break
			}

		}
	}

	return yearAdjustment, month, day, err
}

// GetOrdinalDayNumber - Computes the ordinal day number
// for any given month and day. Input parameter 'isLeapYear'
// indicates whether the year encompassing the specified
// month and day is a 'leap year' containing 366-days
// instead of the standard 365-days.
//
// Reference
//    https://en.wikipedia.org/wiki/Ordinal_date
//
func (calUtil *CalendarUtility) GetOrdinalDayNumber(
	isLeapYear bool,
	month int,
	day int,
	ePrefix string) (ordinalDayNo int, err error) {

	ePrefix += "CalendarUtility.GetOrdinalDayNumber() "

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ordinalDayNo = -1
	err = nil

	if month < 1 || month > 12 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input Parameter 'month' is INVALID!\n" +
			"month='%v'\n", month)
		return ordinalDayNo, err
	}


	ordDays := []int {
		0,
		31,
		59,
		90,
		120,
		151,
		181,
		212,
		243,
		273,
		304,
		334,
	}

	mthDays := []int {
		-1,
		31,
		28,
		31,
		30,
		31,
		30,
		31,
		31,
		30,
		31,
		30,
		31,
	}

	if month == 1 &&
		day == 0 {
		ordinalDayNo = 0
		return ordinalDayNo, err
	}

	monthDays := mthDays[month]

	if month==2 && isLeapYear {
		monthDays++
	}

	if day > monthDays || day < 1 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input parameter 'day' is INVALID!\n" +
			"month='%v'\n day='%v'\n",
			month, day)
		return ordinalDayNo, err
	}


	if month == 1 {
		ordinalDayNo = day
		return ordinalDayNo, err
	} else {

		ordinalDayNo = ordDays[month-1] + day

		if isLeapYear && month > 2 {
			ordinalDayNo++
		}
	}

	return ordinalDayNo, err
}


// GetTimeTotalNanoseconds - Computes the total time in nanoseconds for
// a given time of day expressed in Hours, Minutes, Seconds, and Nanoseconds.
//
func (calUtil *CalendarUtility) GetTimeTotalNanoseconds(
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (totalTimeNanoseconds int64, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "CalendarUtility.GetTimeTotalNanoseconds() "

	totalTimeNanoseconds = math.MaxInt64
	err = nil

	if hour > 24 || hour < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hour' is invalid!\n" +
			"hour='%v'\n", hour)
		return totalTimeNanoseconds, err
	}

	if minute > 59 || minute < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minute' is invalid!\n" +
			"minute='%v'\n", minute)
		return totalTimeNanoseconds, err
	}

	// Watch out for leap seconds
	if second > 60 || second < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'second' is invalid!\n" +
			"second='%v'\n", second)
		return totalTimeNanoseconds, err
	}

	if nanosecond > 999999999 || nanosecond < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecond' is invalid!\n" +
			"nanosecond='%v'\n", nanosecond)
		return totalTimeNanoseconds, err
	}

	totalTimeNanoseconds = 0

	totalTimeNanoseconds += int64(hour) * int64(time.Hour)

	totalTimeNanoseconds += int64(minute) * int64(time.Minute)

	totalTimeNanoseconds += int64(second) * int64(time.Second)

	totalTimeNanoseconds +=  int64(nanosecond)

	if totalTimeNanoseconds > int64(time.Hour) * 24 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Total Time in nanoseconds exceeds 24-hours\n" +
			"Total Nanosecond='%v'\n", totalTimeNanoseconds)
	}

	return totalTimeNanoseconds, err
}

// IsLastNanosecondBeforeMidnight - Determines if a date time is precisely
// equivalent to the nanosecond before Midnight:
//       (23:59:59.999999999).
//
// This moment is literally the last nanosecond of the current day.
//
// In addition this method will validate if a leap second
// if it occurs on the correct month and day:
//       (23:59:60.999999999)
//
// If the combination of date time components is invalid,
// this method will return an error.
//
// If this method returns true for month=12 and day=31 it means that
// this date/time marks the end of a complete year.
//
// For a comparison, see the companion method 'IsMidnight()'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
// isLeapYear          bool
//     - Used to validates days and leap seconds.
//
//
//  month              int
//     - The month time component for this time value.
//       The valid range is 1 - 12 inclusive.
//
//
//  day                int
//     - The day time component for this time value.
//       The valid range is 1 - 31 inclusive.
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of leap seconds. The
//       calling method is responsible for supplying the
//       correct 'second' or 'leap second' value.
//
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isLastNanosecondBeforeMidnight  bool
//     - If successful this boolean flag signals whether the date/time
//       passed by the input parameters precisely represents the last
//       nanosecond before midnight of the next day. As such this means
//       that the passed date/time is literally the last nanosecond of
//       the complete current day.
//
//  err                             error
//     - If successful, and the date/time values specified by the input
//       parameters are valid, the returned error Type is set equal to
//       'nil'.
//
//       If errors are encountered during processing, or if the date/time
//       values specified by the input parameters are invalid, the returned
//       error Type will encapsulate an appropriate error message. Note this
//       error message will be prefixed with the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calUtil *CalendarUtility) IsLastNanosecondBeforeMidnight(
	isLeapYear bool,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	isLastNanosecondBeforeMidnight bool, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	isLastNanosecondBeforeMidnight = false

	err = nil

	ePrefix += "CalendarUtility.IsLastNanosecondBeforeMidnight() "

	calUtil2 := CalendarUtility{}

	err = calUtil2.IsValidDateTimeComponents(
		isLeapYear,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		ePrefix)

	if err != nil {
		return isLastNanosecondBeforeMidnight, err
	}

	if hour == 23 &&
		minute == 59 &&
		(second == 59 || second==60) &&
		nanosecond == 999999999 {
		isLastNanosecondBeforeMidnight = true
	}

	return isLastNanosecondBeforeMidnight, err
}

// IsMidnight - Determines if a date time is precisely
// equivalent to Midnight, the beginning of the current
// day: (00:00:00.000000000).
//
// This moment is literally the first nanosecond of the
// current day.
//
// If the combination of date time components is invalid,
// this method will return an error.
//
// If this method returns true for month=1 and day=1 it means that
// this date/time marks the beginning of the current year.
//
// For a comparison, see the companion method:
//     'IsLastNanosecondBeforeMidnight()'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
// isLeapYear          bool
//     - Used to validates days and leap seconds.
//
//
//  month              int
//     - The month time component for this time value.
//       The valid range is 1 - 12 inclusive.
//
//
//  day                int
//     - The day time component for this time value.
//       The valid range is 1 - 31 inclusive.
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of leap seconds. The
//       calling method is responsible for supplying the
//       correct 'second' or 'leap second' value.
//
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isLastNanosecondBeforeMidnight  bool
//     - If successful this boolean flag signals whether the date/time
//       passed by the input parameters precisely represents the last
//       nanosecond before midnight of the next day. As such this means
//       that the passed date/time is literally the last nanosecond of
//       the complete current day.
//
//  err                             error
//     - If successful, and the date/time values specified by the input
//       parameters are valid, the returned error Type is set equal to
//       'nil'.
//
//       If errors are encountered during processing, or if the date/time
//       values specified by the input parameters are invalid, the returned
//       error Type will encapsulate an appropriate error message. Note this
//       error message will be prefixed with the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calUtil *CalendarUtility) IsMidnight(

	isLeapYear bool,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	isMidnight bool, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	isMidnight = false

	err = nil

	ePrefix += "CalendarUtility.IsMidnight() "

	calUtil2 := CalendarUtility{}

	err = calUtil2.IsValidDateTimeComponents(
		isLeapYear,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		ePrefix)

	if err != nil {
		return isMidnight, err
	}

	if hour == 0 &&
		minute == 0 &&
		second == 0 &&
		nanosecond == 0 {
		isMidnight = true
	}

	return isMidnight, err
}

// IsValidSecond - Returns true if the 'second' qualifies
// as a valid 'leap second'.
//
// The typical range for a 'second' value is:
//    Greater Than or Equal to zero
//         AND
//    Less Than or Equal to '59'
//
// In addition to applying the standard criteria for a
// 'second', this method also tests to determine whether
// 'second' qualifies as a valid 'leap second'. Valid
// 'leap seconds' may have a value of '60'.
//
// A valid leap second must meet the following criteria:
//
//  (1) The value must be equal to '60'.
//
//  (2) The leap second insertion must occur in June,
//      December, March or September on the last day of the
//      month.
//
//  (3) The leap second must occur at the last second of the
//      the day. Example: 23:59:60.
//
//  (4) The time value of the input parameters is assumed to
//      represent Coordinated Universal Time (UTC). UTC leap
//      seconds occur simultaneously worldwide.
//
//      For more information on Coordinated Universal Time,
//      reference:
//      https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// For more information on 'leap second' reference:
//  https://en.wikipedia.org/wiki/Leap_second
//
func (calUtil *CalendarUtility) IsValidSecond(
	month,
	day,
	hour,
	minute,
	second int) bool {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	if second >= 0 && second <= 59 {
		return true
	}

	// Check to see if this is a valid
	// leap second.
	if month == 3 &&
		day == 31 &&
		hour == 23 &&
		minute == 59 &&
		second == 60 {
		return true
	}

	if month == 6 &&
		day == 30 &&
		hour == 23 &&
		minute == 59 &&
		second == 60 {
		return true
	}

	if month == 9 &&
		day == 30 &&
		hour == 23 &&
		minute == 59 &&
		second == 60 {
		return true
	}

	if month == 12 &&
		day == 31 &&
		hour == 23 &&
		minute == 59 &&
		second == 60 {
		return true
	}

	return false
}

// IsValidDateTimeComponents - Returns a boolean flag signaling
// whether the date time values passed as input parameters are
// valid. To guarantee the accuracy of the result, the user must
// provide the true and correct value for parameter 'isLeapYear'.
//
// If the date time parameters are judged to be 'invalid', this
// method will return a type 'error' with an appropriate error
// message.
//
// Note that the 'year' value is not included in this validation.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
// isLeapYear          bool
//     - Used to validates days and leap seconds.
//
//
//  month              int
//     - The month time component for this time value.
//       The valid range is 1 - 12 inclusive.
//
//
//  day                int
//     - The day time component for this time value.
//       The valid range is 1 - 31 inclusive.
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of leap seconds.
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                error
//     - If successful, and the date/time values specified by the input
//       parameters are valid, the returned error Type is set equal to
//       'nil'.
//
//       If errors are encountered during processing, or if the date/time
//       values specified by the input parameters are invalid, the returned
//       error Type will encapsulate an appropriate error message. Note this
//       error message will be prefixed with the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calUtil *CalendarUtility) IsValidDateTimeComponents(
	isLeapYear bool,
	month int,
	day int,
	hour int,
	minute int,
	second int,
	nanosecond int,
	ePrefix string) ( err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "CalendarUtility.IsValidDateTimeComponents() "

	if month < 1 || month > 12 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'month' is invalid!\n" +
			"month='%v'\n", month)
		return err
	}

	if day > 31 || day < 1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'day' is invalid!\n" +
			"day='%v'\n", day)
		return err
	}

	if month == 2 {
		// This is February - check it out!
		if isLeapYear &&
			day > 29 {

			err = fmt.Errorf(ePrefix + "\n" +
				"Error: Input parameter 'month' is invalid!\n" +
				"Febrary only has 29-days.\n" +
				"isLeapYear='%v' month='%v' day='%v'\n",
				isLeapYear, month, day)

			return err

		} else if !isLeapYear &&
			day > 28 {

			err = fmt.Errorf(ePrefix + "\n" +
				"Error: Input parameter 'month' is invalid!\n" +
				"Febrary only has 28-days.\n" +
				"isLeapYear='%v' month='%v' day='%v'\n",
				isLeapYear, month, day)

			return err
		}

	}

	if hour > 24 || hour < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hour' is invalid!\n" +
			"hour='%v'\n", hour)
		return err
	}

	if minute < 0 || minute > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minute' is invalid!\n" +
			"minute='%v'\n", minute)
		return err
	}

	calUtil2 := CalendarUtility{}

	isValidSecond := calUtil2.IsValidSecond(
		month,
		day,
		hour,
		minute,
		second)

	// Watch out for leap seconds
	if !isValidSecond {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'second' is invalid!\n" +
			"second='%v'\n", second)
		return err
	}

	if nanosecond < 0 || nanosecond > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecond' is invalid!\n" +
			"nanosecond='%v'\n", nanosecond)
		return err
	}

	err = nil

	return err
}