package datetime


// TimeZoneConversionType - This type is configured as a series of
// constant integer values describing the two valid types of time
// zone conversion algorithms used to create new instances of type,
// 'TimeZoneDefDto'.
//
// Functionally, 'TimeZoneConversionType' serves as enumeration of
// valid time zone conversion algorithms.
//
// For easy access to these enumeration values, use the global variable
// 'TzConvertType'.
//
//                      Time Zone
//                      Conversion
//   Method               Type
//    Name              Constant        Description
//  ______________________________________________________________________
//
//  Relative()             0            Converts a given time value to
//                                      its equivalent value in a different
//                                      time zone. If the given time value is
//                                      is in time zone 'X' then than equivalent
//                                      time will be computed in time zone 'Y',
//                                      but the hours, minutes and seconds value
//                                      will be different. Example: Using the
//                                      "Relative" conversion method, 10:00AM in
//                                      USA Central Standard Time would be converted
//                                      to equivalent Eastern Standard Time value of
//                                      11:00AM.
//
//  Absolute()             1            Takes the given hours minutes and seconds
//                                      in a given time zone and leaves the hours,
//                                      minutes and seconds value unchanged while
//                                      applying a different time zone. Example:
//                                      Using the "Absolute" conversion method,
//                                      10:00AM in USA Central Standard Time would
//                                      be converted to 10:00AM Eastern Standard
//                                      Time.
//
// 'TimeZoneConversionType' has been adapted to function as an enumeration
// describing the type of time zone assigned to a date time. Since Golang
// does not directly support enumerations, the 'TimeZoneConversionType'
// type has been configured to function in a manner similar to classic
// enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
type TimeZoneConversionType int


func (tzConvertType TimeZoneConversionType) Relative() TimeZoneConversionType {
	return TimeZoneConversionType(0)
}

func (tzConvertType TimeZoneConversionType) Absolute() TimeZoneConversionType {
	return TimeZoneConversionType(1)
}