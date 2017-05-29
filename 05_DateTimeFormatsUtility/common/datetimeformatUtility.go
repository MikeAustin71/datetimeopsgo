package common

import (
	"errors"
	"time"
)

type DateTimeFormatGenerator struct {
	DayOfWeek          string
	DayOfWeekSeparator string
	MthDayYear         string
	DateTimeSeparator  string
	TimeElement        string
	OffsetSeparator    string
	OffsetElement      string
	TimeZoneSeparator  string
	TimeZoneElement    string
}

type DateTimeFormatUtility struct {
	OriginalDateTimeStringIn  string
	FormattedDateTimeStringIn string
	FormatMap                 map[int]map[string]int
	SelectedMapIdx            int
	SelectedFormat            string
	SelectedFormatSource      string
	DictSearches              map[int]int
	TotalNoOfDictSearches     int
	DateTimeOut               time.Time
	NumOfFormatsGenerated     int
}

func (dtf *DateTimeFormatUtility) GetAllDateTimeFormats() (err error) {

	dtf.FormatMap = make(map[int]map[string]int)
	dtf.NumOfFormatsGenerated = 0
	dtf.assemblePreDefinedFormats()
	dtf.assembleMthDayYearFmts()

	return
}

func (dtf *DateTimeFormatUtility) Empty() {
	dtf.OriginalDateTimeStringIn = ""
	dtf.FormattedDateTimeStringIn = ""
	dtf.DateTimeOut = time.Time{}
	dtf.SelectedFormat = ""
	dtf.SelectedFormatSource = ""
	dtf.SelectedMapIdx = -1
	dtf.DictSearches = make(map[int]int)
	dtf.TotalNoOfDictSearches = 0

}

func (dtf *DateTimeFormatUtility) ParseDateTimeString(timeStr string, probableFormat string) (time.Time, error) {

	if timeStr == "" {
		return time.Time{}, errors.New("Empty Time String!")
	}

	dtf.Empty()

	ftimeStr, err := StringUtility{}.TrimEndMultiple(timeStr, ' ')

	if err != nil {
		return time.Time{}, err
	}

	if probableFormat != "" {
		t, err := time.Parse(probableFormat, ftimeStr)

		if err == nil {
			dtf.SelectedFormat = probableFormat
			dtf.SelectedFormatSource = "User Provided"
			dtf.SelectedMapIdx = -1
			dtf.DateTimeOut = t
			dtf.OriginalDateTimeStringIn = timeStr
			dtf.FormattedDateTimeStringIn = ftimeStr
			dtf.TotalNoOfDictSearches = 1
			dtf.DictSearches[0] = 1

			return t, nil
		}

	}

	lenStr := len(ftimeStr)

	lenTests := [7]int{lenStr - 1, lenStr, lenStr - 2, lenStr + 1, lenStr + 2, lenStr - 3, lenStr + 3}

	for _, lTest := range lenTests {

		result, err := dtf.parseFormatMap(ftimeStr, lTest)

		dtf.DictSearches[lTest] = result.TotalNoOfDictSearches
		dtf.TotalNoOfDictSearches += result.TotalNoOfDictSearches

		if err == nil {
			dtf.SelectedFormat = result.SelectedFormat
			dtf.SelectedFormatSource = result.SelectedFormatSource
			dtf.SelectedMapIdx = result.SelectedMapIdx
			dtf.DateTimeOut = result.DateTimeOut
			dtf.OriginalDateTimeStringIn = timeStr
			dtf.FormattedDateTimeStringIn = result.FormattedDateTimeStringIn
			dtf.TotalNoOfDictSearches = result.TotalNoOfDictSearches
			return dtf.DateTimeOut, nil
		}

	}

	return time.Time{}, errors.New("Failed to locate correct time format!")
}

func (dtf *DateTimeFormatUtility) assembleDayMthYears() error {

	dtf.FormatMap = make(map[int]map[string]int)

	fmtStr := ""

	dayOfWeek, _ := dtf.getDayOfWeekElements()
	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()
	mthDayYearFmts, _ := dtf.getMonthDayYearElements()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {

			for _, mmddyyy := range mthDayYearFmts {

				if dowk != "" && mmddyyy != "" {
					fmtStr = dowk + dowkSep + mmddyyy
					dtf.assignFormatStrToMap(fmtStr)
				} else if dowk != "" && mmddyyy == "" {
					dtf.assignFormatStrToMap(dowk)
				} else if dowk == "" && mmddyyy != "" {
					dtf.assignFormatStrToMap(mmddyyy)
				}
			}
		}
	}

	return nil
}

