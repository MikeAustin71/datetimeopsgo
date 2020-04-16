package datetime

import (
	"fmt"
	"strings"
	"sync"
)

var mTDurCalcTypeStringToCode = map[string]TDurCalcType{
	"None"           : TDurCalcType(0),
	"StdYearMth"     : TDurCalcType(1),
	"CumMonths"      : TDurCalcType(2),
	"CumWeeks"       : TDurCalcType(3),
	"CumDays"        : TDurCalcType(4),
	"CumHours"       : TDurCalcType(5),
	"CumMinutes"     : TDurCalcType(6),
	"CumSeconds"     : TDurCalcType(7),
	"CumMilliseconds": TDurCalcType(8),
	"CumMicroseconds": TDurCalcType(9),
	"CumNanoseconds" : TDurCalcType(10),
	"GregorianYears" : TDurCalcType(11),
}

var mTDurCalcTypeLwrCaseStringToCode = map[string]TDurCalcType{
	"none"           :  TDurCalcType(0),
	"stdyearmth"     :  TDurCalcType(1),
	"cummonths"      :  TDurCalcType(2),
	"cumweeks"       :  TDurCalcType(3),
	"cumdays"        :  TDurCalcType(4),
	"cumhours"       :  TDurCalcType(5),
	"cumminutes"     :  TDurCalcType(6),
	"cumseconds"     :  TDurCalcType(7),
	"cummilliseconds":  TDurCalcType(8),
	"cummicroseconds":  TDurCalcType(9),
	"cumnanoseconds" :  TDurCalcType(10),
	"gregorianyears" :  TDurCalcType(11),
}

var mTDurCalcTypeCodeToString = map[TDurCalcType]string{
	 TDurCalcType(0)  : "None",
	 TDurCalcType(1)  : "StdYearMth",
	 TDurCalcType(2)  : "CumMonths",
	 TDurCalcType(3)  : "CumWeeks",
	 TDurCalcType(4)  : "CumDays",
	 TDurCalcType(5)  : "CumHours",
	 TDurCalcType(6)  : "CumMinutes",
	 TDurCalcType(7)  : "CumSeconds",
	 TDurCalcType(8)  : "CumMilliseconds",
	 TDurCalcType(9)  : "CumMicroseconds",
	 TDurCalcType(10) : "CumNanoseconds",
	 TDurCalcType(11) : "GregorianYears",
}

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
// None             (0) - None - Signals that Time Duration Calculation Type
//                        is not initialized. This is an error condition.
//
// StdYearMth       (1) - StdYearMth Calculation type. Allocates time duration by
//                        Standard Year, months, weeks, days, hours, minutes, seconds,
//                        milliseconds, microseconds, and nanoseconds.
//
// CumMonths        (2) - Cumulative Months Calculation. Allocates time duration by
//                       cumulative months plus weeks, week days, date days, hours,
//                       minutes, seconds, milliseconds, microseconds and nanoseconds.
//
// CumWeeks         (3) - Cumulative Weeks Calculation. Allocates time duration by cumulative
//                        weeks, plus week days, hours, minutes, seconds, milliseconds,
//                        microseconds and nanoseconds.
//
// CumDays          (4) - Cumulative Days calculation. Allocates time duration by cumulative
//                        days plus hours, minutes, seconds, milliseconds, microseconds and
//                        nanoseconds.
//
// CumHours         (5) - Cumulative Hours calculation. Allocates time duration by cumulative
//                        hours plus minutes, seconds, milliseconds, microseconds and
//                        nanoseconds.
//
// CumMinutes       (6) - Cumulative Minutes calculation. Allocates time duration by
//                        cumulative minutes plus seconds, milliseconds, microseconds and
//                        nanoseconds.
//
// CumSeconds       (7) - Cumulative Seconds calculation. Allocates time duration by
//                        cumulative seconds plus milliseconds, microseconds and
//                        nanoseconds.
//
// CumMilliseconds  (8) - Cumulative Milliseconds calculation. Allocates time duration
//                        by cumulative milliseconds plus microseconds and nanoseconds.
//
// CumMicroseconds  (9) - Cumulative Microseconds calculation. Allocates time duration
//                        by cumulative microseconds plus nanoseconds.
//
// CumNanoseconds  (10) - Cumulative Nanoseconds calculation. Allocates time duration
//                        by cumulative nanoseconds.
//
// GregorianYears  (11) - Gregorian Year calculation. Allocates time duration by Years,
//                        Months, Weeks, WeekDays, Date Days, Hours, Minutes, Seconds,
//                        Milliseconds, Microseconds, Nanoseconds. However, the Years
//                        allocation is performed using standard Gregorian Years.
//
// For easy access to these enumeration values, use the global variable
// 'TDurCalc'. Example: TDurCalc.CumNanoseconds()
//
// Otherwise you will need to use the formal syntax:
// TDurCalcType(0).CumNanoseconds()
//
type TDurCalcType int

