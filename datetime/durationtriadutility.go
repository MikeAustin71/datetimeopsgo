package datetime

import (
	"fmt"
	"sync"
	"time"
)

type durationTriadUtility struct {

	lock  sync.Mutex

}


// copyIn - Copies the data fields from 'durT2' into
// the data fields of 'durT1'.
//
func(durTUtil *durationTriadUtility) copyIn(
	durT1 *DurationTriad,
	durT2 *DurationTriad,
	ePrefix string) {

	durTUtil.lock.Lock()

	defer durTUtil.lock.Unlock()

	if durT1 == nil{
		panic(ePrefix + "Input Parameter 'durT1' is a nil pointer.")
	}

	if durT2 == nil{
		panic(ePrefix + "Input Parameter 'durT2' is a nil pointer.")
	}

	durTUtil2 := durationTriadUtility{}

	durTUtil2.empty(durT1, ePrefix)

	durT1.BaseTime = durT2.BaseTime.CopyOut()
	durT1.LocalTime = durT2.LocalTime.CopyOut()
	durT1.UTCTime = durT2.UTCTime.CopyOut()
}

// copyOut - Returns a deep copy of input parameter
// 'durT'.
//
func(durTUtil *durationTriadUtility) copyOut(
	durT *DurationTriad,
	ePrefix string) DurationTriad {

	durTUtil.lock.Lock()

	defer durTUtil.lock.Unlock()

	if durT == nil {
		panic(ePrefix +
			"Error: Input parameter 'durT' is a 'nil' pointer!\n")
	}

	durationTriadOut := DurationTriad{}
	durationTriadOut.BaseTime = durT.BaseTime.CopyOut()
	durationTriadOut.LocalTime = durT.LocalTime.CopyOut()
	durationTriadOut.UTCTime = durT.UTCTime.CopyOut()

return durationTriadOut
}

// Empty - This method initializes all of the data fields
// in the input parameter DurationTriad structure ('durT')
// to their zero or uninitialized values.
//
func(durTUtil *durationTriadUtility) empty(
	durT *DurationTriad,
	ePrefix string) {

	durTUtil.lock.Lock()

	defer durTUtil.lock.Unlock()

	if durT == nil {
		panic(ePrefix + "Input parameter 'durT' is a 'nil' pointer!")
	}

	durT.BaseTime.Empty()
	durT.LocalTime.Empty()
	durT.UTCTime.Empty()

	return
}

// isValid - Analyzes input DurationTriad 'durT' to
// determine if it is valid. If the DurationTriad
// instance is found to be 'invalid', an error is
// returned.
//
func(durTUtil *durationTriadUtility) isValid(
	durT *DurationTriad,
	ePrefix string) error {

	durTUtil.lock.Lock()

	defer durTUtil.lock.Unlock()

	ePrefix += "durationTriadUtility.isValid() "

	if durT == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "durT",
			inputParameterValue: "",
			errMsg:              "Input parameter 'durT' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	err := durT.BaseTime.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"INVALID durT.BaseTime. Error='%v'", err.Error())
	}

	err = durT.LocalTime.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"INVALID durT.LocalTime. Error='%v'", err.Error())
	}

	err = durT.UTCTime.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"INVALID durT.UTCTime. Error='%v'", err.Error())
	}

	return nil

}

