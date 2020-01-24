package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// DateTzDto
//
// This source file is located in source code repository:
//    'https://github.com/MikeAustin71/datetimeopsgo.git'
//
// This source code file is located at:
//    MikeAustin71\datetimeopsgo\datetime\datetzdto.go
//
// ------------------------------------------------------------------------
//
// Overview and Usage
//
// The 'DateTzDto' type is used to store and transfer date time information.
// The descriptors contained is this structure are intended to define and
// identify a specific point in time. In addition to date and time identifiers,
// this type also includes information on associated Time Zones and Time Elements.
// Time elements includes years, months, weeks, days, hours, minutes, seconds,
// milliseconds, microseconds and nanoseconds.
//
// 'DateTzDto' is used primarily conjunction with IANA Time Zones. For more information
// on IANA Time Zones, see type 'TimeZones', located in source file:
//
//    Source Repository: 'https://github.com/MikeAustin71/datetimeopsgo.git'
//     Source Code File:  MikeAustin71\datetimeopsgo\datetime\timezonedata.go
//
//
//For Military Time Zones use type, 'MilitaryDateTzDto'.
//
// This Type is NOT used to define time duration; that is, the difference or time
// span between two points in time. For time duration calculations refer to types,
// 'TimeDurationDto' and 'DurationTriad' located in source files:
//
//    'github.com/MikeAustin71/datetimeopsgo/datetime/timedurationdto.go'
//    'github.com/MikeAustin71/datetimeopsgo/datetime/durationtriad.go'
//
// As previously stated, 'DateTzDto' defines a specific point in time using
// a variety of descriptors including year, month, day hour, minute, second,
// millisecond, microsecond and nanosecond. In addition this Type specifies a
// time.Time value as well as time zone location and time zone.
//
// If you are unfamiliar with the concept of a time zone location, reference
// 'https://golang.org/pkg/time/'. The concept of Time Zone Location is important
// and several of the 'DateTzDto' methods use Time Zone Location. Time Zone location
// must be designated as one of two values.
//
//      (1) The string 'Local' - signals the designation of the local time zone
//          location for the host computer.
//
//      (2) IANA Time Zone Location -
//          See https://golang.org/pkg/time/#LoadLocation
//          and https://www.iana.org/time-zones to ensure that
//          the IANA Time Zone Database is properly configured
//          on your system. Note: IANA Time Zone Data base is
//          equivalent to 'tz database'.
//
//          Examples:
//            "America/New_York"
//            "America/Chicago"
//            "America/Denver"
//            "America/Los_Angeles"
//            "Pacific/Honolulu"
//
//          The source file 'constantsdatetime.go' contains a number of
//          constant declarations covering the more frequently used time
//          zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//          time zone constants begin with the prefix 'TzIana'.
//
//
// A requirement for presentation of date time strings is a specific format
// for displaying years, months, days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds. Many 'DateTzDto' methods require calling functions
// to provide a date time format string, ('dateTimeFmtStr'). This format string
// is used to configure date times for display purposes.
//
// dateTimeFmtStr string   - A date time format string which will be used
//                           to format and display 'dateTime'. Example:
//                           "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                           Date time format constants are found in the source
//                           file 'constantsdatetime.go'. These constants represent
//                           the more commonly used date time string formats. All
//                           Date Time format constants begin with the prefix
//                           'FmtDateTime'.
//
//                           If 'dateTimeFmtStr' is submitted as an
//                           'empty string', a default date time format
//                           string will be applied. The default date time
//                           format string is:
//                             FmtDateTimeYrMDayFmtStr =
//                                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// DateTzDto Structure and Methods
//
// ===============================
//
type DateTzDto struct {
	tagDescription string       // Available for tags, classification, labeling or description
	timeComponents TimeDto      // Associated Time Components (years, months, days, hours, minutes,
	                            //    seconds etc.)
	dateTimeValue  time.Time    // DateTime value for this DateTzDto Type
	dateTimeFmt    string       // Date Time Format String. Default is
	                            //    "2006-01-02 15:04:05.000000000 -0700 MST"
	timeZone TimeZoneDefinition // Contains a detailed definition and descriptions of the Time
	//                             Zone and Time Zone Location associated with this date time.
	lock        sync.Mutex      // Mutex used to ensure thread-safe operations.
}

