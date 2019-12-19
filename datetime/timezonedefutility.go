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

	tzdef.offsetMinutes = signedZoneOffsetSeconds / 60 // compute minutes
	signedZoneOffsetSeconds -= tzdef.offsetMinutes * 60

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
		tzdef.zoneSign != tzdef2.zoneSign ||
		tzdef.offsetHours != tzdef2.offsetHours ||
		tzdef.offsetMinutes != tzdef2.offsetMinutes ||
		tzdef.offsetSeconds != tzdef2.offsetSeconds ||
		tzdef.zoneOffset != tzdef2.zoneOffset ||
		tzdef.utcOffset != tzdef2.utcOffset ||
		tzdef.locationName != tzdef2.locationName ||
		tzdef.militaryTimeZoneLetter != tzdef2.militaryTimeZoneLetter ||
		tzdef.militaryTimeZoneName != tzdef2.militaryTimeZoneName ||
		tzdef.tagDescription != tzdef2.tagDescription ||
		tzdef.timeZoneType != tzdef2.timeZoneType {
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

// EqualZoneOffsets - Compares ZoneOffsets for two TimeZoneDefDto's and
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
	tzdef *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.isValidTimeZoneDefDto() ")
	}

	tzDefUtil2 := timeZoneDefUtility{}

	if tzDefUtil2.isEmpty(tzdef) {
		return false
	}

	if strings.TrimLeft(strings.TrimRight(tzdef.locationName, " "), " ") == "" {
		return false
	}

	if tzdef.location.String() != tzdef.locationName {
		return false
	}

	loc, err := time.LoadLocation(tzdef.locationName)

	if err != nil {
		return false
	}

	if loc.String() != tzdef.location.String() {
		return false
	}

	if tzdef.timeZoneType == TzType.Military() &&
		(tzdef.militaryTimeZoneLetter == "" ||
			tzdef.militaryTimeZoneName == "") {
		return false
	}

	return true
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

	tzDefUtil2 := timeZoneDefUtility{}

	tzDefUtil2.empty(tzdef)

	tzdef.zoneName, tzdef.zoneOffsetSeconds = dateTime.Zone()

	tzDefUtil2.allocateZoneOffsetSeconds(tzdef, tzdef.zoneOffsetSeconds)

	// If dateTime.Location() is nil, it returns UTC
	tzdef.location = dateTime.Location()

	tzdef.locationName = dateTime.Location().String()

	tzDefUtil2.setZoneProfile(tzdef)

	tzdef.tagDescription = ""

	testTzName := strings.ToLower(tzdef.locationName)

	if testTzName == "local" {
		tzdef.timeZoneType = TzType.Local()
	} else if testTzName == "utc"  {
		tzdef.timeZoneType = TzType.Utc()
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

	tzDefUtil2 := timeZoneDefUtility{}

	if len(timeZoneName) == 1 {
		return tzDefUtil2.configureMilitaryTimeZone(tzdef, timeZoneName, ePrefix)
	}

	testTzName := strings.ToLower(timeZoneName)

	if testTzName == "utc" ||
			testTzName == "uct" ||
				testTzName == "gmt" {
		return tzDefUtil2.configureUtcZone(tzdef, timeZoneName, ePrefix)
	}

	if testTzName == "local" {
		return tzDefUtil2.configureLocalTimeZone(tzdef, timeZoneName, ePrefix)
	}

	_, err := time.LoadLocation(timeZoneName)

	if err == nil {
		return tzDefUtil2.configureIanaTimeZone(tzdef, timeZoneName, ePrefix)
	}

	return tzDefUtil2.configureMilitaryTimeZone(tzdef, timeZoneName, ePrefix)
}


