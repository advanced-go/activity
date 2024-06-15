package assignment

import (
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

var index = make(map[string]string)

func init() {
	for _, e := range entryData {
		index[core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}.Tag()] = e.EntryId
	}
}

func lookupEntry(values url.Values) (string, bool) {
	id, ok := index[core.NewOrigin(values).Tag()]
	return id, ok
}
