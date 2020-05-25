package datetime

import "sync"

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
	nanoSeconds int64) {

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	ePrefix += "timeMech.computeTimeElements() "

	hours = 0
	minutes = 0
	seconds = 0
	nanoSeconds = 0

	sign := int64(1)

	if grossNanoSeconds < 0 {
		sign = -1
		grossNanoSeconds = grossNanoSeconds * sign
	}

	if grossNanoSeconds == 0 {
		return hours, minutes, seconds, nanoSeconds
	}

	if grossNanoSeconds >= HourNanoSeconds {
		hours = grossNanoSeconds/HourNanoSeconds
		grossNanoSeconds -= hours * HourNanoSeconds
		hours = hours * sign
	}

	if grossNanoSeconds >= MinuteNanoSeconds {
		minutes = grossNanoSeconds/MinuteNanoSeconds
		grossNanoSeconds -= minutes * MinuteNanoSeconds
		minutes = minutes * sign
	}

	if grossNanoSeconds >= SecondNanoseconds {
		seconds = grossNanoSeconds/SecondNanoseconds
		grossNanoSeconds -= seconds * SecondNanoseconds
		seconds = seconds * sign
	}

	nanoSeconds = grossNanoSeconds
	nanoSeconds = nanoSeconds * sign

	return hours, minutes, seconds, nanoSeconds
}

// computeTimeElementsInt - Utility routine to break gross nanoseconds
// int constituent hours, minutes, seconds and remaining nanoseconds. As
// the method name implies, the return values are of type 'int'.
//
func (timeMech *timeMechanics) computeTimeElementsInt(
	grossNanoSeconds int64) (
	hours,
	minutes,
	seconds,
	nanoSeconds int) {

	timeMech.lock.Lock()

	defer timeMech.lock.Unlock()

	hours = 0
	minutes = 0
	seconds = 0
	nanoSeconds = 0

	sign := int(1)

	if grossNanoSeconds < 0 {
		sign = -1
		grossNanoSeconds = grossNanoSeconds * int64(sign)
	}

	if grossNanoSeconds == 0 {
		return hours, minutes, seconds, nanoSeconds
	}

	if grossNanoSeconds >= HourNanoSeconds {
		hours = int(grossNanoSeconds/HourNanoSeconds)
		grossNanoSeconds -= int64(hours) * HourNanoSeconds
		hours = hours * sign
	}

	if grossNanoSeconds >= MinuteNanoSeconds {
		minutes = int(grossNanoSeconds/MinuteNanoSeconds)
		grossNanoSeconds -= int64(minutes) * MinuteNanoSeconds
		minutes = minutes * sign
	}

	if grossNanoSeconds >= SecondNanoseconds {
		seconds = int(grossNanoSeconds/SecondNanoseconds)
		grossNanoSeconds -= int64(seconds) * SecondNanoseconds
		seconds = seconds * sign
	}

	nanoSeconds = int(grossNanoSeconds)
	nanoSeconds = nanoSeconds * sign

	return hours, minutes, seconds, nanoSeconds
}

