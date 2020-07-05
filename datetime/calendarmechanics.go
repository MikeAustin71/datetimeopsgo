package datetime

import (
	"errors"
	"fmt"
	"math/big"
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
//
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
//
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
// Year numbering system which includes year zero. Under this system, the date
// January 1st of year 1 is immediately preceded by the date December 31st, year
// zero.
//
// Proleptic Gregorian
//
// Refers to an extrapolated Gregorian date. The Gregorian Calendar was implemented
// on October 15, 1582. Gregorian dates prior to this implementation date are said
// to be 'proleptic' Gregorian dates because they are projected dates dates prior
// to the introduction of the Gregorian calendar.
//
type calendarMechanics struct {
	input  string
	output string
	lock   sync.Mutex
}

// gregorianDateToJulianDayNoTime - Converts a Gregorian Date to a
// Julian Day Number and Time.
//
// Remember that Julian Day Number Times are valid for all dates
// after noon on Monday, January 1, 4713 BCE, proleptic Julian calendar
// or November 24, 4714 BCE, in the proleptic Gregorian calendar. Therefore,
// using astronomical years encapsulated in the Golang type time.Time,
// this algorithm is valid for all Golang date/times after Gregorian
// calendar (possibly proleptic) values after November 23, −4713.
//
// Reference Wikipedia:
//   https://en.wikipedia.org/wiki/Julian_day
//
// The algorithm employed by this method is based on the work of E.G. Richards.
//
// Reference:
//   Richards, E. G. (1998). Mapping Time: The Calendar and its History.
//   Oxford University Press. ISBN 978-0192862051
//
// However, the original algorithm has been modified to provide for time
// fractions accurate to subMicrosecondNanoseconds.
//
// Taken collectively, the 'input' parameters years, months, days, hours,
// minutes, seconds and subMicrosecondNanoseconds represents a Gregorian date/time using
// the Universal Coordinated Time (UTC Time Zone). Gregorian dates which
// precede November 24, 4714 BCE or 11/24/-4713 (using Astronomical Year
// Numbering System) are invalid and will generate an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  years              int64
//    - The year number expressed as an int64 value. This 'years' value
//      should be formatted using Astronomical Year Numbering; that is,
//      a year numbering system which includes year zero. Year values which
//      are less than -4713, using the using Astronomical Year Numbering System,
//      are invalid and will generate an error.
//
//  months             int
//    - The month number
//
//  days               int
//    - The day number
//
//  hours              int
//    - The hour number expressed on a 24-hour time scale.
//      Example: 3:00PM is passed as the hour 15
//
//  minutes            int
//    - The minutes number
//
//  seconds            int
//    - The number of seconds
//
//  subMicrosecondNanoseconds        int
//    - The number of subMicrosecondNanoseconds
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
//  julianDayNoDto     JulianDayNoDto
//     - This returned type contains the data elements of a Julian Day
//       Number/Time value. Note that key Julian Day Number and Time values
//       are stored as *big.Int and *big.Float
//
//        type JulianDayNoDto struct {
//           julianDayNo             *big.Int   // Julian Day Number expressed as integer value
//           julianDayNoFraction     *big.Float // The Fractional Time value of the Julian
//                                              //   Day No Time
//           julianDayNoTime         *big.Float // JulianDayNo Plus Time Fraction accurate to
//                                              //   within subMicrosecondNanoseconds
//           julianDayNoNumericalSign         int        // Sign of the Julian Day Number/Time value
//           totalJulianNanoSeconds        *big.Int   // Julian Day Number Time Value expressed in nano seconds.
//                                              //   Always represents a value less than 24-hours
//                                              // Julian Hours
//           hours                   int
//           minutes                 int
//           seconds                 int
//           subMicrosecondNanoseconds             int
//           lock                    *sync.Mutex
//        }
//
//   The integer portion of this number (digits to left of
//       the decimal) represents the Julian day number and is
//       stored in 'JulianDayNoDto.julianDayNo'. The fractional
//       digits to the right of the decimal represent elapsed time
//       since noon on the Julian day number and is stored in
//       'JulianDayNoDto.julianDayNoFraction'. The combined Julian
//       Day Number Time value is stored in 'JulianDayNoDto.julianDayNoTime'.
//       All time values are expressed as Universal Coordinated Time (UTC).
//
//
//  err                 error
//     - If this method is successful the returned error Type
//       is set equal to 'nil'. If errors are encountered this
//       error Type will encapsulate an appropriate error message.
//
//
func (calMech *calendarMechanics) gregorianDateToJulianDayNoTime(
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

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calendarMechanics.gregorianDateToJulianDayNoTime() "

	err = nil

	julianDayNoDto = JulianDayNoDto{}

	// Threshold date. Any date before this is
	// invalid!
	// 11/24/-4713 (Astronomical Year Numbering)
	// 12:00:00.000000000
	//thresholdGregorianDate := time.Date(
	//	-4713,
	//	11,
	//	24,
	//	12,
	//	0,
	//	0,
	//	0,
	//	time.UTC)

//	gregorianDateUtc = gregorianDateTime.UTC()

	if years < int64(-4713) {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "years",
			inputParameterValue: "",
			errMsg:              "Gregorian Dates prior to -4713/11/24 12:00:00" +
				" UTC are invalid\nfor Julian Day Number Calculations.",
			err:                 nil,
		}

		return julianDayNoDto, err
	}

	if years == int64(-4713) &&
		months < 11 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "months",
			inputParameterValue: "",
			errMsg:              "Gregorian Dates prior to -4713/11/24 12:00:00" +
				" UTC are invalid\nfor Julian Day Number Calculations.",
			err:                 nil,
		}

		return julianDayNoDto, err
	}

	if years == int64 (-4713) &&
		months == 11 &&
		days < 24 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "days",
			inputParameterValue: "",
			errMsg:              "Gregorian Dates prior to -4713/11/24 12:00:00" +
				" UTC are invalid\nfor Julian Day Number Calculations.",
			err:                 nil,
		}

		return julianDayNoDto, err
	}

	if years == int64 (-4713) &&
		months == 11 &&
		days == 24 &&
		hours < 12 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "hours",
			inputParameterValue: "",
			errMsg:              "Gregorian Dates prior to -4713/11/24 12:00:00" +
				" UTC are invalid\nfor Julian Day Number Calculations.",
			err:                 nil,
		}
		return julianDayNoDto, err
	}

	julianDayNoDto = JulianDayNoDto{}

	err = nil

	Month := int64(months)
	Day := int64(days)

	// JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075

	julianDayNo :=
		(int64(1461) * (years+ int64(4800) +
			(Month - int64(14))/int64(12)))/int64(4) +
			(int64(367) * (Month - int64(2) -
				int64(12) * ((Month - int64(14))/int64(12))))/int64(12) -
			(int64(3) * ((years+ int64(4900) +
				(Month - int64(14))/int64(12))/int64(100)))/int64(4) +
			Day - int64(32075)

