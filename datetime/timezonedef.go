package datetime

import (
	"errors"
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
	//                                         definition.
	originalTimeZone       TimeZoneSpecDto  // The Time Zone Specification originally submitted.
	convertibleTimeZone    TimeZoneSpecDto  // A version of the original time zone with a new fully
	//                                         convertible time zone substituted.
	lock sync.Mutex // Used for implementing thread safe operations.
}

// CopyIn - Copies an incoming TimeZoneDefDto into the
// data fields of the current TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) CopyIn(tzdef2 TimeZoneDefDto) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	tzDefUtil.copyIn(tzdef, &tzdef2)

	return
}

// copyOut - creates and returns a deep copy of the current
// TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) CopyOut() TimeZoneDefDto {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.copyOut(tzdef)
}

// Empty - Resets all field values for the current TimeZoneDefDto
// instance to their uninitialized or 'zero' states.
func (tzdef *TimeZoneDefDto) Empty() {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	tzDefUtil.empty(tzdef)
}

// Equal - Determines if two TimeZoneDefDto are equivalent in
// value. Returns 'true' of two TimeZoneDefDto's are equal in
// all respects.
func (tzdef *TimeZoneDefDto) Equal(tzdef2 TimeZoneDefDto) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.equal(tzdef, &tzdef2)
}

// EqualOffsetSeconds - Compares Zone Offset Seconds for two TimeZoneDefDto's and
// returns 'true' if they are equal.
//
// ZoneOffsetSeconds is a signed number of seconds offset from UTC:
//   + == East of UTC
//   - == West of UTC
func (tzdef *TimeZoneDefDto) EqualOffsetSeconds(tzdef2 TimeZoneDefDto) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.equalOffsetSeconds(tzdef, &tzdef2)
}

// EqualZoneOffsets - Compares ZoneOffsets for two TimeZoneDefDto's and
// returns 'true' if they are equal.
//
// Zone Offset is a text string representing the offset from UTC plus the
// time zone abbreviation.
//
// Example "-0500 CDT"
//
func (tzdef *TimeZoneDefDto) EqualZoneOffsets(tzdef2 TimeZoneDefDto) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.equalZoneOffsets(tzdef, &tzdef2)
}

// EqualLocations - Compares the Time Zone Locations for two TimeZoneDefDto's
// and returns 'true' if they are equal.
//
// Time Zone Location Name Examples:
//   "Local"
//   "America/Chicago"
//   "America/New_York"
//
func (tzdef *TimeZoneDefDto) EqualLocations(tzdef2 TimeZoneDefDto) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.equalLocations(tzdef, &tzdef2)
}

// EqualZoneLocation - Compares two TimeZoneDefDto's and returns
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
func (tzdef *TimeZoneDefDto) EqualZoneLocation(tzdef2 TimeZoneDefDto) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	if !tzDefUtil.equalZoneLocation(tzdef, &tzdef2) {
		return false
	}

	return true
}

// GetLocationPtr - Returns a pointer to a time.Location or
// time zone location (TimeZoneDefDto.Location).
//
func (tzdef *TimeZoneDefDto) GetLocationPtr() *time.Location {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.locationPtr
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

	return tzdef.originalTimeZone.locationName
}

// GetLocationNameType - Returns the value of internal
// private member variable TimeZoneDefDto.locationNameType.
//
// Possible return values:
//  LocationNameType(0).None()
//  LocationNameType(0).ConvertibleAbbreviation()
//  LocationNameType(0).NonConvertibleAbbreviation()
//  LocationNameType(0).ConvertibleTimeZoneName()
//
func (tzdef *TimeZoneDefDto) GetLocationNameType() LocationNameType {

	return tzdef.originalTimeZone.locationNameType
}

// GetMilitaryTimeZoneLetter - Returns the single character which represents
// the Military Time Zone Letter designation.
//
// Examples:
//   "A"                = Alpha Military Time Zone
//   "B"                = Bravo Military Time Zone
//   "C"                = Charlie Military Time Zone
//   "Z"                = Zulu Military Time Zone
//
// If the current 'TimeZoneDefDto' instance is NOT configured as a Military
// Time Zone, this method will return an error.
//
func (tzdef *TimeZoneDefDto) GetMilitaryTimeZoneLetter() (string, error) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefDto.GetMilitaryTimeZoneLetter() "

	if tzdef.originalTimeZone.timeZoneType != TzType.Military() {
		return "",
			errors.New(ePrefix +
				"\nError: This TimeZoneDefDto instance is NOT configured as a Military Time Zone!\n" +
				"Therefore Military Time Zone Letter is invalid!\n")
	}

	return tzdef.originalTimeZone.militaryTimeZoneLetter, nil
}

