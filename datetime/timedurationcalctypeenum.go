package datetime

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)


// mTDurCalcTypeIntToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type TDurCalcType.
var mTDurCalcTypeIntToString = map[int]string{}

// mTDurCalcTypeStringToInt - This map is used to map enumeration names
// stored as strings to enumeration values for Type TDurCalcType.
var mTDurCalcTypeStringToInt = map[string]int{}

// mTDurCalcTypeStringToInt - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type TDurCalcType.
// This map is used for case insensitive look ups.
var mTDurCalcTypeLwrCaseStringToInt = map[string]int{}

// TDurCalcType - An enumeration of time duration calculation types.
// The time duration calculation can allocate elapsed time by multiple
// combinations of Years, Months, Weeks, Days, Hours, Seconds, Milliseconds,
// Microseconds, Nanoseconds and/or GregorianYears.
//
// This time duration calculation type determines how time duration is
// is calculated and accumulated by these categories. The calculation stores
// allocated time duration in a Type TimeDurationDto struct. (See source file
// 'datetimeopsgo/datetime/timedurationdto.go')
//
// Since Go does not directly support enumerations, the 'TDurCalcType' has
// been adapted to function in a manner similar to classic enumerations.
// 'TDurCalcType' is declared as an 'int'. The method names are effectively
// an enumeration of calculation types. These methods are listed as follows:
//
//	StdYearMth      - StdYearMth Calculation type. Allocates time duration by
//	                  Standard Year, months, weeks, days, hours, minutes, seconds,
//	                  milliseconds, microseconds, and nanoseconds.
//
//	CumMonths       - Cumulative Months Calculation. Allocates time duration by
//	                  cumulative months plus weeks, week days, date days, hours,
//	                  minutes, seconds, milliseconds, microseconds and nanoseconds.
//
//	CumWeeks        - Cumulative Weeks Calculation. Allocates time duration by cumulative
//	                  weeks, plus week days, hours, minutes, seconds, milliseconds,
//	                  microseconds and nanoseconds.
//
//	CumDays         - Cumulative Days calculation. Allocates time duration by cumulative
//	                  days plus hours, minutes, seconds, milliseconds, microseconds and
//	                  nanoseconds.
//
//	CumHours        - Cumulative Hours calculation. Allocates time duration by cumulative
//	                  hours plus minutes, seconds, milliseconds, microseconds and
//	                  nanoseconds.
//
//	CumMinutes      - Cumulative Minutes calculation. Allocates time duration by
//	                  cumulative minutes plus seconds, milliseconds, microseconds and
//	                  nanoseconds.
//
//	CumSeconds      - Cumulative Seconds calculation. Allocates time duration by
//	                  cumulative seconds plus milliseconds, microseconds and
//	                  nanoseconds.
//
//	CumMilliseconds - Cumulative Milliseconds calculation. Allocates time duration
//	                  by cumulative milliseconds plus microseconds and nanoseconds.
//
//	CumMicroseconds - Cumulative Microseconds calculation. Allocates time duration
//	                  by cumulative microseconds plus nanoseconds.
//
//	CumNanoseconds - Cumulative Nanoseconds calculation. Allocates time duration
//	                 by cumulative nanoseconds.
//
//	GregorianYears - Gregorian Year calculation. Allocates time duration by Years,
//	                 Months, Weeks, WeekDays, Date Days, Hours, Minutes, Seconds,
//	                 Milliseconds, Microseconds, Nanoseconds. However, the Years
//	                 allocation is performed using standard Gregorian Years.
//
type TDurCalcType int

var lockTDurCalcType sync.Mutex

