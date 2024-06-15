package assignment

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"strconv"
)

func FilterT[T any](values url.Values, entries []T, valid func(url.Values, T) bool) ([]T, *core.Status) {
	if values == nil {
		return entries, core.StatusBadRequest()
	}
	if len(entries) == 0 {
		return entries, core.StatusNotFound()
	}
	var result []T
	id := ""
	ok := false
	if values.Get(core.RegionKey) != "*" {
		id, ok = lookupEntry(values)
		if !ok {
			return nil, core.StatusNotFound()
		}
	}
	values.Add("entry-id", id)
	for _, e := range entries {
		if valid(values, e) {
			result = append(result, e)
		}
	}
	if len(result) == 0 {
		return result, core.StatusNotFound()
	}
	result = Order(values, result)
	return Top(values, result), core.StatusOK()
}

func Order[T any](values url.Values, entries []T) []T {
	if entries == nil || values == nil {
		return entries
	}
	s := values.Get("order")
	if s != "desc" {
		return entries
	}
	var result []T

	for index := len(entries) - 1; index >= 0; index-- {
		result = append(result, entries[index])
	}
	return result
}

func Top[T any](values url.Values, entries []T) []T {
	if entries == nil || values == nil {
		return entries
	}
	s := values.Get("top")
	if s == "" {
		return entries
	}
	cnt, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("top value is not valid: %v", s)
	}
	var result []T
	for i, e := range entries {
		if i < cnt {
			result = append(result, e)
		} else {
			break
		}
	}
	return result
}
