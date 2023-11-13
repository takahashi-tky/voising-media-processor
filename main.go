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
	"strings"
)

func init() {
	functions.CloudEvent("main", main)
}

func main(_ context.Context, e event.Event) error {
	log.Println("Start processing")
	ctx := context.Background()
	var gcsEvent storagedata.StorageObjectData
	if err := protojson.Unmarshal(e.Data(), &gcsEvent); err != nil {
		return fmt.Errorf("protojson.Unmarshal: failed to decode event data: %w", err)
	}
	log.Println(gcsEvent.Bucket)
	log.Println(gcsEvent.Name)
	gcsService := service.NewGCSService(ctx)
	imagickService := service.NewImagickService()

	switch {
	case strings.HasPrefix(gcsEvent.Name, "profiles"):
		profileImageUserCase := usecase.NewProfileImageUseCase(gcsService, imagickService)
		err := profileImageUserCase.ProfileImageProcess(gcsEvent.Bucket, gcsEvent.Name)
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
