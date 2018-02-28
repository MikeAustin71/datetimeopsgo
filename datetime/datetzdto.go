package datetime

import (
	"time"
	"fmt"
	"errors"
	"strings"
)

// DateTzDto - Type
// ================
// Used to store and transfer date times.
// The descriptors contained is this structure are intended
// to define and identify a specific point in time.
//
// This Type is NOT used to define duration; that is, the
// difference or time span between two point in time. For
// these types of operations see:
// DurationTimeUtility/common/durationutil.go
//
// DateTzDto defines a specific point in time using
// a variety of descriptors including year, month, day
// hour, minute, second, millisecond, microsecond and
// and nanosecond. In addition this Type specifies a
// time.Time value as well as time zone location and
// time zone.
//
// If you are unfamiliar with the concept of a time
// zone location, consider the field TimeLoc and
// TimeLocName below:
//
// Time zone location must be designated as one of two values.
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
//
type DateTzDto struct {
	Description			string					// Unused, available for classification, labeling or description
	Time						TimeDto					// Associated Time Components
	DateTime 				time.Time				// DateTime value for this DateTzDto Type
	DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
	TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
	// 		associated with this date time.
}

// AddDate - Adds input parameters 'years, 'months' and 'days' to date time value of the
// current DateTzDto and returns the updated value in a new DateTzDto instance.
//
// Input Parameters
// ================
//
// years		int							- Number of years to add to the current date.
// months		int							- Number of months to add to the current date.
// days			int							- Number of days to add to the current date.
//
// Note: 	Date Component input parameters may be either negative
// 					or positive. Negative values will subtract time from
// 					the current DateTzDto instance.
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
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Description			string					// Unused, available for classification, labeling or description
//		Year       			int							// Year Number
//		Month      			int							// Month Number
//		Day        			int							// Day Number
//		Hour       			int							// Hour Number
//		Minute     			int							// Minute Number
//		Second     			int							// Second Number
//		Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//		Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//		Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																		// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//		TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//		TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//	}
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) AddDate(years, months, days int, dateTimeFormatStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.AddDate() "

	err := dtz.IsValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "The current DateTzDto is INVALID! dtz.DateTime='%v'", dtz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	newDt := dtz.DateTime.AddDate(years, months, days)

	dtz2, err := DateTzDto{}.New(newDt, dtz.DateTimeFmt)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(newDt, dtz.DateTimeFmt). newDt='%v'  Error='%v'", newDt.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}


	return dtz2, nil
}

// AddDateToThis - Adds input parameters 'years, 'months' and 'days' to date time value
// of the current DateTzDto. The updated DateTime is retained in the current
// DateTzDto instance.
//
// Input Parameters
// ================
//
// years		int							- Number of years to add to the current date.
// months		int							- Number of months to add to the current date.
// days			int							- Number of days to add to the current date.
//
// Note: 	Date Component input parameters may be either negative
// 					or positive. Negative values will subtract time from
// 					the current DateTzDto instance.
//
// Returns
// =======
//
//  There one return values: An Error type
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) AddDateToThis(years, months, days int) error {

	ePrefix := "DateTzDto.AddDate() "

	err := dtz.IsValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "The current DateTzDto is INVALID! dtz.DateTime='%v'", dtz.DateTime.Format(FmtDateTimeYrMDayFmtStr))
	}

	newDt := dtz.DateTime.AddDate(years, months, days)

	dtz2, err := DateTzDto{}.New(newDt, dtz.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(newDt, dtz.DateTimeFmt). newDt='%v'  Error='%v'", newDt.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.CopyIn(dtz2)

	return nil

}

