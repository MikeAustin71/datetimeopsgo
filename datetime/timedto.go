package datetime

import (
	"errors"
	"sync"
	"time"
)

/*
 TimeDto
  =======

 This source file is located in source code repository:
   https://github.com/MikeAustin71/datetimeopsgo.git'

 This source code file is located at:
   MikeAustin71\datetimeopsgo\datetime\timedto.go

*/

// TimeDto -   is a collection of time element values. Time
//             element values are represented by Years, Months,
//             Weeks, WeekDays, DateDays, Hours, Minutes, Seconds,
//             Milliseconds, Microseconds and Nanoseconds.
//
// TimeDto data fields are designed to store one of two
// types of time components:
//
//    (1)  A specific point in time (date time).
//                  or
//    (2) Incremental time or time duration which is useful in
//          adding or subtracting time values. Note that this
//          structure does not track time location or time zone.
//          For a fully supported date time structure, review
//          type DateTzDto located in source file 'datetzdto.go'
//          Note: TimeDto is part of the DateTzDto structure.
//
type TimeDto struct {
	Years                int // Number of Years
	Months               int // Number of Months
	Weeks                int // Number of Weeks
	WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
	DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
	Hours                int // Number of Hours.
	Minutes              int // Number of Minutes
	Seconds              int // Number of Seconds
	Milliseconds         int // Number of Milliseconds
	Microseconds         int // Number of Microseconds
	Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
	TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
	                         //  plus remaining Nanoseconds
	TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
	                         //  + Seconds + Milliseconds + Nanoseconds
	lock          sync.Mutex // Used for coordinating thread safe operations.
}

// AddTimeDto - Adds time to the current TimeDto. The amount of time added
// is provided by the input parameter 't2Dto' of type TimeDto.
//
// Date time math uses timezone UTC.
//
//  Input Parameters
//  ================
//
//  t2Dto     TimeDto  - The amount of time to be added to the current TimeDto
//                       data fields.
//
//
func (tDto *TimeDto) AddTimeDto(t2Dto TimeDto) error {

	tDto.lock.Lock()
	
	defer tDto.lock.Unlock()
	
	ePrefix := "TimeDto.AddTimeDto() "

	tDtoUtil := timeDtoUtility{}

	return tDtoUtil.addTimeDto(tDto, &t2Dto, ePrefix)
}

// CopyOut - Creates a new TimeDto instance
// which precisely duplicates the current TimeDto
// instance, and returns it to the calling function.
func (tDto *TimeDto) CopyOut() TimeDto {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	tDtoUtil := timeDtoUtility{}

	ePrefix := "TimeDto.CopyOut() "

	return tDtoUtil.copyOut(tDto, ePrefix)
}

// CopyIn - Receives a TimeDto input parameter, 'tDto2'
// and proceeds to copy all 'tDto2' data fields into
// the current TimeDto data fields. When this method
// completes, 'tDto' will be equivalent to 'tDto2'.
//
func (tDto *TimeDto) CopyIn(t2Dto TimeDto) {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.CopyIn() "

	tDtoUtil := timeDtoUtility{}

	tDtoUtil.copyIn(tDto, &t2Dto, ePrefix)

	return
}

// ConvertToAbsoluteValues - Converts time components
// (Years, months, weeks days, hours, seconds, etc.)
// to absolute values.
func (tDto *TimeDto) ConvertToAbsoluteValues() {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.ConvertToAbsoluteValues() "

	tDtoUtil := timeDtoUtility{}

	tDtoUtil.convertToAbsoluteValues(tDto, ePrefix)

}

// ConvertToNegativeValues - Multiplies time component
// values by -1
//
func (tDto *TimeDto) ConvertToNegativeValues() {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.ConvertToNegativeValues() "

	tDtoUtil := timeDtoUtility{}

	tDtoUtil.convertToAbsoluteValues(tDto, ePrefix)

	tDtoUtil.convertToNegativeValues(tDto, ePrefix)
}

// Empty - returns all TimeDto data fields to their
// uninitialized or zero state.
//
func (tDto *TimeDto) Empty() {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.Empty() "

	tDtoUtil := timeDtoUtility{}

	tDtoUtil.empty(tDto, ePrefix)

	return
}

// Equal - Compares the data fields of input parameter TimeDto, 'tDto2',
// to the data fields of the current TimeDto, 'tDto'. If the two sets of
// data fields are equal in all respects, this method returns 'true'.
// Otherwise, the method returns false.
//
func (tDto *TimeDto) Equal(t2Dto TimeDto) bool {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.Equal() "

	tDtoUtil := timeDtoUtility{}

	return tDtoUtil.equalTimeDtos(tDto, &t2Dto, ePrefix)
}

