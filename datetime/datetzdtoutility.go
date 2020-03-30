package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type dateTzDtoUtility struct {
	input  string
	output string
	lock   sync.Mutex
}

// addDate - Adds input parameters 'years, 'months' and 'days' to date time value
// of the DateTzDto input parameter and returns the updated value in a new
// 'DateTzDto' instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dTz *DateTzDto - Provides the base date to which input parameters 'years',
//                  'months' and 'days' are added.
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
//  years             int   - Number of years to add to the current date.
//  months            int   - Number of months to add to the current date.
//  days              int   - Number of days to add to the current date.
//
//          Note: Date Component input parameters may be either negative
//                or positive. Negative values will subtract time from
//                the current DateTzDto instance.
//
//  dateTimeFmtStr string   - A date time format string which will be used
//                            to format and display 'dateTime'. Example:
//                            "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                            Date time format constants are found in the source
//                            file 'constantsdatetime.go'. These constants represent
//                            the more commonly used date time string formats. All
//                            Date Time format constants begin with the prefix
//                            'FmtDateTime'.
//
//                            If 'dateTimeFmtStr' is submitted as an
//                            'empty string', a default date time format
//                            string will be applied. The default date time
//                            format string is:
//                              FmtDateTimeYrMDayFmtStr =
//                                  "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  ePrefix        string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a new, valid, fully populated
//              DateTzDto type updated to reflect the added input parameters,
//              years, months and days.
//
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) addDate(
	dTz *DateTzDto,
	timeCalcMode TimeMathCalcMode,
	years,
	months,
	days int,
	dateTimeFormatStr,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addDate() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	dTzUtil2 := dateTzDtoUtility{}

	err := dTzUtil2.isValidDateTzDto(dTz, ePrefix)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix +
			"\nInput parameter 'dTz' is Invalid!\n" +
			"Validation Error='%v'\n", err.Error())
	}

	var newDt1 time.Time
	
	if timeCalcMode == TCalcMode.LocalTimeZone() {

		newDt1 = dTz.dateTimeValue.AddDate(
			years,
			months,
			days)

	} else if timeCalcMode == TCalcMode.UtcTimeZone() {

		dtMech := DTimeMechanics{}

		newDt1 = dtMech.AddDateTimeByUtc(
			dTz.dateTimeValue,
			years,
			months,
			days,
			0,
			0,
			0,
			0,
			0,
			0)

	} else {
		return DateTzDto{}, 
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeCalcMode",
				inputParameterValue: "",
				errMsg:              "Input parameter " +
					"'timeCalcMode' is invalid!",
				err:                 nil,
			}
	}
	

	if dateTimeFormatStr == "" {
		dateTimeFormatStr = dTz.dateTimeFmt
	}

	dtz2 := DateTzDto{}

	err = dTzUtil2.setFromDateTime( &dtz2, newDt1, dateTimeFormatStr, ePrefix)

	return dtz2, err
}

// addDateTime - Adds date time components to the date time value of the
// current DateTzDto instance. The updated date time value is returned to
// the calling function as a new DateTzDto instance.
//
// Note that the input parameter 'timeCalcMode' determines whether the
// addition operation will be performed using the local time zone
// or Universal Coordinated Time. For more information on Time Calculation
// Mode, see the type documentation for 'TimeMathCalcMode'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dTz *DateTzDto
//     - Provides the base date to which input parameters 'years',
//       'months', 'days', 'hours', 'minutes', 'seconds',
//       'milliseconds', 'microseconds' and 'nanoseconds' are added.
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
//  years             int   - Number of years to add to the value of parameter 'dTz'.
//  months            int   - Number of months to add to the value of parameter 'dTz'.
//  days              int   - Number of days to add to the value of parameter 'dTz'.
//  hours             int   - Number of hours to add to the value of parameter 'dTz'.
//  minutes           int   - Number of minutes to add to the value of parameter 'dTz'.
//  seconds           int   - Number of seconds to add to the value of parameter 'dTz'.
//  milliseconds      int   - Number of milliseconds to add to the value of parameter 'dTz'.
//  microseconds      int   - Number of microseconds to add to the value of parameter 'dTz'.
//  nanoseconds       int   - Number of nanoseconds to add to the value of parameter 'dTz'.
//
//          Note: Date Component input parameters may be either negative
//                or positive. Negative values will subtract time from
//                the current DateTzDto instance.
//
//  ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a new, valid, fully populated
//              DateTzDto type updated to reflect the added input parameters,
//              years, months, days, hours, minutes, seconds, milliseconds,
//              microseconds and nanoseconds.
//
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) addDateTime(
	dTz *DateTzDto,
	timeCalcMode TimeMathCalcMode,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addDateTime() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	var newDateTime time.Time
	dtMech := DTimeMechanics{}

	if timeCalcMode == TCalcMode.LocalTimeZone() {

		newDateTime = dtMech.AddDateTimeByLocalTimeZone(
										dTz.dateTimeValue,
										years,
										months,
										days,
										hours,
										minutes,
										seconds,
										milliseconds,
										microseconds,
										nanoseconds)

	} else if timeCalcMode == TCalcMode.UtcTimeZone() {

		newDateTime = dtMech.AddDateTimeByUtc(
			dTz.dateTimeValue,
			years,
			months,
			days,
			hours,
			minutes,
			seconds,
			milliseconds,
			microseconds,
			nanoseconds)

	} else {
		return DateTzDto{},
		&InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeCalcMode",
			inputParameterValue: "",
			errMsg:              "Input parameter 'timeCalcMode' " +
				"is invalid!",
			err:                 nil,
		}
	}

	dTzUtil2 := dateTzDtoUtility{}

	dTz2 := DateTzDto{}

	err := dTzUtil2.setFromDateTime(&dTz2, newDateTime, dTz.dateTimeFmt, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dTz2, nil
}

