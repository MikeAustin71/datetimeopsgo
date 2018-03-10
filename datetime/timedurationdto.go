package datetime

import (
	"time"
	"errors"
	"fmt"
	"strings"
)


// TDurCalcType - Time Duration Calculation Type. 
// Specifies how time duration is allocated by 
// Years, Months, Weeks, Days, Hours, Seconds, 
// Milliseconds, Microseconds and Nanoseconds.
type TDurCalcType int

// String - Returns a string equivalent to the 
// integer value of TDurCalcType
func (tDurCalcType TDurCalcType) String() string {
	
	return TDurCalcTypeLabels[tDurCalcType]
}

const (
	
	// TDurCalcTypeYEARMTH - Standard Year, Month, Weeks, Days calculation. All data 
	// fields in the TimeDto are populated in the duration allocation
	TDurCalcTypeSTDYEARMTH	TDurCalcType = iota

	// TDurCalcTypeCUMMONTHS - Cumulative Months Calculation. Years are ignored.
	// Years and Months are consolidated and counted as cumulative months. Years
	// duration is not provided. The entire duration is broken down by cumulative
	// months plus weeks, week days, date days, hours, minutes, seconds, milliseconds,
	// microseconds and nanoseconds. The Data Fields for for Years is set to zero.
	TDurCalcTypeCUMMONTHS

	// TDurCalcTypeCUMWEEKS - Cumulative Weeks calculation. Years and months are ignored.
	// Years, months, weeks are consolidated and counted as cumulative weeks. Years and
	// months duration is not provided. The entire duration is broken down by cumulative 
	// weeks, plus week days, hours, minutes, seconds, milliseconds, microseconds and
	// nanoseconds. Data Fields for Years and Months are always set to zero.
	TDurCalcTypeCUMWEEKS
	
	// TDurCalcTypeCUMDAYS - Cumulative Days calculation. Years, months and weeks are 
	// ignored. Years, months, weeks and days are consolidated and counted as cumulative
	// days. Years, months and weeks duration is not calculated. The entire duration is 
	// broken down by cumulative days plus hours, minutes, seconds, milliseconds,
	// microseconds and nanoseconds. Data Fields for years, months, weeks, and weekdays
	// are always set to zero.
	TDurCalcTypeCUMDAYS
	
	// TDurCalcTypeCUMHOURS - Cumulative Hours calculations. Years, months, weeks, and days
	// are ignored. Years, months weeks, days and hours are consolidated as cumulative hours.
	// Years, months, weeks and days duration is not calculated. The entire duration is 
	// broken down by cumulative hours plus minutes, seconds, milliseconds, microseconds
	// and nanoseconds.  Data Fields for years, months, weeks, and days are always set
	// to zero. 
	TDurCalcTypeCUMHOURS

	// TDurCalcTypeCUMMINUTES - Cumulative Minutes calculations. Years, months, weeks, days
	// and hours are ignored. Years, months weeks, days, hours and minutes are consolidated
	// and counted as cumulative Minutes.
	//
	// Years, months, weeks, days and hours duration is not calculated. The entire duration is
	// broken down by cumulative minutes plus seconds, milliseconds, microseconds
	// and nanoseconds.  Data Fields for years, months, weeks, days and hours are always set
	// to zero.
	TDurCalcTypeCUMMINUTES

	// TDurCalcTypeGregorianYrs - Allocates Years, Months, Weeks, WeekDays, Date Days, Hours
	// Minutes, Seconds, Milliseconds, Microseconds, Nanoseconds. However, the Years allocation
	// is performed using standard Gregorian Years. 
	//			Sources:
	//					https://en.wikipedia.org/wiki/Year
	//					Source: https://en.wikipedia.org/wiki/Gregorian_calendar
	//
	// The Gregorian Average Year is therefore equivalent to 365 days, 5 hours,
	// 49 minutes and 12 seconds. GregorianYearNanoSeconds = 31,556,952,000,000,000 nanoseconds
	//
	TDurCalcTypeGregorianYrs
	
)

// TDurCalcTypeLabels - Text Names associated with TDurCalcType types.
var TDurCalcTypeLabels = [...]string{"StdYearMthCalc","CumMonthsCalc","CumWeeksCalc", "CumDaysCalc",
																			"CumHoursCalc", "CumMinutesCalc", "GregorianYrsCalc"}

// TimeDurationDto - Is designed to work with incremental time or duration.
type TimeDurationDto struct {
	StartTimeDateTz				DateTzDto			// Starting Date Time with Time Zone info
	EndTimeDateTz        	DateTzDto			// Ending Date Time with Time Zone info
	TimeDuration         	time.Duration	// Elapsed time or duration between starting and ending date time
	CalcType              TDurCalcType  // The calculation Type. This controls the allocation of time 
																			// 		duration over years, months, weeks, days and hours.
	Years                	int64					// Number of Years
	YearsNanosecs        	int64					// Number of Years in Nanoseconds
	Months               	int64					// Number of Months
	MonthsNanosecs       	int64					// Number of Months in Nanoseconds
	Weeks                	int64					// Number of Weeks: Date Days / 7
	WeeksNanosecs        	int64					// Number of Weeks in Nanoseconds
	WeekDays             	int64					// WeekDays = DateDays - (Weeks * 7)
	WeekDaysNanosecs     	int64					// Equivalent WeekDays in NanoSeconds
	DateDays             	int64					// Day Number in Month (1-31)
	DateDaysNanosecs     	int64					// DateDays in equivalent nanoseconds
	Hours                	int64					// Number of Hours 
	HoursNanosecs        	int64					// Number of Hours in Nanoseconds
	Minutes              	int64					// Number of Minutes
	MinutesNanosecs      	int64					// Number of Minutes in Nanoseconds
	Seconds              	int64					// Number of Seconds
	SecondsNanosecs      	int64					// Number of Seconds in Nanoseconds
	Milliseconds         	int64					// Number of Milliseconds
	MillisecondsNanosecs 	int64					// Number of Milliseconds in Nanoseconds
	Microseconds         	int64					// Number of Microseconds
	MicrosecondsNanosecs 	int64					// Number of Microseconds in Nanoseconds
	Nanoseconds          	int64					// Number of Nanoseconds (Remainder after Milliseconds & Microseconds) 
	TotSubSecNanoseconds 	int64					// Equivalent Nanoseconds for Milliseconds + Microseconds + Nanoseconds
	TotDateNanoseconds		int64					// Equal to Years + Months + DateDays in equivalent nanoseconds.
	TotTimeNanoseconds		int64					// Equal to Hours + Seconds + Milliseconds + Microseconds + Nanoseconds in
																			// 		in equivalent nanoseconds

}

// CopyIn - Receives a TimeDurationDto as an input parameters
// and proceeds to set all data fields of the current TimeDurationDto
// equal to the incoming TimeDurationDto.
//
// When this method completes, the current TimeDurationDto will
// equal in all respects to the incoming TimeDurationDto.
func (tDur *TimeDurationDto) CopyIn(t2Dur TimeDurationDto) {
	
	tDur.Empty()

	tDur.StartTimeDateTz 				= t2Dur.StartTimeDateTz.CopyOut()
	tDur.EndTimeDateTz 					=	t2Dur.EndTimeDateTz.CopyOut()
	tDur.TimeDuration     			= t2Dur.TimeDuration
	tDur.CalcType								= t2Dur.CalcType
	tDur.Years									= t2Dur.Years
	tDur.YearsNanosecs    			= t2Dur.YearsNanosecs
	tDur.Months           			= t2Dur.Months
	tDur.MonthsNanosecs   			= t2Dur.MonthsNanosecs
	tDur.Weeks            			= t2Dur.Weeks
	tDur.WeeksNanosecs    			= t2Dur.WeeksNanosecs
	tDur.WeekDays								= t2Dur.WeekDays
	tDur.WeekDaysNanosecs				= t2Dur.WeekDaysNanosecs
	tDur.DateDays								= t2Dur.DateDays
	tDur.DateDaysNanosecs				= t2Dur.DateDaysNanosecs
	tDur.Hours									= t2Dur.Hours
	tDur.HoursNanosecs					= t2Dur.HoursNanosecs
	tDur.Minutes								= t2Dur.Minutes
	tDur.MinutesNanosecs				= t2Dur.MinutesNanosecs
	tDur.Seconds								= t2Dur.Seconds
	tDur.SecondsNanosecs				= t2Dur.SecondsNanosecs
	tDur.Milliseconds						= t2Dur.Milliseconds
	tDur.MillisecondsNanosecs		= t2Dur.MillisecondsNanosecs
	tDur.Microseconds						= t2Dur.Microseconds
	tDur.MicrosecondsNanosecs 	= t2Dur.MicrosecondsNanosecs
	tDur.Nanoseconds						= t2Dur.MillisecondsNanosecs
	tDur.TotSubSecNanoseconds 	= t2Dur.TotSubSecNanoseconds
	tDur.TotDateNanoseconds			= t2Dur.TotDateNanoseconds
	tDur.TotTimeNanoseconds			= t2Dur.TotTimeNanoseconds
}

// CopyOut - Returns a deep copy of the current 
// TimeDurationDto instance.
func (tDur *TimeDurationDto) CopyOut() TimeDurationDto {

	t2Dur := TimeDurationDto{}
	
	t2Dur.StartTimeDateTz 			= tDur.StartTimeDateTz.CopyOut()
	t2Dur.EndTimeDateTz 				=	tDur.EndTimeDateTz.CopyOut()
	t2Dur.TimeDuration     			= tDur.TimeDuration
	t2Dur.CalcType							= tDur.CalcType
	t2Dur.Years									= tDur.Years
	t2Dur.YearsNanosecs    			= tDur.YearsNanosecs
	t2Dur.Months           			= tDur.Months
	t2Dur.MonthsNanosecs   			= tDur.MonthsNanosecs
	t2Dur.Weeks            			= tDur.Weeks
	t2Dur.WeeksNanosecs    			= tDur.WeeksNanosecs
	t2Dur.WeekDays							= tDur.WeekDays
	t2Dur.WeekDaysNanosecs			= tDur.WeekDaysNanosecs
	t2Dur.DateDays							= tDur.DateDays
	t2Dur.DateDaysNanosecs			= tDur.DateDaysNanosecs
	t2Dur.Hours									= tDur.Hours
	t2Dur.HoursNanosecs					= tDur.HoursNanosecs
	t2Dur.Minutes								= tDur.Minutes
	t2Dur.MinutesNanosecs				= tDur.MinutesNanosecs
	t2Dur.Seconds								= tDur.Seconds
	t2Dur.SecondsNanosecs				= tDur.SecondsNanosecs
	t2Dur.Milliseconds					= tDur.Milliseconds
	t2Dur.MillisecondsNanosecs	= tDur.MillisecondsNanosecs
	t2Dur.Microseconds					= tDur.Microseconds
	t2Dur.MicrosecondsNanosecs 	= tDur.MicrosecondsNanosecs
	t2Dur.Nanoseconds						= tDur.MillisecondsNanosecs
	t2Dur.TotSubSecNanoseconds 	= tDur.TotSubSecNanoseconds
	t2Dur.TotDateNanoseconds		= tDur.TotDateNanoseconds
	t2Dur.TotTimeNanoseconds		= tDur.TotTimeNanoseconds
	
	return t2Dur
}

