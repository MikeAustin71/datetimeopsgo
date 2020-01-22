package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var mTimeZoneTypeStringToCode = map[string]TimeZoneType{
	"None":     TimeZoneType(0),
	"Iana":     TimeZoneType(1),
	"Local":    TimeZoneType(2),
	"Military": TimeZoneType(3),
}

var mTimeZoneTypeLwrCaseStringToCode = map[string]TimeZoneType{
	"none":     TimeZoneType(0),
	"iana":     TimeZoneType(1),
	"local":    TimeZoneType(2),
	"military": TimeZoneType(3),
}

var mTimeZoneTypeCodeToString = map[TimeZoneType]string{
	TimeZoneType(0) : "None",
	TimeZoneType(1) : "Iana",
	TimeZoneType(2) : "Local",
	TimeZoneType(3) : "Military",
}

// TimeZoneType - This type is configured as a series of
// constant integer values describing the valid types
// of time zones processed by type 'TimeZoneDefinition'.
//
// Functionally, 'TimeZoneType' serves as an enumeration
// of valid time zone types.
//
//                      Time Zone
//   Method             Type Code
//    Name              Constant        Description
//  ______________________________________________________________________
//  None()                  0           Time Zone type is uninitialized
//                                        and has no significant value.
//
//  Iana()                  1           Tests have established that the
//                                        Time Zone is identified in the
//                                        IANA Time Zone database.
//
//  Local()                 2           Tests have established that the
//                                        Time Zone is 'Local'. This is
//                                        the time zone currently configured
//                                        for the host computer.
//
//  Military()              3           Tests have established that the
//                                        Time Zone is a valid, standard
//                                        Military Time Zone.
//
// For easy access to these enumeration values, use the global variable
// 'TzType'.
//
// 'TimeZoneType' has been adapted to function as an enumeration
// describing the type of time zone assigned to a date time. Since
// Golang does not directly support enumerations, the 'TimeZoneType'
// type has been configured to function in a manner similar to classic
// enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type TimeZoneType int

var lockTimeZoneType sync.Mutex

// None - TimeZoneType is uninitialized and has no value.
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) None() TimeZoneType {

	lockTimeZoneType.Lock()

	defer lockTimeZoneType.Unlock()

	return TimeZoneType(0)
}

// Iana - Classifies the time zone as part of the IANA Time Zone database
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) Iana() TimeZoneType {

	lockTimeZoneType.Lock()

	defer lockTimeZoneType.Unlock()

	return TimeZoneType(1)
}

// Local - The 'Local' time zone is construct of the Go programming language.
// It signals that the time zone currently configured on the host computer
// has been selected.
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) Local() TimeZoneType {

	lockTimeZoneType.Lock()

	defer lockTimeZoneType.Unlock()

	return TimeZoneType(2)
}

// Military - Classifies the time zone as a standard military time zone.
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) Military() TimeZoneType {

	lockTimeZoneType.Lock()

	defer lockTimeZoneType.Unlock()

	return TimeZoneType(3)
}

// =========================================================================

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneType'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneType value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= TimeZoneType(0).Military()
// str := t.String()
//     str is now equal to "Military"
//
func (tzType TimeZoneType) String() string {

	lockTimeZoneType.Lock()

	defer lockTimeZoneType.Unlock()

	label, ok := mTimeZoneTypeCodeToString[tzType]

	if !ok {
		return ""
	}

	return label
}

// XParseString - Receives a string and attempts to match it with the string
// value of the supported enumeration. If successful, a new instance of
// TimeZoneType is returned set to the value of the associated enumeration.
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
// caseSensitive   bool - If 'true', the search for enumeration names
//                        will be case sensitive and will require an
//                        exact match. Therefore, 'iana' will NOT
//                        match the enumeration name, 'Iana'.
//
//                        If 'false', a case insensitive search is
//                        conducted for the enumeration name. In
//                        this case, 'iana' will match the
//                        enumeration name 'Iana'.
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
// t, err := TimeZoneType(0).XParseString("Iana", true)
//                            OR
// t, err := TimeZoneType(0).XParseString("Iana()", true)
//                            OR
// t, err := TimeZoneType(0).XParseString("iana", false)
//
// For all of the cases shown above,
//  t is now equal to TimeZoneType(0).Iana()
//
func (tzType TimeZoneType) XParseString(
	valueString string,
	caseSensitive bool) (TimeZoneType, error) {

	lockTimeZoneType.Lock()

	defer lockTimeZoneType.Unlock()

	ePrefix := "TimeZoneType.XParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 4 {
		return TimeZoneType(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"Length Less than 4-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var timeZoneType TimeZoneType

	testStr := strings.ToLower(valueString)

	if testStr == "utc" ||
		testStr == "uct" ||
		testStr == "gmt" {

		// TimeZoneType(0).Iana()
		return TimeZoneType(1), nil

	} else if caseSensitive {

		timeZoneType, ok = mTimeZoneTypeStringToCode[valueString]

		if !ok {
			return TimeZoneType(0),
				errors.New(ePrefix + "\nInvalid TimeZoneType Code!\n")
		}

	} else {
		// caseSensitive must be 'false'

		valueString = strings.ToLower(testStr)

		timeZoneType, ok = mTimeZoneTypeLwrCaseStringToCode[valueString]

		if !ok {
			return TimeZoneType(0),
				errors.New(ePrefix + "\nInvalid TimeZoneType Code!\n")
		}

	}

	return timeZoneType, nil
}

// XValue - Returns the value of the current TimeZoneType instance
// as type TimeZoneType.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (tzType TimeZoneType) XValue() TimeZoneType {

	lockTimeZoneType.Lock()

	defer lockTimeZoneType.Unlock()

	return tzType
}

// XValueInt - Returns the integer value of the current TimeZoneType
// instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (tzType TimeZoneType) XValueInt() int {

	lockTimeZoneType.Lock()

	defer lockTimeZoneType.Unlock()

	return int(tzType)
}

// TzType - public global variable of
// type TimeZoneType.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneType values.
//
// Usage:
//  TzType.None()
//  TzType.Iana()
//  TzType.Military()
//  TzType.Local()
//
var TzType TimeZoneType
