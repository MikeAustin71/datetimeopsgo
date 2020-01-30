package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type DTimeMechanics struct {
	lock sync.Mutex
}

// AbsoluteTimeToTimeZoneDefConversion - Converts a given time to
// another time zone using the 'Absolute' conversion method.
// This means that the years, months, days, hours, minutes,
// seconds and nanoseconds of the original 'dateTime' are not
// changed. That time value is simply assigned to another
// designated time zone. The target time zone is derived from
// input parameter 'timeZoneDefDto', an instance of type
// 'TimeZoneDefinition'.
//
// For example, assume that 'dateTime' represents 10:00AM in USA
// time zone 'Central Standard Time'.  Using the 'Absolute'
// conversion method, and converting this time value to the USA
// Eastern Standard Time Zone would result in a date time of
// 10:00AM EST or Eastern Standard Time. The time value of
// 10:00AM is not changed, it is simply assigned to another
// time zone.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// dateTime            time.Time
//                      - The date time to be converted.
//
// timeZoneDefDto TimeZoneDefinition
//                      - A properly initialized 'TimeZoneDefinition'
//                        object encapsulating the time zone to which
//                       'dateTime' will be converted.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// time.Time  - The date time converted to the time zone specified in
//              in input parameter 'timeZoneDefDto'.
//
// error      - If the method completes successfully this value is set
//              to 'nil'. If an error is encountered, the returned error
//              value encapsulates an appropriate error message.
//
func (dtMech *DTimeMechanics) AbsoluteTimeToTimeZoneDefConversion(
	dateTime time.Time,
	timeZoneDefDto TimeZoneDefinition) (time.Time, error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix := "DTimeMechanics.AbsoluteTimeToTimeZoneDefConversion() "

	if dateTime.IsZero() {
		return time.Time{},
			errors.New(ePrefix +
				"\nError: Input parameter 'dateTime' is zero!")
	}


	if err := timeZoneDefDto.IsValid(); err != nil {
		return time.Time{},
			fmt.Errorf(ePrefix +
				"Input parameter 'timeZoneDefDto' is Invalid!\n" +
				"Error='%v'\n", err.Error())
	}

	return time.Date(dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		dateTime.Hour(),
		dateTime.Minute(),
		dateTime.Second(),
		dateTime.Nanosecond(),
		timeZoneDefDto.GetOriginalLocationPtr()), nil
}

// AbsoluteTimeToTimeZoneNameConversion - Converts a given time to
// another time zone using the 'Absolute' conversion method.
// This means that the years, months, days, hours, minutes,
// seconds and nanoseconds of the original 'dateTime' are not
// changed. That time value is simply assigned to another
// designated time zone. The target time zone is derived from
// input parameter 'timeZoneDefDto', an instance of type
// 'TimeZoneDefinition'.
//
// For example, assume that 'dateTime' represents 10:00AM in USA
// time zone 'Central Standard Time'.  Using the 'Absolute'
// conversion method, and converting this time value to the USA
// Eastern Standard Time Zone would result in a date time of
// 10:00AM EST or Eastern Standard Time. The time value of
// 10:00AM is not changed, it is simply assigned to another
// time zone.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// dateTime            time.Time  - The date time to be converted.
//
// timeZoneName           string  - A string containing a valid IANA,
//                                  Military or "Local" time zone.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// time.Time  - The date time converted to the time zone specified in
//              in input parameter 'timeZoneDefDto'.
//
// error      - If the method completes successfully this value is set
//              to 'nil'. If an error is encountered, the returned error
//              value encapsulates an appropriate error message.
//
func (dtMech *DTimeMechanics) AbsoluteTimeToTimeZoneNameConversion(
	dateTime time.Time,
	timeZoneName string,
	ePrefix string) (time.Time, error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.AbsoluteTimeToTimeZoneNameConversion() "

	if dateTime.IsZero() {
		return time.Time{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "dateTime",
				inputParameterValue: "",
				errMsg:              "Error: Input parameter 'dateTime' is zero!",
				err:                 nil,
			}
	}

	timeZoneName = strings.TrimRight(strings.TrimLeft(timeZoneName, " "), " ")

	if len(timeZoneName) == 0 {
		return time.Time{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeZoneName",
				inputParameterValue: "",
				errMsg:              "Error: 'timeZoneName' is an empty string!",
				err:                 nil,
			}
	}

	var err error
	tzSpec := TimeZoneSpecification{}
	tzMech := TimeZoneMechanics{}

	tzSpec,
	err = tzMech.GetTimeZoneFromName(
		dateTime,
		timeZoneName,
		TzConvertType.Absolute(),
		ePrefix)

	if err != nil {
		return time.Time{}, err
	}

	return tzSpec.referenceDateTime, nil
}