var lockTDurCalcType sync.Mutex


// None - Signals that the Time Duration Calculation Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (durCalc TDurCalcType) None() TDurCalcType {
	
	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(0)
}

// StdYearMth - Allocates time duration by Standard Year, months, weeks,
// weekdays, date days, hours, minutes, seconds, milliseconds, microseconds
// and nanoseconds. For the 'StdYearMth' time duration calculation type, all
// data fields in the TimeDurationDto	structure are populated.
//
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).StdYearMth()
//     Years                       populated
//     YearsNanosecs               populated
//     Months                      populated
//     MonthsNanosecs              populated
//     Weeks                       populated
//     WeeksNanosecs               populated
//     WeekDays                    populated
//     WeekDaysNanosecs            populated
//     DateDays                    populated
//     DateDaysNanosecs            populated
//     Hours                       populated
//     HoursNanosecs               populated
//     Minutes                     populated
//     MinutesNanosecs             populated
//     Seconds                     populated
//     SecondsNanosecs             populated
//     Milliseconds                populated
//     MillisecondsNanosecs        populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).StdYearMth()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).StdYearMth())
//
//
func (durCalc TDurCalcType) StdYearMth() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(1)
}

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
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).CumMonths()
//     Years                       NOT-populated
//     YearsNanosecs               NOT-populated
//     Months                      populated
//     MonthsNanosecs              populated
//     Weeks                       populated
//     WeeksNanosecs               populated
//     WeekDays                    populated
//     WeekDaysNanosecs            populated
//     DateDays                    populated
//     DateDaysNanosecs            populated
//     Hours                       populated
//     HoursNanosecs               populated
//     Minutes                     populated
//     MinutesNanosecs             populated
//     Seconds                     populated
//     SecondsNanosecs             populated
//     Milliseconds                populated
//     MillisecondsNanosecs        populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).CumMonths()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumMonths())
//
func (durCalc TDurCalcType) CumMonths() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(2)
}

