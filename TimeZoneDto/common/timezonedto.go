package common

import (
	"errors"
	"fmt"
	"time"
	"strings"

)

/*
	Overview - Location
	===================

  timezonedto.go is part of the date time operations library. The source code repository
 	for this file is located at:

					https://github.com/MikeAustin71/datetimeopsgo.git



	Dependencies
	============

	None

 */

// NOTE: See https://golang.org/pkg/time/#LoadLocation
// and https://www.iana.org/time-zones to ensure that
// the IANA Time Zone Database is properly configured
// on your system. Note: IANA Time Zone Data base is
// equivalent to 'tz database'.
//
// Reference: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
const (

	TzIanaAfricaCairo = "Africa/Cairo"
	TzIanaAfricaJohannesburg = "Africa/Johannesburg"
	TzIanaAfricaTripoli = "Africa/Tripoli"
	TzIanaAfricaTunis = "Africa/Tunis"
	TzIanaAmericaJamaica = "America/Jamaica"
	TzIanaAmericaMatamoros = "America/Matamoros"
	TzIanaAmericaMexicoCity = "America/Mexico_City"
	TzIanaAmericaPuertoRico = "America/Puerto_Rico"
	TzIanaAmericaTijuana = "America/Tijuana"
	TzIanaAmericaBuenosAires = "America/Argentina/Buenos_Aires"
	TzIanaAmericaBogota = "America/Bogota"
	TzIanaAmericaCancum = "America/Cancun"
	TzIanaAmericaCaracas = "America/Caracas"
	TzIanaAmericaCostaRica = "America/Costa_Rica"
	TzIanaAmericaElSalvador = "America/El_Salvador"
	TzIanaAmericaGooseBay = "America/Goose_Bay"
	TzIanaAmericaPortOfSpain ="America/Port_of_Spain" // Grenada
	TzIanaAmericaGuatemala = "America/Guatemala"
	TzIanaAmericaGuayaquil = "America/Guayaquil" // Ecuador
	TzIanaAmericaGuyana = "America/Guyana"
	TzIanaAmericaHalifax = "America/Halifax"
	TzIanaAmericaHavana = "America/Havana"
	TzIanaAmericaLaPaz = "America/La_Paz"
	TzIanaAmericaLima = "America/Lima"
	TzIanaAmericaManaus = "America/Manaus"  // Amazonas East
	TzIanaAmericaMartinique = "America/Martinique"
	TzIanaAmericaMazatlan = "America/Mazatlan" // Baja California
	TzIanaAmericaMotevideo = "America/Montevideo" //Uruguay
	TzIanaAmericaNassau = "America/Nassau" // Bahamas
	TzIanaAmericaPanama = "America/Panama"
	TzIanaAmericaRecife = "America/Recife"
	TzIanaAmericaSantiago = "America/Santiago"
	TzIanaAmericaSantoDomingo = "America/Santo_Domingo"
	TzIanaAmericaSaoPaulo = "America/Sao_Paulo"
	TzIanaAmericaStJohns = "America/St_Johns" // Newfoundland Labrador
	TzIanaAmericaStThomas = "America/St_Thomas"
	TzIanaAmericaThule = "America/Thule"
	TzIanaAmericaToronto = "America/Toronto" // Eastern - ON, QC (most areas)
	TzIanaAmericaVancouver = "America/Vancouver" // Pacific - BC (most areas)
	TzIanaAmericaWinnipeg = "America/Winnipeg" // Central - ON (west); Manitoba
	TzIanaAmericaWhitehorse = "America/Whitehorse" // Pacific - Yukon (south)
	TzIanaAntarcticaMcMurdo = "Antarctica/McMurdo"
	TzIanaAntarcticaSoutPole = "Pacific/Auckland"
	TzIanaAsiaBankok = "Asia/Bangkok"
	TzIanaAsiaBaghdad = "Asia/Baghdad"
	TzIanaAsiaBahrain = "Asia/Bahrain"
	TzIanaAsiaBaku = "Asia/Baku"
	TzIanaAsia = "Asia/Brunei"
	TzIanaAsiaBeirut = "Asia/Beirut"
	TzIanaAsiaDamasucs = "Asia/Damascus"
	TzIanaAsiaDubai = "Asia/Dubai"
	TzIanaAsiaHoChiMinh = "Asia/Ho_Chi_Minh"			// Saigon Vietnam
	TzIanaAsiaHongKong = "Asia/Hong_Kong"
	TzIanaAsiaIndia = "Asia/Kolkata"						// Formerly Calcutta  - India Time
	TzIanaAsiaJakarta = "Asia/Jakarta"
	TzIanaAsiaJerusalem = "Asia/Jerusalem"
	TzIanaAsiaKabul = "Asia/Kabul"
	TzIanaAsiaKarachi = "Asia/Karachi"
	TzIanaAsiaKualaLumpur = "Asia/Kuala_Lumpur"
	TzIanaAsiaKuwait = "Asia/Kuwait"
	TzIanaAsiaManila = "Asia/Manila"
	TzIanaAsiaPhnomPenh = "Asia/Phnom_Penh"
	TzIanaAsiaPyongyang = "Asia/Pyongyang"
	TzIanaAsiaQatar = "Asia/Qatar"
	TzIanaAsiaRangoon = "Asia/Yangon" 							// Rangoon
	TzIanaAsiaRiyadh = "Asia/Riyadh"
	TzIanaAsiaSeoul = "Asia/Seoul"
	TzIanaAsiaShanghai = "Asia/Shanghai"						// Beijing time
	TzIanaAsiaSigapore = "Asia/Singapore"
	TzIanaAsiaTaipei = "Asia/Taipei"
	TzIanaAsiaTehran = "Asia/Tehran"
	TzIanaAsiaTokyo = "Asia/Tokyo"
	TzIanaAsiaVladivostok = "Asia/Vladivostok"
	TzIanaAtlanticAzores = "Atlantic/Azores"
	TzIanaAtlanticBermuda = "Atlantic/Bermuda"
	TzIanaAtlanticCanary = "Atlantic/Canary"
	TzIanaAtlanticCapeVerde = "Atlantic/Cape_Verde"
	TzIanaAtlanticReykjavik = "Atlantic/Reykjavik"
	TzIanaAtlanticStanley = "Atlantic/Stanley"
	TzIanaAustraliaDarwin = "Australia/Darwin"
	TzIanaAustraliaMelbourne = "Australia/Melbourne"
	TzIanaAustraliaSydney = "Australia/Sydney"
	TzIanaAustraliaPerth = "Australia/Perth"
	TzIanaEuropeAmsterdam = "Europe/Amsterdam"
	TzIanaEuropeAthens = "Europe/Athens"
	TzIanaEuropeBelgrade = "Europe/Belgrade"
	TzIanaEuropeBerlin = "Europe/Berlin"
	TzIanaEuropeBrussels = "Europe/Brussels"
	TzIanaEuropeBucharest = "Europe/Bucharest"
	TzIanaEuropeBudapest = "Europe/Budapest"
	TzIanaEuropeCopenhagen = "Europe/Copenhagen"
	TzIanaEuropeDublin = "Europe/Dublin"
	TzIanaEuropeGibraltar = "Europe/Gibraltar"
	TzIanaEuropeHelsinki = "Europe/Helsinki"
	TzIanaEuropeIstanbul = "Europe/Istanbul"
	TzIanaEuropeKiev = "Europe/Kiev"
	TzIanaEuropeLisbon = "Europe/Lisbon"
	TzIanaEuropeLondon = "Europe/London"
	TzIanaEuropeLuxembourg = "Europe/Luxembourg"
	TzIanaEuropeMadrid = "Europe/Madrid"
	TzIanaEuropeMalta = "Europe/Malta"
	TzIanaEuropeMinsk = "Europe/Minsk"
	TzIanaEuropeMonaco = "Europe/Monaco"
	TzIanaEuropeMoscow = "Europe/Moscow"
	TzIanaEuropeOslo = "Europe/Oslo"
	TzIanaEuropeParis = "Europe/Paris"
	TzIanaEuropePrague = "Europe/Prague"
	TzIanaEuropeRiga = "Europe/Riga"
	TzIanaEuropeRome = "Europe/Rome"
	TzIanaEuropeSofia = "Europe/Sofia"
	TzIanaEuropeStockholm = "Europe/Stockholm"
	TzIanaEuropeVienna = "Europe/Vienna"
	TzIanaEuropeVilnius = "Europe/Vilnius"
	TzIanaEuropeWarsaw = "Europe/Warsaw"
	TzIanaEuropeZurich ="Europe/Zurich"
	TzIanaPacificAuckland = "Pacific/Auckland"
	TzIanaPacificFiji = "Pacific/Fiji"
	TzIanaPacificGuam = "Pacific/Guam"
	TzIanaPacificHonolulu = "Pacific/Honolulu"
	TzIanaPacificPortMoresby = "Pacific/Port_Moresby"
	TzIanaPacificTahiti = "Pacific/Tahiti"

	// TzIanaUsAlaska - USA Alaska
	TzIanaUsAlaskaAnchorage =  "America/Anchorage"
	TzIanaUsAlaskaJuneau = "America/Juneau"
	TzIanaUsAlaskaNome = "America/Nome"
	TzIanaUsAlaskaYakutat = "America/Yakutat"
	TzIanaUsAleutianIslands = "America/Adak"

	// TzIanaUsArizona
	TzIanaUsArizona = "America/Phoenix"

	// TzIanaUsEast - USA Eastern Time Zone
	// IANA database identifier
	TzIanaUsEast = "America/New_York"

	// TzIanaUsCentral - USA Central Time Zone
	// IANA database identifier
	TzIanaUsCentral = "America/Chicago"

	// TzIanaUsMountain - USA Mountain Time Zone
	// IANA database identifier
	TzIanaUsMountain = "America/Denver"

	// TzIanaUsPacific - USA Pacific Time Zone
	// IANA database identifier
	TzIanaUsPacific  = "America/Los_Angeles"

	// TzIanaUsHawaii - USA Hawaiian Time Zone
	// IANA database identifier
	TzIanaUsHawaii = "Pacific/Honolulu"

	// TzIanaZulu - UTC Time Zone IANA database
	// identifier
	TzIanaZulu = "Etc/UCT"

	// TzIanaUTC
	TzIanaUTC = "Etc/UCT"


)