//	fmt.Printf("julianDayNo: %v - calendarMechanics.gregorianDateToJulianDayNoTime\n",
//		julianDayNo)

	gregorianTimeNanoSecs := int64(hours) * HourNanoSeconds
	gregorianTimeNanoSecs += int64(minutes) * MinuteNanoSeconds
	gregorianTimeNanoSecs += int64(seconds) * SecondNanoseconds
	gregorianTimeNanoSecs += int64(nanoseconds)

	if gregorianTimeNanoSecs < NoonNanoSeconds {

		if julianDayNo > 0 {
			julianDayNo -= 1
		}

		gregorianTimeNanoSecs += NoonNanoSeconds

	} else {

		gregorianTimeNanoSecs -= NoonNanoSeconds
	}

	bfGregorianTimeNanoSecs :=
		big.NewFloat(0.0).
				SetMode(big.ToZero).
				SetPrec(0).
				SetInt64(gregorianTimeNanoSecs)

	bfDayNanoSeconds := big.NewFloat(0.0).
		SetMode(big.ToZero).
		SetPrec(0).
		SetInt64(DayNanoSeconds)

	bfTimeFraction := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(0).
		Quo(bfGregorianTimeNanoSecs,
			bfDayNanoSeconds)

	jDNDtoUtil := julianDayNoDtoUtility{}

	err = jDNDtoUtil.setDto(
		&julianDayNoDto,
		julianDayNo,
		bfTimeFraction,
		ePrefix)

	return julianDayNoDto, err
}