// CumWeeks - Cumulative Weeks calculation. Years and months are ignored. Instead,
// years, months, weeks are consolidated and counted as cumulative weeks. Years and
// months duration are not calculated. The entire duration is broken down by cumulative
// weeks, plus week days, hours, minutes, seconds, milliseconds, microseconds and
// nanoseconds. Data Fields for Years and Months are always set to zero.
//
// For the 'CumWeeks' time duration calculation type, the following data fields are
// populated in the TimeDurationDto structure:
//
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).CumWeeks()
//     Years                       NOT-populated
//     YearsNanosecs               NOT-populated
//     Months                      NOT-populated
//     MonthsNanosecs              NOT-populated
//     Weeks                       populated
//     WeeksNanosecs               populated
//     WeekDays                    populated
//     WeekDaysNanosecs            populated
//     DateDays                    NOT-populated
//     DateDaysNanosecs            NOT-populated
//     Hours                       populated
//     HoursNanosecs               populated
//     Minutes                     populated
//     MinutesNanosecs             populated
//     Seconds                     populated
//     SecondsNanosecs             populated
//     Milliseconds                populated
//     MillisecondsNanosecs        populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).CumWeeks()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumWeeks())
//
func (durCalc TDurCalcType) CumWeeks() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(3)
}

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
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).CumDays()
//     Years                       NOT-populated
//     YearsNanosecs               NOT-populated
//     Months                      NOT-populated
//     MonthsNanosecs              NOT-populated
//     Weeks                       NOT-populated
//     WeeksNanosecs               NOT-populated
//     WeekDays                    NOT-populated
//     WeekDaysNanosecs            NOT-populated
//     DateDays                    populated
//     DateDaysNanosecs            populated
//     Hours                       populated
//     HoursNanosecs               populated
//     Minutes                     populated
//     MinutesNanosecs             populated
//     Seconds                     populated
//     SecondsNanosecs             populated
//     Milliseconds                populated
//     MillisecondsNanosecs        populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).CumDays()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumDays())
//
func (durCalc TDurCalcType) CumDays() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(4)
}

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
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).CumHours()
//     Years                       NOT-populated
//     YearsNanosecs               NOT-populated
//     Months                      NOT-populated
//     MonthsNanosecs              NOT-populated
//     Weeks                       NOT-populated
//     WeeksNanosecs               NOT-populated
//     WeekDays                    NOT-populated
//     WeekDaysNanosecs            NOT-populated
//     DateDays                    NOT-populated
//     DateDaysNanosecs            NOT-populated
//     Hours                       populated
//     HoursNanosecs               populated
//     Minutes                     populated
//     MinutesNanosecs             populated
//     Seconds                     populated
//     SecondsNanosecs             populated
//     Milliseconds                populated
//     MillisecondsNanosecs        populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).CumHours()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumHours())
//
func (durCalc TDurCalcType) CumHours() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(5)
}

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
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).CumMinutes()
//     Years                       NOT-populated
//     YearsNanosecs               NOT-populated
//     Months                      NOT-populated
//     MonthsNanosecs              NOT-populated
//     Weeks                       NOT-populated
//     WeeksNanosecs               NOT-populated
//     WeekDays                    NOT-populated
//     WeekDaysNanosecs            NOT-populated
//     DateDays                    NOT-populated
//     DateDaysNanosecs            NOT-populated
//     Hours                       NOT-populated
//     HoursNanosecs               NOT-populated
//     Minutes                     populated
//     MinutesNanosecs             populated
//     Seconds                     populated
//     SecondsNanosecs             populated
//     Milliseconds                populated
//     MillisecondsNanosecs        populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).CumMinutes()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumMinutes())
//
func (durCalc TDurCalcType) CumMinutes() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(6)
}

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
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).CumSeconds()
//     Years                       NOT-populated
//     YearsNanosecs               NOT-populated
//     Months                      NOT-populated
//     MonthsNanosecs              NOT-populated
//     Weeks                       NOT-populated
//     WeeksNanosecs               NOT-populated
//     WeekDays                    NOT-populated
//     WeekDaysNanosecs            NOT-populated
//     DateDays                    NOT-populated
//     DateDaysNanosecs            NOT-populated
//     Hours                       NOT-populated
//     HoursNanosecs               NOT-populated
//     Minutes                     NOT-populated
//     MinutesNanosecs             NOT-populated
//     Seconds                     populated
//     SecondsNanosecs             populated
//     Milliseconds                populated
//     MillisecondsNanosecs        populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).CumSeconds()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumSeconds())
//
func (durCalc TDurCalcType) CumSeconds() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(7)
}

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
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).CumMilliseconds()
//     Years                       NOT-populated
//     YearsNanosecs               NOT-populated
//     Months                      NOT-populated
//     MonthsNanosecs              NOT-populated
//     Weeks                       NOT-populated
//     WeeksNanosecs               NOT-populated
//     WeekDays                    NOT-populated
//     WeekDaysNanosecs            NOT-populated
//     DateDays                    NOT-populated
//     DateDaysNanosecs            NOT-populated
//     Hours                       NOT-populated
//     HoursNanosecs               NOT-populated
//     Minutes                     NOT-populated
//     MinutesNanosecs             NOT-populated
//     Seconds                     NOT-populated
//     SecondsNanosecs             NOT-populated
//     Milliseconds                populated
//     MillisecondsNanosecs        populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).CumMilliseconds()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumMilliseconds())
//
func (durCalc TDurCalcType) CumMilliseconds() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(8)
}

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
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).CumMicroseconds()
//     Years                       NOT-populated
//     YearsNanosecs               NOT-populated
//     Months                      NOT-populated
//     MonthsNanosecs              NOT-populated
//     Weeks                       NOT-populated
//     WeeksNanosecs               NOT-populated
//     WeekDays                    NOT-populated
//     WeekDaysNanosecs            NOT-populated
//     DateDays                    NOT-populated
//     DateDaysNanosecs            NOT-populated
//     Hours                       NOT-populated
//     HoursNanosecs               NOT-populated
//     Minutes                     NOT-populated
//     MinutesNanosecs             NOT-populated
//     Seconds                     NOT-populated
//     SecondsNanosecs             NOT-populated
//     Milliseconds                NOT-populated
//     MillisecondsNanosecs        NOT-populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).CumMicroseconds()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumMicroseconds())
//
func (durCalc TDurCalcType) CumMicroseconds() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(9)
}

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
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).CumNanoseconds()
//     Years                       NOT-populated
//     YearsNanosecs               NOT-populated
//     Months                      NOT-populated
//     MonthsNanosecs              NOT-populated
//     Weeks                       NOT-populated
//     WeeksNanosecs               NOT-populated
//     WeekDays                    NOT-populated
//     WeekDaysNanosecs            NOT-populated
//     DateDays                    NOT-populated
//     DateDaysNanosecs            NOT-populated
//     Hours                       NOT-populated
//     HoursNanosecs               NOT-populated
//     Minutes                     NOT-populated
//     MinutesNanosecs             NOT-populated
//     Seconds                     NOT-populated
//     SecondsNanosecs             NOT-populated
//     Milliseconds                NOT-populated
//     MillisecondsNanosecs        NOT-populated
//     Microseconds                NOT-populated
//     MicrosecondsNanosecs        NOT-populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).CumNanoseconds()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).CumNanoseconds())
//
func (durCalc TDurCalcType) CumNanoseconds() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(10)
}

