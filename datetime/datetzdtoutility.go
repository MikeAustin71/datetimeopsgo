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

 // addDateTime - Adds date time components to the date time value of the
// current DateTzDto instance. The updated date time value is returned to
// the calling function as a new DateTzDto instance.
//
 func (dTzUtil *dateTzDtoUtility) addDateTime(
	dTz *DateTzDto,
	 years,
	 months,
	 days,
	 hours,
	 minutes,
	 seconds,
	 milliseconds,
	 microseconds,
	 nanoseconds int,
	 dateTimeFormatStr,
	 ePrefix string) (DateTzDto, error) {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.addDateTime() "

	newDate := dTz.dateTimeValue.AddDate(years, months, 0)

	totNanoSecs := int64(days) * DayNanoSeconds
	totNanoSecs += int64(hours) * int64(time.Hour)
	totNanoSecs += int64(minutes) * int64(time.Minute)
	totNanoSecs += int64(seconds) * int64(time.Second)
	totNanoSecs += int64(milliseconds) * int64(time.Millisecond)
	totNanoSecs += int64(microseconds) * int64(time.Microsecond)
	totNanoSecs += int64(nanoseconds)

	newDateTime := newDate.Add(time.Duration(totNanoSecs))

	dTz2, err := DateTzDto{}.New(newDateTime, dateTimeFormatStr)

	if err != nil {
		 return DateTzDto{},
			 fmt.Errorf(ePrefix+
			 	"\nError returned from DateTzDto{}.New(newDateTime, dateTimeFormatStr)\n"+
				 "newDateTime='%v'\nError='%v'\n", newDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	return dTz2, nil
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

	if baseDtz == nil {
		return
	}

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

	if dTz == nil {
		return dtz2
	}

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

	if dTz == nil {
		return
	}

	dTz.tagDescription = ""
	dTz.timeComponents.Empty()
	dTz.timeZone = TimeZoneDefDto{}
	dTz.dateTimeValue = time.Time{}
	dTz.dateTimeFmt = ""

	return
}

// isEmptyDateTzDto - Analyzes an instanceof DateTzDto to
// determine if all data fields are uninitialized or zero
// values.
//
func (dTzUtil *dateTzDtoUtility) isEmptyDateTzDto(
		dTz *DateTzDto) bool {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	if dTz == nil {
		return true
	}

	if dTz.tagDescription == "" &&
		dTz.timeComponents.IsEmpty() &&
		dTz.dateTimeValue.IsZero() &&
		dTz.dateTimeFmt == "" &&
		dTz.timeZone.IsEmpty() {

		return true
	}

	return false
}

// isValidDateTzDto - Analyzes an instance of 'DateTzDto' to
// determine if is value. If the instance evaluates as invalid,
// an error is returned.
//
func (dTzUtil *dateTzDtoUtility) isValidDateTzDto(
		dTz *DateTzDto,
		ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.isValidDateTzDto() "

	dTzUtil2 := dateTzDtoUtility{}

	if dTzUtil2.isEmptyDateTzDto(dTz) {
		return errors.New(ePrefix +
			"\nThis 'DateTzDto' instance is EMPTY!\n")
	}

	if dTz.dateTimeValue.IsZero() {
		return errors.New(ePrefix +
			"\nError: DateTzDto.DateTime is ZERO!\n")
	}

	if dTz.timeZone.IsEmpty() {
		return errors.New(ePrefix +
			"\nError: DateTzDto.TimeZone is EMPTY!\n")
	}

	if err := dTz.timeComponents.IsValidDateTime(); err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: dTz.timeComponents is INVALID.\n" +
			"Error='%v'\n", err.Error())
	}

	if !dTz.timeZone.IsValidFromDateTime(dTz.dateTimeValue) {
		return errors.New(ePrefix +
			"\nError: dTz.TimeZone is INVALID!\n")
	}

	return nil
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

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

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

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

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

// setFromTimeTz - Sets the time values for input parameter 'dTz' 
// (type DateTzDto). The new values will be  based on input parameters
// 'dateTime', 'timeZoneLocation' and a date time format string,
// 'dateTimeFmtStr'.
//
func (dTzUtil *dateTzDtoUtility) setFromTimeTz(
			dTz *DateTzDto,
			dateTime time.Time,
			timeZoneLocation,
			dateTimeFmtStr,
			ePrefix string) error {

	dTzUtil.lock.Lock()

	defer dTzUtil.lock.Unlock()

	ePrefix += "dateTzDtoUtility.setFromTimeTz() "

	if dTz == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter dTz (*DateTzDto) is 'nil'!\n")
	}

	if dateTime.IsZero() {
		return errors.New(ePrefix +
			"\nError: Input parameter 'dateTime' is ZERO and INVALID!\n")
	}

	dTzUtil2 := dateTzDtoUtility{}

	timeZoneLocation = dTzUtil2.preProcessTimeZoneLocation(timeZoneLocation)

	tLoc, err := time.LoadLocation(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nINVALID timeZoneLocation. Error returned by time.LoadLocation(timeZoneLocation)\n"+
			"timeZoneLocation='%v'\ntzl='%v'\nError='%v'\n",
			timeZoneLocation, timeZoneLocation, err.Error())
	}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	targetDateTime := dateTime.In(tLoc)

	tZone, err := TimeZoneDefDto{}.New(targetDateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by TimeZoneDefDto{}.New(targetDateTime)\n"+
			"targetDateTime='%v'\nTarget Time Zone Location='%v'\nError='%v'\n",
			targetDateTime.Format(FmtDateTimeYrMDayFmtStr), timeZoneLocation, err.Error())
	}

	tDto, err := TimeDto{}.NewFromDateTime(targetDateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned from TimeDto{}.NewFromDateTime(targetDateTime).\n"+
			"targetDateTime='%v'\nError='%v'\n",
			targetDateTime.Format(FmtDateTimeYrMDayFmtStr), err.Error())
	}

	dTzUtil2.empty(dTz)

	dTz.dateTimeValue = targetDateTime
	dTz.timeZone = tZone.CopyOut()
	dTz.timeComponents = tDto.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}

