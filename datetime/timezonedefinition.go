package datetime

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// TimeZoneDefinition - Time Zone definition.
//
// Overview and General Usage
//
// This type is designed to process IANA Time Zones, Military Time
// Zones and the 'Local' Time Zone implemented by Golang.
//
// For background information on Time Zones, see comments in the
// 'timezonedata.go' source file and the 'TimeZones' Type at:
//
//   MikeAustin71\datetimeopsgo\datetime\timezonedata.go
//
// A Time Zone Definition includes two Time Zone Specifications
// (Type: TimeZoneSpecification). Both describe the same time zone.
// The first 'TimeZoneSpecification' describes the original time
// zone used to initialize the current TimeZoneDefinition instance.
// The second is provided to ensure the configuration of a time zone
// which can be used to convert date times to another time zone
// anywhere in the world.
//
// Why Two Time Zone Specifications?
//
// Providing two time zone specifications deals with errors which
// may be introduced when using the Golang time package to parse
// date time strings in order to generate date time objects (time.Time).
// In the following example, assume you, and your computer, live in
// NewStartEndTimes York, NewYork USA. (I'll explain later). Now, create a new
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
//   live in NewStartEndTimes York, NewStartEndTimes York USA. If in fact you lived in Los
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
//   MikeAustin71\datetimeopsgo\datetime\timezonedefinition.go
//
type TimeZoneDefinition struct {
	originalTimeZone    TimeZoneSpecification // The Time Zone Specification originally submitted.
	convertibleTimeZone TimeZoneSpecification // A version of the original time zone with a new fully
	//                                         convertible time zone substituted.
	lock *sync.Mutex // Used for implementing thread safe operations.
}

// CopyIn - Copies an incoming TimeZoneDefinition into the
// data fields of the current TimeZoneDefinition instance.
func (tzdef *TimeZoneDefinition) CopyIn(tzdef2 TimeZoneDefinition) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	tzDefUtil.copyIn(tzdef, &tzdef2)

	return
}

// copyOut - creates and returns a deep copy of the current
// TimeZoneDefinition instance.
func (tzdef *TimeZoneDefinition) CopyOut() TimeZoneDefinition {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.copyOut(tzdef)
}

// Empty - Resets all field values for the current TimeZoneDefinition
// instance to their uninitialized or 'zero' states.
func (tzdef *TimeZoneDefinition) Empty() {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	tzDefUtil.empty(tzdef)
}

// Equal - Determines if two TimeZoneDefinition are equivalent in
// value. Returns 'true' of two TimeZoneDefinition's are equal in
// all respects.
//
// In order to qualify as equal both the Original Time Zones and
// the Convertible Time Zones must be equal.
//
func (tzdef *TimeZoneDefinition) Equal(tzdef2 TimeZoneDefinition) bool {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

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
//
// In order to qualify as equal, both the Original Time Zones and
// the Convertible Time Zones must be equal.
//
func (tzdef *TimeZoneDefinition) EqualOffsetSeconds(tzdef2 TimeZoneDefinition) bool {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

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
// In order to qualify as 'equal', both the Original Time Zone Offsets and
// the Convertible Time Zone Offsets must be equal.
//
func (tzdef *TimeZoneDefinition) EqualZoneOffsets(tzdef2 TimeZoneDefinition) bool {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

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
// In order to qualify as equal, both the Original Time Zone Locations
// and the Convertible Time Zone Locations must be equal.
//
func (tzdef *TimeZoneDefinition) EqualLocations(tzdef2 TimeZoneDefinition) bool {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

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
// To qualify as 'equal', Location Names, Zone Names and Zone Offsets
// for both Original Time Zones and Convertible Time Zones must be
// equal.
//
func (tzdef *TimeZoneDefinition) EqualZoneLocation(tzdef2 TimeZoneDefinition) bool {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzDefUtil := timeZoneDefUtility{}

	if !tzDefUtil.equalZoneLocation(tzdef, &tzdef2) {
		return false
	}

	return true
}

// GetBestConvertibleTimeZone - If the Original Time Zone qualifies as
// a fully convertible time zone, this method will return the Time Zone
// Specification for the Original Time Zone.
//
// If the Original Time Zone does NOT qualify as a fully convertible Time
// Zone, this method will return the Time Zone Specification for the
// Convertible Time Zone.
//
// A convertible time zone is a date time (time.Time) Location Name
// configured as a fully formed time zone name which can be used to
// accurately convert date times to other time zones across the globe.
//
func (tzdef *TimeZoneDefinition) GetBestConvertibleTimeZone() TimeZoneSpecification {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	if tzdef.originalTimeZone.locationNameType == LocNameType.ConvertibleTimeZone() {
		return tzdef.originalTimeZone.CopyOut()
	}

	return tzdef.convertibleTimeZone.CopyOut()
}


// GetConvertibleDateTime - Returns the Convertible Time Zone reference
// date time value.
//
func (tzdef *TimeZoneDefinition) GetConvertibleDateTime() time.Time {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

return tzdef.convertibleTimeZone.GetReferenceDateTime()
}


// GetConvertibleLocationPtr - Returns a pointer to a time.Location for
// the convertible time zone (TimeZoneDefinition.convertibleTimeZone.locationPtr).
//
// The Time Zone 'Location' represents the Time Zone Name.
//
func (tzdef *TimeZoneDefinition) GetConvertibleLocationPtr() *time.Location {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetLocationPointer()
}

// GetConvertibleLocationName - Returns convertible Time Zone
// Name. This is also referred to as the "Location Name". This
// private member variable is TimeZoneDefinition.convertibleTimeZone.locationName.
//
// Time Zone Location Name Examples: "Local", "America/Chicago",
// "America/New_York".
//
func (tzdef *TimeZoneDefinition) GetConvertibleLocationName() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetLocationName()
}

// GetConvertibleLocationNameType - Returns the Location Name Type
// value for the Convertible Time Zone. This is private
// member variable:
//
//  TimeZoneDefinition.convertibleTimeZone.locationNameType.
//
// Possible return values:
//
//  LocationNameType(0).None()
//                               - The Time Zone is uninitialized. This
//                                 is an error condition.
//
//  LocationNameType(0).NonConvertibleTimeZone()
//                               - The Time Zone Location Name cannot
//                                 be converted to other time zones.
//
//  LocationNameType(0).ConvertibleTimeZone()
//                               - The Time Zone Name is a complete
//                                 and valid time zone name which is
//                                 convertible across all other
//                                 time zones.
//
// For easy access to these enumeration values, use the global variable,
// 'LocNameType'. Example: LocNameType.ConvertibleTimeZone()
//
func (tzdef *TimeZoneDefinition) GetConvertibleLocationNameType() LocationNameType {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetLocationNameType()
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

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetOffsetElements()
}

// GetConvertibleOffsetSignValue - Returns the Convertible Time
// Zone Offset Sign as an integer value. The returned integer
// will have one of two values: '-1' or the positive value
// '1'.
//
// A negative value (-1) signals that the time zone offset from
// UTC is West of UTC.  Conversely, a positive value (+1)
// signals that the offset is East of UTC.
//
// -1 == West of UTC  +1 == East of UTC
//
func (tzdef *TimeZoneDefinition) GetConvertibleOffsetSignValue() int {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetZoneSignValue()
}

// GetConvertibleTagDescription - Returns the tag/description
// for the Convertible Time Zone.
//
// This string is typically used classification, labeling
// or description text by user.
//
func (tzdef *TimeZoneDefinition) GetConvertibleTagDescription() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetTagDescription()
}

// GetConvertibleTimeZone - Returns a deep copy of the convertible
// form of the current time zone. This time zone should be convertible
// across all time zones. To verify convertibility, call
// 'TimeZoneSpecification.GetOriginalLocationNameType()' and inspect the result.
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZone() TimeZoneSpecification {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.CopyOut()
}

// GetTimeZoneAbbreviation - Returns the Time Zone abbreviation for the
// Convertible Time Zone. The term 'Time Zone Abbreviation' is a synonym
// for 'Zone Name'.
//
// Examples of Time Zone Abbreviations are:
//  'EST', 'CST', 'PST', 'EDT', 'CDT', 'PDT'
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZoneAbbreviation() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetTimeZoneAbbreviation()
}

// GetConvertibleTimeZoneCategory - Returns the Time Zone Category for
// the Convertible Time Zone.
//
// Time Zone Category is styled as enumeration. Possible values are:
//  TimeZoneCategory(0).None()       -  Signals that Time Zone Category is uninitialized.
//                                      This represents an error condition.
//
//  TimeZoneCategory(0).TextName()   -  Signals that the Time Zone is identified
//                                      by a standard IANA Text Name. Examples:
//                                        "America/Chicago"
//                                        "Asia/Amman"
//                                        "Atlantic/Bermuda"
//                                        "Australia/Sydney"
//                                        "Europe/Rome"
//
//  TimeZoneCategory(0)..UtcOffset()  -  Signals that the Time Zone is identified
//                                       by a valid UTC Offset and has no associated
//                                       text name. Examples:
//                                         "+07"
//                                         "+10"
//
// For easy access to these enumeration values, use the global variable
// 'TzCat'. Example: TzCat.None()
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZoneCategory() TimeZoneCategory {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetTimeZoneCategory()

}

// GetConvertibleTimeZoneClass - Returns the Time Zone Class associated
// with the Convertible Time Zone.
//
// Time Zone Class is styled as enumeration. Possible values are:
//
// TimeZoneClass(0).None()              - An Error Condition
// TimeZoneClass(0).AlternateTimeZone() - Generated Time Zone from
//                                        Time Zone Abbreviation
// TimeZoneClass(0).OriginalTimeZone()  - Original Valid Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzClass'. Example: TzClass.AlternateTimeZone()
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZoneClass() TimeZoneClass {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetTimeZoneClass()
}

// GetConvertibleTimeZoneName - Returns the time zone name,
// also known as the Time Zone 'Location' Name, for the
// Convertible Time Zone.
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZoneName() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.locationName
}

