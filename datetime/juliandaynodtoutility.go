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

	intSign := int64(1)

	if julianDayNo < 0 {
		intSign = -1

		julianDayNo = julianDayNo * intSign
	}

	if julianDayNoTimeFraction.Sign() < 0 {
		julianDayNoTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(julianDayNoTimeFraction.Prec()).
				Neg(julianDayNoTimeFraction)
	}

	bigJulianDayNo := big.NewFloat(0.0).
		SetMode(big.ToZero).
		SetPrec(julianDayNoTimeFraction.Prec()).
		SetInt64(julianDayNo)

	requestedPrecision :=	julianDayNoTimeFraction.Prec() +
	 bigJulianDayNo.Prec() + 10

	// requestedPrecision := uint(100)

	jDNDto.julianDayNoNanoSecs =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(requestedPrecision).
			Add(bigJulianDayNo,
			julianDayNoTimeFraction)

	jDNDto.julianDayNoFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(requestedPrecision).
			Sub(jDNDto.julianDayNoNanoSecs, bigJulianDayNo)

	originalFraction := julianDayNoTimeFraction.Text('f',int(requestedPrecision))
	fmt.Printf("setDto original Fraction                           %v\n",
		originalFraction)

	fractionStr := jDNDto.julianDayNoFraction.Text('f', int(requestedPrecision))
	fmt.Printf("setDto jDNDto.julianDayNoFraction                  %v\n",
		fractionStr)

	if intSign < 0 {
		jDNDto.julianDayNoNanoSecs =
			big.NewFloat(0.0).
				Neg(jDNDto.julianDayNoNanoSecs)
	}

	bigDayNanoSeconds := big.NewFloat(0.0).
		SetMode(big.ToZero).
		SetPrec(0).
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
		SetPrec(1).
		Mul(bigDayNanoSeconds, jDNDto.julianDayNoFraction)

	grossNanoSecs.SetMode(big.ToZero).
		SetPrec(1)


	if !grossNanoSecs.IsInt() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: grossNanoSecs did NOT convert to an integer!\n" +
			"grossNanoSecs='%v'\n",
			grossNanoSecs.Text('f', 0))

		return err
	}

	var accuracy big.Accuracy

	jDNDto.totalNanoSeconds, accuracy = grossNanoSecs.Int64()

	if accuracy != big.Exact {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Calculation of grossNanoSecInt64 yielded " +
			"an inexact resutl!\n" +
			"Accuracy='%v'\n", accuracy)

		return err
	}

	fmt.Printf("setDto jDNDto.totalNanoSeconds:              %v\n",
		jDNDto.totalNanoSeconds)

	timeMech := timeMechanics{}
	jDNDto.hours,
		jDNDto.minutes,
		jDNDto.seconds,
		jDNDto.nanoseconds = timeMech.computeTimeElementsInt(
		jDNDto.totalNanoSeconds)

	newDayNoNanSecsStr := jDNDto.julianDayNoNanoSecs.Text('f', int(requestedPrecision))
	fmt.Printf("setDto julianDayNoNanoSecs:                  %v\n",
		newDayNoNanSecsStr)

	return  err
}

