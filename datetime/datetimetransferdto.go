package datetime

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

// DateTransferDto - This is a light data transfer object designed expressly
// for transmission of raw date t	ime information.
//
type DateTransferDto struct {
	isLeapYear bool
	year int64
	month int
	day int
	hour int
	minute int
	second int
	nanosecond int
	tag string
	isThisInstanceValid bool
	lock *sync.Mutex
}

// Compare - Receives a pointer to an incoming DateTransferDto instance and
// compares the internal date/time data fields to those of the current DateTransferDto
// instance. If successful, an integer value is returned describing the result
// of this comparison.
//
// IMPORTANT
//
// Complete accuracy can only be guaranteed if the 'isLeapYear' data fields are 
// correctly initialized in both the current DateTransferDto instance and in
// the second instance passed by input parameter, 'timeTransDto2'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dateTimeTransDto2   *DateTransferDto
//     - A pointer to another incoming DateTransferDto instance. The date/time
//       values of this second instance will be compared to those of the current
//       instance.
//
//
//  ePrefix             string
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
//  comparisonResult    int
//     - If successful this method will return will return an integer value
//       describing the result of comparing 'timeTransDto2' to the current 
//       DateTransferDto instance.
//
//       Comparison Result
//
//       -1 - The current DateTransferDto instance is LESS THAN 'timeTransDto2'.
//
//        0 - The current DateTransferDto instance is EQUAL to 'timeTransDto2'.
//
//       +1 - The current DateTransferDto instance is GREATER THAN 'timeTransDto2'.
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (dateTimeTransDto *DateTransferDto) Compare(
	dateTimeTransDto2 *DateTransferDto,
	ePrefix string) (comparisonResult int, err error) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	if dateTimeTransDto2.lock == nil {
		dateTimeTransDto2.lock = &sync.Mutex{}
	}

	dateTimeTransDto2.lock.Lock()

	defer dateTimeTransDto2.lock.Unlock()

	ePrefix += "DateTransferDto.Compare() "

	comparisonResult = -99
	err = nil

	calMech2 := CalendarUtility{}

	var timeTransDtoTotalNanoseconds, timeTransDto2TotalNanoseconds int64

	var timeTransDtoOrdinalDays, timeTrans2OrdinalDays int


	if dateTimeTransDto.year == dateTimeTransDto2.year {

		yearSign := 1

		if dateTimeTransDto.year < 0 {
			yearSign = -1
		}

		timeTransDtoOrdinalDays, err = calMech2.GetOrdinalDayNumber(
			dateTimeTransDto.isLeapYear,
			dateTimeTransDto.month,
			dateTimeTransDto.day,
			ePrefix + "timeTransDtoOrdinalDays- ")

		if err != nil {
			return comparisonResult, err
		}

		timeTrans2OrdinalDays, err = calMech2.GetOrdinalDayNumber(
			dateTimeTransDto2.isLeapYear,
			dateTimeTransDto2.month,
			dateTimeTransDto2.day,
			ePrefix + "timeTrans2OrdinalDays- ")

		if err != nil {
			return comparisonResult, err
		}

		timeTransDtoOrdinalDays *= yearSign
		timeTrans2OrdinalDays *= yearSign

		if timeTransDtoOrdinalDays == timeTrans2OrdinalDays {

			timeTransDtoTotalNanoseconds, err =
				calMech2.GetTimeTotalNanoseconds(
					dateTimeTransDto.hour,
					dateTimeTransDto.month,
					dateTimeTransDto.second,
					dateTimeTransDto.nanosecond,
					ePrefix + "dateTimeTransDto- ")

			if err != nil {
				return comparisonResult, err
			}

			timeTransDto2TotalNanoseconds, err =
				calMech2.GetTimeTotalNanoseconds(
					dateTimeTransDto2.hour,
					dateTimeTransDto2.minute,
					dateTimeTransDto2.second,
					dateTimeTransDto2.nanosecond,
					ePrefix + "dateTimeTransDto2- ")

			if err != nil {
				return comparisonResult, err
			}

			timeTransDtoTotalNanoseconds*= int64(yearSign)

			timeTransDto2TotalNanoseconds *= int64(yearSign)

			if timeTransDtoTotalNanoseconds == timeTransDto2TotalNanoseconds {
				comparisonResult = 0

			} else if timeTransDtoTotalNanoseconds > timeTransDto2TotalNanoseconds {
				comparisonResult = 1
			} else {
				comparisonResult = -1
			}

		} else if timeTransDtoOrdinalDays > timeTrans2OrdinalDays {
			comparisonResult = 1
		} else {
			comparisonResult = -1
		}

	} else if dateTimeTransDto.year > dateTimeTransDto2.year {
		comparisonResult = 1
	} else {
		comparisonResult = -1
	}

	if comparisonResult < -1 || comparisonResult > 1 {
		err = errors.New(ePrefix + "\n" +
			"Error: The final comparison was inconclusive. CODING ERROR!\n")
	}
	
	return comparisonResult, err
}

