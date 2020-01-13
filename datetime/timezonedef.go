package datetime

import (
	"errors"
	"sync"
	"time"
)

// TimeZoneDefinition - Time Zone definition.
//
// Overview and General Usage
//
// A Time Zone Definition includes two Time Zone Specifications
// (TimeZoneSpec
// This structure is designed to store detailed information
// on Time Zones. It used primarily as a 'helper' or subsidiary
// structure for Type, 'TimeZoneDto'.
//
// The TimeZoneDefinition method 'New' may be used to extract
// detailed time zone information from a time.Time date time
// value.
//
//
// TimeZoneDefinition - Time Zone Definition Dto
// Contains detailed parameters describing a specific
// Time Zone and Time Zone Location
//
// Source Code Location
//
// This source file is located in source code repository:
//   https://github.com/MikeAustin71/datetimeopsgo.git'
//
// This source code file is located at:
//   MikeAustin71\datetimeopsgo\datetime\timezonedef.go
//
type TimeZoneDefinition struct {
	originalTimeZone    TimeZoneSpecification // The Time Zone Specification originally submitted.
	convertibleTimeZone TimeZoneSpecification // A version of the original time zone with a new fully
	//                                         convertible time zone substituted.
	lock sync.Mutex // Used for implementing thread safe operations.
}

// CopyIn - Copies an incoming TimeZoneDefinition into the
// data fields of the current TimeZoneDefinition instance.
func (tzdef *TimeZoneDefinition) CopyIn(tzdef2 TimeZoneDefinition) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	tzDefUtil.copyIn(tzdef, &tzdef2)

	return
}

// copyOut - creates and returns a deep copy of the current
// TimeZoneDefinition instance.
func (tzdef *TimeZoneDefinition) CopyOut() TimeZoneDefinition {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.copyOut(tzdef)
}

// Empty - Resets all field values for the current TimeZoneDefinition
// instance to their uninitialized or 'zero' states.
func (tzdef *TimeZoneDefinition) Empty() {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	tzDefUtil.empty(tzdef)
}

// Equal - Determines if two TimeZoneDefinition are equivalent in
// value. Returns 'true' of two TimeZoneDefinition's are equal in
// all respects.
func (tzdef *TimeZoneDefinition) Equal(tzdef2 TimeZoneDefinition) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.equal(tzdef, &tzdef2)
}

// EqualOffsetSeconds - Compares Zone Offset Seconds for two TimeZoneDefinition's and
// returns 'true' if they are equal.
//
// ZoneOffsetSeconds is a signed number of seconds offset from UTC:
//   + == East of UTC
//   - == West of UTC
func (tzdef *TimeZoneDefinition) EqualOffsetSeconds(tzdef2 TimeZoneDefinition) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.equalOffsetSeconds(tzdef, &tzdef2)
}

// EqualZoneOffsets - Compares ZoneOffsets for two TimeZoneDefinition's and
// returns 'true' if they are equal.
//
// Zone Offset is a text string representing the offset from UTC plus the
// time zone abbreviation.
//
// Example "-0500 CDT"
//
func (tzdef *TimeZoneDefinition) EqualZoneOffsets(tzdef2 TimeZoneDefinition) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.equalZoneOffsets(tzdef, &tzdef2)
}

// EqualLocations - Compares the Time Zone Locations for two TimeZoneDefinition's
// and returns 'true' if they are equal.
//
// Time Zone Location Name Examples:
//   "Local"
//   "America/Chicago"
//   "America/New_York"
//
func (tzdef *TimeZoneDefinition) EqualLocations(tzdef2 TimeZoneDefinition) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.equalLocations(tzdef, &tzdef2)
}

// EqualZoneLocation - Compares two TimeZoneDefinition's and returns
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
func (tzdef *TimeZoneDefinition) EqualZoneLocation(tzdef2 TimeZoneDefinition) bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	if !tzDefUtil.equalZoneLocation(tzdef, &tzdef2) {
		return false
	}

	return true
}

// GetLocationPtr - Returns a pointer to a time.Location or
// time zone location (TimeZoneDefinition.Location).
//
func (tzdef *TimeZoneDefinition) GetLocationPtr() *time.Location {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.locationPtr
}

// GetLocationName - Returns TimeZoneDefinition member variable
// LocationName value.
//
// Time Zone Location Name Examples: "Local", "America/Chicago",
// "America/New_York".
//
func (tzdef *TimeZoneDefinition) GetLocationName() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.locationName
}

