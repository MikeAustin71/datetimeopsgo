package datetime

import (
	"math/big"
	"sync"
	"time"
)


type longTimeDurationUtility struct {

	lock              sync.Mutex
}

// addDuration - Add a time.Duration value to the current
// value of longDur.duration.
//
func (lngDurUtil *longTimeDurationUtility) addDuration(
	longDur *LongTimeDuration,
	duration time.Duration,
	ePrefix string) (err error) {

	lngDurUtil.lock.Lock()

	defer lngDurUtil.lock.Unlock()

	ePrefix += "longTimeDurationUtility.addDuration() "

	err = nil

	if longDur == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "longDur",
			inputParameterValue: "",
			errMsg:              "Input parameter 'longDur' " +
				"is nil!",
			err:                 nil,
		}

		return err
	}

	if longDur.lock == nil {
		longDur.lock = new(sync.Mutex)
	}

	lngDurUtil2 := longTimeDurationUtility{}

	var isZero bool

	isZero, err = lngDurUtil2.isValid(longDur, ePrefix)

	if err != nil {
		return err
	}

	if isZero {
		return nil
	}

	var currentDur *big.Int

	if longDur.sign == -1 {
		currentDur = big.NewInt(0).Neg(longDur.duration)
	} else {
		currentDur = big.NewInt(0).Set(longDur.duration)
	}

	incrementalDuration := big.NewInt(int64(duration))

	newDurTotal := big.NewInt(0).Add(incrementalDuration, currentDur)

	bigZero := big.NewInt(0)

	var finalStartDateTime, finalEndDateTime time.Time

	timeZoneLocation := longDur.startDateTimeTz.GetTimeZoneName()
	dateTimeFormat := longDur.startDateTimeTz.GetDateTimeFmt()

	dTzUtil := dateTzDtoUtility{}

	if bigZero.Cmp(newDurTotal) == 0 {
		// New Duration Total is zero
		longDur.sign = 0

		longDur.duration = big.NewInt(0)

		err = dTzUtil.setFromTimeTzName(
			&longDur.endDateTimeTz,
			longDur.startDateTimeTz.dateTimeValue,
			TzConvertType.Absolute(),
			timeZoneLocation,
			dateTimeFormat,
			ePrefix)

		return err

	} else if bigZero.Cmp(newDurTotal) == 1 {
		// New Duration Total is Negative
		longDur.sign = -1

		longDur.duration = big.NewInt(0).Neg(newDurTotal)

		finalStartDateTime,
		finalEndDateTime,
		err =
		lngDurUtil2.getStartDateMinusDuration(
			longDur,
			longDur.startDateTimeTz.dateTimeValue,
			ePrefix)

		if err != nil {
			return err
		}

	} else {
		// New Duration Total is Positive
		longDur.sign = 1
		longDur.duration = big.NewInt(0).Set(newDurTotal)

		finalStartDateTime,
			finalEndDateTime,
			err =
			lngDurUtil2.getStartDateMinusDuration(
				longDur,
				longDur.startDateTimeTz.dateTimeValue,
				ePrefix)

		if err != nil {
			return err
		}

	}

	err = dTzUtil.setFromTimeTzName(
		&longDur.startDateTimeTz,
		finalStartDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFormat,
		ePrefix + "Setting longDur.startDateTimeTz ")

	if err != nil {
		return err
	}

	err = dTzUtil.setFromTimeTzName(
		&longDur.endDateTimeTz,
		finalEndDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFormat,
		ePrefix + "Setting longDur.endDateTimeTz ")

	return err
}

