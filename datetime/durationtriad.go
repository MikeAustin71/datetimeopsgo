package datetime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

/*

The source file for this type is located in source code repository:
	'https://github.com/MikeAustin71/datetimeopsgo.git'

The location of this source file is:
	'MikeAustin71\datetimeopsgo\datetime\durationtriad.go'


Dependencies

	TimeDurationDto - Calculates and stores date time duration information
		          for a single Time Zone. This Type is defined in same
		          source code repository as 'DurationTriad' and is
		          located in source file:

				'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'



Overview and Usage

The 'DurationTriad' Type is used to calculate date time duration across multiple
time zones. It consists principally of three 'TimeDurationDto' types which perform
and store time duration calculations for three different Time Zone Locations. The
first time zone is designated by the calling function, 'Base Time'. From this initial
input, the equivalent 'Local Time' and 'UTC Time' are then calculated. These three
time zone types are therefore defined as follows:

		1. Base Time Zone - Or Input Parameter Time Zone.
		2. Local Time Zone - The local time zone on the host computer.
		3. Coordinated Universal Time - UTC

Calculations for time duration require the identification of a starting date
time, an ending datetime. Thereafter the time duration computed by subtracting
starting date time from ending date time.

In time duration calculations, time zone location is important. If starting and
ending date time use different time zones, it could create errors in the time
duration result. Also, depending on daylight savings time, the same time duration
could produce differing ending date times depending on which time zone is used.

In order to ensure accuracy, the DurationTriad type first calculates date time duration
for the user specified time zone and then proceeds to calculate that same duration
using the 'Local' Time Zone Location and the 'UTC' Time Zone Location.

BaseTime -

DurationTriad.BaseTime is an instance of 'TimeDurationDto' which performs and
stores starting date time duration calculations for a time zone location specified
by the calling function.

In addition to 'BaseTime', a DurationTriad instance automatically calculates
date time duration for two additional time zone locations, 'Local' time and
UTC time.

LocalTime -

DurationTriad.LocalTime is an instance of 'TimeDurationDto' which performs and
stores date time duration calculations for the 'Local' Time Zone.  This represents
the Time Zone configured for the host computer running this code.

UTCTime -

DurationTriad.LocalTime is an instance of 'TimeDurationDto' which performs and
stores date time duration calculations for the 'UTC' Time Zone.  'UTC' stands
for Coordinated Universal Time, a standardized and uniform methodology for computing
time across the globe. 'UTC' is sometimes referred to as 'Zulu', 'GMT' or Greenwich Mean
Time. Reference: https://en.wikipedia.org/wiki/Coordinated_Universal_Time


Time Duration Calculation Types -

Some of the methods provided by the DurationTriad Type allow the user
to specify the format for time duration information. Examples include
the default, 'Year, month, day, time ... ', or alternative formats like,
'Cumulative Weeks', 'Cumulative Days', 'Cumulative Hours' etc.

For details on type TDurCalcType see the source file:

	'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'


References -

	'Local' Time Zone Location - Package Time, The Go Programming Language
		https://golang.org/pkg/time/

	Coordinated Universal Time
		Reference: https://en.wikipedia.org/wiki/Coordinated_Universal_Time

	Date Time Formats
	  A number of methods require the use of a Date Time Format string. These
	  strings provide specific formats for displaying dates, time zones and
	  time components such is hours, minutes, seconds, milliseconds, microseconds
	  and nanoseconds.

	  dateTimeFmtStr string - A date time format string which will be used
	                          to format and display 'dateTime'. Example:
	                          "2006-01-02 15:04:05.000000000 -0700 MST"

	                          Date time format constants are found in the source
	                          file 'constantsdatetime.go'. These constants represent
	                          the more commonly used date time string formats. All
	                          Date Time format constants begin with the prefix
	                          'FmtDateTime'.

	                          If 'dateTimeFmtStr' is submitted as an
	                          'empty string', a default date time format
	                          string will be applied. The default date time
	                          format string is:
	                            FmtDateTimeYrMDayFmtStr =
	                              "2006-01-02 15:04:05.000000000 -0700 MST"


	Time Zone Location
		Reference Package Time, The Go Programming Language
		https://golang.org/pkg/time/

		Time Zone Location Designates the standard Time Zone location by which
		time duration will be compared. This ensures that, "oranges are compared to oranges
		and apples are compared to apples", with respect to start time and end time duration
		calculations.

		Time zone location must be designated as one of two values:

		(1) the string 'Local' - signals the designation of the local time zone
			location for the host computer. 'Local' is a creation of the
			Go Programming Language. Reference Package Time, The Go Programming
			Language:	https://golang.org/pkg/time/

		(2) IANA Time Zone Location -
			See https://golang.org/pkg/time/#LoadLocation
			and https://www.iana.org/time-zones to ensure that
			the IANA Time Zone Database is properly configured
			on your system. Note: IANA Time Zone Data base is
			equivalent to 'tz database'.
			Examples:
					"America/New_York"
					"America/Chicago"
					"America/Denver"
					"America/Los_Angeles"
					"Pacific/Honolulu"
					"Etc/UTC" = ZULU, GMT or UTC - Default

		The source file 'constantsdatetime.go' contains a number of
		constant declarations covering the more frequently used time
		zones. Example: 'TZones.US.Central()' = "America/Chicago". All
		time zone constants begin with the prefix 'TzIana'.


DurationTriad Structure

=========================================================================
*/
type DurationTriad struct {
	BaseTime  TimeDurationDto
	LocalTime TimeDurationDto
	UTCTime   TimeDurationDto
}

// CopyIn - Receives and incoming DurationTriad instance. This method then
// copies the incoming data values into the current DurationTriad data
// structure. This method performs a deep copy on all data elements.
//
// NOTE: This method will alter the data fields of the current DurationTriad
// instance.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//	duIn DurationTriad - An instance of type 'DurationTriad'. The data fields
//	                     of input parameter 'duIn' will be copied to the
//	                     the data fields of the current DurationTriad instance.
//	                     The type of copy operation performed is a 'deep' copy.
//
//	                     When this operation completes, 'duIn' and the current
//	                     'DurationTriad' instance will be equivalent.
//
//	                     A DurationTriad Structure is defined as follows:
//
//	                     type DurationTriad struct {
//	                       BaseTime  TimeDurationDto
//	                       LocalTime TimeDurationDto
//	                       UTCTime   TimeDurationDto
//	                      }
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	None.
//
func (durT *DurationTriad) CopyIn(duIn DurationTriad) {
	durT.Empty()
	durT.BaseTime = duIn.BaseTime.CopyOut()
	durT.LocalTime = duIn.LocalTime.CopyOut()
	durT.UTCTime = duIn.UTCTime.CopyOut()

	return
}

// CopyOut - Creates and returns a new DurationTriad instance. The deep
// copy operation copies all data elements from the current
// DurationTriad instance to the new DurationTriad instance which is
// returned to the calling function.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	None.
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	DurationTriad - Upon completion, this method returns a new instance of
//	                Type DurationTriad which is, in all respects, an exact
//	                copy of the current DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
func (durT *DurationTriad) CopyOut() DurationTriad {
	duOut := DurationTriad{}
	duOut.BaseTime = durT.BaseTime.CopyOut()
	duOut.LocalTime = durT.LocalTime.CopyOut()
	duOut.UTCTime = durT.UTCTime.CopyOut()

	return duOut
}

// Empty - This method initializes
// all of the fields in the current
// DurationTriad structure to their
// zero values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	None.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	None.
//
func (durT *DurationTriad) Empty() {
	durT.BaseTime.Empty()
	durT.LocalTime.Empty()
	durT.UTCTime.Empty()
}

// Equal - This method may be used to determine if two
// DurationTriad data structures are equivalent.
//
// If input parameter 'duIn' is equal in all respects to
// the current DurationTriad instance, a boolean value of
// 'true'. Otherwise, the method returns, 'false'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	durT DurationTriad - An instance of DurationTriad which will be compared
//	                     to the current DurationTriad instance in order to
//	                     determine if all data values are equivalent.
//
//	                     A DurationTriad Structure is defined as follows:
//
//	                     type DurationTriad struct {
//	                       BaseTime  TimeDurationDto
//	                       LocalTime TimeDurationDto
//	                       UTCTime   TimeDurationDto
//	                     }
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool - If the method returns 'true' in signals that both the input parameter
//	       DurationTriad and the current DurationTriad instance have equivalent
//	       data values.
//
//	       If the method returns 'false' the two DurationTriad instances are NOT
//	       equal.
//
func (durT *DurationTriad) Equal(duIn DurationTriad) bool {

	if durT.BaseTime.Equal(duIn.BaseTime) &&
		durT.LocalTime.Equal(duIn.LocalTime) &&
		durT.UTCTime.Equal(duIn.UTCTime) {

		return true
	}

	return false

}

// IsValid - Validates the current DurationTriad instance. If the current
// instance is invalid, an error type is returned with an appropriate error
// message.
//
// If the current instance is valid, this method returns 'nil'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	None.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error - If the current DurationTriad instance is valid, the returned error
//	        Type is set equal to 'nil'. If the current DurationTriad instance is
//	        determined to be invalid, this error Type will encapsulate an appropriate
//	        error message.
//
func (durT *DurationTriad) IsValid() error {

	ePrefix := "DurationTriad.IsValid() "

	err := durT.BaseTime.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"INVALID durT.BaseTime. Error='%v'", err.Error())
	}

	err = durT.LocalTime.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"INVALID durT.LocalTime. Error='%v'", err.Error())
	}

	err = durT.UTCTime.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"INVALID durT.UTCTime. Error='%v'", err.Error())
	}

	return nil
}


// New - Creates and returns a new DurationTriad based on time duration calculations
// using input parameters 'startDateTime' and 'endDateTime'.
//
// The Time Zone Location used for time duration calculations is extracted from input
// parameter 'startDateTime'. 'endDateTime' parameter is converted to this Time Zone
// before computing time duration.
//
// This method automatically applies the standard Time Duration allocation, calculation
// type, 'TDurCalcType(0).StdYearMth()'. This means that time duration is allocated over years,
// months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,	microseconds
// and nanoseconds.	For details, see Type 'TDurCalcType' in this source file:
//				MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime time.Time - Starting date time
//
//	endDateTime   time.Time - Ending date time
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad	- Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//
//	du, err := DurationTriad{}.New(
//			startTime,
//			endTime,
//			FmtDateTimeYrMDayFmtStr)
//
//	    Note: 'FmtDateTimeYrMDayFmtStr' is a date format constant defined in
//	          source file 'constantsdatetime.go'.
//
func (durT DurationTriad) New(
	startDateTime time.Time,
	endDateTime time.Time,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.New() "

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return DurationTriad{},
			errors.New(ePrefix + "Both 'startDateTime' and 'endDateTime' input parameters are ZERO!")
	}

	locName := startDateTime.Location().String()

	t2Dur := DurationTriad{}

	err := t2Dur.SetStartEndTimesCalcTz(startDateTime,
		endDateTime,
		TDurCalcType(0).StdYearMth(),
		locName,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{},
			fmt.Errorf(ePrefix+
				"Error returned by SetStartEndTimesCalcTz(...) "+
				"startDateTime='%v'  endDateTime='%v'  Error='%v'",
				startDateTime.Format(FmtDateTimeYrMDayFmtStr),
				endDateTime.Format(FmtDateTimeYrMDayFmtStr),
				err.Error())
	}

	return t2Dur, nil
}

