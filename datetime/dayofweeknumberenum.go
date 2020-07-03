package datetime

import (
	"fmt"
	"strings"
	"sync"
)

var mUsDayOfWeekNoStringToCode = map[string]UsDayOfWeekNo{
	"None"            : UsDayOfWeekNo(-1),
	"Sunday"          : UsDayOfWeekNo(0),
	"Monday"          : UsDayOfWeekNo(1),
	"Tuesday"         : UsDayOfWeekNo(2),
	"Wednesday"       : UsDayOfWeekNo(3),
	"Thursday"        : UsDayOfWeekNo(4),
	"Friday"          : UsDayOfWeekNo(5),
	"Saturday"        : UsDayOfWeekNo(6),
}

var mUsDayOfWeekNoLwrCaseStringToCode = map[string]UsDayOfWeekNo{
	"none"            : UsDayOfWeekNo(-1),
	"sunday"          : UsDayOfWeekNo(0),
	"monday"          : UsDayOfWeekNo(1),
	"tuesday"         : UsDayOfWeekNo(2),
	"wednesday"       : UsDayOfWeekNo(3),
	"thursday"        : UsDayOfWeekNo(4),
	"friday"          : UsDayOfWeekNo(5),
	"saturday"        : UsDayOfWeekNo(6),
}

var mUsDayOfWeekNoCodeToString = map[UsDayOfWeekNo]string{
	UsDayOfWeekNo(-1)  : "None",
	UsDayOfWeekNo(0)   : "Sunday",
	UsDayOfWeekNo(1)   : "Monday",
	UsDayOfWeekNo(2)   : "Tuesday",
	UsDayOfWeekNo(3)   : "Wednesday",
	UsDayOfWeekNo(4)   : "Thursday",
	UsDayOfWeekNo(5)   : "Friday",
}

var mUsDayOfWeekNoCodeToAbbrvDay = map[UsDayOfWeekNo]string{
	UsDayOfWeekNo(-1)  : "None",
	UsDayOfWeekNo(0)   : "Sun",
	UsDayOfWeekNo(1)   : "Mon",
	UsDayOfWeekNo(2)   : "Tue",
	UsDayOfWeekNo(3)   : "Wed",
	UsDayOfWeekNo(4)   : "Thu",
	UsDayOfWeekNo(5)   : "Fri",
}


// UsDayOfWeekNo - An enumeration of U.S. (United States) day of
// the week number.
//
// Since Go does not directly support enumerations, the 'UsDayOfWeekNo'
// type has been adapted to function in a manner similar to classic
// enumerations. 'UsDayOfWeekNo' is declared as a type 'int'. The method
// names effectively represent an enumeration of U.S. day of the week
// numbers. These methods are listed as follows:
//
//
// None            (-1) - None - Signals that the U.S. Day Of The Week Number
//                        (UsDayOfWeekNo) Type is not initialized. This is
//                        an error condition.
//
// Sunday           (0) - In accordance with the U.S. day of the week numbering
//                        system, Sunday is the first day of the week and is
//                        represented by a week day number value of 'zero'.
//
// Monday           (1) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Monday', with a weekday number value of '1'.
//
//
// Tuesday          (2) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Tuesday', with a weekday number value of '2'.
//
// Wednesday        (3) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Wednesday', with a weekday number value of '3'.
//
// Thursday         (4) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Thursday', with a weekday number value of '4'.
//
// Friday           (5) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Friday', with a weekday number value of '5'.
//
// Saturday         (6) - Signals that the UsDayOfWeekNo instance is equal
//                        to 'Saturday', with a weekday number value of '6'.
//                        Under the U.S. day of the week numbering system,
//                        'Saturday' is the last day of the week.
//
//
// For easy access to these enumeration values, use the global variable 'UsWeekDayNo'.
// Example: UsWeekDayNo.Tuesday()
//
// Otherwise you will need to use the formal syntax.
// Example: UsDayOfWeekNo(0).Tuesday()
//
// Depending on your editor, intellisense (a.k.a. intelligent code completion) may not
// list the 'UsDayOfWeekNo' methods in alphabetical order. Be advised that all 'UsDayOfWeekNo'
// methods beginning with 'X', as well as the methods 'String()' and 'AbbrvDay()', are utility
// methods and not part of the enumeration values.
//
type UsDayOfWeekNo int

