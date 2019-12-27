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
"Africa/Ceuta"                       :{ "CEST+0200","CET+0100"},
"Africa/Conakry"                     :{ "GMT+0000"},
"Africa/Dakar"                       :{ "GMT+0000"},
"Africa/Dar_es_Salaam"               :{ "EAT+0300"},
"Africa/Djibouti"                    :{ "EAT+0300"},
"Africa/Douala"                      :{ "WAT+0100"},
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
"America/Aruba"                      :{ "AST-0400"},
"America/Atikokan"                   :{ "EST-0500"},
"America/Atka"                       :{ "HDT-0900","HST-1000"},
"America/Bahia_Banderas"             :{ "CDT-0500","CST-0600"},
"America/Barbados"                   :{ "AST-0400"},
"America/Belize"                     :{ "CST-0600"},
"America/Blanc-Sablon"               :{ "AST-0400"},
"America/Boise"                      :{ "MDT-0600","MST-0700"},
"America/Cambridge_Bay"              :{ "MDT-0600","MST-0700"},
"America/Cancun"                     :{ "EST-0500"},
"America/Cayman"                     :{ "EST-0500"},
"America/Chicago"                    :{ "CDT-0500","CST-0600"},
"America/Chihuahua"                  :{ "MDT-0600","MST-0700"},
"America/Coral_Harbour"              :{ "EST-0500"},
"America/Costa_Rica"                 :{ "CST-0600"},
"America/Creston"                    :{ "MST-0700"},
"America/Curacao"                    :{ "AST-0400"},
"America/Danmarkshavn"               :{ "GMT+0000"},
"America/Dawson"                     :{ "PDT-0700","PST-0800"},
"America/Dawson_Creek"               :{ "MST-0700"},
"America/Denver"                     :{ "MDT-0600","MST-0700"},
"America/Detroit"                    :{ "EDT-0400","EST-0500"},
"America/Dominica"                   :{ "AST-0400"},
"America/Edmonton"                   :{ "MDT-0600","MST-0700"},
"America/El_Salvador"                :{ "CST-0600"},
"America/Ensenada"                   :{ "PDT-0700","PST-0800"},
"America/Fort_Nelson"                :{ "MST-0700"},
"America/Fort_Wayne"                 :{ "EDT-0400","EST-0500"},
"America/Glace_Bay"                  :{ "ADT-0300","AST-0400"},
"America/Goose_Bay"                  :{ "ADT-0300","AST-0400"},
"America/Grand_Turk"                 :{ "EDT-0400","EST-0500"},
"America/Grenada"                    :{ "AST-0400"},
"America/Guadeloupe"                 :{ "AST-0400"},
"America/Guatemala"                  :{ "CST-0600"},
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
"America/Juneau"                     :{ "AKDT-0800","AKST-0900"},
"America/Kentucky/Louisville"        :{ "EDT-0400","EST-0500"},
"America/Kentucky/Monticello"        :{ "EDT-0400","EST-0500"},
"America/Knox_IN"                    :{ "CDT-0500","CST-0600"},
"America/Kralendijk"                 :{ "AST-0400"},
"America/Los_Angeles"                :{ "PDT-0700","PST-0800"},
"America/Louisville"                 :{ "EDT-0400","EST-0500"},
"America/Lower_Princes"              :{ "AST-0400"},
"America/Managua"                    :{ "CST-0600"},
"America/Marigot"                    :{ "AST-0400"},
"America/Martinique"                 :{ "AST-0400"},
"America/Matamoros"                  :{ "CDT-0500","CST-0600"},
"America/Mazatlan"                   :{ "MDT-0600","MST-0700"},
"America/Menominee"                  :{ "CDT-0500","CST-0600"},
"America/Merida"                     :{ "CDT-0500","CST-0600"},
"America/Metlakatla"                 :{ "AKDT-0800","AKST-0900"},
"America/Mexico_City"                :{ "CDT-0500","CST-0600"},
"America/Moncton"                    :{ "ADT-0300","AST-0400"},
"America/Monterrey"                  :{ "CDT-0500","CST-0600"},
"America/Montreal"                   :{ "EDT-0400","EST-0500"},
"America/Montserrat"                 :{ "AST-0400"},
"America/Nassau"                     :{ "EDT-0400","EST-0500"},
"America/New_York"                   :{ "EDT-0400","EST-0500"},
"America/Nipigon"                    :{ "EDT-0400","EST-0500"},
"America/Nome"                       :{ "AKDT-0800","AKST-0900"},
"America/North_Dakota/Beulah"        :{ "CDT-0500","CST-0600"},
"America/North_Dakota/Center"        :{ "CDT-0500","CST-0600"},
"America/North_Dakota/New_Salem"     :{ "CDT-0500","CST-0600"},
"America/Ojinaga"                    :{ "MDT-0600","MST-0700"},
"America/Panama"                     :{ "EST-0500"},
"America/Pangnirtung"                :{ "EDT-0400","EST-0500"},
"America/Phoenix"                    :{ "MST-0700"},
"America/Port-au-Prince"             :{ "EDT-0400","EST-0500"},
"America/Port_of_Spain"              :{ "AST-0400"},
"America/Puerto_Rico"                :{ "AST-0400"},
"America/Rainy_River"                :{ "CDT-0500","CST-0600"},
"America/Rankin_Inlet"               :{ "CDT-0500","CST-0600"},
"America/Regina"                     :{ "CST-0600"},
"America/Resolute"                   :{ "CDT-0500","CST-0600"},
"America/Santa_Isabel"               :{ "PDT-0700","PST-0800"},
"America/Santo_Domingo"              :{ "AST-0400"},
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
"Antarctica/McMurdo"                 :{ "NZST+1200","NZDT+1300"},
"Antarctica/South_Pole"              :{ "NZST+1200","NZDT+1300"},
"Arctic/Longyearbyen"                :{ "CEST+0200","CET+0100"},
"Asia/Amman"                         :{ "EEST+0300","EET+0200"},
"Asia/Beirut"                        :{ "EEST+0300","EET+0200"},
"Asia/Calcutta"                      :{ "IST+0530"},
"Asia/Chongqing"                     :{ "CST+0800"},
"Asia/Chungking"                     :{ "CST+0800"},
"Asia/Damascus"                      :{ "EEST+0300","EET+0200"},
"Asia/Famagusta"                     :{ "EEST+0300","EET+0200"},
"Asia/Gaza"                          :{ "EEST+0300","EET+0200"},
"Asia/Harbin"                        :{ "CST+0800"},
"Asia/Hebron"                        :{ "EEST+0300","EET+0200"},
"Asia/Hong_Kong"                     :{ "HKT+0800"},
"Asia/Jakarta"                       :{ "WIB+0700"},
"Asia/Jayapura"                      :{ "WIT+0900"},
"Asia/Jerusalem"                     :{ "IDT+0300","IST+0200"},
"Asia/Karachi"                       :{ "PKT+0500"},
"Asia/Kolkata"                       :{ "IST+0530"},
"Asia/Macao"                         :{ "CST+0800"},
"Asia/Macau"                         :{ "CST+0800"},
"Asia/Makassar"                      :{ "WITA+0800"},
"Asia/Manila"                        :{ "PST+0800"},
"Asia/Nicosia"                       :{ "EEST+0300","EET+0200"},
"Asia/Pontianak"                     :{ "WIB+0700"},
"Asia/Pyongyang"                     :{ "KST+0900"},
"Asia/Seoul"                         :{ "KST+0900"},
"Asia/Shanghai"                      :{ "CST+0800"},
"Asia/Taipei"                        :{ "CST+0800"},
"Asia/Tel_Aviv"                      :{ "IDT+0300","IST+0200"},
"Asia/Tokyo"                         :{ "JST+0900"},
"Asia/Ujung_Pandang"                 :{ "WITA+0800"},
"Atlantic/Bermuda"                   :{ "ADT-0300","AST-0400"},
"Atlantic/Canary"                    :{ "WEST+0100","WET+0000"},
"Atlantic/Faeroe"                    :{ "WEST+0100","WET+0000"},
"Atlantic/Faroe"                     :{ "WEST+0100","WET+0000"},
"Atlantic/Jan_Mayen"                 :{ "CEST+0200","CET+0100"},
"Atlantic/Madeira"                   :{ "WEST+0100","WET+0000"},
"Atlantic/Reykjavik"                 :{ "GMT+0000"},
"Atlantic/St_Helena"                 :{ "GMT+0000"},
"Australia/ACT"                      :{ "AEST+1000","AEDT+1100"},
"Australia/Adelaide"                 :{ "ACST+0930","ACDT+1030"},
"Australia/Brisbane"                 :{ "AEST+1000"},
"Australia/Broken_Hill"              :{ "ACST+0930","ACDT+1030"},
"Australia/Canberra"                 :{ "AEST+1000","AEDT+1100"},
"Australia/Currie"                   :{ "AEST+1000","AEDT+1100"},
"Australia/Darwin"                   :{ "ACST+0930"},
"Australia/Hobart"                   :{ "AEST+1000","AEDT+1100"},
"Australia/Lindeman"                 :{ "AEST+1000"},
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
"Cuba"                               :{ "CDT-0400","CST-0500"},
"EET"                                :{ "EEST+0300","EET+0200"},
"EST"                                :{ "EST-0500"},
"EST5EDT"                            :{ "EDT-0400","EST-0500"},
"Egypt"                              :{ "EET+0200"},
"Eire"                               :{ "IST+0100","GMT+0000"},
"Etc/GMT"                            :{ "GMT+0000"},
"Etc/GMT+0"                          :{ "GMT+0000"},
"Etc/GMT-0"                          :{ "GMT+0000"},
"Etc/GMT0"                           :{ "GMT+0000"},
"Etc/Greenwich"                      :{ "GMT+0000"},
"Etc/UCT"                            :{ "UTC+0000"},
"Etc/UTC"                            :{ "UTC+0000"},
"Etc/Universal"                      :{ "UTC+0000"},
"Etc/Zulu"                           :{ "UTC+0000"},
"Europe/Amsterdam"                   :{ "CEST+0200","CET+0100"},
"Europe/Andorra"                     :{ "CEST+0200","CET+0100"},
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
"Europe/Jersey"                      :{ "BST+0100","GMT+0000"},
"Europe/Kaliningrad"                 :{ "EET+0200"},
"Europe/Kiev"                        :{ "EEST+0300","EET+0200"},
"Europe/Lisbon"                      :{ "WEST+0100","WET+0000"},
"Europe/Ljubljana"                   :{ "CEST+0200","CET+0100"},
"Europe/London"                      :{ "BST+0100","GMT+0000"},
"Europe/Luxembourg"                  :{ "CEST+0200","CET+0100"},
"Europe/Madrid"                      :{ "CEST+0200","CET+0100"},
"Europe/Malta"                       :{ "CEST+0200","CET+0100"},
"Europe/Mariehamn"                   :{ "EEST+0300","EET+0200"},
"Europe/Monaco"                      :{ "CEST+0200","CET+0100"},
"Europe/Moscow"                      :{ "MSK+0300"},
"Europe/Nicosia"                     :{ "EEST+0300","EET+0200"},
"Europe/Oslo"                        :{ "CEST+0200","CET+0100"},
"Europe/Paris"                       :{ "CEST+0200","CET+0100"},
"Europe/Podgorica"                   :{ "CEST+0200","CET+0100"},
"Europe/Prague"                      :{ "CEST+0200","CET+0100"},
"Europe/Riga"                        :{ "EEST+0300","EET+0200"},
"Europe/Rome"                        :{ "CEST+0200","CET+0100"},
"Europe/San_Marino"                  :{ "CEST+0200","CET+0100"},
"Europe/Sarajevo"                    :{ "CEST+0200","CET+0100"},
"Europe/Simferopol"                  :{ "MSK+0300"},
"Europe/Skopje"                      :{ "CEST+0200","CET+0100"},
"Europe/Sofia"                       :{ "EEST+0300","EET+0200"},
"Europe/Stockholm"                   :{ "CEST+0200","CET+0100"},
"Europe/Tallinn"                     :{ "EEST+0300","EET+0200"},
"Europe/Tirane"                      :{ "CEST+0200","CET+0100"},
"Europe/Tiraspol"                    :{ "EEST+0300","EET+0200"},
"Europe/Uzhgorod"                    :{ "EEST+0300","EET+0200"},
"Europe/Vaduz"                       :{ "CEST+0200","CET+0100"},
"Europe/Vatican"                     :{ "CEST+0200","CET+0100"},
"Europe/Vienna"                      :{ "CEST+0200","CET+0100"},
"Europe/Vilnius"                     :{ "EEST+0300","EET+0200"},
"Europe/Warsaw"                      :{ "CEST+0200","CET+0100"},
"Europe/Zagreb"                      :{ "CEST+0200","CET+0100"},
"Europe/Zaporozhye"                  :{ "EEST+0300","EET+0200"},
"Europe/Zurich"                      :{ "CEST+0200","CET+0100"},
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
"Indian/Comoro"                      :{ "EAT+0300"},
"Indian/Mayotte"                     :{ "EAT+0300"},
"Israel"                             :{ "IDT+0300","IST+0200"},
"Jamaica"                            :{ "EST-0500"},
"Japan"                              :{ "JST+0900"},
"Libya"                              :{ "EET+0200"},
"MET"                                :{ "MEST+0200","MET+0100"},
"MST"                                :{ "MST-0700"},
"MST7MDT"                            :{ "MDT-0600","MST-0700"},
"Mexico/BajaNorte"                   :{ "PDT-0700","PST-0800"},
"Mexico/BajaSur"                     :{ "MDT-0600","MST-0700"},
"Mexico/General"                     :{ "CDT-0500","CST-0600"},
"NZ"                                 :{ "NZST+1200","NZDT+1300"},
"Navajo"                             :{ "MDT-0600","MST-0700"},
"PRC"                                :{ "CST+0800"},
"PST8PDT"                            :{ "PDT-0700","PST-0800"},
"Pacific/Auckland"                   :{ "NZST+1200","NZDT+1300"},
"Pacific/Guam"                       :{ "ChST+1000"},
"Pacific/Honolulu"                   :{ "HST-1000"},
"Pacific/Johnston"                   :{ "HST-1000"},
"Pacific/Midway"                     :{ "SST-1100"},
"Pacific/Pago_Pago"                  :{ "SST-1100"},
"Pacific/Saipan"                     :{ "ChST+1000"},
"Pacific/Samoa"                      :{ "SST-1100"},
"Poland"                             :{ "CEST+0200","CET+0100"},
"Portugal"                           :{ "WEST+0100","WET+0000"},
"ROC"                                :{ "CST+0800"},
"ROK"                                :{ "KST+0900"},
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