// StdYearMth - Allocates time duration by Standard Year, months, weeks,
// weekdays, date days, hours, minutes, seconds, milliseconds, microseconds
// and nanoseconds. For the 'StdYearMth' time duration calculation type, all
// data fields in the TimeDurationDto	structure are populated.
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).StdYearMth()
//	    Years                       populated
//	    YearsNanosecs               populated
//	    Months                      populated
//	    MonthsNanosecs              populated
//	    Weeks                       populated
//	    WeeksNanosecs               populated
//	    WeekDays                    populated
//	    WeekDaysNanosecs            populated
//	    DateDays                    populated
//	    DateDaysNanosecs            populated
//	    Hours                       populated
//	    HoursNanosecs               populated
//	    Minutes                     populated
//	    MinutesNanosecs             populated
//	    Seconds                     populated
//	    SecondsNanosecs             populated
//	    Milliseconds                populated
//	    MillisecondsNanosecs        populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).StdYearMth()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).StdYearMth())
//
//
func (TDurCalcType) StdYearMth() TDurCalcType { return TDurCalcType(0) }

// CumMonths - CumMonths - Cumulative Months Calculation. Years are ignored.
// Years and Months are consolidated and counted as cumulative months. Years
// duration is not provided. The entire duration is broken down by cumulative
// months plus weeks, week days, date days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
//
// The Data Fields for Years is set to zero.
//
// For the 'CumMonths' time duration calculation type, the following data fields
// are populated in the TimeDurationDto structure:
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).CumMonths()
//	    Years                       NOT-populated
//	    YearsNanosecs               NOT-populated
//	    Months                      populated
//	    MonthsNanosecs              populated
//	    Weeks                       populated
//	    WeeksNanosecs               populated
//	    WeekDays                    populated
//	    WeekDaysNanosecs            populated
//	    DateDays                    populated
//	    DateDaysNanosecs            populated
//	    Hours                       populated
//	    HoursNanosecs               populated
//	    Minutes                     populated
//	    MinutesNanosecs             populated
//	    Seconds                     populated
//	    SecondsNanosecs             populated
//	    Milliseconds                populated
//	    MillisecondsNanosecs        populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).CumMonths()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumMonths())
//
func (TDurCalcType) CumMonths() TDurCalcType { return TDurCalcType(1) }

// CumWeeks - Cumulative Weeks calculation. Years and months are ignored. Instead,
// years, months, weeks are consolidated and counted as cumulative weeks. Years and
// months duration are not calculated. The entire duration is broken down by cumulative
// weeks, plus week days, hours, minutes, seconds, milliseconds, microseconds and
// nanoseconds. Data Fields for Years and Months are always set to zero.
//
// For the 'CumWeeks' time duration calculation type, the following data fields are
// populated in the TimeDurationDto structure:
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).CumWeeks()
//	    Years                       NOT-populated
//	    YearsNanosecs               NOT-populated
//	    Months                      NOT-populated
//	    MonthsNanosecs              NOT-populated
//	    Weeks                       populated
//	    WeeksNanosecs               populated
//	    WeekDays                    populated
//	    WeekDaysNanosecs            populated
//	    DateDays                    NOT-populated
//	    DateDaysNanosecs            NOT-populated
//	    Hours                       populated
//	    HoursNanosecs               populated
//	    Minutes                     populated
//	    MinutesNanosecs             populated
//	    Seconds                     populated
//	    SecondsNanosecs             populated
//	    Milliseconds                populated
//	    MillisecondsNanosecs        populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).CumWeeks()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumWeeks())
//
func (TDurCalcType) CumWeeks() TDurCalcType { return TDurCalcType(2) }

// CumDays - Cumulative Days calculation. Years, months and weeks are ignored.
// Years, months, weeks and days are consolidated and counted as cumulative days.
// Years, months and weeks duration are not calculated. The entire time duration
// is broken down by cumulative days plus hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds. Data Fields for years, months, weeks, and weekdays
// are always set to zero.
//
// For the 'CumDays' time duration calculation type, the following data fields are
// populated in the TimeDurationDto structure:
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).CumDays()
//	    Years                       NOT-populated
//	    YearsNanosecs               NOT-populated
//	    Months                      NOT-populated
//	    MonthsNanosecs              NOT-populated
//	    Weeks                       NOT-populated
//	    WeeksNanosecs               NOT-populated
//	    WeekDays                    NOT-populated
//	    WeekDaysNanosecs            NOT-populated
//	    DateDays                    populated
//	    DateDaysNanosecs            populated
//	    Hours                       populated
//	    HoursNanosecs               populated
//	    Minutes                     populated
//	    MinutesNanosecs             populated
//	    Seconds                     populated
//	    SecondsNanosecs             populated
//	    Milliseconds                populated
//	    MillisecondsNanosecs        populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).CumDays()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumDays())
//
func (TDurCalcType) CumDays() TDurCalcType { return TDurCalcType(3) }

