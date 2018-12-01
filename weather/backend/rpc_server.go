package backend

import (
	"log"
	"strings"

	pb "github.com/mungujn/weather-server/weather/services"

	"golang.org/x/net/context"
)

const (
	ap = 10
)

// Server is used to implement pb.DatabaseService
type Server struct{}

// GetWeather implements pb.WeatherService.Create
func (s *Server) GetWeather(ctx context.Context, in *pb.LocationAndDate) (*pb.Weather, error) {
	p := strings.Repeat("-", ap)
	log.Printf("%s Get weather rpc call started %s", p, p)

	weather, err := GetWeather(in.Location, in.Date)

	if err != nil {
		return nil, err
	}

	log.Printf(strings.Repeat("-", ap*3))

	return weather, nil
}
