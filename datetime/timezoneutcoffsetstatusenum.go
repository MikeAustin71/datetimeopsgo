package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var mTimeZoneUtcOffsetStatusStringToCode = map[string]TimeZoneUtcOffsetStatus{
	"None"    :  TimeZoneUtcOffsetStatus(0),
	"Static"  :  TimeZoneUtcOffsetStatus(1),
	"Variable":  TimeZoneUtcOffsetStatus(2),
}

var mTimeZoneUtcOffsetStatusLwrCaseStringToCode = map[string]TimeZoneUtcOffsetStatus{
	"none"    : TimeZoneUtcOffsetStatus(0),
	"static"  : TimeZoneUtcOffsetStatus(1),
	"variable": TimeZoneUtcOffsetStatus(2),
}

var mTimeZoneUtcOffsetStatusCodeToString = map[TimeZoneUtcOffsetStatus]string{
	TimeZoneUtcOffsetStatus(0):  "None",
	TimeZoneUtcOffsetStatus(1):  "Static",
	TimeZoneUtcOffsetStatus(2):  "Variable",
}

// TimeZoneUtcOffsetStatus - This type is configured as a series of
// constant integer values describing the status of the UTC Offset
// associated with a given time zone.  'TimeZoneUtcOffsetStatus' is
// used in processing time zone for types 'TimeZoneDefinition' and
// 'TimeZoneSpecification'.
//
// Functionally, 'TimeZoneUtcOffsetStatus' serves as an enumeration
// of valid UTC Offset status descriptions associated with a given
// Time Zone.
//
//                      Time Zone
//   Method             Type Code
//    Name              Constant        Description
//  __________________________________________________________________________
//  None()                  0           Time Zone UTC Offset Status is
//                                      uninitialized and contains no
//                                      significant or valid value. This
//                                      is an error condition.
//
//  Static()                1           Signals that the UTC Offset associated
//                                      with a given Time Zone is constant
//                                      throughout the year and never changes.
//                                      Typically, this means that Daylight
//                                      Savings Time is NOT observed in the
//                                      specified Time Zone.
//
//  Variable()              2           Signals that the UTC Offset associated
//                                      with a given Time Zone is not constant,
//                                      and varies at least once during the year.
//                                      This usually means that Daylight Savings
//                                      Time is observed within the designated
//                                      Time Zone.
//
// For easy access to these enumeration values, use the global variable
// 'TzUtcStatus'.
//
// 'TimeZoneUtcOffsetStatus' has been adapted to function as an enumeration
// describing the type of time zone assigned to a date time. Since
// Golang does not directly support enumerations, the 'TimeZoneUtcOffsetStatus'
// type has been configured to function in a manner similar to classic
// enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type TimeZoneUtcOffsetStatus int

var lockTimeZoneUtcOffsetStatus sync.Mutex


// None - TimeZoneUtcOffsetStatus is uninitialized. This is an error
// condition. It signals that no valid Time Zone UTC Offset Status
// description has been identified and assigned to this time zone.
//
// This method is part of the standard enumeration.
//
func (tzUtcStat TimeZoneUtcOffsetStatus) None() TimeZoneUtcOffsetStatus {

	lockTimeZoneUtcOffsetStatus.Lock()

	defer lockTimeZoneUtcOffsetStatus.Unlock()

	return TimeZoneUtcOffsetStatus(0)
}

// Static - Signals that the UTC Offset associated with a given
// Time Zone is constant throughout the year and never changes.
// Typically, this means that Daylight Savings Time is NOT
// observed within the specified Time Zone.
//
func (tzUtcStat TimeZoneUtcOffsetStatus) Static() TimeZoneUtcOffsetStatus {

	lockTimeZoneUtcOffsetStatus.Lock()

	defer lockTimeZoneUtcOffsetStatus.Unlock()

	return TimeZoneUtcOffsetStatus(1)
}

// Variable - Signals that the UTC Offset associated with a given
// Time Zone is not constant, and varies at least once during the
// year. This usually means that Daylight Savings Time is observed
// within the designated Time Zone.
//
func (tzUtcStat TimeZoneUtcOffsetStatus) Variable() TimeZoneUtcOffsetStatus {

	lockTimeZoneUtcOffsetStatus.Lock()

	defer lockTimeZoneUtcOffsetStatus.Unlock()

	return TimeZoneUtcOffsetStatus(2)
}


