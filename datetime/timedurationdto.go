package datetime

import (
	"fmt"
	"sync"
	"time"
)

// TimeDurationDto
//
// Overview and Usage
//
// The 'TimeDurationDto' Type is used to calculate, store and transmit
// time duration information. It is therefore designed to work with
// incremental time or time duration.
//
// The Type includes a starting date time, ending date time, a time
// duration value specifying the time span between starting and ending
// date time, and a breakdown of time duration by a series of time
// components including years, months, weeks, days, hours, minutes,
// seconds, milliseconds, microseconds and nanoseconds.
//
// Dependencies
//
// Starting and ending date times are stored as 'DateTzDto' types. The
// 'DateTzDto' type may be found in source code file:
//   MikeAustin71\datetimeopsgo\datetime\datetzdto.go
//
// Source Code Location
//
// This source file is located in source code repository:
//   https://github.com/MikeAustin71/datetimeopsgo.git
//
// The location of this source file is:
//   MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
type TimeDurationDto struct {
	startDateTimeTz DateTzDto     // Starting Date Time with Time Zone info
	endDateTimeTz   DateTzDto     // Ending Date Time with Time Zone info
	timeDuration    time.Duration // Elapsed time or duration between starting and ending date time
	timeDurCalcType TDurCalcType  // The calculation Type. This controls the allocation of time
	//                               duration over years, months, weeks, days and hours.
	timeMathCalcMode     TimeMathCalcMode // The Time Math algorithm used to calculate time duration
	Years                int64            // Number of Years
	YearsNanosecs        int64            // Number of Years in Nanoseconds
	Months               int64            // Number of Months
	MonthsNanosecs       int64            // Number of Months in Nanoseconds
	Weeks                int64            // Number of Weeks: Date Days / 7
	WeeksNanosecs        int64            // Number of Weeks in Nanoseconds
	WeekDays             int64            // WeekDays = DateDays - (Weeks * 7)
	WeekDaysNanosecs     int64            // Equivalent WeekDays in NanoSeconds
	DateDays             int64            // Day Number in Month (1-31)
	DateDaysNanosecs     int64            // DateDays in equivalent nanoseconds
	Hours                int64            // Number of Hours
	HoursNanosecs        int64            // Number of Hours in Nanoseconds
	Minutes              int64            // Number of Minutes
	MinutesNanosecs      int64            // Number of Minutes in Nanoseconds
	Seconds              int64            // Number of Seconds
	SecondsNanosecs      int64            // Number of Seconds in Nanoseconds
	Milliseconds         int64            // Number of Milliseconds
	MillisecondsNanosecs int64            // Number of Milliseconds in Nanoseconds
	Microseconds         int64            // Number of Microseconds
	MicrosecondsNanosecs int64            // Number of Microseconds in Nanoseconds
	Nanoseconds          int64            // Number of Nanoseconds (Remainder after Milliseconds & Microseconds)
	TotSubSecNanoseconds int64            // Equivalent Nanoseconds for Milliseconds + Microseconds + Nanoseconds
	TotDateNanoseconds   int64            // Equal to Years + Months + DateDays in equivalent nanoseconds.
	TotTimeNanoseconds   int64            // Equal to Hours + Seconds + Milliseconds + Microseconds + Nanoseconds in
	// in equivalent nanoseconds

	lock *sync.Mutex // Used to enforce thread safe operations
}

// CopyIn - Receives a TimeDurationDto as an input parameters
// and proceeds to set all data fields of the current TimeDurationDto
// equal to the incoming TimeDurationDto.
//
// When this method completes, the current TimeDurationDto will
// equal in all respects to the incoming TimeDurationDto.
//
// __________________________________________________________________________
//
// Return Values:
//
//  -- NONE --
//
func (tDur *TimeDurationDto) CopyIn(t2Dur TimeDurationDto) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.copyIn(
		tDur,
		&t2Dur,
		"TimeDurationDto.CopyIn() ")
}

// copyOut - Returns a deep copy of the current
// TimeDurationDto instance.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
func (tDur *TimeDurationDto) CopyOut() TimeDurationDto {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.copyOut(
		tDur,
		"TimeDurationDto.CopyOut() ")
}

// Empty - Resets all current TimeDurationDto data
// fields to their zero or uninitialized values.
//
// __________________________________________________________________________
//
// Return Values:
//
//   -- NONE --
//
func (tDur *TimeDurationDto) Empty() {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.empty(
		tDur,
		"TimeDurationDto.Empty() ")
}

// Equal - Compares two TimeDurationDto instances to determine
// if they are equivalent.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  t2Dur TimeDurationDto - A TimeDurationDto object which will be
//                          compared to the current TimeDurationDto
//                          instance to determine if the two are
//                          equivalent.
// __________________________________________________________________________
//
// Return Value:
//
//  bool - If 'true' it signals that all relevant data fields in
//         'tDur' and 'tDur2' are equivalent.
//
func (tDur *TimeDurationDto) Equal(t2Dur TimeDurationDto) bool {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.Equal() "

	tDurDtoUtil := timeDurationDtoUtility{}

	isEqual, err := tDurDtoUtil.equal(tDur, &t2Dur, ePrefix)

	if err != nil {
		return false
	}

	return isEqual
}

// IsEmpty() - Returns 'true' if the current TimeDurationDto
// instance is uninitialized and consists entirely of zero values.
// __________________________________________________________________________
//
// Return Value:
//
//  bool - If 'true' it signals that all relevant data fields in
//         in the current 'TimeDurationDto' instance (tDur) are
//         empty, or set to their zero values.
//
func (tDur *TimeDurationDto) IsEmpty() bool {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.IsEmpty() "

	tDurDtoUtil := timeDurationDtoUtility{}

	isEmpty, err := tDurDtoUtil.isEmpty(tDur, ePrefix)

	if err != nil {
		return false
	}

	return isEmpty
}

// IsValid - Returns an error value signaling whether
// the current TimeDurationDto data fields are valid.
//
// __________________________________________________________________________
//
// Return Value:
//
//  error - If the current 'TimeDurationDto' instance is valid and populated,
//          the returned 'error' is set to nil. If the instance is invalid,
//          the returned 'error' object contains an appropriate error message.
//
func (tDur *TimeDurationDto) IsValid() error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.IsValid() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.isValid(tDur, ePrefix)
}

// GetCumDaysCalcDto - Returns a new TimeDurationDto which re-calculates
// the values of the current TimeDurationDto and stores them in a
// 'cumulative days' format. This format always shows zero years and
// zero months. It consolidates years, months and days and presents them
// as cumulative days.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a new
//       valid and fully populated 'TimeDurationDto' instance will
//       be returned.
//
//       The new, returned TimeDurationDto instance will have a
//       calculation type of 'TDurCalcType(0).CumDays()'
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Cumulative Days Format:
//
//  97-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetCumDaysCalcDto() (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumDaysCalcDto() "

	if int64(tDur.timeDuration) == 0 {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix + "Error: Time Duration is equal to zero!")
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumDays(),
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// GetCumDaysTimeStr - Returns duration formatted as days, hours,
// minutes, seconds, milliseconds, microseconds, and nanoseconds.
// Years, months and weeks are always excluded and included in
// cumulative 'days'.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) formatted as cumulative days. See Example String below.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Return String:
//
//  97-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetCumDaysTimeStr() (string, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumDaysTimeStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumDays(),
		ePrefix)

	if err != nil {
		return "", err
	}

	str := ""

	str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetCumHoursCalcDto - Returns a new TimeDurationDto. The time
