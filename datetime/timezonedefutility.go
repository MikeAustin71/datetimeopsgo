package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type timeZoneDefUtility struct {
	lock sync.Mutex
}


// allocateZoneOffsetSeconds - allocates a signed value of total offset seconds from
// UTC to the associated fields in the current TimeZoneDefDto instance.
func (tzDefUtil *timeZoneDefUtility) allocateZoneOffsetSeconds(
  tzdef *TimeZoneDefDto,
	signedZoneOffsetSeconds int) {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.allocateZoneOffsetSeconds()\n" +
			"Error: Input parameter 'tzdef' is a 'nil' pointer!\n")
	}

	if signedZoneOffsetSeconds < 0 {
		tzdef.zoneSign = -1
	} else {
		tzdef.zoneSign = 1
	}

	tzdef.zoneOffsetSeconds = signedZoneOffsetSeconds

	signedZoneOffsetSeconds *= tzdef.zoneSign

	tzdef.offsetHours = 0
	tzdef.offsetMinutes = 0
	tzdef.offsetSeconds = 0

	if signedZoneOffsetSeconds == 0 {
		return
	}

	tzdef.offsetHours = signedZoneOffsetSeconds / 3600 // compute hours
	signedZoneOffsetSeconds -= tzdef.offsetHours * 3600

	if signedZoneOffsetSeconds > 0 {
		tzdef.offsetMinutes = signedZoneOffsetSeconds / 60 // compute minutes
		signedZoneOffsetSeconds -= tzdef.offsetMinutes * 60
	}

	tzdef.offsetSeconds = signedZoneOffsetSeconds

	return
}

// CopyIn - Copies an incoming TimeZoneDefDto into the
// data fields of the current TimeZoneDefDto instance.
//
func (tzDefUtil *timeZoneDefUtility) copyIn(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.CopyIn()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.CopyIn()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	tzDefUtil2 := timeZoneDefUtility{}

	tzDefUtil2.empty(tzdef)

	tzdef.zoneName = tzdef2.zoneName
	tzdef.zoneOffsetSeconds = tzdef2.zoneOffsetSeconds
	tzdef.zoneSign = tzdef2.zoneSign
	tzdef.offsetHours = tzdef2.offsetHours
	tzdef.offsetMinutes = tzdef2.offsetMinutes
	tzdef.offsetSeconds = tzdef2.offsetSeconds
	tzdef.zoneOffset = tzdef2.zoneOffset
	tzdef.utcOffset = tzdef2.utcOffset
	tzdef.location = tzdef2.location
	tzdef.locationName = tzdef2.locationName
	tzdef.locationNameType = tzdef2.locationNameType
	tzdef.militaryTimeZoneLetter = tzdef2.militaryTimeZoneLetter
	tzdef.militaryTimeZoneName = tzdef2.militaryTimeZoneName
	tzdef.tagDescription = tzdef2.tagDescription
	tzdef.timeZoneType = tzdef2.timeZoneType

	return
}

// copyOut - creates and returns a deep copy of the current
// TimeZoneDefDto instance.
//
func (tzDefUtil *timeZoneDefUtility) copyOut(
	tzdef *TimeZoneDefDto) TimeZoneDefDto {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.copyOut()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	tzdef2 := TimeZoneDefDto{}

	tzdef2.zoneName = tzdef.zoneName
	tzdef2.zoneOffsetSeconds = tzdef.zoneOffsetSeconds
	tzdef2.zoneSign = tzdef.zoneSign
	tzdef2.offsetHours = tzdef.offsetHours
	tzdef2.offsetMinutes = tzdef.offsetMinutes
	tzdef2.offsetSeconds = tzdef.offsetSeconds
	tzdef2.zoneOffset = tzdef.zoneOffset
	tzdef2.utcOffset = tzdef.utcOffset

	tzdef2.location = tzdef.location
	tzdef2.locationName = tzdef.locationName
	tzdef2.locationNameType = tzdef.locationNameType
	tzdef2.militaryTimeZoneLetter = tzdef.militaryTimeZoneLetter
	tzdef2.militaryTimeZoneName = tzdef.militaryTimeZoneName

	tzdef2.tagDescription = tzdef.tagDescription
	tzdef2.timeZoneType = tzdef.timeZoneType

	return tzdef2
}

