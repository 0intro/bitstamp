// Copyright 2017 David du Colombier. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

// Transaction types
const (
	TransactionTypeDeposit            = 0
	TransactionTypeWithdrawal         = 1
	TransactionTypeMarketTrade        = 2
	TransactionTypeSubAccountTransfer = 14
)

// Order types
const (
	OrderTypeBuy  = 0
	OrderTypeSell = 1
)

// Withdrawal types
const (
	WithdrawalTypeSEPA = 0
	WithdrawalTypeBTC  = 1
	WithdrawalTypeWIRE = 2
	WithdrawalTypeXRP  = 14
	WithdrawalTypeLTC  = 15
	WithdrawalTypeETH  = 15
)
