package utils

import (
	"bytes"
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImageToCloudinary(image []byte) (string, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), bytes.NewReader(image), uploader.UploadParams{
		Folder: "go-sanber64",
	})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}