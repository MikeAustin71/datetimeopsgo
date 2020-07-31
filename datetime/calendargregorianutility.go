package datetime

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"
)

// CalendarGregorianUtility - This type contains methods
// used to process date arithmetic associated with the
// Gregorian Calendar
//
// Reference:
//  https://en.wikipedia.org/wiki/Gregorian_calendar
//
type CalendarGregorianUtility struct {

	lock *sync.Mutex
}


// GetJulianDayNumber - Returns values defining the Julian Day Number
// and Time for a date/time in the Revised Julian Calendar.
//
// All time input parameters are assumed to be expressed in Coordinated
// Universal Time (UTC). For more information on Coordinated Universal
// Time, reference:
//   https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// Julian Day Number is used to define a standard time duration, and
// perform date/time conversions, between differing calendar systems.
//
// The base date/time for Julian Day Number zero on the Gregorian
// Calendar is November 24, -4713 12:00:00.000000000 UTC (Noon) or
// November 24, 4714 BCE 12:00:00.000000000 UTC (Noon).
//
// For more information on the Julian Day Number, reference:
//  https://en.wikipedia.org/wiki/Julian_day
//
// For more information on the Gregorian Calendar, reference:
//  https://en.wikipedia.org/wiki/Gregorian_calendar
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetYear         int64
//     - The year number associated with this date/time specification.
//       The year value may be positive or negative. The year value must
//       conform to the astronomical year numbering system. This means
//       that year zero is valid and recognized. Example: 1/1/0000. The
//       astronomical year value -4712 is therefore equivalent to
//       -4713 BCE. All year values submitted to this method must use
//       the astronomical year numbering system. For more information
//       on the astronomical year numbering system, reference:
//              https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//  targetMonth        int
//     - The month number for this date/time specification.
//       The valid range is 1 - 12 inclusive.
//
//  targetDay          int
//     - The day number for this date/time specification. The day
//       number must fall within the limits of the month number
//       submitted above in input parameter, 'targetMonth'.
//
//  targetHour         int
//     - The hour time component for this date/time specification.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00. All time
//       parameters are assumed to be expressed in Universal
//       Coordinated Time (UTC).
//
//  targetMinute       int
//     - The minute time component for this date/time specification.
//       The valid range is 0 - 59 inclusive.  All time
//       parameters are assumed to be expressed in Universal
//       Coordinated Time (UTC).
//
//  targetSecond       int
//     - The second time component for this date/time specification.
//       The valid range is 0 - 60 inclusive. The value 60 is only
//       used in the case of leap seconds.  All time parameters are
//       assumed to be expressed in Universal Coordinated Time (UTC).
//
//  targetNanosecond   int
//     - The nanosecond time component for this date/time specification.
//       The valid range is 0 - 999,999,999 inclusive.  All time
//       parameters are assumed to be expressed in Universal
//       Coordinated Time (UTC).
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  julianDayNumber             int64
//     - The integer julian day number of the date specified by input
//       parameters 'targetYear', 'targetMonth' and 'targetDay'. This
//       value equals the number of days elapsed between the base date,
//       February 8, 4714 BCE 12:00:00 UTC (Noon), and the target date/time
//       specified by the input parameters. Both base and target date/times
//       represent moments on the Revised Julian Calendar.
//
//
//  julianDayNumberTime         *big.Float
//     - The combined Julian Day number and fractional time value represented
//       by the date/time specified by input parameters, targetYear,
//       targetMonth, targetDay, targetHour, targetMinute, targetSecond
//       and targetNanosecond. This value equals the duration in day numbers
//       calculated from the start date of February 8, 4714 BCE 12:00:00 UTC
//       (Noon) on the Revised Julian Calendar.
//
//
//  julianDayNumberTimeFraction *big.Float
//     - The fractional time value associated with the the ordinal
//       day number. This value does NOT contain the integer ordinal
//       day number. Instead, it only contains the time value
//       represented by the date/time input parameters, targetHour,
//       targetMinute, targetSecond and targetNanosecond. This value
//       equals the number of nanoseconds since midnight of the ordinal
//       day number divided by the number of nanoseconds in a
//       24-hour day.
//
//
//  err                         error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calGreg *CalendarGregorianUtility) GetJDN(
	targetYear int64,
	targetMonth int,
	targetDay int,
	targetHour int,
	targetMinute int,
	targetSecond int,
	targetNanosecond int,
	ePrefix string) (
	julianDayNumber int64,
	julianDayNumberTime *big.Float,
	julianDayNumberTimeFraction *big.Float,
	err error) {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	ePrefix += "CalendarGregorianUtility.GetJDN() "

	julianDayNumber = 0

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	julianDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	calGregMech := calendarGregorianMechanics{}

	isLeapYear := calGregMech.isLeapYear(targetYear)

	calUtil := CalendarUtility{}

	err = calUtil.IsValidDateTimeComponents(
		isLeapYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {

		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	var baseDateTimeDto, targetDateTimeDto DateTransferDto

	targetDateTimeDto, err = DateTransferDto{}.New(
		isLeapYear,
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	calCycleCfg := calGregMech.getCalendarCyclesConfig()

	baseDateTimeDto = calCycleCfg.GetJDNBaseStartYearDateTime()

	// Base Date Time for Gregorian Calendar:
	// November 24, -4713 12:00:00.000000000 UTC (Noon)

	var baseTargetComparisonResult, baseRemainingDaysInYear,
	targetRemainingDaysInYear, targetYearOrdinalNumber,
	julianDayNumberSign int

	var mainCycleAdjustmentDays, remainderYears, ordinalDays *big.Int

	baseTargetComparisonResult, err = baseDateTimeDto.Compare(&targetDateTimeDto, ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	baseDateTimeDto.SetTag("baseDateTimeDto")
	targetDateTimeDto.SetTag("targetDateTimeDto")

	targetYearOrdinalNumber, err = targetDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	targetRemainingDaysInYear, err = targetDateTimeDto.GetRemainingDaysInYear(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	fmt.Printf("Target Days Remaining in Year: %v\n",
		targetRemainingDaysInYear)

	baseRemainingDaysInYear, err = baseDateTimeDto.GetRemainingDaysInYear(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	fmt.Printf("Base Days Remaining in Year: %v\n",
		baseRemainingDaysInYear)

	// Primary Decision Tree
	if baseTargetComparisonResult == 0 {
		// base and target have equivalent date/times
		// Julian Day Number is Zero and time fraction is Zero.

		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err

	} else if baseTargetComparisonResult == 1 {
		// base is greater than target date/time
		// target date/time must be negative
		// Julian Day Number must be negative

		mainCycleStartDateTime :=
			calCycleCfg.GetMainCycleStartDateForNegativeJDNNo()

		mainCycleStartDateTime.SetTag("Main Cycle Start Date Time")

		fmt.Printf("%s\n",
			mainCycleStartDateTime.String())

		julianDayNumberSign = -1

		// RemainderYears =
		// BaseYear - TargetYear -1
		remainderYears =
			big.NewInt(
				mainCycleStartDateTime.GetYear() -
					targetDateTimeDto.GetYear() - 1 )

		// OrdinalDays =
		// BaseDaysRemainingInYear +
		//    TargetOrdinalDayNumber

		ordinalDays =
			big.NewInt(
				int64(
					baseRemainingDaysInYear +
					targetYearOrdinalNumber))

		mainCycleAdjustmentDays =
			calCycleCfg.GetMainCycleAdjustmentDaysForNegativeJDNNo()

	} else {
		// target is greater than base date/time
		// base is less than target date/time
		// target date/time could be positive or negative
		// Use Main Cycle Start Date/Time which is less Than target

		mainCycleStartDateTime :=
			calCycleCfg.GetMainCycleStartDateForPositiveJDNNo()

		mainCycleStartDateTime.SetTag("Main Cycle Start Date Time")

		fmt.Printf("%s\n",
			mainCycleStartDateTime.String())

		julianDayNumberSign = 1

		// RemainderYears or Whole Years Interval =
		// TargetYear - mainCycleStartDateTime -1
		remainderYears =
			big.NewInt(
			targetDateTimeDto.GetYear() -
				mainCycleStartDateTime.GetYear() - 1)

		fmt.Printf("Initial Remainder Years: %v\n",
			remainderYears.Text(10))

		ordinalDays =
			big.NewInt(
				int64(
					baseRemainingDaysInYear +
						targetYearOrdinalNumber) )

	mainCycleAdjustmentDays =
		calCycleCfg.GetMainCycleAdjustmentDaysForPositiveJDNNo()

}

	cycles := calCycleCfg.GetCalendarCycleConfigurations()
	daysDuration := big.NewInt(0)

	if remainderYears.Cmp(big.NewInt(0)) == 1 {

		var cycleCount *big.Int
		totalYears := big.NewInt(0)

		for i:=0; i < len(cycles); i++ {

			cycleCount =
				big.NewInt(0).
					Quo(remainderYears,
						cycles[i].GetYearsInCycle())

			remainderYears =
				big.NewInt(0).
					Sub(remainderYears,
						big.NewInt(0).
						Mul(cycleCount, cycles[i].GetYearsInCycle()))

			fmt.Printf("cycles[%v].cycleCount=%v\n",
				i, cycleCount.Text(10))

			fmt.Printf("cycles[%v].yearsInCycle=%v\n",
				i, cycles[i].GetYearsInCycle().Text(10))

			fmt.Printf("cycles[%v].remainderYears=%v\n",
				i, remainderYears.Text(10))


			cycles[i].SetCycleCount(cycleCount)
			cycles[i].SetRemainderYears(remainderYears)

			totalYears =
				big.NewInt(0).
					Add(totalYears, cycles[i].GetCycleCountTotalYears())

			fmt.Printf("cycles[%v].CycleCountTotalDays=%v\n",
				i, cycles[i].GetCycleCountTotalDays().Text(10))

			daysDuration =
				big.NewInt(0).
					Add(daysDuration, cycles[i].GetCycleCountTotalDays())
		}

		fmt.Printf("daysDuration: %v\n",
			daysDuration.Text(10))

		daysDuration =
			big.NewInt(0).
				Add(daysDuration, mainCycleAdjustmentDays)

		fmt.Printf("Adjusted daysDuration: %v\n",
			daysDuration.Text(10))

	}

	fmt.Printf("ordinalDays: %v\n",
		ordinalDays.Text(10))

	daysDuration =
		big.NewInt(0).
			Add(daysDuration, ordinalDays)

	fmt.Printf("Final daysDuration: %v\n",
		daysDuration.Text(10))

	if !daysDuration.IsInt64() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: 'daysDuration' is too large and cannot convert to int64!\n" +
			"daysDuration='%v'\n",
			daysDuration.Text(10))
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	julianDayNumber = daysDuration.Int64()

	// Compute Time Components

	noonNanoseconds := int64(time.Hour * 12)

	var targetDayTotalNanoseconds int64

	targetDayTotalNanoseconds, err =
		targetDateTimeDto.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}


	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)

	julianDayNoTimeAdjustment := int64(0)

	if targetDayTotalNanoseconds == noonNanoseconds {

		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(0.0)

	} else if targetDayTotalNanoseconds < noonNanoseconds {
		targetDayTotalNanoseconds += noonNanoseconds

		julianDayNoTimeAdjustment--

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)


	} else {
		// targetDayTotalNanoseconds > noonNanoseconds


		targetDayTotalNanoseconds -= noonNanoseconds

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)
	}

	julianDayNumber += julianDayNoTimeAdjustment

	fmt.Printf("julianDayNoTimeAdjustment: %v\n",
		julianDayNoTimeAdjustment)

	julianDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(julianDayNumber)

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Add(julianDayNoFloat, julianDayNumberTimeFraction)


	if julianDayNumberSign == -1 {

		julianDayNumber *= -1

		julianDayNumberTime =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Neg(julianDayNumberTime)
	}

	separator := strings.Repeat("-", 75)
	fmt.Println()
	fmt.Println(separator)

	fmt.Printf("Final Julian Day Number: %v\n",
		julianDayNumber)

	fmt.Println(separator)
	fmt.Println(baseDateTimeDto.String())
	fmt.Println(targetDateTimeDto.String())
	fmt.Println(separator)

	return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
}

// GetOrdinalDayNumberTime - Computes and returns the 'Fixed Day'
// or ordinal day number time under the Revised Julian Calendar.
//
// The ordinal day number time identifies the number of days and
// and fraction times since a fixed start date. The start date
// for the Revised Julian Calendar is Monday, 1 January 1 AD 00:00:00
// (midnight). This is defined as the beginning of the first ordinal
// day. All ordinal day numbers use this start date as the base base
// reference point.
//
// This method receives a series of input parameters specifying
// a target date and time. The times are always assumed to be
// Universal Coordinated Times (UTC). The method then returns
// the ordinal day number as integer (int64) and a type *big.Float
// which defines both the ordinal day number and the time expressed
// as a fraction of a 24-hour day.
//
// Note that the input parameter 'targetYear' is a type int64 which
// must be configured under the astronomical year numbering system.
// This system recognizes the year zero as a legitimate year value.
// Under the astronomical year numbering system the year 4713 BCE is
// formatted as -4712.
//
// Reference:
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//  https://en.wikipedia.org/wiki/Rata_Die
//  https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetYear         int64
//     - The year number associated with this date/time specification.
//       The year value may be positive or negative. The year value must
//       conform to the astronomical year numbering system. This means
//       that year zero is valid and recognized. Example: 1/1/0000. The
//       astronomical year value -4712 is therefore equivalent to
//       -4713 BCE. All year values submitted to this method must use
//       the astronomical year numbering system. For more information
//       on the astronomical year numbering system, reference:
//              https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//  targetMonth        int
//     - The month number for this date/time specification.
//       The valid range is 1 - 12 inclusive.
//
//  targetDay          int
//     - The day number for this date/time specification. The day
//       number must fall within the limits of the month number
//       submitted above in input parameter, 'targetMonth'.
//
//  targetHour         int
//     - The hour time component for this date/time specification.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00. 'targetHour'
//       is assumed to represent Coordinated Universal Time (UTC).
//
//  targetMinute       int
//     - The minute time component for this date/time specification.
//       The valid range is 0 - 59 inclusive. 'targetMinute' is assumed
//       to represent  Coordinated Universal Time (UTC).
//
//  targetSecond       int
//     - The second time component for this date/time specification.
//       The valid range is 0 - 60 inclusive. The value 60 is only
//       used in the case of leap seconds.
//
//  targetNanosecond   int
//     - The nanosecond time component for this date/time specification.
//       The valid range is 0 - 999,999,999 inclusive
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ordinalDayNumber              int64
//     - The ordinal day number of the date represented by the date
//       specified by input parameters targetYear, targetMonth and
//       targetDay. This value equals the number of days elapsed,
//       plus 1-day, since January 1, 1 CE on the Revised Julian
//       Calendar.
//
//
//  ordinalDayNumberTime          *big.Float
//     - The combined ordinal day number and fractional time
//       value for the ordinal day number represented by the
//       date/time specified by input parameters, targetYear,
//       targetMonth, targetDay, targetHour, targetMinute,
//       targetSecond and targetNanosecond. This value equals
//       the ordinal day number calculated from the start date
//       of January 1, 1 CE on the Revised Julian Calendar.
//
//
//  ordinalDayNumberTimeFraction  *big.Float
//     - The fractional time value associated with the the ordinal
//       day number. This value does NOT contain the integer ordinal
//       day number. Instead, it only contains the time value
//       represented by the date/time input parameters, targetHour,
//       targetMinute, targetSecond and targetNanosecond. This value
//       equals the number of nanoseconds since midnight of the ordinal
//       day number divided by the number of nanoseconds in a
//       24-hour day.
//
//
//  err                           error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calGreg *CalendarGregorianUtility) GetOrdinalDayNumberTime(
	targetYear int64,
	targetMonth,
	targetDay,
	targetHour,
	targetMinute,
	targetSecond,
	targetNanosecond int,
	ePrefix string) (
	ordinalDayNumber int64,
	ordinalDayNumberTime *big.Float,
	ordinalDayNumberTimeFraction *big.Float,
	err error) {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	ordinalDayNumber = 0

	ordinalDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	ordinalDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	calGreg2 := CalendarGregorianUtility{}

	isTargetLeapYear := calGreg2.IsLeapYear(targetYear)

	calUtil := CalendarUtility{}

	err = calUtil.IsValidDateTimeComponents(
		isTargetLeapYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {

		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	var baseDateTimeDto, targetDateTimeDto DateTransferDto

	targetDateTimeDto, err = DateTransferDto{}.New(
		isTargetLeapYear,
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	// Base Date Time:  January 1, 0001 00:00:00.000000000 UTC
	baseDateTimeDto, err = DateTransferDto{}.New(
		false,
		int64(1),
		1,
		1,
		0,
		0,
		0,
		0,
		ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	var baseTargetComparisonResult, baseTargetYearsComparison int

	baseTargetComparisonResult, err = baseDateTimeDto.Compare(&targetDateTimeDto, ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	baseTargetYearsComparison = baseDateTimeDto.CompareYears(&targetDateTimeDto)

	// baseDateTimeDto is now less than targetDateTimeDto

	baseDateTimeDto.SetTag("baseDateTimeDto")
	targetDateTimeDto.SetTag("targetDateTimeDto")
	// Compute wholeYearInterval
	var wholeYearsInterval int64

	var targetYearOrdinalNumber, baseYearOrdinalNumber int

	targetYearOrdinalNumber, err = targetDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	wholeYearsInterval = 0

	if baseTargetComparisonResult == 0 {
		// base and target have equivalent date/times
		// Julian Day Number is Zero and time fraction is Zero.

		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err

	} else if baseTargetYearsComparison == 0 {
		// base and target years are equal
		ordinalDayNumber = int64(targetYearOrdinalNumber)




	} else if baseTargetComparisonResult == 1 {
		// base is greater than target date/time
		// target date/time must be negative
		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() + 1

	} else {
		// target is greater than base date/time
		// target date/time could be positive or
		// negative
		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() - 1

	}

	if wholeYearsInterval < 0 {
		wholeYearsInterval *= -1
	}

	var wholeYearDays int64

	if wholeYearsInterval == 0 {

		wholeYearDays = 0

	} else {

		wholeYearDays =
			calGreg2.NumCalendarDaysForWholeYearsInterval(wholeYearsInterval)

	}

	baseYearOrdinalNumber, err = baseDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	var lastWholeTargetYear int64
	var lastWholeTargetYearDeltaDays, targetPartialYearDeltaDays int64


	if baseTargetYearsComparison == 1 {
		// base year is greater than target year
		// target year is negative

		lastWholeTargetYear = targetDateTimeDto.GetYear() + 1

	} else {
		// baseTargetYearsComparison == -1
		// base year is less than target year.
		// target could be positive or negative.
		//
		if targetDateTimeDto.GetYearNumberSign() < 0 {
			// target year is negative
			lastWholeTargetYear = targetDateTimeDto.GetYear() + 1

		} else {
			// target year is positive
			lastWholeTargetYear = targetDateTimeDto.GetYear() - 1
		}
	}

	if calGreg2.IsLeapYear(lastWholeTargetYear) {
		lastWholeTargetYearDeltaDays = int64(366) - int64(baseYearOrdinalNumber)
	} else {
		lastWholeTargetYearDeltaDays = int64(365) - int64(baseYearOrdinalNumber)
	}

	targetPartialYearDeltaDays = int64(targetYearOrdinalNumber)

	ordinalDayNumber = wholeYearDays +
		lastWholeTargetYearDeltaDays +
		targetPartialYearDeltaDays

	var targetDayTotalNanoseconds int64

	targetDayTotalNanoseconds, err =
		targetDateTimeDto.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)

	targetNanosecondFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(targetDayTotalNanoseconds)

	ordinalDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)

	ordinalDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(ordinalDayNumber)

	ordinalDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Add(ordinalDayNoFloat, ordinalDayNumberTimeFraction)

	separator := strings.Repeat("-", 65)
	fmt.Println()
	fmt.Println(separator)

	fmt.Println(separator)
	fmt.Println(baseDateTimeDto.String())
	fmt.Println(targetDateTimeDto.String())
	fmt.Println(separator)

	return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
}

// GetYearDays - Returns the number of days in the year
// identified by input parameter 'year' under the Gregorian
// Calendar.
//
// If the year is a standard year, this method will return 365-days.
// If the year is a leap year, this method will return 365-days.
//
// For more information on the Gregorian Calendar and leap years,
// reference:
//
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//   https://www.timeanddate.com/date/leapyear.html
//
//
func (calGreg *CalendarGregorianUtility) GetYearDays(
	year int64) int {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	calGregMech := calendarGregorianMechanics{}

	isLeapYear := calGregMech.isLeapYear(year)

	if isLeapYear {
		return 366
	}

	return 365
}

// IsLeapYear - Returns a boolean value signaling whether the year
// value passed as an input parameter is a leap year (366-days)
// under the Gregorian Calendar.
//
// If the method returns 'true' the input parameter 'year' qualifies
// as a leap year consisting of 366-days. If the method returns 'false',
// the input parameter 'year' is a standard year consisting of 365-days.
//
// Methodology:
//
// In the Gregorian calendar, three criteria must be taken
// into account to identify leap years:
//
// 1. The year must be evenly divisible by 4;
//
// 2. If the year can also be evenly divided by 100, it is not
//    a leap year, unless...
//
//  3. The year is evenly divisible by 100 and the year is also
//     evenly divisible by 400. Then it is a leap year.
//
// According to these rules, the years 2000 and 2400 are leap years,
// while 1800, 1900, 2100, 2200, 2300, and 2500 are not leap years.
//
// For more information on the Gregorian Calendar and leap years,
// reference:
//
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//   https://www.timeanddate.com/date/leapyear.html
//
//
func (calGreg *CalendarGregorianUtility) IsLeapYear(
	year int64) bool {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()
	calGregMech := calendarGregorianMechanics{}

	return calGregMech.isLeapYear(year)
}

// JulianDayNumberTimeToDateTime - Receives a Julian Day Number Time floating
// point value and returns the equivalent date/time under the Gregorian Calendar.
//
// The start date/time for Julian Day Number calculates performed under the
// Gregorian Calendar is November 24, -4713 12:00:00.000000000 UTC (Noon). All
// calculated Julian Day Numbers after this moment are positive. All Julian Day
// Numbers prior to this moment are negatvie.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  julianDayNumberTime  *big.Float
//     - The Julian Day Number/Time expressed as a floating point value.
//
//
//
//  ePrefix              string
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
//  gregorianDateTime    DateTransferDto
//     - If successful this method will return a new, fully populated instance
//       of type DateTransferDto contain the year, month, day, hour, minute,
//       second and nanosecond date/time value equivalent to the Julian Day Number
//       Time passed as an input parameter.
//
//  err                  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calGreg *CalendarGregorianUtility) JulianDayNumberTimeToDateTime(
	julianDayNumberTime *big.Float,
	ePrefix string) (gregorianDateTime DateTransferDto, err error) {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	gregorianDateTime = DateTransferDto{}
	err = nil

	julianDayNumTimeNumberSign := julianDayNumberTime.Sign()

	// Convert to absolute value
	if julianDayNumTimeNumberSign == -1 {
		julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Neg(julianDayNumberTime)

	}

	// Truncate to integer
	julianDayNoBigInt, _ :=
		julianDayNumberTime.Int(nil)

	julianDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt(julianDayNoBigInt)

	julianDayNumberTimeFraction :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Sub(julianDayNumberTime, julianDayNoFloat)

	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64( int64(time.Hour * 24))

	convertedNanoseconds :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Mul(julianDayNumberTimeFraction, twentyFourHourNanosecondsFloat)

	remainingNanoseconds, _ :=
		convertedNanoseconds.Int64()

	noonNanoseconds := int64(time.Hour * 12)

	if remainingNanoseconds < noonNanoseconds {
		remainingNanoseconds += noonNanoseconds
	} else {
		// remainingNanoseconds >= noonNanoseconds
		remainingNanoseconds -= noonNanoseconds
	}

	hour := int(remainingNanoseconds / int64(time.Hour))

	remainingNanoseconds -= int64(hour) * int64(time.Hour)

	minute := int(remainingNanoseconds / int64(time.Minute))

	remainingNanoseconds -= int64(minute) * int64(time.Minute)

	second := int(remainingNanoseconds / int64(time.Second))

	remainingNanoseconds -= int64(second) * int64(time.Second)

	nanosecond := int(remainingNanoseconds)


	cycleCount := big.NewInt(0)

	remainderDays :=
		big.NewInt(0).
			Set(julianDayNoBigInt)

	yearsDuration := big.NewInt(0)

	calGregMech := calendarGregorianMechanics{}

	calendarConfig := calGregMech.getCalendarCyclesConfig()

	cycleDtos := calendarConfig.GetCalendarCycleConfigurations()

	blank := big.NewInt(0)

	for i:=0; i < len(cycleDtos); i++ {

		cycleCount,
		remainderDays = big.NewInt(0).
			QuoRem(
				remainderDays,
				cycleDtos[i].GetDaysInCycle(),
				blank)

		cycleDtos[i].SetCycleCount(cycleCount)
		cycleDtos[i].SetRemainderDays(remainderDays)

		yearsDuration = big.NewInt(0).
			Add(yearsDuration, cycleDtos[i].GetCycleCountTotalYears())

	}

	fmt.Printf("yearsDuration Raw Total: %s\n",
		yearsDuration.Text(10))

	var yearNumber *big.Int

	if julianDayNumTimeNumberSign == -1 {
		// Julian Day Number is Less Than Base
		yearNumber =
			big.NewInt(0).
				Add(
					yearsDuration,
					big.NewInt(4713))

	} else {
		// Julian Day Number is greater than or
		// equal to Base Year

		yearNumber =
			big.NewInt(0).
				Sub(
				yearsDuration,
				big.NewInt(4713))

	}

	fmt.Printf("Remainder Days Raw: %s\n",
	remainderDays.Text(10))

	if remainderDays.Cmp(big.NewInt(37)) > -1 {

		yearNumber =
			big.NewInt(0).
				Add(yearNumber,
					big.NewInt(1))

		remainderDays =
			big.NewInt(0).
				Sub(
					remainderDays,
					big.NewInt(37))

	}

	//fmt.Printf("Remainder Days After Minus 37: %s\n",
	//	remainderDays.Text(10))

	if !yearNumber.IsInt64() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Year Number cannot be converted to Int64!\n" +
			"Year Number='%v'\n",
			yearNumber.Text(10))
		return gregorianDateTime, err
	}

	year := yearNumber.Int64()

	var isLeapYear bool

	if year % int64(4) == 0 {

		isLeapYear = true

		if year % 100 == 0 {

			isLeapYear = false

			if  year % int64(400) == 0 {
				isLeapYear = true
			}
		}

	} else {

		isLeapYear = false

	}

	if !remainderDays.IsInt64() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Month Days cannot be converted to Int64!\n" +
			"Month Days='%v'\n",
			remainderDays.Text(10))
		return gregorianDateTime, err
	}

	lastYearOrdinalDayNo := remainderDays.Int64()

	if lastYearOrdinalDayNo < 0 ||
		lastYearOrdinalDayNo > 367 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Month Days value is INVALID!\n" +
			"Month Days='%v'\n", lastYearOrdinalDayNo)
		return gregorianDateTime, err
	}

	calUtil := CalendarUtility{}
	var yearAdjustment, month, day int

	yearAdjustment,
	month,
	day,
	err = calUtil.GetMonthDayFromOrdinalDayNo(
		lastYearOrdinalDayNo,
		isLeapYear,
		ePrefix)

	if err != nil {
		return gregorianDateTime, err
	}

	if yearAdjustment == -1 {
		year--
		month = 12
		day = 31
	} else if yearAdjustment == 1 {
		year++
		month = 1
		day = 1
	}

	// JDN Time to gross nanoseconds.


	gregorianDateTime = DateTransferDto{
		isLeapYear:          isLeapYear,
		year:                year,
		month:               month,
		day:                 day,
		hour:                hour,
		minute:              minute,
		second:              second,
		nanosecond:          nanosecond,
		tag:                 "",
		isThisInstanceValid: true,
		lock:                new(sync.Mutex),
	}

	return gregorianDateTime, err
}

