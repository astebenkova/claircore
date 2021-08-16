package java

import (
	"archive/tar"
	"context"
	"path"
	"strings"

	"github.com/quay/zlog"
)

// IsArchive reports whether the file described by the passed tar.Header is some
// form of Java archive.
//
// Assumes slash-separated paths.
func isArchive(ctx context.Context, h *tar.Header) bool {
	base := path.Base(h.Name)
Outer:
	switch {
	case h.Typeflag != tar.TypeReg:
	case strings.HasPrefix(base, ".wh."):
	default:
		switch ext := path.Ext(base); ext {
		case ".jar":
			zlog.Debug(ctx).Str("file", h.Name).Msg("found jar")
		case ".war":
			zlog.Debug(ctx).Str("file", h.Name).Msg("found war")
		case ".ear":
			zlog.Debug(ctx).Str("file", h.Name).Msg("found ear")
		default:
			break Outer
		}
		return true
	}
	return false
}