const (
	// TzDtoNeutralDateFmt - Neutral Date Format without Time Zone
	TzDtoNeutralDateFmt = "2006-01-02 15:04:05.000000000"

	// TzDtoMDYrFmtStr - Month Day Year Date Format String
	TzDtoMDYrFmtStr = "01/02/2006 15:04:05.000000000 -0700 MST"

	// TzDtoYrMDayFmtStr - Year Month Day Date Format String
	TzDtoYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
)

// TimeZoneDefDto - Time Zone Definition Dto
// Contains detailed parameters describing a specific
// Time Zone and Time Zone Location
type TimeZoneDefDto struct {
	ZoneName						string
	ZoneOffsetSeconds		int			// Signed number of seconds offset from UTC. + == East of UTC; - == West of UTC
	ZoneSign						int 		// -1 == West of UTC  +1 == East of UTC
	OffsetHours					int			// Hours offset from UTC. Always a positive number, refer to ZoneSign
	OffsetMinutes				int			// Minutes offset from UTC. Always a positive number, refer to ZoneSign
	OffsetSeconds				int			// Seconds offset from UTC. Always a positive number, refer to ZoneSign
	ZoneOffset					string	// A text string representing the time zone. Example "-0500 CDT"
	Location						*time.Location	// Pointer to a Time Zone Location
	LocationName				string					// Time Zone Location Name Examples: "Local", "America/Chicago", "America/New_York"
	Description					string	// Unused - Available for classification, labeling or description by user.
}

// CopyIn - Copies an incoming TimeZoneDefDto into the
// data fields of the current TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) CopyIn(tzdef2 *TimeZoneDefDto) {

	tzdef.Empty()

	tzdef.ZoneName					= tzdef2.ZoneName
	tzdef.ZoneOffsetSeconds	= tzdef2.ZoneOffsetSeconds
	tzdef.ZoneSign					= tzdef2.ZoneSign
	tzdef.OffsetHours				= tzdef2.OffsetHours
	tzdef.OffsetMinutes			= tzdef2.OffsetMinutes
	tzdef.OffsetSeconds			= tzdef2.OffsetSeconds
	tzdef.ZoneOffset				= tzdef2.ZoneOffset
	tzdef.Location	 				= tzdef2.Location
	tzdef.LocationName			= tzdef2.LocationName
	tzdef.Description				= tzdef2.Description

}

// CopyOut - creates and returns a deep copy of the current
// TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) CopyOut() TimeZoneDefDto {

	tzdef2 := TimeZoneDefDto{}

	tzdef2.ZoneName						= tzdef.ZoneName
	tzdef2.ZoneOffsetSeconds	= tzdef.ZoneOffsetSeconds
	tzdef2.ZoneSign						= tzdef.ZoneSign
	tzdef2.OffsetHours				= tzdef.OffsetHours
	tzdef2.OffsetMinutes			= tzdef.OffsetMinutes
	tzdef2.OffsetSeconds			= tzdef.OffsetSeconds
	tzdef2.ZoneOffset					= tzdef.ZoneOffset
	tzdef2.Location	  				= tzdef.Location
	tzdef2.LocationName				= tzdef.LocationName
	tzdef2.Description				= tzdef.Description

	return tzdef2

}


// Empty - Resets all field values for the current TimeZoneDefDto
// instance to their uninitialized or 'zero' states.
func (tzdef *TimeZoneDefDto) Empty() {
	tzdef.ZoneName						= ""
	tzdef.ZoneOffsetSeconds		= 0
	tzdef.ZoneSign						= 0
	tzdef.OffsetHours					= 0
	tzdef.OffsetMinutes				= 0
	tzdef.OffsetSeconds				= 0
	tzdef.ZoneOffset					= ""
	tzdef.Location						= nil
	tzdef.LocationName				= ""
	tzdef.Description					= ""

}

// Equal - Determines if two TimeZoneDefDto are equivalent in
// value. Returns 'true' of two TimeZoneDefDto's are equal
func(tzdef *TimeZoneDefDto) Equal(tzdef2 *TimeZoneDefDto) bool {

	if tzdef.ZoneName == tzdef2.ZoneName &&
			tzdef.ZoneOffsetSeconds == tzdef2.ZoneOffsetSeconds &&
			tzdef.ZoneSign == tzdef2.ZoneSign &&
			tzdef.OffsetHours == tzdef2.OffsetHours &&
			tzdef.OffsetMinutes == tzdef2.OffsetMinutes &&
			tzdef.OffsetSeconds == tzdef2.OffsetSeconds &&
			tzdef.ZoneOffset == tzdef2.ZoneOffset &&
			tzdef.Location.String() == tzdef2.Location.String() &&
			tzdef.LocationName == tzdef2.LocationName &&
			tzdef.Description == tzdef2.Description {
				return true
	}


	return false
}

// IsEmpty - Determines whether the current TimeZoneDefDto
// instance is Empty.
//
// If the TimeZoneDefDto instance is NOT populated, the method
// returns 'true'. Otherwise, it returns 'false'.
//
func (tzdef *TimeZoneDefDto) IsEmpty() bool {

	if tzdef.ZoneName 					!= "" ||
		 	tzdef.ZoneOffsetSeconds != 0 	||
			tzdef.ZoneSign 					!= 0 	||
			tzdef.OffsetHours 			!= 0 	||
			tzdef.OffsetMinutes 		!= 0 	||
			tzdef.OffsetSeconds 		!= 0	||
			tzdef.ZoneOffset				!= "" ||
			tzdef.LocationName			!= "" {
				return false
	}

	return true
}

// IsValid - Analyzes the current TimeZoneDefDto instance
// to determine validity.
//
// This method returns 'true' if the TimeZoneDefDto is
// valid.  Otherwise, it returns 'false'.
//
func (tzdef *TimeZoneDefDto) IsValid() bool {

	if tzdef.IsEmpty() {
		return false
	}

	if strings.TrimLeft(strings.TrimRight(tzdef.LocationName," "), " ")  == "" {
		return false
	}

	if tzdef.Location.String() != tzdef.LocationName {
		return false
	}

	loc, err := time.LoadLocation(tzdef.LocationName)

	if err != nil {
		return false
	}

	if loc != tzdef.Location {
		return false
	}

	return true
}

// IsValidFromDateTime - Uses a time.Time input parameter, 'dateTime' to
// analyze the current TimeZoneDefDto instance. If the zone and location
// details of 'dateTime' are not perfectly matched to the current TimeZoneDefDto
// instance, the instance is considered INVALID, and this method returns 'false'.
//
// Otherwise, if all zone and location details are perfectly matched, this
// method returns 'true', signaling that the TimeZoneDateDefDto instance
// is VALID.
//
func (tzdef *TimeZoneDefDto) IsValidFromDateTime(dateTime time.Time) bool {

	if tzdef.IsEmpty() {
		return false
	}

	tzdef2, err := TimeZoneDefDto{}.New(dateTime)

	if err != nil {
		return false
	}

	tzdef2.Description = tzdef.Description

	if !tzdef.Equal(&tzdef2) {
		return false
	}

	return true
}

// New - Creates and returns a new TimeZoneDefDto instance based on
// a 'dateTime (time.Time) input parameter.
//
// Input Parameter
// ===============
//
//	dateTime 	time.Time	- A date time value which will be used to construct the
//													elements of a Time Zone Definition Dto instance.
//
// Returns
// =======
//
// This method will return two Types:
//			(1) A Time Zone Definition Dto
//			(2) An 'error' type
//
// (1) If successful, this method will return a valid, populated TimeZoneDefDto instance.
//		 A TimeZoneDefDto is defined as follows:
//			type TimeZoneDefDto struct {
//				ZoneName						string
//				ZoneOffsetSeconds		int			// Signed number of seconds offset from UTC. + == East of UTC; - == West of UTC
//				ZoneSign						int 		// -1 == West of UTC  +1 == East of UTC
//				OffsetHours					int			// Hours offset from UTC. Always a positive number, refer to ZoneSign
//				OffsetMinutes				int			// Minutes offset from UTC. Always a positive number, refer to ZoneSign
//				OffsetSeconds				int			// Seconds offset from UTC. Always a positive number, refer to ZoneSign
//				ZoneOffset					string	// A text string representing the time zone. Example "-0500 CDT"
//				Location						*time.Location	// Pointer to a Time Zone Location
//				LocationName				string					// Time Zone Location Name Examples: "Local", "America/Chicago", "America/New_York"
//				Description					string	// Unused - Available for classification, labeling or description by user.
//			}
//
//
// (2) 	If successful, this method will set the returned error instance to 'nil.
//			If errors are encountered a valid error message will be returned in the
//			error instance.
//
func (tzdef TimeZoneDefDto) New(dateTime time.Time) (TimeZoneDefDto, error) {

	ePrefix := "TimeZoneDefDto.New() "

	if dateTime.IsZero() {
		return TimeZoneDefDto{}, errors.New(ePrefix + "Error: Input parameter 'dateTime' is a ZERO value!")
	}

	tzDef2 := TimeZoneDefDto{}

	err := tzDef2.SetFromDateTime(dateTime)

	if err!=nil {
		return TimeZoneDefDto{}, fmt.Errorf(ePrefix +
			"Error returned by tzDef2.SetFromDateTime(dateTime). " +
			"dateTime='%v'  Error='%v'",
			dateTime.Format(TzDtoYrMDayFmtStr), err.Error())
	}

	return tzDef2, nil
}

// SetFromDateTime - Re-initializes the values of the current
// TimeZoneDefDto instance based on input parameter, 'dateTime'.
func (tzdef *TimeZoneDefDto) SetFromDateTime(dateTime time.Time) error {
	ePrefix := "TimeZoneDefDto.SetFromDateTime() "

	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter 'dateTime' is a ZERO value!")
	}

	tzdef.Empty()

	tzdef.ZoneName, tzdef.ZoneOffsetSeconds = dateTime.Zone()

	tzdef.allocateZoneOffsetSeconds(tzdef.ZoneOffsetSeconds)

	tzdef.Location = dateTime.Location()

	tzdef.LocationName = dateTime.Location().String()

	tzdef.setZoneString()

	tzdef.Description = ""



	return nil
}