// GetConvertibleTimeZoneStatusFlags - Returns all internal status
// flags for the Convertible Time Zone.
//
// Return Values
// --------------------------------------------------------------------
//
// LocationNameType:
//
//  LocationNameType(0).None()
//                               - The Time Zone is uninitialized. This
//                                 is an error condition.
//
//  LocationNameType(0).NonConvertibleTimeZone()
//                               - The Time Zone Location Name cannot
//                                 be converted to other time zones.
//
//  LocationNameType(0).ConvertibleTimeZone()
//                               - The Time Zone Name is a complete
//                                 and valid time zone name which is
//                                 convertible across all other
//                                 time zones.
//
// For easy access to these enumeration values, use the global variable,
// 'LocNameType'. Example: LocNameType.ConvertibleTimeZone()
//
//
// TimeZoneCategory:
//
//  TimeZoneCategory(0).None()       -  Signals that Time Zone Category is uninitialized.
//                                      This represents an error condition.
//
//  TimeZoneCategory(0).TextName()   -  Signals that the Time Zone is identified
//                                      by a standard IANA Text Name. Examples:
//                                        "America/Chicago"
//                                        "Asia/Amman"
//                                        "Atlantic/Bermuda"
//                                        "Australia/Sydney"
//                                        "Europe/Rome"
//
//  TimeZoneCategory(0).UtcOffset()   -  Signals that the Time Zone is identified
//                                       by a valid UTC Offset and has no associated
//                                       text name. Examples:
//                                         "+07"
//                                         "+10"
//
//  For easy access to these enumeration values, use the global variable
//  'TzCat'. Example: TzCat.None()
//
//
// TimeZoneClass:
//
// TimeZoneClass(0).None()              - Signals that Time Zone Class is uninitialized
//                                        This is an Error Condition.
//
// TimeZoneClass(0).AlternateTimeZone() - Generated Time Zone from Time Zone
//                                        Abbreviation
//
// TimeZoneClass(0).OriginalTimeZone()  - Original Valid Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzClass'. Example: TzClass.AlternateTimeZone()
//
//
// TimeZoneType:
//
//  TimeZoneType(0).None()      - Time Zone type is uninitialized
//                                and has no significant value.
//
//  TimeZoneType(0).Iana()      - Identifies an IANA Time Zone
//
//  TimeZoneType(0).Local()     - Identifies this as a 'Local' Time Zone
//
//  TimeZoneType(0).Military()  - Identifies a Military Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzType'. Example: TzType.Military()
//
//
// TimeZoneUtcOffsetStatus:
//
// TimeZoneUtcOffsetStatus(0).None()
//               - Signals that Time Zone UTC Offset
//                 Status is uninitialized and contains
//                 no significant or valid value. This
//                 is an error condition.
//
// TimeZoneUtcOffsetStatus(0).Static()
//               - Signals that the UTC Offset associated
//                 with a given Time Zone is constant
//                 throughout the year and never changes.
//                 Typically, this means that Daylight
//                 Savings Time is NOT observed in the
//                 specified Time Zone.
//
// TimeZoneUtcOffsetStatus(0).Variable()
//               - Signals that the UTC Offset associated
//                 with a given Time Zone is not constant,
//                 and varies at least once during the year.
//                 This usually means that Daylight Savings
//                 Time is observed within the designated
//                 Time Zone.
//
// For easy access to these enumeration values, use the global variable
// 'TzUtcStatus'. Example: TzUtcStatus.Variable()
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZoneStatusFlags() (
	LocationNameType,
	TimeZoneCategory,
	TimeZoneClass,
	TimeZoneType,
	TimeZoneUtcOffsetStatus) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetTimeZoneSpecFlags()
}

// GetConvertibleTimeZoneType - Returns the Time Zone Type associated
// with the Convertible Time Zone.
//
// Time Zone Type is styled as an enumeration. Possible values are:
//
//  TimeZoneType(0).None()      - Time Zone type is uninitialized
//                                and has no significant value.
//
//  TimeZoneType(0).Iana()      - Identifies an IANA Time Zone
//  TimeZoneType(0).Local()     - Identifies this as a 'Local' Time Zone
//  TimeZoneType(0).Military()  - Identifies a Military Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzType'. Example: TzType.Military()
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZoneType() TimeZoneType {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetTimeZoneType()
}

