package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type dateTzDtoUtility struct {
 	input    string
 	output   string
 	lock     sync.Mutex
 }


// copyIn - Receives two parameters which are pointers
// to types DateTzDto. The method then copies all of
// the data field values from 'incomingDtz' into
// 'baseDtz'.
//
func (dTzUtil *dateTzDtoUtility) copyIn(
	baseDtz ,
	incomingDtz *DateTzDto) {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	dTzUtil2 := dateTzDtoUtility{}

	dTzUtil2.empty(baseDtz)

	baseDtz.tagDescription = incomingDtz.tagDescription
	baseDtz.timeComponents = incomingDtz.timeComponents.CopyOut()
	baseDtz.dateTimeFmt = incomingDtz.dateTimeFmt

	if incomingDtz.dateTimeValue.IsZero() {
		baseDtz.timeZone = TimeZoneDefDto{}
		baseDtz.dateTimeValue = time.Time{}
	} else {
		baseDtz.dateTimeValue = incomingDtz.dateTimeValue
		baseDtz.timeZone = incomingDtz.timeZone.CopyOut()
	}

}

// copyOut - Returns a deep copy of input parameter
// 'dTz' which is a pointer to a type 'DateTzDto'.
func (dTzUtil *dateTzDtoUtility) copyOut(
	dTz *DateTzDto) DateTzDto {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	 dtz2 := DateTzDto{}

	 dtz2.tagDescription = dTz.tagDescription
	 dtz2.timeComponents = dTz.timeComponents.CopyOut()
	 dtz2.dateTimeFmt = dTz.dateTimeFmt

	 if dTz.dateTimeValue.IsZero() {
		 dtz2.timeZone = TimeZoneDefDto{}
		 dtz2.dateTimeValue = time.Time{}
	 } else {
		 dtz2.dateTimeValue = dTz.dateTimeValue
		 dtz2.timeZone = dTz.timeZone.CopyOut()
	 }

	 return dtz2
}

// empty - Receives a pointer to a type 'DateTzDto' and
// proceeds to set all internal member variables to their
// 'zero' or uninitialized values.
//
func (dTzUtil *dateTzDtoUtility) empty(dTz *DateTzDto) {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	dTz.tagDescription = ""
	dTz.timeComponents.Empty()
	dTz.timeZone = TimeZoneDefDto{}
	dTz.dateTimeValue = time.Time{}
	dTz.dateTimeFmt = ""

	return
}


func (dTzUtil *dateTzDtoUtility) preProcessDateFormatStr(
		dateTimeFmtStr string) string {

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}

func (dTzUtil *dateTzDtoUtility) preProcessTimeZoneLocation(
		timeZoneLocation string) string {

	if len(timeZoneLocation) == 0 {
		return TZones.Other.UTC()
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		return "Local"
	}

	return timeZoneLocation
}


// setFromDateTime - Sets the values for DateTzDto fields encapsulated
// in input parameter 'dTz'. The field values will be set
// based on an input parameter 'dateTime' (Type time.time) and a
// date time format string.
func (dTzUtil *dateTzDtoUtility) setFromDateTime(
	dTz *DateTzDto,
	dateTime time.Time,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromDateTime() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input Parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO!\n")
	}

	if len(dateTimeFmtStr) == 0 {
		dateTimeFmtStr = FmtDateTimeYrMDayFmtStr
	}

	tDto, err := TimeDto{}.NewFromDateTime(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned from TimeDto{}.NewFromDateTime(dateTime).\n"+
			"dateTime='%v'\nError='%v'\n",
			dateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	timeZone, err := TimeZoneDefDto{}.New(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned from TimeZoneDefDto{}.New(dateTime). "+
			"dateTime='%v'  Error='%v'",
			dateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dateTzUtil := dateTzDtoUtility{}

	dateTzUtil.empty(dTz)

	dTz.dateTimeValue = dateTime
	dTz.timeComponents = tDto.CopyOut()
	dTz.timeZone = timeZone.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}