// configureIanaTimeZone - Configures the specified 'TimeZoneDefDto'
// instance, 'tzdef', as an IANA Time Zone.
//
// The Go Programming Language uses IANA Time Zones in date-time
// calculations.
//
// Reference:
//    https://golang.org/pkg/time/
//    https://golang.org/pkg/time/#LoadLocation
//
// The IANA Time Zone database is widely recognized as a leading
// authority on global time zones.
//
// Reference:
//    https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//    https://en.wikipedia.org/wiki/Tz_database
//    https://en.wikipedia.org/wiki/List_of_military_time_zones
//
// For additional information on the IANA Time Zone Database,
// reference:
//
//    https://www.iana.org/time-zones
//    https://data.iana.org/time-zones/releases/
//
// Also, for easy access to the 600+ time zones, see type
// 'TimeZones' in source file: "datetime/timezonedata.go".
//
func (tzDefUtil *timeZoneDefUtility) configureIanaTimeZone(
	tzdef *TimeZoneDefDto,
	ianaTimeZoneName,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.configureLocalTimeZone() "

	if tzdef == nil {
		return errors.New(ePrefix +
			"\nInput parameter 'tzdef' is nil!\n")
	}

	ianaTimeZoneName = strings.TrimLeft(strings.TrimRight(ianaTimeZoneName, " "), " ")

	locPtr, err := time.LoadLocation(ianaTimeZoneName)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by time.LoadLocation(ianaTimeZoneName)\n" +
			"ianaTimeZoneName='%v'\n" +
			"Error='%v'\n", ianaTimeZoneName, err.Error())
	}

	dateTimeNow := time.Now().In(locPtr)

	tzDefUtil2 := timeZoneDefUtility{}

	err = tzDefUtil2.setFromDateTime(tzdef, dateTimeNow, ePrefix)

	if err != nil {
		return fmt.Errorf("IANA Time Zone Name is invalid!\n" +
			"ianaTimeZoneName='%v'\n" +
			"Error='%v'\n", ianaTimeZoneName, err.Error())
	}

	tzdef.timeZoneType = TzType.Iana()

	return nil
}

// configureLocalTimeZone - Configures the specified 'TimeZoneDefDto'
// instance, 'tzdef', as a 'Local' Time Zone.
//
// The 'Local' time zone is a construct of the Go Programming
// Language. As such, the 'Local' time zone is automatically
// configured using the time zone applied by the host computer
// running this code.
//
// Reference:
//   https://golang.org/pkg/time/#Time.Local
//
func (tzDefUtil *timeZoneDefUtility) configureLocalTimeZone(
	tzdef *TimeZoneDefDto,
	localTimeZoneName,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.configureLocalTimeZone() "

	if tzdef == nil {
		return errors.New(ePrefix +
			"\nInput parameter 'tzdef' is nil!\n")
	}

	tLocalZoneName := strings.TrimLeft(strings.TrimRight(localTimeZoneName, " "), " ")

	if len(localTimeZoneName) == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'localTimeZoneName' is an empty string!\n")
	}

	tLocalZoneName = strings.ToLower(tLocalZoneName)

	if tLocalZoneName != "local" {
		return fmt.Errorf(ePrefix +
			"\nError: Input prameter 'localTimeZoneName' is NOT a Local Time Zone!\n" +
			"localTimeZoneName:='%v'\n", localTimeZoneName)
	}

	tLocalZoneName = "Local"

	tzDefUtil2 := timeZoneDefUtility{}

	locPtr, err := time.LoadLocation(tLocalZoneName)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by time.LoadLocation(\"Local\")\n" +
			"Error='%v'\n", err.Error())
	}

	dateTime := time.Now().In(locPtr)

	err = tzDefUtil2.setFromDateTime(tzdef, dateTime, ePrefix)

	if err != nil {
		return fmt.Errorf("Time Zone Local is invalid!\n" +
			"Input parameter localTimeZoneName='%v'\n" +
			"Error='%v'", localTimeZoneName, err.Error())
	}

	tzdef.timeZoneType = TzType.Local()

	return nil
}