// AllocateSecondsToHrsMinSecs - Useful in calculating offset hours,
// minutes and seconds from UTC+0000. A total signed seconds value
// is passed as an input parameter. This method then breaks down
// hours, minutes and seconds as positive integer values. The sign
// of the hours, minutes and seconds is returned in the 'sign'
// parameter as +1 or -1.
//
func (dtMech *DTimeMechanics) AllocateSecondsToHrsMinSecs(
	signedTotalSeconds int) (hours, minutes, seconds, sign int) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	hours = 0
	minutes = 0
	seconds = 0
	sign = 1

	if signedTotalSeconds == 0 {
		return hours, minutes, seconds, sign
	}

	if signedTotalSeconds < 0 {
		sign = -1
	}

	remainingSeconds := signedTotalSeconds

	remainingSeconds *= sign

	hours = remainingSeconds / 3600

	remainingSeconds -= hours * 3600

	if remainingSeconds > 0 {
		minutes = remainingSeconds / 60
		remainingSeconds -= minutes * 60
	}

	seconds = remainingSeconds

	return hours, minutes, seconds, sign
}

// GetTimeZoneFromDateTime - Analyzes a date time object
// and returns a valid time zone in the form of a
// 'TimeZoneSpecification' instance.
//
// Because date time objects (time.Time) do not support
// Military Time Zones; therefore, Military Time Zones
// are never returned by this method.
//
func (dtMech *DTimeMechanics) GetTimeZoneFromDateTime(
	dateTime time.Time,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.GetTimeZoneFromName() "

	tzMech := TimeZoneMechanics{}

	return tzMech.GetTimeZoneFromDateTime(dateTime, ePrefix)
}

// GetTimeZoneFromName - Analyzes a time zone name passed
// through input parameter, 'timeZoneName'. If valid, the
// method populates and returns a 'TimeZoneSpecification'
// instance.
//
// This method will accept and successfully process one
// of three types of time zones:
//
//   (1) The time zone "Local", which Golang accepts as
//       the time zone currently configured on the host
//       computer.
//
//   (2) IANA Time Zone - A valid IANA Time Zone from the
//       IANA database.
//       See https://golang.org/pkg/time/#LoadLocation
//       and https://www.iana.org/time-zones to ensure that
//       the IANA Time Zone Database is properly configured
//       on your system.
//
//       IANA Time Zone Examples:
//         "America/New_York"
//         "America/Chicago"
//         "America/Denver"
//         "America/Los_Angeles"
//         "Pacific/Honolulu"
//         "Etc/UTC" = GMT or UTC
//
//    (3) A Military Time Zone
//        Reference:
//         https://en.wikipedia.org/wiki/List_of_military_time_zones
//         http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//         https://www.timeanddate.com/time/zones/military
//         https://www.timeanddate.com/worldclock/timezone/alpha
//         https://www.timeanddate.com/time/map/
//
//        Examples:
//          "Alpha"   or "A"
//          "Bravo"   or "B"
//          "Charlie" or "C"
//          "Delta"   or "D"
//          "Zulu"    or "Z"
//
// If the time zone "Zulu" is passed to this method, it will be
// classified as a Military Time Zone.
//
func (dtMech *DTimeMechanics) GetTimeZoneFromName(
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeUtility.DTimeMechanics() "

	tzSpec = TimeZoneSpecification{}
	err = nil

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Input parameter 'dateTime' has " +
				"a Zero value!",
			err:                 nil,
		}

		return tzSpec, err
	}

	if timeConversionType < TzConvertType.Absolute() ||
		timeConversionType > TzConvertType.Relative() {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: timeConversionType.String(),
			errMsg:              "Input Parameter 'timeConversionType' " +
				"contains an invalid value!",
			err:                 nil,
		}
		return tzSpec, err
	}

	if len(timeZoneName) == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "Input parameter is an EMPTY String!",
			err:                 nil,
		}

		return tzSpec, err
	}

	tzMech := TimeZoneMechanics{}

	return tzMech.GetTimeZoneFromName(
		dateTime,
		timeZoneName,
		timeConversionType,
		ePrefix)
}

// GetUtcOffsetTzAbbrvFromDateTime - Receives a time.Time,
// date time, input parameter and extracts and returns the
// 5-character UTC offset and the time zone abbreviation.
//
// UTC Offsets are returned in the format illustrated by the
// following examples:
//
//   +1030
//   -0500
//   +1100
//   -1100
//
// The time zone abbreviation,'tzAbbrv', is formatted as
// shown in the following example ('CST').
//
// Example:
//  Time String:  2019-12-26 00:56:15 -0600 CST
//  Returned UTC Offset:  '-0600'
//  Returned Time Zone Abbreviation: 'CST'
//
func (dtMech *DTimeMechanics) GetUtcOffsetTzAbbrvFromDateTime(
	dateTime time.Time,
	ePrefix string) (
	utcOffset,
	tzAbbrv string,
	err error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	utcOffset = ""
	tzAbbrv = ""
	err = nil

	ePrefix += "DTimeMechanics.GetUtcOffsetTzAbbrvFromDateTime() "

	if dateTime.IsZero() {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO!\n")

		return utcOffset, tzAbbrv, err
	}

	tStr := dateTime.Format("2006-01-02 15:04:05 -0700 MST")

	lenLeadTzAbbrvStr := len("2006-01-02 15:04:05 -0700 ")

	tzAbbrv = tStr[lenLeadTzAbbrvStr:]

	tzAbbrv = strings.TrimRight(tzAbbrv, " ")

	lenLeadOffsetStr := len("2006-01-02 15:04:05 ")

	utcOffset = tStr[lenLeadOffsetStr : lenLeadOffsetStr+5]

	return utcOffset, tzAbbrv, err
}