// addDuration - Adds Duration to the DateTime XValue of the input
// parameter 'dTz' (DateTzDto) and returns a new DateTzDto instance
// with the updated Date Time value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dTz *DateTzDto - Provides the base date to which input parameters 'years',
//                  'months' and 'days' are added.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dTz         *DateTzDto
//     - Provides the base date to which input parameters 'years',
//       'months' and 'days' are added.
//
//  duration time.Duration
//     - Time duration value to be added to parameter 'dTz'.
//
//  dateTimeFmtStr  string
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
//  ePrefix        string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a new, valid, fully populated
//              DateTzDto type updated to reflect the added input parameters,
//              years, months and days.
//
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) addDuration(
	dTz *DateTzDto,
	duration time.Duration,
	dateTimeFmtStr,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addDuration() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	newDateTime := dTz.dateTimeValue.Add(duration)

	dTzUtil2 := dateTzDtoUtility{}

	dtz2 := DateTzDto{}

	err := dTzUtil2.setFromDateTime(&dtz2, newDateTime, dateTimeFmtStr, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// addMinusTimeDto - Creates and returns a new DateTzDto by
// subtracting a TimeDto from the value of the input
// parameter 'dTz' (DateTzDto) instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dTz                   *DateTzDto
//     - Provides the base date from which input parameters
//      'minusTimeDto' will be subtracted.
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
//  minusTimeDto             TimeDto
//     - The value of this 'TimeDto' will be subtracted from parameter
//       'dTz'.
//
//  ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a new, valid, fully populated
//              DateTzDto type incorporating the time value calculated by subtracting
//              'minusTimeDto' from 'dTz'.
//
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) addMinusTimeDto(
	dTz *DateTzDto,
	timeCalcMode TimeMathCalcMode,
	minusTimeDto TimeDto,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addMinusTimeDto() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	tDto := minusTimeDto.CopyOut()

	err := tDto.NormalizeTimeElements()

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"\nError returned by tDto.NormalizeTimeElements().\n"+
				"Error='%v'\n", err.Error())
	}

	_, err = tDto.NormalizeDays()

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"\nError returned by tDto.NormalizeDays().\n"+
				"\nError='%v'\n", err.Error())
	}

	tDto.ConvertToNegativeValues()

	var dt2 time.Time

	dtMech := DTimeMechanics{}

	if timeCalcMode == TCalcMode.LocalTimeZone() {

		dt2 = dtMech.AddDateTimeByLocalTimeZone(
			dTz.dateTimeValue,
			tDto.Years,
			tDto.Months,
			tDto.DateDays,
			tDto.Hours,
			tDto.Minutes,
			tDto.Seconds,
			tDto.Milliseconds,
			tDto.Microseconds,
			tDto.Nanoseconds)

	} else if timeCalcMode == TCalcMode.UtcTimeZone() {

		dt2 = dtMech.AddDateTimeByUtc(
			dTz.dateTimeValue,
			tDto.Years,
			tDto.Months,
			tDto.DateDays,
			tDto.Hours,
			tDto.Minutes,
			tDto.Seconds,
			tDto.Milliseconds,
			tDto.Microseconds,
			tDto.Nanoseconds)

	} else {
		
		return DateTzDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeCalcMode",
				inputParameterValue: "",
				errMsg:              "",
				err:                 nil,
			}

	}

	dtz2 := DateTzDto{}
	dTzUtil2 := dateTzDtoUtility{}

	err = dTzUtil2.setFromTzDef(
		&dtz2,
		dt2,
		dTz.timeZone.CopyOut(),
		TzConvertType.Relative(),
		dTz.dateTimeFmt,
		ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// addPlusTimeDto - Creates and returns a new DateTzDto by adding a TimeDto
// to the value of theDateTzDto instance passed as an input parameter.
//
// The value of the input parameter DateTzDto instance is not be altered.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dTz                   *DateTzDto
//     - Provides the base date to which input parameter 'plusTimeDto'
//       will be added.
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
//  plusTimeDto              TimeDto
//     - The time duration value represented by 'plusTimeDto' will
//       be added to parameter 'dTz'.
//
//  ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a new, valid, fully populated
//              DateTzDto containing the sum of time values for 'dTz' and
//              parameter 'plusTimeDto'.
//
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) addPlusTimeDto(
	dTz *DateTzDto,
	timeCalcMode TimeMathCalcMode,
	plusTimeDto TimeDto,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addPlusTimeDto() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	tDto := plusTimeDto.CopyOut()

	err := tDto.NormalizeTimeElements()

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"\nError returned by tDto.NormalizeTimeElements().\n"+
				"\nError='%v'\n", err.Error())
	}

	_, err = tDto.NormalizeDays()

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"\nError returned by tDto.NormalizeDays().\n"+
				"\nError='%v'\n", err.Error())
	}

	tDto.ConvertToAbsoluteValues()

	var dt2 time.Time

	dtMech := DTimeMechanics{}

	if timeCalcMode == TCalcMode.LocalTimeZone() {

		dt2 = dtMech.AddDateTimeByLocalTimeZone(
			dTz.dateTimeValue,
			tDto.Years,
			tDto.Months,
			tDto.DateDays,
			tDto.Hours,
			tDto.Minutes,
			tDto.Seconds,
			tDto.Milliseconds,
			tDto.Microseconds,
			tDto.Nanoseconds)

	} else if timeCalcMode == TCalcMode.UtcTimeZone() {

		dt2 = dtMech.AddDateTimeByUtc(
			dTz.dateTimeValue,
			tDto.Years,
			tDto.Months,
			tDto.DateDays,
			tDto.Hours,
			tDto.Minutes,
			tDto.Seconds,
			tDto.Milliseconds,
			tDto.Microseconds,
			tDto.Nanoseconds)

	} else {

		return DateTzDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeCalcMode",
				inputParameterValue: "",
				errMsg:              "",
				err:                 nil,
			}

	}

	dTz2 := DateTzDto{}

	dTzUtil2 := dateTzDtoUtility{}

	err = dTzUtil2.setFromDateTime(&dTz2, dt2, dTz.dateTimeFmt, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dTz2, nil
}

