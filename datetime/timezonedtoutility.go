package datetime

import (
	"errors"
	"fmt"
	"sync"
)

type typeZoneUtility struct {
	lock        sync.Mutex
}

// isValidTimeZoneDto - Analyzes the specified 'TimeZoneDto'
// instance and returns an error if the instance is INVALID.
//
func (tZoneUtil *typeZoneUtility) isValidTimeZoneDto(
	tzDto *TimeZoneDto,
	ePrefix string) error {

	tZoneUtil.lock.Lock()

	defer tZoneUtil.lock.Unlock()

	ePrefix += "typeZoneUtility.isValidTimeZoneDto() "

	if tzDto == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter 'tzDto' is a 'nil' pointer!\n")
	}

	errorArray := make([]error, 0)

	if err := tzDto.TimeIn.IsValid(); err != nil {
		errorArray =
			append(errorArray, fmt.Errorf(ePrefix+"\ntzDto.TimeIn is INVALID!\nError='%v'\n", err.Error()))
	}

	if err := tzDto.TimeOut.IsValid(); err != nil {
		errorArray =
			append(errorArray, fmt.Errorf(ePrefix+"\nError: TimeOut is INVALID!\nError='%v'\n", err.Error()))
	}

	if err := tzDto.TimeUTC.IsValid(); err != nil {
		errorArray =
			append(errorArray, fmt.Errorf(ePrefix+"\nError: TimeUTC is INVALID!\nError='%v'\n", err.Error()))
	}

	if err := tzDto.TimeLocal.IsValid(); err != nil {
		errorArray =
			append(errorArray, fmt.Errorf(ePrefix+"Error:\nTimeLocal is INVALID!\nError='%v'\n", err.Error()))
	}

	dtUtil := DTimeUtility{}

	return dtUtil.ConsolidateErrors(errorArray)
}