// values of the current TimeDurationDto are recalculated for
// 'cumulative hours'.
//
// This means that years, months and days are ignored and set to
// a zero value.  Instead, years, months, days and hours are
// consolidated and stored as cumulative hours.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a new
//       valid and fully populated 'TimeDurationDto' instance will
//       be returned.
//
//       The new, returned TimeDurationDto instance will have a
//       calculation type of 'TDurCalcType(0).CumHours()'
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Cumulative Hours Format:
//
//  152-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetCumHoursCalcDto() (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumHoursCalcDto() "

	if int64(tDur.timeDuration) == 0 {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"\nError: Time Duration is ZERO value!\n")
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumHours(),
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// GetCumHoursTimeStr - Returns duration formatted as hours,
// minutes, seconds, milliseconds, microseconds, nanoseconds.
//
// Years, months and days are ignored and set to a zero value.
// Instead, years, months, days and hours are consolidated and
// presented as cumulative hours.
//
// This method will NOT modify the internal data fields of the
// current TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) formatted as cumulative hours. See Example String below.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Return String:
//
//  152-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetCumHoursTimeStr() (string, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumHoursTimeStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumHours(),
		ePrefix)

	if err != nil {
		return "", err
	}

	str := ""

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetCumMinutesCalcDto - Returns a new TimeDurationDto calculated and configured
// for cumulative minutes. This means that years, months, days and hours are
// set to a zero value.
//
// Instead, years, months, days, hours and minutes are all consolidated
// and presented as cumulative minutes.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a new
//       valid and fully populated 'TimeDurationDto' instance will
//       be returned.
//
//       The new, returned TimeDurationDto instance will have a
//       calculation type of 'TDurCalcType(0).CumMinutes()'
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Cumulative Minutes Format:
//
//   "527-Minutes 37-Seconds 18-Milliseconds 256-Microseconds 852-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumMinutesCalcDto() (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumMinutesCalcDto() "

	if int64(tDur.timeDuration) == 0 {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"\nError: Time Duration is ZERO!\n")
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumMinutes(),
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// GetCumMinutesTimeStr - Returns duration formatted as cumulative
// minutes. This format ignores years, months, days and hours.
//
// Instead, years, months, days, hours and minutes are all consolidated
// and presented as minutes.
//
// This method will NOT modify the internal data fields of the
// current TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) formatted as cumulative minutes. See Example String below.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Return String:
//
//  "527-Minutes 37-Seconds 18-Milliseconds 256-Microseconds 852-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumMinutesTimeStr() (string, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumMinutesTimeStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumMinutes(),
		ePrefix)

	if err != nil {
		return "", err
	}

	str := ""

	str += fmt.Sprintf("%v-Minutes ", tDur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", tDur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetCumMonthsDaysCalcDto - Returns a new TimeDurationDto calculated
// for 'cumulative months'.
//
// The time values of the current TimeDurationDto are re-calculated and
// returned as a new TimeDurationDTo as 'cumulative months'.
// This means that Years are ignored and assigned a zero value. Instead,
// Years and Months are consolidated and presented as 'cumulative months'.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a new
//       valid and fully populated 'TimeDurationDto' instance will
//       be returned.
//
//       The new, returned TimeDurationDto instance will have a
//       calculation type of 'TDurCalcType(0).CumMonths()'
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Cumulative Months-Days Format:
//
//  "3-Months 27-Days 19-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumMonthsDaysCalcDto() (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumMonthsDaysCalcDto() "

	if int64(tDur.timeDuration) == 0 {
		return TimeDurationDto{}, nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumMonths(),
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// GetCumMonthsDaysTimeStr - Returns Cumulative Months Display
// showing Months, Days, Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
//
// Years are ignored and assigned a zero value. Instead, years and
// months are consolidated and presented as cumulative months and
// days.
//
// This method will NOT modify the internal data fields of the
// current TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) formatted as cumulative months/days. See Example String
//       below.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Return String:
//
//  "3-Months 27-Days 19-Hours 6-Minutes 46-Seconds 666-Milliseconds 132-Microseconds 70-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumMonthsDaysTimeStr() (string, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumMonthsDaysTimeStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumMonths(),
		ePrefix)

	if err != nil {
		return "", err
	}

	str := ""

	str += fmt.Sprintf("%v-Months ", t2Dur.Months)

	str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetCumNanosecondsCalcDto - Returns a new TimeDurationDto. The time
// values of the current TimeDurationDto are recalculated for
// 'cumulative nanoseconds' and returned as the new TimeDurationDto
// object.
//
// This means that years, months, days, seconds, milliseconds and
// microseconds are ignored and set to a zero value.  Instead,
// months, days, seconds, milliseconds, microseconds and nanoseconds
// are consolidated and stored as cumulative nanoseconds.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a new
//       valid and fully populated 'TimeDurationDto' instance will
//       be returned.
//
//       The new, returned TimeDurationDto instance will have a
//       calculation type of 'TDurCalcType(0).CumNanoseconds()'
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Cumulative Nanoseconds Format:
//
//  "832-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumNanosecondsCalcDto() (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumNanosecondsCalcDto() "

	if int64(tDur.timeDuration) == 0 {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"\nError: Time Duration is ZERO value!\n")
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumNanoseconds(),
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// GetCumNanosecondsTimeStr - Returns duration formatted as
// Nanoseconds. DisplayStr shows Nanoseconds expressed as a
// 64-bit integer value.
//
// This means that years, months, days, seconds, milliseconds and
// microseconds are ignored and set to a zero value.  Instead,
// years, months, days, seconds, milliseconds, microseconds and
// nanoseconds are consolidated and stored as cumulative
// nanoseconds.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) formatted as cumulative nanoseconds. See Example String
//       below.
// __________________________________________________________________________
//
// Example Return String:
//
//  "832-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumNanosecondsTimeStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	str := fmt.Sprintf("%v-Nanoseconds", int64(tDur.timeDuration))

	return str
}

// GetCumSecondsCalcDto - Returns a new TimeDurationDto calculated
// for 'cumulative seconds'.
//
// The time values of the current TimeDurationDto are re-calculated and
// returned as the new TimeDurationDTo as 'cumulative seconds'.
// This means that Years, months, weeks, week days, date days, hours,
// and minutes are ignored and assigned a zero value. Instead,
// time duration is consolidated and presented as 'cumulative seconds'
// including seconds, milliseconds, microseconds and nanoseconds.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a new
//       valid and fully populated 'TimeDurationDto' instance will
//       be returned.
//
//       The new, returned TimeDurationDto instance will have a
//       calculation type of 'TDurCalcType(0).CumSeconds()'
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Cumulative Seconds Format:
//
//  "62-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumSecondsCalcDto() (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumSecondsCalcDto() "

	if int64(tDur.timeDuration) == 0 {
		return TimeDurationDto{}, nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumSeconds(),
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// GetCumSecondsTimeStr - Returns a formatted time string presenting
// time duration as cumulative seconds. The display shows Seconds,
// Milliseconds, Microseconds and Nanoseconds.
//
// This method will NOT modify the internal data fields of the
// current TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) formatted as cumulative seconds. See Example String below.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Return String:
//
//  "62-Seconds 0-Milliseconds 0-Microseconds 0-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumSecondsTimeStr() (string, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumSecondsTimeStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumSeconds(),
		ePrefix)

	if err != nil {
		return "", err
	}

	str := ""

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetCumWeeksCalcDto - Returns a new TimeDurationDto re-calculated for 'Cumulative Weeks'.
// The time values of the current TimeDurationDto are converted to cumulative weeks and
// stored in the returned 'TimeDurationDto' instance.
//
// 'Cumulative Weeks' means that Years and Months are ignored and assigned zero values.
// Instead, Years, Months and Weeks are consolidated and stored as cumulative Weeks,
// WeekDays, Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a new
//       valid and fully populated 'TimeDurationDto' instance will
//       be returned.
//
//       The new, returned TimeDurationDto instance will have a
//       calculation type of 'TDurCalcType(0).CumWeeks()'
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Cumulative Weeks Format:
//
//  126-Weeks 1-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetCumWeeksCalcDto() (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumWeeksCalcDto() "

	if int64(tDur.timeDuration) == 0 {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix + "Error: Time Duration is Zero")
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumWeeks(),
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// GetCumWeeksDaysTimeStr - Returns time duration expressed as Weeks, WeekDays, Hours,
// Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds. Years, Months and
// Days are ignored and assigned a zero value.
//
// Instead, Years, Months and Days are consolidated and presented as cumulative
// Weeks.
//
// This method will NOT modify the internal data fields of the
// current TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) formatted as cumulative weeks. See Example String below.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Return String
//
//  126-Weeks 1-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetCumWeeksDaysTimeStr() (string, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetCumWeeksDaysTimeStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.CumWeeks(),
		ePrefix)

	if err != nil {
		return "", err
	}

	str := ""

	str += fmt.Sprintf("%v-Weeks ", t2Dur.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// GetDefaultDurationStr - Returns duration formatted
// as nanoseconds. The DisplayStr shows the default
// string value for duration.
//
// __________________________________________________________________________
//
// Return Value:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
// __________________________________________________________________________
//
// Example Return:
//
//  "61h26m46.864197832sz"
//
func (tDur *TimeDurationDto) GetDefaultDurationStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return fmt.Sprintf("%v", tDur.timeDuration)
}

// GetDurationFromDays - returns a time Duration value
// based on the number of days passed into this method.
// No changes are made to or stored in the existing
// TimeDurationDto data fields.
//
// This method will NOT modify the internal data fields
// of the current TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  days int64
//     - A number of days which will be converted to
//       a time.Duration value.
//
// __________________________________________________________________________
//
// Return Value:
//
//  time.Duration - The value of this returned 'time.Duration' object is
//                  equivalent to the number of 24-hour days specified
//                  by input parameter 'days'.
//
func (tDur TimeDurationDto) GetDurationFromDays(days int64) time.Duration {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return time.Duration(days*24) * time.Hour

}

// GetDurationFromHours - returns a time Duration value
// based on the number of hours passed into this method.
// No changes are made to or stored in the existing
// TimeDurationDto data structures.
//
// This method will NOT modify the internal data fields
// of the current TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  hours int64 - The number of hours which will be converted to
//                a time.Duration value.
//
// __________________________________________________________________________
//
// Return Value:
//
//  time.Duration - The value of this returned 'time.Duration' object is
//                  equivalent to the number of 60-minute hours specified
//                  by input parameter 'hours'.
//
func (tDur TimeDurationDto) GetDurationFromHours(hours int64) time.Duration {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return time.Duration(hours) * time.Hour

}

// GetDurationFromMinutes - returns a time Duration value
// based on the number of minutes passed into this method.
// No changes are made to or stored in the existing
// TimeDurationDto data structures.
//
// This method will NOT modify the internal data fields of
// the current TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  minutes int64 - The number of minutes which will be converted to
//                  a time.Duration value.
//
// __________________________________________________________________________
//
// Return Value:
//
//  time.Duration - The value of this returned 'time.Duration' object is
//                  equivalent to the number of 60-second minutes specified
//                  by input parameter 'minutes'.
//
func (tDur TimeDurationDto) GetDurationFromMinutes(minutes int64) time.Duration {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return time.Duration(minutes) * time.Minute

}

// GetDurationFromSeconds - returns a time Duration value
// based on the number of seconds passed to this method.
// No changes are made to or stored in the existing
// TimeDurationDto data structures.
//
// This method will NOT modify the internal data fields
// of the current TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  minutes int64 - The number of seconds which will be converted to
//                  a time.Duration value.
//
// __________________________________________________________________________
//
// Return Value:
//
//  time.Duration - The value of this returned 'time.Duration' object is
//                  equivalent to the number seconds specified by input
//                  parameter 'seconds'.
//
func (tDur TimeDurationDto) GetDurationFromSeconds(seconds int64) time.Duration {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return time.Duration(seconds) * time.Second
}

