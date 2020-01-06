package datetime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// MilitaryDateTzDto - This type is used to define date times associated with
// Military Time Zones.
//
// Military time zones are commonly used in aviation as well as at sea.
// They are also known as nautical or maritime time zones.
//
// For easy access to Military Time Zones and their IANA equivalents see the
// enumeration type, 'TimeZones'.
//
// Type 'TimeZones', is located in source file:
//
//    Source Repository: 'https://github.com/MikeAustin71/datetimeopsgo.git'
//    Source Code File:  MikeAustin71\datetimeopsgo\datetime\timezonedata.go
//
// For background information on Military Time Zones reference:
//
//     https://en.wikipedia.org/wiki/List_of_military_time_zones
//     http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//     https://www.timeanddate.com/time/zones/military
//     https://www.timeanddate.com/worldclock/timezone/alpha
//     https://www.timeanddate.com/time/map/
//
// The following list details Military Time Zones, 'A' to 'Z'.  The 'J' (Juliet)
// Time Zone is occasionally used to refer to the observer's local time. Note that
// Time Zone 'J' (Juliet) is not listed below.
///
//    Time Zone       Time Zone        Equivalent IANA          UTC
//   Abbreviation       Name              Time Zone            Offset
//   ------------     --------          ---------------        ------
//
//       A        Alpha Time Zone         Etc/GMT-1            UTC +1
//       B        Bravo Time Zone         Etc/GMT-2            UTC +2
//       C        Charlie Time Zone       Etc/GMT-3            UTC +3
//       D        Delta Time Zone         Etc/GMT-4            UTC +4
//       E        Echo Time Zone          Etc/GMT-5            UTC +5
//       F        Foxtrot Time Zone       Etc/GMT-6            UTC +6
//       G        Golf Time Zone          Etc/GMT-7            UTC +7
//       H        Hotel Time Zone         Etc/GMT-8            UTC +8
//       I        India Time Zone         Etc/GMT-9            UTC +9
//       K        Kilo Time Zone          Etc/GMT-10           UTC +10
//       L        Lima Time Zone          Etc/GMT-11           UTC +11
//       M        Mike Time Zone          Etc/GMT-12           UTC +12
//       N        November Time Zone      Etc/GMT+1            UTC -1
//       O        Oscar Time Zone         Etc/GMT+2            UTC -2
//       P        Papa Time Zone          Etc/GMT+3            UTC -3
//       Q        Quebec Time Zone        Etc/GMT+4            UTC -4
//       R        Romeo Time Zone         Etc/GMT+5            UTC -5
//       S        Sierra Time Zone        Etc/GMT+6            UTC -6
//       T        Tango Time Zone         Etc/GMT+7            UTC -7
//       U        Uniform Time Zone       Etc/GMT+8            UTC -8
//       V        Victor Time Zone        Etc/GMT+9            UTC -9
//       W        Whiskey Time Zone       Etc/GMT+10           UTC -10
//       X        X-ray Time Zone         Etc/GMT+11           UTC -11
//       Y        Yankee Time Zone        Etc/GMT+12           UTC -12
//       Z        Zulu Time Zone          UTC                  UTC +0
//
type MilitaryDateTzDto struct {
	Description                string          // Unused, available for classification, labeling
	                                           //   or description.
	Time                       TimeDto         // Associated Time Components
	DateTime                   time.Time       // DateTime value for this MilitaryDateTzDto Type
	MilitaryTzLetterName       string          // A single letter/character designating the military
	                                           //   time zone
	MilitaryTzTextName         string          // A Text String containing the Military Time Zone Name
	EquivalentIanaTimeZone     TimeZoneDefDto  // The IANA time zone which equates to this
	                                           //   Military Time Zone
	UtcOffset                  string          // UTC Offset for this Military Time Zone
	GeoLocationDesc            string          // Military Time Zone Geographic Location
}

