package datetime

import (
	"fmt"
	"math/big"
	"sync"
)

type julianDayNoDtoUtility struct {

	lock     sync.Mutex

}

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

	jDNDto.julianDayNoSign = 1

	if julianDayNo < 0 {

		jDNDto.julianDayNoSign = -1

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

	// Always less than 24-hours
	jDNDto.totalNanoSeconds, _ = grossNanoSecs.Int(nil)

	timeMech := TimeMechanics{}
	_,
	jDNDto.hours,
		jDNDto.minutes,
		jDNDto.seconds,
		jDNDto.nanoseconds,
		_ = timeMech.ComputeTimeElementsBigInt(
					jDNDto.totalNanoSeconds)

	return  err
}