// AddDuration - Adds Duration to the DateTime Value of the current
// DateTzDto and returns a new DateTzDto instance with the updated
// Date Time value.
//
// Input Parameter
// ===============
//
// duration time.Duration		- A Time duration value which is added to the DateTime
//														value of the current DateTzDto instance to produce and
//														return a new, updated DateTzDto instance.
//
// Note: 	The time.Duration input parameter may be either negative
// 					or positive. Negative values will subtract time from
// 					the current DateTzDto instance.
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
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type updated to reflect the added 'duration'
// 							input parameter. A DateTzDto structure is defined as follows:
//
//	type DateTzDto struct {
//		Description			string					// Unused, available for classification, labeling or description
//		Year       			int							// Year Number
//		Month      			int							// Month Number
//		Day        			int							// Day Number
//		Hour       			int							// Hour Number
//		Minute     			int							// Minute Number
//		Second     			int							// Second Number
//		Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//		Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//		Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																		// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//		TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//		TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//	}
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) AddDuration(duration time.Duration, dateTimeFmtStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.AddDuration() "

	newDateTime := dtz.DateTime.Add(duration)

	dtz2, err := DateTzDto{}.New(newDateTime, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(newDateTime, dateTimeFmtStr). newDateTime='%v'  Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return dtz2, nil

}

// AddDurationToThis - Receives a time.Duration input parameter and adds this
// duration value to the Date Time value of the current DateTzDto. The current
// DateTzDto Date Time values are updated to reflect the added 'duration'.
//
// Input Parameter
// ===============
//
// duration time.Duration		- A Time duration value which is added to the DateTime
//														value of the current DateTzDto instance to produce and
//														return a new, updated DateTzDto instance.
//
// Note: 	The time.Duration input parameter may be either negative
// 					or positive. Negative values will subtract time from
// 					the current DateTzDto instance.
//
// Returns
// =======
//						There is only one return parameter, an 'error' type.
//
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) AddDurationToThis(duration time.Duration) error {

	ePrefix := "DateTzDto.AddDuration() "

	newDateTime := dtz.DateTime.Add(duration)

	dtz2, err := DateTzDto{}.New(newDateTime, dtz.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(newDateTime, dtz.DateTimeFmt). newDateTime='%v'  Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.CopyIn(dtz2)

	return nil
}

// AddTime - Adds time components to the date time value of the current
// DateTzDto instance. The resulting updated date time value is returned
// to the calling function in the form of a new DateTzDto instance.
//
// Input Parameters
// ================
//
// hours				int	- Number of hours to add.
// minutes			int	- Number of minutes to add.
// seconds			int - Number of seconds to add.
// milliseconds	int	- Number of milliseconds to add.
// microseconds	int	- Number of microseconds to add.
// nanoseconds	int - Number of nanoseconds to add.
//
// Note: 	  Time Component input parameters may be either negative
// 					or positive. Negative values will subtract time from
// 					the current DateTzDto instance.
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
// Returns
// =======
// There are two returns	(1) A DateTzDto instance containing the
//															updated time values.
//
//												(2) An 'error' type.
//
// (1)  DateTzDto - If successful the method returns a valid, fully populated
//										DateTzDto type defined as follows:
//
//			type DateTzDto struct {
//				Description			string					// Unused, available for classification, labeling or description
//				Year       			int							// Year Number
//				Month      			int							// Month Number
//				Day        			int							// Day Number
//				Hour       			int							// Hour Number
//				Minute     			int							// Minute Number
//				Second     			int							// Second Number
//				Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//				Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//				Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																				// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//				TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//				DateTime 				time.Time				// DateTime value for this DateTzDto Type
//				DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//				TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//			}
//
// (2) error 			- 	If successful the returned error Type is set equal to 'nil'. If errors are
//											encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) AddTime(hours, minutes, seconds, milliseconds, microseconds,
nanoseconds int, dateTimeFormatStr string ) (DateTzDto, error) {

	ePrefix := "DateTzDto.AddTime() "

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	newDateTime := dtz.DateTime.Add(time.Duration(totNanoSecs))

	dtz2, err := DateTzDto{}.New(newDateTime, dtz.DateTimeFmt)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(newDateTime, dtz.DateTimeFmt) " +
				"newDateTime='%v'  Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz2.SetDateTimeFmt(dateTimeFormatStr)

	return dtz2, nil
}

