package datetime

import (
	"errors"
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

	if durT1.lock == nil {
		durT1.lock = new(sync.Mutex)
	}

	if durT2 == nil{
		panic(ePrefix + "Input Parameter 'durT2' is a nil pointer.")
	}

	if durT2.lock == nil {
		durT2.lock = new(sync.Mutex)
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

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durationTriadOut := DurationTriad{}
	durationTriadOut.BaseTime = durT.BaseTime.CopyOut()
	durationTriadOut.LocalTime = durT.LocalTime.CopyOut()
	durationTriadOut.UTCTime = durT.UTCTime.CopyOut()

return durationTriadOut
}

// Empty - This method initializes all data fields
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

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	durT.BaseTime.Empty()
	durT.LocalTime.Empty()
	durT.UTCTime.Empty()

	return
}

// Equal - This method may be used to determine if two
// DurationTriad data structures are equivalent.
//
// If input parameter 'duIn' is equal in all respects to
// the current DurationTriad instance, a boolean value of
// 'true'. Otherwise, the method returns, 'false'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  durT1 *DurationTriad
//     - An instance of DurationTriad which will be compared
//       to parameter 'durT2' in order to determine if all
//       relevant data values are equivalent.
//
//  durT2 *DurationTriad
//     - An instance of DurationTriad which will be compared
//       to parameter 'durT1' in order to determine if all
//       relevant data values are equivalent.
//
// __________________________________________________________________________
//
// Return Values
//
// bool
//     - If the method returns 'true' in signals that both the input parameter
//       DurationTriad and the current DurationTriad instance have equivalent
//       data values.
//
//        If the method returns 'false' the two DurationTriad instances are NOT
//        equal.
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
func(durTUtil *durationTriadUtility) equal(
	durT1,
	durT2 *DurationTriad,
	ePrefix string) (bool, error) {

durTUtil.lock.Lock()

defer durTUtil.lock.Unlock()

ePrefix += "durationTriadUtility.equal() "

	if durT1 == nil {
		return false,
				&InputParameterError{
					ePrefix:             ePrefix,
					inputParameterName:  "durT1",
					inputParameterValue: "",
					errMsg:              "Input parameter 'durT1' is a 'nil' pointer!",
					err:                 nil,
				}
}

	if durT1.lock == nil {
		durT1.lock = new(sync.Mutex)
	}

	if durT2 == nil {
		return false,
				&InputParameterError{
					ePrefix:             ePrefix,
					inputParameterName:  "durT2",
					inputParameterValue: "",
					errMsg:              "Input parameter 'durT2' is a 'nil' pointer!",
					err:                 nil,
				}
}

	if durT2.lock == nil {
		durT2.lock = new(sync.Mutex)
	}

	if durT1.BaseTime.Equal(durT2.BaseTime) &&
			durT1.LocalTime.Equal(durT2.LocalTime) &&
			durT1.UTCTime.Equal(durT2.UTCTime) {

		return true, nil
	}

	return false, nil
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

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
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

// setEndTimeMinusTimeDto - Calculates duration values based on an Ending Date Time and
// a TimeDto structure consisting of time values (Years, Months, weeks, days, hours,
// minutes etc.). The time values in the 'timeDto' parameter are subtracted
// from 'endDateTime'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time.
//
// As a result, true values for startDateTimeTz, endDateTimeTz and timeDuration
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
//  durT              *DurationTriad
//       - The results of this calculation will be stored
//         in this instance of 'DurationTriad'
//
//
//  endDateTime       time.Time
//       - The ending date time value from which TimeDto
//         parameter 'minusTimeDto' will be subtracted
//         in order to compute the Starting Date Time.
//
//
//  minusTimeDto      TimeDto
//       - An instance of TimeDto containing time values,
//         (Years, Months, weeks, days, hours, minutes etc.)
//         which will be subtracted from input parameter
//         'endDateTime' in order to compute the Starting
//         Date Time and Time Duration.
//
//         type TimeDto struct {
//           Years                  int // Number of Years
//           Months                 int // Number of Months
//           Weeks                  int // Number of Weeks
//           WeekDays               int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//           DateDays               int // Total Number of Days. Weeks x 7 plus WeekDays
//           Hours                  int // Number of Hours.
//           Minutes                int // Number of Minutes
//           Seconds                int // Number of Seconds
//           Milliseconds           int // Number of Milliseconds
//           Microseconds           int // Number of Microseconds
//           Nanoseconds            int // Remaining Nanoseconds after Milliseconds & Microseconds
//           TotSubSecNanoseconds   int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                      // plus remaining Nanoseconds
//         }
//
//         Type 'TimeDto' is located in source file:
//          MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType      TDurCalcType
//       - Specifies the calculation type to be used in allocating
//         time duration:
//
//    TDurCalcType(0).StdYearMth()
//       - Default - standard year, month week,
//         day time calculation.
//
//    TDurCalcTypeCUMMONTHS
//       - Computes cumulative months - no Years.
//
//    TDurCalcTypeCUMWEEKS
//       - Computes cumulative weeks. No Years or months
//
//    TDurCalcTypeCUMDAYS
//       - Computes cumulative days. No Years, months or weeks.
//
//    TDurCalcTypeCUMHOURS
//       - Computes cumulative hours. No Years, months, weeks or days.
//
//    TDurCalcTypeCUMMINUTES
//       - Computes cumulative minutes. No Years, months, weeks, days
//         or hours.
//
//    TDurCalcTypeCUMSECONDS
//       - Computes cumulative seconds. No Years, months, weeks, days,
//         hours or minutes.
//
//    TDurCalcTypeGregorianYrs
//       - Computes Years based on average length of a Gregorian Year
//         Used for very large duration values.
//
//
//  timeZoneLocation  string
//       - Time zone location must be designated as one of
//         three types of time zones:
//
//         (1) the string 'Local' - signals the designation of the
//             time zone location used by the host computer.
//
//         (2) IANA Time Zone Location -
//             See https://golang.org/pkg/time/#LoadLocation
//             and https://www.iana.org/time-zones to ensure that
//             the IANA Time Zone Database is properly configured
//             on your system. Note: IANA Time Zone Data base is
//             equivalent to 'tz database'.
//                Examples:
//                  "America/New_York"
//                  "America/Chicago"
//                  "America/Denver"
//                  "America/Los_Angeles"
//                  "Pacific/Honolulu"
//
//         (3) A Military Time Zone
//             Reference:
//              https://en.wikipedia.org/wiki/List_of_military_time_zones
//              http://www.thefightschool.demon.co.uk/UNMC_Military_Time.htm
//              https://www.timeanddate.com/time/zones/military
//              https://www.timeanddate.com/worldclock/timezone/alpha
//              https://www.timeanddate.com/time/map/
//
//              Examples:
//                "Alpha"   or "A"
//                "Bravo"   or "B"
//                "Charlie" or "C"
//                "Delta"   or "D"
//                "Zulu"    or "Z"
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
//         Date time format constants are found in the source
//         file 'constantsdatetime.go'. These constants represent
//         the more commonly used date time string formats. All
//         Date Time format constants begin with the prefix
//         'FmtDateTime'.
//
//         If 'dateTimeFmtStr' is submitted as an
//         'empty string', a default date time format
//         string will be applied. The default date time
//         format string is:
//           FmtDateTimeYrMDayFmtStr =
//               "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  ePrefix           string
//       - The error prefix containing the names of all
//         the methods executed up to this point.
//
//
// ------------------------------------------------------------------------
//
// Return Value
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
//  dt  := DurationTriad{}
//  err := dt.SetEndTimeMinusTimeDto(
//         durT,
//         timeCalcMode,
//         startDateTime,
//         minusTimeDto,
//         TDurCalcType(0).StdYearMth(),
//         TZones.US.Central(),
//         timeCalcMode,
//         FmtDateTimeYrMDayFmtStr)
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
func(durTUtil *durationTriadUtility) setEndTimeMinusTimeDto(
	durT *DurationTriad,
	endDateTime time.Time,
	minusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string,
	ePrefix string) error {

	durTUtil.lock.Lock()

	defer durTUtil.lock.Unlock()

	ePrefix += "durationTriadUtility.setEndTimeMinusTimeDto() "

	if durT == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "durT",
			inputParameterValue: "",
			errMsg:              "Input parameter 'durT' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
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

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	durT2 := DurationTriad{}

	tDurDtoUtil := timeDurationDtoUtility{}

	err := tDurDtoUtil.setEndTimeMinusTimeDto(
		&durT2.BaseTime,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "Base Time Creation- ")

	if err != nil {
		return err
	}

	err = tDurDtoUtil.setEndTimeMinusTimeDto(
		&durT2.LocalTime,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		TZones.Local(),
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "Local Time Creation- ")

	if err != nil {
		return err
	}

	err = tDurDtoUtil.setEndTimeMinusTimeDto(
		&durT2.UTCTime,
		endDateTime,
		minusTimeDto,
		tDurCalcType,
		TZones.Etc.UTC(),
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "UTC Time Creation- ")

	if err != nil {
		return err
	}

	durTUtil2 := durationTriadUtility{}

	err = durTUtil2.isValid(
		&durT2,
		ePrefix + "durT2 Validity Check - ")

	if err != nil {
		return err
	}

	durTUtil2.empty(durT, ePrefix)

	durTUtil2.copyIn(durT, &durT2, ePrefix)

	return nil
}

// setStartEndTimes - Calculates duration values and save the results in the data fields
// of the current DurationTriad instance. Calculations are based on a starting date time and
// an ending date time passed to the method. This method requires the user to specify a
// 'timeZoneLocation' input parameter which ensures that starting date time and ending date
// time will be converted to a common Time Zone before being used to compute time duration.
//
// The allocation of time duration to years, months, weeks, days, hours etc.
// is controlled by the input parameter calculation type, 'tDurCalcType'.
// For most purposes, the calculation type 'TDurCalcType(0).StdYearMth()' will
// suffice. For details see Type 'TDurCalcType' which is located in
// source file:
//     MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  durT       *DurationTriad
//     - The results of this calculation will be stored
//       in this instance of 'DurationTriad'
//
//  startDateTime   time.Time
//     - Starting date time
//
//  endDateTime     time.Time
//     - Ending date time
//
//  tDurCalcType TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation   string
//     - Designates the standard Time Zone location by which
//       time duration will be compared. This ensures that
//       'oranges are compared to oranges and apples are compared
//       to apples' with respect to start time and end time duration
//       calculations.
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
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
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
//  timeCalcMode       TimeMathCalcMode
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
//  dateTimeFmtStr     string
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
//
//  ePrefix            string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method proceeds to successful completion, the returned
//       error instance is set to 'nil'. If an error is encountered, the
//       error object is populated with an appropriate error message.
//
func(durTUtil *durationTriadUtility) setStartEndTimes(
	durT *DurationTriad,
	startDateTime,
	endDateTime time.Time,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string,
	ePrefix string) error {

	durTUtil.lock.Lock()

	defer durTUtil.lock.Unlock()

	ePrefix += "durationTriadUtility.setStartEndTimes() "

	if durT == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "durT",
			inputParameterValue: "",
			errMsg:              "Input parameter 'durT' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	if startDateTime.IsZero() && endDateTime.IsZero() {
		return errors.New(ePrefix +
			"Error: Input parameters 'startDateTime' and 'endDateTime' are ZERO!")
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

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := DurationTriad{}

	err := tDurDtoUtil.setStartEndTimes(
		&tDur2.BaseTime,
		startDateTime,
		endDateTime,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "BaseTime Failed- ")

	if err != nil {
		return err
	}

	err = tDurDtoUtil.setStartEndTimes(
		&tDur2.LocalTime,
		startDateTime,
		endDateTime,
		tDurCalcType,
		TZones.Local(),
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix+ "LocalTime Failed- ")

	if err != nil {
		return err
	}

	err = tDurDtoUtil.setStartEndTimes(
		&tDur2.UTCTime,
		startDateTime,
		endDateTime,
		tDurCalcType,
		TZones.UTC(),
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix+ "UTC Time Failed - ")

	if err != nil {
		return err
	}

	durTUtil2 := durationTriadUtility{}

	err = durTUtil2.isValid(&tDur2, ePrefix + "'tDur2' Validity Check ")

	if err != nil {
		return err
	}

	durTUtil2.empty(durT, ePrefix)

	durTUtil2.copyIn(durT, &tDur2, ePrefix)

	return nil
}

// setStartTimeDuration - Receives a starting date time and calculates
// a time duration. The method then calculates the ending date time, duration
// and populates the DurationTriad data fields.
//
// The method will except negative time durations. A negative duration means that
// starting date time will be reclassified as ending date time with time duration
// being subtracted from that ending date time to compute staring date time.
//
// The user is also required to submit input parameters to time zone location and
// date time calculation type.
//
// __________________________________________________________________________
//
// Input Parameters:
//
//  durT       *DurationTriad
//     - The results of this calculation will be stored
//       in this instance of 'DurationTriad'
//
//  startDateTime   time.Time
//     - Starting Date Time for duration calculation
//
//  duration    time.Duration
//     - Time Duration added to 'startDateTime' in order to
//       compute Ending Date Time
//
//  tDurCalcType TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation   string
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
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
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
//  dateTimeFmtStr     string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  ePrefix            string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
//
func(durTUtil *durationTriadUtility) setStartTimeDuration(
	durT *DurationTriad,
	startDateTime time.Time,
	duration time.Duration,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr string,
	ePrefix string) error {

	durTUtil.lock.Lock()

	defer durTUtil.lock.Unlock()

	ePrefix += "durationTriadUtility.setStartTimeDuration() "

	if durT == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "durT",
			inputParameterValue: "",
			errMsg:              "Input parameter 'durT' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	if startDateTime.IsZero() && duration == 0 {
		return errors.New(ePrefix +
			"\nError: Both 'startDateTime' and 'duration' are Zero!\n")
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

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := DurationTriad{}

	err := tDurDtoUtil.setStartTimeDuration(
		&tDur2.BaseTime,
		startDateTime,
		duration,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "Base Time - ")

	if err != nil {
		return err
	}


	err = tDurDtoUtil.setStartTimeDuration(
		&tDur2.LocalTime,
		startDateTime,
		duration,
		tDurCalcType,
		TZones.Local(),
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "Local Time - ")

	if err != nil {
		return err
	}

	err = tDurDtoUtil.setStartTimeDuration(
		&tDur2.UTCTime,
		startDateTime,
		duration,
		tDurCalcType,
		TZones.UTC(),
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "UTC Time - ")

	if err != nil {
		return err
	}

	durTUtil2 := durationTriadUtility{}

	err = durTUtil2.isValid(&tDur2, ePrefix + "'tDur2' Validity Check ")

	if err != nil {
		return err
	}

	durTUtil2.empty(durT, ePrefix)

	durTUtil2.copyIn(durT, &tDur2, ePrefix)

	return nil
}


// setStartTimePlusTimeDto - Calculates time duration values based on a Starting Date Time
// plus time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'plusTimeDto' parameter. The 'plusTimeDto' parameter is added to 'startDateTime' to
// calculate ending date time and duration.
//
// Values in the 'plusTimeDto' parameter are automatically converted to positive numeric
// values before being added to parameter 'startDateTime'.
//
// True values for starting date time, ending date time and time duration are then stored in
// the DurationTriad data structure.
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
//       MikeAustin71\datetimeopsgo\datetime\timedurationdto.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  durT       *DurationTriad
//     - The results of this calculation will be stored
//       in this instance of 'DurationTriad'
//
//  startDateTime   time.Time
//     - Starting date time. Input parameter 'plusTimeDto'
//       will be added to this starting date time in order
//       to generate ending date time.
//
//  plusTimeDto       TimeDto
//     - Provides time values which will be added to
//       'startDateTime' in order to calculate duration.
//
//       type TimeDto struct {
//         Years                  int // Number of Years
//         Months                 int // Number of Months
//         Weeks                  int // Number of Weeks
//         WeekDays               int // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
//         DateDays               int // Total Number of Days. Weeks x 7 plus WeekDays
//         Hours                  int // Number of Hours.
//         Minutes                int // Number of Minutes
//         Seconds                int // Number of Seconds
//         Milliseconds           int // Number of Milliseconds
//         Microseconds           int // Number of Microseconds
//         Nanoseconds            int // Remaining Nanoseconds after Milliseconds & Microseconds
//         TotSubSecNanoseconds   int // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
//                                    // plus remaining Nanoseconds
//        }
//
//        Type 'TimeDto' is located in source file:
//          MikeAustin71\datetimeopsgo\datetime\timedto.go
//
//
//  tDurCalcType TDurCalcType
//     - Specifies the calculation type to be used in allocating
//       time duration:
//
//       TDurCalcType(0).StdYearMth()
//         - Default - standard year, month week, day time calculation.
//
//       TDurCalcType(0).CumMonths()
//         - Computes cumulative months - no Years.
//
//       TDurCalcType(0).CumWeeks()
//         - Computes cumulative weeks. No Years or months
//
//       TDurCalcType(0).CumDays()
//         - Computes cumulative days. No Years, months or weeks.
//
//       TDurCalcType(0).CumHours()
//         - Computes cumulative hours. No Years, months, weeks or days.
//
//       TDurCalcType(0).CumMinutes()
//         - Computes cumulative minutes. No Years, months, weeks, days
//           or hours.
//
//       TDurCalcType(0).CumSeconds()
//         - Computes cumulative seconds. No Years, months, weeks, days,
//           hours or minutes.
//
//       TDurCalcType(0).GregorianYears()
//         - Computes Years based on average length of a Gregorian Year
//           Used for very large duration values.
//
//           Type 'TDurCalcType' is located in source file:
//              MikeAustin71\datetimeopsgo\datetime\timedurationcalctypeenum.go
//
//
//  timeZoneLocation   string
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
//             In addition to military operations, Military
//             time zones are commonly used in aviation as
//             well as at sea. They are also known as nautical
//             or maritime time zones.
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
//  dateTimeFmtStr string
//     - A date time format string which will be used
//       to format and display 'dateTime'. Example:
//       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//       Date time format constants are found in the source
//       file 'constantsdatetime.go'. These constants represent
//       the more commonly used date time string formats. All
//       Date Time format constants begin with the prefix
//       'FmtDateTime'.
//
//       If 'dateTimeFmtStr' is submitted as an
//       'empty string', a default date time format
//       string will be applied. The default date time
//       format string is:
//         FmtDateTimeYrMDayFmtStr =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  ePrefix            string
//     - The error prefix containing the names of all
//       the methods executed up to this point.
//
// __________________________________________________________________________
//
// Return Value:
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error Type which encapsulates an
//       appropriate error message.
//
func(durTUtil *durationTriadUtility) setStartTimePlusTimeDto(
	durT *DurationTriad,
	startDateTime time.Time,
	plusTimeDto TimeDto,
	tDurCalcType TDurCalcType,
	timeZoneLocation string,
	timeCalcMode TimeMathCalcMode,
	dateTimeFmtStr,
	ePrefix string) error {

	durTUtil.lock.Lock()

	defer durTUtil.lock.Unlock()

	ePrefix += "durationTriadUtility.setStartTimePlusTimeDto() "

	if durT == nil {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "durT",
			inputParameterValue: "",
			errMsg:              "Input parameter 'durT' is a 'nil' pointer!",
			err:                 nil,
		}
	}

	if durT.lock == nil {
		durT.lock = new(sync.Mutex)
	}

	err := plusTimeDto.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Input Parameter 'plusTimeDto' is INVALID!\n" +
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
			errMsg:  "Input Parameter 'tDurCalcType' is INVALID!",
			err:                 nil,
		}
	}

	dtMech := DTimeMechanics{}

	dateTimeFmtStr = dtMech.PreProcessDateFormatStr(dateTimeFmtStr)

	tzMech := TimeZoneMechanics{}

	timeZoneLocation = tzMech.PreProcessTimeZoneLocation(timeZoneLocation)

	tDurDtoUtil := timeDurationDtoUtility{}

	tDur2 := DurationTriad{}

	err = tDurDtoUtil.setStartTimePlusTimeDto(
		&tDur2.BaseTime,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		timeZoneLocation,
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "Base Time - ")

	if err != nil {
		return err
	}

	err = tDurDtoUtil.setStartTimePlusTimeDto(
		&tDur2.LocalTime,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		TZones.Local(),
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "Local Time - ")

	if err != nil {
		return err
	}

	err = tDurDtoUtil.setStartTimePlusTimeDto(
		&tDur2.UTCTime,
		startDateTime,
		plusTimeDto,
		tDurCalcType,
		TZones.UTC(),
		timeCalcMode,
		dateTimeFmtStr,
		ePrefix + "UTC Time - ")

	if err != nil {
		return err
	}

	durTUtil2 := durationTriadUtility{}

	err = durTUtil2.isValid(&tDur2, ePrefix + "'tDur2' Validity Check ")

	if err != nil {
		return err
	}

	durTUtil2.empty(durT, ePrefix)

	durTUtil2.copyIn(durT, &tDur2, ePrefix)

	return nil
}
