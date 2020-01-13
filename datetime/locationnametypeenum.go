package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var mLocationNameTypeStringToCode = map[string]LocationNameType{
	"None"                    : LocationNameType(0).None(),
	"ConvertibleAbbreviation" : LocationNameType(0).ConvertibleAbbreviation(),
	"ConvertibleTimeZoneName" : LocationNameType(0).ConvertibleTimeZoneName(),
}

var mLocationNameTypeLwrCaseStringToCode = map[string]LocationNameType{
	"none"                    : LocationNameType(0).None(),
	"convertibleabbreviation" : LocationNameType(0).ConvertibleAbbreviation(),
	"convertibletimezonename" : LocationNameType(0).ConvertibleTimeZoneName(),
}

var mLocationNameTypeCodeToString = map[LocationNameType] string {
	LocationNameType(0).None()                    : "None",
	LocationNameType(0).ConvertibleAbbreviation() : "ConvertibleAbbreviation",
	LocationNameType(0).ConvertibleTimeZoneName() : "ConvertibleTimeZoneName",
}

// LocationNameType - This type is configured as a series of
// constant integer values describing different types of Time
// Zone Location Names processed by type 'TimeZoneDefinition'.
//
// Functionally, 'LocationNameType' serves as an enumeration
// of valid Location Name types.
//
//                           Time Zone
//   Method                  Type Code
//    Name                   Constant      Description
//  ______________________________________________________________________
//   None()                     0          LocationNameType is uninitialized
//                                          and has no significant value.
//
//   ConvertibleAbbreviation    1          Tests have established that this
//                                          time zone abbreviation can be
//                                          used to convert date times to
//                                          a valid time zone.
//
//   NonConvertibleAbbreviation 2          Tests have established that this
//                                          time zone abbreviation cannot be
//                                          used to convert date times to
//                                          a valid time zone.
//
//   ConvertibleTimeZoneName    3          Tests have established that this
//                                          Location Name is a valid fully
//                                          formed time zone name which can
//                                          be used to successfully convert
//                                          date times to a valid time zone.
//
// For easy access to these enumeration values, use the global variable
// 'LocNameType'.
//
// 'LocationNameType' has been adapted to function as an enumeration
// describing the type of time zone assigned to a date time. Since
// Golang does not directly support enumerations, the 'LocationNameType'
// type has been configured to function in a manner similar to classic
// enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type LocationNameType int

var lockLocationNameType sync.Mutex

// None - Signifies that Location Name Type is uninitialized and has
// no value.
//
// This method is part of the standard enumeration.
//
func (locType LocationNameType) None() LocationNameType {
	return LocationNameType(0)
}

// ConvertibleAbbreviation - Classifies the Location Name as a Time Zone
// abbreviation which can be used to successfully convert date times to
// the designated time zone.
//
// This method is part of the standard enumeration.
//
func (locType LocationNameType) ConvertibleAbbreviation() LocationNameType {
	return LocationNameType(1)
}

// NonConvertibleAbbreviation - Classifies the Location Name as a Time
// Zone abbreviation which cannot be used to successfully convert date
// times to the designated time zone.
//
// This method is part of the standard enumeration.
//
func (locType LocationNameType) NonConvertibleAbbreviation() LocationNameType {
	return LocationNameType(2)
}

// ConvertibleTimeZoneName - Classifies the Location Name as fully formed
// time zone name which can be used to successfully convert date times
// to a designated time zone.
//
// As a practical matter, Location Names associated with this classification
// are either IANA Time Zones or the special time zone, 'Local'. The 'Local'
// is a Golang construct used to identify the time zone implemented by the
// host computer.
//
// This method is part of the standard enumeration.
//
func (locType LocationNameType) ConvertibleTimeZoneName() LocationNameType {
	return LocationNameType(3)
}


// =========================================================================

// String - Returns a string with the name of the enumeration associated
// with this instance of 'LocationNameType'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the LocationNameType value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= LocationNameType(0).ConvertibleTimeZoneName()
//	str := t.String()
//	    str is now equal to "ConvertibleTimeZoneName"
//
func (locType LocationNameType) String() string {

	lockLocationNameType.Lock()

	defer lockLocationNameType.Unlock()

	label, ok := mLocationNameTypeCodeToString[locType]

	if !ok {
		return ""
	}

	return label
}


// XParseString - Receives a string and attempts to match it with the string
// value of the supported enumeration. If successful, a new instance of
// LocationNameType is returned set to the value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
// valueString   string - A string which will be matched against the
//                        enumeration string values. If 'valueString'
//                        is equal to one of the enumeration names, this
//                        method will proceed to successful completion
//
// caseSensitive   bool - If 'true' the search for enumeration names
//                        will be case sensitive and will require an
//                        exact match. Therefore, 'convertibleabbreviation' will NOT
//                        match the enumeration name, 'ConvertibleAbbreviation'.
//
//                        If 'false' a case insensitive search is
//                        conducted for the enumeration name. In
//                        this case, 'convertibleabbreviation' will match
//                        the enumeration name 'ConvertibleAbbreviation'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
// TimeZoneType        - Upon successful completion, this method will return a new
//                       instance of TimeZoneType set to the value of the
//                       enumeration matched by the string search performed on
//                       input parameter,'valueString'.
//
// error               - If this method completes successfully, the returned error
//                       Type is set equal to 'nil'. If an error condition is encountered,
//                       this method will return an error Type which encapsulates an
//                       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
// t, err := LocationNameType(0).ParseString("ConvertibleTimeZoneName", true)
//                            OR
// t, err := LocationNameType(0).ParseString("ConvertibleTimeZoneName()", true)
//                            OR
// t, err := LocationNameType(0).ParseString("convertibletimezonename", false)
//
// For all of the cases shown above,
//  t is now equal to LocationNameType(0).ConvertibleTimeZoneName()
//
func (locType LocationNameType) XParseString(
	valueString string,
	caseSensitive bool) (LocationNameType, error) {

	lockLocationNameType.Lock()

	defer lockLocationNameType.Unlock()

	ePrefix := "LocationNameType.XParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 4 {
		return LocNameType.None(),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"Length Less than 4-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var locationNameType LocationNameType

	if caseSensitive == true {

		locationNameType, ok = mLocationNameTypeStringToCode[valueString]

		if !ok {
			return LocNameType.None(),
				errors.New(ePrefix + "Invalid LocationNameType Code!")
		}

	} else {
		// caseSensitive must be 'false'
		valueString = strings.ToLower(valueString)

		locationNameType, ok = mLocationNameTypeLwrCaseStringToCode[valueString]

		if !ok {
			return LocNameType.None(),
				errors.New(ePrefix + "Invalid LocationNameType Code!")
		}
	}

	return locationNameType, nil
}

// XValue - Returns the value of the LocationNameType instance
// as type LocationNameType.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (locType LocationNameType) XValue() LocationNameType {

	lockLocationNameType.Lock()

	defer lockLocationNameType.Unlock()

	return locType
}


// LocNameType - public global variable of
// type LocationNameType.
//
// This variable serves as an easier, short hand
// technique for accessing LocationNameType values.
//
// Usage:
//  LocNameType.None()
//  LocNameType.ConvertibleAbbreviation()
//  LocNameType.NonConvertibleAbbreviation()
//  LocNameType.ConvertibleTimeZoneName()
//
var LocNameType LocationNameType