// CompareAreYearsAdjacent - Compares the year value of the current
// 'DateTransferDto' instance, 'dateTimeTransDto.year', to the
// year value encapsulated by a second, input parameter instance,
// dateTimeTransDto2.year, to determine if the two year values
// differ a value of plus or minus and are therefore 'adjacent'
// years. 'Adjacent' year values differ by 1 such that year1 - year2
// is always equal to plus or minus 1 (+1 or -1).
//
// For the purposes of this description, the instance , dateTimeTransDto.year,
// is year1 and the input parameter instance, dateTimeTransDto2.year, is
// year2.
//
// There are three possible outcomes from this comparison:
//
//  1. If the current year1 value, dateTimeTransDto.year,
//     is adjacent to year2, and (year1 - year2) is equal
//     to +1, this method will return an integer value of
//     plus 1 ('+1'). In this case, year1 is GREATER than
//     year2 by a value of '+1'.
//
//  2. If the current year1 value (dateTimeTransDto.year)
//     is NOT adjacent to year2 this method will return an
//     integer value of zero ('0'). In this case (year1 - year2)
//     is Greater than 1, Less than 1 or zero.
//
//  3. If the current year1 value, dateTimeTransDto.year,
//     is adjacent to year2 and  (year1 - year2) is equal
//     to minus 1 ('-1'), this method will return an integer
//     value of minus 1 ('-1'). In this case, year1 is LESS
//     than year2 by a value of '-1'.
//
//
//  For all cases, this method will return one of these three
//  integer values:
//
//                -1, 0 or +1.
//
func (dateTimeTransDto *DateTransferDto) CompareAreYearsAdjacent(
	dateTimeTransDto2 *DateTransferDto) (
	yearAdjacentComparison int ) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	if dateTimeTransDto2.lock == nil {
		dateTimeTransDto2.lock = &sync.Mutex{}
	}

	dateTimeTransDto2.lock.Lock()

	defer dateTimeTransDto2.lock.Unlock()

	variance := dateTimeTransDto2.year - dateTimeTransDto.year

	if variance == -1 {
		// year1 is 1 less than year2
		yearAdjacentComparison = -1
	} else if variance == 1 {
		// year1 is 1 more than year2
		yearAdjacentComparison = 1
	} else {
		// year1 and year2 are NOT adjacent years
		yearAdjacentComparison = 0
	}

	return yearAdjacentComparison
}


