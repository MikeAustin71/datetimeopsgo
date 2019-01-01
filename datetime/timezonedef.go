package datetime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

/*
 TimeZoneDefDto
 ==============
 This source file is located in source code repository:
 		https://github.com/MikeAustin71/datetimeopsgo.git'

 This source code file is located at:
		MikeAustin71\datetimeopsgo\datetime\timezonedef.go


 Overview and General Usage
 ==========================
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
	ZoneName          string
	ZoneOffsetSeconds int            // Signed number of seconds offset from UTC. + == East of UTC; - == West of UTC
	ZoneSign          int            // -1 == West of UTC  +1 == East of UTC
	OffsetHours       int            // Hours offset from UTC. Always a positive number, refer to ZoneSign
	OffsetMinutes     int            // Minutes offset from UTC. Always a positive number, refer to ZoneSign
	OffsetSeconds     int            // Seconds offset from UTC. Always a positive number, refer to ZoneSign
	ZoneOffset        string         // A text string representing the time zone. Example "-0500 CDT"
	Location          *time.Location // Pointer to a Time Zone Location
	LocationName      string         // Time Zone Location Name Examples: "Local", "America/Chicago", "America/New_York"
	Description       string         // Unused - Available for classification, labeling or description by user.
}

// CopyIn - Copies an incoming TimeZoneDefDto into the
// data fields of the current TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) CopyIn(tzdef2 TimeZoneDefDto) {

	tzdef.Empty()

	tzdef.ZoneName = tzdef2.ZoneName
	tzdef.ZoneOffsetSeconds = tzdef2.ZoneOffsetSeconds
	tzdef.ZoneSign = tzdef2.ZoneSign
	tzdef.OffsetHours = tzdef2.OffsetHours
	tzdef.OffsetMinutes = tzdef2.OffsetMinutes
	tzdef.OffsetSeconds = tzdef2.OffsetSeconds
	tzdef.ZoneOffset = tzdef2.ZoneOffset
	tzdef.Location = tzdef2.Location
	tzdef.LocationName = tzdef2.LocationName
	tzdef.Description = tzdef2.Description

}

// CopyOut - creates and returns a deep copy of the current
// TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) CopyOut() TimeZoneDefDto {

	tzdef2 := TimeZoneDefDto{}

	tzdef2.ZoneName = tzdef.ZoneName
	tzdef2.ZoneOffsetSeconds = tzdef.ZoneOffsetSeconds
	tzdef2.ZoneSign = tzdef.ZoneSign
	tzdef2.OffsetHours = tzdef.OffsetHours
	tzdef2.OffsetMinutes = tzdef.OffsetMinutes
	tzdef2.OffsetSeconds = tzdef.OffsetSeconds
	tzdef2.ZoneOffset = tzdef.ZoneOffset
	tzdef2.Location = tzdef.Location
	tzdef2.LocationName = tzdef.LocationName
	tzdef2.Description = tzdef.Description

	return tzdef2

}

// Empty - Resets all field values for the current TimeZoneDefDto
// instance to their uninitialized or 'zero' states.
func (tzdef *TimeZoneDefDto) Empty() {
	tzdef.ZoneName = ""
	tzdef.ZoneOffsetSeconds = 0
	tzdef.ZoneSign = 0
	tzdef.OffsetHours = 0
	tzdef.OffsetMinutes = 0
	tzdef.OffsetSeconds = 0
	tzdef.ZoneOffset = ""
	tzdef.Location = nil
	tzdef.LocationName = ""
	tzdef.Description = ""

}

// Equal - Determines if two TimeZoneDefDto are equivalent in
// value. Returns 'true' of two TimeZoneDefDto's are equal in
// all respects.
func (tzdef *TimeZoneDefDto) Equal(tzdef2 TimeZoneDefDto) bool {

	if tzdef.ZoneName == tzdef2.ZoneName &&
		tzdef.ZoneOffsetSeconds == tzdef2.ZoneOffsetSeconds &&
		tzdef.ZoneSign == tzdef2.ZoneSign &&
		tzdef.OffsetHours == tzdef2.OffsetHours &&
		tzdef.OffsetMinutes == tzdef2.OffsetMinutes &&
		tzdef.OffsetSeconds == tzdef2.OffsetSeconds &&
		tzdef.ZoneOffset == tzdef2.ZoneOffset &&
		tzdef.Location.String() == tzdef2.Location.String() &&
		tzdef.LocationName == tzdef2.LocationName &&
		tzdef.Description == tzdef2.Description {
		return true
	}

	return false
}

// EqualOffsetSeconds - Compares Zone Offset Seconds for two TimeZoneDefDto's and
// returns 'true' if they are equal.
//
// ZoneOffsetSeconds is a signed number of seconds offset from UTC:
// 		+ == East of UTC
// 		- == West of UTC
func (tzdef *TimeZoneDefDto) EqualOffsetSeconds(tzdef2 TimeZoneDefDto) bool {

	if tzdef.ZoneOffsetSeconds == tzdef2.ZoneOffsetSeconds {
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

	if tzdef.ZoneOffset == tzdef2.ZoneOffset {
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

	if tzdef.LocationName == tzdef2.LocationName {
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

// IsEmpty - Determines whether the current TimeZoneDefDto
// instance is Empty.
//
// If the TimeZoneDefDto instance is NOT populated, the method
// returns 'true'. Otherwise, it returns 'false'.
//
func (tzdef *TimeZoneDefDto) IsEmpty() bool {

	if tzdef.ZoneName != "" ||
		tzdef.ZoneOffsetSeconds != 0 ||
		tzdef.ZoneSign != 0 ||
		tzdef.OffsetHours != 0 ||
		tzdef.OffsetMinutes != 0 ||
		tzdef.OffsetSeconds != 0 ||
		tzdef.ZoneOffset != "" ||
		tzdef.LocationName != "" {
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

	if strings.TrimLeft(strings.TrimRight(tzdef.LocationName, " "), " ") == "" {
		return false
	}

	if tzdef.Location.String() != tzdef.LocationName {
		return false
	}

	loc, err := time.LoadLocation(tzdef.LocationName)

	if err != nil {
		return false
	}

	if loc != tzdef.Location {
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

	tzdef2.Description = tzdef.Description

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
//	dateTime 	time.Time	- A date time value which will be used to construct the
//													elements of a Time Zone Definition Dto instance.
//
// Returns
// =======
//
// This method will return two Types:
//			(1) A Time Zone Definition Dto
//			(2) An 'error' type
//
// (1) If successful, this method will return a valid, populated TimeZoneDefDto instance.
//		 A TimeZoneDefDto is defined as follows:
//			type TimeZoneDefDto struct {
//				ZoneName						string
//				ZoneOffsetSeconds		int			// Signed number of seconds offset from UTC. + == East of UTC; - == West of UTC
//				ZoneSign						int 		// -1 == West of UTC  +1 == East of UTC
//				OffsetHours					int			// Hours offset from UTC. Always a positive number, refer to ZoneSign
//				OffsetMinutes				int			// Minutes offset from UTC. Always a positive number, refer to ZoneSign
//				OffsetSeconds				int			// Seconds offset from UTC. Always a positive number, refer to ZoneSign
//				ZoneOffset					string	// A text string representing the time zone. Example "-0500 CDT"
//				Location						*time.Location	// Pointer to a Time Zone Location
//				LocationName				string					// Time Zone Location Name Examples: "Local", "America/Chicago", "America/New_York"
//				Description					string	// Unused - Available for classification, labeling or description by user.
//			}
//
//
// (2) 	If successful, this method will set the returned error instance to 'nil.
//			If errors are encountered a valid error message will be returned in the
//			error instance.
//
func (tzdef TimeZoneDefDto) New(dateTime time.Time) (TimeZoneDefDto, error) {

	ePrefix := "TimeZoneDefDto.New() "

	if dateTime.IsZero() {
		return TimeZoneDefDto{}, errors.New(ePrefix + "Error: Input parameter 'dateTime' is a ZERO value!")
	}

	tzDef2 := TimeZoneDefDto{}

	err := tzDef2.SetFromDateTime(dateTime)

	if err != nil {
		return TimeZoneDefDto{}, fmt.Errorf(ePrefix+
			"Error returned by tzDef2.SetFromDateTime(dateTime). "+
			"dateTime='%v'  Error='%v'",
			dateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return tzDef2, nil
}

// SetFromDateTime - Re-initializes the values of the current
// TimeZoneDefDto instance based on input parameter, 'dateTime'.
func (tzdef *TimeZoneDefDto) SetFromDateTime(dateTime time.Time) error {
	ePrefix := "TimeZoneDefDto.SetFromDateTime() "

	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'dateTime' is a ZERO value!")
	}

	tzdef.Empty()

	tzdef.ZoneName, tzdef.ZoneOffsetSeconds = dateTime.Zone()

	tzdef.allocateZoneOffsetSeconds(tzdef.ZoneOffsetSeconds)

	tzdef.Location = dateTime.Location()

	tzdef.LocationName = dateTime.Location().String()

	tzdef.setZoneString()

	tzdef.Description = ""

	return nil
}

// allocateZoneOffsetSeconds - allocates a signed value of total offset seconds from
// UTC to the associated fields in the current TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) allocateZoneOffsetSeconds(signedZoneOffsetSeconds int) {

	if signedZoneOffsetSeconds < 0 {
		tzdef.ZoneSign = -1
	} else {
		tzdef.ZoneSign = 1
	}

	tzdef.ZoneOffsetSeconds = signedZoneOffsetSeconds

	signedZoneOffsetSeconds *= tzdef.ZoneSign

	tzdef.OffsetHours = 0
	tzdef.OffsetMinutes = 0
	tzdef.OffsetSeconds = 0

	if signedZoneOffsetSeconds == 0 {
		return
	}

	tzdef.OffsetHours = signedZoneOffsetSeconds / 3600 // compute hours
	signedZoneOffsetSeconds -= tzdef.OffsetHours * 3600

	tzdef.OffsetMinutes = signedZoneOffsetSeconds / 60 // compute minutes
	signedZoneOffsetSeconds -= tzdef.OffsetMinutes * 60

	tzdef.OffsetSeconds = signedZoneOffsetSeconds

	return
}

// setZoneString - assembles and assigns the composite zone
// offset and zone name abbreviation in the TimeZoneDefDto.ZoneOffset
// field. Example: "-0500 CST"
func (tzdef *TimeZoneDefDto) setZoneString() {

	tzdef.ZoneOffset = ""

	if tzdef.ZoneSign < 0 {
		tzdef.ZoneOffset += "-"
	}

	tzdef.ZoneOffset += fmt.Sprintf("%02d%02d", tzdef.OffsetHours, tzdef.OffsetMinutes)

	if tzdef.OffsetSeconds > 0 {
		tzdef.ZoneOffset += fmt.Sprintf("%02d", tzdef.OffsetSeconds)
	}

	tzdef.ZoneOffset += " " + tzdef.ZoneName

	return
}