var lockUsDayOfWeekNo sync.Mutex

// None - Signals that the UsDayOfWeekNo Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) None() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(-1)
}

// Sunday - Signals that the UsDayOfWeekNo instance is
// equal to 'Sunday', with a day number value of zero.
// Under the USA weekday numbering system, 'Sunday' is
// the first day of the week.
//
// ISO 8601 specifies that the first day of the week is
// 'Monday' and is numbered as weekday number 1.
//
// The US week day numbering system enumerated here specifies
// that the first day of the week is 'Sunday' and is numbered
// as weekday number zero.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Sunday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(0)
}

// Monday - Signals that the UsDayOfWeekNo instance is equal
// to 'Monday', with a weekday number value of '1'.
//
// ISO 8601 specifies that the first day of the week is
// 'Monday' and is numbered as weekday number 1.
//
// The US week day numbering system enumerated here specifies
// that the first day of the week is 'Sunday' and is numbered
// as weekday number zero.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Monday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(1)
}

// Tuesday - Signals that the UsDayOfWeekNo instance is equal
// to 'Tuesday', with a weekday number value of '2'.
//
// ISO 8601 specifies that the first day of the week is
// 'Monday' and is numbered as weekday number 1.
//
// The US week day numbering system enumerated here specifies
// that the first day of the week is 'Sunday' and is numbered
// as weekday number zero.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Tuesday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(2)
}

// Wednesday - Signals that the UsDayOfWeekNo instance is equal
// to 'Wednesday', with a weekday number value of '3'.
//
// ISO 8601 specifies that the first day of the week is
// 'Monday' and is numbered as weekday number 1.
//
// The US week day numbering system enumerated here specifies
// that the first day of the week is 'Sunday' and is numbered
// as weekday number zero.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Wednesday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(3)
}

// Thursday - Signals that the UsDayOfWeekNo instance is equal
// to 'Thursday', with a weekday number value of '4'.
//
// ISO 8601 specifies that the first day of the week is
// 'Monday' and is numbered as weekday number 1.
//
// The US week day numbering system enumerated here specifies
// that the first day of the week is 'Sunday' and is numbered
// as weekday number zero.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Thursday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(4)
}

// Friday - Signals that the UsDayOfWeekNo instance is equal
// to 'Friday', with a weekday number value of '5'.
//
// ISO 8601 specifies that the first day of the week is
// 'Monday' and is numbered as weekday number 1.
//
// The US week day numbering system enumerated here specifies
// that the first day of the week is 'Sunday' and is numbered
// as weekday number zero.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Friday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(5)
}

// Saturday - Signals that the UsDayOfWeekNo instance is equal
// to 'Saturday', with a weekday number value of '6'. Under the
// USA weekday numbering system, 'Saturday' is the last day of
// the week.
//
// ISO 8601 specifies that the first day of the week is
// 'Monday' and is numbered as weekday number 1.
//
// The US week day numbering system enumerated here specifies
// that the first day of the week is 'Sunday' and is numbered
// as weekday number zero.
//
// This method is part of the standard enumeration.
//
func (usDayOfWk UsDayOfWeekNo) Saturday() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return UsDayOfWeekNo(6)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'UsDayOfWeekNo'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= UsDayOfWeekNo(0).Monday()
// str := t.String()
//     str is now equal to 'Monday'
//
func (usDayOfWk UsDayOfWeekNo) String() string {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	result, ok := mUsDayOfWeekNoCodeToString[usDayOfWk]

	if !ok {
		return ""
	}

	return result
}