// CompareTotalTimeNanoseconds - Compares the total time nanoseconds values for
// the current DateTransferDto instance and another DateTransferDto
// instance passed as an input parameter.
//
// The comparison result is returned as an integer value.
//
// The total time nanoseconds value is the total value of hours, minutes,
// seconds and nanoseconds converted to total nanoseconds.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dateTimeTransDto2   *DateTransferDto
//     - A pointer to another incoming DateTransferDto instance. The total
//       time nanoseconds value for the current DateTransferDto instance
//       will be compared to total time nanoseconds value of this second instance.
//
//
//  ePrefix            string
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
//  comparisonResult    int
//     - If successful this method will return will return an integer value
//       describing the result of comparing 'timeTransDto.year' value to the
//       'dateTimeTransDto2.year' value.
//
//       Comparison Result
//
//       -1 - The current DateTransferDto instance total time nanoseconds
//            value is LESS THAN the 'timeTransDto2' total time nanoseconds value.
//
//        0 - The current DateTransferDto instance total time nanoseconds
//            value is EQUAL to the 'timeTransDto2' total time nanoseconds value.
//
//       +1 - The current DateTransferDto instance total time nanoseconds
//            value is GREATER THAN the 'timeTransDto2' year value.
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
func (dateTimeTransDto *DateTransferDto) CompareTotalTimeNanoseconds(
	dateTimeTransDto2 *DateTransferDto,
	ePrefix string) (comparisonResult int, err error) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	if dateTimeTransDto2.lock == nil {
		dateTimeTransDto2.lock = &sync.Mutex{}
	}

	dateTimeTransDto2.lock.Lock()

	defer dateTimeTransDto2.lock.Unlock()

	ePrefix += "DateTransferDto.CompareTotalTimeNanoseconds() "

	comparisonResult = -99

	totalTimeInNanoseconds := int64(0)

	totalTimeInNanoseconds += int64(dateTimeTransDto.hour) * int64(time.Hour)

	totalTimeInNanoseconds += int64(dateTimeTransDto.minute) * int64(time.Minute)

	totalTimeInNanoseconds += int64(dateTimeTransDto.second) * int64(time.Second)

	totalTimeInNanoseconds +=  int64(dateTimeTransDto.nanosecond)

	if totalTimeInNanoseconds > int64(time.Hour) * 24 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: The Current Instance total time in nanoseconds exceeds 24-hours\n" +
			"Total Nanosecond='%v'\n", totalTimeInNanoseconds)
		return comparisonResult, err
	}


	var totalTimeInNanoseconds2 int64

	totalTimeInNanoseconds2, err = dateTimeTransDto2.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return comparisonResult, err
	}

	if totalTimeInNanoseconds == totalTimeInNanoseconds2 {
		comparisonResult = 0
	} else if totalTimeInNanoseconds > totalTimeInNanoseconds2{
		comparisonResult = 1
	} else {
		// totalTimeInNanoseconds < totalTimeInNanoseconds2
		comparisonResult = -1
	}

	return comparisonResult, err
}

// CompareYears - Compares the years values for the current DateTransferDto
// instance and another DateTransferDto instance passed as an input
// parameter.
//
// The comparison result is returned as an integer value.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dateTimeTransDto2   *DateTransferDto
//     - A pointer to another incoming DateTransferDto instance. The year
//       value for the current DateTransferDto instance will be compared
//       to the year value of this second instance.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  comparisonResult    int
//     - If successful this method will return will return an integer value
//       describing the result of comparing 'timeTransDto.year' value to the
//       'dateTimeTransDto2.year' value.
//
//       Comparison Result
//
//       -1 - The current DateTransferDto instance year value is LESS THAN the
//            'timeTransDto2' year value.
//
//        0 - The current DateTransferDto instance year value is EQUAL to the
//            'timeTransDto2' year value.
//
//       +1 - The current DateTransferDto instance year value is GREATER THAN
//            the 'timeTransDto2' year value.
//
func (dateTimeTransDto *DateTransferDto) CompareYears(
	dateTimeTransDto2 *DateTransferDto) (comparisonResult int) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	if dateTimeTransDto2.lock == nil {
		dateTimeTransDto2.lock = &sync.Mutex{}
	}

	dateTimeTransDto2.lock.Lock()

	defer dateTimeTransDto2.lock.Unlock()

	comparisonResult = -99

	if dateTimeTransDto.year == dateTimeTransDto2.year {
		comparisonResult = 0
	} else if dateTimeTransDto.year > dateTimeTransDto2.year {
		comparisonResult = 1
	} else {
		// dateTimeTransDto.year < dateTimeTransDto2.year
		comparisonResult = -1
	}

	return comparisonResult
}

// CopyIn - Receives a pointer to an incoming DateTransferDto and copies all
// of the internal data fields to the current DateTransferDto instance.
//
func (dateTimeTransDto *DateTransferDto) CopyIn(
	dateTimeTransDto2 *DateTransferDto) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	dateTimeTransDto.isLeapYear = dateTimeTransDto2.isLeapYear
	dateTimeTransDto.year = dateTimeTransDto2.year
	dateTimeTransDto.month = dateTimeTransDto2.month
	dateTimeTransDto.day = dateTimeTransDto2.day
	dateTimeTransDto.hour = dateTimeTransDto2.hour
	dateTimeTransDto.minute = dateTimeTransDto2.minute
	dateTimeTransDto.second = dateTimeTransDto2.second
	dateTimeTransDto.nanosecond = dateTimeTransDto2.nanosecond
	dateTimeTransDto.tag = dateTimeTransDto2.tag
	dateTimeTransDto.isThisInstanceValid = dateTimeTransDto2.isThisInstanceValid

	return
}

