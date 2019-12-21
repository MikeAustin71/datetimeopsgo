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
// 'TimeZoneDefDto'.
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
// timeZoneDefDto TimeZoneDefDto  - A properly initialized 'TimeZoneDto'
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
	timeZoneDefDto TimeZoneDefDto) (time.Time, error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix := "DTimeUtility.AbsoluteTimeToTimeZoneDtoConversion() "

	if dateTime.IsZero() {
		return time.Time{},
			errors.New(ePrefix +
				"\nError: Input parameter 'dateTime' is zero!")
	}

	if !timeZoneDefDto.IsValid() {
		return time.Time{},
			errors.New(ePrefix +
				"Input parameter 'timeZoneDefDto' is Invalid!\n")
	}

	return time.Date(dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		dateTime.Hour(),
		dateTime.Minute(),
		dateTime.Second(),
		dateTime.Nanosecond(),
		timeZoneDefDto.GetLocationPtr()), nil
}

// AbsoluteTimeToTimeZoneDtoConversion - Converts a given time to
// another time zone using the 'Absolute' conversion method.
// This means that the years, months, days, hours, minutes,
// seconds and nanoseconds of the original 'dateTime' are not
// changed. That time value is simply assigned to another
// designated time zone. The target time zone is derived from
// input parameter 'timeZoneDefDto', an instance of type
// 'TimeZoneDefDto'.
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
	timeZoneName string) (time.Time, error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix := "DTimeUtility.AbsoluteTimeToTimeZoneNameConversion() "

	if dateTime.IsZero() {
		return time.Time{},
			errors.New(ePrefix +
				"\nError: Input parameter 'dateTime' is zero!")
	}

	dt2Util := DTimeUtility{}

	_,
	_,
	_,
	tzLocPtr,
	_,
	err := dt2Util.GetTimeZoneFromName(timeZoneName, ePrefix)


	if err != nil {
		return time.Time{}, err
	}

	return time.Date(dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		dateTime.Hour(),
		dateTime.Minute(),
		dateTime.Second(),
		dateTime.Nanosecond(),
		tzLocPtr), nil
}