// GetDurationFromTime - Calculates and returns a cumulative
// time duration based on time component input parameters
// consisting of hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
//
// Any combination of non-zero input parameters will be
// accumulated and converted to a valid time duration value.
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  hours          int
//     - Number of hours to be accumulated in the summary time duration
//
//  minutes        int
//     - Number of minutes to be accumulated in the summary time duration
//
//  seconds        int
//     - Number of seconds to be accumulated in the summary time duration
//
//  milliseconds   int
//     - Number of milliseconds to be accumulated in the summary time duration
//
//  microseconds   int
//     - Number of microseconds to be accumulated in the summary time duration
//
//  nanoseconds    int
//     - Number of nanoseconds to be accumulated in the summary time duration
//
// __________________________________________________________________________
//
// Return Value:
//
//  time.Duration - A time.Duration value which represents the sum of
//                  all input values.
//
func (tDur TimeDurationDto) GetDurationFromTime(
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) time.Duration {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	dtMech := DTimeMechanics{}

	return dtMech.GetDurationFromTimeComponents(
		0,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds)
}

// GetElapsedMinutesStr - Provides a quick means for formatting Years, Months,
// DateDays, Hours, Minutes, Seconds, Milliseconds, Microseconds and
// Nanoseconds. At a minimum, only Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds are included in the returned display
// string.
//
// This method only returns years, months, days or hours if those values
// are greater than zero.
//
// The data fields of the current TimeDurationDto instance (tDur) are
// NOT modified by this method
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
// __________________________________________________________________________
//
// Example Return:
//
//  0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetElapsedMinutesStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetElapsedMinutesStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.StdYearMth(),
		ePrefix)

	if err != nil {
		return fmt.Sprintf("%v\n", err.Error())
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.DateDays > 0 || str != "" {
		str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)
	}

	if t2Dur.Hours > 0 || str != "" {

		str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	}

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetElapsedSecsNanosecsStr - Provides a means of formatting Years,
// Months, DateDays, Hours, Minutes, Seconds and/ Nanoseconds. At a
// minimum, only seconds and nanoseconds are included in the returned
// display string.
//
// Years, months, days, hours and minutes are only included if the
// values are greater than zero.
//
// The data fields of the current TimeDurationDto instance (tDur) are
// NOT modified by this method
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
// __________________________________________________________________________
//
// Example Return:
//
//  0-Seconds 2832-Nanoseconds
//
func (tDur *TimeDurationDto) GetElapsedSecsNanosecsStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetElapsedSecsNanosecsStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.StdYearMth(),
		ePrefix)

	if err != nil {
		return fmt.Sprintf("%v\n", err.Error())
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.DateDays > 0 || str != "" {
		str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)
	}

	if t2Dur.Hours > 0 || str != "" {

		str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	}

	if t2Dur.Minutes > 0 || str != "" {

		str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	}

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.TotSubSecNanoseconds)

	return str

}

// GetElapsedTimeStr - Provides a quick means of formatting Years, Months,
// DateDays, Hours, Minutes, Seconds, Milliseconds, Microseconds and
// Nanoseconds. At a minimum, only Nanoseconds are included in the returned
// display string.
//
// If Years, Months and DateDays, Hours, Minutes, Seconds, Milliseconds
// and Microseconds have zero values, they will be excluded from the
// returned display string.
//
// The time.Duration value used to format this display is taken from the
// internal data field of the current 'TimeDurationDto' object,
// 'tDur.timeDuration'.
//
// This method only returns date time elements with value greater than
// zero. If all values are zero, the string will display Nanoseconds.
//
// The data fields of the current TimeDurationDto instance (tDur) are
// NOT modified by this method
//
// __________________________________________________________________________
//
// Return Value:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
// __________________________________________________________________________
//
// Example Return:
//
//  864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetElapsedTimeStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetElapsedTimeStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.StdYearMth(),
		ePrefix)

	if err != nil {
		return fmt.Sprintf("%v\n", err.Error())
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.DateDays > 0 || str != "" {
		str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)
	}

	if t2Dur.Hours > 0 || str != "" {

		str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	}

	if t2Dur.Minutes > 0 || str != "" {

		str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	}

	if t2Dur.Seconds > 0 || str != "" {

		str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	}

	if t2Dur.Milliseconds > 0 || str != "" {

		str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	}

	if t2Dur.Microseconds > 0 || str != "" {

		str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	}

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetThisEndDateTime - Returns the ending date time for this
// duration as a Type, 'time.Time'.  This value is extracted from
// private member variable TimeDurationDto.endDateTimeTz.
//
func (tDur *TimeDurationDto) GetThisEndDateTime() time.Time {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.endDateTimeTz.dateTimeValue
}

// GetTypeEndDateTimeString - Returns the ending date time for
// this duration as a string.
//
func (tDur *TimeDurationDto) GetThisEndDateTimeString() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.endDateTimeTz.String()
}

// GetTypeStartTimeDateTz - Returns the ending date time for this
// duration as a Type,'DateTzDTo'. This value is stored in private
// member variable TimeDurationDto.endDateTimeTz.
//
func (tDur *TimeDurationDto) GetThisEndDateTimeTz() DateTzDto {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.endDateTimeTz
}

// GetThisStartDateTime - Returns the starting date time for this
// duration as a Type, 'time.Time'.  This value is extracted from
// private member variable TimeDurationDto.startDateTimeTz.
func (tDur *TimeDurationDto) GetThisStartDateTime() time.Time {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.startDateTimeTz.dateTimeValue
}

// GetThisStartDateTimeString - Returns the starting date time for
// this duration as a string.
//
func (tDur *TimeDurationDto) GetThisStartDateTimeString() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.startDateTimeTz.String()
}

// GetThisStartDateTimeTz - Returns the starting date time for this
// duration as a Type,'DateTzDTo'. This value is stored in private
// member variable TimeDurationDto.startDateTimeTz.
//
func (tDur *TimeDurationDto) GetThisStartDateTimeTz() DateTzDto {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.startDateTimeTz.CopyOut()
}

// GetThisTimeDuration - Returns the time duration as Type, 'time.Duration'.
// This value is extracted from private member variable TimeDurationDto.timeDuration.
//
func (tDur *TimeDurationDto) GetThisTimeDuration() time.Duration {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.timeDuration
}

// GetThisTimeDurationCalcType - Returns the time duration calculation type
// associated with this TimeDurationDto instance. This value is extracted
// from private member variable TimeDurationDto.timeDurCalcType.
//
func (tDur *TimeDurationDto) GetThisTimeDurationCalcType() TDurCalcType {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.timeDurCalcType
}

// GetYearMthDaysTimeAbbrvStr - Abbreviated formatting of time duration
// as Years, Months, DateDays, Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds. At a minimum, only Hours, Minutes,
// Seconds, Milliseconds, Microseconds and Nanoseconds will be included
// in the returned display string.
//
// This method only returns date time elements with value greater than
// zero. If all values are zero, the string will still display Hours,
// Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds.
//
// The data fields of the current TimeDurationDto instance (tDur) are
// NOT modified by this method
//
// __________________________________________________________________________
//
// Return Value:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
// __________________________________________________________________________
//
// Example Return:
//
//  0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearMthDaysTimeAbbrvStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYearMthDaysTimeAbbrvStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.StdYearMth(),
		ePrefix)

	if err != nil {
		return fmt.Sprintf("%v\n", err.Error())
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.DateDays > 0 || str != "" {
		str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)
	}

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetYearMthDaysTimeStr - Calculates Duration and breakdowns
// time elements by Years, Months, Date Days, Hours, Minutes,
// Seconds, Milliseconds, Microseconds and Nanoseconds.
//
// Years, Months and Date Days are only included in the
// returned display string if their values are greater than
// zero. In contrast, Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds will always be included in
// the returned display string even if they have zero values.
//
// The data fields of the current TimeDurationDto instance (tDur) are
// NOT modified by this method
//
// __________________________________________________________________________
//
// Return Values:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
// __________________________________________________________________________
//
// Example Return:
//
//  12-Years 3-Months 2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
//  If Years, Months and Days have a zero value, only the time components will be displayed as
//  shown in the following example:
//
//   13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearMthDaysTimeStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYearMthDaysTimeStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.StdYearMth(),
		ePrefix)

	if err != nil {
		return fmt.Sprintf("%v\n", err.Error())
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.DateDays > 0 || str != "" {
		str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)
	}

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetYearsMthsWeeksTimeAbbrvStr - Abbreviated formatting of Years, Months,
// Weeks, WeekDays, Hours, Minutes, Seconds, Milliseconds, Microseconds,
// Nanoseconds.
//
// This method only returns date time elements with value greater than
// zero. If all values are zero, the string will still display Hours,
// Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds.
//
// The data fields of the current TimeDurationDto instance (tDur) are
// NOT modified by this method
//
// __________________________________________________________________________
//
// Return Value:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
// __________________________________________________________________________
//
// Example Return:
//
// 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearsMthsWeeksTimeAbbrvStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYearsMthsWeeksTimeAbbrvStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.StdYearMth(),
		ePrefix)

	if err != nil {
		return fmt.Sprintf("%v\n", err.Error())
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	if t2Dur.Weeks > 0 || str != "" {
		str += fmt.Sprintf("%v-Weeks ", t2Dur.Weeks)
	}

	if t2Dur.WeekDays > 0 || str != "" {
		str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)
	}

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetYearsMthsWeeksTimeStr - Returns a string containing time duration
// formatted as Years, Months, Weeks, WeekDays, Hours, Minutes, Seconds,
// Milliseconds, Microseconds and Nanoseconds.
//
// This method only returns date time elements with value greater than
// zero. However, if all values are zero, the returned string will still
// display Weeks, Week-Days, Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
//
// The data fields of the current TimeDurationDto instance (tDur) are
// NOT modified by this method
//
// __________________________________________________________________________
//
// Return Value:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
// __________________________________________________________________________
//
// Example Return:
//
//  12-Years 3-Months 2-Weeks 1-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
//  At a minimum Weeks, WeekDays, Hours, Minutes, Seconds, Milliseconds,
//  Microseconds and Nanoseconds are always displayed.
//
//  3-Weeks 2-WeekDays 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearsMthsWeeksTimeStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYearsMthsWeeksTimeStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.StdYearMth(),
		ePrefix)

	if err != nil {
		return fmt.Sprintf("%v\n", err.Error())
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	str += fmt.Sprintf("%v-Weeks ", t2Dur.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetYrMthWkDayHrMinSecNanosecsStr - Returns duration formatted
// as Years, Months, Weeks, Week-Days, Hours, Minutes, Seconds
// and Nanoseconds.
//
// If Years and Months have zero values, they are excluded from
// the returned display string. However, Weeks, Week-Days, Hours,
// Minutes, Seconds and Nanoseconds are always displayed even if
// they have zero values.
//
// The data fields of the current TimeDurationDto instance (tDur)
// are NOT modified by this method
//
// __________________________________________________________________________
//
// Return Value:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
// __________________________________________________________________________
//
// Example Return:
//
//  3-Years 2-Months 3-Weeks 2-WeekDays 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYrMthWkDayHrMinSecNanosecsStr() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetYrMthWkDayHrMinSecNanosecsStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds"
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.StdYearMth(),
		ePrefix)

	if err != nil {
		return fmt.Sprintf("%v\n", err.Error())
	}

	str := ""

	if t2Dur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", t2Dur.Years)
	}

	if t2Dur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", t2Dur.Months)
	}

	str += fmt.Sprintf("%v-Weeks ", t2Dur.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.TotSubSecNanoseconds)

	return str
}

