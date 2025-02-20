package resources_test

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/gohugoio/hugo/common/hugio"
	"github.com/gohugoio/hugo/identity"
	"github.com/gohugoio/hugo/resources"
)

func TestNewResource(t *testing.T) {
	c := qt.New(t)

	spec := newTestResourceSpec(specDescriptor{c: c})

	open := hugio.NewOpenReadSeekCloser(hugio.NewReadSeekerNoOpCloserFromString("content"))

	rd := resources.ResourceSourceDescriptor{
		OpenReadSeekCloser:   open,
		TargetPath:           "a/b.txt",
		BasePathRelPermalink: "c/d",
		BasePathTargetPath:   "e/f",
		GroupIdentity:        identity.Anonymous,
	}

	r, err := spec.NewResource(rd)
	c.Assert(err, qt.IsNil)
	c.Assert(r, qt.Not(qt.IsNil))
	c.Assert(r.RelPermalink(), qt.Equals, "/c/d/a/b.txt")

	info := resources.GetTestInfoForResource(r)
	c.Assert(info.Paths.TargetLink(), qt.Equals, "/c/d/a/b.txt")
	c.Assert(info.Paths.TargetPath(), qt.Equals, "/e/f/a/b.txt")
}
