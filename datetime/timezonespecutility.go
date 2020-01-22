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
// TimeZoneSpecification (tzSpec). When completed 'tzSpec' will
// have data field values identical to those of 'tzSpec2'
//
func (tzSpecUtil *typeZoneSpecUtility) copyIn(
	tzSpec *TimeZoneSpecification,
	tzSpec2 TimeZoneSpecification)  {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("typeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	tzSpec.zoneLabel               = tzSpec2.zoneLabel
	tzSpec.referenceDateTime       = tzSpec2.referenceDateTime
	tzSpec.zoneName                = tzSpec2.zoneName
	tzSpec.zoneOffsetTotalSeconds  = tzSpec2.zoneOffsetTotalSeconds
	tzSpec.zoneSignValue           = tzSpec2.zoneSignValue
	tzSpec.offsetHours             = tzSpec2.offsetHours
	tzSpec.offsetMinutes           = tzSpec2.offsetMinutes
	tzSpec.offsetSeconds           = tzSpec2.offsetSeconds
	tzSpec.zoneOffset              = tzSpec2.zoneOffset
	tzSpec.zoneAbbrvLookupId       = tzSpec2.zoneAbbrvLookupId
	tzSpec.utcOffset               = tzSpec2.utcOffset
	tzSpec.locationPtr             = tzSpec2.locationPtr
	tzSpec.locationName            = tzSpec2.locationName
	tzSpec.locationNameType        = tzSpec2.locationNameType
	tzSpec.militaryTimeZoneName    = tzSpec2.militaryTimeZoneName
	tzSpec.militaryTimeZoneLetter  = tzSpec2.militaryTimeZoneLetter
	tzSpec.tagDescription          = tzSpec2.tagDescription
	tzSpec.timeZoneType            = tzSpec2.timeZoneType
	tzSpec.timeZoneClass           = tzSpec2.timeZoneClass
	tzSpec.timeZoneCategory        = tzSpec2.timeZoneCategory
	tzSpec.timeZoneUtcOffsetStatus = tzSpec2.timeZoneUtcOffsetStatus
}
	// CopyOut - Returns a deep copy of the current Time Zone
// Specification object as a new instance of 'TimeZoneSpecification'.
//
func (tzSpecUtil *typeZoneSpecUtility) copyOut(
	tzSpec *TimeZoneSpecification) TimeZoneSpecification {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("typeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	tzSpec2 := TimeZoneSpecification{}

	tzSpec2.zoneLabel               = tzSpec.zoneLabel
	tzSpec2.referenceDateTime       = tzSpec.referenceDateTime
	tzSpec2.zoneName                = tzSpec.zoneName
	tzSpec2.zoneOffsetTotalSeconds  = tzSpec.zoneOffsetTotalSeconds
	tzSpec2.zoneSignValue           = tzSpec.zoneSignValue
	tzSpec2.offsetHours             = tzSpec.offsetHours
	tzSpec2.offsetMinutes           = tzSpec.offsetMinutes
	tzSpec2.offsetSeconds           = tzSpec.offsetSeconds
	tzSpec2.zoneOffset              = tzSpec.zoneOffset
	tzSpec2.zoneAbbrvLookupId       = tzSpec.zoneAbbrvLookupId
	tzSpec2.utcOffset               = tzSpec.utcOffset
	tzSpec2.locationPtr             = tzSpec.locationPtr
	tzSpec2.locationName            = tzSpec.locationName
	tzSpec2.locationNameType        = tzSpec.locationNameType
	tzSpec2.militaryTimeZoneName    = tzSpec.militaryTimeZoneName
	tzSpec2.militaryTimeZoneLetter  = tzSpec.militaryTimeZoneLetter
	tzSpec2.tagDescription          = tzSpec.tagDescription
	tzSpec2.timeZoneType            = tzSpec.timeZoneType
	tzSpec2.timeZoneClass           = tzSpec.timeZoneClass
	tzSpec2.timeZoneCategory        = tzSpec.timeZoneCategory
	tzSpec2.timeZoneUtcOffsetStatus = tzSpec.timeZoneUtcOffsetStatus

	return tzSpec2
}

// empty - This method resets all the member variables
// of a TimeZoneSpecification instance to their uninitialized
// or zero values.
//
func (tzSpecUtil *typeZoneSpecUtility) empty(
	tzSpec *TimeZoneSpecification) {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("typeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	tzSpec.zoneLabel               = ""
	tzSpec.referenceDateTime       = time.Time{}
	tzSpec.zoneName                = ""
	tzSpec.zoneOffsetTotalSeconds  = 0
	tzSpec.zoneSignValue           = 0
	tzSpec.offsetHours             = 0
	tzSpec.offsetMinutes           = 0
	tzSpec.offsetSeconds           = 0
	tzSpec.zoneOffset              = ""
	tzSpec.zoneAbbrvLookupId       = ""
	tzSpec.utcOffset               = ""
	tzSpec.locationPtr             = nil
	tzSpec.locationName            = ""
	tzSpec.locationNameType        = LocNameType.None()
	tzSpec.militaryTimeZoneName    = ""
	tzSpec.militaryTimeZoneLetter  = ""
	tzSpec.tagDescription          = ""
	tzSpec.timeZoneType            = TzType.None()
	tzSpec.timeZoneClass           = TzClass.None()
	tzSpec.timeZoneCategory        = TzCat.None()
	tzSpec.timeZoneUtcOffsetStatus = TzUtcStatus.None()

}

// equal - Returns a boolean value of true if both the current instance
// of TimeZoneSpecification and the input parameter TimeZoneSpecification are
// equivalent in all respects.
//
// Exceptions: Note that the following private member data fields
// are NOT checked for equivalency.
//
// zone label is NOT checked for equivalency
// tagDescription is NOT checked for equivalency
//
func (tzSpecUtil *typeZoneSpecUtility) equal(
	tzSpec *TimeZoneSpecification,
	tzSpec2 TimeZoneSpecification) bool {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("typeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	if !tzSpec.referenceDateTime.Equal(tzSpec2.referenceDateTime) {
		return false
	}

	if tzSpec.zoneName != tzSpec2.zoneName {
		return false
	}

	if tzSpec.zoneOffsetTotalSeconds != tzSpec2.zoneOffsetTotalSeconds{
		return false
	}

	if tzSpec.zoneSignValue != tzSpec2.zoneSignValue {
		return false
	}

	if tzSpec.offsetHours != tzSpec2.offsetHours {
		return false
	}

	if tzSpec.offsetMinutes != tzSpec2.offsetMinutes {
		return false
	}

	if tzSpec.offsetSeconds != tzSpec2.offsetSeconds {
		return false
	}

	if tzSpec.zoneOffset != tzSpec2.zoneOffset {
		return false
	}

	if tzSpec.zoneAbbrvLookupId != tzSpec2.zoneAbbrvLookupId {
		return false
	}

	if tzSpec.utcOffset != tzSpec2.utcOffset {
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

	if tzSpec.militaryTimeZoneLetter != tzSpec2.militaryTimeZoneLetter {
		return false
	}

	if tzSpec.militaryTimeZoneName != tzSpec2.militaryTimeZoneName {
		return false
	}

	if tzSpec.locationNameType != tzSpec.locationNameType {
		return false
	}

	if tzSpec.timeZoneType != tzSpec2.timeZoneType {
		return false
	}

	if tzSpec.timeZoneClass != tzSpec2.timeZoneClass {
		return false
	}

	if tzSpec.timeZoneCategory != tzSpec2.timeZoneCategory {
		return false
	}

	if tzSpec.timeZoneUtcOffsetStatus != tzSpec2.timeZoneUtcOffsetStatus {
		return false
	}

	return true

}