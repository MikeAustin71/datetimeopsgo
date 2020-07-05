package datetime

import (
	"fmt"
	"sync"
	"time"
)


// TimeZoneDto
// ===========
//
// TimeZoneDto is part of the date time operation's library. The source code repository
// for this file is located at:
//
//     https://github.com/MikeAustin71/datetimeopsgo.git
//
// This source code file is located at:
//
//     MikeAustin71\datetimeopsgo\datetime\timezonedto.go
//
//
// Overview and General Usage
// ==========================
//
// TimeZoneDto is used to convert, store and transport time zone information.
// The user will use this Type to convert 'time.Time', date time values, between
// differing time zones.
//
// In addition to generating a date time converted to a time zone specified
// by the user, this Type automatically generates equivalent date time values
// for Time Zone Locations 'Local' and 'UTC'.
//
// If you are unfamiliar with the concept of a Time Zone Location, reference
// 'https://golang.org/pkg/time/'. The concept of Time Zone Location is used
// extensively by Type TimeZoneDto. Time Zone location must be designated as
// one of three types of values.
//
//   (1) The string 'Local' - signals the designation of the local time zone
//       configured for the host computer executing this code.
//
//   (2) IANA Time Zone Location -
//       See https://golang.org/pkg/time/#LoadLocation
//       and https://www.iana.org/time-zones to ensure that
//       the IANA Time Zone Database is properly configured
//       on your system. Note: IANA Time Zone Data base is
//       equivalent to 'tz database'.
//
//          Examples:
//            "America/New_York"
//            "America/Chicago"
//            "America/Denver"
//            "America/Los_Angeles"
//            "Pacific/Honolulu"
//
//   (3) A valid Military Time Zone
//       Military time zones are commonly used in
//       aviation as well as at sea. They are also
//       known as nautical or maritime time zones.
//       Reference:
//           https://en.wikipedia.org/wiki/List_of_military_time_zones
//           http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//           https://www.timeanddate.com/time/zones/military
//           https://www.timeanddate.com/worldclock/timezone/alpha
//           https://www.timeanddate.com/time/map/
//
//   Note:
//       The source file 'timezonedata.go' contains over 600 constant
//       time zone declarations covering all IANA and Military Time
//       Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//       time zone constants begin with the prefix 'TZones'.
//
// Dependencies
// ============
// 
//   DateTzDto    - datetzdto.go
//   TimeZoneDef  - timezonedefinition.go
//
//
// TimeZoneDto - Time Zone Data Transfer Object Type and Methods
// =============================================================
//
type TimeZoneDto struct {
	Description string       // Unused - available for tagging, classification or
	                         //   labeling.
	TimeIn      DateTzDto    // Original input time value
	TimeOut     DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
	TimeUTC     DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
	                         //   equivalent to TimeIn
	TimeLocal   DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
	                         //   'Local' is the Time Zone Location used by the host computer.
	DateTimeFmt string       // Date Time Format String. This format string is used to format
	                         //   Date Time text displays. The Default format string is:
	                         //   "2006-01-02 15:04:05.000000000 -0700 MST"
	lock         *sync.Mutex // Mutex used to manage thread safe operations
}

// AddDate - Adds specified years, months and days values to the
// current time values maintained by this TimeZoneDto
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// years    int    - Number of years to add to current TimeZoneDto instance
// months   int    - Number of months to add to current TimeZoneDto instance
// days     int    - Number of months to add to current TimeZoneDto instance
//
// Note:  The date input parameter may be either negative
//        or positive. Negative values will subtract time
//        from the current TimeZoneDto instance.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// There only one return: An 'error' type.
//
// error - If errors are encountered, this method returns an error object.
//         Otherwise, the error value is 'nil'.
//
func (tzDto *TimeZoneDto) AddDate(years, months, days int) error {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.AddDate() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.addDateTime(
		tzDto,
		years,
		months,
		days,
		0,        // hours
		0,        // minutes
		0,        //  seconds
		0,        // milliseconds
		0,        // microseconds
		0,
		ePrefix)
}

// AddDateTime - Adds input time elements to the time
// value of the current TimeZoneDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// years        int  - Number of years added to current TimeZoneDto
// months       int  - Number of months added to current TimeZoneDto
// days         int  - Number of days added to current TimeZoneDto
// hours        int  - Number of hours added to current TimeZoneDto
// minutes      int  - Number of minutes added to current TimeZoneDto
// seconds      int  - Number of seconds added to current TimeZoneDto
// milliseconds int  - Number of milliseconds added to current TimeZoneDto
// microseconds int  - Number of microseconds added to current TimeZoneDto
// subMicrosecondNanoseconds  int  - Number of subMicrosecondNanoseconds added to current TimeZoneDto
//
// Note:  Date Time input parameters may be either negative or positive.
//        Negative values will subtract time from the current TimeZoneDto
//        instance.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddDateTime(
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) error {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.addDateTime() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.addDateTime(
		tzDto,
		years,
		months,
		days,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds,
		ePrefix)
}

// AddDuration - Adds 'duration' to the time values maintained by the
// current TimeZoneDto.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// duration  time.Duration  - May be a positive or negative duration
//                            value which is added to the time value
//                            of the current TimeZoneDto instance.
//
// Note:   The time.Duration input parameter may be either negative
//         or positive. Negative values will subtract time from
//         the current TimeZoneDt instance.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddDuration(duration time.Duration) error {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.AddDuration() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.addDuration(tzDto, duration, ePrefix)
}

// AddMinusTimeDto - This method receives a TimeDto input parameter. It
// then proceeds to convert all time components to negative values and
// subtracts those time components from the time values of the current
// TimeZoneDto.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// timeDto TimeDto - A TimeDto type containing time components (i.e.
//          years, months, weeks, days, hours, minutes,
//          seconds etc.) to be subtracted from the current
//          TimeZoneDto.
//
//          type TimeDto struct {
//             Years                int // Number of Years
//             Months               int // Number of Months
//             Weeks                int // Number of Weeks
//             WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//             DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//             Hours                int // Number of Hours.
//             Minutes              int // Number of Minutes
//             Seconds              int // Number of Seconds
//             Milliseconds         int // Number of Milliseconds
//             Microseconds         int // Number of Microseconds
//             Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//             TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                      //  plus remaining Nanoseconds
//             TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                      //  + Seconds + Milliseconds + Nanoseconds
//          }
//
//          Type 'TimeDto' is located in source file:
//             datetimeopsgo\datetime\timedto.go
//
// ------------------------------------------------------------------------
//
// Return Values
//
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddMinusTimeDto(timeDto TimeDto) error {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.AddMinusTimeDto() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.addMinusTimeDto(
		tzDto,
		TCalcMode.LocalTimeZone(),
		timeDto,
		ePrefix)
}

// AddPlusTimeDto - This method receives a TimeDto input parameter. It
// then proceeds to convert all time components to positive values and
// adds those time components to the time values of the current TimeZoneDto.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// timeDto  TimeDto - A TimeDto type containing time components (i.e.
//          years, months, weeks, days, hours, minutes,
//          seconds etc.) to be added to the current
//          TimeZoneDto.
//
//          type TimeDto struct {
//             Years                int // Number of Years
//             Months               int // Number of Months
//             Weeks                int // Number of Weeks
//             WeekDays             int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//             DateDays             int // Total Number of Days. Weeks x 7 plus WeekDays
//             Hours                int // Number of Hours.
//             Minutes              int // Number of Minutes
//             Seconds              int // Number of Seconds
//             Milliseconds         int // Number of Milliseconds
//             Microseconds         int // Number of Microseconds
//             Nanoseconds          int // Remaining Nanoseconds after Milliseconds & Microseconds
//             TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                      //  plus remaining Nanoseconds
//             TotTimeNanoseconds int64 // Total Number of equivalent Nanoseconds for Hours + Minutes
//                                      //  + Seconds + Milliseconds + Nanoseconds
//          }
//
//          Type 'TimeDto' is located in source file:
//             datetimeopsgo\datetime\timedto.go
//
// ------------------------------------------------------------------------
//
// Return Values
//
// There is only one return: an 'error' type.
//
// error -  If errors are encountered, this method returns an 'error'
//          instance populated with an error message. If the method completes
//          successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddPlusTimeDto(timeDto TimeDto) error {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.AddPlusTimeDto() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.addPlusTimeDto(
		tzDto,
		TCalcMode.LocalTimeZone(),
		timeDto,
		ePrefix)
}

