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
	"log"
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
	log.Println(gcsEvent.Metadata)
	userImageId, err := strconv.Atoi(gcsEvent.Metadata["user-image-id"])
	if err != nil {
		return fmt.Errorf("user-image-id is not number: %v", gcsEvent.Metadata["user-image-id"])
	}
	switch {
	case strings.HasPrefix(gcsEvent.Name, "profiles"):
		log.Println(gcsEvent.Metadata)
		profileImageUserCase := usecase.NewProfileImageUseCase(gcsService, imagickService, voisingFcAPIService)
		err := profileImageUserCase.ProfileImageProcess(gcsEvent.Bucket, gcsEvent.Name, uint32(userImageId))
		if err != nil {
			return fmt.Errorf("profile image process error: %v", err)
		}
	case strings.HasPrefix(strings.Split(gcsEvent.Name, "/")[0]+"/"+gcsEvent.Name, "reports/cover"):
		fmt.Println("cover")
	case strings.HasPrefix(strings.Split(gcsEvent.Name, "/")[0]+"/"+gcsEvent.Name, "reports/content"):
		fmt.Println("content")
	default:
		return fmt.Errorf("object name is not match: %v", gcsEvent.Name)
	}

	return nil
}