// Empty - Resets all field values for the input parameter
// TimeZoneDefDto to their uninitialized or 'zero' states.
//
func (tzDefUtil *timeZoneDefUtility) empty(
	tzdef *TimeZoneDefDto) {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.empty()\n" +
			"Error: 'tzdef' pointer is nil!\n")
	}

	tzdef.zoneName = ""
	tzdef.zoneOffsetSeconds = 0
	tzdef.zoneSign = 0
	tzdef.offsetHours = 0
	tzdef.offsetMinutes = 0
	tzdef.offsetSeconds = 0
	tzdef.zoneOffset = ""
	tzdef.utcOffset = ""
	tzdef.location = nil
	tzdef.locationName = ""
	tzdef.locationNameType = LocNameType.None()
	tzdef.militaryTimeZoneLetter = ""
	tzdef.militaryTimeZoneName = ""
	tzdef.tagDescription = ""
	tzdef.timeZoneType = TzType.None()

	return
}


// Equal - Determines if two TimeZoneDefDto instances are
// equivalent in value.
//
// This method returns 'true' of two TimeZoneDefDto's are
// equal in all respects.
//
func (tzDefUtil *timeZoneDefUtility) equal(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equal() " +
			"\nError: Input parameter 'tzdef' is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equal() " +
			"\nError: Input parameter 'tzdef2' is nil!")
	}

	tzDefUtil2 := timeZoneDefUtility{}

	tzdefIsEmpty := tzDefUtil2.isEmpty(tzdef)

	tzdef2IsEmpty := tzDefUtil2.isEmpty(tzdef2)

	if tzdefIsEmpty == true &&
		tzdef2IsEmpty == true {
		return true
	}

	if tzdef.zoneName != tzdef2.zoneName ||
		tzdef.zoneOffsetSeconds != tzdef2.zoneOffsetSeconds ||
		tzdef.zoneSign != tzdef2.zoneSign {
		return false
	}

	if tzdef.offsetHours != tzdef2.offsetHours ||
		tzdef.offsetMinutes != tzdef2.offsetMinutes ||
		tzdef.offsetSeconds != tzdef2.offsetSeconds {
		return false
	}

	if tzdef.zoneOffset != tzdef2.zoneOffset ||
		tzdef.utcOffset != tzdef2.utcOffset {
		return false
	}

	if tzdef.locationName != tzdef2.locationName ||
		tzdef.militaryTimeZoneLetter != tzdef2.militaryTimeZoneLetter ||
		tzdef.militaryTimeZoneName != tzdef2.militaryTimeZoneName ||
		tzdef.tagDescription != tzdef2.tagDescription {
		return false
	}

	if tzdef.timeZoneType != tzdef2.timeZoneType ||
		tzdef.locationNameType != tzdef2.locationNameType {
		return false
	}

	if tzdef.location != nil && tzdef2.location == nil ||
		tzdef.location == nil && tzdef2.location != nil ||
		tzdef.location.String() != tzdef2.location.String() {
		return false
	}

	return true
}

// equalLocations - Compares the Time Zone Locations for two TimeZoneDefDto's
// and returns 'true' if they are equal.
//
// Time Zone Location Name Examples:
//   "Local"
//   "America/Chicago"
//   "America/New_York"
//
func (tzDefUtil *timeZoneDefUtility) equalLocations(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	return tzdef.locationName == tzdef2.locationName
}

// equalOffsetSeconds - Compares Zone Offset Seconds for two TimeZoneDefDto's and
// returns 'true' if they are equal.
//
// ZoneOffsetSeconds is a signed number of seconds offset from UTC:
//   + == East of UTC
//   - == West of UTC
//
func (tzDefUtil *timeZoneDefUtility) equalOffsetSeconds(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	return tzdef.zoneOffsetSeconds == tzdef2.zoneOffsetSeconds
}

// equalZoneLocation - Compares two TimeZoneDefDto's and returns
// 'true' if Time Zone Location Name, the Zone Name and Zone
// Offsets match.
//
// Examples Of Time Zone Location Location Name:
//
//   "Local"
//   "America/Chicago"
//   "America/New_York"
//
// Examples of Zone Names:
//   "EST"
//   "CST"
//   "PST"
//
// Examples of Zone Offsets:
//   "-0600 CST"
//   "-0500 EST"
//   "+0200 EET"
//
func (tzDefUtil *timeZoneDefUtility) equalZoneLocation(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	if tzdef.locationName != tzdef2.locationName ||
	 tzdef.zoneName != tzdef2.zoneName ||
		tzdef.zoneOffset != tzdef2.zoneOffset {
		return false
	}

	return true
}