func (dtf *DateTimeFormatUtility) parseFormatMap(timeStr string, idx int) (dtResult DateTimeFormatUtility, err error) {

	if dtf.FormatMap[idx] == nil {
		err = errors.New("Time String Length not found in Format Map!")
		return
	}

	for key := range dtf.FormatMap[idx] {

		dtResult.TotalNoOfDictSearches++

		t, err := time.Parse(key, timeStr)

		if err == nil {
			dtResult.OriginalDateTimeStringIn = timeStr
			dtResult.FormattedDateTimeStringIn = timeStr
			dtResult.SelectedMapIdx = idx
			dtResult.SelectedFormatSource = "Format Dictionary"
			dtResult.DateTimeOut = t
			dtResult.SelectedFormat = key
			return dtResult, err
		}

	}

	err = errors.New("Failed to parse time string")

	return
}

func (dtf *DateTimeFormatUtility) assembleMthDayYearFmts() error {

	dayOfWeek, _ := dtf.getDayOfWeekElements()

	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()

	mthDayYearFmts, _ := dtf.getMonthDayYearElements()

	dateTimeSeparators, _ := dtf.getDateTimeSeparators()

	timeFmts, _ := dtf.getTimeElements()

	offsetSeparators, _ := dtf.getTimeOffsetSeparators()

	offsetFmts, _ := dtf.getTimeOffsets()

	tzSeparators, _ := dtf.getTimeZoneSeparators()

	timeZoneFmts, _ := dtf.getTimeZoneElements()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {
			for _, mmddyyyy := range mthDayYearFmts {
				for _, dtSep := range dateTimeSeparators {
					for _, t := range timeFmts {
						for _, tOffsetSep := range offsetSeparators {
							for _, offFmt := range offsetFmts {
								for _, stdSep := range tzSeparators {
									for _, tzF := range timeZoneFmts {
										fmtGen := DateTimeFormatGenerator{
											DayOfWeek:          dowk,
											DayOfWeekSeparator: dowkSep,
											MthDayYear:         mmddyyyy,
											DateTimeSeparator:  dtSep,
											TimeElement:        t,
											OffsetSeparator:    tOffsetSep,
											OffsetElement:      offFmt,
											TimeZoneSeparator:  stdSep,
											TimeZoneElement:    tzF,
										}

										dtf.analyzeDofWeekMMDDYYYYTimeOffsetTz(fmtGen)
										dtf.analyzeDofWeekMMDDYYYYTzTimeOffset(fmtGen)
									}
								}

							}
						}
					}
				}

			}

		}

	}

	return nil
}

func (dtf *DateTimeFormatUtility) assembleMthDayTimeOffsetTzYearFmts() error {



	return nil
}

func (dtf *DateTimeFormatUtility) assemblePreDefinedFormats() {

	preDefFmts := dtf.getPredefinedFormats()

	for _, pdf := range preDefFmts {

		dtf.assignFormatStrToMap(pdf)

	}

}

func (dtf *DateTimeFormatUtility) analyzeDofWeekMMDDYYYYTimeOffsetTz(dtfGen DateTimeFormatGenerator) {

	fmtStr := ""

	if dtfGen.DayOfWeek != "" {
		fmtStr += dtfGen.DayOfWeek
	}

	if dtfGen.MthDayYear != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.MthDayYear
		} else {
			fmtStr += dtfGen.DayOfWeekSeparator
			fmtStr += dtfGen.MthDayYear
		}
	}

	if dtfGen.TimeElement != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.TimeElement
		} else {
			fmtStr += dtfGen.DateTimeSeparator
			fmtStr += dtfGen.TimeElement
		}
	}

	if dtfGen.OffsetElement != "" {
		if fmtStr == "" || dtfGen.TimeElement == "" {
			return
		} else {
			fmtStr += dtfGen.OffsetSeparator
			fmtStr += dtfGen.OffsetElement
		}
	}

	if dtfGen.TimeZoneElement != "" {
		if fmtStr == "" || dtfGen.TimeElement == "" {
			return
		} else {
			fmtStr += dtfGen.TimeZoneSeparator
			fmtStr += dtfGen.TimeZoneElement
		}

	}

	dtf.assignFormatStrToMap(fmtStr)

	return
}

