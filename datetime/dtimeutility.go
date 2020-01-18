package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type DTimeUtility struct {
	lock sync.Mutex
}

// AbsoluteTimeToTimeZoneDtoConversion - Converts a given time to
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
// timeZoneDefDto TimeZoneDefinition  - A properly initialized 'TimeZoneDto'
//                                  encapsulating the time zone to which
//                                  'dateTime' will be converted.
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
func (dtUtil *DTimeUtility) AbsoluteTimeToTimeZoneDtoConversion(
	dateTime time.Time,
	timeZoneDefDto TimeZoneDefinition) (time.Time, error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix := "DTimeUtility.AbsoluteTimeToTimeZoneDtoConversion() "

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

// AbsoluteTimeToTimeZoneDtoConversion - Converts a given time to
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
func (dtUtil *DTimeUtility) AbsoluteTimeToTimeZoneNameConversion(
	dateTime time.Time,
	timeZoneName string,
	ePrefix string) (time.Time, error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.AbsoluteTimeToTimeZoneNameConversion() "

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

	dtMech := dateTimeMechanics{}
	return 	dtMech.absoluteTimeToTimeZoneNameConversion(
		dateTime,
		timeZoneName,
		ePrefix)
}

// ConsolidateErrors - Receives an array of errors and converts them
// to a single error which is returned to the caller. Multiple errors
// are separated by a new line character.
//
// If the length of the error array is zero, this method returns nil.
//
func (dtUtil *DTimeUtility) ConsolidateErrors(errs []error) error {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	lErrs := len(errs)

	if lErrs == 0 {
		return nil
	}

	errStr := ""

	for i := 0; i < lErrs; i++ {

		if errs[i] == nil {
			continue
		}

		tempStr := fmt.Sprintf("%v", errs[i].Error())

		tempStr = strings.TrimLeft(strings.TrimRight(tempStr, " "), " ")

		strLen := len(tempStr)

		for strings.HasSuffix(tempStr,"\n") &&
			strLen > 1 {

			tempStr = tempStr[0:strLen-1]
			strLen--
		}

		if i == (lErrs - 1) {
			errStr += fmt.Sprintf("%v", tempStr)
		} else if i == 0 {
			errStr = fmt.Sprintf("\n%v\n\n", tempStr)
		} else {
			errStr += fmt.Sprintf("%v\n\n", tempStr)
		}
	}

	return fmt.Errorf("%v", errStr)
}

// GetTimeZoneFromName - Analyzes a time zone name passed
// through input parameter, 'timeZoneName'. If valid, the
// method populates time zone description elements and
// returns them.
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
func (dtUtil *DTimeUtility) GetTimeZoneFromName(
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.GetTimeZoneFromName() "

	tzSpec = TimeZoneSpecification{}
	err = nil

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

// GetConvertibleTimeZoneFromDateTime - Receives a date time
// (type time.Time) as an input parameter. 'dateTime' is parsed
// and a valid, convertible time zone name and location pointer
// are returned.  Note: Due to the structure of 'dateTime', a
// military time zone is never returned. All returned time zones
// are either IANA time zones or the 'Local' time zone designated
// by golang and the host computer.
//
// If the initial time zone extracted from 'dateTime' is invalid,
// the date time time zone abbreviation will be used to look up an
// alternate, convertible time zone. Check the flags on the returned
// 'Time Zone Specification'.
//
func (dtUtil *DTimeUtility) GetConvertibleTimeZoneFromDateTime(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.GetConvertibleTimeZoneFromDateTime() "

	tzMech := TimeZoneMechanics{}

	return tzMech.GetConvertibleTimeZoneFromDateTime(
		dateTime,
		timeConversionType,
		"Convertible Time Zone",
		ePrefix)
}

// GetUtcOffsetTzAbbrvFromDateTime - Receives a time.Time, date
// time, input parameter and extracts and returns the
// 5-character UTC offset and UTC offset.
//
// UTC Offsets are returned in the format illustrated by the
// following examples:
//   +1030
//   -0500
//   +1100
//   -1100
//
// Example:
//  Time String:  2019-12-26 00:56:15 -0600 CST
//  Returned UTC Offset:  '-0600'
//  Returned Time Zone Abbreviation: 'CST'
//
func (dtUtil *DTimeUtility) GetUtcOffsetTzAbbrvFromDateTime(
	dateTime time.Time,
	ePrefix string) (utcOffset, tzAbbrv string, err error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.GetUtcOffsetTzAbbrvFromDateTime() "

	dtMech := dateTimeMechanics{}

	return dtMech.getUtcOffsetTzAbbrvFromDateTime(dateTime, ePrefix)
}

// LoadTzLocation - This is a wrapper method for time.LoadLocation.
// Using this method provides the option to load time zones from
// an internal zoneinfo.zip database in future upgrades.
//
func (dtUtil *DTimeUtility) LoadTzLocation(
	timeZoneName string,
	ePrefix string) (
	tzPointer *time.Location,
	err error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.LoadTzLocation() "

	tzPointer = nil
	err = nil

	if len(timeZoneName) == 0 {
		return tzPointer,
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeZoneName",
				inputParameterValue: "Input parameter 'timeZoneName' is an EMPTY string!",
				errMsg:              "",
				err:                 nil,
			}
	}

	dtMech := dateTimeMechanics{}

	return dtMech.loadTzLocationPtr(timeZoneName, ePrefix)
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
func (dtUtil *DTimeUtility) RelativeTimeToTimeNameZoneConversion(
	dateTime time.Time,
	timeZoneName string,
	ePrefix string) (time.Time, error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.RelativeTimeToTimeNameZoneConversion() "

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

	dtMech := dateTimeMechanics{}

	return dtMech.relativeTimeToTimeNameZoneConversion(
		dateTime,
		timeZoneName,
		ePrefix)
}

// PreProcessTimeZoneLocation - Scans a time zone location
// name string and attempts to correct errors. If input
// parameter 'timeZoneLocation' is an empty string, this
// method set 'timeZoneLocation' to 'UTC'.
//
func (dtUtil *DTimeUtility) PreProcessTimeZoneLocation(
	timeZoneLocation string) string {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	timeZoneLocation =
		strings.TrimLeft(strings.TrimRight(timeZoneLocation, " "), " ")

	if len(timeZoneLocation) == 0 {
		return timeZoneLocation
	}

	testZone := strings.ToLower(timeZoneLocation)

	if testZone == "utc" {

		timeZoneLocation = TZones.UTC()

	} else if testZone == "uct" {

		timeZoneLocation = TZones.UCT()

	} else if testZone == "gmt" {

		timeZoneLocation = TZones.UTC()

	} else if testZone == "local" {

		return TZones.Local()
	}

	return timeZoneLocation

}

// ParseMilitaryTzNameAndLetter - Parses a text string which
// contains either a single letter military time zone designation
// or a multi-character time zone text name.
//
// If successful, three populated strings are returned. The first
// is the valid Military Time Zone Letter designation. The second
// returned string contains the text name of the Military Time
// Zone. The third string contains the name of the equivalent
// IANA Time Zone. This is required because Golang does not
// currently support Military Time Zones.
//
// In addition to the three strings, a successful method completion
// will also return the equivalent IANA Time Zone Location pointer
// (*time.Location).
//
// If an error is encountered, the return value, 'err' is populated
// with an appropriate error message. Otherwise, 'err' is set
// equal to 'nil' signaling no error was encountered.
//
func (dtUtil *DTimeUtility) ParseMilitaryTzNameAndLetter(
	rawTz string,
	ePrefix string) (milTzLetter,
	milTzName,
	equivalentIanaTimeZone string,
	equivalentIanaLocationPtr *time.Location,
	err error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.ParseMilitaryTzNameAndLetter() "

	milTzLetter = ""
	milTzName = ""
	equivalentIanaTimeZone = ""
	equivalentIanaLocationPtr = nil
	err = nil

	rawTz =
		strings.TrimLeft(strings.TrimLeft(rawTz, " "), " ")

	lMilTz := len(rawTz)

	if lMilTz == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "rawTz",
			inputParameterValue: "",
			errMsg:              "Error: Input Parameter 'rawTz' is empty string!",
			err:                 nil,
		}

	return milTzLetter,
			milTzName,
			equivalentIanaTimeZone,
			equivalentIanaLocationPtr,
			err
	}

	tzMech := TimeZoneMechanics{}

	return tzMech.ParseMilitaryTzNameAndLetter(rawTz, ePrefix)
}
