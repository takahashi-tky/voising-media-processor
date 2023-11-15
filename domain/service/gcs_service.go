package service

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"
)

type GCSService interface {
	GetObjectBlob(bucket string, name string) ([]byte, error)
	GetObjectMetaData(bucket string, name string) (metadata map[string]string, err error)
	CreateObject(blob []byte, destBucket string, name string, contentType string) error
	DeleteObject(bucket string, name string) error
}

type gcsService struct {
	ctx           *context.Context
	storageClient *storage.Client
}

func (g *gcsService) GetObjectMetaData(bucket string, name string) (metadata map[string]string, err error) {
	obj := g.storageClient.Bucket(bucket).Object(name)
	attrs, err := obj.Attrs(*g.ctx)
	if err != nil {
		return nil, err
	}
	return attrs.Metadata, err
}

func (g *gcsService) DeleteObject(bucket string, name string) error {
	return g.storageClient.Bucket(bucket).Object(name).Delete(*g.ctx)
}

func (g *gcsService) CreateObject(blob []byte, destBucket string, name string, contentType string) error {
	wc := g.storageClient.Bucket(destBucket).Object(name).NewWriter(*g.ctx)
	wc.ContentType = contentType
	if _, err := wc.Write(blob); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func (g *gcsService) GetObjectBlob(bucket string, name string) (bytes []byte, err error) {
	obj := g.storageClient.Bucket(bucket).Object(name)
	reader, err := obj.NewReader(*g.ctx)
	if err != nil {
		return nil, fmt.Errorf("obj.NewReader: %v", err)
	}
	defer func(reader *storage.Reader) {
		err := reader.Close()
		if err != nil {
			log.Fatalf("reader.Close: %v", err)
		}
	}(reader)

	blob, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v", err)
	}
	if strings.Index(string(blob), ";base64,") > 0 {
		blob, err = base64.StdEncoding.DecodeString(string(blob))
		if err != nil {
			return nil, fmt.Errorf("base64.StdEncoding.DecodeString: %v", err)
		}
	}
	return blob, nil
}

func NewGCSService(ctx context.Context) GCSService {
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
	return &gcsService{
		ctx:           &ctx,
		storageClient: storageClient,
	}
}