/*func (calMech *calendarMechanics) gregorianDateToJulianDayNoTime(
	gregorianDateTime time.Time,
	ePrefix string) (
	gregorianDateUtc time.Time,
	julianDayNoDto JulianDayNoDto,
	err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calendarMechanics.gregorianDateToJulianDayNoTime() "

	gregorianDateUtc = gregorianDateTime.UTC()

	julianDayNoDto = JulianDayNoDto{}

	thresholdGregorianDate := time.Date(
		-4713,
		11,
		24,
		12,
		0,
		0,
		0,
		time.UTC)


	if gregorianDateTime.Before(thresholdGregorianDate) {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "gregorianDateTime",
			inputParameterValue: "",
			errMsg:              "Gregorian Dates prior to -4713/11/24 12:00:00" +
				" UTC are invalid\n for Julian Day Number Calculations.",
			err:                 nil,
		}
		return gregorianDateUtc, julianDayNoDto, err
	}

	julianDayNoDto = JulianDayNoDto{}

	err = nil

	Year := int64(gregorianDateUtc.Year())
	Month := int64(gregorianDateUtc.Month())
	Day := int64(gregorianDateUtc.Day())

	// JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075

	julianDayNo :=
		(int64(1461) * (Year + int64(4800) +
			(Month - int64(14))/int64(12)))/int64(4) +
			(int64(367) * (Month - int64(2) -
				int64(12) * ((Month - int64(14))/int64(12))))/int64(12) -
			(int64(3) * ((Year + int64(4900) +
				(Month - int64(14))/int64(12))/int64(100)))/int64(4) +
			Day - int64(32075)

	fmt.Printf("julianDayNo: %v - calendarMechanics.gregorianDateToJulianDayNoTime\n",
		julianDayNo)

	gregorianTimeNanoSecs := int64(gregorianDateUtc.Hour()) * HourNanoSeconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Minute()) * MinuteNanoSeconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Second()) * SecondNanoseconds
	gregorianTimeNanoSecs += int64(gregorianDateUtc.Nanosecond())

	if gregorianTimeNanoSecs < NoonNanoSeconds {

		if julianDayNo > 0 {
			julianDayNo -= 1
		}

		gregorianTimeNanoSecs += NoonNanoSeconds

	} else {

		gregorianTimeNanoSecs -= NoonNanoSeconds
	}

	bfGregorianTimeNanoSecs :=
		big.NewFloat(0.0).
				SetMode(big.ToZero).
				SetPrec(0).
				SetInt64(gregorianTimeNanoSecs)

	bfDayNanoSeconds := big.NewFloat(0.0).
		SetMode(big.ToZero).
		SetPrec(0).
		SetInt64(DayNanoSeconds)

	bfTimeFraction := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(0).
		Quo(bfGregorianTimeNanoSecs,
			bfDayNanoSeconds)

	jDNDtoUtil := julianDayNoDtoUtility{}

	err = jDNDtoUtil.setDto(
		&julianDayNoDto,
		julianDayNo,
		bfTimeFraction,
		ePrefix)

	return gregorianDateUtc, julianDayNoDto, err
}
*/