// AddDate - Adds input parameters 'years, 'months' and 'days' to date time value of the
// current DateTzDto and returns the updated value in a new DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//
//  du, err := dtz.AddDate(
//                  years,
//                  months,
//                  days,
//                  FmtDateTimeYrMDayFmtStr)
//
//  Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in
//        source file 'constantsdatetime.go'.
//
func (dtz *DateTzDto) AddDate(
	years,
	months,
	days int,
	dateTimeFormatStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddDate() "

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.isValidDateTzDto(dtz, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	newDt1 := dtz.dateTimeValue.AddDate(years, months, 0)

	dur := DayNanoSeconds * int64(days)
	newDt2 := newDt1.Add(time.Duration(dur))

	if dateTimeFormatStr == "" {
		dateTimeFormatStr = dtz.dateTimeFmt
	}

	dtz2 := DateTzDto{}

	err = dTzUtil.setFromDateTime( &dtz2, newDt2, dateTimeFormatStr, ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// AddDateTime - Adds date time components to the date time value of the
// current DateTzDto instance. The updated date time value is returned to
// the calling function as a new DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  years             int   - Number of years to add.
//  months            int   - Number of months to add.
//  days              int   - Number of days to add.
//  hours             int   - Number of hours to add.
//  minutes           int   - Number of minutes to add.
//  seconds           int   - Number of seconds to add.
//  milliseconds      int   - Number of milliseconds to add.
//  microseconds      int   - Number of microseconds to add.
//  nanoseconds       int   - Number of nanoseconds to add.
//
//  Note: Date Time Component input parameters may be either negative
//        or positive. Negative values will subtract time from
//        the current DateTzDto instance.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a new, valid, fully populated
//              DateTzDto type updated to reflect the addition of input
//              parameters to the date time value of the current DateTzDto.
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to a value.
//
//  dtz, err := dtz.addDateTime(
//                   years,
//                   months,
//                   days,
//                   hours,
//                   minutes,
//                   seconds,
//                   milliseconds,
//                   microseconds,
//                   nanoseconds,
//                   FmtDateTimeYrMDayFmtStr)
//
//  Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//        'constantsdatetime.go'.
//
func (dtz *DateTzDto) AddDateTime(
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFormatStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.addDateTime() "

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.addDateTime(
		dtz,
		years,
		months,
		days,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds,
		dateTimeFormatStr,
		ePrefix)

}

// AddDateTimeToThis - Adds date time components to the date time value of the current
// DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  years        int - Number of years to add.
//  months       int - Number of months to add.
//  days         int - Number of days to add.
//  hours        int - Number of hours to add.
//  minutes      int - Number of minutes to add.
//  seconds      int - Number of seconds to add.
//  milliseconds int - Number of milliseconds to add.
//  microseconds int - Number of microseconds to add.
//  nanoseconds  int - Number of nanoseconds to add.
//
//  Note: Date Time Component input parameters may be either negative
//        or positive. Negative values will subtract time from
//        the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//          encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  err := dtz.AddDateTimeToThis(
//                years,
//                months,
//                days,
//                hours,
//                minutes,
//                seconds,
//                milliseconds,
//                microseconds,
//                nanoseconds)
//
func (dtz *DateTzDto) AddDateTimeToThis(
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddDateTimeToThis() "

	dTzUtil := dateTzDtoUtility{}

	dtz2, err :=dTzUtil.addDateTime(
							dtz,
							years,
							months,
							days,
							hours,
							minutes,
							seconds,
							milliseconds,
							microseconds,
							nanoseconds,
							dtz.dateTimeFmt,
							ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error='%v'", err.Error())
	}

	dTzUtil.copyIn(dtz, &dtz2)

	return nil
}

// AddDateToThis - Adds input parameters 'years, 'months' and 'days' to date time value
// of the current DateTzDto. The updated DateTime is retained in the current
// DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  years    int - Number of years to add to the current date.
//  months   int - Number of months to add to the current date.
//  days     int - Number of days to add to the current date.
//
//           Note: Date Component input parameters may be either negative
//                 or positive. Negative values will subtract time from
//                 the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//          encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  err := dtz.AddDateToThis(
//                years,
//                months,
//                days)
//
func (dtz *DateTzDto) AddDateToThis(
	years,
	months,
	days int) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddDateToThis() "

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.isValidDateTzDto(dtz, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nThe current DateTzDto is INVALID!\ndtz.dateTimeValue='%v'\n",
			dtz.dateTimeValue.Format(FmtDateTimeYrMDayFmtStr))
	}

	newDt1 := dtz.dateTimeValue.AddDate(years, months, 0)
	dur := int64(days) * DayNanoSeconds
	newDt2 := newDt1.Add(time.Duration(dur))

	dtz2 := DateTzDto{}

	err = dTzUtil.setFromDateTime( &dtz2, newDt2, dtz.dateTimeFmt, ePrefix)

	if err != nil {
		return err
	}

	dTzUtil.copyIn(dtz, &dtz2)

	return nil

}

// AddDuration - Adds Duration to the DateTime XValue of the current
// DateTzDto and returns a new DateTzDto instance with the updated
// Date Time value.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  duration time.Duration  - A Time duration value which is added to the DateTime
//                            value of the current DateTzDto instance to produce and
//                            return a new, updated DateTzDto instance.
//
//            Note: The time.Duration input parameter may be either negative
//                  or positive. Negative values will subtract time from
//                  the current DateTzDto instance.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a new, valid, fully populated
//              DateTzDto type updated to reflect the addition of input parameter
//              time duration to the date time value of the current DateTzDto
//              instance.
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  dtz2, err := dtz.AddDuration(
//                duration,
//                FmtDateTimeYrMDayFmtStr)
//
//  Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//        'constantsdatetime.go'.
//
func (dtz *DateTzDto) AddDuration(
	duration time.Duration,
	dateTimeFmtStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddDuration() "

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.addDuration(
					dtz,
					duration,
					dateTimeFmtStr,
					ePrefix)
}

// AddDurationToThis - Receives a time.Duration input parameter and adds this
// duration value to the Date Time value of the current DateTzDto. The current
// DateTzDto Date Time values are updated to reflect the added 'duration'.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  duration time.Duration - A Time duration value which is added to the DateTime
//                           value of the current DateTzDto instance to produce and
//                           return a new, updated DateTzDto instance.
//
//           Note: The time.Duration input parameter may be either negative
//                 or positive. Negative values will subtract time from
//                 the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Returns
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//          encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  err := dtz.AddDurationToThis(duration)
//
func (dtz *DateTzDto) AddDurationToThis(duration time.Duration) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddDurationToThis() "

	dTzUtil := dateTzDtoUtility{}

	dtz2, err := dTzUtil.addDuration(
						dtz,
						duration,
						dtz.dateTimeFmt,
						ePrefix)

	if err != nil {
		return err
	}

	dTzUtil.copyIn(dtz, &dtz2)

	return nil
}

// AddMinusTimeDto - Creates and returns a new DateTzDto by subtracting a
// TimeDto from the value of the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  minusTimeDto TimeDto - A TimeDto instance consisting of time components
//                         (years, months, weeks, days, hours, minutes etc.)
//                         which will be subtracted from the date time value
//                         of the current DateTzDto instance.
//
//        type TimeDto struct {
//          Years          int // Number of Years
//          Months         int // Number of Months
//          Weeks          int // Number of Weeks
//          WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//          DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//          Hours          int // Number of Hours.
//          Minutes        int // Number of Minutes
//          Seconds        int // Number of Seconds
//          Milliseconds   int // Number of Milliseconds
//          Microseconds   int // Number of Microseconds
//          Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//          TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                 // plus remaining Nanoseconds
//        }
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a valid, fully populated
//              DateTzDto type updated to reflect the subtracted 'TimeDto'
//              input parameter.
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  dtz2, err := dtz.AddMinusTimeDto(minusTimeDto)
//
func (dtz *DateTzDto) AddMinusTimeDto(minusTimeDto TimeDto) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddMinusTimeDto() "

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.addMinusTimeDto(
								dtz,
								minusTimeDto,
								ePrefix)
}

// AddMinusTimeDtoToThis - Modifies the current DateTzDto instance by subtracting a TimeDto
// from the value of the current DateTzDto Instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  minusTimeDto TimeDto - A TimeDto instance consisting of time components
//                         (years, months, weeks, days, hours, minutes etc.)
//                         which will be subtracted from the date time value
//                         of the current DateTzDto instance.
//
//      type TimeDto struct {
//        Years          int // Number of Years
//        Months         int // Number of Months
//        Weeks          int // Number of Weeks
//        WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//        DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//        Hours          int // Number of Hours.
//        Minutes        int // Number of Minutes
//        Seconds        int // Number of Seconds
//        Milliseconds   int // Number of Milliseconds
//        Microseconds   int // Number of Microseconds
//        Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//        TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//              // plus remaining Nanoseconds
//      }
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//          encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  err := dtz.AddMinusTimeDtoToThis(minusTimeDto)
//
func (dtz *DateTzDto) AddMinusTimeDtoToThis(minusTimeDto TimeDto) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddMinusTimeDtoToThis() "

	dTzUtil := dateTzDtoUtility{}

	dtz2, err := dTzUtil.addMinusTimeDto(
		dtz,
		minusTimeDto,
		ePrefix)

	if err != nil {
		return err
	}

	dTzUtil.copyIn(dtz, &dtz2)

	return nil
}

// AddPlusTimeDto - Creates and returns a new DateTzDto by adding a TimeDto
// to the value of the current DateTzDto instance and returning that new
// value as an of type DateTzDto. The value of the current DateTzDto instance
// will not be altered.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  plusTimeDto TimeDto - A TimeDto instance consisting of time components
//                        (years, months, weeks, days, hours, minutes etc.)
//                        which will be added to the date time value of the
//                        current DateTzDto instance and returned as an instance
//                        of type DateTzDto. Note: The value of the current DateTzDto
//                        will not be altered.
//
//      type TimeDto struct {
//        Years    int // Number of Years
//        Months    int // Number of Months
//        Weeks    int // Number of Weeks
//        WeekDays  int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//        DateDays  int // Total Number of Days. Weeks x 7 plus WeekDays
//        Hours    int // Number of Hours.
//        Minutes    int // Number of Minutes
//        Seconds    int // Number of Seconds
//        Milliseconds  int // Number of Milliseconds
//        Microseconds  int // Number of Microseconds
//        Nanoseconds  int // Remaining Nanoseconds after Milliseconds & Microseconds
//        TotSubSecNanoseconds  int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                  //    plus remaining Nanoseconds
//      }
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a new, valid, fully populated
//              DateTzDto type updated to reflect the added input parameter
//              'plusTimeDto'.
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  dtz2, err := dtz.AddPlusTimeDto(plusTimeDto)
//
func (dtz *DateTzDto) AddPlusTimeDto(plusTimeDto TimeDto) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddPlusTimeDto() "

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.addPlusTimeDto(
								dtz,
								plusTimeDto,
								ePrefix)
}

