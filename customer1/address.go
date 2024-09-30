package customer1

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
