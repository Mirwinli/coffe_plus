package domain

import (
	"context"
	"io"
)

type ImageUploader interface {
	Upload(ctx context.Context, file File) (url string, publicID string, err error)
	Delete(ctx context.Context, fileID string) error
}

type File interface {
	io.Reader
	io.Closer
}