// AddPlusTimeDtoToThis - Modifies the current DateTzDto instance by adding a TimeDto
// to the value of the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  plusTimeDto TimeDto - A TimeDto instance consisting of time components
//                        (years, months, weeks, days, hours, minutes etc.)
//                        which will be added to the date time value of the
//                        current DateTzDto instance. Note: The value of the
//                        current DateTzDto will be modified.
//
//      type TimeDto struct {
//        Years    int // Number of Years
//        Months    int // Number of Months
//        Weeks    int // Number of Weeks
//        WeekDays  int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//        DateDays  int // Total Number of Days. Weeks x 7 plus WeekDays
//        Hours    int // Number of Hours.
//        Minutes    int // Number of Minutes
//        Seconds    int // Number of Seconds
//        Milliseconds  int // Number of Milliseconds
//        Microseconds  int // Number of Microseconds
//        Nanoseconds  int // Remaining Nanoseconds after Milliseconds & Microseconds
//        TotSubSecNanoseconds  int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                  //    plus remaining Nanoseconds
//      }
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//          encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  err := dtz.AddPlusTimeDtoToThis(plusTimeDto)
//
func (dtz *DateTzDto) AddPlusTimeDtoToThis(plusTimeDto TimeDto) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddPlusTimeDtoToThis() "

	dTzUtil := dateTzDtoUtility{}

	dtz2, err := dTzUtil.addPlusTimeDto(
		dtz,
		plusTimeDto,
		ePrefix)

	if err != nil {
		return err
	}

	dTzUtil.copyIn(dtz, &dtz2)

	return nil
}

// AddTime - Adds input parameter time components (hours, minutes, seconds etc.)
// to the date time value of the current DateTzDto instance. The resulting updated
// date time value is returned to the calling function in the form of a new DateTzDto
// instance. The value of the current DateTzDto instance is NOT altered.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//..hours             int   - Number of hours to add.
//..minutes           int   - Number of minutes to add.
//..seconds           int   - Number of seconds to add.
//..milliseconds      int   - Number of milliseconds to add.
//..microseconds      int   - Number of microseconds to add.
//..nanoseconds       int   - Number of nanoseconds to add.
//
//..Note: Time Component input parameters may be either negative
//        or positive. Negative values will subtract time from
//        the current DateTzDto instance.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful the method returns a valid, fully populated
//              DateTzDto type updated to reflect the added time value
//              input parameters.
//
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  dtz2, err := dtz.AddTime(
//                hours,
//                minutes,
//                seconds,
//                milliseconds,
//                microseconds,
//                nanoseconds,
//                FmtDateTimeYrMDayFmtStr)
//
//  Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//        'constantsdatetime.go'.
//
func (dtz *DateTzDto) AddTime(
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFormatStr string) (DateTzDto, error) {

		dtz.lock.Lock()
		defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddTime() "

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.addTime(
					dtz,
					hours,
					minutes,
					seconds,
					milliseconds,
					microseconds,
					nanoseconds,
					dateTimeFormatStr,
					ePrefix)
}

// AddTimeToThis - Modifies the current DateTzTdo instance by adding input
// parameter time components (hours, minutes, seconds etc.) to the current
// value.
//
// Note: This method WILL alter the value of the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  hours        int - Number of hours to add.
//  minutes      int - Number of minutes to add.
//  seconds      int - Number of seconds to add.
//  milliseconds int - Number of milliseconds to add.
//  microseconds int - Number of microseconds to add.
//  nanoseconds  int - Number of nanoseconds to add.
//
//  Note: Time Component input parameters may be either negative
//        or positive. Negative values will subtract time from
//        the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//          encountered this error Type will encapsulate an error message.
//
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  err := dtz.AddTimeToThis(
//                hours,
//                minutes,
//                seconds,
//                milliseconds,
//                microseconds,
//                nanoseconds,
//                FmtDateTimeYrMDayFmtStr)
//
//  Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//        'constantsdatetime.go'.
//
func (dtz *DateTzDto) AddTimeToThis(
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds int) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.AddTimeToThis() "

	dTzUtil := dateTzDtoUtility{}

	dtz2, err := dTzUtil.addTime(
					dtz,
					hours,
					minutes,
					seconds,
					milliseconds,
					microseconds,
					nanoseconds,
					dtz.dateTimeFmt,
					ePrefix)

	if err != nil {
		return err
	}

	dTzUtil.copyIn(dtz, &dtz2)

	return nil
}

// CopyIn - Receives an incoming DateTzDto and copies those data
// fields to the current DateTzDto instance.
//
// When completed, the current DateTzDto will be equal in all
// respects to the incoming DateTaDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  dtz2 DateTzDto  - A DateTzDto instance. This data will be copied
//                    into the data fields of the current DateTzDto
//                    instance.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   None
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  dtz.CopyIn(dtz2)
//
//  Note: dtz and dtz2 are now equivalent.
//
func (dtz *DateTzDto) CopyIn(dtz2 DateTzDto) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	dTzUtil := dateTzDtoUtility{}

	dTzUtil.copyIn(dtz, &dtz2)

	return
}

// copyOut - returns a DateTzDto instance
// which represents a deep copy of the current
// DateTzDto object.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  None
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - This method returns a new, valid, fully populated DateTzDto
//              which is a deep copy of the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  dtz2 := dtz.copyOut()
//
//  Note: dtz and dtz2 are now equivalent.
//
func (dtz *DateTzDto) CopyOut() DateTzDto {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.copyOut(dtz)
}

// Empty - sets all values of the current DateTzDto
// instance to their uninitialized or zero state.
func (dtz *DateTzDto) Empty() {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	dTzUtil := dateTzDtoUtility{}

	dTzUtil.empty(dtz)

	return
}

// Equal - Returns 'true' if input DateTzDto is equal
// in all respects to the current DateTzDto instance.
//
// Otherwise, the method returns 'false'.
//
func (dtz *DateTzDto) Equal(dtz2 DateTzDto) bool {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	if dtz.tagDescription != dtz2.tagDescription ||
		!dtz.timeComponents.Equal(dtz2.timeComponents) ||
		!dtz.dateTimeValue.Equal(dtz2.dateTimeValue) ||
		dtz.dateTimeFmt != dtz2.dateTimeFmt ||
		!dtz.timeZone.Equal(dtz2.timeZone) {

		return false
	}

	return true
}

// EqualUtcOffset - Compares a second instance of 'DateTzDto' to the
// current 'DateTzDto' object and returns a boolean value signaling
// whether the two objects have the same UTC offsets.
//
// If the return value is true, it signals that both 'DateTzDto'
// instances have the same UTC offset value.
//
func (dtz *DateTzDto) EqualUtcOffset(dtz2 DateTzDto) (bool, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.EqualUtcOffset() "

	dtzDateTimeStr := dtz.dateTimeValue.Format(FmtDateTimeYMDHMSTz)

	dtzUtcOffsetAry := strings.Split(dtzDateTimeStr, " ")

	if len(dtzUtcOffsetAry) != 4 {
		return false, fmt.Errorf(ePrefix +
			"Error: Current DateTzDto is INVALID!\n" +
			"Date Time String='%v'", dtzDateTimeStr)
	}

	dtzUtcOffset := dtzUtcOffsetAry[2]

	dtz2DateTimeStr := dtz2.dateTimeValue.Format(FmtDateTimeYMDHMSTz)

	dtz2UtcOffsetAry := strings.Split(dtz2DateTimeStr, " ")

	if len(dtz2UtcOffsetAry) != 4 {
		return false, fmt.Errorf(ePrefix +
			"\nError: Input parameter 'dtz2' is INVALID!\n" +
			"dtz2 Time String='%v'\n", dtz2DateTimeStr)
	}

	dtz2UtcOffset := dtz2UtcOffsetAry[2]

	return dtzUtcOffset == dtz2UtcOffset, nil
}

// GetConvertibleTzAbbreviation - Returns the time zone abbreviation
// for the Convertible Time Zone. The time zone abbreviation for a
// given time zone is also referred to as the 'zone name'.
//
// The Time Zone abbreviation may be  a series of characters,
// like "EST", "CST" and "PDT" - or - if a time zone alphabetic,
// text abbreviation does not exist, the time zone abbreviation
// might be listed simply as the UTC offset ('+04' or '+03').
//
func (dtz *DateTzDto) GetConvertibleTzAbbreviation() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.timeZone.GetConvertibleZoneName()
}

// GetConvertibleTzLocationPtr - Returns a pointer to the Time Zone
// Location for the Convertible Time Zone. The return value is
// of type '*time.Location'.
//
func (dtz *DateTzDto) GetConvertibleTzLocationPtr() *time.Location {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.timeZone.GetConvertibleLocationPtr()
}