// addTime - Adds input parameter time components (hours, minutes, seconds etc.)
// to the date time value of the input parameter 'dTz' (DateTzDto). The resulting
// updated date time value is returned to the calling function in the form of a
// new DateTzDto instance.
//
// The value of the input parameter 'dTz' (DateTzDto) instance is NOT altered.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dTz *DateTzDto - Provides the base date to which input parameters 'hours',
//                  'minutes', 'seconds', 'milliseconds', 'microseconds' and
//                   'nanoseconds' are added.
//
//  hours             int   - Number of hours to add to the value of parameter 'dTz'.
//  minutes           int   - Number of minutes to add to the value of parameter 'dTz'.
//  seconds           int   - Number of seconds to add to the value of parameter 'dTz'.
//  milliseconds      int   - Number of milliseconds to add to the value of parameter 'dTz'.
//  microseconds      int   - Number of microseconds to add to the value of parameter 'dTz'.
//  nanoseconds       int   - Number of hours to add to the value of parameter 'dTz'.
//
//          Note: Date Component input parameters may be either negative
//                or positive. Negative values will subtract time from
//                the current DateTzDto instance.
//
//  dateTimeFmtStr  string
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
//  ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a new, valid, fully populated
//              DateTzDto type updated to reflect the added input parameters,
//              years, months, days, hours, minutes, seconds, milliseconds,
//              microseconds and nanoseconds.
//
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) addTime(
	dTz *DateTzDto,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFormatStr,
	ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addTime() "

	if dTz == nil {
		return DateTzDto{},
			errors.New(ePrefix +
				"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	var dt2 time.Time

	dtMech := DTimeMechanics{}

	dt2 = dtMech.AddDateTimeByLocalTimeZone(
			dTz.dateTimeValue,
			0,
			0,
			0,
			hours,
			minutes,
			seconds,
			milliseconds,
			microseconds,
			nanoseconds)

	dTzUtil2 := dateTzDtoUtility{}

	dtz2 := DateTzDto{}

	err := dTzUtil2.setFromDateTime(&dtz2, dt2, dateTimeFormatStr, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// copyIn - Receives two parameters which are pointers
// to types DateTzDto. The method then copies all of
// the data field values from 'incomingDtz' into
// 'baseDtz'.
//
func (dTzUtil *dateTzDtoUtility) copyIn(
	baseDtz,
	incomingDtz *DateTzDto) {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	if baseDtz == nil {
		panic("dateTzDtoUtility.copyIn() - baseDtz is a 'nil' pointer!")
	}

	if baseDtz.lock == nil {
		baseDtz.lock = new(sync.Mutex)
	}

	if incomingDtz == nil {
		panic("dateTzDtoUtility.copyIn() - incomingDtz is a 'nil' pointer!")
	}

	if incomingDtz.lock == nil {
		incomingDtz.lock = new(sync.Mutex)
	}

	dTzUtil2 := dateTzDtoUtility{}

	dTzUtil2.empty(baseDtz)

	baseDtz.tagDescription = incomingDtz.tagDescription
	baseDtz.timeComponents = incomingDtz.timeComponents.CopyOut()
	baseDtz.dateTimeFmt = incomingDtz.dateTimeFmt

	if incomingDtz.dateTimeValue.IsZero() {
		baseDtz.timeZone = TimeZoneDefinition{}
		baseDtz.dateTimeValue = time.Time{}
	} else {
		baseDtz.dateTimeValue = incomingDtz.dateTimeValue
		baseDtz.timeZone = incomingDtz.timeZone.CopyOut()
	}

}

// copyOut - Returns a deep copy of input parameter
// 'dTz' which is a pointer to a type 'DateTzDto'.
//
func (dTzUtil *dateTzDtoUtility) copyOut(
	dTz *DateTzDto) DateTzDto {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	if dTz == nil {
		panic("dateTzDtoUtility.copyOut() - Input " +
			"parameter 'dTz' is a 'nil' pointer!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	dtz2 := DateTzDto{}

	dtz2.tagDescription = dTz.tagDescription
	dtz2.timeComponents = dTz.timeComponents.CopyOut()
	dtz2.dateTimeFmt = dTz.dateTimeFmt

	if dTz.dateTimeValue.IsZero() {
		dtz2.timeZone = TimeZoneDefinition{}
		dtz2.dateTimeValue = time.Time{}
	} else {
		dtz2.dateTimeValue = dTz.dateTimeValue
		dtz2.timeZone = dTz.timeZone.CopyOut()
	}

	dtz2.lock = new(sync.Mutex)

	return dtz2
}

// compareDateTimeValue - Compares the date time values
// of two DateTzDto objects.
//
// If the 'dTz1' date time value is is less than that of
// 'dTz2', this method returns an integer value of '-1'
// (minus one).
//
// If the 'dTz1' date time value is equal to that of
// 'dTz2', this method returns an integer value of '0'
// (zero).
//
//
// If the 'dTz1' date time values is is greater than
// that of 'dTz2', this method returns an integer value
// of '1' (plus one).
//
// Return Values
// =============
//
// -1 = 'dTz1' is less than the date time value of 'dTz2'
//  0 = 'dTz1' is equal to the date time value of 'dTz2'
//  1 = 'dTz1' is greater than the date time value of 'dTz2'
//
func (dTzUtil *dateTzDtoUtility) compareDateTimeValue(
	dTz1,
	dTz2 *DateTzDto,
	ePrefix string) (int, error) {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.compareDateTimeValue() "

	if dTz1 == nil {
		return -99,
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "dTz1",
				inputParameterValue: "",
				errMsg:              "Input parameter 'dTz1' is a nil pointer!",
				err:                 nil,
			}
	}

	if dTz1.lock == nil {
		dTz1.lock = new(sync.Mutex)
	}


	if dTz2 == nil {
		return -99,
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "dTz2",
				inputParameterValue: "",
				errMsg:              "Input parameter 'dTz2' is a nil pointer!",
				err:                 nil,
			}
	}

	if dTz2.lock == nil {
		dTz2.lock = new(sync.Mutex)
	}

	if dTz1.dateTimeValue.Before(dTz2.dateTimeValue) {
		return -1, nil
	}

	if dTz1.dateTimeValue.After(dTz2.dateTimeValue) {
		return 1, nil
	}

	return 0, nil
}

// empty - Receives a pointer to a type 'DateTzDto' and
// proceeds to set all internal member variables to their
// 'zero' or uninitialized values.
//
func (dTzUtil *dateTzDtoUtility) empty(dTz *DateTzDto) {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	if dTz == nil {
		panic("dateTzDtoUtility.empty() - " +
			"Input parameter 'dTz' is a 'nil' pointer!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	dTz.tagDescription = ""
	dTz.timeComponents.Empty()
	dTz.timeZone = TimeZoneDefinition{}
	dTz.dateTimeValue = time.Time{}
	dTz.dateTimeFmt = ""

	return
}

// isEmptyDateTzDto - Analyzes an instanceof DateTzDto to
// determine if all data fields are uninitialized or zero
// values.
//
func (dTzUtil *dateTzDtoUtility) isEmptyDateTzDto(
	dTz *DateTzDto) bool {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	if dTz == nil {
		panic("dateTzDtoUtility.isEmptyDateTzDto() - " +
			"Input parameter 'dTz' is a 'nil' pointer!")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	if dTz.tagDescription == "" &&
		dTz.timeComponents.IsEmpty() &&
		dTz.dateTimeValue.IsZero() &&
		dTz.dateTimeFmt == "" &&
		dTz.timeZone.IsEmpty() {

		return true
	}

	return false
}

// isValidDateTzDto - Analyzes an instance of 'DateTzDto' to
// determine if is value. If the instance evaluates as invalid,
// an error is returned.
//
func (dTzUtil *dateTzDtoUtility) isValidDateTzDto(
	dTz *DateTzDto,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.isValidDateTzDto() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	dTzUtil2 := dateTzDtoUtility{}

	if dTzUtil2.isEmptyDateTzDto(dTz) {
		return errors.New(ePrefix +
			"\nThis 'DateTzDto' instance is EMPTY!\n")
	}

	if dTz.dateTimeValue.IsZero() {
		return errors.New(ePrefix +
			"\nError: DateTzDto.DateTime is ZERO!\n")
	}

	if dTz.timeZone.IsEmpty() {
		return errors.New(ePrefix +
			"\nError: DateTzDto.TimeZone is EMPTY!\n")
	}

	if err := dTz.timeZone.IsValid(); err != nil {
		return fmt.Errorf(ePrefix +
			"\nError: DateTzDto Time Zone is INVALID!\n" +
			"Error='%v'\n", err.Error())
	}

	if err := dTz.timeComponents.IsValid(); err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: DateTzDto timeComponents is INVALID!\n"+
			"Error='%v'\n", err.Error())
	}

	return nil
}

// preProcessDateFormatStr - Provides a standardized method
// for implementing a default date time format string.
//
func (dTzUtil *dateTzDtoUtility) preProcessDateFormatStr(
	dateTimeFmtStr string) string {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	dateTimeFmtStr = strings.TrimLeft(strings.TrimRight(dateTimeFmtStr, " "), " ")

	if len(dateTimeFmtStr) == 0 {
		lockDefaultDateTimeFormat.Lock()
		dateTimeFmtStr = DEFAULTDATETIMEFORMAT
		lockDefaultDateTimeFormat.Unlock()
	}

	return dateTimeFmtStr
}

// setDateTimeFormat - Sets the Date Time Format
// string in the 'dtz' object.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dTz *DateTzDto
//     - The date time format member variable contained within this
//       'DateTzDto' instance will be set to new value specified by
//        input parameter, 'dateTimeFmtStr'.
//
//  dateTimeFmtStr  string
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
//  ePrefix string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//   None
//
func (dTzUtil *dateTzDtoUtility) setDateTimeFormat(
	dtz *DateTzDto,
	dateTimeFmtStr string,
	ePrefix string) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	if dtz == nil {
		ePrefix += "dateTzDtoUtility.setDateTimeFormat() "
		panic (ePrefix + "\nInput parameter 'dtz' is a 'nil' pointer!\n")
	}

	if dtz.lock == nil {
		dtz.lock = new(sync.Mutex)
	}

	dTzUtil2 := dateTzDtoUtility{}

	dtz.dateTimeFmt =
		dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)
}


// setFromDateTime - Sets the values for DateTzDto fields encapsulated
// in input parameter 'dTz'. The field values will be set
// based on an input parameter 'dateTime' (Type time.time) and a
// date time format string.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dTz *DateTzDto
//     - The member variables contained with in this 'DateTzDto'
//       instance will be set to new values based on the following
//       input parameters.
//
//
//   dateTime    time.Time
//     - A date time value
//
//
//  dateTimeFmtStr  string
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
//  ePrefix string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// __________________________________________________________________________
//
// Return Values:
//
//  DurationTriad
//     - Upon successful completion, this method will return
//       a new, populated DurationTriad instance.
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
func (dTzUtil *dateTzDtoUtility) setFromDateTime(
	dTz *DateTzDto,
	dateTime time.Time,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromDateTime() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input Parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO!\n")
	}

	dTzUtil2 := dateTzDtoUtility{}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	tDto := TimeDto{}

	tDtoUtil := timeDtoUtility{}

	err := tDtoUtil.setFromDateTime(&tDto, dateTime, ePrefix)

	if err != nil {
		return err
	}

	timeZone := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromDateTime(&timeZone, dateTime, ePrefix)

	if err != nil {
		return err
	}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = dateTime
	dTz.timeComponents = tDto.CopyOut()
	dTz.timeZone = timeZone.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}

// setFromDateTimeComponents - Sets the values of the Date Time fields
// for the current DateTzDto instance based on time components
// and a Time Zone Location.
//
// Note that this variation of time elements breaks time down by
// hour, minute, second, millisecond, microsecond and nanosecond.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dTz *DateTzDto
//     - The member variables contained with in this 'DateTzDto'
//       instance will be set to new values based on the following
//       input parameters.
//
//  year              int   - Year value.
//  month             int   - Month value.
//  day               int   - Day value.
//  hour              int   - Hour value .
//  minute            int   - Minutes value.
//  second            int   - Seconds value.
//  millisecond       int   - Milliseconds value.
//  microsecond       int   - Microseconds value.
//  nanosecond        int   - Nanoseconds value.
//
//          Note: Date Component input parameters may be either negative
//                or positive. Negative values will subtract time from
//                the current DateTzDto instance.
//
//   timeZoneLocationName  string
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
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//  dateTimeFmtStr  string
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
//  ePrefix string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) setFromDateTimeComponents(
	dTz *DateTzDto,
	year,
	month,
	day,
	hour,
	minute,
	second,
	millisecond,
	microsecond,
	nanosecond int,
	timeZoneLocationName,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromTimeDto() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	tDto, err := TimeDto{}.NewTimeComponents(year, month, 0, day, hour, minute,
		second, millisecond, microsecond, nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by TimeDto{}.NewStartEndTimes(year, month,...).  "+
			"Error='%v'", err.Error())
	}

	dTzUtil2 := dateTzDtoUtility{}

	fmtStr := dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	dtUtil := TimeZoneMechanics{}

	timeZoneLocationName = dtUtil.PreProcessTimeZoneLocation(timeZoneLocationName)

	if len(timeZoneLocationName) == 0 {
		return fmt.Errorf(ePrefix +
			"\nError: Input Parameter 'timeZoneLocationName' " +
			"resolved to an empty string!\n" +
			"timeZoneLocationName='%v'\n", timeZoneLocationName)
	}

	dtMech := DTimeMechanics{}

	_, err = dtMech.LoadTzLocation(timeZoneLocationName, ePrefix)

	if err != nil {
		return err
	}

	var dt time.Time

	dt, err = tDto.GetDateTime(timeZoneLocationName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto.GetDateTime(timeZoneLocationName).\n"+
			"\ntimeZoneLocationName='%v'\nError='%v'\n",
			timeZoneLocationName, err.Error())
	}

	timeZone := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromTimeZoneName(
		&timeZone,
		dt,
		TzConvertType.Absolute(),
		timeZoneLocationName,
		ePrefix)

	if err != nil {
		return err
	}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = dt
	dTz.timeZone = timeZone.CopyOut()
	dTz.timeComponents = tDto.CopyOut()
	dTz.dateTimeFmt = fmtStr

	return nil
}