// CopyIn - Receives an incoming MilitaryDateTzDto and copies those data
// fields into the current MilitaryDateTzDto instance.
//
// When completed, the current MilitaryDateTzDto will be equal in all
// respects to the incoming MilitaryDateTzDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  milDtDto2 MilitaryDateTzDto - A MilitaryDateTzDto instance.
//                    This data will be copied into the data
//                    fields of the current MilitaryDateTzDto
//                    instance.
//
//      A MilitaryDateTzDto struct is defined as follows:
//
//      type MilitaryDateTzDto struct {
//        Description                string          // Unused, available for classification, labeling
//                                                   //   or description.
//        Time                       TimeDto         // Associated Time Components
//        DateTime                   time.Time       // DateTime value for this
//                                                   //   MilitaryDateTzDto Type
//        MilitaryTzLetterName       string          // A single letter/character
//                                                   //   designating the military
//                                                   //   time zone
//        MilitaryTzTextName         string          // A Text String containing the
//                                                   //   Military Time Zone Name
//        EquivalentIanaTimeZone     TimeZoneDefDto  // The IANA time zone which equates
//                                                   //   to the Military Time Zone
//        UtcOffset                  string          // UTC Offset for this Military
//                                                   //   Time Zone
//        GeoLocationDesc            string          // Military Time Zone Geographic
//                                                   //   Location
//      }
//
// ------------------------------------------------------------------------
//
// Return Values
//
//   None
//
// ------------------------------------------------------------------------
//
// Usage
//
//  milDtzDto := MilitaryDateTzDto{}
//  milDtzDto.CopyIn(milDtzDto2)
//
//  Note: milDtzDto and milDtzDto2 are now equivalent.
//
func (milDtDto *MilitaryDateTzDto) CopyIn(milDtDto2 MilitaryDateTzDto) {

	milDtDto.Empty()
	milDtDto.Description = milDtDto2.Description
	milDtDto.Time = milDtDto2.Time.CopyOut()

	if !milDtDto2.DateTime.IsZero() {
		milDtDto.DateTime = milDtDto2.DateTime
		milDtDto.EquivalentIanaTimeZone = milDtDto2.EquivalentIanaTimeZone.CopyOut()
	} else {
		milDtDto.EquivalentIanaTimeZone = TimeZoneDefDto{}
		milDtDto.DateTime = time.Time{}
	}

	milDtDto.UtcOffset = milDtDto2.UtcOffset
	milDtDto.GeoLocationDesc = milDtDto2.GeoLocationDesc
}

// copyOut - Returns a deep copy of the current 'MilitaryDateTzDto'
// instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  None
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  MilitaryDateTzDto - This method returns a new, valid, fully populated DateTzDto
//                      which is a deep copy of the current DateTzDto instance.
//                      A DateTzDto structure is defined as follows:
//
//      type MilitaryDateTzDto struct {
//        Description                string          // Unused, available for classification, labeling
//                                                   //   or description.
//        Time                       TimeDto         // Associated Time Components
//        DateTime                   time.Time       // DateTime value for this
//                                                   //   MilitaryDateTzDto Type
//        MilitaryTzLetterName       string          // A single letter/character
//                                                   //   designating the military
//                                                   //   time zone
//        MilitaryTzTextName         string          // A Text String containing the
//                                                   //   Military Time Zone Name
//        EquivalentIanaTimeZone     TimeZoneDefDto  // The IANA time zone which equates
//                                                   //   to the Military Time Zone
//        UtcOffset                  string          // UTC Offset for this Military
//                                                   //   Time Zone
//        GeoLocationDesc            string          // Military Time Zone Geographic
//                                                   //   Location
//      }
//
// ------------------------------------------------------------------------
//
// Usage
//
//  milDtz1 := MilitaryDateTzDto{}
//  ... initialize to some value
//
//  milDtz2 := milDtz1.copyOut()
//
//  Note: milDtz1 and milDtz2 are now equivalent.
//
func (milDtDto *MilitaryDateTzDto) CopyOut() MilitaryDateTzDto {

	milDtz2 := MilitaryDateTzDto{}

	milDtz2.Description = milDtDto.Description
	milDtz2.Time = milDtDto.Time.CopyOut()
	milDtz2.DateTime = milDtDto.DateTime
	milDtz2.MilitaryTzTextName = milDtDto.MilitaryTzTextName
	milDtz2.MilitaryTzLetterName = milDtDto.MilitaryTzLetterName
	milDtz2.EquivalentIanaTimeZone = milDtDto.EquivalentIanaTimeZone.CopyOut()
	milDtz2.UtcOffset = milDtDto.UtcOffset
	milDtz2.GeoLocationDesc = milDtDto.GeoLocationDesc

	return milDtz2
}


