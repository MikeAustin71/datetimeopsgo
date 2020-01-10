package datetime

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type timeZoneMechanics struct {
	lock sync.Mutex
}

// allocateOffsetSeconds - Designed to calculate offset hours,
// minutes and seconds from UTC+0000. A total signed seconds
// integer value is passed as an input parameter. This method
// then breaks down the total seconds into hours, minutes and
// seconds as positive integer values. The sign of the hours,
// minutes and seconds is returned in the 'sign' parameter as
// either a value +1, or -1.
//
func (tzMech *timeZoneMechanics) allocateOffsetSeconds(
	signedTotalSeconds int) (
	hours,
	minutes,
	seconds,
	sign int) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

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

	seconds = signedTotalSeconds * sign

	hours = seconds / 3600

	seconds -= hours * 3600

	if seconds > 0 {
		minutes = seconds / 60
		seconds -= minutes * 60
	}

	return hours, minutes, seconds, sign
}


// calcConvertibleTimeZoneStats - Receives and examines a date time
// value to determine if the associated time zone is convertible
// across other time zones.
//
func (tzMech *timeZoneMechanics) calcConvertibleTimeZoneStats(
	dateTime time.Time,
	ePrefix string) (
	tzIsConvertible bool,
	convertibleDateTime time.Time,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	tzIsConvertible = false
	convertibleDateTime = time.Time{}
	err = nil

	ePrefix += "timeZoneMechanics.calcConvertibleTimeZoneStats() "

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "'dateTime' value is Zero!",
			err:                 nil,
		}

		return tzIsConvertible,
			convertibleDateTime,
			err
	}

	dateLocPtr := dateTime.Location()

	if dateLocPtr == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime.Location()",
			inputParameterValue: "",
			errMsg:              "dateTime Location Pointer is 'nil'",
			err:                 nil,
		}
		return tzIsConvertible,
			convertibleDateTime,
			err
	}

	dateTimeLocName := dateTime.Location().String()

	dateTimeLocName = strings.TrimLeft(strings.TrimRight(dateTimeLocName," "), " ")

	dateTimeLocName = strings.ToLower(dateTimeLocName)

	if dateTimeLocName == "local"{

		tzIsConvertible = true

		return tzIsConvertible,
			convertibleDateTime,
			err
	}

	tzMech2 := timeZoneMechanics{}

	var tzAbbrvLookUpId  string
	var ianaLocationPtr *time.Location

	tzAbbrvLookUpId, err =
		tzMech2.getTzAbbrvLookupIdFromDateTime(
			dateTime, ePrefix)

	if err != nil {
		return tzIsConvertible,
			convertibleDateTime,
			err
	}

	_,
	_,
	_,
	ianaLocationPtr,
	err =
	 tzMech2.convertTzAbbreviationToTimeZone(tzAbbrvLookUpId, ePrefix)

	 if err != nil {
		 return tzIsConvertible,
			 convertibleDateTime,
			 err
	 }

	tInputJune := time.Date(
		dateTime.Year(),
		time.Month(6),
		15,
		11,
		0,
		0,
		0,
		dateTime.Location())


	tInputDec := time.Date(
		dateTime.Year(),
		time.Month(12),
		15,
		11,
		0,
		0,
		0,
		dateTime.Location())

	tLookupJune := time.Date(
		dateTime.Year(),
		time.Month(6),
		15,
		11,
		0,
		0,
		0,
		ianaLocationPtr)


	tLookupDec := time.Date(
		dateTime.Year(),
		time.Month(12),
		15,
		11,
		0,
		0,
		0,
		ianaLocationPtr)

	tLookupActual := time.Date(
		dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		dateTime.Hour(),
		dateTime.Minute(),
		dateTime.Second(),
		dateTime.Nanosecond(),
		ianaLocationPtr)

	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tInputActualStr := dateTime.Format(fmtStr)

	tInputJuneStr := tInputJune.Format(fmtStr)

	tInputDecStr := tInputDec.Format(fmtStr)

	tLookupJuneStr := tLookupJune.Format(fmtStr)

	tLookupDecStr := tLookupDec.Format(fmtStr)

	tLookupActualStr := tLookupActual.Format(fmtStr)

	tzIsConvertible = true

	if tInputActualStr != tLookupActualStr {
		tzIsConvertible = false
	} else if tInputJuneStr != tLookupJuneStr {
		tzIsConvertible = false
	} else if tInputDecStr != tLookupDecStr {
		tzIsConvertible = false
	}

	if !tzIsConvertible {
		convertibleDateTime = tLookupActual
	}


	return tzIsConvertible,
		convertibleDateTime,
		err
}