// Empty - Resets all of the current TimeDurationDto
// data fields to their zero or uninitialized values.
func (tDur *TimeDurationDto) Empty() {
	tDur.StartTimeDateTz 			= DateTzDto{}
	tDur.EndTimeDateTz 				=	DateTzDto{}
	tDur.TimeDuration     		= time.Duration(0)
	tDur.CalcType							= TDurCalcTypeSTDYEARMTH
	tDur.Years								= 0
	tDur.YearsNanosecs    		= 0
	tDur.Months           		= 0
	tDur.MonthsNanosecs   		= 0
	tDur.Weeks            		= 0
	tDur.WeeksNanosecs    		= 0
	tDur.WeekDays							= 0
	tDur.WeekDaysNanosecs			= 0
	tDur.DateDays							= 0
	tDur.DateDaysNanosecs			= 0
	tDur.Hours								= 0
	tDur.HoursNanosecs				= 0
	tDur.Minutes							= 0
	tDur.MinutesNanosecs			= 0
	tDur.Seconds							= 0
	tDur.SecondsNanosecs			= 0
	tDur.Milliseconds					= 0
	tDur.MillisecondsNanosecs	= 0
	tDur.Microseconds					= 0
	tDur.MicrosecondsNanosecs = 0
	tDur.Nanoseconds					= 0
	tDur.TotSubSecNanoseconds = 0
	tDur.TotDateNanoseconds		= 0
	tDur.TotTimeNanoseconds		= 0
}

// EmptyTimeFields - Sets all of the data fields
// associated with time duration allocation to zero.
func (tDur *TimeDurationDto) EmptyTimeFields() {

	tDur.Years								= 0
	tDur.YearsNanosecs    		= 0
	tDur.Months           		= 0
	tDur.MonthsNanosecs   		= 0
	tDur.Weeks            		= 0
	tDur.WeeksNanosecs    		= 0
	tDur.WeekDays							= 0
	tDur.WeekDaysNanosecs			= 0
	tDur.DateDays							= 0
	tDur.DateDaysNanosecs			= 0
	tDur.Hours								= 0
	tDur.HoursNanosecs				= 0
	tDur.Minutes							= 0
	tDur.MinutesNanosecs			= 0
	tDur.Seconds							= 0
	tDur.SecondsNanosecs			= 0
	tDur.Milliseconds					= 0
	tDur.MillisecondsNanosecs	= 0
	tDur.Microseconds					= 0
	tDur.MicrosecondsNanosecs = 0
	tDur.Nanoseconds					= 0
	tDur.TotSubSecNanoseconds = 0
	tDur.TotDateNanoseconds		= 0
	tDur.TotTimeNanoseconds		= 0

}

// Equal - Compares two TimeDurationDto instances to determine
// if they are equivalent.
func (tDur *TimeDurationDto) Equal(t2Dur TimeDurationDto) bool {
	
	if	!tDur.StartTimeDateTz.Equal(t2Dur.StartTimeDateTz)				||
		 	!tDur.EndTimeDateTz.Equal(t2Dur.EndTimeDateTz)						||
			tDur.TimeDuration 				!= 	t2Dur.TimeDuration					||
			tDur.CalcType							!=	t2Dur.CalcType							||
			tDur.Years								!= 	t2Dur.Years									||
			tDur.YearsNanosecs    		!= 	t2Dur.YearsNanosecs					||
			tDur.Months           		!= 	t2Dur.Months 								||
			tDur.MonthsNanosecs   		!= 	t2Dur.MonthsNanosecs				||
			tDur.Weeks            		!= 	t2Dur.Weeks									||
			tDur.WeeksNanosecs    		!= 	t2Dur.WeeksNanosecs					||
			tDur.WeekDays							!= 	t2Dur.WeekDays							||
			tDur.WeekDaysNanosecs			!= 	t2Dur.WeekDaysNanosecs			||
			tDur.DateDays							!= 	t2Dur.DateDays							||
			tDur.DateDaysNanosecs 		!= 	t2Dur.DateDaysNanosecs			||
			tDur.Hours								!= 	t2Dur.Hours									||
			tDur.HoursNanosecs				!= 	t2Dur.HoursNanosecs					||
			tDur.Minutes							!=	t2Dur.Minutes								||
			tDur.MinutesNanosecs			!= 	t2Dur.MinutesNanosecs				||
			tDur.Seconds							!= 	t2Dur.Seconds								||
			tDur.SecondsNanosecs			!= 	t2Dur.SecondsNanosecs				||
			tDur.Milliseconds					!= 	t2Dur.Milliseconds					||
			tDur.MillisecondsNanosecs	!=	t2Dur.MillisecondsNanosecs	||
			tDur.Microseconds					!= 	t2Dur.Microseconds					||
			tDur.MicrosecondsNanosecs != 	t2Dur.MicrosecondsNanosecs	||
			tDur.Nanoseconds					!= t2Dur.MillisecondsNanosecs		||
			tDur.TotSubSecNanoseconds != t2Dur.TotSubSecNanoseconds		||
			tDur.TotDateNanoseconds		!= t2Dur.TotDateNanoseconds			||
			tDur.TotTimeNanoseconds		!= t2Dur.TotTimeNanoseconds			{
				
				return false
	}
	
	return true
	
}

// IsEmpty() - Returns 'true' if the current TimeDurationDto
// instance is uninitialized and consists entirely of zero values.
func (tDur *TimeDurationDto) IsEmpty() bool {


	if	tDur.StartTimeDateTz.IsEmpty()				&&
		  tDur.EndTimeDateTz.IsEmpty()					&&
			tDur.TimeDuration 						== 0 		&&
			tDur.Years										== 0 		&&
			tDur.YearsNanosecs   					== 0 		&&
			tDur.Months          					== 0 	 	&&
			tDur.MonthsNanosecs 			 		== 0 		&&
			tDur.Weeks 			          		== 0 		&&
			tDur.WeeksNanosecs		    		== 0 		&&
			tDur.WeekDays									== 0 		&&
			tDur.WeekDaysNanosecs					== 0 		&&
			tDur.DateDays									== 0		&&
			tDur.DateDaysNanosecs 				== 0		&&
			tDur.Hours										== 0		&&
			tDur.HoursNanosecs						== 0		&&
			tDur.Minutes									== 0		&&
			tDur.MinutesNanosecs					== 0		&&
			tDur.Seconds									== 0		&&
			tDur.SecondsNanosecs					== 0		&&
			tDur.Milliseconds							== 0		&&
			tDur.MillisecondsNanosecs			== 0		&&
			tDur.Microseconds							== 0		&&
			tDur.MicrosecondsNanosecs 		== 0		&&
			tDur.Nanoseconds							== 0		&&
			tDur.TotSubSecNanoseconds 		== 0		&&
			tDur.TotDateNanoseconds				== 0		&&
			tDur.TotTimeNanoseconds				== 0			{

		tDur.CalcType = TDurCalcTypeSTDYEARMTH	
		return true
	}

	return false
	
}

// GetYearMthDaysTimeAbbrvStr - Abbreviated formatting of Years, Months,
// DateDays, Hours, Minutes, Seconds, Milliseconds, Microseconds and
// Nanoseconds. At a minimum only Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
//
// Abbreviated Years Mths DateDays Time Duration - Example Return:
//
// 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearMthDaysTimeAbbrvStr() string {

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	str := ""

	if tDur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", tDur.Years)
	}

	if tDur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", tDur.Months)
	}

	if tDur.DateDays > 0 || str != "" {
		str +=  fmt.Sprintf("%v-Days ", tDur.DateDays)
	}

	str += fmt.Sprintf("%v-Hours ", tDur.Hours)

	str += fmt.Sprintf("%v-Minutes ", tDur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", tDur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", tDur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", tDur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", tDur.Nanoseconds)

	return str

}

// GetYearMthDaysTimeStr - Calculates Duration and breakdowns
// time elements by Years, Months, Date Days, hours, minutes,
// seconds, milliseconds, microseconds and nanoseconds.
//
// Example DisplayStr
// ==================
//
// Years Months DateDays Time Duration - Example Return:
//
// 12-Years 3-Months 2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
// If Years, Months and Days have a zero value, only the time components will be displayed. 
// Example:
//		13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (tDur *TimeDurationDto) GetYearMthDaysTimeStr() string {

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	str := ""
	
	if tDur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", tDur.Years)	
	}
	
	if tDur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", tDur.Months)	
	}

	if tDur.DateDays > 0 || str!= "" {
		str += fmt.Sprintf("%v-Days ", tDur.DateDays)
	}


	str += fmt.Sprintf("%v-Hours ", tDur.Hours)

	str += fmt.Sprintf("%v-Minutes ", tDur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", tDur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", tDur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", tDur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", tDur.Nanoseconds)

	return str
}

// GetYearsMthsWeeksTimeAbbrvStr - Abbreviated formatting of Years, Months,
// Weeks, WeekDays, Hours, Minutes, Seconds, Milliseconds, Microseconds,
// Nanoseconds. 
// 
// At a minimum only Hours, Minutes, Seconds, Milliseconds, Microseconds
// Nanoseconds are displayed. Example return when Years, Months, Weeks
// and WeekDays are zero:
//
// 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearsMthsWeeksTimeAbbrvStr() string {

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	str := ""
	
	if tDur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", tDur.Years)
	}

	if tDur.Months > 0 || str != "" {
		str += fmt.Sprintf("%v-Months ", tDur.Months)
	}

	if tDur.Weeks > 0 || str != "" {
		str += fmt.Sprintf("%v-Weeks ", tDur.Weeks)
	}

	if tDur.WeekDays > 0 || str != "" {
		str += fmt.Sprintf("%v-WeekDays ", tDur.WeekDays)
	}

	str += fmt.Sprintf("%v-Hours ", tDur.Hours)

	str += fmt.Sprintf("%v-Minutes ", tDur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", tDur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", tDur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", tDur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", tDur.Nanoseconds)

	return str
}

