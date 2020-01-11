package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// timeZoneSpecDto - Internal data structure used to
// store Time Zone data elements.
type TimeZoneSpecDto struct {
	zoneLabel              string           // Identifying Time Zone Label. A tag or description field.
	referenceDateTime      time.Time        // The date time used in defining the Time Zone
	zoneName               string           // The Time Zone abbreviation. Examples: 'EST', 'CST', 'PST'
	zoneOffsetTotalSeconds int              // Signed number of seconds offset from UTC.
	//                                         + == East of UTC; - == West of UTC
	zoneSignValue          int              // -1 == West of UTC  +1 == East of UTC. Apply this sign
	//                                         to the offset hours, minutes and seconds value
	offsetHours            int              // Normalized Offset Hours from UTC. Always a positive number,
	//                                         refer to ZoneSign for correct sign value.
	offsetMinutes          int              // Normalized Offset Minutes offset from UTC. Always a positive number,
	//                                         refer to ZoneSign for the correct sign value
	offsetSeconds          int              // Normalized Offset Seconds offset from UTC. Always a positive number,
	//                                         refer to ZoneSign for the correct sign value
	zoneOffset             string           // A text string representing offset from UTC for this time zone.
	//                                         Example "-0600 CST" or "+0200 EET"
	zoneAbbrvLookupId      string           // A string representing the Abbreviation Id used in map lookups.
	//                                         Examples: "CST-0600", "EET+0200"
	utcOffset              string           // A text string representing the offset from UTC for this Time Zone.
	//                                         Examples: "-0600" or "+0200"
	locationPtr            *time.Location   // Pointer to a Time Zone Location
	locationName           string           // Time Zone Location Name Examples: "Local", "America/Chicago", "America/New_York"
	locationNameType       LocationNameType // Four possible values:
	//                                           None()
	//                                           ConvertibleAbbreviation()
	//                                           NonConvertibleAbbreviation()
	//                                           ConvertibleTimeZoneName()
	militaryTimeZoneName   string           // Full Military Time Zone text name. Examples: "Alpha", "Bravo", "Charlie", "Zulu"
	militaryTimeZoneLetter string           // Single Alphabetic Character identifying a Military Time Zone.
	tagDescription         string           // Unused - Available for classification, labeling or description by user.
	timeZoneType           TimeZoneType     // Enumeration of Time Zone Type:
	//                                          TzType.None()
	//                                          TzType.Iana()
	//                                          TzType.Military()
	//                                          TzType.Local()
	//                                          TzType.UtcOffset()
	lock sync.Mutex // Used for implementing thread safe operations.
}

// CopyIn - Copies the values of input parameter 'tzSpec2'
// to all of the data fields in the current instance of 
// TimeZoneSpecDto (tzSpec). When completed 'tzSpec' will
// have data field values identical to those of 'tzSpec2'
//
func (tzSpec *TimeZoneSpecDto) CopyIn(tzSpec2 TimeZoneSpecDto) {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpecUtil := typeZoneSpecUtility{}

	tzSpecUtil.copyIn(tzSpec, tzSpec2)

}


// CopyOut - Returns a deep copy of the current Time Zone 
// Specification object as a new instance of 'TimeZoneSpecDto'.
//
func (tzSpec *TimeZoneSpecDto) CopyOut() TimeZoneSpecDto {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpecUtil := typeZoneSpecUtility{}

	return tzSpecUtil.copyOut(tzSpec)
}

// Empty - Sets all the values of the data fields in the
// current TimeZoneSpecDto to their empty or zero values.
//
func (tzSpec *TimeZoneSpecDto) Empty() {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpecUtil := typeZoneSpecUtility{}

	tzSpecUtil.empty(tzSpec)

}