// Empty - sets all values of the current MilitaryDateTzDto
// instance to their uninitialized or zero state.
func (milDtDto *MilitaryDateTzDto) Empty() {

	milDtDto.Description = ""
	milDtDto.Time.Empty()
	milDtDto.DateTime = time.Time{}
	milDtDto.MilitaryTzLetterName = ""
	milDtDto.MilitaryTzTextName = ""
	milDtDto.EquivalentIanaTimeZone = TimeZoneDefDto{}
	milDtDto.UtcOffset = ""
	milDtDto.GeoLocationDesc = ""

	return
}

// IsValid - Analyzes the current MilitaryDateTzDto instance to determine
// if the data fields are valid and correctly initialized. If the current
// 'MilitaryDateTzDto' is determined to by INVALID, an error, populated
// with an appropriate error message, will be returned.
//
// If the current MilitaryDateTzDto instance is VALID, this method returns
// nil.
//
func (milDtDto *MilitaryDateTzDto) IsValid() error {

	ePrefix := "DateTzDto.IsValid() "

	if milDtDto.DateTime.IsZero() {
		return errors.New(ePrefix + "Error: milDtDto.DateTime is ZERO!")
	}

	if milDtDto.EquivalentIanaTimeZone.IsEmpty() {
		return errors.New(ePrefix +
			"\nError: milDtDto.EquivalentIanaTimeZone is EMPTY!")
	}

	if err := milDtDto.Time.IsValid(); err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: milDtDto.Time is INVALID. Error='%v'", err.Error())
	}

	return nil
}

// GetCompactDateTimeGroup - Outputs date time string formatted for
// standard U.S.A. Military date time also referred to as the Military
// Date Time Group (DTG). This form of the Date Time Group is configured
// as the 'Compact' Date Time Group. This means there are no spaces between
// the date time elements.
//
// The time value is taken from the current 'MilitaryDateTzDto' object.
//
// Reference:
//    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//    https://www.timeanddate.com/time/zones/z
//    http://blog.refactortactical.com/blog/military-date-time-group/
//
// Military 2-digit year format or "Date Time Group" is traditionally
// formatted as DDHHMM(Z)MONYY, where 'Z' is the Military Time Zone.
//
// EXAMPLES:
//
//    "011815ZJAN11" = 01/01/2011 18:15 +0000 Zulu
//
//     630pm on January 6th, 2012 in Fayetteville NC would read '061830RJAN12'
//
func (milDtDto *MilitaryDateTzDto) GetCompactDateTimeGroup() (string, error) {

	if milDtDto.DateTime.Equal(time.Time{}) && milDtDto.Time.IsEmpty() {
		return "", errors.New("MilitaryDateTzDto.GetCompactDateTimeGroup() Error:\n" +
			"MilitaryDateTzDto.DateTime is ZERO or Uninitialized!")
	}

	fmtDateTime := milDtDto.DateTime.Format("021504" + milDtDto.MilitaryTzLetterName + "Jan06")

	fmtDateTime = strings.ToUpper(fmtDateTime)

	return fmtDateTime, nil
}

// GetOpenDateTimeGroup - Outputs date time string formatted for
// standard U.S.A. Military date time also referred to as the Military
// Date Time Group (DTG). This form of the Date Time Group is configured
// as the 'Open', easy to read, Date Time Group. This means that spaces
// are inserted between the date time elements.
//
// The time value is taken from the current 'MilitaryDateTzDto' object.
//
// Reference:
//    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//    https://www.timeanddate.com/time/zones/z
//    http://blog.refactortactical.com/blog/military-date-time-group/
//
// Military 2-digit year format or "Date Time Group" is traditionally
// formatted as DDHHMM(Z)MONYY, where 'Z' is the Military Time Zone.
//
// EXAMPLES:
//
//    "01 1815Z JAN 11" = 01/01/2011 18:15 +0000 Zulu
//
//     630pm on January 6th, 2012 in Fayetteville NC would read '06 1830R JAN 12'
//
func (milDtDto *MilitaryDateTzDto) GetOpenDateTimeGroup() (string, error) {

	if milDtDto.DateTime.Equal(time.Time{}) && milDtDto.Time.IsEmpty() {
		return "", errors.New("MilitaryDateTzDto.GetOpenDateTimeGroup() Error:\n" +
			"MilitaryDateTzDto.DateTime is ZERO or Uninitialized!")
	}

	fmtDateTime := milDtDto.DateTime.Format("02 1504" + milDtDto.MilitaryTzLetterName + " Jan 06")

	fmtDateTime = strings.ToUpper(fmtDateTime)

	return fmtDateTime, nil
}

