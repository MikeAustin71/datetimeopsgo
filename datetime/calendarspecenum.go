package datetime

import (
	"fmt"
	"strings"
	"sync"
)

var mCalendarSpecStringToCode = map[string]CalendarSpec{
	"None"           : CalendarSpec(0),
	"Gregorian"      : CalendarSpec(1),
	"Julian"         : CalendarSpec(2),
	"RevisedJulian"  : CalendarSpec(3),
	"RevisedGoucherParker"  : CalendarSpec(4),
}

var mCalendarSpecLwrCaseStringToCode = map[string]CalendarSpec{
	"none"           : CalendarSpec(0),
	"gregorian"      : CalendarSpec(1),
	"julian"         : CalendarSpec(2),
	"revisedjulian"  : CalendarSpec(3),
	"revisedgoucherparker"  : CalendarSpec(4),
}

var mCalendarSpecCodeToString = map[CalendarSpec]string{
	CalendarSpec(0)  : "None",
	CalendarSpec(1)  : "Gregorian",
	CalendarSpec(2)  : "Julian",
	CalendarSpec(3)  : "RevisedJulian",
	CalendarSpec(4)  : "RevisedGoucherParker",
}

// CalendarSpec - An enumeration of calendar designations or types.
// CalendarSpec is used to designate the specific calendar applied
// to a date time operation.
//
// Since Go does not directly support enumerations, the 'CalendarSpec'
// type has been adapted to function in a manner similar to classic
// enumerations. 'CalendarSpec' is declared as a type 'int'. The method
// names effectively represent an enumeration of calendar specification
// types. These methods are listed as follows:
//
//
// None             (0) - None - Signals that the Calendar Specification
//                        (CalendarSpec) Type is not initialized. This is
//                        an error condition.
//
// Gregorian        (1) - Gregorian signals that the Gregorian Calendar is
//                        specified and in effect.
//
//                        The Gregorian calendar is the calendar used for
//                        civilian and commercial purposes throughout most
//                        of the world. It is named after Pope Gregory XIII,
//                        who introduced it Friday, 15 October 1582.
//
// Julian           (2) - Julian signals that the Julian Calendar is specified
//                        and in effect.
//
//                        The Julian calendar, proposed by Julius Caesar in 708
//                        Ab urbe condita (AUC) (46 BC), was a reform of the Roman
//                        calendar. It took effect on 1 January 709 AUC (45 BC),
//                        by edict. It was designed with the aid of Greek mathematicians
//                        and Greek astronomers such as Sosigenes of Alexandria.
//
//                        The calendar was the predominant calendar in the Roman
//                        world, most of Europe, and in European settlements in
//                        the Americas and elsewhere, until it was gradually
//                        replaced by the Gregorian calendar, promulgated in 1582
//                        by Pope Gregory XIII. The Julian calendar is still used
//                        in parts of the Eastern Orthodox Church and in parts of
//                        Oriental Orthodoxy as well as by the Berbers.
//
// RevisedJulian    (3) - Signals that the Revised Julian Calendar is specified and
//                        in effect.
//
//                        The Revised Julian calendar, also known as the Milanković
//                        calendar, or, less formally, new calendar, is a calendar
//                        proposed by the Serbian scientist Milutin Milanković in 1923,
//                        which effectively discontinued the 340 years of divergence
//                        between the naming of dates sanctioned by those Eastern
//                        Orthodox churches adopting it and the Gregorian calendar
//                        that has come to predominate worldwide. This calendar was
//                        intended to replace the ecclesiastical calendar based on
//                        the Julian calendar hitherto in use by all of the Eastern
//                        Orthodox Church. From 1 March 1600 through 28 February 2800,
//                        the Revised Julian calendar aligns its dates with the Gregorian
//                        calendar, which was proclaimed in 1582 by Pope Gregory XIII
//                        for adoption by the Christian world. The calendar has been
//                        adopted by the Orthodox churches of Constantinople, Albania,
//                        Alexandria, Antioch, Bulgaria, Cyprus, Greece, and Romania.
//
// RevisedGoucherParker    (4) - Signals that the RevisedGoucherParker Calendar developed by Adam P.
//                        Goucher, Matt Parker, and modified by Mike Rapp, is specified
//                        and in effect. This calendar type is used by the 'datetime'
//                        package, or library, in order to implement standardized
//                        calendar date time mathematics over a very large number of
//                        years. In its current version, this calendar assumes a mean
//                        solar year of 365.2421897 days and experiences a 1-day calendar
//                        drift approximately every + or - 500-billion years.
//
//
// For easy access to these enumeration values, use the global variable 'CalSpec'.
// Example: CalSpec.Julian()
//
// Otherwise you will need to use the formal syntax.
// Example: CalendarSpec(0).Julian()
//
// Depending on your editor, intellisense (a.k.a. intelligent code completion) may not
// list the CalendarSpec methods in alphabetical order. Be advised that all 'CalendarSpec'
// methods beginning with 'X', as well as the method 'String()', are utility methods and
// not part of the enumeration values.
//
type CalendarSpec int

var lockCalendarSpec sync.Mutex

