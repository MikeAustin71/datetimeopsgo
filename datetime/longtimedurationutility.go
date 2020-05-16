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

	var bigDur *big.Int

	if longDur.sign == -1 {
		bigDur = big.NewInt(0).Neg(longDur.duration)
	} else {
		bigDur = big.NewInt(0).Set(longDur.duration)
	}

	incrementalDuration := big.NewInt(int64(duration))

	newDur := big.NewInt(0).Add(incrementalDuration, bigDur)

	bigZero := big.NewInt(0)

	if bigZero.Cmp(newDur) == 0 {

		longDur.sign = 0
		longDur.duration.Set(newDur)

	} else if bigZero.Cmp(newDur) == 1 {

		longDur.sign = -1
		longDur.duration = big.NewInt(0).Neg(newDur)

	} else {

		longDur.sign = 1
		longDur.duration = big.NewInt(0).Set(newDur)
	}

	return nil
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

	if longDur1.lock == nil {
		longDur1.lock = new(sync.Mutex)
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

	if longDur2.lock == nil {
		longDur2.lock = new(sync.Mutex)
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

	var bigDur *big.Int

		bigDur = big.NewInt(0).Set(longDur1.duration)

	incrementalDuration := big.NewInt(0).Set(longDur2.duration)

	newDur := big.NewInt(0).Add(incrementalDuration, bigDur)

	bigZero := big.NewInt(0)

	if bigZero.Cmp(newDur) == 0 {

		longDur1.sign = 0
		longDur1.duration.Set(newDur)

	} else if bigZero.Cmp(newDur) == 1 {

		longDur1.sign = -1
		longDur1.duration = big.NewInt(0).Neg(newDur)

	} else {

		longDur1.sign = 1
		longDur1.duration = big.NewInt(0).Set(newDur)
	}

	var finalStartDateTime, finalEndDateTime time.Time

	if longDur1.sign == -1 {

		finalStartDateTime,
		finalEndDateTime,
		err = lngDurUtil2.getStartDateMinusDuration(
			longDur1,
			longDur1.endDateTimeTz.dateTimeValue,
			ePrefix)
	} else {

		finalStartDateTime,
			finalEndDateTime,
			err = lngDurUtil2.getStartDatePlusDuration(
			longDur1,
			longDur1.startDateTimeTz.dateTimeValue,
			ePrefix)

	}

	dTzUtil := dateTzDtoUtility{}

	err = dTzUtil.setFromTimeTzName(
		&longDur1.startDateTimeTz,
		finalStartDateTime,
		TzConvertType.Relative(),
		longDur1.startDateTimeTz.GetTimeZoneName(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

	if err != nil {
		return  err
	}

	err = dTzUtil.setFromTimeTzName(
		&longDur1.endDateTimeTz,
		finalEndDateTime,
		TzConvertType.Relative(),
		longDur1.endDateTimeTz.GetTimeZoneName(),
		FmtDateTimeYrMDayFmtStr,
		ePrefix)

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

// getStartDatePlusDuration - Used to add duration to a starting date
// time in those cases where LongTimeDuration.duration is greater than
// zero.
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