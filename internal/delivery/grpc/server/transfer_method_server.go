package server

import (
	"context"

	"github.com/itsLeonB/drex-protos/gen/go/transaction/v1"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/service"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/gerpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transferMethodServer struct {
	transaction.UnimplementedTransferMethodServiceServer
	transferMethodSvc service.TransferMethodService
}

func newTransferMethodServer(transferMethodSvc service.TransferMethodService) transaction.TransferMethodServiceServer {
	return &transferMethodServer{transferMethodSvc: transferMethodSvc}
}

func (tms *transferMethodServer) GetAll(ctx context.Context, _ *emptypb.Empty) (*transaction.GetAllResponse, error) {
	response, err := tms.transferMethodSvc.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	mapFunc := func(resp dto.TransferMethodResponse) *transaction.TransferMethodResponse {
		return &transaction.TransferMethodResponse{
			Id:        resp.ID.String(),
			Name:      resp.Name,
			Display:   resp.Display,
			CreatedAt: gerpc.NullableTimeToProto(resp.CreatedAt),
			UpdatedAt: gerpc.NullableTimeToProto(resp.UpdatedAt),
			DeletedAt: gerpc.NullableTimeToProto(resp.DeletedAt),
		}
	}

	return &transaction.GetAllResponse{TransferMethods: ezutil.MapSlice(response, mapFunc)}, nil
}
