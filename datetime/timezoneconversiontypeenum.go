package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var mTimeZoneConversionTypeStringToCode = map[string]TimeZoneConversionType{
	"None"    :  TimeZoneConversionType(0),
	"Absolute":  TimeZoneConversionType(1),
	"Relative":  TimeZoneConversionType(2),
}

var mTimeZoneConversionTypeLwrCaseStringToCode = map[string]TimeZoneConversionType{
	"none"     : TimeZoneConversionType(0),
	"absolute" : TimeZoneConversionType(1),
	"relative" : TimeZoneConversionType(2),
}

var mTimeZoneConversionTypeCodeToString = map[TimeZoneConversionType]string{
	TimeZoneConversionType(0) : "None",
	TimeZoneConversionType(1) : "Absolute",
	TimeZoneConversionType(2) : "Relative",
}


// TimeZoneConversionType - This type is configured as a series of
// constant integer values describing the two valid types of time
// zone conversion algorithms used to create new instances of type,
// 'TimeZoneDefinition'.
//
// Functionally, 'TimeZoneConversionType' serves as enumeration of
// valid time zone conversion algorithms.
//
// For easy access to these enumeration values, use the global variable
// 'TzConvertType'.
//
//                      Time Zone
//                      Conversion
//   Method               Type
//    Name              Constant        Description
//  ______________________________________________________________________
//
//  None()                 0            No TimeZoneConversionType has been 
//                                      specified. the TimeZoneConversionType
//                                      value is empty.
//
//  Absolute()             1            Takes the given hours minutes and seconds
//                                      in a given time zone and leaves the hours,
//                                      minutes and seconds value unchanged while
//                                      applying a different time zone. Example:
//                                      Using the "Absolute" conversion method,
//                                      10:00AM in USA Central Standard Time would
//                                      be converted to 10:00AM Eastern Standard
//                                      Time.
//
//  Relative()             2            Converts a given time value to
//                                      its equivalent value in a different
//                                      time zone. If the given time value is
//                                      is in time zone 'X' then than equivalent
//                                      time will be computed in time zone 'Y',
//                                      but the hours, minutes and seconds value
//                                      will be different. Example: Using the
//                                      "Relative" conversion method, 10:00AM in
//                                      USA Central Standard Time would be converted
//                                      to equivalent Eastern Standard Time value of
//                                      11:00AM.
//
// 'TimeZoneConversionType' has been adapted to function as an enumeration
// describing the type of time zone assigned to a date time. Since Golang
// does not directly support enumerations, the 'TimeZoneConversionType'
// type has been configured to function in a manner similar to classic
// enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type TimeZoneConversionType int

var lockTimeZoneConversionType sync.Mutex

// None - Signals that the 'TimeZoneConversionType' value is empty. No 
// 'TimeZoneConversionType' value has been specified.
func (tzConvertType TimeZoneConversionType) None() TimeZoneConversionType {

	lockTimeZoneConversionType.Lock()

	defer lockTimeZoneConversionType.Unlock()

	return TimeZoneConversionType(0)
}


// Absolute - Identifies the 'Absolute' time to time zone conversion algorithm.
// This algorithm provides that times in time zone 'X' will be converted to the
// same time in time zone 'Y'.
//
// For example, assume the time 10:00AM is associated with time zone USA Central
// Standard time and that this time is to be converted to USA Eastern Standard time.
// Applying the 'Absolute' algorithm would convert ths time to 10:00AM Eastern
// Standard time.  In this case the hours, minutes and seconds have not been altered.
// 10:00AM in USA Central Standard Time has simply been reclassified as 10:00AM in
// USA Eastern Standard Time.
//
// For a comparison, see method TimeZoneConversionType(0).Relative() above.
//
func (tzConvertType TimeZoneConversionType) Absolute() TimeZoneConversionType {

	lockTimeZoneConversionType.Lock()

	defer lockTimeZoneConversionType.Unlock()

	return TimeZoneConversionType(1)
}