// setFromDateTimeElements - Sets the values of input parameter
// 'dTz' (type DateTzDto). 'dTz' data fields are set based on
// input parameters consisting of date time elements,
// a time zone location and a date time format string.
//
// Date Time elements include year, month, day, hour, minute,
// second and nanosecond.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dTz *DateTzDto
//     - The member variables contained with in this 'DateTzDto'
//       instance will be set to new values based on the following
//       input parameters.
//
//  year              int   - Year value.
//  month             int   - Month value.
//  day               int   - Day value.
//  hour              int   - Hour value .
//  minute            int   - Minutes value.
//  second            int   - Seconds value.
//  nanosecond        int   - Nanoseconds value.
//
//          Note: Date Component input parameters may be either negative
//                or positive. Negative values will subtract time from
//                the current DateTzDto instance.
//
//   timeZoneLocationName  string
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
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//  dateTimeFmtStr  string
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
//  ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) setFromDateTimeElements(
	dTz *DateTzDto,
	year,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	timeZoneLocationName,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromDateTimeElements() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	tDto, err := TimeDto{}.NewTimeComponents(
		year,
		month,
		0,
		day,
		hour,
		minute,
		second,
		0,
		0,
		nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned from TimeDto{}.NewStartEndTimes(year, month, ...).\n"+
			"Error='%v'\n", err.Error())
	}

	dTzUtil2 := dateTzDtoUtility{}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	tzMech := TimeZoneMechanics{}

	timeZoneLocationName = tzMech.PreProcessTimeZoneLocation(timeZoneLocationName)

	if len(timeZoneLocationName) == 0 {
		return errors.New(ePrefix +
			"\nError: Input Parameter 'timeZoneLocationName' is an empty string!\n")
	}

	dtMech := DTimeMechanics{}

	_, err = dtMech.LoadTzLocation(timeZoneLocationName, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by dtMech.LoadTzLocation(timeZoneLocationName, ePrefix).\n"+
			"INVALID 'timeZoneLocationName'!\n"+
			"timeZoneLocationName='%v'\n"+
			"Error='%v'\n",
			timeZoneLocationName, err.Error())
	}

	var dt time.Time

	dt, err = tDto.GetDateTime(timeZoneLocationName)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto.GetDateTime(tzl).\n"+
			"\ntimeZoneLocationName='%v'\n"+
			"Error='%v'\n",
			timeZoneLocationName, err.Error())
	}

	timeZone := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromTimeZoneName(
		&timeZone,
		dt,
		TzConvertType.Absolute(),
		timeZoneLocationName,
		ePrefix)

	if err != nil {
		return err
	}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = dt
	dTz.timeZone = timeZone.CopyOut()
	dTz.timeComponents = tDto.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}