// CopyOut - Returns a deep copy of the current DateTransferDto
// instance.
//
func (dateTimeTransDto *DateTransferDto) CopyOut() (
	newDateTimeDto DateTransferDto) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	newDateTimeDto = DateTransferDto{}

	newDateTimeDto.isLeapYear = dateTimeTransDto.isLeapYear
	newDateTimeDto.year = dateTimeTransDto.year
	newDateTimeDto.month = dateTimeTransDto.month
	newDateTimeDto.day = dateTimeTransDto.day
	newDateTimeDto.hour = dateTimeTransDto.hour
	newDateTimeDto.minute = dateTimeTransDto.minute
	newDateTimeDto.second = dateTimeTransDto.second
	newDateTimeDto.nanosecond = dateTimeTransDto.nanosecond
	newDateTimeDto.tag = dateTimeTransDto.tag
	newDateTimeDto.isThisInstanceValid = dateTimeTransDto.isThisInstanceValid
	newDateTimeDto.lock = &sync.Mutex{}

	return newDateTimeDto
}

// Empty - Resets the internal data fields of the current 
// DateTransferDto instance to their zero values. Effectively,
// the current DateTransferDto instance is rendered blank and
// invalid.
//
func (dateTimeTransDto *DateTransferDto) Empty() {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	dateTimeTransDto.isLeapYear = false
	dateTimeTransDto.year = 0
	dateTimeTransDto.month = 0
	dateTimeTransDto.day = 0
	dateTimeTransDto.hour = 0
	dateTimeTransDto.minute = 0
	dateTimeTransDto.second = 0
	dateTimeTransDto.nanosecond = 0
	dateTimeTransDto.tag = ""
	dateTimeTransDto.isThisInstanceValid = false

}

// ExchangeValues - Performs a data exchange. The data fields
// from incoming DateTransferDto instance 'dateTimeTransDto2'
// are copied to the current DateTransferDto instance and the
// data fields from the current instance are in turn copied to
// 'dateTimeTransDto2'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTimeTransDto2  *DateTransferDto
//     - A pointer to another incoming DateTransferDto instance. The
//       internal data values of this second instance will be populated 
//       with data fields from the current instance.
//
//
//  ePrefix            string
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
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (dateTimeTransDto *DateTransferDto) ExchangeValues(
	dateTimeTransDto2 *DateTransferDto) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	if dateTimeTransDto2.lock == nil {
		dateTimeTransDto2.lock = &sync.Mutex{}
	}

	dateTimeTransDto2.lock.Lock()

	defer dateTimeTransDto2.lock.Unlock()
	
	tempTimeTransDto := DateTransferDto{}
	tempTimeTransDto.lock = &sync.Mutex{}

	tempTimeTransDto.isLeapYear = dateTimeTransDto.isLeapYear
	tempTimeTransDto.year = dateTimeTransDto.year
	tempTimeTransDto.month = dateTimeTransDto.month
	tempTimeTransDto.day = dateTimeTransDto.day
	tempTimeTransDto.hour = dateTimeTransDto.hour
	tempTimeTransDto.minute = dateTimeTransDto.minute
	tempTimeTransDto.second = dateTimeTransDto.second
	tempTimeTransDto.nanosecond = dateTimeTransDto.nanosecond
	tempTimeTransDto.tag = dateTimeTransDto.tag
	tempTimeTransDto.isThisInstanceValid = dateTimeTransDto.isThisInstanceValid

	dateTimeTransDto.isLeapYear = dateTimeTransDto2.isLeapYear
	dateTimeTransDto.year = dateTimeTransDto2.year
	dateTimeTransDto.month = dateTimeTransDto2.month
	dateTimeTransDto.day = dateTimeTransDto2.day
	dateTimeTransDto.hour = dateTimeTransDto2.hour
	dateTimeTransDto.minute = dateTimeTransDto2.minute
	dateTimeTransDto.second = dateTimeTransDto2.second
	dateTimeTransDto.nanosecond = dateTimeTransDto2.nanosecond
	dateTimeTransDto.tag = dateTimeTransDto2.tag
	dateTimeTransDto.isThisInstanceValid = dateTimeTransDto2.isThisInstanceValid

	dateTimeTransDto2.isLeapYear = tempTimeTransDto.isLeapYear
	dateTimeTransDto2.year = tempTimeTransDto.year
	dateTimeTransDto2.month = tempTimeTransDto.month
	dateTimeTransDto2.day = tempTimeTransDto.day
	dateTimeTransDto2.hour = tempTimeTransDto.hour
	dateTimeTransDto2.minute = tempTimeTransDto.minute
	dateTimeTransDto2.second = tempTimeTransDto.second
	dateTimeTransDto2.nanosecond = tempTimeTransDto.nanosecond
	dateTimeTransDto2.tag = tempTimeTransDto.tag
	dateTimeTransDto2.isThisInstanceValid = tempTimeTransDto.isThisInstanceValid

}

