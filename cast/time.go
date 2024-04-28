package cast

import (
	"bytes"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const UTC = "UTC"

func ToDuration(a any) time.Duration {
	v, _ := ToDurationE(a)
	return v
}

func ToDurationE(a any) (time.Duration, error) {
	a = indirectToStringerOrError(a)

	switch v := a.(type) {
	case int:
		return time.Duration(v), nil
	case int8:
		return time.Duration(v), nil
	case int16:
		return time.Duration(v), nil
	case int32:
		return time.Duration(v), nil
	case int64:
		return time.Duration(v), nil
	case uint:
		return time.Duration(v), nil
	case uint8:
		return time.Duration(v), nil
	case uint16:
		return time.Duration(v), nil
	case uint32:
		return time.Duration(v), nil
	case uint64:
		return time.Duration(v), nil
	case float32:
		return time.Duration(v), nil
	case float64:
		return time.Duration(v), nil
	case *big.Int:
		return time.Duration(v.Int64()), nil
	case *big.Float:
		n, _ := v.Int64()
		return time.Duration(n), nil
	case *big.Rat:
		n, _ := v.Float64()
		return time.Duration(n), nil
	case complex64:
		return time.Duration(real(v)), nil
	case complex128:
		return time.Duration(real(v)), nil
	case bool:
		if v {
			return time.Duration(1), nil
		}
		return time.Duration(0), nil
	case time.Duration:
		return v, nil
	case time.Location:
		_, offset := time.Now().In(&v).Zone()
		return time.Duration(offset), nil
	case *time.Location:
		_, offset := time.Now().In(v).Zone()
		return time.Duration(offset), nil
	case string:
		return stringToDurationE(v)
	case []byte:
		return stringToDurationE(string(v))
	case fmt.Stringer:
		return stringToDurationE(v.String())
	case error:
		return stringToDurationE(v.Error())
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("invalid duration type: %T", a)
	}
}

func ToTimeZone(a any) *time.Location {
	v, _ := ToTimeZoneE(a)
	return v
}

func ToTimeZoneE(a any) (*time.Location, error) {
	a = indirectToStringerOrError(a)

	switch v := a.(type) {
	case int:
		return time.FixedZone(UTC, v), nil
	case int8:
		return time.FixedZone(UTC, int(v)), nil
	case int16:
		return time.FixedZone(UTC, int(v)), nil
	case int32:
		return time.FixedZone(UTC, int(v)), nil
	case int64:
		return time.FixedZone(UTC, int(v)), nil
	case uint:
		return time.FixedZone(UTC, int(v)), nil
	case uint8:
		return time.FixedZone(UTC, int(v)), nil
	case uint16:
		return time.FixedZone(UTC, int(v)), nil
	case uint32:
		return time.FixedZone(UTC, int(v)), nil
	case uint64:
		return time.FixedZone(UTC, int(v)), nil
	case float32:
		return time.FixedZone(UTC, int(v)), nil
	case float64:
		return time.FixedZone(UTC, int(v)), nil
	case *big.Int:
		return time.FixedZone(UTC, int(v.Int64())), nil
	case *big.Float:
		n, _ := v.Int64()
		return time.FixedZone(UTC, int(n)), nil
	case *big.Rat:
		n, _ := v.Float64()
		return time.FixedZone(UTC, int(n)), nil
	case complex64:
		return time.FixedZone(UTC, int(real(v))), nil
	case complex128:
		return time.FixedZone(UTC, int(real(v))), nil
	case bool:
		if v {
			return time.FixedZone(UTC, 1), nil
		}
		return time.FixedZone(UTC, 0), nil
	case time.Duration:
		return durationToLocationE(v)
	case time.Location:
		return &v, nil
	case *time.Location:
		return v, nil
	case string:
		return stringToLocationE(v)
	case []byte:
		return stringToLocationE(string(v))
	case fmt.Stringer:
		return stringToLocationE(v.String())
	case error:
		return stringToLocationE(v.Error())
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("invalid time zone type: %T", a)
	}
}

var (
	durationRegExp       *regexp.Regexp
	durationRegExpGroups = []string{
		`<years>[\+|\-]?\d+[Y|y]`,
		`<months>[\+|\-]?\d+M`,
		`<days>[\+|\-]?\d+[D|d]`,
		`<hours>[\+|\-]?\d+[H|h]`,
		`<minutes>[\+|\-]?\d+m`,
		`<seconds>[\+|\-]?\d+[S|s]`,
		`<milliseconds>[\+|\-]?\d+ms`,
		`<microseconds>[\+|\-]?\d+us`,
		`<nanoseconds>[\+|\-]?\d+ns`,
	}
)

func init() {
	var buf = new(bytes.Buffer)
	for _, group := range durationRegExpGroups {
		buf.WriteString(`(?P`)
		buf.WriteString(group)
		buf.WriteString(`)?`)
	}
	durationRegExp = regexp.MustCompile(buf.String())
}

func regexpStringToDurationE(str string) (time.Duration, error) {
	epoch := time.Now().UTC()
	index := strings.Index(str, " ")
	if index != -1 {
		epoch = stringToTime(str[index+1:])
	}

	lastChar := str[len(str)-1]
	if lastChar >= '0' && lastChar <= '9' {
		str += "s"
	}

	matches := durationRegExp.FindStringSubmatch(str)

	if len(matches) == 0 {
		return 0, fmt.Errorf("parse duration `%s` failed, empty match", str)
	}

	nums := []int{}
	for index := 1; index < len(matches); index++ {
		s := matches[index]
		if len(s) == 0 {
			nums = append(nums, 0)
			continue
		}
		for s[len(s)-1] < '0' || s[len(s)-1] > '9' {
			s = s[:len(s)-1]
		}
		n, err := ToInt64E(s)
		if err != nil {
			return 0, fmt.Errorf("parse duration `%s` failed, %v", str, err)
		}
		nums = append(nums, int(n))
	}

	duration := epoch.AddDate(nums[0], nums[1], nums[2]).Add(
		time.Duration(nums[3]) * time.Hour,
	).Add(
		time.Duration(nums[4]) * time.Minute,
	).Add(
		time.Duration(nums[5]) * time.Second,
	).Add(
		time.Duration(nums[6]) * time.Millisecond,
	).Add(
		time.Duration(nums[7]) * time.Microsecond,
	).Add(
		time.Duration(nums[8]) * time.Nanosecond,
	).Sub(epoch)

	return duration, nil
}

func utcStringToDurationE(str string) (time.Duration, error) {
	var sign = 1

	if str == UTC {
		return 0, nil
	} else if strings.HasPrefix(str, "UTC+") {
		str = str[4:]
	} else if strings.HasPrefix(str, "UTC-") {
		str = str[4:]
		sign = -1
	} else {
		return 0, fmt.Errorf("invalid timezone name `%s`", str)
	}

	parts := strings.Split(str, ":")
	if len(parts) == 1 {
		h, err := ToInt64E(parts[0])
		if err != nil {
			return 0, err
		}
		return time.Duration(sign) * time.Duration(h) * time.Hour, nil
	} else if len(parts) == 2 {
		h, err := ToInt64E(parts[0])
		if err != nil {
			return 0, err
		}
		m, err := ToInt64E(parts[1])
		if err != nil {
			return 0, err
		}
		return time.Duration(sign) * (time.Duration(h)*time.Hour + time.Duration(m)*time.Minute), nil
	} else if len(parts) == 3 {
		h, err := ToInt64E(parts[0])
		if err != nil {
			return 0, err
		}
		m, err := ToInt64E(parts[1])
		if err != nil {
			return 0, err
		}
		s, err := ToInt64E(parts[2])
		if err != nil {
			return 0, err
		}
		return time.Duration(sign) * (time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second), nil
	} else {
		return 0, fmt.Errorf("invalid time zone name `%s`", str)
	}
}

func locationStringToDurationE(str string) (time.Duration, error) {
	// Check timezone is valid
	loc, err := time.LoadLocation(str)
	if err != nil {
		return 0, err
	}

	// get time zone offset
	_, offset := time.Now().In(loc).Zone()
	duration := time.Duration(offset)

	return duration, nil
}

func stringToDurationE(str string) (time.Duration, error) {
	// Check first char is + or -, or is digit
	if str[0] == '+' || str[0] == '-' || (str[0] >= '0' && str[0] <= '9') {
		duration, err := regexpStringToDurationE(str)
		if err != nil {
			return 0, err
		}
		return duration, nil
	} else if strings.HasPrefix(str, UTC) {
		duration, err := utcStringToDurationE(str)
		if err != nil {
			return 0, err
		}
		return duration, nil
	} else {
		duration, err := locationStringToDurationE(str)
		if err != nil {
			return 0, err
		}
		return duration, nil
	}
}

func durationToLocationE(duration time.Duration) (*time.Location, error) {
	var seconds = int(duration.Seconds())
	var durationStr string = UTC
	var absSeconds int
	if seconds < 0 {
		durationStr += "-"
		absSeconds = -seconds
	} else {
		durationStr += "+"
		absSeconds = seconds
	}

	h := absSeconds / 3600
	m := absSeconds % 3600 / 60
	s := absSeconds % 3600 % 60
	if h != 0 {
		durationStr += fmt.Sprintf("%02d", h)
	}
	if m != 0 {
		durationStr += fmt.Sprintf(":%02d", m)
	}
	if s != 0 {
		durationStr += fmt.Sprintf(":%02d", s)
	}

	return time.FixedZone(durationStr, seconds), nil
}

func stringToLocationE(str string) (*time.Location, error) {
	duration, err := stringToDurationE(str)
	if err != nil {
		return nil, err
	}

	return durationToLocationE(duration)
}

type TimeFormatType int

const (
	TimeFormatNoTimezone TimeFormatType = iota
	TimeFormatNamedTimezone
	TimeFormatNumericTimezone
	TimeFormatNumericAndNamedTimezone
	TimeFormatTimeOnly
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TimeFormatNoTimezone-0]
	_ = x[TimeFormatNamedTimezone-1]
	_ = x[TimeFormatNumericTimezone-2]
	_ = x[TimeFormatNumericAndNamedTimezone-3]
	_ = x[TimeFormatTimeOnly-4]
}

