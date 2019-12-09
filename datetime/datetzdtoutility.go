package datetime

import "time"

type dateTzDtoUtility struct {
 	input    string
 	output   string
 }


// copyIn - Receives two parameters which are pointers
// to types DateTzDto. The method then copies all of
// the data field values from 'incomingDtz' into
// 'baseDtz'.
//
func (dTzUtil dateTzDtoUtility) copyIn(
	baseDtz ,
	incomingDtz *DateTzDto) {

	dTzUtil.empty(baseDtz)

	baseDtz.tagDescription = incomingDtz.tagDescription
	baseDtz.timeComponents = incomingDtz.timeComponents.CopyOut()
	baseDtz.dateTimeFmt = incomingDtz.dateTimeFmt

	if !baseDtz.dateTimeValue.IsZero() {
		baseDtz.dateTimeValue = incomingDtz.dateTimeValue
		baseDtz.timeZone = incomingDtz.timeZone.CopyOut()
	} else {
		baseDtz.timeZone = TimeZoneDefDto{}
		baseDtz.dateTimeValue = time.Time{}
	}

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

// empty - Receives a pointer to a type 'DateTzDto' and
// proceeds to set all internal member variables to their
// 'zero' or uninitialized values.
//
func (dTzUtil dateTzDtoUtility) empty(dTz *DateTzDto) {

	dTz.tagDescription = ""
	dTz.timeComponents.Empty()
	dTz.timeZone = TimeZoneDefDto{}
	dTz.dateTimeValue = time.Time{}
	dTz.dateTimeFmt = ""

	return
}