// GetAbsoluteYearValue - Returns the internal member variable 'year'
// as an absolute or positive value. See instance method,
// 'GetYearNumberSign()'.
//
func (dateTimeTransDto *DateTransferDto) GetAbsoluteYearValue() int64 {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	year := dateTimeTransDto.year

	if year < 0 {
		year *= -1
	}

	return year
}

// GetDateTime - Returns the individual date/time components of 
// the current DateTransferDto instance.
//
func (dateTimeTransDto *DateTransferDto) GetDateTime() (
	isLeapYear bool,
	year int64,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

return dateTimeTransDto.isLeapYear,
	dateTimeTransDto.year,
	dateTimeTransDto.month,
	dateTimeTransDto.day,
	dateTimeTransDto.hour,
	dateTimeTransDto.minute,
	dateTimeTransDto.second,
	dateTimeTransDto.nanosecond
}

// GetDate - Returns the date components of the current DateTransferDto
// instance.
func (dateTimeTransDto *DateTransferDto) GetDate() (
	year int64, 
	month int, 
	day int) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()


	return dateTimeTransDto.year, 
				dateTimeTransDto.month, 
				dateTimeTransDto.day
}

// GetDay - Returns the day number associated with this
// current DateTransferDto instance.
//
func (dateTimeTransDto *DateTransferDto) GetDay() int {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.day
}

// GetHour - Returns the hour number associated with this
// current DateTransferDto instance.
//
func (dateTimeTransDto *DateTransferDto) GetHour() int {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.hour
}

// GetIsLeapYear - Returns a the internal member variable
// 'isLeapYear' as a boolean value. The user is responsible
// for setting this value. See instance method, 'SetIsLeapYear()'.
//
func (dateTimeTransDto *DateTransferDto) GetIsLeapYear() bool {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.isLeapYear
}

// GetMinute - Returns the minute number associated with this
// current DateTransferDto instance.
//
func (dateTimeTransDto *DateTransferDto) GetMinute() int {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.minute
}

// GetMonth - Returns the month number associated with this
// current DateTransferDto instance.
//
func (dateTimeTransDto *DateTransferDto) GetMonth() int {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.month
}

// GetNanosecond - Returns the nanosecond number associated with this
// current DateTransferDto instance.
//
func (dateTimeTransDto *DateTransferDto) GetNanosecond() int {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.nanosecond
}

// GetOrdinalDayNumber - Returns the ordinal day number of the year
// for the date/time specified by this DateTransferDto instance.
// The accuracy of the result in turns depends on the accuracy of the
// 'isLeapYear' setting. To set a correct value for 'isLeapYear', see
// instance method, 'SetIsLeapYear()'.
//
// For more information on Ordinal Day Number, reference:
//   https://en.wikipedia.org/wiki/Ordinal_date
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  ePrefix            string
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
//  ordinalDayNo        int
//     - If successful this method will return will return an integer value
//       designating the ordinal day number of the year represented by the
//       year, month and day encapsulated by the current DateTransferDto
//       instance. Accuracy depends on the 'isLeapYear' value. To set the
//       'isLeapYear' value see instance method, 'SetIsLeapYear()'.
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
func (dateTimeTransDto *DateTransferDto) GetOrdinalDayNumber(
	ePrefix string) (
	ordinalDayNo int,
	err error) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.GetOrdinalDayNumber() "
	ordinalDayNo = -99
	err = nil

	calUtil := CalendarUtility{}

	ordinalDayNo, err = calUtil.GetOrdinalDayNumber(
		dateTimeTransDto.isLeapYear,
		dateTimeTransDto.month,
		dateTimeTransDto.day,
		ePrefix)

	return ordinalDayNo, err
}

