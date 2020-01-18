package datetime

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type TimeZoneMechanics struct {
	lock sync.Mutex
}

// AllocateOffsetSeconds - Designed to calculate offset hours,
// minutes and seconds from UTC+0000. A total signed seconds
// integer value is passed as an input parameter. This method
// then breaks down the total seconds into hours, minutes and
// seconds as positive integer values. The sign of the hours,
// minutes and seconds is returned in the 'sign' parameter as
// either a value +1, or -1.
//
func (tzMech *TimeZoneMechanics) AllocateOffsetSeconds(
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


// CalcConvertibleTimeZoneStats - Receives and examines a date time
// value to determine if the associated time zone is convertible
// across other time zones.
//
func (tzMech *TimeZoneMechanics) CalcConvertibleTimeZoneStats(
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

	ePrefix += "TimeZoneMechanics.CalcConvertibleTimeZoneStats() "

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

	tzMech2 := TimeZoneMechanics{}

	var tzAbbrvLookUpId  string
	var tzSpec TimeZoneSpecification

	tzAbbrvLookUpId, err =
		tzMech2.GetTzAbbrvLookupIdFromDateTime(
			dateTime, ePrefix)

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

	tzSpec,
	err =
	 tzMech2.ConvertTzAbbreviationToTimeZone(
	 	tInputJune,
	 	TzConvertType.Absolute(),
	 	tzAbbrvLookUpId,
	 	"",
	 	ePrefix)

	 if err != nil {
		 return tzIsConvertible,
			 convertibleDateTime,
			 err
	 }

	tInputJune = time.Date(
		dateTime.Year(),
		time.Month(6),
		15,
		11,
		0,
		0,
		0,
		tzSpec.GetLocationPointer())


	tInputDec := time.Date(
		dateTime.Year(),
		time.Month(12),
		15,
		11,
		0,
		0,
		0,
		tzSpec.GetLocationPointer())

	tLookupJune := time.Date(
		dateTime.Year(),
		time.Month(6),
		15,
		11,
		0,
		0,
		0,
		tzSpec.GetLocationPointer())


	tLookupDec := time.Date(
		dateTime.Year(),
		time.Month(12),
		15,
		11,
		0,
		0,
		0,
		tzSpec.GetLocationPointer())

	tLookupActual := time.Date(
		dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		dateTime.Hour(),
		dateTime.Minute(),
		dateTime.Second(),
		dateTime.Nanosecond(),
		tzSpec.GetLocationPointer())

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

// CalcUtcZoneOffsets - Receives an input parameter, 'dateTime',
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
//  zoneAbbrvLookupId string - Zone Abbreviations are used by other methods
//                             key values for map lookups. The Zone Abbreviation
//                             return value is formatted as follows:
//                             "CST-0600", "EET+0200"
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
func (tzMech *TimeZoneMechanics) CalcUtcZoneOffsets(
	dateTime time.Time,
	ePrefix string) (
	zoneName string,
	zoneOffset string,
	utcOffset string,
	zoneAbbrvLookupId string,
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

	ePrefix += "TimeZoneMechanics.CalcUtcZoneOffsets() "

	zoneName = ""
	zoneOffset = ""
	utcOffset = ""
	zoneAbbrvLookupId = ""
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
			zoneAbbrvLookupId,
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
			zoneAbbrvLookupId,
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
	zoneAbbrvLookupId = zoneName + zoneOffset

	if offsetSeconds > 0 {
		zoneOffset += fmt.Sprintf("%02d", offsetSeconds)
	}

	// Generates final ZoneOffset in the form
	// "-0500 CST" or "+0200 EET"
	zoneOffset += " " + zoneName

	return zoneName,
		zoneOffset,
		utcOffset,
		zoneAbbrvLookupId,
		offsetHours,
		offsetMinutes,
		offsetSeconds,
		offsetSignValue,
		zoneOffsetTotalSeconds,
		locationPtr,
		locationName,
		err
}

// ConvertTzAbbreviationToTimeZone - receives an input parameter,
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
func (tzMech *TimeZoneMechanics) ConvertTzAbbreviationToTimeZone(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	tzAbbrvLookupKey string,
	timeZoneLabel string,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	milTzLetter := ""
	milTzName := ""
	ianaTimeZoneName := ""
	var ianaLocationPtr *time.Location
	tzSpec = TimeZoneSpecification{}
	err = nil

	ePrefix += "dateTimeMechanics.ConvertTzAbbreviationToTimeZone() "

	if len(tzAbbrvLookupKey) == 0 {
		err = &InputParameterError{
			ePrefix:            ePrefix,
			inputParameterName: tzAbbrvLookupKey,
			errMsg:             "tzAbbrvLookKey is a zero length string!",
			err:                nil}

		return tzSpec, err
	}

	firstLtr := tzAbbrvLookupKey[0:1]

	if firstLtr == "+" || firstLtr == "-" {
		if len(tzAbbrvLookupKey) == 10 {
			if tzAbbrvLookupKey[3:5] == "00" {
				newAbbrv :=
					tzAbbrvLookupKey[0:3] + tzAbbrvLookupKey[5:10]
				tzAbbrvLookupKey = newAbbrv
			}
		}
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

		return tzSpec, err
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
		return tzSpec, err
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

		return tzSpec, err
	}

	timeZoneType := TzType.Iana()

	if tzAbbrRef.Location == "Military" {

		milTzLetter = tzAbbrRef.Abbrv
		milTzName = tzAbbrRef.AbbrvDescription
		timeZoneType = TzType.Military()
	}

	dtMech2 := dateTimeMechanics{}
	var err2 error

	//loadTzLocationPtr(
	if lenTZones == 1 {

		ianaLocationPtr, err2 = dtMech2.loadTzLocationPtr(tZones[0], ePrefix)

		if err2 != nil {

			err = fmt.Errorf(ePrefix +
				"\nAttempted loading of Time Zone %v Failed!\n" +
				"Error='%v'\n", tZones[0], err2.Error())

			return tzSpec, err
		}

	if timeConversionType == TzConvertType.Absolute() {
		dateTime = time.Date(
			dateTime.Year(),
			dateTime.Month(),
			dateTime.Day(),
			dateTime.Hour(),
			dateTime.Minute(),
			dateTime.Second(),
			dateTime.Nanosecond(),
			ianaLocationPtr)
	} else {
		dateTime = dateTime.In(ianaLocationPtr)
	}

		tzSpec = TimeZoneSpecification{}

		err = tzSpec.SetTimeZone(
			dateTime,
			milTzLetter,
			milTzName,
			timeZoneLabel,
			"",
			LocNameType.ConvertibleTimeZone(),
			timeZoneType,
			TzClass.AlternateTimeZone(),
			ePrefix)

		return tzSpec, err
	}

	lockTzAbbrvToTimeZonePriorityList.Lock()
	defer lockTzAbbrvToTimeZonePriorityList.Unlock()
	tzPriority := 999999

for i:=0; i < lenTZones; i++ {

	for j := 0; j < lenTzAbbrvToTimeZonePriorityList; j++ {

		if strings.HasPrefix(tZones[i], tzAbbrvToTimeZonePriorityList[j]) {
			if j < tzPriority {
				tzPriority = j
				ianaTimeZoneName = tZones[i]
			}
		}
	}
}

	if len(ianaTimeZoneName) == 0 {
		ianaTimeZoneName = tZones[0]
	}

	ianaLocationPtr, err2 = dtMech2.loadTzLocationPtr(ianaTimeZoneName, ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix +
			"\nAttempted Loading of Time Zone '%v' Failed!\n" +
			"Error='%v'\n", ianaTimeZoneName, err2.Error())

		return tzSpec, err
	}


	if timeConversionType == TzConvertType.Absolute() {
		dateTime = time.Date(
			dateTime.Year(),
			dateTime.Month(),
			dateTime.Day(),
			dateTime.Hour(),
			dateTime.Minute(),
			dateTime.Second(),
			dateTime.Nanosecond(),
			ianaLocationPtr)
	} else {
		dateTime = dateTime.In(ianaLocationPtr)
	}

	tzSpec = TimeZoneSpecification{}

	err = tzSpec.SetTimeZone(
		dateTime,
		milTzLetter,
		milTzName,
		timeZoneLabel,
		"",
		LocNameType.ConvertibleTimeZone(),
		timeZoneType,
		TzClass.AlternateTimeZone(),
		ePrefix)

	return tzSpec, err
}

// ConvertUtcAbbrvToStaticTz - Takes a UTC offset time zone
// abbreviation and converts to a static time zone name.
//
// If an equivalent 'Etc' static time zone is associated with
// this offset, it will be selected. Other wise the standard
// selection criteria will be applied to select a standard
// Time Zone Name.
//
// Examples:
//
// UTC Offset        Selected Time Zone
// ----------        ------------------
//  +10                Etc/GMT-10
//  +1030              Australia/Lord_Howe
//  -10                Etc/GMT+10
//
func (tzMech *TimeZoneMechanics) ConvertUtcAbbrvToStaticTz(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	timeZoneLabel string,
	utcOffsetAbbrv string,
	ePrefix string) (
	staticTimeZone TimeZoneSpecification,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	ePrefix += "TimeZoneMechanics.ConvertUtcAbbrvToStaticTz() "

	err = nil
	staticTimeZone = TimeZoneSpecification{}

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Input parameter 'dateTime' " +
				"has a ZERO value!",
			err:                 nil,
		}

		return staticTimeZone, err
	}

	if timeConversionType != TzConvertType.Relative() &&
		timeConversionType != TzConvertType.Absolute() {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: timeConversionType.String(),
			errMsg:              "Input parameter 'timeConversionType' value is Invalid!",
			err:                 nil,
		}

		return staticTimeZone, err
	}

	utcOffsetAbbrv =
		strings.TrimLeft(strings.TrimRight(utcOffsetAbbrv, " "), " ")

	lenUtcOffsetAbbrv := len(utcOffsetAbbrv)

	if lenUtcOffsetAbbrv != 3 &&
			lenUtcOffsetAbbrv != 5 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "utcOffsetAbbrv",
			inputParameterValue: utcOffsetAbbrv,
			errMsg: "Input parameter 'utcOffsetAbbrv' " +
				"has an invalid string length!",
			err: nil,
		}
		return staticTimeZone, err
	}

	// lenUtcOffsetAbbrv length must be 3 or 5
	firstLetterSignChar := utcOffsetAbbrv[0:1]

	if firstLetterSignChar != "+" &&
		firstLetterSignChar != "-" {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "utcOffsetAbbrv",
			inputParameterValue: utcOffsetAbbrv,
			errMsg: "Input parameter 'utcOffsetAbbrv' " +
				"does NOT qualify as a UTC Offset time zone abbreviation!",
			err: nil,
		}
		return staticTimeZone, err
	}

	dtMech := dateTimeMechanics{}

	if utcOffsetAbbrv == "+00"    ||
			utcOffsetAbbrv == "-00"   ||
			utcOffsetAbbrv == "+0000" ||
			utcOffsetAbbrv == "-0000" {

		// TZones.Etc.UTC()
		locPtr, err2 := dtMech.loadTzLocationPtr(TZones.Etc.UTC(),ePrefix)

		if err2 != nil {
			err = fmt.Errorf(ePrefix + "\n" +
				"Attempted UTC Time Zone Load Failed!\n" +
				"Error: %v", err2.Error())

			return staticTimeZone, err
		}

		if timeConversionType == TzConvertType.Absolute() {
			dateTime = time.Date(
				dateTime.Year(),
				dateTime.Month(),
				dateTime.Day(),
				dateTime.Hour(),
				dateTime.Minute(),
				dateTime.Second(),
				dateTime.Nanosecond(),
				locPtr)
		} else {

			dateTime = dateTime.In(locPtr)
		}

		staticTimeZone = TimeZoneSpecification{}

		err = staticTimeZone.SetTimeZone(
			dateTime,
			"",
			"",
			timeZoneLabel,
			"",
			LocNameType.ConvertibleTimeZone(),
			TzType.Iana(),
			TzClass.AlternateTimeZone(),
			ePrefix)

		return staticTimeZone, err
	}

	var utcOffsetAbbrv2 string

	switch lenUtcOffsetAbbrv {

	case 3:
		utcOffsetAbbrv2 = utcOffsetAbbrv + "00"

	case 5:

		minutes:= utcOffsetAbbrv[3:5]

		if minutes == "00" {
			utcOffsetAbbrv2 = utcOffsetAbbrv
			utcOffsetAbbrv = utcOffsetAbbrv[0:3]
		} else {
			utcOffsetAbbrv2 = utcOffsetAbbrv
		}

	default:
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "utcOffsetAbbrv",
			inputParameterValue: utcOffsetAbbrv,
			errMsg: "Input parameter 'utcOffsetAbbrv' " +
				"has an invalid string length!\n",
			err: nil,
		}
		return staticTimeZone, err
	}

	utcOffsetLookUpId := utcOffsetAbbrv + utcOffsetAbbrv2

	stdAbbrvs := StdTZoneAbbreviations{}

	tZones, ok := stdAbbrvs.AbbrvOffsetToTimeZones(utcOffsetLookUpId)

	if !ok {
		ePrefix += "StdTZoneAbbreviations.AbbrvOffsetToTimeZones() "
		err = &TzAbbrvMapLookupError{
			ePrefix:  ePrefix,
			mapName:  "mapTzAbbrvsToTimeZones",
			lookUpId: utcOffsetLookUpId,
			errMsg:   "",
			err:      nil,
		}

		return staticTimeZone, err
	}

	timeZoneName := ""

	lenTZones := len(tZones)
	// First look for a static time zone beginning with 'Etc'
	for i:=0; i < lenTZones; i++ {
		if strings.HasPrefix(tZones[i], "Etc") {
			timeZoneName = tZones[i]
			break
		}
	}

	if len(timeZoneName) == 0 {
		// No 'Etc' zone was found. Now look for a
		// a time zone name.
		lockTzAbbrvToTimeZonePriorityList.Lock()
		defer lockTzAbbrvToTimeZonePriorityList.Unlock()
		tzPriority := 999999

		for i:=0; i < lenTZones; i++ {

			for j := 0; j < lenTzAbbrvToTimeZonePriorityList; j++ {

				if strings.HasPrefix(tZones[i], tzAbbrvToTimeZonePriorityList[j]) {
					if j < tzPriority {
						tzPriority = j
						timeZoneName = tZones[i]
					}
				}
			}
		}
	}

	if len(timeZoneName) == 0 {
		timeZoneName = tZones[0]
	}

	if timeZoneName == "" {
		err = fmt.Errorf(ePrefix +
			"\nError: Could not locate equivalent 'Etc' " +
			"static time zone for UTC Time Zone abbreviation!\n" +
			"utcOffsetAbbrv='%v'\n", utcOffsetAbbrv)
		return staticTimeZone, err
	}

	locPtr, err2 := dtMech.loadTzLocationPtr(timeZoneName, ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n" +
			"Attempted Load of Time Zone '%v' Failed!\n" +
			"Error: %v",
			timeZoneName,
			err2.Error())

		return staticTimeZone, err
	}

	dateTime = time.Date(
		dateTime.Year(),
		dateTime.Month(),
		dateTime.Day(),
		dateTime.Hour(),
		dateTime.Minute(),
		dateTime.Second(),
		dateTime.Nanosecond(),
		locPtr)

	staticTimeZone = TimeZoneSpecification{}

	err = staticTimeZone.SetTimeZone(
		dateTime,
		"",
		"",
		timeZoneLabel,
		"",
		LocNameType.ConvertibleTimeZone(),
		TzType.Iana(),
		TzClass.AlternateTimeZone(),
		ePrefix)

	return staticTimeZone, err
}