// GetLocationNameType - Returns the value of internal
// private member variable TimeZoneDefinition.locationNameType.
//
// Possible return values:
//  LocationNameType(0).None()
//  LocationNameType(0).ConvertibleAbbreviation()
//  LocationNameType(0).NonConvertibleAbbreviation()
//  LocationNameType(0).ConvertibleTimeZoneName()
//
func (tzdef *TimeZoneDefinition) GetLocationNameType() LocationNameType {

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
// If the current 'TimeZoneDefinition' instance is NOT configured as a Military
// Time Zone, this method will return an error.
//
func (tzdef *TimeZoneDefinition) GetMilitaryTimeZoneLetter() (string, error) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.GetMilitaryTimeZoneLetter() "

	if tzdef.originalTimeZone.timeZoneType != TzType.Military() {
		return "",
			errors.New(ePrefix +
				"\nError: This TimeZoneDefinition instance is NOT configured as a Military Time Zone!\n" +
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
// If the current 'TimeZoneDefinition' instance is NOT configured as a Military
// Time Zone, this method will return an error.
//
func (tzdef *TimeZoneDefinition) GetMilitaryTimeZoneName() (string, error) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.GetMilitaryTimeZoneLetter() "

	if tzdef.originalTimeZone.timeZoneType != TzType.Military() {
		return "",
			errors.New(ePrefix +
				"\nError: This TimeZoneDefinition instance is NOT configured as a Military Time Zone!\n" +
				"Therefore Military Time Zone Name is invalid!\n")
	}

	return tzdef.originalTimeZone.militaryTimeZoneName, nil
}

// GetOffsetHours - Returns TimeZoneDefinition member variable
// ZoneOffset value.
//
// Normalized Offset Hours from UTC. Always a positive number,
// refer to ZoneSign to determine East or West of UTC.
//
func (tzdef *TimeZoneDefinition) GetOffsetHours() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.offsetHours
}

// GetOffsetMinutes - Returns TimeZoneDefinition member variable
// ZoneOffsetMinutes value.
//
// Normalized Offset Minutes offset from UTC. Always a
// positive number, refer to ZoneSign to determine East or
// West of UTC.
//
func (tzdef *TimeZoneDefinition) GetOffsetMinutes() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.offsetMinutes
}

// GetOffsetSeconds - Returns TimeZoneDefinition member variable
// ZoneOffsetSeconds value.
//
// Normalized Offset Seconds offset from UTC. Always a
// positive number, refer to ZoneSign to determine East or
// West of UTC.
//
func (tzdef *TimeZoneDefinition) GetOffsetSeconds() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.offsetSeconds
}

// GetConvertibleTimeZone - Returns the convertible form of
// the current time zone. This time zone should be convertible
// across all time zones. To verify convertibility, call
// TimeZoneSpecification.GetLocationNameType() and inspect the result.
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZone() TimeZoneSpecification {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone

}

// GetOriginalTimeZone - Returns the originally configured
// time zone specification. Check TimeZoneSpecification.GetLocationNameType()
// if this time zone is fully convertible across all time zones.
// If not convertible, use TimeZoneDefinition.GetConvertibleTimeZone().
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZone() TimeZoneSpecification {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone
}

// GetTagDescription - Returns TimeZoneDefinition member variable
// Description value.
//
// Unused - Available to user. This string is typically used
// classification, labeling or description text by user.
//
func (tzdef *TimeZoneDefinition) GetTagDescription() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.tagDescription
}

// GetUtcOffset - Returns the offset from UTC as a string.
// Examples of the UTC offset format are: "-0600" or "+0200".
//
func (tzdef *TimeZoneDefinition) GetUtcOffset() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.utcOffset
}

// GetZoneName - Returns TimeZoneDefinition member variable
// TimeZoneDefinition.zoneName value.
//
// Zone Name is the Time Zone abbreviation. This may
// may be a series of characters, like "EST", "CST"
// and "PDT" - or - if a time zone abbreviation does
// not exist for this time zone, the time zone abbreviation
// might be listed simply as the UTC offset ('-0430')
//
func (tzdef *TimeZoneDefinition) GetZoneName() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneName
}