// =========================================================================

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneUtcOffsetStatus'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the 'TimeZoneUtcOffsetStatus' value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= TimeZoneUtcOffsetStatus(0).Static()
// str := t.String()
//     str is now equal to "Static"
//
func (tzUtcStat TimeZoneUtcOffsetStatus) String() string {

	lockTimeZoneUtcOffsetStatus.Lock()

	defer lockTimeZoneUtcOffsetStatus.Unlock()

	label, ok := mTimeZoneUtcOffsetStatusCodeToString[tzUtcStat]

	if !ok {
		return ""
	}

	return label
}


// XParseString - Receives a string and attempts to match it with the string
// value of the supported enumeration. If successful, a new instance of
// TimeZoneUtcOffsetStatus is returned set to the value of the associated
// enumeration.
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
//                        exact match. Therefore, 'variable' will NOT
//                        match the enumeration name, 'Variable'.
//
//                        If 'false', a case insensitive search is
//                        conducted for the enumeration name. In
//                        this case, 'variable' will match the
//                        enumeration name 'Variable'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
// TimeZoneUtcOffsetStatus - Upon successful completion, this method will return a new
//                           instance of 'TimeZoneUtcOffsetStatus' set to the value of
//                           the enumeration matched by the string search performed on
//                           input parameter,'valueString'.
//
// error                   - If this method completes successfully, the returned error
//                           Type is set equal to 'nil'. If an error condition is encountered,
//                           this method will return an error Type which encapsulates an
//                           appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
// t, err := TimeZoneUtcOffsetStatus(0).XParseString("Variable", true)
//                            OR
// t, err := TimeZoneUtcOffsetStatus(0).XParseString("Variable()", true)
//                            OR
// t, err := TimeZoneUtcOffsetStatus(0).XParseString("variable", false)
//
// For all of the cases shown above,
//  t is now equal to TimeZoneUtcOffsetStatus(0).Variable()
//
func (tzUtcStat TimeZoneUtcOffsetStatus) XParseString(
	valueString string,
	caseSensitive bool) (TimeZoneUtcOffsetStatus, error) {

	lockTimeZoneUtcOffsetStatus.Lock()

	defer lockTimeZoneUtcOffsetStatus.Unlock()

	ePrefix := "TimeZoneUtcOffsetStatus.XParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 4 {
		return TimeZoneUtcOffsetStatus(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"Length Less than 4-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var timeZoneUtcOffsetStatus TimeZoneUtcOffsetStatus

	testStr := strings.ToLower(valueString)

	if caseSensitive {

		timeZoneUtcOffsetStatus, ok = mTimeZoneUtcOffsetStatusStringToCode[valueString]

		if !ok {
			return TimeZoneUtcOffsetStatus(0),
				errors.New(ePrefix + "\nInvalid TimeZoneUtcOffsetStatus Code!\n")
		}

	} else {
		// caseSensitive must be 'false'

		valueString = strings.ToLower(testStr)

		timeZoneUtcOffsetStatus, ok = mTimeZoneUtcOffsetStatusLwrCaseStringToCode[valueString]

		if !ok {
			return TimeZoneUtcOffsetStatus(0),
				errors.New(ePrefix + "\nInvalid TimeZoneUtcOffsetStatus Code!\n")
		}

	}

	return timeZoneUtcOffsetStatus, nil
}

// XValue - Returns the value of the current 'TimeZoneUtcOffsetStatus'
// instance as type 'TimeZoneUtcOffsetStatus'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (tzUtcStat TimeZoneUtcOffsetStatus) XValue() TimeZoneUtcOffsetStatus {

	lockTimeZoneUtcOffsetStatus.Lock()

	defer lockTimeZoneUtcOffsetStatus.Unlock()

	return tzUtcStat
}

// XValueInt - Returns the integer value of the current
// 'TimeZoneUtcOffsetStatus' instance.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
//
func (tzUtcStat TimeZoneUtcOffsetStatus) XValueInt() int {

	lockTimeZoneUtcOffsetStatus.Lock()

	defer lockTimeZoneUtcOffsetStatus.Unlock()

	return int(tzUtcStat)
}


// TzUtcStatus - public global variable of type
// 'TimeZoneUtcOffsetStatus'.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneType values.
//
// Usage:
//  TzUtcStatus.None()
//  TzUtcStatus.Static()
//  TzUtcStatus.Variable()
//
var TzUtcStatus TimeZoneUtcOffsetStatus
