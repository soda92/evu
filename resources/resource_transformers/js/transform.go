package js

import (
	"io"
	"path"
	"path/filepath"

	"github.com/gohugoio/hugo/media"
	"github.com/gohugoio/hugo/noninternal/js/esbuild"
	"github.com/gohugoio/hugo/resources"
	"github.com/gohugoio/hugo/resources/internal"
)

type buildTransformation struct {
	optsm map[string]any
	c     *Client
}

func (t *buildTransformation) Key() internal.ResourceTransformationKey {
	return internal.NewResourceTransformationKey("jsbuild", t.optsm)
}

func (t *buildTransformation) Transform(ctx *resources.ResourceTransformationCtx) error {
	ctx.OutMediaType = media.Builtin.JavascriptType

	var opts esbuild.Options

	if t.optsm != nil {
		optsExt, err := esbuild.DecodeExternalOptions(t.optsm)
		if err != nil {
			return err
		}
		opts.ExternalOptions = optsExt
	}

	if opts.TargetPath != "" {
		ctx.OutPath = opts.TargetPath
	} else {
		ctx.ReplaceOutPathExtension(".js")
	}

	src, err := io.ReadAll(ctx.From)
	if err != nil {
		return err
	}

	opts.SourceDir = filepath.FromSlash(path.Dir(ctx.SourcePath))
	opts.Contents = string(src)
	opts.MediaType = ctx.InMediaType
	opts.Stdin = true

	_, err = t.c.transform(opts, ctx)

	return err
}
