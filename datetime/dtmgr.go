package datetime

import (
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
//	t	    time.Time - A date time value to be formatted as a string
//
//	dateTimeFmtStr string - A date time format string which will be used
//				to format and display 'dateTime'. Example:
//				"2006-01-02 15:04:05.000000000 -0700 MST"
//
//				If 'dateTimeFmtStr' is submitted as an
//				'empty string', a default date time format
//				string will be applied. The default date time
//				format string is:
//					FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	string - A formatted date time string
func (dt DtMgr) GetDateTimeCustomFmt(t time.Time, dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		dateTimeFmtStr = FmtDateTimeYrMDayFmtStr
	}

	return t.Format(dateTimeFmtStr)
}

// GetDateTimeEverything - Receives a time value and formats that
// time value as a date time string in the format:
//	EXAMPLE: Saturday April 29, 2017 19:54:30.123456489 -0500 CDT
func (dt DtMgr) GetDateTimeEverything(t time.Time) string {
	return t.Format(FmtDateTimeEverything)
}

// GetDateTimeNanoSecText - Returns formatted date time string with
// nanoseconds.
//	Example:
//		2006-01-02 15:04:05.000000000.
func (dt DtMgr) GetDateTimeNanoSecText(t time.Time) string {
	// Time Format down to the nanosecond
	return t.Format(FmtDateTimeNanoSecondStr)
}

// GetDateTimeSecText - Returns a formatted date time string with
// seconds.
//	Example:
//		2006-01-02 15:04:05.
func (dt DtMgr) GetDateTimeSecText(t time.Time) string {
	// Time Display Format with seconds
	return t.Format(FmtDateTimeSecText)
}

// GetDateTimeStr - Returns a date time string in the format
//		20170427211307
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
//		20170427211307
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
//	EXAMPLE:
//		Monday 2006-01-02 15:04:05.000000000 -0700 MST
func (dt DtMgr) GetDateTimeTzNanoSecDowYMDText(t time.Time) string {
	return t.Format(FmtDateTimeTzNanoDowYMD)
}

// GetDateTimeTzNanoSecText - Receives input parameter 't' as a type 'time.Time'.
// This date time is then returned as a string formatted with the 'FmtDateTimeTzNano'
// format. This format presents date time down to nanoseconds in addition to the
// associated time zone.
//
//	EXAMPLE:
// 		01/02/2006 15:04:05.000000000 -0700 MST
func (dt DtMgr) GetDateTimeTzNanoSecText(t time.Time) string {
	return t.Format(FmtDateTimeTzNano)
}

// GetDateTimeTzNanoSecYMDText - Receives input parameter 't' as a type 'time.Time'.
// This date time is then returned in string format using the FmtDateTimeTzNanoYMD
// format which presents date time down to nanoseconds in addtition to the
// associated time zone. In this format, the date is expressed as
// Year-Month-Day (Example: 2017-12-06)
//
//	EXAMPLE:
//		2006-01-02 15:04:05.000000000 -0700 MST
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
//	EXAMPLE:
//		"2006-01-02 Mon 15:04:05.000000000 -0700 MST"
func (dt DtMgr) GetDateTimeYMDAbbrvDowNano(t time.Time) string {
	return t.Format(FmtDateTimeYMDAbbrvDowNano)
}

// GetDateTimeTzNanoSecYMDDowText - Outputs date time in string format using
// the FmtDateTimeTzNanoYMDDow format which incorporates date time to nano seconds
// and the associated time zone. In this format, the date is expressed as
// Year-Month-Day (Example: 2017-12-06) followed by the day of the week.
// EXAMPLE: 2006-01-02 Monday 15:04:05.000000000 -0700 MST
func (dt DtMgr) GetDateTimeTzNanoSecYMDDowText(t time.Time) string {
	return t.Format(FmtDateTimeTzNanoYMDDow)
}

// GetTimeStampEverything - Generates and returns a time stamp as
// type 'string'. The current time is computed using time.Now() for
// the 'Local' timezone on the host computer. The time stamp is formatted
// using the format, 'FmtDateTimeEverything'.
//
//	EXAMPLE:
//		"Saturday April 29, 2017 19:54:30.123456489 -0500 CDT"
func (dt DtMgr) GetTimeStampEverything() string {
	return time.Now().Local().Format(FmtDateTimeEverything)
}

// GetTimeStampYMDAbbrvDowNano - Generates and returns a time stamp as
// type 'string'. The current time is computed using time.Now() for the
// 'Local' timezone on the host computer. The time stamp is formatted
// using the format 'FmtDateTimeYMDAbbrvDowNano'.
//
//	EXAMPLE:
//		"2006-01-02 Mon 15:04:05.000000000 -0700 MST"
func (dt DtMgr) GetTimeStampYMDAbbrvDowNano() string {

	return time.Now().Local().Format(FmtDateTimeYMDAbbrvDowNano)

}