// NewAutoEnd - Creates and returns a new DurationTriad instance. The
// starting date time is provided by input parameter, 'startDateTime'.
// The ending date time is automatically assigned by calling time.Now()
//
// The required input parameter, 'timeZoneLocation' specifies the time zone
// used to configure both starting and ending date time.
//
// Note: This method applies the standard Time Duration allocation, 'TDurCalcType(0).StdYearMth()'.
// This means that duration is allocated over years, months, weeks, weekdays, date days,
// hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
// See Type 'TDurCalcType' for details at:
//	MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime time.Time - Starting date time. This method automatically calls
//	                          time.Now() to compute the ending time.
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//
//	du, err := DurationTriad{}.NewAutoEnd(
// 			startDateTime,
//			TZones.US.Central(),
//			FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewAutoEnd(startDateTime time.Time,
	timeZoneLocation string,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewAutoEnd() "

	endDateTime := time.Now().Local()

	durT2 := DurationTriad{}

	err := durT2.SetStartEndTimesCalcTz(startDateTime,
		endDateTime,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{},
			fmt.Errorf(ePrefix+
				"Error returned by SetStartEndTimesCalcTz() "+
				"startDateTime='%v' endDateTime='%v' Error='%v'",
				startDateTime.Format(FmtDateTimeYrMDayFmtStr),
				endDateTime.Format(FmtDateTimeYrMDayFmtStr),
				err.Error())
	}

	return durT2, err
}

// NewAutoStart - Creates and returns a new DurationTriad instance. Starting date time is
// automatically initialized by calling time.Now(). Afterwards, start date time is converted
// to the Time Zone specified in input parameter, 'timeZoneLocation'.
//
// This method will set ending date time to the same value as starting date time resulting in
// a time duration value of zero.
//
// In order to compute the final time duration value, the user must call the method
// DurationTriad.SetAutoEnd().  At that point, the ending date time will be set by a call to
// time.Now().
//
// Use of these two methods, 'NewAutStart' and 'SetAutoEnd', constitutes a stop watch feature which
// can be triggered to measure elapsed time.
//
// Note: This method applies the standard Time Duration allocation, calculation type
// 'TDurCalcType(0).StdYearMth()'. This means that time duration is allocated over years,
// months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.	For details, see Type 'TDurCalcType' in this source
// file:
//     MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	tDurDto, err := DurationTriad{}.NewAutoStart(
//			TZones.US.Central(),
//			FmtDateTimeYrMDayFmtStr)
//
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewAutoStart(
	timeZoneLocation string,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewAutoStart() "

	startDateTime := time.Now().Local()

	endDateTime := startDateTime

	durT2 := DurationTriad{}

	err := durT2.SetStartEndTimesCalcTz(
		startDateTime,
		endDateTime,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{},
			fmt.Errorf(ePrefix+
				"Error returned by SetStartEndTimesCalcTz() "+
				"startDateTime='%v' Error='%v'",
				startDateTime.Format(FmtDateTimeYrMDayFmtStr),
				err.Error())
	}

	return durT2, nil
}

// NewEndTimeMinusTimeDto - Returns a new DurationTriad based on two input parameters,
// 'endDateTime' and 'timeDto'. 'timeDto' is an instance of TimeDto which is
// subtracted from 'endDateTime' in order to calculate starting date time and time duration.
//
// The Time Zone Location used to compute time duration is extracted from 'endDateTime'. This
// same Time Zone Location will be applied to the resulting starting date time generated by
// the time duration calculation.
//
// This method will automatically apply time duration calculation type,'TDurCalcType(0).StdYearMth()'.
// This is the default calculation type which formats time duration as years, months, days and time.
// For a discussion of Duration Calculation types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	endDateTime	time.Time - Ending date time. The TimeDto parameter (minusTimeDto) will
//	                        be subtracted from this date time in order to compute the
//	                        starting date time.
//
//	minusTimeDto  TimeDto - Provides time values which will be subtracted from
//	                        input parameter, 'endDateTime', in order to calculate
//	                        start Date Time and time duration.
//
//	      A TimeDto structure is defined as follows:
//
//	      type TimeDto struct {
//	        Years                int // Number of Years
//	        Months               int // Number of Months
//	        Weeks                int // Number of Weeks
//	        WeekDays             int // Number of Week-WeekDays.
//	                                 //   Total WeekDays/7 + Remainder WeekDays
//	        DateDays             int // Total Number of Days.
//	                                 //   Weeks x 7 plus WeekDays
//	        Hours                int // Number of Hours.
//	        Minutes              int // Number of Minutes
//	        Seconds              int // Number of Seconds
//	        Milliseconds         int // Number of Milliseconds
//	        Microseconds         int // Number of Microseconds
//	        Nanoseconds	         int // Remaining Nanoseconds after Milliseconds
//	                                 //   and Microseconds
//	        TotSubSecNanoseconds int // Total Nanoseconds:
//	                                 //   Millisecond NanoSecs + Microsecond NanoSecs
//	                                 //   plus remaining Nanoseconds
//	      }
//
//	      Type 'TimeDto' is located in source file:
//	         datetimeopsgo\datetime\timedto.go
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad	- Upon successful completion, this method will return
//			  a new, populated DurationTriad instance.
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewEndTimeMinusTimeDto(
//					startDateTime,
//					minusTimeDto,
//					FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' is a date format constant defined in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewEndTimeMinusTimeDto(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewEndTimeMinusTimeDto() "

	du2 := DurationTriad{}

	locName := endDateTime.Location().String()

	err := du2.SetEndTimeMinusTimeDtoCalcTz(endDateTime,
		minusTimeDto,
		TDurCalcType(0).StdYearMth(),
		locName,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+
			"Error returned from du2.SetEndTimeMinusTimeDtoCalcTz(...). "+
			"Error='%v'", err.Error())
	}

	return du2, nil
}

// NewEndTimeMinusTimeDtoCalcTz - Returns a new DurationTriad based on two input parameters,
// 'endDateTime' and 'timeDto'. 'timeDto' is an instance of TimeDto which is
// subtracted from 'endDateTime' in order to calculate starting date time and time duration.
//
// The user is required to provide Time Zone Location as an input parameter in order
// to ensure the accuracy of time duration calculations. This Time Zone Location is
// applied to both starting and ending date times for the DurationTriad.BaseTime value.
//
// The user is also required to provide the time duration calculation type which will
// control the output of the time duration calculation. The standard date time calculation
// type is, 'TDurCalcType(0).StdYearMth()'. For a discussion of Duration Calculation
// types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	endDateTime	time.Time - Ending date time. The TimeDto parameter (minusTimeDto) will
//				    be subtracted from this date time in order to compute the starting
//				    date time.
//
//	minusTimeDto	TimeDto	- Provides time values which will be subtracted from
//				  input parameter, 'endDateTime', in order to calculate duration.
//
//	      A TimeDto structure is defined as follows:
//
//	      type TimeDto struct {
//	        Years                int // Number of Years
//	        Months               int // Number of Months
//	        Weeks                int // Number of Weeks
//	        WeekDays             int // Number of Week-WeekDays.
//	                                 //   Total WeekDays/7 + Remainder WeekDays
//	        DateDays             int // Total Number of Days.
//	                                 //   Weeks x 7 plus WeekDays
//	        Hours                int // Number of Hours.
//	        Minutes              int // Number of Minutes
//	        Seconds              int // Number of Seconds
//	        Milliseconds         int // Number of Milliseconds
//	        Microseconds         int // Number of Microseconds
//	        Nanoseconds	         int // Remaining Nanoseconds after Milliseconds
//	                                 //   and Microseconds
//	        TotSubSecNanoseconds int // Total Nanoseconds:
//	                                 //   Millisecond NanoSecs + Microsecond NanoSecs
//	                                 //   plus remaining Nanoseconds
//	      }
//
//	      Type 'TimeDto' is defined in source file:
//	         datetimeopsgo\datetime\timedto.go
//
//
//
//	tDurCalcType	TDurCalcType	- Specifies the calculation type to be used in allocating
//			   		  time duration:
//
//				TDurCalcType(0).StdYearMth()	- Default - standard year, month week,
//					  			    day time calculation.
//
//				TDurCalcTypeCUMMONTHS	- Computes cumulative months - no Years.
//
//				TDurCalcTypeCUMWEEKS	- Computes cumulative weeks. No Years or months
//
//				TDurCalcTypeCUMDAYS	- Computes cumulative days. No Years, months or weeks.
//
//				TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//				TDurCalcTypeCUMMINUTES	- Computes cumulative minutes. No Years, months, weeks, days
//							  or hours.
//
//				TDurCalcTypeCUMSECONDS	- Computes cumulative seconds. No Years, months, weeks, days,
//							  hours or minutes.
//
//				TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//							   Used for very large duration values.
//
//			Type 'TDurCalcType' is defined in source file:
//				datetimeopsgo\datetime\timedurationdto.go
//
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewEndTimeMinusTimeDtoCalcTz(
// 						startDateTime,
// 						minusTimeDto,
// 						TDurCalcType(0).StdYearMth(),
// 						TZones.US.Central(),
// 						FmtDateTimeYrMDayFmtStr)
//
//	 Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//	        standard year month day time duration allocation.
//
//	         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	         TZones.US.Central() = "America/Chicago"
//
//	         'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	         source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewEndTimeMinusTimeDtoCalcTz(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewEndTimeMinusTimeDtoTz() "

	du2 := DurationTriad{}

	err := du2.SetEndTimeMinusTimeDtoCalcTz(endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+
			"Error returned from du2.SetEndTimeMinusTimeDtoCalcTz(...). "+
			"Error='%v'", err.Error())
	}

	return du2, nil
}

