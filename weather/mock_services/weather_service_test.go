package mock_services_test

import (
	"fmt"
	"testing"
	"time"

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

	// test
	response, err := client.GetWeather(ctx, &pb.LocationAndDate{Location: "location", Date: "today"})

	// assert
	if err != nil || response.Temperature != 28 || response.Location != "kampala" || response.Date != "today" {
		t.Errorf("Get weather mocking failed")
	}
	t.Log("Response : ", response)
}

// get rpc client
func getClient(ctrl *gomock.Controller) pb.WeatherServiceClient {
	if client == nil {
		newClient := pbmock.NewMockWeatherServiceClient(ctrl)

		req := &pb.LocationAndDate{Location: "location", Date: "today"}

		newClient.EXPECT().GetWeather(
			gomock.Any(),
			&rpcMsg{msg: req},
		).Return(&pb.Weather{Temperature: 28, Location: "kampala", Date: "today"}, nil)

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