// GetYearsMthsWeeksTimeStr - Example Return:
// 12-Years 3-Months 2-Weeks 1-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
// At a minimum only Weeks, WeekDays, Hours, Minutes, Seconds,
// Milliseconds, Microseconds and Nanoseconds are displayed.
// 
// Example return when Years, and Months are zero:
//
// 3-Weeks 2-WeekDays 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetYearsMthsWeeksTimeStr() string {

	if int64(tDur.TimeDuration) == 0 {
			return "0-Nanoseconds"
	}	

	str := ""
	
	if tDur.Years > 0 {
		str += fmt.Sprintf("%v-Years ", tDur.Years)	
	}
	
	if tDur.Months > 0 || str != "" {
		str+= fmt.Sprintf("%v-Months ", tDur.Months)
	}

	str+= fmt.Sprintf("%v-Weeks ", tDur.Weeks)	
	
	str+= fmt.Sprintf("%v-WeekDays ", tDur.WeekDays)

	str+= fmt.Sprintf("%v-Hours ", tDur.Hours)

	str+= fmt.Sprintf("%v-Minutes ", tDur.Minutes)

	str+= fmt.Sprintf("%v-Seconds ", tDur.Seconds)

	str+= fmt.Sprintf("%v-Milliseconds ", tDur.Milliseconds)

	str+= fmt.Sprintf("%v-Microseconds ", tDur.Microseconds)

	str+= fmt.Sprintf("%v-Nanoseconds", tDur.Nanoseconds)

	return str
}

// GetWeeksDaysTimeStr - Example DisplayStr
// 126-Weeks 1-WeekDays 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (tDur *TimeDurationDto) GetWeeksDaysTimeStr() string {

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}
	
	t2Dur := tDur.CopyOut()
	
	t2Dur.ReCalcTimeDurationAllocation(TDurCalcTypeCUMWEEKS)
	
	str := ""

	str += fmt.Sprintf("%v-Weeks ", t2Dur.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetDaysTimeStr - Returns duration formatted as
// days, hours, minutes, seconds, milliseconds, microseconds,
// and nanoseconds. Years, months and weeks are always excluded and 
// included in cumulative 'days'.
//
// Example: 
// 
// 97-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
//
func (tDur *TimeDurationDto) GetDaysTimeStr() string {
	
	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	t2Dur := tDur.CopyOut()
	
	t2Dur.ReCalcTimeDurationAllocation(TDurCalcTypeCUMDAYS)
	
	str := ""

	str += fmt.Sprintf("%v-Days ", t2Dur.DateDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)


	return str
}

// GetHoursTimeStr - Returns duration formatted as hours,
// minutes, seconds, milliseconds, microseconds, nanoseconds.
// Example: 152-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (tDur *TimeDurationDto) GetHoursTimeStr() string {

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	t2Dur := tDur.CopyOut()

	t2Dur.ReCalcTimeDurationAllocation(TDurCalcTypeCUMHOURS)

	str := ""

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}

// GetYrMthWkDayHrMinSecNanosecsStr - Returns duration formatted
// as Year, Month, Day, Hour, Second and Nanoseconds.
// Example: 3-Years 2-Months 3-Weeks 2-WeekDays 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds
func (tDur *TimeDurationDto) GetYrMthWkDayHrMinSecNanosecsStr() string {

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}
	
	str := ""

	str += fmt.Sprintf("%v-Years ", tDur.Years)

	str += fmt.Sprintf("%v-Months ", tDur.Months)

	str += fmt.Sprintf("%v-Weeks ", tDur.Weeks)

	str += fmt.Sprintf("%v-WeekDays ", tDur.WeekDays)

	str += fmt.Sprintf("%v-Hours ", tDur.Hours)

	str += fmt.Sprintf("%v-Minutes ", tDur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", tDur.Seconds)

	str += fmt.Sprintf("%v-Nanoseconds", tDur.TotSubSecNanoseconds)

	return str
}

// GetCumMinutesStr - Returns duration formatted as cumulative
// minutes.
//
// Example:
//	"527-Minutes 37-Seconds 18-Milliseconds 256-Microseconds 852-Nanoseconds"
//
func (tDur *TimeDurationDto) GetCumMinutesStr() string {

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	t2Dur := tDur.CopyOut()

	t2Dur.ReCalcTimeDurationAllocation(TDurCalcTypeCUMMINUTES)

	str := ""

	str += fmt.Sprintf("%v-Minutes ", tDur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", tDur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str

}

// GetNanosecondsDurationStr - Returns duration formatted as
// Nanoseconds. DisplayStr shows Nanoseconds expressed as a
// 64-bit integer value.
func (tDur *TimeDurationDto) GetNanosecondsDurationStr() string {
	
	str := fmt.Sprintf("%v-Nanoseconds", int64(tDur.TimeDuration))

	return str

}

// GetDefaultDurationStr - Returns duration formatted
// as nanoseconds. The DisplayStr shows the default
// string value for duration.
// Example: 61h26m46.864197832s
func (tDur *TimeDurationDto) GetDefaultDurationStr() string {

	return fmt.Sprintf("%v", tDur.TimeDuration)
}

// GetGregorianYearDuration - Returns a string showing the
// breakdown of duration by Gregorian Years, WeekDays, Hours, Minutes,
// Seconds, Milliseconds, Microseconds and Nanoseconds. Unlike
// other calculations which use a Standard 365-day year consisting
// of 24-hour days, a Gregorian Year consists of 365 days, 5-hours,
// 59-minutes and 12 Seconds. For the Gregorian calendar the
// average length of the calendar year (the mean year) across
// the complete leap cycle of 400 Years is 365.2425 days.
// Sources:
// https://en.wikipedia.org/wiki/Year
// Source: https://en.wikipedia.org/wiki/Gregorian_calendar
//
func (tDur *TimeDurationDto) GetGregorianYearDuration() string {

	if int64(tDur.TimeDuration) == 0 {
		return "0-Nanoseconds"
	}

	t2Dur := tDur.CopyOut()

	t2Dur.ReCalcTimeDurationAllocation(TDurCalcTypeGregorianYrs)
	
	
	str := fmt.Sprintf("%v-Gregorian Years ", t2Dur.Years)

	str += fmt.Sprintf("%v-WeekDays ", t2Dur.WeekDays)

	str += fmt.Sprintf("%v-Hours ", t2Dur.Hours)

	str += fmt.Sprintf("%v-Minutes ", t2Dur.Minutes)

	str += fmt.Sprintf("%v-Seconds ", t2Dur.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", t2Dur.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", t2Dur.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", t2Dur.Nanoseconds)

	return str
}


// New - Creates and returns a new TimeDurationDto based on starting
// and ending date times.  Because, time zone location is crucial to
// completely accurate duration calculations, the time zone of the
// starting date time, 'startDateTime' is applied to parameter,
// 'endDateTime' before making the duration calculation.
//
// Note: 	This method applies the standard Time Duration allocation, 'TDurCalcTypeSTDYEARMTH'. 
// 				This means that duration is allocated over years, months, weeks, weekdays, date days,
//				hours, minutes, seconds, milliseconds, microseconds and nanoseconds. 
// 				See Type 'TDurCalcType' for details.
//
//	Input Parameters:
//  =================
//
// startDateTime	time.Time	- Starting date time
//
// endDateTime		time.Time - Ending date time
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.New(startTime, endTime, FmtDateTimeYrMDayFmtStr)
//
//		Note: FmtDateTimeYrMDayFmtStr' is a constant available in datetimeconstants.go
//
func (tDur TimeDurationDto) New(startDateTime, endDateTime time.Time,
							dateTimeFmtStr string) (TimeDurationDto, error) {

  ePrefix := "TimeDurationDto.New() "

  if startDateTime.IsZero() && endDateTime.IsZero() {
  	return TimeDurationDto{},
  	errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
  		"input parameters are ZERO!")
	}

	tzStartLocation := startDateTime.Location().String()

	t2Dur := TimeDurationDto{}

	err := t2Dur.SetStartEndTimesTzCalc(startDateTime, endDateTime, TDurCalcTypeSTDYEARMTH, tzStartLocation, dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartEndTimesTzCalc(startDateTime, " +
			"endDateTime, tzStartLocation, dateTimeFmtStr). "+
			"tzStartLocation='%v'  Error='%v'",
				tzStartLocation, err.Error())
	}

	return t2Dur, nil
}

// NewStartEndTimesTz - Creates and returns a new TimeDurationDto populated with 
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
// The user is required to specify a common Time Zone Location four use in converting
// date times to a common frame of reference to subsequent time duration calculations.
//
// Note: 	This method applies the standard Time Duration allocation, calculation type
// 				'TDurCalcTypeSTDYEARMTH'. This means that duration is allocated over years,
// 				months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
// 				microseconds and nanoseconds.	See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// endDateTime		time.Time - Ending time
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesTz(startTime, endTime, TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
// 							datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartEndTimesTz(startDateTime, endDateTime time.Time,
																timeZoneLocation, dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartEndTimesTz() "

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!")
	}

	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
			"Error: 'timeZoneLocation' input parameter is INVALID! " +
			"'timeZoneLocation'='%v'  processed tzLoc= '%v' Error='%v'",
			timeZoneLocation, tzLoc, err.Error())
	}

	t2Dur := TimeDurationDto{}
	
	err = t2Dur.SetStartEndTimesTzCalc(startDateTime, endDateTime, TDurCalcTypeSTDYEARMTH, tzLoc, dateTimeFmtStr)
	
	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix + "Error returned from " + 
			"SetStartEndTimesTzCalc(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)." +
			"Error='%v'", err.Error())
	}
															
	return t2Dur, nil
}

// NewStartEndTimesTzCalc - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
//
// The Time Zone Location used to standardize the time duration calculation is derived
// from input parameter 'startDateTime'.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// endDateTime		time.Time - Ending time
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month week,
// 																		day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesTzCalc(startTime, endTime,
// 																				TDurCalcTypeSTDYEARMTH, FmtDateTimeYrMDayFmtStr)
//
//		Note:	'TDurCalcTypeSTDYEARMTH' is of type 'TDurCalcType' and signals
//						standard year month day time duration allocation.
//
// 						'FmtDateTimeYrMDayFmtStr' is a constant defined in datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartEndTimesCalc(startDateTime,
																endDateTime time.Time,  tDurCalcType TDurCalcType,
																		dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartEndTimesCalc() "

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!")
	}

	timeZoneLocation := startDateTime.Location().String()

	t2Dur := TimeDurationDto{}

	err := t2Dur.SetStartEndTimesTzCalc(startDateTime, endDateTime, tDurCalcType, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix + "Error returned from " +
			"SetStartEndTimesTzCalc(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)." +
			"Error='%v'", err.Error())
	}

	return t2Dur, nil

}

// NewStartEndTimesTzCalc - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters.
//
// The user is required to specify a common Time Zone Location for use in converting
// date times to a common frame of reference for use in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// endDateTime		time.Time - Ending time
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month week,
// 																		day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesTzCalc(startTime, endTime, TzIanaUsCentral,
// 									TDurCalcTypeSTDYEARMTH, FmtDateTimeYrMDayFmtStr)
//
//		Note:	'TDurCalcTypeSTDYEARMTH' is of type 'TDurCalcType' and signals
//						standard year month day time duration allocation.
//
// 						'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants available in
// 							datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartEndTimesTzCalc(startDateTime,
	endDateTime time.Time, timeZoneLocation string, tDurCalcType TDurCalcType,
		dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartEndTimesTzCalc() "

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!")
	}

	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"Error: 'timeZoneLocation' input parameter is INVALID! " +
				"'timeZoneLocation'='%v'  processed tzLoc= '%v' Error='%v'",
				timeZoneLocation, tzLoc, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartEndTimesTzCalc(startDateTime, endDateTime, tDurCalcType, tzLoc, dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix + "Error returned from " +
			"SetStartEndTimesTzCalc(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)." +
			"Error='%v'", err.Error())
	}

	return t2Dur, nil

}

