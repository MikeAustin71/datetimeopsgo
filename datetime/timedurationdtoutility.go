package datetime

import (
	"errors"
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

// calcDateDaysWeeksFromDuration - Calculates the Days associated
// with the duration for this TimeDurationDto.
//
// Calculates 'tDur.DateDays', 'tDur.DateDaysNanosecs', 'tDur.Weeks', 'tDur.WeeksNanosecs',
// 'tDur.WeekDays' and 'tDur.WeekDaysNanosecs'.
//
// NOTE: (1) Before calling this method, ensure that TimeDurationDto.StartTimeDateTz,
//           TimeDurationDto.EndTimeDateTz and tDur.TimeDuration are properly initialized.
//
//       (2) Before calling this method, ensure that the following methods are called
//           first, in sequence:
//             TimeDurationDto.calcYearsFromDuration
//             TimeDurationDto.calcMonthsFromDuration
//
func (tDurDtoUtil *timeDurationDtoUtility) calcDateDaysWeeksFromDuration(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcDateDaysWeeksFromDuration() "

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs

	// Calculate DateDays
	tDur.DateDays = 0
	tDur.DateDaysNanosecs = 0

	if rd >= DayNanoSeconds {
		tDur.DateDays = rd / DayNanoSeconds
		tDur.DateDaysNanosecs = DayNanoSeconds * tDur.DateDays
	}

	// Calculate Weeks and WeekDays
	tDur.Weeks = 0
	tDur.WeeksNanosecs = 0
	tDur.WeekDays = 0
	tDur.WeekDaysNanosecs = 0

	if tDur.DateDays > 0 {

		if tDur.DateDays >= 7 {

			tDur.Weeks = tDur.DateDays / int64(7)
			tDur.WeeksNanosecs = WeekNanoSeconds * tDur.Weeks

		}

		tDur.WeekDays = tDur.DateDays - (tDur.Weeks * 7)
		tDur.WeekDaysNanosecs = tDur.WeekDays * DayNanoSeconds

	}

	return nil
}

// calcHoursMinSecs - Calculates Hours, Minute, and
// Seconds of duration using startTime, tDur.StartTimeDateTz,
// and endTime, tDur.EndTimeDateTz.DateTime.
//
//
// NOTE: (1) Before calling this method, ensure that tDur.StartTimeDateTz,
//           TimeDurationDto.EndTimeDateTz and TimeDurationDto.TimeDuration
//           are properly initialized.
//
//       (2) Before calling this method, ensure that the following methods are called
//           first, in sequence:
//             TimeDurationDto.calcYearsFromDuration
//             TimeDurationDto.calcMonthsFromDuration
//             TimeDurationDto.calcDateDaysWeeksFromDuration
//
func (tDurDtoUtil *timeDurationDtoUtility) calcHoursMinSecs(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcHoursMinSecs() "

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	if tDur.DateDays > 0 {
		rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs +
			tDur.DateDaysNanosecs
	} else {
		rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs +
			tDur.WeeksNanosecs + tDur.WeekDaysNanosecs
	}

	tDur.Hours = 0
	tDur.HoursNanosecs = 0
	tDur.Minutes = 0
	tDur.MinutesNanosecs = 0
	tDur.Seconds = 0
	tDur.SecondsNanosecs = 0

	if rd >= HourNanoSeconds {
		tDur.Hours = rd / HourNanoSeconds
		tDur.HoursNanosecs = HourNanoSeconds * tDur.Hours
		rd -= tDur.HoursNanosecs
	}

	if rd >= MinuteNanoSeconds {
		tDur.Minutes = rd / MinuteNanoSeconds
		tDur.MinutesNanosecs = MinuteNanoSeconds * tDur.Minutes
		rd -= tDur.MinutesNanosecs
	}

	if rd >= SecondNanoseconds {
		tDur.Seconds = rd / SecondNanoseconds
		tDur.SecondsNanosecs = SecondNanoseconds * tDur.Seconds
		rd -= tDur.SecondsNanosecs
	}

	return nil
}

// calcNanoseconds - Calculates 'tDur.Milliseconds', 'tDur.MillisecondsNanosecs',
// 'tDur.Microseconds', 'tDur.MicrosecondsNanosecs',  and 'tDur.Nanoseconds'.
//
// NOTE: (1) Before calling this method, ensure that tDur.StartTimeDateTz,
//           TimeDurationDto.EndTimeDateTz and TimeDurationDto.TimeDuration
//           are properly initialized.
//
//       (2) Before calling this method, ensure that the following methods are called
//           first, in sequence:
//             TimeDurationDto.calcYearsFromDuration
//             TimeDurationDto.calcMonthsFromDuration
//             TimeDurationDto.calcDateDaysWeeksFromDuration
//             TimeDurationDto.calcHoursMinSecs
//
func (tDurDtoUtil *timeDurationDtoUtility) calcNanoseconds(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcNanoseconds() "

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs

	if tDur.DateDaysNanosecs > 0 {
		rd -= tDur.DateDaysNanosecs
	} else {
		rd -= tDur.WeeksNanosecs + tDur.WeekDaysNanosecs
	}

	rd -= tDur.HoursNanosecs +
		tDur.MinutesNanosecs + tDur.SecondsNanosecs

	tDur.Milliseconds = 0
	tDur.MillisecondsNanosecs = 0
	tDur.Microseconds = 0
	tDur.MicrosecondsNanosecs = 0
	tDur.Nanoseconds = 0

	if rd >= MilliSecondNanoseconds {
		tDur.Milliseconds = rd / MilliSecondNanoseconds
		tDur.MillisecondsNanosecs = MilliSecondNanoseconds * tDur.Milliseconds
		rd -= tDur.MillisecondsNanosecs
	}

	if rd >= MicroSecondNanoseconds {
		tDur.Microseconds = rd / MicroSecondNanoseconds
		tDur.MicrosecondsNanosecs = MicroSecondNanoseconds * tDur.Microseconds
		rd -= tDur.MicrosecondsNanosecs
	}

	tDur.Nanoseconds = rd

	return nil
}

// calcMonthsFromDuration - calculates the months duration
// using the start and end dates, 'tDur.StartTimeDateTz' and
// 'tDur.EndTimeDateTz.DateTime'.
//
// NOTE: (1) Before calling this method, ensure that tDur.StartTimeDateTz,
//           tDur.EndTimeDateTz and tDur.TimeDuration are properly initialized.
//
//       (2) Before calling this method, ensure that the following method is called
//           first:
//              TimeDurationDto.calcYearsFromDuration
//
func (tDurDtoUtil *timeDurationDtoUtility) calcMonthsFromDuration(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcMonthsFromDuration() "

	startTime := tDur.StartTimeDateTz.dateTimeValue
	endTime := tDur.EndTimeDateTz.dateTimeValue

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() !=
		endTime.Location().String() {

		return fmt.Errorf(ePrefix +
			"Error: 'startTime' and 'endTime' Time Zone Location do NOT match!\n"+
			"startTimeZoneLocation='%v'\n" +
			"endTimeZoneLocation='%v'\n",
			startTime.Location().String(),
			endTime.Location().String())
	}

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	rd -= tDur.YearsNanosecs

	i := 0

	yearDateTime := startTime.Add(time.Duration(tDur.YearsNanosecs))

	mthDateTime := yearDateTime

	for mthDateTime.Before(endTime) {

		i++

		mthDateTime = yearDateTime.AddDate(0, i, 0)

	}

	i -= 1

	if i > 0 {

		tDur.Months = int64(i)

		mthDateTime = yearDateTime.AddDate(0, i, 0)

		tDur.MonthsNanosecs = int64(mthDateTime.Sub(yearDateTime))

	} else {
		tDur.Months = 0
		tDur.MonthsNanosecs = 0
	}

	return nil
}

// calcSummaryTimeElements - Calculates totals for Date, Time and
// sub-second nanoseconds.
//
// NOTE: (1) Before calling this method, ensure that tDur.StartTimeDateTz,
//           TimeDurationDto.EndTimeDateTz and TimeDurationDto.TimeDuration
//           are properly initialized.
//
//       (2) Before calling this method, ensure that the following methods are called
//           first, in sequence:
//             TimeDurationDto.calcYearsFromDuration
//             TimeDurationDto.calcMonthsFromDuration
//             TimeDurationDto.calcDateDaysWeeksFromDuration
//             TimeDurationDto.calcHoursMinSecs
//             TimeDurationDto.calcNanoseconds
//
func (tDurDtoUtil *timeDurationDtoUtility) calcSummaryTimeElements(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcSummaryTimeElements() "

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	tDur.TotDateNanoseconds = 0
	tDur.TotTimeNanoseconds = 0
	tDur.TotSubSecNanoseconds = 0

	tDur.TotDateNanoseconds = tDur.YearsNanosecs
	tDur.TotDateNanoseconds += tDur.MonthsNanosecs

	if tDur.DateDaysNanosecs == 0 {
		tDur.TotDateNanoseconds += tDur.WeeksNanosecs
		tDur.TotDateNanoseconds += tDur.WeekDaysNanosecs
	} else {
		tDur.TotDateNanoseconds += tDur.DateDaysNanosecs
	}

	tDur.TotSubSecNanoseconds = tDur.MillisecondsNanosecs
	tDur.TotSubSecNanoseconds += tDur.MicrosecondsNanosecs
	tDur.TotSubSecNanoseconds += tDur.Nanoseconds

	tDur.TotTimeNanoseconds = tDur.HoursNanosecs
	tDur.TotTimeNanoseconds += tDur.MinutesNanosecs
	tDur.TotTimeNanoseconds += tDur.SecondsNanosecs
	tDur.TotTimeNanoseconds += tDur.TotSubSecNanoseconds

	return nil
}

// calcTimeDurationAllocations - Examines the input parameter 'calcType' and
// then determines which type of time duration allocation calculation will be
// applied to the data fields of the current TimeDurationDto instance.
func (tDurDtoUtil *timeDurationDtoUtility) calcTimeDurationAllocations(
	tDur *TimeDurationDto,
	calcType TDurCalcType,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTimeDurationAllocations() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	switch calcType {

	case TDurCalcType(0).StdYearMth():
		return tDurDtoUtilX2.calcTypeStdYearMth(
			tDur,
			ePrefix)

	case TDurCalcType(0).CumMonths():
		return tDurDtoUtilX2.calcTypeCumMonths(
			tDur,
			ePrefix)

	case TDurCalcType(0).CumWeeks():
		return tDurDtoUtilX2.calcTypeCumWeeks(
			tDur,
			ePrefix)

	case TDurCalcType(0).CumDays():
		return tDurDtoUtilX2.calcTypeCumDays(tDur, ePrefix)

	case TDurCalcType(0).CumHours():
		return tDurDtoUtilX2.calcTypeCumHours(tDur, ePrefix)

	case TDurCalcType(0).CumMinutes():
		return tDurDtoUtilX2.calcTypeCumMinutes(tDur, ePrefix)

	case TDurCalcType(0).CumSeconds():
		return tDurDtoUtilX2.calcTypeCumSeconds(tDur, ePrefix)

	case TDurCalcType(0).GregorianYears():
		return tDurDtoUtilX2.calcTypeGregorianYears(tDur, ePrefix)
	}

	return fmt.Errorf(ePrefix+
		"\nError: Invalid TDurCalcType.\ncalcType='%v'\n", calcType.String())
}

// calcTypeCumDays - Calculates Cumulative Days. Years, months and weeks are consolidated
// and counted as cumulative days. The Data Fields for years, months, weeks and week days
// are set to zero.  All cumulative days are allocated to the data field, 'DateDays'.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCumDays(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTypeCumDays() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(tDur, ePrefix)

	tDur2.CalcType = TDurCalcType(0).CumDays()

	rd := int64(tDur2.TimeDuration)

	if rd == 0 {
		return nil
	}

	if rd >= DayNanoSeconds {
		tDur2.DateDays = rd / DayNanoSeconds
		tDur2.DateDaysNanosecs = tDur2.DateDays * DayNanoSeconds
	}

	err := tDurDtoUtilX2.calcHoursMinSecs(
		&tDur2, 
		ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcNanoseconds(
		&tDur2, 
		ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcSummaryTimeElements(
		&tDur2, 
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtilX2.copyIn(tDur, &tDur2, ePrefix)
	
	return nil
}

// calcTypeCumHours - Calculates Cumulative Hours. Years, months, weeks, week days,
// date days and hours are consolidated and included in cumulative hours. Values for years,
// months, weeks, week days and date days are ignored and set to zero. Time duration is
// allocated over cumulative hours plus minutes, seconds, milliseconds, microseconds and
// nanoseconds.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCumHours(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTypeCumHours() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDur2.CalcType = TDurCalcType(0).CumHours()

	if tDur2.TimeDuration == 0 {
		return nil
	}

	err := tDurDtoUtilX2.calcHoursMinSecs(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcNanoseconds(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcSummaryTimeElements(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtilX2.copyIn(tDur, &tDur2, ePrefix)
	
	return nil
}


// calcTypeCumMinutes - Calculates Cumulative Minutes. Years, months, weeks, week days,
// date days, hours and minutes are consolidated and included in cumulative minutes.
// Values for years, months, weeks, week days, date days and hours are ignored and set
// to zero. Time duration is allocated over cumulative minutes plus seconds, milliseconds,
// microseconds and nanoseconds.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCumMinutes(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTypeCumMinutes() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDur2.CalcType = TDurCalcType(0).CumMinutes()

	if tDur2.TimeDuration == 0 {
		return nil
	}

	rd := int64(tDur2.TimeDuration)

	if rd >= MinuteNanoSeconds {
		tDur2.Minutes = rd / MinuteNanoSeconds
		tDur2.MinutesNanosecs = tDur2.Minutes * MinuteNanoSeconds
		rd -= tDur2.MinutesNanosecs
	}

	if rd >= SecondNanoseconds {
		tDur2.Seconds = rd / SecondNanoseconds
		tDur2.SecondsNanosecs = tDur2.Seconds * SecondNanoseconds
		rd -= tDur2.SecondsNanosecs
	}

	if rd >= MilliSecondNanoseconds {
		tDur2.Milliseconds = rd / MilliSecondNanoseconds
		tDur2.MillisecondsNanosecs = tDur2.Milliseconds * MilliSecondNanoseconds
		rd -= tDur2.MillisecondsNanosecs
	}

	if rd >= MicroSecondNanoseconds {
		tDur2.Microseconds = rd / MicroSecondNanoseconds
		tDur2.MillisecondsNanosecs = tDur2.Microseconds * MicroSecondNanoseconds
		rd -= tDur2.MicrosecondsNanosecs
	}

	tDur2.Nanoseconds = rd

	err := tDurDtoUtilX2.calcSummaryTimeElements(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtilX2.copyIn(tDur, &tDur2, ePrefix)
	
	return nil
}


// Data Fields for Years is always set to Zero. Years
// and months are consolidated and counted as cumulative
// months.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCumMonths(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTypeCumMonths() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDur2.CalcType = TDurCalcType(0).CumMonths()

	err := tDurDtoUtilX2.calcMonthsFromDuration(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcDateDaysWeeksFromDuration(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcHoursMinSecs(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcNanoseconds(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcSummaryTimeElements(&tDur2, ePrefix)

	if err != nil {
		return err
	}
	
	tDurDtoUtilX2.copyIn(tDur, &tDur2, ePrefix)
	
	return nil
}

// calcTypeCumSeconds - Calculates Cumulative Seconds of
// time duration.
//
// tDur.CalcType = TDurCalcType(0).CumSeconds()
//
// Years, months, weeks, weekdays, date days, hours and
// minutes are ignored and set to zero. Time is accumulated
// in seconds, milliseconds, microseconds and nanoseconds.
//
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCumSeconds(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTypeCumSeconds() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDur2.CalcType = TDurCalcType(0).CumSeconds()

	rd := int64(tDur2.TimeDuration)

	if rd == 0 {
		return nil
	}

	if rd >= SecondNanoseconds {
		tDur2.Seconds = rd / SecondNanoseconds
		tDur2.SecondsNanosecs = SecondNanoseconds * tDur2.Seconds
		rd -= tDur2.SecondsNanosecs
	}

	if rd >= MilliSecondNanoseconds {
		tDur2.Milliseconds = rd / MilliSecondNanoseconds
		tDur2.MillisecondsNanosecs = MilliSecondNanoseconds * tDur2.Milliseconds
		rd -= tDur2.MillisecondsNanosecs
	}

	if rd >= MicroSecondNanoseconds {
		tDur2.Microseconds = rd / MicroSecondNanoseconds
		tDur2.MicrosecondsNanosecs = MicroSecondNanoseconds * tDur2.Microseconds
		rd -= tDur2.MicrosecondsNanosecs
	}

	tDur2.Nanoseconds = rd

	tDurDtoUtilX2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// calcTypeCumWeeks - Data Fields for Years and Months are always set to zero.
// Years and Months are consolidated and counted as equivalent Weeks.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCumWeeks(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTypeCumWeeks() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDur2.CalcType = TDurCalcType(0).CumWeeks()

	rd := int64(tDur2.TimeDuration)

	if rd >= WeekNanoSeconds {

		tDur2.Weeks = rd / WeekNanoSeconds
		tDur2.WeeksNanosecs = tDur2.Weeks * WeekNanoSeconds
		rd -= tDur2.WeeksNanosecs
	}

	if rd >= DayNanoSeconds {
		tDur2.WeekDays = rd / DayNanoSeconds
		tDur2.WeekDaysNanosecs = tDur2.WeekDays * DayNanoSeconds
		rd -= tDur2.WeekDaysNanosecs
	}

	tDur2.DateDays = tDur2.Weeks * int64(7)
	tDur2.DateDays += tDur2.WeekDays
	tDur2.DateDaysNanosecs = tDur2.WeeksNanosecs + tDur2.WeekDaysNanosecs

	err := tDurDtoUtilX2.calcHoursMinSecs(
		&tDur2, 
		ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcNanoseconds(
		&tDur2,
		ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcSummaryTimeElements(
		&tDur2,
		ePrefix)

	if err != nil {
		return err
	}

	// For Cumulative Weeks calculation and presentations, set Date Days to zero
	tDur2.DateDays = 0
	tDur2.DateDaysNanosecs = 0

	tDurDtoUtilX2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// calcTypeGregorianYears - Allocates Years using the number of nanoseconds in a
// standard or average GregorianYear
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeGregorianYears(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTypeGregorianYears() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDur2.CalcType = TDurCalcType(0).GregorianYears()

	rd := int64(tDur2.TimeDuration)

	if rd == 0 {
		return nil
	}

	if rd >= GregorianYearNanoSeconds {
		tDur2.Years = rd / GregorianYearNanoSeconds
		tDur2.YearsNanosecs = tDur2.Years * GregorianYearNanoSeconds
	}

	err := tDurDtoUtilX2.calcMonthsFromDuration(
					&tDur2,
					ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcDateDaysWeeksFromDuration(
		&tDur2,
		ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcHoursMinSecs(
		&tDur2,
		ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcNanoseconds(
		&tDur2,
		ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcSummaryTimeElements(
		&tDur2,
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtilX2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// calcTypeStdYearMth - Performs Duration calculations for
// TDurCalcType == TDurCalcType(0).StdYearMth()
//
// TDurCalcTypeYEARMTH - Standard Year, Month, Weeks, Days calculation.
// All data fields in the TimeDto are populated in the duration
// allocation.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeStdYearMth(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTypeStdYearMth() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}
	
	tDur2 := tDurDtoUtilX2.copyOut(
						tDur,
						ePrefix)

	tDur2.CalcType = TDurCalcType(0).StdYearMth()

	err := tDurDtoUtilX2.calcYearsFromDuration(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcMonthsFromDuration(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcDateDaysWeeksFromDuration(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcHoursMinSecs(&tDur2, ePrefix)

	if err != nil {
		return nil
	}

	err = tDurDtoUtilX2.calcNanoseconds(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	err = tDurDtoUtilX2.calcSummaryTimeElements(&tDur2, ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtilX2.copyIn(tDur, &tDur2, ePrefix)
	
	return err
}

// calcYearsFromDuration - Calculates number of years duration and nanoseconds
// represented by years duration using input parameters 'tDur.StartTimeDateTz' and
// 'tDur.EndTimeDateTz'.
//
// NOTE: Before calling this method, ensure that tDur.StartTimeDateTz,
//       tDur.EndTimeDateTz and tDur.TimeDuration are properly initialized.
//
func (tDurDtoUtil *timeDurationDtoUtility) calcYearsFromDuration(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcYearsFromDuration() "

	years := int64(0)
	yearNanosecs := int64(0)
	startTime := tDur.StartTimeDateTz.dateTimeValue
	endTime := tDur.EndTimeDateTz.dateTimeValue

	if endTime.Before(startTime) {
		return errors.New(ePrefix + 
			"\nError: 'endTime' precedes, is less than, startTime!\n")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix+
			"\nError: 'startTime' and 'endTime' Time Zone Location do NOT match!\n"+
			"startTimeZoneLocation='%v'\n" +
			"endTimeZoneLocation='%v'\n",
			startTime.Location().String(), endTime.Location().String())
	}

	yearDateTime := startTime

	i := 0

	for yearDateTime.Before(endTime) {

		i++

		yearDateTime = startTime.AddDate(i, 0, 0)

	}

	i--

	if i > 0 {

		years = int64(i)

		yearDateTime = startTime.AddDate(i, 0, 0)

		duration := yearDateTime.Sub(startTime)

		yearNanosecs = int64(duration)

	} else {

		years = 0

		yearNanosecs = 0
	}

	tDur.Years = years
	tDur.YearsNanosecs = yearNanosecs

	return nil
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
// Input Parameter
// ===============
//
// tDur     *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
//
// ePrefix  string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) isValid(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.isValid() "

	if tDur == nil {
		return fmt.Errorf(ePrefix +
			"\nError: Input Parameter 'tDur' is a 'nil' pointer!")
	}

	if tDur.StartTimeDateTz.dateTimeValue.IsZero() &&
		tDur.EndTimeDateTz.dateTimeValue.IsZero() {

		return fmt.Errorf(ePrefix +
			"\nError: Both Start and End Times are Zero!\n")

	}

	if tDur.EndTimeDateTz.dateTimeValue.Before(tDur.StartTimeDateTz.dateTimeValue) {
		return fmt.Errorf(ePrefix +
			"\nError: End Time is Before Start Time!\n")
	}

	return nil
}


// reCalcTimeDurationAllocation - Re-calculates and allocates time duration for the current
// TimeDurationDto instance over the various time components (years, months, weeks, weekdays,
// datedays, hour, minutes, seconds, milliseconds, microseconds and nanoseconds) depending
// on the value of the 'TDurCalcType' input parameter.
//
// Input Parameters
// ================
//
// tDur    *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
//
// ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
func (tDurDtoUtil *timeDurationDtoUtility) reCalcTimeDurationAllocation(
	tDur *TimeDurationDto,
	tDurCalcType TDurCalcType,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.reCalcTimeDurationAllocation() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	tDurDtoUtil2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtil2.copyOut(tDur, ePrefix)

	err := tDurDtoUtil2.calcTimeDurationAllocations(
		&tDur2,
		tDurCalcType,
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtil2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// reCalcEndDateTimeToNow - Recomputes time duration values for the
// current TimeDurationDto by setting ending date time to time.Now().
// This is useful in stop watch applications.
//
// The Time Zone Location is derived from the existing starting date
// time, 'tDur.StartTimeDateTz'.  The Calculation type is taken from
// the existing calculation type, 'tDur.CalcType'.
//
// Input Parameter
// ===============
//
// tDur     *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
//
// ePrefix  string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) reCalcEndDateTimeToNow(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.reCalcEndDateTimeToNow() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	tDurDtoUtil2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtil2.copyOut(tDur, ePrefix)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr := dtMech.PreProcessDateFormatStr(tDur.StartTimeDateTz.dateTimeFmt)

	dTzUtil := dateTzDtoUtility{}

	tDur2.StartTimeDateTz = dTzUtil.copyOut(&tDur.StartTimeDateTz)

	err := dTzUtil.setFromTzDef(
		&tDur2.EndTimeDateTz,
		time.Now().UTC(),
		TzConvertType.Relative(),
		tDur2.StartTimeDateTz.timeZone.CopyOut(),
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.TimeDuration =
		tDur2.EndTimeDateTz.dateTimeValue.Sub(
			tDur2.StartTimeDateTz.dateTimeValue)

	err = tDurDtoUtil2.calcTimeDurationAllocations(
		&tDur2,
		tDur.CalcType,
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtil2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// setAutoEnd - When called, this method automatically sets the ending date
// time and re-calculates the time duration for the current TimeDurationDto
// instance.
//
// Ending date time is assigned the value returned by time.Now(). This ending
// date time is converted to the specified Time Zone specified by the Time Zone
// Location associated with the current starting date time value.
//
// When used together, the two methods 'NewAutoStart' and this method, 'SetAutoEnd'
// function as a stop watch feature. Simply calling these functions can set
// the starting date time and later the ending date time to measure elapsed time, or
// time duration.
//
// The time duration calculation type is taken from the current TimeDurationDto
// calculation type setting.
//
func (tDurDtoUtil *timeDurationDtoUtility) setAutoEnd(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setStartEndTimesCalcTz() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	tDurDtoUtil2 := timeDurationDtoUtility{}

	err := tDurDtoUtil2.isValid(tDur, ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nInput Parameter 'tDur' is INVALID!\n" +
			"Validation Error='%v'\n", err.Error())
	}

	startDateTime := tDur.StartTimeDateTz.dateTimeValue

	endDateTime := time.Now().Local()

	tzDef := tDur.StartTimeDateTz.timeZone.CopyOut()

	dtMech := DTimeMechanics{}

	dTzUtil := dateTzDtoUtility{}

	tDur2 := TimeDurationDto{}

	dateTimeFmtStr := dtMech.PreProcessDateFormatStr(tDur.StartTimeDateTz.dateTimeFmt)

	err = dTzUtil.setFromTzDef(
		&tDur2.StartTimeDateTz,
		startDateTime,
		TzConvertType.Relative(),
		tzDef,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	err = dTzUtil.setFromTzDef(
		&tDur2.EndTimeDateTz,
		endDateTime,
		TzConvertType.Relative(),
		tzDef,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.TimeDuration =
			tDur2.EndTimeDateTz.dateTimeValue.Sub(
				tDur2.StartTimeDateTz.dateTimeValue)

	calcType := tDur.CalcType

	err = tDurDtoUtil2.calcTimeDurationAllocations(
		&tDur2,
		calcType,
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtil2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// SetEndTimeMinusTimeDtoCalcTz - Sets start date time, end date time and duration
// based on an ending date time and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// Input Parameters:
// =================
//
// tDur    *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
// endDateTime time.Time
//            - Ending date time. The starting date time will be computed
//              by subtracting minusTimeDto from 'endDateTime'
//
// minusTimeDto TimeDto
//            - Time components (Years, months, weeks, days, hours etc.)
//              which will be subtracted from 'endDateTime' to compute
//              time duration and starting date time.
//
//              type TimeDto struct {
//               Years          int // Number of Years
//               Months         int // Number of Months
//               Weeks          int // Number of Weeks
//               WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//               DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//               Hours          int // Number of Hours.
//               Minutes        int // Number of Minutes
//               Seconds        int // Number of Seconds
//               Milliseconds   int // Number of Milliseconds
//               Microseconds   int // Number of Microseconds
//               Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//               TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                        //  plus remaining Nanoseconds
//              }
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                    Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time comparisons.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setEndTimeMinusTimeDtoCalcTz(
	tDur *TimeDurationDto,
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setEndTimeMinusTimeDtoCalcTz() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	err := minusTimeDto.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Input Parameter 'minusTimeDto' is INVALID!\n" +
			"Validation Error='%v'\n", err.Error())
	}

	if endDateTime.IsZero() && minusTimeDto.IsEmpty() {
		return errors.New(ePrefix +
			"\nError: Both 'endDateTime' and 'minusTimeDto' " +
			"input parameters are ZERO/EMPTY!\n")
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tDur2 := TimeDurationDto {}

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTimeTzName(
		&tDur2.EndTimeDateTz,
		endDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.StartTimeDateTz, err =
		dTzUtil.addMinusTimeDto(
					&tDur2.EndTimeDateTz,
					minusTimeDto,
					ePrefix)

	if err != nil {
		return err
	}

	tDur2.TimeDuration =
		tDur2.EndTimeDateTz.dateTimeValue.Sub(
				tDur2.StartTimeDateTz.dateTimeValue)

	tDurDtoUtil2 := timeDurationDtoUtility{}

	err = tDurDtoUtil2.calcTimeDurationAllocations(
					&tDur2,
					tDurCalcType,
					ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtil2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// setStartEndTimesCalcTz - Sets data field values for the current TimeDurationDto
// instance using a Start Date Time, End Date Time and a time zone specification.
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
//
// All data fields in the current TimeDurationDto instance are overwritten with
// the new time duration values.
//
// Input Parameters:
// =================
//
// tDur *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
//
// startDateTime time.Time - Starting time
//
// endDateTime  time.Time - Ending time
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                     or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//     Type 'TDurCalcType' is located in source file:
//          MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//              FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setStartEndTimesCalcTz(
	tDur *TimeDurationDto,
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) error {

		tDurDtoUtil.lock.Lock()

		defer tDurDtoUtil.lock.Unlock()

		ePrefix += "timeDurationDtoUtility.setStartEndTimesCalcTz() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Both 'startDateTime' and 'endDateTime' " +
			"input parameters are ZERO!\n")
	}

	// If endDateTime is less than startDateTime
	// reverse the order.
	if endDateTime.Before(startDateTime) {
		tempDateTime := startDateTime
		startDateTime = endDateTime
		endDateTime = tempDateTime
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tDurCalcType",
				inputParameterValue: tDurCalcType.String(),
				errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
				err:                 nil,
			}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tDur2 := TimeDurationDto{}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromTimeTzName(
		&tDur2.StartTimeDateTz,
		startDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	err = dTzUtil.setFromTimeTzName(
		&tDur2.EndTimeDateTz,
		endDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.TimeDuration =
		tDur2.EndTimeDateTz.dateTimeValue.Sub(
					tDur2.StartTimeDateTz.dateTimeValue)

	tDurDtoUtil2 := timeDurationDtoUtility{}

	err = tDurDtoUtil2.calcTimeDurationAllocations(
		&tDur2,
		tDurCalcType,
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtil2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// setStartTimeDurationCalcTz - Sets start time, end time and duration for the
// current TimeDurationDto instance. 'startDateTime' is converted to the
// specified 'timeZoneLocation' and the duration value is added to it
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to
// ending date time and the actual starting date time is computed by
// subtracting duration.
//
// Input Parameters:
// =================
//
// tDur *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
//
// startDateTime time.Time
//            - Starting date time for the duration calculation
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()   - Default - standard year, month week,
//                                      day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//            - A date time format string which will be used
//              to format and display 'dateTime'. Example:
//              "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setStartTimeDurationCalcTz(
	tDur *TimeDurationDto,
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setStartTimeDurationCalcTz() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if startDateTime.IsZero() && duration == 0 {
		return errors.New(ePrefix +
			"\nError: Both 'startDateTime' and 'duration' " +
			"input parameters are ZERO!\n")
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	var endDateTime time.Time

	if duration < 0 {
		endDateTime = startDateTime

		startDateTime = endDateTime.Add(duration)

		duration = duration * -1

	} else {

		endDateTime = startDateTime.Add(duration)

	}

	tDur2 := TimeDurationDto{}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromTimeTzName(
		&tDur2.StartTimeDateTz,
		startDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	err = dTzUtil.setFromTimeTzName(
		&tDur2.EndTimeDateTz,
		endDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.TimeDuration = duration

	tDurDtoUtil2 := timeDurationDtoUtility{}

	err = tDurDtoUtil2.calcTimeDurationAllocations(
		&tDur2,
		tDurCalcType,
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtil2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// SetStartTimeDurationDateDtoCalcTz - Sets start time, end time and
// duration for the current TimeDurationDto instance.
//
// The input parameter, 'startDateTime', is of type DateTzDto. It is
// converted to the specified 'timeZoneLocation' and the duration value
// is added to in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to
// ending date time and the actual starting date time is computed by
// subtracting duration.
//
// Input Parameters:
// =================
//
// tDur *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
//
// startDateTime DateTzDto
//            - Provides starting date time for the duration calculation
//
// duration  time.Duration
//            - Amount of time to be added to or subtracted from
//              'startDateTime'. Note: If duration is a negative value
//              'startDateTime' is converted to ending date time and
//              actual starting date time is computed by subtracting
//              duration.
//
// tDurCalcType TDurCalcType
//            - Specifies the calculation type to be used in allocating
//              time duration:
//
//     TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                        day time calculation.
//
//     TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//     TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//     TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//     TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//     TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                        or hours.
//
//     TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                        hours or minutes.
//
//     TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                        Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//            MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
// timeZoneLocation string
//            - Designates the standard Time Zone location by which
//              time duration will be compared. This ensures that
//              'oranges are compared to oranges and apples are compared
//              to apples' with respect to start time and end time duration
//              calculations.
//
//              If 'timeZoneLocation' is passed as an empty string, it
//              will be automatically defaulted to the 'UTC' time zone.
//              Reference Universal Coordinated Time:
//                 https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//              Time zone location must be designated as one of three types of
//              time zones.
//
//              (1) The time zone "Local", which Golang accepts as
//                  the time zone currently configured on the host
//                  computer.
//
//              (2) IANA Time Zone - A valid IANA Time Zone from the
//                  IANA database.
//                  See https://golang.org/pkg/time/#LoadLocation
//                  and https://www.iana.org/time-zones to ensure that
//                  the IANA Time Zone Database is properly configured
//                  on your system.
//
//                  IANA Time Zone Examples:
//                    "America/New_York"
//                    "America/Chicago"
//                    "America/Denver"
//                    "America/Los_Angeles"
//                    "Pacific/Honolulu"
//                    "Etc/UTC" = GMT or UTC
//
//              (3) A Military Time Zone
//                  Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                   Examples:
//                     "Alpha"   or "A"
//                     "Bravo"   or "B"
//                     "Charlie" or "C"
//                     "Delta"   or "D"
//                     "Zulu"    or "Z"
//
//                     If the time zone "Zulu" is passed to this method, it will be
//                     classified as a Military Time Zone.
//
// dateTimeFmtStr string
//             - A date time format string which will be used
//               to format and display 'dateTime'. Example:
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//              If 'dateTimeFmtStr' is submitted as an
//              'empty string', a default date time format
//              string will be applied. The default date time
//              format string is:
//                 FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ePrefix string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setStartTimeDurationDateDtoCalcTz(
	tDur *TimeDurationDto,
	startDateTimeTz DateTzDto,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setStartTimeDurationDateDtoCalcTz() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if startDateTimeTz.dateTimeValue.IsZero() && duration == 0 {
		return errors.New(ePrefix +
			"\nError: Both 'startDateTimeTz' and 'duration' " +
			"input parameters are ZERO!\n")
	}


	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	var startDateTime, endDateTime time.Time

	if duration < 0 {
		endDateTime = startDateTimeTz.dateTimeValue

		startDateTime = endDateTime.Add(duration)

		duration = duration * -1

	} else {
		startDateTime = startDateTimeTz.dateTimeValue
		endDateTime = startDateTime.Add(duration)

	}

	tDur2 := TimeDurationDto{}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.setFromTimeTzName(
		&tDur2.StartTimeDateTz,
		startDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	err = dTzUtil.setFromTimeTzName(
		&tDur2.EndTimeDateTz,
		endDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.TimeDuration = duration

	tDurDtoUtil2 := timeDurationDtoUtility{}

	err = tDurDtoUtil2.calcTimeDurationAllocations(
		&tDur2,
		tDurCalcType,
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtil2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}