// GetConvertibleTzName - Returns a string containing the
// time zone name for the Convertible Time Zone.
//
// The Time Zone Name is also referred to as the 'Location'
// Name.
//
// If this is an IANA time zone, the full IANA Time
// Zone text name will be returned. If this is a Military
// Time Zone, the equivalent IANA Time Zone name will
// be returned. (See DateTzDto.GetMilitaryTzName() )
//
func (dtz *DateTzDto) GetConvertibleTzName() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return  dtz.timeZone.GetConvertibleLocationName()

}

// GetConvertibleTzUtcOffset - Returns the UTC Offset for
// the Convertible Time Zone associated with this date time.
//
// The UTC offset is formatted as shown in the following
// Examples:
//
//   "+0600"
//   "+0500"
//   "-0500"
//   "-0430"
//
func (dtz *DateTzDto) GetConvertibleTzUtcOffset() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.timeZone.GetConvertibleUtcOffset()
}

// GetDateTimeValue - Returns DateTzDto private member variable
// 'dateTimeValue' as a type time.Time.
//
func (dtz *DateTzDto) GetDateTimeValue() time.Time {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue
}

// GetDateTimeEverything - Receives a time value and formats as
// a date time string in the format:
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: Saturday April 29, 2017 19:54:30.123456489 -0500 CDT
//
func (dtz *DateTzDto) GetDateTimeEverything() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Format(FmtDateTimeEverything)
}

// GetDateTimeNanoSecText - Returns formatted
// date time string with nanoseconds
// 	EXAMPLE: 2006-01-02 15:04:05.000000000
//
func (dtz *DateTzDto) GetDateTimeNanoSecText() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	// Time Format down to the nanosecond
	return dtz.dateTimeValue.Format(FmtDateTimeNanoSecondStr)
}

// GetDateTimeFmt - Returns the DateTzDto private member
// variable, DateTzDto.dateTimeFmt.
//
func (dtz *DateTzDto) GetDateTimeFmt() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeFmt
}

// GetDateTimeSecText - Returns formatted
// date time with seconds for display.
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: 2006-01-02 15:04:05
//
func (dtz *DateTzDto) GetDateTimeSecText() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	// Time Display Format with seconds
	return dtz.dateTimeValue.Format(FmtDateTimeSecText)
}

// GetDateTimeStr - Returns a date time string
// in the format '20170427211307'. Useful in naming
// files.
func (dtz *DateTzDto) GetDateTimeStr() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	// Time Format down to the second
	return dtz.dateTimeValue.Format(FmtDateTimeSecondStr)

}

// GetDateTimeTzNanoSecDowYMDText - Outputs date time in string format using
// the FmtDateTimeTzNanoDowYMD format which incorporates date time to the
// nano second and the associated time zone. In this format, the date is
// expressed as Year-Month-Day (Example: 2017-12-06). The string is
// prefixed with the day of the week:
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: Monday 2006-01-02 15:04:05.000000000 -0700 MST
//
func (dtz *DateTzDto) GetDateTimeTzNanoSecDowYMDText() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Format(FmtDateTimeTzNanoDowYMD)
}

// GetDateTimeTzNanoSecText - Outputs date time in string format using
// the FmtDateTimeDMYNanoTz format which incorporates date time to nano seconds
// and the associated time zone.
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: 01/02/2006 15:04:05.000000000 -0700 MST
//
func (dtz *DateTzDto) GetDateTimeTzNanoSecText() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Format(FmtDateTimeDMYNanoTz)
}

// GetDateTimeTzNanoSecYMDDowText - Outputs date time in string format using
// the FmtDateTimeTzNanoYMDDow format which incorporates date time to nano seconds
// and the associated time zone. In this format, the date is expressed as
// Year-Month-Day (Example: 2017-12-06) followed by the day of the week.
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: 2006-01-02 Monday 15:04:05.000000000 -0700 MST
//
func (dtz *DateTzDto) GetDateTimeTzNanoSecYMDDowText() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Format(FmtDateTimeTzNanoYMDDow)
}

// GetDateTimeTzNanoSecYMDText - Outputs date time in string format using
// the FmtDateTimeTzNanoYMD format which incorporates date time to nano seconds
// and the associated time zone. In this format, the date is expressed as
// Year-Month-Day (2017-12-06)
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: 2006-01-02 15:04:05.000000000 -0700 MST
//
func (dtz *DateTzDto) GetDateTimeTzNanoSecYMDText() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Format(FmtDateTimeTzNanoYMD)
}

// GetDateTimeYMDAbbrvDowNano - Outputs date time in string format using
// the FmtDateTimeYMDAbbrvDowNano format which incorporates date time to the
// nano second and the associated time zone. In this format, the date is
// expressed as Year-Month-Day (Example: 2017-12-06). The string includes
// the abbreviated (limited to 3-characters) day of the week:
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
//
func (dtz *DateTzDto) GetDateTimeYMDAbbrvDowNano() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Format(FmtDateTimeYMDAbbrvDowNano)
}

// GetDateTimeYrMDayTzFmtStr - Returns a date time string
// formatted as year-mth-day time and time zone.
// FmtDateTimeYrMDayFmtStr - Year Month Day Date Format String
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: "2006-01-02 15:04:05.000000000 -0700 MST"
func (dtz *DateTzDto) GetDateTimeYrMDayTzFmtStr() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Format(FmtDateTimeYrMDayFmtStr)
}

// GetOriginalTagDescription - Returns DateTzDto private member
// variable, DateTzDto.tagDescription.
//
// 'tagDescription' is available to users for use as
// a tag, label, classification or text description.
//
func (dtz *DateTzDto) GetOriginalTagDescription() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.tagDescription
}

// GetOriginalTzAbbreviation - Returns the time zone abbreviation
// for the Original Time Zone.
//
// The Time Zone abbreviation may be  a series of characters,
// like "EST", "CST" and "PDT" - or - if a time zone alphabetic,
// text abbreviation does not exist, the time zone abbreviation
// might be listed simply as the UTC offset ('+04' or '+03').
//
func (dtz *DateTzDto) GetOriginalTzAbbreviation() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.timeZone.GetOriginalZoneName()
}

// GetOriginalTzLocationPtr - Returns a pointer to the Time Zone
// Location for the Original Time Zone. The return value is
// of type '*time.Location'.
//
func (dtz *DateTzDto) GetOriginalTzLocationPtr() *time.Location {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.timeZone.GetOriginalLocationPtr()
}

// GetOriginalTzName - Returns a string containing the
// time zone name for the Original Time Zone.
//
// The Time Zone Name is also referred to as the 'Location'
// Name.
// If this is an IANA time zone, the full IANA Time
// Zone text name will be returned. If this is a Military
// Time Zone, the equivalent IANA Time Zone name will
// be returned. (See DateTzDto.GetMilitaryTzName() )
//
func (dtz *DateTzDto) GetOriginalTzName() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return  dtz.timeZone.GetOriginalLocationName()

}

// GetOriginalTzUtcOffset - Returns the UTC Offset for
// the Original Time Zone associated with this date time.
//
// The UTC offset is formatted as shown in the following
// Examples:
//
//   "+0600"
//   "+0500"
//   "-0500"
//   "-0430"
//
func (dtz *DateTzDto) GetOriginalTzUtcOffset() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.timeZone.GetOriginalUtcOffset()
}