// addDuration - Add the duration value of another LongTimeDuration
// instance to longDur.
//
func (lngDurUtil *longTimeDurationUtility) addLongDuration(
	longDur1 *LongTimeDuration,
	longDur2 *LongTimeDuration,
	ePrefix string) (err error) {

	lngDurUtil.lock.Lock()

	defer lngDurUtil.lock.Unlock()

	err = nil

	ePrefix += "longTimeDurationUtility.addDuration() "

	if longDur1 == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "longDur1",
			inputParameterValue: "",
			errMsg:              "Input parameter 'longDur1' " +
				"is nil!",
			err:                 nil,
		}

		return err
	}

	if longDur2 == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "longDur2",
			inputParameterValue: "",
			errMsg:              "Input parameter 'longDur2' " +
				"is nil!",
			err:                 nil,
		}

		return err
	}

	lngDurUtil2 := longTimeDurationUtility{}

	var isZero2 bool

	_, err = lngDurUtil2.isValid(longDur1, ePrefix + "longDur1 ")

	if err != nil {
		return err
	}

	isZero2, err = lngDurUtil2.isValid(longDur2, ePrefix + "longDur2 ")

	if err != nil {
		return err
	}

	if isZero2 {
		return err
	}

	var currentDur, incrementalDuration *big.Int

	if longDur1.sign == -1 {
		currentDur = big.NewInt(0).Neg(longDur1.duration)
	} else {
		currentDur = big.NewInt(0).Set(longDur1.duration)
	}

	if longDur2.sign == -1 {
		incrementalDuration = big.NewInt(0).Neg(longDur2.duration)
	} else {
		incrementalDuration = big.NewInt(0).Set(longDur2.duration)
	}

	newDurTotal := big.NewInt(0).Add(incrementalDuration, currentDur)

	bigZero := big.NewInt(0)

	var finalStartDateTime, finalEndDateTime time.Time

	timeZoneLocation := longDur1.startDateTimeTz.GetTimeZoneName()
	dateTimeFormat := longDur1.startDateTimeTz.GetDateTimeFmt()

	dTzUtil := dateTzDtoUtility{}

	if bigZero.Cmp(newDurTotal) == 0 {
		// New Duration Total is zero
		longDur1.sign = 0

		longDur1.duration = big.NewInt(0)

		err = dTzUtil.setFromTimeTzName(
			&longDur1.endDateTimeTz,
			longDur1.startDateTimeTz.dateTimeValue,
			TzConvertType.Absolute(),
			timeZoneLocation,
			dateTimeFormat,
			ePrefix)

		return err

	} else if bigZero.Cmp(newDurTotal) == 1 {
		// New Duration Total is Negative
		longDur1.sign = -1

		longDur1.duration = big.NewInt(0).Neg(newDurTotal)

		finalStartDateTime,
			finalEndDateTime,
			err =
			lngDurUtil2.getStartDateMinusDuration(
				longDur1,
				longDur1.startDateTimeTz.dateTimeValue,
				ePrefix)

		if err != nil {
			return err
		}

	} else {
		// New Duration Total is Positive
		longDur1.sign = 1
		longDur1.duration = big.NewInt(0).Set(newDurTotal)

		finalStartDateTime,
			finalEndDateTime,
			err =
			lngDurUtil2.getStartDateMinusDuration(
				longDur1,
				longDur1.startDateTimeTz.dateTimeValue,
				ePrefix)

		if err != nil {
			return err
		}

	}

	err = dTzUtil.setFromTimeTzName(
		&longDur1.startDateTimeTz,
		finalStartDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFormat,
		ePrefix + "Setting longDur.startDateTimeTz ")

	if err != nil {
		return err
	}

	err = dTzUtil.setFromTimeTzName(
		&longDur1.endDateTimeTz,
		finalEndDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFormat,
		ePrefix + "Setting longDur.endDateTimeTz ")

	return err
}

// isValid - Tests an instance of LongTimeDuration
// returns an error if it is invalid.
//
// In addition the method returns a boolean value
// signaling whether the longDur.duration value is
// zero.
//
func (lngDurUtil *longTimeDurationUtility) isValid(
	longDur *LongTimeDuration,
	ePrefix string) (isZero bool, err error) {

	lngDurUtil.lock.Lock()

	defer lngDurUtil.lock.Unlock()

	ePrefix += "longTimeDurationUtility.isValid() "
	isZero = false
	err = nil

		if longDur == nil {
			err = &InputParameterError{
				ePrefix:             ePrefix,
				inputParameterName:  "longDur",
				inputParameterValue: "",
				errMsg:              "'longDur' is " +
					"nil!",
				err:                 nil,
			}

		return isZero, err
	}

	if longDur.lock == nil {
		longDur.lock = new(sync.Mutex)
	}

	if longDur.duration == nil {

		longDur.duration = big.NewInt(0)

		isZero = true

		return isZero, err
	}

	bigZero := big.NewInt(0)

	if bigZero.Cmp(longDur.duration) == 0 {
		longDur.sign = 0
		isZero = true
		return isZero, err
	}


	if bigZero.Cmp(longDur.duration) == 1 &&
		longDur.sign != -1 {

		longDur.sign = -1

		return isZero, err
	}

	if bigZero.Cmp(longDur.duration) == -1 &&
		longDur.sign != 1 {

		longDur.sign = -1
	}

	return isZero, err
}


