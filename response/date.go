// Copyright 2017 David du Colombier. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import (
	"strconv"
	"strings"
	"time"
)

type date time.Time
type dateTime time.Time

const (
	dateLayout     = "2006-01-02 15:04:05"
	dateTimeLayout = "2006-01-02 15:04:05.9999999"
)

func (d *date) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(strings.Trim(string(b), "\""), 10, 64)
	if err != nil {
		return err
	}
	*d = date(time.Unix(i, 0))
	return nil
}

func (d date) String() string {
	return time.Time(d).Format(dateTimeLayout)
}

func (d *dateTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"`+dateTimeLayout+`"`, string(b))
	if err != nil {
		return err
	}
	*d = dateTime(t)
	return nil
}

func (d dateTime) String() string {
	return time.Time(d).Format(dateTimeLayout)
}
