package client

import (
	"context"
	"log"
	"time"

	gmodels "github.com/lwinmgmg/gmodels/golang/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GConn *grpc.ClientConn = nil
)

func GetClientConn() *grpc.ClientConn {
	var err error
	if GConn == nil || GConn.GetState() != connectivity.Ready {
		GConn, err = grpc.Dial("localhost:8069", grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if err != nil {
		log.Println("Can't connect to the client", err)
	}
	return GConn
}

func GetUserByCode(code string) (*gmodels.User, error) {
	client := gmodels.NewUserServiceClient(GetClientConn())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.GetUserByCode(ctx, &gmodels.GetUserByCodeRequest{Code: code})
}

func GetUserByToken(token string) (*gmodels.User, error) {
	client := gmodels.NewUserServiceClient(GetClientConn())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.GetProfile(ctx, &gmodels.GetProfileRequest{Token: token})
}
