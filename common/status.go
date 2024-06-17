package common

import (
	"fmt"
	"net/url"
)

type StatusT interface {
	GetEntryId() int
	GetStatus() string
}

func ValidStatus(values url.Values, e StatusT) bool {
	if values == nil || values.Get("entry-id") != fmt.Sprintf("%v", e.GetEntryId()) {
		return false
	}

	s := values.Get("status")
	if s != "" && e.GetStatus() != s {
		return false
	}
	return true
}