// setFromTimeTzName - Sets the date and time values for input
// parameter 'dTz' (type DateTzDto). The new values will be based
// on input parameters 'dateTime', 'timeZoneLocationName' and a date
// time format string, 'dateTimeFmtStr'.
//
// 'timeZoneSpec' is a valid instance of TimeZoneSpecification.
// Parameter 'timeZoneConversionType' is an instance of the
// 'TimeZoneConversionType' enumeration and determines how
// date time will be converted to the target time zone
// represented by parameter, 'timeZoneLocationName'.
//
// The parameter, 'timeZoneConversionType', is an instance
// the type enumeration type TimeZoneConversionType.
// This parameter will determine how 'dateTime' will be
// converted to the target time zone.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   dTz                     *DateTzDto
//     - The member variables contained with in this 'DateTzDto'
//       instance will be set to new values based on the following
//       input parameters.
//
//   dateTime                time.Time
//     - A date time value
//
//   timeZoneConversionType  TimeZoneConversionType
//     - This parameter determines the algorithm that will
//       be used to convert parameter 'dateTime' to the time
//       zone specified by parameter 'timeZoneName'.
//
//       TimeZoneConversionType is an enumeration type which
//       be used to convert parameter 'dateTime' to the time
//       must be set to one of two values:
//       This parameter determines the algorithm that will
//          TimeZoneConversionType(0).Absolute()
//          TimeZoneConversionType(0).Relative()
//       Note: You can also use the global variable
//       'TzConvertType' for easier access:
//          TzConvertType.Absolute()
//          TzConvertType.Relative()
//
//       Absolute Time Conversion - Identifies the 'Absolute' time
//       to time zone conversion algorithm. This algorithm provides
//       that a time value in time zone 'X' will be converted to the
//       same time value in time zone 'Y'.
//
//       For example, assume the time 10:00AM is associated with time
//       zone USA Central Standard time and that this time is to be
//       converted to USA Eastern Standard time. Applying the 'Absolute'
//       algorithm would convert ths time to 10:00AM Eastern Standard
//       time.  In this case the hours, minutes and seconds have not been
//       altered. 10:00AM in USA Central Standard Time has simply been
//       reclassified as 10:00AM in USA Eastern Standard Time.
//
//       Relative Time Conversion - Identifies the 'Relative' time to time
//       zone conversion algorithm. This algorithm provides that times in
//       time zone 'X' will be converted to their equivalent time in time
//       zone 'Y'.
//
//       For example, assume the time 10:00AM is associated with time zone
//       USA Central Standard time and that this time is to be converted to
//       USA Eastern Standard time. Applying the 'Relative' algorithm would
//       convert ths time to 11:00AM Eastern Standard time. In this case the
//       hours, minutes and seconds have been changed to reflect an equivalent
//       time in the USA Eastern Standard Time Zone.
//
//
//   timeZoneLocationName    string
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
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//   dateTimeFmtStr          string
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
//   ePrefix                 string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   error - If successful the returned error Type is set equal to 'nil'. If errors are
//           encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) setFromTimeTzName(
	dTz *DateTzDto,
	dateTime time.Time,
	timeZoneConversionType TimeZoneConversionType,
	timeZoneLocationName string,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromTimeTzName() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO and INVALID!\n")
	}

	tZoneDefDto := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	var err error

	err = tzDefUtil.setFromTimeZoneName(
		&tZoneDefDto,
		dateTime,
		timeZoneConversionType,
		timeZoneLocationName,
		ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: Invalid Time Zone Location Name!\n"+
			"timeZoneLocationName='%v'\nError='%v'\n",
			timeZoneLocationName, err.Error())
	}

	var targetDateTime time.Time

	if timeZoneConversionType == TzConvertType.Absolute() {
		// FmtDateTimeTzNanoYMD = "2006-01-02 15:04:05.000000000 -0700 MST"
		dtMech := DTimeMechanics{}
		targetDateTime, err = dtMech.AbsoluteTimeToTimeZoneDefConversion(dateTime, tZoneDefDto)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by dtMech.AbsoluteTimeToTimeZoneDefConversion(dateTime,tZoneDefDto)\n"+
				"dateTime='%v'\ntZoneDefDto='%v'\nError='%v'\n",
				dateTime.Format(FmtDateTimeTzNanoYMD), tZoneDefDto.GetOriginalLocationName(), err.Error())
		}
	} else {
		// Must be TzConvertType.Relative() or TzConvertType.None()
		// This the default.
		targetDateTime = dateTime.In(tZoneDefDto.GetOriginalLocationPtr())
	}

	var tDto TimeDto

	tDtoUtil := timeDtoUtility{}

	err =tDtoUtil.setFromDateTime(&tDto, targetDateTime, ePrefix)

	if err != nil {
		return  err
	}

	dTzUtil2 := dateTzDtoUtility{}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = targetDateTime
	dTz.timeZone = tZoneDefDto
	dTz.timeComponents = tDto
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}