// allocateZoneOffsetSeconds - allocates a signed value of total offset seconds from
// UTC to the associated fields in the current TimeZoneDefDto instance.
func (tzdef *TimeZoneDefDto) allocateZoneOffsetSeconds(signedZoneOffsetSeconds int) {

	if signedZoneOffsetSeconds < 0 {
		tzdef.ZoneSign = -1
	} else {
		tzdef.ZoneSign = 1
	}

	tzdef.ZoneOffsetSeconds = signedZoneOffsetSeconds

	signedZoneOffsetSeconds *= tzdef.ZoneSign

	tzdef.OffsetHours = 0
	tzdef.OffsetMinutes = 0
	tzdef.OffsetSeconds = 0

	if signedZoneOffsetSeconds == 0 {
		return
	}

	tzdef.OffsetHours = signedZoneOffsetSeconds / 3600 // compute hours
	signedZoneOffsetSeconds -= tzdef.OffsetHours * 3600

	tzdef.OffsetMinutes = signedZoneOffsetSeconds / 60 // compute minutes
	signedZoneOffsetSeconds -= tzdef.OffsetMinutes * 60

	tzdef.OffsetSeconds = signedZoneOffsetSeconds

	return
}

// setZoneString - assembles and assigns the composite zone
// offset and zone name abbreviation in the TimeZoneDefDto.ZoneOffset
// field. Example: "-0500 CST"
func (tzdef *TimeZoneDefDto) setZoneString() {

	tzdef.ZoneOffset = ""

	if tzdef.ZoneSign < 0 {
		tzdef.ZoneOffset += "-"
	}

	tzdef.ZoneOffset += fmt.Sprintf("%02d%02d",tzdef.OffsetHours,tzdef.OffsetMinutes)

	if tzdef.OffsetSeconds > 0 {
		tzdef.ZoneOffset += fmt.Sprintf("%02%", tzdef.OffsetSeconds)
	}

	tzdef.ZoneOffset += " " + tzdef.ZoneName

	return
}


// DateTzDto - Type
// ================
// Used to store and transfer date times.
// The descriptors contained is this structure are intended
// to define and identify a specific point in time.
//
// This Type is NOT used to define duration; that is, the
// difference or time span between two point in time. For
// these types of operations see:
// DurationTimeUtility/common/durationutil.go
//
// DateTzDto defines a specific point in time using
// a variety of descriptors including year, month, day
// hour, minute, second, millisecond, microsecond and
// and nanosecond. In addition this Type specifies a
// time.Time value as well as time zone location and
// time zone.
//
// If you are unfamiliar with the concept of a time
// zone location, consider the field TimeLoc and
// TimeLocName below:
//
// Time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
type DateTzDto struct {
	Year       			int							// Year Number
	Month      			int							// Month Number
	Day        			int							// Day Number
	Hour       			int							// Hour Number
	Minute     			int							// Minute Number
	Second     			int							// Second Number
	Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
	Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
	Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
	TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
	DateTime 				time.Time				// DateTime value for this DateTzDto Type
	DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
	TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
																	// 		associated with this date time.
}

// New - returns a new DateTzDto instance based on a time.Time ('dateTime')
// input parameter.
//
// Input Parameter
// ===============
//
// dateTime   time.Time - a date time value
//
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Year       			int							// Year Number
//		Month      			int							// Month Number
//		Day        			int							// Day Number
//		Hour       			int							// Hour Number
//		Minute     			int							// Minute Number
//		Second     			int							// Second Number
//		Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//		Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//		Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																		// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//		TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//		TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//	}
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
// Usage
// =====
//
// Example:
//			dtzDto, err := DateTzDto{}.New(dateTime)
//
func (dtz DateTzDto) New(dateTime time.Time)(DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	if dateTime.IsZero() {
		return DateTzDto{}, errors.New(ePrefix + "Error: Input parameter dateTime is Zero value!")
	}

	dtz2 := DateTzDto{}

	err := dtz2.SetFromTime(dateTime)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned from dtz2.SetFromTime(dateTime). dateTime='%v'  Error='%v'", dateTime, err.Error())
	}

	return dtz2, nil
}

// NewDateTimeElements - creates a new DateTzDto object and populates the data fields based on
// input parameters.
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// minute						int			- minute number	0 - 59
// second						int			- second number	0	-	59
// nanosecond				int			- nanosecond number 0 - 999999999
//
// timeZoneLocation	string	- time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Year       			int							// Year Number
//		Month      			int							// Month Number
//		Day        			int							// Day Number
//		Hour       			int							// Hour Number
//		Minute     			int							// Minute Number
//		Second     			int							// Second Number
//		Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//		Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//		Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																		// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//		TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//		TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//	}
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
// Usage
// =====
//
// Example:
//	dtzDto, err := DateTzDto{}.NewDateTimeElements(year, month, day, hour, minute, second, nanosecond , timeZoneLocation)
//
//
func (dtz DateTzDto) NewDateTimeElements(year, month, day, hour, minute, second, nanosecond int, timeZoneLocation string) (DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	dtz2 := DateTzDto{}

	err := dtz2.SetFromDateTimeElements(year, month, day, hour, minute, second, nanosecond, timeZoneLocation)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned from dtz2.SetFromDateTimeElements(...) " +
			"year='%v' month='%v' day='%v' hour='%v' minute='%v' second='%v' nanosecond='%v' timeZoneLocatin='%v'  Error='%v'",
			year, month, day, hour, minute, second, nanosecond, timeZoneLocation, err.Error())
	}

	return dtz2, nil
}

// NewDateTime - creates a new DateTzDto object and populates the data fields based on
// input parameters.
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// minute						int			- minute number	0 - 59
// second						int			- second number	0	-	59
// millisecond			int			- millisecond number 0 - 999
// microsecond			int			-	microsecond number 0 - 999
// nanosecond				int			- nanosecond number 0 - 999
// timeZoneLocation	string	- time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// Returns
// =======
//
//  There are two return values: 	(1) a DateTzDto Type
//																(2) an Error type
//
//  DateTzDto - If successful the method returns a valid, fully populated
//							DateTzDto type defined as follows:
//
//	type DateTzDto struct {
//		Year       			int							// Year Number
//		Month      			int							// Month Number
//		Day        			int							// Day Number
//		Hour       			int							// Hour Number
//		Minute     			int							// Minute Number
//		Second     			int							// Second Number
//		Millisecond			int							// Number of MilliSeconds - A Millisecond is 1 one-thousandth or 1/1,000 of a second
//		Microsecond			int							// Number of MicroSeconds - A Microsecond is 1 one-millionth or 1/1,000,000 of a second
//		Nanosecond 			int							// Number of Nanoseconds - A Nanosecond is 1 one-billionth or 1/1,000,000,000 of a second.
//																		// Nanosecond = TotalNanoSecs - millisecond nonseconds - microsecond nanoseconds
//		TotalNanoSecs		int64						// Total Nanoseconds = MilliSecond Nanoseconds + MicroSeconds Nanoseconds + Nanoseconds
//		DateTime 				time.Time				// DateTime value for this DateTzDto Type
//		DateTimeFmt			string					// Date Time Format String. Default is "2006-01-02 15:04:05.000000000 -0700 MST"
//		TimeZone				TimeZoneDefDto	// Contains a detailed description of the Time Zone and Time Zone Location
// 																		//		associated with this date time.
//	}
//
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
// Usage
// =====
//
// Example:
//			dtzDto, err := DateTzDto{}.New(year, month, day, hour, min, sec, nanosecond , timeZoneLocation)
//
//
func (dtz DateTzDto) NewDateTime(year, month, day, hour, minute, second,
					millisecond, microsecond, nanosecond int, timeZoneLocation string) (DateTzDto, error) {

	ePrefix := "DateTzDto.New() "

	dtz2 := DateTzDto{}

	err := dtz2.SetFromDateTime(year, month, day, hour, minute, second,
						millisecond, microsecond, nanosecond, timeZoneLocation)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by dtz2.SetFromDateTime(...) " +
			"year='%v', month='%v', day='%v', hour='%v', minute='%v', second='%v', millisecond='%v', microsecond='%v' nanosecond='%v', timeZoneLocation='%v' Error='%v'",
			year, month, day, hour, minute, second, millisecond, microsecond, nanosecond, timeZoneLocation, err.Error())
	}

	return dtz2, nil
}

// CopyOut - returns a DateTzDto  instance
// which represents a deep copy of the current
// DateTzDto object.
func (dtz *DateTzDto) CopyOut() DateTzDto {
	dtz2 := DateTzDto{}

	dtz2.Year 					= dtz.Year
	dtz2.Month 					= dtz.Month
	dtz2.Day						= dtz.Day
	dtz2.Hour						= dtz.Hour
	dtz2.Minute					= dtz.Minute
	dtz2.Second					= dtz.Second
	dtz2.Millisecond		= dtz.Millisecond
	dtz2.Microsecond		= dtz.Microsecond
	dtz2.Nanosecond			= dtz.Nanosecond
	dtz2.TotalNanoSecs	= dtz.TotalNanoSecs

	if !dtz.DateTime.IsZero() {
		dtz2.DateTime = dtz.DateTime
		dtz2.TimeZone = dtz.TimeZone.CopyOut()
	} else {
		dtz2.TimeZone				= TimeZoneDefDto{}
		dtz2.DateTime				= time.Time{}
	}

	return dtz2
}

// Empty - sets all values of the current DateTzDto
// instance to their uninitialized or zero state.
func (dtz *DateTzDto) Empty() {

	dtz.Year 						= 0
	dtz.Month 					= 0
	dtz.Day							= 0
	dtz.Hour						= 0
	dtz.Minute					= 0
	dtz.Second					= 0
	dtz.Millisecond			= 0
	dtz.Microsecond			= 0
	dtz.Nanosecond			= 0
	dtz.TotalNanoSecs		= 0
	dtz.TimeZone				= TimeZoneDefDto{}
	dtz.DateTime				= time.Time{}

	return
}

// SetDateTimeFmt - Sets the DateTzDto data field 'DateTimeFmt'.
// This string is used to format the DateTzDto DateTime field
// when DateTzDto.String() is called.
func (dtz *DateTzDto) SetDateTimeFmt(dateTimeFmtStr string) {

	dtz.DateTimeFmt = dateTimeFmtStr

}

