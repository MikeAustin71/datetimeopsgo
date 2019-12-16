package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

/*
TimeZoneDefDto

 This source file is located in source code repository:
 		https://github.com/MikeAustin71/datetimeopsgo.git'

 This source code file is located at:
		MikeAustin71\datetimeopsgo\datetime\timezonedef.go


Overview and General Usage

 This structure is designed to store detailed information
 on Time Zones. It used primarily as a 'helper' or subsidiary
 structure for Type, 'TimeZoneDto'.

 The TimeZoneDefDto method 'New' may be used to extract
 detailed time zone information from a time.Time date time
 value.

*/

// TimeZoneDefDto - Time Zone Definition Dto
// Contains detailed parameters describing a specific
// Time Zone and Time Zone Location
type TimeZoneDefDto struct {
	zoneName               string         // The Time Zone abbreviation. Examples: 'EST', 'CST', 'PST'
	zoneOffsetSeconds      int            // Signed number of seconds offset from UTC. + == East of UTC; - == West of UTC
	zoneSign               int            // -1 == West of UTC  +1 == East of UTC
	offsetHours            int            // Normalized Offset Hours from UTC. Always a positive number, refer to ZoneSign
	offsetMinutes          int            // Normalized Offset Minutes offset from UTC. Always a positive number, refer to ZoneSign
	offsetSeconds          int            // Normalized Offset Seconds offset from UTC. Always a positive number, refer to ZoneSign
	zoneOffset             string         // A text string representing the time zone. Example "-0600 CST" or "+0200 EET"
	utcOffset              string         // A text string representing the offset for UTC. Example "-0600" or "+0200"
	location               *time.Location // Pointer to a Time Zone Location
	locationName           string         // Time Zone Location Name Examples: "Local", "America/Chicago", "America/New_York"
	militaryTimeZoneName   string         // Full Military Time Zone text name. Examples: "Alpha", "Bravo", "Charlie", "Zulu"
	militaryTimeZoneLetter string         // Single Alphabetic Character identifying a Military Time Zone.
	tagDescription         string         // Unused - Available for classification, labeling or description by user.
	timeZoneType           TimeZoneType   // Enumeration of Time Zone Type:
	                                      //  TzType.None()
	                                      //  TzType.Iana()
	                                      //  TzType.Military()
	                                      //  TzType.Local()
	                                      //  TzType.UtcOffset()
	lock                   sync.Mutex     // Used for implementing thread safe operations.
}

// CopyIn - Copies an incoming TimeZoneDefDto into the
// data fields of the current TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) CopyIn(tzdef2 TimeZoneDefDto) {

	tzdef.Empty()

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
	tzdef.tagDescription = tzdef2.tagDescription

}

// CopyOut - creates and returns a deep copy of the current
// TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) CopyOut() TimeZoneDefDto {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.CopyOut(tzdef)
}

// Empty - Resets all field values for the current TimeZoneDefDto
// instance to their uninitialized or 'zero' states.
func (tzdef *TimeZoneDefDto) Empty() {

	tzDefUtil := timeZoneDefUtility{}

	tzDefUtil.empty(tzdef)

}

// Equal - Determines if two TimeZoneDefDto are equivalent in
// value. Returns 'true' of two TimeZoneDefDto's are equal in
// all respects.
func (tzdef *TimeZoneDefDto) Equal(tzdef2 TimeZoneDefDto) bool {

	if tzdef.zoneName != tzdef2.zoneName ||
		tzdef.zoneOffsetSeconds != tzdef2.zoneOffsetSeconds ||
		tzdef.zoneSign != tzdef2.zoneSign ||
		tzdef.offsetHours != tzdef2.offsetHours ||
		tzdef.offsetMinutes != tzdef2.offsetMinutes ||
		tzdef.offsetSeconds != tzdef2.offsetSeconds ||
		tzdef.zoneOffset != tzdef2.zoneOffset ||
		tzdef.utcOffset != tzdef2.utcOffset ||
		tzdef.locationName != tzdef2.locationName ||
		tzdef.tagDescription != tzdef2.tagDescription {
		return false
	}

	if tzdef.location != nil && tzdef2.location == nil ||
	   tzdef.location == nil && tzdef2.location != nil ||
		tzdef.location.String() != tzdef2.location.String() {
		return false
	}

	return true
}