// GetConvertibleTimeZoneUtcOffsetStatus - Returns the Time Zone UTC Offset
// Status for the Convertible Time Zone.
//
// Time Zone Type is styled as an enumeration. Possible values are:
//
// TzUtcStatus.None()                   Signals that Time Zone UTC Offset
//                                      Status is uninitialized and contains
//                                      no significant or valid value. This
//                                      is an error condition.
//
// TzUtcStatus.Static()                 Signals that the UTC Offset associated
//                                      with a given Time Zone is constant
//                                      throughout the year and never changes.
//                                      Typically, this means that Daylight
//                                      Savings Time is NOT observed in the
//                                      specified Time Zone.
//
// TzUtcStatus.Variable()               Signals that the UTC Offset associated
//                                      with a given Time Zone is not constant,
//                                      and varies at least once during the year.
//                                      This usually means that Daylight Savings
//                                      Time is observed within the designated
//                                      Time Zone.
//
// For easy access to these enumeration values, use the global variable
// 'TzUtcStatus'. Example: TzUtcStatus.Variable()
//
func (tzdef *TimeZoneDefinition) GetConvertibleTimeZoneUtcOffsetStatus() TimeZoneUtcOffsetStatus {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetTimeZoneUtcOffsetStatus()
}

// GetConvertibleUtcOffset - Returns the offset from UTC as a string
// for the Convertible Time Zone.
//
// Examples of the UTC offset format are: "-0600" or "+0200".
//
func (tzdef *TimeZoneDefinition) GetConvertibleUtcOffset() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetUtcOffset()
}

// GetConvertibleZoneOffset - Returns the Convertible Time Zone
// Offset value.
//
// Zone Offset is a text string representing the time zone
// which is formatted with UTC offset followed by time zone
// abbreviation.
//
//   Examples: "-0600 CST" or "+0200 EET"
//
func (tzdef *TimeZoneDefinition) GetConvertibleZoneOffset() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetZoneOffset()
}

// GetOriginalZoneOffsetTotalSeconds - Returns Zone Offset Total Seconds
// for the Convertible Time Zone. Zone offset total seconds is a positive
// or negative integer value presenting the total number of seconds of
// offset from UTC. A positive value signals that the offset is East of
// UTC. A negative value indicates that the offset is West of UTC.
//
func (tzdef *TimeZoneDefinition) GetConvertibleZoneOffsetTotalSeconds() int {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.convertibleTimeZone.GetZoneOffsetTotalSeconds()
}

// GetConvertibleZoneName - Returns the Convertible Time
// Zone Name.
//
// Zone Name is the Time Zone abbreviation. This may
// may be a series of characters, like "EST", "CST"
// and "PDT" - or - if a time zone abbreviation does
// not exist for this time zone, the time zone abbreviation
// might be listed simply as the UTC offset.
//
// UTC Offset Examples: '-0430', '-04', '+05'
//
func (tzdef *TimeZoneDefinition) GetConvertibleZoneName() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetZoneName()
}

// GetOriginalDateTime - Returns the Original Time Zone reference
// date time value.
//
func (tzdef *TimeZoneDefinition) GetOriginalDateTime() time.Time {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetReferenceDateTime()
}

// GetOriginalLocationPtr - Returns a pointer to a time.Location for
// the original time zone (TimeZoneDefinition.originalTimeZone.locationPtr).
//
func (tzdef *TimeZoneDefinition) GetOriginalLocationPtr() *time.Location {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetLocationPointer()
}

// GetOriginalLocationName - Returns original Time Zone Name.
// This is also referred to as the "Location Name". This
// private member variable is:
//   TimeZoneDefinition.convertibleTimeZone.locationName.
//
// Time Zone Location Name Examples: "Local", "America/Chicago",
// "America/New_York".
//
func (tzdef *TimeZoneDefinition) GetOriginalLocationName() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetLocationName()
}

// GetOriginalLocationNameType - Returns the Location Name Type
// value for the Original Time Zone. This is private a member
// variable:
//
//   TimeZoneDefinition.originalTimeZone.locationNameType.
//
// Possible return values:
//
//  LocationNameType(0).None()
//                               - The Time Zone is uninitialized. This
//                                 is an error condition.
//
//  LocationNameType(0).NonConvertibleTimeZone()
//                               - The Time Zone Location Name cannot
//                                 be converted to other time zones.
//
//  LocationNameType(0).ConvertibleTimeZone()
//                               - The Time Zone Name is a complete
//                                 and valid time zone name which is
//                                 convertible across all other
//                                 time zones.
//
// For easy access to these enumeration values, use the global variable,
// 'LocNameType'. Example: LocNameType.ConvertibleTimeZone()
//
func (tzdef *TimeZoneDefinition) GetOriginalLocationNameType() LocationNameType {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetLocationNameType()
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

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

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

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetZoneSignValue()
}

// GetOriginalTagDescription - Returns the tag/description
// for the Original Time Zone.
//
// This string is typically used classification, labeling
// or description text by user.
//
func (tzdef *TimeZoneDefinition) GetOriginalTagDescription() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetTagDescription()
}

// GetOriginalTimeZone - Returns a deep copy of the originally configured
// time zone specification. Check 'TimeZoneSpecification.GetOriginalLocationNameType()'
// to determine if this time zone is fully convertible across all time zones.
// If it not convertible, use TimeZoneDefinition.GetConvertibleTimeZone().
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZone() TimeZoneSpecification {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.CopyOut()
}

// GetOriginalTimeZoneAbbreviation - Returns the Time Zone abbreviation for
// the Original Time Zone.
//
// The term 'Time Zone Abbreviation' is a synonym for 'Zone Name'.
//
// Examples of Time Zone Abbreviations are:
//  'EST', 'CST', 'PST', 'EDT', 'CDT', 'PDT'
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZoneAbbreviation() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetTimeZoneAbbreviation()
}

// GetOriginalTimeZoneCategory - Returns the Time Zone Category for
// the Original Time Zone.
//
// Time Zone Category is styled as enumeration. Possible values are:
//  TimeZoneCategory(0).None()       -  Signals that Time Zone Category is uninitialized.
//                                      This represents an error condition.
//
//  TimeZoneCategory(0).TextName()   -  Signals that the Time Zone is identified
//                                      by a standard IANA Text Name. Examples:
//                                        "America/Chicago"
//                                        "Asia/Amman"
//                                        "Atlantic/Bermuda"
//                                        "Australia/Sydney"
//                                        "Europe/Rome"
//
//  TimeZoneCategory(0)..UtcOffset()  -  Signals that the Time Zone is identified
//                                       by a valid UTC Offset and has no associated
//                                       text name. Examples:
//                                         "+07"
//                                         "+10"
//
// For easy access to these enumeration values, use the global variable
// 'TzCat'. Example: TzCat.None()
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZoneCategory() TimeZoneCategory {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetTimeZoneCategory()

}