// NewStartEndTimesDateDto - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters. The
// 'startDateTime' and 'endDateTime' parameters are of type DateTzDto.
//
// Time Zone Location is derived from input parameter 'startDateTime' and provides a
// common frame of reference for use in subsequent time duration calculations.
//
// Note: 	This method applies the standard Time Duration allocation, calculation type
// 				'TDurCalcTypeSTDYEARMTH'. This means that duration is allocated over years,
// 				months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
// 				microseconds and nanoseconds.	See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	DateTzDto	- Starting date time
//
// endDateTime		DateTzDto - Ending date time
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesDateDto(startTime, endTime, FmtDateTimeYrMDayFmtStr)
//
// NOTE:		FmtDateTimeYrMDayFmtStr' is a constant defined in source file, datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartEndTimesDateDto(startDateTime,
										endDateTime DateTzDto,
											dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartEndTimesDateDto() "

	if startDateTime.DateTime.IsZero() && endDateTime.DateTime.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!")
	}

	timeZoneLocation := startDateTime.TimeZone.LocationName

	t2Dur := TimeDurationDto{}

	err := t2Dur.SetStartEndTimesDateDtoTzCalc(startDateTime,
															endDateTime,
																TDurCalcTypeSTDYEARMTH,
																	timeZoneLocation,
																		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix + "Error returned from " +
			"SetStartEndTimesDateDtoTzCalc(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)." +
			"Error='%v'", err.Error())
	}

	return t2Dur, nil
}

// NewStartEndTimesDateDtoCalc - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters. The
// 'startDateTime' and 'endDateTime' parameters are of type DateTzDto.
//
// Time Zone Location is derived from input parameter 'startDateTime'. If the 'endDateTime'
// time zone is NOT equivalent to 'startDateTime', the 'endDateTime' will be converted to
// the time zone provided by the 'startDateTime' parameter. This will provide a common basis
// for use in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	DateTzDto	- Starting date time
//
// endDateTime		DateTzDto - Ending date time
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month week,
// 																		day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesDateDtoCalc(startTime, endTime,
// 													TDurCalcTypeSTDYEARMTH, FmtDateTimeYrMDayFmtStr)
//
//		Note:	'TDurCalcTypeSTDYEARMTH' is of type 'TDurCalcType' and signals
//						standard year month day time duration allocation.
//
// 					'FmtDateTimeYrMDayFmtStr' is a constant defined in source file
// 						datetimeconstants.go.
//
func (tDur TimeDurationDto) NewStartEndTimesDateDtoCalc(startDateTime,
										endDateTime DateTzDto, tDurCalcType TDurCalcType,
													dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartEndTimesDateDtoCalc() "

	if startDateTime.DateTime.IsZero() && endDateTime.DateTime.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!")
	}

	t2Dur := TimeDurationDto{}

	timeZoneLocation := startDateTime.TimeZone.LocationName

	err := t2Dur.SetStartEndTimesDateDtoTzCalc(startDateTime,
														endDateTime,
															tDurCalcType,
																timeZoneLocation,
																		dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix + "Error returned from " +
			"SetStartEndTimesDateDtoTzCalc(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)." +
			"Error='%v'", err.Error())
	}

	return t2Dur, nil
}

// NewStartEndTimesDateDtoTzCalc - Creates and returns a new TimeDurationDto populated with
// time duration data based on 'startDateTime' and 'endDateTime' input parameters. The
// 'startDateTime' and 'endDateTime' parameters are of type DateTzDto.
//
// The user is required to specify a specific Time Zone Location used to convert date times
// to a common frame of reference for use in subsequent time duration calculations.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	DateTzDto	- Starting date time
//
// endDateTime		DateTzDto - Ending date time
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month week,
// 																		day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartEndTimesDateTzDtoCalc(startTime, endTime, TzIanaUsCentral,
// 									TDurCalcTypeSTDYEARMTH, FmtDateTimeYrMDayFmtStr)
//
//		Note:	'TDurCalcTypeSTDYEARMTH' is of type 'TDurCalcType' and signals
//						standard year month day time duration allocation.
//
// 					'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants defined in
// 						datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartEndTimesDateDtoTzCalc(startDateTime,
										endDateTime DateTzDto, timeZoneLocation string, tDurCalcType TDurCalcType,
													dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartEndTimesDateDtoTzCalc() "

	if startDateTime.DateTime.IsZero() && endDateTime.DateTime.IsZero() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
				"input parameters are ZERO!")
	}

	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"Error: 'timeZoneLocation' input parameter is INVALID! " +
				"'timeZoneLocation'='%v'  processed tzLoc= '%v' Error='%v'",
				timeZoneLocation, tzLoc, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartEndTimesDateDtoTzCalc(startDateTime,
																	endDateTime,
																		tDurCalcType,
																			tzLoc,
																				dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{}, fmt.Errorf(ePrefix + "Error returned from " +
			"SetStartEndTimesDateDtoTzCalc(startDateTime, endDateTime, timeZoneLocation, dateTimeFmtStr)." +
			"Error='%v'", err.Error())
	}

	return t2Dur, nil
}

// NewStartTimeDuration - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime' and time duration. 'startDateTime' is used to derive Time Zone Location.
// The time duration value is added to 'startDateTime' in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time. The
// actual starting date time is computed by subtracting duration from ending date time.
//
// Note: 	This method applies the standard Time Duration allocation, 'TDurCalcTypeSTDYEARMTH'.
// 				This means that duration is allocated over years, months, weeks, weekdays, date days,
//				hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
// 				See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDuration(startTime, duration, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'FmtDateTimeYrMDayFmtStr' is a constant available in source file, datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartTimeDuration(startDateTime time.Time,
															duration time.Duration,
																dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartTimeDuration() "

	if startDateTime.IsZero() && duration==0 {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
				"input parameters are ZERO!")
	}

	timeZoneLocation := startDateTime.Location().String()

	t2Dur := TimeDurationDto{}

	err := t2Dur.SetStartTimeDurationTzCalc(startDateTime, duration, TDurCalcTypeSTDYEARMTH , timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartTimeDurationTzCalc(...) Error='%v'", err.Error())
	}

	return t2Dur, nil

}

// NewStartTimeDurationTz - Creates and returns a new TimeDurationDto based on input parameters
// 'startDateTime', time duration and 'timeZoneLocation'. 'startDateTime' is converted to the
// specified 'timeZoneLocation' and the duration value is added to it in order to compute the
// ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and the
// 	actual starting date time is computed by subtracting duration.
//
// Note: 	This method applies the standard Time Duration allocation, 'TDurCalcTypeSTDYEARMTH'. 
// 				This means that duration is allocated over years, months, weeks, weekdays, date days,
//				hours, minutes, seconds, milliseconds, microseconds and nanoseconds. 
// 				See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationTz(startTime, duration,
// 										TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants defined in
// 							datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartTimeDurationTz(startDateTime time.Time,
	duration time.Duration, timeZoneLocation, dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartTimeDurationTz() "

	if startDateTime.IsZero() && duration==0 {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
				"input parameters are ZERO!")
	}

	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"Error: 'timeZoneLocation' input parameter is INVALID! " +
				"'timeZoneLocation'='%v'  processed tzLoc= '%v' Error='%v'",
				timeZoneLocation, tzLoc, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartTimeDurationTzCalc(startDateTime, duration, TDurCalcTypeSTDYEARMTH , tzLoc, dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartTimeDurationTzCalc(...) Error='%v'", err.Error())
	}

	return t2Dur, nil

}

// NewStartTimeDurationTzCalc - Creates and returns a new TimeDurationDto based on input
// parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation type.
// 'startDateTime' is converted to the specified 'timeZoneLocation' and the duration value
// is added to it in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and
// the	actual starting date time is computed by subtracting duration.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//			TDurCalcTypeSTDYEARMTH 		- Default - standard year, month week,
// 																		day time calculation.
//
//			TDurCalcTypeCUMMONTHS 		- Computes cumulative months - no Years.
//
//			TDurCalcTypeCUMWEEKS  		- Computes cumulative weeks. No Years or months
//
//			TDurCalcTypeCUMDAYS				- Computes cumulative days. No Years, months or weeks.
//
//			TDurCalcTypeCUMHOURS			- Computes cumulative hours. No Years, months, weeks or days.
//
//			TDurCalcTypeGregorianYrs 	- Computes Years based on average length of a Gregorian Year
//																 		Used for very large duration values.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationTzCalc(startTime, duration,
// 										TzIanaUsCentral, TDurCalcTypeSTDYEARMTH, FmtDateTimeYrMDayFmtStr)
//
//		Note:	'TDurCalcTypeSTDYEARMTH' is of type 'TDurCalcType' and signals
//						standard year month day time duration allocation.
//
// 					'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants defined in
// 						datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartTimeDurationTzCalc(startDateTime time.Time,
					duration time.Duration, timeZoneLocation string, tDurCalcType TDurCalcType,
							dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartTimeDurationTzCalc() "

	if startDateTime.IsZero() && duration==0 {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
				"input parameters are ZERO!")
	}

	tlz := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tlz)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"Error: 'timeZoneLocation' input parameter is INVALID! " +
				"'timeZoneLocation'='%v'  processed tzl= '%v' Error='%v'",
				timeZoneLocation, tlz, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartTimeDurationTzCalc(startDateTime,
																	duration,
																		tDurCalcType,
																			timeZoneLocation,
																				dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix +
			"Error returned by t2Dur.SetStartTimeDurationTzCalc(...) Error='%v'",
				err.Error())
	}

	return t2Dur, nil
}

