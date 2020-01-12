package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type timeZoneDefUtility struct {
	lock sync.Mutex
}

// CopyIn - Copies an incoming TimeZoneDefDto into the
// data fields of the current TimeZoneDefDto instance.
//
func (tzDefUtil *timeZoneDefUtility) copyIn(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.CopyIn()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.CopyIn()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	tzDefUtil2 := timeZoneDefUtility{}

	tzDefUtil2.empty(tzdef)

	tzdef.originalTimeZone = tzdef2.originalTimeZone.CopyOut()
	tzdef.convertibleTimeZone = tzdef2.convertibleTimeZone.CopyOut()

	return
}

// copyOut - creates and returns a deep copy of the current
// TimeZoneDefDto instance.
//
func (tzDefUtil *timeZoneDefUtility) copyOut(
	tzdef *TimeZoneDefDto) TimeZoneDefDto {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.copyOut()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	tzdef2 := TimeZoneDefDto{}
	tzdef2.originalTimeZone = tzdef.originalTimeZone.CopyOut()
	tzdef2.convertibleTimeZone = tzdef.convertibleTimeZone.CopyOut()

	return tzdef2
}

// Empty - Resets all field values for the input parameter
// TimeZoneDefDto to their uninitialized or 'zero' states.
//
func (tzDefUtil *timeZoneDefUtility) empty(
	tzdef *TimeZoneDefDto) {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.empty()\n" +
			"Error: 'tzdef' pointer is nil!\n")
	}
	tzdef.originalTimeZone.Empty()
	tzdef.convertibleTimeZone.Empty()

	return
}


// Equal - Determines if two TimeZoneDefDto instances are
// equivalent in value.
//
// This method returns 'true' of two TimeZoneDefDto's are
// equal in all respects.
//
func (tzDefUtil *timeZoneDefUtility) equal(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equal() " +
			"\nError: Input parameter 'tzdef' is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equal() " +
			"\nError: Input parameter 'tzdef2' is nil!")
	}

	tzDefUtil2 := timeZoneDefUtility{}

	tzdefIsEmpty := tzDefUtil2.isEmpty(tzdef)

	tzdef2IsEmpty := tzDefUtil2.isEmpty(tzdef2)

	if tzdefIsEmpty == true &&
		tzdef2IsEmpty == true {
		return true
	}

	if !tzdef.originalTimeZone.Equal(tzdef2.originalTimeZone) ||
		!tzdef.convertibleTimeZone.Equal(tzdef2.convertibleTimeZone) {
		return false
	}

	return true
}

// equalLocations - Compares the Time Zone Locations for two TimeZoneDefDto's
// and returns 'true' if they are equal.
//
// Time Zone Location Name Examples:
//   "Local"
//   "America/Chicago"
//   "America/New_York"
//
func (tzDefUtil *timeZoneDefUtility) equalLocations(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	 if tzdef.originalTimeZone.locationName != tzdef2.originalTimeZone.locationName {
	 	return false
	 }

	 if tzdef.convertibleTimeZone.locationName != tzdef2.convertibleTimeZone.locationName {
	 	return false
	 }

	 return true
}

// equalOffsetSeconds - Compares Zone Offset Seconds for two TimeZoneDefDto's and
// returns 'true' if they are equal.
//
// ZoneOffsetSeconds is a signed number of seconds offset from UTC:
//   + == East of UTC
//   - == West of UTC
//
func (tzDefUtil *timeZoneDefUtility) equalOffsetSeconds(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	if tzdef.originalTimeZone.zoneOffsetTotalSeconds !=
		tzdef2.originalTimeZone.zoneOffsetTotalSeconds {
		return false
	}
	
	if tzdef.convertibleTimeZone.zoneOffsetTotalSeconds !=
		tzdef2.convertibleTimeZone.zoneOffsetTotalSeconds {
		return false
	}
	
	return true
}

// equalZoneLocation - Compares two TimeZoneDefDto's and returns
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
func (tzDefUtil *timeZoneDefUtility) equalZoneLocation(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equalZoneLocation()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	if tzdef.originalTimeZone.locationName != 
			tzdef2.originalTimeZone.locationName ||
	 tzdef.originalTimeZone.zoneName != 
			tzdef2.originalTimeZone.zoneName ||
		tzdef.originalTimeZone.zoneOffset != 
		tzdef2.originalTimeZone.zoneOffset {
		return false
	}

	if tzdef.convertibleTimeZone.locationName != 
			tzdef2.convertibleTimeZone.locationName ||
	 tzdef.convertibleTimeZone.zoneName != 
			tzdef2.convertibleTimeZone.zoneName ||
		tzdef.convertibleTimeZone.zoneOffset != 
		tzdef2.convertibleTimeZone.zoneOffset {
		return false
	}

	return true
}