// GetOriginalTimeZoneClass - Returns the Time Zone Class associated
// with the Original Time Zone.
//
// Time Zone Class is styled as enumeration. Possible values are:
//
// TimeZoneClass(0).None()              - An Error Condition
// TimeZoneClass(0).AlternateTimeZone() - Generated Time Zone from
//                                        Time Zone Abbreviation
// TimeZoneClass(0).OriginalTimeZone()  - Original Valid Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzClass'. Example: TzClass.AlternateTimeZone()
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZoneClass() TimeZoneClass {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetTimeZoneClass()
}

// GetOriginalTimeZoneName - Returns the time zone name, also
// known as the Time Zone 'Location' Name, for the Original
// Time Zone.
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZoneName() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.locationName
}

// GetOriginalTimeZoneStatusFlags - Returns all internal status
// flags for the Original Time Zone.
// Return Values
// --------------------------------------------------------------------
//
// LocationNameType:
//
//  LocationNameType(0).None()
//                               - The Time Zone is uninitialized. This
//                                 is an error condition.
//
//  LocationNameType(0).NonConvertibleTimeZone()
//                               - The Time Zone Location Name cannot
//                                 be converted to other time zones.
//
//  LocationNameType(0).ConvertibleTimeZone()
//                               - The Time Zone Name is a complete
//                                 and valid time zone name which is
//                                 convertible across all other
//                                 time zones.
//
// For easy access to these enumeration values, use the global variable,
// 'LocNameType'. Example: LocNameType.ConvertibleTimeZone()
//
//
// TimeZoneCategory:
//
//  TimeZoneCategory(0).None()       -  Signals that Time Zone Category is uninitialized.
//                                      This represents an error condition.
//
//  TimeZoneCategory(0).TextName()   -  Signals that the Time Zone is identified
//                                      by a standard IANA Text Name. Examples:
//                                        "America/Chicago"
//                                        "Asia/Amman"
//                                        "Atlantic/Bermuda"
//                                        "Australia/Sydney"
//                                        "Europe/Rome"
//
//  TimeZoneCategory(0).UtcOffset()   -  Signals that the Time Zone is identified
//                                       by a valid UTC Offset and has no associated
//                                       text name. Examples:
//                                         "+07"
//                                         "+10"
//
//  For easy access to these enumeration values, use the global variable
//  'TzCat'. Example: TzCat.None()
//
//
// TimeZoneClass:
//
// TimeZoneClass(0).None()              - Signals that Time Zone Class is uninitialized
//                                        This is an Error Condition.
//
// TimeZoneClass(0).AlternateTimeZone() - Generated Time Zone from Time Zone
//                                        Abbreviation
//
// TimeZoneClass(0).OriginalTimeZone()  - Original Valid Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzClass'. Example: TzClass.AlternateTimeZone()
//
//
// TimeZoneType:
//
//  TimeZoneType(0).None()      - Time Zone type is uninitialized
//                                and has no significant value.
//
//  TimeZoneType(0).Iana()      - Identifies an IANA Time Zone
//
//  TimeZoneType(0).Local()     - Identifies this as a 'Local' Time Zone
//
//  TimeZoneType(0).Military()  - Identifies a Military Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzType'. Example: TzType.Military()
//
//
// TimeZoneUtcOffsetStatus:
//
// TimeZoneUtcOffsetStatus(0).None()
//               - Signals that Time Zone UTC Offset
//                 Status is uninitialized and contains
//                 no significant or valid value. This
//                 is an error condition.
//
// TimeZoneUtcOffsetStatus(0).Static()
//               - Signals that the UTC Offset associated
//                 with a given Time Zone is constant
//                 throughout the year and never changes.
//                 Typically, this means that Daylight
//                 Savings Time is NOT observed in the
//                 specified Time Zone.
//
// TimeZoneUtcOffsetStatus(0).Variable()
//               - Signals that the UTC Offset associated
//                 with a given Time Zone is not constant,
//                 and varies at least once during the year.
//                 This usually means that Daylight Savings
//                 Time is observed within the designated
//                 Time Zone.
//
// For easy access to these enumeration values, use the global variable
// 'TzUtcStatus'. Example: TzUtcStatus.Variable()
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZoneStatusFlags() (
	LocationNameType,
	TimeZoneCategory,
	TimeZoneClass,
	TimeZoneType,
	TimeZoneUtcOffsetStatus) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetTimeZoneSpecFlags()
}

// GetOriginalTimeZoneType - Returns the Time Zone Type associated
// with the Original Time Zone.
//
// Time Zone Type is styled as an enumeration. Possible values are:
//
//  TimeZoneType(0).None()      - Time Zone type is uninitialized
//                                and has no significant value.
//
//  TimeZoneType(0).Iana()      - Identifies an IANA Time Zone
//  TimeZoneType(0).Local()     - Identifies this as a 'Local' Time Zone
//  TimeZoneType(0).Military()  - Identifies a Military Time Zone
//
// For easy access to these enumeration values, use the global variable
// 'TzType'. Example: TzType.Military()
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZoneType() TimeZoneType {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetTimeZoneType()
}

// GetOriginalTimeZoneUtcOffsetStatus - Returns the Time Zone UTC Offset
// Status for the Original Time Zone.
//
// Time Zone Type is styled as an enumeration. Possible values are:
//
// TzUtcStatus.None()                   Signals that Time Zone UTC Offset
//                                      Status is uninitialized and contains
//                                      no significant or valid value. This
//                                      is an error condition.
//
// TzUtcStatus.Static()                 Signals that the UTC Offset associated
//                                      with a given Time Zone is constant
//                                      throughout the year and never changes.
//                                      Typically, this means that Daylight
//                                      Savings Time is NOT observed in the
//                                      specified Time Zone.
//
// TzUtcStatus.Variable()               Signals that the UTC Offset associated
//                                      with a given Time Zone is not constant,
//                                      and varies at least once during the year.
//                                      This usually means that Daylight Savings
//                                      Time is observed within the designated
//                                      Time Zone.
//
// For easy access to these enumeration values, use the global variable
// 'TzUtcStatus'. Example: TzUtcStatus.Variable()
//
func (tzdef *TimeZoneDefinition) GetOriginalTimeZoneUtcOffsetStatus() TimeZoneUtcOffsetStatus {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetTimeZoneUtcOffsetStatus()
}

// GetOriginalUtcOffset - Returns the offset from UTC as a string
// for the Original Time Zone.
//
// Examples of the UTC offset format are:
//   "+0600"
//   "-0500"
//   "-0430"
//
func (tzdef *TimeZoneDefinition) GetOriginalUtcOffset() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetUtcOffset()
}

// GetOriginalZoneName - Returns the Original Time Zone
// Name.
//
// Zone Name is the Time Zone abbreviation. This may
// may be a series of characters, like "EST", "CST"
// and "PDT" - or - if a time zone abbreviation does
// not exist for this time zone, the time zone abbreviation
// might be listed simply as the UTC offset.
//
// UTC Offset Examples: '-0430', '-04', '+05'
//
func (tzdef *TimeZoneDefinition) GetOriginalZoneName() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetZoneName()
}

