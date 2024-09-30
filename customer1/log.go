package customer1

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

// Log - access log struct
type Log struct {
	//CustomerId string    `json:"customer-id"`
	Origin core.Origin `json:"origin"`

	StartTime time.Time `json:"start-time"`
	Duration  int64     `json:"duration"`
	Traffic   string    `json:"traffic"`
	//CreatedTS  time.Time `json:"created-ts"`

	//RequestId string `json:"request-id"`
	//RelatesTo string `json:"relates-to"`
	//Location  string `json:"location"`
	//Protocol  string `json:"proto"`
	Method string `json:"method"`
	//From      string `json:"from"`
	//To        string `json:"to"`
	Uri string `json:"uri"`
	//Path      string `json:"path"`
	//Query     string `json:"query"`

	StatusCode int32 `json:"status-code"`
	//Encoding   string `json:"encoding"`
	//Bytes      int64  `json:"bytes"`

	Timeout        int32   `json:"timeout"`
	RateLimit      float64 `json:"rate-limit"`
	RateBurst      int32   `json:"rate-burst"`
	ControllerCode string  `json:"cc"`
}