// NewStartTimeDurationCalc - Creates and returns a new TimeDurationDto based on input
// parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation type.
//
// The duration value is added to 'startDateTime' in order to compute the ending date time.
//
// If 'duration' is a negative value, 'startDateTime' is converted to ending date time and
// the	actual starting date time is computed by subtracting duration.
//
// The time zone location applied to both 'startDateTime' and ending date time is derived
// from input parameter, 'startDateTime'.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//			TDurCalcTypeSTDYEARMTH 		- Default - standard year, month week,
// 																		day time calculation.
//
//			TDurCalcTypeCUMMONTHS 		- Computes cumulative months - no Years.
//
//			TDurCalcTypeCUMWEEKS  		- Computes cumulative weeks. No Years or months
//
//			TDurCalcTypeCUMDAYS				- Computes cumulative days. No Years, months or weeks.
//
//			TDurCalcTypeCUMHOURS			- Computes cumulative hours. No Years, months, weeks or days.
//
//			TDurCalcTypeGregorianYrs 	- Computes Years based on average length of a Gregorian Year
//																 		Used for very large duration values.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationCalc(startTime, duration,
// 						TDurCalcTypeSTDYEARMTH, FmtDateTimeYrMDayFmtStr)
//
//		Note:	'TDurCalcTypeSTDYEARMTH' is of type 'TDurCalcType' and signals
//						standard year month day time duration allocation.
//
// 					'FmtDateTimeYrMDayFmtStr' is a constants defined in the source
// 						file, datetimeconstants.go.
///
func (tDur TimeDurationDto) NewStartTimeDurationCalc(startDateTime time.Time,
					duration time.Duration, tDurCalcType TDurCalcType,
							dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartTimeDurationCalc() "

	if startDateTime.IsZero() && duration==0 {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
				"input parameters are ZERO!")
	}

	timeZoneLocation := startDateTime.Location().String()

	t2Dur := TimeDurationDto{}

	err := t2Dur.SetStartTimeDurationTzCalc(startDateTime, duration, tDurCalcType , timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartTimeDurationTzCalc(...) Error='%v'", err.Error())
	}

	return t2Dur, nil
}

// NewStartTimeDurationDateDto - Creates and returns a new TimeDurationDto based on input
// parameters 'startDateTime' and time duration. 'startDateTime' is of type DateTzDto.
//
// The time duration value is added to 'startDateTime' in order to compute the ending date time.
// If 'duration' is a negative value, 'startDateTime' is converted to ending date time. The
// actual starting date time is computed by subtracting duration from ending date time.
//
// Time Zone location is derived from 'startDateTime'.
//
// Note: 	This method applies the standard Time Duration allocation calculation type,
// 				'TDurCalcTypeSTDYEARMTH'. This means that duration is allocated over years,
// 				months, weeks, weekdays, date days, hours, minutes, seconds, milliseconds,
// 				microseconds and nanoseconds.	See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	DateTzDto	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDto(startTime, duration, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'FmtDateTimeYrMDayFmtStr' is a constant available in source file, datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartTimeDurationDateDto(startDateTime DateTzDto,
	duration time.Duration,	dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartTimeDuration() "

	if startDateTime.DateTime.IsZero() && duration==0 {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
				"input parameters are ZERO!")
	}

	timeZoneLocation := startDateTime.TimeZone.LocationName

	t2Dur := TimeDurationDto{}

	err := t2Dur.SetStartTimeDurationDateDtoTzCalc(startDateTime,
												duration,
													TDurCalcTypeSTDYEARMTH,
														timeZoneLocation,
															dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartTimeDurationTzCalc(...) Error='%v'", err.Error())
	}

	return t2Dur, nil
}


// NewStartTimeDurationDateDtoTz - Creates and returns a new TimeDurationDto based on input
// parameters 'startDateTime', time duration and 'timeZoneLocation'. 'startDateTime' is
// converted to the specified 'timeZoneLocation' and the duration value is added to it
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and the
// 	actual starting date time is computed by subtracting duration.
//
// The user is required to specify a common Time Zone Location for use in converting
// date times to a common frame of reference for use in subsequent time duration calculations.
//
// Note: 	This method applies the standard Time Duration allocation, 'TDurCalcTypeSTDYEARMTH'.
// 				This means that duration is allocated over years, months, weeks, weekdays, date days,
//				hours, minutes, seconds, milliseconds, microseconds and nanoseconds.
// 				See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting time
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDtoTz(startTime, duration,
// 																		TzIanaUsCentral, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants defined in
// 							datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartTimeDurationDateDtoTz(startDateTime DateTzDto,
								duration time.Duration, timeZoneLocation string,
									dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDtoTz() "

	if startDateTime.DateTime.IsZero() && duration==0 {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
				"input parameters are ZERO!")
	}

	tlz := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tlz)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"Error: 'timeZoneLocation' input parameter is INVALID! " +
				"'timeZoneLocation'='%v'  processed tzl= '%v' Error='%v'",
				timeZoneLocation, tlz, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartTimeDurationDateDtoTzCalc(startDateTime,
																	duration,
																		TDurCalcTypeSTDYEARMTH,
																			timeZoneLocation,
																					dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartTimeDurationDateDtoTzCalc(...) Error='%v'", err.Error())
	}

	return t2Dur, nil
}

// NewStartTimeDurationDateDtoTzCalc - Creates and returns a new TimeDurationDto based
// on input parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation
// type.
//
// Input parameter, 'startDateTime' is of Type DateTzDto.
//
// 'startDateTime' is converted to the specified 'timeZoneLocation' and the duration
// value is added to it in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to ending date time and
// the	actual starting date time is computed by subtracting duration.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	DateTzDto	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//			TDurCalcTypeSTDYEARMTH 		- Default - standard year, month week,
// 																		day time calculation.
//
//			TDurCalcTypeCUMMONTHS 		- Computes cumulative months - no Years.
//
//			TDurCalcTypeCUMWEEKS  		- Computes cumulative weeks. No Years or months
//
//			TDurCalcTypeCUMDAYS				- Computes cumulative days. No Years, months or weeks.
//
//			TDurCalcTypeCUMHOURS			- Computes cumulative hours. No Years, months, weeks or days.
//
//			TDurCalcTypeGregorianYrs 	- Computes Years based on average length of a Gregorian Year
//																 		Used for very large duration values.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDtoTzCalc(startTime, duration,
// 										TzIanaUsCentral, TDurCalcTypeSTDYEARMTH, FmtDateTimeYrMDayFmtStr)
//
//		Note:	'TDurCalcTypeSTDYEARMTH' is of type 'TDurCalcType' and signals
//						standard year month day time duration allocation.
//
// 					'TzIanaUsCentral' and 'FmtDateTimeYrMDayFmtStr' are constants defined in
// 						datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartTimeDurationDateDtoTzCalc(startDateTime DateTzDto,
	duration time.Duration, timeZoneLocation string, tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDtoTzCalc() "

	if startDateTime.DateTime.IsZero() && duration==0 {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
				"input parameters are ZERO!")
	}

	tlz := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tlz)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"Error: 'timeZoneLocation' input parameter is INVALID! " +
				"'timeZoneLocation'='%v'  processed tzl= '%v' Error='%v'",
				timeZoneLocation, tlz, err.Error())
	}

	t2Dur := TimeDurationDto{}

	err = t2Dur.SetStartTimeDurationDateDtoTzCalc(startDateTime,
													duration,
														tDurCalcType,
															tlz,
																dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix +
				"Error returned by t2Dur.SetStartTimeDurationDateDtoTzCalc(...) Error='%v'",
				err.Error())
	}

	return t2Dur, nil
}

// NewStartTimeDurationDateDtoCalc - Creates and returns a new TimeDurationDto based on input
// parameters, 'startDateTime', time duration, 'timeZoneLocation' and calculation type.
//
// Input parameter 'startDateTime' is of Type DateTzDto.
//
// The duration value is added to 'startDateTime' in order to compute the ending date time.
//
// If 'duration' is a negative value, 'startDateTime' is converted to ending date time and
// the	actual starting date time is computed by subtracting duration.
//
// The time zone location applied to both 'startDateTime' and ending date time is derived
// from input parameter, 'startDateTime'.
//
// The allocation of time duration to data fields, Years, Months, Weeks, WeekDays, DateDays,
// Hours, Minutes, Seconds, Milliseconds, Microseconds and Nanoseconds is controlled by the
// input parameter calculation type, 'tDurCalcType'. See 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime		 DateTz	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//			TDurCalcTypeSTDYEARMTH 		- Default - standard year, month week,
// 																		day time calculation.
//
//			TDurCalcTypeCUMMONTHS 		- Computes cumulative months - no Years.
//
//			TDurCalcTypeCUMWEEKS  		- Computes cumulative weeks. No Years or months
//
//			TDurCalcTypeCUMDAYS				- Computes cumulative days. No Years, months or weeks.
//
//			TDurCalcTypeCUMHOURS			- Computes cumulative hours. No Years, months, weeks or days.
//
//			TDurCalcTypeGregorianYrs 	- Computes Years based on average length of a Gregorian Year
//																 		Used for very large duration values.
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimeDurationDateDtoCalc(startTime,
// 											duration,
// 												TDurCalcTypeSTDYEARMTH,
// 													FmtDateTimeYrMDayFmtStr)
//
//		Note:	'TDurCalcTypeSTDYEARMTH' is of type 'TDurCalcType' and signals
//						standard year month day time duration allocation.
//
// 					'FmtDateTimeYrMDayFmtStr' is a constants defined in the source
// 						file, datetimeconstants.go.
///
func (tDur TimeDurationDto) NewStartTimeDurationDateDtoCalc(startDateTime DateTzDto,
	duration time.Duration, tDurCalcType TDurCalcType,
	dateTimeFmtStr string) (TimeDurationDto, error) {

	ePrefix := "TimeDurationDto.NewStartTimeDurationDateDtoCalc() "

	if startDateTime.DateTime.IsZero() && duration==0 {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
				"input parameters are ZERO!")
	}

	timeZoneLocation := startDateTime.TimeZone.LocationName

	t2Dur := TimeDurationDto{}

	err := t2Dur.SetStartTimeDurationDateDtoTzCalc(startDateTime,
											duration,
												tDurCalcType ,
													timeZoneLocation,
														dateTimeFmtStr)

	if err != nil {
		return TimeDurationDto{},
			fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartTimeDurationDateDtoTzCalc(...) Error='%v'", err.Error())
	}

	return t2Dur, nil
}

// NewStartTimePlusTimeDto - Creates and returns a new TimeDurationDto setting 
// the start date time, end date time and duration based on a starting date time
// and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// For the purposes of this time duration calculation, the Time Zone Location is
// extracted from the input parameter, 'startDateTime'.
//
// Note: 	This method applies the standard Time Duration allocation, 'TDurCalcTypeSTDYEARMTH'. 
// 				This means that duration is allocated over years, months, weeks, weekdays, date days,
//				hours, minutes, seconds, milliseconds, microseconds and nanoseconds. 
// 				See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	-   Starting date time. The ending date time will be computed
// 															by adding the time components of the 'plusTimeDto' to
// 															'startDateTime'.
//
// plusTimeDto		TimeDto 	- 	Time components (Years, months, weeks, days, hours etc.)
//															which will be added to 'startDateTime' to compute
//															time duration and ending date time.
//
//									type TimeDto struct {
//										Years          int // Number of Years
//										Months         int // Number of Months
//										Weeks          int // Number of Weeks
//										WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//										DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//										Hours          int // Number of Hours.
//										Minutes        int // Number of Minutes
//										Seconds        int // Number of Seconds
//										Milliseconds   int // Number of Milliseconds
//										Microseconds   int // Number of Microseconds
//										Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//										TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//																			// 	plus remaining Nanoseconds
//									}
//
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewStartTimePlusTimeDto(startTime,
// 																			plusTimeDto,
// 																				FmtDateTimeYrMDayFmtStr)
//
//		Note: 'FmtDateTimeYrMDayFmtStr' are constants available in datetimeconstants.go
//
func (tDur TimeDurationDto) NewStartTimePlusTimeDto(startDateTime time.Time,
								plusTimeDto TimeDto, dateTimeFmtStr string)	(TimeDurationDto, error) {
									
	ePrefix := "TimeDurationDto.NewStartTimePlusTimeDto() "

	if startDateTime.IsZero() && plusTimeDto.IsEmpty() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'startDateTime' and 'plusTimeDto' " +
				"input parameters are ZERO/EMPTY!")
	}

	timeZoneLocation := startDateTime.Location().String()

	t2Dur := TimeDurationDto{}
	
	err := t2Dur.SetStartTimePlusTimeDto(startDateTime,
												plusTimeDto,
													TDurCalcTypeSTDYEARMTH,
														timeZoneLocation,
																dateTimeFmtStr)
	
	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix + "Error returned by t2Dur.SetStartTimePlusTimeDto(...) Error='%v'", err.Error())
	}
	
	return t2Dur, nil
}

