// Copyright 2017 David du Colombier. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitstamp

import "errors"

var (
	ErrNotFound     = errors.New("Not Found")
	ErrAuthRequired = errors.New("Authentication Required")
)

var CurrencyPairs = []string{
	"btcusd",
	"btceur",
	"eurusd",
	"xrpusd",
	"xrpeur",
	"xrpbtc",
	"ltcusd",
	"ltceur",
	"ltcbtc",
	"ethusd",
	"etheur",
	"ethbtc",
}