// NewEndTimeMinusTimeDtoTz - Returns a new DurationTriad based on two input parameters,
// 'endDateTime' and 'timeDto'. 'timeDto' is an instance of TimeDto which is
// subtracted from 'endDateTime' in order to calculate starting date time and time duration.
//
// The user is required to provide Time Zone Location as an input parameter in order to
// ensure the accuracy of time duration calculations. This Time Zone Location is applied
// to both starting and ending date times for the DurationTriad.BaseTime value.
//
// The standard date time calculation type, 'TDurCalcType(0).StdYearMth()' is
// automatically applied by this method. For a discussion of Duration Calculation
// types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	endDateTime   time.Time - Ending date time. The TimeDto parameter (minusTimeDto) will
//	                          be subtracted from this date time in order to compute the
//	                          starting date time.
//
//	minusTimeDto    TimeDto - Provides time values which will be subtracted from
//	                          input parameter, 'endDateTime', in order to calculate
//	                          time duration.
//
//	      A TimeDto structure is defined as follows:
//
//	      type TimeDto struct {
//	        Years                int // Number of Years
//	        Months               int // Number of Months
//	        Weeks                int // Number of Weeks
//	        WeekDays             int // Number of Week-WeekDays.
//	                                 //   Total WeekDays/7 + Remainder WeekDays
//	        DateDays             int // Total Number of Days.
//	                                 //   Weeks x 7 plus WeekDays
//	        Hours                int // Number of Hours.
//	        Minutes              int // Number of Minutes
//	        Seconds              int // Number of Seconds
//	        Milliseconds         int // Number of Milliseconds
//	        Microseconds         int // Number of Microseconds
//	        Nanoseconds	         int // Remaining Nanoseconds after Milliseconds
//	                                 //   and Microseconds
//	        TotSubSecNanoseconds int // Total Nanoseconds:
//	                                 //   Millisecond NanoSecs + Microsecond NanoSecs
//	                                 //   plus remaining Nanoseconds
//	      }
//
//	      Type 'TimeDto' is defined in source file:
//	         datetimeopsgo\datetime\timedto.go
//
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewEndTimeMinusTimeDtoTz(
//			startDateTime,
//			minusTimeDto,
//			TZones.US.Central(),
//			FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewEndTimeMinusTimeDtoTz(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	timeZoneLocation,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewEndTimeMinusTimeDtoTz() "

	du2 := DurationTriad{}

	err := du2.SetEndTimeMinusTimeDtoCalcTz(endDateTime,
		minusTimeDto,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+
			"Error returned from du2.SetEndTimeMinusTimeDtoCalcTz(endDateTime, minusTimeDto). "+
			"Error='%v'", err.Error())
	}

	return du2, nil
}

// NewStartTimeDuration - Returns a New DurationTriad based on 'startDateTz'
// and 'duration' input parameters. Time 'duration' is added to 'startDateTz'
// in order to compute ending date time.
//
// Input parameter 'startDateTz' is of Type, 'DateTzDto'.
//
// This method will except negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
//
// Time Zone Location is extracted from input parameter, 'startDateTz', and
// applied to both starting and ending date times. Applying a common Time Zone
// to both starting and ending date times ensures accurate time duration calculations.
//
// This method automatically applies the standard time duration calculation
// type, 'TDurCalcType(0).StdYearMth()'. The standard time duration calculation type
// allocates time duration by years, months, weeks, days, hours, minutes, seconds,
// milliseconds, microseconds and nanoseconds. For a discussion of Duration Calculation
// types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// The Date Time Format string used to format string displays of date time values will
// be extracted from input parameter 'startDateTz'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTz  DateTzDto - Provides starting date time for duration calculation
//
//	  A DateTzDto structure is defined as follows:
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
//			                         //    associated with this date time.
//		}
//
//	duration time.Duration - Time Duration added to 'startDatTime' in order to
//	                         compute Ending Date Time
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	tDurDto, err := TimeDurationDto{}.NewStartDateTzDuration(
//			startTime,
//			duration)
//
func (durT DurationTriad) NewStartDateTzDuration(
	startDateTime DateTzDto,
	duration time.Duration) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartDateTzDuration() "

	timeZoneLocation := startDateTime.TimeZone.LocationName
	dateTimeFmtStr := startDateTime.DateTimeFmt

	du2 := DurationTriad{}

	err := du2.SetStartTimeDurationCalcTz(startDateTime.GetDateTimeValue(),
		duration,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+
			"Error returned from du2.SetStartTimeDurationCalcTz(startDateTime, duration). "+
			"Error='%v'", err.Error())
	}

	return du2, nil
}

// NewStartEndDateTzDto - Returns a New DurationTriad based on two input
// parameters, 'startDateTime' and 'endDateTime'. These two input parameters
// are submitted as type 'DateTzDto'.
//
// Time Zone Location is extracted from input parameter, 'startDateTime'.
//
// Date Time Format string is likewise extracted from input parameter,
// 'startDateTime'.
//
// This method automatically applies the time duration calculation type,
// 'TDurCalcType(0).StdYearMth()'. This standard time duration calculation type
// allocates time duration by years, months, weeks, days, hours, minutes,
// seconds, milliseconds, microseconds and nanoseconds.
//
// For details on Type 'TDurCalcType', see source file:
//			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime	DateTzDto	- Starting date time
//
//	endDateTime	DateTzDto	- Ending date time
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewStartEndDateTzDto(
//				startTimeDateTz,
//				endTimeDateTz)
//
//
func (durT DurationTriad) NewStartEndDateTzDto(
	startDateTime,
	endDateTime DateTzDto) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartEndDateTzDto() "

	du2 := DurationTriad{}

	timeZoneLocation := startDateTime.TimeZone.LocationName
	dateTimeFmtStr := startDateTime.DateTimeFmt

	err := du2.SetStartEndDateTzCalcTz(
		startDateTime,
		endDateTime,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{},
			fmt.Errorf(ePrefix+
				"Error returned from du2.SetStartEndDateTzCalcTz(...)."+
				"Error='%v'", err)
	}

	return du2, nil
}

// NewStartEndDateTzDtoCalcTz - Returns a New DurationTriad based on two input
// parameters, 'startDateTime' and 'endDateTime'. These two parameters are
// formatted as Type, 'DateTzDto'.
//
// Input parameter, 'timeZoneLocation', is applied to both the starting and ending
// date times before computing date time duration. This ensures accuracy in
// time duration calculations.
//
// The allocation of time duration to years, months, weeks, days, hours etc.
// is controlled by the input parameter calculation type, 'tDurCalcType'.
// For most purposes, the calculation type 'TDurCalcType(0).StdYearMth()' will
// suffice. For details see Type 'TDurCalcType' which is located in
// source file:
// 			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime	DateTzDto	- Starting Date time
//
//	endDateTime	  DateTzDto - Ending Date time
//
//	  A DateTzDto structure is defined as follows:
//
//		type DateTzDto struct {
//			Description      string  // Unused, available for classification, labeling or
//			                         //    description
//			Year                int  // Year Number
//			Month               int  // Month Number
//			Day                 int  // Day Number
//			Hour                int  // Hour Number
//			Minute              int  // Minute Number
//			Second              int  // Second Number
//			Millisecond         int  // Number of MilliSeconds - A Millisecond is
//			                         //    1 one-thousandth or 1/1,000 of a second
//			Microsecond         int  // Number of MicroSeconds - A Microsecond is
//			                         //    1 one-millionth or 1/1,000,000 of a second
//			Nanosecond          int  // Number of Nanoseconds - A Nanosecond is
//			                         //    1 one-billionth or 1/1,000,000,000 of a second.
//			                         //    Nanosecond = TotalNanoSecs
//			                         //    - millisecond nanoseconds - microsecond nanoseconds
//			TotalNanoSecs     int64  // Total Nanoseconds = MilliSecond Nanoseconds
//			                         //    + MicroSeconds Nanoseconds + Nanoseconds
//			DateTime      time.Time  // DateTime value for this DateTzDto Type
//			DateTimeFmt      string  // Date Time Format String. Default is
//			                         //    "2006-01-02 15:04:05.000000000 -0700 MST"
//			TimeZone TimeZoneDefDto  // Contains a detailed description of the Time Zone
//			                         //    and Time Zone Location associated with this
//			                         //    date time.
//		}
//
//
//	tDurCalcType	TDurCalcType	- Specifies the calculation type to be used in allocating
//			   		  time duration:
//
//				TDurCalcType(0).StdYearMth()	- Default - standard year, month week,
//					  			    day time calculation.
//
//				TDurCalcTypeCUMMONTHS	- Computes cumulative months - no Years.
//
//				TDurCalcTypeCUMWEEKS	- Computes cumulative weeks. No Years or months
//
//				TDurCalcTypeCUMDAYS	- Computes cumulative days. No Years, months or weeks.
//
//				TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//				TDurCalcTypeCUMMINUTES	- Computes cumulative minutes. No Years, months, weeks, days
//							  or hours.
//
//				TDurCalcTypeCUMSECONDS	- Computes cumulative seconds. No Years, months, weeks, days,
//							  hours or minutes.
//
//				TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//							   Used for very large duration values.
//
//			Type 'TDurCalcType' is located in source file:
//				MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                 A DurationTriad Structure is defined as follows:
//
//	                 type DurationTriad struct {
//	                   BaseTime  TimeDurationDto
//	                   LocalTime TimeDurationDto
//	                   UTCTime   TimeDurationDto
//	                  }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewStartEndDateTzDtoCalcTz(
//				startDateTzDto,
//				endDateTzDto,
//				TDurCalcType(0).StdYearMth(),
//				TZones.US.Central(),
//				FmtDateTimeYrMDayFmtStr)
//
//	Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//	      standard year month day time duration allocation.
//
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartEndDateTzDtoCalcTz(
	startDateTime,
	endDateTime DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartEndDateTzDtoCalcTz() "

	du2 := DurationTriad{}

	err := du2.SetStartEndDateTzCalcTz(
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{},
			fmt.Errorf(ePrefix+
				"Error returned from du2.SetStartEndDateTzCalcTz(...)."+
				"Error='%v'", err)
	}

	return du2, nil
}

// NewStartEndDateTzDtoTz - Returns a New DurationTriad based on two input
// parameters, 'startDateTime' and 'endDateTime'. These parameters are
// configured as Type, 'DateTzDto'.
//
// 'startDateTime' and 'endDateTime' are first converted to the Time Zone
// specified by input parameter, 'timeZoneLocation' before computing date
// time duration.
//
// The standard date time calculation type, 'TDurCalcType(0).StdYearMth()' is
// automatically applied by this method. For a discussion of Duration Calculation
// types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime DateTzDto - Starting date time
//
//	endDateTime   DateTzDto - Ending date time
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                 A DurationTriad Structure is defined as follows:
//
//	                 type DurationTriad struct {
//	                   BaseTime  TimeDurationDto
//	                   LocalTime TimeDurationDto
//	                   UTCTime   TimeDurationDto
//	                  }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewStartEndDateTzDtoTz(
//				        startDateTzDto,
//				        endDateTzDto,
//				        TZones.US.Central(),
//				        FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartEndDateTzDtoTz(
	startDateTime,
	endDateTime DateTzDto,
	timeZoneLocation,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartEndDateTzDtoTz() "

	du2 := DurationTriad{}

	err := du2.SetStartEndDateTzCalcTz(
		startDateTime,
		endDateTime,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{},
			fmt.Errorf(ePrefix+
				"Error returned from du2.SetStartEndDateTzCalcTz(...)."+
				"Error='%v'", err)
	}

	return du2, nil
}

