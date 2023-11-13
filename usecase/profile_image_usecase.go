package usecase

import (
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
	reader, err := p.gcsService.GetObjectReader(bucket, name)
	if err != nil {
		return err
	}
	format, err := p.imagickService.GetFileFormat(reader)
	if err != nil {
		return err
	}
	log.Fatal(format)
	return nil
}

func NewProfileImageUseCase(gcsService service.GCSService, imagickService service.ImagickService) ProfileImageUseCase {
	return &profileImageUseCase{
		gcsService:     gcsService,
		imagickService: imagickService,
	}
}
