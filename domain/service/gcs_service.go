package service

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"log"
)

type GCSService interface {
	GetObjectReader(bucket string, name string) (*storage.Reader, error)
	GetObjectBlob(bucket string, name string) ([]byte, error)
	CreateObject(blob []byte, destBucket string, name string, contentType string) error
	DeleteObject(bucket string, name string) error
}

type gcsService struct {
	ctx           *context.Context
	storageClient *storage.Client
}

func (g *gcsService) GetObjectReader(bucket string, name string) (*storage.Reader, error) {
	bucketH := g.storageClient.Bucket(bucket)
	it := bucketH.Objects(*g.ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Bucket(%q).Objects: %w", bucketH, err)
		}
		log.Println(attrs.Name)
	}
	obj := g.storageClient.Bucket(bucket).Object(name)
	return obj.NewReader(*g.ctx)
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