// CumHours - Cumulative Hours calculation. Years, months, weeks, and days are
// ignored. Years, months weeks, days and hours are consolidated as cumulative
// hours.
//
// Years, months, weeks and days duration are not calculated. The entire duration
// is broken down by cumulative hours plus minutes, seconds, milliseconds,
// microseconds and nanoseconds.  Data Fields for years, months, weeks, and days
// are always set to zero.
//
// For the 'CumHours' time duration calculation type, the following data fields
// are populated in the TimeDurationDto structure:
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).CumHours()
//	    Years                       NOT-populated
//	    YearsNanosecs               NOT-populated
//	    Months                      NOT-populated
//	    MonthsNanosecs              NOT-populated
//	    Weeks                       NOT-populated
//	    WeeksNanosecs               NOT-populated
//	    WeekDays                    NOT-populated
//	    WeekDaysNanosecs            NOT-populated
//	    DateDays                    NOT-populated
//	    DateDaysNanosecs            NOT-populated
//	    Hours                       populated
//	    HoursNanosecs               populated
//	    Minutes                     populated
//	    MinutesNanosecs             populated
//	    Seconds                     populated
//	    SecondsNanosecs             populated
//	    Milliseconds                populated
//	    MillisecondsNanosecs        populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).CumHours()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumHours())
//
func (TDurCalcType) CumHours() TDurCalcType { return TDurCalcType(4) }

// CumMinutes - Cumulative Minutes calculation. Years, months, weeks, days and hours
// are ignored. Years, months weeks, days, hours and minutes are consolidated and
// counted as cumulative Minutes.
//
// Years, months, weeks, days and hours duration are not calculated. The entire
// duration is broken down by cumulative minutes plus seconds, milliseconds,
// microseconds and nanoseconds.  Data Fields for years, months, weeks, days
// and hours are always set to zero.
//
// For the 'CumMinutes' time duration calculation type, the following data fields
// are populated:
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).CumMinutes()
//	    Years                       NOT-populated
//	    YearsNanosecs               NOT-populated
//	    Months                      NOT-populated
//	    MonthsNanosecs              NOT-populated
//	    Weeks                       NOT-populated
//	    WeeksNanosecs               NOT-populated
//	    WeekDays                    NOT-populated
//	    WeekDaysNanosecs            NOT-populated
//	    DateDays                    NOT-populated
//	    DateDaysNanosecs            NOT-populated
//	    Hours                       NOT-populated
//	    HoursNanosecs               NOT-populated
//	    Minutes                     populated
//	    MinutesNanosecs             populated
//	    Seconds                     populated
//	    SecondsNanosecs             populated
//	    Milliseconds                populated
//	    MillisecondsNanosecs        populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).CumMinutes()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumMinutes())
//
func (TDurCalcType) CumMinutes() TDurCalcType { return TDurCalcType(5) }

// CumSeconds - Cumulative Seconds calculation. Years, months, weeks, days, hours
// and minutes are ignored. Years, months weeks, days, hours, minutes and seconds
// are consolidated and counted as cumulative Seconds.
//
// Years, months, weeks, days, hours and minutes duration are not calculated. The
// entire duration is broken down by cumulative seconds plus milliseconds, microseconds
// and nanoseconds. Data Fields for years, months, weeks, days hours and minutes are
// always set to zero.
//
// For the 'CumSeconds' time duration calculation type, the following data fields are
// populated:
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).CumSeconds()
//	    Years                       NOT-populated
//	    YearsNanosecs               NOT-populated
//	    Months                      NOT-populated
//	    MonthsNanosecs              NOT-populated
//	    Weeks                       NOT-populated
//	    WeeksNanosecs               NOT-populated
//	    WeekDays                    NOT-populated
//	    WeekDaysNanosecs            NOT-populated
//	    DateDays                    NOT-populated
//	    DateDaysNanosecs            NOT-populated
//	    Hours                       NOT-populated
//	    HoursNanosecs               NOT-populated
//	    Minutes                     NOT-populated
//	    MinutesNanosecs             NOT-populated
//	    Seconds                     populated
//	    SecondsNanosecs             populated
//	    Milliseconds                populated
//	    MillisecondsNanosecs        populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).CumSeconds()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumSeconds())
//
func (TDurCalcType) CumSeconds() TDurCalcType { return TDurCalcType(6) }