// AddTime - Adds time elements to the time value of the current
// TimeZoneDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  hours        - hours to be added to current TimeZoneDto
//
//  minutes      - minutes to be added to current TimeZoneDto
//
//  seconds      - seconds to be added to current TimeZoneDto
//
//  milliseconds - milliseconds to be added to current TimeZoneDto
//
//  microseconds - microseconds to be added to current TimeZoneDto
//
//  subMicrosecondNanoseconds  - subMicrosecondNanoseconds to be added to current TimeZoneDto
//
//  Note:  The time component input parameter may be either negative
//         or positive. Negative values will subtract time from
//         the current TimeZoneDto instance.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  There is only one return: an 'error' type.
//
//  error -  If errors are encountered, this method returns an 'error'
//           instance populated with an error message. If the method completes
//           successfully, this error value is set to 'nil'.
//
func (tzDto *TimeZoneDto) AddTime(
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int) error {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.AddTime() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.addTime(
		tzDto,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds,
		ePrefix)
}

// AddTimeDurationDto - Adds time duration as expressed by input type 'TimeDurationDto'
// to the time values maintained by the current TimeZoneDto.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  durDto  TimeDurationDto  - Contains the time duration value
//                             to be added to the current TimeZoneDto.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  There is only one return: an 'error' type.
//
//  error -  If errors are encountered, this method returns an 'error'
//           instance populated with an error message. If the method completes
//           successfully, this error value is set to 'nil'
//
func (tzDto *TimeZoneDto) AddTimeDurationDto(durDto TimeDurationDto) error {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.AddTimeDurationDto() "

	tZoneUtil :=timeZoneDtoUtility{}

	return tZoneUtil.addTimeDurationDto(
		tzDto,
		durDto,
		ePrefix)
}

// ConvertTz - Converts 'tIn' Date Time from existing time zone to a 'targetTz'
// or target Time Zone. The results are stored and returned in a new
// TimeZoneDto data structure. The current TimeZoneDto is NOT changed.
//
// The input time and output time are equivalent times adjusted for
// different time zones.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  tIn     time.Time
//     - Initial time values
//
//  targetTz   string
//     - Time zone location must be designated as one of three
//       types of time zones:
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
// ------------------------------------------------------------------------
//
// Return Values
//
// There are two returns:
//             (1) A TimeZoneDto instance
//             (2) An error type
//
// (1) TimeZoneDto
//     If successful, this method creates a new TimeZoneDto,
//     populated with, TimeIn, TimeOut, TimeUTC and TimeLocal
//     date time values plus time zone information.
//
//     A TimeZoneDto structure is defined as follows:
//
//     type TimeZoneDto struct {
//      Description  string     // Unused - available for tagging, classification or
//                              //   labeling.
//      TimeIn       DateTzDto  // Original input time value
//      TimeOut      DateTzDto  // TimeOut - 'TimeIn' value converted to TimeOut
//      TimeUTC      DateTzDto  // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                              //   equivalent to TimeIn
//      TimeLocal    DateTzDto  // TimeIn value converted to the 'Local' Time Zone Location.
//                              //   'Local' is the Time Zone Location used by the host computer.
//      DateTimeFmt  string     // Date Time Format String. This format string is used to format
//                              //  Date Time text displays. The Default format string is:
//                              //   "2006-01-02 15:04:05.000000000 -0700 MST"
//     }
//
//
// (2) error - If errors are encountered, this method returns an error instance populated with
//             a valid 'error' message. If the method completes successfully the returned error
//             error type is set to 'nil'.
//
func (tzDto TimeZoneDto) ConvertTz(
	tIn time.Time,
	targetTimeZoneName,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.ConvertTz() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.convertTz(&tzDto, tIn, targetTimeZoneName, dateTimeFmtStr, ePrefix)
}

// CopyIn - Copies input parameter TimeZoneDto data fields
// into the current TimeZoneDto data fields.
// When the method completes, the current TimeZoneDto and
// the input parameter TimeZoneDto are equivalent.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  tzdto2 TimeZoneDto
//     - A TimeZoneDto instance. The data
//       fields from this incoming TimeZoneDto
//       will be copied to the data fields
//       of the current TimeZoneDto.
//
//       A TimeZoneDto structure is defined as follows:
//
//       type TimeZoneDto struct {
//        Description  string       // Unused - available for tagging, classification or
//                                  //  labeling.
//        TimeIn       DateTzDto    // Original input time value
//        TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//        TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                                  //   equivalent to TimeIn
//        TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                                  //   'Local' is the Time Zone Location used by the host computer.
//        DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                                  //  Date Time text displays. The Default format string is:
//                                  //   "2006-01-02 15:04:05.000000000 -0700 MST"
//       }
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  None
//
func (tzDto *TimeZoneDto) CopyIn(tzDto2 TimeZoneDto) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	tZoneUtil := timeZoneDtoUtility{}

	tZoneUtil.copyIn(tzDto, &tzDto2, "TimeZoneDto.CopyIn() ")

	return
}

// CopyOut - Creates and returns a deep copy of the
// current TimeZoneDto instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  None
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  There is only one return: A TimeZoneDto instance.
//
//       A TimeZoneDto structure is defined as follows:
//
//       type TimeZoneDto struct {
//        Description  string       // Unused - available for tagging, classification or
//                                  //  labeling.
//        TimeIn       DateTzDto    // Original input time value
//        TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//        TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                                  //   equivalent to TimeIn
//        TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                                  //   'Local' is the Time Zone Location used by the host computer.
//        DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                                  //  Date Time text displays. The Default format string is:
//                                  //   "2006-01-02 15:04:05.000000000 -0700 MST"
//       }
//
func (tzDto *TimeZoneDto) CopyOut() TimeZoneDto {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.CopyOut() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.copyOut(tzDto, ePrefix)
}

// Equal - returns a boolean value indicating
// whether the current TimeZoneDto data structure
// is equivalent to the input parameter TimeZoneDto
// data structure.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  tzdto2  TimeZoneDto
//     - This input parameter TimeZoneDto
//       is compared to the current TimeZoneDto
//       to determine if they are equivalent.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  bool  - If the current TimeZoneDto is equivalent to the
//          input parameter TimeZoneDto in all respects,
//         this method returns 'true'.
//
//          If the two TimeZoneDto's are NOT equivalent, this
//          method returns 'false'
//
func (tzDto *TimeZoneDto) Equal(tzdto2 TimeZoneDto) bool {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()
	tZoneUtil := timeZoneDtoUtility{}

	ePrefix := "TimeZoneDto.Equal() "

	return tZoneUtil.equal(tzDto, &tzdto2, ePrefix)
}

// Empty - Clears or returns the current
// TimeZoneDto to an uninitialized or
// 'Empty' state.
//
func (tzDto *TimeZoneDto) Empty() {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	tZoneUtil := timeZoneDtoUtility{}

	tZoneUtil.empty(tzDto, "TimeZoneDto.Empty() ")

	return
}

// IsValid - Analyzes the current TimeZoneDto
// instance and returns an error if the instance is INVALID.
//
func (tzDto *TimeZoneDto) IsValid() error {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	ePrefix := "TimeZoneDto.IsValid() "

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.isValidTimeZoneDto(tzDto, ePrefix)
}

