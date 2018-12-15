echo ./api/build/ ./weather/build/ | xargs -n 1 cp ./build/ca-certificates.crt
echo ./api/build/ ./weather/build/ | xargs -n 1 cp ./build/passwd
cd api
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -installsuffix cgo -ldflags="-w -s" -o ./build/wapi
# docker build -t mungujn/wapi .
cd ..
cd weather
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -installsuffix cgo -ldflags="-w -s" -o ./build/weather
# docker build -t mungujn/weather .
cd ..
