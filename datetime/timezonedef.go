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
// (Type: TimeZoneSpecification). Both describe the same time zone.
// The first 'TimeZoneSpecification' describes the original time
// zone used to initialize the current TimeZoneDefinition instance.
// The second is provided to ensure the configuration of a time zone
// which can be used to convert date times to another time zone
// anywhere in the world.
//
//
// Why Two Time Zone Specifications?
//
// Providing two time zone specifications deals with errors which
// may be introduced when using the Golang time package to parse
// date time strings in order to generate date time objects (time.Time).
// In the following example, assume you, and your computer, live in
// New York, NewYork USA. (I'll explain later). Now, create a new
// date time object (time.Time) by parsing the following date time
// string, "06/20/2019 09:58:32 -0700 PDT". This is a date time
// using 'Pacific Daylight Time' or 'PDT'. Notice how the string
// is parsed to identify a time zone.
//
//   tstr := "06/20/2019 09:58:32 -0700 PDT"
//   fmtStr := "01/02/2006 15:04:05 -0700 MST"
//   t1, err := time.Parse(fmtstr, t1str)
//
//   Note: 'time.Parse' will NOT return an 'error'. The resulting
//   't1' will be created with a time.Location (a.k.a time zone)
//   of "PDT". However, "PDT" is NOT a valid IANA time zone. If
//   you try to create a new date time ('t2') using the 't1' location
//   pointer ('t1.Location()') it could well generate an invalid
//   date time and time zone.
//
//   t2 := time.Date(
//         2019,
//         time.Month(12),
//         30,
//          9,
//          0,
//          0,
//          0,
//          t1.Location())
//
//   't2' is created as "12/30/2019 09:00:00 -0700 PDT". This is
//   WRONG. The correct time value is "12/30/2019 09:00:00 -0800 PST".
//   'PST' stands for 'Pacific Standard Time'.
//
//   Remember that I asked that you assume both you and your computer
//   live in New York, New York USA. If in fact you lived in Los
//   Angeles, California USA, parsing date time string,
//   "06/20/2019 09:58:32 -0700 PDT", Golang would return a valid time
//   zone of "Local", and the error described would not exist.
//
// Because of the possibility of generating invalid time zones, Type
// 'TimeZoneDefinition' contains two Time Zone Specifications. The
// first contains the 'Original Time Zone' values. The second contains
// 'Convertible Time Zone' values. 'Convertible' values are valid times
// which can be safely used across all world time zones. 'Convertible'
// time zones can be used to accurately convert a given date time to
// other time zones.
//
// Each Time Zone Specification object includes a flag indicating
// whether that time zone is 'Convertible'. This flag or indicator
// is labeled "location Name Type". See type 'LocationNameType'
// in source file:
//   'MikeAustin71\datetimeopsgo\datetime\locationnametypeenum.go'
//
// If the time zone is NOT convertible, the Location Name Type will
// be assigned an enumeration value of "NonConvertibleTimeZone".
// Conversely, if the time IS convertible, it will be assigned a
// Location Name Type of 'ConvertibleTimeZone'.
//
// The TimeZoneDefinition Type includes 'Getter' methods which will
// allow calling functions to extract 'Original' Time Zone data or
// 'Convertible' Time Zone data.
//
// Reference:
//   Golang Time Package: https://golang.org/pkg/time/
//   IANA Time Zones:
//     https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//     https://en.wikipedia.org/wiki/Tz_database
//     https://www.iana.org/time-zones
//     https://data.iana.org/time-zones/releases/
//
//
// Dependencies
//
// The 'TimeZoneDefinition' Type is dependent on type
// 'TimeZoneSpecification'.
//  Source Code File:
//    MikeAustin71/datetimeopsgo/datetime/timezonespecification.go
//
// In addition, 'TimeZoneDefinition' methods rely on
// the utility methods found in type 'timeZoneDefUtility'.
//
//  Source Code File:
//    MikeAustin71\datetimeopsgo\datetime\timezonedefutility.go
//
// The 'TimeZoneDefinition' makes use of two enumeration types,
// 'LocationNameType' and 'TimeZoneType'.
// Source Code Files:
//   'MikeAustin71\datetimeopsgo\datetime\locationnametypeenum.go'
//   'MikeAustin71\datetimeopsgo\datetime\timezonetypeenum.go'
//
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

// GetConvertibleDateTime - Returns the Convertible Time Zone reference
// date time value.
//
func (tzdef *TimeZoneDefinition) GetConvertibleDateTime() time.Time {
	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

return tzdef.convertibleTimeZone.referenceDateTime
}


