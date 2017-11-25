[![Build Status](https://travis-ci.org/0intro/bitstamp.svg?branch=master)](https://travis-ci.org/0intro/bitstamp)

Bitstamp
========

Package bitstamp implements the [Bitstamp HTTP API](https://www.bitstamp.net/api/).
For full documentation, see [the godoc page](http://godoc.org/github.com/0intro/bitstamp).

Examples
--------

Ticker (unauthenticated)

```
	conn, err := bitstamp.NewConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	ticker, err := conn.GetTicker(currencyPair)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", ticker)
```

Account Balance (authenticated)

```
	conn, err := bitstamp.NewAuthConn(apiKey, apiSecret, customerID)
	if err != nil {
		return err
	}
	defer conn.Close()

	balance, err := conn.PostBalance(currencyPair)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", balance)
```