// GetTimeZoneFromName - Analyzes a time zone name passed
// through input parameter,
func (dtUtil *DTimeUtility) GetTimeZoneFromName(
	timeZoneName string,
	ePrefix string) (milTzLetter,
	milTzName,
	ianaTzName string,
	ianaLocationPtr *time.Location,
	tzType TimeZoneType,
	err error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.GetTimeZoneFromName() "

	milTzLetter = ""
	milTzName = ""
	ianaTzName = ""
	ianaLocationPtr = nil
	tzType = TzType.None()
	err = nil

	var err2 error

	timeZoneName =
		strings.TrimLeft(strings.TrimRight(timeZoneName, " "), " ")

	if len(timeZoneName) == 0 {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'timeZoneName' is an empty string!\n")

		return milTzLetter,
			milTzName,
			ianaTzName,
			ianaLocationPtr,
			tzType,
			err
	}

	if len(timeZoneName) == 1 {

		timeZoneName = strings.ToUpper(timeZoneName)

		dtUtil2 := DTimeUtility{}

		milTzLetter,
			milTzName,
			ianaTzName,
			ianaLocationPtr,
			err =
			dtUtil2.ParseMilitaryTzNameAndLetter(timeZoneName, ePrefix)

		if err != nil {
			return milTzLetter,
				milTzName,
				ianaTzName,
				ianaLocationPtr,
				tzType,
				err
		}

		tzType = TzType.Military()

		return milTzLetter,
			milTzName,
			ianaTzName,
			ianaLocationPtr,
			tzType,
			err
	}

	testTzName := strings.ToLower(timeZoneName)
	tzType = TzType.Iana()

	if testTzName == "utc" {
		timeZoneName = TZones.UTC()

	} else if testTzName == "uct" {
		timeZoneName = TZones.UCT()
	} else if testTzName == "gmt" {
		timeZoneName = TZones.Other.GMT()
	} else if testTzName == "local" {
		timeZoneName = TZones.Local()
		tzType = TzType.Local()
	}

	ianaLocationPtr, err2 = time.LoadLocation(timeZoneName)

	if err2 == nil {
		milTzLetter = ""
		milTzName = ""
		ianaTzName = timeZoneName
		err = nil

		return milTzLetter,
			milTzName,
			ianaTzName,
			ianaLocationPtr,
			tzType,
			err
	}

	tzType = TzType.Military()

	err = nil

	dtUtil3 := DTimeUtility{}

	milTzLetter,
		milTzName,
		ianaTzName,
		ianaLocationPtr,
		err2 =
		dtUtil3.ParseMilitaryTzNameAndLetter(timeZoneName, ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"\nError: The input parameter 'timeZoneName' is invalid!\n"+
			"It could NOT be classified as either an IANA, Local or\n"+
			"Military time zone.\n"+
			"timeZoneName='%v'\n", timeZoneName)

		milTzLetter = ""
		milTzName = ""
		ianaTzName = ""
		ianaLocationPtr = nil
		tzType = TzType.None()
	}

	return milTzLetter,
		milTzName,
		ianaTzName,
		ianaLocationPtr,
		tzType,
		err
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
//          "Alpha"   or A
//          "Bravo"   or B
//          "Charlie" or C
//          "Delta"   or D
//          "Zulu"    or Z
//
func (dtUtil *DTimeUtility) RelativeTimeToTimeNameZoneConversion(
	dateTime time.Time,
	timeZoneName string,
	ePrefix string) (time.Time, error) {

	dtUtil.lock.Lock()

	defer dtUtil.lock.Unlock()

	ePrefix += "DTimeUtility.RelativeTimeToTimeNameZoneConversion() "

	dt2Util := DTimeUtility{}

	_,
	_,
	_,
	tzLocPtr,
	_,
	err := dt2Util.GetTimeZoneFromName(timeZoneName, ePrefix)

	if err != nil {
		return time.Time{}, err
	}

	return dateTime.In(tzLocPtr), nil
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
		return TZones.UTC()
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
		err = errors.New(ePrefix +
			"Error: Input Parameter 'rawTz' is EMPTY!\n")
		return milTzLetter,
			milTzName,
			equivalentIanaTimeZone,
			equivalentIanaLocationPtr,
			err
	}

	var ok bool
	milTzData := MilitaryTimeZoneData{}

	if lMilTz == 1 {

		milTzLetter = strings.ToUpper(rawTz)

		milTzName, ok =
			milTzData.MilTzLetterToTextName(milTzLetter)

		if !ok {
			err = fmt.Errorf(ePrefix+
				"Error: Input Parameter Value 'militaryTz' is INVALID!\n"+
				"'militaryTz' DOES NOT map to a valid Military Time Zone.\n"+
				"militaryTz='%v'", milTzLetter)

			milTzLetter = ""
			milTzName = ""
			equivalentIanaTimeZone = ""
			equivalentIanaLocationPtr = nil

			return milTzLetter,
				milTzName,
				equivalentIanaTimeZone,
				equivalentIanaLocationPtr,
				err
		}

		equivalentIanaTimeZone, ok = milTzData.MilitaryTzToIanaTz(milTzName)

		if !ok {
			err = fmt.Errorf(ePrefix+
				"Error: Input Parameter Value 'rawTz' is INVALID!\n"+
				"'rawTz' DOES NOT map to a valid IANA Time Zone.\n"+
				"rawTz='%v'", milTzName)

			milTzLetter = ""
			milTzName = ""
			equivalentIanaTimeZone = ""
			equivalentIanaLocationPtr = nil

			return milTzLetter,
				milTzName,
				equivalentIanaTimeZone,
				equivalentIanaLocationPtr,
				err
		}

	} else {
		// lMilTz > 1
		temp1 := rawTz[:1]
		temp2 := rawTz[1:]

		temp1 = strings.ToUpper(temp1)
		temp2 = strings.ToLower(temp2)

		milTzLetter = temp1
		milTzName = temp1 + temp2

		equivalentIanaTimeZone, ok = milTzData.MilitaryTzToIanaTz(milTzName)

		if !ok {
			err = fmt.Errorf(ePrefix+
				"Error: Input Parameter Value 'rawTz' is INVALID!\n"+
				"'rawTz' DOES NOT map to a valid IANA Time Zone.\n"+
				"Military Time Zone Letter='%v'\n"+
				"Military Time Zone Text Name='%v'", milTzLetter, milTzName)

			milTzLetter = ""
			milTzName = ""
			equivalentIanaTimeZone = ""
			equivalentIanaLocationPtr = nil

			return milTzLetter,
				milTzName,
				equivalentIanaTimeZone,
				equivalentIanaLocationPtr,
				err
		}
	}

	var err2 error
	err = nil
	equivalentIanaLocationPtr, err2 = time.LoadLocation(equivalentIanaTimeZone)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"\nError: Input parameter 'timeZoneName' was classified as a Military Time Zone.\n"+
			"However, the equivalent IANA Time Zone Name failed to return a Location Pointer.\n"+
			"timeZoneName='%v'\n"+
			"Military Time Zone Letter     ='%v'\n"+
			"Military Time Zone Name       ='%v'\n"+
			"Equivalent IANA Time Zone Name='%v'\n"+
			"Load Location Error='%v'\n",
			milTzLetter,
			milTzName,
			equivalentIanaTimeZone,
			err2.Error())

		milTzLetter = ""
		milTzName = ""
		equivalentIanaTimeZone = ""
		equivalentIanaLocationPtr = nil
	}

	return milTzLetter,
		milTzName,
		equivalentIanaTimeZone,
		equivalentIanaLocationPtr,
		err
}