// NewStartEndTimesCalcTz - Returns a New DurationTriad based on two input
// parameters, 'startDateTime' and 'endDateTime'.
//
// Input parameter, 'timeZoneLocation', is applied to both the starting and ending
// date times before computing date time duration. This ensures accuracy in
// time duration calculations.
//
// The allocation of time duration to years, months, weeks, days, hours etc.
// is controlled by the input parameter calculation type, 'tDurCalcType'.
// For most purposes, the calculation type 'TDurCalcType(0).StdYearMth()' will
// suffice. For details see Type 'TDurCalcType' which is located in
// source file:
// 			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime time.Time    - Starting date time
//
//	endDateTime   time.Time    - Ending date time
//
//	tDurCalcType  TDurCalcType - Specifies the calculation type to be used in allocating
//				     time duration:
//
//				TDurCalcType(0).StdYearMth()	- Default - standard year, month week,
//					  			    day time calculation.
//
//				TDurCalcTypeCUMMONTHS	- Computes cumulative months - no Years.
//
//				TDurCalcTypeCUMWEEKS	- Computes cumulative weeks. No Years or months
//
//				TDurCalcTypeCUMDAYS	- Computes cumulative days. No Years, months or weeks.
//
//				TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//				TDurCalcTypeCUMMINUTES	- Computes cumulative minutes. No Years, months, weeks, days
//							  or hours.
//
//				TDurCalcTypeCUMSECONDS	- Computes cumulative seconds. No Years, months, weeks, days,
//							  hours or minutes.
//
//				TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//							   Used for very large duration values.
//
//			Type 'TDurCalcType' is located in source file:
//				MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
///
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                 A DurationTriad Structure is defined as follows:
//
//	                 type DurationTriad struct {
//	                   BaseTime  TimeDurationDto
//	                   LocalTime TimeDurationDto
//	                   UTCTime   TimeDurationDto
//	                  }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewStartEndTimesCalcTz(
//				startTime,
//				endTime,
//				TDurCalcType(0).StdYearMth(),
//				TZones.US.Central(),
//				FmtDateTimeYrMDayFmtStr)
//
//	Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//	      standard year month day time duration allocation.
//
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartEndTimesCalcTz(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartEndTimesCalcTz() "

	du2 := DurationTriad{}

	err := du2.SetStartEndTimesCalcTz(startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+"Error returned from du2.SetStartEndTimesCalcTz(startDateTime, endDateTime).\nError='%v'", err)
	}

	return du2, nil
}

// NewStartEndTimesTz - Returns a New DurationTriad based on two input
// parameters, 'startDateTime' and 'endDateTime'.
//
// 'startDateTime' and 'endDateTime' are first converted to the Time Zone
// specified by input parameter, 'timeZoneLocation' before computing date
// time duration.
//
// The standard date time calculation type, 'TDurCalcType(0).StdYearMth()' is
// automatically applied by this method. For a discussion of Duration Calculation
// types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime time.Time	- Starting date time
//
//	endDateTime   time.Time - Ending date time
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewStartEndTimesTz(
//				startTime,
//				endTime,
//				TZones.US.Central(),
//				FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartEndTimesTz(startDateTime,
	endDateTime time.Time, timeZoneLocation, dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartTimeDurationTz() "

	du2 := DurationTriad{}

	err := du2.SetStartEndTimesCalcTz(startDateTime, endDateTime, TDurCalcType(0).StdYearMth(), timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+"Error returned from du2.SetStartEndTimesCalcTz(startDateTime, endDateTime).\nError='%v'", err)
	}

	return du2, nil

}

// NewStartTimeDuration - Returns a New DurationTriad based on 'startDateTime'
// and 'duration' input parameters. Time 'duration' is added to 'startDateTime'
// in order to compute ending date time.
//
// This method will except negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
//
// Time Zone Location is extracted from input parameter, 'startDateTime', and
// applied to both starting and ending date times. Applying a common Time Zone
// to both starting and ending date times ensures accurate time duration calculations.
//
// This method automatically applies the standard time duration calculation
// type, 'TDurCalcType(0).StdYearMth()'. The standard time duration calculation type
// allocates time duration by years, months, weeks, days, hours, minutes, seconds,
// milliseconds, microseconds and nanoseconds. For a discussion of Duration Calculation
// types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime time.Time - Starting Date Time for duration calculation
//
//	duration  time.Duration - Time Duration added to 'startDatTime' in order to
//	                          compute Ending Date Time
//
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	tDurDto, err := TimeDurationDto{}.NewStartTimeDuration(
//				startTime,
//				duration,
//				FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' is a constant defined in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartTimeDuration() "

	timeZoneLocation := startDateTime.Location().String()

	du2 := DurationTriad{}

	err := du2.SetStartTimeDurationCalcTz(startDateTime,
		duration,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+
			"Error returned from du2.SetStartTimeDurationCalcTz(startDateTime, duration). "+
			"Error='%v'", err.Error())
	}

	return du2, nil
}

// NewStartTimeDurationCalcTz - Returns a New DurationTriad based on 'startDateTime'
// and 'duration' input parameters. Time 'duration' is added to 'startDateTime'
// in order to compute ending date time.
//
// The user is required to submit an input parameter for Time Zone Location.
// This Time Zone Location will convert the 'startDateTime' parameter to the
// specified time zone before computing ending date time.
//
// This method will except negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
//
// The user is also required to provide the time duration calculation type which will
// control the output of the time duration calculation. The standard date time calculation
// type is, 'TDurCalcType(0).StdYearMth()'. For a discussion of Duration Calculation
// types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime    time.Time - Starting Date Time for duration calculation
//
//	duration     time.Duration - Time Duration added to 'startDatTime' in order to
//	                              compute Ending Date Time
//
//	tDurCalcType  TDurCalcType - Specifies the calculation type to be used in allocating
//				     time duration:
//
//				TDurCalcType(0).StdYearMth()	- Default - standard year, month week,
//					  			    day time calculation.
//
//				TDurCalcTypeCUMMONTHS	- Computes cumulative months - no Years.
//
//				TDurCalcTypeCUMWEEKS	- Computes cumulative weeks. No Years or months
//
//				TDurCalcTypeCUMDAYS	- Computes cumulative days. No Years, months or weeks.
//
//				TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//				TDurCalcTypeCUMMINUTES	- Computes cumulative minutes. No Years, months, weeks, days
//							  or hours.
//
//				TDurCalcTypeCUMSECONDS	- Computes cumulative seconds. No Years, months, weeks, days,
//							  hours or minutes.
//
//				TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//							   Used for very large duration values.
//
//			Type 'TDurCalcType' is located in source file:
//				MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	tDurDto, err := TimeDurationDto{}.NewStartTimeDurationCalcTz(
//				startTime,
//				duration,
//				TDurCalcType(0).StdYearMth(),
//				TZones.US.Central(),
//				FmtDateTimeYrMDayFmtStr)
//
//	Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//	      standard year month day time duration allocation.
//
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartTimeDurationCalcTz(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartTimeDurationCalcTz() "

	du2 := DurationTriad{}

	err := du2.SetStartTimeDurationCalcTz(startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+
			"Error returned from du2.SetStartTimeDurationCalcTz(startDateTime, duration). "+
			"Error='%v'", err.Error())
	}

	return du2, nil
}

// NewStartTimeDurationTz - Returns a New DurationTriad based on 'startDateTime'
// and 'duration' input parameters. Time 'duration' is added to 'startDateTime'
// in order to compute ending date time.
//
// The user is required to submit an input parameter for Time Zone Location.
// This Time Zone Location will convert the 'startDateTime' parameter to the
// specified time zone before computing ending date time.
//
// The method will except negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
//
// Note: 	This method applies the standard Time Duration allocation, calculation type
// 				'TDurCalcType(0).StdYearMth()'. This means that duration is allocated over years,
// 				months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
// 				microseconds and nanoseconds.	See Type 'TDurCalcType' for details.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime time.Time - Starting Date Time for duration calculation
//
//	duration time.Duration 	- Time Duration added to 'startDatTime' in order to
//	                          compute Ending Date Time
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewStartTimeDurationTz(
//				startDateTime,
//				duration,
//				TZones.US.Central(),
//				FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartTimeDurationTz(
	startDateTime time.Time,
	duration time.Duration,
	timeZoneLocation,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartTimeDurationTz() "

	du2 := DurationTriad{}

	err := du2.SetStartTimeDurationCalcTz(startDateTime,
		duration,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+
			"Error returned from du2.SetStartTimeDurationCalcTz(...). "+
			"Error='%v'", err.Error())

	}

	return du2, nil
}

// NewStartTimePlusTimeDto - Returns a new DurationTriad based on two input
// parameters, 'startDateTime' and 'timeDto'. 'timeDto' is an instance of TimeDto
// which is added to 'startDateTime' in order to calculate ending date time as
// well as time duration.
//
// The Time Zone Location used in configuring both starting and ending datetime is
// extracted from the input parameter, 'startDateTime'. Using a common time zone
// ensures the accuracy of time duration calculations. This extracted Time Zone
// Location is applied to both starting and ending date times for the
// DurationTriad.BaseTime value.
//
// This method automatically applies the The standard date time calculation
// type, 'TDurCalcType(0).StdYearMth()'. This standard date time calculation type
// allocates time duration by years, months, weeks, days, hours, minutes, seconds,
// milliseconds, microseconds and nanoseconds. For a discussion of Time Duration
// Calculation types, see Type 'TDurCalcType' located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime	time.Time - Starting date time. The TimeDto parameter will be added
//	                            to this date time in order to compute the ending date time.
//
//	plusTimeDto     TimeDto - Provides time values which will be added to
//	                          'startDateTime' in order to calculate time duration.
//
//				type TimeDto struct {
//				  Years			int // Number of Years
//				  Months		int // Number of Months
//				  Weeks			int // Number of Weeks
//				  WeekDays		int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//				  DateDays		int // Total Number of Days. Weeks x 7 plus WeekDays
//				  Hours			int // Number of Hours.
//				  Minutes		int // Number of Minutes
//				  Seconds		int // Number of Seconds
//				  Milliseconds		int // Number of Milliseconds
//				  Microseconds		int // Number of Microseconds
//				  Nanoseconds		int // Remaining Nanoseconds after Milliseconds & Microseconds
//				  TotSubSecNanoseconds	int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//							    // 	plus remaining Nanoseconds
//				}
//
//				Type 'TimeDto' is located in source file:
//					MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewStartTimePlusTimeDto(
//					startDateTime,
//					FmtDateTimeYrMDayFmtStr)
//
//	Note: 'FmtDateTimeYrMDayFmtStr' is a constant defined in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartTimePlusTimeDto(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartTimePlusTimeDto() "

	timeZoneLocation := startDateTime.Location().String()

	du2 := DurationTriad{}

	err := du2.SetStartTimePlusTimeDtoCalcTz(startDateTime,
		plusTimeDto,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{},
			fmt.Errorf(ePrefix+
				"Error returned from du2.SetStartTimePlusTimeDtoCalcTz(startDateTime, plusTimeDto). "+
				"Error='%v'", err)
	}

	return du2, nil
}

