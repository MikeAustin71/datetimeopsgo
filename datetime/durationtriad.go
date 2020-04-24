package datetime

import (
	"errors"
	"sync"
	"time"
)

/*

Source Files

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
 time, an ending datetime. Thereafter, the time duration computed by subtracting
 starting date time from ending date time.

 In time duration calculations, time zone location is important. If starting and
 ending date time use different time zones, it could create errors in the time
 duration result. Also, depending on daylight savings time, the same time duration
 could produce differing ending date times depending on which time zone is used.

 In order to ensure accuracy, the DurationTriad type first calculates date time duration
 for the user specified time zone and then proceeds to calculate that same duration
 using the 'Local' Time Zone Location, and the 'UTC' Time Zone Location.

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
 Time.

 Reference:
   https://en.wikipedia.org/wiki/Coordinated_Universal_Time


Time Duration Calculation Types -

 Some methods provided by the DurationTriad Type allow the user
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

  Time zone location must be designated as one of three types of time zones:

  (1) Time Zone 'Local' -
        The string 'Local' signals the designation of the local time zone
        location for the host computer. 'Local' is a creation of the
        Go Programming Language. Reference Package Time, The Go Programming
        Language: https://golang.org/pkg/time/

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

        The source file 'timezonedata.go' contains a number of
        constant declarations covering the more frequently used time
        zones. Example: 'TZones.US.Central()' = "America/Chicago". All
        time zone constants begin with the prefix 'TZones'.

  (3) A Military Time Zone
        In addition to military operations, Military
        time zones are commonly used in aviation as
        well as at sea. They are also known as nautical
        or maritime time zones.
        Reference:
        https://en.wikipedia.org/wiki/List_of_military_time_zones
        http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
        https://www.timeanddate.com/time/zones/military
        https://www.timeanddate.com/worldclock/timezone/alpha
        https://www.timeanddate.com/time/map/

        Examples:
         "Alpha"   or "A"
         "Bravo"   or "B"
         "Charlie" or "C"
         "Delta"   or "D"
         "Zulu"    or "Z"

         If the time zone "Zulu" is passed to this method, it will be
         classified as a Military Time Zone.

       Note:
           The source file 'timezonedata.go' contains over 600 constant
           time zone declarations covering all IANA and Military Time
           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
           time zone constants begin with the prefix 'TZones'.


DurationTriad Structure

=========================================================================

*/
type DurationTriad struct {
	BaseTime  TimeDurationDto
	LocalTime TimeDurationDto
	UTCTime   TimeDurationDto
	lock      *sync.Mutex
}

// CopyIn - Receives and incoming DurationTriad instance. This method then
// copies the incoming data values into the current DurationTriad data
// structure. This method performs a deep copy on all data elements.
//
// NOTE: This method will alter the data fields of the current DurationTriad
// instance.
//
// __________________________________________________________________________
//
// Input Parameter
//
//  duIn DurationTriad - An instance of type 'DurationTriad'. The data fields
//                       of input parameter 'duIn' will be copied to the
//                       the data fields of the current DurationTriad instance.
//                       The type of copy operation performed is a 'deep' copy.
//
//                       When this operation completes, 'duIn' and the current
//                       'DurationTriad' instance will be equivalent.
//
//                       A DurationTriad Structure is defined as follows:
//
//                       type DurationTriad struct {
//                         BaseTime  TimeDurationDto
//                         LocalTime TimeDurationDto
//                         UTCTime   TimeDurationDto
//                        }
//
// __________________________________________________________________________
//
// Return Values
//
//  None.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  durTriad := DurationTriad{}
//
//  durTriad.CopyIn(dTriad2)
//
func (durT *DurationTriad) CopyIn(durTIn DurationTriad) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	durTUtil := durationTriadUtility{}

	durTUtil.copyIn(durT, &durTIn, "DurationTriad.CopyIn() ")

	return
}