// SetFromTime - Sets the values of the current DateTzDto fields
// based on an input parameter 'dateTime' (time.time).
//
// Input Parameter
// ===============
//
// dateTime   time.Time - a date time value
//
// Returns
// =======
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromTime(dateTime time.Time) error {

ePrefix := "DateTzDto.SetFromTime() "

	if dateTime.IsZero() {
		return errors.New(ePrefix + "Error: Input parameter dateTime is Zero value!")
	}

	dtz.Empty()

	dtz.Year  = dateTime.Year()
	dtz.Month = int(dateTime.Month())
	dtz.Day = dateTime.Day()
	dtz.Hour = dateTime.Hour()
	dtz.Minute = dateTime.Minute()
	dtz.Second = dateTime.Second()
	dtz.allocateNanoseconds(int64(dateTime.Nanosecond()))
	dtz.DateTime = dateTime

	var err error
	dtz.TimeZone, err = TimeZoneDefDto{}.New(dateTime)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeZoneDefDto{}.New(dateTime). dateTime='%v'  Error='%v'", dateTime.Format(TzDtoYrMDayFmtStr), err.Error())
	}

	return nil
}

// SetFromDateTimeElements - sets the values of the current DateTzDto
// data fields based on input parameters of date time components and
// a time zone location.
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// minute						int			- minute number	0 - 59
// second						int			- second number	0	-	59
// nanosecond				int			- nanosecond number 0 - 999999999
//														This represents the total number of
//														nanoseconds which is less than one second.
//
// timeZoneLocation	string	- time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// Returns
// =======
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromDateTimeElements(year, month, day, hour, minute, second,
												nanosecond int, timeZoneLocation string) (error) {

	ePrefix := "DateTzDto.SetFromDateTimeElements() "


	if year < 0 {
		return fmt.Errorf(ePrefix + "Error: Input parameter year number is INVALID. 'year' must be greater than or equal to Zero. year='%v'", year)
	}

	if month < 1 || month > 12  {
		return fmt.Errorf(ePrefix + "Error: Input parameter month number is INVALID. Correct range is 1-12. month='%v'", month)
	}


	if day < 1 || day > 31  {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'day' number is INVALID. Correct range is 1-31. day='%v'", day)
	}


	if hour < 0 || hour > 24 {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'hour' number is INVALID. Correct range is 0-24. hour='%v'", hour)
	}

	if minute < 0 || minute > 59 {
		return fmt.Errorf(ePrefix + "Error: Input parameter minute number is INVALID. Correct range is 0 - 59. minute='%v'", minute)
	}

	if second < 0 || second > 59 {
		return fmt.Errorf(ePrefix + "Error: Input parmeter second number is INVALID. Correct range is 0 - 59. second='%v'", second)
	}


	maxNanoSecs := int(time.Second) - int(1)

	if nanosecond < 0 || nanosecond > maxNanoSecs {
		return fmt.Errorf(ePrefix + "Error: Input parameter nanoseconds exceeds maximum limit and is INVLIAD. Correct range is 0 - %v. nanosecond='%v'", maxNanoSecs, nanosecond)
	}

	if year==0 && month==0 && day == 0 && hour ==0 &&
		minute == 0 && second == 0 && nanosecond == 0 {

		return fmt.Errorf(ePrefix + "Error: All input parameter date time elements equal ZERO!")
	}

	if len(timeZoneLocation) == 0 {
		return errors.New(ePrefix + "Error: Input parameter 'timeZoneLocation' is an EMPTY STRING! 'timeZoneLocation' is required!")
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		timeZoneLocation = "Local"
	}

	loc, err := time.LoadLocation(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: Invalid time zone location! timeZoneLocation='%v'", timeZoneLocation)
	}

	dtz.Empty()

	dtz.Year 				= year
	dtz.Month				= month
	dtz.Day 				= day
	dtz.Hour 				= hour
	dtz.Minute			= minute
	dtz.Second			= second
	dtz.DateTime 		= time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, loc)

	dtz.TimeZone, err = TimeZoneDefDto{}.New(dtz.DateTime)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeZoneDefDto{}.New(dtz.DateTime). dtz.DateTime='%v'  Error='%v'", dtz.DateTime.Format(TzDtoYrMDayFmtStr), err.Error())
	}

	dtz.allocateNanoseconds(int64(nanosecond))

	return nil

}

// SetFromDateTime - Sets the values of the Date Time fields
// for the current DateTzDto instance based on time components
// and a Time Zone Location.
//
// Note that this variation of time elements breaks time down by
// hour, minute, second, millisecond, microsecond and nanosecond.
//
// See method SetFromDateTimeElements(), above, which uses a slightly
// different set of time components.
//
//
// Input Parameters
// ================
//
// year 						int			- year number
// month						int			- month number 	1 - 12
// day							int			- day number   	1 - 31
// hour							int			- hour number  	0 - 24
// min							int			- minute number	0 - 59
// sec							int			- second number	0	-	59
// millisecond			int			- millisecond number 0 - 999
// microsecond			int			-	microsecond number 0 - 999
// nanosecond				int			- nanosecond number 0 - 999
// timeZoneLocation	string	- time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the local time zone
//																location for the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// Returns
// =======
//
// error - 		If successful the returned error Type is set equal to 'nil'. If errors are
//						encountered this error Type will encapsulate an error message.
//
func (dtz *DateTzDto) SetFromDateTime(year, month, day, hour, minute, second,
millisecond, microsecond, nanosecond int, timeZoneLocation string) error {
	ePrefix := "DateTzDto.SetFromDateTime() "

	var err error

	if year < 0 {
		return fmt.Errorf(ePrefix + "Error: Input parameter year number is INVALID. 'year' must be greater than or equal to Zero. year='%v'", year)
	}

	if month < 1 || month > 12  {
		return fmt.Errorf(ePrefix + "Error: Input parameter month number is INVALID. Correct range is 1-12. month='%v'", month)
	}

	if day < 1 || day > 31  {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'day' number is INVALID. Correct range is 1-31. day='%v'", day)
	}

	if hour < 0 || hour > 24 {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'hour' number is INVALID. Correct range is 0-24. hour='%v'", hour)
	}

	if minute < 0 || minute > 59 {
		return fmt.Errorf(ePrefix + "Error: Input parameter minute number is INVALID. Correct range is 0 - 59. min='%v'", minute)
	}

	if second < 0 || second > 59 {
		return fmt.Errorf(ePrefix + "Error: Input parmeter second number is INVALID. Correct range is 0 - 59. second='%v'", second)
	}

	if millisecond < 0 || millisecond > 999 {
		return fmt.Errorf(ePrefix + "Error: Input parameter millisecond is INVALID. Correct range is 0 - 999. millisecond='%v'", millisecond)
	}

	if microsecond < 0 || microsecond > 999 {
		return fmt.Errorf(ePrefix + "Error: Input parameter microsecond is INVALID. Correct range is 0 - 999,999. microsecond='%v'", microsecond)
	}

	if nanosecond < 0 || nanosecond > 999 {
		return fmt.Errorf(ePrefix + "Error: Input parameter nanosecond is INVALID. Correct range is 0 - 999. nanosecond='%v'", nanosecond)
	}

	if year == 0 && month == 0 && day == 0 && hour == 0 && minute == 0 && second == 0 &&
		millisecond == 0 && microsecond == 0 && nanosecond == 0 {
			return errors.New(ePrefix + "Error: All time element input parameters are zero!")
	}

	if len(timeZoneLocation) == 0 {
		return errors.New(ePrefix + "Error: Input parameter 'timeZoneLocation' is a ZERO length string!")
	}

	if strings.ToLower(timeZoneLocation) == "local" {
		timeZoneLocation = "Local"
	}

	dtz.Empty()

	dtz.TotalNanoSecs =  int64(millisecond) * int64(time.Millisecond)
	dtz.Millisecond = millisecond
	dtz.TotalNanoSecs += int64(microsecond) * int64(time.Microsecond)
	dtz.Microsecond = microsecond
	dtz.TotalNanoSecs += int64(nanosecond)
	dtz.Nanosecond = nanosecond
	loc, err := time.LoadLocation(timeZoneLocation)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by time.LoadLocation(timeZoneLocation). timeZoneLocation='%v'  Error='%v'", timeZoneLocation, err.Error())
	}

	dtz.DateTime = time.Date(year, time.Month(month),day, hour, minute, second, int(dtz.TotalNanoSecs), loc)

	dtz.TimeZone, err = TimeZoneDefDto{}.New(dtz.DateTime)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDefDto{}.New(dtz.DateTime). dtz.DateTime='%v'  Error=%v", dtz.DateTime.Format(TzDtoYrMDayFmtStr), err.Error())
	}

	dtz.Year = dtz.DateTime.Year()
	dtz.Month = int(dtz.DateTime.Month())
	dtz.Hour = dtz.DateTime.Hour()
	dtz.Minute = dtz.DateTime.Minute()
	dtz.Second = dtz.DateTime.Second()

	return nil
}

// String - This method returns the DateTzDto
// DateTime field value formatted as a string.
// If the DateTzDto field DateTimeFmt is an
// empty string, a default format string will
// be used. The default format is:
//
// TzDtoYrMDayFmtStr = "2006-01-02 15:04:05.000000000 -0700 MST"
//
func (dtz *DateTzDto) String() string {

	fmtStr := dtz.DateTimeFmt

	if len(fmtStr) == 0 {
		fmtStr = TzDtoYrMDayFmtStr
	}

	return dtz.DateTime.Format(fmtStr)
}

// allocateNanoseconds - allocates total Nanoseconds to milliseconds, microseconds
// and nanoseconds.
func (dtz *DateTzDto) allocateNanoseconds(totNanoseconds int64) {

	if totNanoseconds == 0 {
		dtz.TotalNanoSecs = 0
		dtz.Millisecond = 0
		dtz.Microsecond = 0
		dtz.Nanosecond = 0
		return
	}

	r := int(totNanoseconds)

	dtz.Millisecond = r / int(time.Millisecond)

	r -= dtz.Millisecond * int(time.Millisecond)

	if r == 0 {
		return
	}

	dtz.Microsecond = r / int(time.Microsecond)

	r -= dtz.Microsecond * int(time.Microsecond)

	dtz.Nanosecond = r

	dtz.TotalNanoSecs = totNanoseconds

	return
}