// EqualOffsetSeconds - Compares Zone Offset Seconds for two TimeZoneDefDto's and
// returns 'true' if they are equal.
//
// ZoneOffsetSeconds is a signed number of seconds offset from UTC:
// 		+ == East of UTC
// 		- == West of UTC
func (tzdef *TimeZoneDefDto) EqualOffsetSeconds(tzdef2 TimeZoneDefDto) bool {

	if tzdef.zoneOffsetSeconds == tzdef2.zoneOffsetSeconds {
		return true
	}

	return false

}

// EqualZones - Compares ZoneOffsets for two TimeZoneDefDto's and
// returns 'true' if they are equal.
//
// Zone Offset is a text string representing the time zone.
// Example "-0500 CDT"
//
func (tzdef *TimeZoneDefDto) EqualZones(tzdef2 TimeZoneDefDto) bool {

	if tzdef.zoneOffset == tzdef2.zoneOffset {
		return true
	}

	return false

}

// EqualLocations - Compares the Time Zone Locations for two TimeZoneDefDto's
// and returns 'true' if they are equal.
//
// Time Zone Location Name Examples:
// 		"Local"
// 		"America/Chicago"
// 		"America/New_York"
//
func (tzdef *TimeZoneDefDto) EqualLocations(tzdef2 TimeZoneDefDto) bool {

	if tzdef.locationName == tzdef2.locationName {
		return true
	}

	return false

}

// EqualZoneLocation - Compares two TimeZoneDefDto's and returns
// 'true' if both the TimeZoneLocations and Time Zones match.
func (tzdef *TimeZoneDefDto) EqualZoneLocation(tzdef2 TimeZoneDefDto) bool {

	if tzdef.EqualLocations(tzdef2) && tzdef.EqualZoneLocation(tzdef2) {
		return true
	}

	return false

}

// GetTagDescription - Returns TimeZoneDefDto member variable
// Description value.
//
// Unused - Available to user. This string is typically used
// classification, labeling or description text by user.
//
func (tzdef *TimeZoneDefDto) GetTagDescription() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.tagDescription
}

// GetLocationPtr - Returns a pointer to a time.Location or
// time zone location (TimeZoneDefDto.Location).
//
func (tzdef *TimeZoneDefDto) GetLocationPtr() *time.Location {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.location
}

// GetLocationName - Returns TimeZoneDefDto member variable
// LocationName value.
//
// Time Zone Location Name Examples: "Local", "America/Chicago",
// "America/New_York".
//
func (tzdef *TimeZoneDefDto) GetLocationName() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.locationName
}

// GetOffsetHours - Returns TimeZoneDefDto member variable
// ZoneOffset value.
//
// Normalized Offset Hours from UTC. Always a positive number,
// refer to ZoneSign to determine East or West of UTC.
//
func (tzdef *TimeZoneDefDto) GetOffsetHours() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.offsetHours
}

// GetOffsetMinutes - Returns TimeZoneDefDto member variable
// ZoneOffsetMinutes value.
//
// Normalized Offset Minutes offset from UTC. Always a
// positive number, refer to ZoneSign to determine East or
// West of UTC.
//
func (tzdef *TimeZoneDefDto) GetOffsetMinutes() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.offsetMinutes
}

// GetOffsetSeconds - Returns TimeZoneDefDto member variable
// ZoneOffsetSeconds value.
//
// Normalized Offset Seconds offset from UTC. Always a
// positive number, refer to ZoneSign to determine East or
// West of UTC.
//
func (tzdef *TimeZoneDefDto) GetOffsetSeconds() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.offsetSeconds
}

// GetUtcOffset - Returns the offset from UTC as a string.
// Examples of the UTC offset format are: "-0600" or "+0200".
//
func (tzdef *TimeZoneDefDto) GetUtcOffset() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.utcOffset
}

// GetZoneName - Returns TimeZoneDefDto member variable
// TimeZoneDefDto.zoneName value.
//
// Zone Name is the Time Zone abbreviation. This may
// may be a series of characters, like "EST", "CST"
// and "PDT" - or - if a time zone abbreviation does
// not exist for this time zone, the time zone abbreviation
// might be listed simply as the UTC offset ('-0430')
//
func (tzdef *TimeZoneDefDto) GetZoneName() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.zoneName
}

// GetZoneOffset - Returns TimeZoneDefDto member variable
// ZoneOffset value.
func (tzdef *TimeZoneDefDto) GetZoneOffset() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.zoneOffset
}