// GetRemainingDaysInYear - Returns a positive integer value identifying
// the number of days remaining in the year based on the date/time encapsulated
// by the current DateTransferDto instance.
//
// This calculation is performed by subtracting the date/time's ordinal day
// day number from the number of days in the current year. The accuracy of
// the result depends largely on the accuracy of the 'isLeap' year setting
// for the current DateTransferDto instance. The 'isLeap' setting can
// be controlled through instance method, 'SetIsLeapYear()'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
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
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (dateTimeTransDto *DateTransferDto) GetRemainingDaysInYear(
	ePrefix string) (
	remainingDaysOfYear int,
	err error) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.GetRemainingDaysInYear() "

	remainingDaysOfYear = -999
	err = nil

	year := 365

	if dateTimeTransDto.isLeapYear {
		year++
	}

	ordinalDayNo := -99

	calUtil := CalendarUtility{}

	ordinalDayNo, err = calUtil.GetOrdinalDayNumber(
		dateTimeTransDto.isLeapYear,
		dateTimeTransDto.month,
		dateTimeTransDto.day,
		ePrefix)

	if err != nil {
		return remainingDaysOfYear, err
	}


	remainingDaysOfYear = year - ordinalDayNo

	return remainingDaysOfYear, err
}

// GetSecond - Returns the second number associated with this
// current DateTransferDto instance.
//
func (dateTimeTransDto *DateTransferDto) GetSecond() int {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.second
}

// GetTime - Returns the time components of the current DateTransferDto
// instance.
func (dateTimeTransDto *DateTransferDto) GetTime() (
	hour int, 
	minute int, 
	second int,
	nanosecond int) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()


	return dateTimeTransDto.hour, 
				dateTimeTransDto.minute, 
				dateTimeTransDto.second,
				dateTimeTransDto.nanosecond
}

// GetTotalTimeInNanoseconds - For the current DateTransferDto instance
// this method returns the sum of Hours, Minutes, Seconds and Nanoseconds
// expressed as total nanoseconds.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//  totalTimeInNanoseconds     int64
//     - If successful this method will return a the total of
//       Hours, Minutes, Seconds and Nanoseconds for the current
//       instance of DateTransferDto expressed as total
//       nanoseconds.
//
//  err                         error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
func (dateTimeTransDto *DateTransferDto) GetTotalTimeInNanoseconds(
	ePrefix string) (
	totalTimeInNanoseconds int64, err error) {


	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.GetTotalTimeInNanoseconds() "

	totalTimeInNanoseconds = math.MaxInt64

	if !dateTimeTransDto.isThisInstanceValid {
		err = errors.New(ePrefix + "\n" +
			"Error: The current instance of 'DateTransferDto' is INVALID!\n")

		return totalTimeInNanoseconds, err
	}

	totalTimeInNanoseconds = 0

	totalTimeInNanoseconds += int64(dateTimeTransDto.hour) * int64(time.Hour)

	totalTimeInNanoseconds += int64(dateTimeTransDto.minute) * int64(time.Minute)

	totalTimeInNanoseconds += int64(dateTimeTransDto.second) * int64(time.Second)

	totalTimeInNanoseconds +=  int64(dateTimeTransDto.nanosecond)

	if totalTimeInNanoseconds > int64(time.Hour) * 24 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Total Time in nanoseconds exceeds 24-hours\n" +
			"Total Nanosecond='%v'\n", totalTimeInNanoseconds)
	}

	return totalTimeInNanoseconds, err
}

// GetYear - Returns the year value for the current DateTransferDto
// instance.
//
func (dateTimeTransDto *DateTransferDto) GetYear() int64 {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.year
}

// GetYear - Returns the the number of days in the year specified by
// the current DateTransferDto instance.
//
// Important: This method will return 365 or 366 based on the
// the setting for 'isLeapYear'. The accuracy of this method
// result depends entirely on the setting for 'isLeapYear'.
// this value.
//
func (dateTimeTransDto *DateTransferDto) GetYearDays() int64 {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.year
}

