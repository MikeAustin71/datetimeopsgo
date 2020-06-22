package datetime

import "sync"

type CalendarDesignation int

var lockCalendarDesignation sync.Mutex


// None - Signals that the CalendarDesignation Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (calDes CalendarDesignation) None() CalendarDesignation {

	lockCalendarDesignation.Lock()

	defer lockCalendarDesignation.Unlock()

	return CalendarDesignation(0)
}

// Gregorian - Signals that the Gregorian Calendar is specified and
// in effect.
//
// The Gregorian calendar is the calendar used in most of the world.
// It is named after Pope Gregory XIII, who introduced it Friday,
// 15 October 1582.
//
//The calendar spaces leap years to make the average year 365.2425
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
// For additional information reference:
//    https://en.wikipedia.org/wiki/Gregorian_calendar
//
// This method is part of the standard enumeration.
//
func (calDes CalendarDesignation) Gregorian() CalendarDesignation {

	lockCalendarDesignation.Lock()

	defer lockCalendarDesignation.Unlock()

	return CalendarDesignation(1)
}

// Julian - - Signals that the Julian Calendar is specified and
// in effect.
//
// The Julian calendar, proposed by Julius Caesar in 708 Ab urbe condita
// (AUC) (46 BC), was a reform of the Roman calendar.[1] It took effect
// on 1 January 709 AUC (45 BC), by edict. It was designed with the aid
// of Greek mathematicians and Greek astronomers such as Sosigenes of
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
func (calDes CalendarDesignation) Julian() CalendarDesignation {

	lockCalendarDesignation.Lock()

	defer lockCalendarDesignation.Unlock()

	return CalendarDesignation(2)
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
// This method is part of the standard enumeration.
//
func (calDes CalendarDesignation) RevisedJulian() CalendarDesignation {

	lockCalendarDesignation.Lock()

	defer lockCalendarDesignation.Unlock()

	return CalendarDesignation(2)
}

func (calDes CalendarDesignation) Goucher() CalendarDesignation {

	lockCalendarDesignation.Lock()

	defer lockCalendarDesignation.Unlock()

	return CalendarDesignation(2)
}
