package datetime

import (
	"sync"
	"time"
)

// timeZoneSpecDto - Internal data structure used to
// store Time Zone data elements.
type TimeZoneSpecDto struct {
	zoneLabel              string           // Identifying Time Zone Label
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
	zoneOffset             string           // A text string representing the time zone. Example "-0600 CST" or "+0200 EET"
	zoneAbbreviation       string           // A test string representing the Abbreviation Id used in map lookups. Example: "CST-0600"
	utcOffset              string           // A text string representing the offset for UTC. Example "-0600" or "+0200"
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

	tzSpec.zoneLabel              = tzSpec2.zoneLabel
	tzSpec.referenceDateTime      = tzSpec2.referenceDateTime
	tzSpec.zoneName               = tzSpec2.zoneName
	tzSpec.zoneOffsetTotalSeconds = tzSpec2.zoneOffsetTotalSeconds
	tzSpec.zoneSignValue          = tzSpec2.zoneSignValue
	tzSpec.offsetHours            = tzSpec2.offsetHours
	tzSpec.offsetMinutes          = tzSpec2.offsetMinutes
	tzSpec.offsetSeconds          = tzSpec2.offsetSeconds
	tzSpec.zoneOffset             = tzSpec2.zoneOffset
	tzSpec.zoneAbbreviation       = tzSpec2.zoneAbbreviation
	tzSpec.utcOffset              = tzSpec2.utcOffset
	tzSpec.locationPtr            = tzSpec2.locationPtr
	tzSpec.locationName           = tzSpec2.locationName
	tzSpec.locationNameType       = tzSpec2.locationNameType
	tzSpec.militaryTimeZoneName   = tzSpec2.militaryTimeZoneName
	tzSpec.militaryTimeZoneLetter = tzSpec2.militaryTimeZoneLetter
	tzSpec.tagDescription         = tzSpec2.tagDescription
	tzSpec.timeZoneType           = tzSpec2.timeZoneType

}


// CopyOut - Returns a deep copy of the current Time Zone 
// Specification object as a new instance of 'TimeZoneSpecDto'.
//
func (tzSpec *TimeZoneSpecDto) CopyOut() TimeZoneSpecDto {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpec2 := TimeZoneSpecDto{}

	tzSpec2.zoneLabel              = tzSpec.zoneLabel
	tzSpec2.referenceDateTime      = tzSpec.referenceDateTime      
	tzSpec2.zoneName               = tzSpec.zoneName               
	tzSpec2.zoneOffsetTotalSeconds = tzSpec.zoneOffsetTotalSeconds 
	tzSpec2.zoneSignValue          = tzSpec.zoneSignValue          
	tzSpec2.offsetHours            = tzSpec.offsetHours            
	tzSpec2.offsetMinutes          = tzSpec.offsetMinutes          
	tzSpec2.offsetSeconds          = tzSpec.offsetSeconds          
	tzSpec2.zoneOffset             = tzSpec.zoneOffset             
	tzSpec2.zoneAbbreviation       = tzSpec.zoneAbbreviation       
	tzSpec2.utcOffset              = tzSpec.utcOffset              
	tzSpec2.locationPtr            = tzSpec.locationPtr            
	tzSpec2.locationName           = tzSpec.locationName           
	tzSpec2.locationNameType       = tzSpec.locationNameType       
	tzSpec2.militaryTimeZoneName   = tzSpec.militaryTimeZoneName   
	tzSpec2.militaryTimeZoneLetter = tzSpec.militaryTimeZoneLetter 
	tzSpec2.tagDescription         = tzSpec.tagDescription         
	tzSpec2.timeZoneType           = tzSpec.timeZoneType           

	return tzSpec2
}

// Empty - Sets all the values of the data fields in the
// current TimeZoneSpecDto to their empty or zero values.
//
func (tzSpec *TimeZoneSpecDto) Empty() {

	tzSpec.lock.Lock()

	defer tzSpec.lock.Unlock()

	tzSpec.zoneLabel              = ""
	tzSpec.referenceDateTime      = time.Time{}
	tzSpec.zoneName               = ""
	tzSpec.zoneOffsetTotalSeconds = 0
	tzSpec.zoneSignValue          = 0
	tzSpec.offsetHours            = 0
	tzSpec.offsetMinutes          = 0
	tzSpec.offsetSeconds          = 0
	tzSpec.zoneOffset             = ""
	tzSpec.zoneAbbreviation       = ""
	tzSpec.utcOffset              = ""
	tzSpec.locationPtr            = nil
	tzSpec.locationName           = ""
	tzSpec.locationNameType       = LocNameType.None()
	tzSpec.militaryTimeZoneName   = ""
	tzSpec.militaryTimeZoneLetter = ""
	tzSpec.tagDescription         = ""
	tzSpec.timeZoneType           = TzType.None()

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
		tzSpec.zoneAbbreviation != tzSpec2.zoneAbbreviation ||
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
	tzSpec.zoneAbbreviation != "" ||
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

// SetTimeZone - Sets the data values of the current Time Zone
// Specification Structure (TimeZoneSpecDto).
//
func (tzSpec *TimeZoneSpecDto) SetTimeZone(
	referenceDateTime      time.Time,
	militaryTimeZoneName   string,
	militaryTimeZoneLetter string,
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
	
	tzMech := timeZoneMechanics{}
  var err error
	
	tzSpec.zoneName,
		tzSpec.zoneOffset,
		tzSpec.utcOffset,
		tzSpec.zoneAbbreviation,
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