// GetConvertibleLocationPtr - Returns a pointer to a time.Location for
// the convertible time zone (TimeZoneDefinition.convertibleTimeZone.locationPtr).
//
// The Time Zone 'Location' represents the Time Zone Name.
//
func (tzdef *TimeZoneDefinition) GetConvertibleLocationPtr() *time.Location {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.locationPtr
}

// GetConvertibleLocationName - Returns convertible Time Zone
// Name. This is also referred to as the "Location Name". This
// private member variable is TimeZoneDefinition.convertibleTimeZone.locationName.
//
// Time Zone Location Name Examples: "Local", "America/Chicago",
// "America/New_York".
//
func (tzdef *TimeZoneDefinition) GetConvertibleLocationName() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.locationName
}

// GetConvertibleLocationNameType - Returns the Location Name Type
// value for the Convertible Time Zone. This is private
// member variable:
//
//  TimeZoneDefinition.convertibleTimeZone.locationNameType.
//
// Possible return values:
//  LocationNameType(0).None()
//  LocationNameType(0).NonConvertibleTimeZone()
//  LocationNameType(0).ConvertibleTimeZone()
//
func (tzdef *TimeZoneDefinition) GetConvertibleLocationNameType() LocationNameType {

	return tzdef.convertibleTimeZone.locationNameType
}

// GetConvertibleOffsetElements - Returns a series of string and
// integer values which taken collectively identify the offset
// from UTC for the Convertible Time Zone.
//
// ------------------------------------------------------------
//
// Return Values
// =============
//
// offsetSignChar string - Like return value offsetSignValue, this string
//                         value signals whether the offset from UTC is West
//                         or East of UTC. This string will always have one of
//                         two values: "+" or "-". The plus sign ("+") signals
//                         that the offset is East of UTC. The minus sign ("-")
//                         signals that the offset is West of UTC.
//
// offsetSignValue int   - Similar to return value 'offsetSignChar' above except
//                         that sign values are expressed as either a '-1' or positive
//                         '1' integer value. -1 == West of UTC  +1 == East of UTC.
//                         Apply this sign value to the offset hours, minutes and
//                         seconds value returned below.
//
// offsetHours     int   - Normalized Offset Hours from UTC. Always a positive number,
//                         refer to ZoneSign for correct sign value.
//
// offsetMinutes   int   - Normalized Offset Minutes offset from UTC. Always a
//                         positive number, refer to ZoneSign for the correct
//                         sign value.
//
// offsetSeconds   int   - Normalized Offset Seconds offset from UTC. Always a
//                         positive number, refer to ZoneSign for the correct
//                         sign value.
//
func (tzdef *TimeZoneDefinition) GetConvertibleOffsetElements() (
	offsetSignChar string,
	offsetSignValue,
	offsetHours,
	offsetMinutes,
	offsetSeconds int) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetOffsetElements()
}

// GetConvertibleOffsetSignValue - Returns the Original Time Zone
// Offset Sign as an integer value. The returned integer will
// have one of two values: '-1' or the positive value, '1'.
//
// A negative value (-1) signals that the time zone offset
// from UTC is West of UTC.  Conversely, a positive value
// (+1) signals that the offset is East of UTC.
//
// -1 == West of UTC  +1 == East of UTC
//
func (tzdef *TimeZoneDefinition) GetConvertibleOffsetSignValue() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneSignValue
}

// GetConvertibleTimeZone - Returns the convertible form of
// the current time zone. This time zone should be convertible
// across all time zones. To verify convertibility, call
// TimeZoneSpecification.GetOriginalLocationNameType() and inspect the result.
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZone() TimeZoneSpecification {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone
}

// GetOriginalDateTime - Returns the Original Time Zone reference
// date time value.
//
func (tzdef *TimeZoneDefinition) GetOriginalDateTime() time.Time {
	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.referenceDateTime
}

// GetOriginalLocationPtr - Returns a pointer to a time.Location for
// the original time zone (TimeZoneDefinition.originalTimeZone.locationPtr).
//
func (tzdef *TimeZoneDefinition) GetOriginalLocationPtr() *time.Location {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.locationPtr
}

// GetOriginalLocationName - Returns original Time Zone
// Name. This is also referred to as the "Location Name". This
// private member variable is TimeZoneDefinition.convertibleTimeZone.locationName.
//
// Time Zone Location Name Examples: "Local", "America/Chicago",
// "America/New_York".
//
func (tzdef *TimeZoneDefinition) GetOriginalLocationName() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.locationName
}

