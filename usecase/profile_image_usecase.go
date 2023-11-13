package usecase

import (
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
	err = p.imagickService.ReadBlob(blob)
	if err != nil {
		return err
	}
	err = p.imagickService.Resize(ProfileImageWidth, ProfileImageHeight)
	if err != nil {
		return err
	}
	err = p.imagickService.ConvertFormat(ProfileImageFormat)
	if err != nil {
		return err
	}
	newBlob := p.imagickService.GetBlob()

	err = p.gcsService.CreateObject(newBlob, os.Getenv("DEST_BUCKET"), name+"."+ProfileImageFormat, "image/"+ProfileImageFormat)
	if err != nil {
		return err
	}
	err = p.gcsService.DeleteObject(bucket, name)
	return err
}

func NewProfileImageUseCase(gcsService service.GCSService, imagickService service.ImagickService) ProfileImageUseCase {
	return &profileImageUseCase{
		gcsService:     gcsService,
		imagickService: imagickService,
	}
}
