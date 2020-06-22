package datetime


// timeDtoSigns - Helper structure used to
// track the sign of TimeDto Time elements
type timeDtoSigns struct {
	signYears                int
	signMonths               int
	signWeeks                int
	signWeekDays             int
	signDateDays             int
	signHours                int
	signMinutes              int
	signSeconds              int
	signMilliseconds         int
	signMicroseconds         int
	signNanoseconds          int
	signTotSubSecNanoseconds int
	signTotTimeNanoseconds   int
}

// new - creates and returns a new timeDtoSigns
// instance where all the sign values are set
// to +1.
func (tSigns timeDtoSigns) new() timeDtoSigns {

	tSgn := timeDtoSigns{}

	tSgn.setSignsToOne()

	return tSgn
}

// areAnySignsNegative - returns a boolean value signaling
// whether any of the sign values are negative
func (tSigns *timeDtoSigns) areAnySignsNegative() bool {

	if tSigns.signYears < 0 ||
		tSigns.signMonths < 0 ||
		tSigns.signWeeks < 0 ||
		tSigns.signWeekDays < 0 ||
		tSigns.signDateDays < 0 ||
		tSigns.signHours < 0 ||
		tSigns.signMinutes < 0 ||
		tSigns.signSeconds < 0 ||
		tSigns.signMilliseconds < 0 ||
		tSigns.signMicroseconds < 0 ||
		tSigns.signNanoseconds < 0 ||
		tSigns.signTotSubSecNanoseconds < 0 ||
		tSigns.signTotTimeNanoseconds < 0 {

		return true
	}

	return false
}

// applySignsToTimeDto - Receives a pointer to t TimeDto and
// applies stored sign values to the TimeDto data fields.
func (tSigns *timeDtoSigns) applySignsToTimeDto(tDto *TimeDto) {

	tDto.Years *= tSigns.signYears
	tDto.Months *= tSigns.signMonths
	tDto.Weeks *= tSigns.signWeeks
	tDto.WeekDays *= tSigns.signWeekDays
	tDto.DateDays *= tSigns.signDateDays
	tDto.Hours *= tSigns.signHours
	tDto.Minutes *= tSigns.signMinutes
	tDto.Seconds *= tSigns.signSeconds
	tDto.Milliseconds *= tSigns.signMilliseconds
	tDto.Microseconds *= tSigns.signMicroseconds
	tDto.Nanoseconds *= tSigns.signNanoseconds
	tDto.TotSubSecNanoseconds *= tSigns.signTotSubSecNanoseconds
	tDto.TotTimeNanoseconds *= int64(tSigns.signTotTimeNanoseconds)

}

func (tSigns *timeDtoSigns) captureTimeDtoSigns(tDto *TimeDto) {

	tSigns.setSignsToOne()

	if tDto.Years < 0 {
		tSigns.signYears = -1
	}

	if tDto.Months < 0 {
		tSigns.signMonths = -1
	}

	if tDto.Weeks < 0 {
		tSigns.signWeeks = -1
	}

	if tDto.WeekDays < 0 {
		tSigns.signWeekDays = -1
	}

	if tDto.DateDays < 0 {
		tSigns.signDateDays = -1
	}

	if tDto.Hours < 0 {
		tSigns.signHours = -1
	}

	if tDto.Minutes < 0 {
		tSigns.signMinutes = -1
	}

	if tDto.Seconds < 0 {
		tSigns.signSeconds = -1
	}

	if tDto.Milliseconds < 0 {
		tSigns.signMilliseconds = -1
	}

	if tDto.Microseconds < 0 {
		tSigns.signMicroseconds = -1
	}

	if tDto.Nanoseconds < 0 {
		tSigns.signNanoseconds = -1
	}

	if tDto.TotSubSecNanoseconds < 0 {
		tSigns.signTotSubSecNanoseconds = -1
	}

	if tDto.TotTimeNanoseconds < 0 {
		tSigns.signTotTimeNanoseconds = -1
	}

}

func (tSigns *timeDtoSigns) setSignsToOne() {

	tSigns.signYears = 1
	tSigns.signMonths = 1
	tSigns.signWeeks = 1
	tSigns.signWeekDays = 1
	tSigns.signDateDays = 1
	tSigns.signHours = 1
	tSigns.signMinutes = 1
	tSigns.signSeconds = 1
	tSigns.signMilliseconds = 1
	tSigns.signMicroseconds = 1
	tSigns.signNanoseconds = 1
	tSigns.signTotSubSecNanoseconds = 1
	tSigns.signTotTimeNanoseconds = 1

}

func (tSigns *timeDtoSigns) setSignsToMinusOne() {

	tSigns.signYears = -1
	tSigns.signMonths = -1
	tSigns.signWeeks = -1
	tSigns.signWeekDays = -1
	tSigns.signDateDays = -1
	tSigns.signHours = -1
	tSigns.signMinutes = -1
	tSigns.signSeconds = -1
	tSigns.signMilliseconds = -1
	tSigns.signMicroseconds = -1
	tSigns.signNanoseconds = -1
	tSigns.signTotSubSecNanoseconds = -1
	tSigns.signTotTimeNanoseconds = -1

}

