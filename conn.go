// Copyright 2017 David du Colombier. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bitstamp

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/0intro/bitstamp/response"
)

const (
	addr    = "https://www.bitstamp.net/api"
	verbose = false
)

type Conn struct {
	apiKey     string
	apiSecret  string
	customerID string
}

// NewConn returns a new Conn.
func NewConn() (*Conn, error) {
	return &Conn{}, nil
}

// NewConn returns a new Conn with authentication parameters.
// Customer ID: https://www.bitstamp.net/account/balance/
// API key & secret: https://www.bitstamp.net/account/security/api/
func NewAuthConn(apiKey, apiSecret, customerID string) (*Conn, error) {
	conn := &Conn{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		customerID: customerID,
	}
	return conn, nil
}

// Close closes the connection.
func (c *Conn) Close() error {
	return nil
}

// SetAuth sets authentication parameters.
func (c *Conn) SetAuth(apiKey, apiSecret, customerID string) {
	c.apiKey = apiKey
	c.apiSecret = apiSecret
	c.customerID = customerID
}

func (c *Conn) request(meth string, path string, values url.Values, auth bool) ([]byte, error) {
	if auth {
		if c.customerID == "" || c.apiKey == "" || c.apiSecret == "" {
			return nil, ErrAuthRequired
		}
		setAuthValues(values, c.customerID, c.apiKey, c.apiSecret)
	}
	body := strings.NewReader(values.Encode())
	client := &http.Client{}
	req, err := http.NewRequest(meth, addr+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if verbose {
		b, err := httputil.DumpRequest(req, true)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(b))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if verbose {
		fmt.Println(string(b))
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		res := &response.Error{}
		err = json.Unmarshal(b, res)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("HTTP Status: %s, Error: %s, Reason: %s, Status: %s", resp.Status, res.Code, res.Reason, res.Status)
	}
	return b, nil
}

func setAuthValues(values url.Values, customerID, apiKey, apiSecret string) {
	nonce := genNonce()
	signature := genSignature(nonce, customerID, apiKey, apiSecret)
	values.Set("key", apiKey)
	values.Set("nonce", nonce)
	values.Set("signature", signature)
}

func genNonce() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func genSignature(nonce, customerID, apiKey, apiSecret string) string {
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(nonce + customerID + apiKey))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}