// NewEndTimeMinusTimeDto - Creates and returns a new TimeDurationDto setting 
// start date time, end date time and duration based on an ending date time
// and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// Note: 	This method applies the standard Time Duration allocation, 'TDurCalcTypeSTDYEARMTH'. 
// 				This means that duration is allocated over years, months, weeks, weekdays, date days,
//				hours, minutes, seconds, milliseconds, microseconds and nanoseconds. 
// 				See Type 'TDurCalcType' for details.
//
// Input Parameters:
// =================
//
// endDateTime	time.Time	-   Ending date time. The starting date time will be computed
// 														by subtracting minusTimeDto from 'endDateTime'
//
// minusTimeDto	TimeDto 	- 	Time components (Years, months, weeks, days, hours etc.)
//														which will be subtracted from 'endDateTime' to compute
//														time duration and starting date time.
//
//									type TimeDto struct {
//										Years          int // Number of Years
//										Months         int // Number of Months
//										Weeks          int // Number of Weeks
//										WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//										DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//										Hours          int // Number of Hours.
//										Minutes        int // Number of Minutes
//										Seconds        int // Number of Seconds
//										Milliseconds   int // Number of Milliseconds
//										Microseconds   int // Number of Microseconds
//										Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//										TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//																			// 	plus remaining Nanoseconds
//									}
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Example Usage:
// ==============
//
// tDurDto, err := TimeDurationDto{}.NewEndTimeMinusTimeDto(endTime, minusTimeDto, FmtDateTimeYrMDayFmtStr)
//
//		Note: 'FmtDateTimeYrMDayFmtStr' is a constant defined in source file,
// 							datetimeconstants.go.
//
func (tDur TimeDurationDto) NewEndTimeMinusTimeDto(endDateTime time.Time,
								minusTimeDto TimeDto, dateTimeFmtStr string)	(TimeDurationDto, error) {
									
	ePrefix := "TimeDurationDto.NewEndTimeMinusTimeDto() "

	if endDateTime.IsZero() && minusTimeDto.IsEmpty() {
		return TimeDurationDto{},
			errors.New(ePrefix + "Error: Both 'endDateTime' and 'minusTimeDto' " +
				"input parameters are ZERO/EMPTY!")
	}

	timeZoneLocation := endDateTime.Location().String()

	t2Dur := TimeDurationDto{}
	
	err := t2Dur.SetEndTimeMinusTimeDto(endDateTime, minusTimeDto,
										TDurCalcTypeSTDYEARMTH,timeZoneLocation, dateTimeFmtStr)
	
	if err != nil {
		return TimeDurationDto{},
		fmt.Errorf(ePrefix + "Error returned by t2Dur.SetEndTimeMinusTimeDto(...) Error='%v'", err.Error())
	}
	
	return t2Dur, nil
}


// ReCalcTimeDurationAllocation - Re-calculates and allocates time duration for the current
// TimeDurationDto instance over the various time components (years, months, weeks, weekdays,
// datedays, hour, minutes, seconds, milliseconds, microseconds and nanoseconds) depending
// on the value of the 'TDurCalcType' input parameter.
//
// Input Parameter
// ===============
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month
//																	 week day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
//
func (tDur *TimeDurationDto) ReCalcTimeDurationAllocation(calcType TDurCalcType) error {
	
	return tDur.calcTimeDurationAllocations(calcType)
	
}

// SetEndTimeMinusTimeDto - Sets start date time, end date time and duration
// based on an ending date time and the time components contained in a TimeDto.
//
// Starting date time is computed by subtracting the value of the TimeDto from
// the ending date time input parameter, 'endDateTime'.
//
// Input Parameters:
// =================
//
// endDateTime	time.Time	-   Ending date time. The starting date time will be computed
// 														by subtracting minusTimeDto from 'endDateTime'
//
// minusTimeDto	TimeDto 	- 	Time components (Years, months, weeks, days, hours etc.)
//														which will be subtracted from 'endDateTime' to compute
//														time duration and starting date time.
//
//									type TimeDto struct {
//										Years          int // Number of Years
//										Months         int // Number of Months
//										Weeks          int // Number of Weeks
//										WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//										DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//										Hours          int // Number of Hours.
//										Minutes        int // Number of Minutes
//										Seconds        int // Number of Seconds
//										Milliseconds   int // Number of Milliseconds
//										Microseconds   int // Number of Microseconds
//										Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//										TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//																			// 	plus remaining Nanoseconds
//									}
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month
//																	 week day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time comparisons.
//
// 														Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetEndTimeMinusTimeDto(endDateTime time.Time,
	minusTimeDto TimeDto, tDurCalcType TDurCalcType, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "TimeDurationDto.SetEndTimeMinusTimeDto() "

	if endDateTime.IsZero() && minusTimeDto.IsEmpty() {
		return 	errors.New(ePrefix + "Error: Both 'endDateTime' and 'minusTimeDto' " +
				"input parameters are ZERO/EMPTY!")
	}

	dtFormat := tDur.preProcessDateFormatStr(dateTimeFmtStr)
	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return 	fmt.Errorf(ePrefix +
				"Error: 'timeZoneLocation' input parameter is INVALID! " +
				"'timeZoneLocation'='%v'  processed tzLoc= '%v' Error='%v'",
				timeZoneLocation, tzLoc, err.Error())
	}


	eDateTime, err := TimeZoneDto{}.New(endDateTime, tzLoc, dtFormat)

	if err!= nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(endDateTime, tzLoc, " +
			"dtFormat). Error='%v'", err.Error())
	}

	tDur.Empty()

	tDur.EndTimeDateTz = eDateTime.TimeOut.CopyOut()

	tDur.StartTimeDateTz, err = eDateTime.TimeOut.MinusTimeDto(minusTimeDto)

	tDur.TimeDuration = tDur.EndTimeDateTz.DateTime.Sub(tDur.StartTimeDateTz.DateTime)

	err = tDur.calcTimeDurationAllocations(tDurCalcType)

	if err != nil {
		tDur.Empty()
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcTypeSTDYEARMTH(). " +
			"Error='%v'", err.Error())
	}

	return nil
}

//
// SetStartEndTimesDateDtoTzCalc - Sets data field values for the current
// TimeDurationDto instance using a Start Date Time, End Date Time and a
// time zone specification.
//
// The Starting Date Time and Ending Date Time are submitted as type 'DateTzDto'
//
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
// startDateTime	DateTzDto	- Starting time
//
// endDateTime		DateTzDto - Ending time
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month
//																	 week day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartEndTimesDateDtoTzCalc(startDateTime,
									endDateTime DateTzDto, tDurCalcType TDurCalcType,
										timeZoneLocation, dateTimeFmtStr string) error {

										ePrefix := "TimeDurationDto.SetStartEndTimesDateDtoTzCalc() "


		err := tDur.SetStartEndTimesTzCalc(startDateTime.DateTime, endDateTime.DateTime,
								tDurCalcType, timeZoneLocation, dateTimeFmtStr)

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by SetStartEndTimesTzCalc- " +
					"Error:='%v'", err.Error())
		}

		return nil
}

// SetStartEndTimesTzCalc - Sets data field values for the current TimeDurationDto
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
// startDateTime	time.Time	- Starting time
//
// endDateTime		time.Time - Ending time
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month
//																	 week day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartEndTimesTzCalc(startDateTime,
endDateTime time.Time, tDurCalcType TDurCalcType, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "TimeDurationDto.SetStartEndTimes() "

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return 	errors.New(ePrefix + "Error: Both 'startDateTime' and 'endDateTime' " +
			"input parameters are ZERO!")
	}

	dtFormat := tDur.preProcessDateFormatStr(dateTimeFmtStr)
	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: 'timeZoneLocation' input parameter is INVALID! " +
			"'timeZoneLocation'='%v'  processed tzLoc= '%v' Error='%v'",
			timeZoneLocation, tzLoc, err.Error())
	}

	sTime, err := TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat)

	if err != nil {
		return fmt.Errorf(ePrefix + 
			"Error returned by TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat). " +
			"Error='%v'", err.Error())
	}

	eTime, err := TimeZoneDto{}.New(endDateTime, tzLoc, dtFormat)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(endDateTime, tzLoc, dtFormat). " +
			"Error='%v'", err.Error())
	}
	
	if eTime.TimeOut.DateTime.Before(sTime.TimeOut.DateTime) {
		s2 := sTime.CopyOut()
		sTime = eTime.CopyOut()
		eTime = s2.CopyOut()
	}

	tDur.Empty()
	tDur.StartTimeDateTz = sTime.TimeOut.CopyOut()
	tDur.EndTimeDateTz	= eTime.TimeOut.CopyOut()
	tDur.TimeDuration = tDur.EndTimeDateTz.DateTime.Sub(tDur.StartTimeDateTz.DateTime)

	err = tDur.calcTimeDurationAllocations(tDurCalcType)
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcTypeSTDYEARMTH(). " +
				"Error='%v'", err.Error())
	}
	
	return nil
}