// New - Creates and returns a new instance of MilitaryDateTzDto.
// The values for the new returned instance are calculated from
// the two input parameters 't' and 'militaryTz'.
//
// ----------------------------------------------------------
//
// Input Parameters
//
//   t        time.Time - This time value is used to initialize the new
//                        instance of MilitaryDateTzDto returned by this
//                        method.
//
//   militaryTz  string - This string designates the Military Time Zone
//                        used to initialized the new instance of
//                        MilitaryDateTzDto returned by this method.
//
//                        This input parameter is also used to convert
//                        input parameter 't' to the time value in the
//                        the designated Military Time Zone.
//
//                        'militaryTz' may be submitted as either a single
//                        character Military Time Zone Letter Name or as
//                        the Military Time Zone Text Name. For example,
//                        Military Time Zone 'Lima' may be submitted as the
//                        letter, 'L', or as the text string 'Lima'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  MilitaryDateTzDto - An instance of type 'MilitaryDateTzDto' initialized
//                     from input parameters 't' and 'militaryTz'.
//
//    type MilitaryDateTzDto struct {
//      Description                string          // Unused, available for classification, labeling
//                                                 //   or description.
//      Time                       TimeDto         // Associated Time Components
//      DateTime                   time.Time       // DateTime value for this MilitaryDateTzDto Type
//      MilitaryTzLetterName       string          // A single letter/character designating the
//                                                 //   military time zone
//      MilitaryTzTextName         string          // A Text String containing the Military Time
//                                                 //   Zone Name
//      EquivalentIanaTimeZone     TimeZoneDefDto  // The IANA time zone which equates to this
//                                                 //   Military Time Zone
//      UtcOffset                  string          // UTC Offset for this Military Time Zone
//      GeoLocationDesc            string          // Military Time Zone Geographic Location
//    }
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//         encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  milDtTzDto, err := MilitaryDateTzDto{}.New(t, militaryTz)
//
func (milDtDto MilitaryDateTzDto) New(
	t time.Time, militaryTz string) (MilitaryDateTzDto, error) {

	ePrefix := "MilitaryDateTzDto.New() "

	newMilDateDto := MilitaryDateTzDto{}

	if t.IsZero() {
		return newMilDateDto,
			errors.New(ePrefix +
				"\nError: Input parameter t tim.Time is ZERO!\n")
	}


	if len(militaryTz) == 0 {
		return newMilDateDto, errors.New(ePrefix +
			"Error: Input parameter 'militaryTz' is EMPTY!\n")
	}

	var err error

	newMilDateDto.MilitaryTzLetterName,
	newMilDateDto.MilitaryTzTextName,
	newMilDateDto.EquivalentIanaTimeZone,
	err = MilitaryDateTzDto{}.parseMilitaryTzNameAndLetter(militaryTz)

	if err != nil {
		return MilitaryDateTzDto{},
			fmt.Errorf(ePrefix +
				"\nError: '%v'\n", err.Error())
	}

	newMilDateDto.DateTime = t.In(newMilDateDto.EquivalentIanaTimeZone.GetLocationPtr())

	newMilDateDto.Time, err = TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)

	if err != nil {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error returned by TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)\n" +
			"newMilDateDto.DateTime='%v'\n" +
			"Error='%v'\n",
			newMilDateDto.DateTime.Format(FmtDateTimeTzNanoYMDDow), err.Error())
	}

	var ok bool

	milTzDat := MilitaryTimeZoneData{}

	newMilDateDto.GeoLocationDesc, ok = milTzDat.MilitaryTzToLocation(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map geographical location description.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}

	newMilDateDto.UtcOffset, ok = milTzDat.MilitaryTzToUtc(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map UTC Offset.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}


	return newMilDateDto, nil
}


