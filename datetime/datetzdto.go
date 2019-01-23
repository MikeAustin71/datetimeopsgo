package datetime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// DateTzDto
//
// This source file is located in source code repository:
// 		'https://github.com/MikeAustin71/datetimeopsgo.git'
//
// This source code file is located at:
//		MikeAustin71\datetimeopsgo\datetime\datetzdto.go
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
// This Type is NOT used to define time duration; that is, the
// difference or time span between two points in time. For time
// duration calculations refer to types, 'TimeDurationDto' and
// 'DurationTriad' located in source files:
//
//		'github.com/MikeAustin71/datetimeopsgo/datetime/timedurationdto.go'
//		'github.com/MikeAustin71/datetimeopsgo/datetime/durationtriad.go'
//
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
// 														(1) the string 'Local' - signals the designation of the local time zone
//																	location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// A requirement for presentation of date time strings is a specific format
// for displaying years, months, days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds. Many 'DateTzDto' methods require calling functions
// to provide a date time format string, ('dateTimeFmtStr'). This format string
// is used to configure date times for display purposes.
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an 'empty string', a
//				default date time format string will be applied. The default
//				date time format string is:
//						FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// DateTzDto Structure
//
// ===========================
type DateTzDto struct {
	Description string         // Unused, available for classification, labeling or description
	Time        TimeDto        // Associated Time Components
	DateTime    time.Time      // DateTime value for this DateTzDto Type
	DateTimeFmt string         // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
	TimeZone    TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
	// 		associated with this date time.
}

