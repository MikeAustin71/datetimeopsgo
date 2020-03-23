package datetime

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type timeZoneSpecUtility struct {
	lock sync.Mutex
}

// copyIn - Copies the values of input parameter 'tzSpec2'
// to all of the data fields in the tzSpec instance of
// TimeZoneSpecification (tzSpec). When completed 'tzSpec' will
// have data field values identical to those of 'tzSpec2'
//
func (tzSpecUtil *timeZoneSpecUtility) copyIn(
	tzSpec *TimeZoneSpecification,
	tzSpec2 *TimeZoneSpecification)  {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("timeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	if tzSpec2 == nil {
		panic("timeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec2 is a 'nil' pointer!\n")
	}

	if tzSpec2.lock == nil {
		tzSpec2.lock = new(sync.Mutex)
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
func (tzSpecUtil *timeZoneSpecUtility) copyOut(
	tzSpec *TimeZoneSpecification) TimeZoneSpecification {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("timeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
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

	tzSpec2.lock = new(sync.Mutex)

	return tzSpec2
}

// empty - This method resets all the member variables
// of a TimeZoneSpecification instance to their uninitialized
// or zero values.
//
func (tzSpecUtil *timeZoneSpecUtility) empty(
	tzSpec *TimeZoneSpecification) {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("timeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
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
func (tzSpecUtil *timeZoneSpecUtility) equal(
	tzSpec *TimeZoneSpecification,
	tzSpec2 TimeZoneSpecification) bool {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	if tzSpec == nil {
		panic("timeZoneSpecUtility.empty()\n" +
			"Error: Input parameter tzSpec is a 'nil' pointer!\n")
	}

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
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
// isEmpty() returns a boolean value of 'true' if all
// data field values are set to their empty or zero
// values.
//
func (tzSpecUtil *timeZoneSpecUtility) isEmpty(
	tzSpec * TimeZoneSpecification,
	ePrefix string) bool {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	ePrefix += "timeZoneSpecUtility.setFromTimeZoneSpec() "

	if tzSpec == nil {
		panic(ePrefix + "\nError: " +
			"Input parameter 'tzSpec' is a 'nil' pointer!\n")
	}

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

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

	if tzSpec.locationNameType != LocNameType.None(){
		return false
	}

	if tzSpec.timeZoneType != TzType.None() {
		return false
	}

	if tzSpec.timeZoneClass != TzClass.None() {
		return false
	}

	if tzSpec.tagDescription != "" {
		return false
	}

	return true
}

// SetTimeZone - Sets the data values of the input parameter
// 'tzSpec', an instance of type TimeZoneSpecification.
//
func (tzSpecUtil *timeZoneSpecUtility) setTimeZone(
	tzSpec *TimeZoneSpecification,
	referenceDateTime      time.Time,
	militaryTimeZoneLetter string,
	militaryTimeZoneName   string,
	zoneLabel              string,
	tagDescription         string,
	timeZoneClass          TimeZoneClass,
	ePrefix string) error {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	ePrefix += "timeZoneSpecUtility.setFromTimeZoneSpec() "

	if tzSpec == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzSpec",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tzSpec' is a nil pointer!",
			err:                 nil,
		}
	}

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}

	if referenceDateTime.IsZero() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "referenceDateTime",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'referenceDateTime' has a Zero Value!",
			err:                 nil,
		}
	}

	tzSpecUtil2 := timeZoneSpecUtility{}

	tzSpecUtil2.empty(tzSpec)

	tzMech := TimeZoneMechanics{}
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
		err = tzMech.CalcUtcZoneOffsets(referenceDateTime, ePrefix)


	var timeZoneType TimeZoneType
	var ok bool

	if err != nil {
		tzSpecUtil2.empty(tzSpec)
		return err
	}

	locNameType := LocationNameType(0).ConvertibleTimeZone()

	dtMech := DTimeMechanics{}

	// Test For Location Name Type
	_, err = dtMech.LoadTzLocation(tzSpec.locationName, ePrefix)

	if err != nil {
		locNameType = LocationNameType(0).NonConvertibleTimeZone()
	}

	// Test for Time Zone Type
	if len(militaryTimeZoneName) > 0 ||
		len(militaryTimeZoneLetter) > 0 {

		var foundMilTextName string

		milTzDat := MilitaryTimeZoneData{}

		foundMilTextName, ok = milTzDat.MilTzLetterToTextName(militaryTimeZoneLetter)

		if !ok {

			tzSpecUtil2.empty(tzSpec)

			return fmt.Errorf(ePrefix +
				"\nInput parameter 'militaryTimeZoneLetter' is Invalid!\n" +
				"militaryTimeZoneLetter='%v'\n", militaryTimeZoneLetter)
		}

		if foundMilTextName != militaryTimeZoneName {

			tzSpecUtil2.empty(tzSpec)

			return fmt.Errorf(ePrefix +
				"\nInput parameter 'militaryTimeZoneName' is Invalid!\n" +
				"militaryTimeZoneName='%v'\n" +
				"The correct military Time Zone Name is '%v'\n",
				militaryTimeZoneName, foundMilTextName)
		}

		timeZoneType = TzType.Military()

	} else if strings.ToLower(tzSpec.locationName) == "local" {

		timeZoneType = TzType.Local()

	} else {

		timeZoneType = TzType.Iana()

	}

	// Test for Time Zone Utc Offset Status
	var tzUtcOffsetStatus TimeZoneUtcOffsetStatus

	tzUtcOffsetStatus, err = tzMech.GetTimeZoneUtcOffsetStatus(tzSpec.locationPtr, ePrefix)

	if err != nil {
		tzSpecUtil2.empty(tzSpec)
		return err
	}

	// Test for Time Zone Category
	var tzCategory TimeZoneCategory

	firstLetter := tzSpec.locationName[0:1]

	if firstLetter == "+" ||
		firstLetter == "-" {
		tzCategory = TimeZoneCategory(0).UtcOffset()
	} else {
		tzCategory = TimeZoneCategory(0).TextName()
	}

	tzSpec.referenceDateTime = referenceDateTime
	tzSpec.zoneLabel = zoneLabel
	tzSpec.militaryTimeZoneLetter = militaryTimeZoneLetter
	tzSpec.militaryTimeZoneName = militaryTimeZoneName
	tzSpec.tagDescription = tagDescription
	tzSpec.locationNameType = locNameType
	tzSpec.timeZoneCategory = tzCategory
	tzSpec.timeZoneClass = timeZoneClass
	tzSpec.timeZoneType = timeZoneType
	tzSpec.timeZoneUtcOffsetStatus = tzUtcOffsetStatus

	return nil
}

// setFromTimeZoneSpec - Sets the data fields for
// 'tzSpec' based on Time Zone Specification, 'tzSpecIn'.
//
func (tzSpecUtil *timeZoneSpecUtility) setFromTimeZoneSpec(
	tzSpec *TimeZoneSpecification,
	dateTime time.Time,
	timeZoneConversionType TimeZoneConversionType,
	tzSpecIn *TimeZoneSpecification,
	ePrefix string) error {

	tzSpecUtil.lock.Lock()

	defer tzSpecUtil.lock.Unlock()

	ePrefix += "timeZoneSpecUtility.setFromTimeZoneSpec() "

	if tzSpec == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzSpec",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tzSpec' is a nil pointer!",
			err:                 nil,
		}
	}

	if tzSpec.lock == nil {
		tzSpec.lock = new(sync.Mutex)
	}


	if tzSpecIn == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzSpecIn",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tzSpecIn' is a nil pointer!",
			err:                 nil,
		}
	}

	if tzSpecIn.lock == nil {
		tzSpecIn.lock = new(sync.Mutex)
	}

	if dateTime.IsZero() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'dateTime' has a Zero Value!",
			err:                 nil,
		}
	}

	err := tzSpec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nInput Parameter 'tzSpec' is Invalid!\n" +
			"Error='%v'\n", err.Error())
	}

	err = tzSpecIn.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nInput Parameter 'tzSpecIn' is Invalid!\n" +
			"Error='%v'\n", err.Error())
	}

	if timeZoneConversionType < TzConvertType.Absolute() ||
		timeZoneConversionType > TzConvertType.Relative() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneConversionType",
			inputParameterValue: "",
			errMsg:              "Input Parameter timeZoneConversionType must be 'Absolute' or 'Relative' !",
			err:                 nil,
		}
	}

	if timeZoneConversionType == TzConvertType.Absolute() {
		dateTime = time.Date(
			dateTime.Year(),
			dateTime.Month(),
			dateTime.Day(),
			dateTime.Hour(),
			dateTime.Minute(),
			dateTime.Second(),
			dateTime.Nanosecond(),
			tzSpecIn.locationPtr)
	} else {
		// Must be TzConvertType.Relative()

		dateTime = dateTime.In(tzSpecIn.locationPtr)
	}

	tzSpecUtil2 := timeZoneSpecUtility{}

	tzSpecUtil2.empty(tzSpec)

	tzSpecUtil2.copyIn(tzSpec, tzSpecIn)

	tzSpec.referenceDateTime = dateTime

	return nil
}