// GetOriginalZoneOffset - Returns the Original Time Zone
// Offset value.
//
// Zone Offset is a text string representing the time zone
// which is formatted with UTC offset followed by time zone
// abbreviation.
//
//   Examples: "-0600 CST" or "+0200 EET"
//
func (tzdef *TimeZoneDefinition) GetOriginalZoneOffset() string {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetZoneOffset()
}

// GetOriginalZoneOffsetTotalSeconds - Returns Zone Offset Total Seconds
// for the Original Time Zone. Zone offset total seconds is a positive
// or negative integer value presenting the total number of seconds of
// offset from UTC. A positive value signals that the offset is East of
// UTC. A negative value indicates that the offset is West of UTC.
//
func (tzdef *TimeZoneDefinition) GetOriginalZoneOffsetTotalSeconds() int {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	return tzdef.originalTimeZone.GetZoneOffsetTotalSeconds()
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

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

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

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

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

// IsEmpty - Determines whether the current TimeZoneDefinition
// instance is Empty.
//
// If the TimeZoneDefinition instance is NOT populated, the method
// returns 'true'. Otherwise, it returns 'false'.
//
func (tzdef *TimeZoneDefinition) IsEmpty() bool {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

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

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.IsValid() "

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.isValidTimeZoneDef(tzdef, ePrefix)
}

// New - Returns a new TimeZoneDefinition instance with
// member variables initialized to zero values.
//
func (tzdef TimeZoneDefinition) New() (TimeZoneDefinition, error) {

	if tzdef.lock == nil {
	tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.New() "

	tzDef2 := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	err := tzDefUtil.setZeroTimeZoneDef(
		&tzDef2,
		ePrefix)


	return tzDef2, err
}

// NewDateTime - Creates and returns a new TimeZoneDefinition instance based on
// a 'dateTime (time.Time) input parameter.
//
// Input Parameter
// ===============
//
//  dateTime   time.Time  - A date time value which will be used to construct the
//                          an instance of Time Zone Definition (TimeZoneDefinition).
//
// Return Values
// =============
//
// This method will return two values:
//      (1) A Time Zone Definition (TimeZoneDefinition)
//      (2) An 'error' type
//
//  (1) If successful, this method will return a valid, populated 'TimeZoneDefinition'
//      instance.
//
//  (2) If successful, this method will set the returned 'error' instance to 'nil'.
//      If errors are encountered a valid error message will be returned in the
//      error instance.
//
func (tzdef TimeZoneDefinition) NewDateTime(
	dateTime time.Time) (TimeZoneDefinition, error) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.NewDateTime() "

	tzDef2 := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	err := tzDefUtil.setFromDateTime(&tzDef2, dateTime, ePrefix)

	if err != nil {
		return TimeZoneDefinition{}, err
	}

	return tzDef2, nil
}

// NewFromTimeComponents - Creates and returns a new 'TimeZoneDefinition' instance based on
// date time components and Time Zone Name input parameters.
//
// The date time components, years, months, days, hours, seconds and subMicrosecondNanoseconds are used
// in the Time Zone Name to construct an instance of 'TimeZoneDefinition'.

// Input Parameter
// ===============
//
//  years       int        - Years value used to construct a date time object (time.Time).
//
//  months      int        - Months value used to construct a date time object (time.Time).
//
//  days        int        - Days value used to construct a date time object (time.Time).
//
//  hours       int        - Hours value used to construct a date time object (time.Time).
//
//  minutes     int        - Minutes value used to construct a date time object (time.Time).
//
//  seconds     int        - Seconds value used to construct a date time object (time.Time).
//
//  subMicrosecondNanoseconds int        - Nanoseconds value used to construct a date time object (time.Time).
//
//  timeZoneName string    - This string contains the name of a valid time zone.
//                           The 'timeZoneName' string must be set to one of three values:
//
//                           1. A valid IANA Time Zone name.
//
//                           2. The time zone "Local", which Golang accepts as the time
//                              zone currently configured on the host computer.
//
//                           3. A valid Military Time Zone which can be submitted either as
//                              a single alphabetic character Military Time Zone abbreviation
//                              or as a full Military Time Zone name.
//
// Return Values
// =============
//
// This method will return two values:
//      (1) A Time Zone Definition (TimeZoneDefinition)
//      (2) An 'error' type
//
//  (1) If successful, this method will return a valid, populated TimeZoneDefinition
//      instance.
//
//  (2) If successful, this method will set the returned 'error' instance to 'nil'.
//      If errors are encountered a valid error message will be returned in the
//      error instance.
//
func (tzdef TimeZoneDefinition) NewFromTimeComponents(
	years,
	months,
	days,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	timeZoneName string) (TimeZoneDefinition, error) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.NewFromTimeComponents() "

	timeDto, err := TimeDto{}.NewTimeElements(
		years,
		months,
		days,
		hours,
		minutes,
		seconds,
		nanoseconds)

	if err != nil {
		return TimeZoneDefinition{},
				fmt.Errorf(ePrefix +
					"\nError retunred by TimeDto{}.NewTimeElements(...)\n" +
					"Error='%v'\n", err.Error())
	}

	if len(timeZoneName) == 0 {
		return TimeZoneDefinition{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeZoneName",
				inputParameterValue: "",
				errMsg:              "Input parameter 'timeZoneName' is an EMPTY string!",
				err:                 nil,
			}
	}

	tzDefUtil := timeZoneDefUtility{}

	tzDef2 := TimeZoneDefinition{}

	err = tzDefUtil.setFromTimeDto(
		&tzDef2,
		timeDto,
		timeZoneName,
		ePrefix)

	if err != nil {
		return TimeZoneDefinition{}, err
	}

	return tzDef2, nil
}


// NewFromTimeDto - Creates and returns a new TimeZoneDefinition instance based on
// a TimeDto and Time Zone Name input parameters.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeDto      TimeDto
//     - The 'TimeDto' type contains time components such as years,
//       months, days, seconds and subMicrosecondNanoseconds which is used to
//       construct a date time value (time.Time).
//
//
//   timeZoneName string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time duration
//       calculations.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location, or time zone name,
//       must be designated as one of three types
//       of values:
//
//       (1) The string 'Local' - signals the designation of the local time zone
//           configured for the host computer executing this code.
//
//       (2) IANA Time Zone Location -
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system. Note: IANA Time Zone Data base is
//           equivalent to 'tz database'.
//
//              Examples:
//                "America/New_York"
//                "America/Chicago"
//                "America/Denver"
//                "America/Los_Angeles"
//                "Pacific/Honolulu"
//
//       (3) A valid Military Time Zone
//           Military time zones are commonly used in
//           aviation as well as at sea. They are also
//           known as nautical or maritime time zones.
//           Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// This method will return two values:
//      (1) A Time Zone Definition (TimeZoneDefinition)
//      (2) An 'error' type
//
//  (1) If successful, this method will return a valid, populated TimeZoneDefinition
//      instance.
//
//  (2) If successful, this method will set the returned 'error' instance to 'nil'.
//      If errors are encountered a valid error message will be returned in the
//      error instance.
//
func (tzdef TimeZoneDefinition) NewFromTimeDto(
	timeDto TimeDto,
	timeZoneName string) (TimeZoneDefinition, error) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.NewFromTimeDto() "

	if timeDto.IsEmpty() {
		return TimeZoneDefinition{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeDto",
				inputParameterValue: "",
				errMsg:              "Input parameter 'timeDto' is EMPTY!",
				err:                 nil,
		}
	}

	if len(timeZoneName) == 0 {
		return TimeZoneDefinition{},
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "timeZoneName",
				inputParameterValue: "",
				errMsg:              "Input parameter 'timeZoneName' is an EMPTY string!",
				err:                 nil,
		}
	}

	tzDefUtil := timeZoneDefUtility{}

	tzDef2 := TimeZoneDefinition{}

	err := tzDefUtil.setFromTimeDto(
		&tzDef2,
		timeDto,
		timeZoneName,
		ePrefix)

	if err != nil {
		return TimeZoneDefinition{}, err
	}

	return tzDef2, nil
}

