package datetime

import (
  "errors"
  "sync"
)


// TzAbbreviationDto - encapsulates Time Zone abbreviation
// information. A Time Zone Abbreviation must consist entirely
// of alphabetic characters.
// 
// The Id is styled as Abbreviation text plus the UTC offset.
// Example: CST-0600 - Central Standard time with offset UTC-0600.
// 
type TimeZoneAbbreviationDto struct {
  Id                 string  // Example: "CST-0600"
  Abbrv              string  // Example: "CST"
  AbbrvDescription   string  // Example: "Central Standard Time"
  Location           string  // Example: "North America"
  UtcOffset          string  // Example: "-0600"
}


// CopyOut() - Makes and returns a deep copy of the current TimeZoneAbbreviationDto
// object.
// 
func (TzAbbrv *TimeZoneAbbreviationDto) CopyOut() TimeZoneAbbreviationDto {
  
  newDto := TimeZoneAbbreviationDto{}
  newDto.Id = TzAbbrv.Id
  newDto.Abbrv = TzAbbrv.Abbrv
  newDto.AbbrvDescription = TzAbbrv.AbbrvDescription
  newDto.Location = TzAbbrv.Location
  newDto.UtcOffset = TzAbbrv.UtcOffset
  
  return newDto
}


// CopyIn() - Copies the field values from an incoming TimeZoneAbbreviationDto
// object to the current TimeZoneAbbreviationDto object.
// 
func (TzAbbrv *TimeZoneAbbreviationDto) CopyIn(inComing *TimeZoneAbbreviationDto) error {
  
  ePrefix := "TzAbbreviationDto.CopyIn()" 
  
  if inComing == nil {
    return  errors.New(ePrefix +
      "Error: Input parameter 'incoming' is nil!")  }
  
  TzAbbrv.Id = inComing.Id
  TzAbbrv.Abbrv = inComing.Abbrv
  TzAbbrv.AbbrvDescription = inComing.AbbrvDescription
  TzAbbrv.Location = inComing.Location
  TzAbbrv.UtcOffset = inComing.UtcOffset
  return nil
}  


// StdTZoneAbbreviations - Provides thread safe access to
// standard IANA Time Zone abbreviations, abbreviation
// descriptions and UTC Offsets.
//  
type StdTZoneAbbreviations struct {
  Input                  TimeZoneAbbreviationDto
  Output                 TimeZoneAbbreviationDto
}


// AbbrvOffsetToTzReference - This method returns a type
// 'TimeZoneAbbreviation' describing a specific time zone
// abbreviation based on an input parameter consisting of
// an alphabetic time zone abbreviation and an UTC offset
// parameter.
//  
// The Time Zone Abbreviation Offset parameter, 'abbrvOffset',
// must be formatted with a time zone abbreviation in all
// upper case characters followed by the UTC Offset expressed
// in hours and minutes.

// For example, to return a 'TimeZoneAbbreviationDto' describing
// North America Central Standard Time, the 'abbrvOffset' input
// parameter must be formatted as 'CST-0600'. Note: the UTC
// offset for North America Central Standard Time is 'UTC-0600'.

// If the Abbreviation Offset parameter is invalid or if no
// 'TimeZoneAbbreviationDto' exists for the Abbreviation Offset
// parameter, this method will return a boolean value of 'false'.
//  
func (stdTzAbbrvs *StdTZoneAbbreviations) AbbrvOffsetToTzReference(
		abbrvOffset string) (TimeZoneAbbreviationDto, bool) {

	lockMapTzAbbreviationReference.Lock()

	defer lockMapTzAbbreviationReference.Unlock()

	result, ok := mapTzAbbreviationReference[abbrvOffset]

	return result, ok
}

// AbbrvOffsetToTimeZones - Returns a string array consisting of
// all standard time zones associated with a specific time zone
// abbreviation based on an input parameter consisting of an
// alphabetic time zone abbreviation and an UTC offset parameter.
// 
// The Time Zone Abbreviation Offset parameter, 'abbrvOffset'
// must be formatted with a time zone abbreviation in all
// upper case characters followed by the UTC Offset expressed
// in hours and minutes.
// 
// For example, to return a string array containing all standard
// time zones associated with North America Central Standard
// Time, the 'abbrvOffset' input parameter must be formatted as
// 'CST-0600'. Note: the UTC offset for North America Central
// Standard Time is 'UTC-0600'.
// 
// If the Time Zone Abbreviation Offset parameter is invalid or
// if no string array exists for the Abbreviation Offset
// parameter, this method will return a boolean value of 'false'.
// 
func (stdTzAbbrvs *StdTZoneAbbreviations) AbbrvOffsetToTimeZones(
		abbrvOffset string) ([]string, bool) {

	lockMapTzAbbrvsToTimeZones.Lock()

	defer lockMapTzAbbrvsToTimeZones.Unlock()

	result, ok :=  mapTzAbbrvsToTimeZones[abbrvOffset]

	return result, ok
}

// TimeZonesToAbbrvs - Returns a string array consisting of
// all time zone abbreviations associated with a standard,
// IANA time zone name passed as an input parameter.  This
// input parameter, 'timeZone', must be formatted as a
// standard IANA time zone name using upper and lower case
// characters as specified in the IANA Time Zone Database.
// 
// The returned string array actually contains Time Zone
// Abbreviation and UTC Offset pairs.
// 
// For example, the standard IANA Time Zone, 'America/Chicago'
// will return a string array consisting of two strings:
// "CDT-0500" and "CST-0600". These two strings describe the
// two time zone abbreviations associated with 'America/Chicago'
// (a.k.a. North America Central Time). "CDT-0500" stands for
// 'Central Daylight Time' and UTC-0500 (a UTC offset of
// -5 hours). Likewise, "CST-0600" identifies 'Central Standard
// Time' with an UTC offset of -6 hours (UTC-0600).
// 
// If the Time Zone input parameter is invalid or if no string
// array exists for the Time Zone input parameter, this method
// will return a boolean value of 'false'.
// 
func (stdTzAbbrvs *StdTZoneAbbreviations) TimeZonesToAbbrvs(
		timeZone string) ([]string, bool) {

	lockMapTimeZonesToTzAbbrvs.Lock()

	defer lockMapTimeZonesToTzAbbrvs.Unlock()

	result, ok := mapTimeZonesToTzAbbrvs[timeZone]

	return result, ok
}


// mapTzAbbreviationReference - A reference map including all valid
// alphabetic Time Zone abbreviations.
//

var lockMapTzAbbreviationReference sync.Mutex

