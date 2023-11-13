package usecase

import (
	"bytes"
	"irelove.ireisu.com/domain/service"
	"log"
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
	format, err := p.imagickService.GetFileFormat(bytes.NewBuffer(blob))
	if err != nil {
		return err
	}
	//newImageBuffer, err := p.imagickService.ConvertResize(bytes.NewBuffer(blob), ProfileImageWidth, ProfileImageHeight)

	log.Println(format)
	return nil
}

func NewProfileImageUseCase(gcsService service.GCSService, imagickService service.ImagickService) ProfileImageUseCase {
	return &profileImageUseCase{
		gcsService:     gcsService,
		imagickService: imagickService,
	}
}
