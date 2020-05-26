package datetime

import (
	"math/big"
	"sync"
)

type timeMechanics struct {
	lock       sync.Mutex
}

// computeTimeElementsInt64 - Utility routine to break gross nanoseconds
// int constituent hours, minutes, seconds and remaining nanoseconds. As
// the method name implies, the return values are of type Int64.
//
func (timeMech *timeMechanics) computeTimeElementsInt64(
	grossNanoSeconds int64,
	ePrefix string) (
	hours,
	minutes,
	seconds,
	nanoSeconds int64,
	numericalSign int) {

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "timeMech.computeTimeElements() "

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
		hours = grossNanoSeconds/HourNanoSeconds
		grossNanoSeconds -= hours * HourNanoSeconds
	}

	if grossNanoSeconds >= MinuteNanoSeconds {
		minutes = grossNanoSeconds/MinuteNanoSeconds
		grossNanoSeconds -= minutes * MinuteNanoSeconds
	}

	if grossNanoSeconds >= SecondNanoseconds {
		seconds = grossNanoSeconds/SecondNanoseconds
		grossNanoSeconds -= seconds * SecondNanoseconds
	}

	nanoSeconds = grossNanoSeconds

	return hours, minutes, seconds, nanoSeconds, numericalSign
}

// computeTimeElementsInt - Utility routine to break gross nanoseconds
// int constituent hours, minutes, seconds and remaining nanoseconds. As
// the method name implies, the return values are of type 'int'.
//
func (timeMech *timeMechanics) computeTimeElementsBigInt(
	grossNanoSeconds *big.Int) (
	hours,
	minutes,
	seconds,
	nanoSeconds *big.Int,
	numericalSign int) {

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	hours = big.NewInt(0)
	minutes = big.NewInt(0)
	seconds = big.NewInt(0)
	nanoSeconds = big.NewInt(0)
	numericalSign = 1

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
		return hours, minutes, seconds, nanoSeconds, numericalSign
	}

	compareResult = grossNanoSeconds.Cmp(bigHourNanoSecs)
	var temp *big.Int

	if compareResult > -1 {
		hours = big.NewInt(0).Div(grossNanoSeconds,bigHourNanoSecs)
		temp = big.NewInt(0).Mul(hours, bigHourNanoSecs)
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	compareResult = grossNanoSeconds.Cmp(bigMinuteNanoSecs)

	if compareResult > -1 {
		minutes = big.NewInt(0).Div(grossNanoSeconds,bigMinuteNanoSecs)
		temp = big.NewInt(0).Mul(minutes, bigMinuteNanoSecs)
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	compareResult = grossNanoSeconds.Cmp(bigSecondNanoSecs)

	if compareResult > -1 {
		seconds = big.NewInt(0).Div(grossNanoSeconds,bigMinuteNanoSecs)
		temp = big.NewInt(0).Mul(seconds, bigMinuteNanoSecs)
		grossNanoSeconds = big.NewInt(0).Sub(grossNanoSeconds, temp)
	}

	nanoSeconds = big.NewInt(0).Set(grossNanoSeconds)

	return hours, minutes, seconds, nanoSeconds, numericalSign
}

