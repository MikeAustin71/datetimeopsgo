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


// calcTimeDurationAllocations - Examines the input parameter 'calcType' and
// then determines which type of time duration allocation calculation will be
// applied to the data fields of the current TimeDurationDto instance.
func (tDurDtoUtil *timeDurationDtoUtility) calcTimeDurationAllocations(
	tDur *TimeDurationDto,
	calcType TDurCalcType,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "TimeDurationDto.calcTimeDurationAllocations() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	switch calcType {

	case TDurCalcType(0).StdYearMth():
		return tDur.calcTypeSTDYEARMTH()

	case TDurCalcType(0).CumMonths():
		return tDur.calcTypeCUMMONTHS()

	case TDurCalcType(0).CumWeeks():
		return tDur.calcTypeCUMWEEKS()

	case TDurCalcType(0).CumDays():
		return tDurDtoUtilX2.calcTypeCUMDays(tDur, ePrefix)

	case TDurCalcType(0).CumHours():
		return tDur.calcTypeCUMHours()

	case TDurCalcType(0).CumMinutes():
		return tDur.calcTypeCUMMINUTES()

	case TDurCalcType(0).CumSeconds():
		return tDur.calcTypeCUMSECONDS()

	case TDurCalcType(0).GregorianYears():
		return tDur.calcTypeGregorianYears()
	}

	return fmt.Errorf(ePrefix+
		"Error: Invalid TDurCalcType. calcType='%v'", calcType.String())
}

// calcTypeCUMDays - Calculates Cumulative Days. Years, months and weeks are consolidated
// and counted as cumulative days. The Data Fields for years, months, weeks and week days
// are set to zero.  All cumulative days are allocated to the data field, 'DateDays'.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCUMDays(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "TimeDurationDto.calcTypeCUMDays() "

	tDurDtoUtil2 := timeDurationDtoUtility{}

	tDurDtoUtil2.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumDays()

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	if rd >= DayNanoSeconds {
		tDur.DateDays = rd / DayNanoSeconds
		tDur.DateDaysNanosecs = tDur.DateDays * DayNanoSeconds
	}

	err := tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}


// calcTypeCUMHours - Calculates Cumulative Hours. Years, months, weeks, week days,
// date days and hours are consolidated and included in cumulative hours. Values for years,
// months, weeks, week days and date days are ignored and set to zero. Time duration is
// allocated over cumulative hours plus minutes, seconds, milliseconds, microseconds and
// nanoseconds.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCUMHours(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "TimeDurationDto.calcTypeCUMHours() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDurDtoUtilX2.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumHours()

	if tDur.TimeDuration == 0 {
		return nil
	}

	err := tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}

// calcTypeGregorianYears - Allocates Years using the number of nanoseconds in a
// standard or average GregorianYear
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeGregorianYears(
tDur *TimeDurationDto,
ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "TimeDurationDto.calcTypeGregorianYears() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDurDtoUtilX2.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).GregorianYears()

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	if rd >= GregorianYearNanoSeconds {
		tDur.Years = rd / GregorianYearNanoSeconds
		tDur.YearsNanosecs = tDur.Years * GregorianYearNanoSeconds
	}

	err := tDur.calcMonthsFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcDateDaysWeeksFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}


// calcTypeCUMMINUTES - Calculates Cumulative Minutes. Years, months, weeks, week days,
// date days, hours and minutes are consolidated and included in cumulative minutes.
// Values for years, months, weeks, week days, date days and hours are ignored and set
// to zero. Time duration is allocated over cumulative minutes plus seconds, milliseconds,
// microseconds and nanoseconds.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCUMMINUTES(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "TimeDurationDto.calcTypeCUMHours() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDurDtoUtilX2.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumMinutes()

	if tDur.TimeDuration == 0 {
		return nil
	}

	rd := int64(tDur.TimeDuration)

	if rd >= MinuteNanoSeconds {
		tDur.Minutes = rd / MinuteNanoSeconds
		tDur.MinutesNanosecs = tDur.Minutes * MinuteNanoSeconds
		rd -= tDur.MinutesNanosecs
	}

	if rd >= SecondNanoseconds {
		tDur.Seconds = rd / SecondNanoseconds
		tDur.SecondsNanosecs = tDur.Seconds * SecondNanoseconds
		rd -= tDur.SecondsNanosecs
	}

	if rd >= MilliSecondNanoseconds {
		tDur.Milliseconds = rd / MilliSecondNanoseconds
		tDur.MillisecondsNanosecs = tDur.Milliseconds * MilliSecondNanoseconds
		rd -= tDur.MillisecondsNanosecs
	}

	if rd >= MicroSecondNanoseconds {
		tDur.Microseconds = rd / MicroSecondNanoseconds
		tDur.MillisecondsNanosecs = tDur.Microseconds * MicroSecondNanoseconds
		rd -= tDur.MicrosecondsNanosecs
	}

	tDur.Nanoseconds = rd

	err := tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil

}


