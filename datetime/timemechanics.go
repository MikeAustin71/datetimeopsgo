package datetime

import (
	"math"
	"math/big"
	"sync"
)

type TimeMechanics struct {
	lock       * sync.Mutex
}

// ComputeBigIntNanoseconds - Utility method to sum days, hours, minutes,
// seconds and subMicrosecondNanoseconds and return total subMicrosecondNanoseconds as a type *big.Int.
//
func (timeMech *TimeMechanics) ComputeBigIntNanoseconds(
	days *big.Int,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	ePrefix string) (
	totalNanoseconds *big.Int,
	err error) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "TimeMechanics.ComputeBigIntNanoseconds() "

	totalNanoseconds = big.NewInt(0)
	err = nil

	if days == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "days",
			inputParameterValue: "",
			errMsg:              "Error: Input parameter 'days' is nil!",
			err:                 nil,
		}

		return totalNanoseconds, err
	}

	temp := big.NewInt(0).
		Mul(days, big.NewInt(DayNanoSeconds))

	totalNanoseconds.Add(totalNanoseconds, temp)

	temp = big.NewInt(int64(hours) * HourNanoSeconds)

	totalNanoseconds.Add(totalNanoseconds, temp)

	temp = big.NewInt(int64(minutes) * MinuteNanoSeconds)

	totalNanoseconds.Add(totalNanoseconds, temp)

	temp = big.NewInt(int64(seconds) * SecondNanoseconds)

	totalNanoseconds.Add(totalNanoseconds, temp)

	temp = big.NewInt(int64(nanoseconds))

	totalNanoseconds.Add(totalNanoseconds, temp)

	return totalNanoseconds, err
}

// ComputeTimeElementsInt64 - Utility method to break gross subMicrosecondNanoseconds
// int constituent hours, minutes, seconds and remaining subMicrosecondNanoseconds. As
// the method name implies, the return values are of type Int64.
//
func (timeMech *TimeMechanics) ComputeTimeElementsInt64(
	grossNanoSeconds int64) (
	hours,
	minutes,
	seconds,
	nanoSeconds int,
	numericalSign int) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	hours = 0
	minutes = 0
	seconds = 0
	nanoSeconds = 0

	numericalSign = 1

	if grossNanoSeconds < 0 {
		numericalSign = -1
		grossNanoSeconds = grossNanoSeconds *int64(numericalSign)
	}

	if grossNanoSeconds == 0 {
		numericalSign = 0
		return hours, minutes, seconds, nanoSeconds, numericalSign
	}

	if grossNanoSeconds >= HourNanoSeconds {
		hours = int(grossNanoSeconds/HourNanoSeconds)
		grossNanoSeconds -= int64(hours) * HourNanoSeconds
	}

	if grossNanoSeconds >= MinuteNanoSeconds {
		minutes = int(grossNanoSeconds/MinuteNanoSeconds)
		grossNanoSeconds -= int64(minutes) * MinuteNanoSeconds
	}

	if grossNanoSeconds >= SecondNanoseconds {
		seconds = int(grossNanoSeconds/SecondNanoseconds)
		grossNanoSeconds -= int64(seconds) * SecondNanoseconds
	}

	nanoSeconds = int(grossNanoSeconds)

	return hours, minutes, seconds, nanoSeconds, numericalSign
}