func (dtf *DateTimeFormatUtility) analyzeDofWeekMMDDYYYYTzTimeOffset(dtfGen DateTimeFormatGenerator){

	fmtStr := ""

	if dtfGen.DayOfWeek != "" {
		fmtStr += dtfGen.DayOfWeek
	}

	if dtfGen.MthDayYear != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.MthDayYear
		} else {
			fmtStr += dtfGen.DayOfWeekSeparator
			fmtStr += dtfGen.MthDayYear
		}
	}

	if dtfGen.TimeElement != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.TimeElement
		} else {
			fmtStr += dtfGen.DateTimeSeparator
			fmtStr += dtfGen.TimeElement
		}
	}

	if dtfGen.TimeZoneElement != "" {
		if fmtStr == "" || dtfGen.TimeElement == "" {
			return
		} else {
			fmtStr += dtfGen.TimeZoneSeparator
			fmtStr += dtfGen.TimeZoneElement
		}

	}

	if dtfGen.OffsetElement != "" {
		if fmtStr == "" || dtfGen.TimeElement == "" {
			return
		} else {
			fmtStr += dtfGen.OffsetSeparator
			fmtStr += dtfGen.OffsetElement
		}
	}


	dtf.assignFormatStrToMap(fmtStr)

	return

}

func (dtf *DateTimeFormatUtility) assignFormatStrToMap(fmtStr string) {

	l := len(fmtStr)

	if l == 0 {
		return
	}

	if dtf.FormatMap[l] == nil {
		dtf.FormatMap[l] = make(map[string]int)
	}

	if dtf.FormatMap[l][fmtStr] != 0 {
		return
	}

	dtf.FormatMap[l][fmtStr] = l
	dtf.NumOfFormatsGenerated++
}

func (dtf DateTimeFormatUtility) getDayOfWeekElements() ([]string, error) {
	dayOfWeek := make([]string, 0, 10)

	dayOfWeek = append(dayOfWeek, "")
	dayOfWeek = append(dayOfWeek, "Mon")
	dayOfWeek = append(dayOfWeek, "Monday")

	return dayOfWeek, nil
}

