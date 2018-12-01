package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/mungujn/weather-server/common/utils"

	"golang.org/x/net/context"

	pbmock "github.com/mungujn/weather-server/weather/mock_services"
	pb "github.com/mungujn/weather-server/weather/services"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
)

var client pb.WeatherServiceClient

//TestWeatherService tests rpc methods in the WeatherService
func TestWeatherService(t *testing.T) {
	// set up
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := getClient(ctrl)

	expected := make(map[string]string)
	expected["location"] = "Kampala"
	expected["date"] = "Today"
	expected["temperature"] = "28"

	// test
	response, err := client.GetWeather(ctx, &pb.LocationAndDate{Location: "kampala", Date: "today"})

	// assert
	if err != nil || !utils.MapsEqual(response.Data, expected) {
		t.Errorf("Get weather mocking failed, expected: %v, got: %v, error: %v", expected, response.Data, err)
	} else {
		t.Log("Response : ", response)
	}
}

// get rpc client
func getClient(ctrl *gomock.Controller) pb.WeatherServiceClient {
	if client == nil {
		newClient := pbmock.NewMockWeatherServiceClient(ctrl)

		req := &pb.LocationAndDate{Location: "kampala", Date: "today"}

		values := make(map[string]string)
		values["location"] = "Kampala"
		values["date"] = "Today"
		values["temperature"] = "28"

		newClient.EXPECT().GetWeather(
			gomock.Any(),
			&rpcMsg{msg: req},
		).Return(&pb.Weather{Data: values}, nil)

		return newClient
	}
	return client
}

type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)

	if !ok {
		return false
	}

	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}