// IsTzAbbrvUtcOffset - Returns a boolean value indicating
// whether the Time Zone Abbreviation is a UTC offset.
//
// Example UTC Offsets:
//   "+10", "+1000", "-05", "-0500"
//
// If the time zone abbreviation is a UTC offset, this
// method returns 'true'.
//
func (tzMech *TimeZoneMechanics) IsTzAbbrvUtcOffset(
	timeZoneAbbreviation string ) bool {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	lenTzAbbrvStr := len(timeZoneAbbreviation)

	if lenTzAbbrvStr == 0 {
		return false
	}

	firstChar := timeZoneAbbreviation[0:1]

	if firstChar == "+" || firstChar == "-" {
		return true
	}

	return false
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
// alternate, convertible time zone and the returned boolean value,
// 'isAlternateConvertibleTz', will be set to 'true'.
//
func (tzMech *TimeZoneMechanics) GetConvertibleTimeZoneFromDateTime(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	timeZoneLabel string,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	ePrefix += "TimeZoneMechanics.GetConvertibleTimeZoneFromDateTime() "

	ianaTimeZoneName := ""
	var ianaLocationPtr  *time.Location
	timeZoneType := TzType.None()
	tzSpec = TimeZoneSpecification{}
	err = nil

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'dateTime' is ZERO!",
			err:                 nil,
		}

		return tzSpec, err
	}

	if timeConversionType != TzConvertType.Relative() &&
		timeConversionType != TzConvertType.Absolute() {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: timeConversionType.String(),
			errMsg:              "Input parameter 'timeConversionType' value is Invalid!",
			err:                 nil,
		}

		return tzSpec, err
	}

	ianaLocationPtr = dateTime.Location()

	if ianaLocationPtr == nil {
		err = fmt.Errorf(ePrefix +
			"\nAttempt to load dateTime.Location() pointer FAILED!\n" +
			"Returned pointer is nil.\n" +
			"dateTime='%v'",dateTime.Format(FmtDateTimeTzNanoYMD))

		return TimeZoneSpecification{}, err
	}

	ianaTimeZoneName = ianaLocationPtr.String()

	dtMech := dateTimeMechanics{}
	var err2 error

	ianaLocationPtr, err2 = dtMech.loadTzLocationPtr(ianaTimeZoneName, ePrefix)

	if err2 == nil {

		if strings.ToLower(ianaTimeZoneName) == "local" {
			timeZoneType = TzType.Local()
		} else {
			timeZoneType = TzType.Iana()
		}

		tzSpec = TimeZoneSpecification{}
		err = tzSpec.SetTimeZone(
			dateTime,
			"",
			"",
			timeZoneLabel,
			"",
			LocNameType.ConvertibleTimeZone(),
			timeZoneType,
			TzClass.OriginalTimeZone(),
			ePrefix)

		return tzSpec, err
	}

	// The Original Time Zone failed to Load.
	// Try to extract a valid time zone
	// from the Time Zone Abbreviation.

	var utcOffset, tzAbbrv string

	utcOffset,
		tzAbbrv,
		err2 =
		dtMech.getUtcOffsetTzAbbrvFromDateTime(dateTime, ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix +
			"\nError: 'dateTime' Time Zone failed to load. Attempt to create look-up ID FAILED!\n" +
			"dateTime='%v'\n" +
			"Error='%v'\n",
			dateTime.Format("2006-01-02 15:04:05 -0700 MST"),
			err2.Error())

		return tzSpec, err
	}

	tzMech2 := TimeZoneMechanics{}

	if tzMech2.IsTzAbbrvUtcOffset(tzAbbrv) {

		return tzMech2.ConvertUtcAbbrvToStaticTz(
			dateTime,
			timeConversionType,
			timeZoneLabel,
			tzAbbrv,
			ePrefix)
	}

	tzAbbrv = tzAbbrv + utcOffset

	tzSpec,
	err =
			tzMech2.ConvertTzAbbreviationToTimeZone(
				dateTime,
				timeConversionType,
				tzAbbrv,
				timeZoneLabel,
				ePrefix)

	return tzSpec, err
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
func (tzMech *TimeZoneMechanics) GetTimeZoneFromName(
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	ePrefix += "TimeZoneMechanics.GetTimeZoneFromName() "

	milTzLetter := ""
	milTzName := ""
	var ianaLocationPtr *time.Location
	tzSpec = TimeZoneSpecification{}
	err = nil

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Input parameter 'dateTime' " +
				"has a ZERO value!",
			err:                 nil,
		}

		return tzSpec, err
	}

	if timeConversionType != TzConvertType.Relative() &&
		timeConversionType != TzConvertType.Absolute() {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: timeConversionType.String(),
			errMsg:              "Input parameter 'timeConversionType' value is Invalid!",
			err:                 nil,
		}

		return tzSpec, err
	}

	timeZoneName = strings.TrimRight(strings.TrimLeft(timeZoneName, " "), " ")

	if len(timeZoneName) == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "'timeZoneName' is an empty string!",
			err:                 nil,
		}

		return tzSpec, err
	}

	tzMech2 := TimeZoneMechanics{}

	var err2 error

	tzSpec,
	err = tzMech2.ParseMilitaryTzNameAndLetter(
		dateTime,
		timeConversionType,
		timeZoneName,
		ePrefix)

	if err == nil {
		// Loaded a Military Time Zone
		return tzSpec, err
	}

	// This is NOT a Military Time Zone. Try
	// a conventional Time Zone.
	tzSpec = TimeZoneSpecification{}
	err = nil

	dtMech := dateTimeMechanics{}

	ianaLocationPtr, err2 =
		dtMech.loadTzLocationPtr(timeZoneName,ePrefix)

	if err2 != nil {

		// The Time Zone failed to load!
		var utcOffset, tzAbbrv string

		utcOffset, tzAbbrv, err2 =
			tzMech2.GetUtcOffsetTzAbbrvFromDateTime(dateTime, ePrefix)

		if err2 != nil {

			err = fmt.Errorf(ePrefix + "\n" +
				"Load Location Failed. Error returned extracting UTC Offset, Tz Abreviation.\n" +
				"%v", err2.Error())

			return tzSpec, err
		}

		if !tzMech2.IsTzAbbrvUtcOffset(tzAbbrv) {
			err = fmt.Errorf(ePrefix +
				"\nError: Input parameter, 'timeZoneName', failed to load!\n" +
				"Therefore, 'timeZoneName is an invalid time zone." )
			return tzSpec, err
		}

		// This is a UTC offset
		tzSpec,
			err = tzMech2.ConvertUtcAbbrvToStaticTz(
			dateTime,
			TzConvertType.Absolute(),
			"Original Time Zone",
			tzAbbrv,
			ePrefix)

		if err == nil {
			return tzSpec, err
		}

		// An error occurred ConvertUtcAbbrvToStaticTz() Failed
		// Next, try ConvertTzAbbreviationToTimeZone
		tzSpec,
			err = tzMech2.ConvertTzAbbreviationToTimeZone(
			dateTime,
			TzConvertType.Absolute(),
			tzAbbrv + utcOffset,
			"Convertible Time Zone",
			ePrefix)

		if err != nil {
			return tzSpec, err
		}

		tzSpec.timeZoneClass = TzClass.AlternateTimeZone()
		return tzSpec, err
	}

	// The time zone loaded successfully!
	var tzType TimeZoneType

	if strings.ToLower(timeZoneName) == "local" {

		tzType = TzType.Local()

	} else {

		tzType = TzType.Iana()
	}

	if timeConversionType == TzConvertType.Absolute() {

		dateTime = time.Date(
			dateTime.Year(),
			dateTime.Month(),
			dateTime.Day(),
			dateTime.Hour(),
			dateTime.Minute(),
			dateTime.Second(),
			dateTime.Nanosecond(),
			ianaLocationPtr)
	} else {
		dateTime = dateTime.In(ianaLocationPtr)
	}

	err = tzSpec.SetTimeZone(
		dateTime,
		milTzLetter,
		milTzName,
		"Original Time Zone",
		"",
		LocNameType.ConvertibleTimeZone(),
		tzType,
		TzClass.OriginalTimeZone(),
		ePrefix)

	return tzSpec, err
}