func (dtf DateTimeFormatUtility) getDayOfWeekSeparator() ([]string, error) {
	dayOfWeekSeparator := make([]string, 0, 1024)

	dayOfWeekSeparator = append(dayOfWeekSeparator, " ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, ", ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, " - ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, "-")

	return dayOfWeekSeparator, nil
}

func (dtf DateTimeFormatUtility) getMonthDayYearElements() ([]string, error) {
	mthDayYr := make([]string, 0, 1024)

	mthDayYr = append(mthDayYr, "2006-01-02")
	mthDayYr = append(mthDayYr, "2006/01/02")
	mthDayYr = append(mthDayYr, "2006-1-2")
	mthDayYr = append(mthDayYr, "2006/1/2")

	mthDayYr = append(mthDayYr, "2006-1-02")
	mthDayYr = append(mthDayYr, "2006/1/02")
	mthDayYr = append(mthDayYr, "2006-01-2")
	mthDayYr = append(mthDayYr, "2006/01/2")

	// European Date Formats
	mthDayYr = append(mthDayYr, "02.01.06")
	mthDayYr = append(mthDayYr, "02.01.2006")
	mthDayYr = append(mthDayYr, "02.01.'06")

	mthDayYr = append(mthDayYr, "2.1.06")
	mthDayYr = append(mthDayYr, "2.1.2006")
	mthDayYr = append(mthDayYr, "2.1.'06")

	mthDayYr = append(mthDayYr, "2.01.06")
	mthDayYr = append(mthDayYr, "2.01.2006")
	mthDayYr = append(mthDayYr, "2.01.'06")

	mthDayYr = append(mthDayYr, "02.1.06")
	mthDayYr = append(mthDayYr, "02.1.2006")
	mthDayYr = append(mthDayYr, "02.1.'06")


	mthDayYr = append(mthDayYr, "02.January.'06")
	mthDayYr = append(mthDayYr, "02.January.06")
	mthDayYr = append(mthDayYr, "02.January.2006")

	mthDayYr = append(mthDayYr, "02.Jan.'06")
	mthDayYr = append(mthDayYr, "02.Jan.06")
	mthDayYr = append(mthDayYr, "02.Jan.2006")
	// ----------------------------------------------

	mthDayYr = append(mthDayYr, "01-02-06")
	mthDayYr = append(mthDayYr, "01-02-2006")
	mthDayYr = append(mthDayYr, "1-2-06")
	mthDayYr = append(mthDayYr, "1-2-2006")
	mthDayYr = append(mthDayYr, "1-02-06")
	mthDayYr = append(mthDayYr, "1-02-2006")
	mthDayYr = append(mthDayYr, "01/02/06")
	mthDayYr = append(mthDayYr, "01/02/2006")
	mthDayYr = append(mthDayYr, "1/2/06")
	mthDayYr = append(mthDayYr, "1/02/06")
	mthDayYr = append(mthDayYr, "1/2/2006")
	mthDayYr = append(mthDayYr, "1/02/2006")

	mthDayYr = append(mthDayYr, "Jan-02-06")
	mthDayYr = append(mthDayYr, "Jan 02 06")
	mthDayYr = append(mthDayYr, "Jan 02, 06")
	mthDayYr = append(mthDayYr, "Jan/02/06")
	mthDayYr = append(mthDayYr, "Jan-2-06")
	mthDayYr = append(mthDayYr, "Jan 2 06")
	mthDayYr = append(mthDayYr, "Jan 2, 06")
	mthDayYr = append(mthDayYr, "Jan/2/06")
	mthDayYr = append(mthDayYr, "Jan 02 2006")
	mthDayYr = append(mthDayYr, "Jan 2 2006")
	mthDayYr = append(mthDayYr, "Jan-2-2006")
	mthDayYr = append(mthDayYr, "Jan 02, 2006")
	mthDayYr = append(mthDayYr, "Jan 2, 2006")

	mthDayYr = append(mthDayYr, "Jan _2, 2006")
	mthDayYr = append(mthDayYr, "Jan _2 2006")
	mthDayYr = append(mthDayYr, "Jan _2 06")
	mthDayYr = append(mthDayYr, "Jan _2, 06")

	mthDayYr = append(mthDayYr, "January 02, 06")
	mthDayYr = append(mthDayYr, "January 02, 2006")
	mthDayYr = append(mthDayYr, "January 02 06")
	mthDayYr = append(mthDayYr, "January 02 2006")
	mthDayYr = append(mthDayYr, "January-02-2006")
	mthDayYr = append(mthDayYr, "January-02-06")
	mthDayYr = append(mthDayYr, "January 2, 06")
	mthDayYr = append(mthDayYr, "January 2, 2006")
	mthDayYr = append(mthDayYr, "January _2, 2006")
	mthDayYr = append(mthDayYr, "January _2, 06")

	mthDayYr = append(mthDayYr, "2 January, 06")
	mthDayYr = append(mthDayYr, "02 January, 06")
	mthDayYr = append(mthDayYr, "2 January, 2006")
	mthDayYr = append(mthDayYr, "02 January, 2006")
	mthDayYr = append(mthDayYr, "2 January 06")
	mthDayYr = append(mthDayYr, "02 January 06")
	mthDayYr = append(mthDayYr, "_2 January 06")
	mthDayYr = append(mthDayYr, "2 January 2006")
	mthDayYr = append(mthDayYr, "02 January 2006")
	mthDayYr = append(mthDayYr, "_2 January 2006")
	mthDayYr = append(mthDayYr, "02/Jan/2006")
	mthDayYr = append(mthDayYr, "2/Jan/2006")
	mthDayYr = append(mthDayYr, "02/Jan/06")
	mthDayYr = append(mthDayYr, "2/Jan/06")
	mthDayYr = append(mthDayYr, "02 Jan 06")
	mthDayYr = append(mthDayYr, "02 Jan 2006")
	mthDayYr = append(mthDayYr, "2 Jan 06")
	mthDayYr = append(mthDayYr, "2 Jan 2006")

	mthDayYr = append(mthDayYr, "02 Jan, 06")
	mthDayYr = append(mthDayYr, "02 Jan, 2006")
	mthDayYr = append(mthDayYr, "2 Jan, 06")
	mthDayYr = append(mthDayYr, "2 Jan, 2006")

	// ? May cause problems
	mthDayYr = append(mthDayYr, "06-01-02")
	mthDayYr = append(mthDayYr, "06/01/02")
	mthDayYr = append(mthDayYr, "060102")
	mthDayYr = append(mthDayYr, "06-1-2")
	mthDayYr = append(mthDayYr, "06/1/2")

	mthDayYr = append(mthDayYr, "20060102")
	mthDayYr = append(mthDayYr, "01022006")
	mthDayYr = append(mthDayYr, "")

	return mthDayYr, nil
}

func (dtf DateTimeFormatUtility) getMonthDayElements() ([]string, error) {
	mthDayElements := make([]string, 0, 124)

	mthDayElements = append(mthDayElements, "Jan 2")
	mthDayElements = append(mthDayElements, "January 2")
	mthDayElements = append(mthDayElements, "Jan _2")
	mthDayElements = append(mthDayElements, "January _2")
	mthDayElements = append(mthDayElements, "01-02")
	mthDayElements = append(mthDayElements, "01-_2")
	mthDayElements = append(mthDayElements, "01/02")
	mthDayElements = append(mthDayElements, "01/_2")
	mthDayElements = append(mthDayElements, "1-2")
	mthDayElements = append(mthDayElements, "1/2")
	mthDayElements = append(mthDayElements, "01-2")
	mthDayElements = append(mthDayElements, "01/2")
	mthDayElements = append(mthDayElements, "1-02")
	mthDayElements = append(mthDayElements, "1/02")
	mthDayElements = append(mthDayElements, "0102")
	// European Format Day Month
	mthDayElements = append(mthDayElements, "02.01")
	mthDayElements = append(mthDayElements, "2.1")
	mthDayElements = append(mthDayElements, "02.1")
	mthDayElements = append(mthDayElements, "2.01")

	return mthDayElements, nil
}

func (dtf DateTimeFormatUtility) getYears() ([]string, error) {
	yearElements := make([]string, 0, 10)

	yearElements = append(yearElements, "2006")
	yearElements = append(yearElements, "06")
	yearElements = append(yearElements, "'06")

	return yearElements, nil
}

func (dtf DateTimeFormatUtility) getMthDayAfterSeparators() ([]string, error) {
	mthDayAfterSeparators := make([]string, 0, 10)

	mthDayAfterSeparators = append(mthDayAfterSeparators, " ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ", ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, "T")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ":")

	return mthDayAfterSeparators, nil

}

