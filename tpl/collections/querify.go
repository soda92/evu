package collections

import (
	"errors"
	"net/url"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/spf13/cast"
)

var (
	errWrongArgStructure = errors.New("expected a map, a slice with an even number of elements, or an even number of scalar values, and each key must be a string")
	errKeyIsEmptyString  = errors.New("one of the keys is an empty string")
)

// Querify returns a URL query string composed of the given key-value pairs,
// encoded and sorted by key.
func (ns *Namespace) Querify(params ...any) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	if len(params) == 1 {
		switch v := params[0].(type) {
		case map[string]any: // created with collections.Dictionary
			return mapToQueryString(v)
		case maps.Params: // site configuration or page parameters
			return mapToQueryString(v)
		case []string:
			return stringSliceToQueryString(v)
		case []any:
			s, err := interfaceSliceToStringSlice(v)
			if err != nil {
				return "", err
			}
			return stringSliceToQueryString(s)
		default:
			return "", errWrongArgStructure
		}
	}

	if len(params)%2 != 0 {
		return "", errWrongArgStructure
	}

	s, err := interfaceSliceToStringSlice(params)
	if err != nil {
		return "", err
	}
	return stringSliceToQueryString(s)
}

// mapToQueryString returns a URL query string derived from the given string
// map, encoded and sorted by key. The function returns an error if it cannot
// convert an element value to a string.
func mapToQueryString[T map[string]any | maps.Params](m T) (string, error) {
	if len(m) == 0 {
		return "", nil
	}

	qs := url.Values{}
	for k, v := range m {
		if len(k) == 0 {
			return "", errKeyIsEmptyString
		}
		vs, err := cast.ToStringE(v)
		if err != nil {
			return "", err
		}
		qs.Add(k, vs)
	}
	return qs.Encode(), nil
}

// sliceToQueryString returns a URL query string derived from the given slice
// of strings, encoded and sorted by key. The function returns an error if
// there are an odd number of elements.
func stringSliceToQueryString(s []string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if len(s)%2 != 0 {
		return "", errWrongArgStructure
	}

	qs := url.Values{}
	for i := 0; i < len(s); i += 2 {
		if len(s[i]) == 0 {
			return "", errKeyIsEmptyString
		}
		qs.Add(s[i], s[i+1])
	}
	return qs.Encode(), nil
}

// interfaceSliceToStringSlice converts a slice of interfaces to a slice of
// strings, returning an error if it cannot convert an element to a string.
func interfaceSliceToStringSlice(s []any) ([]string, error) {
	if len(s) == 0 {
		return []string{}, nil
	}

	ss := make([]string, 0, len(s))
	for _, v := range s {
		vs, err := cast.ToStringE(v)
		if err != nil {
			return []string{}, err
		}
		ss = append(ss, vs)
	}
	return ss, nil
}
