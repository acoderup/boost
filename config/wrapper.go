package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Get joins args as config key and replies config value
func Get(args ...string) interface{} {
	return c.Get(strings.Join(args, "."))
}

// GetString joins args as config key and replies config value
func GetString(args ...string) string {
	return c.ParseDepth(c.autoParseDepth, c.GetString(strings.Join(args, ".")))
}

// GetBool joins args as config key and replies config value
func GetBool(args ...string) bool {
	return c.GetBool(strings.Join(args, "."))
}

// GetInt joins args as config key and replies config value
func GetInt(args ...string) int {
	return c.GetInt(strings.Join(args, "."))
}

// GetInt32 joins args as config key and replies config value
func GetInt32(args ...string) int32 {
	return c.GetInt32(strings.Join(args, "."))
}

// GetInt64 joins args as config key and replies config value
func GetInt64(args ...string) int64 {
	return c.GetInt64(strings.Join(args, "."))
}

// GetUint joins args as config key and replies config value
func GetUint(args ...string) uint {
	return c.GetUint(strings.Join(args, "."))
}

// GetUint32 joins args as config key and replies config value
func GetUint32(args ...string) uint32 {
	return c.GetUint32(strings.Join(args, "."))
}

// GetUint64 joins args as config key and replies config value
func GetUint64(args ...string) uint64 {
	return c.GetUint64(strings.Join(args, "."))
}

// GetFloat64 joins args as config key and replies config value
func GetFloat64(args ...string) float64 {
	return c.GetFloat64(strings.Join(args, "."))
}

// GetTime joins args as config key and replies config value
func GetTime(args ...string) time.Time {
	return c.GetTime(strings.Join(args, "."))
}

// GetDuration joins args as config key and replies config value
func GetDuration(args ...string) time.Duration {
	return c.GetDuration(strings.Join(args, "."))
}

// GetIntSlice joins args as config key and replies config value
func GetIntSlice(args ...string) []int {
	return c.GetIntSlice(strings.Join(args, "."))
}

// GetStringSlice joins args as config key and replies config value
func GetStringSlice(args ...string) []string {
	strs := c.GetStringSlice(strings.Join(args, "."))
	for i, s := range strs {
		strs[i] = c.ParseDepth(c.autoParseDepth, s)
	}
	return strs
}

// GetStringMap joins args as config key and replies config value
func GetStringMap(args ...string) map[string]interface{} {
	m := c.GetStringMap(strings.Join(args, "."))
	for k, v := range m {
		if v, ok := v.(string); ok {
			m[c.ParseDepth(c.autoParseDepth, k)] = c.ParseDepth(c.autoParseDepth, v)
		}
	}
	return m
}

// GetStringMapString joins args as config key and replies config value
func GetStringMapString(args ...string) map[string]string {
	m := c.GetStringMapString(strings.Join(args, "."))
	for k, v := range m {
		m[c.ParseDepth(c.autoParseDepth, k)] = c.ParseDepth(c.autoParseDepth, v)
	}
	return m
}

// GetStringMapStringSlice joins args as config key and replies config value
func GetStringMapStringSlice(args ...string) map[string][]string {
	m := c.GetStringMapStringSlice(strings.Join(args, "."))
	for k, v := range m {
		for i, s := range v {
			v[i] = c.ParseDepth(c.autoParseDepth, s)
		}
		m[c.ParseDepth(c.autoParseDepth, k)] = v
	}
	return m
}

// GetSizeInBytes joins args as config key and replies config value
func GetSizeInBytes(args ...string) uint {
	return c.GetSizeInBytes(strings.Join(args, "."))
}

// Unmarshal joins args as config key and replies config value
func Unmarshal(rawVal interface{}, args ...string) error {
	return c.UnmarshalKey(strings.Join(args, "."), rawVal)
}

func Sub(args ...string) *viper.Viper {
	return c.Sub(strings.Join(args, "."))
}
