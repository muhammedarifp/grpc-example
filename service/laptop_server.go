package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/muhammedarifp/grpc-example/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	LaptopStore LaptopStore
}

func NewLaptopServer(
	laptopStore LaptopStore,
) *LaptopServer {
	return &LaptopServer{
		LaptopStore: laptopStore,
	}
}

func (l *LaptopServer) CreateLaptop(ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Println("laptop recived : ", laptop.Id)
	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not valid UUID : %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot genarate new laptop id : %v", err)
		}
		req.Laptop.Id = id.String()
	}

	// save laptop to store
	if err := l.LaptopStore.Save(laptop); err != nil {
		code := codes.Internal
		if err == ErrAldredyExist {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to store : %v", err)

	}

	resp := pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return &resp, nil
}