// GetZoneOffsetSeconds - Returns TimeZoneDefDto member variable
// ZoneOffsetSeconds value.
//
func (tzdef *TimeZoneDefDto) GetZoneOffsetSeconds() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.zoneOffsetSeconds
}

// GetZoneSign - Returns TimeZoneDefDto member variable
// ZoneSign value.
//
// -1 == West of UTC  +1 == East of UTC
//
func (tzdef *TimeZoneDefDto) GetZoneSign() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.zoneSign
}

// GetTimeZoneType - Returns the Time Zone Type associated
// with this instance of TimeZoneDefDto.
//
// Time Zone Type is styled as an enumeration.
func (tzdef *TimeZoneDefDto) GetTimeZoneType() TimeZoneType {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.timeZoneType
}

// IsEmpty - Determines whether the current TimeZoneDefDto
// instance is Empty.
//
// If the TimeZoneDefDto instance is NOT populated, the method
// returns 'true'. Otherwise, it returns 'false'.
//
func (tzdef *TimeZoneDefDto) IsEmpty() bool {

	if tzdef.zoneName != "" ||
		tzdef.zoneOffsetSeconds != 0 ||
		tzdef.zoneSign != 0 ||
		tzdef.offsetHours != 0 ||
		tzdef.offsetMinutes != 0 ||
		tzdef.offsetSeconds != 0 ||
		tzdef.zoneOffset != "" ||
		tzdef.utcOffset != "" ||
		tzdef.locationName != "" {
		return false
	}

	return true
}