// AddTimeToThis - Adds time components (hours, minutes, seconds etc.)
// to the current value of this DateTzDto instance.
//
// Input Parameters
// ================
//
// hours				int	- Number of hours to add.
// minutes			int	- Number of minutes to add.
// seconds			int - Number of seconds to add.
// milliseconds	int	- Number of milliseconds to add.
// microseconds	int	- Number of microseconds to add.
// nanoseconds	int - Number of nanoseconds to add.
//
// Note: 		Time Component input parameters may be either negative
// 					or positive. Negative values will subtract time from
// 					the current DateTzDto instance.
//
// Returns
// =======
// There is only one return; an 'error' type.
//
//
// (1) error 			- 	If successful the returned error Type is set equal to 'nil'.
// 											If errors are encountered this error Type will encapsulate
// 											an error message.
//
func (dtz *DateTzDto) AddTimeToThis(hours, minutes, seconds, milliseconds,
microseconds,	nanoseconds int ) error {

	ePrefix := "DateTzDto.AddTimeToThis() "

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	newDateTime := dtz.DateTime.Add(time.Duration(totNanoSecs))

	dtz2, err := DateTzDto{}.New(newDateTime, dtz.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(newDateTime, dtz.DateTimeFmt) " +
			"newDateTime='%v'  Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.CopyIn(dtz2)

	return nil
}

// AddDateTime - Adds date time components to the date time value of the
// current DateTzDto instance. The updated date time value is returned to
// the calling function as a new DateTzDto instance.
//
// Input Parameters
// ================
//
// years				int - Number of years to add.
// months				int - Number of months to add.
// days					int	- Number of days to add.
// hours				int	- Number of hours to add.
// minutes			int	- Number of minutes to add.
// seconds			int - Number of seconds to add.
// milliseconds	int	- Number of milliseconds to add.
// microseconds	int	- Number of microseconds to add.
// nanoseconds	int - Number of nanoseconds to add.
//
// Note: 	Date Time Component input parameters may be either negative
// 					or positive. Negative values will subtract time from
// 					the current DateTzDto instance.
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
// Returns
// =======
// There are two returns	(1) A DateTzDto instance containing the
//															updated date time values.
//
//												(2) An 'error' type.
//
// (1)  DateTzDto - If successful the method returns a valid, fully populated
//										DateTzDto type defined as follows:
//
//			type DateTzDto struct {
//				Description			string					// Unused, available for classification, labeling or description
//				Year       			int							// Year Number
//				Month      			int							// Month Number
//				Day        			int							// Day Number
//				Hour       			int							// Hour Number
//				Minute     			int							// Minute Number
//				Second     			int							// Second Number
//				Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//				Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//				Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																				// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//				TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//				DateTime 				time.Time				// DateTime value for this DateTzDto Type
//				DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//				TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//			}
//
// (2) error 			- 	If successful the returned error Type is set equal to 'nil'. If errors are
//											encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) AddDateTime(years, months, days, hours, minutes, seconds,
milliseconds, microseconds, nanoseconds int,
	dateTimeFormatStr string) (DateTzDto, error ) {

	ePrefix := "DateTzDto.AddDateTime() "

	newDate := dtz.DateTime.AddDate(years, months, days)

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	newDateTime := newDate.Add(time.Duration(totNanoSecs))

	dtz2, err := DateTzDto{}.New(newDateTime, dateTimeFormatStr)

	if err != nil {
		return DateTzDto{},
			fmt.Errorf(ePrefix + "Error returned from DateTzDto{}.New(newDateTime, dateTimeFormatStr) " +
				"newDateTime='%v' Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return dtz2, nil
}

// AddDateTimeToThis - Adds date time components to the date time value of the current
// DateTzDto instance.
//
// Input Parameters
// ================
//
// years				int - Number of years to add.
// months				int - Number of months to add.
// days					int	- Number of days to add.
// hours				int	- Number of hours to add.
// minutes			int	- Number of minutes to add.
// seconds			int - Number of seconds to add.
// milliseconds	int	- Number of milliseconds to add.
// microseconds	int	- Number of microseconds to add.
// nanoseconds	int - Number of nanoseconds to add.
//
// Note: 	Date Time Component input parameters may be either negative
// 					or positive. Negative values will subtract time from
// 					the current DateTzDto instance.
//
// Returns
// =======
// There is only one return; an 'error' type.
//
//  error 	- 	If successful the returned error Type is set equal to 'nil'. If errors are
//							encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) AddDateTimeToThis(years, months, days, hours, minutes, seconds,
milliseconds, microseconds, nanoseconds int,
	dateTimeFormatStr string) error {
	ePrefix := "DateTzDto.AddDateTime() "

	newDate := dtz.DateTime.AddDate(years, months, days)

	totNanoSecs := int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	newDateTime := newDate.Add(time.Duration(totNanoSecs))

	dtz2, err := DateTzDto{}.New(newDateTime, dateTimeFormatStr)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned from DateTzDto{}.New(newDateTime, dateTimeFormatStr) " +
			"newDateTime='%v' Error='%v'", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dtz.CopyIn(dtz2)

	return nil
}

// CopyIn - Receives an incoming DateTzDto and
// copies those data fields to the current DateTzDto
// instance.
//
// When completed, the current DateTzDto will be
// equal in all respects to the incoming DateTaDto
// instance.
//
// Input Parameter
// ===============
// dtz2		*DateTzDto - A pointer to a DateTzDto instance.
//											This data will be copied into the
//											data fields of the current DateTzDto
//											instance.
//
//		A DateTzDto struct is defined as follows:
//
//		type DateTzDto struct {
//			Description			string					// Unused, available for classification, labeling or description
//			Year       			int							// Year Number
//			Month      			int							// Month Number
//			Day        			int							// Day Number
//			Hour       			int							// Hour Number
//			Minute     			int							// Minute Number
//			Second     			int							// Second Number
//			Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//			Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//			Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//			TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//			DateTime 				time.Time				// DateTime value for this DateTzDto Type
//			DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//			TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
//																			// 		associated with this date time.
//		}
//
// Returns
// =======
//
// None
//
func (dtz *DateTzDto) CopyIn(dtz2 DateTzDto) {
	dtz.Empty()

	dtz.Description 		= dtz2.Description
	dtz.Time 					  = dtz2.Time.CopyOut()
	dtz.DateTimeFmt			= dtz2.DateTimeFmt

	if !dtz2.DateTime.IsZero() {
		dtz.DateTime = dtz2.DateTime
		dtz.TimeZone = dtz2.TimeZone.CopyOut()
	} else {
		dtz.TimeZone				= TimeZoneDefDto{}
		dtz.DateTime				= time.Time{}
	}

}

// CopyOut - returns a DateTzDto  instance
// which represents a deep copy of the current
// DateTzDto object.
func (dtz *DateTzDto) CopyOut() DateTzDto {
	dtz2 := DateTzDto{}

	dtz2.Description 		= dtz.Description
	dtz2.Time 					= dtz.Time.CopyOut()
	dtz2.DateTimeFmt		= dtz.DateTimeFmt

	if !dtz.DateTime.IsZero() {
		dtz2.DateTime = dtz.DateTime
		dtz2.TimeZone = dtz.TimeZone.CopyOut()
	} else {
		dtz2.TimeZone				= TimeZoneDefDto{}
		dtz2.DateTime				= time.Time{}
	}

	return dtz2
}

// Empty - sets all values of the current DateTzDto
// instance to their uninitialized or zero state.
func (dtz *DateTzDto) Empty() {

	dtz.Description		 	= ""
	dtz.Time.Empty()
	dtz.TimeZone				= TimeZoneDefDto{}
	dtz.DateTime				= time.Time{}
	dtz.DateTimeFmt			= ""

	return
}

// Equal - Returns true if input DateTzDto is equal
// in all respects to the current DateTzDto instance.
func (dtz *DateTzDto) Equal(dtz2 DateTzDto) bool {

	if 	dtz.Description != dtz2.Description		||
		!dtz.Time.Equal(dtz2.Time) 		 					||
		!dtz.DateTime.Equal(dtz2.DateTime)			||
		dtz.DateTimeFmt != dtz2.DateTimeFmt 		||
		!dtz.TimeZone.Equal(dtz2.TimeZone)	{

		return false
	}

	return true
}

// IsEmpty - Analyzes the current DateTzDto instance to determine
// if the instance is in an 'EMPTY' or uninitialized state.
//
// If the current DateTzDto instance is found to be 'EMPTY', this
// method returns 'true'. Otherwise, if the instance is 'NOT EMPTY',
// this method returns false!
func (dtz *DateTzDto) IsEmpty () bool {

	if dtz.Description 			== "" 	&&
		dtz.Time.IsZero() 						&&
		dtz.DateTime.IsZero() 				&&
		dtz.DateTimeFmt 			== "" 	&&
		dtz.TimeZone.IsEmpty() {

		return true

	}

	return false
}

// IsValid - Analyzes the current DateTzDto instance and returns
// an error if it is found to be INVALID.
//
// If the current DateTzDto instance is VALID, this method returns
// nil.
func (dtz *DateTzDto) IsValid() error {

	ePrefix := "DateTzDto.IsValid() "

	if dtz.IsEmpty() {
		return errors.New(ePrefix+"Error: This DateTzDto instance is EMPTY!")
	}

	if dtz.DateTime.IsZero() {
		return errors.New(ePrefix + "Error: DateTzDto.DateTime is ZERO!")
	}

	if dtz.TimeZone.IsEmpty() {
		return errors.New(ePrefix + "Error: dtz.TimeZone is EMPTY!")
	}

	if err := dtz.Time.IsValid(); err != nil {
		return fmt.Errorf(ePrefix + "Error: dtz.Time is INVALID. Error='%v'", err.Error())
	}

	if !dtz.TimeZone.IsValidFromDateTime(dtz.DateTime) {
		return errors.New(ePrefix + "Error: dtz.TimeZone is INVALID!")
	}

	dtz2, err := DateTzDto{}.New(dtz.DateTime, dtz.DateTimeFmt)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error creating check DateTzDto - Error='%v'", err.Error())
	}

	if !dtz.Equal(dtz2) {
		return errors.New(ePrefix + "Error: Current DateTzDto is NOT EQUAL to Check DateTzDto!")
	}

	return nil
}

