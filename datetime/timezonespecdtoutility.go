package datetime

import (
	"sync"
	"time"
)

type typeZoneSpecUtility struct {
	lock sync.Mutex
}

// copyIn - Copies the values of input parameter 'tzSpec2'
// to all of the data fields in the tzSpec instance of
// TimeZoneSpecDto (tzSpec). When completed 'tzSpec' will
// have data field values identical to those of 'tzSpec2'
//
func (tzSpecUtil *typeZoneSpecUtility) copyIn(
	tzSpec *TimeZoneSpecDto,
	tzSpec2 TimeZoneSpecDto)  {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("typeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	tzSpec.zoneLabel              = tzSpec2.zoneLabel
	tzSpec.referenceDateTime      = tzSpec2.referenceDateTime
	tzSpec.zoneName               = tzSpec2.zoneName
	tzSpec.zoneOffsetTotalSeconds = tzSpec2.zoneOffsetTotalSeconds
	tzSpec.zoneSignValue          = tzSpec2.zoneSignValue
	tzSpec.offsetHours            = tzSpec2.offsetHours
	tzSpec.offsetMinutes          = tzSpec2.offsetMinutes
	tzSpec.offsetSeconds          = tzSpec2.offsetSeconds
	tzSpec.zoneOffset             = tzSpec2.zoneOffset
	tzSpec.zoneAbbrvLookupId      = tzSpec2.zoneAbbrvLookupId
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
func (tzSpecUtil *typeZoneSpecUtility) copyOut(
	tzSpec *TimeZoneSpecDto) TimeZoneSpecDto {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("typeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

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
	tzSpec2.zoneAbbrvLookupId      = tzSpec.zoneAbbrvLookupId
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

// empty - This method resets all the member variables
// of a TimeZoneSpecDto instance to their uninitialized
// or zero values.
//
func (tzSpecUtil *typeZoneSpecUtility) empty(
	tzSpec *TimeZoneSpecDto) {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("typeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	tzSpec.zoneLabel              = ""
	tzSpec.referenceDateTime      = time.Time{}
	tzSpec.zoneName               = ""
	tzSpec.zoneOffsetTotalSeconds = 0
	tzSpec.zoneSignValue          = 0
	tzSpec.offsetHours            = 0
	tzSpec.offsetMinutes          = 0
	tzSpec.offsetSeconds          = 0
	tzSpec.zoneOffset             = ""
	tzSpec.zoneAbbrvLookupId      = ""
	tzSpec.utcOffset              = ""
	tzSpec.locationPtr            = nil
	tzSpec.locationName           = ""
	tzSpec.locationNameType       = LocNameType.None()
	tzSpec.militaryTimeZoneName   = ""
	tzSpec.militaryTimeZoneLetter = ""
	tzSpec.tagDescription         = ""
	tzSpec.timeZoneType           = TzType.None()

}
