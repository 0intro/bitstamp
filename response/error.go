// Copyright 2017 David du Colombier. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

type Error struct {
	Code   string `json:"code"`
	Reason string `json:"reason"`
	Status string `json:"status"`
}