// CumMilliseconds - Cumulative Milliseconds calculation. Years, months, weeks, days,
// hours, minutes and seconds are ignored. Years, months, weeks, days, hours, minutes,
// seconds and milliseconds are consolidated and counted as cumulative Milliseconds.
//
// Years, months, weeks, days, hours, minutes and seconds duration are not calculated.
// The entire time duration is broken down by cumulative milliseconds plus microseconds
// and nanoseconds.  Data Fields for years, months, weeks, days, hours, minutes and seconds
// are always set to zero.
//
// For the 'CumMilliseconds' time duration calculation type, the following data fields are populated:
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).CumMilliseconds()
//	    Years                       NOT-populated
//	    YearsNanosecs               NOT-populated
//	    Months                      NOT-populated
//	    MonthsNanosecs              NOT-populated
//	    Weeks                       NOT-populated
//	    WeeksNanosecs               NOT-populated
//	    WeekDays                    NOT-populated
//	    WeekDaysNanosecs            NOT-populated
//	    DateDays                    NOT-populated
//	    DateDaysNanosecs            NOT-populated
//	    Hours                       NOT-populated
//	    HoursNanosecs               NOT-populated
//	    Minutes                     NOT-populated
//	    MinutesNanosecs             NOT-populated
//	    Seconds                     NOT-populated
//	    SecondsNanosecs             NOT-populated
//	    Milliseconds                populated
//	    MillisecondsNanosecs        populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).CumMilliseconds()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumMilliseconds())
//
func (TDurCalcType) CumMilliseconds() TDurCalcType { return TDurCalcType(7) }

// CumMicroseconds - Cumulative Milliseconds calculation. Years, months, weeks, days,
// hours, minutes, seconds and milliseconds are ignored. Years, months, weeks, days,
// hours, minutes, seconds, milliseconds and microseconds are consolidated and counted
// as cumulative Microseconds.
//
// Years, months, weeks, days, hours, minutes, seconds and milliseconds duration are
// not calculated. The entire time duration is broken down by cumulative microseconds
// plus nanoseconds. Data Fields for years, months, weeks, days, hours, minutes, seconds
// and milliseconds are always set to zero.
//
// For the 'CumMicroseconds' time duration calculation type, the following data fields
// are populated:
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).CumMicroseconds()
//	    Years                       NOT-populated
//	    YearsNanosecs               NOT-populated
//	    Months                      NOT-populated
//	    MonthsNanosecs              NOT-populated
//	    Weeks                       NOT-populated
//	    WeeksNanosecs               NOT-populated
//	    WeekDays                    NOT-populated
//	    WeekDaysNanosecs            NOT-populated
//	    DateDays                    NOT-populated
//	    DateDaysNanosecs            NOT-populated
//	    Hours                       NOT-populated
//	    HoursNanosecs               NOT-populated
//	    Minutes                     NOT-populated
//	    MinutesNanosecs             NOT-populated
//	    Seconds                     NOT-populated
//	    SecondsNanosecs             NOT-populated
//	    Milliseconds                NOT-populated
//	    MillisecondsNanosecs        NOT-populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).CumMicroseconds()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumMicroseconds())
//
func (TDurCalcType) CumMicroseconds() TDurCalcType { return TDurCalcType(8) }