// CopyOut - Creates and returns a new DurationTriad instance. The deep
// copy operation copies all data elements from the current
// DurationTriad instance to the new DurationTriad instance which is
// returned to the calling function.
//
//
// __________________________________________________________________________
//
// Input Parameters
//
//  None.
//
// __________________________________________________________________________
//
// Return Value
//
//  DurationTriad - Upon completion, this method returns a new instance of
//                  Type DurationTriad which is, in all respects, an exact
//                  copy of the current DurationTriad instance.
//
//                  A DurationTriad Structure is defined as follows:
//
//                  type DurationTriad struct {
//                    BaseTime  TimeDurationDto
//                    LocalTime TimeDurationDto
//                    UTCTime   TimeDurationDto
//                  }
//
// __________________________________________________________________________
//
// Example Usage:
//
//  durTriad := DurationTriad{}
//
//  durTriad2 := durTriad.CopyOut()
//
func (durT *DurationTriad) CopyOut() DurationTriad {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	durT.lock.Unlock()

	durTUtil := durationTriadUtility{}

	return durTUtil.copyOut(durT, "DurationTriad.CopyOut() ")
}

// Empty - This method initializes all data fields in the
// current DurationTriad structure to their zero or
// uninitialized values.
//
// __________________________________________________________________________
//
// Input Parameters
//
// None.
//
// __________________________________________________________________________
//
// Return Values
//
// None.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  durTriad := DurationTriad{}
//
//  durTriad.Empty()
//
func (durT *DurationTriad) Empty() {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	durT.lock.Unlock()

	durTUtil := durationTriadUtility{}

	durTUtil.empty(durT, "DurationTriad.Empty() ")

	return
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
//  duIn DurationTriad - An instance of DurationTriad which will be compared
//                       to the current DurationTriad instance in order to
//                       determine if all data values are equivalent.
//
//                       A DurationTriad Structure is defined as follows:
//
//                       type DurationTriad struct {
//                         BaseTime  TimeDurationDto
//                         LocalTime TimeDurationDto
//                         UTCTime   TimeDurationDto
//                       }
//
// __________________________________________________________________________
//
// Return Value
//
// bool - If the method returns 'true' in signals that both the input parameter
//        DurationTriad and the current DurationTriad instance have equivalent
//        data values.
//
//        If the method returns 'false' the two DurationTriad instances are NOT
//        equal.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  durTriad := DurationTriad{}
//
//  areEqual := durTriad.Equal(dTriad2)
//
func (durT *DurationTriad) Equal(duIn DurationTriad) bool {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	durTUtil := durationTriadUtility{}

	areEqual, _ := durTUtil.equal(
		durT,
		&duIn,
		"DurationTriad.Equal() ")

	return areEqual
}

// IsValid - Validates the current DurationTriad instance. If the current
// instance is invalid, an error type is returned with an appropriate error
// message.
//
// If the current instance is valid, this method returns 'nil'.
//
// __________________________________________________________________________
//
// Input Parameters
//
//  None.
//
// __________________________________________________________________________
//
// Return Values
//
//  error - If the current DurationTriad instance is valid, the returned error
//          Type is set equal to 'nil'. If the current DurationTriad instance is
//          determined to be invalid, this error Type will encapsulate an appropriate
//          error message.
//
//
// __________________________________________________________________________
//
// Example Usage:
//
//  durTriad := DurationTriad{}
//
//  err := durTriad.IsValid()
//
//  if err != nil {
//     return err
//  }
//
func (durT *DurationTriad) IsValid() error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.IsValid() "

	durTUtil := durationTriadUtility{}

	return durTUtil.isValid(durT, ePrefix)
}

// NewStartEndTimes - Returns a new DurationTriad instance
// initialized to zero values.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  durTriad := DurationTriad{}.New()
//
//  Note: Member variables in 'durTriad' are now
//        set to their zero values.
//
func (durT DurationTriad) New() DurationTriad {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	durTriad2 := DurationTriad{}
	durTriad2.lock = new(sync.Mutex)
	durTriad2.BaseTime = TimeDurationDto{}.New()
	durTriad2.LocalTime = TimeDurationDto{}.New()
	durTriad2.UTCTime = TimeDurationDto{}.New()

	return durTriad2
}

