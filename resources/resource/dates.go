package resource

import (
	"time"

	"github.com/gohugoio/hugo/common/htime"
)

// Dated wraps a "dated resource". These are the 4 dates that makes
// the date logic in Hugo.
type Dated interface {
	// Date returns the date of the resource.
	Date() time.Time

	// Lastmod returns the last modification date of the resource.
	Lastmod() time.Time

	// PublishDate returns the publish date of the resource.
	PublishDate() time.Time

	// ExpiryDate returns the expiration date of the resource.
	ExpiryDate() time.Time
}

// IsFuture returns whether the argument represents the future.
func IsFuture(d Dated) bool {
	if d.PublishDate().IsZero() {
		return false
	}

	return d.PublishDate().After(htime.Now())
}

// IsExpired returns whether the argument is expired.
func IsExpired(d Dated) bool {
	if d.ExpiryDate().IsZero() {
		return false
	}
	return d.ExpiryDate().Before(htime.Now())
}

// IsZeroDates returns true if all of the dates are zero.
func IsZeroDates(d Dated) bool {
	return d.Date().IsZero() && d.Lastmod().IsZero() && d.ExpiryDate().IsZero() && d.PublishDate().IsZero()
}
