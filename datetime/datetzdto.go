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
	tagDescription string      // Available for tags, classification, labeling or description
	timeComponents TimeDto     // Associated Time Components (years, months, days, hours, minutes,
	                           //    seconds etc.)
	dateTimeValue  time.Time   // DateTime value for this DateTzDto Type
	dateTimeFmt    string      // Date Time Format String. Default is
	                           //    "2006-01-02 15:04:05.000000000 -0700 MST"
	timeZone    TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone
	                           //    Location associated with this date time.
	lock        sync.Mutex     // Mutex used to ensure thread-safe operations.
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
//              A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//        tagDescription  string         // Unused, available for classification,
//                                       //  labeling or description
//        Time            TimeDto        // Associated Time Components
//        DateTime        time.Time      // DateTime value for this DateTzDto Type
//        DateTimeFmt     string         // Date Time Format String.
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//        TimeZone        TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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

	ePrefix := "DateTzDto.AddDate() "

	err := dtz.IsValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+"The current DateTzDto is INVALID! dtz.dateTimeValue='%v'", dtz.dateTimeValue.Format(FmtDateTimeYrMDayFmtStr))
	}

	newDt1 := dtz.dateTimeValue.AddDate(years, months, 0)

	dur := DayNanoSeconds * int64(days)
	newDt2 := newDt1.Add(time.Duration(dur))

	if dateTimeFormatStr == "" {
		dateTimeFormatStr = dtz.dateTimeFmt
	}

	dtz2, err := DateTzDto{}.New(newDt2, dateTimeFormatStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+"Error returned by DateTzDto{}.New(newDt2, dtz.DateTimeFmt). newDt='%v'  Error='%v'", newDt2.Format(FmtDateTimeYrMDayFmtStr), err.Error())
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
//              A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//        tagDescription  string         // Unused, available for classification,
//                                       //  labeling or description
//        Time            TimeDto        // Associated Time Components
//        DateTime        time.Time      // DateTime value for this DateTzDto Type
//        DateTimeFmt     string         // Date Time Format String.
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//        TimeZone        TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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
//  dtz, err := dtz.AddDateTime(
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

	ePrefix := "DateTzDto.AddDateTime() "

	newDate := dtz.dateTimeValue.AddDate(years, months, 0)

	totNanoSecs := int64(days) * DayNanoSeconds
	totNanoSecs += int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	newDateTime := newDate.Add(time.Duration(totNanoSecs))

	dtz2, err := DateTzDto{}.New(newDateTime, dateTimeFormatStr)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+"Error returned from DateTzDto{}.New(newDateTime, dateTimeFormatStr) "+
				"newDateTime='%v' Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return dtz2, nil
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

	ePrefix := "DateTzDto.AddDateTimeToThis() "

	dtz2, err := dtz.AddDateTime(years, months, days, hours, minutes, seconds,
		milliseconds, microseconds, nanoseconds, dtz.dateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error='%v'", err.Error())
	}

	dtz.CopyIn(dtz2)

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

	ePrefix := "DateTzDto.AddDateToThis() "

	err := dtz.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nThe current DateTzDto is INVALID!\ndtz.dateTimeValue='%v'\n",
			dtz.dateTimeValue.Format(FmtDateTimeYrMDayFmtStr))
	}

	newDt1 := dtz.dateTimeValue.AddDate(years, months, 0)
	dur := int64(days) * DayNanoSeconds
	newDt2 := newDt1.Add(time.Duration(dur))

	dtz2, err := DateTzDto{}.New(newDt2, dtz.dateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by DateTzDto{}.New(newDt2, dtz.dateTimeFmt). newDt='%v'  Error='%v'",
			newDt2.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.CopyIn(dtz2)

	return nil

}

