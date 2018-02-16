package common

import (
	"errors"
	"fmt"
	"time"
)

/*
	Overview - Location
	===================

  timezoneutility.go is part of the date time operations library. The source code repository
 	for this file is located at:

					https://github.com/MikeAustin71/datetimeopsgo.git



	Dependencies
	============

	None

 */

// NOTE: See https://golang.org/pkg/time/#LoadLocation
// and https://www.iana.org/time-zones to ensure that
// the IANA Time Zone Database is properly configured
// on your system. Note: IANA Time Zone Data base is
// equivalent to 'tz database'.
const (
	// TzUsEast - USA Eastern Time Zone
	// IANA database identifier
	TzUsEast = "America/New_York"

	// TzUsCentral - USA Central Time Zone
	// IANA database identifier
	TzUsCentral = "America/Chicago"

	// TzUsMountain - USA Mountain Time Zone
	// IANA database identifier
	TzUsMountain = "America/Denver"

	// TzUsPacific - USA Pacific Time Zone
	// IANA database identifier
	TzUsPacific = "America/Los_Angeles"

	// TzUsHawaii - USA Hawaiian Time Zone
	// IANA database identifier
	TzUsHawaii = "Pacific/Honolulu"

	// tzUTC - UTC Time Zone IANA database
	// identifier
	TzUTC = "Zulu"

	NeutralDateFmt = "2006-01-02 15:04:05.000000000"
)

// DateTzDto - `Used to store and transfer
// date times.
type DateTzDto struct {
	Year int
	Month int
	Day int
	Hour int
	Minute int
	Second int
	Nanosecond int
	IANATimeZone string
}

// TimeZoneDto - Time Zone Data and Methods
type TimeZoneUtility struct {
	Description string
	TimeIn      time.Time
	TimeInLoc   *time.Location
	TimeOut     time.Time
	TimeOutLoc  *time.Location
	TimeUTC     time.Time
	TimeLocal		time.Time
}

// AddDate - Adds specified years, months and days values to the
// current time values maintained by this TimeZoneDto
func (tzu *TimeZoneUtility) AddDate(years, months, days int) error {

	ePrefix := "TimeZoneDto.AddDate() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: This Time Zone Utility is INVALID!  Error='%v'", err.Error())
	}

	tzu.TimeIn = tzu.TimeIn.AddDate(years, months, days)
	tzu.TimeOut = tzu.TimeOut.AddDate(years, months, days)
	tzu.TimeUTC = tzu.TimeUTC.AddDate(years, months, days)
	tzu.TimeLocal = tzu.TimeLocal.AddDate(years, months, days)
	tzu.TimeInLoc = tzu.TimeIn.Location()
	tzu.TimeOutLoc = tzu.TimeOut.Location()

	return nil
}


// AddDuration - Adds 'duration' to the time values maintained by the
// current TimeZoneDto
func (tzu *TimeZoneUtility) AddDuration(duration time.Duration) error {

	ePrefix := "TimeZoneDto.AddDuration() "

	if duration == 0 {
		return nil
	}

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "This TimeZoneDto instance is INVALID! Error='%v'", err.Error())
	}

	tzu.TimeIn = tzu.TimeIn.Add(duration)
	tzu.TimeInLoc = tzu.TimeIn.Location()
	tzu.TimeOut = tzu.TimeOut.Add(duration)
	tzu.TimeOutLoc = tzu.TimeOut.Location()
	tzu.TimeUTC = tzu.TimeUTC.Add(duration)
	tzu.TimeLocal = tzu.TimeLocal.Add(duration)

	return nil
}


// ConvertTz - Convert Time from existing time zone to targetTZone.
// The results are stored in the TimeZoneDto data structure.
// The input time and output time are equivalent times adjusted
// for different time zones.
//
// Input Parameters:
// tIn - time.Time initial time
// targetTz - The IANA Time Zone or Time Zone 'Local' to which
// input parameter 'tIn' will be converted.