// getStartDatePlusDuration - Calculates and returns the starting and
// ending date times associated with a positive time duration.
//
// In order to calculate these two dates, the caller must supply the
// 'starting' date time.
//
func (lngDurUtil *longTimeDurationUtility) getStartDatePlusDuration(
	longDur *LongTimeDuration,
	startDateTime time.Time,
	ePrefix string) ( finalStartDateTime, finalEndDateTime time.Time, err error) {

	lngDurUtil.lock.Lock()

	defer lngDurUtil.lock.Unlock()

	ePrefix += "longTimeDurationUtility.getStartDatePlusDuration() "

	finalStartDateTime = startDateTime
	finalEndDateTime = time.Time{}
	err = nil

	lngDurUtil2 := longTimeDurationUtility{}

	var isZero bool

	isZero, err = lngDurUtil2.isValid(longDur, ePrefix)

	if err != nil {
		return finalStartDateTime, finalEndDateTime, err
	}

	if isZero {

		finalEndDateTime = finalStartDateTime
		longDur.sign = 0

		return finalStartDateTime, finalEndDateTime, err
	}

	if longDur.sign == -1 {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "LongTimeDuration.sign",
			inputParameterValue: "",
			errMsg:              "'LongTimeDuration.sign' is '-1'.\n" +
				"This method only processes positive duration values!\n" +
				"'LongTimeDuration.sign' MUST BE '+1'.",
			err:                 nil,
		}

		return finalStartDateTime, finalEndDateTime, err
	}

	// assume duration is positive
	// signVal := lngDur.duration.Sign()

	finalEndDateTime = time.Date(
		startDateTime.Year() + 200,
		startDateTime.Month(),
		startDateTime.Day(),
		startDateTime.Hour(),
		startDateTime.Minute(),
		startDateTime.Second(),
		startDateTime.Nanosecond(),
		startDateTime.Location())

	remainingDuration := big.NewInt(0).Set(longDur.duration)

	incrementalDuration := finalEndDateTime.Sub(startDateTime)

	tempDur := big.NewInt(0).SetInt64(int64(incrementalDuration))

	if tempDur.Cmp(remainingDuration) < 1 {
		// tempDur less than or equal to remainingDuration

		startDateTime = finalEndDateTime

		remainingDuration = big.NewInt(0).Sub(remainingDuration, tempDur)

		finalEndDateTime = time.Date(
			startDateTime.Year() + 200,
			startDateTime.Month(),
			startDateTime.Day(),
			startDateTime.Hour(),
			startDateTime.Minute(),
			startDateTime.Second(),
			startDateTime.Nanosecond(),
			startDateTime.Location())

		incrementalDuration = finalEndDateTime.Sub(startDateTime)

		tempDur = big.NewInt(0).SetInt64(int64(incrementalDuration))
	}

	incrementalDuration = time.Duration(remainingDuration.Int64())

	finalEndDateTime = startDateTime.Add(incrementalDuration)

	return finalStartDateTime, finalEndDateTime, err
}

// getStartDateMinusDuration - Calculates and returns the starting and
// ending date times associated with a negative time duration.
//
// In order to calculate these two dates, the caller must supply the
// 'ending' date time.
//
func (lngDurUtil *longTimeDurationUtility) getStartDateMinusDuration(
	longDur *LongTimeDuration,
	endDateTime time.Time,
	ePrefix string) ( finalStartDateTime, finalEndDateTime time.Time, err error) {

	lngDurUtil.lock.Lock()

	defer lngDurUtil.lock.Unlock()

	ePrefix += "longTimeDurationUtility.getStartDateMinusDuration() "

	finalEndDateTime = endDateTime
	finalStartDateTime = time.Time{}
	err = nil

	lngDurUtil2 := longTimeDurationUtility{}

	var isZero bool

	isZero, err = lngDurUtil2.isValid(longDur, ePrefix)

	if err != nil {
		return finalStartDateTime, finalEndDateTime, err
	}

	if isZero {
		finalStartDateTime = finalEndDateTime
		return finalStartDateTime, finalEndDateTime, err
	}

	if longDur.sign == 1 {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "LongTimeDuration.sign",
			inputParameterValue: "",
			errMsg:              "'LongTimeDuration.sign' is '+1'.\n" +
				"This method only processes negative duration values!\n" +
				"'LongTimeDuration.sign' MUST BE '-1'.",
			err:                 nil,
		}

		return finalStartDateTime, finalEndDateTime, err
	}

	finalStartDateTime = time.Date(
		endDateTime.Year() - 200,
		endDateTime.Month(),
		endDateTime.Day(),
		endDateTime.Hour(),
		endDateTime.Minute(),
		endDateTime.Second(),
		endDateTime.Nanosecond(),
		endDateTime.Location())

	remainingDuration := big.NewInt(0).Set(longDur.duration)

	incrementalDuration := endDateTime.Sub(finalStartDateTime)

	tempDur := big.NewInt(0).SetInt64(int64(incrementalDuration))

	if tempDur.Cmp(remainingDuration) < 1 {
		// tempDur less than or equal to remainingDuration

		endDateTime = finalStartDateTime

		remainingDuration = big.NewInt(0).Sub(remainingDuration, tempDur)

		finalStartDateTime = time.Date(
			endDateTime.Year() - 200,
			endDateTime.Month(),
			endDateTime.Day(),
			endDateTime.Hour(),
			endDateTime.Minute(),
			endDateTime.Second(),
			endDateTime.Nanosecond(),
			endDateTime.Location())

		incrementalDuration = endDateTime.Sub(finalStartDateTime)

		tempDur = big.NewInt(0).SetInt64(int64(incrementalDuration))
	}

	incrementalDuration = time.Duration(remainingDuration.Int64() * -1)

	finalStartDateTime = endDateTime.Add(incrementalDuration)

	return finalStartDateTime, finalEndDateTime, nil
}

