package VoisingMediaProcessor

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/googleapis/google-cloudevents-go/cloud/storagedata"
	"google.golang.org/protobuf/encoding/protojson"
	"irelove.ireisu.com/domain/service"
	"irelove.ireisu.com/usecase"
	"strconv"
	"strings"
)

func init() {
	functions.CloudEvent("main", main)
}

func main(_ context.Context, e event.Event) error {
	ctx := context.Background()
	var gcsEvent storagedata.StorageObjectData
	if err := protojson.Unmarshal(e.Data(), &gcsEvent); err != nil {
		return fmt.Errorf("protojson.Unmarshal: failed to decode event data: %w", err)
	}

	gcsService := service.NewGCSService(ctx)
	imagickService := service.NewImagickService()
	voisingFcAPIService := service.NewVoisingFcAPIService()

	objectMetadata, err := gcsService.GetObjectMetaData(gcsEvent.Bucket, gcsEvent.Name)
	if err != nil {
		return fmt.Errorf("get object metadata error: %v", err)
	}

	userImageIdStr, exists := objectMetadata["user-image-id"]
	if !exists {
		return fmt.Errorf("user-image-id is not exists")
	}
	userImageId, err := strconv.Atoi(userImageIdStr)
	if err != nil {
		return fmt.Errorf("user-image-id is not number")
	}

	switch {
	case strings.HasPrefix(gcsEvent.Name, "profiles"):
		profileImageUserCase := usecase.NewProfileImageUseCase(gcsService, imagickService, voisingFcAPIService)
		err := profileImageUserCase.ProfileImageProcess(gcsEvent.Bucket, gcsEvent.Name, uint32(userImageId))
		if err != nil {
			return fmt.Errorf("profile image process error: %v", err)
		}
	case strings.HasPrefix(strings.Split(gcsEvent.Name, "/")[0]+"/"+strings.Split(gcsEvent.Name, "/")[2]+"/"+gcsEvent.Name, "reports/cover"):
		userReportCoverUseCase := usecase.NewUserReportCoverUseCase(gcsService, imagickService, voisingFcAPIService)
		err := userReportCoverUseCase.UserReportCoverProcess(gcsEvent.Bucket, gcsEvent.Name, uint32(userImageId))
		if err != nil {
			return fmt.Errorf("user report cover process error: %v", err)
		}
	case strings.HasPrefix(strings.Split(gcsEvent.Name, "/")[0]+"/"+strings.Split(gcsEvent.Name, "/")[2]+"/"+gcsEvent.Name, "reports/content"):
		userReportContentImageUseCase := usecase.NewUserReportContentImageUseCase(gcsService, imagickService, voisingFcAPIService)
		err := userReportContentImageUseCase.UserReportContentImageProcess(gcsEvent.Bucket, gcsEvent.Name, uint32(userImageId))
		if err != nil {
			return fmt.Errorf("user report content image process error: %v", err)
		}
	default:
		return fmt.Errorf("object name is not match: %v", gcsEvent.Name)
	}

	return nil
}