// julianDayNoTimeToGregorianCalendar - Converts a Julian Day
// Number and Time value to the corresponding date time in the
// Gregorian Calendar.
//
// The Gregorian Calendar is today applied almost universally across
// planet Earth. It is named after Pope Gregory XIII, who introduced
// it in October 1582.  Because the Gregorian Calendar was instituted
// on Friday, October 15, 1582, all Gregorian Calendar dates prior
// to this are extrapolated or proleptic.
//
// This method uses the 'Richards' algorithm.
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
// The Julian day number is based on the Julian Period proposed
// by Joseph Scaliger, a classical scholar, in 1583 (one year after
// the Gregorian calendar reform) as it is the product of three
// calendar cycles used with the Julian calendar.
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
// For more information on the Julian Day Number/Time see:
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
//  gregorianDateUtc    time.Time
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
//     - If this method is successful the returned error Type
//       is set equal to 'nil'. If errors are encountered this
//       error Type will encapsulate an appropriate error message.
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
func (calMech *calendarMechanics) julianDayNoTimeToGregorianCalendar(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	gregorianDateUtc time.Time,
	err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calendarMechanics.julianDayNoTimeToGregorianCalendar() "

	gregorianDateUtc = time.Time{}
	err = nil

	var err2 error

	var bigJulianDayNo *big.Int

	bigJulianDayNo, err2 = julianDayNoDto.GetJulianDayBigInt()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error returned by julianDayNoDto.GetJulianDayBigInt()\n" +
			"Error='%v'\n", err2.Error())
		return gregorianDateUtc, err
	}

	//numericalSignVal := 1
	//
	//if big.NewInt(0).Cmp(bigJulianDayNo) == 1 {
	//	numericalSignVal = -1
	//}

	if big.NewInt(0).Cmp(bigJulianDayNo) == 1 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNoDto.julianDayNo",
			inputParameterValue: "",
			errMsg:              "'julianDayNoDto.julianDayNo' " +
				"is less than Zero!",
			err:                 nil,
		}
		return gregorianDateUtc, err
	}

	var julianDayNoInt64 int64

	if !bigJulianDayNo.IsInt64() {
		err = errors.New(ePrefix + "\n" +
			"Error: Julian Day Number is too large to be represented\n" +
			"by type int64\n")
		return gregorianDateUtc, err
	}

	julianDayNoInt64 = bigJulianDayNo.Int64()

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

	// Julian Day No as int64
	J := julianDayNoInt64

	f := J + j + ((((4 * J + B) / 146097) * 3) /4) + C

	e := r * f + v // #2

	g := (e % p) / r // #3

	h := u * g + w // #4

	D := ((h % s) / u) + 1  // #5

	M := ((( h / s) + m) % n) + 1  // #6

	Y := (e / p) - y + ((n + m - M)/ n)

	gregorianDateUtc = time.Date(
		int(Y),
		time.Month(M),
		int(D),
		0,
		0,
		0,
		0,
		time.UTC)

	//timeMech := TimeMechanics{}
	//
	//_,
	//hours,
	//minutes,
	//seconds,
	//subMicrosecondNanoseconds,
	//_ := timeMech.ComputeTimeElementsBigInt(julianDayNoDto.totalJulianNanoSeconds)


	//fmt.Printf("julianDayNoDto.totalJulianNanoSeconds: hours=%d minutes=%d seconds=%d subMicrosecondNanoseconds=%d\n",
	//	hours, minutes,seconds, subMicrosecondNanoseconds)

		//fmt.Println("Added 12-hours!")

		timeDifferential := julianDayNoDto.totalJulianNanoSeconds + NoonNanoSeconds

	gregorianDateUtc = gregorianDateUtc.Add(time.Duration(timeDifferential))

	return gregorianDateUtc, err
}

// julianCalendarDateJulianDayNo - Creates and returns a 'JulianDayNoDto'
// type populated with the Julian Day Number/Time equivalent to the input
// parameters for a Julian Calendar Date. The 'Dogget' algorithm discussed
// below is valid for all dates greater than or equal to Julian Day Number
// zero.
//
// Julian Day numbers start on day zero at noon. This means that Julian
// Day Number Times are valid for all dates on or after noon on Monday,
// January 1, 4713 BCE, in the proleptic Julian calendar or November 24,
// 4714 BCE, in the proleptic Gregorian calendar. These Date limits are
// expressed using the 'Common Era' epoch or format ('BCE' Before Common
// Era or 'CE' Common Era).
//
// The algorithm used to convert Julian Calendar Dates to a Julian Day Number
// time was taken from L. E. Doggett, Ch. 12, "Calendars", p. 606, in
// Seidelmann 1992.
//
// Reference:
//       https://en.wikipedia.org/wiki/Julian_day
//
//
func (calMech *calendarMechanics) julianCalendarDateJulianDayNo(
	year int64,
	months,
	days,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	ePrefix string) ( julianDayNo JulianDayNoDto, err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calendarMechanics.julianCalendarDateJulianDayNo() "

	err = nil
	julianDayNo = JulianDayNoDto{}

	var jDN, month, day int64

	month = int64(months)
	day = int64(days)
	// JDN = 367 × Y − (7 × (Y + 5001 + (M − 9)/7))/4 + (275 × M)/9 + D + 1729777

	jDN = 367 * year -
		( int64(7) * (year + int64(5001) +
			(month - int64(9))/int64(7)))/int64(4) +
		(int64(275) * month)/int64(9) + day + int64(1729777)

	if hours < 12 {
		jDN--
	}

	totalTime := big.NewInt(int64(hours) * HourNanoSeconds)

	totalTime = big.NewInt(0).Add(totalTime,
		big.NewInt(int64(minutes) * MinuteNanoSeconds))

	totalTime = big.NewInt(0).Add(totalTime,
		big.NewInt(int64(seconds) * SecondNanoseconds))

	totalTime = big.NewInt(0).Add(totalTime,
		big.NewInt(int64(nanoseconds)))

	requestedPrecision :=	uint(1024)

	var totalTimeFrac *big.Float

	totalTimeFrac =
		big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		SetInt(totalTime)

var err2 error

	julianDayNo, err2 = JulianDayNoDto{}.New(
		jDN,
		totalTimeFrac)

	if err2 != nil {
		err = fmt.Errorf(ePrefix +
			"%v", err2.Error())
	}

	return julianDayNo, err
}