// equalZoneOffsets - Compares ZoneOffsets for two TimeZoneDefDto's and
// returns 'true' if they are equal.
//
// Zone Offset is a text string representing the offset from UTC plus the
// time zone abbreviation.
//
// Example "-0500 CDT"
//
func (tzDefUtil *timeZoneDefUtility) equalZoneOffsets(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equalZoneOffsets()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equalZoneOffsets()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	return tzdef.zoneOffset == tzdef2.zoneOffset
}

// isEmpty - Determines whether the current TimeZoneDefDto
// instance is Empty.
//
// If the TimeZoneDefDto instance (tzdef) is NOT populated,
// this method returns 'true'. Otherwise, it returns 'false'.
//
func (tzDefUtil *timeZoneDefUtility) isEmpty(
	tzdef *TimeZoneDefDto ) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.isValidFromDateTime()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef.zoneName != "" ||
		tzdef.zoneOffsetSeconds != 0 ||
		tzdef.zoneSign != 0 ||
		tzdef.offsetHours != 0 ||
		tzdef.offsetMinutes != 0 ||
		tzdef.offsetSeconds != 0 ||
		tzdef.zoneOffset != "" ||
		tzdef.utcOffset != "" ||
		tzdef.timeZoneType != TzType.None() ||
		tzdef.locationName != "" ||
		tzdef.militaryTimeZoneName != "" ||
		tzdef.militaryTimeZoneLetter != "" {
		return false
	}

	return true
}

// isValidTimeZoneDefDto - Analyzes the TimeZoneDefDto
// parameter, 'tzdef', instance to determine validity.
//
// This method returns 'true' if the TimeZoneDefDto
// instance is valid.  Otherwise, it returns 'false'.
//
func (tzDefUtil *timeZoneDefUtility) isValidTimeZoneDefDto(
	tzdef *TimeZoneDefDto,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.isValidTimeZoneDefDto() "

	if tzdef == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzdef' is a 'nil' pointer!\n")
	}

	tzDefUtil2 := timeZoneDefUtility{}

	if tzDefUtil2.isEmpty(tzdef) {
		return errors.New(ePrefix +
			"\nError: This TimeZoneDefDto instance is empty!\n")
	}

	if strings.TrimLeft(strings.TrimRight(tzdef.locationName, " "), " ") == "" {
		return errors.New(ePrefix +
			"\nError: tzdef.locationName is an empty string!\n")
	}

	if tzdef.location.String() != tzdef.locationName {
		return fmt.Errorf(ePrefix +
			"\nError: The Location Pointer is NOT equal to the Location Name!\n" +
			"Location Pointer String='%v'\n" +
			"Location Name = '%v'\n",
			tzdef.location.String() , tzdef.locationName)
	}

	loc, err := time.LoadLocation(tzdef.locationName)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError: time.LoadLocation(tzdef.locationName) Failed!\n" +
			"tzdef.locationName='%v'\n" +
			"Returned Error='%v'\n", tzdef.locationName, err.Error())
	}

	if loc.String() != tzdef.location.String() {
		return fmt.Errorf(ePrefix +
			"\nError: LoadLocation Pointer string NOT equal to tzdef.location.String() !\n" +
			"tzdef.location.String()='%v'\n" +
			"loc.String()='%v'\n", tzdef.location.String(), loc.String())
	}

	if tzdef.timeZoneType == TzType.Military() &&
		(tzdef.militaryTimeZoneLetter == "" ||
			tzdef.militaryTimeZoneName == "") {
		return fmt.Errorf(ePrefix +
			"\nError: This time zone is classified as a 'Military' Time Zone.\n" +
			"However, one or both of the Military Time Zone name strings is empty.\n" +
			"tzdef.militaryTimeZoneLetter='%v'\n" +
			"tzdef.militaryTimeZoneName='%v'\n", tzdef.militaryTimeZoneLetter ,tzdef.militaryTimeZoneName)
	}

	return nil
}