// AddDate - Adds input parameters 'years, 'months' and 'days' to date time value of the
// current DateTzDto and returns the updated value in a new DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	years   int - Number of years to add to the current date.
//	months  int - Number of months to add to the current date.
//	days    int - Number of days to add to the current date.
//
//	        Note: Date Component input parameters may be either negative
//	              or positive. Negative values will subtract time from
//	              the current DateTzDto instance.
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful the method returns a new, valid, fully populated
//	            DateTzDto type updated to reflect the added input parameters,
//	            years, months and days.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz := DateTzDto{}
//
//	du, err := dtz.AddDate(
//	                years,
//	                months,
//	                days,
//	                FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' is a constant available in
//	      source file 'datetimeconstants.go'.
//
func (dtz *DateTzDto) AddDate(
	years,
	months,
	days int,
	dateTimeFormatStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.AddDate() "

	err := dtz.IsValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+"The current DateTzDto is INVALID! dtz.DateTime='%v'", dtz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	newDt1 := dtz.DateTime.AddDate(years, months, 0)

	dur := DayNanoSeconds * int64(days)
	newDt2 := newDt1.Add(time.Duration(dur))

	dtz2, err := DateTzDto{}.New(newDt2, dtz.DateTimeFmt)

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
//	years         int - Number of years to add.
//	months        int - Number of months to add.
//	days          int - Number of days to add.
//	hours         int - Number of hours to add.
//	minutes       int - Number of minutes to add.
//	seconds       int - Number of seconds to add.
//	milliseconds  int - Number of milliseconds to add.
//	microseconds  int - Number of microseconds to add.
//	nanoseconds   int - Number of nanoseconds to add.
//
//	Note: Date Time Component input parameters may be either negative
//	      or positive. Negative values will subtract time from
//	      the current DateTzDto instance.
//
//	dateTimeFmtStr string - A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful the method returns a new, valid, fully populated
//	            DateTzDto type updated to reflect the addition of input
//	            parameters to the date time value of the current DateTzDto.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz := DateTzDto{}
//	... initialize to a value.
//
//	dtz, err := dtz.AddDateTime(
//	                 years,
//	                 months,
//	                 days,
//	                 hours,
//	                 minutes,
//	                 seconds,
//	                 milliseconds,
//	                 microseconds,
//	                 nanoseconds,
//	                 FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//	      'datetimeconstants.go'.
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

	newDate := dtz.DateTime.AddDate(years, months, 0)

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
//	years        int - Number of years to add.
//	months       int - Number of months to add.
//	days         int - Number of days to add.
//	hours        int - Number of hours to add.
//	minutes      int - Number of minutes to add.
//	seconds      int - Number of seconds to add.
//	milliseconds int - Number of milliseconds to add.
//	microseconds int - Number of microseconds to add.
//	nanoseconds  int - Number of nanoseconds to add.
//
//	Note: Date Time Component input parameters may be either negative
//	      or positive. Negative values will subtract time from
//	      the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error - If successful the returned error Type is set equal to 'nil'. If errors are
//	        encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//	... initialize to some value
//
//	err := dtz.AddDateTimeToThis(
//	              years,
//	              months,
//	              days,
//	              hours,
//	              minutes,
//	              seconds,
//	              milliseconds,
//	              microseconds,
//	              nanoseconds)
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
		milliseconds, microseconds, nanoseconds, dtz.DateTimeFmt)

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
//	years    int - Number of years to add to the current date.
//	months   int - Number of months to add to the current date.
//	days     int - Number of days to add to the current date.
//
//	         Note: Date Component input parameters may be either negative
//	               or positive. Negative values will subtract time from
//	               the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	error - If successful the returned error Type is set equal to 'nil'. If errors are
//	        encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//	... initialize to some value
//
//	err := dtz.AddDateToThis(
//	              years,
//	              months,
//	              days)
//
func (dtz *DateTzDto) AddDateToThis(
	years,
	months,
	days int) error {

	ePrefix := "DateTzDto.AddDateToThis() "

	err := dtz.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"The current DateTzDto is INVALID! dtz.DateTime='%v'", dtz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	newDt1 := dtz.DateTime.AddDate(years, months, 0)
	dur := int64(days) * DayNanoSeconds
	newDt2 := newDt1.Add(time.Duration(dur))

	dtz2, err := DateTzDto{}.New(newDt2, dtz.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by DateTzDto{}.New(newDt2, dtz.DateTimeFmt). newDt='%v'  Error='%v'", newDt2.Format(FmtDateTimeYrMDayFmtStr), err.Error())
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
//	duration time.Duration - A Time duration value which is added to the DateTime
//	                         value of the current DateTzDto instance to produce and
//	                         return a new, updated DateTzDto instance.
//
//	          Note: The time.Duration input parameter may be either negative
//	                or positive. Negative values will subtract time from
//	                the current DateTzDto instance.
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful the method returns a new, valid, fully populated
//	            DateTzDto type updated to reflect the addition of input parameter
//	            time duration to the date time value of the current DateTzDto
//	            instance.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//  dtz := DateTzDto{}
//	... initialize to some value
//
//	dtz2, err := dtz.AddDuration(
//	              duration,
//	              FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//	      'datetimeconstants.go'.
//
func (dtz *DateTzDto) AddDuration(
	duration time.Duration,
	dateTimeFmtStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.AddDuration() "

	newDateTime := dtz.DateTime.Add(duration)

	dtz2, err := DateTzDto{}.New(newDateTime, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+"Error returned by DateTzDto{}.New(newDateTime, dateTimeFmtStr). newDateTime='%v'  Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
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
//	duration time.Duration - A Time duration value which is added to the DateTime
//	                         value of the current DateTzDto instance to produce and
//	                         return a new, updated DateTzDto instance.
//
//	         Note: The time.Duration input parameter may be either negative
//	               or positive. Negative values will subtract time from
//	               the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Returns
//
//	error - If successful the returned error Type is set equal to 'nil'. If errors are
//	        encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//	... initialize to some value
//
//	err := dtz.AddDurationToThis(duration)
//
func (dtz *DateTzDto) AddDurationToThis(duration time.Duration) error {

	ePrefix := "DateTzDto.AddDurationToThis() "

	newDateTime := dtz.DateTime.Add(duration)

	dtz2, err := DateTzDto{}.New(newDateTime, dtz.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by DateTzDto{}.New(newDateTime, dtz.DateTimeFmt). newDateTime='%v'  Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
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
//	minusTimeDto TimeDto - A TimeDto instance consisting of time components
//	                       (years, months, weeks, days, hours, minutes etc.)
//	                       which will be subtracted from the date time value
//	                       of the current DateTzDto instance.
//
//				type TimeDto struct {
//					Years          int // Number of Years
//					Months         int // Number of Months
//					Weeks          int // Number of Weeks
//					WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//					DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//					Hours          int // Number of Hours.
//					Minutes        int // Number of Minutes
//					Seconds        int // Number of Seconds
//					Milliseconds   int // Number of Milliseconds
//					Microseconds   int // Number of Microseconds
//					Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//					TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//								 // plus remaining Nanoseconds
//				}
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful the method returns a valid, fully populated
//	            DateTzDto type updated to reflect the subtracted 'TimeDto'
//	            input parameter. A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dtz := DateTzDto{}
//	... initialize to some value
//
//	dtz2, err := dtz.AddMinusTimeDto(minusTimeDto)
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
//	minusTimeDto TimeDto - A TimeDto instance consisting of time components
//	                       (years, months, weeks, days, hours, minutes etc.)
//	                       which will be subtracted from the date time value
//	                       of the current DateTzDto instance.
//
//			type TimeDto struct {
//				Years          int // Number of Years
//				Months         int // Number of Months
//				Weeks          int // Number of Weeks
//				WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//				DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//				Hours          int // Number of Hours.
//				Minutes        int // Number of Minutes
//				Seconds        int // Number of Seconds
//				Milliseconds   int // Number of Milliseconds
//				Microseconds   int // Number of Microseconds
//				Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//				TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//							// plus remaining Nanoseconds
//			}
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error - If successful the returned error Type is set equal to 'nil'. If errors are
//	        encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz := DateTzDto{}
//	... initialize to some value
//
//	err := dtz.AddMinusTimeDtoToThis(minusTimeDto)
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

	dt1 := dtz.DateTime.AddDate(tDto.Years,
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

	dtz2, err := DateTzDto{}.NewTz(dt2, dtz.TimeZone.LocationName, dtz.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned from DateTzDto{}.New(dt2, dtz.DateTimeFmt). "+
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
//	plusTimeDto TimeDto - A TimeDto instance consisting of time components
//	                      (years, months, weeks, days, hours, minutes etc.)
//	                      which will be added to the date time value of the
//	                      current DateTzDto instance and returned as an instance
//	                      of type DateTzDto. Note: The value of the current DateTzDto
//	                      will not be altered.
//
//			type TimeDto struct {
//				Years		int // Number of Years
//				Months		int // Number of Months
//				Weeks		int // Number of Weeks
//				WeekDays	int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//				DateDays	int // Total Number of Days. Weeks x 7 plus WeekDays
//				Hours		int // Number of Hours.
//				Minutes		int // Number of Minutes
//				Seconds		int // Number of Seconds
//				Milliseconds	int // Number of Milliseconds
//				Microseconds	int // Number of Microseconds
//				Nanoseconds	int // Remaining Nanoseconds after Milliseconds & Microseconds
//				TotSubSecNanoseconds	int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//							    //    plus remaining Nanoseconds
//			}
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful the method returns a new, valid, fully populated
//	            DateTzDto type updated to reflect the added input parameter
//	            'plusTimeDto'. A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz := DateTzDto{}
//	... initialize to some value
//
//	dtz2, err := dtz.AddPlusTimeDto(plusTimeDto)
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
//	plusTimeDto TimeDto - A TimeDto instance consisting of time components
//	                      (years, months, weeks, days, hours, minutes etc.)
//	                      which will be added to the date time value of the
//	                      current DateTzDto instance. Note: The value of the
//	                      current DateTzDto will be modified.
//
//			type TimeDto struct {
//				Years		int // Number of Years
//				Months		int // Number of Months
//				Weeks		int // Number of Weeks
//				WeekDays	int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//				DateDays	int // Total Number of Days. Weeks x 7 plus WeekDays
//				Hours		int // Number of Hours.
//				Minutes		int // Number of Minutes
//				Seconds		int // Number of Seconds
//				Milliseconds	int // Number of Milliseconds
//				Microseconds	int // Number of Microseconds
//				Nanoseconds	int // Remaining Nanoseconds after Milliseconds & Microseconds
//				TotSubSecNanoseconds	int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//							    //    plus remaining Nanoseconds
//			}
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error - If successful the returned error Type is set equal to 'nil'. If errors are
//	        encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz := DateTzDto{}
//	... initialize to some value
//
//	err := dtz.AddPlusTimeDtoToThis(plusTimeDto)
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

	dt1 := dtz.DateTime.AddDate(tDto.Years,
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

	dtz2, err := DateTzDto{}.New(dt2, dtz.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned from DateTzDto{}.New(dt2, dtz.DateTimeFmt). "+
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
//	hours         int - Number of hours to add.
//	minutes       int - Number of minutes to add.
//	seconds       int - Number of seconds to add.
//	milliseconds  int - Number of milliseconds to add.
//	microseconds  int - Number of microseconds to add.
//	nanoseconds   int - Number of nanoseconds to add.
//
// 		Note: Time Component input parameters may be either negative
//		or positive. Negative values will subtract time from
//		the current DateTzDto instance.
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//				FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful the method returns a valid, fully populated
//	            DateTzDto type updated to reflect the added time value
//	            input parameters. A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz := DateTzDto{}
//	... initialize to some value
//
//	dtz2, err := dtz.AddTime(
//	              hours,
//	              minutes,
//	              seconds,
//	              milliseconds,
//	              microseconds,
//	              nanoseconds,
//	              FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//	      'datetimeconstants.go'.
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

	newDateTime := dtz.DateTime.Add(time.Duration(totNanoSecs))

	dtz2, err := DateTzDto{}.New(newDateTime, dtz.DateTimeFmt)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix+"Error returned by DateTzDto{}.New(newDateTime, dtz.DateTimeFmt) "+
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
//	hours        int - Number of hours to add.
//	minutes      int - Number of minutes to add.
//	seconds      int - Number of seconds to add.
//	milliseconds int - Number of milliseconds to add.
//	microseconds int - Number of microseconds to add.
//	nanoseconds  int - Number of nanoseconds to add.
//
//	Note: Time Component input parameters may be either negative
//	      or positive. Negative values will subtract time from
//	      the current DateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error - If successful the returned error Type is set equal to 'nil'. If errors are
//	        encountered this error Type will encapsulate an error message.
//
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz := DateTzDto{}
//	... initialize to some value
//
//	err := dtz.AddTimeToThis(
//	              hours,
//	              minutes,
//	              seconds,
//	              milliseconds,
//	              microseconds,
//	              nanoseconds,
//	              FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//	      'datetimeconstants.go'.
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
		microseconds, nanoseconds, dtz.DateTimeFmt)

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
//	dtz2 *DateTzDto - A pointer to a DateTzDto instance.
//	                  This data will be copied into the
//	                  data fields of the current DateTzDto
//	                  instance.
//
//	   A DateTzDto struct is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
// ------------------------------------------------------------------------
//
// Return Values
//
// 	None
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz := DateTzDto{}
//	dtz.CopyIn(dtz2)
//
//	Note: dtz and dtz2 are now equivalent.
//
func (dtz *DateTzDto) CopyIn(dtz2 DateTzDto) {
	dtz.Empty()

	dtz.Description = dtz2.Description
	dtz.Time = dtz2.Time.CopyOut()
	dtz.DateTimeFmt = dtz2.DateTimeFmt

	if !dtz2.DateTime.IsZero() {
		dtz.DateTime = dtz2.DateTime
		dtz.TimeZone = dtz2.TimeZone.CopyOut()
	} else {
		dtz.TimeZone = TimeZoneDefDto{}
		dtz.DateTime = time.Time{}
	}

}

// CopyOut - returns a DateTzDto instance
// which represents a deep copy of the current
// DateTzDto object.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	None
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - This method returns a new, valid, fully populated DateTzDto
//	            which is a deep copy of the current DateTzDto instance.
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz := DateTzDto{}
//	... initialize to some value
//
//	dtz2 := dtz.CopyOut()
//
//	Note: dtz and dtz2 are now equivalent.
//
func (dtz *DateTzDto) CopyOut() DateTzDto {
	dtz2 := DateTzDto{}

	dtz2.Description = dtz.Description
	dtz2.Time = dtz.Time.CopyOut()
	dtz2.DateTimeFmt = dtz.DateTimeFmt

	if !dtz.DateTime.IsZero() {
		dtz2.DateTime = dtz.DateTime
		dtz2.TimeZone = dtz.TimeZone.CopyOut()
	} else {
		dtz2.TimeZone = TimeZoneDefDto{}
		dtz2.DateTime = time.Time{}
	}

	return dtz2
}

// Empty - sets all values of the current DateTzDto
// instance to their uninitialized or zero state.
func (dtz *DateTzDto) Empty() {

	dtz.Description = ""
	dtz.Time.Empty()
	dtz.TimeZone = TimeZoneDefDto{}
	dtz.DateTime = time.Time{}
	dtz.DateTimeFmt = ""

	return
}

// Equal - Returns 'true' if input DateTzDto is equal
// in all respects to the current DateTzDto instance.
//
// Otherwise, the method returns 'false'.
//
func (dtz *DateTzDto) Equal(dtz2 DateTzDto) bool {

	if dtz.Description != dtz2.Description ||
		!dtz.Time.Equal(dtz2.Time) ||
		!dtz.DateTime.Equal(dtz2.DateTime) ||
		dtz.DateTimeFmt != dtz2.DateTimeFmt ||
		!dtz.TimeZone.Equal(dtz2.TimeZone) {

		return false
	}

	return true
}

// GetDateTimeEverything - Receives a time value and formats as
// a date time string in the format:
//
// ------------------------------------------------------------------------
//
//	EXAMPLE: Saturday April 29, 2017 19:54:30.123456489 -0500 CDT
//
func (dtz *DateTzDto) GetDateTimeEverything() string {
	return dtz.DateTime.Format(FmtDateTimeEverything)
}

// GetDateTimeNanoSecText - Returns formatted
// date time string with nanoseconds
// 	EXAMPLE: 2006-01-02 15:04:05.000000000
//
func (dtz *DateTzDto) GetDateTimeNanoSecText() string {
	// Time Format down to the nanosecond
	return dtz.DateTime.Format(FmtDateTimeNanoSecondStr)
}

// GetDateTimeSecText - Returns formatted
// date time with seconds for display,
//
// ------------------------------------------------------------------------
//
//	EXAMPLE: 2006-01-02 15:04:05
//
func (dtz *DateTzDto) GetDateTimeSecText() string {
	// Time Display Format with seconds
	return dtz.DateTime.Format(FmtDateTimeSecText)
}

// GetDateTimeStr - Returns a date time string
// in the format 20170427211307. Useful in naming
// files.
func (dtz *DateTzDto) GetDateTimeStr() string {

	// Time Format down to the second
	return dtz.DateTime.Format(FmtDateTimeSecondStr)

}

// GetDateTimeTzNanoSecDowYMDText - Outputs date time in string format using
// the FmtDateTimeTzNanoDowYMD format which incorporates date time to the
// nano second and the associated time zone. In this format, the date is
// expressed as Year-Month-Day (Example: 2017-12-06). The string is
// prefixed with the day of the week:
//
// ------------------------------------------------------------------------
//
//	EXAMPLE: Monday 2006-01-02 15:04:05.000000000 -0700 MST
func (dtz *DateTzDto) GetDateTimeTzNanoSecDowYMDText() string {
	return dtz.DateTime.Format(FmtDateTimeTzNanoDowYMD)
}

// GetDateTimeTzNanoSecText - Outputs date time in string format using
// the FmtDateTimeTzNano format which incorporates date time to nano seconds
// and the associated time zone.
//
// ------------------------------------------------------------------------
//
//	EXAMPLE: 01/02/2006 15:04:05.000000000 -0700 MST
func (dtz *DateTzDto) GetDateTimeTzNanoSecText() string {
	return dtz.DateTime.Format(FmtDateTimeTzNano)
}

// GetDateTimeTzNanoSecYMDDowText - Outputs date time in string format using
// the FmtDateTimeTzNanoYMDDow format which incorporates date time to nano seconds
// and the associated time zone. In this format, the date is expressed as
// Year-Month-Day (Example: 2017-12-06) followed by the day of the week.
//
// ------------------------------------------------------------------------
//
//	EXAMPLE: 2006-01-02 Monday 15:04:05.000000000 -0700 MST
func (dtz *DateTzDto) GetDateTimeTzNanoSecYMDDowText() string {
	return dtz.DateTime.Format(FmtDateTimeTzNanoYMDDow)
}

// GetDateTimeTzNanoSecYMDText - Outputs date time in string format using
// the FmtDateTimeTzNanoYMD format which incorporates date time to nano seconds
// and the associated time zone. In this format, the date is expressed as
// Year-Month-Day (2017-12-06)
//
// ------------------------------------------------------------------------
//
//	EXAMPLE: 2006-01-02 15:04:05.000000000 -0700 MST
//
func (dtz *DateTzDto) GetDateTimeTzNanoSecYMDText() string {
	return dtz.DateTime.Format(FmtDateTimeTzNanoYMD)
}

// GetDateTimeYMDAbbrvDowNano - Outputs date time in string format using
// the FmtDateTimeYMDAbbrvDowNano format which incorporates date time to the
// nano second and the associated time zone. In this format, the date is
// expressed as Year-Month-Day (Example: 2017-12-06). The string includes
// the abbreviated (limited to 3-characters) day of the week:
//
// ------------------------------------------------------------------------
//
//	EXAMPLE: "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
func (dtz *DateTzDto) GetDateTimeYMDAbbrvDowNano() string {
	return dtz.DateTime.Format(FmtDateTimeYMDAbbrvDowNano)
}

// GetDateTimeYrMDayTzFmtStr - Returns a date time string
// formatted as year-mth-day time and time zone.
// FmtDateTimeYrMDayFmtStr - Year Month Day Date Format String
//
// ------------------------------------------------------------------------
//
//	EXAMPLE: "2006-01-02 15:04:05.000000000 -0700 MST"
func (dtz *DateTzDto) GetDateTimeYrMDayTzFmtStr() string {
	return dtz.DateTime.Format(FmtDateTimeYrMDayFmtStr)
}

// GetTimeDto - Converts the current DateTzDto instance
// date time information into an instance of TimeDto
// and returns that TimeDto to the caller.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	None.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TimeDto - A TimeDto structure is defined as follows:
//
//			type TimeDto struct {
//				Years		int // Number of Years
//				Months		int // Number of Months
//				Weeks		int // Number of Weeks
//				WeekDays	int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//				DateDays	int // Total Number of Days. Weeks x 7 plus WeekDays
//				Hours		int // Number of Hours.
//				Minutes		int // Number of Minutes
//				Seconds		int // Number of Seconds
//				Milliseconds	int // Number of Milliseconds
//				Microseconds	int // Number of Microseconds
//				Nanoseconds	int // Remaining Nanoseconds after Milliseconds & Microseconds
//				TotSubSecNanoseconds	int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//							    //    plus remaining Nanoseconds
//			}
//
//	error	- If successful the returned error Type is set equal to 'nil'. If errors are
//		encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) GetTimeDto() (TimeDto, error) {

	ePrefix := "DateTzDto.GetTimeDto() "

	tDto, err := TimeDto{}.NewFromDateTime(dtz.DateTime)

	if err != nil {
		return TimeDto{}, fmt.Errorf(ePrefix+
			"Error returned by TimeDto{}.NewFromDateTime(dtz.DateTime) "+
			"dtz.DateTime ='%v'  Error='%v'",
			dtz.DateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return tDto, nil
}

// GetTimeStampEverything - Generates and returns a time stamp as
// type string. The time stamp is formatted using the format,
// 'FmtDateTimeEverything'.
//
// ------------------------------------------------------------------------
//
//	Example output:
//		"Saturday April 29, 2017 19:54:30.123456489 -0500 CDT"
func (dtz *DateTzDto) GetTimeStampEverything() string {
	return dtz.DateTime.Format(FmtDateTimeEverything)
}

// GetTimeStampYMDAbbrvDowNano - Generates and returns a time stamp as
// type string. The time stamp is formatted using the format
// 'FmtDateTimeYMDAbbrvDowNano'.
//
// ------------------------------------------------------------------------
//
//	Example Output:
//	"2006-01-02 Mon 15:04:05.000000000 -0700 MST"
func (dtz *DateTzDto) GetTimeStampYMDAbbrvDowNano() string {

	return dtz.DateTime.Format(FmtDateTimeYMDAbbrvDowNano)

}

// IsEmpty - Analyzes the current DateTzDto instance to determine
// if the instance is in an 'EMPTY' or uninitialized state.
//
// If the current DateTzDto instance is found to be 'EMPTY', this
// method returns 'true'. Otherwise, if the instance is 'NOT EMPTY',
// this method returns 'false'.
func (dtz *DateTzDto) IsEmpty() bool {

	if dtz.Description == "" &&
		dtz.Time.IsEmpty() &&
		dtz.DateTime.IsZero() &&
		dtz.DateTimeFmt == "" &&
		dtz.TimeZone.IsEmpty() {

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

	if dtz.DateTime.IsZero() {
		return errors.New(ePrefix + "Error: DateTzDto.DateTime is ZERO!")
	}

	if dtz.TimeZone.IsEmpty() {
		return errors.New(ePrefix + "Error: dtz.TimeZone is EMPTY!")
	}

	if err := dtz.Time.IsValidDateTime(); err != nil {
		return fmt.Errorf(ePrefix+"Error: dtz.Time is INVALID. Error='%v'", err.Error())
	}

	if !dtz.TimeZone.IsValidFromDateTime(dtz.DateTime) {
		return errors.New(ePrefix + "Error: dtz.TimeZone is INVALID!")
	}

	dtz2, err := DateTzDto{}.New(dtz.DateTime, dtz.DateTimeFmt)

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
//	dateTime    time.Time - A date time value
//
//	dateTimeFmtStr string - A date time format string which will be used
//	                       to format and display 'dateTime'. Example:
//	                       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                       Date time format constants are found in the source file
//	                       'datetimeconstants.go'. These constants represent the more
//	                       commonly used date time string formats. All Date Time format
//	                       constants begin with the prefix 'FmtDateTime'.
//
//	                       If 'dateTimeFmtStr' is submitted as an 'empty string', a
//	                       default date time format string will be applied. The default
//	                       date time format string is:
//	                         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful this method returns a new DateTzDto instance.
//	            The data fields of this new instance are initialized to zero
//	            values.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtzDto, err := DateTzDto{}.New(dateTime, FmtDateTimeYrMDayFmtStr)
//
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//	      'datetimeconstants.go'.
//
func (dtz DateTzDto) New(dateTime time.Time, dateTimeFmtStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	if dateTime.IsZero() {
		return DateTzDto{}, errors.New(ePrefix + "Error: Input parameter dateTime is Zero value!")
	}

	dtz2 := DateTzDto{}

	err := dtz2.SetFromTime(dateTime, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+"Error returned from dtz2.SetFromTime(dateTime). dateTime='%v'  Error='%v'", dateTime, err.Error())
	}

	return dtz2, nil
}

// NewDateTime - creates a new DateTzDto object and populates the data fields based on
// input parameters.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//	year			int - year number
//	month			int - month number 	1 - 12
//	day			int - day number   	1 - 31
//	hour			int - hour number  	0 - 24
//	minute			int - minute number	0 - 59
//	second			int - second number	0 - 59
//	millisecond		int - millisecond number 0 - 999
//	microsecond		int - microsecond number 0 - 999
//	nanosecond		int - nanosecond number  0 - 999
//
//	timeZoneLocation	string - time zone location must be designated as one of two values.
//				(1) the string 'Local' - signals the designation of the local time zone
//				location for the host computer.
//
//				(2) IANA Time Zone Location -
//				See https://golang.org/pkg/time/#LoadLocation
//				and https://www.iana.org/time-zones to ensure that
//				the IANA Time Zone Database is properly configured
//				on your system. Note: IANA Time Zone Data base is
//				equivalent to 'tz database'.
//					Examples:
//						"America/New_York"
//						"America/Chicago"
//						"America/Denver"
//						"America/Los_Angeles"
//						"Pacific/Honolulu"
//
//
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful this method returns a new DateTzDto instance.
//	            The data fields of this new instance are initialized to zero
//	            values.
//
//	            A DateTzDto structure is defined as follows:
//
//		type DateTzDto struct {
//			Description      string  // Unused, available for classification, labeling or description
//			Year                int  // Year Number
//			Month               int  // Month Number
//			Day                 int  // Day Number
//			Hour                int  // Hour Number
//			Minute              int  // Minute Number
//			Second              int  // Second Number
//			Millisecond         int  // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//			Microsecond         int  // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//			Nanosecond          int  // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//			                         //   Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//			TotalNanoSecs     int64  // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//			DateTime      time.Time  // DateTime value for this DateTzDto Type
//			DateTimeFmt      string  // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//			TimeZone TimeZoneDefDto  // Contains a detailed description of the Time Zone and Time Zone Location
//				                 //    associated with this date time.
//		}
//
//	error - If successful the returned error Type is set equal to 'nil'. If errors are
//	        encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//		dtzDto, err := DateTzDto{}.New(
//	                     year,
//	                     month,
//	                     day,
//	                     hour,
//	                     min,
//	                     sec,
//	                     nanosecond,
//	                     TzIanaUsCentral,
//	                     FmtDateTimeYrMDayFmtStr)
//
//	Note: TzIanaUsCentral = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'datetimeconstants.go'.
//
func (dtz DateTzDto) NewDateTime(
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

	ePrefix := "DateTzDto.NewDateTime() "

	dtz2 := DateTzDto{}

	err := dtz2.SetFromDateTime(year, month, day, hour, minute, second,
		millisecond, microsecond, nanosecond, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix+"Error returned by dtz2.SetFromDateTime(...) "+
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
//	year        int - year number
//	month       int - month number 	1 - 12
//	day         int - day number   	1 - 31
//	hour        int - hour number  	0 - 24
//	minute      int - minute number	0 - 59
//	second      int - second number	0 - 59
//	nanosecond  int - nanosecond number  0 - 999,999,999
//
//
//	timeZoneLocation	string - time zone location must be designated as one of two values.
//				(1) the string 'Local' - signals the designation of the local time zone
//				location for the host computer.
//
//				(2) IANA Time Zone Location -
//				See https://golang.org/pkg/time/#LoadLocation
//				and https://www.iana.org/time-zones to ensure that
//				the IANA Time Zone Database is properly configured
//				on your system. Note: IANA Time Zone Data base is
//				equivalent to 'tz database'.
//					Examples:
//						"America/New_York"
//						"America/Chicago"
//						"America/Denver"
//						"America/Los_Angeles"
//						"Pacific/Honolulu"
//
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful, this method returns a new DateTzDto instance.
//	            The data fields of this new instance are initialized to zero
//	            values.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//		dtzDto, err := DateTzDto{}.NewDateTimeElements(
//			year,
//			month,
//			day,
//			hour,
//			minute,
//			second,
//			nanosecond ,
//			TzIanaUsCentral,
//			FmtDateTimeYrMDayFmtStr)
//
//	Note: TzIanaUsCentral = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'datetimeconstants.go'.
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

// NewNowLocal - returns a new DateTzDto instance based on a date time value
// which is automatically assigned by time.Now(). Effectively, this means that
// the time selected is equal to the current value of the host computer clock.
//
// The Time Zone Location is automatically set to 'Local'.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful, this method returns a new DateTzDto instance.
//	            The data fields of this new instance are initialized to zero
//	            values.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtzDto, err := DateTzDto{}.NewNowLocal(FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      FmtDateTimeYrMDayFmtStr' is a constants available in source file
//	      'datetimeconstants.go'.
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
//	timeZoneLocation	string - time zone location must be designated as one of two values.
//				(1) the string 'Local' - signals the designation of the local time zone
//				location for the host computer.
//
//				(2) IANA Time Zone Location -
//				See https://golang.org/pkg/time/#LoadLocation
//				and https://www.iana.org/time-zones to ensure that
//				the IANA Time Zone Database is properly configured
//				on your system. Note: IANA Time Zone Data base is
//				equivalent to 'tz database'.
//					Examples:
//						"America/New_York"
//						"America/Chicago"
//						"America/Denver"
//						"America/Los_Angeles"
//						"Pacific/Honolulu"
//
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful, this method returns a new DateTzDto instance.
//	            The data fields of this new instance are initialized to zero
//	            values.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//		dtzDto, err := DateTzDto{}.NewNowTz(
//			TzIanaUsCentral,
//			FmtDateTimeYrMDayFmtStr)
//
//	Note: TzIanaUsCentral = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'datetimeconstants.go'.
//
func (dtz DateTzDto) NewNowTz(timeZoneLocation, dateTimeFmtStr string) (DateTzDto, error) {
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
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful, this method returns a new DateTzDto instance.
//	            The data fields of this new instance are initialized to zero
//	            values.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtzDto, err := DateTzDto{}.NewNowUTC(
//	                   FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' is a constant available in source file
//	      'datetimeconstants.go'.
//
func (dtz DateTzDto) NewNowUTC(dateTimeFmtStr string) (DateTzDto, error) {
	ePrefix := "DateTzDto.NewNowUTC() "

	dt := time.Now().Local()

	dTz := DateTzDto{}

	err := dTz.SetFromTimeTz(dt, TzIanaUTC, dateTimeFmtStr)

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
//	tDto	TimeDto - Time values used to construct the DateTzDto instance
//
//			type TimeDto struct {
//				Years          int // Number of Years
//				Months         int // Number of Months
//				Weeks          int // Number of Weeks
//				WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//				DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//				Hours          int // Number of Hours.
//				Minutes        int // Number of Minutes
//				Seconds        int // Number of Seconds
//				Milliseconds   int // Number of Milliseconds
//				Microseconds   int // Number of Microseconds
//				Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//				TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//							// plus remaining Nanoseconds
//			}
//
//
//	timeZoneLocation	string - time zone location must be designated as one of two values.
//				(1) the string 'Local' - signals the designation of the local time zone
//				location for the host computer.
//
//				(2) IANA Time Zone Location -
//				See https://golang.org/pkg/time/#LoadLocation
//				and https://www.iana.org/time-zones to ensure that
//				the IANA Time Zone Database is properly configured
//				on your system. Note: IANA Time Zone Data base is
//				equivalent to 'tz database'.
//					Examples:
//						"America/New_York"
//						"America/Chicago"
//						"America/Denver"
//						"America/Los_Angeles"
//						"Pacific/Honolulu"
//
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DateTzDto - If successful, this method returns a new DateTzDto instance.
//	            The data fields of this new instance are initialized to zero
//	            values.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtz, err := DateTzDto{}.NewTimeDto(
//				tDto,
//				TzIanaUsCentral,
//				FmtDateTimeYrMDayFmtStr)
//
//	Note: TzIanaUsCentral = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'datetimeconstants.go'.
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
//	dateTime      time.Time - A date time value
//
//	timeZoneLocation string - time zone location must be designated as one of two values.
//				(1) the string 'Local' - signals the designation of the local time zone
//				location for the host computer.
//
//				(2) IANA Time Zone Location -
//				See https://golang.org/pkg/time/#LoadLocation
//				and https://www.iana.org/time-zones to ensure that
//				the IANA Time Zone Database is properly configured
//				on your system. Note: IANA Time Zone Data base is
//				equivalent to 'tz database'.
//					Examples:
//						"America/New_York"
//						"America/Chicago"
//						"America/Denver"
//						"America/Los_Angeles"
//						"Pacific/Honolulu"
//
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				Date time format constants are found in the source file
//				'datetimeconstants.go'. These constants represent the more
//				commonly used date time string formats. All Date Time format
//				constants begin with the prefix 'FmtDateTime'.
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
// Return Values
//
//	DateTzDto - If successful, this method returns a new DateTzDto instance.
//	            The data fields of this new instance are initialized to zero
//	            values.
//
//	            A DateTzDto structure is defined as follows:
//
//	      type DateTzDto struct {
//	        Description      string // Unused, available for classification, labeling or description
//	        Year                int // Year Number
//	        Month               int // Month Number
//	        Day                 int // Day Number
//	        Hour                int // Hour Number
//	        Minute              int // Minute Number
//	        Second              int // Second Number
//	        Millisecond         int // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//	        Microsecond         int // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//	        Nanosecond          int // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//	                                //    Nanosecond = TotalNanoSecs - millisecond nanoseconds - microsecond nanoseconds
//	        TotalNanoSecs     int64 // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//	        DateTime      time.Time // DateTime value for this DateTzDto Type
//	        DateTimeFmt      string // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//	        TimeZone TimeZoneDefDto // Contains a detailed description of the Time Zone and Time Zone Location
//	                                //    associated with this date time.
//	      }
//
//
//	error     - If successful the returned error Type is set equal to 'nil'. If errors are
//	            encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dtzDto, err := DateTzDto{}.NewTz(
//			dateTime,
//			TzIanaUsCentral,
//			FmtDateTimeYrMDayFmtStr)
//
//	Note: TzIanaUsCentral = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'datetimeconstants.go'.
//
func (dtz DateTzDto) NewTz(dateTime time.Time, timeZoneLocation,
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
// This string is used to format the DateTzDto DateTime field
// when DateTzDto.String() is called.
func (dtz *DateTzDto) SetDateTimeFmt(dateTimeFmtStr string) {

	dtz.DateTimeFmt = dateTimeFmtStr

}

// SetFromDateTime - Sets the values of the Date Time fields
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
//	year		int - year number
//	month		int - month number 	1 - 12
//	day		int - day number   	1 - 31
//	hour		int - hour number  	0 - 24
//	min		int - minute number	0 - 59
//	sec		int - second number	0 - 59
//	millisecond	int - millisecond number  0 - 999
//	microsecond	int - microsecond number  0 - 999
//	nanosecond	int - nanosecond number   0 - 999
//
//	timeZoneLocation string - time zone location must be designated as one of two values.
//				(1) the string 'Local' - signals the designation of the local time zone
//				location for the host computer.
//
//				(2) IANA Time Zone Location -
//				See https://golang.org/pkg/time/#LoadLocation
//				and https://www.iana.org/time-zones to ensure that
//				the IANA Time Zone Database is properly configured
//				on your system. Note: IANA Time Zone Data base is
//				equivalent to 'tz database'.
//					Examples:
//						"America/New_York"
//						"America/Chicago"
//						"America/Denver"
//						"America/Los_Angeles"
//						"Pacific/Honolulu"
//
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error	- If successful the returned error Type is set equal to 'nil'. If errors are
//		encountered this error Type will encapsulate an error message.
func (dtz *DateTzDto) SetFromDateTime(
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

	ePrefix := "DateTzDto.SetFromDateTime() "

	tDto, err := TimeDto{}.New(year, month, 0, day, hour, minute,
		second, millisecond, microsecond, nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by TimeDto{}.New(year, month,...).  "+
			"Error='%v'", err.Error())
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
			"timeZoneLocation='%v' tzl='%v'  Error='%v'",
			timeZoneLocation, tzl, err.Error())
	}

	timeZone, err := TimeZoneDefDto{}.New(dt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by TimeZoneDefDto{}.New(dt). "+
			"dt='%v'  Error=%v",
			dt.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.Empty()

	dtz.DateTime = dt
	dtz.TimeZone = timeZone.CopyOut()
	dtz.Time = tDto.CopyOut()
	dtz.DateTimeFmt = fmtStr

	return nil
}

// SetFromDateTimeElements - Sets the values of the current DateTzDto
// data fields based on input parameters consisting of date time
// components, a time zone location and a date time format string.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	year		int - year number
//	month		int - month number 	1 - 12
//	day		int - day number   	1 - 31
//	hour		int - hour number  	0 - 24
//	min		int - minute number	0 - 59
//	sec		int - second number	0 - 59
//	nanosecond	int - nanosecond number   0 - 999,999,999
//
//	timeZoneLocation string - time zone location must be designated as one of two values.
//				(1) the string 'Local' - signals the designation of the local time zone
//				location for the host computer.
//
//				(2) IANA Time Zone Location -
//				See https://golang.org/pkg/time/#LoadLocation
//				and https://www.iana.org/time-zones to ensure that
//				the IANA Time Zone Database is properly configured
//				on your system. Note: IANA Time Zone Data base is
//				equivalent to 'tz database'.
//					Examples:
//						"America/New_York"
//						"America/Chicago"
//						"America/Denver"
//						"America/Los_Angeles"
//						"Pacific/Honolulu"
//
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error	- If successful the returned error Type is set equal to 'nil'. If errors are
//		encountered this error Type will encapsulate an error message.
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

	dtz.DateTime = dt
	dtz.TimeZone = timeZone.CopyOut()
	dtz.Time = tDto.CopyOut()
	dtz.DateTimeFmt = fmtStr

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
//	dateTime    time.Time	- A date time value
//
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error	- If successful the returned error Type is set equal to 'nil'. If errors are
//		encountered this error Type will encapsulate an error message.
func (dtz *DateTzDto) SetFromTime(dateTime time.Time, dateTimeFmtStr string) error {

	ePrefix := "DateTzDto.SetFromTime() "

	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter dateTime is Zero value!")
	}

	tDto, err := TimeDto{}.NewFromDateTime(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from TimeDto{}.NewFromDateTime(dateTime). "+
			" dateTime='%v' Error='%v'",
			dateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	fmtStr := dtz.preProcessDateFormatStr(dateTimeFmtStr)

	timeZone, err := TimeZoneDefDto{}.New(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from TimeZoneDefDto{}.New(dateTime). "+
			"dateTime='%v'  Error='%v'",
			dateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.Empty()

	dtz.DateTime = dateTime
	dtz.Time = tDto.CopyOut()
	dtz.TimeZone = timeZone.CopyOut()
	dtz.DateTimeFmt = fmtStr

	return nil
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
//	tDto	      TimeDto	- A populated TimeDto instance
//
//
//				type TimeDto struct {
//					Years          int // Number of Years
//					Months         int // Number of Months
//					Weeks          int // Number of Weeks
//					WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//					DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//					Hours          int // Number of Hours.
//					Minutes        int // Number of Minutes
//					Seconds        int // Number of Seconds
//					Milliseconds   int // Number of Milliseconds
//					Microseconds   int // Number of Microseconds
//					Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//					TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//								 // plus remaining Nanoseconds
//				}
//
//
//	timeZoneLocation string - time zone location must be designated as one of two values.
//				(1) the string 'Local' - signals the designation of the local time zone
//				location for the host computer.
//
//				(2) IANA Time Zone Location -
//				See https://golang.org/pkg/time/#LoadLocation
//				and https://www.iana.org/time-zones to ensure that
//				the IANA Time Zone Database is properly configured
//				on your system. Note: IANA Time Zone Data base is
//				equivalent to 'tz database'.
//					Examples:
//						"America/New_York"
//						"America/Chicago"
//						"America/Denver"
//						"America/Los_Angeles"
//						"Pacific/Honolulu"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error	- If successful the returned error Type is set equal to 'nil'. If errors are
//		encountered this error Type will encapsulate an error message.
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

	fmtStr := dtz.DateTimeFmt

	dtz.Empty()
	dtz.DateTime = dateTime
	dtz.TimeZone = timeZoneDef.CopyOut()
	dtz.Time = t2Dto.CopyOut()
	dtz.DateTimeFmt = fmtStr

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
//	dateTime     time.Time	- A date time value
//
//
//	timeZoneLocation string - time zone location must be designated as one of two values.
//				(1) the string 'Local' - signals the designation of the local time zone
//				location for the host computer.
//
//				(2) IANA Time Zone Location -
//				See https://golang.org/pkg/time/#LoadLocation
//				and https://www.iana.org/time-zones to ensure that
//				the IANA Time Zone Database is properly configured
//				on your system. Note: IANA Time Zone Data base is
//				equivalent to 'tz database'.
//					Examples:
//						"America/New_York"
//						"America/Chicago"
//						"America/Denver"
//						"America/Los_Angeles"
//						"Pacific/Honolulu"
//
//
//	dateTimeFmtStr string	- A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error	- If successful the returned error Type is set equal to 'nil'. If errors are
//		encountered this error Type will encapsulate an error message.
func (dtz *DateTzDto) SetFromTimeTz(dateTime time.Time, timeZoneLocation,
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
	dtz.DateTime = targetDateTime
	dtz.TimeZone = tZone.CopyOut()
	dtz.Time = tDto.CopyOut()
	dtz.DateTimeFmt = dateTimeFmtStr

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
//	newTimeZoneLocation string - Designates the standard Time Zone location used to
//				     compute date time. The existing DateTzDto Date Time will be converted
//				     to an equivalent time this 'newTimeZoneLocation'. This 'newTimeZoneLocation'
//				     must be designated as one of two values.
//
//				     (1) the string 'Local' - signals the designation of the local time zone
//				     location for the host computer.
//
//				     (2) IANA Time Zone Location -
//				     See https://golang.org/pkg/time/#LoadLocation
//				     and https://www.iana.org/time-zones to ensure that
//				     the IANA Time Zone Database is properly configured
//				     on your system. Note: IANA Time Zone Data base is
//				     equivalent to 'tz database'.
//					     Examples:
//						     "America/New_York"
//						     "America/Chicago"
//						     "America/Denver"
//						     "America/Los_Angeles"
//						     "Pacific/Honolulu"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error	- If successful the returned error Type is set equal to 'nil'. If errors are
//		encountered this error Type will encapsulate an error message.
func (dtz *DateTzDto) SetNewTimeZone(newTimeZoneLocation string) error {

	ePrefix := "DateTzDto.SetNewTimeZone() "
	tzl := dtz.preProcessTimeZoneLocation(newTimeZoneLocation)

	loc, err := time.LoadLocation(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by time.LoadLocation(tzl). "+
			"tzl='%v' newTimeZoneLocation='%v' Error='%v'",
			tzl, newTimeZoneLocation, err.Error())
	}

	newDateTime := dtz.DateTime.In(loc)
	newFmtStr := dtz.DateTimeFmt

	err = dtz.SetFromTime(newDateTime, newFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by SetFromTime(newDateTime, newFmtStr). "+
			"newDateTime='%v' Error='%v'",
			newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return nil
}

// String - This method returns the DateTzDto
// DateTime field value formatted as a string.
// If the DateTzDto data field, 'DateTimeFmt'
// is an empty string, a default format string
// will be used. The default format is:
//
//	FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (dtz *DateTzDto) String() string {

	fmtStr := dtz.DateTimeFmt

	if len(fmtStr) == 0 {
		fmtStr = FmtDateTimeYrMDayFmtStr
	}

	return dtz.DateTime.Format(fmtStr)
}

// Sub - Subtracts the DateTime value of the incoming DateTzDto
// from the DateTime value of the current DateTzDto. The result
// is returned as a Type 'time.Duration'.
func (dtz *DateTzDto) Sub(dtz2 DateTzDto) time.Duration {

	return dtz.DateTime.Sub(dtz2.DateTime)

}

// SubDateTime - Subtracts a date time value (Type: 'time.Time')
// from the date time value of the current DateTzDto. The result
// is returned as a Type 'time.Duration'.
func (dtz *DateTzDto) SubDateTime(t2 time.Time) time.Duration {
	return dtz.DateTime.Sub(t2)
}

func (dtz *DateTzDto) preProcessDateFormatStr(dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}

func (dtz *DateTzDto) preProcessTimeZoneLocation(timeZoneLocation string) string {

	if len(timeZoneLocation) == 0 {
		return TzIanaUTC
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return "Local"
	}

	return timeZoneLocation
}