var mapTzAbbreviationReference = map[string]TimeZoneAbbreviationDto{
"+00+0000"    :{"+00+0000","+00","Numerical Time Zone +00","+00","+0000"},
"+01+0100"    :{"+01+0100","+01","Numerical Time Zone +01","+01","+0100"},
"+02+0200"    :{"+02+0200","+02","Numerical Time Zone +02","+02","+0200"},
"+03+0300"    :{"+03+0300","+03","Numerical Time Zone +03","+03","+0300"},
"+04+0400"    :{"+04+0400","+04","Numerical Time Zone +04","+04","+0400"},
"+0430+0430"  :{"+0430+0430","+0430","Numerical Time Zone +0430","+0430","+0430"},
"+05+0500"    :{"+05+0500","+05","Numerical Time Zone +05","+05","+0500"},
"+0530+0530"  :{"+0530+0530","+0530","Numerical Time Zone +0530","+0530","+0500"},
"+0545+0545"  :{"+0545+0545","+0545","Numerical Time Zone +0545","+0545","+0545"},
"+06+0600"    :{"+06+0600","+06","Numerical Time Zone +06","+06","+0600"},
"+0630+0630"  :{"+0630+0630","+0630","Numerical Time Zone +0630","+0630","+0630"},
"+07+0700"    :{"+07+0700","+07","Numerical Time Zone +07","+07","+0700"},
"+08+0800"    :{"+08+0800","+08","Numerical Time Zone +08","+08","+0800"},
"+0845+0845"  :{"+0845+0845","+0845","Numerical Time Zone +0845","+0845","+0845"},
"+09+0900"    :{"+09+0900","+09","Numerical Time Zone +09","+09","+0900"},
"+10+1000"    :{"+10+1000","+10","Numerical Time Zone +10","+10","+1000"},
"+1030+1030"  :{"+1030+1030","+1030","Numerical Time Zone +1030","+1030","+1030"},
"+11+1100"    :{"+11+1100","+11","Numerical Time Zone +11","+11","+1100"},
"+12+1200"    :{"+12+1200","+12","Numerical Time Zone +12","+12","+1200"},
"+1245+1245"  :{"+1245+1245","+1245","Numerical Time Zone +1245","+1245","+1245"},
"+13+1300"    :{"+13+1300","+13","Numerical Time Zone +13","+13","+1300"},
"+14+1400"    :{"+14+1400","+14","Numerical Time Zone +14","+14","+1400"},
"-00+0000"    :{"-00+0000","-00","Numerical Time Zone -00","-00","+0000"},
"-01-0100"    :{"-01-0100","-01","Numerical Time Zone -01","-01","-0100"},
"-02-0200"    :{"-02-0200","-02","Numerical Time Zone -02","-02","-0200"},
"-03-0300"    :{"-03-0300","-03","Numerical Time Zone -03","-03","-0300"},
"-04-0400"    :{"-04-0400","-04","Numerical Time Zone -04","-04","-0400"},
"-05-0500"    :{"-05-0500","-05","Numerical Time Zone -05","-05","-0500"},
"-06-0600"    :{"-06-0600","-06","Numerical Time Zone -06","-06","-0600"},
"-07-0700"    :{"-07-0700","-07","Numerical Time Zone -07","-07","-0700"},
"-08-0800"    :{"-08-0800","-08","Numerical Time Zone -08","-08","-0800"},
"-09-0900"    :{"-09-0900","-09","Numerical Time Zone -09","-09","-0900"},
"-0930-0930"  :{"-0930-0930","-0930","Numerical Time Zone -0930","-0930","-0930"},
"-10-1000"    :{"-10-1000","-10","Numerical Time Zone -10","-10","-1000"},
"-11-1100"    :{"-11-1100","-11","Numerical Time Zone -11","-11","-1100"},
"-12-1200"    :{"-12-1200","-12","Numerical Time Zone -12","-12","-1200"},
"ACDT+1030"   :{"ACDT+1030","ACDT","Australian Central Daylight Time","Australia","+1030"},
"ACST+0930"   :{"ACST+0930","ACST","Australian Central Standard Time","Australia","+0930"},
"ADT-0300"    :{"ADT-0300","ADT","Atlantic Daylight Time","North America","-0300"},
"AEDT+1100"   :{"AEDT+1100","AEDT","Australian Eastern Daylight Time","Australia","+1100"},
"AEST+1000"   :{"AEST+1000","AEST","Australian Eastern Standard Time","Australia","+1000"},
"AKDT-0800"   :{"AKDT-0800","AKDT","Alaska Daylight Time","North America","-0800"},
"AKST-0900"   :{"AKST-0900","AKST","Alaska Standard Time","North America","-0900"},
"AST-0400"    :{"AST-0400","AST","Atlantic Standard Time","Atlantic","-0400"},
"AWST+0800"   :{"AWST+0800","AWST","Australian Western Standard Time","Australia","+0800"},
"BST+0100"    :{"BST+0100","BST","British Summer Time","Europe","+0100"},
"CAT+0200"    :{"CAT+0200","CAT","Central Africa Time","Africa","+0200"},
"CDT-0400"    :{"CDT-0400","CDT","Cuba Daylight Time","Caribbean","-0400"},
"CDT-0500"    :{"CDT-0500","CDT","Central Daylight Time","North America","-0500"},
"CEST+0200"   :{"CEST+0200","CEST","Central European Summer Time","Europe","+0200"},
"CET+0100"    :{"CET+0100","CET","Central European Time","Europe","+0100"},
"CST+0800"    :{"CST+0800","CST","China Standard Time","Asia","+0800"},
"CST-0500"    :{"CST-0500","CST","Cuba Standard Time","Caribbean","-0500"},
"CST-0600"    :{"CST-0600","CST","Central Standard Time","North America","-0600"},
"ChST+1000"   :{"ChST+1000","ChST","Chamorro Standard Time","Pacific","+1000"},
"EAT+0300"    :{"EAT+0300","EAT","Eastern Africa Time","Africa","+0300"},
"EDT-0400"    :{"EDT-0400","EDT","Eastern Daylight Time","North America","-0400"},
"EEST+0300"   :{"EEST+0300","EEST","Eastern European Summer Time","Europe","+0300"},
"EET+0200"    :{"EET+0200","EET","Eastern European Time","Europe","+0200"},
"EST-0500"    :{"EST-0500","EST","Eastern Standard Time","North America","-0500"},
"GMT+0000"    :{"GMT+0000","GMT","Greenwich Mean Time","Europe","+0000"},
"HDT-0900"    :{"HDT-0900","HDT","Hawaii-Aleutian Daylight Time","Hawaii, Aleutians","-0900"},
"HKT+0800"    :{"HKT+0800","HKT","Hong Kong Time","Asia","+0800"},
"HST-1000"    :{"HST-1000","HST","Hawaii-Aleutian Standard Time","Hawaii, Aleutians","-1000"},
"IDT+0300"    :{"IDT+0300","IDT","Israel Daylight Time","Asia","+0300"},
"IST+0100"    :{"IST+0100","IST","Irish Standard Time","Europe","+0100"},
"IST+0200"    :{"IST+0200","IST","Israel Standard Time","Asia","+0200"},
"IST+0530"    :{"IST+0530","IST","India Standard Time","Asia","+0530"},
"JST+0900"    :{"JST+0900","JST","Japan Standard Time","Asia","+0900"},
"KST+0900"    :{"KST+0900","KST","Korea Standard Time","Asia","+0900"},
"MDT-0600"    :{"MDT-0600","MDT","Mountain Daylight Time","North America","-0600"},
"MEST+0200"   :{"MEST+0200","MEST","Middle European Summer Time","Europe","+0200"},
"MET+0100"    :{"MET+0100","MET","Middle European Time","Europe","+0100"},
"MSK+0300"    :{"MSK+0300","MSK","Moscow Standard Time","Europe","+0300"},
"MST-0700"    :{"MST-0700","MST","Mountain Standard Time","North America","-0700"},
"NDT-0230"    :{"NDT-0230","NDT","Newfoundland Daylight Time","North America","-0230"},
"NST-0330"    :{"NST-0330","NST","Newfoundland Standard Time","North America","-0330"},
"NZDT+1300"   :{"NZDT+1300","NZDT","New Zealand Daylight Time","Pacific","+1300"},
"NZST+1200"   :{"NZST+1200","NZST","New Zealand Standard Time","Pacific","+1200"},
"PDT-0700"    :{"PDT-0700","PDT","Pacific Daylight Time","North America","-0700"},
"PKT+0500"    :{"PKT+0500","PKT","Pakistan Standard Time","Asia","+0500"},
"PST+0800"    :{"PST+0800","PST","Philippine Standard Time","Asia","+0800"},
"PST-0800"    :{"PST-0800","PST","Pacific Standard Time","North America","-0800"},
"SAST+0200"   :{"SAST+0200","SAST","South Africa Standard Time","Africa","+0200"},
"SST-1100"    :{"SST-1100","SST","Samoa Standard Time","Pacific","-1100"},
"UTC+0000"    :{"UTC+0000","UTC","Universal Time Coordinated","Universal","+0000"},
"WAT+0100"    :{"WAT+0100","WAT","West Africa Time","Africa","+0100"},
"WEST+0100"   :{"WEST+0100","WEST","Western European Summer Time","Europe","+0100"},
"WET+0000"    :{"WET+0000","WET","Western European Time","Europe","+0000"},
"WIB+0700"    :{"WIB+0700","WIB","Western Indonesian Time","Asia","+0700"},
"WIT+0900"    :{"WIT+0900","WIT","Eastern Indonesian Time","Asia","+0900"},
"WITA+0800"   :{"WITA+0800","WITA","Central Indonesian Time","Asia","+0800"},
}


// mapTzAbbrvsToTimeZones - A cross reference that maps
// Time Zone Abbreviations to Time Zone Canonical Values.
// 

var lockMapTzAbbrvsToTimeZones sync.Mutex

