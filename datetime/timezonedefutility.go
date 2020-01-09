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

// SetFromDateTimeComponents - Re-initializes the values of the
// 'TimeZoneDefDto' instance based on input parameter, 'dateTime'.
//
/*
func (tzDefUtil *timeZoneDefUtility) setFromDateTime(
	tzdef *TimeZoneDefDto,
	dateTime time.Time,
	ePrefix string) error {

	ePrefix += "timeZoneDefUtility.setFromDateTime() "

	locPtr := dateTime.Location()
	fmtStr := "01/02/2006 15:04:05.000000000 -0700 MST"

	if locPtr == nil {
		return fmt.Errorf(ePrefix +
			"Error: Input Parameter 'dateTime' has a nil Location Pointer!\n" +
			"dateTime='%v'\n", dateTime.Format(fmtStr))
	}

	zoneName, zoneSeconds := dateTime.Zone()

	locName := locPtr.String()

	if locName != zoneName {
		// Maybe a good Iana Time Zone!
	}

	timeStr := dateTime.Format(fmtStr)

	offsetLeadLen := len("01/02/2006 15:04:05.000000000 ")

	t2AbbrvLookup := locName + timeStr[offsetLeadLen:offsetLeadLen+5]

	stdAbbrvs := StdTZoneAbbreviations{}

	tZones, ok := stdAbbrvs.AbbrvOffsetToTimeZones(t2AbbrvLookup)

	if !ok {
		return fmt.Errorf(ePrefix +
			"\nError: Could NOT location Time Zone Abbreviation!\n" +
			"Mapping Time Zone Abbreviation to Time Zones Failed.\n" +
			"Lookup key='%v'\n", t2AbbrvLookup)
	}

	var newTZone string

	if len(tZones) == 1 {
		newTZone = tZones[0]
	}

	if len(newTZone) == 0 {
		// tzAbbrvToTimeZonePriorityList
		for i:=0; i < len(tzAbbrvToTimeZonePriorityList) && len(newTZone)== 0 ; i++ {

			for j:=0; j < len(tZones); j++ {

				if strings.HasPrefix(tZones[j], priorityList[i]) {
					newTZone = tZones[j]
					break
				}
			}
		}
	}

	if len(newTZone) == 0 {
		newTZone = tZones[0]
	}


	return nil
}


func (tzDefUtil *timeZoneDefUtility) isConvertibleTimeZone(
	dateTime time.Time,
	ePrefix string) error {

		ePrefix += "timeZoneDefUtility.isConvertibleTimeZone()"



	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'dateTime' is a ZERO value!")
	}

	fmtStr := "2006-01-02 15:04:05 -0700 MST"

	tzAbbrFmtStr := "2006-01-02 15:04:05 -0700 "

	lenAbbrFmtStr := len(tzAbbrFmtStr)

	d1TzAbbrv := dateTime.Format(fmtStr)[lenAbbrFmtStr:]

	d1TzAbbrv = strings.TrimLeft(strings.TrimRight(d1TzAbbrv, " "), " ")

	_,
	_,
	ianaTimeZoneName,
	ianaLocationPtr,
	err := dtMech.convertTzAbbreviationToTimeZone(d1TzAbbrv, ePrefix)

	if err != nil {
		return err
	}

	dt1June := time.Date(
		time.Now().Year(),
		time.Month(06),
		30,
		9,
		0,
		0,
		0,
		dateTime.Location())


	dt1Dec :=  

	return nil
}
*/

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

	dtMech :=  dateTimeMechanics{}

	ianaTimeZonePtr := dateTime.Location()

	if ianaTimeZonePtr == nil {
		return fmt.Errorf(ePrefix +
			"\nAttempt to load dateTime.Location() pointer FAILED!\n" +
			"Returned pointer is nil.\n" +
			"dateTime='%v'",dateTime.Format(FmtDateTimeTzNanoYMD))
	}

	var err error

	ianaTimeZoneName := ianaTimeZonePtr.String()

	_, err =  dtMech.loadTzLocationPtr(ianaTimeZoneName, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: Attempt to load dateTime Time Zone Name FAILED!\n" +
			"ianaTimeZoneName='%v'\ndateTime='%v'\n",
			ianaTimeZoneName, dateTime.Format(FmtDateTimeTzNanoYMD))
	}

	milTzLetter := ""
	milTzName := ""

	var timeZoneType TimeZoneType

	testTzName := strings.ToLower(ianaTimeZoneName)

	if testTzName == "local" {

		timeZoneType = TzType.Local()

	} else {

		timeZoneType = TzType.Iana()

	}

	zoneLabel := "Original Time Zone"

	tzSpec := TimeZoneSpecDto{}

	err = tzSpec.SetTimeZone(
		dateTime,
		milTzLetter,
		milTzName,
		zoneLabel,
		"",
		LocNameType.ConvertibleTimeZoneName(),
		timeZoneType,
		ePrefix)

	if err != nil {
		return err
	}

	tzdef.originalTimeZone = tzSpec.CopyOut()

	tzdef.convertibleTimeZone = tzSpec.CopyOut()

	return nil
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

	if zoneLabel != ianaTimeZoneName {
		tzdef.originalTimeZone.locationNameType = LocNameType.ConvertibleTimeZoneName()
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

	}

	return err
}

