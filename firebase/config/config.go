package config

import (
	"time"
)

const FireBaseKeyPath = "../key/firebase.json"

var (
	AndroidTTL = time.Duration(1) * time.Hour
	Badge      = 1
	Icon       = "stock_ticker_update"
	Color      = "#f45342"
)