// GetMilitaryCompactDateTimeGroup - Outputs date time string formatted for
// standard U.S.A. Military date time also referred to as the Military
// Date Time Group (DTG). This form of the Date Time Group is configured
// as the 'Compact' Date Time Group. This means there are no spaces between
// the date time elements.
//
// This "Compact Date Time Group" format differs from the "Open Date Time
// Group" format returned by method DateTzDto.GetMilitaryOpenDateTimeGroup().
// This "Compact Date Time Group" deletes spaces between critical time
// components. The "Open Date Time Group" uses the same basic format but
// inserts spaces between critical date time components.
//
// Note: The Compact Date Time Group is only applicable to Military Time Zones.
// If the current time zone is not configured as a Military Time Zone,
// an error will be returned.
//
// Reference:
//    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//    https://www.timeanddate.com/time/zones/z
//    http://blog.refactortactical.com/blog/military-date-time-group/
//
// Military 2-digit year format or "Date Time Group" is traditionally
// formatted as DDHHMM(Z)MONYY, where 'Z' is the Military Time Zone.
//
// EXAMPLES:
//
//    "011815ZJAN11" = 01/01/2011 18:15 +0000 Zulu
//
//     630pm on January 6th, 2012 in Fayetteville NC would read '061830RJAN12'
//
func (dtz *DateTzDto) GetMilitaryCompactDateTimeGroup() (string, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "TimeZoneDefinition.GetMilitaryCompactDateTimeGroup() "

	tzType := dtz.timeZone.GetOriginalTimeZoneType()

	if tzType != TzType.Military() {
		return "",
			fmt.Errorf(ePrefix +
				"\nError: This Time Zone for this DateTzDto instance is NOT configured as a\n" +
				"Military Time Zone. The Compact Date Time Group is only applicable to \n" +
				"Military Time Zones. Therefore, this time zone is invaid as a Military\n" +
				"Time Zone.\n" +
				"TimeZoneDefinition Time Zone Type='%v'\n", tzType.String())
	}

	milTzLetter, err := dtz.timeZone.GetMilitaryTimeZoneLetter()

	if err != nil {
		return "",
			fmt.Errorf(ePrefix +
				"Error returned by dtz.timeZone.GetMilitaryTimeZoneLetter()\n" +
				"Error='%v'\n", err.Error())
	}

	fmtDateTime := dtz.dateTimeValue.Format("021504" + milTzLetter + "Jan06")

	fmtDateTime = strings.ToUpper(fmtDateTime)

	return fmtDateTime, nil
}

// GetMilitaryOpenDateTimeGroup - Outputs date time string formatted for
// standard U.S.A. Military date time also referred to as the Military
// Date Time Group (DTG). This form of the Date Time Group is configured
// as the 'Open', easy to read, Date Time Group. This means that spaces
// are inserted between the critical date time components.
//
// This "Open Date Time Group" format differs from the "Compact Date Time
// Group" format returned by method DateTzDto.GetMilitaryCompactDateTimeGroup().
// This "Open Date Time Group" format inserts spaces between critical date
// time components in order to improve overall readability. The "Compact
// Date Time Group" uses the same basic format but removes all internal
// spaces.
//
// Note: The Open Date Time Group is only applicable to Military Time Zones.
// If the current time zone is not configured as a Military Time Zone,
// an error will be returned.
//
// Reference:
//    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//    https://www.timeanddate.com/time/zones/z
//    http://blog.refactortactical.com/blog/military-date-time-group/
//
// The Military 2-digit year format or "Date Time Group" is traditionally
// formatted as DDHHMM(Z)MONYY, where 'Z' is the Military Time Zone.
// The "Open Date Time Group" format inserts spaces between critical
// date time components as shown in the following examples.
//
// EXAMPLES:
//
//    "01 1815Z JAN 11" = 01/01/2011 18:15 +0000 Zulu
//
//     630pm on January 6th, 2012 in Fayetteville NC would read '06 1830R JAN 12'
//
func (dtz *DateTzDto) GetMilitaryOpenDateTimeGroup() (string, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "TimeZoneDefinition.GetMilitaryOpenDateTimeGroup() "

	tzType := dtz.timeZone.GetOriginalTimeZoneType()

	if tzType != TzType.Military() {
		return "",
			fmt.Errorf(ePrefix +
				"\nError: This Time Zone for this DateTzDto instance is NOT configured as a\n" +
					"Military Time Zone. The Open Date Time Group is only applicable to \n" +
					"Military Time Zones. Therefore, this time zone is invalid as a Military\n" +
					"Time Zone.\n" +
					"TimeZoneDefinition Time Zone Type='%v'\n", tzType.String())
	}

	milTzLetter, err := dtz.timeZone.GetMilitaryTimeZoneLetter()

	if err != nil {
		return "",
			fmt.Errorf(ePrefix +
				"Error returned by dtz.timeZone.GetMilitaryTimeZoneLetter()\n" +
				"Error='%v'\n", err.Error())
	}

	fmtDateTime :=dtz.dateTimeValue.Format("02 1504" + milTzLetter + " Jan 06")

	fmtDateTime = strings.ToUpper(fmtDateTime)

	return fmtDateTime, nil
}

// GetTimeComponents - Returns a deep copy of DateTzDto
// private member variable DateTzDto.timeComponents.
// The private member variable is returned as a type
// 'TimeDto'.
//
func (dtz *DateTzDto) GetTimeComponents() TimeDto {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.timeComponents.CopyOut()
}

// GetTimeDto - Converts the current DateTzDto instance
// date time information into an instance of TimeDto
// and returns that TimeDto to the caller.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  None.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
// TimeDto - A TimeDto structure is defined as follows:
//
//    type TimeDto struct {
//       Years                int // Number of Years
//       Months               int // Number of Months
//       Weeks                int // Number of Weeks
//       WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//       DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//       Hours                int // Number of Hours.
//       Minutes              int // Number of Minutes
//       Seconds              int // Number of Seconds
//       Milliseconds         int // Number of Milliseconds
//       Microseconds         int // Number of Microseconds
//       Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//       TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                //  plus remaining Nanoseconds
//       TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                //  + Seconds + Milliseconds + Nanoseconds
//    }
//
//
// error - If successful the returned error Type is set equal to 'nil'. If errors are
//         encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) GetTimeDto() (TimeDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.GetTimeDto() "

	tDto, err := TimeDto{}.NewFromDateTime(dtz.dateTimeValue)

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix+
			"\nError returned by TimeDto{}.NewFromDateTime(dtz.DateTime)\n"+
			"dtz.DateTime ='%v'\nError='%v'\n",
			dtz.dateTimeValue.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return tDto, nil
}

// GetTimeStampEverything - Generates and returns a time stamp as
// type string. The time stamp is formatted using the format,
// 'FmtDateTimeEverything'.
//
// ------------------------------------------------------------------------
//
//  Example output:
//    "Saturday April 29, 2017 19:54:30.123456489 -0500 CDT"
func (dtz *DateTzDto) GetTimeStampEverything() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Format(FmtDateTimeEverything)
}

// GetTimeStampYMDAbbrvDowNano - Generates and returns a time stamp as
// type string. The time stamp is formatted using the format
// 'FmtDateTimeYMDAbbrvDowNano'.
//
// ------------------------------------------------------------------------
//
//  Example Output:
//  "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
func (dtz *DateTzDto) GetTimeStampYMDAbbrvDowNano() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Format(FmtDateTimeYMDAbbrvDowNano)
}

// GetTimeZoneDef - Returns a deep copy of the 'DateTzDto' private
// member variable, 'timeZone', of type TimeZoneDefinition.
//
func (dtz *DateTzDto) GetTimeZoneDef() TimeZoneDefinition {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.timeZone.CopyOut()
}

// IsEmpty - Analyzes the current DateTzDto instance to determine
// if the instance is in an 'EMPTY' or uninitialized state.
//
// If the current DateTzDto instance is found to be 'EMPTY', this
// method returns 'true'. Otherwise, if the instance is 'NOT EMPTY',
// this method returns 'false'.
//
func (dtz *DateTzDto) IsEmpty() bool {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	if dtz.tagDescription == "" &&
		dtz.timeComponents.IsEmpty() &&
		dtz.dateTimeValue.IsZero() &&
		dtz.dateTimeFmt == "" &&
		dtz.timeZone.IsEmpty() {

		return true

	}

	return false
}

// IsValid - Analyzes the current DateTzDto instance and returns
// an error, populated with an appropriate error message, if the instance
// is found to be INVALID.
//
// If the current DateTzDto instance is VALID, this method returns
// nil.
func (dtz *DateTzDto) IsValid() error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.IsValid() "

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.isValidDateTzDto(dtz, ePrefix)
}

