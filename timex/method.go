package timex

import (
	"time"
)

// TODO: upgrade all api

func Time(ts int64) time.Time {
	return timex.Time(ts)
}

func TimeDelta(ts int64) time.Time {
	return timex.Time(ts - timex.Delta)
}

func Now() time.Time {
	return timex.Now()
}

func NowDelta() time.Time {
	return Time(timex.Now().Unix() - timex.Delta)
}

func ParseTime(s string) time.Time {
	tm, err := timex.Parse(s)
	if err != nil {
		panic(err)
	}
	return tm
}

func FormatTime(tm time.Time) string {
	return timex.Format(tm)
}

func TimeLocation() *time.Location {
	return timex.Location()
}

func FakeDuration() time.Duration {
	return timex.FakeDuration()
}

// IsSameDay returns true if timestamp ts1 & ts2 in same day
func IsSameDay(ts1 int64, ts2 int64) bool {
	dt1 := Time(ts1)
	dt2 := Time(ts2)

	year1, month1, day1 := dt1.Date()
	year2, month2, day2 := dt2.Date()

	return year1 == year2 && month1 == month2 && day1 == day2
}

// IsSameDay returns true if timestamp ts1 & ts2 in same day
func IsYesterday(time int64) bool {
	return IsSameDay(time, Now().Unix()-24*3600)
}

func IsSameDayDelta(ts1 int64, ts2 int64) bool {
	dt1 := TimeDelta(ts1)
	dt2 := TimeDelta(ts2)

	year1, month1, day1 := dt1.Date()
	year2, month2, day2 := dt2.Date()

	return year1 == year2 && month1 == month2 && day1 == day2
}

// IsToday returns true if timestamp ts in today
func IsToday(ts int64) bool {
	dt := Time(ts)
	now := Now()

	year, month, day := dt.Date()
	yearNow, monthNow, dayNow := now.Date()
	return year == yearNow && month == monthNow && day == dayNow
}

func IsTodayDelta(ts int64) bool {
	dt := TimeDelta(ts)
	now := NowDelta()

	year, month, day := dt.Date()
	yearNow, monthNow, dayNow := now.Date()
	return year == yearNow && month == monthNow && day == dayNow
}

// IsSameDay returns true if timestamp ts1 & ts2 in same day
func IsSameWeek(ts1 int64, ts2 int64) bool {
	dt1 := Time(ts1)
	dt2 := Time(ts2)

	year1, week1 := dt1.ISOWeek()
	year2, week2 := dt2.ISOWeek()

	return year1 == year2 && week1 == week2
}

func IsSameWeekDelta(ts1 int64, ts2 int64) bool {
	dt1 := TimeDelta(ts1)
	dt2 := TimeDelta(ts2)

	year1, week1 := dt1.ISOWeek()
	year2, week2 := dt2.ISOWeek()

	return year1 == year2 && week1 == week2
}

// IsThisWeek returns true if timestamp t in this day
func IsThisWeek(ts int64) bool {
	dt := Time(ts)
	now := Now()

	year, week := dt.ISOWeek()
	yearNow, weekNow := now.ISOWeek()

	return year == yearNow && week == weekNow
}

func IsThisWeekDelta(ts int64) bool {
	dt := TimeDelta(ts)
	now := NowDelta()

	year, week := dt.ISOWeek()
	yearNow, weekNow := now.ISOWeek()

	return year == yearNow && week == weekNow
}

// SecondsOfThisWeek gets seconds passed in this week
func SecondsOfThisWeek() int64 {
	now := Now()
	weekDay := now.Weekday()
	day := int64(0)
	if weekDay == time.Sunday {
		day = 6
	} else {
		day = int64(weekDay) - 1
	}
	return day*86400 + SecondsOfToday()
}

func SecondsOfThisWeekDelta() int64 {
	now := NowDelta()
	weekDay := now.Weekday()
	day := int64(0)
	if weekDay == time.Sunday {
		day = 6
	} else {
		day = int64(weekDay) - 1
	}
	return day*86400 + SecondsOfToday()
}

// SecondsOfToday get seconds passed of today
func SecondsOfToday() int64 {
	now := Now()
	return 60*60*int64(now.Hour()) + 60*int64(now.Minute()) + int64(now.Second())
}

func SecondsOfTodayDelta() int64 {
	now := NowDelta()
	return 60*60*int64(now.Hour()) + 60*int64(now.Minute()) + int64(now.Second())
}

// SecondsOfDay gets the beginning time of that day
func SecondsOfDay(ts int64) int64 {
	tm := Time(ts)
	return 60*60*int64(tm.Hour()) + 60*int64(tm.Minute()) + int64(tm.Second())
}

func SecondsOfDayDelta(ts int64) int64 {
	tm := TimeDelta(ts)
	return 60*60*int64(tm.Hour()) + 60*int64(tm.Minute()) + int64(tm.Second())
}

// IntervalDays gets interval day from the beginning of that day to now
func IntervalDays(ts int64) int64 {
	if ts == 0 {
		return 0
	}
	now := Now()
	tm := Time(ts)
	secondsOfDay := SecondsOfDay(ts)

	return (now.Unix() - tm.Unix() + secondsOfDay) / 86400
}

func IntervalDaysDelta(ts int64) int64 {
	now := NowDelta()
	tm := TimeDelta(ts)
	secondsOfDay := SecondsOfDayDelta(ts)

	return (now.Unix() - tm.Unix() + secondsOfDay) / 86400
}

func BeginningOfDay(ts int64) int64 {
	tm := Time(ts)
	secondsOfDay := SecondsOfDay(ts)
	return tm.Unix() - secondsOfDay
}

func BeginningOfDayDelta(ts int64) int64 {
	return ts - SecondsOfDayDelta(ts)
}

func EndOfDay(ts int64) int64 {
	return BeginningOfDay(ts) + 3600*24
}

func EndOfDayDelta(ts int64) int64 {
	return BeginningOfDayDelta(ts) + 3600*24
}

// 获取时间戳所在周的开始和结束时间, 周以指定周号开始(如周一开始的一周)
// param t-时间戳 w-周号(0-6)
func CycleWeek(t int64, w int64) (int64, int64) {
	// 获取对应的周号
	weekDay := int64(Time(t).Weekday())
	curDay := weekDay - w
	if curDay < 0 {
		curDay += 7
	}
	// 获取当天已过去的时间
	dayBegin := BeginningOfDay(t)
	startTime := dayBegin - curDay*86400
	endTime := startTime + 7*86400
	return startTime, endTime
}

// 偏移量
func CycleWeekDelta(t int64, w int64) (int64, int64) {
	// 获取对应的周号
	weekDay := int64(TimeDelta(t).Weekday())
	curDay := weekDay - w
	if curDay < 0 {
		curDay += 7
	}
	// 获取当天已过去的时间
	dayBegin := BeginningOfDayDelta(t)
	startTime := dayBegin - curDay*86400
	endTime := startTime + 7*86400
	return startTime, endTime
}