// NewFromTimeZoneName - Creates and returns a new instance 'TimeZoneDefinition'.
// The new instance is based on input parameter 'timeZoneName'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   dateTime time.Time
//     - A reference date time value which will be used in conjunction with
//       parameter 'timeZoneName' to create the returned Time Zone Definition
//       instance.
//
//
//   timeZoneName  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time duration
//       calculations.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location, or time zone name,
//       must be designated as one of three types
//       of values:
//
//       (1) The string 'Local' - signals the designation of the local time zone
//           configured for the host computer executing this code.
//
//       (2) IANA Time Zone Location -
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system. Note: IANA Time Zone Data base is
//           equivalent to 'tz database'.
//
//              Examples:
//                "America/New_York"
//                "America/Chicago"
//                "America/Denver"
//                "America/Los_Angeles"
//                "Pacific/Honolulu"
//
//       (3) A valid Military Time Zone
//           Military time zones are commonly used in
//           aviation as well as at sea. They are also
//           known as nautical or maritime time zones.
//           Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//   timeConversionType TimeZoneConversionType -
//          This parameter determines the algorithm that will
//          be used to convert parameter 'dateTime' to the time
//          zone specified by parameter 'tzSpec'.
//
//          TimeZoneConversionType is an enumeration type which
//          must be set to one of two values:
//             TimeZoneConversionType(0).Absolute()
//             TimeZoneConversionType(0).Relative()
//          Note: You can also use the global variable
//          'TzConvertType' for easier access:
//             TzConvertType.Absolute()
//             TzConvertType.Relative()
//
//          Absolute Time Conversion - Identifies the 'Absolute' time
//          to time zone conversion algorithm. This algorithm provides
//          that a time value in time zone 'X' will be converted to the
//          same time value in time zone 'Y'.
//
//          For example, assume the time 10:00AM is associated with time
//          zone USA Central Standard time and that this time is to be
//          converted to USA Eastern Standard time. Applying the 'Absolute'
//          algorithm would convert ths time to 10:00AM Eastern Standard
//          time.  In this case the hours, minutes and seconds have not been
//          altered. 10:00AM in USA Central Standard Time has simply been
//          reclassified as 10:00AM in USA Eastern Standard Time.
//
//          Relative Time Conversion - Identifies the 'Relative' time to time
//          zone conversion algorithm. This algorithm provides that times in
//          time zone 'X' will be converted to their equivalent time in time
//          zone 'Y'.
//
//          For example, assume the time 10:00AM is associated with time zone
//          USA Central Standard time and that this time is to be converted to
//          USA Eastern Standard time. Applying the 'Relative' algorithm would
//          convert ths time to 11:00AM Eastern Standard time. In this case the
//          hours, minutes and seconds have been changed to reflect an equivalent
//          time in the USA Eastern Standard Time Zone.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
// This method will return two values:
//      (1) A Time Zone Definition (TimeZoneDefinition)
//      (2) An 'error' type
//
//  (1) If successful, this method will return a valid, populated TimeZoneDefinition
//      instance.
//
//  (2) If successful, this method will set the returned 'error' instance to 'nil'.
//      If errors are encountered a valid error message will be returned in the
//      error instance.
//
func (tzdef TimeZoneDefinition) NewFromTimeZoneName(
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType) (
	tzDefDto TimeZoneDefinition,
	err error) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.NewFromTimeZoneName() "

	err = nil
	tzDefDto = TimeZoneDefinition{}

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
		timeConversionType,
		timeZoneName,
		ePrefix)

	return tzDefDto, err
}