// AddDuration - Adds Duration to the DateTime Value of the current
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
//              A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//        tagDescription  string         // Unused, available for classification,
//                                       //  labeling or description
//        Time            TimeDto        // Associated Time Components
//        DateTime        time.Time      // DateTime value for this DateTzDto Type
//        DateTimeFmt     string         // Date Time Format String.
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//        TimeZone        TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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

	ePrefix := "DateTzDto.AddDuration() "

	newDateTime := dtz.dateTimeValue.Add(duration)

	dtz2, err := DateTzDto{}.New(newDateTime, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+
			"\nError returned by DateTzDto{}.New(newDateTime, dateTimeFmtStr).\n" +
			"newDateTime='%v'\nError='%v'\n",
			newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return dtz2, nil

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

	ePrefix := "DateTzDto.AddDurationToThis() "

	newDateTime := dtz.dateTimeValue.Add(duration)

	dtz2, err := DateTzDto{}.New(newDateTime, dtz.dateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by DateTzDto{}.New(newDateTime, dtz.dateTimeFmt).\n" +
			"newDateTime='%v'\nError='%v'\n",
			newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.CopyIn(dtz2)

	return nil
}

// AddMinusTimeDto - Creates and returns a new DateTzDto by subtracting a TimeDto
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
//              input parameter. A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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

	ePrefix := "DateTzDto.AddMinusTimeDto() "
	dtz2 := dtz.CopyOut()

	err := dtz2.AddMinusTimeDtoToThis(minusTimeDto)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"Error returned from dtz2.AddMinusTimeDtoToThis(minusTimeDto). "+
				" Error='%v'", err.Error())
	}

	return dtz2, nil
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

	ePrefix := "DateTzDto.AddMinusTimeDtoToThis() "

	tDto := minusTimeDto.CopyOut()

	err := tDto.NormalizeTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDto.NormalizeTimeElements(). "+
			"Error='%v' ", err.Error())
	}

	_, err = tDto.NormalizeDays()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tDto.NormalizeDays(). "+
			"Error='%v' ", err.Error())

	}

	tDto.ConvertToNegativeValues()

	dt1 := dtz.dateTimeValue.AddDate(tDto.Years,
		tDto.Months,
		0)

	totNanosecs := int64(tDto.DateDays) * DayNanoSeconds
	totNanosecs += int64(tDto.Hours) * HourNanoSeconds
	totNanosecs += int64(tDto.Minutes) * MinuteNanoSeconds
	totNanosecs += int64(tDto.Seconds) * SecondNanoseconds
	totNanosecs += int64(tDto.Milliseconds) * MilliSecondNanoseconds
	totNanosecs += int64(tDto.Microseconds) * MicroSecondNanoseconds
	totNanosecs += int64(tDto.Nanoseconds)

	dt2 := dt1.Add(time.Duration(totNanosecs))

	dtz2, err := DateTzDto{}.NewTz(dt2, dtz.timeZone.LocationName, dtz.dateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned from DateTzDto{}.New(dt2, dtz.dateTimeFmt). "+
			" Error='%v'", err.Error())
	}

	dtz.CopyIn(dtz2)

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
//              'plusTimeDto'. A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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

	ePrefix := "DateTzDto.AddPlusTimeDto() "

	dtz2 := dtz.CopyOut()

	err := dtz2.AddPlusTimeDtoToThis(plusTimeDto)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+"Error returned from dtz2.AddPlusTimeDtoToThis(plusTimeDto). "+
				" Error='%v'", err.Error())
	}

	return dtz2, nil
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

	ePrefix := "DateTzDto.AddPlusTimeDtoToThis() "
	tDto := plusTimeDto.CopyOut()

	err := tDto.NormalizeTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tDto.NormalizeTimeElements(). "+
			"Error='%v' ", err.Error())
	}

	_, err = tDto.NormalizeDays()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tDto.NormalizeDays(). "+
			"Error='%v' ", err.Error())
	}

	tDto.ConvertToAbsoluteValues()

	dt1 := dtz.dateTimeValue.AddDate(tDto.Years,
		tDto.Months,
		0)

	incrementalDur := int64(tDto.DateDays) * DayNanoSeconds
	incrementalDur += int64(tDto.Hours) * HourNanoSeconds
	incrementalDur += int64(tDto.Minutes) * MinuteNanoSeconds
	incrementalDur += int64(tDto.Seconds) * SecondNanoseconds
	incrementalDur += int64(tDto.Milliseconds) * MilliSecondNanoseconds
	incrementalDur += int64(tDto.Microseconds) * MicroSecondNanoseconds
	incrementalDur += int64(tDto.Nanoseconds)

	dt2 := dt1.Add(time.Duration(incrementalDur))

	dtz2, err := DateTzDto{}.New(dt2, dtz.dateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned from DateTzDto{}.New(dt2, dtz.dateTimeFmt). "+
			" Error='%v'", err.Error())
	}

	dtz.CopyIn(dtz2)

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
//              input parameters. A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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

	ePrefix := "DateTzDto.AddTime() "

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	newDateTime := dtz.dateTimeValue.Add(time.Duration(totNanoSecs))

	dtz2, err := DateTzDto{}.New(newDateTime, dtz.dateTimeFmt)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+"Error returned by DateTzDto{}.New(newDateTime, dtz.dateTimeFmt) "+
				"newDateTime='%v'  Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz2.SetDateTimeFmt(dateTimeFormatStr)

	return dtz2, nil
}