// GetYearNumberSign - Returns a code for the numeric sign of the
// year value for the current DateTransferDto instance.
//
// Return Values:
//  +1 = Year Value is Positive
//   0 = Year Value is Zero
//  -1 = Year Value is Negative
func (dateTimeTransDto *DateTransferDto) GetYearNumberSign() int {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	if dateTimeTransDto.year == 0 {
		return 0
	} else if dateTimeTransDto.year > 0 {
		return 1
	}


	// Value must be negative; less than zero.
	return -1
}

// IsBeginningOfYear - Returns a boolean value signaling whether
// the date/time represented by the current DateTransferDto
// instance is equal to the beginning of the year.
//
// Effectively, this method only returns 'true' when month== 1,
// day== 1, hour=0, minute=0, second=0, and nanosecond
// equals 0. (January 1st 00:00:00.000000000 a.k.a "Midnight")
//
func (dateTimeTransDto *DateTransferDto) IsBeginningOfYear() bool {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	if dateTimeTransDto.month == 1 &&
		dateTimeTransDto.day == 1 &&
		dateTimeTransDto.hour == 0 &&
		dateTimeTransDto.minute == 0 &&
		dateTimeTransDto.second== 0 &&
		dateTimeTransDto.nanosecond == 0 {

		return true
	}

	return false
}

// IsEndOfYear - Returns a boolean value signaling whether the
// date/time represented by the current DateTransferDto
// instance is equal to the end of the year.
//
// Effectively, this method only returns 'true' when month==12,
// day==31, hour=23, minute=59, second=59/60, and nanosecond
// equals 999,999,999. (December 31st 23:59:59.999999999)
//
// The use of 60-seconds occurs when 'leap seconds' are applied.
// For a discussion of 'leap seconds', reference:
//  https://en.wikipedia.org/wiki/Leap_second
//
func (dateTimeTransDto *DateTransferDto) IsEndOfYear() bool {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	if dateTimeTransDto.month == 12 &&
		dateTimeTransDto.day == 31 &&
		dateTimeTransDto.hour == 23 &&
		dateTimeTransDto.minute == 59 &&
		(dateTimeTransDto.second== 59 ||
			dateTimeTransDto.second == 60) &&
		dateTimeTransDto.nanosecond == 999999999 {

		return true
	}

	return false
}

// IsValidDateTime - Analyzes the internal date fields of the
// current DateTransferDto instance to determine validity.
//
// Important: Do NOT call this method without first setting a
// correct value for the internal member variable,'isLeapYear'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//  err                error
//     - If successful, and the date/time values specified by the current
//       DateTransferDto instance are valid, the returned error Type
//       is set equal to 'nil'.
//
//       If errors are encountered during processing, or if the date/time
//       values specified by the current DateTransferDto instance are
//       invalid, the returned error Type will encapsulate an appropriate
//       error message. Note this error message will be prefixed with the
//       method chain and text passed by input parameter, 'ePrefix'.
//
func (dateTimeTransDto *DateTransferDto) IsValidDateTime(
	ePrefix string) ( err error) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	err = nil

	ePrefix += "DateTransferDto.IsValidDateTime() "

	if !dateTimeTransDto.isThisInstanceValid {
		return 	errors.New(ePrefix + "\n" +
			"Error: The current DateTransferDto instance was never properly initialized!\n")
	}

	calUtil := CalendarUtility{}

	return calUtil.IsValidDateTimeComponents(
		dateTimeTransDto.isLeapYear,
		dateTimeTransDto.month,
		dateTimeTransDto.day,
		dateTimeTransDto.hour,
		dateTimeTransDto.minute,
		dateTimeTransDto.second,
		dateTimeTransDto.nanosecond,
		ePrefix)
}

// IsValidInitialize - If this instance returns 'true' it signals that
// the current DateTransferDto instance has been correctly
// initialized and populated using one of the 'New' methods.
//
// Importantly, this method does NOT validate the internal data fields
// which comprise the date time specification. To validate the date time
// values, invoke instance method IsValidDateTime().
//
func (dateTimeTransDto *DateTransferDto) IsValidInitialize() bool {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	return dateTimeTransDto.isThisInstanceValid
}