var mapTzAbbrvsToTimeZones = map[string][]string {
"+00+0000"    :{ "America/Scoresbysund","Atlantic/Azores"},
"+01+0100"    :{ "Africa/Casablanca","Africa/El_Aaiun","Etc/GMT-1"},
"+02+0200"    :{ "Antarctica/Troll","Etc/GMT-2"},
"+03+0300"    :{ "Antarctica/Syowa","Asia/Aden","Asia/Baghdad","Asia/Bahrain","Asia/Istanbul","Asia/Kuwait","Asia/Qatar","Asia/Riyadh","Etc/GMT-3","Europe/Istanbul","Europe/Kirov","Europe/Minsk","Turkey"},
"+04+0400"    :{ "Asia/Baku","Asia/Dubai","Asia/Muscat","Asia/Tbilisi","Asia/Yerevan","Etc/GMT-4","Europe/Astrakhan","Europe/Samara","Europe/Saratov","Europe/Ulyanovsk","Europe/Volgograd","Indian/Mahe","Indian/Mauritius","Indian/Reunion"},
"+0430+0430"  :{ "Asia/Kabul","Asia/Tehran","Iran"},
"+05+0500"    :{ "Antarctica/Mawson","Asia/Aqtau","Asia/Aqtobe","Asia/Ashgabat","Asia/Ashkhabad","Asia/Atyrau","Asia/Dushanbe","Asia/Oral","Asia/Qyzylorda","Asia/Samarkand","Asia/Tashkent","Asia/Yekaterinburg","Etc/GMT-5","Indian/Kerguelen","Indian/Maldives"},
"+0530+0530"  :{ "Asia/Colombo"},
"+0545+0545"  :{ "Asia/Kathmandu","Asia/Katmandu"},
"+06+0600"    :{ "Antarctica/Vostok","Asia/Almaty","Asia/Bishkek","Asia/Dacca","Asia/Dhaka","Asia/Kashgar","Asia/Omsk","Asia/Qostanay","Asia/Thimbu","Asia/Thimphu","Asia/Urumqi","Etc/GMT-6","Indian/Chagos"},
"+0630+0630"  :{ "Asia/Rangoon","Asia/Yangon","Indian/Cocos"},
"+07+0700"    :{ "Antarctica/Davis","Asia/Bangkok","Asia/Barnaul","Asia/Ho_Chi_Minh","Asia/Hovd","Asia/Krasnoyarsk","Asia/Novokuznetsk","Asia/Novosibirsk","Asia/Phnom_Penh","Asia/Saigon","Asia/Tomsk","Asia/Vientiane","Etc/GMT-7","Indian/Christmas"},
"+08+0800"    :{ "Antarctica/Casey","Asia/Brunei","Asia/Choibalsan","Asia/Irkutsk","Asia/Kuala_Lumpur","Asia/Kuching","Asia/Singapore","Asia/Ulaanbaatar","Asia/Ulan_Bator","Etc/GMT-8","Singapore"},
"+0845+0845"  :{ "Australia/Eucla"},
"+09+0900"    :{ "Asia/Chita","Asia/Dili","Asia/Khandyga","Asia/Yakutsk","Etc/GMT-9","Pacific/Palau"},
"+10+1000"    :{ "Antarctica/DumontDUrville","Asia/Ust-Nera","Asia/Vladivostok","Etc/GMT-10","Pacific/Chuuk","Pacific/Port_Moresby","Pacific/Truk","Pacific/Yap"},
"+1030+1030"  :{ "Australia/LHI","Australia/Lord_Howe"},
"+11+1100"    :{ "Antarctica/Macquarie","Asia/Magadan","Asia/Sakhalin","Asia/Srednekolymsk","Etc/GMT-11","Pacific/Bougainville","Pacific/Efate","Pacific/Guadalcanal","Pacific/Kosrae","Pacific/Norfolk","Pacific/Noumea","Pacific/Pohnpei","Pacific/Ponape"},
"+12+1200"    :{ "Asia/Anadyr","Asia/Kamchatka","Etc/GMT-12","Kwajalein","Pacific/Fiji","Pacific/Funafuti","Pacific/Kwajalein","Pacific/Majuro","Pacific/Nauru","Pacific/Tarawa","Pacific/Wake","Pacific/Wallis"},
"+1245+1245"  :{ "NZ-CHAT","Pacific/Chatham"},
"+13+1300"    :{ "Etc/GMT-13","Pacific/Apia","Pacific/Enderbury","Pacific/Fakaofo","Pacific/Tongatapu"},
"+14+1400"    :{ "Etc/GMT-14","Pacific/Kiritimati"},
"-00+0000"    :{ "Factory"},
"-01-0100"    :{ "Atlantic/Cape_Verde","Etc/GMT+1"},
"-02-0200"    :{ "America/Godthab","America/Miquelon","America/Noronha","Atlantic/South_Georgia","Brazil/DeNoronha","Etc/GMT+2"},
"-03-0300"    :{ "America/Araguaina","America/Bahia","America/Belem","America/Buenos_Aires","America/Catamarca","America/Cayenne","America/Cordoba","America/Fortaleza","America/Jujuy","America/Maceio","America/Mendoza","America/Montevideo","America/Paramaribo","America/Punta_Arenas","America/Recife","America/Rosario","America/Santarem","America/Sao_Paulo","Antarctica/Palmer","Antarctica/Rothera","Atlantic/Stanley","Brazil/East","Etc/GMT+3","America/Argentina/Buenos_Aires","America/Argentina/Catamarca","America/Argentina/ComodRivadavia","America/Argentina/Cordoba","America/Argentina/Jujuy","America/Argentina/La_Rioja","America/Argentina/Mendoza","America/Argentina/Rio_Gallegos","America/Argentina/Salta","America/Argentina/San_Juan","America/Argentina/San_Luis","America/Argentina/Tucuman","America/Argentina/Ushuaia"},
"-04-0400"    :{ "America/Asuncion","America/Boa_Vista","America/Campo_Grande","America/Caracas","America/Cuiaba","America/Guyana","America/La_Paz","America/Manaus","America/Porto_Velho","America/Santiago","Brazil/West","Chile/Continental","Etc/GMT+4"},
"-05-0500"    :{ "America/Bogota","America/Eirunepe","America/Guayaquil","America/Lima","America/Porto_Acre","America/Rio_Branco","Brazil/Acre","Etc/GMT+5"},
"-06-0600"    :{ "Chile/EasterIsland","Etc/GMT+6","Pacific/Easter","Pacific/Galapagos"},
"-07-0700"    :{ "Etc/GMT+7"},
"-08-0800"    :{ "Etc/GMT+8","Pacific/Pitcairn"},
"-09-0900"    :{ "Etc/GMT+9","Pacific/Gambier"},
"-0930-0930"  :{ "Pacific/Marquesas"},
"-10-1000"    :{ "Etc/GMT+10","Pacific/Rarotonga","Pacific/Tahiti"},
"-11-1100"    :{ "Etc/GMT+11","Pacific/Niue"},
"-12-1200"    :{ "Etc/GMT+12"},
"ACDT+1030"   :{ "Australia/Adelaide","Australia/Broken_Hill","Australia/South","Australia/Yancowinna"},
"ACST+0930"   :{ "Australia/Adelaide","Australia/Broken_Hill","Australia/Darwin","Australia/North","Australia/South","Australia/Yancowinna"},
"ADT-0300"    :{ "America/Glace_Bay","America/Goose_Bay","America/Halifax","America/Moncton","America/Thule","Atlantic/Bermuda","Canada/Atlantic"},
"AEDT+1100"   :{ "Australia/ACT","Australia/Canberra","Australia/Currie","Australia/Hobart","Australia/Melbourne","Australia/NSW","Australia/Sydney","Australia/Tasmania","Australia/Victoria"},
"AEST+1000"   :{ "Australia/ACT","Australia/Brisbane","Australia/Canberra","Australia/Currie","Australia/Hobart","Australia/Lindeman","Australia/Melbourne","Australia/NSW","Australia/Queensland","Australia/Sydney","Australia/Tasmania","Australia/Victoria"},
"AKDT-0800"   :{ "America/Anchorage","America/Juneau","America/Metlakatla","America/Nome","America/Sitka","America/Yakutat","US/Alaska"},
"AKST-0900"   :{ "America/Anchorage","America/Juneau","America/Metlakatla","America/Nome","America/Sitka","America/Yakutat","US/Alaska"},
"AST-0400"    :{ "America/Anguilla","America/Antigua","America/Aruba","America/Barbados","America/Blanc-Sablon","America/Curacao","America/Dominica","America/Glace_Bay","America/Goose_Bay","America/Grenada","America/Guadeloupe","America/Halifax","America/Kralendijk","America/Lower_Princes","America/Marigot","America/Martinique","America/Moncton","America/Montserrat","America/Port_of_Spain","America/Puerto_Rico","America/Santo_Domingo","America/St_Barthelemy","America/St_Kitts","America/St_Lucia","America/St_Thomas","America/St_Vincent","America/Thule","America/Tortola","America/Virgin","Atlantic/Bermuda","Canada/Atlantic"},
"AWST+0800"   :{ "Australia/Perth","Australia/West"},
"BST+0100"    :{ "Europe/Belfast","Europe/Guernsey","Europe/Isle_of_Man","Europe/Jersey","Europe/London","GB","GB-Eire"},
"CAT+0200"    :{ "Africa/Blantyre","Africa/Bujumbura","Africa/Gaborone","Africa/Harare","Africa/Khartoum","Africa/Kigali","Africa/Lubumbashi","Africa/Lusaka","Africa/Maputo","Africa/Windhoek"},
"CDT-0400"    :{ "America/Havana","Cuba"},
"CDT-0500"    :{ "America/Bahia_Banderas","America/Chicago","America/Knox_IN","America/Matamoros","America/Menominee","America/Merida","America/Mexico_City","America/Monterrey","America/Rainy_River","America/Rankin_Inlet","America/Resolute","America/Winnipeg","Canada/Central","Mexico/General","CST6CDT","US/Central","US/Indiana-Starke","America/Indiana/Knox","America/Indiana/Tell_City","America/North_Dakota/Beulah","America/North_Dakota/Center","America/North_Dakota/New_Salem"},
"CEST+0200"   :{ "Africa/Ceuta","Arctic/Longyearbyen","Atlantic/Jan_Mayen","Europe/Amsterdam","Europe/Andorra","Europe/Belgrade","Europe/Berlin","Europe/Bratislava","Europe/Brussels","Europe/Budapest","Europe/Busingen","Europe/Copenhagen","Europe/Gibraltar","Europe/Ljubljana","Europe/Luxembourg","Europe/Madrid","Europe/Malta","Europe/Monaco","Europe/Oslo","Europe/Paris","Europe/Podgorica","Europe/Prague","Europe/Rome","Europe/San_Marino","Europe/Sarajevo","Europe/Skopje","Europe/Stockholm","Europe/Tirane","Europe/Vaduz","Europe/Vatican","Europe/Vienna","Europe/Warsaw","Europe/Zagreb","Europe/Zurich","CET","Poland"},
"CET+0100"    :{ "Africa/Algiers","Africa/Ceuta","Africa/Tunis","Arctic/Longyearbyen","Atlantic/Jan_Mayen","Europe/Amsterdam","Europe/Andorra","Europe/Belgrade","Europe/Berlin","Europe/Bratislava","Europe/Brussels","Europe/Budapest","Europe/Busingen","Europe/Copenhagen","Europe/Gibraltar","Europe/Ljubljana","Europe/Luxembourg","Europe/Madrid","Europe/Malta","Europe/Monaco","Europe/Oslo","Europe/Paris","Europe/Podgorica","Europe/Prague","Europe/Rome","Europe/San_Marino","Europe/Sarajevo","Europe/Skopje","Europe/Stockholm","Europe/Tirane","Europe/Vaduz","Europe/Vatican","Europe/Vienna","Europe/Warsaw","Europe/Zagreb","Europe/Zurich","CET","Poland"},
"CST+0800"    :{ "Asia/Chongqing","Asia/Chungking","Asia/Harbin","Asia/Macao","Asia/Macau","Asia/Shanghai","Asia/Taipei","PRC","ROC"},
"CST-0500"    :{ "America/Havana","Cuba"},
"CST-0600"    :{ "America/Bahia_Banderas","America/Belize","America/Chicago","America/Costa_Rica","America/El_Salvador","America/Guatemala","America/Knox_IN","America/Managua","America/Matamoros","America/Menominee","America/Merida","America/Mexico_City","America/Monterrey","America/Rainy_River","America/Rankin_Inlet","America/Regina","America/Resolute","America/Swift_Current","America/Tegucigalpa","America/Winnipeg","Canada/Central","Canada/Saskatchewan","Mexico/General","CST6CDT","US/Central","US/Indiana-Starke","America/Indiana/Knox","America/Indiana/Tell_City","America/North_Dakota/Beulah","America/North_Dakota/Center","America/North_Dakota/New_Salem"},
"ChST+1000"   :{ "Pacific/Guam","Pacific/Saipan"},
"EAT+0300"    :{ "Africa/Addis_Ababa","Africa/Asmara","Africa/Asmera","Africa/Dar_es_Salaam","Africa/Djibouti","Africa/Juba","Africa/Kampala","Africa/Mogadishu","Africa/Nairobi","Indian/Antananarivo","Indian/Comoro","Indian/Mayotte"},
"EDT-0400"    :{ "America/Detroit","America/Fort_Wayne","America/Grand_Turk","America/Indianapolis","America/Iqaluit","America/Louisville","America/Montreal","America/Nassau","America/New_York","America/Nipigon","America/Pangnirtung","America/Port-au-Prince","America/Thunder_Bay","America/Toronto","Canada/Eastern","EST5EDT","US/East-Indiana","US/Eastern","US/Michigan","America/Indiana/Indianapolis","America/Indiana/Marengo","America/Indiana/Petersburg","America/Indiana/Vevay","America/Indiana/Vincennes","America/Indiana/Winamac","America/Kentucky/Louisville","America/Kentucky/Monticello"},
"EEST+0300"   :{ "Asia/Amman","Asia/Beirut","Asia/Damascus","Asia/Famagusta","Asia/Gaza","Asia/Hebron","Asia/Nicosia","Europe/Athens","Europe/Bucharest","Europe/Chisinau","Europe/Helsinki","Europe/Kiev","Europe/Mariehamn","Europe/Nicosia","Europe/Riga","Europe/Sofia","Europe/Tallinn","Europe/Tiraspol","Europe/Uzhgorod","Europe/Vilnius","Europe/Zaporozhye","EET"},
"EET+0200"    :{ "Africa/Cairo","Africa/Tripoli","Asia/Amman","Asia/Beirut","Asia/Damascus","Asia/Famagusta","Asia/Gaza","Asia/Hebron","Asia/Nicosia","Europe/Athens","Europe/Bucharest","Europe/Chisinau","Europe/Helsinki","Europe/Kaliningrad","Europe/Kiev","Europe/Mariehamn","Europe/Nicosia","Europe/Riga","Europe/Sofia","Europe/Tallinn","Europe/Tiraspol","Europe/Uzhgorod","Europe/Vilnius","Europe/Zaporozhye","EET","Egypt","Libya"},
"EST-0500"    :{ "America/Atikokan","America/Cancun","America/Cayman","America/Coral_Harbour","America/Detroit","America/Fort_Wayne","America/Grand_Turk","America/Indianapolis","America/Iqaluit","America/Jamaica","America/Louisville","America/Montreal","America/Nassau","America/New_York","America/Nipigon","America/Panama","America/Pangnirtung","America/Port-au-Prince","America/Thunder_Bay","America/Toronto","Canada/Eastern","EST","EST5EDT","Jamaica","US/East-Indiana","US/Eastern","US/Michigan","America/Indiana/Indianapolis","America/Indiana/Marengo","America/Indiana/Petersburg","America/Indiana/Vevay","America/Indiana/Vincennes","America/Indiana/Winamac","America/Kentucky/Louisville","America/Kentucky/Monticello"},
"GMT+0000"    :{ "Africa/Abidjan","Africa/Accra","Africa/Bamako","Africa/Banjul","Africa/Bissau","Africa/Conakry","Africa/Dakar","Africa/Freetown","Africa/Lome","Africa/Monrovia","Africa/Nouakchott","Africa/Ouagadougou","Africa/Sao_Tome","Africa/Timbuktu","America/Danmarkshavn","Atlantic/Reykjavik","Atlantic/St_Helena","Etc/GMT","Etc/GMT+0","Etc/GMT-0","Etc/GMT0","Etc/Greenwich","Europe/Belfast","Europe/Dublin","Europe/Guernsey","Europe/Isle_of_Man","Europe/Jersey","Europe/London","Eire","GB","GB-Eire","GMT","GMT+0","GMT-0","GMT0","Greenwich","Iceland"},
"HDT-0900"    :{ "America/Adak","America/Atka","US/Aleutian"},
"HKT+0800"    :{ "Asia/Hong_Kong","Hongkong"},
"HST-1000"    :{ "America/Adak","America/Atka","HST","Pacific/Honolulu","Pacific/Johnston","US/Aleutian","US/Hawaii"},
"IDT+0300"    :{ "Asia/Jerusalem","Asia/Tel_Aviv","Israel"},
"IST+0100"    :{ "Europe/Dublin","Eire"},
"IST+0200"    :{ "Asia/Jerusalem","Asia/Tel_Aviv","Israel"},
"IST+0530"    :{ "Asia/Calcutta","Asia/Kolkata"},
"JST+0900"    :{ "Asia/Tokyo","Japan"},
"KST+0900"    :{ "Asia/Pyongyang","Asia/Seoul","ROK"},
"MDT-0600"    :{ "America/Boise","America/Cambridge_Bay","America/Chihuahua","America/Denver","America/Edmonton","America/Inuvik","America/Mazatlan","America/Ojinaga","America/Shiprock","America/Yellowknife","Canada/Mountain","Mexico/BajaSur","MST7MDT","Navajo","US/Mountain"},
"MEST+0200"   :{ "MET"},
"MET+0100"    :{ "MET"},
"MSK+0300"    :{ "Europe/Moscow","Europe/Simferopol","W-SU"},
"MST-0700"    :{ "America/Boise","America/Cambridge_Bay","America/Chihuahua","America/Creston","America/Dawson_Creek","America/Denver","America/Edmonton","America/Fort_Nelson","America/Hermosillo","America/Inuvik","America/Mazatlan","America/Ojinaga","America/Phoenix","America/Shiprock","America/Yellowknife","Canada/Mountain","Mexico/BajaSur","MST","MST7MDT","Navajo","US/Arizona","US/Mountain"},
"NDT-0230"    :{ "America/St_Johns","Canada/Newfoundland"},
"NST-0330"    :{ "America/St_Johns","Canada/Newfoundland"},
"NZDT+1300"   :{ "Antarctica/McMurdo","Antarctica/South_Pole","NZ","Pacific/Auckland"},
"NZST+1200"   :{ "Antarctica/McMurdo","Antarctica/South_Pole","NZ","Pacific/Auckland"},
"PDT-0700"    :{ "America/Dawson","America/Ensenada","America/Los_Angeles","America/Santa_Isabel","America/Tijuana","America/Vancouver","America/Whitehorse","Canada/Pacific","Canada/Yukon","Mexico/BajaNorte","PST8PDT","US/Pacific"},
"PKT+0500"    :{ "Asia/Karachi"},
"PST+0800"    :{ "Asia/Manila"},
"PST-0800"    :{ "America/Dawson","America/Ensenada","America/Los_Angeles","America/Santa_Isabel","America/Tijuana","America/Vancouver","America/Whitehorse","Canada/Pacific","Canada/Yukon","Mexico/BajaNorte","PST8PDT","US/Pacific"},
"SAST+0200"   :{ "Africa/Johannesburg","Africa/Maseru","Africa/Mbabane"},
"SST-1100"    :{ "Pacific/Midway","Pacific/Pago_Pago","Pacific/Samoa","US/Samoa"},
"UTC+0000"    :{ "Etc/UCT","Etc/Universal","Etc/UTC","Etc/Zulu","UCT","Universal","UTC","Zulu"},
"WAT+0100"    :{ "Africa/Bangui","Africa/Brazzaville","Africa/Douala","Africa/Kinshasa","Africa/Lagos","Africa/Libreville","Africa/Luanda","Africa/Malabo","Africa/Ndjamena","Africa/Niamey","Africa/Porto-Novo"},
"WEST+0100"   :{ "Atlantic/Canary","Atlantic/Faeroe","Atlantic/Faroe","Atlantic/Madeira","Europe/Lisbon","Portugal","WET"},
"WET+0000"    :{ "Atlantic/Canary","Atlantic/Faeroe","Atlantic/Faroe","Atlantic/Madeira","Europe/Lisbon","Portugal","WET"},
"WIB+0700"    :{ "Asia/Jakarta","Asia/Pontianak"},
"WIT+0900"    :{ "Asia/Jayapura"},
"WITA+0800"   :{ "Asia/Makassar","Asia/Ujung_Pandang"},
}