// Equal - Returns a boolean value of true if both the current instance
// of TimeZoneSpecDto and the input parameter TimeZoneSpecDto are
// equivalent in all respects.
//
func (tzSpec *TimeZoneSpecDto) Equal( tzSpec2 TimeZoneSpecDto) bool {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	if !tzSpec.referenceDateTime.Equal(tzSpec2.referenceDateTime) {
		return false
	}

	if tzSpec.zoneLabel != tzSpec2.zoneLabel ||
		tzSpec.zoneName != tzSpec2.zoneName ||
		tzSpec.zoneOffsetTotalSeconds != tzSpec2.zoneOffsetTotalSeconds ||
		tzSpec.zoneSignValue != tzSpec2.zoneSignValue {
		return false
	}

	if tzSpec.offsetHours != tzSpec2.offsetHours ||
		tzSpec.offsetMinutes != tzSpec2.offsetMinutes ||
		tzSpec.offsetSeconds != tzSpec2.offsetSeconds {
		return false
	}


	if tzSpec.zoneOffset != tzSpec2.zoneOffset ||
		tzSpec.zoneAbbrvLookupId != tzSpec2.zoneAbbrvLookupId ||
		tzSpec.utcOffset != tzSpec2.utcOffset {
		return false
	}

	if tzSpec.locationPtr == nil && tzSpec2.locationPtr != nil{
		return false
	}

	if tzSpec.locationPtr != nil && tzSpec2.locationPtr == nil {
		return false
	}

	if tzSpec.locationPtr != nil && tzSpec2.locationPtr != nil &&
		tzSpec.locationPtr.String() != tzSpec2.locationPtr.String() {
		return false
	}

	if tzSpec.locationName != tzSpec2.locationName {
		return false
	}

	if tzSpec.militaryTimeZoneLetter != tzSpec2.militaryTimeZoneLetter ||
		tzSpec.militaryTimeZoneName != tzSpec2.militaryTimeZoneName {
		return false
	}

	if tzSpec.locationNameType != tzSpec.locationNameType {
		return false
	}

	if tzSpec.timeZoneType != tzSpec2.timeZoneType {
		return false
	}

	if tzSpec.tagDescription != tzSpec2.tagDescription {
		return false
	}

	return true

}

// IsEmpty() returns a boolean value of 'true' if all
// data field values are set to their empty or zero
// values.
func (tzSpec *TimeZoneSpecDto) IsEmpty() bool {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	if 	!tzSpec.referenceDateTime.IsZero() {
		return false
	}

	if 	tzSpec.zoneOffsetTotalSeconds != 0 ||
		tzSpec.zoneSignValue != 0 ||
		tzSpec.offsetHours != 0 ||
		tzSpec.offsetMinutes != 0 ||
		tzSpec.offsetSeconds != 0 {
		return false
	}


	if tzSpec.zoneLabel != "" ||
	tzSpec.zoneName != "" ||
	tzSpec.zoneOffset != "" ||
	tzSpec.zoneAbbrvLookupId != "" ||
	tzSpec.utcOffset != "" {
		return false
	}

	if tzSpec.locationPtr != nil ||
		tzSpec.locationName != "" {
		return false
	}

	if tzSpec.militaryTimeZoneName != "" ||
		tzSpec.militaryTimeZoneLetter != "" {
		return false
	}

	if tzSpec.locationNameType != LocNameType.None() ||
		tzSpec.timeZoneType != TzType.None(){
		return false
	}

	if tzSpec.tagDescription != "" {
		return false
	}

	return true
}

// IsValid - Examines the data fields of the current
// TimeZoneSpecDto instance are valid.
//
func (tzSpec *TimeZoneSpecDto) IsValid(ePrefix string) error {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	ePrefix += "TimeZoneSpecDto.IsValid() "

	if strings.TrimLeft(strings.TrimRight(tzSpec.locationName, " "), " ") == "" {
		return errors.New(ePrefix +
			"\nError: locationName is an empty string!\n")
	}

	if tzSpec.locationPtr == nil {
		return errors.New(ePrefix +
			"\nError: Location Pointer is 'nil'!\n")
	}

	if tzSpec.locationPtr.String() != tzSpec.locationName {
		return fmt.Errorf(ePrefix +
			"\nError: The Location Pointer is NOT equal to the Location Name!\n" +
			"Location Pointer String='%v'\n" +
			"Location Name = '%v'\n",
			tzSpec.locationPtr.String() , tzSpec.locationName)
	}

	dtMech := dateTimeMechanics{}

	locPtr, err := dtMech.loadTzLocationPtr(tzSpec.locationName, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: Location Name is NOT a valid time zone!\n"+
			"tzdef.locationName='%v'\n"+
			"Returned Error='%v'\n", tzSpec.locationName, err.Error())
	}

	if locPtr.String() != tzSpec.locationName {
		return fmt.Errorf(ePrefix +
			"\nError: LoadLocation Pointer string NOT equal to tzSpec.locationName !\n" +
			"tzSpec.locationName='%v'\n" +
			"loc.String()='%v'\n", tzSpec.locationName, locPtr.String())
	}

	if tzSpec.timeZoneType == TzType.Military() &&
		(tzSpec.militaryTimeZoneLetter == "" ||
			tzSpec.militaryTimeZoneName == "") {
		return fmt.Errorf(ePrefix +
			"\nError: This time zone is classified as a 'Military' Time Zone.\n" +
			"However, one or both of the Military Time Zone name strings are empty.\n" +
			"tzSpec.militaryTimeZoneLetter='%v'\n" +
			"tzSpec.militaryTimeZoneName='%v'\n",
			tzSpec.militaryTimeZoneLetter , tzSpec.militaryTimeZoneName)
	}

	return nil
}