// AddTimeToThis - Modifies the current DateTzTdo instance by adding input parameter
// time components (hours, minutes, seconds etc.) to the current value.
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

	ePrefix := "DateTzDto.AddTimeToThis() "

	dtz2, err := dtz.AddTime(hours, minutes, seconds, milliseconds,
		microseconds, nanoseconds, dtz.dateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: '%v'", err.Error())
	}

	dtz.CopyIn(dtz2)

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
//      A DateTzDto struct is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//            Time        TimeDto        // Associated Time Components
//            DateTime    time.Time      // DateTime value for this DateTzDto Type
//            DateTimeFmt string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//            TimeZone    TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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

// CopyOut - returns a DateTzDto instance
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
//              A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description string          // Unused, available for classification,
//                                       //  labeling or description
//            Time        TimeDto        // Associated Time Components
//            DateTime    time.Time      // DateTime value for this DateTzDto Type
//            DateTimeFmt string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//            TimeZone    TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//  ... initialize to some value
//
//  dtz2 := dtz.CopyOut()
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

// GetDateTimeValue - Returns DateTzDto private member variable
// 'dateTimeValue' as a type time.Time.
//
func (dtz *DateTzDto) GetDateTimeValue() time.Time {
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
	return dtz.dateTimeValue.Format(FmtDateTimeEverything)
}

// GetDateTimeNanoSecText - Returns formatted
// date time string with nanoseconds
// 	EXAMPLE: 2006-01-02 15:04:05.000000000
//
func (dtz *DateTzDto) GetDateTimeNanoSecText() string {
	// Time Format down to the nanosecond
	return dtz.dateTimeValue.Format(FmtDateTimeNanoSecondStr)
}

// GetDateTimeFmt - Returns the DateTzDto private member
// variable, DateTzDto.dateTimeFmt.
//
func (dtz *DateTzDto) GetDateTimeFmt() string {
	return dtz.dateTimeFmt
}

// GetDateTimeSecText - Returns formatted
// date time with seconds for display,
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: 2006-01-02 15:04:05
//
func (dtz *DateTzDto) GetDateTimeSecText() string {
	// Time Display Format with seconds
	return dtz.dateTimeValue.Format(FmtDateTimeSecText)
}