// mapTimeZonesToTzAbbrvs - A cross reference that maps
// Time Zone Canonical Values to Time Zone Abbreviations.
// 

var lockMapTimeZonesToTzAbbrvs sync.Mutex

var mapTimeZonesToTzAbbrvs = map[string][]string {
"Africa/Abidjan"                     :{ "GMT+0000"},
"Africa/Accra"                       :{ "GMT+0000"},
"Africa/Addis_Ababa"                 :{ "EAT+0300"},
"Africa/Algiers"                     :{ "CET+0100"},
"Africa/Asmara"                      :{ "EAT+0300"},
"Africa/Asmera"                      :{ "EAT+0300"},
"Africa/Bamako"                      :{ "GMT+0000"},
"Africa/Bangui"                      :{ "WAT+0100"},
"Africa/Banjul"                      :{ "GMT+0000"},
"Africa/Bissau"                      :{ "GMT+0000"},
"Africa/Blantyre"                    :{ "CAT+0200"},
"Africa/Brazzaville"                 :{ "WAT+0100"},
"Africa/Bujumbura"                   :{ "CAT+0200"},
"Africa/Cairo"                       :{ "EET+0200"},
"Africa/Casablanca"                  :{ "+01+0100"},
"Africa/Ceuta"                       :{ "CEST+0200","CET+0100"},
"Africa/Conakry"                     :{ "GMT+0000"},
"Africa/Dakar"                       :{ "GMT+0000"},
"Africa/Dar_es_Salaam"               :{ "EAT+0300"},
"Africa/Djibouti"                    :{ "EAT+0300"},
"Africa/Douala"                      :{ "WAT+0100"},
"Africa/El_Aaiun"                    :{ "+01+0100"},
"Africa/Freetown"                    :{ "GMT+0000"},
"Africa/Gaborone"                    :{ "CAT+0200"},
"Africa/Harare"                      :{ "CAT+0200"},
"Africa/Johannesburg"                :{ "SAST+0200"},
"Africa/Juba"                        :{ "EAT+0300"},
"Africa/Kampala"                     :{ "EAT+0300"},
"Africa/Khartoum"                    :{ "CAT+0200"},
"Africa/Kigali"                      :{ "CAT+0200"},
"Africa/Kinshasa"                    :{ "WAT+0100"},
"Africa/Lagos"                       :{ "WAT+0100"},
"Africa/Libreville"                  :{ "WAT+0100"},
"Africa/Lome"                        :{ "GMT+0000"},
"Africa/Luanda"                      :{ "WAT+0100"},
"Africa/Lubumbashi"                  :{ "CAT+0200"},
"Africa/Lusaka"                      :{ "CAT+0200"},
"Africa/Malabo"                      :{ "WAT+0100"},
"Africa/Maputo"                      :{ "CAT+0200"},
"Africa/Maseru"                      :{ "SAST+0200"},
"Africa/Mbabane"                     :{ "SAST+0200"},
"Africa/Mogadishu"                   :{ "EAT+0300"},
"Africa/Monrovia"                    :{ "GMT+0000"},
"Africa/Nairobi"                     :{ "EAT+0300"},
"Africa/Ndjamena"                    :{ "WAT+0100"},
"Africa/Niamey"                      :{ "WAT+0100"},
"Africa/Nouakchott"                  :{ "GMT+0000"},
"Africa/Ouagadougou"                 :{ "GMT+0000"},
"Africa/Porto-Novo"                  :{ "WAT+0100"},
"Africa/Sao_Tome"                    :{ "GMT+0000"},
"Africa/Timbuktu"                    :{ "GMT+0000"},
"Africa/Tripoli"                     :{ "EET+0200"},
"Africa/Tunis"                       :{ "CET+0100"},
"Africa/Windhoek"                    :{ "CAT+0200"},
"America/Adak"                       :{ "HDT-0900","HST-1000"},
"America/Anchorage"                  :{ "AKDT-0800","AKST-0900"},
"America/Anguilla"                   :{ "AST-0400"},
"America/Antigua"                    :{ "AST-0400"},
"America/Araguaina"                  :{ "-03-0300"},
"America/Argentina/Buenos_Aires"     :{ "-03-0300"},
"America/Argentina/Catamarca"        :{ "-03-0300"},
"America/Argentina/ComodRivadavia"   :{ "-03-0300"},
"America/Argentina/Cordoba"          :{ "-03-0300"},
"America/Argentina/Jujuy"            :{ "-03-0300"},
"America/Argentina/La_Rioja"         :{ "-03-0300"},
"America/Argentina/Mendoza"          :{ "-03-0300"},
"America/Argentina/Rio_Gallegos"     :{ "-03-0300"},
"America/Argentina/Salta"            :{ "-03-0300"},
"America/Argentina/San_Juan"         :{ "-03-0300"},
"America/Argentina/San_Luis"         :{ "-03-0300"},
"America/Argentina/Tucuman"          :{ "-03-0300"},
"America/Argentina/Ushuaia"          :{ "-03-0300"},
"America/Aruba"                      :{ "AST-0400"},
"America/Asuncion"                   :{ "-04-0400"},
"America/Atikokan"                   :{ "EST-0500"},
"America/Atka"                       :{ "HDT-0900","HST-1000"},
"America/Bahia"                      :{ "-03-0300"},
"America/Bahia_Banderas"             :{ "CDT-0500","CST-0600"},
"America/Barbados"                   :{ "AST-0400"},
"America/Belem"                      :{ "-03-0300"},
"America/Belize"                     :{ "CST-0600"},
"America/Blanc-Sablon"               :{ "AST-0400"},
"America/Boa_Vista"                  :{ "-04-0400"},
"America/Bogota"                     :{ "-05-0500"},
"America/Boise"                      :{ "MDT-0600","MST-0700"},
"America/Buenos_Aires"               :{ "-03-0300"},
"America/Cambridge_Bay"              :{ "MDT-0600","MST-0700"},
"America/Campo_Grande"               :{ "-04-0400"},
"America/Cancun"                     :{ "EST-0500"},
"America/Caracas"                    :{ "-04-0400"},
"America/Catamarca"                  :{ "-03-0300"},
"America/Cayenne"                    :{ "-03-0300"},
"America/Cayman"                     :{ "EST-0500"},
"America/Chicago"                    :{ "CDT-0500","CST-0600"},
"America/Chihuahua"                  :{ "MDT-0600","MST-0700"},
"America/Coral_Harbour"              :{ "EST-0500"},
"America/Cordoba"                    :{ "-03-0300"},
"America/Costa_Rica"                 :{ "CST-0600"},
"America/Creston"                    :{ "MST-0700"},
"America/Cuiaba"                     :{ "-04-0400"},
"America/Curacao"                    :{ "AST-0400"},
"America/Danmarkshavn"               :{ "GMT+0000"},
"America/Dawson"                     :{ "PDT-0700","PST-0800"},
"America/Dawson_Creek"               :{ "MST-0700"},
"America/Denver"                     :{ "MDT-0600","MST-0700"},
"America/Detroit"                    :{ "EDT-0400","EST-0500"},
"America/Dominica"                   :{ "AST-0400"},
"America/Edmonton"                   :{ "MDT-0600","MST-0700"},
"America/Eirunepe"                   :{ "-05-0500"},
"America/El_Salvador"                :{ "CST-0600"},
"America/Ensenada"                   :{ "PDT-0700","PST-0800"},
"America/Fort_Nelson"                :{ "MST-0700"},
"America/Fort_Wayne"                 :{ "EDT-0400","EST-0500"},
"America/Fortaleza"                  :{ "-03-0300"},
"America/Glace_Bay"                  :{ "ADT-0300","AST-0400"},
"America/Godthab"                    :{ "-02-0200"},
"America/Goose_Bay"                  :{ "ADT-0300","AST-0400"},
"America/Grand_Turk"                 :{ "EDT-0400","EST-0500"},
"America/Grenada"                    :{ "AST-0400"},
"America/Guadeloupe"                 :{ "AST-0400"},
"America/Guatemala"                  :{ "CST-0600"},
"America/Guayaquil"                  :{ "-05-0500"},
"America/Guyana"                     :{ "-04-0400"},
"America/Halifax"                    :{ "ADT-0300","AST-0400"},
"America/Havana"                     :{ "CDT-0400","CST-0500"},
"America/Hermosillo"                 :{ "MST-0700"},
"America/Indiana/Indianapolis"       :{ "EDT-0400","EST-0500"},
"America/Indiana/Knox"               :{ "CDT-0500","CST-0600"},
"America/Indiana/Marengo"            :{ "EDT-0400","EST-0500"},
"America/Indiana/Petersburg"         :{ "EDT-0400","EST-0500"},
"America/Indiana/Tell_City"          :{ "CDT-0500","CST-0600"},
"America/Indiana/Vevay"              :{ "EDT-0400","EST-0500"},
"America/Indiana/Vincennes"          :{ "EDT-0400","EST-0500"},
"America/Indiana/Winamac"            :{ "EDT-0400","EST-0500"},
"America/Indianapolis"               :{ "EDT-0400","EST-0500"},
"America/Inuvik"                     :{ "MDT-0600","MST-0700"},
"America/Iqaluit"                    :{ "EDT-0400","EST-0500"},
"America/Jamaica"                    :{ "EST-0500"},
"America/Jujuy"                      :{ "-03-0300"},
"America/Juneau"                     :{ "AKDT-0800","AKST-0900"},
"America/Kentucky/Louisville"        :{ "EDT-0400","EST-0500"},
"America/Kentucky/Monticello"        :{ "EDT-0400","EST-0500"},
"America/Knox_IN"                    :{ "CDT-0500","CST-0600"},
"America/Kralendijk"                 :{ "AST-0400"},
"America/La_Paz"                     :{ "-04-0400"},
"America/Lima"                       :{ "-05-0500"},
"America/Los_Angeles"                :{ "PDT-0700","PST-0800"},
"America/Louisville"                 :{ "EDT-0400","EST-0500"},
"America/Lower_Princes"              :{ "AST-0400"},
"America/Maceio"                     :{ "-03-0300"},
"America/Managua"                    :{ "CST-0600"},
"America/Manaus"                     :{ "-04-0400"},
"America/Marigot"                    :{ "AST-0400"},
"America/Martinique"                 :{ "AST-0400"},
"America/Matamoros"                  :{ "CDT-0500","CST-0600"},
"America/Mazatlan"                   :{ "MDT-0600","MST-0700"},
"America/Mendoza"                    :{ "-03-0300"},
"America/Menominee"                  :{ "CDT-0500","CST-0600"},
"America/Merida"                     :{ "CDT-0500","CST-0600"},
"America/Metlakatla"                 :{ "AKDT-0800","AKST-0900"},
"America/Mexico_City"                :{ "CDT-0500","CST-0600"},
"America/Miquelon"                   :{ "-02-0200"},
"America/Moncton"                    :{ "ADT-0300","AST-0400"},
"America/Monterrey"                  :{ "CDT-0500","CST-0600"},
"America/Montevideo"                 :{ "-03-0300"},
"America/Montreal"                   :{ "EDT-0400","EST-0500"},
"America/Montserrat"                 :{ "AST-0400"},
"America/Nassau"                     :{ "EDT-0400","EST-0500"},
"America/New_York"                   :{ "EDT-0400","EST-0500"},
"America/Nipigon"                    :{ "EDT-0400","EST-0500"},
"America/Nome"                       :{ "AKDT-0800","AKST-0900"},
"America/Noronha"                    :{ "-02-0200"},
"America/North_Dakota/Beulah"        :{ "CDT-0500","CST-0600"},
"America/North_Dakota/Center"        :{ "CDT-0500","CST-0600"},
"America/North_Dakota/New_Salem"     :{ "CDT-0500","CST-0600"},
"America/Ojinaga"                    :{ "MDT-0600","MST-0700"},
"America/Panama"                     :{ "EST-0500"},
"America/Pangnirtung"                :{ "EDT-0400","EST-0500"},
"America/Paramaribo"                 :{ "-03-0300"},
"America/Phoenix"                    :{ "MST-0700"},
"America/Port-au-Prince"             :{ "EDT-0400","EST-0500"},
"America/Port_of_Spain"              :{ "AST-0400"},
"America/Porto_Acre"                 :{ "-05-0500"},
"America/Porto_Velho"                :{ "-04-0400"},
"America/Puerto_Rico"                :{ "AST-0400"},
"America/Punta_Arenas"               :{ "-03-0300"},
"America/Rainy_River"                :{ "CDT-0500","CST-0600"},
"America/Rankin_Inlet"               :{ "CDT-0500","CST-0600"},
"America/Recife"                     :{ "-03-0300"},
"America/Regina"                     :{ "CST-0600"},
"America/Resolute"                   :{ "CDT-0500","CST-0600"},
"America/Rio_Branco"                 :{ "-05-0500"},
"America/Rosario"                    :{ "-03-0300"},
"America/Santa_Isabel"               :{ "PDT-0700","PST-0800"},
"America/Santarem"                   :{ "-03-0300"},
"America/Santiago"                   :{ "-04-0400"},
"America/Santo_Domingo"              :{ "AST-0400"},
"America/Sao_Paulo"                  :{ "-03-0300"},
"America/Scoresbysund"               :{ "+00+0000"},
"America/Shiprock"                   :{ "MDT-0600","MST-0700"},
"America/Sitka"                      :{ "AKDT-0800","AKST-0900"},
"America/St_Barthelemy"              :{ "AST-0400"},
"America/St_Johns"                   :{ "NDT-0230","NST-0330"},
"America/St_Kitts"                   :{ "AST-0400"},
"America/St_Lucia"                   :{ "AST-0400"},
"America/St_Thomas"                  :{ "AST-0400"},
"America/St_Vincent"                 :{ "AST-0400"},
"America/Swift_Current"              :{ "CST-0600"},
"America/Tegucigalpa"                :{ "CST-0600"},
"America/Thule"                      :{ "ADT-0300","AST-0400"},
"America/Thunder_Bay"                :{ "EDT-0400","EST-0500"},
"America/Tijuana"                    :{ "PDT-0700","PST-0800"},
"America/Toronto"                    :{ "EDT-0400","EST-0500"},
"America/Tortola"                    :{ "AST-0400"},
"America/Vancouver"                  :{ "PDT-0700","PST-0800"},
"America/Virgin"                     :{ "AST-0400"},
"America/Whitehorse"                 :{ "PDT-0700","PST-0800"},
"America/Winnipeg"                   :{ "CDT-0500","CST-0600"},
"America/Yakutat"                    :{ "AKDT-0800","AKST-0900"},
"America/Yellowknife"                :{ "MDT-0600","MST-0700"},
"Antarctica/Casey"                   :{ "+08+0800"},
"Antarctica/Davis"                   :{ "+07+0700"},
"Antarctica/DumontDUrville"          :{ "+10+1000"},
"Antarctica/Macquarie"               :{ "+11+1100"},
"Antarctica/Mawson"                  :{ "+05+0500"},
"Antarctica/McMurdo"                 :{ "NZST+1200","NZDT+1300"},
"Antarctica/Palmer"                  :{ "-03-0300"},
"Antarctica/Rothera"                 :{ "-03-0300"},
"Antarctica/South_Pole"              :{ "NZST+1200","NZDT+1300"},
"Antarctica/Syowa"                   :{ "+03+0300"},
"Antarctica/Troll"                   :{ "+02+0200"},
"Antarctica/Vostok"                  :{ "+06+0600"},
"Arctic/Longyearbyen"                :{ "CEST+0200","CET+0100"},
"Asia/Aden"                          :{ "+03+0300"},
"Asia/Almaty"                        :{ "+06+0600"},
"Asia/Amman"                         :{ "EEST+0300","EET+0200"},
"Asia/Anadyr"                        :{ "+12+1200"},
"Asia/Aqtau"                         :{ "+05+0500"},
"Asia/Aqtobe"                        :{ "+05+0500"},
"Asia/Ashgabat"                      :{ "+05+0500"},
"Asia/Ashkhabad"                     :{ "+05+0500"},
"Asia/Atyrau"                        :{ "+05+0500"},
"Asia/Baghdad"                       :{ "+03+0300"},
"Asia/Bahrain"                       :{ "+03+0300"},
"Asia/Baku"                          :{ "+04+0400"},
"Asia/Bangkok"                       :{ "+07+0700"},
"Asia/Barnaul"                       :{ "+07+0700"},
"Asia/Beirut"                        :{ "EEST+0300","EET+0200"},
"Asia/Bishkek"                       :{ "+06+0600"},
"Asia/Brunei"                        :{ "+08+0800"},
"Asia/Calcutta"                      :{ "IST+0530"},
"Asia/Chita"                         :{ "+09+0900"},
"Asia/Choibalsan"                    :{ "+08+0800"},
"Asia/Chongqing"                     :{ "CST+0800"},
"Asia/Chungking"                     :{ "CST+0800"},
"Asia/Colombo"                       :{ "+0530+0530"},
"Asia/Dacca"                         :{ "+06+0600"},
"Asia/Damascus"                      :{ "EEST+0300","EET+0200"},
"Asia/Dhaka"                         :{ "+06+0600"},
"Asia/Dili"                          :{ "+09+0900"},
"Asia/Dubai"                         :{ "+04+0400"},
"Asia/Dushanbe"                      :{ "+05+0500"},
"Asia/Famagusta"                     :{ "EEST+0300","EET+0200"},
"Asia/Gaza"                          :{ "EEST+0300","EET+0200"},
"Asia/Harbin"                        :{ "CST+0800"},
"Asia/Hebron"                        :{ "EEST+0300","EET+0200"},
"Asia/Ho_Chi_Minh"                   :{ "+07+0700"},
"Asia/Hong_Kong"                     :{ "HKT+0800"},
"Asia/Hovd"                          :{ "+07+0700"},
"Asia/Irkutsk"                       :{ "+08+0800"},
"Asia/Istanbul"                      :{ "+03+0300"},
"Asia/Jakarta"                       :{ "WIB+0700"},
"Asia/Jayapura"                      :{ "WIT+0900"},
"Asia/Jerusalem"                     :{ "IDT+0300","IST+0200"},
"Asia/Kabul"                         :{ "+0430+0430"},
"Asia/Kamchatka"                     :{ "+12+1200"},
"Asia/Karachi"                       :{ "PKT+0500"},
"Asia/Kashgar"                       :{ "+06+0600"},
"Asia/Kathmandu"                     :{ "+0545+0545"},
"Asia/Katmandu"                      :{ "+0545+0545"},
"Asia/Khandyga"                      :{ "+09+0900"},
"Asia/Kolkata"                       :{ "IST+0530"},
"Asia/Krasnoyarsk"                   :{ "+07+0700"},
"Asia/Kuala_Lumpur"                  :{ "+08+0800"},
"Asia/Kuching"                       :{ "+08+0800"},
"Asia/Kuwait"                        :{ "+03+0300"},
"Asia/Macao"                         :{ "CST+0800"},
"Asia/Macau"                         :{ "CST+0800"},
"Asia/Magadan"                       :{ "+11+1100"},
"Asia/Makassar"                      :{ "WITA+0800"},
"Asia/Manila"                        :{ "PST+0800"},
"Asia/Muscat"                        :{ "+04+0400"},
"Asia/Nicosia"                       :{ "EEST+0300","EET+0200"},
"Asia/Novokuznetsk"                  :{ "+07+0700"},
"Asia/Novosibirsk"                   :{ "+07+0700"},
"Asia/Omsk"                          :{ "+06+0600"},
"Asia/Oral"                          :{ "+05+0500"},
"Asia/Phnom_Penh"                    :{ "+07+0700"},
"Asia/Pontianak"                     :{ "WIB+0700"},
"Asia/Pyongyang"                     :{ "KST+0900"},
"Asia/Qatar"                         :{ "+03+0300"},
"Asia/Qostanay"                      :{ "+06+0600"},
"Asia/Qyzylorda"                     :{ "+05+0500"},
"Asia/Rangoon"                       :{ "+0630+0630"},
"Asia/Riyadh"                        :{ "+03+0300"},
"Asia/Saigon"                        :{ "+07+0700"},
"Asia/Sakhalin"                      :{ "+11+1100"},
"Asia/Samarkand"                     :{ "+05+0500"},
"Asia/Seoul"                         :{ "KST+0900"},
"Asia/Shanghai"                      :{ "CST+0800"},
"Asia/Singapore"                     :{ "+08+0800"},
"Asia/Srednekolymsk"                 :{ "+11+1100"},
"Asia/Taipei"                        :{ "CST+0800"},
"Asia/Tashkent"                      :{ "+05+0500"},
"Asia/Tbilisi"                       :{ "+04+0400"},
"Asia/Tehran"                        :{ "+0430+0430"},
"Asia/Tel_Aviv"                      :{ "IDT+0300","IST+0200"},
"Asia/Thimbu"                        :{ "+06+0600"},
"Asia/Thimphu"                       :{ "+06+0600"},
"Asia/Tokyo"                         :{ "JST+0900"},
"Asia/Tomsk"                         :{ "+07+0700"},
"Asia/Ujung_Pandang"                 :{ "WITA+0800"},
"Asia/Ulaanbaatar"                   :{ "+08+0800"},
"Asia/Ulan_Bator"                    :{ "+08+0800"},
"Asia/Urumqi"                        :{ "+06+0600"},
"Asia/Ust-Nera"                      :{ "+10+1000"},
"Asia/Vientiane"                     :{ "+07+0700"},
"Asia/Vladivostok"                   :{ "+10+1000"},
"Asia/Yakutsk"                       :{ "+09+0900"},
"Asia/Yangon"                        :{ "+0630+0630"},
"Asia/Yekaterinburg"                 :{ "+05+0500"},
"Asia/Yerevan"                       :{ "+04+0400"},
"Atlantic/Azores"                    :{ "+00+0000"},
"Atlantic/Bermuda"                   :{ "ADT-0300","AST-0400"},
"Atlantic/Canary"                    :{ "WEST+0100","WET+0000"},
"Atlantic/Cape_Verde"                :{ "-01-0100"},
"Atlantic/Faeroe"                    :{ "WEST+0100","WET+0000"},
"Atlantic/Faroe"                     :{ "WEST+0100","WET+0000"},
"Atlantic/Jan_Mayen"                 :{ "CEST+0200","CET+0100"},
"Atlantic/Madeira"                   :{ "WEST+0100","WET+0000"},
"Atlantic/Reykjavik"                 :{ "GMT+0000"},
"Atlantic/South_Georgia"             :{ "-02-0200"},
"Atlantic/St_Helena"                 :{ "GMT+0000"},
"Atlantic/Stanley"                   :{ "-03-0300"},
"Australia/ACT"                      :{ "AEST+1000","AEDT+1100"},
"Australia/Adelaide"                 :{ "ACST+0930","ACDT+1030"},
"Australia/Brisbane"                 :{ "AEST+1000"},
"Australia/Broken_Hill"              :{ "ACST+0930","ACDT+1030"},
"Australia/Canberra"                 :{ "AEST+1000","AEDT+1100"},
"Australia/Currie"                   :{ "AEST+1000","AEDT+1100"},
"Australia/Darwin"                   :{ "ACST+0930"},
"Australia/Eucla"                    :{ "+0845+0845"},
"Australia/Hobart"                   :{ "AEST+1000","AEDT+1100"},
"Australia/LHI"                      :{ "+1030+1030"},
"Australia/Lindeman"                 :{ "AEST+1000"},
"Australia/Lord_Howe"                :{ "+1030+1030"},
"Australia/Melbourne"                :{ "AEST+1000","AEDT+1100"},
"Australia/NSW"                      :{ "AEST+1000","AEDT+1100"},
"Australia/North"                    :{ "ACST+0930"},
"Australia/Perth"                    :{ "AWST+0800"},
"Australia/Queensland"               :{ "AEST+1000"},
"Australia/South"                    :{ "ACST+0930","ACDT+1030"},
"Australia/Sydney"                   :{ "AEST+1000","AEDT+1100"},
"Australia/Tasmania"                 :{ "AEST+1000","AEDT+1100"},
"Australia/Victoria"                 :{ "AEST+1000","AEDT+1100"},
"Australia/West"                     :{ "AWST+0800"},
"Australia/Yancowinna"               :{ "ACST+0930","ACDT+1030"},
"Brazil/Acre"                        :{ "-05-0500"},
"Brazil/DeNoronha"                   :{ "-02-0200"},
"Brazil/East"                        :{ "-03-0300"},
"Brazil/West"                        :{ "-04-0400"},
"CET"                                :{ "CEST+0200","CET+0100"},
"CST6CDT"                            :{ "CDT-0500","CST-0600"},
"Canada/Atlantic"                    :{ "ADT-0300","AST-0400"},
"Canada/Central"                     :{ "CDT-0500","CST-0600"},
"Canada/Eastern"                     :{ "EDT-0400","EST-0500"},
"Canada/Mountain"                    :{ "MDT-0600","MST-0700"},
"Canada/Newfoundland"                :{ "NDT-0230","NST-0330"},
"Canada/Pacific"                     :{ "PDT-0700","PST-0800"},
"Canada/Saskatchewan"                :{ "CST-0600"},
"Canada/Yukon"                       :{ "PDT-0700","PST-0800"},
"Chile/Continental"                  :{ "-04-0400"},
"Chile/EasterIsland"                 :{ "-06-0600"},
"Cuba"                               :{ "CDT-0400","CST-0500"},
"EET"                                :{ "EEST+0300","EET+0200"},
"EST"                                :{ "EST-0500"},
"EST5EDT"                            :{ "EDT-0400","EST-0500"},
"Egypt"                              :{ "EET+0200"},
"Eire"                               :{ "IST+0100","GMT+0000"},
"Etc/GMT"                            :{ "GMT+0000"},
"Etc/GMT+0"                          :{ "GMT+0000"},
"Etc/GMT+1"                          :{ "-01-0100"},
"Etc/GMT+10"                         :{ "-10-1000"},
"Etc/GMT+11"                         :{ "-11-1100"},
"Etc/GMT+12"                         :{ "-12-1200"},
"Etc/GMT+2"                          :{ "-02-0200"},
"Etc/GMT+3"                          :{ "-03-0300"},
"Etc/GMT+4"                          :{ "-04-0400"},
"Etc/GMT+5"                          :{ "-05-0500"},
"Etc/GMT+6"                          :{ "-06-0600"},
"Etc/GMT+7"                          :{ "-07-0700"},
"Etc/GMT+8"                          :{ "-08-0800"},
"Etc/GMT+9"                          :{ "-09-0900"},
"Etc/GMT-0"                          :{ "GMT+0000"},
"Etc/GMT-1"                          :{ "+01+0100"},
"Etc/GMT-10"                         :{ "+10+1000"},
"Etc/GMT-11"                         :{ "+11+1100"},
"Etc/GMT-12"                         :{ "+12+1200"},
"Etc/GMT-13"                         :{ "+13+1300"},
"Etc/GMT-14"                         :{ "+14+1400"},
"Etc/GMT-2"                          :{ "+02+0200"},
"Etc/GMT-3"                          :{ "+03+0300"},
"Etc/GMT-4"                          :{ "+04+0400"},
"Etc/GMT-5"                          :{ "+05+0500"},
"Etc/GMT-6"                          :{ "+06+0600"},
"Etc/GMT-7"                          :{ "+07+0700"},
"Etc/GMT-8"                          :{ "+08+0800"},
"Etc/GMT-9"                          :{ "+09+0900"},
"Etc/GMT0"                           :{ "GMT+0000"},
"Etc/Greenwich"                      :{ "GMT+0000"},
"Etc/UCT"                            :{ "UTC+0000"},
"Etc/UTC"                            :{ "UTC+0000"},
"Etc/Universal"                      :{ "UTC+0000"},
"Etc/Zulu"                           :{ "UTC+0000"},
"Europe/Amsterdam"                   :{ "CEST+0200","CET+0100"},
"Europe/Andorra"                     :{ "CEST+0200","CET+0100"},
"Europe/Astrakhan"                   :{ "+04+0400"},
"Europe/Athens"                      :{ "EEST+0300","EET+0200"},
"Europe/Belfast"                     :{ "BST+0100","GMT+0000"},
"Europe/Belgrade"                    :{ "CEST+0200","CET+0100"},
"Europe/Berlin"                      :{ "CEST+0200","CET+0100"},
"Europe/Bratislava"                  :{ "CEST+0200","CET+0100"},
"Europe/Brussels"                    :{ "CEST+0200","CET+0100"},
"Europe/Bucharest"                   :{ "EEST+0300","EET+0200"},
"Europe/Budapest"                    :{ "CEST+0200","CET+0100"},
"Europe/Busingen"                    :{ "CEST+0200","CET+0100"},
"Europe/Chisinau"                    :{ "EEST+0300","EET+0200"},
"Europe/Copenhagen"                  :{ "CEST+0200","CET+0100"},
"Europe/Dublin"                      :{ "IST+0100","GMT+0000"},
"Europe/Gibraltar"                   :{ "CEST+0200","CET+0100"},
"Europe/Guernsey"                    :{ "BST+0100","GMT+0000"},
"Europe/Helsinki"                    :{ "EEST+0300","EET+0200"},
"Europe/Isle_of_Man"                 :{ "BST+0100","GMT+0000"},
"Europe/Istanbul"                    :{ "+03+0300"},
"Europe/Jersey"                      :{ "BST+0100","GMT+0000"},
"Europe/Kaliningrad"                 :{ "EET+0200"},
"Europe/Kiev"                        :{ "EEST+0300","EET+0200"},
"Europe/Kirov"                       :{ "+03+0300"},
"Europe/Lisbon"                      :{ "WEST+0100","WET+0000"},
"Europe/Ljubljana"                   :{ "CEST+0200","CET+0100"},
"Europe/London"                      :{ "BST+0100","GMT+0000"},
"Europe/Luxembourg"                  :{ "CEST+0200","CET+0100"},
"Europe/Madrid"                      :{ "CEST+0200","CET+0100"},
"Europe/Malta"                       :{ "CEST+0200","CET+0100"},
"Europe/Mariehamn"                   :{ "EEST+0300","EET+0200"},
"Europe/Minsk"                       :{ "+03+0300"},
"Europe/Monaco"                      :{ "CEST+0200","CET+0100"},
"Europe/Moscow"                      :{ "MSK+0300"},
"Europe/Nicosia"                     :{ "EEST+0300","EET+0200"},
"Europe/Oslo"                        :{ "CEST+0200","CET+0100"},
"Europe/Paris"                       :{ "CEST+0200","CET+0100"},
"Europe/Podgorica"                   :{ "CEST+0200","CET+0100"},
"Europe/Prague"                      :{ "CEST+0200","CET+0100"},
"Europe/Riga"                        :{ "EEST+0300","EET+0200"},
"Europe/Rome"                        :{ "CEST+0200","CET+0100"},
"Europe/Samara"                      :{ "+04+0400"},
"Europe/San_Marino"                  :{ "CEST+0200","CET+0100"},
"Europe/Sarajevo"                    :{ "CEST+0200","CET+0100"},
"Europe/Saratov"                     :{ "+04+0400"},
"Europe/Simferopol"                  :{ "MSK+0300"},
"Europe/Skopje"                      :{ "CEST+0200","CET+0100"},
"Europe/Sofia"                       :{ "EEST+0300","EET+0200"},
"Europe/Stockholm"                   :{ "CEST+0200","CET+0100"},
"Europe/Tallinn"                     :{ "EEST+0300","EET+0200"},
"Europe/Tirane"                      :{ "CEST+0200","CET+0100"},
"Europe/Tiraspol"                    :{ "EEST+0300","EET+0200"},
"Europe/Ulyanovsk"                   :{ "+04+0400"},
"Europe/Uzhgorod"                    :{ "EEST+0300","EET+0200"},
"Europe/Vaduz"                       :{ "CEST+0200","CET+0100"},
"Europe/Vatican"                     :{ "CEST+0200","CET+0100"},
"Europe/Vienna"                      :{ "CEST+0200","CET+0100"},
"Europe/Vilnius"                     :{ "EEST+0300","EET+0200"},
"Europe/Volgograd"                   :{ "+04+0400"},
"Europe/Warsaw"                      :{ "CEST+0200","CET+0100"},
"Europe/Zagreb"                      :{ "CEST+0200","CET+0100"},
"Europe/Zaporozhye"                  :{ "EEST+0300","EET+0200"},
"Europe/Zurich"                      :{ "CEST+0200","CET+0100"},
"Factory"                            :{ "-00+0000"},
"GB"                                 :{ "BST+0100","GMT+0000"},
"GB-Eire"                            :{ "BST+0100","GMT+0000"},
"GMT"                                :{ "GMT+0000"},
"GMT+0"                              :{ "GMT+0000"},
"GMT-0"                              :{ "GMT+0000"},
"GMT0"                               :{ "GMT+0000"},
"Greenwich"                          :{ "GMT+0000"},
"HST"                                :{ "HST-1000"},
"Hongkong"                           :{ "HKT+0800"},
"Iceland"                            :{ "GMT+0000"},
"Indian/Antananarivo"                :{ "EAT+0300"},
"Indian/Chagos"                      :{ "+06+0600"},
"Indian/Christmas"                   :{ "+07+0700"},
"Indian/Cocos"                       :{ "+0630+0630"},
"Indian/Comoro"                      :{ "EAT+0300"},
"Indian/Kerguelen"                   :{ "+05+0500"},
"Indian/Mahe"                        :{ "+04+0400"},
"Indian/Maldives"                    :{ "+05+0500"},
"Indian/Mauritius"                   :{ "+04+0400"},
"Indian/Mayotte"                     :{ "EAT+0300"},
"Indian/Reunion"                     :{ "+04+0400"},
"Iran"                               :{ "+0430+0430"},
"Israel"                             :{ "IDT+0300","IST+0200"},
"Jamaica"                            :{ "EST-0500"},
"Japan"                              :{ "JST+0900"},
"Kwajalein"                          :{ "+12+1200"},
"Libya"                              :{ "EET+0200"},
"MET"                                :{ "MEST+0200","MET+0100"},
"MST"                                :{ "MST-0700"},
"MST7MDT"                            :{ "MDT-0600","MST-0700"},
"Mexico/BajaNorte"                   :{ "PDT-0700","PST-0800"},
"Mexico/BajaSur"                     :{ "MDT-0600","MST-0700"},
"Mexico/General"                     :{ "CDT-0500","CST-0600"},
"NZ"                                 :{ "NZST+1200","NZDT+1300"},
"NZ-CHAT"                            :{ "+1245+1245"},
"Navajo"                             :{ "MDT-0600","MST-0700"},
"PRC"                                :{ "CST+0800"},
"PST8PDT"                            :{ "PDT-0700","PST-0800"},
"Pacific/Apia"                       :{ "+13+1300"},
"Pacific/Auckland"                   :{ "NZST+1200","NZDT+1300"},
"Pacific/Bougainville"               :{ "+11+1100"},
"Pacific/Chatham"                    :{ "+1245+1245"},
"Pacific/Chuuk"                      :{ "+10+1000"},
"Pacific/Easter"                     :{ "-06-0600"},
"Pacific/Efate"                      :{ "+11+1100"},
"Pacific/Enderbury"                  :{ "+13+1300"},
"Pacific/Fakaofo"                    :{ "+13+1300"},
"Pacific/Fiji"                       :{ "+12+1200"},
"Pacific/Funafuti"                   :{ "+12+1200"},
"Pacific/Galapagos"                  :{ "-06-0600"},
"Pacific/Gambier"                    :{ "-09-0900"},
"Pacific/Guadalcanal"                :{ "+11+1100"},
"Pacific/Guam"                       :{ "ChST+1000"},
"Pacific/Honolulu"                   :{ "HST-1000"},
"Pacific/Johnston"                   :{ "HST-1000"},
"Pacific/Kiritimati"                 :{ "+14+1400"},
"Pacific/Kosrae"                     :{ "+11+1100"},
"Pacific/Kwajalein"                  :{ "+12+1200"},
"Pacific/Majuro"                     :{ "+12+1200"},
"Pacific/Marquesas"                  :{ "-0930-0930"},
"Pacific/Midway"                     :{ "SST-1100"},
"Pacific/Nauru"                      :{ "+12+1200"},
"Pacific/Niue"                       :{ "-11-1100"},
"Pacific/Norfolk"                    :{ "+11+1100"},
"Pacific/Noumea"                     :{ "+11+1100"},
"Pacific/Pago_Pago"                  :{ "SST-1100"},
"Pacific/Palau"                      :{ "+09+0900"},
"Pacific/Pitcairn"                   :{ "-08-0800"},
"Pacific/Pohnpei"                    :{ "+11+1100"},
"Pacific/Ponape"                     :{ "+11+1100"},
"Pacific/Port_Moresby"               :{ "+10+1000"},
"Pacific/Rarotonga"                  :{ "-10-1000"},
"Pacific/Saipan"                     :{ "ChST+1000"},
"Pacific/Samoa"                      :{ "SST-1100"},
"Pacific/Tahiti"                     :{ "-10-1000"},
"Pacific/Tarawa"                     :{ "+12+1200"},
"Pacific/Tongatapu"                  :{ "+13+1300"},
"Pacific/Truk"                       :{ "+10+1000"},
"Pacific/Wake"                       :{ "+12+1200"},
"Pacific/Wallis"                     :{ "+12+1200"},
"Pacific/Yap"                        :{ "+10+1000"},
"Poland"                             :{ "CEST+0200","CET+0100"},
"Portugal"                           :{ "WEST+0100","WET+0000"},
"ROC"                                :{ "CST+0800"},
"ROK"                                :{ "KST+0900"},
"Singapore"                          :{ "+08+0800"},
"Turkey"                             :{ "+03+0300"},
"UCT"                                :{ "UTC+0000"},
"US/Alaska"                          :{ "AKDT-0800","AKST-0900"},
"US/Aleutian"                        :{ "HDT-0900","HST-1000"},
"US/Arizona"                         :{ "MST-0700"},
"US/Central"                         :{ "CDT-0500","CST-0600"},
"US/East-Indiana"                    :{ "EDT-0400","EST-0500"},
"US/Eastern"                         :{ "EDT-0400","EST-0500"},
"US/Hawaii"                          :{ "HST-1000"},
"US/Indiana-Starke"                  :{ "CDT-0500","CST-0600"},
"US/Michigan"                        :{ "EDT-0400","EST-0500"},
"US/Mountain"                        :{ "MDT-0600","MST-0700"},
"US/Pacific"                         :{ "PDT-0700","PST-0800"},
"US/Samoa"                           :{ "SST-1100"},
"UTC"                                :{ "UTC+0000"},
"Universal"                          :{ "UTC+0000"},
"W-SU"                               :{ "MSK+0300"},
"WET"                                :{ "WEST+0100","WET+0000"},
"Zulu"                               :{ "UTC+0000"},
}


