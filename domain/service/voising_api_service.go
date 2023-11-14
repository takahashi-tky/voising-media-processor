package service

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"irelove.ireisu.com/api/proto/gen/media"
	"os"
	"time"
)

type VoisingFcAPIService interface {
	PatchUserImageStatus(userImageId uint32, status media.UserImageStatus) error
	PatchUserImageName(userImageId uint32, name string) error
	Close() error
}

type voisingFcAPIService struct {
	client media.MediaClient
	conn   *grpc.ClientConn
}

func (v *voisingFcAPIService) PatchUserImageName(userImageId uint32, name string) error {
	request := &media.PatchUserImageNameRequest{
		UserImageId: userImageId,
		Name:        name,
	}
	ctx, cancel, err := getAuthContext()
	defer cancel()
	if err != nil {
		return err
	}
	_, err = v.client.PathUserImageName(ctx, request)
	if err != nil {
		return err
	}
	return err

}

func (v *voisingFcAPIService) PatchUserImageStatus(userImageId uint32, status media.UserImageStatus) error {
	request := &media.PatchUserImageStatusRequest{
		UserImageId: userImageId,
		Status:      status,
	}
	ctx, cancel, err := getAuthContext()
	defer cancel()
	if err != nil {
		return err
	}
	_, err = v.client.PatchUserImageStatus(ctx, request)
	if err != nil {
		return err
	}
	return err
}

func getAuthContext() (ctx context.Context, cancel context.CancelFunc, err error) {
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	tokenSource, err := idtoken.NewTokenSource(ctx, os.Getenv("GRPC_SERVER_AUDIENCE"))
	if err != nil {
		cancel()
		return nil, nil, fmt.Errorf("idtoken.NewTokenSource: %w", err)
	}
	token, err := tokenSource.Token()
	if err != nil {
		cancel()
		return nil, nil, fmt.Errorf("TokenSource.Token: %w", err)
	}
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token.AccessToken)
	return ctx, cancel, err
}

func (v *voisingFcAPIService) Close() error {
	return v.conn.Close()
}

func NewVoisingFcAPIService() VoisingFcAPIService {
	apiServerHost := os.Getenv("GRPC_SERVER_HOST")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(apiServerHost))
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred))
	conn, err := grpc.Dial(apiServerHost, opts...)
	if err != nil {
		panic(err)
	}
	return &voisingFcAPIService{
		client: media.NewMediaClient(conn),
		conn:   conn,
	}
}