// GetLocationPointer - Returns the time zone location in the form of
// a pointer to 'time.Location'.
//
func (tzSpec *TimeZoneSpecDto) GetLocationPointer() *time.Location {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.locationPtr
}

// GetLocationName - Returns the time zone name or time zone location.
// Examples: "Local", "America/Chicago", "America/New_York"
//
func (tzSpec *TimeZoneSpecDto) GetLocationName() string {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.locationName
}

// GetLocationNameType - Describes and classifies the Time Zone
// Location. The return value is a LocationNameType value which
// is an enumeration time zone location name classifications.
//
// Possible return values:
//
//    ConvertibleAbbreviation    - Time Zone Zone Location Name is an
//                                 abbreviation which is valid and
//                                 convertible across all other time
//                                 zones.
//
//    NonConvertibleAbbreviation - The Time Zone Location Name cannot
//                                 be converted to other time zones.
//
//    ConvertibleTimeZoneName    - The Time Zone Name is a complete
//                                 and valid time zone name which is
//                                 convertible across all other
//                                 time zones.
//
func (tzSpec *TimeZoneSpecDto) GetLocationNameType() LocationNameType {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.locationNameType
}

// GetMilitaryTimeZoneName - Returns a string containing the military
// time zone name, if applicable. If the current TimeZoneSpecDto
// instance does not define a military time zone, this return value
// is an empty string.
//
func (tzSpec *TimeZoneSpecDto) GetMilitaryTimeZoneName() string {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.militaryTimeZoneName
}

// GetMilitaryTimeZoneLetter - Returns a string containing the military
// time zone letter or abbreviation. If the current TimeZoneSpecDto
// instance does not define a military time zone, this return value
// is an empty string.
//
func (tzSpec *TimeZoneSpecDto) GetMilitaryTimeZoneLetter() string {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.militaryTimeZoneLetter
}

// GetOffsetElements - Returns a series of integer values which
// taken collectively identify the offset from UTC for this time
// zone.
//
// ------------------------------------------------------------
//
// Return Values
// =============
//
// zoneSignValue int  -1 == West of UTC  +1 == East of UTC. Apply this sign
//                    to the offset hours, minutes and seconds value
//
// offsetHours   int  Normalized Offset Hours from UTC. Always a positive number,
//                    refer to ZoneSign for correct sign value.
//
// offsetMinutes int  Normalized Offset Minutes offset from UTC. Always a
//                    positive number, refer to ZoneSign for the correct
//                    sign value.
//
// offsetSeconds int  Normalized Offset Seconds offset from UTC. Always a
//                    positive number, refer to ZoneSign for the correct
//                    sign value.
//
func (tzSpec *TimeZoneSpecDto) GetOffsetElements() (
	zoneSignValue,
	offsetHours,
	offsetMinutes,
	offsetSeconds int) {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneSignValue,
		tzSpec.offsetHours,
		tzSpec.offsetMinutes,
		tzSpec.offsetSeconds
}

// GetReferenceDateTime - Returns the reference Date Time
//
func (tzSpec *TimeZoneSpecDto) GetReferenceDateTime() time.Time {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.referenceDateTime
}

// GetTagDescription - Returns the private member variable
// "tagDescription". This field is available for users to
// tag, classify or otherwise attach descriptive information
// to this TimeZoneSpecDto instance.
//
func (tzSpec *TimeZoneSpecDto) GetTagDescription() string {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.tagDescription
}

// GetTimeZoneType - Returns the Time Zone Type classification.
// Time Zone Type is an enumeration identifying the time zone
// source.
//
// Possible return types.
//
//  TzType.Iana()      - Identifies an IANA Time Zone
//  TzType.Military()  - Identifies a Military Time Zone
//  TzType.Local()     - Identifies this as a 'Local' Time Zone
//  TzType.UtcOffset() - Identifies this time zone a UTC Offset
//
func (tzSpec *TimeZoneSpecDto) GetTimeZoneType() TimeZoneType {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.timeZoneType
}

// GetUtcOffset - returns a text string representing the
// offset from UTC for this time zone.
//
//  Examples: "-0600", "+0200"
//
func (tzSpec *TimeZoneSpecDto) GetUtcOffset() string {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.utcOffset
}

// GetZoneAbbrvLookupId - Returns a text string containing the
// Time Zone abbreviation plus the UTC offset. This text value
// is used to look up time zone data in various internal data
// maps. Examples: "CST-0600", "EET+0200"
//
// Note: To access the time zone abbreviation, see method
// TimeZoneSpecDto.GetZoneName()
//
func (tzSpec *TimeZoneSpecDto) GetZoneAbbrvLookupId() string {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneAbbrvLookupId
}

