package js

import (
	"context"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/resources/resource"
)

// BatcherClient is used to do JS batch operations.
type BatcherClient interface {
	New(id string) (Batcher, error)
	Store() *maps.Cache[string, Batcher]
}

// BatchPackage holds a group of JavaScript resources.
type BatchPackage interface {
	Groups() map[string]resource.Resources
}

// Batcher is used to build JavaScript packages.
type Batcher interface {
	Build(context.Context) (BatchPackage, error)
	Config(ctx context.Context) OptionsSetter
	Group(ctx context.Context, id string) BatcherGroup
}

// BatcherGroup is a group of scripts and instances.
type BatcherGroup interface {
	Instance(sid, iid string) OptionsSetter
	Runner(id string) OptionsSetter
	Script(id string) OptionsSetter
}

// OptionsSetter is used to set options for a batch, script or instance.
type OptionsSetter interface {
	SetOptions(map[string]any) string
}