// NewTzSpecFromTzName - Creates a 'TimeZoneDefinition' object
// and then returns the first convertible time zone specification.
//
// A 'TimeZoneDefinition' consists of two time zone specifications,
// an Original Time Zone and a Convertible Time Zone. The
// 'TimeZoneDefinition' instance is created based on input parameter
// 'timeZoneName'.
//
// If the Original Time Zone is NOT convertible, the Convertible
// Time Zone Specification is returned.  Otherwise, the Original
// Time Zone Specification is returned.
//
// A convertible time zone means that the Time Zone Location Name
// is a fully formed time zone name which can be used to accurately
// convert date times to other time zones across the globe.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   dateTime time.Time
//     - A reference date time value which will be used in conjunction with
//       parameter 'timeZoneName' to create the returned Time Zone Definition
//       instance.
//
//
//   timeZoneName  string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time duration
//       calculations.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location, or time zone name,
//       must be designated as one of three types
//       of values:
//
//       (1) The string 'Local' - signals the designation of the local time zone
//           configured for the host computer executing this code.
//
//       (2) IANA Time Zone Location -
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system. Note: IANA Time Zone Data base is
//           equivalent to 'tz database'.
//
//              Examples:
//                "America/New_York"
//                "America/Chicago"
//                "America/Denver"
//                "America/Los_Angeles"
//                "Pacific/Honolulu"
//
//       (3) A valid Military Time Zone
//           Military time zones are commonly used in
//           aviation as well as at sea. They are also
//           known as nautical or maritime time zones.
//           Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//       Note:
//           The source file 'timezonedata.go' contains over 600 constant
//           time zone declarations covering all IANA and Military Time
//           Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//           time zone constants begin with the prefix 'TZones'.
//
//
//   timeConversionType TimeZoneConversionType -
//          This parameter determines the algorithm that will
//          be used to convert parameter 'dateTime' to the time
//          zone specified by parameter 'tzSpec'.
//
//          TimeZoneConversionType is an enumeration type which
//          must be set to one of two values:
//             TimeZoneConversionType(0).Absolute()
//             TimeZoneConversionType(0).Relative()
//          Note: You can also use the global variable
//          'TzConvertType' for easier access:
//             TzConvertType.Absolute()
//             TzConvertType.Relative()
//
//          Absolute Time Conversion - Identifies the 'Absolute' time
//          to time zone conversion algorithm. This algorithm provides
//          that a time value in time zone 'X' will be converted to the
//          same time value in time zone 'Y'.
//
//          For example, assume the time 10:00AM is associated with time
//          zone USA Central Standard time and that this time is to be
//          converted to USA Eastern Standard time. Applying the 'Absolute'
//          algorithm would convert ths time to 10:00AM Eastern Standard
//          time.  In this case the hours, minutes and seconds have not been
//          altered. 10:00AM in USA Central Standard Time has simply been
//          reclassified as 10:00AM in USA Eastern Standard Time.
//
//          Relative Time Conversion - Identifies the 'Relative' time to time
//          zone conversion algorithm. This algorithm provides that times in
//          time zone 'X' will be converted to their equivalent time in time
//          zone 'Y'.
//
//          For example, assume the time 10:00AM is associated with time zone
//          USA Central Standard time and that this time is to be converted to
//          USA Eastern Standard time. Applying the 'Relative' algorithm would
//          convert ths time to 11:00AM Eastern Standard time. In this case the
//          hours, minutes and seconds have been changed to reflect an equivalent
//          time in the USA Eastern Standard Time Zone.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
// This method will return two values:
//      (1) A Time Zone Specification (TimeZoneSpecification)
//      (2) An 'error' type
//
//  (1) If successful, this method will return a valid, populated
//      'TimeZoneSpecification' object representing a 'convertible'
//      time zone.
//
//  (2) If successful, this method will set the returned 'error' instance to 'nil'.
//      If errors are encountered a valid error message will be returned in the
//      error instance.
//
//      If neither the Original Time Zone nor the Convertible Time Zone qualify as
//      'Convertible', an error is returned.
//
//      A 'Convertible' Time Zone Name is defined as a complete and valid time zone
//      name which is convertible across all other global time zones.
//
func (tzdef TimeZoneDefinition) NewTzSpecFromTzName(
	dateTime time.Time,
	timeZoneName string,
	timeConversionType TimeZoneConversionType) (
	TimeZoneSpec TimeZoneSpecification,
	err error) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.NewTzSpecFromTzName() "

	err = nil
	TimeZoneSpec = TimeZoneSpecification{}

	if len(timeZoneName) == 0 {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "Input parameter 'timeZoneName' is an empty string!",
			err:                 nil,
		}

		return TimeZoneSpec, err
	}

	tzDefUtil := timeZoneDefUtility{}
	tzDef := TimeZoneDefinition{}

	err = tzDefUtil.setFromTimeZoneName(
		&tzDef,
		dateTime,
		timeConversionType,
		timeZoneName,
		ePrefix)

	if err != nil {
		return TimeZoneSpec, err
	}

	if tzDef.originalTimeZone.locationNameType ==
			LocNameType.ConvertibleTimeZone() {

		TimeZoneSpec = tzDef.originalTimeZone.CopyOut()
	} else {
		TimeZoneSpec = tzDef.convertibleTimeZone.CopyOut()
	}

	return TimeZoneSpec, err
}



// NewForTimeZoneSpec - Creates and returns a new instance of
// 'TimeZoneDefinition'. The new instance is based on a
// 'TimeZoneSpecification' input parameter, 'tzSpec'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTime        time.Time    -
//          This date time will be converted and
//          used in creating the returned
//          'TimeZoneDefinition' object. The type
//          of time conversion performed on this
//          date time value is determined by input
//          parameter, 'timeConversionType'.
//
//  timeConversionType TimeZoneConversionType -
//          This parameter determines the algorithm that will
//          be used to convert parameter 'dateTime' to the time
//          zone specified by parameter 'tzSpec'.
//
//          TimeZoneConversionType is an enumeration type which
//          must be set to one of two values:
//             TimeZoneConversionType(0).Absolute()
//             TimeZoneConversionType(0).Relative()
//          Note: You can also use the global variable
//          'TzConvertType' for easier access:
//             TzConvertType.Absolute()
//             TzConvertType.Relative()
//
//          Absolute Time Conversion - Identifies the 'Absolute' time
//          to time zone conversion algorithm. This algorithm provides
//          that a time value in time zone 'X' will be converted to the
//          same time value in time zone 'Y'.
//
//          For example, assume the time 10:00AM is associated with time
//          zone USA Central Standard time and that this time is to be
//          converted to USA Eastern Standard time. Applying the 'Absolute'
//          algorithm would convert ths time to 10:00AM Eastern Standard
//          time.  In this case the hours, minutes and seconds have not been
//          altered. 10:00AM in USA Central Standard Time has simply been
//          reclassified as 10:00AM in USA Eastern Standard Time.
//
//          Relative Time Conversion - Identifies the 'Relative' time to time
//          zone conversion algorithm. This algorithm provides that times in
//          time zone 'X' will be converted to their equivalent time in time
//          zone 'Y'.
//
//          For example, assume the time 10:00AM is associated with time zone
//          USA Central Standard time and that this time is to be converted to
//          USA Eastern Standard time. Applying the 'Relative' algorithm would
//          convert ths time to 11:00AM Eastern Standard time. In this case the
//          hours, minutes and seconds have been changed to reflect an equivalent
//          time in the USA Eastern Standard Time Zone.
//
//  tzSpec TimeZoneSpecification -
//          A valid 'TimeZoneSpecification' object.
//          The 'TimeZoneSpecification' object must
//          be set to one of these three types of time
//          zones:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                configured for the host computer executing this code.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                   Examples:
//                     "America/New_York"
//                     "America/Chicago"
//                     "America/Denver"
//                     "America/Los_Angeles"
//                     "Pacific/Honolulu"
//
//            (3) A valid Military Time Zone
//                Military time zones are commonly used in
//                aviation as well as at sea. They are also
//                known as nautical or maritime time zones.
//                Reference:
//                  https://en.wikipedia.org/wiki/List_of_military_time_zones
//                  http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                  https://www.timeanddate.com/time/zones/military
//                  https://www.timeanddate.com/worldclock/timezone/alpha
//                  https://www.timeanddate.com/time/map/
//
//             Note:
//                 The source file 'timezonedata.go' contains over 600 constant
//                 time zone declarations covering all IANA and Military Time
//                 Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TZones'.
//
// Return Values
// =============
//
// This method will return two values:
//      (1) A Time Zone Definition (TimeZoneDefinition)
//      (2) An 'error' type
//
//  (1) If successful, this method will return a valid, populated TimeZoneDefinition
//      instance.
//
//  (2) If successful, this method will set the returned 'error' instance to 'nil'.
//      If errors are encountered a valid error message will be returned in the
//      error instance.
//
func (tzdef TimeZoneDefinition) NewForTimeZoneSpec(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	tzSpec TimeZoneSpecification) (
	tzDefDto TimeZoneDefinition,
	err error) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.NewForTimeZoneSpec() "

	tzDefDto = TimeZoneDefinition{}

	err = nil

	err2 := tzSpec.IsValid(ePrefix)

	if err2 != nil {
		err =
		 &InputParameterError{
			 ePrefix:             ePrefix,
			 inputParameterName:  "tzSpec",
			 inputParameterValue: "",
			 errMsg:              fmt.Sprintf(
			 	"Validation Error='%v'", err2.Error()),
			 err:                 nil,
		 }

		 return tzDefDto, err
	}

	if timeConversionType < TzConvertType.Absolute() ||
		timeConversionType > TzConvertType.Relative() {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeConversionType",
			inputParameterValue: "",
			errMsg:              "Input 'timeConversionType' must be 'Absolute' or 'Relative'.",
			err:                 nil,
		}

		return tzDefDto, err
	}

	tzDefUtil := timeZoneDefUtility{}

	err = tzDefUtil.setFromTimeZoneSpecification(
		&tzDefDto,
		dateTime,
		timeConversionType,
		tzSpec,
		ePrefix)

	return tzDefDto, err
}

