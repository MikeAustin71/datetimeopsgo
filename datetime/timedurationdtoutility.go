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
		panic(ePrefix +
			"\nError: Input Parameter 'tDur1' is a 'nil' pointer!")
	}

	if tDur1.lock == nil {
		tDur1.lock = new(sync.Mutex)
	}

	if tDur2 == nil {
		ePrefix += "timeDurationDtoUtility.copyIn() "
		panic(ePrefix +
			"\nError: Input Parameter 'tDur2' is a 'nil' pointer!")
	}

	if tDur2.lock == nil {
		tDur2.lock = new(sync.Mutex)
	}

	tDurDtoUtil2 := timeDurationDtoUtility{}

	tDurDtoUtil2.empty(tDur1, ePrefix)

	tDur1.startDateTimeTz = tDur2.startDateTimeTz.CopyOut()
	tDur1.endDateTimeTz = tDur2.endDateTimeTz.CopyOut()
	tDur1.timeDuration = tDur2.timeDuration
	tDur1.timeDurCalcType = tDur2.timeDurCalcType
	tDur1.timeMathCalcMode = tDur2.timeMathCalcMode
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
	tDur1.Nanoseconds = tDur2.Nanoseconds
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

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	t2Dur := TimeDurationDto{}

	t2Dur.startDateTimeTz = tDur.startDateTimeTz.CopyOut()
	t2Dur.endDateTimeTz = tDur.endDateTimeTz.CopyOut()
	t2Dur.timeDuration = tDur.timeDuration
	t2Dur.timeDurCalcType = tDur.timeDurCalcType
	t2Dur.timeMathCalcMode = tDur.timeMathCalcMode
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
	t2Dur.lock = new(sync.Mutex)

	return t2Dur
}

// calcDateDaysWeeksFromDuration - Calculates the Days associated
// with the duration for this TimeDurationDto.
//
// Calculates 'tDur.DateDays', 'tDur.DateDaysNanosecs', 'tDur.Weeks', 'tDur.WeeksNanosecs',
// 'tDur.WeekDays' and 'tDur.WeekDaysNanosecs'.
//
// NOTE: (1) Before calling this method, ensure that TimeDurationDto.startDateTimeTz,
//           TimeDurationDto.endDateTimeTz and tDur.timeDuration are properly initialized.
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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	rd := int64(tDur.timeDuration)

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
// Seconds of duration using startTime, tDur.startDateTimeTz,
// and endTime, tDur.endDateTimeTz.DateTime.
//
//
// NOTE: (1) Before calling this method, ensure that tDur.startDateTimeTz,
//           TimeDurationDto.endDateTimeTz and TimeDurationDto.timeDuration
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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	rd := int64(tDur.timeDuration)

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
// NOTE: (1) Before calling this method, ensure that tDur.startDateTimeTz,
//           TimeDurationDto.endDateTimeTz and TimeDurationDto.timeDuration
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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	rd := int64(tDur.timeDuration)

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
// using the start and end dates, 'tDur.startDateTimeTz' and
// 'tDur.endDateTimeTz.DateTime'.
//
// NOTE: (1) Before calling this method, ensure that tDur.startDateTimeTz,
//           tDur.endDateTimeTz and tDur.timeDuration are properly initialized.
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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	startTime := tDur.startDateTimeTz.dateTimeValue
	endTime := tDur.endDateTimeTz.dateTimeValue

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() !=
		endTime.Location().String() {

		return fmt.Errorf(ePrefix+
			"Error: 'startTime' and 'endTime' Time Zone Location do NOT match!\n"+
			"startTimeZoneLocation='%v'\n"+
			"endTimeZoneLocation='%v'\n",
			startTime.Location().String(),
			endTime.Location().String())
	}

	rd := int64(tDur.timeDuration)

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
// NOTE: (1) Before calling this method, ensure that tDur.startDateTimeTz,
//           TimeDurationDto.endDateTimeTz and TimeDurationDto.timeDuration
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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	rd := int64(tDur.timeDuration)

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