// NewFromDateTzDto - Creates and returns a new instance of 'MilitaryDateTzDto'.
// This new instance is initialized from the 'DateTzDto' instance passed as an
// input parameter.
//
// ----------------------------------------------------------
//
// Input Parameters
//
//  DateTzDto - A valid, fully populated 'DateTzDto' instance.
//
//              A DateTzDto structure is defined as follows:
//
//      type DateTzDto struct {
//           Description  string         // Unused, available for classification,
//                                       //  labeling or description
//           Time         TimeDto        // Associated Time Components
//           DateTime     time.Time      // DateTime value for this DateTzDto Type
//           DateTimeFmt  string         // Date Time Format String.
//                                       //  Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//           TimeZone     TimeZoneDefDto // Contains a detailed description of the Time Zone
//                                       //  and Time Zone Location
//                                       // associated with this date time.
//      }
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  MilitaryDateTzDto - An instance of type 'MilitaryDateTzDto' initialized
//                     from input parameters 't' and 'militaryTz'.
//
//    type MilitaryDateTzDto struct {
//      Description                string          // Unused, available for classification, labeling
//                                                 //   or description.
//      Time                       TimeDto         // Associated Time Components
//      DateTime                   time.Time       // DateTime value for this MilitaryDateTzDto Type
//      MilitaryTzLetterName       string          // A single letter/character designating the
//                                                 //   military time zone
//      MilitaryTzTextName         string          // A Text String containing the Military Time
//                                                 //   Zone Name
//      EquivalentIanaTimeZone     TimeZoneDefDto  // The IANA time zone which equates to this
//                                                 //   Military Time Zone
//      UtcOffset                  string          // UTC Offset for this Military Time Zone
//      GeoLocationDesc            string          // Military Time Zone Geographic Location
//    }
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//         encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  milDtTzDto, err := MilitaryDateTzDto{}.NewFromDateTzDto(dtzDto)
//
func (milDtDto MilitaryDateTzDto) NewFromDateTzDto(
	dtzDto DateTzDto) (MilitaryDateTzDto, error) {

	ePrefix := "MilitaryDateTzDto.NewFromDateTzDto() "
	newMilDateDto := MilitaryDateTzDto{}

	err := dtzDto.IsValid()

	if err != nil {
		return newMilDateDto,
			fmt.Errorf(ePrefix +
				"\nInput Parameter 'dtzDto' is INVALID!\n" +
				"Error='%v'\n", err.Error())
	}

	// FmtDateTimeTzSec = "01/02/2006 15:04:05 -0700 MST"
	dateTimeArray := strings.Split(dtzDto.GetDateTimeValue().Format(FmtDateTimeTzSec), " ")

	if len(dateTimeArray) != 4 {
		return newMilDateDto,
			fmt.Errorf(ePrefix +
				"\nError: dtzDto.DateTime resolves to an INVALID date time string!\n" +
				"dtzDto.DateTime='%v'\n" +
				"dtzDto.DateTime Array Length='%v'\n",
				dtzDto.GetDateTimeValue().Format(FmtDateTimeTzSec),
				len(dateTimeArray))
	}

	milTzDat := MilitaryTimeZoneData{}

	militaryTz, ok := milTzDat.UtcOffsetToMilitaryTimeZone(dateTimeArray[2])

	if !ok {
		return newMilDateDto,
			fmt.Errorf(ePrefix +
				"\nError: Could not locate Military Time Zone from UTC offset.\n" +
				"UTC Offset='%v'\n", dateTimeArray[2])
	}

	newMilDateDto.MilitaryTzLetterName,
		newMilDateDto.MilitaryTzTextName,
		newMilDateDto.EquivalentIanaTimeZone,
		err = MilitaryDateTzDto{}.parseMilitaryTzNameAndLetter(militaryTz)

	newMilDateDto.DateTime =
		dtzDto.GetDateTimeValue().In(newMilDateDto.EquivalentIanaTimeZone.GetLocationPtr())

	newMilDateDto.Time, err = TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)

	if err != nil {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error returned by TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)\n" +
			"newMilDateDto.DateTime='%v'\n" +
			"Error='%v'\n",
			newMilDateDto.DateTime.Format(FmtDateTimeTzNanoYMDDow), err.Error())
	}

	milTzDat = MilitaryTimeZoneData{}

	newMilDateDto.GeoLocationDesc, ok = milTzDat.MilitaryTzToLocation(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map geographical location description.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}


	newMilDateDto.UtcOffset, ok = milTzDat.MilitaryTzToUtc(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map UTC Offset.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}

	return newMilDateDto, nil
}