// calcUtcZoneOffsets - Receives an input parameter, 'dateTime',
// of type 'time.Time' and proceeds to extract and return time
// a variety of zone components and descriptions.
//
// Input Parameter
// ===============
//
//  dateTime   time.Time  - A date time value which will be analyzed
//                          to extract zone, location and offset
//                          components.
//
// Return Values
// =============
//
//  zoneName         string - The Zone Name which is actually the zone
//                            abbreviation. Examples:
//                               "CST", "EST", "CDT", "EDT"
//
//  zoneOffset       string - The Zone Offset consists of the UTC offset
//                            plus the zone name or abbreviation. Examples:
//                               "-0600 CST", "+0200 EET"
//
//  utcOffset        string - The UTC Offset presents the hours and minutes
//                            offset from UTC TIME. iT is returned as a
//                            5-character string formatted as follows:
//                               "-0400", "-0500", "+0500", "+1000"
//
//  zoneAbbrv        string - Zone Abbreviations are used by other methods
//                            key values for map lookups. The Zone Abbreviation
//                            return value is formatted as follows:
//                               "CST-0600", "EET+0200"
//
//  offsetHours      int    - A positive value indicating the number of hours
//                            offset from UTC. For the sign value of hours,
//                            minutes and seconds of offset, see return value,
//                            'offsetSignValue'.
//
//  offsetMinutes    int    - A positive value indicating the number of minutes
//                            offset from UTC. For the sign value of hours,
//                            minutes and seconds of offset, see return value,
//                            'offsetSignValue'.
//
//  offsetSeconds    int    - A positive value indicating the number of seconds
//                            offset from UTC. For the sign value of hours,
//                            minutes and seconds of offset, see return value,
//                            'offsetSignValue'.

//  offsetSignValue  int    - This value is either +1 or -1. +1 == East of UTC,
//                            -1 == West of UTC. This sign value is applied to
//                            offset hours, minutes and seconds.
//
//  zoneTotalSeconds int    - A positive or negative value indicating the total
//                            number of seconds offset from UTC. A positive value
//                            signals East of UTC and a negative values signals
//                            West of UTC.
//
//  locationPtr      *time.Location - Pointer to the time zone 'location'
//                                     specified by input parameter 'dateTime'.
//
//  locationName     string         - Contains the text name of the time zone
//                                    location specified by input parameter 'dateTime'
//
//  err              error          - If this method completes successfully,
//                                    this error value is set to 'nil'. Otherwise,
//                                    'err' is configured with an appropriate error
//                                    message.
//
func (tzMech *timeZoneMechanics) calcUtcZoneOffsets(
	dateTime time.Time,
	ePrefix string) (
	zoneName string,
	zoneOffset string,
	utcOffset string,
	zoneAbbrv string,
	offsetHours int,
	offsetMinutes int,
	offsetSeconds int,
	offsetSignValue int,
	zoneOffsetTotalSeconds int,
	locationPtr *time.Location,
	locationName string,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	ePrefix += "timeZoneMechanics.calcUtcZoneOffsets() "

	zoneName = ""
	zoneOffset = ""
	utcOffset = ""
	zoneAbbrv = ""
	offsetHours = 0
	offsetMinutes = 0
	offsetSeconds = 0
	offsetSignValue = 0
	zoneOffsetTotalSeconds = 0
	locationPtr = nil
	locationName = ""
	err = nil

	if dateTime.IsZero() {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "'dateTime' has a zero value.",
			err:                 nil,
		}

		return zoneName,
			zoneOffset,
			utcOffset,
			zoneAbbrv,
			offsetHours,
			offsetMinutes,
			offsetSeconds,
			offsetSignValue,
			zoneOffsetTotalSeconds,
			locationPtr,
			locationName,
			err
	}

	locationPtr = dateTime.Location()

	if locationPtr == nil {

		err = &TimeZoneError{
			ePrefix: ePrefix,
			errMsg: fmt.Sprintf("dateTime.Location() returned a nil Location Pointer!\n"+
				"dateTime='%v'\n", dateTime.Format(FmtDateTimeTzNanoYMD)),
			err: nil,
		}

		return zoneName,
			zoneOffset,
			utcOffset,
			zoneAbbrv,
			offsetHours,
			offsetMinutes,
			offsetSeconds,
			offsetSignValue,
			zoneOffsetTotalSeconds,
			locationPtr,
			locationName,
			err
	}

	locationName = locationPtr.String()

	zoneName, zoneOffsetTotalSeconds = dateTime.Zone()

	zoneName = strings.TrimRight(strings.TrimLeft(zoneName, " "), " ")

	offsetSignValue = 1

	if zoneOffsetTotalSeconds < 0 {
		offsetSignValue = -1
	}

	offsetSeconds = zoneOffsetTotalSeconds * offsetSignValue

	if offsetSeconds > 0 {
		offsetHours = offsetSeconds / 3600

		offsetSeconds -= offsetHours * 3600
	}

	if offsetSeconds > 0 {
		offsetMinutes = offsetSeconds / 60
		offsetSeconds -= offsetMinutes * 60
	}

	signStr := "+"

	if offsetSignValue == -1 {
		signStr = "-"
	}

	zoneOffset += fmt.Sprintf("%v%02d%02d",
		signStr, offsetHours, offsetMinutes)

	// Generates final UTC offset in the form
	// "-0500" or "+0200"
	utcOffset = zoneOffset

	// Generates final zone abbreviation in the
	// format "CST-0500" or " EET+0200"
	zoneAbbrv = zoneName + zoneOffset

	if offsetSeconds > 0 {
		zoneOffset += fmt.Sprintf("%02d", offsetSeconds)
	}

	// Generates final ZoneOffset in the form
	// "-0500 CST" or "+0200 EET"
	zoneOffset += " " + zoneName

	return zoneName,
		zoneOffset,
		utcOffset,
		zoneAbbrv,
		offsetHours,
		offsetMinutes,
		offsetSeconds,
		offsetSignValue,
		zoneOffsetTotalSeconds,
		locationPtr,
		locationName,
		err
}

