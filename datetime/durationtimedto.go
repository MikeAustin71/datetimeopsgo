package datetime

import "time"

// DurationDto - This type holds duration
// values as time components.
type DurationDto struct {
	StartTimeTzu         TimeZoneDto
	EndTimeTzu           TimeZoneDto
	TimeDuration         time.Duration
	Years                int64
	YearsNanosecs        int64
	Months               int64
	MonthsNanosecs       int64
	Weeks                int64
	WeeksNanosecs        int64
	Days                 int64
	DaysNanosecs         int64
	Hours                int64
	HoursNanosecs        int64
	Minutes              int64
	MinutesNanosecs      int64
	Seconds              int64
	SecondsNanosecs      int64
	Milliseconds         int64
	MillisecondsNanosecs int64
	Microseconds         int64
	MicrosecondsNanosecs int64
	Nanoseconds          int64
	NanosecondsNanosecs  int64
	DisplayStr           string
}


// CalcTotalNanoSecs - Adds up all the time elements
// of the current DurationDto struct and converts
// the value to nanoseconds.
func (dDto *DurationDto) CalcTotalNanoSecs() int64 {
	ns := dDto.YearsNanosecs
	ns += dDto.MonthsNanosecs
	ns += dDto.WeeksNanosecs
	ns += dDto.DaysNanosecs
	ns += dDto.HoursNanosecs
	ns += dDto.MinutesNanosecs
	ns += dDto.SecondsNanosecs
	ns += dDto.MillisecondsNanosecs
	ns += dDto.MicrosecondsNanosecs
	ns += dDto.NanosecondsNanosecs

	return ns
}

// InitializeTime - Initializes a DurationDto Structure
// from input parameter type TimeDto.
func (dDto *DurationDto) InitializeTime(tDto TimeDto) {
	dDto.EmptyTimeValues()
	dDto.Years = int64(tDto.Years)
	dDto.Months = int64(tDto.Months)
	dDto.Weeks = int64(tDto.Weeks)
	dDto.Days = int64(tDto.WeekDays)
	dDto.Hours = int64(tDto.Hours)
	dDto.Minutes = int64(tDto.Minutes)
	dDto.Seconds = int64(tDto.Seconds)
	dDto.Milliseconds = int64(tDto.Milliseconds)
	dDto.Microseconds = int64(tDto.Microseconds)
	dDto.Nanoseconds = int64(tDto.Nanoseconds)
}

// Copy - Makes a deep copy of the current
// DurationDto struct and returns it as a new
// DurationDto instance.
func (dDto *DurationDto) Copy() DurationDto {
	d := DurationDto{}

	d.StartTimeTzu = dDto.StartTimeTzu
	d.EndTimeTzu = dDto.EndTimeTzu
	d.TimeDuration = dDto.TimeDuration
	d.Years = dDto.Years
	d.YearsNanosecs = dDto.YearsNanosecs
	d.Months = dDto.Months
	d.MonthsNanosecs = dDto.MonthsNanosecs
	d.Weeks = dDto.Weeks
	d.WeeksNanosecs = dDto.WeeksNanosecs
	d.Days = dDto.Days
	d.DaysNanosecs = dDto.DaysNanosecs
	d.Hours = dDto.Hours
	d.HoursNanosecs = dDto.HoursNanosecs
	d.Minutes = dDto.Minutes
	d.MinutesNanosecs = dDto.MinutesNanosecs
	d.Seconds = dDto.Seconds
	d.SecondsNanosecs = dDto.SecondsNanosecs
	d.Milliseconds = dDto.Milliseconds
	d.MillisecondsNanosecs = dDto.MillisecondsNanosecs
	d.Microseconds = dDto.Microseconds
	d.MicrosecondsNanosecs = dDto.MicrosecondsNanosecs
	d.Nanoseconds = dDto.Nanoseconds
	d.NanosecondsNanosecs = dDto.MicrosecondsNanosecs
	d.DisplayStr = dDto.DisplayStr

	return d
}

