package carbon

import (
	stdtime "time"

	"github.com/golang-module/carbon/v2"
)

type Carbon = carbon.Carbon

// timezone constants
const (
	Local = carbon.Local
	UTC   = carbon.UTC
	GMT   = carbon.GMT
	EET   = carbon.EET
	WET   = carbon.WET
	CET   = carbon.CET
	EST   = carbon.EST
	MST   = carbon.MST

	Cuba      = carbon.Cuba
	Egypt     = carbon.Egypt
	Eire      = carbon.Eire
	Greenwich = carbon.Greenwich
	Iceland   = carbon.Iceland
	Iran      = carbon.Iran
	Israel    = carbon.Israel
	Jamaica   = carbon.Jamaica
	Japan     = carbon.Japan
	Libya     = carbon.Libya
	Poland    = carbon.Poland
	Portugal  = carbon.Portugal
	PRC       = carbon.PRC
	Singapore = carbon.Singapore
	Turkey    = carbon.Turkey

	Shanghai   = carbon.Shanghai
	Chongqing  = carbon.Chongqing
	Harbin     = carbon.Harbin
	Urumqi     = carbon.Urumqi
	HongKong   = carbon.HongKong
	Macao      = carbon.Macao
	Taipei     = carbon.Taipei
	Tokyo      = carbon.Tokyo
	Saigon     = carbon.Saigon
	Seoul      = carbon.Seoul
	Bangkok    = carbon.Bangkok
	Dubai      = carbon.Dubai
	NewYork    = carbon.NewYork
	LosAngeles = carbon.LosAngeles
	Chicago    = carbon.Chicago
	Moscow     = carbon.Moscow
	London     = carbon.London
	Berlin     = carbon.Berlin
	Paris      = carbon.Paris
	Rome       = carbon.Rome
	Sydney     = carbon.Sydney
	Melbourne  = carbon.Melbourne
	Darwin     = carbon.Darwin
)

// month constants
const (
	January   = carbon.January
	February  = carbon.February
	March     = carbon.March
	April     = carbon.April
	May       = carbon.May
	June      = carbon.June
	July      = carbon.July
	August    = carbon.August
	September = carbon.September
	October   = carbon.October
	November  = carbon.November
	December  = carbon.December
)

// week constants
const (
	Monday    = carbon.Monday
	Tuesday   = carbon.Tuesday
	Wednesday = carbon.Wednesday
	Thursday  = carbon.Thursday
	Friday    = carbon.Friday
	Saturday  = carbon.Saturday
	Sunday    = carbon.Sunday
)

// number constants
const (
	YearsPerMillennium = 1000
	YearsPerCentury    = 100
	YearsPerDecade     = 10
	QuartersPerYear    = 4
	MonthsPerYear      = 12
	MonthsPerQuarter   = 3
	WeeksPerNormalYear = 52
	weeksPerLongYear   = 53
	WeeksPerMonth      = 4
	DaysPerLeapYear    = 366
	DaysPerNormalYear  = 365
	DaysPerWeek        = 7
	HoursPerWeek       = 168
	HoursPerDay        = 24
	MinutesPerDay      = 1440
	MinutesPerHour     = 60
	SecondsPerWeek     = 604800
	SecondsPerDay      = 86400
	SecondsPerHour     = 3600
	SecondsPerMinute   = 60
)

// layout constants
const (
	ANSICLayout              = stdtime.ANSIC
	UnixDateLayout           = stdtime.UnixDate
	RubyDateLayout           = stdtime.RubyDate
	RFC822Layout             = stdtime.RFC822
	RFC822ZLayout            = stdtime.RFC822Z
	RFC850Layout             = stdtime.RFC850
	RFC1123Layout            = stdtime.RFC1123
	RFC1123ZLayout           = stdtime.RFC1123Z
	RssLayout                = stdtime.RFC1123Z
	KitchenLayout            = stdtime.Kitchen
	RFC2822Layout            = stdtime.RFC1123Z
	CookieLayout             = carbon.CookieLayout
	RFC3339Layout            = carbon.RFC3339Layout
	RFC3339MilliLayout       = carbon.RFC3339MilliLayout
	RFC3339MicroLayout       = carbon.RFC3339MicroLayout
	RFC3339NanoLayout        = carbon.RFC3339NanoLayout
	ISO8601Layout            = carbon.ISO8601Layout
	ISO8601MilliLayout       = carbon.ISO8601MilliLayout
	ISO8601MicroLayout       = carbon.ISO8601MicroLayout
	ISO8601NanoLayout        = carbon.ISO8601NanoLayout
	RFC1036Layout            = carbon.RFC1036Layout
	RFC7231Layout            = carbon.RFC7231Layout
	DayDateTimeLayout        = carbon.DayDateTimeLayout
	DateTimeLayout           = carbon.DateTimeLayout
	DateTimeMilliLayout      = carbon.DateTimeMilliLayout
	DateTimeMicroLayout      = carbon.DateTimeMicroLayout
	DateTimeNanoLayout       = carbon.DateTimeNanoLayout
	ShortDateTimeLayout      = carbon.ShortDateTimeLayout
	ShortDateTimeMilliLayout = carbon.ShortDateTimeMilliLayout
	ShortDateTimeMicroLayout = carbon.ShortDateTimeMicroLayout
	ShortDateTimeNanoLayout  = carbon.ShortDateTimeNanoLayout
	DateLayout               = carbon.DateLayout
	DateMilliLayout          = carbon.DateMilliLayout
	DateMicroLayout          = carbon.DateMicroLayout
	DateNanoLayout           = carbon.DateNanoLayout
	ShortDateLayout          = carbon.ShortDateLayout
	ShortDateMilliLayout     = carbon.ShortDateMilliLayout
	ShortDateMicroLayout     = carbon.ShortDateMicroLayout
	ShortDateNanoLayout      = carbon.ShortDateNanoLayout
	TimeLayout               = carbon.TimeLayout
	TimeMilliLayout          = carbon.TimeMilliLayout
	TimeMicroLayout          = carbon.TimeMicroLayout
	TimeNanoLayout           = carbon.TimeNanoLayout
	ShortTimeLayout          = carbon.ShortTimeLayout
	ShortTimeMilliLayout     = carbon.ShortTimeMilliLayout
	ShortTimeMicroLayout     = carbon.ShortTimeMicroLayout
	ShortTimeNanoLayout      = carbon.ShortTimeNanoLayout
)
