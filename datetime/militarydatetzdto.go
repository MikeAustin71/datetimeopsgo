package datetime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type MilitaryDateTzDto struct {
	Time                       TimeDto
	DateTime                   time.Time
	MilitaryTzLetterName       string
	MilitaryTzTextName         string
	EquivalentIanaTimeZone     TimeZoneDefDto
	UtcOffset                  string
	GeoLocationDesc            string
}

// New - Creates and returns a instance of MilitaryDateTzDto.
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
func (milDtDto MilitaryDateTzDto) New(
	t time.Time, militaryTz string) (MilitaryDateTzDto, error) {

	ePrefix := "MilitaryDateTzDto.New() "

	lMilTz := len(militaryTz)

	newMilDateDto := MilitaryDateTzDto{}

	if lMilTz == 0 {
		return newMilDateDto, errors.New(ePrefix +
			"Error: Input parameter 'militaryTz' is EMPTY!\n")
	}

	var equivalentIanaTimeZone string
	var ok bool

	if lMilTz == 1 {

		newMilDateDto.MilitaryTzLetterName = strings.ToUpper(militaryTz)

		newMilDateDto.MilitaryTzTextName , ok =
			MilitaryTzNameToTxtNameMap[newMilDateDto.MilitaryTzLetterName]

		if !ok {
			return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'militaryTz' is INVALID!\n" +
				"'militaryTz' DOES NOT map to a valid Military Time Zone.\n" +
				"militaryTz='%v'", newMilDateDto.MilitaryTzLetterName)
		}

		equivalentIanaTimeZone, ok = MilitaryTzToIanaTzMap[newMilDateDto.MilitaryTzTextName]

		if !ok {
			return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'militaryTz' is INVALID!\n" +
				"'militaryTz' DOES NOT map to a valid IANA Time Zone.\n" +
				"militaryTz='%v'", newMilDateDto.MilitaryTzTextName)
		}


	} else {
		// lMilTz > 1
		temp1 := militaryTz[:1]
		temp2 := militaryTz[1:]

		temp1 = strings.ToUpper(temp1)
		temp2 = strings.ToLower(temp2)

		newMilDateDto.MilitaryTzTextName = temp1 + temp2

		equivalentIanaTimeZone, ok = MilitaryTzToIanaTzMap[newMilDateDto.MilitaryTzTextName]

		if !ok {
			return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
				"Error: Input Parameter Value 'militaryTz' is INVALID!\n" +
				"'militaryTz' DOES NOT map to a valid IANA Time Zone.\n" +
				"militaryTz='%v'", newMilDateDto.MilitaryTzTextName)
		}

		newMilDateDto.MilitaryTzLetterName = newMilDateDto.MilitaryTzTextName[:1]
	}

	tzLoc, err := time.LoadLocation(equivalentIanaTimeZone)

	if err != nil {
		return newMilDateDto, fmt.Errorf(ePrefix +
			"Error returned by time.LoadLocation(equivalentIanaTimeZone)\n" +
			"equivalentIanaTimeZone='%v'\n" +
			"Error='%v'\n", equivalentIanaTimeZone, err.Error())
	}

	newMilDateDto.DateTime = t.In(tzLoc)

	newMilDateDto.EquivalentIanaTimeZone, err = TimeZoneDefDto{}.New(newMilDateDto.DateTime)

	if err != nil {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error returned by TimeZoneDefDto{}.New(newMilDateDto.DateTime)\n" +
			"newMilDateDto.DateTime='%v'\n" +
			"Error='%v'\n",
			newMilDateDto.DateTime.Format(FmtDateTimeTzNanoYMDDow), err.Error())
	}

	newMilDateDto.Time, err = TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)

	if err != nil {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error returned by TimeDto{}.NewFromDateTime(newMilDateDto.DateTime)\n" +
			"newMilDateDto.DateTime='%v'\n" +
			"Error='%v'\n",
			newMilDateDto.DateTime.Format(FmtDateTimeTzNanoYMDDow), err.Error())
	}

	newMilDateDto.GeoLocationDesc, ok = MilitaryTzLocationMap[newMilDateDto.MilitaryTzTextName]

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map geographical location description.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}

	newMilDateDto.UtcOffset, ok = MilitaryTzToUTCMap[newMilDateDto.MilitaryTzTextName]

	if !ok {
		return MilitaryDateTzDto{}, fmt.Errorf(ePrefix +
			"Error: Military Time Zone Text Name is INVALID!\n" +
			"Could NOT map UTC Offset.\n" +
			"newMilDateDto.MilitaryTzTextName='%v'\n", newMilDateDto.MilitaryTzTextName)
	}


	return newMilDateDto, nil
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