// setFromTzDef - Uses a 'TimeZoneDefinition' instance
// to configure time and time zone data in parameter,
// 'dTz'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   dTz                    *DateTzDto
//     - The member variables contained with in this 'DateTzDto'
//       instance will be set to new values based on the following
//       input parameters.
//
//
//   dateTime               time.Time
//     - A date time value
//
//
//   timeZoneDef            TimeZoneDefinition
//     - A detailed time zone definition containing specifications for both an
//       original time zone and a convertible time zone.  This time zone definition
//       will be used to set the time zone for the 'dTz' instance.
//
//
//   timeZoneConversionType TimeZoneConversionType -
//           This parameter determines the algorithm that will
//           be used to convert parameter 'dateTime' to the time
//           zone specified by parameter 'timeZoneName'.
//
//           TimeZoneConversionType is an enumeration type which
//           must be set to one of two values:
//              TimeZoneConversionType(0).Absolute()
//              TimeZoneConversionType(0).Relative()
//           Note: You can also use the global variable
//           'TzConvertType' for easier access:
//              TzConvertType.Absolute()
//              TzConvertType.Relative()
//
//           Absolute Time Conversion - Identifies the 'Absolute' time
//           to time zone conversion algorithm. This algorithm provides
//           that a time value in time zone 'X' will be converted to the
//           same time value in time zone 'Y'.
//
//           For example, assume the time 10:00AM is associated with time
//           zone USA Central Standard time and that this time is to be
//           converted to USA Eastern Standard time. Applying the 'Absolute'
//           algorithm would convert ths time to 10:00AM Eastern Standard
//           time.  In this case the hours, minutes and seconds have not been
//           altered. 10:00AM in USA Central Standard Time has simply been
//           reclassified as 10:00AM in USA Eastern Standard Time.
//
//           Relative Time Conversion - Identifies the 'Relative' time to time
//           zone conversion algorithm. This algorithm provides that times in
//           time zone 'X' will be converted to their equivalent time in time
//           zone 'Y'.
//
//           For example, assume the time 10:00AM is associated with time zone
//           USA Central Standard time and that this time is to be converted to
//           USA Eastern Standard time. Applying the 'Relative' algorithm would
//           convert ths time to 11:00AM Eastern Standard time. In this case the
//           hours, minutes and seconds have been changed to reflect an equivalent
//           time in the USA Eastern Standard Time Zone.
//
//   dateTimeFmtStr         string
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
//   ePrefix                string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) setFromTzDef(
	dTz *DateTzDto,
	dateTime time.Time,
	timeZoneDef TimeZoneDefinition,
	timeConversionType TimeZoneConversionType,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromTzDef() "

	if dTz == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dTz",
			inputParameterValue: "Input parameter 'dTz' is a 'nil' pointer!",
			errMsg:              "",
			err:                 nil,
		}
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	if dateTime.IsZero() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Input parameter 'dateTime' has a Zero Value!",
			err:                 nil,
		}
	}

	if timeConversionType < TzConvertType.Absolute() ||
		timeConversionType > TzConvertType.Relative() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: timeConversionType.String(),
			errMsg:              "Input parameter " +
				"'timeConversionType' MUST be either 'Absolute' or 'Relative'!",
			err:                 nil,
		}
	}

	dTzUtil2 := dateTzDtoUtility{}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	tzDefUtil :=timeZoneDefUtility{}

	timeZoneDefOut := TimeZoneDefinition{}

	err := tzDefUtil.setFromTimeZoneDefinition(
		&timeZoneDefOut,
		dateTime,
		timeConversionType,
		timeZoneDef,
		ePrefix)

	if err != nil {
		return err
	}


	tDto := TimeDto{}

	tDtoUtil := timeDtoUtility{}

	err = tDtoUtil.setFromDateTime(
		&tDto,
		dateTime,
		ePrefix)

	if err != nil {
		return err
	}

	dTz.timeZone = timeZoneDefOut.CopyOut()
	dTz.timeComponents = tDto.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr
	dTz.dateTimeValue = timeZoneDefOut.originalTimeZone.referenceDateTime

	return nil
}