// equalZoneOffsets - Compares ZoneOffsets for two TimeZoneDefDto's and
// returns 'true' if they are equal.
//
// Zone Offset is a text string representing the offset from UTC plus the
// time zone abbreviation.
//
// Example "-0500 CDT"
//
func (tzDefUtil *timeZoneDefUtility) equalZoneOffsets(
	tzdef *TimeZoneDefDto,
	tzdef2 *TimeZoneDefDto) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.equalZoneOffsets()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef2 == nil {
		panic("timeZoneDefUtility.equalZoneOffsets()\n" +
			"Error: Input parameter 'tzdef2' pointer is nil!\n")
	}

	if tzdef.originalTimeZone.zoneOffset !=
			tzdef2.originalTimeZone.zoneOffset {
		return false
	}

	if tzdef.convertibleTimeZone.zoneOffset !=
			tzdef2.convertibleTimeZone.zoneOffset {
		return false
	}

	return true
}

// isEmpty - Determines whether the current TimeZoneDefDto
// instance is Empty.
//
// If the TimeZoneDefDto instance (tzdef) is NOT populated,
// this method returns 'true'. Otherwise, it returns 'false'.
//
func (tzDefUtil *timeZoneDefUtility) isEmpty(
	tzdef *TimeZoneDefDto ) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.isValidFromDateTime()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if tzdef.originalTimeZone.IsEmpty() &&
		tzdef.convertibleTimeZone.IsEmpty() {
		return true
	}

	return false
}

// isValidTimeZoneDefDto - Analyzes the TimeZoneDefDto
// parameter, 'tzdef', instance to determine validity.
//
// This method returns 'true' if the TimeZoneDefDto
// instance is valid.  Otherwise, it returns 'false'.
//
func (tzDefUtil *timeZoneDefUtility) isValidTimeZoneDefDto(
	tzdef *TimeZoneDefDto,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.isValidTimeZoneDefDto() "

	if tzdef == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzdef' is a 'nil' pointer!\n")
	}

	controlErrors := make([]error, 0)

  err := tzdef.originalTimeZone.IsValid(ePrefix)

  if err!= nil {
  	controlErrors = append(controlErrors, err)
	}

  err = tzdef.convertibleTimeZone.IsValid(ePrefix)

  if err!= nil {
  	controlErrors = append(controlErrors, err)
	}

	dtUtil := DTimeUtility{}

	return dtUtil.ConsolidateErrors(controlErrors)
}

// isValidFromDateTime - Uses a time.Time input parameter, 'dateTime' to
// analyze the specified TimeZoneDefDto instance (tzdef). If the zone and
// location details of 'dateTime' are not perfectly matched to the current
// TimeZoneDefDto instance, the instance is considered INVALID, and this
// method returns 'false'.
//
// Otherwise, if all zone and location details are perfectly matched, this
// method returns 'true', signaling that the TimeZoneDateDefDto instance
// (tzdef) is VALID.
//
func (tzDefUtil *timeZoneDefUtility) isValidFromDateTime(
	tzdef *TimeZoneDefDto,
	dateTime time.Time) bool {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.isValidFromDateTime()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	if dateTime.IsZero() {
		return false
	}

	return true
}

// SetFromDateTimeComponents - Re-initializes the values of a
// 'TimeZoneDefDto' instance based on time components (i.e.
// years, months, days, hours, minutes, seconds and nanoseconds)
// passed through input parameter 'TimeDto' ('tDto').
//
func (tzDefUtil *timeZoneDefUtility) setFromDateTimeComponents(
	tzdef *TimeZoneDefDto,
	tDto TimeDto,
	timeZoneName string,
	ePrefix string) error {

	ePrefix += "timeZoneDefUtility.setFromDateTimeComponents() "

	if tzdef == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzdef",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tzdef' pointer is nil!",
			err:                 nil,
		}
	}

	utcLocPtr, err := time.LoadLocation(TZones.UTC())

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by time.LoadLocation(TZones.UTC())\n" +
			"Error='%v'\n", err.Error())
	}

	dateTime := time.Date(tDto.Years,
		time.Month(tDto.Months),
		tDto.DateDays,
		tDto.Hours,
		tDto.Minutes,
		tDto.Seconds,
		tDto.TotSubSecNanoseconds,
		utcLocPtr)

	tzDefUtil2 := timeZoneDefUtility{}

	return tzDefUtil2.setFromTimeZoneName(
		tzdef,
		dateTime,
		timeZoneName,
		TzConvertType.Absolute(),
		ePrefix)
}

// setFromDateTime - Sets the values of a TimeZoneDefDto
// based on input parameter 'dateTime'.
//
func (tzDefUtil *timeZoneDefUtility) setFromDateTime(
	tzdef *TimeZoneDefDto,
	dateTime time.Time,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.setFromDateTime() "

	if tzdef == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tzdef",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tzdef' pointer is nil!",
			err:                 nil,
		}
	}

	if dateTime.IsZero() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "dateTime",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'dateTime' is ZERO!",
			err:                 nil,
		}
	}

	tzMech := timeZoneMechanics{}

	var ianaTimeZoneName string
	var err error

	ianaTimeZoneName,
	_,
	_,
	err =
		tzMech.getConvertibleTimeZoneFromDateTime(
			dateTime,
			ePrefix)

	if err != nil {
		return err
	}

	tzDefUtil2 := timeZoneDefUtility{}

	return tzDefUtil2.setFromTimeZoneName(
		tzdef,
		dateTime,
		ianaTimeZoneName,
		TzConvertType.Absolute(),
		ePrefix)
}

