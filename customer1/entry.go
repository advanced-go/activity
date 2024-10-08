package customer1

import (
	"github.com/advanced-go/stdlib/core"
	"time"
)

type Entry struct {
	Customer address `json:"customer"`
	Activity []log   `json:"activity"`
}

func (e Entry) CustomerId() string {
	return e.Customer.CustomerId
}

func (e Entry) State() string {
	return e.Customer.State
}

// Address - customer address1 struct
type address struct {
	CustomerId string `json:"customer-id"`
	//CreatedTS  time.Time `json:"created-ts"`

	AddressLine1 string `json:"address-1"`
	AddressLine2 string `json:"address-2"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal-code"`
	Email        string `json:"email"`
}

// Log - access log struct
type log struct {
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