// LoadTzLocation - Provides a single method for calling
// time.LoadLocation(). This method may be altered in the future
// to load time zones from an internal file thus affording
// consistency in time zone definitions without relying on
// zoneinfo.zip databases residing on host computers.
//
// If successful, this method returns a *time.Location or
// location pointer to a valid time zone.
//
func (dtMech *DTimeMechanics) LoadTzLocation(
	timeZoneName string,
	ePrefix string) (*time.Location, error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.LoadTzLocation() "

	if len(timeZoneName) == 0 {
		return nil,
			&TimeZoneError{
				ePrefix: ePrefix,
				errMsg:  "Input parameter 'timeZoneName' is a empty!",
				err:     nil,
			}
	}

	locPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		return nil,
			&TimeZoneError{
				ePrefix: ePrefix,
				errMsg:  fmt.Sprintf("Error returned by time.LoadLocation(timeZoneName)!\n" +
					"timeZoneName='%v'\nError='%v'\n", timeZoneName, err.Error()),
				err:     nil,
			}
	}

	return locPtr, nil
}

// ReclassifyTimeWithNewTz - Receives a valid time (time.Time)
// value and changes the existing time zone to that specified
// in parameter 'tZoneLocationName'.
//
// Input Parameters
// ================
//
//   dateTime time.Time
//          - Initial time whose time zone will be changed to
//            target time zone input parameter, 'tZoneLocationName'
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
func (dtMech *DTimeMechanics) ReclassifyTimeWithNewTz(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	tZoneLocationName string) (time.Time, error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix := "DTimeMechanics.ReclassifyTimeWithNewTz() "

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

// RelativeTimeToTimeNameZoneConversion - Converts a time value
// to its equivalent time in another time zone specified by input
// parameter string, 'timeZoneName'.
//
// The 'timeZoneName' string must specify one of three types of
// time zones:
//
//   (1) The string 'Local' - selects the local time zone
//                            location for the host computer.
//
//   (2) IANA Time Zone Location -
//      See https://golang.org/pkg/time/#LoadLocation
//      and https://www.iana.org/time-zones to ensure that
//      the IANA Time Zone Database is properly configured
//      on your system. Note: IANA Time Zone Data base is
//      equivalent to 'tz database'.
//     Examples:
//      "America/New_York"
//      "America/Chicago"
//      "America/Denver"
//      "America/Los_Angeles"
//      "Pacific/Honolulu"
//      "Etc/UTC" = GMT or UTC
//
//    (3) A Military Time Zone
//        Reference:
//         https://en.wikipedia.org/wiki/List_of_military_time_zones
//         http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//         https://www.timeanddate.com/time/zones/military
//         https://www.timeanddate.com/worldclock/timezone/alpha
//         https://www.timeanddate.com/time/map/
//
//        Examples:
//          "Alpha"   or "A"
//          "Bravo"   or "B"
//          "Charlie" or "C"
//          "Delta"   or "D"
//          "Zulu"    or "Z"
//
func (dtMech *DTimeMechanics) RelativeTimeToTimeNameZoneConversion(
	dateTime time.Time,
	timeZoneName string,
	ePrefix string) (time.Time, error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "DTimeMechanics.RelativeTimeToTimeNameZoneConversion() "

	if dateTime.IsZero() {
		return time.Time{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "dateTime",
				inputParameterValue: "",
				errMsg:              "Error: Input parameter 'dateTime' is zero!",
				err:                 nil,
			}
	}

	timeZoneName = strings.TrimRight(strings.TrimLeft(timeZoneName, " "), " ")

	if len(timeZoneName) == 0 {
		return time.Time{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeZoneName",
				inputParameterValue: "",
				errMsg:              "Error: 'timeZoneName' is an empty string!",
				err:                 nil,
			}
	}

	var err error
	tzSpec := TimeZoneSpecification{}
	tzMech := TimeZoneMechanics{}

	tzSpec,
	err = tzMech.GetTimeZoneFromName(
		dateTime,
		timeZoneName,
		TzConvertType.Relative(),
		ePrefix)

	if err != nil {
		return time.Time{}, err
	}

	return tzSpec.referenceDateTime, nil
}
