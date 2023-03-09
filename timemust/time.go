// Package timemust is a must-style wrapper over the standard library time package.
// All functions panic in response to any errors encountered.
package timemust

import (
	"time"

	"github.com/mdwhatcott/must/must"
)

func ParseDuration(s string) time.Duration {
	return must.Value(time.ParseDuration(s))
}
func LoadLocation(name string) *time.Location {
	return must.Value(time.LoadLocation(name))
}
func LoadLocationFromTZData(name string, data []byte) *time.Location {
	return must.Value(time.LoadLocationFromTZData(name, data))
}
func Parse(layout, value string) time.Time {
	return must.Value(time.Parse(layout, value))
}
func ParseInLocation(layout, value string, loc *time.Location) time.Time {
	return must.Value(time.ParseInLocation(layout, value, loc))
}