// Output Values are returned in the tzu (TimeZoneDto)
// data fields. tzu.TimeOut contains the correct time in the 'target' time
// zone.
func (tzu TimeZoneUtility) ConvertTz(tIn time.Time, targetTz string) (TimeZoneUtility, error) {

	ePrefix := "TimeZoneDto.ConvertTz() "

	tzuOut := TimeZoneUtility{}

	if isValidTz, _, _ := tzu.IsValidTimeZone(targetTz); !isValidTz {
		return tzuOut, errors.New(fmt.Sprintf("%v Error: targetTz is INVALID!! Input Time Zone == %v", ePrefix, targetTz))
	}

	if tIn.IsZero() {
		return tzuOut, errors.New(ePrefix + "Error: Input parameter time, 'tIn' is zero and INVALID")
	}

	tzOut, err := time.LoadLocation(targetTz)

	if err != nil {
		return tzuOut, fmt.Errorf("%vError Loading Target IANA Time Zone 'targetTz', %v. Errors: %v ",ePrefix, targetTz, err.Error())
	}


	tzuOut.SetTimeIn(tIn)

	tzuOut.SetTimeOut(tIn.In(tzOut))

	tzuOut.SetUTCTime(tIn)

	err = tzuOut.SetLocalTime(tIn)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.SetLocalTime(tIn). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// CopyOut - Creates and returns a deep copy of the
// current TimeZoneDto instance.
func (tzu *TimeZoneUtility) CopyOut() TimeZoneUtility {
	tzu2 := TimeZoneUtility{}
	tzu2.Description = tzu.Description
	tzu2.TimeIn = tzu.TimeIn
	tzu2.TimeInLoc = tzu.TimeInLoc
	tzu2.TimeOut = tzu.TimeOut
	tzu2.TimeOutLoc = tzu.TimeOutLoc
	tzu2.TimeUTC = tzu.TimeUTC
	tzu2.TimeLocal = tzu.TimeLocal

	return tzu2
}

// CopyToThis - Copies another TimeZoneDto
// to the current TimeZoneDto data fields.
func (tzu *TimeZoneUtility) CopyToThis(tzu2 TimeZoneUtility) {
	tzu.Empty()

	tzu.Description = tzu2.Description
	tzu.TimeIn = tzu2.TimeIn
	tzu.TimeInLoc = tzu2.TimeInLoc
	tzu.TimeOut = tzu2.TimeOut
	tzu.TimeOutLoc = tzu2.TimeOutLoc
	tzu.TimeUTC = tzu2.TimeUTC
	tzu.TimeLocal = tzu2.TimeLocal
}

// Equal - returns a boolean value indicating
// whether two TimeZoneDto data structures
// are equivalent.
func (tzu *TimeZoneUtility) Equal(tzu2 TimeZoneUtility) bool {
	if tzu.TimeIn != tzu2.TimeIn ||
		tzu.TimeInLoc != tzu2.TimeInLoc ||
		tzu.TimeOut != tzu2.TimeOut ||
		tzu.TimeOutLoc != tzu2.TimeOutLoc ||
		tzu.TimeUTC != tzu2.TimeUTC  ||
		tzu.TimeLocal != tzu2.TimeLocal	 {

		return false
	}

	return true
}

// Empty - Clears or returns this
// TimeZoneDto to an uninitialized
// state.
func (tzu *TimeZoneUtility) Empty() {
	tzu.Description = ""
	tzu.TimeIn = time.Time{}
	tzu.TimeInLoc = nil
	tzu.TimeOut = time.Time{}
	tzu.TimeOutLoc = nil
	tzu.TimeUTC = time.Time{}
	tzu.TimeLocal = time.Time{}
}

// GetLocationIn - Returns the time zone location for the
// TimeInLoc data field which is part of the current TimeZoneDto
// structure.
func (tzu *TimeZoneUtility) GetLocationIn() (string, error) {
	ePrefix := "TimeZoneDto.GetLocationIn() "

	if tzu.TimeIn.IsZero() {
		return "", errors.New(ePrefix + "Error: TimeIn is Zero and Uninitialized!")
	}

	return tzu.TimeInLoc.String(), nil
}

// Get LocationOut - - Returns the time zone location for the
// TimeInLoc data field which is part of the current TimeZoneDto
// structure.
func (tzu *TimeZoneUtility) GetLocationOut() (string, error) {

	ePrefix := "TimeZoneDto.GetLocationOut() "

	if tzu.TimeOut.IsZero() {
		return "", errors.New(ePrefix + "Error: TimeOut is Zero and Uninitialized!")
	}

	return tzu.TimeOutLoc.String(), nil
}


// GetZoneIn - Returns The Time Zone for the TimeIn
// data field contained in the current TimeZoneDto
// structure.
func (tzu *TimeZoneUtility) GetZoneIn() (string, error) {

	ePrefix := "TimeZoneDto.GetZoneIn() "

	if tzu.TimeOut.IsZero() {
		return "", errors.New(ePrefix + "Error: TimeOut is Zero and Uninitialized!")
	}

	tzZone, _ := tzu.TimeIn.Zone()

	return tzZone, nil

}

// GetZoneOut - Returns The Time Zone for the TimeOut
// data field contained in the current TimeZoneDto
// structure.
func (tzu *TimeZoneUtility) GetZoneOut() (string, error) {

	ePrefix := "TimeZoneDto.GetZoneOut() "

	if tzu.TimeOut.IsZero() {
		return "", errors.New(ePrefix + "Error: TimeOut is Zero and Uninitialized!")
	}

	tzZone, _ := tzu.TimeOut.Zone()

	return tzZone, nil

}

// IsTimeZoneUtilityValid - Analyzes the current TimeZoneDto
// instance and returns an error if the instance is Invalid.
func (tzu *TimeZoneUtility) IsTimeZoneUtilityValid() error {

	ePrefix := "TimeZoneDto.IsTimeZoneUtilityValid() "

	if tzu.TimeIn.IsZero() {
		return errors.New(ePrefix + "Error: TimeIn is Zero!")
	}

	if tzu.TimeOut.IsZero() {
		return errors.New(ePrefix + "Error: TimeOut is Zero!")
	}

	if tzu.TimeUTC.IsZero() {
		return errors.New(ePrefix + "Error: TimeUTC is Zero!")
	}

	if tzu.TimeLocal.IsZero() {
		return errors.New(ePrefix + "Error: TimeLocal is Zero!")
	}

	if tzu.TimeInLoc == nil {
		tzu.TimeInLoc = tzu.TimeIn.Location()
	}

	if tzu.TimeOutLoc == nil {
		tzu.TimeOutLoc = tzu.TimeOut.Location()
	}

	return nil

}

// IsValidTimeZone - Tests a Time Zone string and returns three boolean values
// designating whether the passed Time Zone string is:
// (1.) a valid time zone
// (2.) a valid iana time zone
// (3.) a valid Local time zone
func (tzu *TimeZoneUtility) IsValidTimeZone(tZone string) (isValidTz, isValidIanaTz, isValidLocalTz bool) {

	isValidTz = false

	isValidIanaTz = false

	isValidLocalTz = false

	if tZone == "" {
		return
	}

	if tZone == "Local" {
		isValidTz = true
		isValidLocalTz = true
		return
	}

	_, err := time.LoadLocation(tZone)

	if err != nil {
		return
	}

	isValidTz = true

	isValidIanaTz = true

	isValidLocalTz = false

	return

}

// MakeDateTz allows one to create a date time object (time.Time) by
// passing in a DateTzDto structure. Within this structure, the time
// zone is designated either by the IANA Time Zone (DateTzDto.TimeZone)
// or by the string "Local" which specifies the the time zone local to the
// user computer.
//
// Note: If dtTzDto.TimeZone is an empty string, this method will default
// the time zone to "Local".
func (tzu *TimeZoneUtility) MakeDateTz(dtTzDto DateTzDto) (time.Time, error) {

	var err error
	var tzLoc *time.Location
	tOut := time.Time{}


	if dtTzDto.IANATimeZone == "" {

		dtTzDto.IANATimeZone = "Local"

	} else {


		if isValid ,_,_ := tzu.IsValidTimeZone(dtTzDto.IANATimeZone); !isValid {
			return tOut, fmt.Errorf("TimeZoneDto.MakeDateTz() Invalid Time Zone Error. Tz = %v.", dtTzDto.IANATimeZone )

		}
	}

	tzLoc, err = time.LoadLocation(dtTzDto.IANATimeZone)

	if err!= nil {
		return tOut, fmt.Errorf("TimeZoneDto.MakeDateTz() Error Loading Location! Invalid Time Zone Error. Tz = %v. Error: %v", dtTzDto.IANATimeZone, err.Error())
	}

	tOut = time.Date(dtTzDto.Year, time.Month(dtTzDto.Month), dtTzDto.Day, dtTzDto.Hour, dtTzDto.Minute, dtTzDto.Second, dtTzDto.Nanosecond, tzLoc)

	return tOut, nil
}

// New - Initializes and returns a new TimeZoneDto object.
//
// Input Parameters
// ----------------
//
// tIn					time.Time	- The input time object.
//
// timeZoneOut	string		- The IANA Time Zone or Time Zone 'Local' to which
// 													input parameter 'tIn' will be converted.
//													Reference https://golang.org/pkg/time/#LoadLocation
// 													and https://www.iana.org/time-zones for information
//													on IANA Time Zones.
//
// The two input parameters are used to populate and return
// a TimeZoneDto structure
//				type TimeZoneDto struct {
//									Description string
//									TimeIn      time.Time
//									TimeInLoc   *time.Location
//									TimeOut     time.Time
//									TimeOutLoc  *time.Location
//									TimeUTC     time.Time
//				}
//
//
func (tzu TimeZoneUtility) New(tIn time.Time, timeZoneOut string) (TimeZoneUtility, error) {

	tzuOut := TimeZoneUtility{}

	return tzuOut.ConvertTz(tIn, timeZoneOut)
}

// NewAddDate - receives four parameters: a TimeZoneDto 'tzuIn' and integer values for
// 'years', 'months' and 'days'.  The 'years', 'months' and 'days' values are added to the
// 'tzuIn' date time values and returned as a new TimeZoneDto instance.
func (tzu TimeZoneUtility) NewAddDate(tzuIn TimeZoneUtility, years int, months int, days int) (TimeZoneUtility, error) {
	ePrefix := "TimeZoneDto.NewAddDate()"

	err:= tzuIn.IsTimeZoneUtilityValid()

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error: Input parameter tzuIn (TimeZoneDto) is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	if years == 0 && months == 0 && days == 0 {
		return tzuOut, nil
	}

	err = tzuOut.AddDate(years, months, days)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDate(years, months, days) years='%v' months='%v' days='%v'  Error='%v'",years, months, days, err.Error())
	}

	return tzuOut, nil
}