// CumNanoSeconds - Cumulative Nanoseconds calculation. Years, months, weeks, days,
// hours, minutes, seconds milliseconds and microseconds are ignored. Years, months,
// weeks, days, hours, minutes, seconds, milliseconds, microseconds and nanoseconds
// are consolidated and counted as cumulative Nanoseconds.
//
// Years, months, weeks, days, hours, minutes, seconds, milliseconds and microseconds
// duration are not calculated. The entire time duration is broken down by cumulative
// nanoseconds. Data Fields for years, months, weeks, days, hours, minutes, seconds,
// milliseconds and microseconds are always set to zero.
//
// For the 'CumNanoseconds' time duration calculation type, the following data fields
// are populated:
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).CumNanoseconds()
//	    Years                       NOT-populated
//	    YearsNanosecs               NOT-populated
//	    Months                      NOT-populated
//	    MonthsNanosecs              NOT-populated
//	    Weeks                       NOT-populated
//	    WeeksNanosecs               NOT-populated
//	    WeekDays                    NOT-populated
//	    WeekDaysNanosecs            NOT-populated
//	    DateDays                    NOT-populated
//	    DateDaysNanosecs            NOT-populated
//	    Hours                       NOT-populated
//	    HoursNanosecs               NOT-populated
//	    Minutes                     NOT-populated
//	    MinutesNanosecs             NOT-populated
//	    Seconds                     NOT-populated
//	    SecondsNanosecs             NOT-populated
//	    Milliseconds                NOT-populated
//	    MillisecondsNanosecs        NOT-populated
//	    Microseconds                NOT-populated
//	    MicrosecondsNanosecs        NOT-populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).CumNanoseconds()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumNanoseconds())
//
func (TDurCalcType) CumNanoseconds() TDurCalcType { return TDurCalcType(9) }