// SetConvertibleTagDescription - Sets the Convertible Time Zone
// tag/description. This string field is available to users for use
// as a tag, label, classification or text description.
//
func (tzdef *TimeZoneDefinition) SetConvertibleTagDescription(tagDesc string) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzdef.convertibleTimeZone.SetTagDescription(tagDesc)
}

// SetOriginalTagDescription - Sets the Original Time Zone
// tag/description. This string field is available to users for use
// as a tag, label, classification or text description.
//
func (tzdef *TimeZoneDefinition) SetOriginalTagDescription(tagDesc string) {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	tzdef.originalTimeZone.SetTagDescription(tagDesc)
}

// SetFromDateTimeComponents - Re-initializes the values of the current
// TimeZoneDefinition instance based on input parameter, 'dateTime'.
//
func (tzdef *TimeZoneDefinition) SetFromDateTime(
	dateTime time.Time) error {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.SetFromDateTimeComponents() "

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.setFromDateTime(tzdef, dateTime, ePrefix)

}

// SetFromTimeZoneDef - Sets the data fields of the current
// TimeZoneDefinition instance based on another Time Zone
// Definition object passed as an input parameter.
//
// Input Parameters
// ================
//
// dateTime        time.Time
//       -  This date time will be converted and
//          used in creating the returned
//          'TimeZoneDefinition' object. The type
//          of time conversion performed on this
//          date time value is determined by input
//          parameter, 'timeConversionType'.
//
// timeConversionType TimeZoneConversionType
//       -  This parameter determines the algorithm that will
//          be used to convert parameter 'dateTime' to the time
//          zone specified by parameter 'TimeZoneDefinition'.
//
//          TimeZoneConversionType is an enumeration type which
//          must be set to one of two values:
//             TimeZoneConversionType(0).Absolute()
//             TimeZoneConversionType(0).Relative()
//          Note: You can also use the global variable
//          'TzConvertType' for easier access:
//             TzConvertType.Absolute()
//             TzConvertType.Relative()
//
//          Absolute Time Conversion - Identifies the 'Absolute' time
//          to time zone conversion algorithm. This algorithm provides
//          that a time value in time zone 'X' will be converted to the
//          same time value in time zone 'Y'.
//
//          For example, assume the time 10:00AM is associated with time
//          zone USA Central Standard time and that this time is to be
//          converted to USA Eastern Standard time. Applying the 'Absolute'
//          algorithm would convert ths time to 10:00AM Eastern Standard
//          time.  In this case the hours, minutes and seconds have not been
//          altered. 10:00AM in USA Central Standard Time has simply been
//          reclassified as 10:00AM in USA Eastern Standard Time.
//
//          Relative Time Conversion - Identifies the 'Relative' time to time
//          zone conversion algorithm. This algorithm provides that times in
//          time zone 'X' will be converted to their equivalent time in time
//          zone 'Y'.
//
//          For example, assume the time 10:00AM is associated with time zone
//          USA Central Standard time and that this time is to be converted to
//          USA Eastern Standard time. Applying the 'Relative' algorithm would
//          convert ths time to 11:00AM Eastern Standard time. In this case the
//          hours, minutes and seconds have been changed to reflect an equivalent
//          time in the USA Eastern Standard Time Zone.
//
// timeZoneDef TimeZoneDefinition
//       -  A valid 'TimeZoneDefinition' object.
//          The 'TimeZoneDefinition' object must
//          be set to one of these three types of time
//          zones:
//
//            (1) The string 'Local' - signals the designation of the local time zone
//                configured for the host computer executing this code.
//
//            (2) IANA Time Zone Location -
//                See https://golang.org/pkg/time/#LoadLocation
//                and https://www.iana.org/time-zones to ensure that
//                the IANA Time Zone Database is properly configured
//                on your system. Note: IANA Time Zone Data base is
//                equivalent to 'tz database'.
//
//                   Examples:
//                     "America/New_York"
//                     "America/Chicago"
//                     "America/Denver"
//                     "America/Los_Angeles"
//                     "Pacific/Honolulu"
//
//            (3) A valid Military Time Zone
//                Military time zones are commonly used in
//                aviation as well as at sea. They are also
//                known as nautical or maritime time zones.
//                Reference:
//                  https://en.wikipedia.org/wiki/List_of_military_time_zones
//                  http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                  https://www.timeanddate.com/time/zones/military
//                  https://www.timeanddate.com/worldclock/timezone/alpha
//                  https://www.timeanddate.com/time/map/
//
//             Note:
//                 The source file 'timezonedata.go' contains over 600 constant
//                 time zone declarations covering all IANA and Military Time
//                 Zones. Example: 'TZones.US.Central()' = "America/Chicago". All
//                 time zone constants begin with the prefix 'TZones'.
//
func (tzdef *TimeZoneDefinition) SetFromTimeZoneDef(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	timeZoneDef TimeZoneDefinition,
) error {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.SetFromTimeZoneName() "

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.setFromTimeZoneDefinition(
		tzdef,
		dateTime,
		timeConversionType,
		timeZoneDef,
		ePrefix)
}

// SetFromTimeZoneName - Sets the data fields of the current
// TimeZoneDefinition instance based on the time zone text name
// passed as an input parameter.
//
func (tzdef *TimeZoneDefinition) SetFromTimeZoneName(
	dateTime time.Time,
	timeConversionType TimeZoneConversionType,
	timeZoneName string,
	) error {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.SetFromTimeZoneName() "

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.setFromTimeZoneName(
		tzdef,
		dateTime,
		timeConversionType,
		timeZoneName,
		ePrefix)
}


// SetFromTimeZoneSpec - Sets the current TimeZoneDefinition
// object using a date time, time zone specification and
// a time zone conversion type flag.
//
func (tzdef *TimeZoneDefinition) SetFromTimeZoneSpec(
	dateTime time.Time,
	tzSpec TimeZoneSpecification,
	timeConversionType TimeZoneConversionType) error {

	if tzdef.lock == nil {
		tzdef.lock = new(sync.Mutex)
	}

	tzdef.lock.Lock()

	defer tzdef.lock.Unlock()

	ePrefix := "TimeZoneDefinition.SetFromTimeZoneName() "

	tzDefUtil := timeZoneDefUtility{}

	return tzDefUtil.setFromTimeZoneSpecification(
		tzdef,
		dateTime,
		timeConversionType,
		tzSpec,
		ePrefix)

}