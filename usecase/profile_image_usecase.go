package usecase

import (
	"bytes"
	"fmt"
	"irelove.ireisu.com/domain/service"
	"os"
)

const (
	ProfileImageWidth  = 400
	ProfileImageHeight = 400
	ProfileImageFormat = "webp"
)

type ProfileImageUseCase interface {
	ProfileImageProcess(bucket string, name string) (err error)
}

type profileImageUseCase struct {
	gcsService     service.GCSService
	imagickService service.ImagickService
}

func (p *profileImageUseCase) ProfileImageProcess(bucket string, name string) (err error) {
	blob, err := p.gcsService.GetObjectBlob(bucket, name)
	if err != nil {
		return err
	}
	newImageBuffer, err := p.imagickService.ConvertResize(bytes.NewBuffer(blob), ProfileImageWidth, ProfileImageHeight)
	newImageBuffer, err = p.imagickService.ConvertFormat(&newImageBuffer, ProfileImageFormat)
	if err != nil {
		return err
	}
	err = p.gcsService.CreateObject(newImageBuffer.Bytes(), os.Getenv("DEST_BUCKET"), name+"."+ProfileImageFormat, fmt.Sprintf("image/%s", ProfileImageFormat))
	if err != nil {
		return err
	}
	err = p.gcsService.DeleteObject(bucket, name)
	if err != nil {
		return err
	}
	return err
}

func NewProfileImageUseCase(gcsService service.GCSService, imagickService service.ImagickService) ProfileImageUseCase {
	return &profileImageUseCase{
		gcsService:     gcsService,
		imagickService: imagickService,
	}
}
