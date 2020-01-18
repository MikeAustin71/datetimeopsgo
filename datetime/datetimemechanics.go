package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type dateTimeMechanics struct {
	lock sync.Mutex
}


// absoluteTimeToTimeZoneNameConversion - Converts a given time to
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
func (dtMech *dateTimeMechanics) absoluteTimeToTimeZoneNameConversion(
	dateTime time.Time,
	timeZoneName string,
	ePrefix string) (time.Time, error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "dateTimeMechanics.absoluteTimeToTimeZoneNameConversion() "

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

// allocateSecondsToHrsMinSecs - Useful in calculating offset hours,
// minutes and seconds from UTC+0000. A total signed seconds value
// is passed as an input parameter. This method then breaks down
// hours, minutes and seconds as positive integer values. The sign
// of the hours, minutes and seconds is returned in the 'sign'
// parameter as +1 or -1.
//
func (dtMech *dateTimeMechanics) allocateSecondsToHrsMinSecs(
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
func (dtMech *dateTimeMechanics) getUtcOffsetTzAbbrvFromDateTime(
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

	ePrefix += "dateTimeMechanics.GetUtcOffsetTzAbbrvFromDateTime() "

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

// loadTzLocationPtr - Provides a single method for calling
// time.LoadLocation(). This method may be altered in the future
// to load time zones from an internal file thus affording
// consistency in time zone definitions without relying on
// zoneinfo.zip databases residing on host computers.
//
// If successful, this method returns a *time.Location or
// location pointer to a valid time zone.
//
func (dtMech *dateTimeMechanics) loadTzLocationPtr(
	timeZoneName string,
	ePrefix string) (*time.Location, error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "dateTimeMechanics.loadTzLocationPtr() "

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


// relativeTimeToTimeNameZoneConversion - Converts a time value
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
func (dtMech *dateTimeMechanics) relativeTimeToTimeNameZoneConversion(
	dateTime time.Time,
	timeZoneName string,
	ePrefix string) (time.Time, error) {

	dtMech.lock.Lock()

	defer dtMech.lock.Unlock()

	ePrefix += "dateTimeMechanics.relativeTimeToTimeNameZoneConversion() "

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