// TimeZoneDto - Time Zone Data and Methods
// ============================================
type TimeZoneDto struct {
	Description 	string					// Unused - available for tagging, classification or
																//		labeling.
	TimeIn      	time.Time				// Original input time value
	TimeInZone  	TimeZoneDefDto	// Time Zone Definition Dto associated with 'TimeIn'
	TimeOut     	time.Time				// TimeOut - 'TimeIn' value converted to TimeOut
	TimeOutZone 	TimeZoneDefDto  // Time Zone Definition Dto associated with TimeOut
	TimeUTC     	time.Time				// TimeUTC (Universal Coordinated Time aka 'Zulu') value
																// 		equivalent to TimeIn
	TimeUTCZone		TimeZoneDefDto	// Time Zone Definition Dto associated with Time UTC
	TimeLocal			time.Time				// TimeIn value converted to the 'Local' Time Zone Location.
																// 		'Local' is the Time Zone Location	used by the host computer.
	TimeLocalZone	TimeZoneDefDto	// Time Zone Definition Dto associated with 'TimeLocal'
}

// AddDate - Adds specified years, months and days values to the
// current time values maintained by this TimeZoneDto
//
// Input Parameters
// ================
// years		int		- Number of years to add to current TimeZoneDto instance
// months		int		- Number of months to add to current TimeZoneDto instance
// days			int		- Number of months to add to current TimeZoneDto instance
//
// Returns
// ======
// If successful, this method adds input date values to the current TimeZoneDto.
//
// error	- If errors are encountered, this method returns an error object. Otherwise,
//					the error value is 'nil'.
//
func (tzu *TimeZoneDto) AddDate(years, months, days int) error {

	ePrefix := "TimeZoneDto.AddDate() "

	err := tzu.IsTimeZoneDtoValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error: This Time Zone Utility is INVALID!  Error='%v'", err.Error())
	}

	tzu.TimeIn = tzu.TimeIn.AddDate(years, months, days)

	tzu.TimeInZone, err = TimeZoneDefDto{}.New(tzu.TimeIn)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by TimeZoneDefDto{}.New(tzu.TimeIn). tzu.TimeIn")
	}

	tzu.TimeOut = tzu.TimeOut.AddDate(years, months, days)
	tzu.TimeUTC = tzu.TimeUTC.AddDate(years, months, days)
	tzu.TimeLocal = tzu.TimeLocal.AddDate(years, months, days)

	return nil
}


// AddDateTime - Adds input time elements to the time
// value of the current TimeZoneDto instance.
//
// Input Parameters
// ================
// years				int		- Number of years added to current TimeZoneDto
// months				int		- Number of months added to current TimeZoneDto
// days					int		- Number of days added to current TimeZoneDto
// hours				int		- Number of hours added to current TimeZoneDto
// minutes			int		- Number of minutes added to current TimeZoneDto
// seconds			int		- Number of seconds added to current TimeZoneDto
// milliseconds	int		- Number of milliseconds added to current TimeZoneDto
// microseconds	int		- Number of microseconds added to current TimeZoneDto
// nanoseconds	int		- Number of nanoseconds added to current TimeZoneDto
//
// Note: 	Input parameters may be either negative or positive. Negative
// 				values will subtract time from the current TimeZoneDto instance.
//
// Returns
// =======
//
// This method returns an error instance if errors are encountered. There
// are no other returns. If successful, the method updates
// the values of the current TimeZoneDto instance.
//
func (tzu *TimeZoneDto) AddDateTime(years, months, days, hours, minutes,
												seconds, milliseconds, microseconds, nanoseconds int) error {

	ePrefix := "TimeZoneDto.AddDateTime() "

	err := tzu.IsTimeZoneDtoValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "This TimeZoneDto instance is INVALID! Error='%v'", err.Error())
	}

	err = tzu.AddDate(years, months, days)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tzu.AddDate(years, months, days). Error='%v'", err.Error())
	}

	err = tzu.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tzu.AddTime(...). Error='%v'", err.Error())
	}

	return nil
}

// AddDuration - Adds 'duration' to the time values maintained by the
// current TimeZoneDto.
//
// Input Parameters
// ================
//
// duration		time.Duration		- May be a positive or negative duration
//															value which is added to the time value
//															of the current TimeZoneDto instance.
//
func (tzu *TimeZoneDto) AddDuration(duration time.Duration) error {

	ePrefix := "TimeZoneDto.AddDuration() "

	if duration == 0 {
		return nil
	}

	err := tzu.IsTimeZoneDtoValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "This TimeZoneDto instance is INVALID! Error='%v'", err.Error())
	}

	tzu.TimeIn = tzu.TimeIn.Add(duration)
	
	tzu.TimeInZone, err = TimeZoneDefDto{}.New(tzu.TimeIn)
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeZoneDefDto{}.New(tzu.TimeIn). tzu.TimeIn='%v'  Error='%v'", tzu.TimeIn.Format(TzDtoYrMDayFmtStr), err.Error())
	}
	
	tzu.TimeOut = tzu.TimeOut.Add(duration)
	tzu.TimeOutZone, err = TimeZoneDefDto{}.New(tzu.TimeOut)
	
	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeZoneDefDto{}.New(tzu.TimeOut). tzu.TimeOut='%v'  Error='%v'", tzu.TimeOut.Format(TzDtoYrMDayFmtStr), err.Error())
	}
	
	
	tzu.TimeUTC = tzu.TimeUTC.Add(duration)
	tzu.TimeUTCZone, err = TimeZoneDefDto{}.New(tzu.TimeUTC)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeZoneDefDto{}.New(tzu.TimeUTC). tzu.TimeUTC='%v'  Error='%v'", tzu.TimeUTC.Format(TzDtoYrMDayFmtStr), err.Error())
	}
	
	
	tzu.TimeLocal = tzu.TimeLocal.Add(duration)
	tzu.TimeLocalZone, err = TimeZoneDefDto{}.New(tzu.TimeLocal)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from TimeZoneDefDto{}.New(tzu.TimeLocal). tzu.TimeLocal='%v'  Error='%v'", tzu.TimeLocal.Format(TzDtoYrMDayFmtStr), err.Error())
	}

	return nil
}

// AddTime - Adds time elements to the time value of the current
// TimeZoneDto instance.
//
// Input Parameters:
// =================
//
// hours				- hours to be added to current TimeZoneDto
// minutes			- minutes to be added to current TimeZoneDto
// seconds			- seconds to be added to current TimeZoneDto
// milliseconds	- milliseconds to be added to current TimeZoneDto
// microseconds	- microseconds to be added to current TimeZoneDto
// nanoseconds	- nanoseconds to be added to current TimeZoneDto
//
// Note: 	Negative time values may be entered to subtract time from the
// 				current TimeZoneDto time values.
//
// Returns
// =======
//
// If successful this method updates the time value fields of the current TimeZoneDto instance.
//
// error - 	If errors are encountered, the returned 'error' object is populated. Otherwise, 'error'
//					is set to 'nil'.
//
func (tzu *TimeZoneDto) AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds int) error {

	ePrefix := "TimeZoneDto.AddTime() "

	err := tzu.IsTimeZoneDtoValid()

	if err != nil {
		return fmt.Errorf(ePrefix + "This TimeZoneDto instance is INVALID! Error='%v'", err.Error())
	}

	if hours < 0  {
		return fmt.Errorf(ePrefix + "Error: Input parameter 'hours' number is INVALID. Correct range equal to or greater than Zero. hours='%v'", hours)
	}


	var totNanoSecs  int64

	totNanoSecs = int64(time.Hour) * int64(hours)
	totNanoSecs += int64(time.Minute) * int64(minutes)
	totNanoSecs += int64(time.Second) * int64(seconds)
	totNanoSecs += int64(time.Millisecond) * int64(milliseconds)
	totNanoSecs += int64(time.Microsecond) * int64(microseconds)
	totNanoSecs += int64(nanoseconds)

	err = tzu.AddDuration(time.Duration(totNanoSecs))

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by tzu.AddDuration(time.Duration(totNanoSecs)). Error='%v'", err.Error())
	}

	return nil
}

// ConvertTz - Convert Time from existing time zone to targetTZone.
// The results are stored in the TimeZoneDto data structure.
// The input time and output time are equivalent times adjusted
// for different time zones.
//
// Input Parameters:
//
// tIn 				time.Time 	- initial time values
// targetTz 	string			- time zone location must be designated as one of two values.
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// Output Values are returned in the tzu (TimeZoneDto)
// data fields. tzu.TimeOut contains the correct time in the 'target' time
// zone.
//
func (tzu TimeZoneDto) ConvertTz(tIn time.Time, targetTz string) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.ConvertTz() "

	tzuOut := TimeZoneDto{}

	if isValidTz, _, _ := tzu.IsValidTimeZone(targetTz); !isValidTz {
		return tzuOut, errors.New(fmt.Sprintf("%v Error: targetTz is INVALID!! Input Time Zone == %v", ePrefix, targetTz))
	}

	if tIn.IsZero() {
		return tzuOut, errors.New(ePrefix + "Error: Input parameter time, 'tIn' is zero and INVALID")
	}

	tzOut, err := time.LoadLocation(targetTz)

	if err != nil {
		return tzuOut, fmt.Errorf("%vError Loading Target IANA Time Zone 'targetTz', %v. Errors: %v ",ePrefix, targetTz, err.Error())
	}


	tzuOut.setTimeIn(tIn)

	tzuOut.setTimeOut(tIn.In(tzOut))

	tzuOut.setUTCTime(tIn)

	err = tzuOut.setLocalTime(tIn)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.SetLocalTime(tIn). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// CopyOut - Creates and returns a deep copy of the
// current TimeZoneDto instance.
func (tzu *TimeZoneDto) CopyOut() TimeZoneDto {
	tzu2 := TimeZoneDto{}
	tzu2.Description 		= tzu.Description
	tzu2.TimeIn 				= tzu.TimeIn
	tzu2.TimeInZone 		= tzu.TimeInZone.CopyOut()
	tzu2.TimeOut 				= tzu.TimeOut
	tzu2.TimeOutZone 		= tzu.TimeOutZone.CopyOut()
	tzu2.TimeUTC 				= tzu.TimeUTC
	tzu2.TimeUTCZone 		= tzu.TimeUTCZone.CopyOut()
	tzu2.TimeLocal 			= tzu.TimeLocal
	tzu2.TimeLocalZone 	= tzu.TimeLocalZone.CopyOut()

	return tzu2
}