// IsValidTimeZone - Tests a Time Zone Location string and
// returns two values:
//
// Input Parameters
// ============================================================================
//
// tZone   string  - The name of a valid time zone. This time zone must
//                   specify one of three types of time zones.
//
//                   (1) The 'Local' Time Zone.  Time Zone 'Local' specifies
//                       the time zone configured and applied on the host
//                       computer.
//
//                   (2) A IANA Time zone. This time must exist in the IANA
//                       database. The IANA database is widely recognized as
//                       a leading authority of global time zones.
//
//                       Reference:
//
//                          https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//                          https://en.wikipedia.org/wiki/Tz_database
//                          https://www.iana.org/time-zones
//
//                   (3) A single character or the full text name of a valid Military
//                       time zone. Military time zones are commonly used in aviation
//                       as well as at sea. They are also known as nautical or maritime
//                       time zones.
//                       Reference:
//                           https://en.wikipedia.org/wiki/List_of_military_time_zones
//                           http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                           https://www.timeanddate.com/time/zones/military
//
// Return Values
// ============================================================================
//
// isValidTz            bool - If true it signals the the time zone name contained in
//                             input parameter 'tZone' is a valid time zone.
//
// timeZoneType TimeZoneType - If return value 'isValidTz' is 'true', the 'timeZoneType'
//                             value will describe the valid time zone as one of three
//                             types: 'Local', 'IANA' or Military.
//
func (tzDto *TimeZoneDto) IsValidTimeZone(
	timeZoneName string) (
	isValidTimeZone bool,
	timeZoneType TimeZoneType,
	err error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.IsValidTimeZone() "

	tZoneUtil := timeZoneDtoUtility{}

	isValidTimeZone,
	timeZoneType,
	err =
	tZoneUtil.isValidTimeZoneName(
		timeZoneName,
		ePrefix)

	return isValidTimeZone, timeZoneType, err
}

// NewStartEndTimes - Initializes and returns a new TimeZoneDto object.
//
// Input Parameters
// ----------------
//
// tIn      time.Time   - The input time object.
//
// tZoneOutLocationName - string -
//                        The first input time value, 'tIn' will have its time zone
//                        changed to a new time zone location specified by this
//                        second parameter, 'tZoneOutLocationName'. The new time
//                        associated with 'tZoneOutLocationName' is assigned to
//                        the TimeZoneDto data field. The 'tZoneOutLocationName'
//                        time zone location must be designated as one of three
//                        types of time zones:
//
//                        (1) the string 'Local' - signals the designation of the
//                            time zone location used by the host computer.
//
//                        (2) IANA Time Zone Location -
//                           See https://golang.org/pkg/time/#LoadLocation
//                           and https://www.iana.org/time-zones to ensure that
//                           the IANA Time Zone Database is properly configured
//                           on your system. Note: IANA Time Zone Data base is
//                           equivalent to 'tz database'.
//                             Examples:
//                              "America/New_York"
//                              "America/Chicago"
//                              "America/Denver"
//                              "America/Los_Angeles"
//                              "Pacific/Honolulu"
//
//                        (3) A Military Time Zone
//                            Reference:
//                             https://en.wikipedia.org/wiki/List_of_military_time_zones
//                             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                             https://www.timeanddate.com/time/zones/military
//                             https://www.timeanddate.com/worldclock/timezone/alpha
//                             https://www.timeanddate.com/time/map/
//
//                            Examples:
//                              "Alpha"   or "A"
//                              "Bravo"   or "B"
//                              "Charlie" or "C"
//                              "Delta"   or "D"
//                              "Zulu"    or "Z"
//
// dateTimeFmtStr string  - A date time format string which will be used
//                to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                If 'dateTimeFmtStr' is submitted as an
//                'empty string', a default date time format
//                string will be applied. The default date time
//                format string is:
//                TZDtoDefaultDateTimeFormatStr =
//                "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Return Values
// =============
//
//  There are two return values:  (1) A TimeZoneDto Type
//                                (2) An Error type
//
//  (1) TimeZoneDto
//         - The two input parameters are used to populate and return
//           a TimeZoneDto structure:
//
//           type TimeZoneDto struct {
//            Description  string     // Unused - available for tagging, classification or
//                                    //  labeling.
//            TimeIn       DateTzDto  // Original input time value
//            TimeOut      DateTzDto  // TimeOut - 'TimeIn' value converted to TimeOut
//            TimeUTC      DateTzDto  // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                                    //   equivalent to TimeIn
//            TimeLocal   DateTzDto   // TimeIn value converted to the 'Local' Time Zone Location.
//                                    //   'Local' is the Time Zone Location used by the host computer.
//            DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                                    //  Date Time text displays. The Default format string is:
//                                    //   "2006-01-02 15:04:05.000000000 -0700 MST"
//           }
//
//
//  (2) error
//         - If the method completes successfully, the returned error instance is
//           set to nil. If errors are encountered, the returned error instance is populated
//           with an error message.
//
func (tzDto TimeZoneDto) New(
	tIn time.Time,
	timeZoneOutLocationName string,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewStartEndTimes() "
	tZoneDtoUtil := timeZoneDtoUtility{}

	if len(timeZoneOutLocationName) == 0 {
		return TimeZoneDto{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeZoneOutLocationName",
				inputParameterValue: "Input parameter 'timeZoneOutLocationName' is an EMPTY string!",
				errMsg:              "",
				err:                 nil,
			}
	}

	return tZoneDtoUtil.newTzDto( tIn, timeZoneOutLocationName, dateTimeFmtStr, ePrefix)
}

// NewAddDate - receives four parameters: a TimeZoneDto 'tzuIn' and integer values for
// 'years', 'months' and 'days'.  The 'years', 'months' and 'days' values are added to the
// 'tzuIn' date time values and returned as a new TimeZoneDto instance.
//
// Input Parameters
// ================
//
// years     int  - Number of years added to tzuIn value.
// months    int  - Number of months added to tzuIn value.
// days      int  - Number of days added to tzuIn value.
//
// Note: Negative date values may be used to subtract date values from the
//    tzuIn value.
//
// dateTimeFmtStr string  - A date time format string which will be used
//                          to format and display 'dateTime'. Example:
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                          If 'dateTimeFmtStr' is submitted as an
//                          'empty string', a default date time format
//                          string will be applied. The default date time
//                          format string is:
//                          TZDtoDefaultDateTimeFormatStr =
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                                (2) an Error type
//
//  (1) TimeZoneDto -  The date input parameters are added to 'tzuIn to produce,
//                     populate and return a TimeZoneDto structure defined as follows:
//
//    Description  string       // Unused - available for tagging, classification or
//                              //  labeling.
//    TimeIn       DateTzDto    // Original input time value
//
//   type TimeZoneDto struct {
//    TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//    TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                              //   equivalent to TimeIn
//    TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                              //   'Local' is the Time Zone Location used by the host computer.
//    DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                              //  Date Time text displays. The Default format string is:
//                              //   "2006-01-02 15:04:05.000000000 -0700 MST"
//  }
//
//
// (2) error - If the method completes successfully, the returned error instance is
//       set to nil. If errors are encountered, the returned error object is
//        populated with an error message.
//
func (tzDto TimeZoneDto) NewAddDate(
	tzuIn TimeZoneDto,
	years,
	months,
	days int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewAddDate() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.newAddDateTime(
		&tzuIn,
		years,
		months,
		days,
		0,          // hours
		0,          // minutes
		0,          // seconds
		0,          // milliseconds
		0,          // microseconds
		0,          // subMicrosecondNanoseconds
		dateTimeFmtStr,
		ePrefix)
}

