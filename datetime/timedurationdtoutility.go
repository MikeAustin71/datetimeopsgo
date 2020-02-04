package datetime

import (
	"fmt"
	"sync"
	"time"
)

type timeDurationDtoUtility struct {
	lock sync.Mutex
}

// copyIn - Copies the data fields from 'tDur2' into
// 'tDur1'
func (tDurDtoUtil *timeDurationDtoUtility) copyIn(
	tDur1 *TimeDurationDto,
	tDur2 *TimeDurationDto,
	ePrefix string) {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	if tDur1 == nil {
		ePrefix += "timeDurationDtoUtility.copyIn() "
		panic (ePrefix +
			"\nError: Input Parameter 'tDur1' is a 'nil' pointer!")
	}

	if tDur2 == nil {
		ePrefix += "timeDurationDtoUtility.copyIn() "
		panic (ePrefix +
			"\nError: Input Parameter 'tDur2' is a 'nil' pointer!")
	}

	tDurDtoUtil2 := timeDurationDtoUtility{}

	tDurDtoUtil2.empty(tDur1, ePrefix)

	tDur1.StartTimeDateTz = tDur2.StartTimeDateTz.CopyOut()
	tDur1.EndTimeDateTz = tDur2.EndTimeDateTz.CopyOut()
	tDur1.TimeDuration = tDur2.TimeDuration
	tDur1.CalcType = tDur2.CalcType
	tDur1.Years = tDur2.Years
	tDur1.YearsNanosecs = tDur2.YearsNanosecs
	tDur1.Months = tDur2.Months
	tDur1.MonthsNanosecs = tDur2.MonthsNanosecs
	tDur1.Weeks = tDur2.Weeks
	tDur1.WeeksNanosecs = tDur2.WeeksNanosecs
	tDur1.WeekDays = tDur2.WeekDays
	tDur1.WeekDaysNanosecs = tDur2.WeekDaysNanosecs
	tDur1.DateDays = tDur2.DateDays
	tDur1.DateDaysNanosecs = tDur2.DateDaysNanosecs
	tDur1.Hours = tDur2.Hours
	tDur1.HoursNanosecs = tDur2.HoursNanosecs
	tDur1.Minutes = tDur2.Minutes
	tDur1.MinutesNanosecs = tDur2.MinutesNanosecs
	tDur1.Seconds = tDur2.Seconds
	tDur1.SecondsNanosecs = tDur2.SecondsNanosecs
	tDur1.Milliseconds = tDur2.Milliseconds
	tDur1.MillisecondsNanosecs = tDur2.MillisecondsNanosecs
	tDur1.Microseconds = tDur2.Microseconds
	tDur1.MicrosecondsNanosecs = tDur2.MicrosecondsNanosecs
	tDur1.Nanoseconds = tDur2.MillisecondsNanosecs
	tDur1.TotSubSecNanoseconds = tDur2.TotSubSecNanoseconds
	tDur1.TotDateNanoseconds = tDur2.TotDateNanoseconds
	tDur1.TotTimeNanoseconds = tDur2.TotTimeNanoseconds

	return
}

// copyOut - Makes a deep copy of input parameter
// 'tDur' and returns the data field values in a
// new TimeDurationDto instance
func (tDurDtoUtil *timeDurationDtoUtility) copyOut(
	tDur *TimeDurationDto,
	ePrefix string) TimeDurationDto {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	if tDur == nil {
		ePrefix += "timeDurationDtoUtility.copyOut() "
		panic(ePrefix +
			"\nError: Input Parameter 'tDur' is a 'nil' pointer!")
	}

	t2Dur := TimeDurationDto{}

	t2Dur.StartTimeDateTz = tDur.StartTimeDateTz.CopyOut()
	t2Dur.EndTimeDateTz = tDur.EndTimeDateTz.CopyOut()
	t2Dur.TimeDuration = tDur.TimeDuration
	t2Dur.CalcType = tDur.CalcType
	t2Dur.Years = tDur.Years
	t2Dur.YearsNanosecs = tDur.YearsNanosecs
	t2Dur.Months = tDur.Months
	t2Dur.MonthsNanosecs = tDur.MonthsNanosecs
	t2Dur.Weeks = tDur.Weeks
	t2Dur.WeeksNanosecs = tDur.WeeksNanosecs
	t2Dur.WeekDays = tDur.WeekDays
	t2Dur.WeekDaysNanosecs = tDur.WeekDaysNanosecs
	t2Dur.DateDays = tDur.DateDays
	t2Dur.DateDaysNanosecs = tDur.DateDaysNanosecs
	t2Dur.Hours = tDur.Hours
	t2Dur.HoursNanosecs = tDur.HoursNanosecs
	t2Dur.Minutes = tDur.Minutes
	t2Dur.MinutesNanosecs = tDur.MinutesNanosecs
	t2Dur.Seconds = tDur.Seconds
	t2Dur.SecondsNanosecs = tDur.SecondsNanosecs
	t2Dur.Milliseconds = tDur.Milliseconds
	t2Dur.MillisecondsNanosecs = tDur.MillisecondsNanosecs
	t2Dur.Microseconds = tDur.Microseconds
	t2Dur.MicrosecondsNanosecs = tDur.MicrosecondsNanosecs
	t2Dur.Nanoseconds = tDur.Nanoseconds
	t2Dur.TotSubSecNanoseconds = tDur.TotSubSecNanoseconds
	t2Dur.TotDateNanoseconds = tDur.TotDateNanoseconds
	t2Dur.TotTimeNanoseconds = tDur.TotTimeNanoseconds

	return t2Dur
}

