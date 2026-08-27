package core

import (
	"time"
)

const dateLayout = "2006-01"

type Date string

func (d *Date) ToTime() time.Time {
	newTime, err := time.Parse(dateLayout, string(*d))
	if err != nil {
		return time.Time{}
	}

	return newTime
}
