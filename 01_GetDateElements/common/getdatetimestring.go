package common

import "time"

const (
	// DateTimeFmtSecondStr - Date Time format used
	// for file names and directory names
	dateTimeFmtSecondStr   = "20060102150405"
	dateTimeFmtNanoSecText = "2006-01-02 15:04:05.000000000"
	dateTimeFmtSecText     = "2006-01-02 15:04:05"
	dateTimeFmtEverything  = "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
)

// DateTimeUtility - struct used to export
// Date Time Management methods.
type DateTimeUtility struct {
	TimeIn     time.Time
	TimeOut    time.Time
	Duration   time.Time
	TimeInStr  string
	TimeOutStr string
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
	return t.Format(dateTimeFmtSecondStr)

}

// GetDateTimeSecText - Returns formatted
// date time with seconds for display,
// 2006-01-02 15:04:05.
func (dt DateTimeUtility) GetDateTimeSecText(t time.Time) string {
	// Time Display Format with seconds
	return t.Format(dateTimeFmtSecText)
}

// GetDateTimeNanoSecText - Returns formated
// date time string with nanoseconds
// 2006-01-02 15:04:05.000000000.
func (dt DateTimeUtility) GetDateTimeNanoSecText(t time.Time) string {
	// Time Format down to the nanosecond
	return t.Format(dateTimeFmtNanoSecText)
}

// GetDateTimeEverything - Receives a time value and formats as
// a date time string in the format:
//  Saturday April 29, 2017 19:54:30.123456489 -0500 CDT
func (dt DateTimeUtility) GetDateTimeEverything(t time.Time) string {
	return t.Format(dateTimeFmtEverything)
}
