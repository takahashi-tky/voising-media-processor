package usecase

import (
	"bytes"
	"fmt"
	"irelove.ireisu.com/api/proto/gen/media"
	"irelove.ireisu.com/domain/service"
	"os"
)

type UserReportContentImageUseCase interface {
	UserReportContentImageProcess(bucket string, name string, userImageId uint32) (err error)
}

type userReportContentImageUseCase struct {
	gcsService          service.GCSService
	imagickService      service.ImagickService
	voisingFcAPIService service.VoisingFcAPIService
}

func (u *userReportContentImageUseCase) UserReportContentImageProcess(bucket string, name string, userImageId uint32) (err error) {
	err = u.voisingFcAPIService.PatchUserImageStatus(userImageId, media.UserImageStatus_PROCEED)
	if err != nil {
		return err
	}
	blob, err := u.gcsService.GetObjectBlob(bucket, name)
	if err != nil {
		return err
	}

	buffer, err := u.imagickService.DecodeBase64(bytes.NewBuffer(blob))
	if err != nil {
		return err
	}

	objectFormat, err := u.imagickService.GetFileFormat(&buffer)
	if err != nil {
		return err
	}

	switch objectFormat {
	case "jpeg", "jpg", "png", "gif":
	default:
		return fmt.Errorf(fmt.Sprintf("object format is not match: %s", objectFormat))
	}
	err = u.gcsService.CreateObject(buffer.Bytes(), os.Getenv("DEST_BUCKET"), name+"."+objectFormat, fmt.Sprintf("image/%s", objectFormat))
	if err != nil {
		return err
	}
	err = u.gcsService.DeleteObject(bucket, name)
	if err != nil {
		return err
	}
	err = u.voisingFcAPIService.PatchUserImageName(userImageId, name+"."+objectFormat)
	if err != nil {
		return err
	}
	err = u.voisingFcAPIService.PatchUserImageStatus(userImageId, media.UserImageStatus_COMPLETED)
	if err != nil {
		return err
	}
	return err
}

func NewUserReportContentImageUseCase(gcsService service.GCSService, imagickService service.ImagickService, voisingFcAPIService service.VoisingFcAPIService) UserReportContentImageUseCase {
	return &userReportContentImageUseCase{
		gcsService:          gcsService,
		imagickService:      imagickService,
		voisingFcAPIService: voisingFcAPIService,
	}
}
