package datetime

import (
	"math/big"
	"sync"
)

type CalendarCycleDto struct {
	yearsInCycle         *big.Int
	daysInCycle          *big.Int
	cycleCount           *big.Int
	cycleCountTotalDays  *big.Int // cycleCountTotalDays = cycleCount X daysInCycle
	cycleCountTotalYears *big.Int // cycleCountTotalYears = cycleCount X yearsInCycle
	remainderDays        *big.Int
	remainderYears       *big.Int
	lock                 *sync.Mutex
}

// CopyOut - Returns a deep copy of the current CalendarCycleDto
// instance
func (calCycDto *CalendarCycleDto) CopyOut() (
	newCalCycle CalendarCycleDto) {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.cycleCount == nil {
		calCycDto.cycleCount = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalDays == nil {
		calCycDto.cycleCountTotalDays = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalYears == nil {
		calCycDto.cycleCountTotalYears = big.NewInt(0)
	}

	if calCycDto.remainderDays == nil {
		calCycDto.remainderDays = big.NewInt(0)
	}

	if calCycDto.remainderYears == nil {
		calCycDto.remainderYears = big.NewInt(0)
	}

	newCalCycle = CalendarCycleDto{
		yearsInCycle:         big.NewInt(0).Set(calCycDto.yearsInCycle),
		daysInCycle:          big.NewInt(0).Set(calCycDto.daysInCycle),
		cycleCount:           big.NewInt(0).Set(calCycDto.cycleCount),
		cycleCountTotalDays:  big.NewInt(0).Set(calCycDto.cycleCountTotalDays),
		cycleCountTotalYears: big.NewInt(0).Set(calCycDto.cycleCountTotalYears),
		remainderDays:        big.NewInt(0).Set(calCycDto.remainderDays),
		remainderYears:       big.NewInt(0).Set(calCycDto.remainderYears),
		lock:                 new(sync.Mutex),
	}

	return newCalCycle
}

// GetCycleCountTotalDays - Returns the total number of Days
// associated with this cycle count.
//
func (calCycDto *CalendarCycleDto) GetCycleCountTotalDays() *big.Int {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.cycleCount == nil {
		calCycDto.cycleCount = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalDays == nil {
		calCycDto.cycleCountTotalDays = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalYears == nil {
		calCycDto.cycleCountTotalYears = big.NewInt(0)
	}

	if calCycDto.remainderDays == nil {
		calCycDto.remainderDays = big.NewInt(0)
	}

	if calCycDto.remainderYears == nil {
		calCycDto.remainderYears = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycDto.cycleCountTotalDays)
}

// GetCycleCountTotalYears - Returns the total number of years
// associated with this cycle count.
//
func (calCycDto *CalendarCycleDto) GetCycleCountTotalYears() *big.Int {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.cycleCount == nil {
		calCycDto.cycleCount = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalDays == nil {
		calCycDto.cycleCountTotalDays = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalYears == nil {
		calCycDto.cycleCountTotalYears = big.NewInt(0)
	}

	if calCycDto.remainderDays == nil {
		calCycDto.remainderDays = big.NewInt(0)
	}

	if calCycDto.remainderYears == nil {
		calCycDto.remainderYears = big.NewInt(0)
	}


	return big.NewInt(0).
		Set(calCycDto.cycleCountTotalYears)
}

// GetDaysInCycle - Returns the number of days in this cycle.
//
func (calCycDto *CalendarCycleDto) GetDaysInCycle() *big.Int {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.cycleCount == nil {
		calCycDto.cycleCount = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalDays == nil {
		calCycDto.cycleCountTotalDays = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalYears == nil {
		calCycDto.cycleCountTotalYears = big.NewInt(0)
	}

	if calCycDto.remainderDays == nil {
		calCycDto.remainderDays = big.NewInt(0)
	}

	if calCycDto.remainderYears == nil {
		calCycDto.remainderYears = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycDto.daysInCycle)
}

// GetRemainderDays - Returns the number of days remaining
// subtracting this cycle count. This must be set by user.
//
func (calCycDto *CalendarCycleDto) GetRemainderDays() *big.Int {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.cycleCount == nil {
		calCycDto.cycleCount = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalDays == nil {
		calCycDto.cycleCountTotalDays = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalYears == nil {
		calCycDto.cycleCountTotalYears = big.NewInt(0)
	}

	if calCycDto.remainderDays == nil {
		calCycDto.remainderDays = big.NewInt(0)
	}

	if calCycDto.remainderYears == nil {
		calCycDto.remainderYears = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycDto.remainderDays)
}

// GetRemainderYears - Returns the number of years remaining
// subtracting this cycle count. This value must be set by
// the user.
//
func (calCycDto *CalendarCycleDto) GetRemainderYears() *big.Int {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.cycleCount == nil {
		calCycDto.cycleCount = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalDays == nil {
		calCycDto.cycleCountTotalDays = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalYears == nil {
		calCycDto.cycleCountTotalYears = big.NewInt(0)
	}

	if calCycDto.remainderDays == nil {
		calCycDto.remainderDays = big.NewInt(0)
	}

	if calCycDto.remainderYears == nil {
		calCycDto.remainderYears = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycDto.remainderYears)
}

// GetYearsInCycle - Returns the number of years in this
// cycle.
//
func (calCycDto *CalendarCycleDto) GetYearsInCycle() *big.Int {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.cycleCount == nil {
		calCycDto.cycleCount = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalDays == nil {
		calCycDto.cycleCountTotalDays = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalYears == nil {
		calCycDto.cycleCountTotalYears = big.NewInt(0)
	}

	if calCycDto.remainderDays == nil {
		calCycDto.remainderDays = big.NewInt(0)
	}

	if calCycDto.remainderYears == nil {
		calCycDto.remainderYears = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycDto.yearsInCycle)
}

// SetCycleCount - Sets the 'Cycle Count'
//
func (calCycDto *CalendarCycleDto) SetCycleCount(cycleCount *big.Int) {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.remainderDays == nil {
		calCycDto.remainderDays = big.NewInt(0)
	}

	if calCycDto.remainderYears == nil {
		calCycDto.remainderYears = big.NewInt(0)
	}

	calCycDto.cycleCount =
	 	big.NewInt(0).
	 		Set(cycleCount)

	calCycDto.cycleCountTotalDays =
		big.NewInt(0).
			Mul(calCycDto.cycleCount, calCycDto.daysInCycle)

	calCycDto.cycleCountTotalYears =
		big.NewInt(0).
			Mul(calCycDto.cycleCount, calCycDto.yearsInCycle)

	return
}

// SetRemainderDays - Sets the Remaining Days member variable.
//
func (calCycDto *CalendarCycleDto) SetRemainderDays(remainderDays *big.Int) {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.cycleCount == nil {
		calCycDto.cycleCount = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalDays == nil {
		calCycDto.cycleCountTotalDays = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalYears == nil {
		calCycDto.cycleCountTotalYears = big.NewInt(0)
	}

	if calCycDto.remainderYears == nil {
		calCycDto.remainderYears = big.NewInt(0)
	}

	calCycDto.remainderDays =
		big.NewInt(0).
			Set(remainderDays)

}


// SetRemainderYears - Sets the Remaining Years member variable.
//
func (calCycDto *CalendarCycleDto) SetRemainderYears(remainderYears *big.Int) {

	if calCycDto.lock == nil {
		calCycDto.lock = new(sync.Mutex)
	}

	calCycDto.lock.Lock()

	defer calCycDto.lock.Unlock()

	if calCycDto.yearsInCycle == nil {
		calCycDto.yearsInCycle = big.NewInt(0)
	}

	if calCycDto.daysInCycle == nil {
		calCycDto.daysInCycle = big.NewInt(0)
	}

	if calCycDto.cycleCount == nil {
		calCycDto.cycleCount = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalDays == nil {
		calCycDto.cycleCountTotalDays = big.NewInt(0)
	}

	if calCycDto.cycleCountTotalYears == nil {
		calCycDto.cycleCountTotalYears = big.NewInt(0)
	}

	if calCycDto.remainderDays == nil {
		calCycDto.remainderDays = big.NewInt(0)
	}

	calCycDto.remainderYears =
		big.NewInt(0).
			Set(remainderYears)


}