// GetZoneOffset - Returns TimeZoneDefinition member variable
// ZoneOffset value.
//
// ZoneOffset is a text string representing the time zone.
// Example "-0600 CST" or "+0200 EET"
//
func (tzdef *TimeZoneDefinition) GetZoneOffset() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneOffset
}

// GetZoneOffsetSeconds - Returns TimeZoneDefinition member variable
// ZoneOffsetSeconds value.
//
func (tzdef *TimeZoneDefinition) GetZoneOffsetSeconds() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneOffsetTotalSeconds
}

// GetZoneSign - Returns TimeZoneDefinition member variable
// ZoneSign value.
//
// -1 == West of UTC  +1 == East of UTC
//
func (tzdef *TimeZoneDefinition) GetZoneSign() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneSignValue
}

// GetTimeZoneType - Returns the Time Zone Type associated
// with this instance of TimeZoneDefinition.
//
// Time Zone Type is styled as an enumeration.
func (tzdef *TimeZoneDefinition) GetTimeZoneType() TimeZoneType {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.timeZoneType
}

// IsEmpty - Determines whether the current TimeZoneDefinition
// instance is Empty.
//
// If the TimeZoneDefinition instance is NOT populated, the method
// returns 'true'. Otherwise, it returns 'false'.
//
func (tzdef *TimeZoneDefinition) IsEmpty() bool {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.isEmpty(tzdef)

}

// IsValid - Analyzes the current TimeZoneDefinition instance
// to determine validity.
//
// This method returns an error if the TimeZoneDefinition instance
// is INVALID.  Otherwise, it returns 'nil'.
//
func (tzdef *TimeZoneDefinition) IsValid() error {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.IsValid() "

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.isValidTimeZoneDefDto(tzdef, ePrefix)
}

// New - Creates and returns a new TimeZoneDefinition instance based on
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
//      (1) A Time Zone Definition Dto (TimeZoneDefinition)
//      (2) An 'error' type
//
//  (1) If successful, this method will return a valid, populated TimeZoneDefinition instance.
//      A TimeZoneDefinition is defined as follows:
//      type TimeZoneDefinition struct {
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
func (tzdef TimeZoneDefinition) New(dateTime time.Time) (TimeZoneDefinition, error) {

	ePrefix := "TimeZoneDefinition.New() "

	if dateTime.IsZero() {
		return TimeZoneDefinition{}, errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is a ZERO value!\n")
	}

	tzDef2 := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}


	err := tzDefUtil.setFromDateTime(&tzDef2, dateTime, ePrefix)

	// err := tzDefUtil.setFromTimeZoneName(&tzDef2, locationName, ePrefix)
	if err != nil {
		return TimeZoneDefinition{}, err
	}

	return tzDef2, nil
}

// NewFromTimeZoneName - Creates and returns a new instance 'TimeZoneDefinition'.
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
func (tzdef TimeZoneDefinition) NewFromTimeZoneName(
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType,) (
	tzDefDto TimeZoneDefinition,
	err error) {

	ePrefix := "TimeZoneDefinition.NewFromTimeZoneName() "

// 	var tzLoc *time.Location
	err = nil
	tzDefDto = TimeZoneDefinition{}

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

// SetTagDescription - Sets TimeZoneDefinition private member variable
// TimeZoneDefinition.tagDescription to the value passed in 'tagDesc'.
//
// The TimeZoneDefinition.tagDescription string is available to users
// for use as a tag, label, classification or text description.
//
func (tzdef *TimeZoneDefinition) SetTagDescription(tagDesc string) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzdef.originalTimeZone.tagDescription = tagDesc
	tzdef.convertibleTimeZone.tagDescription = tagDesc
}

// SetFromDateTimeComponents - Re-initializes the values of the current
// TimeZoneDefinition instance based on input parameter, 'dateTime'.
//
func (tzdef *TimeZoneDefinition) SetFromDateTime(dateTime time.Time) error {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.SetFromDateTimeComponents() "

	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'dateTime' is a ZERO value!")
	}

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.setFromDateTime(tzdef, dateTime, ePrefix)

}

// SetFromTimeZoneName - Sets the data fields of the current
// TimeZoneDefinition instance based on the time zone text name
// passed as an input parameter.
//
func (tzdef *TimeZoneDefinition) SetFromTimeZoneName(
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType) error {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.SetFromTimeZoneName() "

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.setFromTimeZoneName(
		tzdef,
		dateTime,
		timeZoneName,
		timeConversionType,
		ePrefix)
}
