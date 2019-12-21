package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var lockMapTimeZoneTypeStringToCode = sync.Mutex{}

var mTimeZoneTypeStringToCode = map[string]TimeZoneType{
	"None":          TimeZoneType(0).None(),
	"Iana":          TimeZoneType(0).Iana(),
	"Military":      TimeZoneType(0).Military(),
	"Local":         TimeZoneType(0).Local(),
	"Utc":           TimeZoneType(0).Utc(),
	"UtcOffset":     TimeZoneType(0).UtcOffset(),
}

var lockMapTimeZoneTypeLwrCaseStringToCode = sync.Mutex{}

var mTimeZoneTypeLwrCaseStringToCode = map[string]TimeZoneType{
	"none":          TimeZoneType(0).None(),
	"iana":          TimeZoneType(0).Iana(),
	"military":      TimeZoneType(0).Military(),
	"local":         TimeZoneType(0).Local(),
	"utc":           TimeZoneType(0).Utc(),
	"utcoffset":     TimeZoneType(0).UtcOffset(),
}

var lockMapTimeZoneTypeCodeToString = sync.Mutex{}

var mTimeZoneTypeCodeToString = map[TimeZoneType]string{
	TimeZoneType(0).None():          "None",
	TimeZoneType(0).Iana():          "Iana",
	TimeZoneType(0).Military():      "Military",
	TimeZoneType(0).Local():         "Local",
	TimeZoneType(0).Utc():           "Utc",
	TimeZoneType(0).UtcOffset():     "UtcOffset",
}



// TimeZoneType - This type is configured as a series of
// constant integer values describing the valid types
// of time zones processed by type 'TimeZoneDefDto'.
//
// Functionally, 'TimeZoneType' serves as enumeration of
// valid time zone types.
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
//  Military()              2           Tests have established that the
//                                        Time Zone is a valid, standard
//                                        Military Time Zone.
//
//  Local()                 3           Tests have established that the
//                                        Time Zone is 'Local'. This is
//                                        the time zone currently configured
//                                        for the host computer.
//
//  Utc                     4           Tests have established that this
//                                        Time Zone is Coordinated Universal
//                                        Time or 'UTC'. In some countries,
//                                        the term "Greenwich Mean Time (GMT)"
//                                        is used as an equivalent for 'UTC'.
//
//  UtcOffset()             5           Tests have established that no
//                                        known time zone has been applied.
//                                        Instead, the time zone is identified
//                                        only by its UTC offset.
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

// None - TimeZoneType is uninitialized and has no value.
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) None() TimeZoneType { return TimeZoneType(0) }

// Iana - Classifies the time zone as part of the IANA Time Zone database
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) Iana() TimeZoneType {return TimeZoneType(1)}

// Military - Classifies the time zone as a standard military time zone.
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) Military() TimeZoneType {return TimeZoneType(2)}

// Local - The 'Local' time zone is construct of the Go programming language.
// It signals that the time zone currently configured on the host computer
// has been selected.
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) Local() TimeZoneType {return TimeZoneType(3)}

// Utc - Coordinated Universal Time (UTC or UTC) is the primary
// time standard by which the world regulates clocks and time.
// It is within about 1-second of mean solar time at 0° longitude,
// and is not adjusted for daylight saving time. In some countries,
// the term "Greenwich Mean Time (GMT)" is used as an equivalent
// for 'UTC'.
//
// UTC is equivalent to a zero offset: UTC+0000. For additional
// information, reference:
//
//     https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) Utc() TimeZoneType {return TimeZoneType(4)}

// UtcOffset - This type indicates that no standard IANA, Military or
// Local time zone applies. Instead, this time zone is identified only
// by its UTC offset.
//
// This method is part of the standard enumeration.
//
func (tzType TimeZoneType) UtcOffset() TimeZoneType {return TimeZoneType(5)}

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
//	t:= TimeZoneType(0).Military()
//	str := t.String()
//	    str is now equal to "Military"
//
func (tzType TimeZoneType) String() string {

	lockMapTimeZoneTypeCodeToString.Lock()

	defer lockMapTimeZoneTypeCodeToString.Unlock()

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
// caseSensitive   bool - If 'true' the search for enumeration names
//                        will be case sensitive and will require an
//                        exact match. Therefore, 'valid' will NOT
//                        match the enumeration name, 'Valid'.
//
//                        If 'false' a case insensitive search is
//                        conducted for the enumeration name. In
//                        this case, 'valid' will match the
//                        enumeration name 'Valid'.
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
// t, err := TimeZoneType(0).ParseString("Iana", true)
//                            OR
// t, err := TimeZoneType(0).ParseString("Iana()", true)
//                            OR
// t, err := TimeZoneType(0).ParseString("iana", false)
//
// For all of the cases shown above,
//  t is now equal to TimeZoneType(0).Iana()
//
func (tzType TimeZoneType) XParseString(
	valueString string,
	caseSensitive bool) (TimeZoneType, error) {

	ePrefix := "TimeZoneType.XParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 3 {
		return TzType.None(),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n" +
				"Length Less than 3-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var timeZoneType TimeZoneType

	testStr := strings.ToLower(valueString)

	if testStr == "utc" ||
			testStr == "uct"  ||
				testStr == "gmt" {

		return TzType.Utc(), nil

	} else if caseSensitive {

		lockMapTimeZoneTypeStringToCode.Lock()

		defer lockMapTimeZoneTypeStringToCode.Unlock()

		timeZoneType, ok = mTimeZoneTypeStringToCode[valueString]

		if !ok {
			return TzType.None(),
				errors.New(ePrefix + "Invalid TimeZoneType Code!")
		}

	} else {
		// caseSensitive must be 'false'
		valueString = strings.ToLower(valueString)

		lockMapTimeZoneTypeLwrCaseStringToCode.Lock()

		defer lockMapTimeZoneTypeLwrCaseStringToCode.Unlock()

		timeZoneType, ok = mTimeZoneTypeLwrCaseStringToCode[valueString]

		if !ok {
			return TzType.None(),
				errors.New(ePrefix + "Invalid TimeZoneType Code!")
		}

	}

	return timeZoneType, nil
}

// XValue - Returns the value of the TimeZoneType instance
// as type TimeZoneType.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (tzType TimeZoneType) XValue() TimeZoneType {

	return tzType
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
//  TzType.UtcOffset()
//
var TzType TimeZoneType
