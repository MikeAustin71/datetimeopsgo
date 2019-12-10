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

// preProcessDateFormatStr - Provides a standardized method
// for implementing a default date time format string.
//
func (dTzUtil *dateTzDtoUtility) preProcessDateFormatStr(
		dateTimeFmtStr string) string {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

	if len(dateTimeFmtStr) == 0 {
		return FmtDateTimeYrMDayFmtStr
	}

	return dateTimeFmtStr
}

// preProcessTimeZoneLocation - Provides a method for implementing
// standardized processing of time zone location strings.
func (dTzUtil *dateTzDtoUtility) preProcessTimeZoneLocation(
		timeZoneLocation string) string {

	dTzUtil.lock.Lock()
	defer dTzUtil.lock.Unlock()

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

	dTzUtil2 := dateTzDtoUtility{}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

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
			"\nError returned from TimeZoneDefDto{}.New(dateTime).\n"+
			"dateTime='%v'\nError='%v'\n",
			dateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}


	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = dateTime
	dTz.timeComponents = tDto.CopyOut()
	dTz.timeZone = timeZone.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}

// SetFromDateTimeComponents - Sets the values of the Date Time fields
// for the current DateTzDto instance based on time components
// and a Time Zone Location.
//
// Note that this variation of time elements breaks time down by
// hour, minute, second, millisecond, microsecond and nanosecond.
//
func (dTzUtil *dateTzDtoUtility) setFromDateTimeComponents(
	dTz *DateTzDto,
	year,
	month,
	day,
	hour,
	minute,
	second,
	millisecond,
	microsecond,
	nanosecond int,
	timeZoneLocation,
	dateTimeFmtStr,
	ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromDateTimeComponents() "

	tDto, err := TimeDto{}.New(year, month, 0, day, hour, minute,
		second, millisecond, microsecond, nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by TimeDto{}.New(year, month,...).  "+
			"Error='%v'", err.Error())
	}

	dTzUtil2 := dateTzDtoUtility{}

	fmtStr := dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	tzl := dTzUtil2.preProcessTimeZoneLocation(timeZoneLocation)

	_, err = time.LoadLocation(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by time.LoadLocation(tzl).\nINVALID 'timeZoneLocation'!\n"+
			"tzl='%v'\ntimeZoneLocation='%v'\nError='%v'\n",
			tzl, timeZoneLocation, err.Error())
	}

	dt, err := tDto.GetDateTime(tzl)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto.GetDateTime(tzl).\n"+
			"\ntimeZoneLocation='%v'\ntzl='%v'\nError='%v'\n",
			timeZoneLocation, tzl, err.Error())
	}

	timeZone, err := TimeZoneDefDto{}.New(dt)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by TimeZoneDefDto{}.New(dt).\n"+
			"dt='%v'\nError=%v\n",
			dt.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = dt
	dTz.timeZone = timeZone.CopyOut()
	dTz.timeComponents = tDto.CopyOut()
	dTz.dateTimeFmt = fmtStr

	return nil
}


// setFromDateTimeElements - Sets the values of input parameter
// 'dTz' (type DateTzDto). 'dTz' data fields are set based on
// input parameters consisting of date time elements,
// a time zone location and a date time format string.
//
// Date Time elements include year, month, day, hour, minute,
// second and nanosecond.
//
func (dTzUtil *dateTzDtoUtility) setFromDateTimeElements(
			dTz *DateTzDto,
			year,
			month,
			day,
			hour,
			minute,
			second,
			nanosecond int,
			timeZoneLocation,
			dateTimeFmtStr,
			ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromDateTimeElements() "

	tDto, err := TimeDto{}.New(year, month, 0, day, hour, minute, second,
		0, 0, nanosecond)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned from TimeDto{}.New(year, month, ...).\n"+
			"Error='%v'\n", err.Error())
	}

	dTzUtil2 := dateTzDtoUtility{}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	timeZoneLocation = dTzUtil2.preProcessTimeZoneLocation(timeZoneLocation)

	_, err = time.LoadLocation(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by time.LoadLocation(tzl).\n" +
			"INVALID 'timeZoneLocation'!\n"+
			"tzl='%v'\ntimeZoneLocation='%v'\n" +
			"Error='%v'\n",
			timeZoneLocation, timeZoneLocation, err.Error())
	}

	dt, err := tDto.GetDateTime(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto.GetDateTime(tzl).\n"+
			"\ntimeZoneLocation='%v'\ntzl='%v'\n" +
			"Error='%v'\n",
			timeZoneLocation, timeZoneLocation, err.Error())
	}

	timeZone, err := TimeZoneDefDto{}.New(dt)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by TimeZoneDefDto{}.New(dt).\n"+
			"tzl='%v'\ntimeZonelocation='%v'\ndt='%v'\n" +
			"Error='%v'\n",
			timeZoneLocation,
			timeZoneLocation,
			dt.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = dt
	dTz.timeZone = timeZone.CopyOut()
	dTz.timeComponents = tDto.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}