// GregorianYears - Gregorian Year calculation. Allocates time duration by Years,
// Months, Weeks, WeekDays, Date Days, Hours, Minutes, Seconds, Milliseconds,
// Microseconds, Nanoseconds. This calculation type is very similar to 'StdYearMth()'
// described above. However, in this calculation, the Years allocation is performed
// using standard Gregorian Years.
//
// The Gregorian Average Year is equivalent to 365-days, 5-hours, 49-minutes and
// 12-seconds. Therefore Gregorian Year Nanoseconds = 31,556,952,000,000,000 .
//  Sources:
//  https://en.wikipedia.org/wiki/Year
//  https://en.wikipedia.org/wiki/Gregorian_calendar
//
// For the 'GregorianYears' time duration calculation type, the following data fields
// are populated:
//
//
// type TimeDurationDto struct {
//     startDateTimeTz             populated
//     endDateTimeTz               populated
//     timeDuration                populated
//     calcType                    = TDurCalcType(0).GregorianYears()
//     Years                       populated
//     YearsNanosecs               populated
//     Months                      populated
//     MonthsNanosecs              populated
//     Weeks                       populated
//     WeeksNanosecs               populated
//     WeekDays                    populated
//     WeekDaysNanosecs            populated
//     DateDays                    populated
//     DateDaysNanosecs            populated
//     Hours                       populated
//     HoursNanosecs               populated
//     Minutes                     populated
//     MinutesNanosecs             populated
//     Seconds                     populated
//     SecondsNanosecs             populated
//     Milliseconds                populated
//     MillisecondsNanosecs        populated
//     Microseconds                populated
//     MicrosecondsNanosecs        populated
//     Nanoseconds                 populated
//     TotSubSecNanoseconds        populated
//     TotDateNanoseconds          populated
//     TotTimeNanoseconds          populated
// }
//
// This method is part of the standard enumeration.
//
// ------------------------------------------------------------------------
//
// Usage
//
// Example 1: t:= TDurCalcType(0).GregorianYears()
//
// Example 2: SomeFuncRequiringTimeDurCalcType(TDurCalcType(0).GregorianYears())
//
func (durCalc TDurCalcType) GregorianYears() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(11)
}

