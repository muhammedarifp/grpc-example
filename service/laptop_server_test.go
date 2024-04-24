package service_test

import (
	"context"
	"testing"

	"github.com/muhammedarifp/grpc-example/pb"
	"github.com/muhammedarifp/grpc-example/sample"
	"github.com/muhammedarifp/grpc-example/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopWithInvalidID := sample.NewLaptop()
	laptopWithInvalidID.Id = "invalid-uuid"

	laptopDuplicateID := sample.NewLaptop()
	store := service.NewInmemoryStore()
	err := store.Save(laptopDuplicateID)
	require.Nil(t, err)

	t.Parallel()
	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInmemoryStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  service.NewInmemoryStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopWithInvalidID,
			store:  service.NewInmemoryStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: laptopDuplicateID,
			store:  store,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			srvr := service.NewLaptopServer(tc.store)
			res, err := srvr.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