// GetDateTimeStr - Returns a date time string
// in the format 20170427211307. Useful in naming
// files.
func (dtz *DateTzDto) GetDateTimeStr() string {

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
func (dtz *DateTzDto) GetDateTimeTzNanoSecDowYMDText() string {
	return dtz.dateTimeValue.Format(FmtDateTimeTzNanoDowYMD)
}

// GetDateTimeTzNanoSecText - Outputs date time in string format using
// the FmtDateTimeDMYNanoTz format which incorporates date time to nano seconds
// and the associated time zone.
//
// ------------------------------------------------------------------------
//
//  EXAMPLE: 01/02/2006 15:04:05.000000000 -0700 MST
func (dtz *DateTzDto) GetDateTimeTzNanoSecText() string {
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
func (dtz *DateTzDto) GetDateTimeTzNanoSecYMDDowText() string {
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
func (dtz *DateTzDto) GetDateTimeYMDAbbrvDowNano() string {
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
	return dtz.dateTimeValue.Format(FmtDateTimeYrMDayFmtStr)
}

// GetDescription - Returns DateTzDto private member
// variable, DateTzDto.tagDescription.
//
// 'tagDescription' is available to users for use as
// a tag, label, classification or text description.
//
func (dtz *DateTzDto) GetDescription() string {
	return dtz.tagDescription
}

// GetMilitaryDateTzDto - Returns an instance of 'MilitaryDateTzDto' which
// is equivalent to the current date, time and time zone represented by the
// current 'DateTzDto' object.
//
func (dtz *DateTzDto) GetMilitaryDateTzDto() (MilitaryDateTzDto, error) {

	ePrefix := "DateTzDto.GetMilitaryDateTzDto() "

	err := dtz.IsValid()

	if err != nil {
		return MilitaryDateTzDto{},
			fmt.Errorf(ePrefix +
				"\nCurrent DateTzDto is INVALID!\n" +
				"%v", err.Error())
	}

	dtzDateTimeStr := dtz.dateTimeValue.Format(FmtDateTimeYMDHMSTz)

	dtzDateTimeArray := strings.Split(dtzDateTimeStr, " ")

	if len(dtzDateTimeArray) != 4 {
		return MilitaryDateTzDto{},
			fmt.Errorf(ePrefix +
				"\nError: Current DateTzDto is INVALID!\n" +
				"Date Time String='%v'\n", dtzDateTimeStr)
	}

	utcOffset := dtzDateTimeArray[2]

	utcPrefix := utcOffset[:3]

	utcOffset = utcPrefix + "00"

	milTzDat := MilitaryTimeZoneData{}

	militaryTz, ok := milTzDat.UtcOffsetToMilitaryTimeZone(utcOffset)

	if !ok {
		return MilitaryDateTzDto{},
			fmt.Errorf(ePrefix +
				"\nError: UTC Offset is INVALID!\n" +
				"utcOffset='%v'\n", utcOffset)
	}

	var militaryTzDto MilitaryDateTzDto

	militaryTzDto, err = MilitaryDateTzDto{}.New(dtz.dateTimeValue, militaryTz)

	if err != nil {
		return MilitaryDateTzDto{},
			fmt.Errorf(ePrefix +
				"\nError returned by MilitaryDateTzDto{}.New(dtz.DateTime," +
				" militaryTz)\n" +
				"dtz.DateTime='%v'\n" +
				"militaryTz='%v'\n" +
				"Error='%v'\n",
				dtz.dateTimeValue.Format(FmtDateTimeYMDHMSTz),militaryTz, err.Error() )
	}

	return militaryTzDto, nil
}

// GetTimeComponents - Returns a deep copy of DateTzDto
// private member variable DateTzDto.timeComponents.
// The private member variable is returned as a type
// 'TimeDto'.
//
func (dtz *DateTzDto) GetTimeComponents() TimeDto {
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
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
//
// error - If successful the returned error Type is set equal to 'nil'. If errors are
//         encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) GetTimeDto() (TimeDto, error) {

	ePrefix := "DateTzDto.GetTimeDto() "

	tDto, err := TimeDto{}.NewFromDateTime(dtz.dateTimeValue)

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix+
			"Error returned by TimeDto{}.NewFromDateTime(dtz.DateTime) "+
			"dtz.DateTime ='%v'  Error='%v'",
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

	return dtz.dateTimeValue.Format(FmtDateTimeYMDAbbrvDowNano)

}

// GetTimeZone - Returns a deep copy of the 'DateTzDto' private
// member variable, 'timeZone', of type TimeZoneDefDto.
//
func (dtz *DateTzDto) GetTimeZone() TimeZoneDefDto {
	return dtz.timeZone.CopyOut()
}

// IsEmpty - Analyzes the current DateTzDto instance to determine
// if the instance is in an 'EMPTY' or uninitialized state.
//
// If the current DateTzDto instance is found to be 'EMPTY', this
// method returns 'true'. Otherwise, if the instance is 'NOT EMPTY',
// this method returns 'false'.
func (dtz *DateTzDto) IsEmpty() bool {

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

	ePrefix := "DateTzDto.IsValidDateTime() "

	if dtz.IsEmpty() {
		return errors.New(ePrefix + "Error: This DateTzDto instance is EMPTY!")
	}

	if dtz.dateTimeValue.IsZero() {
		return errors.New(ePrefix + "Error: DateTzDto.DateTime is ZERO!")
	}

	if dtz.timeZone.IsEmpty() {
		return errors.New(ePrefix + "Error: dtz.TimeZone is EMPTY!")
	}

	if err := dtz.timeComponents.IsValidDateTime(); err != nil {
		return fmt.Errorf(ePrefix+"Error: dtz.timeComponents is INVALID. Error='%v'", err.Error())
	}

	if !dtz.timeZone.IsValidFromDateTime(dtz.dateTimeValue) {
		return errors.New(ePrefix + "Error: dtz.TimeZone is INVALID!")
	}

	dtz2, err := DateTzDto{}.New(dtz.dateTimeValue, dtz.dateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error creating check DateTzDto - Error='%v'", err.Error())
	}

	if !dtz.Equal(dtz2) {
		return errors.New(ePrefix + "Error: Current DateTzDto is NOT EQUAL to Check DateTzDto!")
	}

	return nil
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
//              A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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
//               A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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

	ePrefix := "DateTzDto.NewDateTimeComponents() "

	dtz2 := DateTzDto{}

	err := dtz2.SetFromDateTimeComponents(year, month, day, hour, minute, second,
		millisecond, microsecond, nanosecond, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+"Error returned by dtz2.SetFromDateTimeComponents(...) "+
			"year='%v', month='%v', day='%v', hour='%v', minute='%v', second='%v', millisecond='%v', microsecond='%v' nanosecond='%v', timeZoneLocation='%v' Error='%v'",
			year, month, day, hour, minute, second, millisecond, microsecond, nanosecond, timeZoneLocation, err.Error())
	}

	return dtz2, nil
}

// NewDateTimeElements - creates a new DateTzDto object and populates the data fields based on
// input parameters.
//
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
//               A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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

	ePrefix := "DateTzDto.NewDateTimeElements() "

	dtz2 := DateTzDto{}

	err := dtz2.SetFromDateTimeElements(year, month, day, hour, minute, second,
		nanosecond, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"Error returned from dtz2.SetFromDateTimeElements(...) "+
				"year='%v' month='%v' day='%v' hour='%v' minute='%v' second='%v' "+
				"nanosecond='%v' timeZoneLocatin='%v'  Error='%v'",
				year, month, day, hour, minute, second, nanosecond, timeZoneLocation, err.Error())
	}

	return dtz2, nil
}


