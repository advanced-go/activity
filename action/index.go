package action

import (
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

var index = make(map[string]Entry)

func init() {
	for _, e := range entryData {
		index[core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}.Tag()] = e
	}
}

func lookupEntry(t any) (Entry, bool) {
	var e Entry
	ok := false
	if values, ok1 := t.(url.Values); ok1 {
		e, ok = index[core.NewOrigin(values).Tag()]
	}
	if origin, ok1 := t.(core.Origin); ok1 {
		e, ok = index[origin.Tag()]
	}
	return e, ok
}