// New - returns a new DateTzDto instance based on a time.Time ('dateTime')
// input parameter.
//
// Input Parameter
// ===============
//
// dateTime   		time.Time - A date time value
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
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Description			string					// Unused, available for classification, labeling or description
//		Time       			TimeDto					// Time Components
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//		TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//	}
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
// Usage
// =====
//
// Example:
//			fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
//			dtzDto, err := DateTzDto{}.New(dateTime, fmtString)
//
func (dtz DateTzDto) New(dateTime time.Time, dateTimeFmtStr string)(DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	if dateTime.IsZero() {
		return DateTzDto{}, errors.New(ePrefix + "Error: Input parameter dateTime is Zero value!")
	}

	dtz2 := DateTzDto{}

	err := dtz2.SetFromTime(dateTime, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned from dtz2.SetFromTime(dateTime). dateTime='%v'  Error='%v'", dateTime, err.Error())
	}

	return dtz2, nil
}

// NewDateTimeElements - creates a new DateTzDto object and populates the data fields based on
// input parameters.
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// minute						int			- minute number	0 - 59
// second						int			- second number	0	-	59
// nanosecond				int			- nanosecond number 0 - 999999999
//
// timeZoneLocation	string	- time zone location must be designated as one of two values.
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
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Time      			TimeDto					// Time Components
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//		TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//	}
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
// Usage
// =====
//
// Example:
//			fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
//	dtzDto, err := DateTzDto{}.NewDateTimeElements(year, month, day, hour, minute, second, nanosecond ,
// 										timeZoneLocation, fmtStr)
//
//
func (dtz DateTzDto) NewDateTimeElements(year, month, day, hour, minute, second, nanosecond int,
	timeZoneLocation, dateTimeFmtStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	dtz2 := DateTzDto{}

	err := dtz2.SetFromDateTimeElements(year, month, day, hour, minute, second, nanosecond, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned from dtz2.SetFromDateTimeElements(...) " +
			"year='%v' month='%v' day='%v' hour='%v' minute='%v' second='%v' nanosecond='%v' timeZoneLocatin='%v'  Error='%v'",
			year, month, day, hour, minute, second, nanosecond, timeZoneLocation, err.Error())
	}

	return dtz2, nil
}

