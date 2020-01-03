package datetime

import "fmt"

type TzAbbrvMapLookupError struct {
	ePrefix string
	mapName string
	lookUpId string
	errMsg string
	err error
}


func (e *TzAbbrvMapLookupError) Error() string{

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

func (e *TzAbbrvMapLookupError) As(err error) bool {

	t, ok := err.(*TzAbbrvMapLookupError)

	if !ok {
		return false
	}

	t.ePrefix = e.ePrefix
	t.mapName = e.mapName
	t.lookUpId = e.lookUpId
	t.errMsg = e.errMsg
	t.err = e.err

	return true
}

func (e *TzAbbrvMapLookupError) Is(target error) bool {

	_, ok := target.(*TzAbbrvMapLookupError)

	if !ok {
		return false
	}

	return true
}

func (e *TzAbbrvMapLookupError) Unwrap() error {
	return e.err
}


type TzAbbrvError struct {
	ePrefix string
	errMsg string
	err error
}

func (e *TzAbbrvError) Error() string{
	return fmt.Sprintf(e.ePrefix +
		"\n%v\n", e.errMsg)
}

func (e *TzAbbrvError) As(err error) bool {

	t, ok := err.(*TzAbbrvError)

	if !ok {
		return false
	}

	t.ePrefix = e.ePrefix
	t.errMsg = e.errMsg
	t.err =  e.err

	return true
}

func (e *TzAbbrvError) Is(target error) bool {

	_, ok := target.(*TzAbbrvError)

	if !ok {
		return false
	}

	return true
}

func (e *TzAbbrvError) Unwrap() error {
	return e.err
}
