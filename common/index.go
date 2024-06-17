package common

import (
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"sync"
)

type OriginTag interface {
	//Tag() string
	Origin() core.Origin
}

type OriginIndex[T any] struct {
	m  map[string]T
	mu sync.Mutex
}

func NewOriginIndex[T any](items []T) *OriginIndex[T] {
	i := new(OriginIndex[T])
	i.m = make(map[string]T)
	for _, e := range items {
		if ot, ok := any(e).(OriginTag); ok {
			i.m[ot.Origin().Tag()] = e
		}
	}
	return i

}

func (i *OriginIndex[T]) AddEntry(e T) *core.Status {
	i.mu.Lock()
	defer i.mu.Unlock()
	if ot, ok := any(e).(OriginTag); ok {
		if _, ok1 := i.m[ot.Origin().Tag()]; ok1 {
			return core.StatusBadRequest()
		} else {
			i.m[ot.Origin().Tag()] = e
			return core.StatusOK()
		}
	}
	return core.StatusBadRequest()
}

func (i *OriginIndex[T]) LookupEntry(t any) (T, bool) {
	var e T
	ok := false

	i.mu.Lock()
	defer i.mu.Unlock()
	if values, ok1 := t.(url.Values); ok1 {
		e, ok = i.m[core.NewOrigin(values).Tag()]
	}
	if origin, ok1 := t.(core.Origin); ok1 {
		e, ok = i.m[origin.Tag()]
	}
	return e, ok
}