// NewNow() - Calculates the current date time and converts that time to
// military time zone contained in parameter 'militaryTz'. The resulting
// military date time value is returned in a new instance of type,
// 'MilitaryDateTzDto'.
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   militaryTz  string - This string designates the Military Time Zone
//                        used to initialized the new instance of
//                        MilitaryDateTzDto returned by this method.
//
//                        This input parameter is also used to convert
//                        input parameter 't' to the time value in the
//                        the designated Military Time Zone.
//
//                        'militaryTz' may be submitted as either a single
//                        character Military Time Zone Letter Name or as
//                        the Military Time Zone Text Name. For example,
//                        Military Time Zone 'Lima' may be submitted as the
//                        letter, 'L', or as the text string 'Lima'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  MilitaryDateTzDto - An instance of type 'MilitaryDateTzDto' initialized
//                     from input parameters 't' and 'militaryTz'.
//
//    type MilitaryDateTzDto struct {
//      Description                string          // Unused, available for classification, labeling
//                                                 //   or description.
//      Time                       TimeDto         // Associated Time Components
//      DateTime                   time.Time       // DateTime value for this MilitaryDateTzDto Type
//      MilitaryTzLetterName       string          // A single letter/character designating the
//                                                 //   military time zone
//      MilitaryTzTextName         string          // A Text String containing the Military Time
//                                                 //   Zone Name
//      EquivalentIanaTimeZone     TimeZoneDefDto  // The IANA time zone which equates to this
//                                                 //   Military Time Zone
//      UtcOffset                  string          // UTC Offset for this Military Time Zone
//      GeoLocationDesc            string          // Military Time Zone Geographic Location
//    }
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//         encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  milDtTzDto, err := MilitaryDateTzDto{}.NewNow(militaryTz)
//
func (milDtDto MilitaryDateTzDto) NewNow(
	militaryTz string) (MilitaryDateTzDto, error) {

	ePrefix := "MilitaryDateTzDto.NewNow() "

	newMilDateDto := MilitaryDateTzDto{}

	if len(militaryTz) == 0 {
		return newMilDateDto,
			errors.New(ePrefix +
				"\nError: Input parameter 'militaryTz' is EMPTY!\n")
	}

	var err error

	newMilDateDto.MilitaryTzLetterName,
	newMilDateDto.MilitaryTzLetterName,
	newMilDateDto.EquivalentIanaTimeZone,
	err = MilitaryDateTzDto{}.parseMilitaryTzNameAndLetter(militaryTz)

	if err != nil {
		return MilitaryDateTzDto{},
			fmt.Errorf(ePrefix +
				"\nInvalid Military Time Zone!\n" +
				"'%v'\n", err.Error())
	}

	t := time.Now().UTC()
	newMilDateDto.DateTime = t.In(newMilDateDto.EquivalentIanaTimeZone.GetLocationPtr())

	newMilDateDto.Time, err = TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)

	if err != nil {
		return MilitaryDateTzDto{},
			fmt.Errorf(ePrefix +
				"\nError returned by TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)\n" +
				"newMilDateDto.DateTime='%v'\n" +
				"Error='%v'\n", newMilDateDto.DateTime.Format( "2006-01-02 15:04:05.000000000 -0700"),
				err.Error())
	}
	var ok bool

	milTzDat := MilitaryTimeZoneData{}

	newMilDateDto.GeoLocationDesc, ok = milTzDat.MilitaryTzToLocation(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map geographical location description.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}

	newMilDateDto.UtcOffset, ok = milTzDat.MilitaryTzToUtc(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map UTC Offset.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}

	return newMilDateDto, nil
}