// SetStartTimeDurationTzCalc - Sets start time, end time and duration for the
// current TimeDurationDto instance. 'startDateTime' is converted to the
// specified 'timeZoneLocation' and the duration value is added to it
// in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to
// ending date time and the	actual starting date time is computed by
// subtracting duration.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	- Starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month
//																	 week day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
func (tDur *TimeDurationDto) SetStartTimeDurationTzCalc(startDateTime time.Time,
															duration time.Duration, tDurCalcType TDurCalcType,
																	timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "TimeDurationDto.SetStartTimeDurationTzCalc() "

	if startDateTime.IsZero() && duration==0 {
		return 	errors.New(ePrefix + "Error: Both 'startDateTime' and 'duration' " +
			"input parameters are ZERO!")
	}

	dtFormat := tDur.preProcessDateFormatStr(dateTimeFmtStr)
	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: 'timeZoneLocation' input parameter is INVALID! " +
			"'timeZoneLocation'='%v'  processed tzLoc= '%v' Error='%v'",
			timeZoneLocation, tzLoc, err.Error())
	}

	xTime, err := TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat)

	if err!= nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime, tzLoc, " +
			"dtFormat). Error='%v'", err.Error())
	}

	tDur.Empty()

	if duration < 0 {

		tDur.EndTimeDateTz = xTime.TimeOut.CopyOut()

		tDur.StartTimeDateTz, err = tDur.EndTimeDateTz.AddDuration(duration, dtFormat)

		if err != nil {
			tDur.Empty()
			return fmt.Errorf(ePrefix + "Error returned from tDur.EndTimeDateTz."+
				"AddDuration(duration, dtFormat) " +
				" Error='%v'", err.Error())

		}

		tDur.TimeDuration = duration * -1

	} else {

		tDur.StartTimeDateTz = xTime.TimeOut.CopyOut()

		tDur.EndTimeDateTz, err = tDur.StartTimeDateTz.AddDuration(duration, dtFormat)

		if err != nil {
			tDur.Empty()
			return fmt.Errorf(ePrefix + "Error returned from tDur.StartTimeDateTz."+
				"AddDuration(duration, dtFormat) " +
				" Error='%v'", err.Error())
		}

		tDur.TimeDuration = duration

	}

	err = tDur.calcTimeDurationAllocations(tDurCalcType)

	if err != nil {
		tDur.Empty()
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcTypeSTDYEARMTH(). " +
			"Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimeDurationDateDtoTzCalc - Sets start time, end time and
// duration for the current TimeDurationDto instance.
//
// The input parameter, 'startDateTime', is of type DateTzDto. It is
// converted to the specified 'timeZoneLocation' and the duration value
// is added to in order to compute the ending date time.
//
// If 'duration' is a negative value 'startDateTime' is converted to
// ending date time and the	actual starting date time is computed by
// subtracting duration.
//
// Input Parameters:
// =================
//
// startDateTime	DateTzDto	- Provides starting date time for the duration calculation
//
// duration		time.Duration - Amount of time to be added to or subtracted from
//														'startDateTime'. Note: If duration is a negative value
//														'startDateTime' is converted to ending date time and
//														actual starting date time is computed by subtracting
//														duration.
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month
//																	 week day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time duration
// 														calculations.
//
// 														Time zone location must be designated as one of two values.
//
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
func (tDur *TimeDurationDto) SetStartTimeDurationDateDtoTzCalc(startDateTime DateTzDto,
	duration time.Duration, tDurCalcType TDurCalcType,
	timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "TimeDurationDto.SetStartTimeDurationDateDtoTzCalc() "

	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: 'timeZoneLocation' input parameter is INVALID! " +
			"'timeZoneLocation'='%v'  processed tzLoc= '%v' Error='%v'",
			timeZoneLocation, tzLoc, err.Error())
	}


	err = tDur.SetStartTimeDurationTzCalc(startDateTime.DateTime,
																				duration,
																					tDurCalcType,
																						timeZoneLocation,
																							dateTimeFmtStr)

	if err != nil {
		fmt.Errorf(ePrefix + "Error returned by SetStartTimeDurationTzCalc: Error='%v'", err.Error())
	}

	return nil
}

// SetStartTimePlusTimeDto - Sets start date time, end date time and duration
// based on a starting date time and the time components contained in a TimeDto.
//
// The time components of the TimeDto are added to the starting date time to compute
// the ending date time and the duration.
//
// Input Parameters:
// =================
//
// startDateTime	time.Time	-   Starting date time. The ending date time will be computed
// 															by adding the time components of the 'plusTimeDto' to
// 															'startDateTime'.
//
// plusTimeDto		TimeDto 	- 	Time components (Years, months, weeks, days, hours etc.)
//															which will be added to 'startDateTime' to compute
//															time duration and ending date time.
//
//									type TimeDto struct {
//										Years          int // Number of Years
//										Months         int // Number of Months
//										Weeks          int // Number of Weeks
//										WeekDays       int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//										DateDays       int // Total Number of Days. Weeks x 7 plus WeekDays
//										Hours          int // Number of Hours.
//										Minutes        int // Number of Minutes
//										Seconds        int // Number of Seconds
//										Milliseconds   int // Number of Milliseconds
//										Microseconds   int // Number of Microseconds
//										Nanoseconds    int // Remaining Nanoseconds after Milliseconds & Microseconds
//										TotSubSecNanoseconds int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//																			// 	plus remaining Nanoseconds
//									}
//
// tDurCalcType TDurCalcType-	Specifies the calculation type to be used in allocating
//														time duration:
//
//					TDurCalcTypeSTDYEARMTH - Default - standard year, month
//																	 week day time calculation.
//
//					TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//					TDurCalcTypeCUMWEEKS  - Computes cumulative weeks. No Years or months
//
//					TDurCalcTypeCUMDAYS		- Computes cumulative days. No Years, months or weeks.
//
//					TDurCalcTypeCUMHOURS	- Computes cumulative hours. No Years, months, weeks or days.
//
//					TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//																		 Used for very large duration values.
//
//
// timeZoneLocation	string	- Designates the standard Time Zone location by which
//														time duration will be compared. This ensures that
//														'oranges are compared to oranges and apples are compared
//														to apples' with respect to start time and end time comparisons.
//
// 														Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//																	"Etc/UTC" = ZULU, GMT or UTC - Default
//
//														 (3)	If 'timeZoneLocation' is submitted as an empty string,
//																	it will default to "Etc/UTC" = ZULU, GMT, UTC
//
// dateTimeFmtStr string		- A date time format string which will be used
//															to format and display 'dateTime'. Example:
//															"2006-01-02 15:04:05.000000000 -0700 MST"
//
//														If 'dateTimeFmtStr' is submitted as an
//															'empty string', a default date time format
//															string will be applied. The default date time
//															format string is:
//															FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (tDur *TimeDurationDto) SetStartTimePlusTimeDto(startDateTime time.Time,
	plusTimeDto TimeDto, tDurCalcType TDurCalcType, timeZoneLocation, dateTimeFmtStr string) error {
	
	ePrefix := "TimeDurationDto.SetStartTimePlusTimeDto() "

	if startDateTime.IsZero() && plusTimeDto.IsEmpty() {
		return 	errors.New(ePrefix + "Error: Both 'startDateTime' and 'plusTimeDto' " +
			"input parameters are ZERO/EMPTY!")
	}

	dtFormat := tDur.preProcessDateFormatStr(dateTimeFmtStr)
	tzLoc := tDur.preProcessTimeZoneLocation(timeZoneLocation)

	_, err := time.LoadLocation(tzLoc)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: 'timeZoneLocation' input parameter is INVALID! " +
			"'timeZoneLocation'='%v'  processed tzLoc= '%v' Error='%v'",
			timeZoneLocation, tzLoc, err.Error())
	}


	sDateTime, err := TimeZoneDto{}.New(startDateTime, tzLoc, dtFormat)

	if err!= nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDto{}.New(startDateTime, tzLoc, " +
			"dtFormat). Error='%v'", err.Error())
	}

	tDur.Empty()
	
	tDur.StartTimeDateTz = sDateTime.TimeOut.CopyOut()
	
	tDur.EndTimeDateTz, err = sDateTime.TimeOut.PlusTimeDto(plusTimeDto)
	
	tDur.TimeDuration = tDur.EndTimeDateTz.DateTime.Sub(tDur.StartTimeDateTz.DateTime)

	err = tDur.calcTimeDurationAllocations(tDurCalcType)

	if err != nil {
		tDur.Empty()
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcTypeSTDYEARMTH(). " +
			"Error='%v'", err.Error())
	}

	return nil		
}

func (tDur *TimeDurationDto) calcTimeDurationAllocations(calcType TDurCalcType) error {

	ePrefix := "TimeDurationDto.calcTimeDurationAllocations() "

	switch calcType {

	case TDurCalcTypeSTDYEARMTH :
		return tDur.calcTypeSTDYEARMTH()

	case TDurCalcTypeCUMMONTHS :
		return tDur.calcTypeCUMMONTHS()

	case TDurCalcTypeCUMWEEKS :
		return tDur.calcTypeCUMWEEKS()

	case TDurCalcTypeCUMDAYS :
		return tDur.calcTypeCUMDays()

	case TDurCalcTypeCUMHOURS :
		return tDur.calcTypeCUMHours()

	case TDurCalcTypeCUMMINUTES:
		return tDur.calcTypeCUMMINUTES()

	case TDurCalcTypeGregorianYrs :
		return tDur.calcTypeGregorianYears()

	default:
		return fmt.Errorf(ePrefix + "Error: Invalid TDurCalcType. calcType='%v'", calcType.String())
	}

	return nil
}

// calcTypeSTDYEARMTH - Performs Duration calculations for
// TDurCalcType == TDurCalcTypeSTDYEARMTH
//
// TDurCalcTypeYEARMTH - Standard Year, Month, Weeks, Days calculation.
// All data fields in the TimeDto are populated in the duration
// allocation.
func (tDur *TimeDurationDto) calcTypeSTDYEARMTH() error {
	
	ePrefix := "TimeDurationDto.calcTypeSTDYEARMTH() "

	tDur.EmptyTimeFields()
	
	tDur.CalcType = TDurCalcTypeSTDYEARMTH
	
	err := tDur.calcYearsFromDuration()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcYearsFromDuration(). Error='%v'", err.Error())
	}
	
	err = tDur.calcMonthsFromDuration()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}
	
	err = tDur.calcDateDaysWeeksFromDuration()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}
	
	err = tDur.calcHoursMinSecs()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}
	
	err = tDur.calcNanoseconds()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}
	
	err = tDur.calcSummaryTimeElements()
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}
	
	return nil
}

// Data Fields for Years is always set to Zero. Years
// and months are consolidated and counted as cumulative
// months.
func (tDur *TimeDurationDto) calcTypeCUMMONTHS() error {

	ePrefix := "TimeDurationDto.calcTypeCUMWEEK() "

	tDur.EmptyTimeFields()

	tDur.CalcType = TDurCalcTypeCUMMONTHS
	
	err := tDur.calcMonthsFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcDateDaysWeeksFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil

}

// calcTypeCUMWEEKS - Data Fields for Years and Months are always set to zero.
// Years and Months are consolidated and counted as equivalent Weeks.
func (tDur *TimeDurationDto) calcTypeCUMWEEKS() error {

	ePrefix := "TimeDurationDto.calcTypeCUMWEEKS() "

	tDur.EmptyTimeFields()
	
	tDur.CalcType = TDurCalcTypeCUMWEEKS

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
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	// For Cumulative Weeks calculation and presentations, set Date Days to zero
	tDur.DateDays = 0
	tDur.DateDaysNanosecs = 0

	return nil
}