// GetMilitaryTimeZoneName - Returns the a text name which represents
// the Military Time Zone designation.
//
// Examples:
//   "Alpha"                = Military Time Zone 'A'
//   "Bravo"                = Military Time Zone 'B'
//   "Charlie"              = Military Time Zone 'C'
//   "Zulu"                 = Military Time Zone 'Z'
//
// If the current 'TimeZoneDefDto' instance is NOT configured as a Military
// Time Zone, this method will return an error.
//
func (tzdef *TimeZoneDefDto) GetMilitaryTimeZoneName() (string, error) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefDto.GetMilitaryTimeZoneLetter() "

	if tzdef.originalTimeZone.timeZoneType != TzType.Military() {
		return "",
			errors.New(ePrefix +
				"\nError: This TimeZoneDefDto instance is NOT configured as a Military Time Zone!\n" +
				"Therefore Military Time Zone Name is invalid!\n")
	}

	return tzdef.originalTimeZone.militaryTimeZoneName, nil
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

	return tzdef.originalTimeZone.offsetHours
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

	return tzdef.originalTimeZone.offsetMinutes
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

	return tzdef.originalTimeZone.offsetSeconds
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

	return tzdef.originalTimeZone.tagDescription
}

// GetUtcOffset - Returns the offset from UTC as a string.
// Examples of the UTC offset format are: "-0600" or "+0200".
//
func (tzdef *TimeZoneDefDto) GetUtcOffset() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.utcOffset
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

	return tzdef.originalTimeZone.zoneName
}

// GetZoneOffset - Returns TimeZoneDefDto member variable
// ZoneOffset value.
//
// ZoneOffset is a text string representing the time zone.
// Example "-0600 CST" or "+0200 EET"
//
func (tzdef *TimeZoneDefDto) GetZoneOffset() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneOffset
}

// GetZoneOffsetSeconds - Returns TimeZoneDefDto member variable
// ZoneOffsetSeconds value.
//
func (tzdef *TimeZoneDefDto) GetZoneOffsetSeconds() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneOffsetTotalSeconds
}

// GetZoneSign - Returns TimeZoneDefDto member variable
// ZoneSign value.
//
// -1 == West of UTC  +1 == East of UTC
//
func (tzdef *TimeZoneDefDto) GetZoneSign() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneSignValue
}

// GetTimeZoneType - Returns the Time Zone Type associated
// with this instance of TimeZoneDefDto.
//
// Time Zone Type is styled as an enumeration.
func (tzdef *TimeZoneDefDto) GetTimeZoneType() TimeZoneType {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.timeZoneType
}

// IsEmpty - Determines whether the current TimeZoneDefDto
// instance is Empty.
//
// If the TimeZoneDefDto instance is NOT populated, the method
// returns 'true'. Otherwise, it returns 'false'.
//
func (tzdef *TimeZoneDefDto) IsEmpty() bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.isEmpty(tzdef)

}

// IsValid - Analyzes the current TimeZoneDefDto instance
// to determine validity.
//
// This method returns an error if the TimeZoneDefDto instance
// is INVALID.  Otherwise, it returns 'nil'.
//
func (tzdef *TimeZoneDefDto) IsValid() error {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefDto.IsValid() "

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.isValidTimeZoneDefDto(tzdef, ePrefix)
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
// This method will return two values:
//      (1) A Time Zone Definition Dto (TimeZoneDefDto)
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
//  (2)   If successful, this method will set the returned 'error' instance to 'nil'.
//        If errors are encountered a valid error message will be returned in the
//        error instance.
//
func (tzdef TimeZoneDefDto) New(dateTime time.Time) (TimeZoneDefDto, error) {

	ePrefix := "TimeZoneDefDto.New() "

	if dateTime.IsZero() {
		return TimeZoneDefDto{}, errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is a ZERO value!\n")
	}

	tzDef2 := TimeZoneDefDto{}

	tzDefUtil := timeZoneDefUtility{}


	err := tzDefUtil.setFromDateTime(&tzDef2, dateTime, ePrefix)

	// err := tzDefUtil.setFromTimeZoneName(&tzDef2, locationName, ePrefix)
	if err != nil {
		return TimeZoneDefDto{}, err
	}

	return tzDef2, nil
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
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType,) (
	tzDefDto TimeZoneDefDto,
	err error) {

	ePrefix := "TimeZoneDefDto.NewFromTimeZoneName() "

// 	var tzLoc *time.Location
	err = nil
	tzDefDto = TimeZoneDefDto{}

//	var err2 error

	if len(timeZoneName) == 0 {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'timeZoneName' is EMPTY!\n")
		return tzDefDto, err
	}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromTimeZoneName(
		&tzDefDto,
		dateTime,
		timeZoneName,
		timeConversionType,
		ePrefix)

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

	tzdef.originalTimeZone.tagDescription = tagDesc
	tzdef.convertibleTimeZone.tagDescription = tagDesc
}

// SetFromDateTimeComponents - Re-initializes the values of the current
// TimeZoneDefDto instance based on input parameter, 'dateTime'.
//
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

// SetFromTimeZoneName - Sets the data fields of the current
// TimeZoneDefDto instance based on the time zone text name
// passed as an input parameter.
//
func (tzdef *TimeZoneDefDto) SetFromTimeZoneName(
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType) error {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefDto.SetFromTimeZoneName() "

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.setFromTimeZoneName(
		tzdef,
		dateTime,
		timeZoneName,
		timeConversionType,
		ePrefix)
}