// New - returns a new DateTzDto instance based on a time.Time ('dateTime')
// input parameter. The Time Zone Location is extracted from input parameter
// 'dateTime'.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  dateTime    time.Time   - A date time value
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  DateTzDto - If successful this method returns a new DateTzDto instance.
//
//
//
//  error     - If successful the returned error Type is set equal to 'nil'. If errors are
//              encountered this error Type will encapsulate an error message.
//
//
// ------------------------------------------------------------------------
//
// Usage
//
//   dtzDto, err := DateTzDto{}.New(dateTime, FmtDateTimeYrMDayFmtStr)
//
//
//   Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//         'constantsdatetime.go'.
//
func (dtz DateTzDto) New(
	dateTime time.Time,
	dateTimeFmtStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.New() "

	if dateTime.IsZero() {
		return DateTzDto{}, errors.New(ePrefix +
			"\nError: Input parameter dateTime is Zero value!\n")
	}

	dTzUtil := dateTzDtoUtility{}

	dtz2 := DateTzDto{}

	err := dTzUtil.setFromDateTime( &dtz2, dateTime, dateTimeFmtStr, ePrefix)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+
			"\nError returned from dTzUtil.setFromDateTime( &dtz2, dateTime, dateTimeFmtStr, ePrefix).\n" +
			"dateTime='%v'\nError='%v'\n", dateTime, err.Error())
	}

	return dtz2, nil
}

