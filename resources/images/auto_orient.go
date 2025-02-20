package images

import (
	"image"
	"image/draw"

	"github.com/disintegration/gift"
	"github.com/gohugoio/hugo/resources/images/exif"
	"github.com/spf13/cast"
)

var _ gift.Filter = (*autoOrientFilter)(nil)

var transformationFilters = map[int]gift.Filter{
	2: gift.FlipHorizontal(),
	3: gift.Rotate180(),
	4: gift.FlipVertical(),
	5: gift.Transpose(),
	6: gift.Rotate270(),
	7: gift.Transverse(),
	8: gift.Rotate90(),
}

type autoOrientFilter struct{}

type ImageFilterFromOrientationProvider interface {
	AutoOrient(exifInfo *exif.ExifInfo) gift.Filter
}

func (f autoOrientFilter) Draw(dst draw.Image, src image.Image, options *gift.Options) {
	panic("not supported")
}

func (f autoOrientFilter) Bounds(srcBounds image.Rectangle) image.Rectangle {
	panic("not supported")
}

func (f autoOrientFilter) AutoOrient(exifInfo *exif.ExifInfo) gift.Filter {
	if exifInfo != nil {
		if v, ok := exifInfo.Tags["Orientation"]; ok {
			orientation := cast.ToInt(v)
			if filter, ok := transformationFilters[orientation]; ok {
				return filter
			}
		}
	}

	return nil
}
