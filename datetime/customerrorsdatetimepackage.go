package datetime

import "fmt"

type InputParameterError struct {
	ePrefix             string // Contains a chain of called methods leading to error
	inputParameterName  string // Name of invalid input parameter
	inputParameterValue string // The value of the invalid input parameter
	errMsg              string // Error message
	err                 error  // Next error in error chain
}

func (e *InputParameterError) Error() string{

	if len(e.inputParameterValue) > 0 &&
		len(e.errMsg) > 0 {
		return fmt.Sprintf(e.ePrefix + "\n" +
			"Error: Input Parameter, '%v' is invalid!\n" +
			"%v='%v'\n" +
			"%v",
			e.inputParameterName,
			e.inputParameterName,
			e.inputParameterValue,
		  e.errMsg)
	}

	if len(e.inputParameterValue) > 0 &&
		len(e.errMsg) == 0 {
		return fmt.Sprintf(e.ePrefix + "\n" +
			"Error: Input Parameter, '%v' is invalid!\n" +
			"%v='%v'\n",
			e.inputParameterName,
			e.inputParameterName,
			e.inputParameterValue)
	}

	return fmt.Sprintf(e.ePrefix + "\n" +
		"\nError: Input Parameter, '%v' is invalid!\n" +
		"%v\n",
		e.inputParameterName, e.errMsg)
}

func (e *InputParameterError) As(err error) bool {

	t, ok := err.(*InputParameterError)

	if !ok {
		return false
	}

	t.ePrefix = e.ePrefix
	t.inputParameterName = e.inputParameterName
	t.errMsg = e.errMsg
	t.err =  e.err

	return true
}

func (e *InputParameterError) Is(target error) bool {

	_, ok := target.(*InputParameterError)

	if !ok {
		return false
	}

	return true
}

func (e *InputParameterError) Unwrap() error {
	return e.err
}
