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

// gregorianDateToJulianDayNoTime - Converts a Gregorian Date to a
// Julian Day Number and Time.
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
//       Day Number/Time.
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
//                                              //   within nanoseconds
//           julianDayNoNumericalSign         int        // Sign of the Julian Day Number/Time value
//           totalJulianNanoSeconds        *big.Int   // Julian Day Number Time Value expressed in nano seconds.
//                                              //   Always represents a value less than 24-hours
//                                              // Julian Hours
//           hours                   int
//           minutes                 int
//           seconds                 int
//           nanoseconds             int
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
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered this error Type will encapsulate
//       an error message.
//
//
func (calMech *calendarMechanics) gregorianDateToJulianDayNoTime(
	gregorianDateTime time.Time,
	ePrefix string) (
	gregorianDateUtc time.Time,
	julianDayNoDto JulianDayNoDto,
	err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calMech.gregorianDateToJulianDayNoTime() "

	gregorianDateUtc = gregorianDateTime.UTC()
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

	//fmt.Printf("julianDayNo: %v\n",
	//	julianDayNo)

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

/*
func (calMech *calendarMechanics) gregorianDateToJulianDayNoTime(
	gregorianDateTime time.Time,
	digitsAfterDecimal int,
	ePrefix string) (
	gregorianDateUtc time.Time,
	julianDayNoTime *big.Float,
	err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calMech.gregorianDateToJulianDayNoTime() "

	gregorianDateUtc = gregorianDateTime.UTC()
	julianDayNoTime = big.NewFloat(0.0)
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

	// JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075

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

	newDigitsAfterDecimal := uint(digitsAfterDecimal)

	bfGregorianTimeNanoSecs := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(newDigitsAfterDecimal).
		SetInt64(gregorianTimeNanoSecs)

	bfDayNanoSeconds := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(newDigitsAfterDecimal).
		SetInt64(DayNanoSeconds)

	bfTimeFraction := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(newDigitsAfterDecimal).
		Quo(bfGregorianTimeNanoSecs,
			bfDayNanoSeconds)

	bfTimeFractionStr :=
		bfTimeFraction.Text('f', int(newDigitsAfterDecimal))

	fmt.Printf("Internal Time Fraction:                            %v\n",
		bfTimeFraction.Text('f', int(newDigitsAfterDecimal)))

	julianDayNoTimeStr :=
		fmt.Sprintf("%v", julianDayNo) +
			bfTimeFractionStr[1:]

	fmt.Printf("julianDayNoTimeStr                           %v\n",
		julianDayNoTimeStr)

	var err2 error

	bfJulianDayNoTime := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(newDigitsAfterDecimal)

	b := 0

	julianDayNoTime,
	b,
	err2 =
			big.ParseFloat(
					julianDayNoTimeStr,
					10,
					0,
					big.ToNearestAway)

	if err2 != nil {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error returned by " +
			"big.ParseFloat(julianDayNoTimeStr)\n" +
			"julianDayNoTimeStr='%v'\n" +
			"Error='%v'\n",
			julianDayNoTimeStr, err2)

		return gregorianDateUtc, julianDayNoTime, err
	}

	if b != 10 {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Expected base 10 to be returned from " +
			"big.ParseFloat(julianDayNoTimeStr)\n" +
			"julianDayNoTimeStr='%v'\n" +
			"base='%v'\n",
			julianDayNoTimeStr, b)

		return gregorianDateUtc, julianDayNoTime, err
	}

	fmt.Printf("julianDayNoTime                            %v\n",
		bfJulianDayNoTime.Text('f', int(newDigitsAfterDecimal)))

	return gregorianDateUtc, julianDayNoTime, err
}
*/

// richardsJulianDayNoTimeToGregorianCalendar - Converts a Julian Day
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
func (calMech *calendarMechanics) richardsJulianDayNoTimeToGregorianCalendar(
	julianDayNoDto JulianDayNoDto,
	ePrefix string) (
	gregorianDateUtc time.Time,
	err error) {

	calMech.lock.Lock()

	defer calMech.lock.Unlock()

	ePrefix += "calMech.richardsJulianDayNoTimeToGregorianCalendar() "

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
	//nanoseconds,
	//_ := timeMech.ComputeTimeElementsBigInt(julianDayNoDto.totalJulianNanoSeconds)


	//fmt.Printf("julianDayNoDto.totalJulianNanoSeconds: hours=%d minutes=%d seconds=%d nanoseconds=%d\n",
	//	hours, minutes,seconds, nanoseconds)

		//fmt.Println("Added 12-hours!")

		timeDifferential := julianDayNoDto.totalJulianNanoSeconds + NoonNanoSeconds

	gregorianDateUtc = gregorianDateUtc.Add(time.Duration(timeDifferential))

	return gregorianDateUtc, err
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
		0,
		0,
		0,
		0,
		time.UTC)

	fmt.Printf(ePrefix + "\n" +
		"Julian Base Date: %v\n",
		julianDateUtc.Format(FmtDateTimeYrMDayFmtStr))
	//timeDifferential := julianDayNoDto.totalJulianNanoSeconds.Int64() + (HourNanoSeconds * 12)
	// timeDifferential := julianDayNoDto.totalJulianNanoSeconds
//	if Y >= 0 {
////		julianDateUtc = julianDateUtc.Add(time.Duration(julianDayNoDto.netGregorianNanoSeconds))
//		julianDateUtc = julianDateUtc.Add(time.Duration(julianDayNoDto.totalJulianNanoSeconds + NoonNanoSeconds))
//	} else {
//		julianDateUtc = julianDateUtc.Add(time.Duration(julianDayNoDto.totalJulianNanoSeconds + NoonNanoSeconds))
//	}

	julianDateUtc = julianDateUtc.Add(time.Duration(julianDayNoDto.totalJulianNanoSeconds + NoonNanoSeconds))


	return julianDateUtc, err
}