// Empty() - Resets all the time elements
// of the current DurationDto struct to their
// initial or 'zero' values.
func (dDto *DurationDto) Empty() {
	dDto.StartTimeTzu = TimeZoneDto{}
	dDto.EndTimeTzu = TimeZoneDto{}
	dDto.TimeDuration = time.Duration(0)
	dDto.Years = 0
	dDto.YearsNanosecs = 0
	dDto.Months = 0
	dDto.MonthsNanosecs = 0
	dDto.Weeks = 0
	dDto.WeeksNanosecs = 0
	dDto.Days = 0
	dDto.DaysNanosecs = 0
	dDto.Hours = 0
	dDto.HoursNanosecs = 0
	dDto.Minutes = 0
	dDto.MinutesNanosecs = 0
	dDto.Seconds = 0
	dDto.SecondsNanosecs = 0
	dDto.Milliseconds = 0
	dDto.MillisecondsNanosecs = 0
	dDto.Microseconds = 0
	dDto.MicrosecondsNanosecs = 0
	dDto.Nanoseconds = 0
	dDto.NanosecondsNanosecs = 0
	dDto.DisplayStr = ""
}

// EmptyTimeValues - Resets only the time values
// in the current DurationDto struct to their
// initial or 'zero' values.
func (dDto *DurationDto) EmptyTimeValues() {

	dDto.Years = 0
	dDto.YearsNanosecs = 0
	dDto.Months = 0
	dDto.MonthsNanosecs = 0
	dDto.Weeks = 0
	dDto.WeeksNanosecs = 0
	dDto.Days = 0
	dDto.DaysNanosecs = 0
	dDto.Hours = 0
	dDto.HoursNanosecs = 0
	dDto.Minutes = 0
	dDto.MinutesNanosecs = 0
	dDto.Seconds = 0
	dDto.SecondsNanosecs = 0
	dDto.Milliseconds = 0
	dDto.MillisecondsNanosecs = 0
	dDto.Microseconds = 0
	dDto.MicrosecondsNanosecs = 0
	dDto.Nanoseconds = 0
	dDto.NanosecondsNanosecs = 0
	dDto.DisplayStr = ""
}

// EmptyNanosecs - Resets all the Nanosecond
// values in the current DurationDto struct to
// zero.
func (dDto *DurationDto) EmptyNanosecs() {
	dDto.YearsNanosecs = 0
	dDto.MonthsNanosecs = 0
	dDto.WeeksNanosecs = 0
	dDto.DaysNanosecs = 0
	dDto.HoursNanosecs = 0
	dDto.MinutesNanosecs = 0
	dDto.SecondsNanosecs = 0
	dDto.MillisecondsNanosecs = 0
	dDto.MicrosecondsNanosecs = 0
	dDto.NanosecondsNanosecs = 0
	dDto.DisplayStr = ""
}

// Equal - Determines whether the input DurationDto is
// equal to the current DurationDto structure.
func (dDto *DurationDto) Equal(dto2 DurationDto) bool {

	if dDto.StartTimeTzu != dto2.StartTimeTzu ||
		dDto.EndTimeTzu != dto2.EndTimeTzu ||
		dDto.TimeDuration != dto2.TimeDuration ||
		dDto.Years != dto2.Years ||
		dDto.YearsNanosecs != dto2.YearsNanosecs ||
		dDto.Months != dto2.Months ||
		dDto.MonthsNanosecs != dto2.MonthsNanosecs ||
		dDto.Weeks != dto2.Weeks ||
		dDto.WeeksNanosecs != dto2.WeeksNanosecs ||
		dDto.Days != dto2.Days ||
		dDto.DaysNanosecs != dto2.DaysNanosecs ||
		dDto.Hours != dto2.Hours ||
		dDto.HoursNanosecs != dto2.HoursNanosecs ||
		dDto.Minutes != dto2.Minutes ||
		dDto.MinutesNanosecs != dto2.MinutesNanosecs ||
		dDto.Seconds != dto2.Seconds ||
		dDto.SecondsNanosecs != dto2.SecondsNanosecs ||
		dDto.Milliseconds != dto2.Milliseconds ||
		dDto.MillisecondsNanosecs != dto2.MillisecondsNanosecs ||
		dDto.Microseconds != dto2.Microseconds ||
		dDto.MicrosecondsNanosecs != dto2.MicrosecondsNanosecs ||
		dDto.Nanoseconds != dto2.Nanoseconds ||
		dDto.NanosecondsNanosecs != dto2.NanosecondsNanosecs ||
		dDto.DisplayStr != dto2.DisplayStr {

		return false
	}

	return true
}