// convertTzAbbreviationToTimeZone - receives an input parameter,
// 'tzAbbrvLookupKey' which is used to look up a time zone abbreviation
// and return an associated IANA Time Zone Name.
//
// The method uses the global variable, 'tzAbbrvToTimeZonePriorityList'
// to assign the IANA Time Zone in cases of multiple time zones
// associated with the Time Zone Abbreviation.
//
// The 'tzAbbrvLookupKey' is formatted the Time Zone Abbreviation
// followed by the UTC offsets as illustrated by the following
// examples:
//   "EDT-0400"
//   "EST-0500"
//   "CDT-0500"
//   "CST-0600"
//   "PDT-0700"
//   "PST-0800"
//
// The associated IANA Time Zone name is identified using the
// global variable 'mapTzAbbrvsToTimeZones' which is accessed
// through method StdTZoneAbbreviations{}.AbbrvOffsetToTimeZones().
//
// If an associated IANA Time Zone is not found the returned
// boolean value, 'isValidTzAbbreviation', is set to 'false'.
//
func (tzMech *timeZoneMechanics) convertTzAbbreviationToTimeZone(
	tzAbbrvLookupKey string,
	ePrefix string) (
	milTzLetter,
	milTzName,
	ianaTimeZoneName string,
	ianaLocationPtr *time.Location,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	milTzLetter = ""
	milTzName = ""
	ianaTimeZoneName = ""
	ianaLocationPtr = nil
	err = nil

	ePrefix += "dateTimeMechanics.convertTzAbbreviationToTimeZone() "

	if len(tzAbbrvLookupKey) == 0 {
		err = &InputParameterError{
			ePrefix:            ePrefix,
			inputParameterName: tzAbbrvLookupKey,
			errMsg:             "tzAbbrvLookKey is a zero length string!",
			err:                nil}

		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err

	}

	stdAbbrvs := StdTZoneAbbreviations{}

	tZones, ok := stdAbbrvs.AbbrvOffsetToTimeZones(tzAbbrvLookupKey)

	if !ok {
		ePrefix += "StdTZoneAbbreviations.AbbrvOffsetToTimeZones() "
		err = &TzAbbrvMapLookupError{
			ePrefix:  ePrefix,
			mapName:  "mapTzAbbrvsToTimeZones",
			lookUpId: tzAbbrvLookupKey,
			errMsg:   "",
			err:      nil,
		}
		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err
	}

	lenTZones := len(tZones)

	if lenTZones == 0 {
		err = &TzAbbrvMapLookupError{
			ePrefix:  ePrefix,
			mapName:  "mapTzAbbrvsToTimeZones",
			lookUpId: tzAbbrvLookupKey,
			errMsg:   "Map returned a zero length time zones string array!",
			err:      nil,
		}
		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err
	}

	var tzAbbrRef TimeZoneAbbreviationDto

	tzAbbrRef, ok = stdAbbrvs.AbbrvOffsetToTzReference(tzAbbrvLookupKey)

	if !ok {
		ePrefix += "StdTZoneAbbreviations.AbbrvOffsetToTzReference() "
		err = &TzAbbrvMapLookupError{
			ePrefix:  ePrefix,
			mapName:  "mapTzAbbreviationReference",
			lookUpId: tzAbbrvLookupKey,
			errMsg:   "",
			err:      nil,
		}
		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err
	}

	if tzAbbrRef.Location == "Military" {

		milTzLetter = tzAbbrRef.Abbrv
		milTzName = tzAbbrRef.AbbrvDescription

	}

	dtMech2 := dateTimeMechanics{}

	//loadTzLocationPtr(
	if lenTZones == 1 {

		ianaLocationPtr, err = dtMech2.loadTzLocationPtr(tZones[0], ePrefix)

		if err != nil {
			milTzLetter = ""
			milTzName = ""
			ianaLocationPtr = nil

			return milTzLetter,
				milTzName,
				ianaTimeZoneName,
				ianaLocationPtr,
				err
		}

		ianaTimeZoneName = tZones[0]

		return milTzLetter,
			milTzName,
			ianaTimeZoneName,
			ianaLocationPtr,
			err
	}

	lockTzAbbrvToTimeZonePriorityList.Lock()
	defer lockTzAbbrvToTimeZonePriorityList.Unlock()
tzPriority := 999999

for i:=0; i < lenTzAbbrvToTimeZonePriorityList; i++ {

	for j := 0; j < lenTZones; j++ {

		if strings.HasPrefix(tZones[j], tzAbbrvToTimeZonePriorityList[i]) {
			if i < tzPriority {
				tzPriority = i
				ianaTimeZoneName = tZones[j]
			}
		}
	}
}

	if len(ianaTimeZoneName) == 0 {
		ianaTimeZoneName = tZones[0]
	}

	ianaLocationPtr, err = dtMech2.loadTzLocationPtr(ianaTimeZoneName, ePrefix)

	if err != nil {
		milTzLetter = ""
		milTzName = ""
		ianaTimeZoneName = ""
		ianaLocationPtr = nil
	}

	return milTzLetter,
		milTzName,
		ianaTimeZoneName,
		ianaLocationPtr,
		err
}

