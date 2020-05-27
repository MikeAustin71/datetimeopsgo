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

		julianDayNo = julianDayNo * int64(jDNDto.julianDayNoSign)
	}


	requestedPrecision :=	uint(200)

	if julianDayNoTimeFraction.Prec() > requestedPrecision {
		requestedPrecision = julianDayNoTimeFraction.Prec()
	}

	if requestedPrecision > 1024 {
		requestedPrecision = 1024
	}

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

	julianDayNoTimeFracStr :=
		julianDayNoTimeFraction.Text('f',
			int(julianDayNoTimeFraction.Prec()))

	julianDayNoTimeStr :=
		fmt.Sprintf("%v", julianDayNo) +
			julianDayNoTimeFracStr[1:]

	fmt.Printf("julianDayNoTimeStr:                   %v\n",
		julianDayNoTimeStr)

	var b int
	var err2 error

	jDNDto.julianDayNoNanoSecs,
	b,
	err2 = big.NewFloat(0).
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		Parse(julianDayNoTimeStr, 10)

	if err2 != nil {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error returned by Parse(julianDayNoTimeStr, 10)\n" +
			"julianDayNoTimeStr='%v'\n" +
			"Error='%v'\n", julianDayNoTimeStr, err2.Error())
		return err
	}

	if b != 10 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Parse(julianDayNoTimeStr, 10) did NOT return b=10.\n" +
			"b='%v'\n" +
			"julianDayNoTimeStr='%v'\n", b, julianDayNoTimeStr)
		return err
	}

	/*
	bigJulianDayNo := big.NewFloat(0.0).
		SetMode(big.ToZero).
		SetPrec(requestedPrecision).
		SetInt64(julianDayNo)

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
*/

	fmt.Printf("setDto original Fraction:           %80.70f\n",
		julianDayNoTimeFraction)


	fmt.Printf("setDto jDNDto.julianDayNoFraction:  %80.70f\n",
		jDNDto.julianDayNoFraction)

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

	fmt.Printf("setDto bigDayNanoSeconds:      %80.70f\n",
		bigDayNanoSeconds)

	grossNanoSecs := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		Mul(bigDayNanoSeconds, jDNDto.julianDayNoFraction)

	fmt.Printf("setDto grossNanoSecs 1:        %80.70f\n",
		grossNanoSecs)

	grossNanoSecs.
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		Add(big.NewFloat(0.5), grossNanoSecs)

	fmt.Printf("setDto grossNanoSecs 2:        %80.70f\n",
		grossNanoSecs)

	jDNDto.totalNanoSeconds, _ = grossNanoSecs.Int64()

	fmt.Printf("setDto jDNDto.totalNanoSeconds:%v\n",
		jDNDto.totalNanoSeconds)

	timeMech := timeMechanics{}
	jDNDto.hours,
		jDNDto.minutes,
		jDNDto.seconds,
		jDNDto.nanoseconds,
		_ = timeMech.computeTimeElementsInt(
					jDNDto.totalNanoSeconds)

	fmt.Printf("setDto julianDayNoNanoSecs:         %80.70f\n",
		jDNDto.julianDayNoNanoSecs)

	return  err
}