func (tDurDtoUtil *timeDurationDtoUtility) empty(
	tDur *TimeDurationDto,
	ePrefix string) {
		
		tDurDtoUtil.lock.Lock()
		
		defer tDurDtoUtil.lock.Unlock()
		
	if tDur == nil {
		ePrefix += "timeDurationDtoUtility.empty() "
		panic (ePrefix + 
			"\nError: Input Parameter 'tDur' is a 'nil' pointer!")
	}

	tDur.StartTimeDateTz = DateTzDto{}
	tDur.EndTimeDateTz = DateTzDto{}
	tDur.TimeDuration = time.Duration(0)
	tDur.CalcType = TDurCalcType(0).None()
	tDur.Years = 0
	tDur.YearsNanosecs = 0
	tDur.Months = 0
	tDur.MonthsNanosecs = 0
	tDur.Weeks = 0
	tDur.WeeksNanosecs = 0
	tDur.WeekDays = 0
	tDur.WeekDaysNanosecs = 0
	tDur.DateDays = 0
	tDur.DateDaysNanosecs = 0
	tDur.Hours = 0
	tDur.HoursNanosecs = 0
	tDur.Minutes = 0
	tDur.MinutesNanosecs = 0
	tDur.Seconds = 0
	tDur.SecondsNanosecs = 0
	tDur.Milliseconds = 0
	tDur.MillisecondsNanosecs = 0
	tDur.Microseconds = 0
	tDur.MicrosecondsNanosecs = 0
	tDur.Nanoseconds = 0
	tDur.TotSubSecNanoseconds = 0
	tDur.TotDateNanoseconds = 0
	tDur.TotTimeNanoseconds = 0
	
}

// emptyTimeFields - Sets time data fields to their zero
// or uninitialized values.
//
func (tDurDtoUtil *timeDurationDtoUtility) emptyTimeFields(
	tDur *TimeDurationDto,
	ePrefix string) {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	if tDur == nil {
		ePrefix += "timeDurationDtoUtility.empty() "
		panic(ePrefix +
			"\nError: Input Parameter 'tDur' is a 'nil' pointer!")
	}

	tDur.Years = 0
	tDur.YearsNanosecs = 0
	tDur.Months = 0
	tDur.MonthsNanosecs = 0
	tDur.Weeks = 0
	tDur.WeeksNanosecs = 0
	tDur.WeekDays = 0
	tDur.WeekDaysNanosecs = 0
	tDur.DateDays = 0
	tDur.DateDaysNanosecs = 0
	tDur.Hours = 0
	tDur.HoursNanosecs = 0
	tDur.Minutes = 0
	tDur.MinutesNanosecs = 0
	tDur.Seconds = 0
	tDur.SecondsNanosecs = 0
	tDur.Milliseconds = 0
	tDur.MillisecondsNanosecs = 0
	tDur.Microseconds = 0
	tDur.MicrosecondsNanosecs = 0
	tDur.Nanoseconds = 0
	tDur.TotSubSecNanoseconds = 0
	tDur.TotDateNanoseconds = 0
	tDur.TotTimeNanoseconds = 0

}

// isValid - Returns an error value signaling whether
// the data fields of input parameter 'tDur' are valid.
//
func (tDurDtoUtil *timeDurationDtoUtility) isValid(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.isValid() "

	if tDur == nil {
		ePrefix += "timeDurationDtoUtility.isValid() "
		panic(ePrefix +
			"\nError: Input Parameter 'tDur' is a 'nil' pointer!")
	}
	if tDur.StartTimeDateTz.GetDateTimeValue().IsZero() &&
		tDur.EndTimeDateTz.GetDateTimeValue().IsZero() {

		return fmt.Errorf(ePrefix +
			"\nError: Both Start and End Times are Zero!\n")

	}

	if tDur.EndTimeDateTz.GetDateTimeValue().Before(tDur.StartTimeDateTz.GetDateTimeValue()) {
		return fmt.Errorf(ePrefix +
			"\nError: End Time is Before Start Time!\n")
	}

	return nil
}