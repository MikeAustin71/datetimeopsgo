package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var mTimeZoneCategoryStringToCode = map[string]TimeZoneCategory{
	"None":     TimeZoneCategory(0).None(),
	"TextName":     TimeZoneCategory(0).TextName(),
	"UtcOffset":    TimeZoneCategory(0).UtcOffset(),
}

var mTimeZoneCategoryLwrCaseStringToCode = map[string]TimeZoneCategory{
	"none":     TimeZoneCategory(0).None(),
	"textname":     TimeZoneCategory(0).TextName(),
	"utcoffset":    TimeZoneCategory(0).UtcOffset(),
}

var mTimeZoneCategoryCodeToString = map[TimeZoneCategory]string{
	TimeZoneCategory(0).None():     "None",
	TimeZoneCategory(0).TextName(): "TextName",
	TimeZoneCategory(0).UtcOffset():    "UtcOffset",
}

// TimeZoneCategory - This type is configured as a series of constant integer values
// describing the valid categories of time zones processed by types, 'TimeZoneDefinition'
// and 'TimeZoneSpecification'.
//
// Functionally, 'TimeZoneCategory' serves as an enumeration of valid time zone categories.
//
//                              Time Zone
//                              Category
//    Method                      Code
//    Name                      Constant        Description
//  _______________________________________________________________________________
//
//  None()                          0           Time Zone Category is uninitialized.
//                                              This represents an error condition.
//
//  TextName()                      1           Signals that the Time Zone is identified
//                                              by a standard IANA Text Name.
//
//  UtcOffset()                     2           Signals that the Time Zone is identified
//                                              by a valid UTC Offset and has no associated
//                                              text name.
//
// For easy access to these enumeration values, use the global variable
// 'TzCat'.
//
// 'TimeZoneCategory' has been adapted to function as an enumeration
// describing the classification of time zone assigned to a date time.
// Since Golang does not directly support enumerations, the 'TimeZoneCategory'
// type has been configured to function in a manner similar to classic
// enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type TimeZoneCategory int

var lockTimeZoneCategory sync.Mutex

// None - TimeZoneCategory is uninitialized. This is an error
// condition. It signals that no valid time zone category
// has been identified and assigned to this time zone.
//
// This method is part of the standard enumeration.
//
func (tzCat TimeZoneCategory) None() TimeZoneCategory {
	return TimeZoneCategory(0)
}

// TextName - Signals that the associated time zone is identified
// by a text name like:
//   "America/Chicago"
//   "Asia/Amman"
//   "Atlantic/Bermuda"
//   "Australia/Sydney"
//   "Europe/Rome"
//
func (tzCat TimeZoneCategory) TextName() TimeZoneCategory {
	return TimeZoneCategory(1)
}

// UtcOffset - Signals that the text name associated with the
// current time zone is a UTC Offset consisting of a leading
// plus ('+') or minus ('-') sign a series of numeric digits
// showing the hours or hours and minutes of offset from UTC.
//
// Examples:
//   "+07"
//   "+10"
//
func (tzCat TimeZoneCategory) UtcOffset() TimeZoneCategory {
	return TimeZoneCategory(2)
}

// =========================================================================

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeZoneCategory'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the TimeZoneCategory value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TimeZoneCategory(0).TextName()
//	str := t.String()
//	    str is now equal to "TextName"
//
func (tzCat TimeZoneCategory) String() string {

	lockTimeZoneCategory.Lock()

	defer lockTimeZoneCategory.Unlock()

	label, ok := mTimeZoneCategoryCodeToString[tzCat]

	if !ok {
		return ""
	}

	return label
}

// XParseString - Receives a string and attempts to match it with the string
// value of the supported enumeration. If successful, a new instance of
// TimeZoneCategory is returned set to the value of the associated enumeration.
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
//                        exact match. Therefore, 'textname' will NOT
//                        match the enumeration name, 'TextName'.
//
//                        If 'false' a case insensitive search is
//                        conducted for the enumeration name. In
//                        this case, 'textname' will match the
//                        enumeration name 'TextName'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
// TimeZoneType        - Upon successful completion, this method will return a new
//                       instance of TimeZoneCategory set to the value of the
//                       enumeration matched by the string search performed on
//                       input parameter,'valueString'.
//
// error               - If this method completes successfully, the returned 'error'
//                       Type is set equal to 'nil'. If an error condition is encountered,
//                       this method will return an 'error' Type which encapsulates an
//                       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
// t, err := TimeZoneCategory(0).XParseString("TextName", true)
//                            OR
// t, err := TimeZoneCategory(0).XParseString("TextName()", true)
//                            OR
// t, err := TimeZoneCategory(0).XParseString("textname", false)
//
// For all of the cases shown above,
//  t is now equal to TimeZoneCategory(0).TextName()
//
func (tzCat TimeZoneCategory) XParseString(
	valueString string,
	caseSensitive bool) (TimeZoneCategory, error) {

	lockTimeZoneCategory.Lock()

	defer lockTimeZoneCategory.Unlock()

	ePrefix := "TimeZoneCategory.XParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 8 {
		return TzCat.None(),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"Length Less than 8-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var timeZoneCategory TimeZoneCategory

	testStr := strings.ToLower(valueString)

	if caseSensitive {

		timeZoneCategory, ok = mTimeZoneCategoryStringToCode[valueString]

		if !ok {
			return TzCat.None(),
				errors.New(ePrefix + "Invalid TimeZoneCategory Code!")
		}

	} else {
		// caseSensitive must be 'false'

		valueString = strings.ToLower(testStr)

		timeZoneCategory, ok = mTimeZoneCategoryLwrCaseStringToCode[valueString]

		if !ok {
			return TzCat.None(),
				errors.New(ePrefix + "Invalid TimeZoneCategory Code!")
		}

	}

	return timeZoneCategory, nil
}

// XValue - Returns the value of the current TimeZoneCategory instance
// as type TimeZoneCategory.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (tzCat TimeZoneCategory) XValue() TimeZoneCategory {

	lockTimeZoneCategory.Lock()

	defer lockTimeZoneCategory.Unlock()

	return tzCat
}

// TzCat - public global variable of
// type TimeZoneClass.
//
// This variable serves as an easier, short hand
// technique for accessing TimeZoneClass values.
//
// Usage:
//  TzCat.None()
//  TzCat.AlternateTimeZone()
//  TzCat.OriginalTimeZone()
//
var TzCat TimeZoneCategory