// NewFromMilitaryDateTz - Creates and returns a new DateTzDto initialized from an
// instance of type 'MilitaryDateTzDto' passed as an input parameter.
//
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//   militaryDtDto  MilitaryDateTzDto - A valid instance of type 'MilitaryDateTzDto'
//                                      which is converted to and returned as
//                                      a type 'DateTzDto'.
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
//
//               A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String.
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
//
//
//   error     - If successful the returned error Type is set equal to 'nil'.
//               If errors are encountered this error Type will encapsulate
//               an appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtTzDto, err := DateTzDto{}.NewFromMilitaryDateTz(milDtTzDto, dtTimeFmtStr)
//
func (dtz DateTzDto) NewFromMilitaryDateTz(
		militaryDtDto MilitaryDateTzDto,
		dateTimeFmtStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.NewFromMilitaryDateTz() "

	newDateTz := DateTzDto{}

	err := militaryDtDto.IsValid()

	if err != nil {
		return newDateTz,
			fmt.Errorf(ePrefix +
				"\nInput parameter 'militaryDtDto' is INVALID!\n" +
				"Error='%v'\n", err.Error())
	}

	err = newDateTz.SetFromTimeTz(
		militaryDtDto.DateTime,
		militaryDtDto.EquivalentIanaTimeZone.LocationName,
		dateTimeFmtStr)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix +
				"\nError returned by newDateTz.SetFromTimeTz(dateTime, timeZone).\n" +
				"Error='%v'\n", err.Error())
	}

	return newDateTz, nil
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
//               A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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
func (dtz DateTzDto) NewNowLocal(dateTimeFmtStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.NewNowLocal() "

	dt := time.Now().Local()

	dTz := DateTzDto{}

	err := dTz.SetFromTimeTz(dt, "Local", dateTimeFmtStr)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+"Error returned by SetFromTimeTz(). Error='%v'", err.Error())
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
//               A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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

	ePrefix := "DateTzDto.NewNowTz() "

	dt := time.Now().Local()

	dTz := DateTzDto{}

	err := dTz.SetFromTimeTz(dt, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+"Error returned by SetFromTimeTz(). Error='%v'", err.Error())
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
//	https://en.wikipedia.org/wiki/Coordinated_Universal_Time
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
//               A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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
func (dtz DateTzDto) NewNowUTC(dateTimeFmtStr string) (DateTzDto, error) {
	ePrefix := "DateTzDto.NewNowUTC() "

	dt := time.Now().Local()

	dTz := DateTzDto{}

	// TZones.UTC()
	err := dTz.SetFromTimeTz(dt, TZones.Other.UTC() , dateTimeFmtStr)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+"Error returned by SetFromTimeTz(). Error='%v'", err.Error())
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
//               A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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
	timeZoneLocation string,
	dateTimeFormatStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.NewTimeDto() "

	dtz2 := DateTzDto{}

	err := dtz2.SetFromTimeDto(tDto, timeZoneLocation)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+"Error returned by dtz2.SetFromTimeDto(tDto, timeZoneLocation). Error='%v'", err.Error())
	}

	dtz2.SetDateTimeFmt(dateTimeFormatStr)

	return dtz2, nil
}