// NewNowZulu() - Calculates the current date time and converts that
// time to military time zone Zulu (UTC offset +0000). The resulting
// military date time value is returned in a new instance of type,
// 'MilitaryDateTzDto'.
// ------------------------------------------------------------------------
//
// Input Parameters
//
// None
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  MilitaryDateTzDto - An instance of type 'MilitaryDateTzDto' initialized
//                      with the current time in Time Zone 'Zulu' (UTC Offset +0000).
//
//    type MilitaryDateTzDto struct {
//      Description                string          // Unused, available for classification, labeling
//                                                 //   or description.
//      Time                       TimeDto         // Associated Time Components
//      DateTime                   time.Time       // DateTime value for this MilitaryDateTzDto Type
//      MilitaryTzLetterName       string          // A single letter/character designating the
//                                                 //   military time zone
//      MilitaryTzTextName         string          // A Text String containing the Military Time
//                                                 //   Zone Name
//      EquivalentIanaTimeZone     TimeZoneDefDto  // The IANA time zone which equates to this
//                                                 //   Military Time Zone
//      UtcOffset                  string          // UTC Offset for this Military Time Zone
//      GeoLocationDesc            string          // Military Time Zone Geographic Location
//    }
//
//  error - If successful the returned error Type is set equal to 'nil'. If errors are
//         encountered this error Type will encapsulate an error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  milDtTzDto, err := MilitaryDateTzDto{}.NewNowZulu(militaryTz)
//
func (milDtDto MilitaryDateTzDto) NewNowZulu() (MilitaryDateTzDto, error) {

	ePrefix := "MilitaryDateTzDto.NewNow() "

	newMilDateDto := MilitaryDateTzDto{}

	var err error

	newMilDateDto.MilitaryTzLetterName,
	newMilDateDto.MilitaryTzLetterName,
	newMilDateDto.EquivalentIanaTimeZone,
	err = MilitaryDateTzDto{}.parseMilitaryTzNameAndLetter("Zulu")

	if err != nil {
		return MilitaryDateTzDto{},
			fmt.Errorf(ePrefix +
				"\nInvalid Military Time Zone, 'Zulu' !\n" +
				"'%v'\n", err.Error())
	}

	newMilDateDto.DateTime = time.Now().UTC()

	newMilDateDto.Time, err = TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)

	if err != nil {
		return MilitaryDateTzDto{},
			fmt.Errorf(ePrefix +
				"\nError returned by TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)\n" +
				"newMilDateDto.DateTime='%v'\n" +
				"Error='%v'\n", newMilDateDto.DateTime.Format( "2006-01-02 15:04:05.000000000 -0700"),
				err.Error())
	}
	var ok bool

	milTzDat := MilitaryTimeZoneData{}

	newMilDateDto.GeoLocationDesc, ok = milTzDat.MilitaryTzToLocation(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map geographical location description.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}

	newMilDateDto.UtcOffset, ok = milTzDat.MilitaryTzToUtc(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map UTC Offset.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}

	return newMilDateDto, nil
}

// SetFromTimeTz - Resets the data field values for the
// current 'MilitaryDateTzDto' instance with values created
// from the input parameters, 'dateTime' and 'militaryTz'.
//
func (milDtDto *MilitaryDateTzDto) SetFromTimeTz(
	dateTime time.Time,
	militaryTz string) error {

	ePrefix := "MilitaryDateTzDto.SetFromTimeTz() "

	if dateTime.IsZero() {
		return errors.New( ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO and INVALID!\n")
	}

	if len(militaryTz) == 0 {
		return errors.New( ePrefix +
			"\nError: Input parameter 'militaryTz' is an empty string!\n")
	}

	newMilDateDto := MilitaryDateTzDto{}

	var err error

	newMilDateDto.MilitaryTzLetterName,
	newMilDateDto.MilitaryTzTextName,
	newMilDateDto.EquivalentIanaTimeZone,
	err = MilitaryDateTzDto{}.parseMilitaryTzNameAndLetter(militaryTz)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError: '%v'\n", err.Error())
	}

	newMilDateDto.DateTime = dateTime.In(newMilDateDto.EquivalentIanaTimeZone.GetLocationPtr())

	newMilDateDto.Time, err = TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)\n" +
			"newMilDateDto.DateTime='%v'\n" +
			"Error='%v'\n",
			newMilDateDto.DateTime.Format(FmtDateTimeTzNanoYMDDow), err.Error())
	}

	var ok bool

	milTzDat := MilitaryTimeZoneData{}

	newMilDateDto.GeoLocationDesc, ok = milTzDat.MilitaryTzToLocation(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map geographical location description.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}

	newMilDateDto.UtcOffset, ok = milTzDat.MilitaryTzToUtc(newMilDateDto.MilitaryTzTextName)

	if !ok {
		return fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map UTC Offset.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}

	milDtDto.CopyIn(newMilDateDto)

	return nil
}