// configureMilitaryTimeZone - Configures the specified 'TimeZoneDefDto'
// instance, 'tzdef', as a Military Time Zone.
//
//
// For information on Military Time Zones reference:
//     https://en.wikipedia.org/wiki/List_of_military_time_zones
//     http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//     https://www.timeanddate.com/time/zones/military
//     https://www.timeanddate.com/worldclock/timezone/alpha
//     https://www.timeanddate.com/time/map/
//
// Military time zones are commonly used in aviation as well as at sea.
// They are also known as nautical or maritime time zones.
//
// The 'J' (Juliet) Time Zone is occasionally used to refer to the observer's
// local time. Note that Time Zone 'J' (Juliet) is not listed below.
//
//
//    Time Zone       Time Zone        Equivalent IANA          UTC
//   Abbreviation       Name              Time Zone            Offset
//   ------------     --------          ---------------        ------
//
//       A        Alpha Time Zone         Etc/GMT-1            UTC +1
//       B        Bravo Time Zone         Etc/GMT-2            UTC +2
//       C        Charlie Time Zone       Etc/GMT-3            UTC +3
//       D        Delta Time Zone         Etc/GMT-4            UTC +4
//       E        Echo Time Zone          Etc/GMT-5            UTC +5
//       F        Foxtrot Time Zone       Etc/GMT-6            UTC +6
//       G        Golf Time Zone          Etc/GMT-7            UTC +7
//       H        Hotel Time Zone         Etc/GMT-8            UTC +8
//       I        India Time Zone         Etc/GMT-9            UTC +9
//       K        Kilo Time Zone          Etc/GMT-10           UTC +10
//       L        Lima Time Zone          Etc/GMT-11           UTC +11
//       M        Mike Time Zone          Etc/GMT-12           UTC +12
//       N        November Time Zone      Etc/GMT+1            UTC -1
//       O        Oscar Time Zone         Etc/GMT+2            UTC -2
//       P        Papa Time Zone          Etc/GMT+3            UTC -3
//       Q        Quebec Time Zone        Etc/GMT+4            UTC -4
//       R        Romeo Time Zone         Etc/GMT+5            UTC -5
//       S        Sierra Time Zone        Etc/GMT+6            UTC -6
//       T        Tango Time Zone         Etc/GMT+7            UTC -7
//       U        Uniform Time Zone       Etc/GMT+8            UTC -8
//       V        Victor Time Zone        Etc/GMT+9            UTC -9
//       W        Whiskey Time Zone       Etc/GMT+10           UTC -10
//       X        X-ray Time Zone         Etc/GMT+11           UTC -11
//       Y        Yankee Time Zone        Etc/GMT+12           UTC -12
//       Z        Zulu Time Zone          UTC                  UTC +0
//
// For more information on Military Time Zones, reference the type,
// 'TimeZones' and member variable 'TimeZones.militaryTimeZones' in source
//  file, 'datetime/timezonedata.go'.
//
func (tzDefUtil *timeZoneDefUtility) configureMilitaryTimeZone(
	tzdef *TimeZoneDefDto,
	militaryTimeZoneName,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.configureMilitaryTimeZone() "

	if tzdef == nil {
		return errors.New(ePrefix +
			"\nInput parameter 'tzdef' is nil!\n")
	}

	militaryTimeZoneName = strings.TrimLeft(strings.TrimRight(militaryTimeZoneName," "), " ")

	if len(militaryTimeZoneName) == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'militaryTimeZoneName' is an empty string!\n")
	}

	tzDefUtil2 := timeZoneDefUtility{}

	milTzLetter,
	milTzName,
	equivalentIanaTimeZone,
	err := tzDefUtil2.parseMilitaryTzNameAndLetter(
		militaryTimeZoneName,
		ePrefix)

	if err != nil {
		return err
	}

	var locPtr *time.Location

	locPtr, err = time.LoadLocation(equivalentIanaTimeZone)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError return by time.LoadLocation(equivalentIanaTimeZone)\n" +
			"\nMilitary Time Zone='%v'\n" +
			"\nequivalentIanaTimeZone='%v'\n" +
			"\nError='%v'\n",
				militaryTimeZoneName,
					equivalentIanaTimeZone,
						err.Error())
	}

	dateTime := time.Now().In(locPtr)

	err = tzDefUtil2.setFromDateTime(tzdef, dateTime, ePrefix)

	if err != nil {
		return fmt.Errorf("Military Time Zone Name is invalid!\n" +
			"\nMilitary Time Zone='%v'\n" +
			"\nequivalentIanaTimeZone='%v'\n" +
			"\nError='%v'\n",
			militaryTimeZoneName,
			equivalentIanaTimeZone,
			err.Error())
	}

	tzdef.timeZoneType = TzType.Military()
	tzdef.militaryTimeZoneLetter = milTzLetter
	tzdef.militaryTimeZoneName = milTzName

	return nil
}