// isValidFromDateTime - Uses a time.Time input parameter, 'dateTime' to
// analyze the specified TimeZoneDefDto instance (tzdef). If the zone and
// location details of 'dateTime' are not perfectly matched to the current
// TimeZoneDefDto instance, the instance is considered INVALID, and this
// method returns 'false'.
//
// Otherwise, if all zone and location details are perfectly matched, this
// method returns 'true', signaling that the TimeZoneDateDefDto instance
// (tzdef) is VALID.
//
func (tzDefUtil *timeZoneDefUtility) isValidFromDateTime(
	tzdef *TimeZoneDefDto,
	dateTime time.Time) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.isValidFromDateTime()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if dateTime.IsZero() {
		return false
	}

	tzDefUtil2 := timeZoneDefUtility{}

	if tzDefUtil2.isEmpty(tzdef) {
		return false
	}

	tzDef2 := TimeZoneDefDto{}

	err := tzDefUtil2.setFromDateTime( &tzDef2, dateTime, "timeZoneDefUtility.isValidFromDateTime() ")

	if err != nil {
		return false
	}

	tzDef2.tagDescription = tzdef.tagDescription

	return tzDefUtil2.equal(tzdef, &tzDef2)
}

// SetFromDateTimeComponents - Re-initializes the values of the
// 'TimeZoneDefDto' instance based on input parameter, 'dateTime'.
//
/*
func (tzDefUtil *timeZoneDefUtility) setFromDateTime(
	tzdef *TimeZoneDefDto,
	dateTime time.Time,
	ePrefix string) error {

	ePrefix += "timeZoneDefUtility.setFromDateTime() "

	locPtr := dateTime.Location()
	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	if locPtr == nil {
		return fmt.Errorf(ePrefix +
			"Error: Input Parameter 'dateTime' has a nil Location Pointer!\n" +
			"dateTime='%v'\n", dateTime.Format(fmtStr))
	}

	zoneName, zoneSeconds := dateTime.Zone()

	locName := locPtr.String()

	if locName != zoneName {
		// Maybe a good Iana Time Zone!
	}

	timeStr := dateTime.Format(fmtStr)

	offsetLeadLen := len("01/02/2006 15:04:05.000000000 ")

	t2AbbrvLookup := locName + timeStr[offsetLeadLen:offsetLeadLen+5]

	stdAbbrvs := StdTZoneAbbreviations{}

	tzones, ok := stdAbbrvs.AbbrvOffsetToTimeZones(t2AbbrvLookup)

	if !ok {
		return fmt.Errorf(ePrefix +
			"\nError: Could NOT location Time Zone Abbreviation!\n" +
			"Mapping Time Zone Abbreviation to Time Zones Failed.\n" +
			"Lookup key='%v'\n", t2AbbrvLookup)
	}

	var newTZone string

	if len(tzones) == 1 {
		newTZone = tzones[0]
	}

	if len(newTZone) == 0 {
		// tzAbbrvToTimeZonePriorityList
		for i:=0; i < len(tzAbbrvToTimeZonePriorityList) && len(newTZone)== 0 ; i++ {

			for j:=0; j < len(tzones); j++ {

				if strings.HasPrefix(tzones[j], priorityList[i]) {
					newTZone = tzones[j]
					break
				}
			}
		}
	}

	if len(newTZone) == 0 {
		newTZone = tzones[0]
	}


	return nil
}

 */