// CopyIn - Copies another TimeZoneDto
// to the current TimeZoneDto data fields.
func (tzu *TimeZoneDto) CopyIn(tzu2 TimeZoneDto) {
	
tzu.Empty()

	tzu.Description 		= tzu2.Description
	tzu.TimeIn 					= tzu2.TimeIn
	tzu.TimeInZone 			= tzu2.TimeInZone.CopyOut()
	tzu.TimeOut 				= tzu2.TimeOut
	tzu.TimeOutZone 		= tzu2.TimeOutZone.CopyOut()
	tzu.TimeUTC 				= tzu2.TimeUTC
	tzu.TimeUTCZone 		= tzu2.TimeUTCZone.CopyOut()
	tzu.TimeLocal 			= tzu2.TimeLocal
	tzu.TimeLocalZone 	= tzu2.TimeLocalZone.CopyOut()
	
}

// Equal - returns a boolean value indicating
// whether two TimeZoneDto data structures
// are equivalent.
func (tzu *TimeZoneDto) Equal(tzu2 TimeZoneDto) bool {
	
	if !tzu.TimeIn.Equal(tzu2.TimeIn) 					||
		!tzu.TimeInZone.Equal(&tzu2.TimeInZone) 	||
		!tzu.TimeOut.Equal(tzu2.TimeOut) 					||
		!tzu.TimeOutZone.Equal(&tzu2.TimeOutZone) 		||
		!tzu.TimeUTC.Equal(tzu2.TimeUTC)  				||
		!tzu.TimeUTCZone.Equal(&tzu2.TimeUTCZone) ||
		!tzu.TimeLocal.Equal(tzu2.TimeLocal)		 	||
		!tzu.TimeLocalZone.Equal(&tzu2.TimeLocalZone)	{

		return false
	}
	
	return true
}

// Empty - Clears or returns this
// TimeZoneDto to an uninitialized
// state.
func (tzu *TimeZoneDto) Empty() {
	tzu.Description 	= ""
	tzu.TimeIn 				= time.Time{}
	tzu.TimeInZone 		= TimeZoneDefDto{}
	tzu.TimeOut 			= time.Time{}
	tzu.TimeOutZone 	= TimeZoneDefDto{}
	tzu.TimeUTC 			= time.Time{}
	tzu.TimeUTCZone 	= TimeZoneDefDto{}
	tzu.TimeLocal 		= time.Time{}
	tzu.TimeLocalZone = TimeZoneDefDto{}
}


// GetTimeInDto - returns a DateTzDto instance representing the value
// of the TimeIn data field for the current TimeZoneDto.
func (tzu *TimeZoneDto) GetTimeInDto() (DateTzDto, error) {

	ePrefix := "TimeZoneDto) GetTimeInDto() "

	err := tzu.IsTimeZoneDtoValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "This TimeZoneUtiltiy instance is INVALID! Error='%v'", err.Error())
	}

	dtzDto, err := DateTzDto{}.New(tzu.TimeIn)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(tzu.TimeIn). tzu.TimeIn='%v', Error='%v'", tzu.TimeIn, err.Error())
	}

	return dtzDto, nil
}

// GetTimeOutDto - returns a DateTzDto instance representing the value
// of the TimeOut data field for the current TimeZoneDto.
func (tzu *TimeZoneDto) GetTimeOutDto() (DateTzDto, error) {

	ePrefix := "TimeZoneDto) GetTimeOutDto() "

	err := tzu.IsTimeZoneDtoValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "This TimeZoneUtiltiy instance is INVALID! Error='%v'", err.Error())
	}

	dtzDto, err := DateTzDto{}.New(tzu.TimeOut)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(tzu.TimeOut). tzu.TimeOut='%v', Error='%v'", tzu.TimeOut, err.Error())
	}

	return dtzDto, nil
}

// GetTimeLocalDto - returns a DateTzDto instance representing the value
// of the TimeLocal data field for the current TimeZoneDto.
func (tzu *TimeZoneDto) GetTimeLocalDto() (DateTzDto, error) {

	ePrefix := "TimeZoneDto) GetTimeLocalDto() "

	err := tzu.IsTimeZoneDtoValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "This TimeZoneUtiltiy instance is INVALID! Error='%v'", err.Error())
	}

	dtzDto, err := DateTzDto{}.New(tzu.TimeLocal)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(tzu.TimeLocal). tzu.TimeLocal='%v', Error='%v'", tzu.TimeLocal, err.Error())
	}

	return dtzDto, nil
}

// GetTimeUtcDto - returns a DateTzDto instance representing the value
// of the TimeUTC data field for the current TimeZoneDto.
func (tzu *TimeZoneDto) GetTimeUtcDto() (DateTzDto, error) {

	ePrefix := "TimeZoneDto) GetTimeLocalDto() "

	err := tzu.IsTimeZoneDtoValid()

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "This TimeZoneUtiltiy instance is INVALID! Error='%v'", err.Error())
	}

	dtzDto, err := DateTzDto{}.New(tzu.TimeUTC)

	if err != nil {
		return DateTzDto{}, fmt.Errorf(ePrefix + "Error returned by DateTzDto{}.New(tzu.TimeUTC). tzu.TimeUTC='%v', Error='%v'", tzu.TimeUTC, err.Error())
	}

	return dtzDto, nil
}

// IsTimeZoneDtoValid - Analyzes the current TimeZoneDto
// instance and returns an error if the instance is Invalid.
func (tzu *TimeZoneDto) IsTimeZoneDtoValid() error {

	ePrefix := "TimeZoneDto.IsTimeZoneDtoValid() "

	if tzu.TimeIn.IsZero() {
		return errors.New(ePrefix + "Error: TimeIn is Zero!")
	}

	if !tzu.TimeInZone.IsValidFromDateTime(tzu.TimeIn) {
		return errors.New(ePrefix + "Error: TimeIn Zone Definition is INVALID!")
	}

	if tzu.TimeOut.IsZero() {
		return errors.New(ePrefix + "Error: TimeOut is Zero!")
	}

	if !tzu.TimeOutZone.IsValidFromDateTime(tzu.TimeOut) {
		return errors.New(ePrefix + "Error: TimeOut Zone Definition is INVALID!")
	}

	if tzu.TimeUTC.IsZero() {
		return errors.New(ePrefix + "Error: TimeUTC is Zero!")
	}

	if !tzu.TimeUTCZone.IsValidFromDateTime(tzu.TimeUTC) {
		return errors.New(ePrefix + "Error: TimeUTC Zone Definition is INVALID!")
	}

	if tzu.TimeLocal.IsZero() {
		return errors.New(ePrefix + "Error: TimeLocal is Zero!")
	}

	if !tzu.TimeLocalZone.IsValidFromDateTime(tzu.TimeLocal) {
		return errors.New(ePrefix + "Error: TimeLocal Zone Definition is INVALID!")
	}

	return nil
}

// IsValidTimeZone - Tests a Time Zone string and returns three boolean values
// designating whether the passed Time Zone string is:
// (1.) a valid time zone
// (2.) a valid iana time zone
// (3.) a valid Local time zone
func (tzu *TimeZoneDto) IsValidTimeZone(tZone string) (isValidTz, isValidIanaTz, isValidLocalTz bool) {

	isValidTz = false

	isValidIanaTz = false

	isValidLocalTz = false

	if tZone == "" {
		return
	}

	if tZone == "Local" {
		isValidTz = true
		isValidLocalTz = true
		return
	}

	_, err := time.LoadLocation(tZone)

	if err != nil {
		return
	}

	isValidTz = true

	isValidIanaTz = true

	isValidLocalTz = false

	return

}


// New - Initializes and returns a new TimeZoneDto object.
//
// Input Parameters
// ----------------
//
// tIn					 time.Time	- The input time object.
//
// tZoneOutLocation string	- The first input time value, 'tIn' will have its time zone
// 														changed to a new time zone location specified by this second
// 														parameter, 'tZoneOutLocation'. The new time associated with
// 														'tZoneOutLocation' is assigned to the TimeZoneDto data
// 														field. The 'tZoneOutLocation' time zone location must be
// 														designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneDto Type
//																(2) an Error type
//
// 	TimeZoneDto - The two input parameters are used to populate and return
// 										a TimeZoneDto structure:

//				type TimeZoneDto struct {
//									Description string
//									TimeIn      time.Time				// Original input time value
//									TimeOut     time.Time       // TimeOut - 'tIn' value converted to TimeOut
//																							// 		based on 'timeZoneOutLocation' parameter
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value equivalent to TimeIn
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error object is populated
//					with an error message.
//
func (tzu TimeZoneDto) New(tIn time.Time, timeZoneOutLocation string) (TimeZoneDto, error) {

	tzuOut := TimeZoneDto{}

	return tzuOut.ConvertTz(tIn, timeZoneOutLocation)
}

// NewAddDate - receives four parameters: a TimeZoneDto 'tzuIn' and integer values for
// 'years', 'months' and 'days'.  The 'years', 'months' and 'days' values are added to the
// 'tzuIn' date time values and returned as a new TimeZoneDto instance.
//
// Input Parameters
// ================
//
// years				int		- Number of years added to tzuIn value.
// months				int		- Number of months added to tzuIn value.
// days					int		- Number of days added to tzuIn value.
//
// Note: Negative date values may be used to subtract date values from the
// 			tzuIn value.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneDto Type
//																(2) an Error type
//
//  TimeZoneDto - 	The date input parameters are added to 'tzuIn to produce, populate and return
// 											a TimeZoneDto structure defined as follows:
//
//				type TimeZoneDto struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error object is populated
//					with an error message.
//
func (tzu TimeZoneDto) NewAddDate(tzuIn TimeZoneDto, years int, months int, days int) (TimeZoneDto, error) {
	ePrefix := "TimeZoneDto.NewAddDate()"

	err:= tzuIn.IsTimeZoneDtoValid()

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error: Input parameter tzuIn (TimeZoneDto) is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	if years == 0 && months == 0 && days == 0 {
		return tzuOut, nil
	}

	err = tzuOut.AddDate(years, months, days)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDate(years, months, days) years='%v' months='%v' days='%v'  Error='%v'",years, months, days, err.Error())
	}

	return tzuOut, nil
}