// NewAutoEnd - Creates and returns a new DurationTriad instance. The
// starting date time is provided by input parameter, 'startDateTime'.
// The ending date time is automatically assigned by calling time.Now()
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
//     - Starting date time. This method automatically calls
//       time.Now() to compute the ending time.
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
//             datetime\timemathcalcmode.go
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
// Return Values
//
//  DurationTriad
//     - Upon successful completion, this method will return
//       a new, populated DurationTriad instance.
//
//       A DurationTriad Structure is defined as follows:
//
//       type DurationTriad struct {
//         BaseTime  TimeDurationDto
//         LocalTime TimeDurationDto
//         UTCTime   TimeDurationDto
//       }
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  du, err := DurationTriad{}.NewAutoEnd(
//                    startDateTime,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT DurationTriad) NewAutoEnd(
	startDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewAutoEnd() "

	endDateTime := time.Now().Local()

	durT2 := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setStartEndTimes(
		&durT2,
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DurationTriad{}, err
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
// Use of these two methods, 'NewAutoStart' and 'SetAutoEnd', constitutes a stop watch feature
// which can be triggered to measure elapsed time.
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
//  tDurCalcType TDurCalcType
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
//  timeZoneLocation   string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
//  DurationTriad
//     - Upon successful completion, this method will return
//       a new, populated DurationTriad instance.
//
//       A DurationTriad Structure is defined as follows:
//
//       type DurationTriad struct {
//         BaseTime  TimeDurationDto
//         LocalTime TimeDurationDto
//         UTCTime   TimeDurationDto
//       }
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//  tDurDto, err := DurationTriad{}.NewAutoStart(
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT DurationTriad) NewAutoStart(
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewAutoStart() "

	startDateTime := time.Now().Local()

	endDateTime := startDateTime

	durT2 := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setStartEndTimes(
		&durT2,
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DurationTriad{}, err
	}

	return durT2, nil
}

// NewEndTimeMinusTimeDto - Returns a new DurationTriad based on two input parameters,
// 'endDateTime' and 'timeDto'. 'timeDto' is an instance of TimeDto which is
// subtracted from 'endDateTime' in order to calculate starting date time and time duration.
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
//     - Ending date time. The TimeDto parameter (minusTimeDto) will
//       be subtracted from this date time in order to compute the
//       starting date time.
//
//
//  minusTimeDto    TimeDto
//     - Provides time values which will be subtracted from
//       input parameter, 'endDateTime', in order to calculate
//       start Date Time and time duration.
//
//       A TimeDto structure is defined as follows:
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
//       time duration. This Type is configured as an enumeration.
//       Member values may be accessed directly using the syntax
//       TDurCalcType(0).StdYearMth(). Alternatively, an abbreviated
//       syntax may be used by means of the global variable, 'TDurCalc'.
//       Example: TDurCalc.StdYearMth()
//
//       Valid enumerations are listed as follows:
//
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
//  timeZoneLocation string
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
//               datetime\timemathcalcmode.go
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
// __________________________________________________________________________
//
// Example Usage:
//
//  du, err := DurationTriad{}.NewEndTimeMinusTimeDto(
//                    endDateTime,
//                    minusTimeDto,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT DurationTriad) NewEndTimeMinusTimeDto(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewEndTimeMinusTimeDto() "

	durT2 := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setEndTimeMinusTimeDto(
		&durT2,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DurationTriad{}, err
	}

	return durT2, nil
}

// NewEndTimeTzMinusTimeDto - Returns a new DurationTriad based on two input parameters,
// 'endDateTimeTz' and 'timeDto'. 'timeDto' is an instance of TimeDto which is
// subtracted from 'endDateTimeTz' in order to calculate starting date time and time
// duration.
//
// Input parameter 'endDateTimeTz' is formatted as an instance of 'DateTzDto'.
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
//  endDateTimeTz     DateTzDto
//     - Ending date time. The TimeDto parameter (minusTimeDto) will
//       be subtracted from this date time in order to compute the starting
//       date time and time duration.
//
//  minusTimeDto      TimeDto
//     - Provides time values which will be subtracted from
//       input parameter, 'endDateTime', in order to calculate duration.
//
//       A TimeDto structure is defined as follows:
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
//            datetimeopsgo\datetime\timedto.go
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
//           Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
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
//             datetime\timemathcalcmode.go
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
//  dateTimeFmtStr string
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
//  DurationTriad
//     - Upon successful completion, this method will return
//       a new, populated DurationTriad instance.
//
//       A DurationTriad Structure is defined as follows:
//
//         type DurationTriad struct {
//           BaseTime  TimeDurationDto
//           LocalTime TimeDurationDto
//           UTCTime   TimeDurationDto
//         }
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
//
// __________________________________________________________________________
//
// Example Usage:
//
//  du, err := DurationTriad{}.NewEndTimeTzMinusTimeDto(
//                    startDateTime,
//                    minusTimeDto,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT DurationTriad) NewEndTimeTzMinusTimeDto(
	endDateTimeTz DateTzDto,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewEndTimeTzMinusTimeDto() "

	durT2 := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setEndTimeMinusTimeDto(
		&durT2,
		endDateTimeTz.dateTimeValue,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DurationTriad{}, err
	}

	return durT2, nil
}