// richardsJulianDayNoTimeToJulianCalendar - Converts a Julian Day Number and
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
// zero. Dates are expressed in the 'Common Era' format ('BCE' Before
// Common Era or 'CE' Common Era). Therefore a 'time.Time' year of '-4713'
// is equal to the year '4714 BCE'
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
//     - If this method is successful the returned error Type
//       is set equal to 'nil'. If errors are encountered this
//       error Type will encapsulate an appropriate error message.
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
func (calMech *calendarMechanics) richardsJulianDayNoTimeToJulianCalendar(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	julianDateUtc time.Time,
	err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calendarMechanics.richardsJulianDayNoTimeToJulianCalendar() "

	julianDateUtc = time.Time{}
	err = nil

	var err2 error

	var bigJulianDayNo *big.Int

	bigJulianDayNo, err2 = julianDayNoDto.GetJulianDayBigInt()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error returned by julianDayNoDto.GetJulianDayBigInt()\n" +
			"Error='%v'\n", err2.Error())
		return julianDateUtc, err
	}

	if big.NewInt(0).Cmp(bigJulianDayNo) == 1 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNoDto.julianDayNo",
			inputParameterValue: "",
			errMsg:              "'julianDayNoDto.julianDayNo' " +
				"is less than Zero!",
			err:                 nil,
		}
		return julianDateUtc, err
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
// B := int64(274277)
// C := int64(-38)

	julianDayNumInt := bigJulianDayNo.Int64()

	// Julian Day No as integer
	J := julianDayNumInt

	// Julian Day No + 1401
	f := J + j


	e := r * f + v // #2

	g := (e % p) / r // #3

	h := u * g + w // #4

	D := ((h % s) / u) + 1  // #5

	M := ((( h / s) + m) % n) + 1  // #6

	Y := (e / p) - y + ((n + m - M)/ n)


	julianDateUtc = time.Date(
		int(Y),
		time.Month(M),
		int(D),
		12,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf("Julian Base Date: %v -calendarMechanics.richardsJulianDayNoTimeToJulianCalendar\n",
		julianDateUtc.Format(FmtDateTimeYrMDayFmtStr))

	if julianDayNoDto.netGregorianNanoSeconds == NoonNanoSeconds {

		return julianDateUtc, err

	} else if julianDayNoDto.netGregorianNanoSeconds < NoonNanoSeconds {

		if Y < 0 {

			julianDateUtc = julianDateUtc.Add(time.Duration(
				julianDayNoDto.totalJulianNanoSeconds))

		} else {

			julianDateUtc = julianDateUtc.Add(time.Duration(
				(DayNanoSeconds * -1) + julianDayNoDto.totalJulianNanoSeconds))
		}

	} else {

		julianDateUtc = julianDateUtc.Add(time.Duration(
			julianDayNoDto.totalJulianNanoSeconds))

	}

	return julianDateUtc, err
}