// Relative - Identifies the 'Relative' time to time zone conversion algorithm.
// This algorithm provides that times in time zone 'X' will be converted to their 
// equivalent time in time zone 'Y'. 
//
// For example, assume the time 10:00AM is associated with time zone USA Central
// Standard time and that this time is to be converted to USA Eastern Standard time.
// Applying the 'Relative' algorithm would convert ths time to 11:00AM Eastern 
// Standard time. In this case the hours, minutes and seconds have been changed to
// reflect an equivalent time in the USA Eastern Standard Time Zone. 
// 
// For a comparison, see method TimeZoneConversionType(0).Absolute() below.
//
func (tzConvertType TimeZoneConversionType) Relative() TimeZoneConversionType {

	lockTimeZoneConversionType.Lock()

	defer lockTimeZoneConversionType.Unlock()

	return TimeZoneConversionType(2)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneConversionType'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return XValue:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneConversionType value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneConversionType(0).Absolute()
//	str := t.String()
//	    str is now equal to "Absolute"
//
func (tzConvertType TimeZoneConversionType) String() string {

	lockTimeZoneConversionType.Lock()

	defer lockTimeZoneConversionType.Unlock()

	label, ok := mTimeZoneConversionTypeCodeToString[tzConvertType]

	if !ok {
		return ""
	}

	return label
}

// XParseString - Receives a string and attempts to match it with the string
// value of the supported enumeration. If successful, a new instance of
// TimeZoneConversionType is returned set to the value of the associated
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
// TimeZoneConversionType - Upon successful completion, this method will return a new
//                          instance of TimeZoneType set to the value of the
//                          enumeration matched by the string search performed on
//                          input parameter,'valueString'.
//
// error                  - If this method completes successfully, the returned error
//                          Type is set equal to 'nil'. If an error condition is encountered,
//                          this method will return an error Type which encapsulates an
//                          appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
// t, err := TimeZoneConversionType(0).XParseString("Relative", true)
//                            OR
// t, err := TimeZoneConversionType(0).XParseString("Relative()", true)
//                            OR
// t, err := TimeZoneConversionType(0).XParseString("relative", false)
//
// For all of the cases shown above,
//  t is now equal to TimeZoneConversionType(0).Relative()
//
func (tzConvertType TimeZoneConversionType) XParseString(
	valueString string,
	caseSensitive bool) (TimeZoneConversionType, error) {

	lockTimeZoneConversionType.Lock()

	defer lockTimeZoneConversionType.Unlock()

	ePrefix := "TimeZoneConversionType.XParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 4 {
		return TzConvertType.None(),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n" +
				"Length Less than 4-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var timeZoneConvertType TimeZoneConversionType

	if caseSensitive {

		timeZoneConvertType, ok = mTimeZoneConversionTypeStringToCode[valueString]

		if !ok {
			return TzConvertType.None(),
				errors.New(ePrefix +
					"\nInvalid TimeZoneConversionType Code!\n")
		}

	} else {

		valueString = strings.ToLower(valueString)

		timeZoneConvertType, ok = mTimeZoneConversionTypeLwrCaseStringToCode[valueString]

		if !ok {
			return TzConvertType.None(),
				errors.New(ePrefix +
					"\nInvalid TimeZoneConversionType Code!\n")
		}

	}

	return timeZoneConvertType, nil
}

// XValue - Returns the value of the TimeZoneConversionType instance
// as type TimeZoneConversionType.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (tzConvertType TimeZoneConversionType) XValue() TimeZoneConversionType {

	lockTimeZoneConversionType.Lock()

	defer lockTimeZoneConversionType.Unlock()

	return tzConvertType
}

// XValueInt - Returns the integer value of the TimeZoneConversionType
// instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (tzConvertType TimeZoneConversionType) XValueInt() int {

	lockTimeZoneConversionType.Lock()

	defer lockTimeZoneConversionType.Unlock()

	return int(tzConvertType)
}

// TzConvertType - public global variable of
// type TimeZoneConversionType.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneConversionType values.
//
// Usage:
//  TzConvertType.None()
//  TzConvertType.Relative()
//  TzConvertType.Absolute()
//
var TzConvertType TimeZoneConversionType
