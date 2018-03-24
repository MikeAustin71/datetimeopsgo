# Date Time Operations in Go

The source code for the Date Time Operations library is located in the 
source code repository:

https://github.com/MikeAustin71/datetimeopsgo.git


### Examples of Date Time operations in 'golang', the go programming language.

The Date Time Utilities currently consist of separate libraries maintained
in the sub-directory, 'datetime'. This directory also contains tests used to 
validate these libraries. 


1. DateTzDto - A type used to identify a specific point in time by date time,
    time zone and time element (year, month, day, hours etc.).
    MikeAustin71\datetimeopsgo\datetime\datetzdto.go 

2. TimeDurationDto - A type used to manage time duration information for 
    a specific set of starting and ending date times. 
    Location: MikeAustin71\datetimeopsgo\datetime\timedurationdto.go 

3. DurationTriad - A group of TimeDurationDto types used to automatically 
    calculate time durations for a user designated BaseTime, LocalTime
    and UTCTime.
    Location: MikeAustin71\datetimeopsgo\datetime\durationtriad.go 


4. ../DateTimeFormatsUtility - This utility is designed to receive date time strings
 and convert them into time values. There are a number of dependencies associated
 with source file, 'formatdatetimeutility.go'. See source file comments.
  

