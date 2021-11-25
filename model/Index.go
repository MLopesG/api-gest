package model

import (
	"errors"
	"time"
)

type DateFormattedTime time.Time

func (s DateFormattedTime) MarshalJSON() ([]byte, error) {
	t := time.Time(s)
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	return []byte(t.Format(`"Jan 02, 2006"`)), nil
}
