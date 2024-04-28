package timex

import (
	"time"
)

// FakeTime is a simulator for `FakeTime` in C language
type FakeTime struct {
	location     *time.Location
	fakeDuration time.Duration
}

func NewFakeTime() *FakeTime {
	v := &FakeTime{}
	v.initTimeLocation()
	v.initFakeTime()
	return v
}

func (v *FakeTime) initTimeLocation() {
	v.location = time.FixedZone("SYS", int(timex.Zone))
}

func (v *FakeTime) initFakeTime() {
	v.fakeDuration = time.Duration(timex.Fake)
}

func (v *FakeTime) Time(ts int64) time.Time {
	return time.Unix(ts, 0).In(v.location)
}

func (v *FakeTime) Now() time.Time {
	return time.Now().Add(v.fakeDuration).In(v.location)
}

func (v *FakeTime) Parse(s string) (time.Time, error) {
	var format string
	if len(s) > 10 {
		format = datetimeParse
	} else {
		format = dateParse
	}
	return time.ParseInLocation(format, s, v.location)
}

func (v *FakeTime) Format(tm time.Time) string {
	return tm.Format(datetimeFormat)
}

func (v *FakeTime) Location() *time.Location {
	return v.location
}

func (v *FakeTime) FakeDuration() time.Duration {
	return v.fakeDuration
}
