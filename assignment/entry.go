package assignment

// Entry - host
type Entry struct {
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	SubZone string `json:"sub-zone"`
	Host    string `json:"host"`

	Assignment string
	// ???
}

