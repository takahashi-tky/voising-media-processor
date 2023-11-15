package usecase

import (
	"bytes"
	"fmt"
	"irelove.ireisu.com/api/proto/gen/media"
	"irelove.ireisu.com/domain/service"
	"os"
)

type UserReportCoverUseCase interface {
	UserReportCoverProcess(bucket string, name string, userImageId uint32) (err error)
}

type userReportCoverUseCase struct {
	gcsService          service.GCSService
	imagickService      service.ImagickService
	voisingFcAPIService service.VoisingFcAPIService
}

func (u *userReportCoverUseCase) UserReportCoverProcess(bucket string, name string, userImageId uint32) (err error) {
	err = u.voisingFcAPIService.PatchUserImageStatus(userImageId, media.UserImageStatus_PROCEED)
	if err != nil {
		return err
	}
	blob, err := u.gcsService.GetObjectBlob(bucket, name)
	if err != nil {
		return err
	}
	objectFormat, err := u.imagickService.GetFileFormat(bytes.NewBuffer(blob))

	switch objectFormat {
	case "jpeg", "jpg", "png", "gif":
	default:
		return fmt.Errorf(fmt.Sprintf("object format is not match: %s", objectFormat))
	}
	err = u.gcsService.CreateObject(bytes.NewBuffer(blob).Bytes(), os.Getenv("DEST_BUCKET"), name+"."+objectFormat, fmt.Sprintf("image/%s", objectFormat))
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

func NewUserReportCoverUseCase(gcsService service.GCSService, imagickService service.ImagickService, voisingFcAPIService service.VoisingFcAPIService) UserReportCoverUseCase {
	return &userReportCoverUseCase{
		gcsService:          gcsService,
		imagickService:      imagickService,
		voisingFcAPIService: voisingFcAPIService,
	}
}