const _timeFormatTypeName = "timeFormatNoTimezonetimeFormatNamedTimezonetimeFormatNumericTimezonetimeFormatNumericAndNamedTimezonetimeFormatTimeOnly"

var _timeFormatTypeIndex = [...]uint8{0, 20, 43, 68, 101, 119}

func (i TimeFormatType) String() string {
	if i < 0 || i >= TimeFormatType(len(_timeFormatTypeIndex)-1) {
		return "timeFormatType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _timeFormatTypeName[_timeFormatTypeIndex[i]:_timeFormatTypeIndex[i+1]]
}

type TimeFormat struct {
	Format string
	Type   TimeFormatType
}

func (f TimeFormat) hasTimezone() bool {
	// We don't include the formats with only named timezones, see
	// https://github.com/golang/go/issues/19694#issuecomment-289103522
	return f.Type >= TimeFormatNumericTimezone && f.Type <= TimeFormatNumericAndNamedTimezone
}

var (
	defaultTimeFormats = []TimeFormat{
		{time.RFC3339, TimeFormatNumericTimezone},
		{"2006-01-02T15:04:05", TimeFormatNoTimezone}, // iso8601 without timezone
		{time.RFC1123Z, TimeFormatNumericTimezone},
		{time.RFC1123, TimeFormatNamedTimezone},
		{time.RFC822Z, TimeFormatNumericTimezone},
		{time.RFC822, TimeFormatNamedTimezone},
		{time.RFC850, TimeFormatNamedTimezone},
		{"2006-01-02 15:04:05.999999999 -0700 MST", TimeFormatNumericAndNamedTimezone}, // Time.String()
		{"2006-01-02T15:04:05-0700", TimeFormatNumericTimezone},                        // RFC3339 without timezone hh:mm colon
		{"2006-01-02 15:04:05Z0700", TimeFormatNumericTimezone},                        // RFC3339 without T or timezone hh:mm colon
		{"2006-01-02 15:04:05", TimeFormatNoTimezone},
		{time.ANSIC, TimeFormatNoTimezone},
		{time.UnixDate, TimeFormatNamedTimezone},
		{time.RubyDate, TimeFormatNumericTimezone},
		{"2006-01-02 15:04:05Z07:00", TimeFormatNumericTimezone},
		{"2006-01-02", TimeFormatNoTimezone},
		{"02 Jan 2006", TimeFormatNoTimezone},
		{"2006-01-02 15:04:05 -07:00", TimeFormatNumericTimezone},
		{"2006-01-02 15:04:05 -0700", TimeFormatNumericTimezone},
		{time.Kitchen, TimeFormatTimeOnly},
		{time.Stamp, TimeFormatTimeOnly},
		{time.StampMilli, TimeFormatTimeOnly},
		{time.StampMicro, TimeFormatTimeOnly},
		{time.StampNano, TimeFormatTimeOnly},
	}
)

func timeStringToTime(s string) (d time.Time, err error) {
	var (
		location    = time.Local
		timeFormats = defaultTimeFormats
	)

	if n, err := dec.ToInt(s); err == nil {
		return time.Unix(n, 0), nil
	}

	for _, timeFormat := range timeFormats {
		if d, err = time.Parse(timeFormat.Format, s); err == nil {
			// Some time formats have a zone name, but no offset, so it gets
			// put in that zone name (not the default one passed in to us), but
			// without that zone's offset. So set the location manually.
			if timeFormat.Type <= TimeFormatNamedTimezone {
				if location == nil {
					location = time.Local
				}
				year, month, day := d.Date()
				hour, min, sec := d.Clock()
				d = time.Date(year, month, day, hour, min, sec, d.Nanosecond(), location)
			}

			return
		}
	}
	return d, fmt.Errorf("unable to parse date: %s", s)
}

func stringToTime(s string) time.Time {
	d, e := timeStringToTime(s)
	if e != nil {
		panic(e)
	}
	return d
}