// NewAddDateTime - Receives a TimeZoneDto input parameter, 'tzuIn'
// and returns a new TimeZoneDto instance equal to 'tzuIn' plus the
// time value of the remaining time element input parameters.
//
// Input Parameters
// ================
//
//   tzdtoIn   TimeZoneDto - Base TimeZoneDto object to
//                  which time elements will be added.
//
//   years    int  - Number of years added to 'tzuIn'
//
//   months    int  - Number of months added to 'tzuIn'
//
//   days     int  - Number of days added to 'tzuIn'
//
//   hours    int  - Number of hours added to 'tzuIn'
//
//   minutes   int  - Number of minutes added to 'tzuIn'
//
//   seconds   int  - Number of seconds added to 'tzuIn'
//
//   milliseconds int  - Number of milliseconds added to 'tzuIn'
//
//   microseconds int  - Number of microseconds added to 'tzuIn'
//
//   subMicrosecondNanoseconds int  - Number of subMicrosecondNanoseconds added to 'tzuIn'
//
//   Note:  Input time element parameters may be either negative or positive.
//       Negative values will subtract time from the returned TimeZoneDto instance.
//
//   dateTimeFmtStr string
//             - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//               If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               TZDtoDefaultDateTimeFormatStr =
//                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//
//   There are two return values:  (1) a TimeZoneDto Type
//                                 (2) an Error type
//
//   (1) TimeZoneDto -  If successful, this method returns a valid, populated TimeZoneDto
//            instance which is equal to the time value of 'tzuIn' plus the other
//            input parameter date-time elements. The TimeZoneDto structure
//           is defined as follows:
//
//         type TimeZoneDto struct {
//          Description  string       // Unused - available for tagging, classification or
//                                    //  labeling.
//          TimeIn       DateTzDto    // Original input time value
//          TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//          TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                                    //   equivalent to TimeIn
//          TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                                    //   'Local' is the Time Zone Location used by the host computer.
//          DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                                    //  Date Time text displays. The Default format string is:
//                                    //   "2006-01-02 15:04:05.000000000 -0700 MST"
//         }
//
//   (2) error - If errors are encountered, this method returns an error instance populated with
//               a valid 'error' message. If the method completes successfully the returned error
//               error type is set to 'nil'.
//
func (tzDto TimeZoneDto) NewAddDateTime(
	tzDtoIn TimeZoneDto,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewAddDateTime() "

	tZoneUtil := timeZoneDtoUtility{}

	return tZoneUtil.newAddDateTime(
		&tzDtoIn,
		years,
		months,
		days,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds,
		dateTimeFmtStr,
		ePrefix)
}