func (dtf DateTimeFormatUtility) getStandardSeparators() ([]string, error) {
	standardSeparators := make([]string, 0, 10)

	standardSeparators = append(standardSeparators, " ")

	return standardSeparators, nil
}

func (dtf DateTimeFormatUtility) getDateTimeSeparators() ([]string, error) {
	dtTimeSeparators := make([]string, 0, 10)

	dtTimeSeparators = append(dtTimeSeparators, " ")
	dtTimeSeparators = append(dtTimeSeparators, ":")
	dtTimeSeparators = append(dtTimeSeparators, "T")

	return dtTimeSeparators, nil
}

func (dtf DateTimeFormatUtility) getTimeElements() ([]string, error) {
	timeElements := make([]string, 0, 512)

	timeElements = append(timeElements, "15:04:05")
	timeElements = append(timeElements, "15:04")
	timeElements = append(timeElements, "15:04:05.000")
	timeElements = append(timeElements, "15:04:05.000000")
	timeElements = append(timeElements, "15:04:05.000000000")

	timeElements = append(timeElements, "15:4:5")
	timeElements = append(timeElements, "15:4")
	timeElements = append(timeElements, "15:4:5.000")
	timeElements = append(timeElements, "15:4:5.000000")
	timeElements = append(timeElements, "15:4:5.000000000")

	timeElements = append(timeElements, "03:04:05pm")
	timeElements = append(timeElements, "03:04:05p.m.")
	timeElements = append(timeElements, "03:04:05 pm")
	timeElements = append(timeElements, "03:04:05 p.m.")

	timeElements = append(timeElements, "03:04:05PM")
	timeElements = append(timeElements, "03:04:05P.M.")
	timeElements = append(timeElements, "03:04:05 PM")
	timeElements = append(timeElements, "03:04:05 P.M.")

	timeElements = append(timeElements, "3:4:5pm")
	timeElements = append(timeElements, "3:4:5p.m.")
	timeElements = append(timeElements, "3:4:5 pm")
	timeElements = append(timeElements, "3:4:5 p.m.")

	timeElements = append(timeElements, "3:4:5PM")
	timeElements = append(timeElements, "3:4:5P.M.")
	timeElements = append(timeElements, "3:4:5 PM")
	timeElements = append(timeElements, "3:4:5 P.M.")

	timeElements = append(timeElements, "03:04pm")
	timeElements = append(timeElements, "03:04p.m.")
	timeElements = append(timeElements, "03:04 pm")
	timeElements = append(timeElements, "03:04 p.m.")

	timeElements = append(timeElements, "03:04PM")
	timeElements = append(timeElements, "03:04P.M.")
	timeElements = append(timeElements, "03:04 PM")
	timeElements = append(timeElements, "03:04 P.M.")

	timeElements = append(timeElements, "3:4pm")
	timeElements = append(timeElements, "3:4p.m.")
	timeElements = append(timeElements, "3:4 pm")
	timeElements = append(timeElements, "3:4 p.m.")

	timeElements = append(timeElements, "3:4PM")
	timeElements = append(timeElements, "3:4P.M.")
	timeElements = append(timeElements, "3:4 PM")
	timeElements = append(timeElements, "3:4 P.M.")

	timeElements = append(timeElements, "")

	return timeElements, nil
}