// NewDateTimeComponents - creates a new DateTzDto object and populates the
// data fields based on input parameters.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//   year               int  - year number
//   month              int  - month number       1 - 12
//   day                int  - day number         1 - 31
//   hour               int  - hour number        0 - 24
//   minute             int  - minute number      0 - 59
//   second             int  - second number      0 - 59
//   millisecond        int  - millisecond number 0 - 999
//   microsecond        int  - microsecond number 0 - 999
//   nanosecond         int  - nanosecond number  0 - 999
//
//   timeZoneLocation string - time zone location must be designated as one of
//                             two values:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                location for the host computer.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//                 The source file 'constantsdatetime.go' contains a number of
//                 constant declarations covering the more frequently used time
//                 zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TzIana'.
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   DateTzDto - If successful this method returns a new DateTzDto instance.
//
//
//   error - If successful the returned error Type is set equal to 'nil'. If errors are
//           encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//      dtzDto, err := DateTzDto{}.New(
//                        year,
//                        month,
//                        day,
//                        hour,
//                        min,
//                        sec,
//                        nanosecond,
//                        TZones.US.Central(),
//                        FmtDateTimeYrMDayFmtStr)
//
//   Note: TZones.US.Central() = "America/Chicago"
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//         source file 'constantsdatetime.go'.
//
func (dtz DateTzDto) NewDateTimeComponents(
	year,
	month,
	day,
	hour,
	minute,
	second,
	millisecond,
	microsecond,
	nanosecond int,
	timeZoneLocation,
	dateTimeFmtStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.NewDateTimeComponents() "

	dtz2 := DateTzDto{}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromDateTimeComponents(
		&dtz2,
		year,
		month,
		day,
		hour,
		minute,
		second,
		millisecond,
		microsecond,
		nanosecond,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// NewDateTimeElements - creates a new DateTzDto object and populates
// the data fields based on date time elements.
//
// Date Time elements include year, month, day, hour, minute, second and
// nanosecond.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//   year                int - year number
//   month               int - month number       1 - 12
//   day                 int - day number         1 - 31
//   hour                int - hour number        0 - 24
//   minute              int - minute number      0 - 59
//   second              int - second number      0 - 59
//   nanosecond          int - nanosecond number  0 - 999,999,999
//
//   timeZoneLocation string - time zone location must be designated as one of
//                             two values:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                location for the host computer.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//                 The source file 'constantsdatetime.go' contains a number of
//                 constant declarations covering the more frequently used time
//                 zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TzIana'.
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   DateTzDto - If successful, this method returns a new, populated 'DateTzDto'
//               instance.
//
//
//   error     - If successful the returned error Type is set equal to 'nil'. If errors are
//               encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//      dtzDto, err := DateTzDto{}.NewDateTimeElements(
//         year,
//         month,
//         day,
//         hour,
//         minute,
//         second,
//         nanosecond ,
//         TZones.US.Central(),
//         FmtDateTimeYrMDayFmtStr)
//
//   Note: TZones.US.Central() = "America/Chicago"
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//         source file 'constantsdatetime.go'.
//
func (dtz DateTzDto) NewDateTimeElements(
		year,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond int,
		timeZoneLocation,
		dateTimeFmtStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.NewDateTimeElements() "

	dtz2 := DateTzDto{}

	dtUtil := dateTzDtoUtility{}

	err := dtUtil.setFromDateTimeElements(
		&dtz2,
		year,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// NewNowLocal - Creates and returns a new DateTzDto instance based on a date
// time value which is automatically assigned by time.Now(). The time zone 'Local'
// is used by the Go Programming Language to assign the time zone configured
// on the host computer executing this code. Effectively, this means that the
// time selected is equal to the current value of the host computer clock.
//
// The Time Zone Location is automatically set to 'Local'.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   DateTzDto - If successful, this method returns a new DateTzDto instance.
//               The data fields of this new instance are initialized to zero
//               values.
//
//
//   error     - If successful the returned error Type is set equal to 'nil'. If errors are
//               encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//   dtzDto, err := DateTzDto{}.NewNowLocal(FmtDateTimeYrMDayFmtStr)
//
//   Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         FmtDateTimeYrMDayFmtStr' is a constants available in source file
//         'constantsdatetime.go'.
//
func (dtz DateTzDto) NewNowLocal(
	dateTimeFmtStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.NewNowLocal() "

	dt := time.Now().Local()

	dTz := DateTzDto{}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromDateTime(
		&dTz,
		dt,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dTz, nil
}

// NewNowTz - returns a new DateTzDto instance based on a date time value
// which is automatically assigned by time.Now(). Effectively, this means
// that the time is set equal to the current value of the host computer
// clock.
//
// The user is required to provide an input parameter, 'timeZoneLocation',
// which is used to configure the date time value. In essence, the current
// local time is converted to the timezone specified by 'timeZoneLocation'.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//   timeZoneLocation string - time zone location must be designated as one of
//                             two values:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                location for the host computer.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//                 The source file 'constantsdatetime.go' contains a number of
//                 constant declarations covering the more frequently used time
//                 zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TzIana'.
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   DateTzDto - If successful, this method returns a new DateTzDto instance.
//               The data fields of this new instance are initialized to zero
//               values.
//
//
//   error     - If successful the returned error Type is set equal to 'nil'. If errors are
//               encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//      dtzDto, err := DateTzDto{}.NewNowTz(
//         TZones.US.Central(),
//         FmtDateTimeYrMDayFmtStr)
//
//   Note: TZones.US.Central() = "America/Chicago"
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//         source file 'constantsdatetime.go'.
//
func (dtz DateTzDto) NewNowTz(
	timeZoneLocation,
	dateTimeFmtStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.NewNowTz() "

	dt := time.Now().Local()

	dTz := DateTzDto{}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromTimeTz(
		&dTz,
		dt,
		timeZoneLocation,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dTz, nil
}

// NewNowUTC - returns a new DateTzDto instance based on a date time value
// which is automatically assigned by time.Now(). Effectively, this means
// that the time selected is equal to the current value of the host computer
// clock.
//
// The Time Zone Location is automatically set to 'UTC'. UTC refers to Universal
// Coordinated Time and is sometimes referred to as 'Zulu', GMT or Greenwich Mean
// Time.
//
// Reference Universal Coordinated Time:
//   https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// The net effect is that the current local time as provided by the host computer
// is converted into Universal Coordinated Time ('UTC').
//
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   DateTzDto - If successful, this method returns a new DateTzDto instance.
//               The data fields of this new instance are initialized to zero
//               values.
//
//
//   error     - If successful the returned error Type is set equal to 'nil'. If errors are
//               encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//   dtzDto, err := DateTzDto{}.NewNowUTC(
//                      FmtDateTimeYrMDayFmtStr)
//
//   Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//         'constantsdatetime.go'.
//
func (dtz DateTzDto) NewNowUTC(
	dateTimeFmtStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.NewNowUTC() "

	dt := time.Now().In(time.UTC)

	dTz := DateTzDto{}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromDateTime(
		&dTz,
		dt,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dTz, nil
}

// NewTimeDto - Receives input parameters type TimeDto, 'timeZoneLocation' and 'dateTimeFormatStr'.
// These parameters are used to construct and return a new DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   tDto             TimeDto - Time values used to construct the DateTzDto instance
//
//         type TimeDto struct {
//            Years          int // Number of Years
//            Months         int // Number of Months
//            Weeks          int // Number of Weeks
//            WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//            DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//            Hours          int // Number of Hours.
//            Minutes        int // Number of Minutes
//            Seconds        int // Number of Seconds
//            Milliseconds   int // Number of Milliseconds
//            Microseconds   int // Number of Microseconds
//            Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//            TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                     // plus remaining Nanoseconds
//         }
//
//
//   timeZoneName string - time zone location must be designated as one of
//                             two values:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                location for the host computer.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//                 The source file 'constantsdatetime.go' contains a number of
//                 constant declarations covering the more frequently used time
//                 zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TzIana'.
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   DateTzDto - If successful, this method returns a new DateTzDto instance.
//               The data fields of this new instance are initialized to zero
//               values.
//
//
//   error     - If successful the returned error Type is set equal to 'nil'. If errors are
//               encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//   dtz, err := DateTzDto{}.NewTimeDto(
//            tDto,
//            TZones.US.Central(),
//            FmtDateTimeYrMDayFmtStr)
//
//   Note: TZones.US.Central() = "America/Chicago"
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//         source file 'constantsdatetime.go'.
//
func (dtz DateTzDto) NewTimeDto(
	tDto TimeDto,
	timeZoneName string,
	dateTimeFormatStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.NewTimeDto() "

	if len(timeZoneName) == 0 {
		return DateTzDto{},
		&InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'timeZoneName' is an empty string!",
			err:                 nil,
		}
	}

	dtz2 := DateTzDto{}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromTimeDto(
		&dtz2,
		tDto,
		timeZoneName,
		dateTimeFormatStr,
		ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// NewTz - returns a new DateTzDto instance based on a time.Time input parameter ('dateTime').
// The caller is required to provide a Time Zone Location or Name. Input parameter 'dateTime'
// will be converted to this Time Zone before storing the converted 'dateTime' in the newly
// created DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   dateTime          time.Time - A date time value
//
//   timeZoneName         string - time zone location must be designated as one of
//                                 two values:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                location for the host computer.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//                 The source file 'constantsdatetime.go' contains a number of
//                 constant declarations covering the more frequently used time
//                 zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TzIana'.
//
//   timeZoneConversionType TimeZoneConversionType
//                           - TimeZoneConversionType is an enumeration with two
//                             possible values: TimeZoneConversionType(0).Absolute()
//                             or TimeZoneConversionType(0).Relative(). The 'Absolute'
//                             conversion will produce same time value in a different
//                             time zone. The 'Relative' conversion type will yield
//                             a different equivalent time in a different time zone.
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   DateTzDto - If successful, this method returns a new DateTzDto instance.
//               The data fields of this new instance are initialized to zero
//               values.
//
//
//   error     - If successful the returned error Type is set equal to 'nil'. If errors are
//               encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//   dtzDto, err := DateTzDto{}.NewTz(
//         dateTime,
//         TZones.US.Central(),
//         FmtDateTimeYrMDayFmtStr)
//
//   Note: TZones.US.Central() = "America/Chicago"
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants.
//         'TZones.US.Central() is located in 'timezonedata.dto'. 'FmtDateTimeYrMDayFmtStr'
//         is located in source file 'constantsdatetime.go'.
//
func (dtz DateTzDto) NewTz(
	dateTime time.Time,
	timeZoneName string,
	timeZoneConversionType TimeZoneConversionType,
	dateTimeFmtStr string) (DateTzDto, error) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.NewTz() "

	if dateTime.IsZero() {
		return DateTzDto{}, errors.New(ePrefix + "Error: Input parameter dateTime is Zero value!")
	}

	dtz2 := DateTzDto{}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromTimeTz(
		&dtz2,
		dateTime,
		timeZoneName,
		timeZoneConversionType,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DateTzDto{}, err
	}

	return dtz2, nil
}

// SetDateTimeFmt - Sets the DateTzDto data field 'DateTimeFmt'.
// This string is used to format the DateTzDto DateTimeFmt field
// when DateTzDto.String() is called.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   None - This method sets the internal data field DateTzDto.dateTimeFmt
//
func (dtz *DateTzDto) SetDateTimeFmt(dateTimeFmtStr string) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	if len(dateTimeFmtStr) == 0 {
		dateTimeFmtStr = FmtDateTimeYrMDayFmtStr
	}

	dtz.dateTimeFmt = dateTimeFmtStr

}

// SetFromDateTimeComponents - Sets the values of the Date Time fields
// for the current DateTzDto instance based on time components
// and a Time Zone Location.
//
// Note that this variation of time elements breaks time down by
// hour, minute, second, millisecond, microsecond and nanosecond.
//
// See method SetFromDateTimeElements(), above, which uses a slightly
// different set of time components.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   year                int - year number
//   month               int - month number        1 -  12
//   day                 int - day number          1 -  31
//   hour                int - hour number         0 -  24
//   min                 int - minute number       0 -  59
//   sec                 int - second number       0 -  59
//   millisecond         int - millisecond number  0 - 999
//   microsecond         int - microsecond number  0 - 999
//   nanosecond          int - nanosecond number   0 - 999
//
//   timeZoneLocation string - time zone location must be designated as one of
//                             two values:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                location for the host computer.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//                 The source file 'constantsdatetime.go' contains a number of
//                 constant declarations covering the more frequently used time
//                 zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TzIana'.
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   error - If successful the returned error Type is set equal to 'nil'. If errors are
//           encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromDateTimeComponents(
			year,
			month,
			day,
			hour,
			minute,
			second,
			millisecond,
			microsecond,
			nanosecond int,
			timeZoneLocation,
			dateTimeFmtStr string) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.SetFromDateTimeComponents() "

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.setFromDateTimeComponents(
						dtz,
						year,
						month,
						day,
						hour,
						minute,
						second,
						millisecond,
						microsecond,
						nanosecond,
						timeZoneLocation,
						dateTimeFmtStr,
						ePrefix)
}

// SetFromDateTimeElements - Sets the values of the current DateTzDto
// data fields based on input parameters consisting of date time
// elements, a time zone location and a date time format string.
//
// Date Time elements include year, month, day, hour, minute,
// second and nanosecond.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   year                int - year number
//   month               int - month number        1 -  12
//   day                 int - day number          1 -  31
//   hour                int - hour number         0 -  24
//   min                 int - minute number       0 -  59
//   sec                 int - second number       0 -  59
//   millisecond         int - millisecond number  0 - 999
//   microsecond         int - microsecond number  0 - 999
//   nanosecond          int - nanosecond number   0 - 999,999,999
//
//   timeZoneLocation string - time zone location must be designated as one of
//                             two values:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                location for the host computer.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//                 The source file 'constantsdatetime.go' contains a number of
//                 constant declarations covering the more frequently used time
//                 zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TzIana'.
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   error - If successful the returned error Type is set equal to 'nil'. If errors are
//           encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromDateTimeElements(
	year,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	timeZoneLocationName,
	dateTimeFmtStr string) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.SetFromDateTimeElements() "

	dtUtil := dateTzDtoUtility{}

	return dtUtil.setFromDateTimeElements(
			dtz,
			year,
			month,
			day,
			hour,
			minute,
			second,
			nanosecond,
			timeZoneLocationName,
			dateTimeFmtStr,
			ePrefix)
}

// SetFromDateTime - Sets the values of the current DateTzDto fields
// based on an input parameter 'dateTime' (Type time.time) and a
// date time format string.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   dateTime    time.Time   - A date time value
//
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   error - If successful the returned error Type is set equal to 'nil'. If errors are
//           encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromDateTime(dateTime time.Time, dateTimeFmtStr string) error {

	dtz.lock.Lock()
	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.SetFromDateTime() "

	dTzUtility := dateTzDtoUtility{}

	return dTzUtility.setFromDateTime(dtz, dateTime, dateTimeFmtStr, ePrefix)
}

// SetFromTimeDto - Receives data from a TimeDto input parameter
// and sets all data fields of the current DateTzDto instance
// accordingly. When the method completes, the values of the
// current DateTzDto will equal the values of the input parameter
// TimeDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   tDto            TimeDto - A populated TimeDto instance
//                             A TimeDto struct is defined as follows:
//     type TimeDto struct {
//       Years                int // Number of Years
//       Months               int // Number of Months
//       Weeks                int // Number of Weeks
//       WeekDays             int // Number of Week-WeekDays.
//                                //   Total WeekDays/7 + Remainder WeekDays
//       DateDays             int // Total Number of Days.
//                                //   Weeks x 7 plus WeekDays
//       Hours                int // Number of Hours.
//       Minutes              int // Number of Minutes
//       Seconds              int // Number of Seconds
//       Milliseconds         int // Number of Milliseconds
//       Microseconds         int // Number of Microseconds
//       Nanoseconds            int // Remaining Nanoseconds after Milliseconds & Microseconds
//       TotSubSecNanoseconds int // Total Nanoseconds:
//                                //   Millisecond NanoSecs + Microsecond NanoSecs
//                                //   plus remaining Nanoseconds
//     }
//
//
//   timeZoneLocation string - time zone location must be designated as one of
//                             two values:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                location for the host computer.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//                 The source file 'constantsdatetime.go' contains a number of
//                 constant declarations covering the more frequently used time
//                 zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TzIana'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   error - If successful the returned error Type is set equal to 'nil'. If errors are
//           encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromTimeDto(tDto TimeDto, timeZoneLocation string) error {

	ePrefix := "DateTzDto.SetFromTimeDto() "

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.setFromTimeDto(
					dtz,
					tDto,
					timeZoneLocation,
					dtz.dateTimeFmt,
					ePrefix)
}

