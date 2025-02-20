package siteidentities

import (
	"github.com/gohugoio/hugo/identity"
)

const (
	// Identifies site.Data.
	// The change detection in /data is currently very coarse grained.
	Data = identity.StringIdentity("site.Data")
)

// FromString returns the identity from the given string,
// or identity.Anonymous if not found.
func FromString(name string) (identity.Identity, bool) {
	switch name {
	case "Data":
		return Data, true
	}
	return identity.Anonymous, false
}