// GetZoneLabel - Returns the Zone Label, a tag or text
// description field available for use by the user.
//
func (tzSpec *TimeZoneSpecDto) GetZoneLabel() string {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneLabel
}

// GetZoneName - Returns the 'Zone Name'. 'Zone Name' is the
// the Time Zone abbreviation. Examples: 'EST', 'CST', 'PST'
//
func (tzSpec *TimeZoneSpecDto) GetZoneName() string {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneName
}

// GetZoneOffset - Returns data field 'zoneOffset'. This is
// a text string representing the offset from UTC for this
// time zone. The returned offset string consists of two
// components, the hours and minutes of offset and the time
// zone abbreviation.
//
// Example: "-0600 CST" or "+0200 EET"
//
func (tzSpec *TimeZoneSpecDto) GetZoneOffset() string {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneOffset
}

// GetZoneOffsetTotalSeconds - Returns the total offset seconds
// from 'UTC' for this time zone. The returned value is a signed
// value. Positive ('+') values identify seconds East of UTC.
// Negative ('-') values identify seconds West of UTC.
//
func (tzSpec *TimeZoneSpecDto) GetZoneOffsetTotalSeconds() int {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneOffsetTotalSeconds
}

// GetZoneSignValue - Returns the data field 'zoneSignValue'. This
// is a signed integer value identifying whether the time zone
// offset from UTC is east or west of UTC. The returned integer
// will hold one of only two values: a positive '1' or a negative
// '-1'.  '-1' indicates an offset West of UTC while a positive
// '1' identifies and offset East of UTC. Apply this sign to the
// unsigned values for offset hours, minutes and seconds.
//
func (tzSpec *TimeZoneSpecDto) GetZoneSignValue() int {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	return tzSpec.zoneSignValue
}

// New - Returns a new instance of TimeZoneSpecDto.
//
func (tzSpec TimeZoneSpecDto) New(
	referenceDateTime      time.Time,
	militaryTimeZoneName   string,
	militaryTimeZoneLetter string,
	zoneLabel              string,
	tagDescription         string,
	locationNameType       LocationNameType,
	timeZoneType           TimeZoneType,
	ePrefix string) (TimeZoneSpecDto, error) {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	ePrefix += "TimeZoneSpecDto.New() "

	tzSpecOut := TimeZoneSpecDto{}

	err := tzSpecOut.SetTimeZone(
		referenceDateTime,
		militaryTimeZoneName,
		militaryTimeZoneLetter,
		zoneLabel,
		tagDescription,
		locationNameType,
		timeZoneType,
		ePrefix)


	return tzSpecOut, err
}
// SetZoneLabel - Sets the value of data field "Zone Label". 'Zone
// Label' a tag or text description field available for use
// by the user.
//
func (tzSpec *TimeZoneSpecDto) SetZoneLabel() string {
	return tzSpec.zoneLabel
}

// SetTimeZone - Sets the data values of the current Time Zone
// Specification Structure (TimeZoneSpecDto).
//
func (tzSpec *TimeZoneSpecDto) SetTimeZone(
	referenceDateTime      time.Time,
	militaryTimeZoneLetter string,
	militaryTimeZoneName   string,
	zoneLabel              string,
	tagDescription         string,
	locationNameType       LocationNameType,
	timeZoneType           TimeZoneType,
	ePrefix string) error {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	ePrefix += "TimeZoneSpecDto.SetTimeZone() "

	if referenceDateTime.IsZero() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "referenceDateTime",
			inputParameterValue: "",
			errMsg:              "'referenceDateTime' has a zero value.",
			err:                 nil,
		}
	}

	tzSpecUtil := typeZoneSpecUtility{}

	tzSpecUtil.empty(tzSpec)

	tzMech := timeZoneMechanics{}
  var err error
	
	tzSpec.zoneName,
		tzSpec.zoneOffset,
		tzSpec.utcOffset,
		tzSpec.zoneAbbrvLookupId,
		tzSpec.offsetHours,
		tzSpec.offsetMinutes,
		tzSpec.offsetSeconds,
		tzSpec.zoneSignValue,
		tzSpec.zoneOffsetTotalSeconds,
		tzSpec.locationPtr,
		tzSpec.locationName,
		err = tzMech.calcUtcZoneOffsets(referenceDateTime, ePrefix)
	
	 if err != nil {
	 	return err
	 }
	
	tzSpec.referenceDateTime = referenceDateTime
	tzSpec.zoneLabel = zoneLabel
	tzSpec.militaryTimeZoneLetter = militaryTimeZoneLetter
	tzSpec.militaryTimeZoneName = militaryTimeZoneName
	tzSpec.tagDescription = tagDescription
	tzSpec.locationNameType = locationNameType
	tzSpec.timeZoneType = timeZoneType

	return nil
}