// NewAddDateTime - Receives a TimeZoneDto input parameter, 'tzuIn'
// and returns a new TimeZoneDto instance equal to 'tzuIn' plus the
// time value of the remaining time element input parameters.
//
// Input Parameters
// ================
// tzuIn				TimeZoneDto - Base TimeZoneDto object to
//																which time elements will be added.
// years				int		- Number of years added to 'tzuIn'
// months				int		- Number of months added to 'tzuIn'
// days					int		- Number of days added to 'tzuIn'
// hours				int		- Number of hours added to 'tzuIn'
// minutes			int		- Number of minutes added to 'tzuIn'
// seconds			int		- Number of seconds added to 'tzuIn'
// milliseconds	int		- Number of milliseconds added to 'tzuIn'
// microseconds	int		- Number of microseconds added to 'tzuIn'
// nanoseconds	int		- Number of nanoseconds added to 'tzuIn'
//
// Note: 	Input time element parameters may be either negative or positive.
// 				Negative values will subtract time from the returned TimeZoneDto instance.
//
// Returns
// =======
//  There are two return values: 	(1) a TimeZoneDto Type
//																(2) an Error type
//
// TimeZoneDto - 	If successful, this method returns a valid,	populated TimeZoneDto
// 										instance which is equal to the time value of 'tzuIn' plus the other
// 										input parameter date-time elements. The TimeZoneDto structure
//										is defined as follows:
//
//				type TimeZoneDto struct {
//									Description string
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error					 -  The method will return an 'error' object if errors
//										are encountered.
//
func (tzu TimeZoneDto) NewAddDateTime(tzuIn TimeZoneDto, years, months, days, hours, minutes,
seconds, milliseconds, microseconds, nanoseconds int) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewAddDateTime() "

	err := tzuIn.IsTimeZoneDtoValid()

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error: Input Parameter 'tzuIn' is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	err = tzuOut.AddDateTime(years, months, days, hours, minutes,
		seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDateTime(...). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewAddDuration - receives two input parameters, a TimeZoneDto 'tzuIn' and a
// time 'duration'. 'tzuIn' is adjusted for the specified 'duration' and the resulting
// new TimeZoneDto is returned.
//
// Input Parameters
// ================
//
// tzuIn	TimeZoneDto - The second parameter, 'duration', will be added
//													to this TimeZoneDto.
//
// duration	time.Duration	- This duration value will be added to the
//													'tzuIn' input parameter to create, populate and
//													return a new updated TimeZoneDto instance.
//
// Note: 	Input parameter 'duration' will accept both positive and negative values.
// 				Negative values will effectively subtract the duration from 'tzuIn' time
// 				values.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneDto Type
//																(2) an Error type
//
//  TimeZoneDto - 	The input parameter 'duration' is added to 'tzuIn to produce, populate and return
// 											a TimeZoneDto structure:
//
//				type TimeZoneDto struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error object is populated
//					with an error message.
//
func (tzu TimeZoneDto) NewAddDuration(tzuIn TimeZoneDto, duration time.Duration) (TimeZoneDto, error) {
	ePrefix := "TimeZoneDto.NewAddDuration() "

	err := tzuIn.IsTimeZoneDtoValid()

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error: Input Parameter 'tzuIn' is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	err = tzuOut.AddDuration(duration)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDuration(duration). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewAddTime - returns a new TimeZoneDto equivalent to the input TimeZoneDto Plus time elements.
//
// Input Parameters:
// =================
//
// tzuIn TimeZoneDto - 	The base TimeZoneDto to which
//													time values will be added.
// hours				int				- Number of hours to be added to tzuIn
// minutes			int 			- Number of minutes to be added to tzuIn
// seconds			int 			- Number of seconds to be added to tzuIn
// milliseconds	int 			- Number of milliseconds to be added to tzuIn
// microseconds	int				- Number of microseconds to be added to tzuIn
// nanoseconds	int				- Number of nanoseconds to be added to tzuIn
//
// Note: Negative time values may be used to subtract time from 'tzuIn'.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneDto Type
//																(2) an Error type
//
//  TimeZoneDto - 	The time input parameters are added to 'tzuIn to produce, populate and return
// 											a TimeZoneDto structure:
//
//				type TimeZoneDto struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error object is populated
//					with an error message.
//
func (tzu TimeZoneDto) NewAddTime(tzuIn TimeZoneDto, hours, minutes, seconds, milliseconds, microseconds, nanoseconds int) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewAddTime() "

	err := tzuIn.IsTimeZoneDtoValid()

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error: Input Parameter 'tzuIn' is INVALID! Error='%v'", err.Error())
	}

	tzuOut := tzuIn.CopyOut()

	err = tzuOut.AddTime(hours, minutes, seconds, milliseconds, microseconds, nanoseconds)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf("Error returned by tzuOut.AddTime(...). Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddDate - returns a new TimeZoneDto. The TimeZoneDto is initialized
// with the 'tIn' time parameter. The 'TimeOut' data field will be set to the 'tIn'
// value and the time zone location specified by the second parameter, 'tZoneLocation'.
// The method will then add the remaining date element parameters to the new TimeZoneDto
// instance and return it to the calling function.
//
// Input Parameters
// ================
// tIn			time.Time 			- Initial time value assigned to 'TimeIn' field
//														of the new TimeZoneDto.
//
// tZoneOutLocation string	- The first input time value, 'tIn' will have its time zone
// 														changed to a new time zone location specified by this second
// 														parameter, 'tZoneOutLocation'. The new time associated with
// 														'tZoneOutLocation' is assigned to the TimeZoneDto data
// 														field. The 'tZoneOutLocation' time zone location must be
// 														designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// years				int		- Number of years added to initial TimeZoneDto value.
// months				int		- Number of months added to initial TimeZoneDto value.
// days					int		- Number of days added to initial TimeZoneDto value.
//
// Note: Negative date values may be used to subtract date values from the
// 			initial TimeZoneDto.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneDto Type
//																(2) an Error type
//
//  TimeZoneDto - 	The date input parameters are added to a TimeZoneDto created from
//											input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//											instance is then returned to the calling function. A TimeZoneDto structure
//											is defined as follows:
//
//				type TimeZoneDto struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error instance is populated
//					with an error message.
//
func (tzu TimeZoneDto) NewTimeAddDate(tIn time.Time, tZoneOutLocation string, years,
																						months, days int) (TimeZoneDto, error) {
	ePrefix := "TimeZoneDto.NewTimeAddDate() "

	tzuOut, err := tzu.ConvertTz(tIn, tZoneOutLocation)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returne by tzu.ConvertTz(tIn, tZoneOutLocation). tIn='%v' tZoneOutLocation='%v'  Error='%v'", tIn, tZoneOutLocation, err.Error())
	}

	err = tzuOut.AddDate(years, months, days)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDate(years, months, days) years='%v' months='%v' days='%v' Error='%v'", years, months, days, err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddDateTime - returns a new TimeZoneDto. The TimeZoneDto is initialized
// with the 'tIn' time parameter. The 'TimeOut' data field will be set to the 'tIn'
// value adjusted for the time zone location specified by the second parameter, 'tZoneLocation'.
// The method will then add the remaining date-time element parameters to the new TimeZoneDto
// instance and return it to the calling function.
//
// Input Parameters
// ================
// tIn			time.Time 		- Initial time value assigned to 'TimeIn' field
//													of the new TimeZoneDto.
//
// tZoneLocation string		- The first input time value, 'tIn' will have its time zone
// 													changed to a new time zone location specified by this second
// 													parameter, 'tZoneLocation'. This time zone location must be
// 													designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// years				int		- Number of years added to initial TimeZoneDto value.
// months				int		- Number of months added to initial TimeZoneDto value.
// days					int		- Number of days added to initial TimeZoneDto value.
// hours				int		- Number of hours to be added to initial TimeZoneDto value.
// minutes			int		- Number of minutes to be added to initial TimeZoneDto value.
// seconds			int 	- Number of seconds to be added to initial TimeZoneDto value.
// milliseconds	int		- Number of milliseconds to be added to initial TimeZoneDto value.
// microseconds	int		- Number of microseconds to be added to initial TimeZoneDto value.
// nanoseconds	int 	- Number of nanoseconds to be added to initial TimeZoneDto value.
//
// Note: Negative date-time values may be used to subtract date-time from the initial TimeZoneDto.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneDto Type
//																(2) an Error type
//
//  TimeZoneDto - 	The date-time input parameters are added to a TimeZoneDto created from
//											input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//											instance is then returned to the calling function. A TimeZoneDto structure
//											is defined as follows:
//
//				type TimeZoneDto struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error instance is populated
//					with an error message.
//
func (tzu TimeZoneDto) NewTimeAddDateTime(tIn time.Time, tZoneLocation string, years, months,
															days, hours, minutes, seconds, milliseconds, microseconds,
																	nanoseconds int) (TimeZoneDto, error) {

	ePrefix := "TimeZoneDto.NewTimeAddDateTime() "

	tzuOut, err := tzu.ConvertTz(tIn, tZoneLocation)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returne by tzu.ConvertTz(tIn, tZoneLocation). tIn='%v' tZoneLocation='%v'  Error='%v'", tIn, tZoneLocation, err.Error())
	}

	err = tzuOut.AddDateTime(years, months, days, hours, minutes, seconds, milliseconds,
														microseconds, nanoseconds)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDateTime(...)  Error='%v'", err.Error())
	}

	return tzuOut, nil
}


// NewTimeAddDuration - receives a 'tIn' time.Time parameter and a 'tZoneLocation' parameter
// which are used to construct an initial TimeZoneDto instance. The 'TimeOut'
// data field of this initial TimeZoneDto will contain the value of 'tIn'
// converted to a different time zone specified by 'tZoneLocation'.
//
// The 'duration' parameter will be added to this initial TimeZoneDto and
// an updated TimeZoneDto instance will be returned to the calling function.
//
// Input Parameters
// ================
// tIn				time.Time 	- Initial time value assigned to 'TimeIn' field
//													of the new TimeZoneDto.
//
// tZoneLocation string		- The first input time value, 'tIn' will have its time zone
// 													changed to a new time zone location specified by this second
// 													parameter, 'tZoneLocation'. This time zone location must be
// 													designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// duration			time.Duration	- an int64 duration value which is added to the date time
//							value of the initial TimeZoneDto created from 'tIn' and 'tZoneLocation'.
//
// Note: Negative duration values may be used to subtract time duration from the initial TimeZoneDto
// 			 date time values.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneDto Type
//																(2) an Error type
//
//  TimeZoneDto - 	The duration input parameter is added to a TimeZoneDto created from
//											input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//											instance is then returned to the calling function. A TimeZoneDto structure
//											is defined as follows:
//
//				type TimeZoneDto struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error instance is populated
//					with an error message.
//
func (tzu TimeZoneDto) NewTimeAddDuration(tIn time.Time, tZoneLocation string, duration time.Duration) (TimeZoneDto, error) {
	ePrefix := "TimeZoneDto.NewTimeAddDuration() "

	tzuOut, err := tzu.ConvertTz(tIn, tZoneLocation)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returne by tzu.ConvertTz(tIn, tZoneLocation). tIn='%v' tZoneLocation='%v'  Error='%v'", tIn, tZoneLocation, err.Error())
	}

	err = tzuOut.AddDuration(duration)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddDuration(duration). duration='%v'  Error='%v'",duration, err.Error())
	}

	return tzuOut, nil
}