// NumCalendarDaysForWholeYearsInterval - Computes the total
// number of 24-hour days in a period of years specified by
// input parameter 'wholeYearsInterval'. The number of total
// days is calculated in accordance with the Gregorian Calendar.
//
// Methodology:
//
// In the Gregorian calendar, three criteria must be taken
// into account to correctly identify leap years:
//
// 1. The year must be evenly divisible by 4;
//
// 2. If the year can also be evenly divided by 100, it is not
//    a leap year, unless...
//
//  3. The year is evenly divisible by 100 and the year is also
//     evenly divisible by 400. Then it is a leap year.
//
// According to these rules, the years 2000 and 2400 are leap years,
// while 1800, 1900, 2100, 2200, 2300, and 2500 are not leap years.
//
// For more information on the Gregorian Calendar and leap years,
// reference:
//
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//   https://www.timeanddate.com/date/leapyear.html
//
// The input parameter 'wholeYearsInterval' is defined as a series of
// contiguous whole, or complete, years consisting of either 365-days
// or 366-days (in the case of leap years).
//
// No partial years should be included in this interval.
//
//
func (calGreg *CalendarGregorianUtility) NumCalendarDaysForWholeYearsInterval(
	wholeYearsInterval int64) (totalDays int64) {

	if calGreg.lock == nil {
		calGreg.lock = &sync.Mutex{}
	}

	calGreg.lock.Lock()

	defer calGreg.lock.Unlock()

	totalDays = 0

	if wholeYearsInterval < 0 {
		wholeYearsInterval *= -1
	}

	if wholeYearsInterval == 0 {
		return 0
	}

	separator := strings.Repeat("*", 65)

	fmt.Println()
	fmt.Println("NumCalendarDaysForWholeYearsInterval() ")
	fmt.Println(separator)
	fmt.Printf("       Whole Years Interval: %v\n", wholeYearsInterval)

	if wholeYearsInterval >= 900 {

		numOfCycles := wholeYearsInterval / 400

		totalDays = numOfCycles * 146097

		fmt.Printf("  Number of 400-Year Cycles: %v\n", numOfCycles)
		fmt.Printf("Number of Days in %v-Cycles: %v\n", numOfCycles, totalDays)

		wholeYearsInterval = wholeYearsInterval - (numOfCycles * 400)

		fmt.Printf("  Number of Remainder Years: %v\n", wholeYearsInterval)
		fmt.Println(separator)
		fmt.Println()

	}

	totalDays += wholeYearsInterval * 365

	leapDays := wholeYearsInterval / 4

	skipLeapDays := wholeYearsInterval / 100

	totalDays += leapDays - skipLeapDays

	fmt.Println(separator)
	fmt.Printf("Total Days In wholeYearsInterval: %v\n",
		totalDays)
	fmt.Println(separator)
	fmt.Println()

	return totalDays
}