// IsValidDateTime - Analyzes the current TimeZoneDefDto instance
// to determine validity.
//
// This method returns 'true' if the TimeZoneDefDto is
// valid.  Otherwise, it returns 'false'.
//
func (tzdef *TimeZoneDefDto) IsValid() bool {

	if tzdef.IsEmpty() {
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

	if loc != tzdef.location {
		return false
	}

	return true
}

// IsValidFromDateTime - Uses a time.Time input parameter, 'dateTime' to
// analyze the current TimeZoneDefDto instance. If the zone and location
// details of 'dateTime' are not perfectly matched to the current TimeZoneDefDto
// instance, the instance is considered INVALID, and this method returns 'false'.
//
// Otherwise, if all zone and location details are perfectly matched, this
// method returns 'true', signaling that the TimeZoneDateDefDto instance
// is VALID.
//
func (tzdef *TimeZoneDefDto) IsValidFromDateTime(dateTime time.Time) bool {

	if tzdef.IsEmpty() {
		return false
	}

	tzdef2, err := TimeZoneDefDto{}.New(dateTime)

	if err != nil {
		return false
	}

	tzdef2.tagDescription = tzdef.tagDescription

	if !tzdef.Equal(tzdef2) {
		return false
	}

	return true
}

// New - Creates and returns a new TimeZoneDefDto instance based on
// a 'dateTime (time.Time) input parameter.
//
// Input Parameter
// ===============
//
//  dateTime   time.Time  - A date time value which will be used to construct the
//                          elements of a Time Zone Definition Dto instance.
//
// Returns
// =======
//
// This method will return two Types:
//      (1) A Time Zone Definition Dto
//      (2) An 'error' type
//
//  (1) If successful, this method will return a valid, populated TimeZoneDefDto instance.
//      A TimeZoneDefDto is defined as follows:
//      type TimeZoneDefDto struct {
//        ZoneName            string
//        ZoneOffsetSeconds    int     // Signed number of seconds offset from UTC. + == East of UTC; - == West of UTC
//        ZoneSign             int     // -1 == West of UTC  +1 == East of UTC
//        OffsetHours          int     // Hours offset from UTC. Always a positive number, refer to ZoneSign
//        OffsetMinutes        int     // Minutes offset from UTC. Always a positive number, refer to ZoneSign
//        OffsetSeconds        int     // Seconds offset from UTC. Always a positive number, refer to ZoneSign
//        ZoneOffset           string  // A text string representing the time zone. Example "-0500 CDT"
//        Location             *time.Location  // Pointer to a Time Zone Location
//        LocationName         string          // Time Zone Location Name Examples: "Local", "America/Chicago", "America/New_York"
//        Description          string          // Unused - Available for classification, labeling or description by user.
//      }
//
//
//  (2)   If successful, this method will set the returned error instance to 'nil.
//        If errors are encountered a valid error message will be returned in the
//        error instance.
//
func (tzdef TimeZoneDefDto) New(dateTime time.Time) (TimeZoneDefDto, error) {

	ePrefix := "TimeZoneDefDto.New() "

	if dateTime.IsZero() {
		return TimeZoneDefDto{}, errors.New(ePrefix + "Error: Input parameter 'dateTime' is a ZERO value!")
	}

	tzDef2 := TimeZoneDefDto{}

	tzDefUtil := timeZoneDefUtility{}

	err := tzDefUtil.setFromDateTime( &tzDef2, dateTime, ePrefix)

	if err != nil {
		return TimeZoneDefDto{},err
	}

	return tzDef2, nil
}

// NewFromTimeZoneName - Creates and returns a new instance 'TimeZoneDefDto'.
// The new instance is based on input parameter 'tzLocPtr', a pointer to
// and instance to time.Location.
//
func (tzdef TimeZoneDefDto) NewFromTimeZoneLocationPtr(
	tzLocPtr *time.Location) (tzDefDto TimeZoneDefDto, err error) {

	ePrefix := "TimeZoneDefDto.NewFromTimeZoneLocationPtr() "

	err = nil
	tzDefDto = TimeZoneDefDto{}

	if tzLocPtr == nil {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'tzLocPtr' is 'nil'!\n")
		return tzDefDto, err
	}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromDateTime( &tzDefDto, time.Now().In(tzLocPtr), ePrefix)

	if err != nil {
		tzDefDto = TimeZoneDefDto{}
		return tzDefDto, err
	}

	return tzDefDto, err
}

// NewFromTimeZoneName - Creates and returns a new instance 'TimeZoneDefDto'.
// The new instance is based on input parameter 'timeZoneName'.
//
// The 'timeZoneName' string must be set to one of three values:
//
//   1. A valid IANA Time Zone name.
//
//   2. The time zone "Local", which Golang accepts as the time
//      zone currently configured on the host computer.
//
//   3. A valid Military Time Zone which can be submitted either as
//      a single alphabetic character or as a full Military Time
//      zone name.
//
func (tzdef TimeZoneDefDto) NewFromTimeZoneName(
	timeZoneName string) (tzDefDto TimeZoneDefDto, err error) {

	ePrefix := "TimeZoneDefDto.NewFromTimeZoneName() "

	var tzLoc *time.Location
	err = nil
	tzDefDto = TimeZoneDefDto{}

	var err2 error

	if len(timeZoneName) == 0 {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'timeZoneName' is EMPTY!\n")
		return tzDefDto, err
	}

	tzDefUtil := timeZoneDefUtility{}

	tzLoc, err2 = time.LoadLocation(timeZoneName)

	if err2 != nil {
		err = fmt.Errorf(ePrefix +
			"\nError returned by time.LoadLocation(timeZoneName).\n" +
			"timeZoneName='%v'\n" +
			"Error='%v'\n", timeZoneName, err2.Error())
		return tzDefDto, err
	}


	err = tzDefUtil.setFromDateTime(&tzDefDto, time.Now().In(tzLoc), ePrefix)

	return tzDefDto, err
}


// SetTagDescription - Sets TimeZoneDefDto private member variable
// TimeZoneDefDto.tagDescription to the value passed in 'tagDesc'.
//
// The TimeZoneDefDto.tagDescription string is available to users
// for use as a tag, label, classification or text description.
//
func (tzdef *TimeZoneDefDto) SetTagDescription(tagDesc string) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzdef.tagDescription = tagDesc

}

// SetFromDateTimeComponents - Re-initializes the values of the current
// TimeZoneDefDto instance based on input parameter, 'dateTime'.
func (tzdef *TimeZoneDefDto) SetFromDateTime(dateTime time.Time) error {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefDto.SetFromDateTimeComponents() "

	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'dateTime' is a ZERO value!")
	}

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.setFromDateTime(tzdef, dateTime, ePrefix)

}


// allocateZoneOffsetSeconds - allocates a signed value of total offset seconds from
// UTC to the associated fields in the current TimeZoneDefDto instance.
/*

func (tzdef *TimeZoneDefDto) allocateZoneOffsetSeconds(signedZoneOffsetSeconds int) {

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

// setZoneProfile - assembles and assigns the composite zone
// offset, zone names, zone abbreviation and UTC offsets.
//
// The TimeZoneDefDto.ZoneOffset field formatted in accordance
// with the following examples:
//      "-0600 CST"
//      "+0200 EET"
//
func (tzdef *TimeZoneDefDto) setZoneProfile() {

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

*/