// New - Creates and returns an new instance of DateTransferDto. This is a
// light data transfer object designed expressly for transmission of raw 
// date time information.
//
// The input parameter 'isLeapYear' cannot be validated since the applicable
// calendar is unknown. The user is therefore responsible for providing a
// true and correct value. This value may be subsequently updated at a later
// time using the instance method, 'SetLeapYear()'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  isLeapYear         bool
//     - If 'true' it signals that the input parameter 'targetYear' is a leap year.
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
//       The valid range is 1 - 31 inclusive.
//
//  targetHour         int
//     - The hour time component for this date/time specification.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  targetMinute       int
//     - The minute time component for this date/time specification.
//       The valid range is 0 - 59 inclusive
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
//
//  ePrefix            string
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
//  newDateTimeDto     DateTransferDto
//     - If successful this method will return a new, fully populated instance
//       of type DateTransferDto.
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (dateTimeTransDto DateTransferDto) New (
	isLeapYear bool,
	year int64,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (newDateTimeDto DateTransferDto, err error) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	ePrefix += "DateTransferDto.New () "

	newDateTimeDto = DateTransferDto{}
	newDateTimeDto.isThisInstanceValid = false

	err = nil

	if month < 1 || month > 12 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'month' is invalid!\n" +
			"month='%v'\n", month)
		return newDateTimeDto, err
	}

	if day > 31 || day < 1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'day' is invalid!\n" +
			"day='%v'\n", day)
		return newDateTimeDto, err
	}

	if hour > 24 || hour < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hour' is invalid!\n" +
			"hour='%v'\n", hour)
		return newDateTimeDto, err
	}

	if minute < 0 || minute > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minute' is invalid!\n" +
			"minute='%v'\n", minute)
		return newDateTimeDto, err
	}

	calUtil := CalendarUtility{}

	isValidSecond := calUtil.IsValidSecond(
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
		return newDateTimeDto, err
	}

	if nanosecond < 0 || nanosecond > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecond' is invalid!\n" +
			"nanosecond='%v'\n", nanosecond)
		return newDateTimeDto, err
	}

	newDateTimeDto.isLeapYear = isLeapYear
	newDateTimeDto.year = year
	newDateTimeDto.month = month
	newDateTimeDto.day = day
	newDateTimeDto.hour = hour
	newDateTimeDto.minute = minute
	newDateTimeDto.second = second
	newDateTimeDto.nanosecond = nanosecond
	newDateTimeDto.lock = &sync.Mutex{}
	newDateTimeDto.isThisInstanceValid = true

	return newDateTimeDto, err
}

// SetIsLeapYear - Sets the internal value for 'isLeapYear'.
//
func (dateTimeTransDto *DateTransferDto) SetIsLeapYear( isLeapYear bool) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	dateTimeTransDto.isLeapYear = isLeapYear

}

// SetTag - Sets the internal member variable 'tag'.
// This text string is used to identify or describe
// the current DateTransferDto instance.
//
func (dateTimeTransDto *DateTransferDto) SetTag(tag string) {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	dateTimeTransDto.tag = tag
}

func (dateTimeTransDto *DateTransferDto) String() string {

	if dateTimeTransDto.lock == nil {
		dateTimeTransDto.lock = &sync.Mutex{}
	}

	dateTimeTransDto.lock.Lock()

	defer dateTimeTransDto.lock.Unlock()

	outStr := "\n"
	separator := strings.Repeat("-", 65)
	outStr+= separator + "\n"

	if len(dateTimeTransDto.tag) > 0 {
		outStr += dateTimeTransDto.tag + "\n"
	}
	outStr += "DateTransferDto Instance\n"
	outStr += separator + "\n"
	outStr +=  fmt.Sprintf("         isLeapYear: %v\n", dateTimeTransDto.isLeapYear)
	outStr +=  fmt.Sprintf("               year: %v\n", dateTimeTransDto.year)
	outStr +=  fmt.Sprintf("              month: %v\n", dateTimeTransDto.month)
	outStr +=  fmt.Sprintf("                day: %v\n", dateTimeTransDto.day)
	outStr +=  fmt.Sprintf("               hour: %v\n", dateTimeTransDto.hour)
	outStr +=  fmt.Sprintf("             minute: %v\n", dateTimeTransDto.minute)
	outStr +=  fmt.Sprintf("             second: %v\n", dateTimeTransDto.second)
	outStr +=  fmt.Sprintf("         nanosecond: %v\n", dateTimeTransDto.nanosecond)
	outStr +=  fmt.Sprintf("isThisInstanceValid: %v\n", dateTimeTransDto.isThisInstanceValid)
	outStr +=  separator + "\n\n"

	return outStr
}