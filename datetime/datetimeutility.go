package datetime

import (
	"time"
)

/*
datetimeutility.go is part of the date time operations library. The source code repository
 	for this file is located at:

					https://github.com/MikeAustin71/datetimeopsgo.git

 */


// DateTimeUtility - struct used to export
// Date Time Management methods.
type DateTimeUtility struct {
	TimeIn     time.Time
	TimeOut    time.Time
	TimeStart  time.Time
	TimeEnd    time.Time
	Duration   time.Duration
	TimeInStr  string
	TimeOutStr string
	TimeFmtStr string
}

// GetDateTimeStrNowLocal - Gets current
// local time and formats it as a date time
// string
func (dt DateTimeUtility) GetDateTimeStrNowLocal() string {

	return dt.GetDateTimeStr(time.Now().Local())

}

// GetDateTimeStr - Returns a date time string
// in the format 20170427211307
func (dt DateTimeUtility) GetDateTimeStr(t time.Time) string {

	// Time Format down to the second
	return t.Format(FmtDateTimeSecondStr)

}

// GetDateTimeSecText - Returns formatted
// date time with seconds for display,
// 2006-01-02 15:04:05.
func (dt DateTimeUtility) GetDateTimeSecText(t time.Time) string {
	// Time Display Format with seconds
	return t.Format(FmtDateTimeSecText)
}

// GetDateTimeNanoSecText - Returns formated
// date time string with nanoseconds
// 2006-01-02 15:04:05.000000000.
func (dt DateTimeUtility) GetDateTimeNanoSecText(t time.Time) string {
	// Time Format down to the nanosecond
	return t.Format(FmtDateTimeNanoSecondStr)
}

// GetDateTimeTzNanoSecText - Outputs date time in string format using
// the FmtDateTimeTzNano format which incorporates date time to nano seconds
// and the associated time zone.
// EXAMPLE: 01/02/2006 15:04:05.000000000 -0700 MST
func (dt DateTimeUtility) GetDateTimeTzNanoSecText(t time.Time) string {
	return t.Format(FmtDateTimeTzNano)
}

// GetDateTimeTzNanoSecYMDText - Outputs date time in string format using
// the FmtDateTimeTzNanoYMD format which incorporates date time to nano seconds
// and the associated time zone. In this format, the date is expressed as
// Year-Month-Day (Example: 2017-12-06)
// EXAMPLE: 2006-01-02 15:04:05.000000000 -0700 MST
func (dt DateTimeUtility) GetDateTimeTzNanoSecYMDText(t time.Time) string {
	return t.Format(FmtDateTimeTzNanoYMD)
}

// GetDateTimeTzNanoSecDowYMDText - Outputs date time in string format using
// the FmtDateTimeTzNanoDowYMD format which incorporates date time to the
// nano second and the associated time zone. In this format, the date is
// expressed as Year-Month-Day (Example: 2017-12-06). The string is
// prefixed with the day of the week:
// EXAMPLE: Monday 2006-01-02 15:04:05.000000000 -0700 MST
func (dt DateTimeUtility) GetDateTimeTzNanoSecDowYMDText(t time.Time) string {
	return t.Format(FmtDateTimeTzNanoDowYMD)
}

// GetDateTimeYMDAbbrvDowNano - Outputs date time in string format using
// the FmtDateTimeYMDAbbrvDowNano format which incorporates date time to the
// nano second and the associated time zone. In this format, the date is
// expressed as Year-Month-Day (Example: 2017-12-06). The string includes
// the abbreviated (limited to 3-characters) day of the week:
// EXAMPLE: "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
func (dt DateTimeUtility) GetDateTimeYMDAbbrvDowNano(t time.Time) string {
	return t.Format(FmtDateTimeYMDAbbrvDowNano)
}

// GetDateTimeTzNanoSecYMDDowText - Outputs date time in string format using
// the FmtDateTimeTzNanoYMDDow format which incorporates date time to nano seconds
// and the associated time zone. In this format, the date is expressed as
// Year-Month-Day (Example: 2017-12-06) followed by the day of the week.
// EXAMPLE: 2006-01-02 Monday 15:04:05.000000000 -0700 MST
func (dt DateTimeUtility) GetDateTimeTzNanoSecYMDDowText(t time.Time) string {
	return t.Format(FmtDateTimeTzNanoYMDDow)
}

// GetDateTimeEverything - Receives a time value and formats as
// a date time string in the format:
// EXAMPLE: Saturday April 29, 2017 19:54:30.123456489 -0500 CDT
func (dt DateTimeUtility) GetDateTimeEverything(t time.Time) string {
	return t.Format(FmtDateTimeEverything)
}

// GetDateTimeCustomFmt - Returns time string
// formatted according to passed in format
// string. Example format string:
// 'Monday January 2, 2006 15:04:05.000000000 -0700 MST'
func (dt DateTimeUtility) GetDateTimeCustomFmt(t time.Time, fmt string) string {
	return t.Format(fmt)
}


// GetTimeStampEverything - Generates and returns a time stamp as
// type string. The current time is computed using time.Now() for the
// 'Local' timezone on the host machine. The time stamp is formatted
// using the format, 'FmtDateTimeEverything'. Example output:
// "Saturday April 29, 2017 19:54:30.123456489 -0500 CDT"
func (dt DateTimeUtility) GetTimeStampEverything() string {
	return time.Now().Local().Format(FmtDateTimeEverything)
}

// GetTimeStampYMDAbbrvDowNano - Generates and returns a time stamp as
// type string. The current time is computed using time.Now() for the
// 'Local' timezone on the host machine. The time stamp is formatted
// using the format 'FmtDateTimeYMDAbbrvDowNano'. Example output:
// "2006-01-02 Mon 15:04:05.000000000 -0700 MST"
func (dt DateTimeUtility) GetTimeStampYMDAbbrvDowNano() string {

	return time.Now().Local().Format(FmtDateTimeYMDAbbrvDowNano)

}