// setFromTimeZoneName - Sets the data fields of the specified
// TimeZoneDefDto instance base on the time zone text name.
//
func (tzDefUtil *timeZoneDefUtility) setFromTimeZoneName(
	tzdef *TimeZoneDefDto,
	dateTime time.Time,
	timeZoneName string,
	timeZoneConversionType TimeZoneConversionType,
	ePrefix string) error {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.setFromTimeZoneName() "

	if tzdef == nil {
		return errors.New(ePrefix +
			"\nInput parameter 'tzdef' is nil!\n")
	}

	timeZoneName = strings.TrimLeft(strings.TrimRight(timeZoneName, " "), " ")

	if len(timeZoneName) == 0 {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneName",
			inputParameterValue: "",
			errMsg:              "Input parameter 'timeZoneName' is an empty string!",
			err:                 nil,
		}
	}

	tzMech := timeZoneMechanics{}

	milTzLetter,
	milTzName,
	ianaTimeZoneName,
	ianaLocationPtr,
	tzType,
	err := tzMech.getTimeZoneFromName(
		timeZoneName,
		ePrefix)

	if err != nil {
		return err
	}

	zoneLabel := "Original Time Zone"

	switch timeZoneConversionType {

	case TzConvertType.Relative():

		dateTime = dateTime.In(ianaLocationPtr)

	case TzConvertType.Absolute():

		dateTime = time.Date(dateTime.Year(),
			dateTime.Month(),
			dateTime.Day(),
			dateTime.Hour(),
			dateTime.Minute(),
			dateTime.Second(),
			dateTime.Nanosecond(),
			ianaLocationPtr)

	default:
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeZoneConversionType",
			inputParameterValue: timeZoneConversionType.String(),
			errMsg:              "Input parameter 'timeZoneName' is not equal to 'Absolute' and not equal to 'Relative'!",
			err:                 nil,
		}
	}

	tzSpec := TimeZoneSpecDto{}

	err = tzSpec.SetTimeZone(
		dateTime,
		milTzLetter,
		milTzName,
		zoneLabel,
		"",
		LocNameType.ConvertibleTimeZoneName(),
		tzType,
		ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: tzdef.originalTimeZone failed to initialize.\n" +
			"Error='%v'", err.Error())
	}

	tzdef.originalTimeZone = tzSpec.CopyOut()

	zoneLabel, _ = dateTime.Zone()

	firstTzLetter := ianaTimeZoneName[0:1]

	if zoneLabel != ianaTimeZoneName &&
		firstTzLetter != "+" &&
		firstTzLetter != "-" {
		tzdef.originalTimeZone.locationNameType = LocNameType.ConvertibleTimeZoneName()
		tzdef.convertibleTimeZone = tzdef.originalTimeZone.CopyOut()
		return nil
	}

	if ianaTimeZoneName == "UTC" {
		tzdef.originalTimeZone.locationNameType = LocNameType.ConvertibleAbbreviation()
		tzdef.convertibleTimeZone = tzdef.originalTimeZone.CopyOut()
		return nil
	}

	tzdef.originalTimeZone.locationNameType = LocNameType.NonConvertibleAbbreviation()

	var tzAbbrvLookupId string

	tzAbbrvLookupId, err = tzMech.getTzAbbrvLookupIdFromDateTime(dateTime, ePrefix)

	if err != nil {
		tzdef.originalTimeZone.Empty()

		return fmt.Errorf(ePrefix +
			"\nAttempted creation of Time Zone Abbreviation Look-Up Id Failed!\n" +
			"Error='%v'\n", err.Error())
	}

	milTzLetter,
	milTzName,
	ianaTimeZoneName,
	ianaLocationPtr,
	err = tzMech.convertTzAbbreviationToTimeZone(tzAbbrvLookupId, ePrefix)

	if err != nil {
		tzdef.originalTimeZone.Empty()
		return fmt.Errorf(ePrefix +
			"\nError: Non-Convertible Time Zone Abbreviation Lookup Failed!\n" +
			"Error='%v'\n", err.Error())
	}

	tzSpec = TimeZoneSpecDto{}

	dateTime = dateTime.In(ianaLocationPtr)

	zoneLabel = "Convertible Time Zone"

	err = tzSpec.SetTimeZone(
		dateTime,
		milTzLetter,
		milTzName,
		zoneLabel,
		"",
		LocNameType.ConvertibleTimeZoneName(),
		tzType,
		ePrefix)

	if err != nil {

		tzdef.originalTimeZone.Empty()
		tzdef.convertibleTimeZone.Empty()

		return err
	}

	tzdef.convertibleTimeZone = tzSpec.CopyOut()
	tzdef.convertibleTimeZone.locationNameType = LocNameType.ConvertibleAbbreviation()

	return nil
}
