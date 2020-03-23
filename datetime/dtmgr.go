package datetime

import (
	"fmt"
	"strings"
	"time"
)

/*

DtMgr - Date Manager

DtMgr is part of the date time operations library.
The source code repository for this file is located at:

	https://github.com/MikeAustin71/datetimeopsgo.git

This source code file is located at:

	MikeAustin71\datetimeopsgo\datetime\dtmgr.go


Overview and General Usage

The DtMgr Type provides generalized methods for generating
date time strings using a variety of date time formats.

It can also be used as a data transfer object to convey
date time and time duration values.

DtMgr struct

=========================
*/
type DtMgr struct {
	TimeIn     time.Time
	TimeOut    time.Time
	TimeStart  time.Time
	TimeEnd    time.Time
	Duration   time.Duration
	TimeInStr  string
	TimeOutStr string
	TimeFmtStr string
}

// GetDateTimeCustomFmt - Returns a time string formatted according
// to input parameter 'fmt', a time format string.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  t      time.Time - A date time value to be formatted as a string
//
//  dateTimeFmtStr string - A date time format string which will be used
//                          to format and display 'dateTime'. Example:
//                          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                          Date time format constants are found in the source
//                          file 'constantsdatetime.go'. These constants represent
//                          the more commonly used date time string formats. All
//                          Date Time format constants begin with the prefix
//                          'FmtDateTime'.
//
//                          If 'dateTimeFmtStr' is submitted as an
//                          'empty string', a default date time format
//                          string will be applied. The default date time
//                          format string is:
//                          FmtDateTimeYrMDayFmtStr =
//                              "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string - A formatted date time string
//
func (dt DtMgr) GetDateTimeCustomFmt(t time.Time, dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		dateTimeFmtStr = FmtDateTimeYrMDayFmtStr
	}

	return t.Format(dateTimeFmtStr)
}

// GetDateTimeEverything - Receives a time value and formats that
// time value as a date time string in the format:
//   EXAMPLE: Saturday April 29, 2017 19:54:30.123456489 -0500 CDT
func (dt DtMgr) GetDateTimeEverything(t time.Time) string {
	return t.Format(FmtDateTimeEverything)
}

// GetDateTimeNanoSecText - Returns formatted date time string with
// nanoseconds.
//
//  EXAMPLE: "2006-01-02 15:04:05.000000000"
//
func (dt DtMgr) GetDateTimeNanoSecText(t time.Time) string {
	// Time Format down to the nanosecond
	return t.Format(FmtDateTimeNanoSecondStr)
}

// GetDateTimeSecText - Returns a formatted date time string with
// seconds.
//
//  EXAMPLE: "2006-01-02 15:04:05"
//
func (dt DtMgr) GetDateTimeSecText(t time.Time) string {
	// Time Display Format with seconds
	return t.Format(FmtDateTimeSecText)
}

// GetDateTimeStr - Returns a date time string in the format
//
//    "20170427211307"
//
// The string contains year, month, day, hour, minute and
// seconds.
func (dt DtMgr) GetDateTimeStr(t time.Time) string {

	// Time Format down to the second
	return t.Format(FmtDateTimeSecondStr)

}

// GetDateTimeStrNowLocal - Gets current
// local time and formats it as a date time
// string. The current time is formatted as:
//
//    "20170427211307"
//
// This time string contains year, month, day,
// hour, minute and seconds.
func (dt DtMgr) GetDateTimeStrNowLocal() string {

	return dt.GetDateTimeStr(time.Now().Local())

}

// GetDateTimeTzNanoSecDowYMDText - Returns input parameter 't' (time.Time)
// as a date time in string format using the FmtDateTimeTzNanoDowYMD format
// which displays date time to the nanosecond along with the associated time
// zone. In this format, the date is expressed as Year-Month-Day (Example: 2017-12-06).
// The string is prefixed with the day of the week.
//
//  EXAMPLE: "Monday 2006-01-02 15:04:05.000000000 -0700 MST"
//
func (dt DtMgr) GetDateTimeTzNanoSecDowYMDText(t time.Time) string {
	return t.Format(FmtDateTimeTzNanoDowYMD)
}