// NewStartEndTimes - Creates and returns a new DurationTriad based on time duration calculations
// using input parameters 'startDateTime' and 'endDateTime'.
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  startDateTime time.Time
//     - Starting date time
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
// Return Values
//
//  DurationTriad
//     - Upon successful completion, this method will return
//       a new, populated DurationTriad instance.
//
//       A DurationTriad Structure is defined as follows:
//
//       type DurationTriad struct {
//         BaseTime  TimeDurationDto
//         LocalTime TimeDurationDto
//         UTCTime   TimeDurationDto
//       }
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage
//
//
//  du, err := DurationTriad{}.NewStartEndTimes(
//                    startTime,
//                    endTime,
//                    TDurCalc.StdYearMth(),
//                    TZones.US.Central(),
//                    TCalcMode.LocalTimeZone(),
//                    FmtDateTimeYrMDayFmtStr)
//
// Note:
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT DurationTriad) NewStartEndTimes(
	startDateTime time.Time,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewStartEndTimes() "

	t2Dur := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setStartEndTimes(
		&t2Dur,
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DurationTriad{}, err
	}

	return t2Dur, nil
}

// NewStartEndTimesTz - Returns a new DurationTriad based on two input
// parameters, 'startDateTime' and 'endDateTime'. These two input parameters
// are submitted as instances of type 'DateTzDto'.
//
// Time Zone Location is extracted from input parameter, 'startDateTime'.
//
// Date Time Format string is likewise extracted from input parameter,
// 'startDateTime'.
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
//  startDateTime  DateTzDto
//     - Starting date time
//
//
//  endDateTime    DateTzDto
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
//       TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                          day time calculation.
//
//       TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                          or hours.
//
//       TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                          hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//                                        - Computes Years based on average length of a Gregorian Year
//                                          Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//        MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr    string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// __________________________________________________________________________
//
// Return Values:
//
//  DurationTriad
//     - Upon successful completion, this method will return
//       a new, populated DurationTriad instance.
//
//       A DurationTriad Structure is defined as follows:
//
//         type DurationTriad struct {
//           BaseTime  TimeDurationDto
//           LocalTime TimeDurationDto
//           UTCTime   TimeDurationDto
//         }
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
//
// __________________________________________________________________________
//
// Example Usage:
//
//
//  du, err := DurationTriad{}.NewStartEndTimesTz(
//                             startDateTimeTz,
//                             endTimeDateTz,
//                             TDurCalc.StdYearMth(),
//                             TZones.US.Central(),
//                             TCalcMode.LocalTimeZone(),
//                             FmtDateTimeYrMDayFmtStr)
//
// Note:
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT DurationTriad) NewStartEndTimesTz(
	startDateTimeTz,
	endDateTimeTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewStartEndTimesTz() "

	durT2 := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setStartEndTimes(
		&durT2,
		startDateTimeTz.dateTimeValue,
		endDateTimeTz.dateTimeValue,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DurationTriad{}, err
	}

	return durT2, nil
}

// NewStartTimeDuration - Returns a NewStartEndTimes DurationTriad based on 'startDateTime'
// and 'duration' input parameters. Time 'duration' is added to 'startDateTime'
// in order to compute ending date time.
//
// This method will accept negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
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
//
//  startDateTime time.Time
//     - Starting Date Time for duration calculation
//
//
//  duration  time.Duration
//     - Time Duration added to 'startDatTime' in order to
//       compute Ending Date Time
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
//  DurationTriad
//     - Upon successful completion, this method will return
//       a new, populated DurationTriad instance.
//
//       A DurationTriad Structure is defined as follows:
//
//         type DurationTriad struct {
//           BaseTime  TimeDurationDto
//           LocalTime TimeDurationDto
//           UTCTime   TimeDurationDto
//         }
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//
//  tDurDto, err := TimeDurationDto{}.NewStartTimeDuration(
//                    startTime,
//                    duration,
//                    TDurCalc.StdYearMth(),
//                    TZones.US.Central(),
//                    TCalcMode.LocalTimeZone(),
//                    FmtDateTimeYrMDayFmtStr)
//
// Note:
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT DurationTriad) NewStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewStartTimeDuration() "

	durT2 := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setStartTimeDuration(
		&durT2,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DurationTriad{}, err
	}

	return durT2, nil
}