// setStartEndDateDuration - Calculates long duration from starting
// and ending dates
func (lngDurUtil *longTimeDurationUtility) setStartEndDateDuration(
	longDur *LongTimeDuration,
	startDateTime,
	endDateTime time.Time,
	timeZoneLocation string,
	dateTimeFormat string,
	ePrefix string) error {

	lngDurUtil.lock.Lock()

	defer lngDurUtil.lock.Unlock()

	ePrefix += "longTimeDurationUtility.setStartEndDateDuration() "

	var err error

	lngDurUtil2 := longTimeDurationUtility{}

	_, err = lngDurUtil2.isValid(longDur, ePrefix)

	if err != nil {
		return err
	}

	longDur.duration = big.NewInt(0)
	longDur.sign = 0

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTimeTzName(
		&longDur.startDateTimeTz,
		startDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFormat,
		ePrefix + "Setting longDur.startDateTimeTz ")

	if err != nil {
		return err
	}

	err = dTzUtil.setFromTimeTzName(
		&longDur.endDateTimeTz,
		endDateTime,
		TzConvertType.Relative(),
		timeZoneLocation,
		dateTimeFormat,
		ePrefix + "Setting longDur.endDateTimeTz ")

	if err != nil {
		return err
	}

	if longDur.startDateTimeTz.dateTimeValue.Equal(
		longDur.endDateTimeTz.dateTimeValue) {
		return nil
	}

	newStartDateTime := longDur.startDateTimeTz.dateTimeValue

	finalEndDateTime := longDur.endDateTimeTz.dateTimeValue

	tempEndDateTime := time.Date(
		newStartDateTime.Year() + 250,
		newStartDateTime.Month(),
		newStartDateTime.Day(),
		newStartDateTime.Hour(),
		newStartDateTime.Minute(),
		newStartDateTime.Second(),
		newStartDateTime.Nanosecond(),
		newStartDateTime.Location())

	var incrementalDuration time.Duration
	var currentLngDur, tempLngDur *big.Int

	currentLngDur = big.NewInt(0)

	for tempEndDateTime.Before(finalEndDateTime) {

		incrementalDuration = tempEndDateTime.Sub(newStartDateTime)

		tempLngDur = big.NewInt(int64(incrementalDuration))

		currentLngDur = big.NewInt(0).Add(
			longDur.duration,
			tempLngDur)

		longDur.duration = big.NewInt(0).Set(currentLngDur)

		newStartDateTime = tempEndDateTime

		tempEndDateTime = time.Date(
			newStartDateTime.Year() + 250,
			newStartDateTime.Month(),
			newStartDateTime.Day(),
			newStartDateTime.Hour(),
			newStartDateTime.Minute(),
			newStartDateTime.Second(),
			newStartDateTime.Nanosecond(),
			newStartDateTime.Location())
	}

	incrementalDuration = finalEndDateTime.Sub(newStartDateTime)

	tempLngDur = big.NewInt(int64(incrementalDuration))

	currentLngDur = big.NewInt(0).Add(
		longDur.duration,
		tempLngDur)

	longDur.duration = big.NewInt(0).Set(currentLngDur)

	return nil
}

// setZeroLongTermDuration - Receives a pointer to a LongTimeDuration
// instance and proceeds to set all of the data values to zero.
//
func (lngDurUtil *longTimeDurationUtility) setZeroLongTermDuration(
	longDur *LongTimeDuration,
	ePrefix string) error {

	lngDurUtil.lock.Lock()

	defer lngDurUtil.lock.Unlock()

	ePrefix += "longTimeDurationUtility.setZeroLongTermDuration() "

	var err error

	if longDur == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "longDur",
			inputParameterValue: "",
			errMsg:              "'longDur' is " +
				"nil!",
			err:                 nil,
		}

		return err
	}

	if longDur.lock == nil {
		longDur.lock = new(sync.Mutex)
	}

	longDur.duration = big.NewInt(0)

	longDur.sign = 0

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setZeroDateTimeTz(
		&longDur.startDateTimeTz,
		ePrefix + "Setting Starting Date Time ")

	if err != nil {
		return err
	}

	err = dTzUtil.setZeroDateTimeTz(
		&longDur.endDateTimeTz,
		ePrefix + "Setting Ending Date Time ")

	return err
}