func (dtf DateTimeFormatUtility) getTimeOffsets() ([]string, error) {
	timeOffsetElements := make([]string, 0, 20)

	timeOffsetElements = append(timeOffsetElements, "-0700")
	timeOffsetElements = append(timeOffsetElements, "-07:00")
	timeOffsetElements = append(timeOffsetElements, "-07")
	timeOffsetElements = append(timeOffsetElements, "Z0700")
	timeOffsetElements = append(timeOffsetElements, "Z07:00")
	timeOffsetElements = append(timeOffsetElements, "Z07")
	timeOffsetElements = append(timeOffsetElements, "")

	return timeOffsetElements, nil
}

func (dtf DateTimeFormatUtility) getTimeOffsetSeparators() ([]string, error) {
	timeOffsetSeparators := make([]string, 0, 20)

	timeOffsetSeparators = append(timeOffsetSeparators, " ")
	timeOffsetSeparators = append(timeOffsetSeparators, "-")

	return timeOffsetSeparators, nil
}

func (dtf DateTimeFormatUtility) getTimeZoneElements() ([]string, error) {
	tzElements := make([]string, 0, 20)

	tzElements = append(tzElements, "MST")
	tzElements = append(tzElements, "")

	return tzElements, nil
}

func (dtf *DateTimeFormatUtility) getTimeZoneSeparators() ([]string, error) {
	tzElements := make([]string, 0, 20)

	tzElements = append(tzElements, " ")
	tzElements = append(tzElements, "-")

	return tzElements, nil
}

func (dtf *DateTimeFormatUtility) getPredefinedFormats() []string {

	preDefinedFormats := make([]string, 0, 20)

	preDefinedFormats = append(preDefinedFormats, time.ANSIC)
	preDefinedFormats = append(preDefinedFormats, time.UnixDate)
	preDefinedFormats = append(preDefinedFormats, time.RubyDate)
	preDefinedFormats = append(preDefinedFormats, time.RFC822)
	preDefinedFormats = append(preDefinedFormats, time.RFC822Z)
	preDefinedFormats = append(preDefinedFormats, time.RFC850)
	preDefinedFormats = append(preDefinedFormats, time.RFC1123)
	preDefinedFormats = append(preDefinedFormats, time.RFC1123Z)
	preDefinedFormats = append(preDefinedFormats, time.RFC3339)
	preDefinedFormats = append(preDefinedFormats, time.RFC3339Nano)
	preDefinedFormats = append(preDefinedFormats, time.Kitchen)
	preDefinedFormats = append(preDefinedFormats, time.Stamp)
	preDefinedFormats = append(preDefinedFormats, time.StampMilli)
	preDefinedFormats = append(preDefinedFormats, time.StampMicro)
	preDefinedFormats = append(preDefinedFormats, time.StampNano)

	return preDefinedFormats
}
