package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var mTimeZoneClassStringToCode = map[string]TimeZoneClass{
	"None"             : TimeZoneClass(0),
	"AlternateTimeZone": TimeZoneClass(1),
	"OriginalTimeZone" : TimeZoneClass(2),
}

var mTimeZoneClassLwrCaseStringToCode = map[string]TimeZoneClass{
	"none"             : TimeZoneClass(0),
	"alternatetimezone": TimeZoneClass(1),
	"originaltimezone" : TimeZoneClass(2),
}

var mTimeZoneClassCodeToString = map[TimeZoneClass]string{
	TimeZoneClass(0) : "None",
	TimeZoneClass(1) : "AlternateTimeZone",
	TimeZoneClass(2) : "OriginalTimeZone",
}


// TimeZoneClass - This type is configured as a series of constant integer values
// describing the valid classes of time zones processed by types, 'TimeZoneDefinition'
// and 'TimeZoneSpecification'.
//
// Functionally, 'TimeZoneClass' serves as an enumeration of valid time zone classes.
//
//                              Time Zone
//   Method                     Type Code
//    Name                      Constant        Description
//  _______________________________________________________________________________
//
//  None()                          0           Time Zone Class is uninitialized.
//                                              This represents an error condition.
//
//  AlternateTimeZone()             1           Signals that the original Time Zone
//                                              was invalid and that a valid Time
//                                              Zone was substituted.
//
//  OriginalTimeZone()              2           Signals that the original Time Zone
//                                              was valid and that no substitute
//                                              time zone was required.
//
// For easy access to these enumeration values, use the global variable
// 'TzClass'.
//
// 'TimeZoneClass' has been adapted to function as an enumeration
// describing the classification of time zone assigned to a date time.
// Since Golang does not directly support enumerations, the 'TimeZoneClass'
// type has been configured to function in a manner similar to classic
// enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type TimeZoneClass int

var lockTimeZoneClass sync.Mutex


// None - TimeZoneClass is uninitialized. This is an error
// condition. It signals that no valid, convertible time
// zone could be identified or generated.
//
// As used here a 'convertible' time zone is one which is
// convertible across all world time zones. This means that
// the 'convertible' time zone may be used to accurately
// convert date times in that time zone to any other time
// zone. Conversely, a non-convertible time zone means that
// date times may NOT be accurately converted to equivalent
// date times in other time zones.
//
// This method is part of the standard enumeration.
//
func (tzClass TimeZoneClass) None() TimeZoneClass {

	lockTimeZoneClass.Lock()

	defer lockTimeZoneClass.Unlock()

	return TimeZoneClass(0)
}

// AlternateTimeZone - Signals that the original time zone
// was invalid and non-convertible. Therefore, a valid
// convertible time zone was generated from the time zone
// abbreviation and substituted for the original time zone.
//
// As used here a 'convertible' time zone is one which is
// convertible across all world time zones. This means that
// the 'convertible' time zone may be used to accurately
// convert date times in that time zone to any other time
// zone. Conversely, a non-convertible time zone means that
// date times may NOT be accurately converted to equivalent
// date times in other time zones.
//
// This method is part of the standard enumeration.
//
func (tzClass TimeZoneClass) AlternateTimeZone() TimeZoneClass {

	lockTimeZoneClass.Lock()

	defer lockTimeZoneClass.Unlock()

	return TimeZoneClass(1)
}

// OriginalTimeZone - Signals that the original time zone
// was a valid, convertible time zone and generation of
// an alternate substitute time zone was not required.
// This time zone is the original valid time zone and
// may be safely used to calculate equivalent date times
// in other time zones.
//
// As used here a 'convertible' time zone is one which is
// convertible across all world time zones. This means that
// the 'convertible' time zone may be used to accurately
// convert date times in that time zone to any other time
// zone. Conversely, a non-convertible time zone means that
// date times may NOT be accurately converted to equivalent
// date times in other time zones.
//
// This method is part of the standard enumeration.
//
func (tzClass TimeZoneClass) OriginalTimeZone() TimeZoneClass {

	lockTimeZoneClass.Lock()

	defer lockTimeZoneClass.Unlock()

	return TimeZoneClass(2)
}

// =========================================================================

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneClass'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return XValue:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneClass value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneClass(0).AlternateTimeZone()
//	str := t.String()
//	    str is now equal to "AlternateTimeZone"
//
func (tzClass TimeZoneClass) String() string {

	lockTimeZoneClass.Lock()

	defer lockTimeZoneClass.Unlock()

	label, ok := mTimeZoneClassCodeToString[tzClass]

	if !ok {
		return ""
	}

	return label
}


// XParseString - Receives a string and attempts to match it with the string
// value of the supported enumeration. If successful, a new instance of
// TimeZoneClass is returned set to the value of the associated enumeration.
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
//                        exact match. Therefore, 'none' will NOT
//                        match the enumeration name, 'None'.
//
//                        If 'false' a case insensitive search is
//                        conducted for the enumeration name. In
//                        this case, 'none' will match the
//                        enumeration name 'None'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
// TimeZoneClass       - Upon successful completion, this method will return a new
//                       instance of TimeZoneClass set to the value of the
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
// t, err := TimeZoneClass(0).XParseString("AlternateTimeZone", true)
//                            OR
// t, err := TimeZoneClass(0).XParseString("AlternateTimeZone()", true)
//                            OR
// t, err := TimeZoneClass(0).XParseString("alternatetimezone", false)
//
// For all of the cases shown above,
//  t is now equal to TimeZoneClass(0).AlternateTimeZone()
//
func (tzClass TimeZoneClass) XParseString(
	valueString string,
	caseSensitive bool) (TimeZoneClass, error) {

	lockTimeZoneClass.Lock()

	defer lockTimeZoneClass.Unlock()

	ePrefix := "TimeZoneClass.XParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 4 {
		return TimeZoneClass(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"Length Less than 4-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var timeZoneClass TimeZoneClass

	testStr := strings.ToLower(valueString)

	if caseSensitive {

		timeZoneClass, ok = mTimeZoneClassStringToCode[valueString]

		if !ok {
			return TimeZoneClass(0),
				errors.New(ePrefix +
					"\nInvalid TimeZoneClass Code!\n")
		}

	} else {
		// caseSensitive must be 'false'

		valueString = strings.ToLower(testStr)

		timeZoneClass, ok = mTimeZoneClassLwrCaseStringToCode[valueString]

		if !ok {
			return TimeZoneClass(0),
				errors.New(ePrefix +
					"\nInvalid TimeZoneClass Code!\n")
		}

	}

	return timeZoneClass, nil
}

// XValue - Returns the value of the TimeZoneClass instance
// as type TimeZoneClass.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (tzClass TimeZoneClass) XValue() TimeZoneClass {

	lockTimeZoneClass.Lock()

	defer lockTimeZoneClass.Unlock()

	return tzClass
}

// XValueInt - Returns the integer value of the current TimeZoneClass
// instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (tzClass TimeZoneClass) XValueInt() int {

	lockTimeZoneClass.Lock()

	defer lockTimeZoneClass.Unlock()

	return int(tzClass)
}

// TzClass - public global variable of
// type TimeZoneClass.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneClass values.
//
// Usage:
//  TzClass.None()
//  TzClass.AlternateTimeZone()
//  TzClass.OriginalTimeZone()
//
var TzClass TimeZoneClass