// setFromTzSpec - Sets the date and time values for input
// parameter 'dTz' (type DateTzDto). The new values will be
// based on input parameters 'dateTime', 'timeZoneSpec' and
// a date time format string, 'dateTimeFmtStr'.
//
// 'timeZoneSpec' is a valid instance of TimeZoneSpecification.
// Parameter 'timeZoneConversionType' is an instance of the
// 'TimeZoneConversionType' enumeration and determines how
// date time will be converted to the target time zone
// represented by parameter, 'timeZoneSpec'.
//
// The parameter, 'timeZoneConversionType', is an instance
// the type enumeration type TimeZoneConversionType.
// This parameter will determine how 'dateTime' will be
// converted to the target time zone.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dTz                     *DateTzDto
//     - The member variables contained with in this 'DateTzDto'
//       instance will be set to new values based on the following
//       input parameters.
//
//
//   dateTime               time.Time
//     - A date time value
//
//
//   timeZoneSpec           TimeZoneSpecification
//     - A detailed time zone specification which will be
//       used to set the time zone for the 'dTz' instance.
//
//
//   timeZoneConversionType TimeZoneConversionType -
//           This parameter determines the algorithm that will
//           be used to convert parameter 'dateTime' to the time
//           zone specified by parameter 'timeZoneName'.
//
//           TimeZoneConversionType is an enumeration type which
//           must be set to one of two values:
//              TimeZoneConversionType(0).Absolute()
//              TimeZoneConversionType(0).Relative()
//           Note: You can also use the global variable
//           'TzConvertType' for easier access:
//              TzConvertType.Absolute()
//              TzConvertType.Relative()
//
//           Absolute Time Conversion - Identifies the 'Absolute' time
//           to time zone conversion algorithm. This algorithm provides
//           that a time value in time zone 'X' will be converted to the
//           same time value in time zone 'Y'.
//
//           For example, assume the time 10:00AM is associated with time
//           zone USA Central Standard time and that this time is to be
//           converted to USA Eastern Standard time. Applying the 'Absolute'
//           algorithm would convert ths time to 10:00AM Eastern Standard
//           time.  In this case the hours, minutes and seconds have not been
//           altered. 10:00AM in USA Central Standard Time has simply been
//           reclassified as 10:00AM in USA Eastern Standard Time.
//
//           Relative Time Conversion - Identifies the 'Relative' time to time
//           zone conversion algorithm. This algorithm provides that times in
//           time zone 'X' will be converted to their equivalent time in time
//           zone 'Y'.
//
//           For example, assume the time 10:00AM is associated with time zone
//           USA Central Standard time and that this time is to be converted to
//           USA Eastern Standard time. Applying the 'Relative' algorithm would
//           convert ths time to 11:00AM Eastern Standard time. In this case the
//           hours, minutes and seconds have been changed to reflect an equivalent
//           time in the USA Eastern Standard Time Zone.
//
//   dateTimeFmtStr         string
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
//   ePrefix                string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) setFromTzSpec(
	dTz *DateTzDto,
	dateTime time.Time,
	timeZoneSpec TimeZoneSpecification,
	timeZoneConversionType TimeZoneConversionType,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromTzSpec() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO and INVALID!\n")
	}

	err := timeZoneSpec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nInput parameter 'timeZoneSpec' is Invalid!\n" +
			"Validation Error='%v'\n", err.Error())
	}

	tzDefUtil := timeZoneDefUtility{}

	tzDef := TimeZoneDefinition{}

	err = tzDefUtil.setFromTimeZoneSpecification(
		&tzDef,
		dateTime,
		timeZoneConversionType,
		timeZoneSpec,
		ePrefix)

	if err != nil {
		return err
	}

	var tDto TimeDto

	tDtoUtil := timeDtoUtility{}

	err =tDtoUtil.setFromDateTime(
		&tDto,
		tzDef.GetOriginalDateTime(),
		ePrefix)

	if err != nil {
		return  err
	}

	dTzUtil2 := dateTzDtoUtility{}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = tzDef.GetOriginalDateTime()
	dTz.timeZone = tzDef
	dTz.timeComponents = tDto
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}

