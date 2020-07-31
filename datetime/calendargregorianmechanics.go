package datetime

import (
	"math/big"
	"sync"
)


type calendarGregorianMechanics struct {
	lock *sync.Mutex
}

// getCalendarCyclesConfig - returns a CalendarCycleConfiguration instance
// contain all the information necessary for julian day number calculations
// using the Gregorian Calendar
func (calGregMech *calendarGregorianMechanics) getCalendarCyclesConfig() (
	calCyclesCfg CalendarCycleConfiguration) {

	if calGregMech.lock == nil {
		calGregMech.lock = new(sync.Mutex)
	}

	calGregMech.lock.Lock()

	defer calGregMech.lock.Unlock()

	calCycles := CalendarCycleConfiguration{}

	calCycles.mainCycleStartDateForPositiveJDNNo =
		DateTransferDto{
			isLeapYear:          true,
			year:                -4800,
			month:               11,
			day:                 24,
			hour:                12,
			minute:              0,
			second:              0,
			nanosecond:          0,
			tag:                 "Positive JDN No Main Cycle Start",
			isThisInstanceValid: true,
			lock:                new(sync.Mutex),
		}

	calCycles.mainCycleAdjustmentYearsForPositiveJDNNo =
		big.NewInt(-87)

	calCycles.mainCycleAdjustmentDaysForPositiveJDNNo =
		big.NewInt(-31776)

	calCycles.mainCycleStartDateForNegativeJDNNo = DateTransferDto{
		isLeapYear:          true,
		year:                -4400,
		month:               11,
		day:                 24,
		hour:                12,
		minute:              0,
		second:              0,
		nanosecond:          0,
		tag:                 "Negative JDN No Main Cycle Start",
		isThisInstanceValid: true,
		lock:                new(sync.Mutex),
	}

	calCycles.mainCycleAdjustmentYearsForNegativeJDNNo =
		big.NewInt(-313)

	calCycles.mainCycleAdjustmentDaysForNegativeJDNNo =
		big.NewInt(-114320)

	calCycles.jdnBaseStartYearDateTime = DateTransferDto{
		isLeapYear:          false,
		year:                -4713,
		month:               11,
		day:                 24,
		hour:                12,
		minute:              0,
		second:              0,
		nanosecond:          0,
		tag:                 "Julian Day Number Base Start Date/Time",
		isThisInstanceValid: true,
		lock:                new(sync.Mutex),
	}

	calCycles.ordinalFixedDateStartYearDateTime = DateTransferDto{
		isLeapYear:          false,
		year:                1,
		month:               1,
		day:                 1,
		hour:                0,
		minute:              0,
		second:              0,
		nanosecond:          0,
		tag:                 "Gregorian Ordinal Day Number Start Date/Time",
		isThisInstanceValid: true,
		lock:                new(sync.Mutex),
	}

	calCycles.mainCycleConfig =
		CalendarCycleDto{
			yearsInCycle:         big.NewInt(400),
			daysInCycle:          big.NewInt(146097),
			cycleCount:           big.NewInt(0),
			cycleCountTotalDays:  big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:                 new(sync.Mutex),
		}

	calCycles.calendarCyclesConfig = []CalendarCycleDto{
		{ // 0
			yearsInCycle:         big.NewInt(400),
			daysInCycle:          big.NewInt(146097),
			cycleCount:           big.NewInt(0),
			cycleCountTotalDays:  big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:                 new(sync.Mutex),
		},

		{ // 1
			yearsInCycle: big.NewInt(100),
			daysInCycle:  big.NewInt(36524),
			cycleCount:   big.NewInt(0),
			cycleCountTotalDays: big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:         new(sync.Mutex),
		},

		{// 2
			yearsInCycle: big.NewInt(4),
			daysInCycle:  big.NewInt(1461	),
			cycleCount:   big.NewInt(0),
			cycleCountTotalDays: big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:         new(sync.Mutex),
		},

		{// 3
			yearsInCycle: big.NewInt(1),
			daysInCycle:  big.NewInt(365),
			cycleCount:   big.NewInt(0),
			cycleCountTotalDays: big.NewInt(0),
			cycleCountTotalYears: big.NewInt(0),
			remainderYears: big.NewInt(0),
			remainderDays: big.NewInt(0),
			lock:         new(sync.Mutex),
		},
	}

	calCycles.lock = new(sync.Mutex)

	return calCycles
}

// isLeapYear - Returns a boolean value signaling whether the year
// value passed as an input parameter is a leap year (366-days)
// under the Gregorian Calendar.
//
// If the method returns 'true' the input parameter 'year' qualifies
// as a leap year consisting of 366-days. If the method returns 'false',
// the input parameter 'year' is a standard year consisting of 365-days.
//
// Methodology:
//
// In the Gregorian calendar, three criteria must be taken
// into account to identify leap years:
//
// 1. The year must be evenly divisible by 4;
//
// 2. If the year can also be evenly divided by 100, it is not
//    a leap year, unless...
//
//  3. The year is evenly divisible by 100 and the year is also
//     evenly divisible by 400. Then it is a leap year.
//
// According to these rules, the years 2000 and 2400 are leap years,
// while 1800, 1900, 2100, 2200, 2300, and 2500 are not leap years.
//
// For more information on the Gregorian Calendar and leap years,
// reference:
//
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//   https://www.timeanddate.com/date/leapyear.html
//
//
func (calGregMech *calendarGregorianMechanics) isLeapYear(
	year int64) bool {

	if calGregMech.lock == nil {
		calGregMech.lock = new(sync.Mutex)
	}

	calGregMech.lock.Lock()

	defer calGregMech.lock.Unlock()

	var isLeapYear bool

	if year < 0 {
		year *= -1
	}

	if year % int64(4) == 0 {

		isLeapYear = true

		if year % 100 == 0 {

			isLeapYear = false

			if  year % int64(400) == 0 {
				isLeapYear = true
			}
		}

	} else {
		isLeapYear = false
	}

	return isLeapYear
}