// configureUtcZone - Configures the specified 'TimeZoneDefDto'
// instance, 'tzdef', as a UTC Time Zone.
//
// Coordinated Universal Time (or UTC) is the primary time standard
// by which the world regulates clocks and time. It is within about
// 1-second of mean solar time at 0° longitude, and is not adjusted
// for daylight saving time. In some countries, the term "Greenwich
// Mean Time (GMT)" is used as an equivalent for 'UTC'.
//
// UTC is equivalent to a zero offset: UTC+0000. For additional
// information, reference:
//
//     https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// Equivalent time zones configured as 'UTC' are:
//   "UTC"
//   "UCT"
//   "GMT"
//
func (tzDefUtil *timeZoneDefUtility) configureUtcZone(
	tzdef *TimeZoneDefDto,
	utcTimeZoneName,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.configureUtcZone() "

	if tzdef == nil {
		return errors.New(ePrefix +
			"\nInput parameter 'tzdef' is nil!\n")
	}

	utcTzName := strings.TrimLeft(strings.TrimRight(utcTimeZoneName, " "), " ")

	if len(utcTzName) == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'utcTimeZoneName' is an empty string!\n" )
	}

	utcTzName = strings.ToLower(utcTimeZoneName)

	switch utcTzName {

	case "utc":
		utcTimeZoneName = "UTC"
	case "uct":
		utcTimeZoneName = "UTC"
	case "gmt":
		utcTimeZoneName = "UTC"
	default:
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter utcTimeZoneName is invalid!\n" +
			"utcTimeZone='%v'\n", utcTimeZoneName)
	}

	tzDefUtil2 := timeZoneDefUtility{}

	dateTime := time.Now().In(time.UTC)

	err := tzDefUtil2.setFromDateTime(tzdef, dateTime, ePrefix)

	if err != nil {
		return fmt.Errorf("Time Zone UTC is invalid!\n" +
			"Input parameter utcTimeZoneName='%v'\n" +
			"Error='%v'", utcTimeZoneName, err.Error())
	}

	tzdef.timeZoneType = TzType.Utc()

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
func (tzDefUtil *timeZoneDefUtility) setZoneProfile(
	tzdef *TimeZoneDefDto) {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("Error: timeZoneDefUtility.setZoneProfile() - tzdef pointer is nil!")
	}

	tzdef.zoneOffset = ""

	// Generates an offset in the form of "+0330" or "-0330"
	if tzdef.zoneSign < 0 {
		tzdef.zoneOffset += "-"
	} else {
		tzdef.zoneOffset += "+"
	}

	tzdef.zoneOffset += fmt.Sprintf("%02d%02d", tzdef.offsetHours, tzdef.offsetMinutes)

	tzdef.utcOffset = tzdef.zoneOffset

	if tzdef.offsetSeconds > 0 {
		tzdef.zoneOffset += fmt.Sprintf("%02d", tzdef.offsetSeconds)
	}

	// Generates final ZoneOffset in the form
	// "-0500 CST" or "+0200 EET"
	tzdef.zoneOffset += " " + tzdef.zoneName

	return
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