// getTzAbbrvLookupIdFromDateTime - Returns a Time Zone Abbreviation
// Lookup Id. This Time Zone Abbreviation Lookup Id is used to lookup
// alternative time zones from 'mapTzAbbrvsToTimeZones'.
//
func (tzMech *timeZoneMechanics) getTzAbbrvLookupIdFromDateTime(
	dateTime time.Time,
	ePrefix string) (
	tzAbbrvLookupId string,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	tzAbbrvLookupId = ""
	err = nil

	ePrefix += "timeZoneMechanics.getTzAbbrvLookupIdFromDateTime() "

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:            ePrefix,
			inputParameterName: "dateTime",
			errMsg:             "dateTime is Zero!",
			err:                nil,
		}
		return tzAbbrvLookupId, err
	}

	tStr :=
		dateTime.Format("2006-01-02 15:04:05 -0700 MST")

	lenLeadOffsetStr := len("2006-01-02 15:04:05 ")

	tzAbbrvLookupId = tStr[len("2006-01-02 15:04:05 -0700 "):]

	tzAbbrvLookupId =
		strings.TrimLeft(strings.TrimRight(tzAbbrvLookupId, " "), " ")

	tzAbbrvLookupId = tzAbbrvLookupId + tStr[lenLeadOffsetStr : lenLeadOffsetStr+5]

	return tzAbbrvLookupId, err
}