// NewDateTime - creates a new DateTzDto object and populates the data fields based on
// input parameters.
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// minute						int			- minute number	0 - 59
// second						int			- second number	0	-	59
// millisecond			int			- millisecond number 0 - 999
// microsecond			int			-	microsecond number 0 - 999
// nanosecond				int			- nanosecond number 0 - 999
// timeZoneLocation	string	- time zone location must be designated as one of two values.
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
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Time       			TimeDto					// Year Number
//		Month      			int							// Month Number
//		Day        			int							// Day Number
//		Hour       			int							// Hour Number
//		Minute     			int							// Minute Number
//		Second     			int							// Second Number
//		Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//		Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//		Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																		// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//		TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//		TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//	}
//
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
// Usage
// =====
//
// Example:
//			fmtStr := "2006-01-02 15:04:05.000000000 -0700 MST"
//			dtzDto, err := DateTzDto{}.New(year, month, day, hour, min, sec, nanosecond , timeZoneLocation, fmtStr)
//
//
func (dtz DateTzDto) NewDateTime(year, month, day, hour, minute, second, millisecond, microsecond,
nanosecond int, timeZoneLocation, dateTimeFmtStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	dtz2 := DateTzDto{}

	err := dtz2.SetFromDateTime(year, month, day, hour, minute, second,
		millisecond, microsecond, nanosecond, timeZoneLocation, dateTimeFmtStr)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by dtz2.SetFromDateTime(...) " +
			"year='%v', month='%v', day='%v', hour='%v', minute='%v', second='%v', millisecond='%v', microsecond='%v' nanosecond='%v', timeZoneLocation='%v' Error='%v'",
			year, month, day, hour, minute, second, millisecond, microsecond, nanosecond, timeZoneLocation, err.Error())
	}

	return dtz2, nil
}