// NewStartTimePlusTimeDtoCalcTz - Returns a new DurationTriad based on two input
// parameters, 'startDateTime' and 'timeDto'. 'timeDto' is an instance of TimeDto
// which is added to 'startDateTime' in order to calculate ending date time as
// well as time duration.
//
// The user is required to provide Time Zone Location as an input parameter in order
// to ensure the accuracy of time duration calculations. This Time Zone Location is
// applied to both starting and ending date times for the DurationTriad.BaseTime value.
//
// The user is also required to provide the time duration calculation type which will
// control the output of the time duration calculation. The standard date time calculation
// type is, 'TDurCalcType(0).StdYearMth()'. For a discussion of Duration Calculation
// types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime time.Time - Starting date time. The TimeDto parameter, 'plusTimeDto'
//	                          will be added to this date time in order to compute the
//	                          ending date time.
//
//	plusTimeDto TimeDto     - An instance of TimeDto containing time values,
//	                          (Years, Months, weeks, days, hours, minutes etc.)
//	                          which will be added to input parameter 'startDateTime'
//	                          in order to compute the Ending Date Time and Time
//	                          Duration.
//
//	                          type TimeDto struct {
//	                            Years                  int // Number of Years
//	                            Months                 int // Number of Months
//	                            Weeks                  int // Number of Weeks
//	                            WeekDays               int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//	                            DateDays               int // Total Number of Days. Weeks x 7 plus WeekDays
//	                            Hours                  int // Number of Hours.
//	                            Minutes                int // Number of Minutes
//	                            Seconds                int // Number of Seconds
//	                            Milliseconds           int // Number of Milliseconds
//	                            Microseconds           int // Number of Microseconds
//	                            Nanoseconds            int // Remaining Nanoseconds after Milliseconds & Microseconds
//	                            TotSubSecNanoseconds   int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//	                                                       // plus remaining Nanoseconds
//				  }
//
//				  Type 'TimeDto' is located in source file:
//				 	 MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//	tDurCalcType  TDurCalcType - Specifies the calculation type to be used in allocating
//				     time duration:
//
//				TDurCalcType(0).StdYearMth()	- Default - standard year, month week,
//					  			    day time calculation.
//
//				TDurCalcTypeCUMMONTHS	- Computes cumulative months - no Years.
//
//				TDurCalcTypeCUMWEEKS	- Computes cumulative weeks. No Years or months
//
//				TDurCalcTypeCUMDAYS	- Computes cumulative days. No Years, months or weeks.
//
//				TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//				TDurCalcTypeCUMMINUTES	- Computes cumulative minutes. No Years, months, weeks, days
//							  or hours.
//
//				TDurCalcTypeCUMSECONDS	- Computes cumulative seconds. No Years, months, weeks, days,
//							  hours or minutes.
//
//				TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//							   Used for very large duration values.
//
//				Type 'TDurCalcType' is located in source file:
//					MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewStartTimePlusTimeDtoCalcTz(
//				startDateTime,
//				plusTimeDto,
//				TDurCalcType(0).StdYearMth(),
//				TZones.US.Central(),
//				FmtDateTimeYrMDayFmtStr)
//
//	Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' and signals
//	      standard year month day time duration allocation.
//
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartTimePlusTimeDtoCalcTz(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartTimePlusTimeDtoCalcTz() "

	du2 := DurationTriad{}

	err := du2.SetStartTimePlusTimeDtoCalcTz(startDateTime,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return DurationTriad{},
			fmt.Errorf(ePrefix+
				"Error returned from du2.SetStartTimePlusTimeDtoCalcTz(startDateTime, plusTimeDto). "+
				"Error='%v'", err)
	}

	return du2, nil
}

// NewStartTimePlusTimeDtoTz - Returns a new DurationTriad based on two input parameters,
// 'startDateTime' and 'timeDto'. 'timeDto' is an instance of Type TimeDto which is
// added to 'startDateTime' in order to calculate ending date time as well as time
// duration.
//
// The user is required to provide Time Zone Location as an input parameter in order
// to ensure that time duration calculations are performed using equivalent time zones.
//
// The standard date time calculation type, 'TDurCalcType(0).StdYearMth()' is
// automatically applied by this method. For a discussion of Duration Calculation
// types, see Type TDurCalcType located in source file:
// 					'MikeAustin71\datetimeopsgo\datetime\timedurationdto.go'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime time.Time - Starting date time
//
//	plusTimeDto     TimeDto - Provides time values which will be added to
//	                          'startDateTime' in order to calculate duration.
//
//				type TimeDto struct {
//				  Years			int // Number of Years
//				  Months		int // Number of Months
//				  Weeks			int // Number of Weeks
//				  WeekDays		int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//				  DateDays		int // Total Number of Days. Weeks x 7 plus WeekDays
//				  Hours			int // Number of Hours.
//				  Minutes		int // Number of Minutes
//				  Seconds		int // Number of Seconds
//				  Milliseconds		int // Number of Milliseconds
//				  Microseconds		int // Number of Microseconds
//				  Nanoseconds		int // Remaining Nanoseconds after Milliseconds & Microseconds
//				  TotSubSecNanoseconds	int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//							    // 	plus remaining Nanoseconds
//				}
//
//				Type 'TimeDto' is located in source file:
//					MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	DurationTriad - Upon successful completion, this method will return
//	                a new, populated DurationTriad instance.
//
//	                A DurationTriad Structure is defined as follows:
//
//	                type DurationTriad struct {
//	                  BaseTime  TimeDurationDto
//	                  LocalTime TimeDurationDto
//	                  UTCTime   TimeDurationDto
//	                }
//
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which encapsulates an
//	                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	du, err := DurationTriad{}.NewStartTimePlusTimeDtoTz(
//				startDateTime,
//				plusTimeDto,
//				TZones.US.Central(),
//				FmtDateTimeYrMDayFmtStr)
//
//	Note: TZones.US.Central() = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT DurationTriad) NewStartTimePlusTimeDtoTz(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	timeZoneLocation,
	dateTimeFmtStr string) (DurationTriad, error) {

	ePrefix := "DurationTriad.NewStartTimePlusTimeDtoTz() "

	du2 := DurationTriad{}

	err := du2.SetStartTimePlusTimeDtoCalcTz(startDateTime, plusTimeDto, TDurCalcType(0).StdYearMth(), timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DurationTriad{}, fmt.Errorf(ePrefix+"Error returned from du2.SetStartTimePlusTimeDtoTz(startDateTime, plusTimeDto).\nError='%v'", err)
	}

	return du2, nil
}

// SetAutoEnd - When called, this method automatically sets the ending date
// time and re-calculates the time duration for the current DurationTriad
// instance.
//
// Ending date time is assigned the value returned by time.Now(). This ending
// date time is converted to the specified Time Zone specified by the Time Zone
// Location associated with the current starting date time value.
//
// When used together, the two methods 'NewAutoStart' and this method, 'SetAutoEnd'
// function as a stop watch feature. Simply calling these functions can set
// the starting date time and later, the ending date time to measure elapsed time
// or time duration.
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetAutoEnd()
//
func (durT *DurationTriad) SetAutoEnd() error {
	ePrefix := "DurationTriad.SetAutoEnd() "

	endDateTime := time.Now().Local()

	calcType := durT.BaseTime.CalcType
	startDateTime := durT.BaseTime.StartTimeDateTz.GetDateTimeValue()
	tzLocName := durT.BaseTime.StartTimeDateTz.TimeZone.LocationName
	fmtStr := durT.BaseTime.StartTimeDateTz.DateTimeFmt

	err := durT.SetStartEndTimesCalcTz(startDateTime, endDateTime, calcType, tzLocName, fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from SetStartEndTimesCalcTz() "+
			"startDateTime='%v'  endDateTime='%v'  Error='%v'",
			startDateTime.Format(FmtDateTimeYrMDayFmtStr),
			endDateTime.Format(FmtDateTimeYrMDayFmtStr),
			err.Error())
	}

	return nil
}

// SetEndTimeMinusTimeDto - Calculates duration values based on an Ending Date Time and
// a TimeDto structure consisting of time values (Years, Months, weeks, days, hours,
// minutes etc.). The time values in the 'timeDto' parameter are subtracted
// from 'endDateTime'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time. As a result, true values
// for 'StartTimeDateTz', 'EndTimeDateTz' and 'TimeDuration' are stored in
// the DurationTriad data structure.
//
// Time Zone Location is extracted from input parameter, 'endDateTime'. The extracted
// time zone is applied to both the starting and ending date times in order to
// ensure the accuracy of time duration calculations.
//
// This method automatically applies the time duration calculation type,
// 'TDurCalcType(0).StdYearMth()'. The standard time duration calculation type allocates
// time duration over years, months, weeks, days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds. For a details on 'TDurCalcType', see the source file:
//																MikeAustin71\datetimeopsgo\datetime\timedto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	endDateTime   time.Time - The ending date time value from which TimeDto
//	                          parameter 'minusTimeDto' will be subtracted
//	                          in order to compute the Starting Date Time.
//
//	minusTimeDto    TimeDto - An instance of TimeDto containing time values,
//	                          (Years, Months, weeks, days, hours, minutes etc.)
//	                          which will be subtracted from input parameter
//	                          'endDateTime' in order to compute the Starting
//	                          Date Time and Time Duration.
//
//				type TimeDto struct {
//				  Years			int // Number of Years
//				  Months		int // Number of Months
//				  Weeks			int // Number of Weeks
//				  WeekDays		int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//				  DateDays		int // Total Number of Days. Weeks x 7 plus WeekDays
//				  Hours			int // Number of Hours.
//				  Minutes		int // Number of Minutes
//				  Seconds		int // Number of Seconds
//				  Milliseconds		int // Number of Milliseconds
//				  Microseconds		int // Number of Microseconds
//				  Nanoseconds		int // Remaining Nanoseconds after Milliseconds & Microseconds
//				  TotSubSecNanoseconds	int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//							    // 	plus remaining Nanoseconds
//				}
//
//				Type 'TimeDto' is located in source file:
//					MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetEndTimeMinusTimeDto(
//	       endDateTime,
//	       minusTimeDto,
//	       FmtDateTimeYrMDayFmtStr)
//
//	Note: 'FmtDateTimeYrMDayFmtStr' are constants defined in
//	      source file 'constantsdatetime.go'
//
func (durT *DurationTriad) SetEndTimeMinusTimeDto(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetEndTimeMinusTimeDtoCalcTz() "

	timeZoneLocation := endDateTime.Location().String()

	err := durT.SetEndTimeMinusTimeDtoCalcTz(
		endDateTime,
		minusTimeDto,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by SetEndTimeMinusTimeDtoCalcTz Error='%v'",
			err.Error())
	}

	return nil
}