// setFromTimeDto - Receives data from a TimeDto input parameter
// and sets all data fields of the current DateTzDto instance
// accordingly. When the method completes, the values of the
// current DateTzDto will equal the values of the input parameter
// TimeDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dTz              *DateTzDto
//     - The member variables contained with in this 'DateTzDto'
//       instance will be set to new values based on the following
//       input parameters.
//
//
//   tDto            TimeDto
//     - A populated TimeDto instance.
//
//       The TimeDto structure is defined as follows:
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
//   timeZoneLocationName  string
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
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//  ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   error - If successful the returned error Type is set equal to 'nil'. If errors are
//           encountered this error Type will encapsulate an error message.
//
func (dTzUtil *dateTzDtoUtility) setFromTimeDto(
	dTz *DateTzDto,
	tDto TimeDto,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dTz.lock == nil {
		dTz.lock = new(sync.Mutex)
	}

	ePrefix += "dateTzDtoUtility.setFromTimeDto() "

	if tDto.IsEmpty() {

		return fmt.Errorf(ePrefix + "\nError: Input parameter 'tDto' date time elements equal ZERO!\n")
	}

	tDto2 := tDto.CopyOut()

	var err error

	err = tDto2.NormalizeTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto2.NormalizeTimeElements().\nError='%v'\n",
			err.Error())
	}

	tDto2.ConvertToAbsoluteValues()

	if err = tDto2.IsValid(); err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: Input Parameter tDto (TimeDto) is INVALID.\nError='%v'\n",
			err.Error())
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	if len(timeZoneLocation) == 0 {
		return errors.New(ePrefix +
			"\nError: Input Parameter 'timeZoneLocation' is an empty string!\n")
	}

	dtMech := DTimeMechanics{}
	
	_, err = dtMech.LoadTzLocation(timeZoneLocation, ePrefix)

	if err != nil {
		return err
	}

	dTzUtil2 := dateTzDtoUtility{}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	var dateTime time.Time

	dateTime, err = tDto.GetDateTime(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto.GetDateTime(timeZoneLocation).\n"+
			"timeZoneLocation='%v'\nError='%v'\n",
			timeZoneLocation, err.Error())
	}

	timeZoneDef := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromDateTime(&timeZoneDef, dateTime, ePrefix)

	if err != nil {
		return err
	}

	dTzUtil2.empty(dTz)
	dTz.dateTimeValue = dateTime
	dTz.timeZone = timeZoneDef.CopyOut()
	dTz.timeComponents = tDto2.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}
