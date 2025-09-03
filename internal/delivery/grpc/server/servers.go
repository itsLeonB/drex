package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/itsLeonB/drex-protos/gen/go/debt/v1"
	"github.com/itsLeonB/drex-protos/gen/go/transaction/v1"
	"github.com/itsLeonB/drex/internal/provider"
	"github.com/rotisserie/eris"
	"google.golang.org/grpc"
)

type Servers struct {
	TransferMethod  transaction.TransferMethodServiceServer
	DebtTransaction debt.DebtServiceServer
}

func ProvideServers(services *provider.Services) *Servers {
	validate := validator.New()

	return &Servers{
		TransferMethod:  newTransferMethodServer(services.TransferMethod),
		DebtTransaction: newDebtTransactionServer(validate, services.DebtTransaction),
	}
}

func (s *Servers) Register(grpcServer *grpc.Server) error {
	if s.TransferMethod == nil {
		return eris.New("transfer method server is nil")
	}
	if s.DebtTransaction == nil {
		return eris.New("debt transaction server is nil")
	}

	transaction.RegisterTransferMethodServiceServer(grpcServer, s.TransferMethod)
	debt.RegisterDebtServiceServer(grpcServer, s.DebtTransaction)

	return nil
}