func (tzDefUtil *timeZoneDefUtility) setFromDateTime(
	tzdef *TimeZoneDefDto,
	dateTime time.Time,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.setFromDateTime() "

	if tzdef == nil {
		return errors.New(ePrefix +
			"\nInput parameter 'tzdef' is nil!\n")
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'dateTime' is a ZERO value!")
	}

	zoneName, zoneOffsetSeconds := dateTime.Zone()

	dtMech := dateTimeMechanics{}

	offsetHours,
		offsetMinutes,
		offsetSeconds,
		zoneSign := dtMech.allocateSecondsToHrsMinSecs(zoneOffsetSeconds)

	tzDefUtil2 := timeZoneDefUtility{}

	zoneOffset,
	utcOffset,
	err := tzDefUtil2.calcZoneProfile(
		offsetHours,
		offsetMinutes,
		offsetSeconds,
		zoneSign,
		zoneName,
		ePrefix)

	if err != nil {
		return err
	}

	locationPtr := dateTime.Location()
	locationName := dateTime.Location().String()

	if locationName != "Local" {
		abbrvId := locationName + utcOffset

		_,
		_,
		ianaTimeZoneName,
		ianaLocationPtr,
		err := dtMech.convertTzAbbreviationToTimeZone(abbrvId, ePrefix)

		if err == nil {
			locationPtr = ianaLocationPtr
			locationName = ianaTimeZoneName
		}
	}

	tzDefUtil2.empty(tzdef)
	tzdef.zoneName = zoneName
	tzdef.zoneOffsetSeconds = offsetSeconds
	tzdef.zoneSign = zoneSign
	tzdef.offsetHours = offsetHours
	tzdef.offsetMinutes = offsetMinutes
	tzdef.offsetSeconds = offsetSeconds
	tzdef.zoneOffset = zoneOffset
	tzdef.utcOffset = utcOffset
	tzdef.location = locationPtr
	tzdef.locationName = locationName
	tzdef.locationNameType = LocNameType.ConvertibleTimeZoneName()
	tzdef.militaryTimeZoneName = ""
	tzdef.militaryTimeZoneLetter = ""
	tzdef.tagDescription = ""

	testTzName := strings.ToLower(tzdef.locationName)

	if testTzName == "local" {
		tzdef.timeZoneType = TzType.Local()
	} else {
		tzdef.timeZoneType = TzType.Iana()
	}

	return nil
}

// setFromTimeZoneName - Sets the data fields of the specified
// TimeZoneDefDto instance base on the time zone text name.
//
func (tzDefUtil *timeZoneDefUtility) setFromTimeZoneName(
	tzdef *TimeZoneDefDto,
	timeZoneName,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.setFromTimeZoneName() "

	if tzdef == nil {
		return errors.New(ePrefix +
			"\nInput parameter 'tzdef' is nil!\n")
	}

	timeZoneName = strings.TrimLeft(strings.TrimRight(timeZoneName, " "), " ")

	if len(timeZoneName) == 0 {
		return errors.New(ePrefix +
			"\nInput parameter 'timeZoneName' is an empty string!\n")
	}

	dtUtil := DTimeUtility{}

	var milTzLetter, milTzName string
	var ianaLocationPtr *time.Location
	var tzType TimeZoneType
	var err error

	milTzLetter,
	milTzName,
	_,
	ianaLocationPtr,
	tzType,
	err = dtUtil.GetTimeZoneFromName(
		timeZoneName,
		ePrefix)

	if err != nil {
		return fmt.Errorf("'timeZoneName' is INVALID!\n" +
			"%v", err.Error())
	}

	dateTime := time.Now().In(ianaLocationPtr)

	tzDefUtil2 := timeZoneDefUtility{}

	tzDefUtil2.empty(tzdef)

	tzdef.zoneName, tzdef.zoneOffsetSeconds = dateTime.Zone()

	tzDefUtil2.allocateZoneOffsetSeconds(tzdef, tzdef.zoneOffsetSeconds)

	tzdef.location = dateTime.Location()

	tzdef.locationName = dateTime.Location().String()

	tzDefUtil2.setZoneProfile(tzdef)

	tzdef.tagDescription = ""
	tzdef.timeZoneType = tzType

	if tzdef.timeZoneType == TzType.Military() {
		tzdef.militaryTimeZoneLetter = milTzLetter
		tzdef.militaryTimeZoneName = milTzName
	}

	return nil
}

