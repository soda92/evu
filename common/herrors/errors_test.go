package herrors

import (
	"errors"
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/spf13/afero"
)

func TestIsNotExist(t *testing.T) {
	c := qt.New(t)

	c.Assert(IsNotExist(afero.ErrFileNotFound), qt.Equals, true)
	c.Assert(IsNotExist(afero.ErrFileExists), qt.Equals, false)
	c.Assert(IsNotExist(afero.ErrDestinationExists), qt.Equals, false)
	c.Assert(IsNotExist(nil), qt.Equals, false)

	c.Assert(IsNotExist(fmt.Errorf("foo")), qt.Equals, false)

	// os.IsNotExist returns false for wrapped errors.
	c.Assert(IsNotExist(fmt.Errorf("foo: %w", afero.ErrFileNotFound)), qt.Equals, true)
}

func TestIsFeatureNotAvailableError(t *testing.T) {
	c := qt.New(t)

	c.Assert(IsFeatureNotAvailableError(ErrFeatureNotAvailable), qt.Equals, true)
	c.Assert(IsFeatureNotAvailableError(&FeatureNotAvailableError{}), qt.Equals, true)
	c.Assert(IsFeatureNotAvailableError(errors.New("asdf")), qt.Equals, false)
}