// GetDateTime - Analyzes the current TimeDto instance and computes
// an equivalent date time (time.Time). The calling function must
// pass in a valid time zone location.
//
// Input Parameter
// ===============
//
// timeZoneLocation string - time zone location must be designated as one of three values.
//
//   (1) The string 'Local' - selects the local time zone
//                            location for the host computer.
//
//   (2) IANA Time Zone Location -
//      See https://golang.org/pkg/time/#LoadLocation
//      and https://www.iana.org/time-zones to ensure that
//      the IANA Time Zone Database is properly configured
//      on your system. Note: IANA Time Zone Data base is
//      equivalent to 'tz database'.
//     Examples:
//      "America/New_York"
//      "America/Chicago"
//      "America/Denver"
//      "America/Los_Angeles"
//      "Pacific/Honolulu"
//      "Etc/UTC" = GMT or UTC
//
//    (3) A Military Time Zone
//        Reference:
//         https://en.wikipedia.org/wiki/List_of_military_time_zones
//         http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//         https://www.timeanddate.com/time/zones/military
//         https://www.timeanddate.com/worldclock/timezone/alpha
//         https://www.timeanddate.com/time/map/
//
//        Examples:
//          "Alpha"   or A
//          "Bravo"   or B
//          "Charlie" or C
//          "Delta"   or D
//          "Zulu"    or Z
//
func (tDto *TimeDto) GetDateTime(timeZoneLocationName string) (time.Time, error) {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.GetDateTime() "

	dtUtil := DTimeUtility{}

	_,
	_,
	_,
	tzLocPtr,
	_,
	err := dtUtil.GetTimeZoneFromName(timeZoneLocationName, ePrefix)

	if err != nil {
		return time.Time{}, err
	}

return time.Date(tDto.Years,
		time.Month(tDto.Months),
		tDto.DateDays,
		tDto.Hours,
		tDto.Minutes,
		tDto.Seconds,
		tDto.TotSubSecNanoseconds,
		tzLocPtr), nil
}

// IsEmpty - Returns 'true' if all data fields in the current
// TimeDto instance are equal to zero or equal to their
// uninitialized values.
func (tDto *TimeDto) IsEmpty() bool {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.GetDateTime() "

	tDtoUtil := timeDtoUtility{}

	return tDtoUtil.isEmpty(tDto, ePrefix)
}

// IsValid - Returns an error if the current tDto instance is invalid.
// Otherwise, if successful, this method returns 'nil'.
func (tDto *TimeDto) IsValid() error {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.IsValid() "

	tDtoUtil := timeDtoUtility{}

	return tDtoUtil.isValidDateTimeDto(tDto, ePrefix)
}

// New - Returns a new TimeDto instance based on time element
// input parameters.
//
// Be advised that all time elements are normalized. That is, negative
// time values are converted and stored as positive time elements suitable
// for conversion to a date time.
//
// Example: Assume you entered a value of -8 weeks and all other
// New() input parameters were zero value. The normalized TimeDto
// value would be converted and stored as:
//
//                   Years:  -1
//                  Months:  11
//                   Weeks:  4
//                WeekDays:  2
//                DateDays:  30
//                   Hours:  0
//                 Minutes:  0
//                 Seconds:  0
//            Milliseconds:  0
//            Microseconds:  0
//             Nanoseconds:  0
//
func (tDto TimeDto) New(years, months, weeks, days, hours, minutes,
	seconds, milliseconds, microseconds,
	nanoseconds int) (TimeDto, error) {

	ePrefix := "TimeDto.New(...) "

	t2Dto := TimeDto{}

	tDtoUtil := timeDtoUtility{}

	err := tDtoUtil.setTimeElements(
				&t2Dto,
				years,
				months,
				weeks,
				days,
				hours,
				minutes,
				seconds,
				milliseconds,
				microseconds,
				nanoseconds,
				ePrefix)

	if err != nil {
		return TimeDto{}, err
	}

	return t2Dto, nil
}

// NewTimeElements - Creates and returns a new TimeDto using basic
// time components as input parameters.
//
func (tDto TimeDto) NewTimeElements(
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	nanoseconds int) (TimeDto, error) {

	ePrefix := "TimeDto.NewTimeElements(...) "

	t2Dto := TimeDto{}

	tDtoUtil := timeDtoUtility{}

	err := tDtoUtil.setTimeElements(
						&t2Dto,
						years,
						months,
						0,
						days,
						hours,
						minutes,
						seconds,
						0,
						0,
						nanoseconds,
						ePrefix)

	if err != nil {
		return TimeDto{}, err
	}

	return t2Dto, nil
}