// NewTimeDto - Receives input parameters type TimeDto, 'timeZoneLocation' and 'dateTimeFormatStr'.
// These parameters are used to construct and return a new DateTzDto instance.
func (dtz DateTzDto) NewTimeDto(tDto TimeDto, timeZoneLocation string, dateTimeFormatStr string) (DateTzDto, error) {

	ePrefix := "DateTzDto.NewTimeDto() "

	dtz2 := DateTzDto{}

	err := dtz2.SetFromTimeDto(tDto, timeZoneLocation)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by dtz2.SetFromTimeDto(tDto, timeZoneLocation). Error='%v'", err.Error())
	}

	dtz2.SetDateTimeFmt(dateTimeFormatStr)

	return dtz2, nil
}


// SetDateTimeFmt - Sets the DateTzDto data field 'DateTimeFmt'.
// This string is used to format the DateTzDto DateTime field
// when DateTzDto.String() is called.
func (dtz *DateTzDto) SetDateTimeFmt(dateTimeFmtStr string) {

	dtz.DateTimeFmt = dateTimeFmtStr

}

// SetFromTime - Sets the values of the current DateTzDto fields
// based on an input parameter 'dateTime' (time.time).
//
// Input Parameters
// ================
//
// dateTime   		time.Time - A date time value
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
// Returns
// =======
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromTime(dateTime time.Time, dateTimeFmtStr string) error {

	ePrefix := "DateTzDto.SetFromTime() "

	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter dateTime is Zero value!")
	}

	dtz.Empty()

	var err error

	dtz.Time, err  = TimeDto{}.NewFromDateTime(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeDto{}.NewFromDateTime(dateTime). Error='%v'", err.Error())
	}

	dtz.DateTime = dateTime

	if len(dateTimeFmtStr) == 0 {
		dateTimeFmtStr = FmtDateTimeYrMDayFmtStr
	}

	dtz.DateTimeFmt = dateTimeFmtStr


	dtz.TimeZone, err = TimeZoneDefDto{}.New(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeZoneDefDto{}.New(dateTime). dateTime='%v'  Error='%v'", dateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return nil
}

