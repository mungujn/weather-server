package mock_services_test
package mock_services_test

import (
	"time"
	"golang.org/x/net/context"
	"fmt"
	"testing"

	pbmock "github.com/mungujn/weather/common/mock_services"
	pb "github.com/mungujn/weather/common/services"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
)

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

//TestWeatherService tests rpc methods in the WeatherService
func TestWeatherService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeatherClient := pbmock.NewMockWeatherServiceClient(ctrl)

	req := &pb.LocationAndDate{Location: "location", Date: "today"}
	mockWeatherClient.EXPECT().GetWeather(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&pb.Weather{Temperature: 28, Location: "kampala", Date: "today"}, nil)

	testGetWeather(t, mockWeatherClient)
}

func testGetWeather(t *testing.T, client pb.WeatherServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetWeather(ctx, &pb.LocationAndDate{Location: "location", Date: "today"})

	if err != nil || response.Temperature != 28 || response.Location != "kampala" || response.Date != "today" {
		t.Errorf("Get weather mocking failed")
	}
	t.Log("Response : ", response)
}
