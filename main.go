package VoisingMediaProcessor

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/googleapis/google-cloudevents-go/cloud/storagedata"
	"google.golang.org/protobuf/encoding/protojson"
	"strings"
	"voising/domain/service"
	"voising/usecase"
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
	defer imagickService.Close()

	switch {
	case strings.HasPrefix(gcsEvent.Name, "profiles"):
		profileImageUserCase := usecase.NewProfileImageUseCase(gcsService, imagickService)
		err := profileImageUserCase.ProfileImageProcess(gcsEvent.Bucket, gcsEvent.Name)
		if err != nil {
			return fmt.Errorf("profile image process error: %v", err)
		}
	case strings.HasPrefix(strings.Split(gcsEvent.Name, "/")[0]+"/"+strings.Split(gcsEvent.Name, "/")[2], "reports/cover"):
		fmt.Println("cover")
	case strings.HasPrefix(strings.Split(gcsEvent.Name, "/")[0]+"/"+strings.Split(gcsEvent.Name, "/")[2], "reports/content"):
		fmt.Println("content")
	default:
		return fmt.Errorf("object name is not match: %v", gcsEvent.Name)
	}

	return nil
}
