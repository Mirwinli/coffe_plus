package core_infrastructure_cloudinary

import (
	"context"
	"fmt"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryUploader struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryUploader(config Config) (*CloudinaryUploader, error) {
	cld, err := cloudinary.NewFromParams(config.CloudName, config.ApiKey, config.ApiSecret)
	if err != nil {
		return &CloudinaryUploader{}, fmt.Errorf(
			"new from params cloudinary: %w",
			err,
		)
	}

	return &CloudinaryUploader{cld: cld}, nil
}

func (c *CloudinaryUploader) Upload(ctx context.Context, file domain.File) (string, string, error) {
	result, err := c.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "products",
	})
	if err != nil {
		return "", "", fmt.Errorf(
			"upload cloudinary: %w",
			err,
		)
	}

	return result.SecureURL, result.PublicID, nil
}

func (c *CloudinaryUploader) Delete(ctx context.Context, fileID string) error {
	_, err := c.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: fileID,
	})
	if err != nil {
		return fmt.Errorf(
			"delete file cloudinary: %w",
			err,
		)
	}

	return nil
}