// GetGregorianYearCalcDto - Returns a new TimeDurationDto in which years are
// calculated as 'Gregorian Years'.
//
// Unlike other calculations which use a Standard 365-day year consisting of
// 24-hour days which take into account correct leap years, a Gregorian Year
// consists of 365 days, 5-hours, 59-minutes and 12 Seconds. For the Gregorian
// calendar the average length of the calendar year (the mean year) across
// the complete leap cycle of 400 Years is 365.2425 days.
//
// The Gregorian calendar was first applied on 15 October 1582. It represents
// the primary calendar in use today across the globe. The name 'Gregorian'
// derives from  Pope Gregory XIII who introduced the calendar.
//
// Sources:
//  https://en.wikipedia.org/wiki/Year
//  https://en.wikipedia.org/wiki/Gregorian_calendar
//
// This method will NOT modify the internal data fields of the current
// TimeDurationDto instance, 'tDur'.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a new
//       valid and fully populated 'TimeDurationDto' instance will
//       be returned.
//
//       The new, returned TimeDurationDto instance will have a
//       calculation type of 'TDurCalcType(0).GregorianYears()'
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
func (tDur *TimeDurationDto) GetGregorianYearCalcDto() (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetGregorianYearCalcDto() "

	if int64(tDur.timeDuration) == 0 {
		return TimeDurationDto{}, nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.GregorianYears(),
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// GetGregorianYearDurationStr - Returns a string showing the
// breakdown of duration by Gregorian Years, WeekDays, Hours, Minutes,
// Seconds, Milliseconds, Microseconds and Nanoseconds. Unlike
// other calculations which use a Standard 365-day year consisting
// of 24-hour days, a Gregorian Year consists of 365 days, 5-hours,
// 59-minutes and 12 Seconds. For the Gregorian calendar the
// average length of the calendar year (the mean year) across
// the complete leap cycle of 400 Years is 365.2425 days.
//
// The Gregorian calendar was first applied on 15 October 1582. It
// represents the primary calendar in use today across the globe.
// The name 'Gregorian' derives from  Pope Gregory XIII who
// introduced the calendar.
//
// Sources:
//   https://en.wikipedia.org/wiki/Year
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//
// This method returns a string listing time values for Gregorian
// Years, Months, Days, Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
//
// All specified time values are included in the returned display
// string regardless of whether they have a zero value.
//
// The data fields of the current TimeDurationDto instance (tDur)
// are NOT modified by this method
//
// __________________________________________________________________________
//
// Return Value:
//
//  string
//     - A string containing the time duration for the current TimeDurationDto
//       object (tDur) listing all non-zero time components as shown in the
//       example below.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Return:
//
//  3-Gregorian Years 2-Months 5-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetGregorianYearDurationStr() (string, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.GetGregorianYearDurationStr() "

	if int64(tDur.timeDuration) == 0 {
		return "0-Nanoseconds", nil
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := tDurDtoUtil.copyOut(tDur, ePrefix)

	err := tDurDtoUtil.reCalcTimeDurationAllocation(
		&t2Dur,
		TDurCalc.GregorianYears(),
		ePrefix)

	if err != nil {
		return "", err
	}

	str := fmt.Sprintf("%v-Gregorian Years ", t2Dur.Years)

	str += fmt.Sprintf("%v-Months ", t2Dur.Months)

	str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str, nil
}

// New - Returns a new TimeDurationDto instance
// with member variables initialized to their
// zero values.
//
func (tDur TimeDurationDto) New() TimeDurationDto {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	timeDur2 := TimeDurationDto{}

	timeDur2.lock = new(sync.Mutex)

	timeDur2.timeDuration = time.Duration(0)

	timeDur2.timeDurCalcType = TDurCalc.None()

	timeDur2.startDateTimeTz = DateTzDto{}.New()

	timeDur2.endDateTimeTz = DateTzDto{}.New()

	timeDur2.timeMathCalcMode = TCalcMode.None()

	return timeDur2
}

// NewAutoEnd - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'timeZoneLocation' input parameters.
//
// The input parameter 'startDateTime' is first converted to the Time Zone specified
// by the input parameter, 'timeZoneLocation'.
//
// The ending date time is calculated automatically by assigning the date time returned
// by calling time.Now(). The ending date time that is assigned by time.Now() is converted
// to the time zone specified by input parameter, 'timeZoneLocation'.
//
// If the calculated ending date time is prior to 'startDateTime', the values are
// reversed and ending date time is assigned to 'startDateTime' while 'startDateTime'
// is assigned to ending date time.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both the starting date time and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     time.Time
//     - Starting date time
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//       - If this method proceeds to successful completion, a valid
//         and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//       - If this method proceeds to successful completion, the returned
//         error instance is set to 'nil'. If an error is encountered, the
//         error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//   tDurDto, err := TimeDurationDto{}.NewAutoEnd(
//                                     startTime,
//                                     TDurCalc.StdYearMth(),
//                                     TZones.US.Central(),
//                                     TCalcMode.LocalTimeZone(),
//                                     FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewAutoEnd(
	startDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewAutoEnd() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimes(
		&tDur2,
		startDateTime,
		time.Now().UTC(),
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}.New(), err
	}

	return tDur2, nil
}

