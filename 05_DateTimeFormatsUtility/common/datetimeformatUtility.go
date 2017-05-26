package common

import (
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
	DateTimeStringIn string
	Formats          []string
	SelectedFormat   string
	DateTimeOut      time.Time
}

func (dtf *DateTimeFormatUtility) GetAllDateTimeFormats() (err error) {
	dtf.Formats = make([]string, 0, 500000)
	dtf.assemblePreDefinedFormats()
	dtf.assembleDayMthYearFmts()

	return
}

func (dtf *DateTimeFormatUtility) AssembleDayMthYears() error {

	dtf.Formats = make([]string, 0, 2048)
	fmtStr := ""

	dayOfWeek, _ := dtf.getDayOfWeekElements()
	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()
	mthDayYearFmts, _ := dtf.getMonthDayYearElements()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {

			for _, mmddyyy := range mthDayYearFmts {

				if dowk != "" && mmddyyy != "" {
					fmtStr = (dowk + dowkSep + mmddyyy)
					dtf.Formats = append(dtf.Formats, fmtStr)
				} else if dowk != "" && mmddyyy == "" {
					dtf.Formats = append(dtf.Formats, dowk)
				} else if dowk == "" && mmddyyy != "" {
					dtf.Formats = append(dtf.Formats, mmddyyy)
				}
			}
		}
	}

	return nil
}

func (dtf *DateTimeFormatUtility) assembleDayMthYearFmts() error {

	dayOfWeek, _ := dtf.getDayOfWeekElements()

	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()

	mthDayYearFmts, _ := dtf.getMonthDayYearElements()

	dateTimeSeparators, _ := dtf.getDateTimeSeparators()

	timeFmts, _ := dtf.getTimeElements()

	offsetSeparators, _ := dtf.getTimeOffsetSeparators()

	offsetFmts, _ := dtf.getTimeOffsets()

	stdSeparators, _ := dtf.getStandardSeparators()

	timeZoneFmts, _ := dtf.getTimeZoneElements()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {
			for _, mmddyyyy := range mthDayYearFmts {
				for _, dtSep := range dateTimeSeparators {
					for _, t := range timeFmts {
						for _, tOffsetSep := range offsetSeparators {
							for _, offFmt := range offsetFmts {
								for _, stdSep := range stdSeparators {
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

func (dtf *DateTimeFormatUtility) assemblePreDefinedFormats() {
	preDefFmts, _ := dtf.getPredefinedFormats()

	for _, pdf := range preDefFmts {
		dtf.Formats = append(dtf.Formats, pdf)
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
		if fmtStr == "" || dtfGen.TimeElement==""{
			return
		} else {
			fmtStr += dtfGen.OffsetSeparator
			fmtStr += dtfGen.OffsetElement
		}
	}

	if dtfGen.TimeZoneElement != "" {
		if fmtStr == "" || dtfGen.TimeElement=="" {
			return
		} else {
			fmtStr += dtfGen.TimeZoneSeparator
			fmtStr += dtfGen.TimeZoneElement
		}

	}

	dtf.Formats = append(dtf.Formats, fmtStr)

	return
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
	mthDayYr = append(mthDayYr, "2006.01.02")
	mthDayYr = append(mthDayYr, "2006-1-2")
	mthDayYr = append(mthDayYr, "2006/1/2")
	mthDayYr = append(mthDayYr, "2006.1.2")

	mthDayYr = append(mthDayYr, "02.01.06")
	mthDayYr = append(mthDayYr, "02.01.2006")
	mthDayYr = append(mthDayYr, "02.01.'06")

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

	mthDayYr = append(mthDayYr, "2 January, 06")
	mthDayYr = append(mthDayYr, "02 January, 06")
	mthDayYr = append(mthDayYr, "2 January, 2006")
	mthDayYr = append(mthDayYr, "02 January, 2006")
	mthDayYr = append(mthDayYr, "2 January 06")
	mthDayYr = append(mthDayYr, "02 January 06")
	mthDayYr = append(mthDayYr, "2 January 2006")
	mthDayYr = append(mthDayYr, "02 January 2006")
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
	mthDayElements = append(mthDayElements, "01-02")
	mthDayElements = append(mthDayElements, "01/02")
	mthDayElements = append(mthDayElements, "1-2")
	mthDayElements = append(mthDayElements, "1/2")
	mthDayElements = append(mthDayElements, "0102")

	return mthDayElements, nil
}

func (dtf DateTimeFormatUtility) getYears() ([]string, error) {
	yearElements := make([]string, 0, 10)

	yearElements = append(yearElements, "2006")
	yearElements = append(yearElements, "06")
	yearElements = append(yearElements, "'06")

	return yearElements, nil
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
	timeElements = append(timeElements, "15:04:05.000")
	timeElements = append(timeElements, "15:04:05.000000")
	timeElements = append(timeElements, "15:04:05.000000000")

	timeElements = append(timeElements, "15:4:5")
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

	return timeOffsetSeparators, nil
}

func (dtf DateTimeFormatUtility) getTimeZoneElements() ([]string, error) {
	tzElements := make([]string, 0, 20)

	tzElements = append(tzElements, "MST")
	tzElements = append(tzElements, "")

	return tzElements, nil
}

func (dtf DateTimeFormatUtility) getPredefinedFormats() ([]string, error) {
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

	return preDefinedFormats, nil
}
