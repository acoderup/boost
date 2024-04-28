package timex

import (
	"fmt"
	"time"
)

// SpecificTime use specific timestamp
type SpecificTime struct {
	FakeTime
	now time.Time
}

func NewSpecificTime(tm int64) *SpecificTime {
	v := &SpecificTime{}
	v.initTimeLocation()
	v.now = time.Unix(tm, 0)
	fmt.Printf("Specific times is set, time is %s.\n", v.Format(v.Now()))
	return v
}

func (v *SpecificTime) Now() time.Time {
	return v.now.In(v.location)
}

func (v *SpecificTime) FakeDuration() time.Duration {
	return 0
}