// Data Fields for Years is always set to Zero. Years
// and months are consolidated and counted as cumulative
// months.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCUMMONTHS(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "TimeDurationDto.calcTypeCUMWEEK() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDurDtoUtilX2.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumMonths()

	err := tDur.calcMonthsFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcDateDaysWeeksFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil

}

// calcTypeCUMSECONDS - Calculates Cumulative Seconds of
// time duration.
//
// tDur.CalcType = TDurCalcType(0).CumSeconds()
//
// Years, months, weeks, weekdays, date days, hours and
// minutes are ignored and set to zero. Time is accumulated
// in seconds, milliseconds, microseconds and nanoseconds.
//
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCUMSECONDS(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "TimeDurationDto.calcTypeCUMSECONDS() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDurDtoUtilX2.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumSeconds()

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return nil
	}

	if rd >= SecondNanoseconds {
		tDur.Seconds = rd / SecondNanoseconds
		tDur.SecondsNanosecs = SecondNanoseconds * tDur.Seconds
		rd -= tDur.SecondsNanosecs
	}

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

// calcTypeSTDYEARMTH - Performs Duration calculations for
// TDurCalcType == TDurCalcType(0).StdYearMth()
//
// TDurCalcTypeYEARMTH - Standard Year, Month, Weeks, Days calculation.
// All data fields in the TimeDto are populated in the duration
// allocation.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeSTDYEARMTH(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "TimeDurationDto.calcTypeSTDYEARMTH() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDurDtoUtilX2.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).StdYearMth()

	err := tDur.calcYearsFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcYearsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcMonthsFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcDateDaysWeeksFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}

// calcTypeCUMWEEKS - Data Fields for Years and Months are always set to zero.
// Years and Months are consolidated and counted as equivalent Weeks.
func (tDurDtoUtil *timeDurationDtoUtility) calcTypeCUMWEEKS(
	tDur *TimeDurationDto,
	ePrefix string) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

	ePrefix += "TimeDurationDto.calcTypeCUMWEEKS() "

	tDurDtoUtilX2 := timeDurationDtoUtility{}

	tDurDtoUtilX2.emptyTimeFields(
		tDur,
		ePrefix)

	tDur.CalcType = TDurCalcType(0).CumWeeks()

	rd := int64(tDur.TimeDuration)

	if rd >= WeekNanoSeconds {

		tDur.Weeks = rd / WeekNanoSeconds
		tDur.WeeksNanosecs = tDur.Weeks * WeekNanoSeconds
		rd -= tDur.WeeksNanosecs
	}

	if rd >= DayNanoSeconds {
		tDur.WeekDays = rd / DayNanoSeconds
		tDur.WeekDaysNanosecs = tDur.WeekDays * DayNanoSeconds
		rd -= tDur.WeekDaysNanosecs
	}

	tDur.DateDays = tDur.Weeks * int64(7)
	tDur.DateDays += tDur.WeekDays
	tDur.DateDaysNanosecs = tDur.WeeksNanosecs + tDur.WeekDaysNanosecs

	err := tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	// For Cumulative Weeks calculation and presentations, set Date Days to zero
	tDur.DateDays = 0
	tDur.DateDaysNanosecs = 0

	return nil
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

	ePrefix += "TimeDurationDto.calcYearsFromDuration() "

	years := int64(0)
	yearNanosecs := int64(0)
	startTime := tDur.StartTimeDateTz.dateTimeValue
	endTime := tDur.EndTimeDateTz.dateTimeValue

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix+"Error: 'startTime' and 'endTime' Time Zone Location do NOT match! "+
			"startTimeZoneLocation='%v'  endTimeZoneLocation='%v'",
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

	ePrefix += "TimeDurationDto.calcMonthsFromDuration() "

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
	tDur *TimeDurationDto) error {

	tDurDtoUtil.lock.Lock()

	defer tDurDtoUtil.lock.Unlock()

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
func (tDurDtoUtil *timeDurationDtoUtility) setStartEndTimesCalcTz(
	tDur *TimeDurationDto,
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr string) error {

		tDurDtoUtil.lock.Lock()

		tDurDtoUtil.lock.Unlock()

		ePrefix := "timeDurationDtoUtility.setStartEndTimesCalcTz() "

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

	if tDurCalcType < TDurCalc.StdYearMth() ||
		tDurCalcType > TDurCalc.GregorianYears() {
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

	var err error
	var startDateTzDto, endDateTzDto DateTzDto

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTimeTzName(
		&startDateTzDto,
		startDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	err = dTzUtil.setFromTimeTzName(
		&endDateTzDto,
		startDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFmtStr,
		ePrefix)

	if err != nil {
		return err
	}

	var timeDuration time.Duration

	timeDuration =
		endDateTzDto.dateTimeValue.Sub(
			startDateTzDto.dateTimeValue)

	tDurDtoUtil2 := timeDurationDtoUtility{}

	tDur2 := TimeDurationDto{}


	tDur2.StartTimeDateTz =
		startDateTzDto.CopyOut()

	tDur2.EndTimeDateTz =
		endDateTzDto.CopyOut()

	tDur2.TimeDuration = timeDuration



	tDurDtoUtil2.empty(tDur, ePrefix)

	return nil
}