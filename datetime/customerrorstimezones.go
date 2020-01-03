package datetime

import "fmt"

type TimeZoneMapLookupError struct {
	ePrefix string
	mapName string
	lookUpId string
	errMsg string
	err error
}


func (e *TimeZoneMapLookupError) Error() string{

	if len(e.errMsg) > 0 {

		return fmt.Sprintf(e.ePrefix +
			"\n%v\n" +
			"Map Name=%v\nLookUpId='%v'\n",
			e.errMsg, e.mapName, e.lookUpId)
	}

	return fmt.Sprintf(e.ePrefix +
		"\nTime Zone Abbreviation map Look Up Id Not Found!\n" +
		"Map Name=%v\nLookUpId='%v'\n",
		e.mapName, e.lookUpId)
}

func (e *TimeZoneMapLookupError) As(err error) bool {

	t, ok := err.(*TimeZoneMapLookupError)

	if !ok {
		return false
	}

	t.ePrefix = e.ePrefix
	t.mapName = e.mapName
	t.lookUpId = e.lookUpId
	t.errMsg = e.errMsg
	t.err =  e.err

	return true
}

func (e *TimeZoneMapLookupError) Is(target error) bool {

	_, ok := target.(*TimeZoneMapLookupError)

	if !ok {
		return false
	}

	return true
}

func (e *TimeZoneMapLookupError) Unwrap() error {
	return e.err
}


type TimeZoneError struct {
	ePrefix string
	errMsg string
	err error
}

func (e *TimeZoneError) Error() string{
	return fmt.Sprintf(e.ePrefix +
		"\n%v\n", e.errMsg)
}

func (e *TimeZoneError) As(err error) bool {

	t, ok := err.(*TimeZoneError)

	if !ok {
		return false
	}

	t.ePrefix = e.ePrefix
	t.errMsg = e.errMsg
	t.err =  e.err

	return true
}

func (e *TimeZoneError) Is(target error) bool {

	_, ok := target.(*TimeZoneError)

	if !ok {
		return false
	}

	return true
}

func (e *TimeZoneError) Unwrap() error {
	return e.err
}