// getTimeZoneFromName - Analyzes a time zone name passed
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
func (tzMech *timeZoneMechanics) getTimeZoneFromName(
	timeZoneName string,
	ePrefix string) (
	milTzLetter,
	milTzName,
	ianaTzName string,
	ianaLocationPtr *time.Location,
	tzType TimeZoneType,
	err error) {

	ePrefix += "timeZoneMechanics.getTimeZoneFromName() "

	milTzLetter = ""
	milTzName = ""
	ianaTzName = ""
	ianaLocationPtr = nil
	tzType = TzType.None()
	err = nil

	timeZoneName = strings.TrimRight(strings.TrimLeft(timeZoneName, " "), " ")

	if len(timeZoneName) == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "'timeZoneName' is an empty string!",
			err:                 nil,
		}

		return milTzLetter,
		milTzName,
		ianaTzName,
		ianaLocationPtr,
		tzType,
		err
	}

	tzMech2 := timeZoneMechanics{}

	milTzLetter,
	milTzName,
	ianaTzName,
	ianaLocationPtr,
	err  = tzMech2.parseMilitaryTzNameAndLetter(timeZoneName, ePrefix)

	if err == nil {
		tzType = TzType.Military()

		return milTzLetter,
			milTzName,
			ianaTzName,
			ianaLocationPtr,
			tzType,
			err
	}

	dtMech := dateTimeMechanics{}

	var err2 error

	ianaLocationPtr, err2 = dtMech.loadTzLocationPtr(timeZoneName,ePrefix)

	if err2!=nil {
		ianaLocationPtr = nil
		err = fmt.Errorf(ePrefix +
			"\nError: Input parameter, 'timeZoneName', failed to load!\n" +
			"Therefore, 'timeZoneName is an invalid time zone." )
	} else {

		ianaTzName = timeZoneName

		testTzName := strings.ToLower(ianaTzName)

		if testTzName == "local" {

			tzType = TzType.Local()

		} else {

			tzType = TzType.Iana()
		}
	}

	return milTzLetter,
		milTzName,
		ianaTzName,
		ianaLocationPtr,
		tzType,
		err
}

// getUtcOffsetTzAbbrvFromDateTime - Receives a time.Time, date
// time, input parameter and extracts and returns the
// 5-character UTC offset and the time zone abbreviation.
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
//
//  Returned UTC Offset:  '-0600'
//
//  Returned Time Zone Abbreviation: 'CST'
//
func (tzMech *timeZoneMechanics) getUtcOffsetTzAbbrvFromDateTime(
	dateTime time.Time,
	ePrefix string) (
	utcOffset,
	tzAbbrv string, err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	utcOffset = ""
	tzAbbrv = ""
	err = nil

	ePrefix += "timeZoneMechanics.getUtcOffsetTzAbbrvFromDateTime() "

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:            ePrefix,
			inputParameterName: "dateTime",
			errMsg:             "dateTime is Zero!",
			err:                nil,
		}
		return utcOffset, tzAbbrv, err
	}

	tStr :=
		dateTime.Format("2006-01-02 15:04:05 -0700 MST")

	tzAbbrv = tStr[len("2006-01-02 15:04:05 -0700 "):]

	tzAbbrv =
		strings.TrimLeft(strings.TrimRight(tzAbbrv, " "), " ")

	lenLeadOffsetStr := len("2006-01-02 15:04:05 ")

	utcOffset = tStr[lenLeadOffsetStr : lenLeadOffsetStr+5]

	return utcOffset, tzAbbrv, err
}


// parseMilitaryTzNameAndLetter - Parses a text string which
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
func (tzMech *timeZoneMechanics) parseMilitaryTzNameAndLetter(
	rawTz string,
	ePrefix string) (milTzLetter,
	milTzName,
	equivalentIanaTimeZone string,
	equivalentIanaLocationPtr *time.Location,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	ePrefix += "timeZoneMechanics.parseMilitaryTzNameAndLetter() "

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
	dtMech := dateTimeMechanics{}

	equivalentIanaLocationPtr, err2 =
		dtMech.loadTzLocationPtr(equivalentIanaTimeZone, ePrefix)

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