// GetTzAbbrvLookupIdFromDateTime - Returns a Time Zone Abbreviation
// Lookup Id. This Time Zone Abbreviation Lookup Id is used to lookup
// alternative time zones from 'mapTzAbbrvsToTimeZones'.
//
func (tzMech *TimeZoneMechanics) GetTzAbbrvLookupIdFromDateTime(
	dateTime time.Time,
	ePrefix string) (
	tzAbbrvLookupId string,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	tzAbbrvLookupId = ""
	err = nil

	ePrefix += "TimeZoneMechanics.GetTzAbbrvLookupIdFromDateTime() "

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

// GetUtcOffsetTzAbbrvFromDateTime - Receives a time.Time, date
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
func (tzMech *TimeZoneMechanics) GetUtcOffsetTzAbbrvFromDateTime(
	dateTime time.Time,
	ePrefix string) (
	utcOffset,
	tzAbbrv string, err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	utcOffset = ""
	tzAbbrv = ""
	err = nil

	ePrefix += "TimeZoneMechanics.GetUtcOffsetTzAbbrvFromDateTime() "

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
func (tzMech *TimeZoneMechanics) ParseMilitaryTzNameAndLetter(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	timeZoneName string,
	ePrefix string) (
	tzSpec TimeZoneSpecification,
	err error) {

	tzMech.lock.Lock()

	defer tzMech.lock.Unlock()

	ePrefix += "TimeZoneMechanics.ParseMilitaryTzNameAndLetter() "

	milTzLetter := ""
	milTzName := ""
	equivalentIanaTimeZone := ""
	var equivalentIanaLocationPtr *time.Location
	err = nil
	tzSpec = TimeZoneSpecification{}

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Error: Input Parameter 'dateTime' has a ZERO value!",
			err:                 nil,
		}

		return tzSpec, err
	}

	if timeConversionType != TzConvertType.Relative() &&
		timeConversionType != TzConvertType.Absolute() {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: timeConversionType.String(),
			errMsg:              "Input parameter 'timeConversionType' value is Invalid!",
			err:                 nil,
		}

		return tzSpec, err
	}

	timeZoneName =
		strings.TrimLeft(strings.TrimLeft(timeZoneName, " "), " ")

	lMilTz := len(timeZoneName)

	if lMilTz == 0 {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "Error: Input Parameter 'timeZoneName' is empty string!",
			err:                 nil,
		}

		return tzSpec, err
	}

	var ok bool
	milTzData := MilitaryTimeZoneData{}

	if lMilTz == 1 {

		milTzLetter = strings.ToUpper(timeZoneName)

		milTzName, ok =
			milTzData.MilTzLetterToTextName(milTzLetter)

		if !ok {
			err = fmt.Errorf(ePrefix+
				"\nError: Input Parameter Value 'militaryTz' is INVALID!\n"+
				"'militaryTz' DOES NOT map to a valid Military Time Zone.\n"+
				"militaryTz='%v'", milTzLetter)

			return tzSpec, err
		}

		equivalentIanaTimeZone, ok = milTzData.MilitaryTzToIanaTz(milTzName)

		if !ok {
			err = fmt.Errorf(ePrefix+
				"Error: Input Parameter Value 'timeZoneName' is INVALID!\n"+
				"'timeZoneName' DOES NOT map to a valid IANA Time Zone.\n"+
				"timeZoneName='%v'", milTzName)

			return tzSpec, err
		}

	} else {
		// lMilTz > 1 - The length of the Military Time Zone Name
		// string is greater than '1'

		temp1 := timeZoneName[:1]
		temp2 := timeZoneName[1:]

		temp1 = strings.ToUpper(temp1)
		temp2 = strings.ToLower(temp2)

		milTzLetter = temp1
		milTzName = temp1 + temp2

		equivalentIanaTimeZone, ok = milTzData.MilitaryTzToIanaTz(milTzName)

		if !ok {
			err = fmt.Errorf(ePrefix+
				"Error: Input Parameter Value 'timeZoneName' is INVALID!\n"+
				"'timeZoneName' DOES NOT map to a valid IANA Time Zone.\n"+
				"Military Time Zone Letter='%v'\n"+
				"Military Time Zone Text Name='%v'", milTzLetter, milTzName)

			return tzSpec, err
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

		return tzSpec, err
	}

	if timeConversionType == TzConvertType.Absolute() {

		dateTime = time.Date(
			dateTime.Year(),
			dateTime.Month(),
			dateTime.Day(),
			dateTime.Hour(),
			dateTime.Minute(),
			dateTime.Second(),
			dateTime.Nanosecond(),
			equivalentIanaLocationPtr)

	} else {

		dateTime = dateTime.In(equivalentIanaLocationPtr)

	}

	err = tzSpec.SetTimeZone(
		dateTime,
		milTzLetter,
		milTzName,
		"Original Time Zone",
		"",
		LocNameType.ConvertibleTimeZone(),
		TzType.Military(),
		TzClass.OriginalTimeZone(),
		ePrefix)

	return tzSpec, err
}
