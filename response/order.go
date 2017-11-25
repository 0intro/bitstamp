// Copyright 2017 David du Colombier. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import (
	"encoding/json"
	"strconv"
)

type Order struct {
	Price  float64
	Amount float64
}

func (o *Order) UnmarshalJSON(data []byte) error {
	var v []interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.Price, err = strconv.ParseFloat(v[0].(string), 64)
	if err != nil {
		return err
	}
	o.Amount, err = strconv.ParseFloat(v[1].(string), 64)
	if err != nil {
		return err
	}
	return nil
}