// NewStartTimeTzDuration - Returns a NewStartEndTimes DurationTriad based on 'startDateTz'
// and 'duration' input parameters. Time 'duration' is added to 'startDateTz'
// in order to compute ending date time.
//
// Input parameter 'startDateTz' is of Type, 'DateTzDto'.
//
// This method will except negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
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
//  startDateTz       DateTzDto
//     - Provides starting date time for duration calculation
//
//
//  duration          time.Duration
//     - Time Duration added to 'startDatTime' in order to
//       compute Ending Date Time
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
//  DurationTriad
//     - Upon successful completion, this method will return
//       a new, populated DurationTriad instance.
//
//       A DurationTriad Structure is defined as follows:
//
//         type DurationTriad struct {
//           BaseTime  TimeDurationDto
//           LocalTime TimeDurationDto
//           UTCTime   TimeDurationDto
//         }
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
//
// __________________________________________________________________________
//
// Example Usage:
//
//  tDurDto, err := TimeDurationDto{}.NewStartTimeTzDuration(
//                    startDateTimeTz,
//                    duration,
//                    TDurCalc.StdYearMth(),
//                    TZones.US.Central(),
//                    TCalcMode.LocalTimeZone(),
//                    FmtDateTimeYrMDayFmtStr)
//
// Note:
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT DurationTriad) NewStartTimeTzDuration(
	startDateTimeTz DateTzDto,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewStartTimeTzDuration() "

	durT2 := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setStartTimeDuration(
		&durT2,
		startDateTimeTz.dateTimeValue,
		duration,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return DurationTriad{}, err
	}

	return durT2, nil
}

// NewStartTimePlusTimeDto - Returns a new DurationTriad based on two input
// parameters, 'startDateTime' and 'timeDto'. 'timeDto' is an instance of TimeDto
// which is added to 'startDateTime' in order to calculate ending date time as
// well as time duration.
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
//  startDateTime  time.Time
//     - Starting date time. The TimeDto parameter will be added
//       to this date time in order to compute the ending date time
//       and the time duration.
//
//
//  plusTimeDto      TimeDto
//     - Provides time values which will be added to
//      'startDateTime' in order to calculate time duration
//      and ending date time.
//
//      type TimeDto struct {
//         Years                int // Number of Years
//         Months               int // Number of Months
//         Weeks                int // Number of Weeks
//         WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//         DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//         Hours                int // Number of Hours.
//         Minutes              int // Number of Minutes
//         Seconds              int // Number of Seconds
//         Milliseconds         int // Number of Milliseconds
//         Microseconds         int // Number of Microseconds
//         Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//         TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                  //  plus remaining Nanoseconds
//         TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                  //  + Seconds + Milliseconds + Nanoseconds
//      }
//
//      Type 'TimeDto' is located in source file:
//           datetimeopsgo\datetime\timedto.go
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
//  DurationTriad
//     - Upon successful completion, this method will return
//       a new, populated DurationTriad instance.
//
//       A DurationTriad Structure is defined as follows:
//
//         type DurationTriad struct {
//           BaseTime  TimeDurationDto
//           LocalTime TimeDurationDto
//           UTCTime   TimeDurationDto
//         }
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//
//  du, err := DurationTriad{}.NewStartTimePlusTimeDto(
//                    startDateTime,
//                    plusTimeDto,
//                    TDurCalc.StdYearMth(),
//                    TZones.US.Central(),
//                    TCalcMode.LocalTimeZone(),
//                    FmtDateTimeYrMDayFmtStr)
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
func (durT DurationTriad) NewStartTimePlusTimeDto(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewStartTimePlusTimeDto() "

	durT2 := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setStartTimePlusTimeDto(
		&durT2,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	return durT2, err
}

// NewStartTimeTzPlusTimeDto - Returns a new DurationTriad based on two input parameters,
// 'startDateTimeTz' and 'timeDto'.
//
// 'timeDto' is an instance of Type TimeDto which is added to 'startDateTimeTz' in order
// to calculate ending date time as well as time duration. Type 'TimeDto' stores time
// values by granular components such as years, months, days, hours, minutes, seconds,
// milliseconds, microseconds and nanoseconds.
//
// 'startDateTimeTz' is passed as an instance of 'DateTzDto' and marks
// the starting date time for the time duration calculation.
//
// The user is required to provide Time Zone Location as an input parameter in order
// to ensure that time duration calculations are performed using equivalent time zones.
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  startDateTimeTz   DateTzDto
//     - Starting date time encapsulated by an instance of type
//       'DateTzDto'.
//
//
//  plusTimeDto       TimeDto
//     - Provides time values which will be added to
//       'startDateTime' in order to calculate duration.
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
//     Type 'TimeDto' is located in source file:
//     datetimeopsgo\datetime\timedto.go
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  DurationTriad - Upon successful completion, this method will return
//                  a new, populated DurationTriad instance.
//
//                  A DurationTriad Structure is defined as follows:
//
//                  type DurationTriad struct {
//                    BaseTime  TimeDurationDto
//                    LocalTime TimeDurationDto
//                    UTCTime   TimeDurationDto
//                  }
//
//
//  error         - If this method completes successfully, the returned error
//                  Type is set equal to 'nil'. If an error condition is encountered,
//                  this method will return an error Type which encapsulates an
//                  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  du, err := DurationTriad{}.NewStartTimeTzPlusTimeDto(
//                    startDateTimeTz,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT DurationTriad) NewStartTimeTzPlusTimeDto(
	startDateTimeTz DateTzDto,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) (DurationTriad, error) {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.NewStartTimeTzPlusTimeDto() "

	durT2 := DurationTriad{}

	durTUtil := durationTriadUtility{}

	err := durTUtil.setStartTimePlusTimeDto(
		&durT2,
		startDateTimeTz.dateTimeValue,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)

	return durT2, err
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
// Input Parameters
//
//  None
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  error    - If this method completes successfully, the returned error
//             Type is set equal to 'nil'. If an error condition is encountered,
//             this method will return an error Type which encapsulates an
//             appropriate error message.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//
//  dt, err := DurationTriad{}.NewAutoStart(
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
func (durT *DurationTriad) SetAutoEnd() error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.SetAutoEnd() "

	calcType := durT.BaseTime.timeDurCalcType
	startDateTime := durT.BaseTime.startDateTimeTz.GetDateTimeValue()

	tzLocName := durT.BaseTime.startDateTimeTz.GetOriginalTzName()
	fmtStr := durT.BaseTime.startDateTimeTz.GetDateTimeFmt()
	timeMathCalcMode := durT.BaseTime.timeMathCalcMode

	durTUtil := durationTriadUtility{}

	return durTUtil.setStartEndTimes(
		durT,
		startDateTime,
		time.Now().UTC(),
		calcType,
		tzLocName,
		timeMathCalcMode,
		fmtStr,
		ePrefix)
}