// setZoneProfile - assembles and assigns the composite zone
// offset, zone names, zone abbreviation and UTC offsets.
//
// The TimeZoneDefDto.ZoneOffset field formatted in accordance
// with the following examples:
//      "-0600 CST"
//      "+0200 EET"
//
func (tzDefUtil *timeZoneDefUtility) calcZoneProfile(
	offsetHours int,
	offsetMinutes int,
	offsetSeconds int,
	zoneSign int,
	zoneName string,
	ePrefix string) (
	zoneOffset string,
	utcOffset string,
	err error) {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.calcZoneProfile() "
	zoneOffset = ""
	utcOffset = ""
	err = nil

	if zoneSign < -1 ||
		zoneSign > 1 ||
		zoneSign == 0 {
		return zoneOffset, utcOffset,
		fmt.Errorf(ePrefix,
			"Error: Input parameter 'zoneSign' must be equal to -1 or +1.\n" +
			"zoneSign='%v'\n", zoneSign)
	}

	if offsetHours < 0 {
		return zoneOffset, utcOffset,
			fmt.Errorf(ePrefix,
				"Error: Input parameter 'offsetHours' is less than zero.\n" +
					"offsetHours='%v'\n", offsetHours)
	}

	if offsetMinutes < 0 {
		return zoneOffset, utcOffset,
			fmt.Errorf(ePrefix,
				"Error: Input parameter 'offsetMinutes' is less than zero.\n" +
					"offsetMinutes='%v'\n", offsetMinutes)
	}

	if offsetSeconds < 0 {
		return zoneOffset, utcOffset,
			fmt.Errorf(ePrefix,
				"Error: Input parameter 'offsetSeconds' is less than zero.\n" +
					"offsetSeconds='%v'\n", offsetSeconds)
	}

	if len(zoneName) == 0 {
		return zoneOffset, utcOffset,
			errors.New(ePrefix +
				"Error: Input parameter 'zoneName' is an empty string.\n")
	}

	// Generates an offset in the form of "+0330" or "-0330"
	if zoneSign == -1 {
		zoneOffset += "-"
	} else {
		zoneOffset += "+"
	}

	zoneOffset += fmt.Sprintf("%02d%02d", offsetHours, offsetMinutes)

	// Generates final UTC offset in the form
	// "-0500" or "+0200"
	utcOffset = zoneOffset

	if offsetSeconds > 0 {
		zoneOffset += fmt.Sprintf("%02d", offsetSeconds)
	}

	// Generates final ZoneOffset in the form
	// "-0500 CST" or "+0200 EET"
	zoneOffset += " " + zoneName

	return zoneOffset, utcOffset, err
}

// parseMilitaryTzNameAndLetter - Parses a text string which
// contains either a single letter military time zone designation
// or a multi-character military time zone text name.
//
// If successful, three string values are returned. The first
// is a string containing the valid Military Time Zone
// Letter designation. The second returned string contains
// the text name of the Military Time Zone. The third string,
//
// If an error is encountered, the return value, 'err' is populated
// with an appropriate error message. Otherwise, 'err' is set
// to 'nil' signaling no error was encountered.
//
func (tzDefUtil *timeZoneDefUtility) parseMilitaryTzNameAndLetter(
	rawTz,
	ePrefix string) (milTzLetter, milTzName, equivalentIanaTimeZone string, err error) {

		tzDefUtil.lock.Lock()

		defer tzDefUtil.lock.Unlock()

	milTzLetter = ""
	milTzName = ""
	equivalentIanaTimeZone = ""
	err = nil

	ePrefix += "timeZoneDefUtility.parseMilitaryTzNameAndLetter() "

	lMilTz := len(rawTz)

	if lMilTz == 0 {
		err = errors.New(ePrefix +
			"\nError: Input Parameter 'rawTz' is EMPTY!\n")
		return milTzLetter, milTzName, equivalentIanaTimeZone, err
	}

	var ok bool
	milTzData := MilitaryTimeZoneData{}

	if lMilTz == 1 {

		milTzLetter = strings.ToUpper(rawTz)

		milTzName , ok =
			milTzData.MilTzLetterToTextName(milTzLetter)

		if !ok {
			err = fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'militaryTz' is INVALID!\n" +
				"'militaryTz' DOES NOT map to a valid Military Time Zone.\n" +
				"militaryTz='%v'", milTzLetter)
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
		}

		equivalentIanaTimeZone, ok = milTzData.MilitaryTzToIanaTz(milTzName)

		if !ok {
			err = fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'rawTz' is INVALID!\n" +
				"'rawTz' DOES NOT map to a valid IANA Time Zone.\n" +
				"rawTz='%v'", milTzName)
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
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
			err = fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'rawTz' is INVALID!\n" +
				"'rawTz' DOES NOT map to a valid IANA Time Zone.\n" +
				"Military Time Zone Letter='%v'\n" +
				"Military Time Zone Text Name='%v'", milTzLetter ,milTzName)
			milTzLetter = ""
			milTzName = ""
			equivalentIanaTimeZone = ""
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
		}
	}

	err = nil

	return milTzLetter, milTzName, equivalentIanaTimeZone, err
}