// SetEndTimeMinusTimeDtoCalcTz - Calculates duration values based on an Ending Date Time and
// a TimeDto structure consisting of time values (Years, Months, weeks, days, hours,
// minutes etc.). The time values in the 'timeDto' parameter are subtracted
// from 'endDateTime'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time.
//
// As a result, true values for StartTimeDateTz, EndTimeDateTz and TimeDuration
// are stored in the DurationTriad data structure.
//
// Input parameter, 'timeZoneLocation', is applied to both the starting and ending
// date times before computing date time duration. This ensures accuracy in
// time duration calculations.
//
// The allocation of time duration to years, months, weeks, days, hours etc.
// is controlled by the input parameter calculation type, 'tDurCalcType'.
// For most purposes, the calculation type 'TDurCalcType(0).StdYearMth()' will
// suffice. For details see Type 'TDurCalcType' which is located in
// source file:
// 			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	endDateTime   time.Time - The ending date time value from which TimeDto
//	                           parameter 'minusTimeDto' will be subtracted
//	                           in order to compute the Starting Date Time.
//
//	minusTimeDto  TimeDto   - An instance of TimeDto containing time values,
//	                           (Years, Months, weeks, days, hours, minutes etc.)
//	                           which will be subtracted from input parameter
//	                           'endDateTime' in order to compute the Starting
//	                           Date Time and Time Duration.
//
//	                          type TimeDto struct {
//	                            Years                  int // Number of Years
//	                            Months                 int // Number of Months
//	                            Weeks                  int // Number of Weeks
//	                            WeekDays               int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//	                            DateDays               int // Total Number of Days. Weeks x 7 plus WeekDays
//	                            Hours                  int // Number of Hours.
//	                            Minutes                int // Number of Minutes
//	                            Seconds                int // Number of Seconds
//	                            Milliseconds           int // Number of Milliseconds
//	                            Microseconds           int // Number of Microseconds
//	                            Nanoseconds            int // Remaining Nanoseconds after Milliseconds & Microseconds
//	                            TotSubSecNanoseconds   int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//	                                                       // plus remaining Nanoseconds
//	                           }
//
//	                           Type 'TimeDto' is located in source file:
//	                           	MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//	tDurCalcType  TDurCalcType - Specifies the calculation type to be used in allocating
//				     time duration:
//
//				TDurCalcType(0).StdYearMth()	- Default - standard year, month week,
//					  			    day time calculation.
//
//				TDurCalcTypeCUMMONTHS	- Computes cumulative months - no Years.
//
//				TDurCalcTypeCUMWEEKS	- Computes cumulative weeks. No Years or months
//
//				TDurCalcTypeCUMDAYS	- Computes cumulative days. No Years, months or weeks.
//
//				TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//				TDurCalcTypeCUMMINUTES	- Computes cumulative minutes. No Years, months, weeks, days
//							  or hours.
//
//				TDurCalcTypeCUMSECONDS	- Computes cumulative seconds. No Years, months, weeks, days,
//							  hours or minutes.
//
//				TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//							   Used for very large duration values.
//
//				Type 'TDurCalcType' is located in source file:
//					MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	         The source file 'constantsdatetime.go' contains a number of
//	         constant declarations covering the more frequently used time
//	         zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//	         time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error    - If this method completes successfully, the returned error
//	           Type is set equal to 'nil'. If an error condition is encountered,
//	           this method will return an error Type which encapsulates an
//	           appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetEndTimeMinusTimeDtoCalcTz(
//	       startDateTime,
//	       minusTimeDto,
//	       TDurCalcType(0).StdYearMth(),
//	       TZones.US.Central(),
//	       FmtDateTimeYrMDayFmtStr)
//
//
//	Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' which is defined in
//				source file 'timedurationdto.go'.
//
//	      TZones.US.Central() = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT *DurationTriad) SetEndTimeMinusTimeDtoCalcTz(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetEndTimeMinusTimeDtoCalcTz() "

	fmtStr := durT.preProcessDateFormatStr(dateTimeFmtStr)

	tzLoc := durT.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: TimeZoneLocation input parameter is INVALID! "+
			"timeZoneLocation='%v' tzLoc='%v'  Error='%v'",
			timeZoneLocation, tzLoc, err.Error())
	}

	baseTime, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoCalcTz(endDateTime,
		minusTimeDto,
		tDurCalcType,
		tzLoc,
		fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error computing baseTime. TimeDurationDto{}.NewEndTimeMinusTimeDtoTz(). "+
			"endDateTime='%v' Error='%v'",
			endDateTime, err.Error())
	}

	_, err = time.LoadLocation(TZones.Local())

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: Local TimeZoneLocation input parameter is INVALID! "+
			"timeZoneLocation='%v' Error='%v'",
			TZones.Local(), err.Error())
	}

	localTime, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoCalcTz(endDateTime,
		minusTimeDto,
		tDurCalcType,
		TZones.Local(),
		fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error computing localTime. TimeDurationDto{}.NewEndTimeMinusTimeDtoTz(). "+
			"endDateTime='%v' Error='%v'",
			endDateTime, err.Error())
	}

	_, err = time.LoadLocation(TZones.Local())

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: UTC TimeZoneLocation input parameter is INVALID! "+
			"timeZoneLocation='%v' Error='%v'",
			TZones.Local(), err.Error())
	}

	utcTime, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoCalcTz(
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		TZones.Local(),
		fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error computing localTime. TimeDurationDto{}.NewEndTimeMinusTimeDtoTz(). "+
			"endDateTime='%v' Error='%v'",
			endDateTime, err.Error())
	}

	durT.Empty()
	durT.BaseTime = baseTime.CopyOut()
	durT.LocalTime = localTime.CopyOut()
	durT.UTCTime = utcTime.CopyOut()

	err = durT.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: Duration Base Data INVALID! Error='%v'", err.Error())
	}

	return nil

}

// SetEndTimeMinusTimeDtoTz - Calculate duration values based on an Ending Date Time and
// a TimeDto structure consisting of time values (Years, Months, weeks, days, hours,
// minutes etc.). The time values in the 'timeDto' parameter are subtracted
// from 'endDateTime'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time.
//
// As a result, true values for StartTimeDateTz, EndTimeDateTz and TimeDuration
// are stored in the DurationTriad data structure.
//
// Input parameter, 'timeZoneLocation', is applied to both the starting and ending
// date times before computing date time duration. This ensures accuracy in
// time duration calculations.
//
// This method automatically applies the time duration calculation type, 'TDurCalcType(0).StdYearMth()'.
// The standard time duration calculation type allocates time duration over years, months,
// weeks, days, hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//
// For a details on 'TDurCalcType', see the source file:
//									MikeAustin71\datetimeopsgo\datetime\timedto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	endDateTime   time.Time - The ending date time value from which TimeDto
//	                          parameter 'minusTimeDto' will be subtracted
//	                          in order to compute the Starting Date Time.
//
//	minusTimeDto  TimeDto   - An instance of TimeDto containing time values,
//	                          (Years, Months, weeks, days, hours, minutes etc.)
//	                          which will be subtracted from input parameter
//	                          'endDateTime' in order to compute the Starting
//	                          Date Time and Time Duration.
//
//	                          type TimeDto struct {
//	                            Years                  int // Number of Years
//	                            Months                 int // Number of Months
//	                            Weeks                  int // Number of Weeks
//	                            WeekDays               int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//	                            DateDays               int // Total Number of Days. Weeks x 7 plus WeekDays
//	                            Hours                  int // Number of Hours.
//	                            Minutes                int // Number of Minutes
//	                            Seconds                int // Number of Seconds
//	                            Milliseconds           int // Number of Milliseconds
//	                            Microseconds           int // Number of Microseconds
//	                            Nanoseconds            int // Remaining Nanoseconds after Milliseconds & Microseconds
//	                            TotSubSecNanoseconds   int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//	                                                       // plus remaining Nanoseconds
//	                           }
//
//	                           Type 'TimeDto' is located in source file:
//	                           	MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetEndTimeMinusTimeDtoTz(
//	       endDateTime,
//	       minusTimeDto,
//	       TZones.US.Central(),
//	       FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT *DurationTriad) SetEndTimeMinusTimeDtoTz(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetEndTimeMinusTimeDtoCalcTz() "

	err := durT.SetEndTimeMinusTimeDtoCalcTz(endDateTime, minusTimeDto,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by SetEndTimeMinusTimeDtoCalcTz() "+
			"Error='%v'", err.Error())
	}

	return nil
}

// SetStartEndDateTzCalcTz - Calculates duration values and save the results in the current DurationTriad
// data fields. Calculations are based on a starting date time and an ending date time passed
// to the method as Type DateTzDto.
//
// Input parameter, 'timeZoneLocation', is applied to both the starting and ending
// date times before computing date time duration. This ensures accuracy in
// time duration calculations.
//
// The allocation of time duration to years, months, weeks, days, hours etc.
// is controlled by the input parameter calculation type, 'tDurCalcType'.
// For most purposes, the calculation type 'TDurCalcType(0).StdYearMth()' will
// suffice. For details see Type 'TDurCalcType' which is located in
// source file:
// 			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime   DateTzDto - Starting date time
//
//	endDateTime     DateTzDto - Ending date time
//
//	tDurCalcType TDurCalcType - Specifies the calculation type to be used in allocating
//	                            time duration:
//
//					TDurCalcType(0).StdYearMth()    - Default - standard year, month
//	                                    week day time calculation.
//
//					TDurCalcTypeCUMMONTHS     - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS      - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS       - Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS      - Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeCUMMINUTES    - Computes cumulative minutes. No Years, months, weeks, days
//					                            or hours.
//
//					TDurCalcTypeCUMSECONDS    - Computes cumulative seconds. No Years, months, weeks, days,
//					                            hours or minutes.
//
//					TDurCalcTypeGregorianYrs  - Computes Years based on average length of a Gregorian Year
//					                            Used for very large duration values.
//
//					Type 'TDurCalcType' is located in source file:
//					  MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error     - If this method completes successfully, the returned error
//	            Type is set equal to 'nil'. If an error condition is encountered,
//	            this method will return an error Type which encapsulates an
//	            appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartEndDateTzCalcTz(
//	       startDateTime,
//	       endDateTime,
//	       TDurCalcType(0).StdYearMth(),
//	       TZones.US.Central(),
//	       FmtDateTimeYrMDayFmtStr)
//
//
//	Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' which is defined in
//				source file 'timedurationdto.go'.
//
//	Note: TZones.US.Central() = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT *DurationTriad) SetStartEndDateTzCalcTz(
	startDateTime,
	endDateTime DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartEndDateTzCalcTz() "

	err := durT.SetStartEndTimesCalcTz(startDateTime.GetDateTimeValue(),
		endDateTime.GetDateTimeValue(),
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by SetStartEndTimesCalcTz(...). Error='%v'", err.Error())
	}

	return nil

}