// NewAutoStart - Creates and returns a new TimeDurationDto instance. Starting date time is
// automatically initialized by calling time.Now(). Afterwards, start date time is converted
// to the Time Zone specified in input parameter, 'timeZoneLocation'.
//
// This method will set an arbitrary ending date time to be used as a place holder.  Thereafter
// a call to method TimeDurationDto.SetAutoEnd() will automatically set the final
// ending date time and compute the associated time duration.
//
// Use of these two methods, 'NewAutStart' and 'SetAutoEnd', constitutes a stop watch feature which
// can be triggered to measure elapsed time.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//   tDurDto, err := TimeDurationDto{}.NewAutoStart(
//                    TDurCalc.StdYearMth(),
//                    TZones.US.Central(),
//                    TCalcMode.LocalTimeZone(),
//                    FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewAutoStart(
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewAutoStart() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	startEndTime := time.Now().UTC().
		AddDate(0,0,2)

	err := tDurDtoUtil.setStartEndTimes(
		&tDur2,
		time.Now().UTC(),
		startEndTime,
		tDurCalcType,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewDefaultStartEndTimes - Creates and returns a new TimeDurationDto based
// on starting and ending date times.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location and Time Math Calculation Mode.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTime'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime time.Time
//     - Starting date time
//
//
//  endDateTime   time.Time
//     - Ending date time
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewDefaultStartEndTimes(
//                                   startDateTime,
//
func (tDur TimeDurationDto) NewDefaultStartEndTimes(
	startDateTime,
	endDateTime time.Time) (TimeDurationDto, error) {


	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewDefaultStartEndTimes() "
	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimes(
		&tDur2,
		startDateTime,
		endDateTime,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewDefaultStartEndTimesTz - Creates and returns a new TimeDurationDto based
// on starting and ending date times.
//
// Starting and ending date times are encapsulated through input parameters,
// 'startDateTimeTz' and 'endDateTimeTz'. Both parameters are instances of type
// 'DateTzDto'.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location and Time Math Calculation Mode.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTimeTz'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTimeTz DateTzDto
//     - Starting date time
//
//
//  endDateTimeTz   DateTzDto
//     - Ending date time
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewDefaultStartEndTimes(
//                                   startDateTimeTz,
//                                   endDateTimeTz)
//
func (tDur TimeDurationDto) NewDefaultStartEndTimesTz(
	startDateTimeTz,
	endDateTimeTz DateTzDto) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewDefaultStartEndTimesTz() "
	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimes(
		&tDur2,
		startDateTimeTz.dateTimeValue,
		endDateTimeTz.dateTimeValue,
		TDurCalc.StdYearMth(),
		startDateTimeTz.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewDefaultStartTimeDuration - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime' and time 'duration'. The time duration value is added to 'startDateTime' in order
// to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time. The
// actual starting date time is computed by subtracting duration from ending date time.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location and Time Math Calculation Mode.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTime'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime time.Time
//     - Starting date time for the duration calculation.
//
//  duration  time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewDefaultStartTimeDuration(
//                                   startDateTime,
//                                   duration)
//
func (tDur TimeDurationDto) NewDefaultStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration) (TimeDurationDto, error) {


	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewDefaultStartTimeDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDuration(
		&tDur2,
		startDateTime,
		duration,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewDefaultStartTimeTzDuration - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime' and time 'duration'. The time duration value is added to 'startDateTime' in order
// to compute the ending date time.
//
// Starting date time is passed as an input parameter of type, 'DateTzDto'.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time. The
// actual starting date time is computed by subtracting duration from ending date time.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location and Time Math Calculation Mode.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTimeTz'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTimeTz DateTzDto
//     - Starting date time for the duration calculation.
//
//  duration        time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewDefaultStartTimeTzDuration(
//                                   startDateTimeTz,
//                                   duration)
//
func (tDur TimeDurationDto) NewDefaultStartTimeTzDuration(
	startDateTimeTz DateTzDto,
	duration time.Duration) (TimeDurationDto, error) {


	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewDefaultStartTimeTzDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDuration(
		&tDur2,
		startDateTimeTz.dateTimeValue,
		duration,
		TDurCalc.StdYearMth(),
		startDateTimeTz.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewDefaultStartTimePlusTimeDto - Creates and returns a new TimeDurationDto setting
// the start date time, end date time and duration based on a starting date time
// and the time components contained in an instance of TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the time duration. TimeDto components are typically
// expressed as years, months, days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location and Time Math Calculation Mode.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTime'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     time.Time
//     - Starting date time. The ending date time will be computed
//       by adding the time components of the 'plusTimeDto' to
//       'startDateTime'.
//
//  plusTimeDto       TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be added to 'startDateTime' to compute
//       time duration and ending date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewDefaultStartTimePlusTimeDto(
//                    startTime,
//                    plusTimeDto)
//
func (tDur TimeDurationDto) NewDefaultStartTimePlusTimeDto(
	startDateTime time.Time,
	plusTimeDto TimeDto) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewDefaultStartTimePlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimePlusTimeDto(
		&tDur2,
		startDateTime,
		plusTimeDto,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewStartTimeTzPlusTimeDto - Creates and returns a new TimeDurationDto setting
// the start date time, end date time and duration based on a starting date time
// and the time components contained in a TimeDto.
//
// Starting date time is passed to this method as a type 'DateTzDto'.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the time duration. TimeDto components are typically
// expressed as years, months, days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location and Time Math Calculation Mode.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTimeTz'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     DateTzDto
//     - Starting date time. The ending date time will be computed
//       by adding the time components of the 'plusTimeDto' to
//       'startDateTime'.
//
//  plusTimeDto       TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be added to 'startDateTime' to compute
//       time duration and ending date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewDefaultStartTimeTzPlusTimeDto(
//                    startTimeTz,
//                    plusTimeDto)
//
func (tDur TimeDurationDto) NewDefaultStartTimeTzPlusTimeDto(
	startDateTimeTz DateTzDto,
	plusTimeDto TimeDto) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewDefaultStartTimeTzPlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimePlusTimeDto(
		&tDur2,
		startDateTimeTz.dateTimeValue,
		plusTimeDto,
		TDurCalc.StdYearMth(),
		startDateTimeTz.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewEndTimeMinusTimeDto - Creates and returns a new TimeDurationDto instance
// populated with starting date time, ending date time and time duration. Starting
// date time is calculated by subtracting a TimeDto time duration value from the
// ending date time. Ending date time is passed as a 'time.Time' value through
// input parameter 'endDateTime'.  Input parameter 'minusTimeDto' is passed as
// an instance of 'TimeDto' which encapsulates time duration formatted as a series
// time components such as years, months, days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds. To compute starting date time, the 'minusTimeDto'
// time duration value is subtracted from 'endDateTime'.
//
// Input parameter 'minusTimeDto' is first converted to absolute, or positive,
// time duration values before it is subsequently subtracted from ending date
// time in order to compute starting date time.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location and Time Math Calculation Mode.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'endDateTime'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  endDateTime   time.Time
//     - Ending date time. The starting date time will be computed
//       by subtracting minusTimeDto from 'endDateTime'
//
//  minusTimeDto  TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be subtracted from 'endDateTime' to compute
//       time duration and starting date time. 'minusTimeDto' is
//       first converted to absolute, or positive' duration value
//       before being subtracted from 'endDateTime'.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewDefaultEndTimeMinusTimeDto(
//                           endTime,
//                           minusTimeDto)
//
func (tDur TimeDurationDto) NewDefaultEndTimeMinusTimeDto(
	endDateTime time.Time,
	minusTimeDto TimeDto) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewDefaultEndTimeMinusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setEndTimeMinusTimeDto(
		&tDur2,
		endDateTime,
		minusTimeDto,
		TDurCalc.StdYearMth(),
		endDateTime.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewDefaultEndTimeTzMinusTimeDto - Creates and returns a new TimeDurationDto instance
// populated with starting date time, ending date time and time duration. Starting
// date time is calculated by subtracting a TimeDto time duration value from the
// ending date time. Ending date time is passed as an instance of 'DateTzDto' through
// input parameter 'endDateTimeTz'.  Input parameter 'minusTimeDto' is passed as
// an instance of 'TimeDto' which encapsulates time duration formatted as a series
// time components such as years, months, days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds. To compute starting date time, the 'minusTimeDto'
// time duration value is subtracted from 'endDateTimeTz'.
//
// Input parameter 'minusTimeDto' is first converted to absolute, or positive,
// time duration values before it is subsequently subtracted from ending date
// time in order to compute starting date time.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location and Time Math Calculation Mode.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'endDateTimeTz'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  endDateTimeTz  DateTzDto
//     - Ending date time. The starting date time will be computed
//       by subtracting minusTimeDto from 'endDateTime'
//
//  minusTimeDto   TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be subtracted from 'endDateTime' to compute
//       time duration and starting date time. 'minusTimeDto' is
//       first converted to absolute, or positive' duration value
//       before being subtracted from 'endDateTimeTz'.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewDefaultEndTimeTzMinusTimeDto(
//                           endDateTimeTz,
//                           minusTimeDto)
//
func (tDur TimeDurationDto) NewDefaultEndTimeTzMinusTimeDto(
	endDateTimeTz DateTzDto,
	minusTimeDto TimeDto) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewDefaultEndTimeTzMinusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setEndTimeMinusTimeDto(
		&tDur2,
		endDateTimeTz.dateTimeValue,
		minusTimeDto,
		TDurCalc.StdYearMth(),
		endDateTimeTz.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewStartEndTimes - Creates and returns a new TimeDurationDto based on starting
// and ending date times.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime time.Time
//     - Starting date time
//
//
//  endDateTime   time.Time
//     - Ending date time
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewStartEndTimes(
//                                   startTime,
//                                   endTime,
//                                     TDurCalc.StdYearMth(),
//                                     TZones.US.Central(),
//                                     TCalcMode.LocalTimeZone(),
//                                     FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartEndTimes(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimes() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimes(
		&tDur2,
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndTimesTz - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTimeTz' and 'endDateTimeTz' input
// parameters.  Both parameters are instances of type, 'DateTzDto'.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     DateTzDto
//     - Starting time
//
//
//  endDateTime       DateTzDto
//     - Ending time
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewStartEndTimesTz(
//                                    startTimeTz,
//                                    endTimeTz,
//                                     TDurCalc.StdYearMth(),
//                                     TZones.US.Central(),
//                                     TCalcMode.LocalTimeZone(),
//                                     FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartEndTimesTz(
	startDateTimeTz,
	endDateTimeTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimes(
		&tDur2,
		startDateTimeTz.dateTimeValue,
		endDateTimeTz.dateTimeValue,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDuration - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime' and time duration. The time duration value is added to 'startDateTime' in
// order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time. The
// actual starting date time is computed by subtracting duration from ending date time.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime time.Time
//     - Starting date time for the duration calculation.
//
//  duration  time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewStartTimeTzDuration(
//                                     startTime,
//                                     duration,
//                                     TDurCalc.StdYearMth(),
//                                     TZones.US.Central(),
//                                     TCalcMode.LocalTimeZone(),
//                                     FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDuration(
		&tDur2,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeTzDuration - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTimeTz', and time duration. 'startDateTimeTz' is  an instance of type, 'DateTzDto'.
//
// The duration value is added to the starting date time in order to compute the ending date time
// and the associated time duration.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and the
//  actual starting date time is computed by subtracting duration.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both the starting date time and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTimeTz   DateTzDto
//     - Starting date time for the duration calculation encapsulated
//       in a 'DateTzDto' object.
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewStartTimeTzDuration(
//                                    startTimeTz,
//                                    duration,
//                                    TDurCalc.StdYearMth(),
//                                    TZones.US.Central(),
//                                    TCalcMode.LocalTimeZone(),
//                                    FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartTimeTzDuration(
	startDateTimeTz DateTzDto,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeTzDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDuration(
		&tDur2,
		startDateTimeTz.dateTimeValue,
		duration,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimePlusTimeDto - Creates and returns a new TimeDurationDto setting
// the start date time, end date time and duration based on a starting date time
// and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the time duration. TimeDto components are typically
// expressed as years, months, days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both the starting date time and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     time.Time
//     - Starting date time. The ending date time will be computed
//       by adding the time components of the 'plusTimeDto' to
//       'startDateTime'.
//
//  plusTimeDto       TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be added to 'startDateTime' to compute
//       time duration and ending date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewStartTimePlusTimeDto(
//                    startTime,
//                    plusTimeDto,
//                    TDurCalc.StdYearMth(),
//                    TZones.US.Central(),
//                    TCalcMode.LocalTimeZone(),
//                    FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartTimePlusTimeDto(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimePlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimePlusTimeDto(
		&tDur2,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewStartTimeTzPlusTimeDto - Creates and returns a new TimeDurationDto setting
// the starting date time, ending date time and duration based on a starting date time
// and the time components contained in a TimeDto which are treated as a time duration
// value.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both the starting date time and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTimeTz   DateTzDto
//     - Starting date time. The ending date time will be computed
//       by adding the time components of the 'plusTimeDto' to
//       'startDateTimeTz'.
//
//  plusTimeDto       TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be added to 'startDateTime' to compute
//       time duration and ending date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewStartTimeTzPlusTimeDto(
//                    startTimeTz,
//                    plusTimeDto,
//                                     TDurCalc.StdYearMth(),
//                                     TZones.US.Central(),
//                                     TCalcMode.LocalTimeZone(),
//                                     FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartTimeTzPlusTimeDto(
	startDateTimeTz DateTzDto,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeTzPlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimePlusTimeDto(
		&tDur2,
		startDateTimeTz.dateTimeValue,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewEndTimeMinusTimeDto - Creates and returns a new TimeDurationDto instance
// populated with starting date time, ending date time and time duration. Starting
// date time is calculated by subtracting a TimeDto time duration value from
// ending date time. Ending date time is passed as a 'time.Time' value through
// input parameter 'endDateTime'.  Input parameter 'minusTimeDto' is passed as
// an instance of 'TimeDto' which encapsulates time duration formatted as a series
// time components such as years, months, days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds. To compute starting date time, the 'minusTimeDto'
// time duration value is subtracted from 'endDateTime'.
//
// Input parameter 'minusTimeDto' is first converted to absolute, or positive,
// time duration values before it is subsequently subtracted from ending date
// time in order to compute starting date time.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both the starting date time and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  endDateTime   time.Time
//     - Ending date time. The starting date time will be computed
//       by subtracting minusTimeDto from 'endDateTime'
//
//  minusTimeDto  TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be subtracted from 'endDateTime' to compute
//       time duration and starting date time. 'minusTimeDto' is
//       first converted to absolute, or positive' duration value
//       before being subtracted from 'endDateTime'.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewEndTimeMinusTimeDto(
//                           endDateTime,
//                           minusTimeDto,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           TCalcMode.LocalTimeZone(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewEndTimeMinusTimeDto(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewEndTimeMinusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setEndTimeMinusTimeDto(
		&tDur2,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewEndTimeTzMinusTimeDto - Creates and returns a new TimeDurationDto instance
// populated with starting date time, ending date time and time duration. Starting
// date time is calculated by subtracting a TimeDto time duration value from
// ending date time. Ending date time is passed as an instance of 'DateTzDto'
// through input parameter 'endDateTimeTz'.  Input parameter 'minusTimeDto' is
// passed as an instance of 'TimeDto' which encapsulates time duration formatted
// as a series time components such as years, months, days, hours, minutes, seconds,
// milliseconds, microseconds and nanoseconds. To compute starting date time, the
// 'minusTimeDto' time duration value is subtracted from 'endDateTimeTz'.
//
// Input parameter 'minusTimeDto' is first converted to absolute, or positive,
// time duration values before it is subsequently subtracted from ending date
// time in order to compute starting date time.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both the starting date time and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  endDateTimeTz     DateTzDto
//     - Ending date time. The starting date time will be computed
//       by subtracting 'minusTimeDto' from 'endDateTimeTz'
//
//  minusTimeDto      TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be subtracted from 'endDateTimeTz' to compute
//       time duration and starting date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Values:
//
//  TimeDurationDto
//     - If this method proceeds to successful completion, a valid
//       and fully populated 'TimeDurationDto' instance is returned.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewEndTimeTzMinusTimeDto(
//                           endDateTimeTz,
//                           minusTimeDto,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           TCalcMode.LocalTimeZone(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewEndTimeTzMinusTimeDto(
	endDateTimeTz DateTzDto,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewEndTimeTzMinusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := TimeDurationDto{}

	err := tDurDtoUtil.setEndTimeMinusTimeDto(
		&t2Dur,
		endDateTimeTz.dateTimeValue,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return t2Dur, nil
}

// ReCalcTimeDurationAllocation - Re-calculates and allocates time duration for the current
// TimeDurationDto instance over the various time components (years, months, weeks, weekdays,
// datedays, hour, minutes, seconds, milliseconds, microseconds and nanoseconds) depending
// on the value of the 'TDurCalcType' input parameter.
//
// Type 'TDurCalcType' is located in source file:
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  tDurCalcType TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
func (tDur *TimeDurationDto) ReCalcTimeDurationAllocation(
	tDurCalcType TDurCalcType) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.ReCalcTimeDurationAllocation() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.reCalcTimeDurationAllocation(
		tDur,
		tDurCalcType,
		ePrefix)
}

// ReCalcEndDateTimeToNow - Recomputes time duration values for the
// current TimeDurationDto by setting ending date time to time.Now().
// This is useful in stop watch applications.
//
// The Time Zone Location is derived from the existing starting date
// time, 'tDur.startDateTimeTz'.  The Calculation type is taken from
// the existing calculation type, 'tDur.timeDurCalcType'.
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
func (tDur *TimeDurationDto) ReCalcEndDateTimeToNow() error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.ReCalcEndDateTimeToNow() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setAutoEnd(
		tDur,
		tDur.timeDurCalcType,
		tDur.startDateTimeTz.GetBestConvertibleTimeZone().locationName,
		tDur.timeMathCalcMode,
		tDur.startDateTimeTz.dateTimeFmt,
		ePrefix)
}

// SetAutoEnd - When called, this method automatically sets the ending date
// time and re-calculates the time duration for the current TimeDurationDto
// instance.
//
// Ending date time is assigned the value returned by time.Now(). This ending
// date time is converted to the specified Time Zone specified by the Time Zone
// Location associated with the current starting date time value.
//
// When used together, the two methods 'NewAutoStart' and this method, 'SetAutoEnd'
// function as a stop watch feature. Simply calling these functions can set
// the starting date time and later the ending date time to measure elapsed time, or
// time duration.
//
// The time duration calculation type is taken from the current TimeDurationDto
// calculation type setting.
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//
//  dt, err := TimeDurationDto{}.NewAutoStart(
//                    TDurCalc.StdYearMth(),
//                    TZones.US.Central(),
//                    TCalcMode.LocalTimeZone(),
//                    FmtDateTimeYrMDayFmtStr)
//
//  if err != nil {
//     'Do Something'
//  }
//
//
//  err := dt.SetAutoEnd()
//
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
func (tDur *TimeDurationDto) SetAutoEnd() error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetAutoEnd() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setAutoEnd(
		tDur,
		tDur.timeDurCalcType,
		tDur.startDateTimeTz.GetBestConvertibleTimeZone().locationName,
		tDur.timeMathCalcMode,
		tDur.startDateTimeTz.dateTimeFmt,
		ePrefix)
}

// SetEndTimeMinusTimeDto - Sets the member data variables for the current instance
// of TimeDurationDto.  The current TimeDurationDto instance is populated with
// starting date time, ending date time and time duration. Starting date time
// is calculated by subtracting a TimeDto time duration value from ending date
// time. Ending date time is passed as a 'time.Time' value through input parameter
// 'endDateTime'.  Input parameter 'minusTimeDto' is passed as an instance of
// 'TimeDto' which encapsulates time duration formatted as a series time components
// such as years, months, days, hours, minutes, seconds, milliseconds, microseconds
// and nanoseconds. To compute starting date time, the 'minusTimeDto' time duration
// value is subtracted from 'endDateTime'.
//
// Input parameter 'minusTimeDto' is first converted to absolute, or positive,
// time duration values before it is subsequently subtracted from ending date
// time in order to compute starting date time.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both the starting date time and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  endDateTime     time.Time
//     - Ending date time. The starting date time will be computed
//       by subtracting minusTimeDto from 'endDateTime'
//
//  minusTimeDto      TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be subtracted from 'endDateTime' to compute
//       time duration and starting date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetEndTimeMinusTimeDto(
//                           endTime,
//                           minusTimeDto,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           TCalcMode.LocalTimeZone(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetEndTimeMinusTimeDto(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetEndTimeMinusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setEndTimeMinusTimeDto(
		tDur,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetDefaultEndTimeMinusTimeDto - Sets the member data variables for the current
// instance of TimeDurationDto.  The current TimeDurationDto instance is populated
// with starting date time, ending date time and time duration. Starting date time
// is calculated by subtracting a TimeDto time duration value from ending date
// time. Ending date time is passed as a 'time.Time' value through input parameter
// 'endDateTime'.  Input parameter 'minusTimeDto' is passed as an instance of
// 'TimeDto' which encapsulates time duration formatted as a series time components
// such as years, months, days, hours, minutes, seconds, milliseconds, microseconds
// and nanoseconds. To compute starting date time, the 'minusTimeDto' time duration
// value is subtracted from 'endDateTime'.
//
// Input parameter 'minusTimeDto' is first converted to absolute, or positive,
// time duration values before it is subsequently subtracted from ending date
// time in order to compute starting date time.
//
// This method will provide default values for Time Duration Calculation Type,
// Time Math Calculation Mode and Date Time Format.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location Name:        Extracted from input parameter, 'endDateTime'
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// For granular control over these default parameters, see method:
//       TimeDurationDto.SetEndTimeMinusTimeDto()
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  endDateTime     time.Time
//     - Ending date time. The starting date time will be computed
//       by subtracting minusTimeDto from 'endDateTime'
//
//  minusTimeDto      TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be subtracted from 'endDateTime' to compute
//       time duration and starting date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetDefaultEndTimeMinusTimeDto(
//                           endTime,
//                           minusTimeDto)
//
func (tDur *TimeDurationDto) SetDefaultEndTimeMinusTimeDto(
	endDateTime time.Time,
	minusTimeDto TimeDto) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetDefaultEndTimeMinusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setEndTimeMinusTimeDto(
		tDur,
		endDateTime,
		minusTimeDto,
		TDurCalc.StdYearMth(),
		endDateTime.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)
}

// SetDefaultStartEndTimes - Sets data field values for the current TimeDurationDto
// instance using a Start Date Time, End Date Time and a time zone specification.
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
//
// All data fields in the current TimeDurationDto instance are overwritten with
// the new time duration values.
//
// Input parameter 'minusTimeDto' is first converted to absolute, or positive,
// time duration values before it is subsequently subtracted from ending date
// time in order to compute starting date time.
//
// This method will provide default values for Time Duration Calculation Type,
// Time Math Calculation Mode and Date Time Format.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location Name:        Extracted from input parameter, 'startDateTime'
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// For granular control over these default parameters, see method:
//       TimeDurationDto.SetStartEndTimes()
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   time.Time
//     - Starting time
//
//
//  endDateTime     time.Time
//     - Ending time
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetDefaultStartEndTimes(
//                           startDateTime,
//                           endDateTime)
//
func (tDur *TimeDurationDto) SetDefaultStartEndTimes(
	startDateTime,
	endDateTime time.Time) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetDefaultStartEndTimes() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartEndTimes(
		tDur,
		startDateTime,
		endDateTime,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)
}

// SetDefaultStartEndTimesTz - Sets data field values for the current TimeDurationDto
// instance using a Start Date Time, End Date Time and a time zone specification.
//
// The Starting Date Time and Ending Date Time are submitted as type 'DateTzDto'.
//
// First, input parameters 'startDateTimeTz' and 'endDateTimeTz' are converted to
// the designated Time Zone Location. Next, 'startDateTimeTz' is subtracted from
// 'endDateTimeTz' to compute time duration.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location, Time Math Calculation Mode and Date Time Format.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTimeTz'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        Extracts the date time format from 'startDateTimeTz'.
//
//
// For granular control over these parameters, see:
//    TimeDurationDto.SetStartEndTimesTz()
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTimeTz DateTzDto
//     - Starting date time
//
//  endDateTimeTz   DateTzDto
//     - Ending date time
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetDefaultStartEndTimesTz(
//                           startDateTimeTz,
//                           endDateTimeTz)
//
func (tDur *TimeDurationDto) SetDefaultStartEndTimesTz(
	startDateTimeTz,
	endDateTimeTz DateTzDto) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetDefaultStartEndTimesTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartEndTimesTz(
		tDur,
		startDateTimeTz,
		endDateTimeTz,
		TDurCalc.StdYearMth(),
		startDateTimeTz.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		startDateTimeTz.GetDateTimeFmt(),
		ePrefix)
}

// SetSDefaultStartTimeDuration - Sets start time, end time and duration for the
// current TimeDurationDto instance. 'startDateTime' is converted to the
// specified by 'timeZoneLocation'. The duration value is added to 'startDateTime'
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date
// time. Thereafter, the actual starting date time is computed by subtracting
// duration.
//
// This method will provide default values for Time Duration Calculation Type,
// Time Math Calculation Mode and Date Time Format.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location Name:        Extracted from input parameter, 'startDateTime'
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// For granular control over these default parameters, see method:
//       TimeDurationDto.SetStartTimeDuration()
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     time.Time
//     - Starting date time for the duration calculation
//
//
//  duration        time.Duration
//     - Time Duration added to 'startDateTime' in order to
//       compute Ending Date Time.
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetDefaultStartTimeDuration(
//                           startDateTime,
//                           duration)
//
func (tDur *TimeDurationDto) SetDefaultStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetDefaultStartTimeDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimeDuration(
		tDur,
		startDateTime,
		duration,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)
}

// SetDefaultStartTimeTzDuration - Sets start time, end time and
// duration for the current TimeDurationDto instance.
//
// The input parameter, 'startDateTimeTz', is of type DateTzDto. It
// is converted to the specified 'timeZoneLocation', and added to
// the duration value in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTimeTz' is converted
// to ending date time. Thereafter, the actual starting date time is
// computed by subtracting duration.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location, Time Math Calculation Mode and Date Time Format.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTimeTz'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        Extracts the date time format from 'startDateTimeTz'.
//
//
// For granular control over these parameters, see method:
//    TimeDurationDto.SetStartTimeTzDuration()
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     DateTzDto
//     - Provides starting date time for the duration calculation
//
//
//  duration          time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetDefaultStartTimeTzDuration(
//                           startDateTimeTz,
//                           duration)
//
func (tDur *TimeDurationDto) SetDefaultStartTimeTzDuration(
	startDateTimeTz DateTzDto,
	duration time.Duration) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetDefaultStartTimeTzDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimeTzDuration(
		tDur,
		startDateTimeTz,
		duration,
		TDurCalc.StdYearMth(),
		startDateTimeTz.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		startDateTimeTz.GetDateTimeFmt(),
		ePrefix)
}

// SetDefaultStartTimePlusTimeDto - Sets start date time, end date time and duration
// based on a starting date time. The results of this calculation are used to
// populate the current TimeDurationDto instance.
//
// The time components of the TimeDto are added to the starting date time to
// compute the ending date time and the duration.
//
// This method will supply default values for Time Duration Calculation Type,
// Time Zone Location, Time Math Calculation Mode and Date Time Format.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTime'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        FmtDateTimeYrMDayFmtStr
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// For granular control over these parameters, see method:
//    TimeDurationDto.SetStartTimePlusTimeDto()
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     time.Time
//     - Starting date time. The ending date time will be computed
//       by adding the time components of the 'plusTimeDto' to
//       'startDateTime'.
//
//
//  plusTimeDto       TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be added to 'startDateTime' to compute
//       time duration and ending date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetDefaultStartTimePlusTimeDto(
//                           startDateTime,
//                           plusTimeDto)
//
func (tDur *TimeDurationDto) SetDefaultStartTimePlusTimeDto(
	startDateTime time.Time,
	plusTimeDto TimeDto) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetDefaultStartTimePlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimePlusTimeDto(
		tDur,
		startDateTime,
		plusTimeDto,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		TCalcMode.LocalTimeZone(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)
}

// SetDefaultStartTimeTzPlusTimeDto - Sets start date time, end date time and duration
// based on a starting date time. The results of this calculation are used to populate
// the current TimeDurationDto instance.
//
// The starting date time is passed to this method as an instance of 'DateTzDto'.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// This method will supply default values for Time Duration Calculation Type, Time
// Zone Location, Time Math Calculation Mode and Date Time Format.
//
// Default Values:
// Time Duration Calculation Type: TDurCalc.StdYearMth()
// Time Zone Location:             Extracts the Time Zone Location from 'startDateTimeTz'.
// Time Math Calculation Mode:     TCalcMode.LocalTimeZone()
// Date Time Format String:        Extracts the date time format from 'startDateTimeTz'.
//
//
// For granular control over these parameters, see:
//    TimeDurationDto.SetStartTimeTzPlusTimeDto()
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTimeTz   DateTzDto
//     - Starting date time. The ending date time will be computed
//       by adding the time components of the 'plusTimeDto' to
//       'startDateTime'.
//
//  plusTimeDto       TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be added to 'startDateTime' to compute
//       time duration and ending date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetDefaultStartTimeTzPlusTimeDto(
//                           startDateTimeTz,
//                           plusTimeDto)
//
func (tDur *TimeDurationDto) SetDefaultStartTimeTzPlusTimeDto(
	startDateTimeTz DateTzDto,
	plusTimeDto TimeDto) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetDefaultStartTimeTzPlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimePlusTimeDto(
		tDur,
		startDateTimeTz.dateTimeValue,
		plusTimeDto,
		TDurCalc.StdYearMth(),
		startDateTimeTz.GetTimeZoneName(),
		TCalcMode.LocalTimeZone(),
		startDateTimeTz.GetDateTimeFmt(),
		ePrefix)
}

// SetStartEndTimes - Sets data field values for the current TimeDurationDto
// instance using a Start Date Time, End Date Time and a time zone specification.
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
//
// All data fields in the current TimeDurationDto instance are overwritten with
// the new time duration values.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   time.Time
//     - Starting time
//
//
//  endDateTime     time.Time
//     - Ending time
//
//
//  tDurCalcType TDurCalcType
//       - Specifies the calculation type to be used in allocating
//         time duration:
//
//         TDurCalcType(0).StdYearMth()
//           - Default - standard year, month week, day time calculation.
//
//         TDurCalcType(0).CumMonths()
//           - Computes cumulative months - no Years.
//
//         TDurCalcType(0).CumWeeks()
//           - Computes cumulative weeks. No Years or months
//
//         TDurCalcType(0).CumDays()
//           - Computes cumulative days. No Years, months or weeks.
//
//         TDurCalcType(0).CumHours()
//           - Computes cumulative hours. No Years, months, weeks or days.
//
//         TDurCalcType(0).CumMinutes()
//           - Computes cumulative minutes. No Years, months, weeks, days
//             or hours.
//
//         TDurCalcType(0).CumSeconds()
//           - Computes cumulative seconds. No Years, months, weeks, days,
//             hours or minutes.
//
//         TDurCalcType(0).GregorianYears()
//           - Computes Years based on average length of a Gregorian Year
//             Used for very large duration values.
//
//             Type 'TDurCalcType' is located in source file:
//                MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//       - Designates the standard Time Zone location by which
//         time duration will be compared. This ensures that
//         'oranges are compared to oranges and apples are compared
//         to apples' with respect to start time and end time duration
//         calculations.
//
//         If 'timeZoneLocation' is passed as an empty string, it
//         will be automatically defaulted to the 'UTC' time zone.
//         Reference Universal Coordinated Time:
//            https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//         Time zone location, or time zone name,
//         must be designated as one of three types
//         of values:
//
//         (1) The string 'Local' - signals the designation of the local time zone
//             configured for the host computer executing this code.
//
//         (2) IANA Time Zone Location -
//             See https://golang.org/pkg/time/#LoadLocation
//             and https://www.iana.org/time-zones to ensure that
//             the IANA Time Zone Database is properly configured
//             on your system. Note: IANA Time Zone Data base is
//             equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//         (3) A valid Military Time Zone
//             Military time zones are commonly used in
//             aviation as well as at sea. They are also
//             known as nautical or maritime time zones.
//             Reference:
//                 https://en.wikipedia.org/wiki/List_of_military_time_zones
//                 http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                 https://www.timeanddate.com/time/zones/military
//                 https://www.timeanddate.com/worldclock/timezone/alpha
//                 https://www.timeanddate.com/time/map/
//
//         Note:
//             The source file 'timezonedata.go' contains over 600 constant
//             time zone declarations covering all IANA and Military Time
//             Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//             time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode  TimeMathCalcMode
//       - TimeMathCalcMode is an enumeration which specifies the algorithm
//         which will be used when computing time spans or time duration.
//
//         If 'LocalTimeZone' is specified, days are defined as local time
//         zone days which may be less than, or greater than, 24-hours due
//         to local conventions like daylight savings time.
//         (TCalcMode.LocalTimeZone())
//
//         If 'UtcTimeZone' is specified, days are uniformly defined as
//         a time span consisting of 24-consecutive hours.
//         (TCalcMode.UtcTimeZone())
//
//         For additional information see the type documentation at
//               datetime\timemathcalcmodeenum.go
//
//
//   dateTimeFmtStr string
//       - A date time format string which will be used
//         to format and display 'dateTime'. Example:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         Date time format constants are found in the source
//         file 'constantsdatetime.go'. These constants represent
//         the more commonly used date time string formats. All
//         Date Time format constants begin with the prefix
//         'FmtDateTime'.
//
//         If 'dateTimeFmtStr' is submitted as an
//         'empty string', a default date time format
//         string will be applied. The default date time
//         format string is:
//           FmtDateTimeYrMDayFmtStr =
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetStartEndTimes(
//                           startDateTime,
//                           endDateTime,
//                           TZones.US.Central(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TZones.US.Central()' is a constant available int source file,
//         'timezonedata.go'
//
//         TZones.US.Central() is equivalent to "America/Chicago"
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartEndTimes(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartEndTimes() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartEndTimes(
		tDur,
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartEndTimesTz - Sets data field values for the current TimeDurationDto
// instance using a Start Date Time, End Date Time and a time zone specification.
//
// The Starting Date Time and Ending Date Time are submitted as type 'DateTzDto'.
//
// First, 'startDateTimeTz' and 'endDateTimeTz' are converted to the designated
// Time Zone Location. Next, 'startDateTimeTz' is subtracted from 'endDateTime'
// to compute time duration.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTimeTz DateTzDto
//     - Starting date time
//
//  endDateTimeTz   DateTzDto
//     - Ending date time
//
//  tDurCalcType TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time duration
//       calculations.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location, or time zone name,
//       must be designated as one of three types
//       of values:
//
//       (1) The string 'Local' - signals the designation of the local time zone
//           configured for the host computer executing this code.
//
//       (2) IANA Time Zone Location -
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system. Note: IANA Time Zone Data base is
//           equivalent to 'tz database'.
//
//              Examples:
//                "America/New_York"
//                "America/Chicago"
//                "America/Denver"
//                "America/Los_Angeles"
//                "Pacific/Honolulu"
//
//       (3) A valid Military Time Zone
//           Military time zones are commonly used in
//           aviation as well as at sea. They are also
//           known as nautical or maritime time zones.
//           Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode  TimeMathCalcMode
//       - TimeMathCalcMode is an enumeration which specifies the algorithm
//         which will be used when computing time spans or time duration.
//
//         If 'LocalTimeZone' is specified, days are defined as local time
//         zone days which may be less than, or greater than, 24-hours due
//         to local conventions like daylight savings time.
//         (TCalcMode.LocalTimeZone())
//
//         If 'UtcTimeZone' is specified, days are uniformly defined as
//         a time span consisting of 24-consecutive hours.
//         (TCalcMode.UtcTimeZone())
//
//         For additional information see the type documentation at
//               datetime\timemathcalcmodeenum.go
//
//
//   dateTimeFmtStr string
//       - A date time format string which will be used
//         to format and display 'dateTime'. Example:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         Date time format constants are found in the source
//         file 'constantsdatetime.go'. These constants represent
//         the more commonly used date time string formats. All
//         Date Time format constants begin with the prefix
//         'FmtDateTime'.
//
//         If 'dateTimeFmtStr' is submitted as an
//         'empty string', a default date time format
//         string will be applied. The default date time
//         format string is:
//           FmtDateTimeYrMDayFmtStr =
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetStartEndTimesTz(
//                           startDateTimeTz,
//                           endDateTimeTz,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//        standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartEndTimesTz(
	startDateTimeTz,
	endDateTimeTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartEndTimesTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartEndTimesTz(
		tDur,
		startDateTimeTz,
		endDateTimeTz,
		tDurCalcType,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimeDuration - Sets start time, end time and duration for the
// current TimeDurationDto instance. 'startDateTime' is converted to the
// specified by 'timeZoneLocation'. The duration value is added to 'startDateTime'
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date
// time. Thereafter, the actual starting date time is computed by subtracting
// duration.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both the starting date time and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     time.Time
//     - Starting date time for the duration calculation
//
//
//  duration          time.Duration
//     - Time Duration added to 'startDateTime' in order to
//       compute Ending Date Time.
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetStartTimeDuration(
//                           startDateTime,
//                           duration,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           TCalcMode.LocalTimeZone(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartTimeDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimeDuration(
		tDur,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimeTzDuration - Sets start time, end time and duration
// for the current TimeDurationDto instance.
//
// The input parameter, 'startDateTimeTz', is of type DateTzDto. It is
// converted to the specified 'timeZoneLocation', and added to the
// duration value in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTimeTz' is converted to
// ending date time. Thereafter, the actual starting date time is
// computed by subtracting duration.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTimeTz   DateTzDto
//     - Provides starting date time for the duration calculation
//
//
//  duration          time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time duration
//       calculations.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location, or time zone name,
//       must be designated as one of three types
//       of values:
//
//       (1) The string 'Local' - signals the designation of the local time zone
//           configured for the host computer executing this code.
//
//       (2) IANA Time Zone Location -
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system. Note: IANA Time Zone Data base is
//           equivalent to 'tz database'.
//
//              Examples:
//                "America/New_York"
//                "America/Chicago"
//                "America/Denver"
//                "America/Los_Angeles"
//                "Pacific/Honolulu"
//
//       (3) A valid Military Time Zone
//           Military time zones are commonly used in
//           aviation as well as at sea. They are also
//           known as nautical or maritime time zones.
//           Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//   dateTimeFmtStr   string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetStartTimeTzDuration(
//                           startDateTimeTz,
//                           duration,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           TCalcMode.LocalTimeZone(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartTimeTzDuration(
	startDateTimeTz DateTzDto,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartTimeTzDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimeTzDuration(
		tDur,
		startDateTimeTz,
		duration,
		tDurCalcType,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimePlusTimeDto - Sets start date time, end date time and duration
// based on a starting date time. The results of this calculation are used to
// populate the current TimeDurationDto instance.
//
// The time components of the TimeDto are added to the starting date time to
// compute the ending date time and the duration.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     time.Time
//     - Starting date time. The ending date time will be computed
//       by adding the time components of the 'plusTimeDto' to
//       'startDateTime'.
//
//  plusTimeDto       TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be added to 'startDateTime' to compute
//       time duration and ending date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetStartTimePlusTimeDto(
//                           startDateTime,
//                           plusTimeDto,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           TCalcMode.LocalTimeZone(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartTimePlusTimeDto(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartTimePlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimePlusTimeDto(
		tDur,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimeTzPlusTimeDto - Sets start date time, end date time and duration
// based on a starting date time. The results of this calculation are used to
// populate the current TimeDurationDto instance.
//
// The starting date time is passed to this method as an instance of 'DateTzDto'.
//
// The time components of the TimeDto are added to the starting date time to
// compute the ending date time and the duration.
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// The user is also required to provide the time duration calculation type which
// will control the output of the time duration calculation. The standard date
// time calculation type is, 'TDurCalcType(0).StdYearMth()'. This means that
// time duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds. For a
// discussion of Time Duration Calculation type, see Type TDurCalcType located
// in source file:
//
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTimeTz   DateTzDto
//     - Starting date time. The ending date time will be computed
//       by adding the time components of the 'plusTimeDto' to
//       'startDateTime'.
//
//  plusTimeDto       TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be added to 'startDateTime' to compute
//       time duration and ending date time.
//
//       type TimeDto struct {
//          Years                int // Number of Years
//          Months               int // Number of Months
//          Weeks                int // Number of Weeks
//          WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours                int // Number of Hours.
//          Minutes              int // Number of Minutes
//          Seconds              int // Number of Seconds
//          Milliseconds         int // Number of Milliseconds
//          Microseconds         int // Number of Microseconds
//          Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                   //  plus remaining Nanoseconds
//          TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                   //  + Seconds + Milliseconds + Nanoseconds
//       }
//
//       Type 'TimeDto' is located in source file:
//          datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType      TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//  timeCalcMode      TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmodeenum.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto := TimeDurationDto{}
//
//  tDurDto, err := tDurDto.SetStartTimeTzPlusTimeDto(
//                           startDateTimeTz,
//                           plusTimeDto,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           TCalcMode.LocalTimeZone(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'. TZones.US.Central() is equivalent to
//        "America/Chicago".
//
//        TCalcMode.LocalTimeZone() specifies that time duration will be
//        computed in the context of local time zones. Reference Type
//        'TDurCalcType' located in source file:
//            'datetime\timemathcalcmodeenum.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartTimeTzPlusTimeDto(
	startDateTimeTz DateTzDto,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartTimeTzPlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimePlusTimeDto(
		tDur,
		startDateTimeTz.dateTimeValue,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}
