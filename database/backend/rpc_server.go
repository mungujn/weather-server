package backend

import (
	"log"
	"strings"

	pb "github.com/mungujn/weather-server/database/services"

	"golang.org/x/net/context"
)

const (
	ap = 10
)

// Server is used to implement pb.DatabaseService
type Server struct{}

// Create implements pb.DatabaseService.Create
func (s *Server) Create(ctx context.Context, in *pb.Data) (*pb.Result, error) {
	p := strings.Repeat("-", ap)
	log.Printf("%s Db create rpc call started %s", p, p)

	err := CreateData(in.Location, in.Values)

	if err != nil {
		return getError(err)
	}

	log.Printf(strings.Repeat("-", ap*3))

	return &pb.Result{
		Status:  true,
		Message: "Created",
	}, nil
}

// Create implements pb.DatabaseService.Read
func (s *Server) Read(ctx context.Context, in *pb.Data) (*pb.Data, error) {
	p := strings.Repeat("-", ap)
	log.Printf("%s Db Read rpc call started %s", p, p)

	data, err := ReadData(in.Location)
	if err != nil {
		return nil, err
	}

	log.Printf(strings.Repeat("-", ap*3))

	return &pb.Data{
		Values: data,
	}, nil
}

// Update implements pb.DatabaseService.Create
func (s *Server) Update(ctx context.Context, in *pb.Data) (*pb.Result, error) {
	p := strings.Repeat("-", ap)
	log.Printf("%s Db update rpc call started %s", p, p)

	err := UpdateData(in.Location, in.Values)

	if err != nil {
		return getError(err)
	}

	log.Printf(strings.Repeat("-", ap*3))

	return &pb.Result{
		Status:  true,
		Message: "Data updated",
	}, nil
}

// Delete implements pb.DatabaseService.Create
func (s *Server) Delete(ctx context.Context, in *pb.Data) (*pb.Result, error) {
	p := strings.Repeat("-", ap)
	log.Printf("%s Db delete rpc call started %s", p, p)

	err := DeleteData(in.Location)

	if err != nil {
		return getError(err)
	}

	log.Printf(strings.Repeat("-", ap*3))

	return &pb.Result{
		Status:  false,
		Message: "Data updated",
	}, nil
}

func getError(err error) (*pb.Result, error) {
	return &pb.Result{
		Status:  false,
		Message: err.Error(),
	}, err
}