// SetStartEndTimes - Calculates duration values and save the results in the current DurationTriad
// data fields. Calculations are based on a starting date time and an ending date time passed
// to the method.
//
// Time Zone Location is extracted from the input parameter, 'startDateTime'.
// The extracted time zone is applied to both starting and ending date time in order
// to ensure the accuracy of time duration calculations.
//
// This method automatically applies the time duration calculation type, 'TDurCalcType(0).StdYearMth()'.
// The standard time duration calculation type allocates time duration by years, months, weeks,
// days, hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//
// For details on Type 'TDurCalcType', see source file:
//			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	startDateTime time.Time - Starting date time
//
//	endDateTime   time.Time - Ending date time
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartEndTimes(
//	       startDateTime,
//	       endDateTime,
//	       FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' are constants defined in
//	      source file 'constantsdatetime.go'
//
func (durT *DurationTriad) SetStartEndTimes(
	startDateTime,
	endDateTime time.Time,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartEndTimes() "

	locName := startDateTime.Location().String()

	err := durT.SetStartEndTimesCalcTz(startDateTime,
		endDateTime,
		TDurCalcType(0).StdYearMth(),
		locName,
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by SetStartEndTimesCalcTz(...). Error='%v'", err.Error())
	}

	return nil
}

// SetStartEndTimesCalcTz - Calculates duration values and save the results in the data fields
// of the current DurationTriad instance. Calculations are based on a starting date time and
// an ending date time passed to the method. This method requires the user to specify a
// 'timeZoneLocation' input parameter which ensures that starting date time and ending date
// time will be converted to a common Time Zone before being used to compute time duration.
//
// The allocation of time duration to years, months, weeks, days, hours etc.
// is controlled by the input parameter calculation type, 'tDurCalcType'.
// For most purposes, the calculation type 'TDurCalcType(0).StdYearMth()' will
// suffice. For details see Type 'TDurCalcType' which is located in
// source file:
// 			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	startDateTime   time.Time - Starting date time
//
//	endDateTime     time.Time - Ending date time
//
//	tDurCalcType TDurCalcType - Specifies the calculation type to be used in allocating
//	                            time duration:
//
//					TDurCalcType(0).StdYearMth()    - Default - standard year, month
//	                                    week day time calculation.
//
//					TDurCalcTypeCUMMONTHS     - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS      - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS       - Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS      - Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeCUMMINUTES    - Computes cumulative minutes. No Years, months, weeks, days
//					                            or hours.
//
//					TDurCalcTypeCUMSECONDS    - Computes cumulative seconds. No Years, months, weeks, days,
//					                            hours or minutes.
//
//					TDurCalcTypeGregorianYrs  - Computes Years based on average length of a Gregorian Year
//					                            Used for very large duration values.
//
//					Type 'TDurCalcType' is located in source file:
//					  MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartEndTimesCalcTz(
//	       startDateTime,
//	       endDateTime,
//	       TDurCalcType(0).StdYearMth(),
//	       TZones.US.Central(),
//	       FmtDateTimeYrMDayFmtStr)
//
//	Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' which is defined in
//				source file 'timedurationdto.go'.
//
//	      TZones.US.Central() = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT *DurationTriad) SetStartEndTimesCalcTz(
		startDateTime,
		endDateTime time.Time,
		tDurCalcType TDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartEndTimesCalcTz() "

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return errors.New(ePrefix +
			"Error: Input parameters 'startDateTime' and 'endDateTime' are ZERO!")
	}

	tzLoc := durT.preProcessTimeZoneLocation(timeZoneLocation)

	fmtStr := durT.preProcessDateFormatStr(dateTimeFmtStr)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Input paramenter 'timeZoneLocation' is INVALID. "+
			" time.LoadLocation(tzLoc). timeZoneLocation='%v', tzLoc='%v', Error='%v'",
			timeZoneLocation, tzLoc, err.Error())
	}

	baseTime, err := TimeDurationDto{}.NewStartEndTimesCalcTz(startDateTime,
		endDateTime, tDurCalcType, tzLoc, fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"baseTime error returned by TimeDurationDto{}.NewStartEndTimesCalcTz(). "+
			"Error='%v' ", err.Error())
	}

	localTime, err := TimeDurationDto{}.NewStartEndTimesCalcTz(startDateTime,
		endDateTime, tDurCalcType, TZones.Local(), fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"localTime error returned by TimeDurationDto{}.NewStartEndTimesCalcTz(). "+
			"Error='%v' ", err.Error())
	}

	utcTime, err := TimeDurationDto{}.NewStartEndTimesCalcTz(startDateTime,
		endDateTime, tDurCalcType, TZones.UTC(), fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by TimeDurationDto{}.NewStartEndTimesCalcTz(startDateTime," +
			" endDateTime, tDurCalcType, TZones.UTC(), fmtStr)\n" +
			"Error='%v'\n", err.Error())
	}

	durT.Empty()
	durT.BaseTime = baseTime.CopyOut()
	durT.LocalTime = localTime.CopyOut()
	durT.UTCTime = utcTime.CopyOut()

	err = durT.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"ERROR: Duration Base Data is INVALID! Error='%v'", err.Error())
	}

	return nil
}

// SetStartEndTimesTz - Calculates duration values and save the results in the DurationTriad
// data fields. Calculations are based on a starting date time and an ending date time passed
// to the method. This method requires the user to input a 'timeZoneLocation' thus ensuring
// that both starting date time and ending date time are calculated using a common time zone.
//
// Input parameter, 'timeZoneLocation', is applied to both the starting and ending
// date times before computing date time duration. This ensures accuracy in
// time duration calculations.
//
// This method automatically applies the time duration calculation type, 'TDurCalcType(0).StdYearMth()'.
// The standard time duration calculation type allocates time duration by years, months, weeks,
// days, hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//
// For details on Type 'TDurCalcType', see source file:
//			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	startDateTime time.Time - Starting date time
//
//	endDateTime   time.Time - Ending date time
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartEndTimesTz(
//	       startDateTime,
//	       endDateTime
//	       TZones.US.Central(),
//	       FmtDateTimeYrMDayFmtStr)
//
//	Note: TZones.US.Central() = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT *DurationTriad) SetStartEndTimesTz(
	startDateTime,
	endDateTime time.Time,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartEndTimesTz() "

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return errors.New(ePrefix +
			"Error: Input parameters 'startDateTime' and 'endDateTime' are ZERO!")
	}

	err := durT.SetStartEndTimesCalcTz(startDateTime, endDateTime, TDurCalcType(0).StdYearMth(),
		timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by SetStartEndTimesCalcTz()  Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimeDuration - Receives a starting date time and proceeds to calculate
// the ending date time, duration and populates the DurationTriad data fields.
//
// The method will except negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
//
// Time Zone Location is extracted from the input parameter, 'startDateTime'.
// The extracted time zone is applied to both starting and ending date time in order
// to ensure the accuracy of time duration calculations.
//
// This method automatically applies the time duration calculation type, 'TDurCalcType(0).StdYearMth()'.
// The standard time duration calculation type allocates time duration by years, months, weeks,
// days, hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//
// For details on Type 'TDurCalcType', see source file:
//			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	startDateTime time.Time - Starting Date Time for duration calculation
//
//	duration time.Duration  - Time Duration added to 'startDateTime' in order to
//														compute Ending Date Time.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartTimeDuration(
//	       startDateTime,
//	       duration,
//	       FmtDateTimeYrMDayFmtStr)
//
//
//	Note: 'FmtDateTimeYrMDayFmtStr' are constants defined in
//	      source file 'constantsdatetime.go'
//
func (durT *DurationTriad) SetStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartTimeDuration() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error - Start Time is Zero!")
	}

	timeZoneLocation := startDateTime.Location().String()

	err := durT.SetStartTimeDurationCalcTz(startDateTime,
		duration,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned bySetStartTimeDurationCalcTz()  Error='%v'",
			err.Error())
	}

	return nil
}

// SetStartTimeDurationCalcTz - Receives a starting date time and calculates
// a time duration. The method then calculates the ending date time, duration
// and populates the DurationTriad data fields.
//
// The method will except negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
//
// The user is also required to submit input parameters to time zone location and
// date time calculation type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	startDateTime   time.Time - Starting Date Time for duration calculation
//
//	duration    time.Duration - Time Duration added to 'startDateTime' in order to
//	                            compute Ending Date Time
//
//	tDurCalcType TDurCalcType - Specifies the calculation type to be used in allocating
//	                            time duration:
//
//					TDurCalcType(0).StdYearMth()    - Default - standard year, month
//					                            week day time calculation.
//
//					TDurCalcTypeCUMMONTHS     - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS      - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS       - Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS      - Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeCUMMINUTES    - Computes cumulative minutes. No Years, months, weeks, days
//					                            or hours.
//
//					TDurCalcTypeCUMSECONDS    - Computes cumulative seconds. No Years, months, weeks, days,
//					                            hours or minutes.
//
//					TDurCalcTypeGregorianYrs  - Computes Years based on average length of a Gregorian Year
//					                            Used for very large duration values.
//
//					Type 'TDurCalcType' is located in source file:
//					  MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartTimeDurationCalcTz(
//	       startDateTime,
//	       duration,
//	       TDurCalcType(0).StdYearMth(),
//	       TZones.US.Central(),
//	       FmtDateTimeYrMDayFmtStr)
//
//
//	Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' which is defined in
//				source file 'timedurationdto.go'.
//
//	      TZones.US.Central() = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT *DurationTriad) SetStartTimeDurationCalcTz(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartTimeDurationCalcTz() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error - Start Time is Zero!")
	}

	tzLoc := durT.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: Input Parameter 'timeZoneLocation' INVALID. "+
			"timeZoneLocation='%v' tzLoc='%v' Error='%v'",
			timeZoneLocation, tzLoc, err.Error())
	}

	fmtStr := durT.preProcessDateFormatStr(dateTimeFmtStr)

	baseTime, err := TimeDurationDto{}.NewStartTimeDurationCalcTz(startDateTime, duration,
		tDurCalcType, tzLoc, fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"baseTime calculation Error returned by "+
			"TimeDurationDto{}.NewStartTimeDurationCalcTz() Error='%v'",
			err.Error())

	}

	localTime, err := TimeDurationDto{}.NewStartTimeDurationCalcTz(startDateTime, duration,
		tDurCalcType, TZones.Local(), fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"localTime calculation Error returned by "+
			"TimeDurationDto{}.NewStartTimeDurationCalcTz() Error='%v'",
			err.Error())

	}

	utcTime, err := TimeDurationDto{}.NewStartTimeDurationCalcTz(startDateTime, duration,
		tDurCalcType, TZones.Local(), fmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"utcTime calculation Error returned by "+
			"TimeDurationDto{}.NewStartTimeDurationCalcTz() Error='%v'",
			err.Error())

	}

	durT.Empty()
	durT.BaseTime = baseTime.CopyOut()
	durT.LocalTime = localTime.CopyOut()
	durT.UTCTime = utcTime.CopyOut()

	err = durT.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"ERROR: Duration Base Data is INVALID! Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimeDurationTz - Receives a starting date time and a time duration.
