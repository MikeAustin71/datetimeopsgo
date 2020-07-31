package datetime

import (
	"math/big"
	"sync"
)

type CalendarCycleConfiguration struct {
	mainCycleStartDateForPositiveJDNNo DateTransferDto // Main Cycle Start Date is less than JDN Base Date
	//                                                            //  Used for positive JDN Number calculations
	mainCycleAdjustmentYearsForPositiveJDNNo *big.Int             // Adjustment Years for JDN Base Date and positive
	//                                                            //  JDN Number calculations. Always a negative value.
	mainCycleAdjustmentDaysForPositiveJDNNo  *big.Int             // Adjustment Years for JDN Base Date and positive
	//                                                            //  JDN Number Calculations. Always a negative value.
	mainCycleStartDateForNegativeJDNNo DateTransferDto // Main Cycle Start Date is greater than JDN Base
	//                                                            //  Date. Used for Negative JDN Number calculations
	mainCycleAdjustmentYearsForNegativeJDNNo  *big.Int            // Adjustment Years for JDN Base Date and negative
	//                                                            //  JDN Number calculations. Always a negative value.
	mainCycleAdjustmentDaysForNegativeJDNNo   *big.Int            // Adjustment Years for JDN Base Date and negative
	//                                                            //  JDN Number calculations.
	jdnBaseStartYearDateTime          DateTransferDto // Base Start date time for JDN calculations
	ordinalFixedDateStartYearDateTime DateTransferDto // Base Start date time for Ordinal Fixed Date
	//                                                            //  calculations.
	mainCycleConfig                           CalendarCycleDto   // Calendar cycle for main calendar cycle.
	calendarCyclesConfig                      []CalendarCycleDto // Ths array contains all of the calendar cycles
	//                                                            //  applied to JDN and Ordinal/Fixed Date
	//                                                            //  calculations.
	lock                                      *sync.Mutex
}

// GetMainCycleStartDateForPositiveJDNNo - Returns Main Cycle starting
// date/time information to be used in calculating positive Julian Day
// Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleStartDateForPositiveJDNNo() DateTransferDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	return calCycCfg.mainCycleStartDateForPositiveJDNNo.CopyOut()
}

// GetMainCycleAdjustmentYearsForPositiveJDNNo - Returns Main Cycle Adjustment
// Years used in calculating positive Julian Day Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleAdjustmentYearsForPositiveJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if calCycCfg.mainCycleAdjustmentYearsForPositiveJDNNo == nil {
		calCycCfg.mainCycleAdjustmentYearsForPositiveJDNNo = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycCfg.mainCycleAdjustmentYearsForPositiveJDNNo)
}

// GetMainCycleAdjustmentDaysForPositiveJDNNo - Returns Main Cycle Adjustment
// Days used in calculating positive Julian Day Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleAdjustmentDaysForPositiveJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if calCycCfg.mainCycleAdjustmentDaysForPositiveJDNNo == nil {
		calCycCfg.mainCycleAdjustmentDaysForPositiveJDNNo = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycCfg.mainCycleAdjustmentDaysForPositiveJDNNo)
}

// GetMainCycleStartDateForNegativeJDNNo - Returns Main Cycle starting
// date/time information to be used in calculating negative Julian Day
// Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleStartDateForNegativeJDNNo() DateTransferDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	return calCycCfg.mainCycleStartDateForNegativeJDNNo.CopyOut()
}

// GetMainCycleAdjustmentYearsForNegativeJDNNo - Returns Main Cycle Adjustment
// Years used in calculating negative Julian Day Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleAdjustmentYearsForNegativeJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if calCycCfg.mainCycleAdjustmentYearsForNegativeJDNNo == nil {
		calCycCfg.mainCycleAdjustmentYearsForNegativeJDNNo = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycCfg.mainCycleAdjustmentYearsForNegativeJDNNo)
}

// GetMainCycleAdjustmentDaysForNegativeJDNNo - Returns Main Cycle Adjustment
// Days used in calculating negative Julian Day Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleAdjustmentDaysForNegativeJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if calCycCfg.mainCycleAdjustmentDaysForNegativeJDNNo == nil {
		calCycCfg.mainCycleAdjustmentDaysForNegativeJDNNo = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycCfg.mainCycleAdjustmentDaysForNegativeJDNNo)
}

// GetJDNBaseStartYearDateTime - Returns the base starting year and
// date/time for Julian Day Number calculations.
//
func( calCycCfg *CalendarCycleConfiguration) GetJDNBaseStartYearDateTime() DateTransferDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	return calCycCfg.jdnBaseStartYearDateTime.CopyOut()
}

// GetOrdinalFixedDateBaseStartYearDateTime - Returns the base starting year and
// date/time for Ordinal or Fixed Day Number calculations.
//
func( calCycCfg *CalendarCycleConfiguration) GetOrdinalFixedDateBaseStartYearDateTime() DateTransferDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	return calCycCfg.ordinalFixedDateStartYearDateTime.CopyOut()
}

// GetMainCycleConfiguration - Returns calendar main cycle configuration
// information.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleConfiguration() CalendarCycleDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	return calCycCfg.mainCycleConfig.CopyOut()
}