// ComputeTimeElementsBigInt - Utility routine to break gross subMicrosecondNanoseconds
// int constituent hours, minutes, seconds and remaining subMicrosecondNanoseconds. As
// the method name implies, the return values are of type 'int'.
//
func (timeMech *TimeMechanics) ComputeTimeElementsBigInt(
	grossNanoSeconds *big.Int) (
	days *big.Int,
	hours,
	minutes,
	seconds,
	nanoseconds,
	numericalSign int) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	days = big.NewInt(0)
	hours = 0
	minutes = 0
	seconds = 0
	nanoseconds = 0
	hoursBig := big.NewInt(0)
	minutesBig := big.NewInt(0)
	secondsBig := big.NewInt(0)
	numericalSign = 1

	bigDayNanoSecs := big.NewInt(DayNanoSeconds)
	bigHourNanoSecs := big.NewInt(HourNanoSeconds)
	bigMinuteNanoSecs := big.NewInt(MinuteNanoSeconds)
	bigSecondNanoSecs := big.NewInt(SecondNanoseconds)

	compareResult := big.NewInt(0).Cmp(grossNanoSeconds)

	if compareResult < 0 {
		numericalSign = -1
		grossNanoSeconds = big.NewInt(0).Abs(grossNanoSeconds)
	}

	if compareResult == 0 {
		numericalSign = 0
		return days, hours, minutes, seconds, nanoseconds, numericalSign
	}

	var temp *big.Int

	compareResult =  grossNanoSeconds.Cmp(bigDayNanoSecs)

	if compareResult > -1 {
		days = big.NewInt(0).Div(grossNanoSeconds, bigDayNanoSecs)
		temp = big.NewInt(0).Mul(days, bigDayNanoSecs)
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	compareResult = grossNanoSeconds.Cmp(bigHourNanoSecs)

	if compareResult > -1 {
		hoursBig = big.NewInt(0).Div(grossNanoSeconds,bigHourNanoSecs)
		temp = big.NewInt(0).Mul(hoursBig, bigHourNanoSecs)
		hours = int(hoursBig.Int64())
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	compareResult = grossNanoSeconds.Cmp(bigMinuteNanoSecs)

	if compareResult > -1 {
		minutesBig = big.NewInt(0).Div(grossNanoSeconds,bigMinuteNanoSecs)
		temp = big.NewInt(0).Mul(minutesBig, bigMinuteNanoSecs)
		minutes = int(minutesBig.Int64())
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	compareResult = grossNanoSeconds.Cmp(bigSecondNanoSecs)

	if compareResult > -1 {
		secondsBig = big.NewInt(0).Div(grossNanoSeconds,bigSecondNanoSecs)
		temp = big.NewInt(0).Mul(secondsBig, bigSecondNanoSecs)
		seconds = int(secondsBig.Int64())
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	nanoseconds = int(grossNanoSeconds.Int64())

	return days, hours, minutes, seconds, nanoseconds, numericalSign
}

// ComputeTimeElementsInt - Utility routine to break gross subMicrosecondNanoseconds
// int constituent hours, minutes, seconds and remaining subMicrosecondNanoseconds. As
// the method name implies, the return values are of type 'int'.
//
func (timeMech *TimeMechanics) ComputeTimeElementsInt(
	grossNanoSeconds int64) (
	hours,
	minutes,
	seconds,
	nanoSeconds int,
	numericalSign int) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	hours = 0
	minutes = 0
	seconds = 0
	nanoSeconds = 0
	numericalSign = 1

	if grossNanoSeconds < 0 {
		numericalSign = -1
		grossNanoSeconds *= -1
	}

	if grossNanoSeconds == 0 {
		numericalSign = 0
		return hours, minutes, seconds, nanoSeconds, numericalSign
	}

	if grossNanoSeconds >= HourNanoSeconds {
		hours = int(grossNanoSeconds/HourNanoSeconds)
		grossNanoSeconds -= int64(hours) * HourNanoSeconds
	}

	if grossNanoSeconds >= MinuteNanoSeconds {
		minutes = int(grossNanoSeconds/MinuteNanoSeconds)
		grossNanoSeconds -= int64(minutes) * MinuteNanoSeconds
	}

	if grossNanoSeconds >= SecondNanoseconds {
		seconds = int(grossNanoSeconds/SecondNanoseconds)
		grossNanoSeconds -= int64(seconds) * SecondNanoseconds
	}

	nanoSeconds = int(grossNanoSeconds)

	return hours, minutes, seconds, nanoSeconds, numericalSign
}

// ComputeFloat64TimeFracToGregorianSeconds - Utility routine to
// compute time elements to nearest second from a float64
// Julian Day Number Time. Constituent hours, minutes and
// seconds are returned as type int in Gregorian Calendar
// time.
//
// Julian Days start at noon. Gregorian days start at
// midnight. This method adjusts the hours to reflect
// Gregorian days.
func (timeMech *TimeMechanics) ComputeFloat64TimeFracToGregorianSeconds(
	julianDayNoTime float64) (
	days int64,
	hours,
	minutes,
	seconds,
	numericalSign int) {

	if timeMech.lock == nil {
		timeMech.lock = new(sync.Mutex)
	}

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	days = 0
	hours = 0
	minutes = 0
	seconds = 0
	numericalSign = 1

	if julianDayNoTime < 0.0 {
		numericalSign = -1
		julianDayNoTime = math.Abs(julianDayNoTime)
	}

	fracSeconds :=
		int64((julianDayNoTime * 86400.0) + 0.5)

	// 86400 seconds in a 24-hour day
	if fracSeconds >= 86400 {
		days = fracSeconds / int64(86400)

		fracSeconds -= days * 86400
	}

	if fracSeconds >= 3600 {
		hours = int(fracSeconds/3600)
		fracSeconds -= int64(hours) * int64(3600)
		if hours >= 12 {
			hours -= 12
		}
	}

	if fracSeconds >= 60 {
		minutes = int(fracSeconds/60)
		fracSeconds -= int64(minutes) * int64(60)
	}

	seconds = int(fracSeconds)

	return days, hours, minutes, seconds, numericalSign
}