// This method then calculates the ending date time, duration and populates the
// DurationTriad data fields.
//
// The Method will except negative time durations. This means that the starting
// date time will be reclassified as the ending date time and the duration will
// be subtracted from that ending date time to calculate the correct starting
// date time.
//
// In addition, the user is also required to specify the Time Zone Location.
// Both starting and ending date times will be configured for this standard
// time zone in order to ensure accuracy of time duration calculations.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	startDateTime time.Time - Starting Date Time for duration calculation
//
//	duration  time.Duration - Time Duration added to 'startDatTime' in order to
//	                          compute Ending Date Time
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartTimeDurationTz(
//	       startDateTime,
//	       TZones.US.Central(),
//	       FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//	      TZones.US.Central() = "America/Chicago"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT *DurationTriad) SetStartTimeDurationTz(
	startDateTime time.Time,
	duration time.Duration,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartTimeDurationTz() "

	err := durT.SetStartTimeDurationCalcTz(startDateTime, duration,
		TDurCalcType(0).StdYearMth(), timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+"ERROR: Duration Base Data is INVALID! Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimePlusTimeDto - Calculates time duration values based on a Starting Date Time
// plus time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'plusTimeDto' parameter. The 'plusTimeDto' parameter is added to 'startDateTime' in
// order to calculate ending date time and duration.
//
// Values in the 'plusTimeDto' parameter are automatically converted to positive numeric
// values before being added to parameter 'startDateTime'.
//
// True values for starting date time, ending date time and time duration are then stored in
// the DurationTriad data structure.
//
// Time Zone Location is extracted from the input parameter, 'startDateTime'.
// The extracted time zone is applied to both starting and ending date time in order
// to ensure the accuracy of time duration calculations.
//
// This method automatically applies the time duration calculation type, 'TDurCalcType(0).StdYearMth()'.
// The standard time duration calculation type allocates time duration by years, months, weeks,
// days, hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//
// For details on Type 'TDurCalcType', see source file:
//			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	startDateTime time.Time - Starting time
//
//	plusTimeDto     TimeDto - Provides time values which will be subtracted from
//	                          'startDateTime' in order to calculate duration.
//
//
//				type TimeDto struct {
//				  Years			int // Number of Years
//				  Months		int // Number of Months
//				  Weeks			int // Number of Weeks
//				  WeekDays		int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//				  DateDays		int // Total Number of Days. Weeks x 7 plus WeekDays
//				  Hours			int // Number of Hours.
//				  Minutes		int // Number of Minutes
//				  Seconds		int // Number of Seconds
//				  Milliseconds		int // Number of Milliseconds
//				  Microseconds		int // Number of Microseconds
//				  Nanoseconds		int // Remaining Nanoseconds after Milliseconds & Microseconds
//				  TotSubSecNanoseconds	int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//							    // 	plus remaining Nanoseconds
//				}
//
//				Type 'TimeDto' is located in source file:
//					MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartTimePlusTimeDto(
//	       startDateTime,
//	       plusTimeDto,
//	       FmtDateTimeYrMDayFmtStr)
//
//	Note: FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'FmtDateTimeYrMDayFmtStr' are constants defined in
//	      source file 'constantsdatetime.go'
//
func (durT *DurationTriad) SetStartTimePlusTimeDto(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartTimePlusTimeDto() "

	timeZoneLocation := startDateTime.Location().String()

	err := durT.SetStartTimePlusTimeDtoCalcTz(
		startDateTime,
		plusTimeDto,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by SetStartTimePlusTimeDtoCalcTz().  Error='%v'",
			err.Error())
	}

	return nil
}

// SetStartTimePlusTimeDtoCalcTz - Calculates time duration values based on a Starting Date Time
// plus time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'plusTimeDto' parameter. The 'plusTimeDto' parameter is added to 'startDateTime' in
// order to calculate ending date time and duration.
//
// Values in the 'plusTimeDto' parameter are automatically converted to positive numeric
// values before being added to parameter 'startDateTime'.
//
// True values for starting date time, ending date time and time duration are then stored in
// the DurationTriad data structure.
//
// Input parameter, 'timeZoneLocation', is applied to both the starting and ending
// date times before computing date time duration. This ensures accuracy in
// time duration calculations.
//
// The allocation of time duration to years, months, weeks, days, hours etc.
// is controlled by the input parameter calculation type, 'tDurCalcType'.
// For most purposes, the calculation type 'TDurCalcType(0).StdYearMth()' will
// suffice. For details see Type 'TDurCalcType' which is located in
// source file:
// 			MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime   time.Time - Starting date time. Input parameter 'plusTimeDto'
//	                            will be added to this starting date time in order
//	                            to generate ending date time.
//
//	plusTimeDto       TimeDto - Provides time values which will be added to
//	                            'startDateTime' in order to calculate duration.
//
//	                          type TimeDto struct {
//	                            Years                  int // Number of Years
//	                            Months                 int // Number of Months
//	                            Weeks                  int // Number of Weeks
//	                            WeekDays               int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//	                            DateDays               int // Total Number of Days. Weeks x 7 plus WeekDays
//	                            Hours                  int // Number of Hours.
//	                            Minutes                int // Number of Minutes
//	                            Seconds                int // Number of Seconds
//	                            Milliseconds           int // Number of Milliseconds
//	                            Microseconds           int // Number of Microseconds
//	                            Nanoseconds            int // Remaining Nanoseconds after Milliseconds & Microseconds
//	                            TotSubSecNanoseconds   int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//	                                                       // plus remaining Nanoseconds
//	                           }
//
//	                           Type 'TimeDto' is located in source file:
//	                           	MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//
//	tDurCalcType TDurCalcType - Specifies the calculation type to be used in allocating
//	                            time duration:
//
//					TDurCalcType(0).StdYearMth()    - Default - standard year, month
//					                            week day time calculation.
//
//					TDurCalcTypeCUMMONTHS     - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS      - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS       - Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS      - Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeCUMMINUTES    - Computes cumulative minutes. No Years, months, weeks, days
//					                            or hours.
//
//					TDurCalcTypeCUMSECONDS    - Computes cumulative seconds. No Years, months, weeks, days,
//					                            hours or minutes.
//
//					TDurCalcTypeGregorianYrs  - Computes Years based on average length of a Gregorian Year
//					                            Used for very large duration values.
//
//					Type 'TDurCalcType' is located in source file:
//					  MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartTimePlusTimeDtoCalcTz(
//	       startDateTime,
//	       plusTimeDto,
//	       TDurCalcType(0).StdYearMth(),
//	       TZones.US.Central(),
//	       FmtDateTimeYrMDayFmtStr)
//
//	Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' which is defined in
//				source file 'timedurationdto.go'.
//
//	      TZones.US.Central() = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT *DurationTriad) SetStartTimePlusTimeDtoCalcTz(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartTimePlusTimeDtoTz() "

	tzLoc := durT.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: TimeZoneLocation is INVALID! "+
			"timeZoneLocation='%v'  tzLoc='%v'  Error='%v'",
			timeZoneLocation, tzLoc, err.Error())
	}

	baseTime, err := TimeDurationDto{}.NewStartTimePlusTimeDtoCalcTz(
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		tzLoc,
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+"baseTime calculation error returned by "+
			"TimeDurationDto{}.NewStartTimePlusTimeDtoCalcTz(). Error=%v'",
			err.Error())
	}

	_, err = time.LoadLocation(TZones.Local())

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: Local TimeZoneLocation is INVALID! "+
			"timeZoneLocation='%v' Error='%v'",
			TZones.Local(), err.Error())
	}

	localTime, err := TimeDurationDto{}.NewStartTimePlusTimeDtoCalcTz(
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		TZones.Local(),
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+"localTime calculation error returned by "+
			"TimeDurationDto{}.NewStartTimePlusTimeDtoCalcTz(). Error=%v'",
			err.Error())
	}

	_, err = time.LoadLocation(TZones.UTC())

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: UTC TimeZoneLocation is INVALID! "+
			"timeZoneLocation='%v' Error='%v'",
			TZones.UTC(), err.Error())
	}

	utcTime, err := TimeDurationDto{}.NewStartTimePlusTimeDtoCalcTz(
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		TZones.UTC(),
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+"utcTime calculation error returned by "+
			"TimeDurationDto{}.NewStartTimePlusTimeDtoCalcTz(). Error=%v'",
			err.Error())
	}

	durT.Empty()
	durT.BaseTime = baseTime.CopyOut()
	durT.LocalTime = localTime.CopyOut()
	durT.UTCTime = utcTime.CopyOut()

	err = durT.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: Duration Base Data INVALID! Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimePlusTimeDtoTz - Calculate duration values based on a Starting Date Time and
// time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'times' parameter. The 'timeDto' parameter is added to 'StartTimeDateTz'.
//
// Values in the 'timeDto' parameter are automatically converted to positive numeric
// values before being added to 'StartTimeDateTz'.
//
// True values for StartTimeDateTz, EndTimeDateTz and TimeDuration are then stored in
// the DurationTriad data structure.
//
// Time Zone Location is extracted from input parameter, 'startDateTime'. The extracted
// time zone is applied to both the starting and ending date times in order to
// ensure the accuracy of time duration calculations.
//
// This method automatically applies the time duration calculation type, 'TDurCalcType(0).StdYearMth()'.
// The standard time duration calculation type allocates time duration by years, months, weeks,
// days, hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	startDateTime time.Time - Starting date time. This date time value is
//	                          added to
//
//	plusTimeDto     TimeDto - Provides time values which will be subtracted from
//	                          'startDateTime' in order to calculate duration.
//
//	                          type TimeDto struct {
//	                            Years                  int // Number of Years
//	                            Months                 int // Number of Months
//	                            Weeks                  int // Number of Weeks
//	                            WeekDays               int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//	                            DateDays               int // Total Number of Days. Weeks x 7 plus WeekDays
//	                            Hours                  int // Number of Hours.
//	                            Minutes                int // Number of Minutes
//	                            Seconds                int // Number of Seconds
//	                            Milliseconds           int // Number of Milliseconds
//	                            Microseconds           int // Number of Microseconds
//	                            Nanoseconds            int // Remaining Nanoseconds after Milliseconds & Microseconds
//	                            TotSubSecNanoseconds   int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//	                                                       // plus remaining Nanoseconds
//	                           }
//
//	                           Type 'TimeDto' is located in source file:
//	                           	MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//	timeZoneLocation string - time zone location must be designated as one of
//	                          two values:
//
//				(1) The string 'Local' - signals the designation of the local time zone
//				    location for the host computer.
//
//				(2) IANA Time Zone Location -
//				    See https://golang.org/pkg/time/#LoadLocation
//				    and https://www.iana.org/time-zones to ensure that
//				    the IANA Time Zone Database is properly configured
//				    on your system. Note: IANA Time Zone Data base is
//				    equivalent to 'tz database'.
//
//				    Examples:
//				      "America/New_York"
//				      "America/Chicago"
//				      "America/Denver"
//				      "America/Los_Angeles"
//				      "Pacific/Honolulu"
//
//				     The source file 'constantsdatetime.go' contains a number of
//				     constant declarations covering the more frequently used time
//				     zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//				     time zone constants begin with the prefix 'TzIana'.
//
//	dateTimeFmtStr string   - A date time format string which will be used
//	                          to format and display 'dateTime'. Example:
//	                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	                          Date time format constants are found in the source
//	                          file 'constantsdatetime.go'. These constants represent
//	                          the more commonly used date time string formats. All
//	                          Date Time format constants begin with the prefix
//	                          'FmtDateTime'.
//
//	                          If 'dateTimeFmtStr' is submitted as an
//	                          'empty string', a default date time format
//	                          string will be applied. The default date time
//	                          format string is:
//	                            FmtDateTimeYrMDayFmtStr =
//	                                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Value
//
//	error		- If this method completes successfully, the returned error
//			  Type is set equal to 'nil'. If an error condition is encountered,
//			  this method will return an error Type which encapsulates an
//			  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	dt  := DurationTriad{}
//	err := dt.SetStartTimePlusTimeDtoTz(
//	       startDateTime,
//	       plusTimeDto,
//	       TZones.US.Central(),
//	       FmtDateTimeYrMDayFmtStr)
//
//	Note: TZones.US.Central() = "America/Chicago"
//	      FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//	      source file 'constantsdatetime.go'.
//
func (durT *DurationTriad) SetStartTimePlusTimeDtoTz(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	timeZoneLocation,
	dateTimeFmtStr string) error {

	ePrefix := "DurationTriad.SetStartTimePlusTimeDtoTz() "

	err := durT.SetStartTimePlusTimeDtoCalcTz(
		startDateTime,
		plusTimeDto,
		TDurCalcType(0).StdYearMth(),
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by SetStartTimePlusTimeDtoCalcTz(). Error='%v'",
			err.Error())
	}

	return nil
}

func (durT *DurationTriad) preProcessDateFormatStr(dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}

func (durT *DurationTriad) preProcessTimeZoneLocation(timeZoneLocation string) string {

	if len(timeZoneLocation) == 0 {
		return TZones.UTC()
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return "Local"
	}

	return timeZoneLocation
}