// NewTz - returns a new DateTzDto instance based on a time.Time input parameter ('dateTime').
// The caller is required to provide a Time Zone Location. Input parameter 'dateTime' will be
// converted to this Time Zone before storing the converted 'dateTime' in the newly created
// DateTzDto instance.
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
//               A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String. 
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
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
	timeZoneLocation,
	dateTimeFmtStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	if dateTime.IsZero() {
		return DateTzDto{}, errors.New(ePrefix + "Error: Input parameter dateTime is Zero value!")
	}

	dtz2 := DateTzDto{}

	err := dtz2.SetFromTimeTz(dateTime, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+
				"Error returned from SetFromTimeTz(dateTime, timeZoneLocation, dateTimeFmtStr). "+
				"dateTime='%v' timeZoneLocation='%v' dateTimeFmtStr='%v'  Error='%v'",
				dateTime.Format(FmtDateTimeYrMDayFmtStr), timeZoneLocation, dateTimeFmtStr, err.Error())
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

	ePrefix := "DateTzDto.SetFromDateTimeComponents() "


	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromDateTimeComponents(
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

	return err
}

// SetFromDateTimeElements - Sets the values of the current DateTzDto
// data fields based on input parameters consisting of date time
// components, a time zone location and a date time format string.
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
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DateTzDto.SetFromDateTimeElements() "

	tDto, err := TimeDto{}.New(year, month, 0, day, hour, minute, second,
		0, 0, nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from TimeDto{}.New(year, month, ...). "+
			" Error='%v'", err.Error())
	}

	fmtStr := dtz.preProcessDateFormatStr(dateTimeFmtStr)

	tzl := dtz.preProcessTimeZoneLocation(timeZoneLocation)

	_, err = time.LoadLocation(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by time.LoadLocation(tzl). INVALID 'timeZoneLocation'! "+
			"tzl='%v' timeZoneLocation='%v' Error='%v' ",
			tzl, timeZoneLocation, err.Error())
	}

	dt, err := tDto.GetDateTime(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tDto.GetDateTime(tzl). "+
			"timeZoneLocation='%v' tzl='%v' Error='%v'",
			timeZoneLocation, tzl, err.Error())
	}

	timeZone, err := TimeZoneDefDto{}.New(dt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by TimeZoneDefDto{}.New(dt). "+
			"tzl='%v' timeZonelocation='%v' dt='%v' Error='%v'",
			tzl, timeZoneLocation, dt.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.Empty()

	dtz.dateTimeValue = dt
	dtz.timeZone = timeZone.CopyOut()
	dtz.timeComponents = tDto.CopyOut()
	dtz.dateTimeFmt = fmtStr

	return nil
}