// parseMilitaryTzNameAndLetter - Parses a text string which
// contains either a single letter military time zone designation
// or a multi-character time zone text name.
//
// If successful, two populated strings are returned. The first
// is the valid Military Time Zone Letter designation. The second
// returned strings contains the text name of the Military Time
// Zone.
//
// If an error is encountered, the return value, 'err' is populated
// with an appropriate error message. Otherwise, 'err' is set
// equal to 'nil' signaling no error was encountered.
//
func (milDtDto MilitaryDateTzDto) parseMilitaryTzNameAndLetter(
	rawTz string) (milTzLetter, milTzName string, equivalentIanaTimeZone TimeZoneDefDto, err error) {

	milTzLetter = ""
	milTzName = ""
	equivalentIanaTimeZone = TimeZoneDefDto{}
	err = nil

		ePrefix := "MilitaryDateTzDto.parseMilitaryTzNameAndLetter() "

		lMilTz := len(rawTz)

		if lMilTz == 0 {
			err = errors.New(ePrefix +
				"Error: Input Parameter 'rawTz' is EMPTY!\n")
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
		}

	var ok bool
	var equivalentTzStr string
	milTzData := MilitaryTimeZoneData{}

	if lMilTz == 1 {

		milTzLetter = strings.ToUpper(rawTz)

		milTzName , ok =
			milTzData.MilTzLetterToTextName(milTzLetter)

		if !ok {
			err = fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'militaryTz' is INVALID!\n" +
				"'militaryTz' DOES NOT map to a valid Military Time Zone.\n" +
				"militaryTz='%v'", milTzLetter)
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
		}

		equivalentTzStr, ok = milTzData.MilitaryTzToIanaTz(milTzName)

		if !ok {
			err = fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'rawTz' is INVALID!\n" +
				"'rawTz' DOES NOT map to a valid IANA Time Zone.\n" +
				"rawTz='%v'", milTzName)
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
		}


	} else {
		// lMilTz > 1
		temp1 := rawTz[:1]
		temp2 := rawTz[1:]

		temp1 = strings.ToUpper(temp1)
		temp2 = strings.ToLower(temp2)

		milTzLetter = temp1
		milTzName = temp1 + temp2

		equivalentTzStr, ok = milTzData.MilitaryTzToIanaTz(milTzName)

		if !ok {
			err = fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'rawTz' is INVALID!\n" +
				"'rawTz' DOES NOT map to a valid IANA Time Zone.\n" +
				"Military Time Zone Letter='%v'\n" +
				"Military Time Zone Text Name='%v'", milTzLetter ,milTzName)
			milTzLetter = ""
			milTzName = ""
			return milTzLetter, milTzName, equivalentIanaTimeZone, err
		}
	}

	var err2 error

	equivalentIanaTimeZone, err2 = TimeZoneDefDto{}.NewFromTimeZoneName(
		time.Now(),
		equivalentTzStr,
		TzConvertType.Relative())

	if err2 != nil {
		err = fmt.Errorf(ePrefix +
			"\nError returned by TimeZoneDefDto{}.NewFromTimeZoneName(equivalentTzStr)\n" +
			"equivalentTzStr='%v'\n" +
			"Error='%v'\n", equivalentTzStr, err2.Error())

		milTzLetter = ""
		milTzName = ""
		equivalentIanaTimeZone = TimeZoneDefDto{}

		return milTzLetter, milTzName, equivalentIanaTimeZone, err
	}

	err = nil

	return milTzLetter, milTzName, equivalentIanaTimeZone, err
}