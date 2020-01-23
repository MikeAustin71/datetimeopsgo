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

// CopyIn - Copies an incoming TimeZoneDefinition into the
// data fields of the current TimeZoneDefinition instance.
//
func (tzDefUtil *timeZoneDefUtility) copyIn(
	tzdef *TimeZoneDefinition,
	tzdef2 *TimeZoneDefinition) {

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
// TimeZoneDefinition instance.
//
func (tzDefUtil *timeZoneDefUtility) copyOut(
	tzdef *TimeZoneDefinition) TimeZoneDefinition {

	tzDefUtil.lock.Lock()

	defer tzDefUtil.lock.Unlock()

	if tzdef == nil {
		panic("timeZoneDefUtility.copyOut()\n" +
			"Error: Input parameter 'tzdef' pointer is nil!\n")
	}

	tzdef2 := TimeZoneDefinition{}
	tzdef2.originalTimeZone = tzdef.originalTimeZone.CopyOut()
	tzdef2.convertibleTimeZone = tzdef.convertibleTimeZone.CopyOut()

	return tzdef2
}

// Empty - Resets all field values for the input parameter
// TimeZoneDefinition to their uninitialized or 'zero' states.
//
func (tzDefUtil *timeZoneDefUtility) empty(
	tzdef *TimeZoneDefinition) {

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


// Equal - Determines if two TimeZoneDefinition instances are
// equivalent in value.
//
// This method returns 'true' of two TimeZoneDefinition's are
// equal in all respects.
//
func (tzDefUtil *timeZoneDefUtility) equal(
	tzdef *TimeZoneDefinition,
	tzdef2 *TimeZoneDefinition) bool {

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

// equalLocations - Compares the Time Zone Locations for two TimeZoneDefinition's
// and returns 'true' if they are equal.
//
// Time Zone Location Name Examples:
//   "Local"
//   "America/Chicago"
//   "America/New_York"
//
func (tzDefUtil *timeZoneDefUtility) equalLocations(
	tzdef *TimeZoneDefinition,
	tzdef2 *TimeZoneDefinition) bool {

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

// equalOffsetSeconds - Compares Zone Offset Seconds for two TimeZoneDefinition's and
// returns 'true' if they are equal.
//
// ZoneOffsetSeconds is a signed number of seconds offset from UTC:
//   + == East of UTC
//   - == West of UTC
//
func (tzDefUtil *timeZoneDefUtility) equalOffsetSeconds(
	tzdef *TimeZoneDefinition,
	tzdef2 *TimeZoneDefinition) bool {

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

// equalZoneLocation - Compares two TimeZoneDefinition's and returns
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
	tzdef *TimeZoneDefinition,
	tzdef2 *TimeZoneDefinition) bool {

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

// equalZoneOffsets - Compares ZoneOffsets for two TimeZoneDefinition's and
// returns 'true' if they are equal.
//
// Zone Offset is a text string representing the offset from UTC plus the
// time zone abbreviation.
//
// Example "-0500 CDT"
//
func (tzDefUtil *timeZoneDefUtility) equalZoneOffsets(
	tzdef *TimeZoneDefinition,
	tzdef2 *TimeZoneDefinition) bool {

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

// isEmpty - Determines whether the current TimeZoneDefinition
// instance is Empty.
//
// If the TimeZoneDefinition instance (tzdef) is NOT populated,
// this method returns 'true'. Otherwise, it returns 'false'.
//
func (tzDefUtil *timeZoneDefUtility) isEmpty(
	tzdef *TimeZoneDefinition) bool {

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

// isValidTimeZoneDefDto - Analyzes the TimeZoneDefinition
// parameter, 'tzdef', instance to determine validity.
//
// This method returns 'true' if the TimeZoneDefinition
// instance is valid.  Otherwise, it returns 'false'.
//
func (tzDefUtil *timeZoneDefUtility) isValidTimeZoneDefDto(
	tzdef *TimeZoneDefinition,
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
// analyze the specified TimeZoneDefinition instance (tzdef). If the zone and
// location details of 'dateTime' are not perfectly matched to the current
// TimeZoneDefinition instance, the instance is considered INVALID, and this
// method returns 'false'.
//
// Otherwise, if all zone and location details are perfectly matched, this
// method returns 'true', signaling that the TimeZoneDateDefDto instance
// (tzdef) is VALID.
//
func (tzDefUtil *timeZoneDefUtility) isValidFromDateTime(
	tzdef *TimeZoneDefinition,
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
// 'TimeZoneDefinition' instance based on time components (i.e.
// years, months, days, hours, minutes, seconds and nanoseconds)
// passed through input parameter 'TimeDto' ('tDto').
//
func (tzDefUtil *timeZoneDefUtility) setFromTimeDto(
	tzdef *TimeZoneDefinition,
	tDto TimeDto,
	timeZoneName string,
	ePrefix string) error {

	ePrefix += "timeZoneDefUtility.setFromTimeDto() "

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

// setFromDateTime - Sets the values of a TimeZoneDefinition
// based on input parameter 'dateTime'. Note: TimeZoneDefinition
// objects set from date times may NOT be configured as Military
// Time Zones.
//
func (tzDefUtil *timeZoneDefUtility) setFromDateTime(
	tzdef *TimeZoneDefinition,
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

	if dateTime.Location() == nil {
		return &TimeZoneError{
			ePrefix: ePrefix,
			errMsg:  "Error: dateTime.Location() returned a 'nil' pointer!",
			err:     nil,
		}
	}

	tzMech := TimeZoneMechanics{}
	var tzAbbrv, utcOffset string
	var err error
	var tzSpec1, tzSpec2 TimeZoneSpecification

	dtMech := DateTimeMechanics{}

	utcOffset, tzAbbrv, err =
		tzMech.GetUtcOffsetTzAbbrvFromDateTime(dateTime, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"\n"+
			"Load Location Failed. Error returned extracting UTC Offset, Tz Abreviation.\n"+
			"%v", err.Error())
	}

	_, err = dtMech.LoadTzLocationPtr(
		dateTime.Location().String(),
		ePrefix)

	if err == nil {
		// The Original Time Zone Loaded Successfully!

		tzSpec1,
			err =
			tzMech.GetConvertibleTimeZoneFromDateTime(
				dateTime,
				TzConvertType.Absolute(),
				"Original Time Zone",
				ePrefix)

		if err != nil {
			return err
		}

		tzSpec1.timeZoneClass.OriginalTimeZone()
		tzdef.originalTimeZone = tzSpec1.CopyOut()

		tzSpec1.zoneLabel = "Convertible Time Zone"
		tzdef.convertibleTimeZone = tzSpec1.CopyOut()

		return nil
	}

	// Original Time Zone will NOT Load!
	// Try to find an equivalent substitute
	// Time Zone!

	if !tzMech.IsTzAbbrvUtcOffset(tzAbbrv) {
		// Original Time Zone Did NOT Load AND,
		// The Time Zone Abbreviation is NOT a UTC Offset
		tzSpec1,
			err = tzMech.ConvertTzAbbreviationToTimeZone(
			dateTime,
			TzConvertType.Absolute(),
			tzAbbrv+utcOffset,
			"Convertible Time Zone",
			ePrefix)

		if err != nil {
			// Original Time Zone failed to Load AND,
			// all efforts to find a satisfactory substitution
			// FAILED!
			return fmt.Errorf(ePrefix+"\n"+
				"Load Location Failed. The time zone name is invalid!\n"+
				"Time Zone Name: '%v'\n"+
				"dateTime= '%v'\n",
				dateTime.Location().String(),
				dateTime.Format(FmtDateTimeTzNanoYMD))
		}

		tzSpec1.locationNameType = LocNameType.ConvertibleTimeZone()
		tzSpec1.timeZoneClass = TzClass.AlternateTimeZone()

	} else {
		// Original Time Zone Did NOT Load AND,
		// The Time Zone Abbreviation IS a UTC Offset

		tzSpec1,
			err = tzMech.ConvertUtcAbbrvToStaticTz(
			dateTime,
			TzConvertType.Absolute(),
			"Original Time Zone",
			tzAbbrv,
			ePrefix)

		if err == nil {
			// Successfully Loaded Static UTC Time Zone
			tzSpec1.locationNameType = LocNameType.ConvertibleTimeZone()
			tzSpec1.timeZoneClass = TzClass.OriginalTimeZone()


			tzSpec2,
				err = tzMech.ConvertTzAbbreviationToTimeZone(
				dateTime,
				TzConvertType.Absolute(),
				tzAbbrv+utcOffset,
				"Convertible Time Zone",
				ePrefix)

			if err != nil {
				return err
			}

			tzSpec2.locationNameType = LocNameType.ConvertibleTimeZone()
			tzSpec2.timeZoneClass = TzClass.AlternateTimeZone()

			tzdef.originalTimeZone = tzSpec1.CopyOut()
			tzdef.convertibleTimeZone = tzSpec2.CopyOut()

			return nil

		} else {
			// Original Time Zone Did NOT Load AND,
			// Attempted loading of UTC Offset, static time zone
			// Failed. Now, try an alternate equivalent time zone.

			tzSpec1,
				err = tzMech.ConvertTzAbbreviationToTimeZone(
				dateTime,
				TzConvertType.Absolute(),
				tzAbbrv+utcOffset,
				"Convertible Time Zone",
				ePrefix)

			if err != nil {
				return err
			}

			tzSpec1.locationNameType = LocNameType.ConvertibleTimeZone()
			tzSpec1.timeZoneClass = TzClass.AlternateTimeZone()

		}
	}

	tzSpec2,
		err = tzMech.ConvertTzAbbreviationToTimeZone(
		dateTime,
		TzConvertType.Absolute(),
		tzAbbrv+utcOffset,
		"Convertible Time Zone",
		ePrefix)

	if err != nil {
		return err
	}

	tzdef.originalTimeZone = tzSpec1.CopyOut()

	tzSpec2.locationNameType = LocNameType.ConvertibleTimeZone()
	tzSpec2.timeZoneClass = TzClass.AlternateTimeZone()

	tzdef.convertibleTimeZone = tzSpec1.CopyOut()

	return nil
}

// setFromTimeZoneName - Sets the data fields of the specified
// TimeZoneDefinition instance base on the time zone text name.
//
func (tzDefUtil *timeZoneDefUtility) setFromTimeZoneName(
	tzdef *TimeZoneDefinition,
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

	var tzSpec TimeZoneSpecification
	var err error

	tzMech := TimeZoneMechanics{}

	tzSpec,
	err = tzMech.GetTimeZoneFromName(
		dateTime,
		timeZoneName,
		timeZoneConversionType,
		ePrefix)

	if err != nil {
		return err
	}

	tzSpec.zoneLabel = "Original Time Zone"

	tzdef.originalTimeZone = tzSpec.CopyOut()

	tzSpec.zoneLabel = "Convertible Time Zone"
	tzSpec.locationNameType.ConvertibleTimeZone()

	tzdef.convertibleTimeZone = tzSpec.CopyOut()

	return nil
}
