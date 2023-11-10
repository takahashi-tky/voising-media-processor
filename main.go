package CampersImageProcessor

import (
	"campers.fan/domain/model"
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("main", main)
}

func main(_ context.Context, e event.Event) error {
	var objectData model.StorageObjectData
	if err := e.DataAs(&objectData); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}
	return nil
}