// setEndTimeMinusTimeDtoCalcTz - Calculates duration values based on an Ending Date Time and
// a TimeDto structure consisting of time values (Years, Months, weeks, days, hours,
// minutes etc.). The time values in the 'timeDto' parameter are subtracted
// from 'endDateTime'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time.
//
// As a result, true values for StartTimeDateTz, EndTimeDateTz and TimeDuration
// are stored in the DurationTriad data structure.
//
// Input parameter, 'timeZoneLocation', is applied to both the starting and ending
// date times before computing date time duration. This ensures accuracy in
// time duration calculations.
//
// The allocation of time duration to years, months, weeks, days, hours etc.
// is controlled by the input parameter calculation type, 'tDurCalcType'.
// For most purposes, the calculation type 'TDurCalcType(0).StdYearMth()' will
// suffice. For details see Type 'TDurCalcType' which is located in
// source file:
//   MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// durT          *DurationTriad -
//                           The results of this calculation will be stored
//                           in this instance of 'DurationTriad'
//
// endDateTime   time.Time - The ending date time value from which TimeDto
//                            parameter 'minusTimeDto' will be subtracted
//                            in order to compute the Starting Date Time.
//
// minusTimeDto  TimeDto   - An instance of TimeDto containing time values,
//                            (Years, Months, weeks, days, hours, minutes etc.)
//                            which will be subtracted from input parameter
//                            'endDateTime' in order to compute the Starting
//                            Date Time and Time Duration.
//
//                           type TimeDto struct {
//                             Years                  int // Number of Years
//                             Months                 int // Number of Months
//                             Weeks                  int // Number of Weeks
//                             WeekDays               int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//                             DateDays               int // Total Number of Days. Weeks x 7 plus WeekDays
//                             Hours                  int // Number of Hours.
//                             Minutes                int // Number of Minutes
//                             Seconds                int // Number of Seconds
//                             Milliseconds           int // Number of Milliseconds
//                             Microseconds           int // Number of Microseconds
//                             Nanoseconds            int // Remaining Nanoseconds after Milliseconds & Microseconds
//                             TotSubSecNanoseconds   int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                                        // plus remaining Nanoseconds
//                            }
//
//                            Type 'TimeDto' is located in source file:
//                             MikeAustin71\datetimeopsgo\datetime\timedto.go
//
// tDurCalcType  TDurCalcType - Specifies the calculation type to be used in allocating
//         time duration:
//
//    TDurCalcType(0).StdYearMth() - Default - standard year, month week,
//              day time calculation.
//
//    TDurCalcTypeCUMMONTHS - Computes cumulative months - no Years.
//
//    TDurCalcTypeCUMWEEKS - Computes cumulative weeks. No Years or months
//
//    TDurCalcTypeCUMDAYS - Computes cumulative days. No Years, months or weeks.
//
//    TDurCalcTypeCUMHOURS - Computes cumulative hours. No Years, months, weeks or days.
//
//    TDurCalcTypeCUMMINUTES - Computes cumulative minutes. No Years, months, weeks, days
//         or hours.
//
//    TDurCalcTypeCUMSECONDS - Computes cumulative seconds. No Years, months, weeks, days,
//         hours or minutes.
//
//    TDurCalcTypeGregorianYrs - Computes Years based on average length of a Gregorian Year
//          Used for very large duration values.
//
//    Type 'TDurCalcType' is located in source file:
//     MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
//
// timeZoneLocation string - time zone location must be designated as one of
//                           three types of time zones:
//
//               (1) the string 'Local' - signals the designation of the
//                   time zone location used by the host computer.
//
//               (2) IANA Time Zone Location -
//                   See https://golang.org/pkg/time/#LoadLocation
//                   and https://www.iana.org/time-zones to ensure that
//                   the IANA Time Zone Database is properly configured
//                   on your system. Note: IANA Time Zone Data base is
//                   equivalent to 'tz database'.
//                      Examples:
//                        "America/New_York"
//                        "America/Chicago"
//                        "America/Denver"
//                        "America/Los_Angeles"
//                        "Pacific/Honolulu"
//
//               (3) A Military Time Zone
//                   Reference:
//                    https://en.wikipedia.org/wiki/List_of_military_time_zones
//                    http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//                    https://www.timeanddate.com/time/zones/military
//                    https://www.timeanddate.com/worldclock/timezone/alpha
//                    https://www.timeanddate.com/time/map/
//
//                    Examples:
//                      "Alpha"   or "A"
//                      "Bravo"   or "B"
//                      "Charlie" or "C"
//                      "Delta"   or "D"
//                      "Zulu"    or "Z"
//
// dateTimeFmtStr string   - A date time format string which will be used
//                           to format and display 'dateTime'. Example:
//                           "2006-01-02 15:04:05.000000000 -0700 MST"
//
//                           Date time format constants are found in the source
//                           file 'constantsdatetime.go'. These constants represent
//                           the more commonly used date time string formats. All
//                           Date Time format constants begin with the prefix
//                           'FmtDateTime'.
//
//                           If 'dateTimeFmtStr' is submitted as an
//                           'empty string', a default date time format
//                           string will be applied. The default date time
//                           format string is:
//                             FmtDateTimeYrMDayFmtStr =
//                                 "2006-01-02 15:04:05.000000000 -0700 MST"
//
// ------------------------------------------------------------------------
//
// Return XValue
//
// error    - If this method completes successfully, the returned error
//            Type is set equal to 'nil'. If an error condition is encountered,
//            this method will return an error Type which encapsulates an
//            appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
// dt  := DurationTriad{}
// err := dt.SetEndTimeMinusTimeDtoCalcTz(
//        startDateTime,
//        minusTimeDto,
//        TDurCalcType(0).StdYearMth(),
//        TZones.US.Central(),
//        FmtDateTimeYrMDayFmtStr)
//
//
// Note: 'TDurCalcType(0).StdYearMth()' is of type 'TDurCalcType' which is defined in
//    source file 'timedurationdto.go'.
//
//       TZones.US.Central() = "America/Chicago"
//       FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       'TZones.US.Central()' and 'FmtDateTimeYrMDayFmtStr' are constants available in
//       source file 'constantsdatetime.go'.
//
func(durTUtil *durationTriadUtility) setEndTimeMinusTimeDtoCalcTz(
	durT *DurationTriad,
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) error {

	durTUtil.lock.Lock()

	defer durTUtil.lock.Unlock()

	ePrefix += "durationTriadUtility.setEndTimeMinusTimeDtoCalcTz() "

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	tzDefUtil := timeZoneDefUtility{}

	tzDef, err := tzDefUtil.newFromTimeZoneName(
		time.Now().UTC(),
		timeZoneLocation,
		TzConvertType.Relative(),
		ePrefix)

	if err != nil {
		return err
	}

	timeZoneLocation = tzDef.originalTimeZone.locationName

	if tzDef.originalTimeZone.locationNameType == LocNameType.NonConvertibleTimeZone() {
		timeZoneLocation = tzDef.convertibleTimeZone.locationName
	}

	baseTime, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoCalcTz(endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error computing baseTime. TimeDurationDto{}.NewEndTimeMinusTimeDtoTz(). "+
			"endDateTime='%v' Error='%v'",
			endDateTime, err.Error())
	}

	localTime, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoCalcTz(endDateTime,
		minusTimeDto,
		tDurCalcType,
		TZones.Local(),
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error computing localTime. TimeDurationDto{}.NewEndTimeMinusTimeDtoTz(). "+
			"endDateTime='%v' Error='%v'",
			endDateTime, err.Error())
	}

	utcTime, err := TimeDurationDto{}.NewEndTimeMinusTimeDtoCalcTz(
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		TZones.UTC(),
		dateTimeFmtStr)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error computing localTime. TimeDurationDto{}.NewEndTimeMinusTimeDtoTz(). "+
			"endDateTime='%v' Error='%v'",
			endDateTime, err.Error())
	}

	durTUtil2 := durationTriadUtility{}

	durTUtil2.empty(durT, ePrefix)

	durT.BaseTime = baseTime.CopyOut()
	durT.LocalTime = localTime.CopyOut()
	durT.UTCTime = utcTime.CopyOut()

	err = durT.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: Duration Base Data INVALID! Error='%v'", err.Error())
	}

	return nil
}