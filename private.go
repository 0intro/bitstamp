// Copyright 2017 David du Colombier. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitstamp

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/0intro/bitstamp/response"
)

func (c *Conn) PostBalanceAll() (*response.Balance, error) {
	return c.PostBalance("")
}

func (c *Conn) PostBalance(currencyPair string) (*response.Balance, error) {
	v := url.Values{}
	path := fmt.Sprint("/v2/balance/")
	if currencyPair != "" {
		path += fmt.Sprintf("%s/", currencyPair)
	}
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &response.Balance{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Conn) PostUserTransactionsAll(offset int, limit int, sort string) (*[]response.UserTransaction, error) {
	return c.PostUserTransactions("", offset, limit, sort)
}

func (c *Conn) PostUserTransactions(currencyPair string, offset int, limit int, sort string) (*[]response.UserTransaction, error) {
	v := url.Values{}
	if offset > 0 {
		v.Set("offset", fmt.Sprint(offset))
	}
	if limit > 0 {
		v.Set("limit", fmt.Sprint(limit))
	}
	if sort != "" {
		v.Set("sort", sort)
	}
	path := fmt.Sprint("/v2/user_transactions/")
	if currencyPair != "" {
		path += fmt.Sprintf("%s/", currencyPair)
	}
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &[]response.UserTransaction{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Conn) PostOpenOrdersAll() (*[]response.OpenOrder, error) {
	return c.PostOpenOrders("all")
}

func (c *Conn) PostOpenOrders(currencyPair string) (*[]response.OpenOrder, error) {
	v := url.Values{}
	path := fmt.Sprintf("/v2/open_orders/%s/", currencyPair)
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &[]response.OpenOrder{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Conn) PostOrderStatus(id int64) (*response.OrderStatus, error) {
	v := url.Values{}
	v.Set("id", fmt.Sprint(id))
	path := fmt.Sprint("/order_status/")
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &response.OrderStatus{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Conn) PostCancelOrder(id int64) (*response.CancelOrder, error) {
	v := url.Values{}
	v.Set("id", fmt.Sprint(id))
	path := fmt.Sprint("/v2/cancel_order/")
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &response.CancelOrder{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Conn) PostBuyLimitOrder(currencyPair string, amount float64, price float64, limitPrice float64, dailyOrder bool) (*response.BuyLimitOrder, error) {
	v := url.Values{}
	v.Set("amount", fmt.Sprint(amount))
	v.Set("price", fmt.Sprint(price))
	if limitPrice > 0 {
		v.Set("limit_price", fmt.Sprint(limitPrice))
	}
	if dailyOrder == true {
		v.Set("daily_order", fmt.Sprint(dailyOrder))
	}
	path := fmt.Sprintf("/v2/buy/%s/", currencyPair)
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &response.BuyLimitOrder{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Conn) PostBuyMarketOrder(currencyPair string, amount float64) (*response.BuyMarketOrder, error) {
	v := url.Values{}
	v.Set("amount", fmt.Sprint(amount))
	path := fmt.Sprintf("/v2/buy/market/%s/", currencyPair)
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &response.BuyMarketOrder{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Conn) PostSellLimitOrder(currencyPair string, amount float64, price float64, limitPrice float64, dailyOrder bool) (*response.SellLimitOrder, error) {
	v := url.Values{}
	v.Set("amount", fmt.Sprint(amount))
	v.Set("price", fmt.Sprint(price))
	if limitPrice > 0 {
		v.Set("limit_price", fmt.Sprint(limitPrice))
	}
	if dailyOrder == true {
		v.Set("daily_order", fmt.Sprint(dailyOrder))
	}
	path := fmt.Sprintf("/v2/sell/%s/", currencyPair)
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &response.SellLimitOrder{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Conn) PostSellMarketOrder(currencyPair string, amount float64) (*response.SellMarketOrder, error) {
	v := url.Values{}
	v.Set("amount", fmt.Sprint(amount))
	path := fmt.Sprintf("/v2/sell/market/%s/", currencyPair)
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &response.SellMarketOrder{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Conn) PostWithdrawalRequests(timedelta int) (*[]response.WithdrawalRequest, error) {
	v := url.Values{}
	if timedelta > 0 {
		v.Set("timedelta", fmt.Sprint(timedelta))
	}
	path := fmt.Sprintf("/v2/withdrawal-requests/")
	b, err := c.request("POST", path, v, true)
	if err != nil {
		return nil, err
	}
	res := &[]response.WithdrawalRequest{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