// GetOriginalLocationNameType - Returns the Location Name Type
// value for the Original Time Zone. This is private
// member variable:
//
//  TimeZoneDefinition.originalTimeZone.locationNameType.
//
// Possible return values:
//  LocationNameType(0).None()
//  LocationNameType(0).NonConvertibleTimeZone()
//  LocationNameType(0).ConvertibleTimeZone()
//
func (tzdef *TimeZoneDefinition) GetOriginalLocationNameType() LocationNameType {

	return tzdef.originalTimeZone.locationNameType
}

// GetOriginalOffsetElements - Returns a series of string and
// integer values which taken collectively identify the offset
// from UTC for the Original Time Zone.
//
// ------------------------------------------------------------
//
// Return Values
// =============
//
// offsetSignChar string - Like return value offsetSignValue, this string
//                         value signals whether the offset from UTC is West
//                         or East of UTC. This string will always have one of
//                         two values: "+" or "-". The plus sign ("+") signals
//                         that the offset is East of UTC. The minus sign ("-")
//                         signals that the offset is West of UTC.
//
// offsetSignValue int  - Similar to return value 'offsetSignChar' above except
//                        that sign values are expressed as either a '-1' or positive
//                        '1' integer value. -1 == West of UTC  +1 == East of UTC.
//                        Apply this sign value to the offset hours, minutes and
//                        seconds value returned below.
//
// offsetHours     int  - Normalized Offset Hours from UTC. Always a positive number,
//                        refer to ZoneSign for correct sign value.
//
// offsetMinutes   int  - Normalized Offset Minutes offset from UTC. Always a
//                        positive number, refer to ZoneSign for the correct
//                        sign value.
//
// offsetSeconds   int  - Normalized Offset Seconds offset from UTC. Always a
//                        positive number, refer to ZoneSign for the correct
//                        sign value.
//
func (tzdef *TimeZoneDefinition) GetOriginalOffsetElements() (
	offsetSignChar string,
	offsetSignValue,
	offsetHours,
	offsetMinutes,
	offsetSeconds int) {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetOffsetElements()
}

// GetOriginalOffsetSignValue - Returns the Original Time Zone
// Offset Sign as an integer value. The returned integer will
// have one of two values: '-1' or the positive value, '1'.
//
// A negative value (-1) signals that the time zone offset
// from UTC is West of UTC.  Conversely, a positive value
// (+1) signals that the offset is East of UTC.
//
// -1 == West of UTC  +1 == East of UTC
//
func (tzdef *TimeZoneDefinition) GetOriginalOffsetSignValue() int {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneSignValue
}

// GetOriginalTagDescription - Returns the tag/description
// for the Original Time Zone.
//
// This string is typically used classification, labeling
// or description text by user.
//
func (tzdef *TimeZoneDefinition) GetOriginalTagDescription() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.tagDescription
}

// GetOriginalTimeZone - Returns the originally configured
// time zone specification. Check TimeZoneSpecification.GetOriginalLocationNameType()
// if this time zone is fully convertible across all time zones.
// If not convertible, use TimeZoneDefinition.GetConvertibleTimeZone().
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZone() TimeZoneSpecification {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone
}

// GetOriginalZoneName - Returns TimeZoneDefinition member variable
// TimeZoneDefinition.zoneName value.
//
// Zone Name is the Time Zone abbreviation. This may
// may be a series of characters, like "EST", "CST"
// and "PDT" - or - if a time zone abbreviation does
// not exist for this time zone, the time zone abbreviation
// might be listed simply as the UTC offset.
//
// Examples: '-0430', '-04', '+05'
//
func (tzdef *TimeZoneDefinition) GetOriginalZoneName() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.zoneName
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

// GetOriginalUtcOffset - Returns the offset from UTC as a string.
// Examples of the UTC offset format are: "-0600" or "+0200".
//
func (tzdef *TimeZoneDefinition) GetOriginalUtcOffset() string {

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.utcOffset
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

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.New() "

	if dateTime.IsZero() {
		return TimeZoneDefinition{}, errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is a ZERO value!\n")
	}

	tzDef2 := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}


	err := tzDefUtil.setFromDateTime(&tzDef2, dateTime, ePrefix)

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

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.NewFromTimeZoneName() "

	err = nil
	tzDefDto = TimeZoneDefinition{}

	if dateTime.IsZero() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Input parameter 'dateTime' has a Zero value!",
			err:                 nil,
		}

		return tzDefDto, err
	}

	if len(timeZoneName) == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "Input parameter 'timeZoneName' is an empty string!",
			err:                 nil,
		}

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

// SetTagDescription - Sets the Original and Convertible time zone
// tag/description. This string field is available to users for use
// as a tag, label, classification or text description.
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