// SetEndTimeMinusTimeDto - Calculates duration values based on an Ending Date Time and
// a TimeDto structure consisting of time values (Years, Months, weeks, days, hours,
// minutes etc.). The time values in the 'timeDto' parameter are subtracted
// from 'endDateTime'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time. As a result, true values
// for 'startDateTimeTz', 'endDateTimeTz' and 'timeDuration' are stored in
// the DurationTriad data structure.
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  endDateTime    time.Time
//     - The ending date time value from which TimeDto
//       parameter 'minusTimeDto' will be subtracted
//       in order to compute the Starting Date Time.
//
//
//  minusTimeDto   TimeDto
//     - An instance of TimeDto containing time values,
//       (Years, Months, weeks, days, hours, minutes etc.)
//       which will be subtracted from input parameter
//       'endDateTime' in order to compute the Starting
//       Date Time and Time Duration.
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  error    - If this method completes successfully, the returned error
//             Type is set equal to 'nil'. If an error condition is encountered,
//             this method will return an error Type which encapsulates an
//             appropriate error message.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//
//  dt, err := DurationTriad{}.NewAutoStart(
//                    TDurCalc.StdYearMth(),
//                    TZones.US.Central(),
//                    TCalcMode.LocalTimeZone(),
//                    FmtDateTimeYrMDayFmtStr)
//
//  err := dt.SetEndTimeMinusTimeDto(
//              endDateTime,
//              minusTimeDto,
//                    TDurCalc.StdYearMth(),
//                    TZones.US.Central(),
//                    TCalcMode.LocalTimeZone(),
//                    FmtDateTimeYrMDayFmtStr)
//
//
// Note:
//
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT *DurationTriad) SetEndTimeMinusTimeDto(
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.SetEndTimeMinusTimeDto() "

	durTUtil := durationTriadUtility{}

	return durTUtil.setEndTimeMinusTimeDto(
		durT,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetEndTimeTzMinusTimeDto - Calculates duration values based on an Ending Date Time
// and a TimeDto structure consisting of time values (Years, Months, weeks, days,
// hours, minutes etc.). The time values in the 'timeDto' parameter are subtracted
// from the Ending Date Time.  Ending Date Time is formatted as an instance of
// 'DateTzDto'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time.
//
// As a result, true values for starting date time, ending date time and time duration
// are stored in the DurationTriad data structure.
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  endDateTimeTz   DateTzDto
//     - This instance of type 'DateTzDto' encapsulates
//       the ending date time value from which TimeDto
//       parameter 'minusTimeDto' will be subtracted
//       in order to compute the Starting Date Time value.
//
//  minusTimeDto    TimeDto
//     - An instance of TimeDto containing time values,
//       (Years, Months, weeks, days, hours, minutes etc.)
//       which will be subtracted from input parameter
//       'endDateTime' in order to compute the Starting
//       Date Time and Time Duration.
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
// ------------------------------------------------------------------------
//
// Return Value
//
//  error   - If this method completes successfully, the returned error
//            Type is set equal to 'nil'. If an error condition is encountered,
//            this method will return an error Type which encapsulates an
//            appropriate error message.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
// dt  := DurationTriad{}
// err := dt.SetEndTimeTzMinusTimeDto(
//                      endDateTimeTz,
//                      minusTimeDto,
//                      TDurCalc.StdYearMth(),
//                      TZones.US.Central(),
//                      TCalcMode.LocalTimeZone(),
//                      FmtDateTimeYrMDayFmtStr)
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT *DurationTriad) SetEndTimeTzMinusTimeDto(
	endDateTimeTz DateTzDto,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.SetEndTimeTzMinusTimeDto() "

	durTUtil := durationTriadUtility{}

	return durTUtil.setEndTimeMinusTimeDto(
		durT,
		endDateTimeTz.dateTimeValue,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartEndTimes - Calculates duration values and save the results in the current DurationTriad
// data fields. Calculations are based on a starting date time and an ending date time values passed
// as input parameters.
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
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  startDateTime    time.Time
//     - Starting date time
//
//
//  endDateTime      time.Time
//     - Ending date time
//
//
//  tDurCalcType TDurCalcType
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
// ------------------------------------------------------------------------
//
// Return Value
//
//  error    - If this method completes successfully, the returned error
//        Type is set equal to 'nil'. If an error condition is encountered,
//        this method will return an error Type which encapsulates an
//        appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dt  := DurationTriad{}
//  err := dt.SetStartEndTimes(
//                    startDateTime,
//                    endDateTime,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT *DurationTriad) SetStartEndTimes(
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.SetStartEndTimes() "

	durTUtil := durationTriadUtility{}

	return durTUtil.setStartEndTimes(
		durT,
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartEndTimesTz - Calculates duration values and saves the results
// in the current DurationTriad data fields. Calculations are based on a starting
// date time and an ending date time passed to the method as input parameters of
// Type 'DateTzDto'.
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  startDateTimeTz   DateTzDto
//     - Starting date time
//
//
//  endDateTimeTz     DateTzDto
//     - Ending date time
//
//
//  tDurCalcType TDurCalcType
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
// ------------------------------------------------------------------------
//
// Return Value
//
//  error     - If this method completes successfully, the returned error
//              Type is set equal to 'nil'. If an error condition is encountered,
//              this method will return an error Type which encapsulates an
//              appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dt  := DurationTriad{}
//  err := dt.SetStartEndTimesTz(
//                    startDateTimeTz,
//                    endDateTimeTz,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT *DurationTriad) SetStartEndTimesTz(
	startDateTimeTz,
	endDateTimeTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.SetStartEndTimesTz() "

	durTUtil := durationTriadUtility{}

	return durTUtil.setStartEndTimes(
		durT,
		startDateTimeTz.GetDateTimeValue(),
		endDateTimeTz.GetDateTimeValue(),
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimeDuration - Receives a starting date time and proceeds to calculate
// the ending date time and time duration. These results are then saved to the
// current DurationTriad data fields.
//
// This method will except negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
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
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  startDateTime   time.Time
//     - Starting Date Time for duration calculation
//
//
//  duration        time.Duration
//     - Time Duration added to 'startDateTime' in order to
//       compute Ending Date Time.
//
//
//  tDurCalcType TDurCalcType
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
// ------------------------------------------------------------------------
//
// Return Value
//
//  error    - If this method completes successfully, the returned error
//             Type is set equal to 'nil'. If an error condition is encountered,
//             this method will return an error Type which encapsulates an
//             appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dt  := DurationTriad{}
//  err := dt.SetStartTimeDuration(
//                    startDateTime,
//                    duration,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT *DurationTriad) SetStartTimeDuration(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.SetStartTimeDuration() "

	if startDateTime.IsZero() {
		return errors.New(ePrefix + "Error - Start Time is Zero!")
	}

	durTUtil := durationTriadUtility{}

	return durTUtil.setStartTimeDuration(
		durT,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimeTzDuration - Receives a starting date time and then proceeds
// to calculate the ending date time and time duration. These results are then
// saved to the DurationTriad data fields.  The starting date time is formatted
// as an instance of type 'DateTzDto'.
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
//  startDateTimeTz   DateTzDto
//     - Starting Date Time used in the calculation of ending date time
//       and time duration.
//
//  duration          time.Duration
//     - Time Duration added to 'startDateTimeTz' in order to compute Ending
//       Date Time. If duration is a negative value, 'startDateTimeTz' will be
//       treated as an ending date time.
//
//
//  tDurCalcType TDurCalcType
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
//
// __________________________________________________________________________
//
// Example Usage:
//
//
//  dt  := DurationTriad{}
//  err := dt.SetStartTimeTzDuration(
//                    startDateTimeTz,
//                    duration,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT *DurationTriad) SetStartTimeTzDuration(
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.SetStartTimeTzDuration() "

	durTUtil := durationTriadUtility{}

	return durTUtil.setStartTimeDuration(
		durT,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
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
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  startDateTime   time.Time
//     - Starting time
//
//
//  plusTimeDto     TimeDto
//     - Provides time values which will be subtracted from
//       'startDateTime' in order to calculate duration.
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
// ------------------------------------------------------------------------
//
// Return  Value
//
//  error    - If this method completes successfully, the returned error
//        Type is set equal to 'nil'. If an error condition is encountered,
//        this method will return an error Type which encapsulates an
//        appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  dt  := DurationTriad{}
//  err := dt.SetStartTimePlusTimeDto(
//                startDateTime,
//                plusTimeDto,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT *DurationTriad) SetStartTimePlusTimeDto(
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.SetStartTimePlusTimeDto() "

	durTUtil := durationTriadUtility{}

	return durTUtil.setStartTimePlusTimeDto(
		durT,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}

// SetStartTimeTzPlusTimeDto - Calculates time duration values based on a Starting Date Time
// plus time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'plusTimeDto' parameter. The 'plusTimeDto' parameter is added to 'startDateTime' to
// calculate ending date time and duration.
//
// Starting date time is formatted as an instance of type 'DateTzDto'.
//
// Values in the 'plusTimeDto' parameter are automatically converted to positive numeric
// values before being added to parameter 'startDateTime'.
//
// True values for starting date time, ending date time and time duration are then stored in
// the DurationTriad data structure.
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  startDateTimeTz   DateTzDto
//     - Starting date time. Input parameter 'plusTimeDto'
//       will be added to this starting date time value in
//       order to generate ending date time and time duration.
//
//  plusTimeDto       TimeDto
//     - Provides time values which will be added to
//       'startDateTime' in order to calculate duration.
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
//    Type 'TimeDto' is located in source file:
//       datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType TDurCalcType
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
//  timeZoneLocation string
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
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
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
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
// __________________________________________________________________________
//
// Example Usage:
//
//  dt  := DurationTriad{}
//  err := dt.SetStartTimeTzPlusTimeDto(
//         startDateTimeTz,
//         plusTimeDto,
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
//            'datetime\timemathcalcmode.go'
//
//        'FmtDateTimeYrMDayFmtStr' is a constant available in source file,
//        'constantsdatetime.go'.
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (durT *DurationTriad) SetStartTimeTzPlusTimeDto(
	startDateTimeTz DateTzDto,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string) error {

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.lock.Lock()

	defer durT.lock.Unlock()

	ePrefix := "DurationTriad.SetStartTimeTzPlusTimeDto() "

	durTUtil := durationTriadUtility{}

	return durTUtil.setStartTimePlusTimeDto(
		durT,
		startDateTimeTz.dateTimeValue,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix)
}