// SetFromDateTimeElements - sets the values of the current DateTzDto
// data fields based on input parameters of date time components and
// a time zone location.
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// minute						int			- minute number	0 - 59
// second						int			- second number	0	-	59
// nanosecond				int			- nanosecond number 0 - 999999999
//														This represents the total number of
//														nanoseconds which is less than one second.
//
// timeZoneLocation	string	- time zone location must be designated as one of two values.
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
// Returns
// =======
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromDateTimeElements(year, month, day, hour, minute, second,
nanosecond int, timeZoneLocation, dateTimeFmtStr string) (error) {

	ePrefix := "DateTzDto.SetFromDateTimeElements() "


	if len(timeZoneLocation) == 0 {
		return errors.New(ePrefix + "Error: Input parameter 'timeZoneLocation' is an EMPTY STRING! 'timeZoneLocation' is required!")
	}

	dtz.Empty()

	var err error

	dtz.Time, err = TimeDto{}.New(year, month, 0, day, hour, minute, second, 0, 0, nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeDto{}.New(year, month, ...). Error='%v'", err.Error())
	}

	dtz.DateTime, err	= dtz.Time.GetDateTime(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by dtz.Time.GetDateTime(timeZoneLocation). timeZoneLocation='%v'  Error='%v'", timeZoneLocation, err.Error())
	}

	dtz.TimeZone, err = TimeZoneDefDto{}.New(dtz.DateTime)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeZoneDefDto{}.New(dtz.DateTime). dtz.DateTime='%v'  Error='%v'", dtz.DateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	if len(dateTimeFmtStr) == 0 {
		dateTimeFmtStr = FmtDateTimeYrMDayFmtStr
	}

	dtz.DateTimeFmt = dateTimeFmtStr

	return nil

}

// SetFromDateTime - Sets the values of the Date Time fields
// for the current DateTzDto instance based on time components
// and a Time Zone Location.
//
// Note that this variation of time elements breaks time down by
// hour, minute, second, millisecond, microsecond and nanosecond.
//
// See method SetFromDateTimeElements(), above, which uses a slightly
// different set of time components.
//
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// min							int			- minute number	0 - 59
// sec							int			- second number	0	-	59
// millisecond			int			- millisecond number 0 - 999
// microsecond			int			-	microsecond number 0 - 999
// nanosecond				int			- nanosecond number 0 - 999
// timeZoneLocation	string	- time zone location must be designated as one of two values.
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
// Returns
// =======
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromDateTime(year, month, day, hour, minute, second,
millisecond, microsecond, nanosecond int, timeZoneLocation, dateTimeFmtStr string) error {

	ePrefix := "DateTzDto.SetFromDateTime() "

	var err error


	dtz.Empty()

	dtz.Time, err = TimeDto{}.New(year, month,0, day, hour, minute,
															second, millisecond, microsecond, nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeDto{}.New(year, month,...).  Error='%v'", err.Error())
	}

	dtz.DateTime, err = dtz.Time.GetDateTime(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by dtz.Time.GetDateTime(timeZoneLocation). Error='%v'", err.Error())
	}

	dtz.TimeZone, err = TimeZoneDefDto{}.New(dtz.DateTime)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDefDto{}.New(dtz.DateTime). dtz.DateTime='%v'  Error=%v", dtz.DateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	if len(dateTimeFmtStr) == 0 {
		dateTimeFmtStr = FmtDateTimeYrMDayFmtStr
	}

	dtz.DateTimeFmt = dateTimeFmtStr

	return nil
}