// setFromTimeDto - Receives data from a TimeDto input parameter
// and sets all data fields of the current DateTzDto instance
// accordingly. When the method completes, the values of the
// current DateTzDto will equal the values of the input parameter
// TimeDto instance.
//
func (dTzUtil *dateTzDtoUtility) setFromTimeDto(
			dTz *DateTzDto,
			tDto TimeDto, 
			timeZoneLocation,
			dateTimeFmtStr,
			ePrefix string) error {

	dTzUtil.lock.Lock()
	
	defer dTzUtil.lock.Unlock()
	
	ePrefix += "dateTzDtoUtility.setFromTimeDto() "

	if tDto.IsEmpty() {

		return fmt.Errorf(ePrefix + "\nError: Input parameter 'tDto' date time elements equal ZERO!\n")
	}

	t2Dto := tDto.CopyOut()

	err := t2Dto.NormalizeTimeElements()

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by t2Dto.NormalizeTimeElements().\nError='%v'\n",
			err.Error())
	}

	t2Dto.ConvertToAbsoluteValues()

	if err = t2Dto.IsValidDateTime(); err != nil {
		return fmt.Errorf(ePrefix+
			"\nError: Input Parameter tDto (TimeDto) is INVALID.\nError='%v'\n",
			err.Error())
	}

	dTzUtil2 := dateTzDtoUtility{}

	timeZoneLocation = dTzUtil2.preProcessTimeZoneLocation(timeZoneLocation)

	_, err = time.LoadLocation(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by time.LoadLocation(timeZoneLocation).\n"+
			"timeZoneLocation='%v'\nError='%v'\n",  timeZoneLocation, err.Error())
	}

	dateTimeFmtStr = dTzUtil2.preProcessDateFormatStr(dateTimeFmtStr)

	dateTime, err := tDto.GetDateTime(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by tDto.GetDateTime(timeZoneLocation).\n"+
			"timeZoneLocation='%v'\nError='%v'\n",
			timeZoneLocation, err.Error())
	}

	timeZoneDef, err := TimeZoneDefDto{}.New(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by TimeZoneDefDto{}.New(dateTime).\n" +
			"dateTime='%v'\nError='%v'\n",
			dateTime.Format(FmtDateTimeYrMDayFmtStr) , err.Error())
	}

	dTzUtil2.empty(dTz)
	dTz.dateTimeValue = dateTime
	dTz.timeZone = timeZoneDef.CopyOut()
	dTz.timeComponents = t2Dto.CopyOut()
	dTz.dateTimeFmt = dateTimeFmtStr

	return nil
}