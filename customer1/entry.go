package customer1

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