// calcTypeCUMDays - Calculates Cumulative Days. Years, months and weeks are consolidated
// and counted as cumulative days. The Data Fields for years, months, weeks and week days
// are set to zero.  All cumulative days are allocated to the data field, 'DateDays'.
func (tDur *TimeDurationDto) calcTypeCUMDays() error {

	ePrefix := "TimeDurationDto.calcTypeCUMDays() "

	tDur.EmptyTimeFields()

	tDur.CalcType = TDurCalcTypeCUMDAYS
	
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
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}

// calcTypeCUMHours - Calculates Cumulative Hours. Years, months, weeks, week days,
// date days and hours are consolidated and included in cumulative hours. Values for years,
// months, weeks, week days and date days are ignored and set to zero. Time duration is
// allocated over cumulative hours plus minutes, seconds, milliseconds, microseconds and
// nanoseconds.
func (tDur *TimeDurationDto) calcTypeCUMHours() error {

	ePrefix := "TimeDurationDto.calcTypeCUMHours() "

	tDur.EmptyTimeFields()

	tDur.CalcType = TDurCalcTypeCUMHOURS

	if tDur.TimeDuration == 0 {
		return nil
	}

	err := tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}

// calcTypeCUMMINUTES - Calculates Cumulative Minutes. Years, months, weeks, week days,
// date days, hours and minutes are consolidated and included in cumulative minutes.
// Values for years, months, weeks, week days, date days and hours are ignored and set
// to zero. Time duration is allocated over cumulative minutes plus seconds, milliseconds,
// microseconds and nanoseconds.
func (tDur *TimeDurationDto) calcTypeCUMMINUTES() error {

	ePrefix := "TimeDurationDto.calcTypeCUMHours() "

	tDur.EmptyTimeFields()

	tDur.CalcType = TDurCalcTypeCUMMINUTES

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
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil

}


// calcTypeGregorianYears - Allocates Years using the number of nanoseconds in a
// standard or average GregorianYear
func (tDur *TimeDurationDto) calcTypeGregorianYears() error {
	ePrefix := "TimeDurationDto.calcTypeGregorianYears() "

	tDur.EmptyTimeFields()

	tDur.CalcType = TDurCalcTypeGregorianYrs
	
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
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcMonthsFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcDateDaysWeeksFromDuration()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcDateDaysWeeksFromDuration(). Error='%v'", err.Error())
	}

	err = tDur.calcHoursMinSecs()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcHoursMinSecs(). Error='%v'", err.Error())
	}

	err = tDur.calcNanoseconds()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcNanoseconds(). Error='%v'", err.Error())
	}

	err = tDur.calcSummaryTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tDur.calcSummaryTimeElements(). Error='%v'", err.Error())
	}

	return nil
}

// calcYearsFromDuration - Calculates number of years duration and nanoseconds
// represented by years duration using input parameters 'tDur.StartTimeDateTz' and
// 'tDur.EndTimeDateTz'.  
//
// NOTE:	Before calling this method, ensure that tDur.StartTimeDateTz,
//				tDur.EndTimeDateTz and tDur.TimeDuration are properly initialized.
//
func (tDur *TimeDurationDto) calcYearsFromDuration() error {

	ePrefix := "TimeDurationDto.calcYearsFromDuration() "

	years := int64(0)
	yearNanosecs := int64(0)
	startTime := tDur.StartTimeDateTz.DateTime
	endTime := tDur.EndTimeDateTz.DateTime

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix + "Error: 'startTime' and 'endTime' Time Zone Location do NOT match! " +
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
// NOTE:	(1) Before calling this method, ensure that tDur.StartTimeDateTz,
//						tDur.EndTimeDateTz and tDur.TimeDuration are properly initialized.
// 
//				(2) Before calling this method, ensure that the following method is called
//						first:
//										TimeDurationDto.calcYearsFromDuration
//
func (tDur *TimeDurationDto) calcMonthsFromDuration() error {

	ePrefix := "TimeDurationDto.calcMonthsFromDuration() "

	startTime := tDur.StartTimeDateTz.DateTime
	endTime := tDur.EndTimeDateTz.DateTime

	if endTime.Before(startTime) {
		return errors.New(ePrefix + "Error: 'endTime' precedes, is less than, startTime!")
	}

	if startTime.Location().String() != endTime.Location().String() {
		return fmt.Errorf(ePrefix + "Error: 'startTime' and 'endTime' Time Zone Location do NOT match! " +
			"startTimeZoneLocation='%v'  endTimeZoneLocation='%v'",
			startTime.Location().String(), endTime.Location().String())
	}

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return errors.New(ePrefix + "Error: tDur.TimeDuration is ZERO!")
	}

	rd -= tDur.YearsNanosecs

	i := 0

	yearDateTime := startTime.Add(time.Duration(tDur.YearsNanosecs))

	mthDateTime := yearDateTime

	for mthDateTime.Before(endTime) {

		i++

		mthDateTime = yearDateTime.AddDate(0,i,0)

	}

	i -= 1

	if i > 0 {

		tDur.Months = int64(i)

		mthDateTime = yearDateTime.AddDate( 0, i, 0)

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
// NOTE:	(1) Before calling this method, ensure that TimeDurationDto.StartTimeDateTz,
//						TimeDurationDto.EndTimeDateTz and tDur.TimeDuration are properly initialized.
// 
//				(2) Before calling this method, ensure that the following methods are called
//						first, in sequence:
//							TimeDurationDto.calcYearsFromDuration
//							TimeDurationDto.calcMonthsFromDuration
//
func (tDur *TimeDurationDto) calcDateDaysWeeksFromDuration() error {

	ePrefix := "TimeDurationDto.calcDateDaysFromDuration() "

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return errors.New(ePrefix + "Error: tDur.TimeDuration is ZERO!")
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

		tDur.WeekDays = tDur.DateDays -  (tDur.Weeks * 7)
		tDur.WeekDaysNanosecs = tDur.WeekDays * DayNanoSeconds
		
	}

	return nil
}

// calcHoursMinSecs - Calculates Hours, Minute, and 
// Seconds of duration using startTime, tDur.StartTimeDateTz, 
// and endTime, tDur.EndTimeDateTz.DateTime.
//
//
// NOTE:	(1) Before calling this method, ensure that tDur.StartTimeDateTz,
//						TimeDurationDto.EndTimeDateTz and TimeDurationDto.TimeDuration
// 						are properly initialized.
// 
//				(2) Before calling this method, ensure that the following methods are called
//						first, in sequence:
//							TimeDurationDto.calcYearsFromDuration
//							TimeDurationDto.calcMonthsFromDuration
//							TimeDurationDto.calcDateDaysWeeksFromDuration
//
func (tDur *TimeDurationDto) calcHoursMinSecs() error {
	
	ePrefix := "TimeDurationDto.calcHoursMinSecs() "

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return errors.New(ePrefix + "Error: tDur.TimeDuration is ZERO!")
	}

	if tDur.DateDays > 0 {
		rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs +
			tDur.DateDaysNanosecs
	} else {
		rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs +
			tDur.WeeksNanosecs + tDur.WeekDaysNanosecs
	}

	tDur.Hours 						= 0
	tDur.HoursNanosecs 		= 0
	tDur.Minutes 					= 0
	tDur.MinutesNanosecs 	= 0
	tDur.Seconds 					= 0 
	tDur.SecondsNanosecs 	= 0

	if rd >= HourNanoSeconds {
		tDur.Hours = rd / HourNanoSeconds
		tDur.HoursNanosecs = HourNanoSeconds * tDur.Hours
		rd -= tDur.HoursNanosecs
	}

	if rd >= MinuteNanoSeconds {
		tDur.Minutes = rd / MinuteNanoSeconds
		tDur.MinutesNanosecs = MinuteNanoSeconds * tDur.Minutes
		rd -=tDur.MinutesNanosecs
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
//
// NOTE:	(1) Before calling this method, ensure that tDur.StartTimeDateTz,
//						TimeDurationDto.EndTimeDateTz and TimeDurationDto.TimeDuration
// 						are properly initialized.
// 
//				(2) Before calling this method, ensure that the following methods are called
//						first, in sequence:
//							TimeDurationDto.calcYearsFromDuration
//							TimeDurationDto.calcMonthsFromDuration
//							TimeDurationDto.calcDateDaysWeeksFromDuration
//							TimeDurationDto.calcHoursMinSecs
// 
func (tDur *TimeDurationDto) calcNanoseconds() error {

	ePrefix := "TimeDurationDto.calcNanoseconds() "

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return errors.New(ePrefix + "Error: tDur.TimeDuration is ZERO!")
	}


	rd -= tDur.YearsNanosecs + tDur.MonthsNanosecs

	if tDur.DateDaysNanosecs > 0 {
		rd -= tDur.DateDaysNanosecs
	} else {
		rd -= tDur.WeeksNanosecs + tDur.WeekDaysNanosecs
	}

	rd -= tDur.HoursNanosecs +
					tDur.MinutesNanosecs + tDur.SecondsNanosecs
					
	tDur.Milliseconds 					= 0
	tDur.MillisecondsNanosecs 	= 0
	tDur.Microseconds						= 0
	tDur.MicrosecondsNanosecs		= 0
	tDur.Nanoseconds 						= 0
	
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

// calcSummaryTimeElements - Calculates totals for Date, Time and 
// sub-second nanoseconds. 
//
// NOTE:	(1) Before calling this method, ensure that tDur.StartTimeDateTz,
//						TimeDurationDto.EndTimeDateTz and TimeDurationDto.TimeDuration
// 						are properly initialized.
// 
//				(2) Before calling this method, ensure that the following methods are called
//						first, in sequence:
//							TimeDurationDto.calcYearsFromDuration
//							TimeDurationDto.calcMonthsFromDuration
//							TimeDurationDto.calcDateDaysWeeksFromDuration
//							TimeDurationDto.calcHoursMinSecs
//							TimeDurationDto.calcNanoseconds
//
func (tDur *TimeDurationDto) calcSummaryTimeElements() error {
	
	ePrefix := "TimeDurationDto.calcSummaryTimeElements() "

	rd := int64(tDur.TimeDuration)

	if rd == 0 {
		return errors.New(ePrefix + "Error: tDur.TimeDuration is ZERO!")
	}
	
	tDur.TotDateNanoseconds = 0
	tDur.TotTimeNanoseconds = 0
	tDur.TotSubSecNanoseconds = 0

	tDur.TotDateNanoseconds = tDur.YearsNanosecs
	tDur.TotDateNanoseconds += tDur.MonthsNanosecs

	if tDur.DateDaysNanosecs == 0  {
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


func (tDur *TimeDurationDto) preProcessDateFormatStr(dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}


func (tDur *TimeDurationDto) preProcessTimeZoneLocation(timeZoneLocation string) string {

	if len(timeZoneLocation) == 0 {
		return TzIanaUTC
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return "Local"
	}

	return timeZoneLocation
}