// None - Signals that the CalendarSpec Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (calSpec CalendarSpec) None() CalendarSpec {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	return CalendarSpec(0)
}

// Gregorian - Signals that the Gregorian Calendar is specified and
// in effect.
//
// The Gregorian calendar is the calendar used for civilian and commercial
// purposes throughout most of the world. It is named after Pope Gregory XIII,
// who introduced it on Friday, 15 October 1582.
//
// The calendar spaces leap years to make the average year 365.2425
// days long, approximating the 365.2422-day tropical year that is
// determined by the Earth's revolution around the Sun. The rule
// for leap years is:
//
// Every year that is exactly divisible by four is a leap year,
// except for years that are exactly divisible by 100, but these
// centurial years are leap years if they are exactly divisible
// by 400. For example, the years 1700, 1800, and 1900 are not
// leap years, but the years 1600 and 2000 are.
//
//
// In the Gregorian calendar, three criteria must be taken
// into account to identify leap years:
//
// The year must be evenly divisible by 4;
//
// If the year can also be evenly divided by 100, it is not a leap year;
// unless...
//
// The year is also evenly divisible by 400. Then it is a leap year.
//
//
// According to these rules, the years 2000 and 2400 are leap years,
// while 1800, 1900, 2100, 2200, 2300, and 2500 are not leap years.
//
// Reference:
//   https://www.timeanddate.com/date/leapyear.html
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//
// This method is part of the standard enumeration.
//
func (calSpec CalendarSpec) Gregorian() CalendarSpec {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	return CalendarSpec(1)
}

// Julian - - Signals that the Julian Calendar is specified and
// in effect.
//
// The Julian calendar, proposed by Julius Caesar in 708 Ab urbe condita
// (AUC) (46 BC), was a reform of the Roman calendar. It took effect on
// 1 January 709 AUC (45 BC), by edict. It was designed with the aid of
// Greek mathematicians and Greek astronomers such as Sosigenes of
// Alexandria.
//
// The calendar was the predominant calendar in the Roman world, most of
// Europe, and in European settlements in the Americas and elsewhere,
// until it was gradually replaced by the Gregorian calendar, promulgated
// in 1582 by Pope Gregory XIII. The Julian calendar is still used in
// parts of the Eastern Orthodox Church and in parts of Oriental Orthodoxy
// as well as by the Berbers.
//
// The Julian calendar has two types of year: a normal year of 365 days
// and a leap year of 366 days. They follow a simple cycle of three normal
// years and one leap year, giving an average year that is 365.25 days long.
// That is more than the actual solar year value of 365.24219 days, which
// means the Julian calendar gains a day every 128 years.
//
// For any given event during the years from 1901 to 2099 inclusive, its
// date according to the Julian calendar is 13 days behind its corresponding
// Gregorian date.
//
// For additional information, reference:
//    https://en.wikipedia.org/wiki/Julian_calendar
//
// This method is part of the standard enumeration.
//
func (calSpec CalendarSpec) Julian() CalendarSpec {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	return CalendarSpec(2)
}

// RevisedJulian - Signals that the Revised Julian Calendar is specified
// and in effect.
//
// The Revised Julian calendar, also known as the Milanković calendar, or,
// less formally, new calendar, is a calendar proposed by the Serbian scientist
// Milutin Milanković in 1923, which effectively discontinued the 340 years
// of divergence between the naming of dates sanctioned by those Eastern
// Orthodox churches adopting it and the Gregorian calendar that has come
// to predominate worldwide. This calendar was intended to replace the
// ecclesiastical calendar based on the Julian calendar hitherto in use
// by all of the Eastern Orthodox Church. From 1 March 1600 through 28 February 2800,
// the Revised Julian calendar aligns its dates with the Gregorian calendar,
// which was proclaimed in 1582 by Pope Gregory XIII for adoption by the
// Christian world. The calendar has been adopted by the Orthodox churches
// of Constantinople, Albania, Alexandria, Antioch, Bulgaria, Cyprus,
// Greece, and Romania.
//
// The Revised Julian calendar has the same months and month lengths as the
// Julian calendar, but, in the Revised Julian calendar, years evenly
// divisible by 100 are not leap years, except that years with remainders
// of 200 or 600 when divided by 900 remain leap years, e.g. 2000 and 2400
// as in the Gregorian Calendar.
//
// For additional information, reference:
//    https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
//
// Summary
//
// 1. Years evenly divisible by 4 are leap years unless they are
//    century years.
//
// 2. Years evenly divisible by 100 are not leap years unless when
//    divided by 900 those years have remainders of 200 or 600 in
//    which case they are leap years.
//
// This method is part of the standard enumeration.
//
func (calSpec CalendarSpec) RevisedJulian() CalendarSpec {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	return CalendarSpec(3)
}