// ===============================================================================
//                     UTILITY METHODS
// ===============================================================================

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TDurCalcType'. 
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
func (durCalc TDurCalcType) String() string {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	result, ok := mTDurCalcTypeCodeToString[durCalc]

	if !ok {
		return ""
	}

	return result
}

// XFirstValidCalcType - Returns the value of the First TDurCalcType.
// The first TDurCalcType is 'StdYearMth'.
//
func (durCalc TDurCalcType) XFirstValidCalcType() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(1)
}


// XLastValidCalcType - Returns the value of the Last TDurCalcType.
// The Last TDurCalcType Value is 'GregorianYears'.
//
func (durCalc TDurCalcType) XLastValidCalcType() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return TDurCalcType(11)
}


// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of TDurCalcType is returned set to the value of the
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
//                        exact match. Therefore, 'stdyearmth' will NOT
//                        match the enumeration name, 'StdYearMth'.
//
//                        If 'false' a case insensitive search is conducted
//                        for the enumeration name. In this case, 'stdyearmth'
//                        will match match enumeration name 'StdYearMth'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// TDurCalcType - Upon successful completion, this method will return a new
//                instance of TDurCalcType set to the value of the enumeration
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
// t, err := TDurCalcType(0).XParseString("StdYearMth", true)
//
//     t is now equal to TDurCalcType(0).StdYearMth()
//
func (durCalc TDurCalcType) XParseString(
	valueString string,
	caseSensitive bool) (TDurCalcType, error) {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	ePrefix := "TDurCalcType.XParseString() "

	if len(valueString) < 4 {
		return TDurCalcType(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n" +
				"String length is less than '4'.\n" +
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var timeDurCalcType TDurCalcType

	if caseSensitive {

		timeDurCalcType, ok = mTDurCalcTypeStringToCode[valueString]

		if !ok {
			return TDurCalcType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a TDurCalcType.\n" +
					"valueString='%v'\n", valueString)
		}

	} else {

		timeDurCalcType, ok = mTDurCalcTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return TDurCalcType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a TDurCalcType.\n" +
					"valueString='%v'\n", valueString)
		}
	}

	return timeDurCalcType, nil
}

// XValue - This method returns the enumeration value of the current TDurCalcType
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (durCalc TDurCalcType) XValue() TDurCalcType {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return durCalc
}

// XValueInt - This method returns the integer value of the current TDurCalcType
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (durCalc TDurCalcType) XValueInt() int {

	lockTDurCalcType.Lock()

	defer lockTDurCalcType.Unlock()

	return int(durCalc)
}

// TDurCalc - public global variable of
// type TDurCalcType.
//
// This variable serves as an easier, short hand
// technique for accessing TDurCalcType values.
//
// Usage:
// TDurCalc.None(),
// TDurCalc.StdYearMth(),
// TDurCalc.CumMonths(),
// TDurCalc.CumWeeks(),
// TDurCalc.CumDays(),
// TDurCalc.CumHours(),
// TDurCalc.CumMinutes(),
// TDurCalc.CumSeconds(),
// TDurCalc.CumMilliseconds(),
// TDurCalc.CumMicroseconds(),
// TDurCalc.CumNanoseconds(),
// TDurCalc.GregorianYears(),
//
var TDurCalc TDurCalcType