// GetCalendarCycleConfigurations - Returns a array of CalendarCycleDto objects
// covering all of the calendar cycles associated with the Gregorian
// calendar.
//
func( calCycCfg *CalendarCycleConfiguration) GetCalendarCycleConfigurations() []CalendarCycleDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	lenArray := len(calCycCfg.calendarCyclesConfig)

	calConfigs := make([]CalendarCycleDto, lenArray)

	if lenArray == 0 {
		return calConfigs
	}

	for i:=0; i < lenArray; i++ {
		calConfigs[i] = calCycCfg.calendarCyclesConfig[i].CopyOut()
	}

	return calConfigs
}

// SetCalendarCycleConfigurations - Receives an array of CalendarCycleDto objects
// and proceeds to copy configuration to the internal member variable
// CalendarCycleConfiguration.calendarCyclesConfig.
//
func( calCycCfg *CalendarCycleConfiguration) SetCalendarCycleConfigurations(calConfigs []CalendarCycleDto )  {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	lenArray := len(calConfigs)

	if lenArray == 0 {
		return
	}

	calCycCfg.calendarCyclesConfig = make([]CalendarCycleDto, lenArray)

	for i:=0; i < lenArray; i++ {
		calCycCfg.calendarCyclesConfig[i] = calConfigs[i].CopyOut()
	}

	return
}

// SetJDNBaseStartYearDateTime - Sets internal member variable
// CalendarCycleConfiguration.jdnBaseStartYearDateTime.
//
func( calCycCfg *CalendarCycleConfiguration) SetJDNBaseStartYearDateTime(
	jdnBaseStartYearDateTime DateTransferDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.jdnBaseStartYearDateTime =
		jdnBaseStartYearDateTime.CopyOut()

	return
}

// SetMainCycleAdjustmentDaysForNegativeJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleAdjustmentDaysForNegativeJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleAdjustmentDaysForNegativeJDNNo(
	mainCycleAdjustmentDaysForNegativeJDNNo *big.Int) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if mainCycleAdjustmentDaysForNegativeJDNNo == nil {
		mainCycleAdjustmentDaysForNegativeJDNNo = big.NewInt(0)
	}

	calCycCfg.mainCycleAdjustmentDaysForNegativeJDNNo =
		big.NewInt(0).
			Set(mainCycleAdjustmentDaysForNegativeJDNNo)

	return
}

// SetMainCycleAdjustmentDaysForPositiveJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleAdjustmentDaysForPositiveJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleAdjustmentDaysForPositiveJDNNo(
	mainCycleAdjustmentDaysForPositiveJDNNo *big.Int) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if mainCycleAdjustmentDaysForPositiveJDNNo == nil {
		mainCycleAdjustmentDaysForPositiveJDNNo = big.NewInt(0)
	}

	calCycCfg.mainCycleAdjustmentDaysForPositiveJDNNo =
		big.NewInt(0).
			Set(mainCycleAdjustmentDaysForPositiveJDNNo)
}

// SetMainCycleAdjustmentYearsForNegativeJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleAdjustmentYearsForNegativeJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleAdjustmentYearsForNegativeJDNNo(
	mainCycleAdjustmentYearsForNegativeJDNNo *big.Int) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if mainCycleAdjustmentYearsForNegativeJDNNo == nil {
		mainCycleAdjustmentYearsForNegativeJDNNo = big.NewInt(0)
	}

	calCycCfg.mainCycleAdjustmentYearsForNegativeJDNNo =
		big.NewInt(0).
			Set(mainCycleAdjustmentYearsForNegativeJDNNo)
}

// SetMainCycleAdjustmentYearsForPositiveJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleAdjustmentYearsForPositiveJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleAdjustmentYearsForPositiveJDNNo(
	mainCycleAdjustmentYearsForPositiveJDNNo *big.Int) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if mainCycleAdjustmentYearsForPositiveJDNNo == nil {
		mainCycleAdjustmentYearsForPositiveJDNNo = big.NewInt(0)
	}
	
	calCycCfg.mainCycleAdjustmentYearsForPositiveJDNNo =
		big.NewInt(0).
			Set(mainCycleAdjustmentYearsForPositiveJDNNo)
}

// SetMainCycleConfig - Sets internal member variable
// CalendarCycleConfiguration.mainCycleConfig.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleConfig(
	mainCycleConfig CalendarCycleDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.mainCycleConfig =
		mainCycleConfig.CopyOut()

	return
}

// SetMainCycleStartDateForNegativeJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleStartDateForNegativeJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleStartDateForNegativeJDNNo(
	mainCycleStartDateForNegativeJDNNo DateTransferDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.mainCycleStartDateForNegativeJDNNo =
		mainCycleStartDateForNegativeJDNNo.CopyOut()
}

// SetMainCycleStartDateForPositiveJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleStartDateForPositiveJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleStartDateForPositiveJDNNo(
	mainCycleStartDateForPositiveJDNNo DateTransferDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.mainCycleStartDateForPositiveJDNNo =
		mainCycleStartDateForPositiveJDNNo.CopyOut()
}

// SetOrdinalFixedDateStartYearDateTime - Sets internal member variable
// CalendarCycleConfiguration.ordinalFixedDateStartYearDateTime.
//
func( calCycCfg *CalendarCycleConfiguration) SetOrdinalFixedDateStartYearDateTime(
	ordinalFixedDateStartYearDateTime DateTransferDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.ordinalFixedDateStartYearDateTime =
		ordinalFixedDateStartYearDateTime.CopyOut()

	return
}