// revisedGoucherParkerToJulianDayNo - Uses an algorithm by Mike Rapp
// to compute the julian day number for a date on the Revised Goucher-Parker
// calendar.
//
// The Revised Goucher-Parker Calendar implements the Julian Calendar date
// system. Therefore, Julian Day numbers start on day zero at noon. This
// means that Julian Day Number Zero begins at noon on Monday, January 1, 4713 BCE,
// in the proleptic Julian calendar. Using astronomical year numbering this translates
// to Monday, January 1, -4712.
//
// Reference:
//  See documentation for Type, 'CalendarSpec' at datetime\calendarspecenum.go
//
//
// Summary
//
// 1. Apply the Julian Calendar.
// 2. If year is divisible by 4, it IS a leap year; add on day to February
// 3. If year is divisible by 128 it is NOT a leap year and no day is added to February.
// 4. If year is divisible by 454,545, it IS a leap year; add a day to February.
//
func (calMech *calendarMechanics) revisedGoucherParkerToJulianDayNo(
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

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calendarMechanics.revisedGoucherParkerToJulianDayNo() "

	err = nil

	julianDayNoDto = JulianDayNoDto{}

	calDtMech := calendarDateTimeMechanics{}

	numSign := int64(1)

	// Generate absolute years value
	if years < 0 {
		numSign = -1
		years *= -1
	}

	// Julian Day numbers start on day zero at noon. For the Julian Calendar,
	// this means that Julian Day Number Zero begins at noon on Monday,
	// January 1, 4713 BCE, in the proleptic Julian calendar. Using
	// astronomical year numbering this is Monday, January 1, -4712
	//
	// Leap Year Formula
	//
	// 1. Apply the Julian Calendar.
	// 2. If year is divisible by 4, it IS a leap year; add on day to February
	// 3. If year is divisible by 128 it is NOT a leap year and no day is added to February.
	// 4. If year is divisible by 454,545, it IS a leap year; add a day to February.

	// baseYear is year prior to years - 4712

	baseYear := years - 4712 - 1

	baseDays := baseYear * 365

	plus4YrLeapYrs := baseYear / 4

	less128YrNonLeapYrs := baseYear / 128

	plusCycleLeapYrs := baseYear / 454545

	LeapYrDays := plus4YrLeapYrs -
		less128YrNonLeapYrs +
		plusCycleLeapYrs

	isCurrentLeapYear := calDtMech.isRevisedGoucherParkerLeapYear(years)

	var ordinalDayNum int

	ordinalDayNum, err = calDtMech.ordinalDayNumber(
		months,
		days,
		isCurrentLeapYear,
		ePrefix)

	if err != nil {
		return julianDayNoDto, err
	}

	julianDayNo :=
		baseDays + LeapYrDays + int64(ordinalDayNum)

	//	fmt.Printf("julianDayNo: %v - calendarMechanics.gregorianDateToJulianDayNoTime\n",
	//		julianDayNo)

	gregorianTimeNanoSecs := int64(hours) * HourNanoSeconds
	gregorianTimeNanoSecs += int64(minutes) * MinuteNanoSeconds
	gregorianTimeNanoSecs += int64(seconds) * SecondNanoseconds
	gregorianTimeNanoSecs += int64(nanoseconds)

	if gregorianTimeNanoSecs < NoonNanoSeconds {

		if julianDayNo > 0 {
			julianDayNo -= 1
		}

		gregorianTimeNanoSecs += NoonNanoSeconds

	} else {

		gregorianTimeNanoSecs -= NoonNanoSeconds
	}

	bfGregorianTimeNanoSecs :=
		big.NewFloat(0.0).
			SetMode(big.ToZero).
			SetPrec(0).
			SetInt64(gregorianTimeNanoSecs)

	bfDayNanoSeconds := big.NewFloat(0.0).
		SetMode(big.ToZero).
		SetPrec(0).
		SetInt64(DayNanoSeconds)

	bfTimeFraction := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(0).
		Quo(bfGregorianTimeNanoSecs,
			bfDayNanoSeconds)

	jDNDtoUtil := julianDayNoDtoUtility{}

	julianDayNo = julianDayNo * numSign

	err = jDNDtoUtil.setDto(
		&julianDayNoDto,
		julianDayNo,
		bfTimeFraction,
		ePrefix)

	return julianDayNoDto, err
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
func (calMech *calendarMechanics) revisedJulianCalendarToJulianDayNo(
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

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calendarMechanics.revisedJulianCalendarToJulianDayNo() "

	err = nil

	julianDayNoDto = JulianDayNoDto{}

	calDtMech := calendarDateTimeMechanics{}

	numSign := int64(1)

	// Generate absolute years value
	if years < 0 {
		numSign = -1
		years *= -1
	}

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

	baseYear := years - 4712 - 1

	baseDays := baseYear * 365

	plus4YrLeapYrs := baseYear / 4

	less100YrLeapYrs := baseYear / 100

	plus100YrLeapYrsBy900

	return julianDayNoDto, err
}

// Week Number
// https://www.timeanddate.com/date/week-numbers.html

// usDayOfWeekNumber - Receives a Julian Day Number and returns
// the equivalent U.S. Day of the Week number. The U.S. Day
// of the Week Number begins numbering week days with 'Sunday',
// which is assigned day number '0' (zero).
//
// The algorithm used to calculate the U.S. Day of the Week number
// is taken from:
//   https://en.wikipedia.org/wiki/Julian_day
//   Richards 2013, pp. 592, 618
//   Richards, E. G. (2013). Calendars. In S. E. Urban & P. K.
//   Seidelmann, eds. Explanatory Supplement to the Astronomical Almanac,
//   3rd ed. (pp. 585–624). Mill Valley, Calif.: University Science Books.
//   ISBN 978-1-89138-985-6
//
// Reference:
//   https://en.wikipedia.org/wiki/Determination_of_the_day_of_the_week

//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  julianDayNumber    JulianDayNoDto
//    - The Julian Day Number (and time) which will be converted
//      to a day of the week number.
//
//        type JulianDayNoDto struct {
//          julianDayNo         *big.Int   // Julian Day Number expressed as integer value
//          julianDayNoFraction *big.Float // The Fractional Time value of the Julian
//                                             Day No Time
//          julianDayNoTime *big.Float     // Julian Day Number Plus Time Fraction accurate to
//                                             within subMicrosecondNanoseconds
//          julianDayNoNumericalSign int   // Sign of the Julian Day Number/Time value
//          totalJulianNanoSeconds   int64 // Julian Day Number Time Value expressed in subMicrosecondNanoseconds.
//                                             Always represents a positive value less than 36-hours
//          netGregorianNanoSeconds int64  // Gregorian subMicrosecondNanoseconds. Always represents a value in
//                                             subMicrosecondNanoseconds which is less than 24-hours.
//          hours       int                // Gregorian Hours
//          minutes     int
//          seconds     int
//          subMicrosecondNanoseconds int
//          lock        *sync.Mutex
//        }
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
// usDayOfWeekNo       UsDayOfWeekNo
//    - If the method completes successfully this enumeration
//      type is returned specifying the U. S. day of the week
//      number associated with the input parameter, 'julianDayNumber'.
//
//
//  err                 error
//     - If this method is successful the returned error Type
//       is set equal to 'nil'. If errors are encountered this
//       error Type will encapsulate an appropriate error message.
//
//
func (calMech *calendarMechanics) usDayOfWeekNumber(
	julianDayNumber JulianDayNoDto,
	ePrefix string) (usDayOfWeekNo UsDayOfWeekNo, err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calendarMechanics.usDayOfWeekNumber() "

	err = nil

	usDayOfWeekNo = UsDayOfWeekNo(0).None()

	jDN, err2 := julianDayNumber.GetJulianDayBigInt()

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n" +
			"Input parameter julianDayNumber is INVALID!\n" +
			"Error='%v'\n", err2.Error())

		return usDayOfWeekNo, err
	}

	// Richards Algorithm
	//  W1 = mod(J + 1, 7)

	bigOne := big.NewInt(1)

	bigSeven := big.NewInt(7)

	jDN = big.NewInt(0).Add(jDN, bigOne)

	w1 := big.NewInt(0).Mod(jDN, bigSeven)

	if ! w1.IsInt64() {
		err = errors.New(ePrefix + "\n" +
			"Error: Algorithm returned invalid result!\n" +
			"'w1' cannot be represented by an int64.")

		return usDayOfWeekNo, err
	}

	w1Int64 := w1.Int64()

	if w1Int64 > 6 ||
		w1Int64 < 0 {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: w1Int64 result is out of bounds.\n" +
			"Day of week number is invalid!\n" +
			"Day Of Week='%v'\n", w1Int64)

		return usDayOfWeekNo, err
	}

	usDayOfWeekNo = UsDayOfWeekNo(int(w1Int64))

	return usDayOfWeekNo, err
}