// GetDateTimeTzNanoSecText - Receives input parameter 't' as a type 'time.Time'.
// This date time is then returned as a string formatted with the 'FmtDateTimeDMYNanoTz'
// format. This format presents date time down to nanoseconds in addition to the
// associated time zone.
//
//  EXAMPLE: "01/02/2006 15:04:05.000000000 -0700 MST"
//
func (dt DtMgr) GetDateTimeTzNanoSecText(t time.Time) string {
	return t.Format(FmtDateTimeDMYNanoTz)
}

// GetDateTimeTzNanoSecYMDText - Receives input parameter 't' as a type 'time.Time'.
// This date time is then returned in string format using the FmtDateTimeTzNanoYMD
// format which presents date time down to nanoseconds in addition to the
// associated time zone. In this format, the date is expressed as
// Year-Month-Day (Example: 2017-12-06)
//
//  EXAMPLE: "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (dt DtMgr) GetDateTimeTzNanoSecYMDText(t time.Time) string {
	return t.Format(FmtDateTimeTzNanoYMD)
}

// GetDateTimeYMDAbbrvDowNano -Receives input parameter 't' as a type 'time.Time'.
// This date time is then returned in string format using the FmtDateTimeYMDAbbrvDowNano
// format which presents date time down to the nanosecond in addition to the
// associated time zone. In this format, the date is expressed as Year-Month-Day
// (Example: 2017-12-06). The string includes the abbreviated (limited to 3-characters)
// day of the week.
//
//  EXAMPLE: "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
//
func (dt DtMgr) GetDateTimeYMDAbbrvDowNano(t time.Time) string {
	return t.Format(FmtDateTimeYMDAbbrvDowNano)
}

// GetDateTimeTzNanoSecYMDDowText - Outputs date time in string format using
// the FmtDateTimeTzNanoYMDDow format which incorporates date time to nano seconds
// and the associated time zone. In this format, the date is expressed as
// Year-Month-Day (Example: 2017-12-06) followed by the day of the week.
//
//  EXAMPLE: "2006-01-02 Monday 15:04:05.000000000 -0700 MST"
//
func (dt DtMgr) GetDateTimeTzNanoSecYMDDowText(t time.Time) string {
	return t.Format(FmtDateTimeTzNanoYMDDow)
}

// GetMilitaryCompactDateTimeGroup - Outputs date time string formatted for
// standard U.S.A. Military date time also referred to as the Military
// Date Time Group (DTG). This form of the Date Time Group is configured
// as the 'Compact' Date Time Group because there are no spaces between the
// date time elements.
//
// The Military time zone is computed from the UTC offset associated
// with the input parameter 't' (time.Time).
//
// Reference:
//    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//    https://www.timeanddate.com/time/zones/z
//    http://blog.refactortactical.com/blog/military-date-time-group/
//
// Military 2-digit year format or  Date Time Group is traditionally
// formatted as DDHHMM(Z)MONYY, where 'Z' is the Military Time Zone.
//
// EXAMPLES:
//
//    "011815ZJAN11" = 01/01/2011 18:15 +0000 Zulu
//
//     630pm on January 6th, 2012 in Fayetteville NC would read '061830RJAN12'
//
// Note: For more granular management of Military Date Times, use the
// MilitaryDateTzDto type located in source file 'militarydatetzdto.go'.
//
func (dt DtMgr) GetMilitaryCompactDateTimeGroup(t time.Time) (fmtDateTime string, err error) {

	ePrefix := "DtMgr.GetMilitaryCompactDateTimeGroup() "

	fmtDateTime = ""
	err = nil

	// FmtDateTimeYMDHMSTz= "2006-01-02 15:04:05 -0700 MST"
	tempStr := t.Format(FmtDateTimeYMDHMSTz)

	dateElementArray := strings.Split(tempStr," ")

	if len(dateElementArray) != 4 {
		err = fmt.Errorf(ePrefix +
			"Error: Expected Date Element Array length='4'.\n" +
			"Instead, Date Element Array length='%v'.", len(dateElementArray))
	}

	utcOffset := dateElementArray[2][0:3] + "00"

	milTzDat := MilitaryTimeZoneData{}

	militaryTzName, ok := milTzDat.UtcOffsetToMilitaryTimeZone(utcOffset)

	if !ok {
		err = fmt.Errorf(ePrefix +
			"Error: Could not find Military Time Zone for UTC Offset %v.", utcOffset)
		return fmtDateTime, err
	}

	militaryTimeZoneLetter := militaryTzName[0:1]

	// Military 2-digit year format: DDHHMMZMONYY
	// 630pm on January 6th, 2012 in Fayetteville NC = 061830RJAN12

	// FmtDateTimeYMDHMSTz= "2006-01-02 15:04:05 -0700 MST

	fmtDateTime = t.Format("021504" + militaryTimeZoneLetter + "Jan06")

	fmtDateTime = strings.ToUpper(fmtDateTime)

		return fmtDateTime, err
	}

