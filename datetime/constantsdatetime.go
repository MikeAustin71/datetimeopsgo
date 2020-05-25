/*
package datetime provides a variety of types and associated methods designed
to manage dates, times, timezones, date time string formats and time durations.
In addition, associated methods perform date time addition and subtraction.

All types in this package rely on the IANA Time Zone Database.

Reference:

	https://www.iana.org/time-zones
	https://en.wikipedia.org/wiki/Tz_database

*/
package datetime

import (
  "sync"
  "time"
)

/*
 Date Time Constants


 This source file is located in source code repository:
 		'https://github.com/MikeAustin71/datetimeopsgo.git'

 This source code file is located at:
		'MikeAustin71\datetimeopsgo\datetime\constantsdatetime.go'



Overview and General Usage


This source file contains a series of constants useful in managing
date time.

Types of constants defined here include:
  1. Date Time string formats
  2. Common Time conversion constants

*/
const (
  // Date Time Format Constants
  // ================================================================================
  // FmtDateTimeSecondStr - Date Time format used
  // for file names and directory names
  FmtDateTimeSecondStr = "20060102150405"

  // FmtDateTimeNanoSecondStr - Custom Date Time Format
  FmtDateTimeNanoSecondStr = "2006-01-02 15:04:05.000000000"

  // FmtDateTimeSecText - Custom Date Time Format
  FmtDateTimeSecText = "2006-01-02 15:04:05"

  // FmtDateTimeDMYNanoTz - Outputs date time to nano seconds with associated time zone
  FmtDateTimeDMYNanoTz = "01/02/2006 15:04:05.000000000 -0700 MST"

  // FmtDateTimeTzNanoYMD - Outputs date time to nano seconds with Year-Month-Date
  FmtDateTimeTzNanoYMD = "2006-01-02 15:04:05.000000000 -0700 MST"

  // FmtDateTimeYMDHMSTz - Outputs Date time to seconds with Year-Month-Date and Time Zone.
  FmtDateTimeYMDHMSTz = "2006-01-02 15:04:05 -0700 MST"

  // FmtDateTimeTzNanoDowYMD - Output date time to nano seconds with Year-Month-Date
  // prefixed by day of the week
  FmtDateTimeTzNanoDowYMD = "Monday 2006-01-02 15:04:05.000000000 -0700 MST"

  // FmtDateTimeTzNanoYMDDow - Output date time to nano seconds with Year-Month-Date
  // prefixed by day of the week
  FmtDateTimeTzNanoYMDDow = "2006-01-02 Monday 15:04:05.000000000 -0700 MST"

  // FmtDateTimeYMDAbbrvDowNano - Output date time to nano seconds with abbreviated
  // day of week.
  FmtDateTimeYMDAbbrvDowNano = "2006-01-02 Mon 15:04:05.000000000 -0700 MST"

  // FmtDateTimeTzSec - Outputs date time to seconds with associated time zone
  FmtDateTimeTzSec = "01/02/2006 15:04:05 -0700 MST"

  // FmtDateTimeEverything - Custom Date Time Format showing virtually
  // all elements of a date time string.
  FmtDateTimeEverything = "Monday January 2, 2006 15:04:05.000000000 -0700 MST"

  // FmtDateTimeNeutralDateFmt - Neutral Date Format without Time Zone
  FmtDateTimeNeutralDateFmt = "2006-01-02 15:04:05.000000000"

  // FmtDateTimeMDYrFmtStr - Month Day Year Date Format String
  FmtDateTimeMDYrFmtStr = "01/02/2006 15:04:05.000000000 -0700 MST"

  // FmtDateTimeUsMilitaryDate2DYr
  FmtDateTimeUsMilitaryDate2DYr = "021504Z 06"

  // FmtDateTimeYrMDayFmtStr - Year Month Day Date Format String
  FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
)

const (
  // Time Constants
  // ================================================================================
  // Note: A Nanosecond is equal to 1 one-billionth or
  //       1/1,000,000,000 of a second.
  //
  // MilliSecondsPerSecond - Number of Milliseconds in a Second
  MilliSecondsPerSecond = int64(1000)

  // MicroSecondsPerMilliSecond - The Number of Microseconds in
  // a Millisecond.
  MicroSecondsPerMilliSecond = int64(1000)

  // NanoSecondsPerMicroSecond - The number of nanoseconds in
  // a microsecond.
  NanoSecondsPerMicroSecond = int64(1000)

  // MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
  //	 A millisecond is 1/1,000 or 1 one-thousandth of a second
  MilliSecondNanoseconds = int64(time.Millisecond)

  // MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond
  // 	A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
  MicroSecondNanoseconds = int64(time.Microsecond)

  // SecondNanoseconds - Number of Nanoseconds in a Second
  SecondNanoseconds = int64(time.Second)

  HalfSecondNanoseconds = int64(time.Second) / int64(2)

  // MinuteNanoSeconds - Number of Nanoseconds in a minute
  MinuteNanoSeconds = int64(time.Minute)

  // HourNanoSeconds - Number of Nanoseconds in an hour
  HourNanoSeconds = int64(time.Hour)

  // DayNanoSeconds - Number of Nanoseconds in a 24-hour day
  DayNanoSeconds = int64(24) * int64(time.Hour)

  // WeekNanoSeconds - Number of Nanoseconds in a 7-day week
  WeekNanoSeconds = int64(7*24) * int64(time.Hour)

  // NoonNanoSeconds = Nanoseconds from 0-hours to 12-hours
  NoonNanoSeconds = int64(12) * int64(time.Hour)




  // GregorianYearNanoSeconds - Number of Nano Seconds in a
  // Gregorian Year.
  //
  // For the Gregorian calendar the average length of the calendar year
  // (the mean year) across the complete leap cycle of 400 Years is 365.2425 days.
  // The Gregorian Average Year is therefore equivalent to 365 days, 5 hours,
  // 49 minutes and 12 seconds.
  //
  // Sources:
  //	https://en.wikipedia.org/wiki/Year
  //	https://en.wikipedia.org/wiki/Gregorian_calendar
  GregorianYearNanoSeconds = int64(31556952000000000)
)

var lockDefaultDateTimeFormat sync.Mutex

const (
  DEFAULTDATETIMEFORMAT = "2006-01-02 15:04:05.000000000 -0700 MST"
)

