# Date Time Operations in Go

The source code for the Date Time Operations library is located in the 
source code repository:

https://github.com/MikeAustin71/datetimeopsgo.git


### Examples of Date Time operations in 'golang', the go programming language.

The Date Time Utilities currently consist of separate libraries maintained
in the sub-directory, 'datetime'. This directory also contains tests used to 
validate these libraries. 

##### Date Time Operations Component Libraries

 1. DateTimeUtilitiy - The DateTimeUtility structure provides generalized methods for 
     managing the formatting and display of time.Time date time values.
     Location: MikeAustin71\datetimeopsgo\datetime\datetimeutility.go 

 2. Date Time Constants - This file contains various time constants used
     in other libraries. 
     Location: MikeAustin71\datetimeopsgo\datetime\datetimeconstants.go
     
 3. DateTzDto - A type used to identify a specific point in time by date time,
     time zone and time element (year, month, day, hours etc.).
     Location: MikeAustin71\datetimeopsgo\datetime\datetzdto.go 

 4. TimeDto - is a collection of time element values. Time
     element values are represented by Years, Months, Weeks,
     WeekDays, DateDays, Hours, Minutes, Seconds, Milliseconds,
     Microseconds and Nanoseconds. Note that TimeElementDto does
     NOT contain Time Zone information.
     Location: MikeAustin71\datetimeopsgo\datetime\timedto.go

 5. TimeDurationDto - A type used to manage time duration information for 
     a specific set of starting and ending date times. 
     Location: MikeAustin71\datetimeopsgo\datetime\timedurationdto.go 

 6. DurationTriad - A group of TimeDurationDto types used to automatically 
     calculate time durations for a user designated BaseTime, LocalTime
     and UTCTime.
     Location: MikeAustin71\datetimeopsgo\datetime\durationtriad.go 

 7. FormatDateTimeUtility - This utility is designed to receive user generated
     date time strings and convert them into numeric date time values. 
     Location: MikeAustin71\datetimeopsgo\datetime\formatdatetimeutility.go

 8. TimeZoneDefDto - This structure is designed to store detailed information
     on Time Zones. It used primarily as a 'helper' or subsidiary structure
     for Type, 'TimeZoneDto'.
     Location: MikeAustin71\datetimeopsgo\datetime\timezonedef.go
     
 9.  TimeZoneDto - This structure is used to convert, store and transport time
      zone information. The user will employ this Type to convert time.Time,
      date time values, between differing time zones.
      Location:  MikeAustin71\datetimeopsgo\datetime\timezonedto.go