// SetFromTimeDto - Receives data from a TimeDto input parameter
// and sets all data fields of the current DateTzDto instance
// accordingly. When the method completes, the values of the
// current DateTzDto will equal the values of the input parameter
// TimeDto instance.
//
// Input Parameters
// ================
// tDto		TimeDto						- A populated TimeDto instance
//
//														type TimeDto struct {
//															Years        		int64
//															Months       		int64
//															Weeks        		int64
//															WeekDays     		int64
//															DateDays		 		int64
//															Hours        		int64
//															Minutes      		int64
//															Seconds      		int64
//															Milliseconds 		int64
//															Microseconds 		int64
//															Nanoseconds  		int64
//															TotNanoseconds	int64
//														}
//
// timeZoneLocation	string	- time zone location must be designated as one of two values.
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
//
func (dtz *DateTzDto) SetFromTimeDto(tDto TimeDto, timeZoneLocation string) error {

	ePrefix := "DateTzDto.SetFromTimeDto() "

	if tDto.IsZero() {

		return fmt.Errorf(ePrefix + "Error: All input parameter date time elements equal ZERO!")
	}

	t2Dto := tDto.CopyOut()
	
	t2Dto.ConvertToAbsoluteValues()

	if err := t2Dto.IsValid(); err != nil {
		return fmt.Errorf(ePrefix + "Error: Input Parameter tDto (TimeDto) is INVALID. Error='%v'", err.Error())
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		timeZoneLocation = "Local"
	}

	loc, err := time.LoadLocation(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by time.LoadLocation(timeZoneLocation). timeZoneLocation='%v'  Error='%v' ", timeZoneLocation, err.Error())
	}

	dateTime := time.Date(int(t2Dto.Years), time.Month(int(t2Dto.Months)), int(t2Dto.DateDays), int(t2Dto.Hours), int(t2Dto.Minutes), int(t2Dto.Seconds), int(t2Dto.TotNanoseconds), loc)

	timeZoneDef, err := TimeZoneDefDto{}.New(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDefDto{}.New(dateTime). dateTime='%v' Error='%v'", dateTime, err.Error())
	}

	dtz.Empty()
	dtz.DateTime 		= dateTime
	dtz.TimeZone 		= timeZoneDef.CopyOut()
	dtz.Time				= tDto.CopyOut()

	return nil
}

// String - This method returns the DateTzDto
// DateTime field value formatted as a string.
// If the DateTzDto field DateTimeFmt is an
// empty string, a default format string will
// be used. The default format is:
//
// FmtDateTimeYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (dtz *DateTzDto) String() string {

	fmtStr := dtz.DateTimeFmt

	if len(fmtStr) == 0 {
		fmtStr = FmtDateTimeYrMDayFmtStr
	}

	return dtz.DateTime.Format(fmtStr)
}

// Sub - Subtracts the DateTime value of the incoming DateTzDto
// from the DateTime value of the current DateTzDto and returns
// the duration.
func (dtz *DateTzDto) Sub(dtz2 DateTzDto) time.Duration {

	return dtz.DateTime.Sub(dtz2.DateTime)

}

// SubDateTime - Subtracts a date time value from the date time
// value of the current DateTzDto. The result is returned as
// a time.Duration.
func (dtz *DateTzDto) SubDateTime(t2 time.Time) time.Duration {
	return dtz.DateTime.Sub(t2)
}