// NewAddDuration - receives two input parameters, a TimeZoneDto 'tzuIn' and a
// time 'duration'. 'tzuIn' is adjusted for the specified 'duration' and the resulting
// new TimeZoneDto is returned.
//
// Input Parameters
// ================
//
// tzdtoIn TimeZoneDto
//           - The second parameter, 'duration', will be added
//             to this TimeZoneDto.
//
// duration time.Duration
//           - This duration value will be added to the
//             'tzuIn' input parameter to create, populate and
//             return a new updated TimeZoneDto instance.
//
// dateTimeFmtStr string
//             - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//               If 'dateTimeFmtStr' is submitted as an
//               'empty string', a default date time format
//               string will be applied. The default date time
//               format string is:
//               TZDtoDefaultDateTimeFormatStr =
//                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Note:  Input parameter 'duration' will accept both positive and negative values.
//        Negative values will effectively subtract the duration from 'tzuIn' time
//        values.
//
// Returns
// =======
//  There are two return values:  (1) a TimeZoneDto Type
//                                (2) an Error type
//
//   (1) TimeZoneDto - The input parameter 'duration' is added to 'tzuIn to produce, populate and return
//                     a TimeZoneDto structure:
//
//      type TimeZoneDto struct {
//       Description  string       // Unused - available for tagging, classification or
//                                 //  labeling.
//       TimeIn       DateTzDto    // Original input time value
//       TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//       TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                                 //   equivalent to TimeIn
//       TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                                 //   'Local' is the Time Zone Location used by the host computer.
//       DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                                 //  Date Time text displays. The Default format string is:
//                                 //   "2006-01-02 15:04:05.000000000 -0700 MST"
//      }
//
//   (2) error - If errors are encountered, this method returns an error instance populated with
//               a valid 'error' message. If the method completes successfully the returned error
//               error type is set to 'nil'.
//
func (tzDto TimeZoneDto) NewAddDuration(
	tzdtoIn TimeZoneDto,
	duration time.Duration,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewAddDuration() "

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut := tZoneUtil.copyOut(&tzdtoIn, ePrefix)

	tZoneUtil.setDateTimeFormat(&tzuOut, dateTimeFmtStr, ePrefix)

	err := tZoneUtil.addDuration(
		&tzuOut,
		duration,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	return tzuOut, nil
}

// NewAddTime - returns a new TimeZoneDto equivalent to the input TimeZoneDto Plus time elements.
//
// Input Parameters:
// =================
//
//   tzdtoIn TimeZoneDto - The base TimeZoneDto to which
//                         time values will be added.
//
//   hours        int    - Number of hours to be added to tzuIn
//
//   minutes      int    - Number of minutes to be added to tzuIn
//
//   seconds      int    - Number of seconds to be added to tzuIn
//
//   milliseconds int    - Number of milliseconds to be added to tzuIn
//
//   microseconds int    - Number of microseconds to be added to tzuIn
//
//   subMicrosecondNanoseconds  int    - Number of subMicrosecondNanoseconds to be added to tzuIn
//
//   Note: Negative time values may be used to subtract time from 'tzuIn'.
//
//   dateTimeFmtStr string - A date time format string which will be used
//                to format and display 'dateTime'. Example:
//                "2006-01-02 15:04:05.000000000 -0700 MST"
//
//               If 'dateTimeFmtStr' is submitted as an
//                'empty string', a default date time format
//                string will be applied. The default date time
//                format string is:
//                TZDtoDefaultDateTimeFormatStr =
//                  "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Return Values
// =============
//
//  There are two return values:  (1) a TimeZoneDto Type
//                                (2) an Error type
//
//   TimeZoneDto -  The time input parameters are added to 'tzuIn to produce, populate and return
//             a TimeZoneDto structure:
//
//             type TimeZoneDto struct {
//                  Description string     // Unused. Available for tagging and classification.
//                  TimeIn      time.Time    // Original input time value
//                  TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//                  TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//                                //   based on a specific Time Zone Location.
//
//                  TimeOutLoc  *time.Location // Time Zone Location associated with TimeOut
//                  TimeUTC     time.Time    // TimeUTC (Universal Coordinated Time) value
//                                    equivalent to TimeIn
//
//                  TimeLocal  time.Time    // Equivalent to TimeIn value converted to the 'Local'
//                                // Time Zone Location. 'Local' is the Time Zone Location
//                                //  used by the host computer.
//             }
//
//   error - If the method completes successfully, the returned error instance is
//           set to nil. If errors are encountered, the returned error object is populated
//           with an error message.
//
func (tzDto TimeZoneDto) NewAddTime(
	tzdtoIn TimeZoneDto,
	hours, minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewAddTime() "

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut := tZoneUtil.copyOut(&tzdtoIn, ePrefix)

	tZoneUtil.setDateTimeFormat(&tzuOut, dateTimeFmtStr, ePrefix)

	err := tZoneUtil.addTime(
		&tzuOut,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	return tzuOut, nil
}

// NewDateTz - Receives a DateTzDto instance and converts the DateTime associated with this
// DateTzDto instance to a TimeZoneDto instance. The DateTzDto.DateTime value is converted to
// TimeZoneDto.TimeOut using the input parameter, 'tZoneOutLocation', and returned as part
// of the newly created TimeZoneDto instance.
//
// Input Parameters
// ================
//
//  dateTzDto  DateTzDto - Input parameter from which dateTzDto.DateTime
//             will be extracted to form the TimeZoneDto.TimeIn
//             value for the returned TimeZoneDto instance.
//
//             A DateTzDto structure is defined as follows:
//              type DateTzDto struct {
//               Description   string    // Unused, available for classification, labeling or description
//               Year          int       // Year Number
//               Month         int       // Month Number
//               Day           int       // Day Number
//               Hour          int       // Hour Number
//               Minute        int       // Minute Number
//               Second        int       // Second Number
//               Millisecond   int       // Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//               Microsecond   int       // Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//               Nanosecond    int       // Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//                                       // Nanosecond = TotalNanoSecs - millisecond subMicrosecondNanoseconds - microsecond subMicrosecondNanoseconds
//               TotalNanoSecs int64     // Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//               DateTime      time.Time // DateTime value for this DateTzDto Type
//               DateTimeFmt   string    // Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//               TimeZone      TimeZoneDefinition // Contains a detailed description of the Time Zone and Time Zone Location
//                                            //  associated with this date time.
//              }
//
//
//  tZoneOutLocationName - string -
//                         The first input time value, 'tIn' will have its time zone
//                         changed to a new time zone location specified by this
//                         second parameter, 'tZoneOutLocationName'. The new time
//                         associated with 'tZoneOutLocationName' is assigned to
//                         the TimeZoneDto data field. The 'tZoneOutLocationName'
//                         time zone location must be designated as one of three
//                         types of time zones:
//
//                         (1) the string 'Local' - signals the designation of the
//                             time zone location used by the host computer.
//
//                         (2) IANA Time Zone Location -
//                            See https://golang.org/pkg/time/#LoadLocation
//                            and https://www.iana.org/time-zones to ensure that
//                            the IANA Time Zone Database is properly configured
//                            on your system. Note: IANA Time Zone Data base is
//                            equivalent to 'tz database'.
//                              Examples:
//                               "America/New_York"
//                               "America/Chicago"
//                               "America/Denver"
//                               "America/Los_Angeles"
//                               "Pacific/Honolulu"
//
//                         (3) A Military Time Zone
//                             Reference:
//                              https://en.wikipedia.org/wiki/List_of_military_time_zones
//                              http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                              https://www.timeanddate.com/time/zones/military
//                              https://www.timeanddate.com/worldclock/timezone/alpha
//                              https://www.timeanddate.com/time/map/
//
//                             Examples:
//                               "Alpha"   or "A"
//                               "Bravo"   or "B"
//                               "Charlie" or "C"
//                               "Delta"   or "D"
//                               "Zulu"    or "Z"
//
//  dateTimeFmtStr string  - A date time format string which will be used
//                           to format and display 'dateTime'. Example:
//                           "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                           If 'dateTimeFmtStr' is submitted as an
//                           'empty string', a default date time format
//                           string will be applied. The default date time
//                           format string is:
//                            TZDtoDefaultDateTimeFormatStr =
//                              "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Return Values
// =============
//
// There are two returns:
//             (1) A TimeZoneDto instance
//             (2) An error type
//
// (1) TimeZoneDto
//    If successful, this method creates a new TimeZoneDto,
//    populated with, TimeIn, TimeOut, TimeUTC and TimeLocal
//    date time values plus time zone information.
//
//    A TimeZoneDto structure is defined as follows:
//
//    type TimeZoneDto struct {
//     Description  string       // Unused - available for tagging, classification or
//                               //  labeling.
//     TimeIn       DateTzDto    // Original input time value
//     TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//     TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                               //   equivalent to TimeIn
//     TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                               //   'Local' is the Time Zone Location used by the host computer.
//     DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                               //  Date Time text displays. The Default format string is:
//                               //   "2006-01-02 15:04:05.000000000 -0700 MST"
//    }
//
//
// (2) error - If errors are encountered, this method returns an error instance populated with
//             a valid 'error' message. If the method completes successfully the returned error
//             error type is set to 'nil'.
//
//
func (tzDto TimeZoneDto) NewDateTz(
	dateTzDtoIn DateTzDto,
	tZoneOutLocationName,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewDateTz() "

	if err:=dateTzDtoIn.IsValid(); err!=nil {
		return TimeZoneDto{},
			fmt.Errorf(ePrefix + "\nError: Input parameter 'dateTzDtoIn' is INVALID!\n" +
				"Error='%v'\n", err.Error())
	}

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut, err := tZoneUtil.newTzDto(dateTzDtoIn.GetDateTimeValue(), tZoneOutLocationName, dateTimeFmtStr, ePrefix)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf("NewStartEndTimes Time Zone Dto creation Failed!\n" +
				"%v", err.Error())
	}

	return tzuOut, nil
}

// NewFromTzSpec - Receives a date time parameter and creates
// a new 'TimeZoneDto' instance using the input parameter
// 'tzSpec', which is an instance of time zone specification.
//
// Input Parameters
// ================
//
//  dateTimeIn time.Time - The input date time which will be
//                         used as a basis for calculating the
//                         equivalent time in another time zone.
//
//  tzSpec TimeZoneSpecification -
//                         The first input time parameter, 'dateTimeIn' will have
//                         its time zone changed to a new time zone location
//                         specified by this second parameter, 'tzSpec'. This
//                         second parameter is an instance of 'TimeZoneSpecification'
//                         which is used to specify one of three types of time zones.
//
//                         (1) the string 'Local' - signals the designation of the
//                             time zone location used by the host computer.
//
//                         (2) IANA Time Zone Location -
//                            See https://golang.org/pkg/time/#LoadLocation
//                            and https://www.iana.org/time-zones to ensure that
//                            the IANA Time Zone Database is properly configured
//                            on your system. Note: IANA Time Zone Data base is
//                            equivalent to 'tz database'.
//                              Examples:
//                               "America/New_York"
//                               "America/Chicago"
//                               "America/Denver"
//                               "America/Los_Angeles"
//                               "Pacific/Honolulu"
//
//                         (3) A Military Time Zone
//                             Reference:
//                              https://en.wikipedia.org/wiki/List_of_military_time_zones
//                              http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                              https://www.timeanddate.com/time/zones/military
//                              https://www.timeanddate.com/worldclock/timezone/alpha
//                              https://www.timeanddate.com/time/map/
//
//                             Examples:
//                               "Alpha"   or "A"
//                               "Bravo"   or "B"
//                               "Charlie" or "C"
//                               "Delta"   or "D"
//                               "Zulu"    or "Z"
//
//  dateTimeFmtStr string  - A date time format string which will be used
//                 to format and display 'dateTime'. Example:
//                "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                 If 'dateTimeFmtStr' is submitted as an
//                 'empty string', a default date time format
//                 string will be applied. The default date time
//                 format string is:
//                 TZDtoDefaultDateTimeFormatStr =
//                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Return Values
// =============
//
//  There are two return values:  (1) A TimeZoneDto Type
//                                (2) An Error type
//
//  (1) TimeZoneDto
//         - The two input parameters are used to populate and return
//           a TimeZoneDto structure:
//
//           type TimeZoneDto struct {
//            Description  string     // Unused - available for tagging, classification or
//                                    //  labeling.
//            TimeIn       DateTzDto  // Original input time value
//            TimeOut      DateTzDto  // TimeOut - 'TimeIn' value converted to TimeOut
//            TimeUTC      DateTzDto  // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                                    //   equivalent to TimeIn
//            TimeLocal   DateTzDto   // TimeIn value converted to the 'Local' Time Zone Location.
//                                    //   'Local' is the Time Zone Location used by the host computer.
//            DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                                    //  Date Time text displays. The Default format string is:
//                                    //   "2006-01-02 15:04:05.000000000 -0700 MST"
//           }
//
//
//  (2) error
//         - If the method completes successfully, the returned error instance is
//           set to nil. If errors are encountered, the returned error instance is populated
//           with an error message.
//
func (tzDto TimeZoneDto) NewFromTzSpec(
	dateTimeIn time.Time,
	tzSpec TimeZoneSpecification,
	dateTimeFmtStr string) (
			TimeZoneDto,
			error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewFromTzSpec() "

	tzDto2 := TimeZoneDto{}

	tZoneUtil := timeZoneDtoUtility{}

	tzDto2.DateTimeFmt = tZoneUtil.preProcessDateFormatStr(dateTimeFmtStr)

	err := tZoneUtil.setTimeIn(
		&tzDto2,
		dateTimeIn,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	err = tZoneUtil.setTimeOutTzSpec(
		&tzDto2,
		dateTimeIn,
		TzConvertType.Relative(),
		tzSpec,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	err = tZoneUtil.setUTCTime(
		&tzDto2,
		dateTimeIn,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}


	err = tZoneUtil.setUTCTime(
		&tzDto2,
		dateTimeIn,
		TzConvertType.Relative(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	return tzDto2, nil
}

// NewTimeAddDate - returns a new TimeZoneDto. The TimeZoneDto is initialized
// with the 'tIn' time parameter. The 'TimeOut' data field will be set to the 'tIn'
// value and the time zone location specified by the second parameter, 'tZoneLocation'.
// The method will then add the remaining date element parameters to the new TimeZoneDto
// instance and return it to the calling function.
//
// Input Parameters
// ================
//  tIn   time.Time
//             - Initial time value assigned to 'TimeIn' field
//               of the new TimeZoneDto.
//
//  tZoneOutLocationName string
//              - The first input time value, 'tIn' will have its time zone
//                changed to a new time zone location specified by this second
//                parameter, 'tZoneOutLocationName'. The new time associated
//                with 'tZoneOutLocationName' is assigned to a TimeZoneDto
//                data field. The 'tZoneOutLocationName' must specify a time
//                zone location which qualifies as one of the three types of
//                time zones shown below:
//
//                (1) the string 'Local' - signals the designation of the
//                    time zone location used by the host computer.
//
//                (2) IANA Time Zone Location -
//                   See https://golang.org/pkg/time/#LoadLocation
//                   and https://www.iana.org/time-zones to ensure that
//                   the IANA Time Zone Database is properly configured
//                   on your system. Note: IANA Time Zone Data base is
//                   equivalent to 'tz database'.
//                     Examples:
//                      "America/New_York"
//                      "America/Chicago"
//                      "America/Denver"
//                      "America/Los_Angeles"
//                      "Pacific/Honolulu"
//
//                (3) A Military Time Zone
//                    Reference:
//                     https://en.wikipedia.org/wiki/List_of_military_time_zones
//                     http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                     https://www.timeanddate.com/time/zones/military
//                     https://www.timeanddate.com/worldclock/timezone/alpha
//                     https://www.timeanddate.com/time/map/
//
//                    Examples:
//                      "Alpha"   or "A"
//                      "Bravo"   or "B"
//                      "Charlie" or "C"
//                      "Delta"   or "D"
//                      "Zulu"    or "Z"
//
//
//  years    int  - Number of years added to initial TimeZoneDto value.
//  months    int  - Number of months added to initial TimeZoneDto value.
//  days     int  - Number of days added to initial TimeZoneDto value.
//
//  Note: Negative date values may be used to subtract date values from the
//     initial TimeZoneDto.
//
//  dateTimeFmtStr string  - A date time format string which will be used
//                to format and display 'dateTime'. Example:
//                "2006-01-02 15:04:05.000000000 -0700 MST"
//
//               If 'dateTimeFmtStr' is submitted as an
//                'empty string', a default date time format
//                string will be applied. The default date time
//                format string is:
//                TZDtoDefaultDateTimeFormatStr =
//                  "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Return Values
// =============
//
//  There are two return values:  (1) a TimeZoneDto Type
//                                (2) an Error type
//
//  TimeZoneDto
//         - The date input parameters are added to a TimeZoneDto created from
//           input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//           instance is then returned to the calling function. A TimeZoneDto structure
//           is defined as follows:
//
//            type TimeZoneDto struct {
//             Description  string       // Unused - available for tagging, classification or
//                                       //  labeling.
//             TimeIn       DateTzDto    // Original input time value
//             TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//             TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                                       //   equivalent to TimeIn
//             TimeLocal   DateTzDto     // TimeIn value converted to the 'Local' Time Zone Location.
//                                       //   'Local' is the Time Zone Location used by the host computer.
//             DateTimeFmt   string      // Date Time Format String. This format string is used to format
//                                       //  Date Time text displays. The Default format string is:
//                                       //   "2006-01-02 15:04:05.000000000 -0700 MST"
//            }
//
// error
//         - If the method completes successfully, the returned error instance is
//           set to nil. If errors are encountered, the returned error instance is populated
//           with an error message.
//
func (tzDto TimeZoneDto) NewTimeAddDate(
	tIn time.Time,
	tZoneOutLocationName string,
	years,
	months,
	days int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewTimeAddDate() "

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut,
	err := tZoneUtil.newTzDto(
		tIn,
		tZoneOutLocationName,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{},
		fmt.Errorf("Creation of TimeZoneDto, 'tzuOut', FAILED!\n" +
			"%v", err.Error())
	}

	err = tZoneUtil.addDateTime(
		&tzuOut,
		years,
		months,
		days,
		0,
		0,
		0,
		0,
		0,
		0,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	return tzuOut, nil
}

// NewTimeAddDateTime - returns a new TimeZoneDto. The TimeZoneDto is initialized
// with the 'tIn' time parameter. The 'TimeOut' data field will be set to the 'tIn'
// value adjusted for the time zone location specified by the second parameter, 'tZoneLocation'.
// The method will then add the remaining date-time element parameters to the new TimeZoneDto
// instance and return it to the calling function.
//
// Input Parameters
// ================
//
//  tIn   time.Time
//              - Initial time value assigned to 'TimeIn' field
//                of the new TimeZoneDto.
//
//  tZoneLocation string
//              - The first input time value, 'tIn' will have its time zone
//                changed to a new time zone location specified by this second
//                parameter, 'tZoneLocation'. This time zone location must be
//                designated as one of two values:
//
//                (1) the string 'Local' - signals the designation of the
//                  time zone location used by the host computer.
//
//               (2) IANA Time Zone Location -
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system. Note: IANA Time Zone Data base is
//                  equivalent to 'tz database'.
//                 Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//  years    int  - Number of years added to initial TimeZoneDto value.
//  months    int  - Number of months added to initial TimeZoneDto value.
//  days     int  - Number of days added to initial TimeZoneDto value.
//  hours    int  - Number of hours to be added to initial TimeZoneDto value.
//  minutes   int  - Number of minutes to be added to initial TimeZoneDto value.
//  seconds   int  - Number of seconds to be added to initial TimeZoneDto value.
//  milliseconds int  - Number of milliseconds to be added to initial TimeZoneDto value.
//  microseconds int  - Number of microseconds to be added to initial TimeZoneDto value.
//  subMicrosecondNanoseconds int  - Number of subMicrosecondNanoseconds to be added to initial TimeZoneDto value.
//
//  Note: Negative date-time values may be used to subtract date-time from the initial TimeZoneDto.
//
//  dateTimeFmtStr string  - A date time format string which will be used
//                to format and display 'dateTime'. Example:
//                "2006-01-02 15:04:05.000000000 -0700 MST"
//
//               If 'dateTimeFmtStr' is submitted as an
//                'empty string', a default date time format
//                string will be applied. The default date time
//                format string is:
//                TZDtoDefaultDateTimeFormatStr =
//                  "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Returns
// =======
//
//  There are two return values:  (1) a TimeZoneDto Type
//                (2) an Error type
//
//  (1) TimeZoneDto
//        -  The date-time input parameters are added to a TimeZoneDto created from
//           input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//           instance is then returned to the calling function. A TimeZoneDto structure
//           is defined as follows:
//
//           type TimeZoneDto struct {
//            Description  string     // Unused - available for tagging, classification or
//                           //  labeling.
//            TimeIn       DateTzDto    // Original input time value
//            TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//            TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                           //   equivalent to TimeIn
//            TimeLocal   DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
//                           //   'Local' is the Time Zone Location used by the host computer.
//            DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                           //  Date Time text displays. The Default format string is:
//                           //   "2006-01-02 15:04:05.000000000 -0700 MST"
//           }
//
//  (2) error
//        - If the method completes successfully, the returned error instance is
//          set to nil. If errors are encountered, the returned error instance is populated
//          with an error message.
//
func (tzDto TimeZoneDto) NewTimeAddDateTime(
	tIn time.Time,
	tZoneLocation string,
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewTimeAddDateTime() "

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut, err := tZoneUtil.newTzDto(
		tIn,
		tZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	err = tZoneUtil.addDateTime(
		&tzuOut,
		years,
		months,
		days,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	return tzuOut, nil
}

// NewTimeAddDuration - receives a 'tIn' time.Time parameter and a 'tZoneLocation' parameter
// which are used to construct an initial TimeZoneDto instance. The 'TimeOut'
// data field of this initial TimeZoneDto will contain the value of 'tIn'
// converted to a different time zone specified by 'tZoneLocation'.
//
// The 'duration' parameter will be added to this initial TimeZoneDto and
// an updated TimeZoneDto instance will be returned to the calling function.
//
// Input Parameters
// ================
// tIn    time.Time
//            - Initial time value assigned to 'TimeIn' field
//              of the new TimeZoneDto.
//
// tZoneLocationName string
//            - The first input time value, 'tIn' will have its time zone
//              changed to a new time zone location specified by this second
//              parameter, 'tZoneLocationName'. This time zone location or
//              time zone name, must be designated as one of three types of
//              time zones:
//
//                (1) the string 'Local' - signals the designation of the
//                    time zone location used by the host computer.
//
//                (2) IANA Time Zone Location -
//                   See https://golang.org/pkg/time/#LoadLocation
//                   and https://www.iana.org/time-zones to ensure that
//                   the IANA Time Zone Database is properly configured
//                   on your system. Note: IANA Time Zone Data base is
//                   equivalent to 'tz database'.
//                     Examples:
//                      "America/New_York"
//                      "America/Chicago"
//                      "America/Denver"
//                      "America/Los_Angeles"
//                      "Pacific/Honolulu"
//
//                (3) A Military Time Zone
//                    Reference:
//                     https://en.wikipedia.org/wiki/List_of_military_time_zones
//                     http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                     https://www.timeanddate.com/time/zones/military
//                     https://www.timeanddate.com/worldclock/timezone/alpha
//                     https://www.timeanddate.com/time/map/
//
//                    Examples:
//                      "Alpha"   or "A"
//                      "Bravo"   or "B"
//                      "Charlie" or "C"
//                      "Delta"   or "D"
//                      "Zulu"    or "Z"
//
// duration  time.Duration
//            - An int64 duration value which is added to the date time
//              value of the initial TimeZoneDto created from 'tIn' and 'tZoneLocation'.
//
//              Note: Negative duration values may be used to subtract time duration
//              from the initial TimeZoneDto date time values.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                TZDtoDefaultDateTimeFormatStr =
//                   "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Return Values
// =============
//
//  There are two return values:  (1) a TimeZoneDto Type
//                                (2) an Error type
//
//  (1) TimeZoneDto
//            - The duration input parameter is added to a TimeZoneDto created from
//              input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//              instance is then returned to the calling function.
//
//              A TimeZoneDto structure is defined as follows:
//
//               type TimeZoneDto struct {
//                Description  string     // Unused - available for tagging, classification or
//                               //  labeling.
//                TimeIn       DateTzDto    // Original input time value
//                TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//                TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                               //   equivalent to TimeIn
//                TimeLocal   DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
//                               //   'Local' is the Time Zone Location used by the host computer.
//                DateTimeFmt   string    // Date Time Format String. This format string is used to format
//                               //  Date Time text displays. The Default format string is:
//                               //   "2006-01-02 15:04:05.000000000 -0700 MST"
//               }
//
//
// (2) error
//            - If errors are encountered, this method returns an error instance populated with
//              a valid 'error' message. If the method completes successfully the returned error
//              error type is set to 'nil'.
//
func (tzDto TimeZoneDto) NewTimeAddDuration(
	tIn time.Time,
	tZoneLocationName string,
	duration time.Duration,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewTimeAddDuration() "

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut, err := tZoneUtil.newTzDto(tIn, tZoneLocationName, dateTimeFmtStr, ePrefix)

	if err != nil {
		return TimeZoneDto{},
			fmt.Errorf("Creation of a NewStartEndTimes Time Zone Dto FAILED!\n" +
			"%v", err.Error())
	}

	err = tZoneUtil.addDuration(
		&tzuOut,
		duration,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	return tzuOut, nil
}

// NewTimeAddTime - receives a 'tIn' time.Time parameter and a 'tZoneLocation' parameter
// which are used to construct an initial TimeZoneDto instance. The 'TimeOut'
// data field of this initial TimeZoneDto will contain the value of 'tIn'
// converted to a different time zone specified by 'tZoneLocation'.
//
// The remaining time parameters will be added to this initial TimeZoneDto and
// the updated TimeZoneDto will be returned to the calling function.
//
// Input Parameters
// ================
//
//  tIn    time.Time
//         - Initial time value assigned to 'TimeIn' field
//           of the new TimeZoneDto.
//
//  tZoneLocationName string
//         - The first input time value, 'tIn' will have its time zone
//           changed to a new time zone location specified by this second
//           parameter, 'tZoneLocation'. This time zone location, or time
//           zone name, must be designated as one of three types of time
//           zones:
//
//                 (1) the string 'Local' - signals the designation of the
//                     time zone location used by the host computer.
//
//                 (2) IANA Time Zone Location -
//                    See https://golang.org/pkg/time/#LoadLocation
//                    and https://www.iana.org/time-zones to ensure that
//                    the IANA Time Zone Database is properly configured
//                    on your system. Note: IANA Time Zone Data base is
//                    equivalent to 'tz database'.
//                      Examples:
//                       "America/New_York"
//                       "America/Chicago"
//                       "America/Denver"
//                       "America/Los_Angeles"
//                       "Pacific/Honolulu"
//
//                 (3) A Military Time Zone
//                     Reference:
//                      https://en.wikipedia.org/wiki/List_of_military_time_zones
//                      http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                      https://www.timeanddate.com/time/zones/military
//                      https://www.timeanddate.com/worldclock/timezone/alpha
//                      https://www.timeanddate.com/time/map/
//
//                     Examples:
//                       "Alpha"   or "A"
//                       "Bravo"   or "B"
//                       "Charlie" or "C"
//                       "Delta"   or "D"
//                       "Zulu"    or "Z"
//
//  hours        int    - Number of hours to be added to initial time, 'tIn'
//
//  minutes      int    - Number of minutes to be added to initial time, 'tIn'
//
//  seconds      int    - Number of seconds to be added to initial 'time, tIn'
//
//  milliseconds int    - Number of milliseconds to be added to initial time, 'tIn'
//
//  microseconds int    - Number of microseconds to be added to initial time, 'tIn'
//
//  subMicrosecondNanoseconds  int    - Number of subMicrosecondNanoseconds to be added to initial time, 'tIn'
//
//              Note: Negative time values may be used to subtract time from
//                    initial time parameter, 'tIn'.
//
//  dateTimeFmtStr string
//         - A date time format string which will be used
//           to format and display 'dateTime'. Example:
//           "2006-01-02 15:04:05.000000000 -0700 MST"
//
//           If 'dateTimeFmtStr' is submitted as an
//           'empty string', a default date time format
//           string will be applied. The default date time
//           format string is:
//             TZDtoDefaultDateTimeFormatStr =
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Return Values
// =============
//  There are two return values:  (1) a TimeZoneDto Type
//                                (2) an Error type
//
//  (1) TimeZoneDto
//         - The time input parameters are added to a TimeZoneDto created from
//           input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//           instance is then returned to the calling function.
//
//           A TimeZoneDto structure is defined as follows:
//
//             type TimeZoneDto struct {
//              Description  string       // Unused - available for tagging, classification or
//                                        //  labeling.
//              TimeIn       DateTzDto    // Original input time value
//              TimeOut      DateTzDto    // TimeOut - 'TimeIn' value converted to TimeOut
//              TimeUTC      DateTzDto    // TimeUTC (Universal Coordinated Time aka 'Zulu') value
//                                        //   equivalent to TimeIn
//              TimeLocal    DateTzDto    // TimeIn value converted to the 'Local' Time Zone Location.
//                                        //   'Local' is the Time Zone Location used by the host computer.
//              DateTimeFmt  string       // Date Time Format String. This format string is used to format
//                                        //  Date Time text displays. The Default format string is:
//                                        //   "2006-01-02 15:04:05.000000000 -0700 MST"
//             }
//
//
//  (2) error
//         - If errors are encountered, this method returns an error instance populated with
//           a valid 'error' message. If the method completes successfully the returned error
//           error type is set to 'nil'.
//
func (tzDto TimeZoneDto) NewTimeAddTime(
	tIn time.Time,
	tZoneLocationName string,
	hours,
	minutes,
	seconds,
	milliseconds,
	microseconds,
	nanoseconds int,
	dateTimeFmtStr string) (TimeZoneDto, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.NewTimeAddTime() "

	tZoneUtil := timeZoneDtoUtility{}

	tzuOut, err := tZoneUtil.newTzDto(
		tIn,
		tZoneLocationName,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return TimeZoneDto{},
		fmt.Errorf("Creation of new TimeZoneDto, 'tzuOut', FAILED!\n" +
			"%v", err.Error())
	}

	err = tZoneUtil.addTime(
		&tzuOut,
		hours,
		minutes,
		seconds,
		milliseconds,
		microseconds,
		nanoseconds,
		ePrefix)

	if err != nil {
		return TimeZoneDto{}, err
	}

	return tzuOut, nil
}

// ReclassifyTimeWithNewTz - Receives a valid time (time.Time)
// value and changes the existing time zone to that specified
// in the 'tZoneLocationName' parameter.
//
// Input Parameters
// ================
//
//   dateTime time.Time
//          - Initial time whose time zone will be changed to
//            second input parameter, 'tZoneLocation'
//
//
//   timeConversionType TimeZoneConversionType
//          - This parameter determines the algorithm that will
//            be used to convert parameter 'dateTime' to the time
//            zone specified by parameter 'timeZoneName'.
//
//            TimeZoneConversionType is an enumeration type which
//            be used to convert parameter 'dateTime' to the time
//            must be set to one of two values:
//            This parameter determines the algorithm that will
//               TimeZoneConversionType(0).Absolute()
//               TimeZoneConversionType(0).Relative()
//            Note: You can also use the global variable
//            'TzConvertType' for easier access:
//               TzConvertType.Absolute()
//               TzConvertType.Relative()
//
//            Absolute Time Conversion - Identifies the 'Absolute' time
//            to time zone conversion algorithm. This algorithm provides
//            that a time value in time zone 'X' will be converted to the
//            same time value in time zone 'Y'.
//
//            For example, assume the time 10:00AM is associated with time
//            zone USA Central Standard time and that this time is to be
//            converted to USA Eastern Standard time. Applying the 'Absolute'
//            algorithm would convert ths time to 10:00AM Eastern Standard
//            time.  In this case the hours, minutes and seconds have not been
//            altered. 10:00AM in USA Central Standard Time has simply been
//            reclassified as 10:00AM in USA Eastern Standard Time.
//
//            Relative Time Conversion - Identifies the 'Relative' time to time
//            zone conversion algorithm. This algorithm provides that times in
//            time zone 'X' will be converted to their equivalent time in time
//            zone 'Y'.
//
//            For example, assume the time 10:00AM is associated with time zone
//            USA Central Standard time and that this time is to be converted to
//            USA Eastern Standard time. Applying the 'Relative' algorithm would
//            convert ths time to 11:00AM Eastern Standard time. In this case the
//            hours, minutes and seconds have been changed to reflect an equivalent
//            time in the USA Eastern Standard Time Zone.
//
// tZoneLocationName string
//          - The first input time value, 'tIn' will have its time zone
//            changed to a new time zone location specified by this second
//            parameter, 'tZoneLocation'. This time zone location must be
//            designated as one of three types of time zones:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                configured for the host computer executing this code.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                   Examples:
//                     "America/New_York"
//                     "America/Chicago"
//                     "America/Denver"
//                     "America/Los_Angeles"
//                     "Pacific/Honolulu"
//
//            (3) A valid Military Time Zone
//                Military time zones are commonly used in
//                aviation as well as at sea. They are also
//                known as nautical or maritime time zones.
//                Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//
//             Note:
//                 The source file 'timezonedata.go' contains over 600 constant
//                 time zone declarations covering all IANA and Military Time
//                 Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TZones'.
//
func (tzDto *TimeZoneDto) ReclassifyTimeWithNewTz(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	tZoneLocationName string) (time.Time, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.ReclassifyTimeWithNewTz() "

	tzMech := TimeZoneMechanics{}

	tzSpec,
	err := tzMech.GetTimeZoneFromName(
		dateTime,
		tZoneLocationName,
		timeConversionType,
		ePrefix)

	if err != nil {
		return time.Time{}, err
	}

	return tzSpec.GetReferenceDateTime(), nil
}

// SetDateTimeFormatStr - Sets the value of the TimeZoneDto.DateTimeFmt field.
//
// Input Parameter
// ===============
//
//
// dateTimeFmtStr string  - A date time format string which will be used
//                          to format and display 'dateTime'. Example:
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                          If 'dateTimeFmtStr' is submitted as an
//                          'empty string', a default date time format
//                          string will be applied. The default date time
//                          format string is:
//                          DEFAULTDATETIMEFORMAT =
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tzDto *TimeZoneDto) SetDateTimeFormatStr(dateTimeFmtStr string) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.SetDateTimeFormatStr() "

	tZoneUtil := timeZoneDtoUtility{}

	tZoneUtil.setDateTimeFormat(tzDto, dateTimeFmtStr, ePrefix)

	return
}

// Sub - Subtracts the input TimeZoneDto from the current TimeZoneDto
// and returns the duration. Duration is calculated as:
//       tzu.TimeLocal.Sub(tzu2.TimeUTC)
//
// The TimeUTC field is used to compute duration for this method.
//
func (tzDto *TimeZoneDto) Sub(tzu2 TimeZoneDto) (time.Duration, error) {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	ePrefix := "TimeZoneDto.Sub() "

	tZoneUtil := timeZoneDtoUtility{}

	err := tZoneUtil.isValidTimeZoneDto(tzDto, ePrefix)

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix +
			"\nError: Current TimeZoneDto (tzDto) is INVALID.\n" +
			"Error='%v'\n", err.Error())
	}

	err = tZoneUtil.isValidTimeZoneDto(&tzu2, ePrefix)

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix+"Error: Input Parameter 'tzu2' is INVALID! Error='%v'", err.Error())
	}

	return tzDto.TimeUTC.Sub(tzu2.TimeUTC), nil
}

// TimeWithoutTimeZone - Returns a Time String containing
// NO time zone. - When the returned string is converted to
// time.Time, it will default to a UTC time zone.
func (tzDto *TimeZoneDto) TimeWithoutTimeZone(t time.Time) string {

	if tzDto.lock == nil {
		tzDto.lock = new(sync.Mutex)
	}

	tzDto.lock.Lock()

	defer tzDto.lock.Unlock()

	return t.Format(FmtDateTimeNeutralDateFmt)
}
