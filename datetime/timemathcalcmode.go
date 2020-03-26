package datetime

import "sync"

var mTimeMathCalcModeStringToCode = map[string]TimeMathCalcMode{
	"None"           : TimeMathCalcMode(0),
	"LocalTimeZone"  : TimeMathCalcMode(1),
	"UtcTimeZone"    : TimeMathCalcMode(2),
}

var mTimeMathCalcModeToString = map[TimeMathCalcMode]string{
	TimeMathCalcMode(0)  : "None",
	TimeMathCalcMode(1)  : "LocalTimeZone",
	TimeMathCalcMode(2)  : "UtcTimeZone",
}

// TimeMathCalcMode - An enumeration of time calculation modes,
// or algorithms, used in performing addition or subtraction
// operations on date time values.
//
// The 'TimeMathCalcMode' addresses a problem which arises from
// adding or subtracting date times using 'days'. Depending on
// whether the date times being added or subtracted exist in a
// time zone which observes daylight savings time, the term
// 'days'might not be defined as a time span of 24-consecutive
// hours.
//
// Consider the example of a date time like:
//    "2014-02-15 19:54:30.038175584 -0600 CST"
//
// In this case, 'CST' stands for the Central Standard Time Zone
// located in the USA. The Central Standard Time Zone observes
// daylight savings. Generally, Central Standard Time transitions
// to Central Daylight time ('CDT') in March. Following the example
// above, Central Standard Time transitioned to Central Daylight time
// on Sunday, March 9, 2014. Likewise, Central Daylight Time transitioned
// to Central Standard Time on Sunday, November 2, 2019.
//
// The day of March 9, 2014, the day that Central Standard Time transitioned
// to Central Daylight Time, had only 23 hours. At 2:00AM, Sunday, March 9,
// 2014, clocks were moved forward 1-hour to 3:00AM (Spring Forward!). The 1-hour
// time span from 2:00AM to 3:00AM was effectively deleted.
//
// With that background consider the datetime mathematics for our example:
//      "2014-02-15 19:54:30.038175584 -0600 CST"
//
// If we add 5-Years, 6-months and 12-days to this date, what is the
// correct date time. If you use the Golang 'time' package and add
// "5-Years, 6-months and 12-days" to "2014-02-15 19:54:30.038175584 -0600 CST",
// you will generate a result of "2019-08-27 19:54:30.038175584 -0500 CDT".
// This result takes into account days which are NOT 24-hours in length and
// is in fact adjusted for the time zone associated with the original date time.
//
// Now, if you assume that days are always defined as 24-consecutive hours in
// length, you will get a different result.
//
// Coordinated Universal Time (UTC) defines all days as 24-consecutive hours in
// length and therefore does NOT observe daylight savings time.
//
// Reference Coordinated Universal Time (UTC):
//    https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// Therefore, if you assume a day is always defined as 24-hours in length
// the correct way to add "5-Years, 6-months and 12-days" to
// "2014-02-15 19:54:30.038175584 -0600 CST" is to first convert the
// datetime to UTC time.
//
// "2014-02-15 19:54:30.038175584 -0600 CST" is the equivalent of
// "2014-02-16 01:54:30.038175584 +0000 UTC". If you now use the the Golang
// time package to add "5-Years, 6-months and 12-days" to
// "2014-02-16 01:54:30.038175584 +0000 UTC" it will generate a result of
// "2019-08-28 01:54:30.038175584 +0000 UTC". Converting this result back
// to the Central Time Zone, USA will yield a result of
// "2019-08-27 20:54:30.038175584 -0500 CDT". Now compare the results:
//
// "2019-08-27 19:54:30.038175584 -0500 CDT" - Result With Local Time Zone
// "2019-08-27 20:54:30.038175584 -0500 CDT" - Coordinated Universal Time (UTC)
//
// Notice the result of a adding "5-Years, 6-months and 12-days" differs
// by 1-hour depending on whether you use a local time zone datetime or a
// Coordinated Universal Time (UTC) time zone.
//
// The type 'TimeMathCalcMode' allows the calling routine to specify which
// time calculation mode will be used to perform the addition or subtraction
// of time values from a given datetime. Notice that 'TimeMathCalcMode' is
// configured as an enumeration.
//
// Since Go does not directly support enumerations, the 'TimeMathCalcMode' has
// been adapted to function in a manner similar to classic enumerations.
// 'TimeMathCalcMode' is declared as an 'int'. The method names are effectively
// an enumeration of available calculation modes. These methods are listed as
// follows:
//
// None             (0) - None - Signals that Time Math Calculation Mode
//                        is not initialized. This is an error condition.
//
// LocalTimeZone    (1) - LocalTimeZone Time Calculation Mode. This mode
//                        specifies that time spans added or subtracted
//                        from a date time value will use the local time
//                        zone. This means that days MAY be lesser than
//                        greater than 24-consecutive hours.
//
// UtcTimeZone      (2) - Coordinated Universal Time (UTC) Mode. This mode
//                        specifies that time spans added to or subtracted
//                        from date time values will always use days defined
//                        as 24-consecutive hours. At a detail level, the
//                        original date time is first converted to UTC time
//                        before adding the time span to yield an intermediate
//                        result. This intermediate result is then converted
//                        back to the equivalent time zone of the original
//                        date time.
//
// For easy access to these enumeration values, use the global variable
// 'TCalcMode'. Example: TCalcMode.UtcTimeZone()
//
// Otherwise you will need to use the formal syntax:
// TimeMathCalcMode(0).UtcTimeZone()
//
type TimeMathCalcMode int

var lockTimeMathCalcMode sync.Mutex

func (tMathMode TimeMathCalcMode) None() TimeMathCalcMode {

	lockTimeMathCalcMode.Lock()

	defer lockTimeMathCalcMode.Unlock()

	return TimeMathCalcMode(0)
}


func (tMathMode TimeMathCalcMode) LocalTimeZone() TimeMathCalcMode {

	lockTimeMathCalcMode.Lock()

	defer lockTimeMathCalcMode.Unlock()

	return TimeMathCalcMode(1)
}


func (tMathMode TimeMathCalcMode) UtcTimeZone() TimeMathCalcMode {

	lockTimeMathCalcMode.Lock()

	defer lockTimeMathCalcMode.Unlock()

	return TimeMathCalcMode(1)
}


// ===============================================================================
//                     UTILITY METHODS
// ===============================================================================

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TimeMathCalcMode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= TDurCalcType(0).StdYearMth()
// str := t.String()
//     str is now equal to 'StdYearMth'
//
func (tMathMode TimeMathCalcMode) String() string {

	lockTimeMathCalcMode.Lock()

	defer lockTimeMathCalcMode.Unlock()

	result, ok := mTimeMathCalcModeToString[tMathMode]

	if !ok {
		return ""
	}

	return result
}



// TCalcMode - public global variable of
// type TimeMathCalcMode.
//
// This variable serves as an easier, short hand
// technique for accessing TimeMathCalcMode values.
//
// Usage:
// TCalcMode.None(),
// TCalcMode.LocalTimeZone(),
// TCalcMode.UtcTimeZone(),
//
var TCalcMode TimeMathCalcMode