// setZoneProfile - assembles and assigns the composite zone
// offset, zone names, zone abbreviation and UTC offsets.
//
// The TimeZoneDefDto.ZoneOffset field formatted in accordance
// with the following examples:
//      "-0600 CST"
//      "+0200 EET"
//
func (tzDefUtil *timeZoneDefUtility) calcZoneOffsets(
	offsetHours int,
	offsetMinutes int,
	offsetSeconds int,
	zoneSign int,
	zoneName string,
	ePrefix string) (
	zoneOffset string,
	utcOffset string,
	err error) {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	ePrefix += "timeZoneDefUtility.calcZoneOffsets() "
	zoneOffset = ""
	utcOffset = ""
	err = nil

	if zoneSign < -1 ||
		zoneSign > 1 ||
		zoneSign == 0 {
		return zoneOffset, utcOffset,
		fmt.Errorf(ePrefix,
			"Error: Input parameter 'zoneSign' must be equal to -1 or +1.\n" +
			"zoneSign='%v'\n", zoneSign)
	}

	if offsetHours < 0 {
		return zoneOffset, utcOffset,
			fmt.Errorf(ePrefix,
				"Error: Input parameter 'offsetHours' is less than zero.\n" +
					"offsetHours='%v'\n", offsetHours)
	}

	if offsetMinutes < 0 {
		return zoneOffset, utcOffset,
			fmt.Errorf(ePrefix,
				"Error: Input parameter 'offsetMinutes' is less than zero.\n" +
					"offsetMinutes='%v'\n", offsetMinutes)
	}

	if offsetSeconds < 0 {
		return zoneOffset, utcOffset,
			fmt.Errorf(ePrefix,
				"Error: Input parameter 'offsetSeconds' is less than zero.\n" +
					"offsetSeconds='%v'\n", offsetSeconds)
	}

	if len(zoneName) == 0 {
		return zoneOffset, utcOffset,
			errors.New(ePrefix +
				"Error: Input parameter 'zoneName' is an empty string.\n")
	}

	// Generates an offset in the form of "+0330" or "-0330"
	if zoneSign == -1 {
		zoneOffset += "-"
	} else {
		zoneOffset += "+"
	}

	zoneOffset += fmt.Sprintf("%02d%02d", offsetHours, offsetMinutes)

	// Generates final UTC offset in the form
	// "-0500" or "+0200"
	utcOffset = zoneOffset

	if offsetSeconds > 0 {
		zoneOffset += fmt.Sprintf("%02d", offsetSeconds)
	}

	// Generates final ZoneOffset in the form
	// "-0500 CST" or "+0200 EET"
	zoneOffset += " " + zoneName

	return zoneOffset, utcOffset, err
}
