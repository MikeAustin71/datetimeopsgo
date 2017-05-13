package common

import (
	"errors"
	"fmt"
	"time"
)

// NOTE: See https://golang.org/pkg/time/#LoadLocation
// and https://www.iana.org/time-zones to ensure that
// the IANA Time Zone Database is properly configured
// on your system.
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
)

// TimeZoneUtility - Time Zone Data and Methods
type TimeZoneUtility struct {
	TimeIn         time.Time
	TimeInLoc      *time.Location
	TimeInZoneStr  string
	TimeOut        time.Time
	TimeOutLoc     *time.Location
	TimeOutZoneStr string
	TimeUTC        time.Time
}

// ConvertTz - Convert Time from existing time zone to targetTZone.
// The results are stored in the TimeZoneUtility data structure
func (tzu *TimeZoneUtility) ConvertTz(tIn time.Time, targetTZone string) error {

	if targetTZone == "" {
		return errors.New("TimeZoneUtility:ConvertTz() Error: targetTZone is empty")
	}

	tzUTC, err := time.LoadLocation("Zulu")

	if err != nil {
		return errors.New("TimeZoneUtility:ConvertTz() - Error Loading UTC Time Zone " + err.Error())
	}

	tzOut, err := time.LoadLocation(targetTZone)

	if err != nil {
		return fmt.Errorf("TimeZoneUtility:ConvertTz() - Error Loading Target Time Zone %v. Errors: %v ", targetTZone, err.Error())
	}

	tzu.setTimeIn(tIn)

	tzu.setTimeOut(tIn.In(tzOut))

	tzu.setUTCTime(tIn, tzUTC)

	return nil
}

func (tzu *TimeZoneUtility) setTimeIn(tIn time.Time) {
	tzu.TimeIn = tIn
	tzu.TimeInLoc = tIn.Location()
	tzu.TimeInZoneStr = tzu.TimeInLoc.String()
}

func (tzu *TimeZoneUtility) setTimeOut(tOut time.Time) {
	tzu.TimeOut = tOut
	tzu.TimeOutLoc = tOut.Location()
	tzu.TimeOutZoneStr = tzu.TimeOutLoc.String()
}

func (tzu *TimeZoneUtility) setUTCTime(t time.Time, utcLoc *time.Location) {

	tzu.TimeUTC = t.In(utcLoc)
}