// NewAddDuration - receives two input parameters, a TimeZoneDto 'tzuIn' and a
// time 'duration'. 'tzuIn' is adjusted for the specified 'duration' and the resulting
// new TimeZoneDto is returned.
//
func (tzu TimeZoneUtility) NewAddDuration(tzuIn TimeZoneUtility, duration time.Duration) (TimeZoneUtility, error) {
	ePrefix := "TimeZoneDto.NewAddDuration() "

	err := tzuIn.IsTimeZoneUtilityValid()

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error: Input Parameter 'tzIn' is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	err = tzuOut.AddDuration(duration)

	if err != nil {
		return TimeZoneUtility{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDuration(duration). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// ReclassifyTimeWithNewTz - Receives a valid time (time.Time) value and changes the existing time zone
// to that specified in the 'tZone' parameter. During this time reclassification operation, the time
// zone is changed but the time value remains unchanged.
func (tzu *TimeZoneUtility) ReclassifyTimeWithNewTz(tIn time.Time, tZone string) (time.Time, error) {
	strTime := tzu.TimeWithoutTimeZone(tIn)

	isValidTz, _, _ := tzu.IsValidTimeZone(tZone)

	if !isValidTz {
		return time.Time{}, fmt.Errorf("TimeZoneDto:ReclassifyTimeWithNewTz() Error: Input Time Zone is INVALID!")
	}

	tzNew, err := time.LoadLocation(tZone)

	if err != nil {
		return time.Time{}, fmt.Errorf("TimeZoneDto:ReclassifyTimeWithNewTz() - Error from time.Location('%v') - Error: %v", tZone, err.Error())
	}

	tOut, err := time.ParseInLocation(NeutralDateFmt, strTime, tzNew)

	if err != nil {
		return tOut, fmt.Errorf("TimeZoneDto:ReclassifyTimeWithNewTz() - Error from time.Parse - Error: %v", err.Error())
	}

	return tOut, nil
}

// SetTimeIn - Assigns value to field 'TimeIn'
func (tzu *TimeZoneUtility) SetTimeIn(tIn time.Time) {
	tzu.TimeIn = tIn
	tzu.TimeInLoc = tIn.Location()
}

// SetTimeOut - Assigns value to field 'TimeOut'
func (tzu *TimeZoneUtility) SetTimeOut(tOut time.Time) {
	tzu.TimeOut = tOut
	tzu.TimeOutLoc = tOut.Location()
}

// SetUTCTime - Assigns UTC Time to field 'TimeUTC'
func (tzu *TimeZoneUtility) SetUTCTime(t time.Time) {

	tzu.TimeUTC = t.UTC()
}

// SetLocalTime - Assigns Local Time to field 'TimeLocal'
func (tzu *TimeZoneUtility) SetLocalTime(t time.Time) error {
	ePrefix := "TimeZoneDto.SetLocalTime() "

	tzLocal, err := time.LoadLocation("Local")

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by time.LoadLocation(\"Local\") Error='%v'", err.Error())
	}

	tzu.TimeLocal = t.In(tzLocal)

	return nil
}

// Sub - Subtracts the input TimeZoneDto from the current TimeZoneDto
// and returns the duration.
func (tzu *TimeZoneUtility) Sub(tzu2 TimeZoneUtility) (time.Duration, error) {

	ePrefix := "TimeZoneDto.Sub() "

	err := tzu.IsTimeZoneUtilityValid()

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix + "Error: Current TimeZoneDto (tzu) is INVALID. Error='%v'", err.Error())
	}

	err = tzu2.IsTimeZoneUtilityValid()

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix + "Error: Input Parameter 'tzu2' is INVALID! Error='%v'", err.Error())
	}

	return tzu.TimeLocal.Sub(tzu2.TimeLocal), nil
}

// TimeWithoutTimeZone - Returns a Time String containing
// NO time zone. - When the returned string is converted to
// time.Time - in defaults to a UTC time zone.
func (tzu *TimeZoneUtility) TimeWithoutTimeZone(t time.Time) string {
	return t.Format(NeutralDateFmt)
}
