package mock_services_test

import (
	"testing"
	"time"

	"golang.org/x/net/context"

	pbmock "github.com/mungujn/weather-server/database/mock_services"
	pb "github.com/mungujn/weather-server/database/services"

	"github.com/golang/mock/gomock"
)

var client pb.DatabaseServiceClient

//TestCreate tests the create rpc method in pb.DatabaseService
func TestCreate(t *testing.T) {
	// set up
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := getClient(ctrl)

	values := make(map[string]string)
	values["temperature"] = "28"

	// test
	response, err := client.Create(ctx, &pb.Data{Location: "kampala/today", Values: values})

	// assert
	if err != nil || response.Status != true || response.Message != "created" {
		t.Errorf("Create mocking failed")
	}
	t.Log("Response : ", response)
}

// get rpc client
func getClient(ctrl *gomock.Controller) pb.DatabaseServiceClient {
	if client == nil {
		newClient := pbmock.NewMockDatabaseServiceClient(ctrl)

		newClient.EXPECT().Create(
			gomock.Any(),
			gomock.Any(),
		).Return(&pb.Result{Status: true, Message: "created"}, nil)

		return newClient
	}

	return client
}