// NewTimeAddTime - receives a 'tIn' time.Time parameter and a 'tZoneLocation' parameter
// which are used to construct an initial TimeZoneDto instance. The 'TimeOut'
// data field of this initial TimeZoneDto will contain the value of 'tIn'
// converted to a different time zone specified by 'tZoneLocation'.
//
// The remaining time parameters will be added to this initial TimeZoneDto and
// the updated TimeZoneDto will be returned to the calling function.
//
// Input Parameters
// ================
// tIn				time.Time 	- Initial time value assigned to 'TimeIn' field
//													of the new TimeZoneDto.
//
// tZoneLocation string		- The first input time value, 'tIn' will have its time zone
// 													changed to a new time zone location specified by this second
// 													parameter, 'tZoneLocation'. This time zone location must be
// 													designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
// hours				int				- Number of hours to be added to initial TimeZoneDto
// minutes			int 			- Number of minutes to be added to initial TimeZoneDto
// seconds			int 			- Number of seconds to be added to initial TimeZoneDto
// milliseconds	int 			- Number of milliseconds to be added to initial TimeZoneDto
// microseconds	int				- Number of microseconds to be added to initial TimeZoneDto
// nanoseconds	int				- Number of nanoseconds to be added to initial TimeZoneDto
//
// Note: Negative time values may be used to subtract time from initial TimeZoneDto.
//
//	Returns
//	=======
//  There are two return values: 	(1) a TimeZoneDto Type
//																(2) an Error type
//
//  TimeZoneDto - 	The time input parameters are added to a TimeZoneDto created from
//											input parameters, 'tIn' and 'tZoneOutLocation'. The updated TimeZoneDto
//											instance is then returned to the calling function. A TimeZoneDto structure
//											is defined as follows:
//
//				type TimeZoneDto struct {
//									Description string					// Unused. Available for tagging and classification.
//									TimeIn      time.Time				// Original input time value
//									TimeInLoc   *time.Location  // Time Zone Location associated with TimeIn
//									TimeOut     time.Time       // TimeOut - TimeIn value converted to TimeOut
//																							// 		based on a specific Time Zone Location.
//
//									TimeOutLoc  *time.Location	// Time Zone Location associated with TimeOut
//									TimeUTC     time.Time				// TimeUTC (Universal Coordinated Time) value
// 																										equivalent to TimeIn
//
//									TimeLocal		time.Time				// Equivalent to TimeIn value converted to the 'Local'
//																							// Time Zone Location. 'Local' is the Time Zone Location
//																							// 	used by the host computer.
//				}
//
//	error	-	If the method completes successfully, the returned error instance is
//					set to nil. If errors are encountered, the returned error instance is populated
//					with an error message.
//
func (tzu TimeZoneDto) NewTimeAddTime(tIn time.Time, tZoneLocation string, hours, minutes, seconds, milliseconds, microseconds, nanoseconds int) (TimeZoneDto, error) {

ePrefix := "TimeZoneDto.NewTimeAddTime() "

	tzuOut, err := tzu.ConvertTz(tIn, tZoneLocation)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returne by tzu.ConvertTz(tIn, tZoneLocation). tIn='%v' tZoneLocation='%v'  Error='%v'", tIn, tZoneLocation, err.Error())
	}

	err = tzuOut.AddTime(hours, minutes, seconds, milliseconds,
		microseconds, nanoseconds)

	if err != nil {
		return TimeZoneDto{}, fmt.Errorf(ePrefix + "Error returned by tzuOut.AddTime(...)  Error='%v'", err.Error())
	}

	return tzuOut, nil
}

// ReclassifyTimeWithNewTz - Receives a valid time (time.Time) value and changes the existing time zone
// to that specified in the 'tZone' parameter. During this time reclassification operation, the time
// zone is changed but the time value remains unchanged.
// Input Parameters:
//
// tIn time.Time 					- initial time whose time zone will be changed to
//													second input parameter, 'tZoneLocation'
//
// tZoneLocation string		- The first input time value, 'tIn' will have its time zone
// 													changed to a new time zone location specified by this second
// 													parameter, 'tZoneLocation'. This time zone location must be
// 													designated as one of two values:
//
// 														(1) the string 'Local' - signals the designation of the
// 																time zone	location used by the host computer.
//
//														(2) IANA Time Zone Location -
// 																See https://golang.org/pkg/time/#LoadLocation
// 																and https://www.iana.org/time-zones to ensure that
// 																the IANA Time Zone Database is properly configured
// 																on your system. Note: IANA Time Zone Data base is
// 																equivalent to 'tz database'.
//																Examples:
//																	"America/New_York"
//																	"America/Chicago"
//																	"America/Denver"
//																	"America/Los_Angeles"
//																	"Pacific/Honolulu"
//
func (tzu *TimeZoneDto) ReclassifyTimeWithNewTz(tIn time.Time, tZoneLocation string) (time.Time, error) {
	ePrefix := "TimeZoneDto.ReclassifyTimeWithNewTz() "

	strTime := tzu.TimeWithoutTimeZone(tIn)

	if len(tZoneLocation) == 0 {
		return time.Time{}, errors.New(ePrefix + "Error: Time Zone Location, 'tZoneLocation', is an EMPTY string!")
	}

	if strings.ToLower(tZoneLocation) == "local" {
		tZoneLocation = "Local"
	}

	isValidTz, _, _ := tzu.IsValidTimeZone(tZoneLocation)

	if !isValidTz {
		return time.Time{}, fmt.Errorf(ePrefix + "Error: Input Time Zone Location is INVALID! tZoneLocation='%v'", tZoneLocation)
	}

	tzNew, err := time.LoadLocation(tZoneLocation)

	if err != nil {
		return time.Time{}, fmt.Errorf(ePrefix + "Error returned by time.Location('%v') - Error: %v", tZoneLocation, err.Error())
	}

	tOut, err := time.ParseInLocation(TzDtoNeutralDateFmt, strTime, tzNew)

	if err != nil {
		return tOut, fmt.Errorf(ePrefix + "Error returned by time.Parse - Error: %v", err.Error())
	}

	return tOut, nil
}

// setTimeIn - Assigns time and zone values to field 'TimeIn'
func (tzu *TimeZoneDto) setTimeIn(tIn time.Time) error {

	ePrefix := "TimeZoneDto.SetTimeIn() "

	tzu.TimeIn = tIn

	var err error

	tzu.TimeInZone, err = TimeZoneDefDto{}.New(tIn)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error retrned by TimeZoneDefDto{}.New(tIn). tIn='%v'  Error='%v'", tIn.Format(TzDtoYrMDayFmtStr), err.Error())
	}

	return nil
}

// setTimeOut - Assigns time and zone values to field 'TimeOut'
func (tzu *TimeZoneDto) setTimeOut(tOut time.Time) error {

	ePrefix := "TimeZoneDto.setTimeOut()"

	var err error

	tzu.TimeOut = tOut

	tzu.TimeOutZone, err  = TimeZoneDefDto{}.New(tzu.TimeOut)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error retrned by TimeZoneDefDto{}.New(tzu.TimeOut). tzu.TimeOut='%v'  Error='%v'", tOut.Format(TzDtoYrMDayFmtStr), err.Error())
	}

	return nil

}

// setUTCTime - Assigns UTC Time and zone values to fields 'TimeUTC' and 'TimeUTCZone'
func (tzu *TimeZoneDto) setUTCTime(t time.Time) error {

	ePrefix := "TimeZoneDto.setTimeOut()"

	var err error

	tzu.TimeUTC = t.UTC()

	tzu.TimeUTCZone, err  = TimeZoneDefDto{}.New(tzu.TimeUTC)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error retrned by TimeZoneDefDto{}.New(tzu.TimeUTC). tzu.TimeUTC='%v'  Error='%v'", tzu.TimeUTC.Format(TzDtoYrMDayFmtStr), err.Error())
	}

	return nil

}

// setLocalTime - Assigns Local Time to field 'TimeLocal'
func (tzu *TimeZoneDto) setLocalTime(t time.Time) error {
	ePrefix := "TimeZoneDto.SetLocalTime() "

	tzLocal, err := time.LoadLocation("Local")

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by time.LoadLocation(\"Local\") Error='%v'", err.Error())
	}

	tzu.TimeLocal = t.In(tzLocal)
	tzu.TimeLocalZone, err  = TimeZoneDefDto{}.New(tzu.TimeLocal)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error retrned by TimeZoneDefDto{}.New(tzu.TimeLocal). tzu.TimeLocal='%v'  Error='%v'", tzu.TimeLocal.Format(TzDtoYrMDayFmtStr), err.Error())
	}

	return nil
}

// Sub - Subtracts the input TimeZoneDto from the current TimeZoneDto
// and returns the duration. Duration is calculated as:
// 						tzu.TimeLocal.Sub(tzu2.TimeLocal)
//
// The TimeLocal field is used to compute duration for this method.
//
func (tzu *TimeZoneDto) Sub(tzu2 TimeZoneDto) (time.Duration, error) {

	ePrefix := "TimeZoneDto.Sub() "

	err := tzu.IsTimeZoneDtoValid()

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix + "Error: Current TimeZoneDto (tzu) is INVALID. Error='%v'", err.Error())
	}

	err = tzu2.IsTimeZoneDtoValid()

	if err != nil {
		return time.Duration(0), fmt.Errorf(ePrefix + "Error: Input Parameter 'tzu2' is INVALID! Error='%v'", err.Error())
	}

	return tzu.TimeLocal.Sub(tzu2.TimeLocal), nil
}

// TimeWithoutTimeZone - Returns a Time String containing
// NO time zone. - When the returned string is converted to
// time.Time - in defaults to a UTC time zone.
func (tzu *TimeZoneDto) TimeWithoutTimeZone(t time.Time) string {
	return t.Format(TzDtoNeutralDateFmt)
}