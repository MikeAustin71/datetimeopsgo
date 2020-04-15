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
	TimeDuration    time.Duration // Elapsed time or duration between starting and ending date time
	CalcType        TDurCalcType  // The calculation Type. This controls the allocation of time
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

// EmptyTimeFields - Sets all data fields associated
// with time duration allocation to zero.
//
// __________________________________________________________________________
//
// Return Values:
//
//   -- NONE --
//
func (tDur *TimeDurationDto) EmptyTimeFields() {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	tDurDtoUtil := timeDurationDtoUtility{}

	tDurDtoUtil.emptyTimeFields(
		tDur,
		"TimeDurationDto.EmptyTimeFields() ")

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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	str := fmt.Sprintf("%v-Nanoseconds", int64(tDur.TimeDuration))

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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	return fmt.Sprintf("%v", tDur.TimeDuration)
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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
// 'tDur.TimeDuration'.
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

	if int64(tDur.TimeDuration) == 0 {
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

// GetTypeStartDateTimeSting - Returns the starting date time for
// this duration as a string.
//
func (tDur *TimeDurationDto) GetTypeStartDateTimeSting() string {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.startDateTimeTz.String()
}

// GetTypeStartTimeDateTz - Returns the starting date time for this
// duration as a Type,'DateTzDTo'. This value is stored in private
// member variable TimeDurationDto.startDateTimeTz.
//
func (tDur *TimeDurationDto) GetTypeStartDateTimeTz() DateTzDto {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	return tDur.startDateTimeTz
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	if int64(tDur.TimeDuration) == 0 {
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

	timeDur2.TimeDuration = time.Duration(0)

	timeDur2.CalcType = TDurCalc.None()

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
// Note:   This method applies the standard Time Duration allocation, calculation type
//         'TDurCalcType(0).StdYearMth()'. This means that time duration is allocated over years,
//         months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//         microseconds and nanoseconds.  For details, see Type 'TDurCalcType' in this source
//         file: MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//
//  startDateTime  time.Time
//     - Starting date time
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
//                                     TZones.US.Central(),
//                                     FmtDateTimeYrMDayFmtStr)
//
//   Note:
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
func (tDur TimeDurationDto) NewAutoEnd(
	startDateTime time.Time,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewAutoEnd() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimesCalcTz(
		&tDur2,
		startDateTime,
		time.Now().UTC(),
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewAutoStart - Creates and returns a new TimeDurationDto instance. Starting date time is
// automatically initialized by calling time.Now(). Afterwards, start date time is converted
// to the Time Zone specified in input parameter, 'timeZoneLocation'.
//
// This method will set ending date time to the same value as starting date time resulting in
// a time duration value of zero.
//
// In order to compute the final time duration value, call the method TimeDurationDto.SetAutoEnd()
// when ready.  At that point, ending date time will be set by a call to time.Now().
//
// Use of these two methods, 'NewAutStart' and 'SetAutoEnd', constitutes a stop watch feature which
// can be triggered to measure elapsed time.
//
// Note:  This method applies the standard Time Duration allocation, calculation type
//        'TDurCalcType(0).StdYearMth()'. This means that time duration is allocated over years,
//        months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//        microseconds and nanoseconds. For details, see Type 'TDurCalcType' in this source
//        file: MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
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
//                                      TZones.US.Central(),
//                                      FmtDateTimeYrMDayFmtStr)
//
//   Note:
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
func (tDur TimeDurationDto) NewAutoStart(
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewAutoStart() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	startEndTime := time.Now().UTC()

	err := tDurDtoUtil.setStartEndTimesCalcTz(
		&tDur2,
		startEndTime,
		startEndTime,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndDateTzDto - Creates and returns a new TimeDurationDto populated with
// time duration data. The input parameters for starting and ending date time are
// submitted as type 'DateTzDto'.
//
// The Time Zone Location is extracted from input parameter, 'startDateTz'. The extracted
// time zone is used to configure the ending date time thereby providing a common basis for
// subsequent time duration calculations.
//
// This method automatically applies the standard Time Duration allocation, calculation type
// 'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years, months, weeks,
//  weekdays, date days, hours, minutes, seconds, milliseconds,	microseconds and nanoseconds.
// For details, see Type 'TDurCalcType' in this source file:
//        MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTz DateTzDto
//     - Contains the starting date time used in time duration calculations
//
//  endDateTz   DateTzDto
//     - Contains the ending date time used in time duration calculations.
//
//  dateTimeFmtStr string
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
//  tDurDto, err := TimeDurationDto{}.NewStartEndDateTzDto(
//                                   startDateTz,
//                                   endDateTz,
//                                   FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartEndDateTzDto(
	startDateTz,
	endDateTz DateTzDto,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndDateTzDto() "

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.isValidDateTzDto(
		&startDateTz,
		ePrefix)

	if err != nil {
		return TimeDurationDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "startDateTz",
				inputParameterValue: "",
				errMsg: fmt.Sprintf(
					"Input Parameter 'startDateTz' is INVALID!\n"+
						"Validation Error='%v'", err.Error()),
				err: nil,
			}
	}

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	timeZoneLocation :=
		startDateTz.GetMilitaryOrStdTimeZoneName()

	err = tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
		&tDur2,
		startDateTz,
		endDateTz,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix,
	)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndDateTzDtoCalcTz - Creates and returns a new TimeDurationDto populated with
// time duration data. The input parameters for starting and ending date time are
// submitted as type 'DateTzDto'.
//
// The user is required to specify a Time Zone Location for use in converting date
// times to a common frame of reference used in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. For details see
// Type 'TDurCalcType' which is located in source file:
//       MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTz     DateTzDto
//     - Contains the starting date time used in time duration calculations
//
//  endDateTz       DateTzDto
//     - Contains the ending date time used in time duration calculations.
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
//  tDurDto, err := TimeDurationDto{}.NewStartEndDateTzDtoCalcTz(
//                                    startDateTz,
//                                    endDateTz,
//                                    TDurCalc.StdYearMth(),
//                                    TZones.US.Central(),
//                                    FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartEndDateTzDtoCalcTz(
	startDateTz,
	endDateTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndDateTzDtoCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
		&tDur2,
		startDateTz,
		endDateTz,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndDateTzDtoTz - Creates and returns a new TimeDurationDto populated with
// time duration data. The input parameters for starting and ending date time are
// submitted as type 'DateTzDto'.
//
// The user is required to specify a Time Zone Location for use in converting date
// times to a common frame of reference used in subsequent time duration calculations.
//
// This method automatically applies the standard Time Duration allocation, calculation type
// 'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years, months, weeks,
//  weekdays, date days, hours, minutes, seconds, milliseconds,	microseconds and nanoseconds.
// For details, see Type 'TDurCalcType' in this source file:
//        MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTz  DateTzDto
//     - Contains the starting date time used in time duration calculations
//
//  endDateTz    DateTzDto
//     - Contains the ending date time used in time duration calculations.
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
//  tDurDto, err := TimeDurationDto{}.NewStartEndDateTzDtoTz(
//                       startDateTz,
//                       endDateTz,
//                       TZones.US.Central(),
//                       FmtDateTimeYrMDayFmtStr)
//
//      Note:
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
func (tDur TimeDurationDto) NewStartEndDateTzDtoTz(
	startDateTz,
	endDateTz DateTzDto,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndDateTzDtoTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
		&tDur2,
		startDateTz,
		endDateTz,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndTimes - Creates and returns a new TimeDurationDto based on starting
// and ending date times.
//
// Because, time zone location is crucial to completely accurate
// duration calculations, the time zone of the starting date time,
// 'startDateTime' is applied to parameter, 'endDateTime' before
// making the duration calculation.
//
// This method applies the standard Time Duration allocation,
// 'TDurCalcType(0).StdYearMth()'. This means that duration
// is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and
// nanoseconds. For details, see Type 'TDurCalcType' in this
// source file:
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime time.Time
//     - Starting date time
//
//  endDateTime   time.Time
//     - Ending date time
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
//                                   FmtDateTimeYrMDayFmtStr)
//
//  Note: FmtDateTimeYrMDayFmtStr' is a constant available in source file:
//        'constantsdatetime.go'
//
func (tDur TimeDurationDto) NewStartEndTimes(
	startDateTime,
	endDateTime time.Time,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimes() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimesCalcTz(
		&tDur2,
		startDateTime,
		endDateTime,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndTimesTz - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
// The user is required to specify a common Time Zone Location for use in converting
// date times to a common frame of reference to subsequent time duration calculations.
//
// Note:    This method applies the standard Time Duration allocation, calculation type
//          'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years,
//          months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//          microseconds and nanoseconds.   See Type 'TDurCalcType' for details.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime    time.Time
//     - Starting time
//
//  endDateTime      time.Time
//     - Ending time
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
//                                    startTime,
//                                    endTime,
//                                    TZones.US.Central(),
//                                    FmtDateTimeYrMDayFmtStr)
//
//      Note:
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
func (tDur TimeDurationDto) NewStartEndTimesTz(
	startDateTime,
	endDateTime time.Time,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimesCalcTz(
		&tDur2,
		startDateTime,
		endDateTime,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndTimesCalc - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// The user is required to submit input parameters for date time calculation type. The
// Time Zone Location is extracted from the 'startDateTime' parameter
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime    time.Time
//     - Starting time
//
//  endDateTime      time.Time
//     - Ending time
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
//  tDurDto, err := TimeDurationDto{}.NewStartEndTimesCalcTz(
//                    startTime,
//                    endTime,
//                    TDurCalcType(0).StdYearMth(),
//                    FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartEndTimesCalc(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesCalc() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartEndTimesCalcTz(
		&tDur2,
		startDateTime,
		endDateTime,
		tDurCalcType,
		startDateTime.Location().String(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndTimesCalcTz - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
//
// The user is required to specify a common Time Zone Location for use in converting
// date times to a common frame of reference for use in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. For details see
// Type 'TDurCalcType' which is located in source file:
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   time.Time
//     - Starting time
//
//  endDateTime     time.Time
//     - Ending time
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
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesCalcTz(
//                    startTime,
//                    endTime,
//                    TDurCalcType(0).StdYearMth(),
//                    TZones.US.Central(),
//                    FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartEndTimesCalcTz(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesCalcTz() "

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartEndTimesCalcTz(
		&tDur2,
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndTimesDateDto - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters. The
// 'startDateTime' and 'endDateTime' parameters are of type DateTzDto.
//
// Time Zone Location is derived from input parameter 'startDateTime'. This time zone
// provides a common frame of reference for use in subsequent time duration calculations.
//
// Note: This method applies the standard Time Duration allocation, calculation type
//       'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years,
//       months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//       microseconds and nanoseconds.   See Type 'TDurCalcType' for details.
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
//  tDurDto, err := TimeDurationDto{}.NewStartEndTimesDateDto(
//                                    startTime,
//                                    endTime,
//                                    FmtDateTimeYrMDayFmtStr)
//
//      Note:
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartEndTimesDateDto(
	startDateTimeTz,
	endDateTimeTz DateTzDto,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesDateDto() "

	timeZoneLocation :=
		startDateTimeTz.GetMilitaryOrStdTimeZoneName()

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
		&tDur2,
		startDateTimeTz,
		endDateTimeTz,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndTimesDateDtoCalc - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters. The
// 'startDateTime' and 'endDateTime' parameters are of type DateTzDto.
//
// Time Zone Location is derived from input parameter 'startDateTime'. If the 'endDateTime'
// time zone is NOT equivalent to 'startDateTime', the 'endDateTime' will be converted to
// the time zone provided by the 'startDateTime' parameter. This will provide a common basis
// for use in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Type 'TDurCalcType' is located in source file:
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
//  tDurDto, err := TimeDurationDto{}.NewStartEndTimesDateDtoCalc(
//                            startTime,
//                            endTime,
//                            TDurCalcType(0).StdYearMth(),
//                            FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartEndTimesDateDtoCalc(
	startDateTimeTz,
	endDateTimeTz DateTzDto,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesDateDtoCalc() "

	timeZoneLocation :=
		startDateTimeTz.GetMilitaryOrStdTimeZoneName()

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
		&tDur2,
		startDateTimeTz,
		endDateTimeTz,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartEndTimesDateTzDtoCalcTz - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters. The
// 'startDateTime' and 'endDateTime' parameters are of type DateTzDto.
//
// The user is required to specify a specific Time Zone Location used to convert date times
// to a common frame of reference for use in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Type 'TDurCalcType' is located in source file:
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   DateTzDto
//     - Starting date time
//
//  endDateTime     DateTzDto
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
//  tDurDto, err := TimeDurationDto{}.NewStartEndTimesDateTzDtoCalc(
//                                   startTime,
//                                   endTime,
//                                   TZones.US.Central(),
//                                   TDurCalcType(0).StdYearMth(),
//                                   FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartEndTimesDateTzDtoCalcTz(
	startDateTz,
	endDateTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartEndTimesDateTzDtoCalcTz() "

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
		&tDur2,
		startDateTz,
		endDateTz,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationTz - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime' and time duration. 'startDateTime' is used to derive Time Zone Location.
// The time duration value is added to 'startDateTime' in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time. The
// actual starting date time is computed by subtracting duration from ending date time.
//
// Note:  This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//        This means that duration is allocated over years, months, weeks, weekdays, date days,
//        hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//        See Type 'TDurCalcType' for details.
//
//        Type 'TDurCalcType' is located in source file:
//          MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime time.Time
//     - Starting date time for the duration calculation.
//       Note: The Time Zone extracted from 'startDateTime'
//       is used in calculating both starting date time and
//       ending date time.
//
//  duration  time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
//  tDurDto, err := TimeDurationDto{}.NewStartTimeDurationTz(
//                                     startTime,
//                                     duration,
//                                     FmtDateTimeYrMDayFmtStr)
//   Note:
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDuration() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDurationCalcTz(
		&tDur2,
		startDateTime,
		duration,
		TDurCalc.StdYearMth(),
		startDateTime.Location().String(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationTz - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime', time duration and 'timeZoneLocation'. 'startDateTime' is converted to the
// specified 'timeZoneLocation'. The duration value is added to 'startDateTime' in order to
// compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and the
//  actual starting date time is computed by subtracting duration.
//
// Note: This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//       This means that duration is allocated over years, months, weeks, weekdays, date days,
//       hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//       See Type 'TDurCalcType' for details.
//
//       Type 'TDurCalcType' is located in source file:
//        MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime time.Time
//     - Starting date time for the duration calculation
//
//  duration  time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
//  tDurDto, err := TimeDurationDto{}.NewStartTimeDurationTz(
//                                    startTime,
//                                    duration,
//                                    TZones.US.Central(),
//                                    FmtDateTimeYrMDayFmtStr)
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
func (tDur TimeDurationDto) NewStartTimeDurationTz(
	startDateTime time.Time,
	duration time.Duration,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDurationCalcTz(
		&tDur2,
		startDateTime,
		duration,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationCalcTz - Creates and returns a new TimeDurationDto based on input
// parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation type.
// 'startDateTime' is converted to the specified 'timeZoneLocation' and the duration value
// is added to it in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and
// the actual starting date time is computed by subtracting duration.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Type 'TDurCalcType' is located in source file:
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   time.Time
//     - Starting date time for the duration calculation
//
//  duration    time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
//  tDurDto, err := TimeDurationDto{}.NewStartTimeDurationCalcTz(
//                                   startTime,
//                                   duration,
//                                   TZones.US.Central(),
//                                   TDurCalcType(0).StdYearMth(),
//                                   FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartTimeDurationCalcTz(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationCalcTz() "

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartTimeDurationCalcTz(
		&tDur2,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationCalc - Creates and returns a new TimeDurationDto based on input
// parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation type.
//
// The duration value is added to 'startDateTime' in order to compute the ending date time.
//
// If 'duration' is a negative value, 'startDateTime' is converted to ending date time and
// the actual starting date time is computed by subtracting duration.
//
// The time zone location applied to both 'startDateTime' and ending date time is derived
// from input parameter, 'startDateTime'.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Type 'TDurCalcType' is located in source file:
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   time.Time
//     - Starting date time for the duration calculation
//
//  duration    time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
//  tDurDto, err := TimeDurationDto{}.NewStartTimeDurationCalc(
//                                   startTime,
//                                   duration,
//                                   TDurCalc.StdYearMth(),
//                                   FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartTimeDurationCalc(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationCalc() "

	timeZoneLocation := startDateTime.Location().String()

	t2Dur := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartTimeDurationCalcTz(
		&t2Dur,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, nil
	}

	return t2Dur, nil
}

// NewStartTimeDurationDateDto - Creates and returns a new TimeDurationDto based on input
// parameters 'startDateTime' and time duration. 'startDateTime' is of type DateTzDto.
//
// The time duration value is added to 'startDateTime' in order to compute the ending date time.
// If 'duration' is a negative value, 'startDateTime' is converted to ending date time. The
// actual starting date time is computed by subtracting duration from ending date time.
//
// Time Zone location is derived from 'startDateTime'.
//
// Note: This method applies the standard Time Duration allocation calculation type,
//       'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years,
//       months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
//       microseconds and nanoseconds. See Type 'TDurCalcType' for details.
//
//       Type 'TDurCalcType' is located in source file:
//        MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime DateTzDto
//     - Starting date time for the duration calculation
//
//  duration  time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
//  tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDto(
//                   startTime,
//                   duration,
//                   FmtDateTimeYrMDayFmtStr)
//
// Note:
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartTimeDurationDateDto(
	startDateTimeTz DateTzDto,
	duration time.Duration,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDto() "

	timeZoneLocation :=
		startDateTimeTz.GetMilitaryOrStdTimeZoneName()

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
		&tDur2,
		startDateTimeTz,
		duration,
		TDurCalc.StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationDateDtoTz - Creates and returns a new TimeDurationDto based on input
// parameters 'startDateTime', time duration and 'timeZoneLocation'. 'startDateTime' is
// converted to the specified 'timeZoneLocation'. The 'duration' value is added to 'startDateTime'
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and the
// actual starting date time is computed by subtracting duration.
//
// The user is required to specify a common Time Zone Location for use in converting
// date times to a common frame of reference for use in subsequent time duration calculations.
//
// Note: This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//       This means that duration is allocated over years, months, weeks, weekdays, date days,
//       hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//       See Type 'TDurCalcType' for details.
//
//       Type 'TDurCalcType' is located in source file:
//         MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime time.Time
//     - Starting date time
//
//  duration  time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDtoTz(startTime, duration,
//                   TZones.US.Central(), FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartTimeDurationDateDtoTz(
	startDateTime DateTzDto,
	duration time.Duration,
	timeZoneLocation string,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDtoTz() "

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
		&tDur2,
		startDateTime,
		duration,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewStartTimeDurationDateDtoTzCalc - Creates and returns a new TimeDurationDto based
// on input parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation
// type.
//
// Input parameter, 'startDateTime' is of Type DateTzDto.
//
// 'startDateTime' is converted to the specified 'timeZoneLocation' and the duration
// value is added to it in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and
// the actual starting date time is computed by subtracting duration.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Type 'TDurCalcType' is located in source file:
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   DateTzDto
//     - Starting date time for the duration calculation
//
//  duration    time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
//  tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDtoTzCalc(
//                                   startTime,
//                                   duration,
//                                   TZones.US.Central(),
//                                   TDurCalcType(0).StdYearMth(),
//                                   FmtDateTimeYrMDayFmtStr)
//
//
// Note:
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartTimeDurationDateDtoTzCalc(
	startDateTime DateTzDto,
	duration time.Duration,
	timeZoneLocation string,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDtoTzCalc() "

	tDur2 := TimeDurationDto{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
		&tDur2,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, nil
	}

	return tDur2, nil
}

// NewStartTimeDurationDateDtoCalc - Creates and returns a new TimeDurationDto based on input
// parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation type.
//
// Input parameter 'startDateTime' is of Type DateTzDto.
//
// The duration value will be added to 'startDateTime' in order to compute the ending date time.
//
// If 'duration' is a negative value, 'startDateTime' is converted to ending date time and
// the actual starting date time will be computed by subtracting duration.
//
// The time zone location applied to both 'startDateTime' and ending date time is derived
// from input parameter, 'startDateTime'.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Type 'TDurCalcType' is located in source file:
//   MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime      DateTz
//     - Starting date time for the duration calculation
//
//  duration    time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
//  tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDtoCalc(startTime,
//            duration,
//             TDurCalcType(0).StdYearMth(),
//              FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartTimeDurationDateDtoCalc(
	startDateTime DateTzDto,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDtoCalc() "

	tDurDtoUtil := timeDurationDtoUtility{}

	timeZone := startDateTime.GetBestConvertibleTimeZone()

	timeZoneLocation := timeZone.GetMilitaryOrStdTimeZoneName()

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
		&tDur2,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
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
// the ending date time and the duration.
//
// For the purposes of this time duration calculation, the Time Zone Location is
// extracted from the input parameter, 'startDateTime'.
//
// Note: This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//       This means that duration is allocated over years, months, weeks, weekdays, date days,
//       hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//       See Type 'TDurCalcType' for details.
//
//       Type 'TDurCalcType' is located in source file:
//        MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  timeCalcMode    TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the
//       addition algorithm which will be used when adding time
//       components to the current DateTzDto date time value.
//
//       If days are defined as local time zone days (which may be
//       less than or greater than 24-hours) use TCalcMode.LocalTimeZone().
//
//       If days are always defined as having a time span of 24-consecutive
//       hours, use TCalcMode.UtcTimeZone().
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmode.go
//
//       Valid values are:
//             TCalcMode.LocalTimeZone()
//             TCalcMode.UtcTimeZone()
//
//  startDateTime   time.Time
//     - Starting date time. The ending date time will be computed
//       by adding the time components of the 'plusTimeDto' to
//       'startDateTime'.
//
//  plusTimeDto     TimeDto
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
//                    TCalcMode.LocalTimeZone(),
//                    startTime,
//                    plusTimeDto,
//                    FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TCalcMode.LocalTimeZone()' is of type 'TimeMathCalcMode'.
//        Reference 'timeCalcMode' input parameter documentation above
//        and source code documentation at:
//            datetime\timemathcalcmode.go
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewStartTimePlusTimeDto(
	timeCalcMode TimeMathCalcMode,
	startDateTime time.Time,
	plusTimeDto TimeDto,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimePlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimePlusTimeDtoCalcTz(
		&tDur2,
		timeCalcMode,
		startDateTime,
		plusTimeDto,
		TDurCalcType(0).StdYearMth(),
		startDateTime.Location().String(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewStartTimePlusTimeDtoCalcTz - Creates and returns a new TimeDurationDto setting
// the start date time, end date time and duration based on a starting date time
// and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// The user is required to submit input parameters for time zone location and
// date time calculation type.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  timeCalcMode  TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the
//       addition algorithm which will be used when adding time
//       components to the current DateTzDto date time value.
//
//       If days are defined as local time zone days (which may be
//       less than or greater than 24-hours) use TCalcMode.LocalTimeZone().
//
//       If days are always defined as having a time span of 24-consecutive
//       hours, use TCalcMode.UtcTimeZone().
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmode.go
//
//       Valid values are:
//             TCalcMode.LocalTimeZone()
//             TCalcMode.UtcTimeZone()
//
//  startDateTime   time.Time
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
//  tDurDto, err := TimeDurationDto{}.NewStartTimePlusTimeDtoCalcTz(
//                    TCalcMode.LocalTimeZone(),
//                    startTime,
//                    plusTimeDto,
//                    TDurCalcType(0).StdYearMth(),
//                    TZones.US.Central(),
//                    FmtDateTimeYrMDayFmtStr)
//
// Note:
//        'TCalcMode.LocalTimeZone()' is of type 'TimeMathCalcMode'.
//        Reference 'timeCalcMode' input parameter documentation above
//        and source code documentation at:
//            datetime\timemathcalcmode.go
//
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewStartTimePlusTimeDtoCalcTz(
	timeCalcMode TimeMathCalcMode,
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewStartTimePlusTimeDto() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setStartTimePlusTimeDtoCalcTz(
		&tDur2,
		timeCalcMode,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, nil
}

// NewEndTimeMinusTimeDtoTz - Creates and returns a new TimeDurationDto setting
// start date time, end date time and duration based on an ending date time
// and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// The Time Zone Location used in the subsequent date time duration calculations
// is extracted from input parameter, 'endDateTime'. This Time Zone Location is
// applied to both the starting and ending date times.
//
// Note: This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
//       This means that duration is allocated over years, months, weeks, weekdays, date days,
//       hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//       See Type 'TDurCalcType' for details.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  timeCalcMode  TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the
//       addition algorithm which will be used when adding time
//       components to the current DateTzDto date time value.
//
//       If days are defined as local time zone days (which may be
//       less than or greater than 24-hours) use TCalcMode.LocalTimeZone().
//
//       If days are always defined as having a time span of 24-consecutive
//       hours, use TCalcMode.UtcTimeZone().
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmode.go
//
//       Valid values are:
//             TCalcMode.LocalTimeZone()
//             TCalcMode.UtcTimeZone()
//
//  endDateTime   time.Time
//     - Ending date time. The starting date time will be computed
//       by subtracting minusTimeDto from 'endDateTime'
//
//  minusTimeDto  TimeDto
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
//  tDurDto, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoTz(
//                           TCalcMode.LocalTimeZone(),
//                           endTime,
//                           minusTimeDto,
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TCalcMode.LocalTimeZone()' is of type 'TimeMathCalcMode'.
//        Reference 'timeCalcMode' input parameter documentation above
//        and source code documentation at:
//            datetime\timemathcalcmode.go
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur TimeDurationDto) NewEndTimeMinusTimeDto(
	timeCalcMode TimeMathCalcMode,
	endDateTime time.Time,
	minusTimeDto TimeDto,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewEndTimeMinusTimeDtoTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}

	err := tDurDtoUtil.setEndTimeMinusTimeDtoCalcTz(
		&tDur2,
		timeCalcMode,
		endDateTime,
		minusTimeDto,
		TDurCalcType(0).StdYearMth(),
		endDateTime.Location().String(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeDurationDto{}, err
	}

	return tDur2, err
}

// NewEndTimeMinusTimeDtoCalcTz - Creates and returns a new TimeDurationDto setting
// start date time, end date time and duration based on an ending date time
// and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// The user is required to submit input parameters for time zone location and
// date time calculation type.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  timeCalcMode    TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the
//       addition algorithm which will be used when adding time
//       components to the current DateTzDto date time value.
//
//       If days are defined as local time zone days (which may be
//       less than or greater than 24-hours) use TCalcMode.LocalTimeZone().
//
//       If days are always defined as having a time span of 24-consecutive
//       hours, use TCalcMode.UtcTimeZone().
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmode.go
//
//       Valid values are:
//             TCalcMode.LocalTimeZone()
//             TCalcMode.UtcTimeZone()
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
//  tDurDto, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoCalcTz(endTime,
//                    TCalcMode.LocalTimeZone(),
//                    minusTimeDto,
//                    TDurCalcType(0).StdYearMth(),
//                    TZones.US.Central()
//                    FmtDateTimeYrMDayFmtStr)
//
// Note:
//         'TCalcMode.LocalTimeZone()' is of type 'TimeMathCalcMode'.
//         Reference source code documentation at:
//            datetime\timemathcalcmode.go
//
//         'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//         standard year month day time duration allocation.
//
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
func (tDur TimeDurationDto) NewEndTimeMinusTimeDtoCalcTz(
	timeCalcMode TimeMathCalcMode,
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.NewEndTimeMinusTimeDtoTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	t2Dur := TimeDurationDto{}

	err := tDurDtoUtil.setEndTimeMinusTimeDtoCalcTz(
		&t2Dur,
		timeCalcMode,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
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
// the existing calculation type, 'tDur.CalcType'.
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

	return tDurDtoUtil.reCalcEndDateTimeToNow(tDur, ePrefix)
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
		ePrefix)
}

// SetEndTimeMinusTimeDtoCalcTz - Sets start date time, end date time and duration
// based on an ending date time, and the time components contained in a TimeDto
// object.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  timeCalcMode    TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the
//       addition algorithm which will be used when adding time
//       components to the current DateTzDto date time value.
//
//       If days are defined as local time zone days (which may be
//       less than or greater than 24-hours) use TCalcMode.LocalTimeZone().
//
//       If days are always defined as having a time span of 24-consecutive
//       hours, use TCalcMode.UtcTimeZone().
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmode.go
//
//       Valid values are:
//             TCalcMode.LocalTimeZone()
//             TCalcMode.UtcTimeZone()
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
//  tDurDto, err := tDurDto.SetEndTimeMinusTimeDtoCalcTz(
//                           TCalcMode.LocalTimeZone(),
//                           endTime,
//                           minusTimeDto,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TCalcMode.LocalTimeZone()' is of type 'TimeMathCalcMode'.
//        Reference 'timeCalcMode' input parameter documentation above
//        and source code documentation at:
//            datetime\timemathcalcmode.go
//
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
func (tDur *TimeDurationDto) SetEndTimeMinusTimeDtoCalcTz(
	timeCalcMode TimeMathCalcMode,
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetEndTimeMinusTimeDtoCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setEndTimeMinusTimeDtoCalcTz(
		tDur,
		timeCalcMode,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartEndTimesDateDtoCalcTz - Sets data field values for the current
// TimeDurationDto instance using a Start Date Time, End Date Time and a
// time zone specification.
//
// The Starting Date Time and Ending Date Time are submitted as type 'DateTzDto'
//
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
//
// The user is required to submit input parameters for time zone location and
// date time calculation type.
//
// All data fields in the current TimeDurationDto instance are overwritten with
// the new time duration values.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   DateTzDto
//     - Starting date time
//
//  endDateTime     DateTzDto
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
//  tDurDto, err := tDurDto.SetStartEndTimesDateDtoCalcTz(
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
func (tDur *TimeDurationDto) SetStartEndTimesDateDtoCalcTz(
	startDateTimeTz,
	endDateTimeTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartEndTimesDateDtoCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartEndTimesDateDtoCalcTz(
		tDur,
		startDateTimeTz,
		endDateTimeTz,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartEndTimesCalcTz - Sets data field values for the current TimeDurationDto
// instance using a Start Date Time, End Date Time and a time zone specification.
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
//
// All data fields in the current TimeDurationDto instance are overwritten with
// the new time duration values.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   time.Time
//     - Starting time
//
//  endDateTime     time.Time
//     - Ending time
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
//  tDurDto, err := tDurDto.SetStartEndTimesCalcTz(
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
func (tDur *TimeDurationDto) SetStartEndTimesCalcTz(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartEndTimesCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartEndTimesCalcTz(
		tDur,
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimeDurationCalcTz - Sets start time, end time and duration for the
// current TimeDurationDto instance. 'startDateTime' is converted to the
// specified by 'timeZoneLocation'. The duration value is added to 'startDateTime'
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date
// time. Thereafter, the actual starting date time is computed by subtracting
// duration.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime     time.Time
//     - Starting date time for the duration calculation
//
//  duration          time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
//  tDurDto, err := tDurDto.SetStartTimeDurationCalcTz(
//                           startDateTime,
//                           duration,
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
func (tDur *TimeDurationDto) SetStartTimeDurationCalcTz(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartTimeDurationCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimeDurationCalcTz(
		tDur,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimeDurationDateDtoCalcTz - Sets start time, end time and
// duration for the current TimeDurationDto instance.
//
// The input parameter, 'startDateTime', is of type DateTzDto. It is
// converted to the specified 'timeZoneLocation', and added to the
// duration value in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to
// ending date time. Thereafter, the actual starting date time is
// computed by subtracting duration.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  startDateTime   DateTzDto
//     - Provides starting date time for the duration calculation
//
//  duration    time.Duration
//     - Amount of time to be added to or subtracted from
//       'startDateTime'. Note: If duration is a negative value
//       'startDateTime' is converted to ending date time and
//       actual starting date time is computed by subtracting
//       duration.
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
//  tDurDto, err := tDurDto.SetStartTimeDurationDateDtoCalcTz(
//                           startDateTimeTz,
//                           duration,
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
func (tDur *TimeDurationDto) SetStartTimeDurationDateDtoCalcTz(
	startDateTimeTz DateTzDto,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartTimeDurationDateDtoCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimeDurationDateDtoCalcTz(
		tDur,
		startDateTimeTz,
		duration,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimePlusTimeDtoCalcTz - Sets start date time, end date time and duration
// based on a starting date time.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  timeCalcMode    TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the
//       addition algorithm which will be used when adding time
//       components to the current DateTzDto date time value.
//
//       If days are defined as local time zone days (which may be
//       less than or greater than 24-hours) use TCalcMode.LocalTimeZone().
//
//       If days are always defined as having a time span of 24-consecutive
//       hours, use TCalcMode.UtcTimeZone().
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmode.go
//
//       Valid values are:
//             TCalcMode.LocalTimeZone()
//             TCalcMode.UtcTimeZone()
//
//  startDateTime   time.Time
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
//  tDurDto, err := tDurDto.SetStartTimePlusTimeDtoCalcTz(
//                           TCalcMode.LocalTimeZone(),
//                           startDateTime,
//                           plusTimeDto,
//                           TDurCalc.StdYearMth(),
//                           TZones.US.Central(),
//                           FmtDateTimeYrMDayFmtStr)
//
//  Note:
//        'TCalcMode.LocalTimeZone()' is of type 'TimeMathCalcMode'.
//        Reference 'timeCalcMode' input parameter documentation above
//        and source code documentation at:
//            datetime\timemathcalcmode.go
//
//        'TDurCalc.StdYearMth()' is of type 'TDurCalcType' and signals
//        standard year month day time duration allocation.
//
//        'TZones.US.Central()' is a constant available int source file,
//        'timezonedata.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'
//
//        FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartTimePlusTimeDtoCalcTz(
	timeCalcMode TimeMathCalcMode,
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.lock.Lock()

	defer tDur.lock.Unlock()

	ePrefix := "TimeDurationDto.SetStartTimePlusTimeDtoCalcTz() "

	tDurDtoUtil := timeDurationDtoUtility{}

	return tDurDtoUtil.setStartTimePlusTimeDtoCalcTz(
		tDur,
		timeCalcMode,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)
}
