package core

import (
	"fmt"
	"strings"
	"time"
)

const dateLayout = "2006-01"

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Format(dateLayout) + `"`), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	s := string(data)

	s = strings.Trim(s, `"`)

	if s == "" || s == "null" {
		d.Time = time.Time{}
		return nil
	}

	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return fmt.Errorf("invalid date %q, expected format YYYY-MM: %w", s, err)
	}

	d.Time = t
	return nil
}
