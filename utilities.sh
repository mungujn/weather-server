// protoc. Make sure you have protoc on your path
cd weather
protoc -I weather/services/ weather/services/weather.proto --go_out=plugins=grpc:weather/services
cd ..

// TLS/SSL
cd data
openssl genrsa -out wapi.key 2048
openssl req -new -x509 -sha256 -key wapi.key -out wapi.crt -days 30
openssl genrsa -out weather.key 2048
openssl req -new -x509 -sha256 -key weather.key -out weather.crt -days 30
cd ..

// mockgen
mockgen github.com/mungujn/weather-server/weather/services WeatherServiceClient > weather/mock_services/weather_service_mock.go


// All Tests
go test github.com/mungujn/weather-server/weather/tests -v
go test github.com/mungujn/weather-server/api/tests -v




