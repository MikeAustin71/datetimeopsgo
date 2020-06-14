package datetime

import (
	"fmt"
	"math"
	"math/big"
	"sync"
)

type julianDayNoDtoUtility struct {

	lock     sync.Mutex

}

// setDtoFromFloat64 - Receives a pointer to a type JulianDayNoDto
// and proceeds to compute an populate its Julian Day Number
// and Time data elements using a float64 input parameter.
//
func (jDNDtoUtil *julianDayNoDtoUtility) setDtoFromFloat64(
	jDNDto *JulianDayNoDto,
	julianDayNoTime float64,
	ePrefix string) (err error) {

	jDNDtoUtil.lock.Lock()

	defer jDNDtoUtil.lock.Unlock()

	err = nil

	ePrefix += "julianDayNoDtoUtility.setDto() "

	if jDNDto == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "jDNDto",
			inputParameterValue: "",
			errMsg:              "Input parameter 'jDNDto' " +
				"is a nil pointer!",
			err:                 nil,
		}

		return err
	}

	numericalSign := 1

	if math.Signbit(julianDayNoTime) {
		numericalSign = -1
		julianDayNoTime = math.Abs(julianDayNoTime)
	}

	julianDayNoFloat64, julianFracFloat64 :=
		math.Modf(julianDayNoTime)

	julianDayNoInt64 := int64(julianDayNoFloat64)

	julianTimeFracBigFloat := big.NewFloat(0).
		SetMode(big.ToNearestAway).
		SetPrec(1024).
		SetFloat64(julianFracFloat64)

	dayNoNanoSecsBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(1024).
		SetInt64(DayNanoSeconds)

	rawNanoSecsBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(1024).
		Mul(julianTimeFracBigFloat, dayNoNanoSecsBigFloat)

	rawNanoSecsBigInt, _ := rawNanoSecsBigFloat.Int(nil)

	timeMech := TimeMechanics{}

	_,
	hours,
	minutes,
	seconds,
	nanoseconds,
	_ := timeMech.ComputeTimeElementsBigInt(rawNanoSecsBigInt)

	// Round to nearest second
	if nanoseconds >= 500000000 {
		seconds++
	}

	rawNanoSecsBigInt,
		err =	timeMech.ComputeBigIntNanoseconds(
		big.NewInt(0),
		hours,
		minutes,
		seconds,
		0,
		ePrefix)

	if err != nil {
		return err
	}

	rawNanoSecsBigFloat = big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(1024).
		SetInt(rawNanoSecsBigInt)

	julianTimeFracBigFloat = big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(1024).
		Quo(rawNanoSecsBigFloat, dayNoNanoSecsBigFloat)

	jDNDtoUtil2 := julianDayNoDtoUtility{}

	if numericalSign == -1 {
		julianDayNoInt64 *= -1
	}

	return jDNDtoUtil2.setDto(
		jDNDto,
		julianDayNoInt64,
		julianTimeFracBigFloat,
		ePrefix)
}

// setDto - Receives an instance of JulianDayNoDto and
// proceeds to compute and populate its data fields using
// a Julian Day Number integer and Time Fraction of type
// *big.Float.
//
func (jDNDtoUtil *julianDayNoDtoUtility) setDto(
	jDNDto *JulianDayNoDto,
	julianDayNo int64,
	julianDayNoTimeFraction *big.Float,
	ePrefix string) (err error) {

	jDNDtoUtil.lock.Lock()

	defer jDNDtoUtil.lock.Unlock()

	err = nil

	ePrefix += "julianDayNoDtoUtility.setDto() "

	if jDNDto == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "jDNDto",
			inputParameterValue: "",
			errMsg:              "Input parameter 'jDNDto' " +
				"is a nil pointer!",
			err:                 nil,
		}

		return err
	}

	if julianDayNoTimeFraction == nil {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNoTimeFraction",
			inputParameterValue: "",
			errMsg:              "Error: 'julianDayNoTimeFraction' is a 'nil' pointer!",
			err:                 nil,
		}

		return err
	}

	jDNDto.julianDayNoNumericalSign = 1

	if julianDayNo < 0 {

		jDNDto.julianDayNoNumericalSign = -1

		julianDayNo = julianDayNo * int64(-1)
	}

	jDNDto.julianDayNo = big.NewInt(julianDayNo)

	requestedPrecision :=	uint(1024)

	if julianDayNoTimeFraction.Prec() > requestedPrecision {
		requestedPrecision = julianDayNoTimeFraction.Prec()
	}

	//fmt.Printf("setDto Original Fraction            %80.70f\n",
	//	julianDayNoTimeFraction)

	if julianDayNoTimeFraction.Sign() < 0 {
		julianDayNoTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(requestedPrecision).
				Neg(julianDayNoTimeFraction)
	}

	jDNDto.julianDayNoFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(requestedPrecision).
			Set(julianDayNoTimeFraction)

	bigJulianDayNo :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(requestedPrecision).
			SetInt64(julianDayNo)

	jDNDto.julianDayNoTime =
	big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		Add(bigJulianDayNo, julianDayNoTimeFraction)

	bigDayNanoSeconds := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		SetInt64(DayNanoSeconds)

	if !bigDayNanoSeconds.IsInt() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: bigDayNanoSeconds did NOT convert to an integer!\n" +
			"bigDayNanoSeconds='%v'\n",
			bigDayNanoSeconds.Text('f', 0))

		return err
	}

	grossNanoSecs := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		Mul(bigDayNanoSeconds, jDNDto.julianDayNoFraction)

	grossNanoSecs.
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		Add(big.NewFloat(0.5), grossNanoSecs)

	// Always less than or equal to 36-hours
	jDNDto.totalJulianNanoSeconds, _ = grossNanoSecs.Int64()

	// Always less than or equal to 24-hours
	jDNDto.netGregorianNanoSeconds = jDNDto.totalJulianNanoSeconds

	if jDNDto.netGregorianNanoSeconds >= NoonNanoSeconds {
		jDNDto.netGregorianNanoSeconds -= NoonNanoSeconds
	} else {
		jDNDto.netGregorianNanoSeconds += NoonNanoSeconds
	}

	timeMech := TimeMechanics{}
	jDNDto.hours,
		jDNDto.minutes,
		jDNDto.seconds,
		jDNDto.nanoseconds,
		_ = timeMech.ComputeTimeElementsInt64(
					jDNDto.netGregorianNanoSeconds)

	return  err
}