// SetFromTimeTz - Sets the time values of the current DateTzDto instance
// based on input parameters 'dateTime', 'timeZoneLocation' and a date time
// format string.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   dateTime      time.Time - A date time value
//
//   timeZoneLocation string - time zone location must be designated as one of
//                             two values:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                location for the host computer.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//
//   dateTimeFmtStr string   - A date time format string which will be used
//                             to format and display 'dateTime'. Example:
//                             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                             Date time format constants are found in the source
//                             file 'constantsdatetime.go'. These constants represent
//                             the more commonly used date time string formats. All
//                             Date Time format constants begin with the prefix
//                             'FmtDateTime'.
//
//                             If 'dateTimeFmtStr' is submitted as an
//                             'empty string', a default date time format
//                             string will be applied. The default date time
//                             format string is:
//                               FmtDateTimeYrMDayFmtStr =
//                                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   error - If successful the returned error Type is set equal to 'nil'. If errors are
//           encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromTimeTz(
	dateTime time.Time,
	timeZoneLocation string,
	timeZoneConversionType TimeZoneConversionType,
	dateTimeFmtStr string) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.SetFromTimeTz() "

	dTzUtil := dateTzDtoUtility{}

	return dTzUtil.setFromTimeTz(
		dtz,
		dateTime,
		timeZoneLocation,
		timeZoneConversionType,
		dateTimeFmtStr,
		ePrefix)
}

// SetNewTimeZone - Changes the time zone information for the current
// DateTzDto Date Time.  If the value of input parameter 'newTimeZoneLocation'
// is different from the existing Time Zone Location, all values in the
// current DateTzDto data fields will be replaced with the new date time and
// time zone information.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   newTimeZoneLocation string - Designates the standard Time Zone location used to
//                 compute date time. The existing DateTzDto Date Time will be converted
//                 to an equivalent time this 'newTimeZoneLocation'.
//
//        This 'newTimeZoneLocation' must be designated as one of two values:
//
//        (1) The string 'Local' - signals the designation of the local time zone
//            location for the host computer.
//
//        (2) IANA Time Zone Location -
//            See https://golang.org/pkg/time/#LoadLocation
//            and https://www.iana.org/time-zones to ensure that
//            the IANA Time Zone Database is properly configured
//            on your system. Note: IANA Time Zone Data base is
//            equivalent to 'tz database'.
//
//            Examples:
//              "America/New_York"
//              "America/Chicago"
//              "America/Denver"
//              "America/Los_Angeles"
//              "Pacific/Honolulu"
//
//            The source file 'constantsdatetime.go' contains a number of
//            constant declarations covering the more frequently used time
//            zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//            time zone constants begin with the prefix 'TzIana'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   error - If successful the returned error Type is set equal to 'nil'. If errors are
//           encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetNewTimeZone(newTimeZoneLocation string) error {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.SetNewTimeZone() "

	tzMech := TimeZoneMechanics{}

	tzl := tzMech.PreProcessTimeZoneLocation(newTimeZoneLocation)

	if len(tzl) == 0 {
		return errors.New("Error: Input Parameter, 'newTimeZoneLocation' " +
			"resolved to an empty string!\n")
	}

	dtMech := DTimeMechanics{}

	loc, err := dtMech.LoadTzLocation(tzl, ePrefix)

	if err != nil {
		return err
	}

	newDateTime := dtz.dateTimeValue.In(loc)

	newFmtStr := dtz.dateTimeFmt

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromDateTime(dtz, newDateTime, newFmtStr, ePrefix)

	return err
}

// SetOriginalTagDescription - Sets DateTzDto private member variable
// DateTzDto.tagDescription to the value passed in 'tagDesc'.
//
// The DateTzDto.tagDescription is available to users for use
// as a tag, label, classification or text description.
//
func (dtz *DateTzDto) SetTagDescription(tagDesc string) {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	dtz.tagDescription = tagDesc
}

// String - This method returns the DateTzDto DateTime field value
// formatted as a string. If the DateTzDto data field, 'DateTimeFmt'
// is an empty string, a default format string will be used. The
// default format is:
//
//     FmtDateTimeYrMDayFmtStr =
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
// To set the internal data field, 'DateTzDto.dateTimeFmt' reference
// method DateTzDto.SetDateTimeFmt(), above.
//
func (dtz *DateTzDto) String() string {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	fmtStr := dtz.dateTimeFmt

	if len(fmtStr) == 0 {
		fmtStr = FmtDateTimeYrMDayFmtStr
	}

	return dtz.dateTimeValue.Format(fmtStr)
}

// Sub - Subtracts the DateTime value of the incoming DateTzDto
// from the DateTime value of the current DateTzDto. The result
// is returned as a Type 'time.Duration'.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//   dtz2 DateTzDto - A valid and populated instance of type DateTzDto.
//                    A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefinition // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   time.Duration - A Type time.duration which represents the value of input
//                   parameter 'dtz2' subtracted from the current DateTzDto
//                   instance.
//
func (dtz *DateTzDto) Sub(dtz2 DateTzDto) time.Duration {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Sub(dtz2.dateTimeValue)

}

// SubDateTime - Subtracts a date time value (Type: 'time.Time')
// from the date time value of the current DateTzDto. The result
// is returned as a Type 'time.Duration'.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//   t2  time.Time - A date time value which will be subtracted from the
//                   the time value of the current DateTzDto instance.
//
//
// ------------------------------------------------------------------------
//
// Return XValue
//   time.Duration - A time duration value representing the subtraction of the
//                   input parameter t2 time value from the time value of the
//                   current DateTzDto time value.
//
func (dtz *DateTzDto) SubDateTime(t2 time.Time) time.Duration {

	dtz.lock.Lock()

	defer dtz.lock.Unlock()

	return dtz.dateTimeValue.Sub(t2)
}