// GregorianYears - Gregorian Year calculation. Allocates time duration by Years,
// Months, Weeks, WeekDays, Date Days, Hours, Minutes, Seconds, Milliseconds,
// Microseconds, Nanoseconds. This calculation type is very similar to 'StdYearMth()'
// described above. However, in this calculation, the Years allocation is performed
// using standard Gregorian Years.
//
// The Gregorian Average Year is equivalent to 365-days, 5-hours, 49-minutes and
// 12-seconds. Therefore Gregorian Year Nanoseconds = 31,556,952,000,000,000 .
// 	Sources:
//		https://en.wikipedia.org/wiki/Year
//		https://en.wikipedia.org/wiki/Gregorian_calendar
//
// For the 'CumNanoseconds' time duration calculation type, the following data fields
// are populated:
//
//
//	type TimeDurationDto struct {
//	    StartTimeDateTz             populated
//	    EndTimeDateTz               populated
//	    TimeDuration                populated
//	    CalcType                    = TDurCalcType(0).GregorianYears()
//	    Years                       populated
//	    YearsNanosecs               populated
//	    Months                      populated
//	    MonthsNanosecs              populated
//	    Weeks                       populated
//	    WeeksNanosecs               populated
//	    WeekDays                    populated
//	    WeekDaysNanosecs            populated
//	    DateDays                    populated
//	    DateDaysNanosecs            populated
//	    Hours                       populated
//	    HoursNanosecs               populated
//	    Minutes                     populated
//	    MinutesNanosecs             populated
//	    Seconds                     populated
//	    SecondsNanosecs             populated
//	    Milliseconds                populated
//	    MillisecondsNanosecs        populated
//	    Microseconds                populated
//	    MicrosecondsNanosecs        populated
//	    Nanoseconds                 populated
//	    TotSubSecNanoseconds        populated
//	    TotDateNanoseconds          populated
//	    TotTimeNanoseconds          populated
//	}
//
// ------------------------------------------------------------------------
//
// Usage
//
//	Example 1: t:= TDurCalcType(0).GregorianYears()
//
//	Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).GregorianYears())
//
func (TDurCalcType) GregorianYears() TDurCalcType { return TDurCalcType(10) }

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TDurCalcType'. This is a standard utility method
// and is not part of the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= TDurCalcType(0).StdYearMth()
//	str := t.String()
//	    str is now equal to 'StdYearMth'
//
func (c TDurCalcType) String() string {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	c.checkInitializeMaps(false)

	result, ok := mTDurCalcTypeIntToString[int(c)]

	if !ok {
		return ""
	}

	return result
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of TDurCalcType is returned set to the value of the
// associated enumeration.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case sensitive and will require an
//	                       exact match. Therefore, 'stdyearmth' will NOT
//	                       match the enumeration name, 'StdYearMth'.
//
//	                       If 'false' a case insensitive search is conducted
//	                       for the enumeration name. In this case, 'stdyearmth'
//	                       will match match enumeration name 'StdYearMth'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TDurCalcType - Upon successful completion, this method will return a new
//	               instance of TDurCalcType set to the value of the enumeration
//	               matched by the string search performed on input parameter,
//	               'valueString'.
//
//	error        - If this method completes successfully, the returned error
//	               Type is set equal to 'nil'. If an error condition is encountered,
//	               this method will return an error Type which encapsulates an
//	               appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t, err := TDurCalcType(0).ParseString("StdYearMth")
//
//	    t is now equal to TDurCalcType(0).StdYearMth()
//
func (c TDurCalcType) ParseString(
	valueString string,
	caseSensitive bool) (TDurCalcType, error) {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	ePrefix := "TDurCalcType.ParseString() "

	c.checkInitializeMaps(false)

	result := TDurCalcType(0)
	if len(valueString) < 3 {
		return result,
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
	}

	var ok bool
	var idx int

	if caseSensitive {

		idx, ok = mTDurCalcTypeStringToInt[valueString]

		if !ok {
			return TDurCalcType(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a TDurCalcType. valueString='%v' ", valueString)
		}

		result = TDurCalcType(idx)

	} else {

		idx, ok = mTDurCalcTypeLwrCaseStringToInt[strings.ToLower(valueString)]

		if !ok {
			return TDurCalcType(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a TDurCalcType. valueString='%v' ", valueString)
		}

		result =
			TDurCalcType(idx)
	}

	return result, nil
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current TDurCalcType
// instance.
//
func (c TDurCalcType) Value() int {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return int(c)
}

// checkInitializeMaps - String and value comparisons performed on enumerations
// supported by this Type, utilizes a series of 3-map types. These maps are used
// internally to perform 'string to value' or 'value to string' look ups on
// enumerations supported by this type. Each time TDurCalcType.String() or
// TDurCalcType.ParseString() a call is made to this method to determine if
// these maps have been initialized. If the maps and look up data have been
// properly initialized and indexed, this method returns without taking action.
//
// On the other hand, if the maps have not yet been initialized, this method will
// initialize all associated map slices.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	reInitialize     bool - If 'true', this will force initialization of
//	                        all associated maps.
//
func (c TDurCalcType) checkInitializeMaps(reInitialize bool) {

	if !reInitialize &&
		mTDurCalcTypeIntToString != nil &&
		len(mTDurCalcTypeIntToString) > 3 &&
		mTDurCalcTypeStringToInt != nil &&
		len(mTDurCalcTypeStringToInt) > 3 &&
		mTDurCalcTypeLwrCaseStringToInt != nil &&
		len(mTDurCalcTypeLwrCaseStringToInt) > 3 {
		return
	}

	var t = TDurCalcType(0).StdYearMth()

	mTDurCalcTypeIntToString = make(map[int]string, 0)
	mTDurCalcTypeStringToInt = make(map[string]int, 0)
	mTDurCalcTypeLwrCaseStringToInt = make(map[string]int, 0)

	s := reflect.TypeOf(t)

	r := reflect.TypeOf(0)
	args := [1]reflect.Value{reflect.Zero(s)}

	for i := 0; i < s.NumMethod(); i++ {

		f := s.Method(i).Name

		if f == "String" ||
			f == "ParseString" ||
			f == "Value" ||
			f == "checkInitializeMaps" {
			continue
		}

		value := s.Method(i).Func.Call(args[:])[0].Convert(r).Int()
		x := int(value)
		mTDurCalcTypeIntToString[x] = f
		mTDurCalcTypeStringToInt[f] = x
		mTDurCalcTypeLwrCaseStringToInt[strings.ToLower(f)] = x
	}

}

// TimeDurationCalcType - Public Global Variable of type 'TDurCalcType'.
// Used to access the 'TDurCalcType' enumeration by dot operator.
//
//   Example:
//      TimeDurationCalcType.StdYearMth()
//
var TimeDurationCalcType TDurCalcType

