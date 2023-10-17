package client

import (
	"context"
	"time"

	gmodels "github.com/lwinmgmg/gmodels/golang/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GConn *grpc.ClientConn = nil
)

func init() {
	if GConn == nil {
		var err error
		if GConn, err = grpc.Dial("localhost:8069", grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
			panic(err)
		}
	}
}

func GetUserByCode(code string) (*gmodels.User, error) {
	client := gmodels.NewUserServiceClient(GConn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.GetUserByCode(ctx, &gmodels.GetUserByCodeRequest{Code: code})
}

func GetUserByToken(token string) (*gmodels.User, error) {
	client := gmodels.NewUserServiceClient(GConn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.GetProfile(ctx, &gmodels.GetProfileRequest{Token: token})
}