// SetFromTime - Sets the values of the current DateTzDto fields
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
func (dtz *DateTzDto) SetFromTime(dateTime time.Time, dateTimeFmtStr string) error {

	dtz.lock.Lock()
	defer dtz.lock.Unlock()

	ePrefix := "DateTzDto.SetFromTime() "

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

	if tDto.IsEmpty() {

		return fmt.Errorf(ePrefix + "Error: All input parameter date time elements equal ZERO!")
	}

	t2Dto := tDto.CopyOut()

	err := t2Dto.NormalizeTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by t2Dto.NormalizeTimeElements(). Error='%v' ",
			err.Error())
	}

	t2Dto.ConvertToAbsoluteValues()

	if err = t2Dto.IsValidDateTime(); err != nil {
		return fmt.Errorf(ePrefix+
			"Error: Input Parameter tDto (TimeDto) is INVALID. Error='%v'",
			err.Error())
	}

	tzl := dtz.preProcessTimeZoneLocation(timeZoneLocation)

	_, err = time.LoadLocation(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by time.LoadLocation(tzl). "+
			"timeZoneLocation='%v' tzl='%v'  Error='%v' ", timeZoneLocation, tzl, err.Error())
	}

	dateTime, err := tDto.GetDateTime(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by tDto.GetDateTime(tzl). "+
			"timeZoneLocation='%v' tzl='%v' Error='%v'",
			timeZoneLocation, tzl, err.Error())
	}

	timeZoneDef, err := TimeZoneDefDto{}.New(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by TimeZoneDefDto{}.New(dateTime). dateTime='%v' Error='%v'", dateTime, err.Error())
	}

	fmtStr := dtz.dateTimeFmt

	dtz.Empty()
	dtz.dateTimeValue = dateTime
	dtz.timeZone = timeZoneDef.CopyOut()
	dtz.timeComponents = t2Dto.CopyOut()
	dtz.dateTimeFmt = fmtStr

	return nil
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
func (dtz *DateTzDto) SetFromTimeTz(
	dateTime time.Time,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DateTzDto.SetFromTimeTz() "

	tzl := dtz.preProcessTimeZoneLocation(timeZoneLocation)

	tLoc, err := time.LoadLocation(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"INVALID timeZoneLocation. Error returned by time.LoadLocation(tzl) "+
			"timeZoneLocation='%v' tzl='%v'  Error='%v'",
			timeZoneLocation, tzl, err.Error())
	}

	targetDateTime := dateTime.In(tLoc)

	tZone, err := TimeZoneDefDto{}.New(targetDateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by TimeZoneDefDto{}.New(targetDateTime) "+
			"targetDateTime='%v' Target Time Zone Location='%v' Error='%v'",
			targetDateTime.Format(FmtDateTimeYrMDayFmtStr), tzl, err.Error())
	}

	tDto, err := TimeDto{}.NewFromDateTime(targetDateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from TimeDto{}.NewFromDateTime(targetDateTime). "+
			" targetDateTime='%v'  Error='%v'",
			targetDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.Empty()
	dtz.dateTimeValue = targetDateTime
	dtz.timeZone = tZone.CopyOut()
	dtz.timeComponents = tDto.CopyOut()
	dtz.dateTimeFmt = dateTimeFmtStr

	return nil
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

	ePrefix := "DateTzDto.SetNewTimeZone() "
	tzl := dtz.preProcessTimeZoneLocation(newTimeZoneLocation)

	loc, err := time.LoadLocation(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by time.LoadLocation(tzl). "+
			"tzl='%v' newTimeZoneLocation='%v' Error='%v'",
			tzl, newTimeZoneLocation, err.Error())
	}

	newDateTime := dtz.dateTimeValue.In(loc)
	newFmtStr := dtz.dateTimeFmt

	err = dtz.SetFromTime(newDateTime, newFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by SetFromTime(newDateTime, newFmtStr). "+
			"newDateTime='%v' Error='%v'",
			newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return nil
}

// SetTagDescription - Sets DateTzDto private member variable
// DateTzDto.tagDescription to the value passed in 'tagDesc'.
//
// DateTzDto.tagDescription is available to users for use as
// a tag, label, classification or description.
//
func (dtz *DateTzDto) SetTagDescription(tagDesc string) {
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
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
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
// Return Value
//   time.Duration - A time duration value representing the subtraction of the
//                   input parameter t2 time value from the time value of the
//                   current DateTzDto time value.
//
func (dtz *DateTzDto) SubDateTime(t2 time.Time) time.Duration {
	return dtz.dateTimeValue.Sub(t2)
}

func (dtz *DateTzDto) preProcessDateFormatStr(dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}

func (dtz *DateTzDto) preProcessTimeZoneLocation(timeZoneLocation string) string {

	if len(timeZoneLocation) == 0 {
		return TZones.Other.UTC()
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return "Local"
	}

	return timeZoneLocation
}
