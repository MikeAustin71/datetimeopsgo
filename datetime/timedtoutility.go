package datetime

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type timeDtoUtility struct {
	lock   sync.Mutex
}

// AddTimeDto - Adds time to the current TimeDto. The amount of time added
// is provided by the input parameter 't2Dto' of type TimeDto.
//
// Date time math uses timezone UTC.
//
//  Input Parameters
//  ================
//
//  tDto      TimeDto  - These time element or component values are add to those
//                       encapsulated in 't2Dto' to generate a total. The total
//                       time element values then replace those in 'tDto' and
//                       comprise the new 'tDto' values.
//
//  t2Dto     TimeDto  - The amount of time to be added to parameter 'tDto'
//                       data fields and generate a total.
//
//

func (tDtoUtil *timeDtoUtility) addTimeDto(
	tDto *TimeDto,
	t2Dto *TimeDto,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.addTimeDto() "

	if tDto == nil {
		panic(ePrefix +
			"\nInput parameter 'tDto' is a nil pointer!")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	if t2Dto == nil {
		panic(ePrefix +
			"\nInput parameter 't2Dto' is a nil pointer!")
	}

	if t2Dto.lock == nil {
		t2Dto.lock = new(sync.Mutex)
	}

	tDtoUtil2 := timeDtoUtility{}

	return tDtoUtil2.setTimeElements( tDto,
		tDto.Years + t2Dto.Years,
		tDto.Months + t2Dto.Months,
		0,
		tDto.DateDays + t2Dto.DateDays,
		tDto.Hours + t2Dto.Hours,
		tDto.Minutes + t2Dto.Minutes,
		tDto.Seconds + t2Dto.Seconds,
		tDto.Milliseconds + t2Dto.Milliseconds,
		tDto.Microseconds + t2Dto.Microseconds,
		tDto.Nanoseconds + t2Dto.Nanoseconds,
		true, // Normalize data
		ePrefix)

}

// allocateWeeksAndDays - This method receives a total number of
// days and allocates those days to Weeks and WeekDays. The result
// is stored in the Weeks and WeekDays data fields of the specified
// TimeDto instance.
func (tDtoUtil *timeDtoUtility) allocateWeeksAndDays(
	tDto *TimeDto,
	totalDays int,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.allocateWeeksAndDays() "

	if tDto == nil {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter 'tDto' is nil!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	sign := 1

	if totalDays < 0 {
		sign = -1
		totalDays *= -1
	}

	tDto.Weeks = 0
	tDto.WeekDays = 0
	tDto.DateDays = totalDays

	if totalDays >= 7 {
		tDto.Weeks = totalDays / 7
		totalDays -= tDto.Weeks * 7
	}

	tDto.WeekDays = totalDays

	if sign == -1 {
		tDto.Weeks *= sign
		tDto.WeekDays *= sign
		tDto.DateDays *= sign
	}

	return nil
}

// allocateSeconds - Receives totalSeconds and proceeds to
// allocate Hours, Minutes and Seconds. The result is stored
// the Hours, Minutes and Seconds data fields of the specified
// TimeDto instance.
//
func (tDtoUtil *timeDtoUtility) allocateSeconds(
	tDto *TimeDto,
	totalSeconds int,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.allocateSeconds() "

	if tDto == nil {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter 'tDto' is nil!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	sign := 1

	if totalSeconds < 0 {
		sign = -1
		totalSeconds *= -1
	}

	tDto.Hours = 0
	tDto.Minutes = 0
	tDto.Seconds = 0

	if totalSeconds >= 3600 {
		tDto.Hours = totalSeconds / 3600
		totalSeconds -= tDto.Hours * 3600
	}

	if totalSeconds >= 60 {
		tDto.Minutes = totalSeconds / 60
		totalSeconds -= tDto.Minutes * 60
	}

	tDto.Seconds = totalSeconds

	if sign == -1 {

		tDto.Hours *= sign
		tDto.Minutes *= sign
		tDto.Seconds *= sign

	}

	return nil
}

// allocateTotalNanoseconds - Allocates total subMicrosecondNanoseconds to current
// TimeDto instance data fields: milliseconds, microseconds and
// subMicrosecondNanoseconds.
//
// In addition, this method calculates TimeDto.TotTimeNanoseconds which
// is the sum of hours, minutes, seconds, milliseconds, microseconds and
// subMicrosecondNanoseconds. Before calling this method, TimeDto Hours, Minutes and
// Seconds must be properly initialized.
//
func (tDtoUtil *timeDtoUtility) allocateTotalNanoseconds(
	tDto *TimeDto,
	totalNanoSeconds int,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.allocateTotalNanoseconds() "

	if tDto == nil {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter 'tDto' is nil!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	sign := 1

	if totalNanoSeconds < 0 {
		sign = -1
		totalNanoSeconds *= -1
	}

	tDto.Milliseconds = 0
	tDto.Microseconds = 0
	tDto.Nanoseconds = 0
	tDto.TotSubSecNanoseconds = totalNanoSeconds
	tDto.TotTimeNanoseconds = int64(totalNanoSeconds)

	if totalNanoSeconds >= int(time.Millisecond) {
		tDto.Milliseconds = totalNanoSeconds / int(time.Millisecond)
		totalNanoSeconds -= tDto.Milliseconds * int(time.Millisecond)
	}

	if totalNanoSeconds >= int(time.Microsecond) {
		tDto.Microseconds = totalNanoSeconds / int(time.Microsecond)
		totalNanoSeconds -= tDto.Microseconds * int(time.Microsecond)
	}

	tDto.Nanoseconds = totalNanoSeconds

	// calculate total time subMicrosecondNanoseconds
	tDto.TotTimeNanoseconds += int64(time.Hour) * int64(tDto.Hours)
	tDto.TotTimeNanoseconds += int64(time.Minute) * int64(tDto.Minutes)
	tDto.TotTimeNanoseconds += int64(time.Second) * int64(tDto.Seconds)

	if sign == -1 {

		tDto.Milliseconds *= sign
		tDto.Microseconds *= sign
		tDto.Nanoseconds *= sign
		tDto.TotSubSecNanoseconds *= sign
		tDto.TotTimeNanoseconds *= int64(sign)
	}

	return nil
}


// copyIn - Receives a TimeDto input parameter, 'tDto2'
// and proceeds to copy all 'tDto2' data fields into
// the current TimeDto data fields. When this method
// completes, 'tDto' will be equivalent to 'tDto2'.
//
func (tDtoUtil *timeDtoUtility) copyIn(
	tDto *TimeDto,
	t2Dto *TimeDto,
	ePrefix string) {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.copyIn() "

	if tDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tDto' is nil!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	if t2Dto == nil {
		panic(ePrefix +
			"\nError: Input parameter 't2Dto' is nil!\n")
	}

	if t2Dto.lock == nil {
		t2Dto.lock = new(sync.Mutex)
	}

	tDtoUtil2 := timeDtoUtility{}

	tDtoUtil2.empty(tDto, ePrefix)

	tDto.Years = t2Dto.Years
	tDto.Months = t2Dto.Months
	tDto.Weeks = t2Dto.Weeks
	tDto.WeekDays = t2Dto.WeekDays
	tDto.DateDays = t2Dto.DateDays
	tDto.Hours = t2Dto.Hours
	tDto.Minutes = t2Dto.Minutes
	tDto.Seconds = t2Dto.Seconds
	tDto.Milliseconds = t2Dto.Milliseconds
	tDto.Microseconds = t2Dto.Microseconds
	tDto.Nanoseconds = t2Dto.Nanoseconds
	tDto.TotSubSecNanoseconds = t2Dto.TotSubSecNanoseconds
	tDto.TotTimeNanoseconds = t2Dto.TotTimeNanoseconds

	return
}


// copyOut - Creates a new TimeDto instance
// which precisely duplicates the current TimeDto
// instance, and returns it to the calling function.
func (tDtoUtil *timeDtoUtility) copyOut(
	tDto *TimeDto,
	ePrefix string) TimeDto {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.copyOut() "

	if tDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tDto' is nil!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	t2Dto := TimeDto{}

	t2Dto.Years = tDto.Years
	t2Dto.Months = tDto.Months
	t2Dto.Weeks = tDto.Weeks
	t2Dto.WeekDays = tDto.WeekDays
	t2Dto.DateDays = tDto.DateDays
	t2Dto.Hours = tDto.Hours
	t2Dto.Minutes = tDto.Minutes
	t2Dto.Seconds = tDto.Seconds
	t2Dto.Milliseconds = tDto.Milliseconds
	t2Dto.Microseconds = tDto.Microseconds
	t2Dto.Nanoseconds = tDto.Nanoseconds
	t2Dto.TotSubSecNanoseconds = tDto.TotSubSecNanoseconds
	t2Dto.TotTimeNanoseconds = tDto.TotTimeNanoseconds
	t2Dto.lock = new(sync.Mutex)

	return t2Dto
}

// computeDuration - Receives a TimeDto instance and converts
// the constituent time components into a single time duration
// value.
//
// Note that duration is computed using standard 24-hour days
// and does not take into account time zones and daylight savings
// time.
//
func (tDtoUtil *timeDtoUtility) computeDuration(
	tDto *TimeDto,
	ePrefix string) (time.Duration, error) {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.computeDuration() "

	if tDto == nil {
		return time.Duration(0),
		&InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDto",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDto' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	return time.Duration(tDto.TotTimeNanoseconds), nil
}

// convertToAbsoluteValues - Converts time components
// (Years, months, weeks days, hours, seconds, etc.)
// to absolute values.
//
// In other words, after this method completes, all
// time component values will be positive.
//
func (tDtoUtil *timeDtoUtility) convertToAbsoluteValues(
	tDto *TimeDto,
	ePrefix string) {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.convertToAbsoluteValues() "

	if tDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tDto' is nil!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	if tDto.Years < 0 {
		tDto.Years *= -1
	}

	if tDto.Months < 0 {
		tDto.Months *= -1
	}

	if tDto.Weeks < 0 {
		tDto.Weeks *= -1
	}

	if tDto.WeekDays < 0 {
		tDto.WeekDays *= -1
	}

	if tDto.DateDays < 0 {
		tDto.DateDays *= -1
	}

	if tDto.Hours < 0 {
		tDto.Hours *= -1
	}

	if tDto.Minutes < 0 {
		tDto.Minutes *= -1
	}

	if tDto.Seconds < 0 {
		tDto.Seconds *= -1
	}

	if tDto.Milliseconds < 0 {
		tDto.Milliseconds *= -1
	}

	if tDto.Milliseconds < 0 {
		tDto.Milliseconds *= -1
	}

	if tDto.Microseconds < 0 {
		tDto.Microseconds *= -1
	}

	if tDto.Nanoseconds < 0 {
		tDto.Nanoseconds *= -1
	}

	if tDto.TotSubSecNanoseconds < 0 {
		tDto.TotSubSecNanoseconds *= -1
	}

	if tDto.TotTimeNanoseconds < 0 {
		tDto.TotTimeNanoseconds *= -1
	}

	return
}

// convertToNegativeValues - Multiplies time component
// values by -1
//
func (tDtoUtil *timeDtoUtility) convertToNegativeValues(
	tDto *TimeDto,
	ePrefix string) {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.convertToNegativeValues() "

	if tDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tDto' is nil!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	tDtoUtil2 := timeDtoUtility{}

	tDtoUtil2.convertToAbsoluteValues(tDto, ePrefix)

	tDto.Years *= -1
	tDto.Months *= -1
	tDto.Weeks *= -1
	tDto.WeekDays *= -1
	tDto.DateDays *= -1
	tDto.Hours *= -1
	tDto.Minutes *= -1
	tDto.Seconds *= -1
	tDto.Milliseconds *= -1
	tDto.Microseconds *= -1
	tDto.Nanoseconds *= -1
	tDto.TotSubSecNanoseconds *= -1
	tDto.TotTimeNanoseconds *= -1
}

// empty - returns all TimeDto data fields to their
// uninitialized or zero state.
func (tDtoUtil *timeDtoUtility) empty(
	tDto *TimeDto,
	ePrefix string) {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.empty() "

	if tDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tDto' is nil!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	tDto.Years = 0
	tDto.Months = 0
	tDto.Weeks = 0
	tDto.WeekDays = 0
	tDto.DateDays = 0
	tDto.Hours = 0
	tDto.Minutes = 0
	tDto.Seconds = 0
	tDto.Milliseconds = 0
	tDto.Microseconds = 0
	tDto.Nanoseconds = 0
	tDto.TotSubSecNanoseconds = 0
	tDto.TotTimeNanoseconds = 0

	return
}

// Equal - Compares the data fields of TimeDto input parameters, 'tDto'
// and 'tDto2'. If the data fields are equal in all respects,
// this method returns 'true'.
//
func (tDtoUtil *timeDtoUtility) equalTimeDtos(
	tDto *TimeDto,
	t2Dto *TimeDto,
	ePrefix string) bool {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.equalTimeDtos() "

	if tDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tDto' is a 'nil' pointer!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	if t2Dto == nil {
		panic(ePrefix +
			"\nError: Input parameter 't2Dto' is a 'nil' pointer!\n")
	}

	if t2Dto.lock == nil {
		t2Dto.lock = new(sync.Mutex)
	}

	if tDto.Years != t2Dto.Years ||
		tDto.Months != t2Dto.Months ||
		tDto.Weeks != t2Dto.Weeks ||
		tDto.WeekDays != t2Dto.WeekDays ||
		tDto.DateDays != t2Dto.DateDays ||
		tDto.Hours != t2Dto.Hours ||
		tDto.Minutes != t2Dto.Minutes ||
		tDto.Seconds != t2Dto.Seconds ||
		tDto.Milliseconds != t2Dto.Milliseconds ||
		tDto.Microseconds != t2Dto.Microseconds ||
		tDto.Nanoseconds != t2Dto.Nanoseconds ||
		tDto.TotSubSecNanoseconds != t2Dto.TotSubSecNanoseconds ||
		tDto.TotTimeNanoseconds != t2Dto.TotTimeNanoseconds {

		return false
	}

	return true
}

// isEmpty - Returns 'true' if all data fields in the current
// TimeDto instance are equal to zero or equal to their
// uninitialized values.
func (tDtoUtil *timeDtoUtility) isEmpty(
	tDto *TimeDto,
	ePrefix string) bool {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.isEmpty() "

	if tDto == nil {
		panic(ePrefix +
			"\nError: Input parameter 'tDto' pointer is nil!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	if tDto.Years == 0 &&
		tDto.Months == 0 &&
		tDto.Weeks == 0 &&
		tDto.WeekDays == 0 &&
		tDto.DateDays == 0 &&
		tDto.Hours == 0 &&
		tDto.Minutes == 0 &&
		tDto.Seconds == 0 &&
		tDto.Milliseconds == 0 &&
		tDto.Microseconds == 0 &&
		tDto.Nanoseconds == 0 &&
		tDto.TotSubSecNanoseconds == 0 &&
		tDto.TotTimeNanoseconds == 0 {
		return true
	}

	return false
}

// isValidDateTimeDto - Returns an error if the current tDto instance is invalid.
// Otherwise, if successful, this method returns 'nil'.
func (tDtoUtil *timeDtoUtility) isValidDateTimeDto(
	tDto *TimeDto,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.isValidDateTimeDto() "

	if tDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tDto' is a 'nil' pointer!")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	if tDto.Months < 1 || tDto.Months > 12 {
		return fmt.Errorf(ePrefix+
			"\nError: Months value is INVALID!\n" +
			"tDto.Months='%v'\n", tDto.Months)
	}

	if tDto.Weeks < 0 || tDto.Weeks > 4 {
		return fmt.Errorf(ePrefix +
			"\nError: Weeks value is INVALID!\n" +
			"tDto.Weeks='%v'\n", tDto.Weeks)
	}

	if tDto.WeekDays < 0 || tDto.WeekDays > 6 {
		return fmt.Errorf( ePrefix +
			"\nError: WeekDays value is INVALID!\n" +
			"tDto.WeekDays='%v'\n", tDto.WeekDays)
	}

	if tDto.DateDays < 0 || tDto.DateDays > 31 {
		return fmt.Errorf(ePrefix +
			"\nError: Total WeekDays value is INVALID!\n" +
			"tDto.DateDays='%v'\n", tDto.DateDays)
	}

	if tDto.Hours < 0 || tDto.Hours > 24 {
		return fmt.Errorf(ePrefix +
			"\nError: Hours value is INVALID!\n" +
			"tDto.Hours='%v'\n", tDto.Hours)
	}

	if tDto.Minutes < 0 || tDto.Minutes > 59 {
		return fmt.Errorf( ePrefix +
			"\nError: Minutes value is INVALID!\n" +
			"tDto.Minutes='%v'\n", tDto.Minutes)
	}

	if tDto.Seconds < 0 || tDto.Seconds > 59 {
		return fmt.Errorf(ePrefix +
			"\nError: Seconds value is INVALID!\n" +
			"tDto.Seconds='%v'\n", tDto.Seconds)
	}

	if tDto.Milliseconds < 0 || tDto.Milliseconds > int(MilliSecondsPerSecond-1) {
		return fmt.Errorf( ePrefix +
			"\nError: Milliseconds value is INVALID!\n" +
			"tDto.Milliseconds='%v'\n", tDto.Milliseconds)
	}

	if tDto.Microseconds < 0 ||
			tDto.Microseconds > int(MicroSecondsPerMilliSecond-1) {
		return fmt.Errorf( ePrefix +
			"\nError: Microseconds value is INVALID!\n" +
			"tDto.Microseconds='%v'\n", tDto.Microseconds)
	}

	if tDto.Nanoseconds < 0 ||
			tDto.Nanoseconds > int(NanoSecondsPerMicroSecond-1) {
		return fmt.Errorf(ePrefix +
			"\nError: Nanoseconds value is INVALID!\n" +
			"tDto.Nanoseconds='%v'\n", tDto.Nanoseconds)
	}

	if tDto.TotSubSecNanoseconds < 0 ||
			tDto.TotSubSecNanoseconds > int(SecondNanoseconds-1) {
		return fmt.Errorf(ePrefix +
			"\nError: Total Nanoseconds value is INVALID!\n" +
			"tDto.TotSubSecNanoseconds='%v'\n", tDto.TotSubSecNanoseconds)
	}

	if tDto.TotTimeNanoseconds < 0 {
		return fmt.Errorf( ePrefix +
			"\nError: Total Time Nanoseconds value is INVALID!\n" +
			"tDto.TotTimeNanoseconds='%v'\n", tDto.TotTimeNanoseconds)
	}

	return nil
}


// NormalizeDays - Attempts to normalize days. This handles cases
// where the number of days is greater than the number of days
// in a month.
//
// If the number of days did require normalization, the boolean
// return value is set to 'true'.
//
func (tDtoUtil *timeDtoUtility) normalizeDays(
	tDto *TimeDto,
	ePrefix string) (bool, error) {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.normalizeDays() "

	if tDto == nil {
		return false, errors.New(ePrefix +
			"\nError: Input parameter 'tDto' is a 'nil' pointer!")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	if tDto.Years == 0 && tDto.Months == 0 {
		return false, nil
	}

	if tDto.Months < 0 {
		return false, nil
	}

	tDtoUtil2 := timeDtoUtility{}

	t2Dto := tDtoUtil2.copyOut(tDto, ePrefix)

	locUTC, err := time.LoadLocation(TZones.UTC())

	if err != nil {
		return false,
			fmt.Errorf(ePrefix+
				"\nError returned by time.LoadLocation(TZones.UTC()).\n" +
				"Error='%v'\n",
				err.Error())
	}

	weekDays := (t2Dto.Weeks * 7) + t2Dto.WeekDays
	dateDays := t2Dto.DateDays

	if dateDays == weekDays {
		weekDays = 0
	} else if dateDays == 0 && weekDays != 0 {
		dateDays = weekDays
		weekDays = 0
	} else if weekDays != 0 && dateDays != 0 &&
		weekDays != dateDays {
		dateDays += weekDays
		weekDays = 0
	}

	// Date Days are already normalized!
	if dateDays < 29 {
		return false, nil
	}

	years := t2Dto.Years

	months := t2Dto.Months

	if months == 0 {
		months = 1
	}

	dt1 := time.Date(years, time.Month(months), 0, 0, 0, 0, 0, locUTC)

	dur := int64(dateDays) * DayNanoSeconds
	dur += int64(t2Dto.Hours) * HourNanoSeconds
	dur += int64(t2Dto.Minutes) * MinuteNanoSeconds
	dur += int64(t2Dto.Seconds) * SecondNanoseconds
	dur += int64(t2Dto.Milliseconds) * MilliSecondNanoseconds
	dur += int64(t2Dto.Microseconds) * MicroSecondNanoseconds
	dur += int64(t2Dto.Nanoseconds)

	dateTime := dt1.Add(time.Duration(dur))

	tDtoUtil2.empty(&t2Dto, ePrefix)

	t2Dto.Years = dateTime.Year()
	t2Dto.Months = int(dateTime.Month())

	err = tDtoUtil2.allocateWeeksAndDays(&t2Dto, dateTime.Day(), ePrefix)

	if err != nil {
		return false, err
	}

	totSeconds := dateTime.Hour() * 3600
	totSeconds += dateTime.Minute() * 60
	totSeconds += dateTime.Second()

	err = tDtoUtil2.allocateSeconds(&t2Dto, totSeconds, ePrefix)

	if err != nil {
		return false, err
	}

	err = tDtoUtil2.allocateTotalNanoseconds(&t2Dto, dateTime.Nanosecond(), ePrefix)

	if err != nil {
		return false, fmt.Errorf(ePrefix+
			"\nError returned by t2Dto.allocateTotalNanoseconds(dateTime.Nanosecond()).\n"+
			"Error='%v'\n", err.Error())
	}

	err = tDtoUtil2.isValidDateTimeDto(&t2Dto, ePrefix)

	if err != nil {
		return false, err
	}

	tDtoUtil2.copyIn(tDto, &t2Dto, ePrefix)

	return true, nil
}

// normalizeTimeElements - Surveys the time elements of the current
// TimeDto and normalizes time values. Example: Hours between 0 and 23,
// Minutes between 0 and 59, Seconds between 0 and 59, etc.
//
func (tDtoUtil *timeDtoUtility) normalizeTimeElements(
	tDto *TimeDto,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.normalizeTimeElements() "

	if tDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tDto' is a 'nil' pointer!")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	carry := tDto.Nanoseconds / 1000
	tDto.Nanoseconds -= carry * 1000

	tDto.Microseconds += carry
	carry = tDto.Microseconds / 1000
	tDto.Microseconds -= carry * 1000

	tDto.Milliseconds += carry
	carry = tDto.Milliseconds / 1000
	tDto.Milliseconds -= carry * 1000

	tDto.Seconds += carry
	carry = tDto.Seconds / 60
	tDto.Seconds -= carry * 60

	tDto.Minutes += carry
	carry = tDto.Minutes / 60
	tDto.Minutes -= carry * 60

	tDto.Hours += carry
	carry = tDto.Hours / 24
	tDto.Hours -= carry * 24

	weekDays := (tDto.Weeks * 7) + tDto.WeekDays
	dateDays := tDto.DateDays

	if dateDays == weekDays {
		weekDays = 0
	} else if dateDays == 0 && weekDays != 0 {
		dateDays = weekDays
	} else if weekDays != 0 && dateDays != 0 &&
		weekDays != dateDays {
		dateDays += weekDays
	}

	tDto.DateDays = dateDays

	tDto.DateDays += carry

	carry = tDto.Months / 12
	tDto.Months -= carry * 12

	tDto.Years += carry

	tDtoUtil2 := timeDtoUtility{}

	err := tDtoUtil2.allocateWeeksAndDays(tDto, tDto.DateDays, ePrefix)

	if err != nil {
		return err
	}

	totSeconds := tDto.Hours * 3600
	totSeconds += tDto.Minutes * 60
	totSeconds += tDto.Seconds

	err = tDtoUtil2.allocateSeconds(tDto, totSeconds, ePrefix)

	if err != nil {
		return err
	}

	totSubNanoSecs := int(int64(tDto.Milliseconds) * MilliSecondNanoseconds)
	totSubNanoSecs += int(int64(tDto.Microseconds) * MicroSecondNanoseconds)
	totSubNanoSecs += tDto.Nanoseconds

	err = tDtoUtil2.allocateTotalNanoseconds(tDto, totSubNanoSecs, ePrefix)

	return err
}

func (tDtoUtil *timeDtoUtility) setFromDuration(
	tDto *TimeDto,
	duration time.Duration,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.setFromDateTime() "

	if tDto == nil {
		return errors.New(ePrefix +
			"Error: Input parameter 'tDto' is a 'nil' pointer!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	tDtoUtil2 := timeDtoUtility{}

	tDtoUtil2.empty(tDto, ePrefix)

	nanoSecs := int64(duration)


	if nanoSecs == 0 {
		return nil
	}

	tDto.TotTimeNanoseconds = nanoSecs

	signVal := int64(1)

	if nanoSecs < signVal {
		signVal = -1
		nanoSecs = nanoSecs * signVal
	}

	temp := int64(0)

	if nanoSecs >= HourNanoSeconds {
		temp = nanoSecs/HourNanoSeconds

		nanoSecs -= temp * HourNanoSeconds

		tDto.Hours = int(temp * signVal)

	}

	if nanoSecs >= MinuteNanoSeconds {

		temp = nanoSecs/MinuteNanoSeconds

		nanoSecs -= temp * MinuteNanoSeconds

		tDto.Minutes = int(temp * signVal)
	}

	if nanoSecs >= SecondNanoseconds {

		temp = nanoSecs/SecondNanoseconds

		nanoSecs -= temp * SecondNanoseconds

		tDto.Seconds = int(temp * signVal)
	}

	tDto.TotSubSecNanoseconds = int(nanoSecs * signVal)

	if nanoSecs >= MilliSecondNanoseconds {

		temp = nanoSecs/MilliSecondNanoseconds

		nanoSecs -= temp * MilliSecondNanoseconds

		tDto.Milliseconds = int(temp * signVal)
	}

	if nanoSecs >= MicroSecondNanoseconds {

		temp = nanoSecs/MicroSecondNanoseconds

		nanoSecs -= temp * MicroSecondNanoseconds

		tDto.Microseconds = int(temp * signVal)
	}

	tDto.Nanoseconds = int(nanoSecs * signVal)

	return nil
}

// SetFromDateTime - Populates the specified 'TimeDto' instance with new
// data field values based on input parameter 'dateTime' (time.Time)
//
func (tDtoUtil *timeDtoUtility) setFromDateTime(
	tDto *TimeDto,
	dateTime time.Time,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.setFromDateTime() "

	if tDto == nil {
		return errors.New(ePrefix +
			"Error: Input parameter 'tDto' is a 'nil' pointer!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}


	tDtoUtil2 := timeDtoUtility{}

	tDtoUtil2.empty(tDto, ePrefix)

	tDto.Years = dateTime.Year()
	tDto.Months = int(dateTime.Month())

	err := tDtoUtil2.allocateWeeksAndDays(tDto, dateTime.Day(), ePrefix)

	if err != nil {
		return err
	}

	tDto.Hours = dateTime.Hour()
	tDto.Minutes = dateTime.Minute()
	tDto.Seconds = dateTime.Second()

	err = tDtoUtil2.allocateTotalNanoseconds(tDto, dateTime.Nanosecond(), ePrefix)

	if err != nil {
		return err
	}

	err = tDtoUtil2.isValidDateTimeDto(tDto, ePrefix)

	if err != nil {
		return err
	}

	return nil
}


// SetFromDateTzDto - Sets the data field values of the current TimeDto
// instance based on a DateTzDto input parameter.
//
func (tDtoUtil *timeDtoUtility) setFromDateTzDto(
	tDto *TimeDto,
	dTzDto DateTzDto,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.setFromDateTzDto() "

	if tDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tDto' is a 'nil' pointer!")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	dTzUtil2 := dateTzDtoUtility{}

	if dTzUtil2.isEmptyDateTzDto(&dTzDto) {
		return errors.New(ePrefix + "Error: Input parameter 'dTzDto' (DateTzDto) is EMPTY!")
	}

	err := dTzUtil2.isValidDateTzDto(&dTzDto, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter 'dTzDto' (DateTzDto) is INVALID!\n" +
			"Error='%v'\n", err.Error())
	}

	tDtoUtil2 := timeDtoUtility{}

	tDtoUtil2.empty(tDto, ePrefix)

	tDto.Years = dTzDto.GetDateTimeValue().Year()
	tDto.Months = int(dTzDto.GetDateTimeValue().Month())

	err = tDtoUtil2.allocateWeeksAndDays(tDto, dTzDto.GetDateTimeValue().Day(), ePrefix)

	if err != nil {
		return err
	}

	tDto.Hours = dTzDto.GetDateTimeValue().Hour()
	tDto.Minutes = dTzDto.GetDateTimeValue().Minute()
	tDto.Seconds = dTzDto.GetDateTimeValue().Second()

	err = tDtoUtil2.allocateTotalNanoseconds(tDto, dTzDto.GetDateTimeValue().Nanosecond(), ePrefix)

	if err != nil {
		return err
	}

	err = tDtoUtil2.isValidDateTimeDto(tDto, ePrefix)

	return err
}

// setTimeElements - Sets the value of date fields for the current TimeDto instance
// based on time element input parameters.
//
func (tDtoUtil *timeDtoUtility) setTimeElements(
					tDto *TimeDto,
					years,
					months,
					weeks,
					days,
					hours,
					minutes,
					seconds,
					milliseconds,
					microseconds,
					nanoseconds int,
					normalizeData bool,
					ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.setTimeElements(...) "

	if tDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tDto' is a 'nil' pointer!\n")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	if years == 0 &&
		months == 0 &&
		weeks == 0 &&
		days == 0 &&
		hours == 0 &&
		minutes == 0 &&
		seconds == 0 &&
		milliseconds == 0 &&
		microseconds == 0 &&
		nanoseconds == 0 {

		return fmt.Errorf(ePrefix +
			"\nError: All input parameters (years, months, weeks, days etc.) are ZERO XValue!\n")
	}

	t1Dto := TimeDto{}

	t1Dto.Years = years
	t1Dto.Months = months
	t1Dto.DateDays = (weeks * 7) + days

	t1Dto.Hours = hours
	t1Dto.Minutes = minutes
	t1Dto.Seconds = seconds
	t1Dto.Milliseconds = milliseconds
	t1Dto.Microseconds = microseconds
	t1Dto.Nanoseconds = nanoseconds
	t1Dto.lock = new(sync.Mutex)

	tDtoUtil2 := timeDtoUtility{}
	var err error

	if normalizeData {

		err = tDtoUtil2.normalizeTimeElements(&t1Dto, ePrefix)

		if err != nil {
			return err
		}

		_, err = tDtoUtil2.normalizeDays(&t1Dto, ePrefix)

		if err != nil {
			return fmt.Errorf(ePrefix+"Error returned by err := t1Dto.NormalizeDays() "+
				"Error='%v'", err.Error())
		}
	}

	tDtoUtil2.copyIn(tDto, &t1Dto, ePrefix)

	return nil
}

// setZeroTimeDto - Sets the incoming TimeDto instance
// to zero values.
//
func (tDtoUtil *timeDtoUtility) setZeroTimeDto(
	tDto *TimeDto,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.setZeroTimeDto() "

	if tDto == nil {
		return errors.New(ePrefix +
			"\nInput parameter 'tDto' is a nil pointer!")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	tDtoUtil2 := timeDtoUtility{}

	tDtoUtil2.empty(tDto, ePrefix)

	// 0001-01-01 00:00:00.000000000 +0000 UTC
	tDto.Years = 1
	tDto.Months = 1
	tDto.DateDays = 1

	return nil
}

// AddTimeDto - Adds time to the current TimeDto. The amount of time added
// is provided by the input parameter 't2Dto' of type TimeDto.
//
// Date time math uses timezone UTC.
//
//  Input Parameters
//  ================
//
//  tDto      TimeDto  - These time element or component values are add to those
//                       encapsulated in 't2Dto' to generate a total. The total
//                       time element values then replace those in 'tDto' and
//                       comprise the new 'tDto' values.
//
//  t2Dto     TimeDto  - The amount of time to be added to parameter 'tDto'
//                       data fields and generate a total.
//
//

func (tDtoUtil *timeDtoUtility) subTimeDto(
	tDto *TimeDto,
	t2Dto *TimeDto,
	ePrefix string) error {

	tDtoUtil.lock.Lock()

	defer tDtoUtil.lock.Unlock()

	ePrefix += "timeDtoUtility.subTimeDto() "

	if tDto == nil {
		panic(ePrefix +
			"\nInput parameter 'tDto' is a nil pointer!")
	}

	if tDto.lock == nil {
		tDto.lock = new(sync.Mutex)
	}

	if t2Dto == nil {
		panic(ePrefix +
			"\nInput parameter 't2Dto' is a nil pointer!")
	}

	if t2Dto.lock == nil {
		t2Dto.lock = new(sync.Mutex)
	}

	tDtoUtil2 := timeDtoUtility{}

	return tDtoUtil2.setTimeElements( tDto,
		tDto.Years - t2Dto.Years,
		tDto.Months - t2Dto.Months,
		0,
		tDto.DateDays - t2Dto.DateDays,
		tDto.Hours - t2Dto.Hours,
		tDto.Minutes - t2Dto.Minutes,
		tDto.Seconds - t2Dto.Seconds,
		tDto.Milliseconds - t2Dto.Milliseconds,
		tDto.Microseconds - t2Dto.Microseconds,
		tDto.Nanoseconds - t2Dto.Nanoseconds,
		true, // Normalize data
		ePrefix)

}