// calcTimeDurationAllocations - Examines the input parameter 'timeDurCalcType' and
// then determines which type of time duration allocation calculation will be
// applied to the data fields of the current TimeDurationDto instance.
func (tDurDtoUtil *timeDurationDtoUtility) calcTimeDurationAllocations(
	tDur *TimeDurationDto,
	calcType TDurCalcType,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcTimeDurationAllocations() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(tDur, ePrefix)

	tDurDtoUtilX2.emptyTimeFields(
		&tDur2,
		ePrefix)

	tDur2.timeDurCalcType = TDurCalcType(0).CumDays()

	if tDur2.timeDuration == 0 {
		return nil
	}

	rd := int64(tDur2.timeDuration)

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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDurDtoUtilX2.emptyTimeFields(
		&tDur2,
		ePrefix)

	tDur2.timeDurCalcType = TDurCalcType(0).CumHours()

	if tDur2.timeDuration == 0 {
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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDurDtoUtilX2.emptyTimeFields(
		&tDur2,
		ePrefix)

	tDur2.timeDurCalcType = TDurCalcType(0).CumMinutes()

	if tDur2.timeDuration == 0 {
		return nil
	}

	rd := int64(tDur2.timeDuration)

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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDurDtoUtilX2.emptyTimeFields(&tDur2, ePrefix)

	tDur2.timeDurCalcType = TDurCalcType(0).CumMonths()

	if tDur2.timeDuration == 0 {
		return nil
	}

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
// tDur.timeDurCalcType = TDurCalcType(0).CumSeconds()
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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDurDtoUtilX2.emptyTimeFields(
		&tDur2,
		ePrefix)

	tDur2.timeDurCalcType = TDurCalcType(0).CumSeconds()

	if tDur2.timeDuration == 0 {
		return nil
	}

	rd := int64(tDur2.timeDuration)

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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDurDtoUtilX2.emptyTimeFields(&tDur2, ePrefix)

	tDur2.timeDurCalcType = TDurCalcType(0).CumWeeks()

	if tDur2.timeDuration == 0 {
		return nil
	}

	rd := int64(tDur2.timeDuration)

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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDurDtoUtilX2.emptyTimeFields(
		&tDur2,
		ePrefix)

	tDur2.timeDurCalcType = TDurCalcType(0).GregorianYears()

	if tDur2.timeDuration == 0 {
		return nil
	}

	rd := int64(tDur2.timeDuration)

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

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtilX2.copyOut(
		tDur,
		ePrefix)

	tDurDtoUtilX2.emptyTimeFields(
		&tDur2,
		ePrefix)

	tDur2.timeDurCalcType = TDurCalcType(0).StdYearMth()

	if tDur2.timeDuration == 0 {
		return nil
	}

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
// represented by years duration using input parameters 'tDur.startDateTimeTz' and
// 'tDur.endDateTimeTz'.
//
// NOTE: Before calling this method, ensure that tDur.startDateTimeTz,
//       tDur.endDateTimeTz and tDur.timeDuration are properly initialized.
//
func (tDurDtoUtil *timeDurationDtoUtility) calcYearsFromDuration(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.calcYearsFromDuration() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'tDur' is a nil pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	years := int64(0)
	yearNanosecs := int64(0)
	startTime := tDur.startDateTimeTz.dateTimeValue
	endTime := tDur.endDateTimeTz.dateTimeValue

	if endTime.Before(startTime) {
		return errors.New(ePrefix +
			"\nError: 'endTime' precedes, is less than, startTime!\n")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix+
			"\nError: 'startTime' and 'endTime' Time Zone Location do NOT match!\n"+
			"startTimeZoneLocation='%v'\n"+
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
		panic(ePrefix +
			"\nError: Input Parameter 'tDur' is a 'nil' pointer!")
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	tDur.startDateTimeTz = DateTzDto{}
	tDur.endDateTimeTz = DateTzDto{}
	tDur.timeDuration = time.Duration(0)
	tDur.timeDurCalcType = TDurCalcType(0).None()
	tDur.timeMathCalcMode = TCalcMode.None()
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

	return
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

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
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

// Equal - Compares two TimeDurationDto instances to determine
// if they are equivalent.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  t1Dur TimeDurationDto
//     - A TimeDurationDto object which will be
//       compared to the t2Dur TimeDurationDto
//       instance to determine if the two are
//       equivalent.
//
//  t2Dur TimeDurationDto
//     - A TimeDurationDto object which will be
//       compared to the t1Dur TimeDurationDto
//       instance to determine if the two are
//       equivalent.
//
//  ePrefix        string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// __________________________________________________________________________
//
// Return Values:
//
//  bool  - If 'true' it signals that all relevant data fields in
//          't1Dur' and 'tDur2' are equivalent. If data fields are
//          NOT equal or if an error is encountered, this value is
//          set to 'false'.
//
//  error - If either or both of 't1Dur' and 't2Dur' constitute nil pointers
//          the returned error object is configured with an appropriate
//          error message.
//
func (tDurDtoUtil *timeDurationDtoUtility) equal(
	t1Dur,
	t2Dur *TimeDurationDto,
	ePrefix string) (bool, error) {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.equal() "

	if t1Dur == nil {
		return false,
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "t1Dur",
				inputParameterValue: "",
				errMsg:              "Input parameter 't1Dur' is a 'nil' pointer!",
				err:                 nil,
			}
	}

	if t1Dur.lock == nil {
		t1Dur.lock = new(sync.Mutex)
	}

	if t2Dur == nil {
		return false,
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "t2Dur",
				inputParameterValue: "",
				errMsg:              "Input parameter 't2Dur' is a 'nil' pointer!",
				err:                 nil,
			}
	}

	if t2Dur.lock == nil {
		t2Dur.lock = new(sync.Mutex)
	}

	if !t1Dur.startDateTimeTz.Equal(t2Dur.startDateTimeTz) ||
		!t1Dur.endDateTimeTz.Equal(t2Dur.endDateTimeTz) ||
		t1Dur.timeDuration != t2Dur.timeDuration ||
		t1Dur.timeDurCalcType != t2Dur.timeDurCalcType ||
		t1Dur.Years != t2Dur.Years ||
		t1Dur.YearsNanosecs != t2Dur.YearsNanosecs ||
		t1Dur.Months != t2Dur.Months ||
		t1Dur.MonthsNanosecs != t2Dur.MonthsNanosecs ||
		t1Dur.Weeks != t2Dur.Weeks ||
		t1Dur.WeeksNanosecs != t2Dur.WeeksNanosecs ||
		t1Dur.WeekDays != t2Dur.WeekDays ||
		t1Dur.WeekDaysNanosecs != t2Dur.WeekDaysNanosecs ||
		t1Dur.DateDays != t2Dur.DateDays ||
		t1Dur.DateDaysNanosecs != t2Dur.DateDaysNanosecs ||
		t1Dur.Hours != t2Dur.Hours ||
		t1Dur.HoursNanosecs != t2Dur.HoursNanosecs ||
		t1Dur.Minutes != t2Dur.Minutes ||
		t1Dur.MinutesNanosecs != t2Dur.MinutesNanosecs ||
		t1Dur.Seconds != t2Dur.Seconds ||
		t1Dur.SecondsNanosecs != t2Dur.SecondsNanosecs ||
		t1Dur.Milliseconds != t2Dur.Milliseconds ||
		t1Dur.MillisecondsNanosecs != t2Dur.MillisecondsNanosecs ||
		t1Dur.Microseconds != t2Dur.Microseconds ||
		t1Dur.MicrosecondsNanosecs != t2Dur.MicrosecondsNanosecs ||
		t1Dur.Nanoseconds != t2Dur.Nanoseconds ||
		t1Dur.TotSubSecNanoseconds != t2Dur.TotSubSecNanoseconds ||
		t1Dur.TotDateNanoseconds != t2Dur.TotDateNanoseconds ||
		t1Dur.TotTimeNanoseconds != t2Dur.TotTimeNanoseconds {

		return false, nil
	}

	if t1Dur.timeMathCalcMode != t2Dur.timeMathCalcMode {
		return false, nil
	}

	return true, nil
}

// isEmpty() - Returns 'true' if the passed TimeDurationDto instance
// is uninitialized and consists entirely of zero values.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  t Dur TimeDurationDto
//     - A TimeDurationDto object which will be
//       will be analyzed to determine if all of
//       its constituent data fields are empty or
//       contain a zero value.
//
//  ePrefix        string
//     - The error prefix containing the names of
//       all the methods executed up to this point.
//
// __________________________________________________________________________
//
// Return Values:
//
//  bool  - If 'true' it signals that all relevant data fields in
//          'tDur'are empty or contain a zero value. If any of the
//          data fields are populated, or if an error is encountered,
//          this value is set to 'false'.
//
//  error - If 'tDur' is a nil pointer the returned error object is
//          configured with an appropriate error message.
//
func (tDurDtoUtil *timeDurationDtoUtility) isEmpty(
	tDur *TimeDurationDto,
	ePrefix string) (bool, error) {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.isEmpty() "

	if tDur == nil {
		return false,
			&InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "tDur",
				inputParameterValue: "",
				errMsg:              "Input parameter 'tDur' is a 'nil' pointer!",
				err:                 nil,
			}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	if tDur.startDateTimeTz.IsEmpty() &&
		tDur.endDateTimeTz.IsEmpty() &&
		tDur.timeDuration == 0 &&
		tDur.Years == 0 &&
		tDur.YearsNanosecs == 0 &&
		tDur.Months == 0 &&
		tDur.MonthsNanosecs == 0 &&
		tDur.Weeks == 0 &&
		tDur.WeeksNanosecs == 0 &&
		tDur.WeekDays == 0 &&
		tDur.WeekDaysNanosecs == 0 &&
		tDur.DateDays == 0 &&
		tDur.DateDaysNanosecs == 0 &&
		tDur.Hours == 0 &&
		tDur.HoursNanosecs == 0 &&
		tDur.Minutes == 0 &&
		tDur.MinutesNanosecs == 0 &&
		tDur.Seconds == 0 &&
		tDur.SecondsNanosecs == 0 &&
		tDur.Milliseconds == 0 &&
		tDur.MillisecondsNanosecs == 0 &&
		tDur.Microseconds == 0 &&
		tDur.MicrosecondsNanosecs == 0 &&
		tDur.Nanoseconds == 0 &&
		tDur.TotSubSecNanoseconds == 0 &&
		tDur.TotDateNanoseconds == 0 &&
		tDur.TotTimeNanoseconds == 0 {

		tDur.timeDurCalcType = TDurCalcType(0).None()
		return true, nil
	}

	return false, nil
}

// isValid - Returns an error value signaling whether
// the data fields of input parameter 'tDur' are valid.
// __________________________________________________________________________
//
// Input Parameters:
//
// tDur     *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
//
// ePrefix  string
//            - The error prefix containing the names of all
//              the methods executed up to this point.
//
// __________________________________________________________________________
//
// Return Values:
//
//  error - If the 'tDur' TimeDurationDto is invalid, this returned
//          error object is configured with an appropriate error
//          message.
//
//          If the 'tDur' TimeDurationDto is valid, this returned
//          error object is set to 'nil'.
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

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	if tDur.startDateTimeTz.dateTimeValue.IsZero() &&
		tDur.endDateTimeTz.dateTimeValue.IsZero() {

		return fmt.Errorf(ePrefix +
			"\nError: Both Start and End Times are Zero!\n")

	}

	if tDur.endDateTimeTz.dateTimeValue.Before(tDur.startDateTimeTz.dateTimeValue) {
		return fmt.Errorf(ePrefix +
			"\nError: End Time is Before Start Time!\n")
	}

	if tDur.timeMathCalcMode < TCalcMode.XFirstValidCalcType() ||
		tDur.timeMathCalcMode > TCalcMode.XLastValidCalcType() {
		return fmt.Errorf(ePrefix +
			"\nError: Time Math Calculation Mode is INVALID!\n")
	}

	return nil
}

// reCalcTimeDurationAllocation - Re-calculates and allocates time duration for the current
// TimeDurationDto instance over the various time components (years, months, weeks, weekdays,
// datedays, hour, minutes, seconds, milliseconds, microseconds and nanoseconds) depending
// on the value of the 'TDurCalcType' input parameter.
// __________________________________________________________________________
//
// Input Parameters
//
//
// tDur    *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
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
// ePrefix string
//           - The error prefix containing the names of all
//             the methods executed up to this point.
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
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

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:              "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	tDurDtoUtil2 := timeDurationDtoUtility{}

	tDur2 := tDurDtoUtil2.copyOut(tDur, ePrefix)

	tDur2.timeDurCalcType = tDurCalcType

	err := tDurDtoUtil2.calcTimeDurationAllocations(
		&tDur2,
		tDur2.timeDurCalcType,
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
// IMPORTANT:
// This method assumes that the starting date time, and time duration calculation
// type were previous set to valid values.
//
//
// Input Parameters
// ================
//
//  tDur     *TimeDurationDto
//     - The data fields of this TimeDurationDto object will be
//       set with the results generated by this method.
//
//
//  tDurCalcType TDurCalcType
//       - Specifies the calculation type to be used in allocating
//         time duration:
//
//         TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                            day time calculation.
//
//         TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//         TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//         TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//         TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//         TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                         or hours.
//
//         TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                            hours or minutes.
//
//         TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                            Used for very large duration values.
//
//         Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation string
//       - Designates the standard Time Zone location by which
//         time duration will be compared. This ensures that
//         'oranges are compared to oranges and apples are compared
//         to apples' with respect to start time and end time duration
//         calculations.
//
//         If 'timeZoneLocation' is passed as an empty string, it
//         will be automatically defaulted to the 'UTC' time zone.
//         Reference Universal Coordinated Time:
//            https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//         Time zone location must be designated as one of three types of
//         time zones.
//
//         (1) The time zone "Local", which Golang accepts as
//             the time zone currently configured on the host
//             computer.
//
//         (2) IANA Time Zone - A valid IANA Time Zone from the
//             IANA database.
//             See https://golang.org/pkg/time/#LoadLocation
//             and https://www.iana.org/time-zones to ensure that
//             the IANA Time Zone Database is properly configured
//             on your system.
//
//             IANA Time Zone Examples:
//               "America/New_York"
//               "America/Chicago"
//               "America/Denver"
//               "America/Los_Angeles"
//               "Pacific/Honolulu"
//               "Etc/UTC" = GMT or UTC
//
//         (3) A Military Time Zone
//             Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//              Examples:
//                "Alpha"   or "A"
//                "Bravo"   or "B"
//                "Charlie" or "C"
//                "Delta"   or "D"
//                "Zulu"    or "Z"
//
//                If the time zone "Zulu" is passed to this method, it will be
//                classified as a Military Time Zone.
//
//
//  timeCalcMode  TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
//       - A date time format string which will be used
//         to format and display 'dateTime'. Example:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         If 'dateTimeFmtStr' is submitted as an
//         'empty string', a default date time format
//         string will be applied. The default date time
//         format string is:
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  ePrefix  string
//       - The error prefix containing the names of all
//         the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setAutoEnd(
	tDur *TimeDurationDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setAutoEnd() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:              "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	if timeMathCalcMode < TCalcMode.XFirstValidCalcType() ||
		timeMathCalcMode > TCalcMode.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeMathCalcMode",
			inputParameterValue: timeMathCalcMode.String(),
			errMsg:              "Input Parameter 'timeMathCalcMode' is INVALID!",
			err:                 nil,
		}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tDur2 := TimeDurationDto{}

	tDur2.lock = new(sync.Mutex)

	var err error

	tDur2.timeDuration,
		tDur2.startDateTimeTz,
		tDur2.endDateTimeTz,
		err = dtMech.ComputeDurationFromStartEndTimes(
		tDur.startDateTimeTz.dateTimeValue,
		time.Now().UTC(),
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.timeDurCalcType = tDurCalcType

	tDur2.timeMathCalcMode = timeMathCalcMode

	tDurDtoUtil2 := timeDurationDtoUtility{}

	err = tDurDtoUtil2.calcTimeDurationAllocations(
		&tDur2,
		tDur2.timeDurCalcType,
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtil2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// setAutoStart - Receives a pointer to a TimeDurationDto instance
// and proceeds to set the starting date time.  The caller is
// responsible for setting the ending date time and computing the
// time duration between starting date time and ending date time.
//
// As a temporary place holder, this method assigns time.Now() plus
// 2-days as the ending date time.
//
// Input Parameters
// ================
//
//  tDur *TimeDurationDto
//       - The data fields of this TimeDurationDto object will be
//         set with the results generated by this method.
//
//
//  tDurCalcType TDurCalcType
//       - Specifies the calculation type to be used in allocating
//         time duration:
//
//         TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                            day time calculation.
//
//         TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//         TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//         TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//         TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//         TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                         or hours.
//
//         TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                            hours or minutes.
//
//         TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                            Used for very large duration values.
//
//         Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation string
//       - Designates the standard Time Zone location by which
//         time duration will be compared. This ensures that
//         'oranges are compared to oranges and apples are compared
//         to apples' with respect to start time and end time duration
//         calculations.
//
//         If 'timeZoneLocation' is passed as an empty string, it
//         will be automatically defaulted to the 'UTC' time zone.
//         Reference Universal Coordinated Time:
//            https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//         Time zone location must be designated as one of three types of
//         time zones.
//
//         (1) The time zone "Local", which Golang accepts as
//             the time zone currently configured on the host
//             computer.
//
//         (2) IANA Time Zone - A valid IANA Time Zone from the
//             IANA database.
//             See https://golang.org/pkg/time/#LoadLocation
//             and https://www.iana.org/time-zones to ensure that
//             the IANA Time Zone Database is properly configured
//             on your system.
//
//             IANA Time Zone Examples:
//               "America/New_York"
//               "America/Chicago"
//               "America/Denver"
//               "America/Los_Angeles"
//               "Pacific/Honolulu"
//               "Etc/UTC" = GMT or UTC
//
//         (3) A Military Time Zone
//             Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//              Examples:
//                "Alpha"   or "A"
//                "Bravo"   or "B"
//                "Charlie" or "C"
//                "Delta"   or "D"
//                "Zulu"    or "Z"
//
//                If the time zone "Zulu" is passed to this method, it will be
//                classified as a Military Time Zone.
//
//
//  timeCalcMode  TimeMathCalcMode
//       - TimeMathCalcMode is an enumeration which specifies the algorithm
//         which will be used when computing time spans or time duration.
//
//         If 'LocalTimeZone' is specified, days are defined as local time
//         zone days which may be less than, or greater than, 24-hours due
//         to local conventions like daylight savings time.
//         (TCalcMode.LocalTimeZone())
//
//         If 'UtcTimeZone' is specified, days are uniformly defined as
//         a time span consisting of 24-consecutive hours.
//         (TCalcMode.UtcTimeZone())
//
//         For additional information see the type documentation at
//               datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
//       - A date time format string which will be used
//         to format and display 'dateTime'. Example:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         If 'dateTimeFmtStr' is submitted as an
//         'empty string', a default date time format
//         string will be applied. The default date time
//         format string is:
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// ePrefix string
//       - The error prefix containing the names of all
//         the methods executed up to this point.
//
//
func (tDurDtoUtil *timeDurationDtoUtility) setAutoStart(
	tDur *TimeDurationDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setAutoStart() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:              "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	if timeMathCalcMode < TCalcMode.XFirstValidCalcType() ||
		timeMathCalcMode > TCalcMode.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeMathCalcMode",
			inputParameterValue: timeMathCalcMode.String(),
			errMsg:              "Input Parameter 'timeMathCalcMode' is INVALID!",
			err:                 nil,
		}
	}

	dtMech := DTimeMechanics{}

	tDur2 := TimeDurationDto{}

	tDur2.lock = new(sync.Mutex)

	var err error

	// Place holder for ending date time
	endDateTime := time.Now().AddDate(
		0,
		0,
		2)

	startDateTimeNow := time.Now().UTC()

	tDur2.timeDuration,
		tDur2.startDateTimeTz,
		tDur2.endDateTimeTz,
		err = dtMech.ComputeDurationFromStartEndTimes(
		startDateTimeNow,
		endDateTime,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.timeDurCalcType = tDurCalcType

	tDur2.timeMathCalcMode = timeMathCalcMode

	tDurDtoUtil2 := timeDurationDtoUtility{}

	err = tDurDtoUtil2.calcTimeDurationAllocations(
		&tDur2,
		tDur2.timeDurCalcType,
		ePrefix)

	if err != nil {
		return err
	}

	tDurDtoUtil2.copyIn(tDur, &tDur2, ePrefix)

	return nil
}

// setEndTimeMinusTimeDtoCalcTz - Sets start date time, end date time and duration
// based on an ending date time, and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  tDur            *TimeDurationDto
//     - The data fields of this TimeDurationDto object will be
//       set with the results generated by this method.
//
//
// endDateTime      time.Time
//     - Ending date time. The starting date time will be computed
//       by subtracting minusTimeDto from 'endDateTime'
//
//
// minusTimeDto     TimeDto
//     - Time components (Years, months, weeks, days, hours etc.)
//       which will be subtracted from 'endDateTime' to compute
//       time duration and starting date time.
//
//       type TimeDto struct {
//        Years          int // Number of Years
//        Months         int // Number of Months
//        Weeks          int // Number of Weeks
//        WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//        DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//        Hours          int // Number of Hours.
//        Minutes        int // Number of Minutes
//        Seconds        int // Number of Seconds
//        Milliseconds   int // Number of Milliseconds
//        Microseconds   int // Number of Microseconds
//        Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//        TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                 //  plus remaining Nanoseconds
//       }
//
//
// tDurCalcType     TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration:
//
//       TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                          day time calculation.
//
//       TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                          or hours.
//
//       TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                          hours or minutes.
//
//       TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                      Used for very large duration values.
//
//       Type 'TDurCalcType' is located in source file:
//        MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
// timeZoneLocation string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time comparisons.
//
//       If 'timeZoneLocation' is passed as an empty string, it
//       will be automatically defaulted to the 'UTC' time zone.
//       Reference Universal Coordinated Time:
//          https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//       Time zone location must be designated as one of three types of
//       time zones.
//
//       (1) The time zone "Local", which Golang accepts as
//           the time zone currently configured on the host
//           computer.
//
//       (2) IANA Time Zone - A valid IANA Time Zone from the
//           IANA database.
//           See https://golang.org/pkg/time/#LoadLocation
//           and https://www.iana.org/time-zones to ensure that
//           the IANA Time Zone Database is properly configured
//           on your system.
//
//           IANA Time Zone Examples:
//             "America/New_York"
//             "America/Chicago"
//             "America/Denver"
//             "America/Los_Angeles"
//             "Pacific/Honolulu"
//             "Etc/UTC" = GMT or UTC
//
//       (3) A Military Time Zone
//           Reference:
//             https://en.wikipedia.org/wiki/List_of_military_time_zones
//             http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//             https://www.timeanddate.com/time/zones/military
//             https://www.timeanddate.com/worldclock/timezone/alpha
//             https://www.timeanddate.com/time/map/
//
//            Examples:
//              "Alpha"   or "A"
//              "Bravo"   or "B"
//              "Charlie" or "C"
//              "Delta"   or "D"
//              "Zulu"    or "Z"
//
//              If the time zone "Zulu" is passed to this method, it will be
//              classified as a Military Time Zone.
//
//
//  timeCalcMode    TimeMathCalcMode
//     - TimeMathCalcMode is an enumeration which specifies the algorithm
//       which will be used when computing time spans or time duration.
//
//       If 'LocalTimeZone' is specified, days are defined as local time
//       zone days which may be less than, or greater than, 24-hours due
//       to local conventions like daylight savings time.
//       (TCalcMode.LocalTimeZone())
//
//       If 'UtcTimeZone' is specified, days are uniformly defined as
//       a time span consisting of 24-consecutive hours.
//       (TCalcMode.UtcTimeZone())
//
//       For additional information see the type documentation at
//             datetime\timemathcalcmode.go
//
//
// dateTimeFmtStr   string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ePrefix          string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setEndTimeMinusTimeDtoCalcTz(
	tDur *TimeDurationDto,
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
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

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	err := minusTimeDto.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Input Parameter 'minusTimeDto' is INVALID!\n"+
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
			errMsg:              "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	if timeCalcMode < TCalcMode.XFirstValidCalcType() ||
		timeCalcMode > TCalcMode.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeCalcMode",
			inputParameterValue: "",
			errMsg: "Input parameter 'timeCalcMode' " +
				"is invalid!",
			err: nil,
		}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tDur2 := TimeDurationDto{}

	tDur2.lock = new(sync.Mutex)

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTimeTzName(
		&tDur2.endDateTimeTz,
		endDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.startDateTimeTz, err =
		dTzUtil.addMinusTimeDto(
			&tDur2.endDateTimeTz,
			timeCalcMode,
			minusTimeDto,
			ePrefix)

	if err != nil {
		return err
	}

	tDur2.timeDuration =
		tDur2.endDateTimeTz.dateTimeValue.Sub(
			tDur2.startDateTimeTz.dateTimeValue)

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

// setStartEndTimes - Sets data field values for the current TimeDurationDto
// instance using a Start Date Time, End Date Time and a time zone specification.
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
//
// All data fields in the current TimeDurationDto instance are overwritten with
// the new time duration values.
//
// Input Parameters
// ================
//
//  tDur *TimeDurationDto
//       - The data fields of this TimeDurationDto object will be
//         set with the results generated by this method.
//
//  startDateTime time.Time
//       - Starting time
//
//  endDateTime  time.Time
//       - Ending time
//
//
//  tDurCalcType TDurCalcType
//       - Specifies the calculation type to be used in allocating
//         time duration:
//
//         TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                            day time calculation.
//
//         TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//         TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//         TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//         TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//         TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                         or hours.
//
//         TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                            hours or minutes.
//
//         TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                            Used for very large duration values.
//
//         Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//  timeZoneLocation string
//       - Designates the standard Time Zone location by which
//         time duration will be compared. This ensures that
//         'oranges are compared to oranges and apples are compared
//         to apples' with respect to start time and end time duration
//         calculations.
//
//         If 'timeZoneLocation' is passed as an empty string, it
//         will be automatically defaulted to the 'UTC' time zone.
//         Reference Universal Coordinated Time:
//            https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//         Time zone location must be designated as one of three types of
//         time zones.
//
//         (1) The time zone "Local", which Golang accepts as
//             the time zone currently configured on the host
//             computer.
//
//         (2) IANA Time Zone - A valid IANA Time Zone from the
//             IANA database.
//             See https://golang.org/pkg/time/#LoadLocation
//             and https://www.iana.org/time-zones to ensure that
//             the IANA Time Zone Database is properly configured
//             on your system.
//
//             IANA Time Zone Examples:
//               "America/New_York"
//               "America/Chicago"
//               "America/Denver"
//               "America/Los_Angeles"
//               "Pacific/Honolulu"
//               "Etc/UTC" = GMT or UTC
//
//         (3) A Military Time Zone
//             Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//              Examples:
//                "Alpha"   or "A"
//                "Bravo"   or "B"
//                "Charlie" or "C"
//                "Delta"   or "D"
//                "Zulu"    or "Z"
//
//                If the time zone "Zulu" is passed to this method, it will be
//                classified as a Military Time Zone.
//
//
//  timeCalcMode  TimeMathCalcMode
//       - TimeMathCalcMode is an enumeration which specifies the algorithm
//         which will be used when computing time spans or time duration.
//
//         If 'LocalTimeZone' is specified, days are defined as local time
//         zone days which may be less than, or greater than, 24-hours due
//         to local conventions like daylight savings time.
//         (TCalcMode.LocalTimeZone())
//
//         If 'UtcTimeZone' is specified, days are uniformly defined as
//         a time span consisting of 24-consecutive hours.
//         (TCalcMode.UtcTimeZone())
//
//         For additional information see the type documentation at
//               datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
//       - A date time format string which will be used
//         to format and display 'dateTime'. Example:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         If 'dateTimeFmtStr' is submitted as an
//         'empty string', a default date time format
//         string will be applied. The default date time
//         format string is:
//         FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// ePrefix string
//       - The error prefix containing the names of all
//         the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setStartEndTimes(
	tDur *TimeDurationDto,
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setStartEndTimes() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Both 'startDateTime' and 'endDateTime' " +
			"input time parameters are ZERO!\n")
	}

	var tempStartDateTime, tempEndDateTime time.Time

	// If endDateTime is less than startDateTime
	// reverse the order.
	if endDateTime.Before(startDateTime) {
		tempStartDateTime = endDateTime
		tempEndDateTime = startDateTime
	} else {
		tempStartDateTime = startDateTime
		tempEndDateTime = endDateTime
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:              "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	if timeMathCalcMode < TCalcMode.XFirstValidCalcType() ||
		timeMathCalcMode > TCalcMode.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeMathCalcMode",
			inputParameterValue: timeMathCalcMode.String(),
			errMsg:              "Input Parameter 'timeMathCalcMode' is INVALID!",
			err:                 nil,
		}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tDur2 := TimeDurationDto{}

	tDur2.lock = new(sync.Mutex)

	var err error

	tDur2.timeDuration,
		tDur2.startDateTimeTz,
		tDur2.endDateTimeTz,
		err = dtMech.ComputeDurationFromStartEndTimes(
		tempStartDateTime,
		tempEndDateTime,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.timeDurCalcType = tDurCalcType

	tDur2.timeMathCalcMode = timeMathCalcMode

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

// setStartEndTimesTz - Sets data field values for the current
// TimeDurationDto instance using a Start Date Time, End Date Time and a
// time zone specification.
//
// The Starting Date Time and Ending Date Time are submitted as type 'DateTzDto'
//
// First, 'startDateTime' and 'endDateTime' are converted to the designate Time
// Zone Location. Next, 'startDateTime' is subtracted from 'endDateTime' to compute
// time duration.
//
// The user is required to submit input parameters for time zone location and
// date time calculation type.
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
//  startDateTime DateTzDto
//       - Starting date time
//
//
//  endDateTime   DateTzDto
//       - Ending date time
//
//
//  tDurCalcType TDurCalcType
//       - Specifies the calculation type to be used in allocating
//         time duration:
//
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
//  timeZoneLocation string
//       - Designates the standard Time Zone location by which
//         time duration will be compared. This ensures that
//         'oranges are compared to oranges and apples are compared
//         to apples' with respect to start time and end time duration
//         calculations.
//
//         If 'timeZoneLocation' is passed as an empty string, it
//         will be automatically defaulted to the 'UTC' time zone.
//         Reference Universal Coordinated Time:
//            https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//         Time zone location must be designated as one of three types of
//         time zones.
//
//         (1) The time zone "Local", which Golang accepts as
//             the time zone currently configured on the host
//             computer.
//
//         (2) IANA Time Zone - A valid IANA Time Zone from the
//             IANA database.
//             See https://golang.org/pkg/time/#LoadLocation
//             and https://www.iana.org/time-zones to ensure that
//             the IANA Time Zone Database is properly configured
//             on your system.
//
//             IANA Time Zone Examples:
//               "America/New_York"
//               "America/Chicago"
//               "America/Denver"
//               "America/Los_Angeles"
//               "Pacific/Honolulu"
//               "Etc/UTC" = GMT or UTC
//
//         (3) A Military Time Zone
//             Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//              Examples:
//                "Alpha"   or "A"
//                "Bravo"   or "B"
//                "Charlie" or "C"
//                "Delta"   or "D"
//                "Zulu"    or "Z"
//
//                If the time zone "Zulu" is passed to this method, it will be
//                classified as a Military Time Zone.
//
//
//  timeCalcMode  TimeMathCalcMode
//       - TimeMathCalcMode is an enumeration which specifies the algorithm
//         which will be used when computing time spans or time duration.
//
//         If 'LocalTimeZone' is specified, days are defined as local time
//         zone days which may be less than, or greater than, 24-hours due
//         to local conventions like daylight savings time.
//         (TCalcMode.LocalTimeZone())
//
//         If 'UtcTimeZone' is specified, days are uniformly defined as
//         a time span consisting of 24-consecutive hours.
//         (TCalcMode.UtcTimeZone())
//
//         For additional information see the type documentation at
//               datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
//       - A date time format string which will be used
//         to format and display 'dateTime'. Example:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         If 'dateTimeFmtStr' is submitted as an
//         'empty string', a default date time format
//         string will be applied. The default date time
//         format string is:
//            FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ePrefix string
//       - The error prefix containing the names of all
//         the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setStartEndTimesTz(
	tDur *TimeDurationDto,
	startDateTimeTz,
	endDateTimeTz DateTzDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setStartEndTimesTz() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	if startDateTimeTz.dateTimeValue.IsZero() && endDateTimeTz.dateTimeValue.IsZero() {
		return errors.New(ePrefix +
			"\nError: Both 'startDateTimeTz' and 'endDateTimeTz' " +
			"input time parameters have ZERO time values!\n")
	}

	dTzUtil := dateTzDtoUtility{}

	err := dTzUtil.isValidDateTzDto(
		&startDateTimeTz,
		ePrefix)

	if err != nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "startDateTimeTz",
			inputParameterValue: "",
			errMsg: fmt.Sprintf(
				"Input Parameter 'startDateTimeTz' is INVALID!\n"+
					"Validation Error='%v'", err.Error()),
			err: nil,
		}
	}

	err = dTzUtil.isValidDateTzDto(
		&endDateTimeTz,
		ePrefix)

	if err != nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "endDateTimeTz",
			inputParameterValue: "",
			errMsg: fmt.Sprintf(
				"Input Parameter 'endDateTimeTz' is INVALID!\n"+
					"Validation Error='%v'", err.Error()),
			err: nil,
		}
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:              "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	if timeMathCalcMode < TCalcMode.XFirstValidCalcType() ||
		timeMathCalcMode > TCalcMode.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeMathCalcMode",
			inputParameterValue: timeMathCalcMode.String(),
			errMsg:              "Input Parameter 'timeMathCalcMode' is INVALID!",
			err:                 nil,
		}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tDur2 := TimeDurationDto{}

	tDur2.lock = new(sync.Mutex)

	tDur2.timeDuration,
		tDur2.startDateTimeTz,
		tDur2.endDateTimeTz,
		err = dtMech.ComputeDurationFromStartEndTimes(
		startDateTimeTz.dateTimeValue,
		endDateTimeTz.dateTimeValue,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.timeDurCalcType = tDurCalcType

	tDur2.timeMathCalcMode = timeMathCalcMode

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

// setStartTimeDuration - Sets start time, end time and duration for the
// current TimeDurationDto instance. 'startDateTime' is converted to the
// specified 'timeZoneLocation'. The duration value is added to 'startDateTime'
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time,
// while the actual starting date time is computed by
// subtracting duration.
//
// Input Parameters:
// =================
//
//  tDur *TimeDurationDto
//       - The data fields of this TimeDurationDto object will be
//         set with the results generated by this method.
//
//
//  startDateTime time.Time
//       - Starting date time for the duration calculation
//
//
//  duration  time.Duration
//       - Amount of time to be added to or subtracted from
//         'startDateTime'. Note: If duration is a negative value
//         'startDateTime' is converted to ending date time and
//         actual starting date time is computed by subtracting
//         duration.
//
//
//  tDurCalcType TDurCalcType
//       - Specifies the calculation type to be used in allocating
//         time duration:
//
//         TDurCalcType(0).StdYearMth()   - Default - standard year, month week,
//                                          day time calculation.
//
//         TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//         TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//         TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//         TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//         TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                            or hours.
//
//         TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                            hours or minutes.
//
//         TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                            Used for very large duration values.
//
//         Type 'TDurCalcType' is located in source file:
//          MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation string
//       - Designates the standard Time Zone location by which
//         time duration will be compared. This ensures that
//         'oranges are compared to oranges and apples are compared
//         to apples' with respect to start time and end time duration
//         calculations.
//
//         If 'timeZoneLocation' is passed as an empty string, it
//         will be automatically defaulted to the 'UTC' time zone.
//         Reference Universal Coordinated Time:
//            https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//         Time zone location must be designated as one of three types of
//         time zones.
//
//         (1) The time zone "Local", which Golang accepts as
//             the time zone currently configured on the host
//             computer.
//
//         (2) IANA Time Zone - A valid IANA Time Zone from the
//             IANA database.
//             See https://golang.org/pkg/time/#LoadLocation
//             and https://www.iana.org/time-zones to ensure that
//             the IANA Time Zone Database is properly configured
//             on your system.
//
//             IANA Time Zone Examples:
//               "America/New_York"
//               "America/Chicago"
//               "America/Denver"
//               "America/Los_Angeles"
//               "Pacific/Honolulu"
//               "Etc/UTC" = GMT or UTC
//
//         (3) A Military Time Zone
//             Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//              Examples:
//                "Alpha"   or "A"
//                "Bravo"   or "B"
//                "Charlie" or "C"
//                "Delta"   or "D"
//                "Zulu"    or "Z"
//
//                If the time zone "Zulu" is passed to this method, it will be
//                classified as a Military Time Zone.
//
//
//  timeCalcMode  TimeMathCalcMode
//       - TimeMathCalcMode is an enumeration which specifies the algorithm
//         which will be used when computing time spans or time duration.
//
//         If 'LocalTimeZone' is specified, days are defined as local time
//         zone days which may be less than, or greater than, 24-hours due
//         to local conventions like daylight savings time.
//         (TCalcMode.LocalTimeZone())
//
//         If 'UtcTimeZone' is specified, days are uniformly defined as
//         a time span consisting of 24-consecutive hours.
//         (TCalcMode.UtcTimeZone())
//
//         For additional information see the type documentation at
//               datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
//       - A date time format string which will be used
//         to format and display 'dateTime'. Example:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         If 'dateTimeFmtStr' is submitted as an
//         'empty string', a default date time format
//         string will be applied. The default date time
//         format string is:
//           FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// ePrefix string
//       - The error prefix containing the names of all
//         the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setStartTimeDuration(
	tDur *TimeDurationDto,
	baseDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr string,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setStartTimeDuration() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	if baseDateTime.IsZero() && duration == 0 {
		return errors.New(ePrefix +
			"\nError: Both 'baseDateTime' and 'duration' " +
			"input parameters are ZERO!\n")
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:              "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tDur2 := TimeDurationDto{}

	tDur2.lock = new(sync.Mutex)

	var err error

	tDur2.startDateTimeTz,
		tDur2.endDateTimeTz,
		err = dtMech.ComputeDurationFromBaseTime(
		baseDateTime,
		duration,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.timeDuration = duration

	tDur2.timeDurCalcType = tDurCalcType

	tDur2.timeMathCalcMode = timeMathCalcMode

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

// setStartTimeTzDuration - Sets start time, end time and
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
//  tDur *TimeDurationDto
//       - The data fields of this TimeDurationDto object will be
//         set with the results generated by this method.
//
//
//  startDateTime DateTzDto
//       - Provides starting date time for the duration calculation
//
//
//  duration  time.Duration
//       - Amount of time to be added to or subtracted from
//         'startDateTime'. Note: If duration is a negative value
//         'startDateTime' is converted to ending date time and
//         actual starting date time is computed by subtracting
//         duration.
//
//
//  tDurCalcType TDurCalcType
//       - Specifies the calculation type to be used in allocating
//         time duration:
//
//         TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                            day time calculation.
//
//         TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//         TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//         TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//         TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//         TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                            or hours.
//
//         TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                            hours or minutes.
//
//         TDurCalcType(0).GregorianYears() - Computes Years based on average length of a Gregorian Year
//                                            Used for very large duration values.
//
//         Type 'TDurCalcType' is located in source file:
//          MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
// timeZoneLocation string
//       - Designates the standard Time Zone location by which
//         time duration will be compared. This ensures 'oranges
//         are compared to oranges and apples are compared to apples'
//         with respect to start time and end time duration calculations.
//
//         If 'timeZoneLocation' is passed as an empty string, it
//         will be automatically defaulted to the 'UTC' time zone.
//         Reference Universal Coordinated Time:
//            https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//         Time zone location must be designated as one of three types of
//         time zones.
//
//         (1) The time zone "Local", which Golang accepts as
//             the time zone currently configured on the host
//             computer.
//
//         (2) IANA Time Zone - A valid IANA Time Zone from the
//             IANA database.
//             See https://golang.org/pkg/time/#LoadLocation
//             and https://www.iana.org/time-zones to ensure that
//             the IANA Time Zone Database is properly configured
//             on your system.
//
//             IANA Time Zone Examples:
//               "America/New_York"
//               "America/Chicago"
//               "America/Denver"
//               "America/Los_Angeles"
//               "Pacific/Honolulu"
//               "Etc/UTC" = GMT or UTC
//
//         (3) A Military Time Zone
//             Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//              Examples:
//                "Alpha"   or "A"
//                "Bravo"   or "B"
//                "Charlie" or "C"
//                "Delta"   or "D"
//                "Zulu"    or "Z"
//
//                If the time zone "Zulu" is passed to this method, it will be
//                classified as a Military Time Zone.
//
//
//  timeCalcMode  TimeMathCalcMode
//       - TimeMathCalcMode is an enumeration which specifies the algorithm
//         which will be used when computing time spans or time duration.
//
//         If 'LocalTimeZone' is specified, days are defined as local time
//         zone days which may be less than, or greater than, 24-hours due
//         to local conventions like daylight savings time.
//         (TCalcMode.LocalTimeZone())
//
//         If 'UtcTimeZone' is specified, days are uniformly defined as
//         a time span consisting of 24-consecutive hours.
//         (TCalcMode.UtcTimeZone())
//
//         For additional information see the type documentation at
//               datetime\timemathcalcmode.go
//
//
//  timeCalcMode  TimeMathCalcMode
//       - TimeMathCalcMode is an enumeration which specifies the algorithm
//         which will be used when computing time spans or time duration.
//
//         If 'LocalTimeZone' is specified, days are defined as local time
//         zone days which may be less than, or greater than, 24-hours due
//         to local conventions like daylight savings time.
//         (TCalcMode.LocalTimeZone())
//
//         If 'UtcTimeZone' is specified, days are uniformly defined as
//         a time span consisting of 24-consecutive hours.
//         (TCalcMode.UtcTimeZone())
//
//         For additional information see the type documentation at
//               datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr string
//       - A date time format string which will be used
//         to format and display 'dateTime'. Example:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//        If 'dateTimeFmtStr' is submitted as an
//        'empty string', a default date time format
//        string will be applied. The default date time
//        format string is:
//           FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  ePrefix string
//       - The error prefix containing the names of all
//         the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setStartTimeTzDuration(
	tDur *TimeDurationDto,
	startDateTimeTz DateTzDto,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeMathCalcMode TimeMathCalcMode,
	dateTimeFmtStr,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setStartTimeTzDuration() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
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
			errMsg:              "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	if timeMathCalcMode < TCalcMode.XFirstValidCalcType() ||
		timeMathCalcMode > TCalcMode.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeMathCalcMode",
			inputParameterValue: timeMathCalcMode.String(),
			errMsg:              "Input Parameter 'timeMathCalcMode' is INVALID!",
			err:                 nil,
		}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tDur2 := TimeDurationDto{}

	tDur2.lock = new(sync.Mutex)

	var err error

	tDur2.startDateTimeTz,
		tDur2.endDateTimeTz,
		err = dtMech.ComputeDurationFromBaseTime(
		startDateTimeTz.dateTimeValue,
		duration,
		timeZoneLocation,
		timeMathCalcMode,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.timeDuration = duration

	tDur2.timeDurCalcType = tDurCalcType

	tDur2.timeMathCalcMode = timeMathCalcMode

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

// setStartTimePlusTimeDto - Sets start date time, end date time and duration
// based on a starting date time and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// Input Parameters:
// =================
//
//  tDur              *TimeDurationDto
//           - The data fields of this TimeDurationDto object will be
//             set with the results generated by this method.
//
//
//  startDateTime     time.Time
//            - Starting date time. The ending date time will be computed
//              by adding the time components of the 'plusTimeDto' to
//              'startDateTime'.
//
//
//  plusTimeDto      TimeDto
//            -  Time components (Years, months, weeks, days, hours etc.)
//               which will be added to 'startDateTime' to compute
//               time duration and ending date time.
//
//               type TimeDto struct {
//                Years          int // Number of Years
//                Months         int // Number of Months
//                Weeks          int // Number of Weeks
//                WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//                DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//                Hours          int // Number of Hours.
//                Minutes        int // Number of Minutes
//                Seconds        int // Number of Seconds
//                Milliseconds   int // Number of Milliseconds
//                Microseconds   int // Number of Microseconds
//                Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//                TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                         //  plus remaining Nanoseconds
//               }
//
//
//  tDurCalcType      TDurCalcType
//       - Specifies the calculation type to be used in allocating
//         time duration:
//
//         TDurCalcType(0).StdYearMth()     - Default - standard year, month week,
//                                            day time calculation.
//
//         TDurCalcType(0).CumMonths()      - Computes cumulative months - no Years.
//
//         TDurCalcType(0).CumWeeks()       - Computes cumulative weeks. No Years or months
//
//         TDurCalcType(0).CumDays()        - Computes cumulative days. No Years, months or weeks.
//
//         TDurCalcType(0).CumHours()       - Computes cumulative hours. No Years, months, weeks or days.
//
//         TDurCalcType(0).CumMinutes()     - Computes cumulative minutes. No Years, months, weeks, days
//                                            or hours.
//
//         TDurCalcType(0).CumSeconds()     - Computes cumulative seconds. No Years, months, weeks, days,
//                                            hours or minutes.
//
//         TDurCalcType(0).GregorianYears()
//                                          - Computes Years based on average length of a Gregorian Year
//                                            Used for very large duration values.
//
//         Type 'TDurCalcType' is located in source file:
//          MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation  string
//       - Designates the standard Time Zone location by which
//         time duration will be compared. This ensures that
//         'oranges are compared to oranges and apples are compared
//         to apples' with respect to start time and end time comparisons.
//
//         If 'timeZoneLocation' is passed as an empty string, it
//         will be automatically defaulted to the 'UTC' time zone.
//         Reference Universal Coordinated Time:
//            https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
//         Time zone location must be designated as one of three types of
//         time zones.
//
//         (1) The time zone "Local", which Golang accepts as
//             the time zone currently configured on the host
//             computer.
//
//         (2) IANA Time Zone - A valid IANA Time Zone from the
//             IANA database.
//             See https://golang.org/pkg/time/#LoadLocation
//             and https://www.iana.org/time-zones to ensure that
//             the IANA Time Zone Database is properly configured
//             on your system.
//
//             IANA Time Zone Examples:
//               "America/New_York"
//               "America/Chicago"
//               "America/Denver"
//               "America/Los_Angeles"
//               "Pacific/Honolulu"
//               "Etc/UTC" = GMT or UTC
//
//         (3) A Military Time Zone
//             Reference:
//               https://en.wikipedia.org/wiki/List_of_military_time_zones
//               http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//               https://www.timeanddate.com/time/zones/military
//               https://www.timeanddate.com/worldclock/timezone/alpha
//               https://www.timeanddate.com/time/map/
//
//              Examples:
//                "Alpha"   or "A"
//                "Bravo"   or "B"
//                "Charlie" or "C"
//                "Delta"   or "D"
//                "Zulu"    or "Z"
//
//                If the time zone "Zulu" is passed to this method, it will be
//                classified as a Military Time Zone.
//
//
//  timeCalcMode      TimeMathCalcMode
//       - TimeMathCalcMode is an enumeration which specifies the algorithm
//         which will be used when computing time spans or time duration.
//
//         If 'LocalTimeZone' is specified, days are defined as local time
//         zone days which may be less than, or greater than, 24-hours due
//         to local conventions like daylight savings time.
//         (TCalcMode.LocalTimeZone())
//
//         If 'UtcTimeZone' is specified, days are uniformly defined as
//         a time span consisting of 24-consecutive hours.
//         (TCalcMode.UtcTimeZone())
//
//         For additional information see the type documentation at
//               datetime\timemathcalcmode.go
//
//
//  dateTimeFmtStr    string
//       - A date time format string which will be used
//         to format and display 'dateTime'. Example:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//         If 'dateTimeFmtStr' is submitted as an
//         'empty string', a default date time format
//         string will be applied. The default date time
//         format string is:
//           FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  ePrefix string
//       - The error prefix containing the names of all
//         the methods executed up to this point.
//
func (tDurDtoUtil *timeDurationDtoUtility) setStartTimePlusTimeDto(
	tDur *TimeDurationDto,
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "timeDurationDtoUtility.setStartTimePlusTimeDto() "

	if tDur == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDur",
			inputParameterValue: "",
			errMsg:              "Input Parameter 'tDur' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if tDur.lock == nil {
		tDur.lock = new(sync.Mutex)
	}

	err := plusTimeDto.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Input Parameter 'plusTimeDto' is INVALID!\n"+
			"Validation Error='%v'\n", err.Error())
	}

	if startDateTime.IsZero() && plusTimeDto.IsEmpty() {
		return errors.New(ePrefix +
			"\nError: Both 'startDateTime' and 'plusTimeDto' " +
			"input parameters are ZERO/EMPTY!\n")
	}

	if tDurCalcType < TDurCalc.XFirstValidCalcType() ||
		tDurCalcType > TDurCalc.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "tDurCalcType",
			inputParameterValue: tDurCalcType.String(),
			errMsg:              "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	if timeCalcMode < TCalcMode.XFirstValidCalcType() ||
		timeCalcMode > TCalcMode.XLastValidCalcType() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "timeCalcMode",
			inputParameterValue: "",
			errMsg: "Input parameter 'timeCalcMode' " +
				"is invalid!",
			err: nil,
		}
	}

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tDur2 := TimeDurationDto{}

	tDur2.lock = new(sync.Mutex)

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTimeTzName(
		&tDur2.startDateTimeTz,
		startDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	tDur2.endDateTimeTz, err = dTzUtil.addPlusTimeDto(
		&tDur2.startDateTimeTz,
		timeCalcMode,
		plusTimeDto,
		ePrefix)

	tDur2.timeDuration =
		tDur2.endDateTimeTz.dateTimeValue.Sub(
			tDur2.startDateTimeTz.dateTimeValue)

	tDur2.timeDurCalcType = tDurCalcType

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
