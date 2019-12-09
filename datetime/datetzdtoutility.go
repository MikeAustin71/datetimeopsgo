package datetime

import "time"

type dateTzDtoUtility struct {
 	input    string
 	output   string
 }

// copyOut - Returns a deep copy of input parameter
// 'dTz' which is a pointer to a type 'DateTzDto'.
func (dTzUtil dateTzDtoUtility) copyOut(
	dTz *DateTzDto) DateTzDto {

	 dtz2 := DateTzDto{}

	 dtz2.tagDescription = dTz.tagDescription
	 dtz2.timeComponents = dTz.timeComponents.CopyOut()
	 dtz2.dateTimeFmt = dTz.dateTimeFmt

	 if !dTz.dateTimeValue.IsZero() {
		 dtz2.dateTimeValue = dTz.dateTimeValue
		 dtz2.timeZone = dTz.timeZone.CopyOut()
	 } else {
		 dtz2.timeZone = TimeZoneDefDto{}
		 dtz2.dateTimeValue = time.Time{}
	 }

	 return dtz2
}