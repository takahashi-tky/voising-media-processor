package usecase

import (
	"bytes"
	"fmt"
	"irelove.ireisu.com/api/proto/gen/media"
	"irelove.ireisu.com/domain/service"
	"os"
	"strings"
)

const (
	ProfileImageWidth  = 400
	ProfileImageHeight = 400
	ProfileImageFormat = "webp"
)

type ProfileImageUseCase interface {
	ProfileImageProcess(bucket string, name string, userImageId uint32, userId uint32) (err error)
}

type profileImageUseCase struct {
	gcsService          service.GCSService
	imagickService      service.ImagickService
	voisingFcAPIService service.VoisingFcAPIService
}

func (p *profileImageUseCase) ProfileImageProcess(bucket string, name string, userImageId uint32, userId uint32) (err error) {
	err = p.voisingFcAPIService.PatchUserImageStatus(userImageId, media.UserImageStatus_PROCEED)
	if err != nil {
		return err
	}
	blob, err := p.gcsService.GetObjectBlob(bucket, name)
	if err != nil {
		return err
	}
	buffer, err := p.imagickService.DecodeBase64(bytes.NewBuffer(blob))
	if err != nil {
		return err
	}
	newImageBuffer, err := p.imagickService.ConvertResize(&buffer, ProfileImageWidth, ProfileImageHeight)
	newImageBuffer, err = p.imagickService.ConvertFormat(&newImageBuffer, ProfileImageFormat)
	if err != nil {
		return err
	}
	err = p.gcsService.CreateObject(newImageBuffer.Bytes(), os.Getenv("DEST_BUCKET"), name+"."+ProfileImageFormat, fmt.Sprintf("image/%s", ProfileImageFormat))
	if err != nil {
		return err
	}
	objectName := strings.Split(name, "/")[len(strings.Split(name, "/"))]
	err = p.voisingFcAPIService.PatchUserImageName(userImageId, objectName+"."+ProfileImageFormat)
	if err != nil {
		return err
	}
	err = p.gcsService.DeleteObject(bucket, name)
	if err != nil {
		return err
	}
	err = p.voisingFcAPIService.PatchUserImageStatus(userImageId, media.UserImageStatus_COMPLETED)
	if err != nil {
		return err
	}
	err = p.voisingFcAPIService.CreateUserProfileImage(userImageId, userId)
	if err != nil {
		return err
	}
	return err
}

func NewProfileImageUseCase(gcsService service.GCSService, imagickService service.ImagickService, voisingFcAPIService service.VoisingFcAPIService) ProfileImageUseCase {
	return &profileImageUseCase{
		gcsService:          gcsService,
		imagickService:      imagickService,
		voisingFcAPIService: voisingFcAPIService,
	}
}