// AbbrvDay - Returns a string with the abbreviated name of the day of the
// week associated with this instance of 'UsDayOfWeekNo'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= UsDayOfWeekNo(0).Monday()
// str := t.AbbrvDay()
//     str is now equal to 'Mon'
//
func (usDayOfWk UsDayOfWeekNo) AbbrvDay() string {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	result, ok := mUsDayOfWeekNoCodeToAbbrvDay[usDayOfWk]

	if !ok {
		return ""
	}

	return result
}

// XIsValid - Returns a boolean value signaling
// whether the current U.S. Day of Week Number
// value is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  usDayOfWeek := UsDayOfWeekNo(0).Sunday()
//
//  isValid := usDayOfWeek.XIsValid()
//
func (usDayOfWk UsDayOfWeekNo) XIsValid() bool {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	if usDayOfWk > 6 ||
		usDayOfWk < 0 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of UsDayOfWeekNo is returned set to the value of
// the associated enumeration.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// valueString   string - A string which will be matched against the
//                        enumeration string values. If 'valueString'
//                        is equal to one of the enumeration names, this
//                        method will proceed to successful completion
//                        and return the correct enumeration value.
//
// caseSensitive   bool - If 'true' the search for enumeration names
//                        will be case sensitive and will require an
//                        exact match. Therefore, 'monday' will NOT
//                        match the enumeration name, 'Monday'.
//
//                        If 'false' a case insensitive search is conducted
//                        for the enumeration name. In this case, 'monday'
//                        will match match enumeration name 'Monday'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// UsDayOfWeekNo - Upon successful completion, this method will return a new
//                 instance of UsDayOfWeekNo set to the value of the enumeration
//                 matched by the string search performed on input parameter,
//                'valueString'.
//
// error        - If this method completes successfully, the returned error
//                Type is set equal to 'nil'. If an error condition is encountered,
//                this method will return an error type which encapsulates an
//                appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t, err := UsDayOfWeekNo(0).XParseString("Tuesday", true)
//
//     t is now equal to UsDayOfWeekNo(0).Tuesday()
//
func (usDayOfWk UsDayOfWeekNo) XParseString(
	valueString string,
	caseSensitive bool) (UsDayOfWeekNo, error) {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	ePrefix := "UsDayOfWeekNo.XParseString() "

	if len(valueString) < 4 {
		return UsDayOfWeekNo(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n" +
				"String length is less than '4'.\n" +
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var usDayOfWeek UsDayOfWeekNo

	if caseSensitive {

		usDayOfWeek, ok = mUsDayOfWeekNoStringToCode[valueString]

		if !ok {
			return UsDayOfWeekNo(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CalendarSpec Value.\n" +
					"valueString='%v'\n", valueString)
		}

	} else {

		usDayOfWeek, ok = mUsDayOfWeekNoLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return UsDayOfWeekNo(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CalendarSpec Value.\n" +
					"valueString='%v'\n", valueString)
		}
	}

	return usDayOfWeek, nil
}

// XValue - This method returns the enumeration value of the current CalendarSpec
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (usDayOfWk UsDayOfWeekNo) XValue() UsDayOfWeekNo {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return usDayOfWk
}

// XValueInt - This method returns the integer value of the current CalendarSpec
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (usDayOfWk UsDayOfWeekNo) XValueInt() int {

	lockUsDayOfWeekNo.Lock()

	defer lockUsDayOfWeekNo.Unlock()

	return int(usDayOfWk)
}

// UsWeekDayNo - public global variable of
// type UsDayOfWeekNo.
//
// This variable serves as an easier, short hand
// technique for accessing UsDayOfWeekNo values.
//
// Usage:
// UsWeekDayNo.None(),
// UsWeekDayNo.Sunday(),
// UsWeekDayNo.Monday(),
// UsWeekDayNo.Tuesday(),
// UsWeekDayNo.Wednesday(),
// UsWeekDayNo.Thursday(),
// UsWeekDayNo.Friday(),
// UsWeekDayNo.Saturday(),
//
var UsWeekDayNo UsDayOfWeekNo