// RevisedGoucherParker - Signals that the RevisedGoucherParker Calendar
// developed by Adam P. Goucher, Matt Parker, and modified by Mike Rapp,
// is specified and in effect. This calendar type is used by the
// 'datetimeops' library to implement standardized calendar date time
// mathematics over a very large number of years. In its current version,
// this calendar cycle is 454,545. The calendar assumes a mean solar year
// of 365.2421897 days and therefore experiences a 1-day drift every + or
// - 500-billion years.
//
// Be advised that the Revised Goucher Parker Calendar implements
// Astronomical year numbering. This means that the Revised Goucher
// Parker Calendar has a year zero; that is the date immediately
// preceding the date January 1, year 1 is December 31, year 0.
//
// Future versions of this calendar may incorporate a sliding scale
// mean solar year to account for variances in the earth's rotational
// dynamics.
//
// "While Matt Parker came up with this system on his own, he does
// acknowledge that British mathematician Adam P. Goucher had suggested
// the idea before. Parker suggests a second alternative — that we go
// back to the Julian calendar, but adjust it so that we skip one leap
// day every 128 years. This system would have it so that the calendar
// would only be off by a day every 625,000 years; essentially, humankind
// would go half a million years before we drifted a day.
//
// Admittedly, calculating what year makes for every 128 years isn’t the
// easiest, so Parker takes it one step further, suggesting we just write
// all our years in binary. If the last seven digits of a binary year are
// zeros, then we skip the leap year. This year is 11111100000 in binary,
// so we’re good."
//
// Reference:
//   https://www.inverse.com/article/12152-how-to-make-a-better-leap-year-with-math
//   https://www.theguardian.com/science/2011/feb/28/leap-year-alex-bellos
//   https://www.youtube.com/watch?v=qkt_wmRKYNQ
//
// The Goucher Parker Calendar as used in the 'datetimeops' library has been
// modified by Mike Rapp to provide an accurate calendar system for computing
// date time mathematics over a very large number of years.
//
// Methodology
//
// 1. Apply the Julian Calendar.
// 2. If year is divisible by 4, it IS a leap year; add on day to February
// 3. If year is divisible by 128 it is NOT a leap year and no day is added to February.
// 4. If year is divisible by 454,545, it IS a leap year; add a day to February.
//
// This will lead to a calendar drift of 1-day every 500-billion (500,000,000,000)
// years. The current version of this calendar assumes a mean solar year of 365.2421897
// days.
//
func (calSpec CalendarSpec) RevisedGoucherParker() CalendarSpec {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	return CalendarSpec(4)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'CalendarSpec'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= CalendarSpec(0).Gregorian()
// str := t.String()
//     str is now equal to 'Gregorian'
//
func (calSpec CalendarSpec) String() string {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	result, ok := mCalendarSpecCodeToString[calSpec]

	if !ok {
		return ""
	}

	return result
}

// XIsValid - Returns a boolean value signaling
// whether the current Calendar Specification
// value is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  CalendarSpecification := CalendarSpec(0).Gregorian()
//
//  isValid := CalendarSpecification.XIsValid()
//
func (calSpec CalendarSpec) XIsValid() bool {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	if calSpec > 4 ||
		calSpec < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of CalendarSpec is returned set to the value of the
// associated enumeration.
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
//                        exact match. Therefore, 'gregorian' will NOT
//                        match the enumeration name, 'Gregorian'.
//
//                        If 'false' a case insensitive search is conducted
//                        for the enumeration name. In this case, 'gregorian'
//                        will match match enumeration name 'Gregorian'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// CalendarSpec - Upon successful completion, this method will return a new
//                instance of CalendarSpec set to the value of the enumeration
//                matched by the string search performed on input parameter,
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
// t, err := CalendarSpec(0).XParseString("Gregorian", true)
//
//     t is now equal to CalendarSpec(0).Gregorian()
//
func (calSpec CalendarSpec) XParseString(
	valueString string,
	caseSensitive bool) (CalendarSpec, error) {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	ePrefix := "CalendarSpec.XParseString() "

	if len(valueString) < 4 {
		return CalendarSpec(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n" +
				"String length is less than '4'.\n" +
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var calendarSpecification CalendarSpec

	if caseSensitive {

		calendarSpecification, ok = mCalendarSpecStringToCode[valueString]

		if !ok {
			return CalendarSpec(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CalendarSpec Value.\n" +
					"valueString='%v'\n", valueString)
		}

	} else {

		calendarSpecification, ok = mCalendarSpecLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return CalendarSpec(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CalendarSpec Value.\n" +
					"valueString='%v'\n", valueString)
		}
	}

	return calendarSpecification, nil
}

// XValue - This method returns the enumeration value of the current CalendarSpec
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (calSpec CalendarSpec) XValue() CalendarSpec {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	return calSpec
}

// XValueInt - This method returns the integer value of the current CalendarSpec
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (calSpec CalendarSpec) XValueInt() int {

	lockCalendarSpec.Lock()

	defer lockCalendarSpec.Unlock()

	return int(calSpec)
}

// CalSpec - public global variable of
// type CalendarSpec.
//
// This variable serves as an easier, short hand
// technique for accessing CalendarSpec values.
//
// Usage:
// CalSpec.None(),
// CalSpec.Gregorian(),
// CalSpec.Julian(),
// CalSpec.RevisedJulian(),
// CalSpec.RevisedGoucherParker(),
//
var CalSpec CalendarSpec