// NewAddTimeDtos - Creates a new TimeDto which adds the values of the two TimeDto's
// passed as input parameters. Date time math is performed using time zone 'UTC'.
//
// Note: This method is called with a pointer.
//
// Input Parameters
// ================
//
// t1Dto      TimeDto - The value of this TimeDto will be added to the second
//                      input parameter to create and return a summary TimeDto.
//
// t2Dto      TimeDto - The value of this TimeDto will be added to the first
//                      input parameter to create and return a summary TimeDto.
//
// Return Values
// =============
//
// TimeDto    - If successful, this method will return an instance of 'TimeDto'
//              populated with the total time element values calculated by adding
//              parameters 't1Dto' and 't2Dto'.
//
// Usage
// =====
//
// Method 'NewAddTimeDtos' must be called with a pointer. Example:
//  tDto := TimeDto{}
// tResultDto, err := tDto.NewAddTimeDtos(t1Dto, t2Dto)
//
func (tDto *TimeDto) NewAddTimeDtos(t1Dto, t2Dto TimeDto) (TimeDto, error) {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.NewAddTimeDtos(...) "

	tDtoUtil := timeDtoUtility{}
	
	tOutDto := tDtoUtil.copyOut(&t1Dto, ePrefix)
	
	err := tDtoUtil.addTimeDto(&tOutDto, &t2Dto, ePrefix)

	if err != nil {
		return TimeDto{}, err
	}

	return tOutDto, nil
}

// NewFromDateTime - Creates and returns a new TimeDto based on
// a date time (time.Time) input parameter.
//
func (tDto TimeDto) NewFromDateTime(dateTime time.Time) (TimeDto, error) {

	ePrefix := "TimeDto.NewFromDateTime() "

	if dateTime.IsZero() {
		return TimeDto{}, errors.New(ePrefix +
			"\nError: Input Parameter 'dateTime' has a ZERO Value!\n")
	}

	t2Dto := TimeDto{}

	tDtoUtil := timeDtoUtility{}

	err:=tDtoUtil.setFromDateTime(&t2Dto, dateTime, ePrefix)

	if err != nil {
		return TimeDto{}, err
	}

	return t2Dto, nil
}

// NewFromDateTzDto - Creates and returns a new TimeDto instance based on
// a DateTzDto input parameter.
func (tDto TimeDto) NewFromDateTzDto(dTzDto DateTzDto) (TimeDto, error) {

	ePrefix := "TimeDto.NewFromDateTzDto() "

	tDto2 := TimeDto{}

	tDtoUtil := timeDtoUtility{}

	err := tDtoUtil.setFromDateTzDto(&tDto2, dTzDto, ePrefix)

	if err != nil {
		return TimeDto{}, err
	}

	return tDto2, nil
}

// NormalizeTimeElements - Surveys the time elements of the current
// TimeDto and normalizes time values. Example: Hours between 0 and 23,
// Minutes between 0 and 59, Seconds between 0 and 59, etc.
//
func (tDto *TimeDto) NormalizeTimeElements() error {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.NormalizeTimeElements() "

	tDtoUtil := timeDtoUtility{}

	return tDtoUtil.normalizeTimeElements(tDto, ePrefix)
}

// NormalizeDays - Attempts to normalize days. This handles cases
// where the number of days is greater than the number of days
// in a month.
//
// If the number of days required normalization, the boolean
// return value is set to true.
//
func (tDto *TimeDto) NormalizeDays() (bool, error) {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.NormalizeDays() "

	tDtoUtil := timeDtoUtility{}

	return tDtoUtil.normalizeDays(tDto, ePrefix)
}

// SetTimeElements - Sets the value of date fields for the current TimeDto instance
// based on time element input parameters.
//
func (tDto *TimeDto) SetTimeElements(
	years,
	months,
	weeks,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) error {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.SetTimeElements(...) "

	tDtoUtil := timeDtoUtility{}

	return tDtoUtil.setTimeElements(
		tDto,
		years,
		months,
		weeks,
		days,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds,
		ePrefix)
}

// SetFromDateTime - Sets the current TimeDto instance to new
// data field values based on input parameter 'dateTime' (time.Time)
//
func (tDto *TimeDto) SetFromDateTime(dateTime time.Time) error {

	ePrefix := "TimeDto.SetFromDateTimeComponents() "

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Input Parameter 'dateTime' has a ZERO Value!\n")
	}

	tDtoUtil := timeDtoUtility{}

	return tDtoUtil.setFromDateTime(tDto, dateTime, ePrefix)
}

// SetFromDateTzDto - Sets the data field values of the current TimeDto
// instance based on a DateTzDto input parameter.
//
func (tDto *TimeDto) SetFromDateTzDto(dTzDto DateTzDto) error {

	tDto.lock.Lock()

	defer tDto.lock.Unlock()

	ePrefix := "TimeDto.SetFromDateTzDto() "

	tDtoUtil := timeDtoUtility{}

	return tDtoUtil.setFromDateTzDto(tDto, dTzDto, ePrefix)
}
