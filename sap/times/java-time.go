package times

import (
	"encoding/json"
	"strconv"
	"time"
)

var nilTime, _ = time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")

type JavaTime time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JavaTime) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	tm := time.Unix(0, millis*int64(time.Millisecond))
	*j = JavaTime(tm)
	return nil
}

func (j JavaTime) MarshalJSON() ([]byte, error) {
	t := time.Time(j)
	if nilTime == t {
		return json.Marshal(nil)
	}
	return json.Marshal(t)
}

// Maybe a Format function for printing your date
func (j JavaTime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