// GetMilitaryOpenDateTimeGroup - Outputs date time string formatted for
// standard U.S.A. Military date time also referred to as the Military
// Date Time Group (DTG). This form of the Date Time Group is configured
// as the 'Open' or easy to read Date Time Group because spaces are inserted
// between the date time elements.
//
// The Military time zone is computed from the UTC offset associated
// with the input parameter 't' (time.Time).
//
// Reference:
//    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//    https://www.timeanddate.com/time/zones/z
//    http://blog.refactortactical.com/blog/military-date-time-group/
//
// Military 2-digit year format or  Date Time Group is traditionally
// formatted as DD HHMM(Z) MON YY, where 'Z' is the Military Time Zone.
//
// EXAMPLES:
//
//    "01 1815Z JAN 11" = 01/01/2011 18:15 +0000 Zulu
//
//     630pm on January 6th, 2012 in Fayetteville NC would read '06 1830R JAN 12'
//
// Note: For more granular management of Military Date Times, use the
// MilitaryDateTzDto type located in source file 'militarydatetzdto.go'.
//
func (dt DtMgr) GetMilitaryOpenDateTimeGroup(t time.Time) (fmtDateTime string, err error) {

	ePrefix := "DtMgr.GetMilitaryCompactDateTimeGroup() "

	fmtDateTime = ""
	err = nil

	// FmtDateTimeYMDHMSTz= "2006-01-02 15:04:05 -0700 MST"
	tempStr := t.Format(FmtDateTimeYMDHMSTz)

	dateElementArray := strings.Split(tempStr," ")

	if len(dateElementArray) != 4 {
		err = fmt.Errorf(ePrefix +
			"Error: Expected Date Element Array length='4'.\n" +
			"Instead, Date Element Array length='%v'.", len(dateElementArray))
	}

	utcOffset := dateElementArray[2][0:3] + "00"

	milTzDat := MilitaryTimeZoneData{}

	militaryTzName, ok := milTzDat.UtcOffsetToMilitaryTimeZone(utcOffset)

	if !ok {
		err = fmt.Errorf(ePrefix +
			"Error: Could not find Military Time Zone for UTC Offset %v.", utcOffset)
		return fmtDateTime, err
	}

	militaryTimeZoneLetter := militaryTzName[0:1]

	// Military 2-digit year format: DD HHMMZ MON YY
	// 630pm on January 6th, 2012 in Fayetteville NC = 06 1830R JAN 12

	// FmtDateTimeYMDHMSTz= "2006-01-02 15:04:05 -0700 MST

	fmtDateTime = t.Format("02 1504" + militaryTimeZoneLetter + " Jan 06")

	fmtDateTime = strings.ToUpper(fmtDateTime)

		return fmtDateTime, err
	}



// GetTimeStampEverything - Generates and returns a time stamp as
// type 'string'. The current time is computed using time.Now() for
// the 'Local' timezone on the host computer. The time stamp is formatted
// using the format, 'FmtDateTimeEverything'.
//
//  EXAMPLE: "Saturday April 29, 2017 19:54:30.123456489 -0500 CDT"
//
func (dt DtMgr) GetTimeStampEverything() string {
	return time.Now().Local().Format(FmtDateTimeEverything)
}

// GetTimeStampYMDAbbrvDowNano - Generates and returns a time stamp as
// type 'string'. The current time is computed using time.Now() for the
// 'Local' timezone on the host computer. The time stamp is formatted
// using the format 'FmtDateTimeYMDAbbrvDowNano'.
//
//  EXAMPLE: "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
//
func (dt DtMgr) GetTimeStampYMDAbbrvDowNano() string {

	return time.Now().Local().Format(FmtDateTimeYMDAbbrvDowNano)

}
