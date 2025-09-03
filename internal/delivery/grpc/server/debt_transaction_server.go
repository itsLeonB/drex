package server

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/itsLeonB/drex-protos/gen/go/debt/v1"
	"github.com/itsLeonB/drex/internal/delivery/grpc/mapper"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/service"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/ungerr"
	"google.golang.org/protobuf/types/known/emptypb"
)

type debtTransactionServer struct {
	debt.UnimplementedDebtServiceServer
	validate           *validator.Validate
	debtTransactionSvc service.DebtTransactionService
}

func newDebtTransactionServer(validate *validator.Validate, debtTransactionSvc service.DebtTransactionService) debt.DebtServiceServer {
	return &debtTransactionServer{
		validate:           validate,
		debtTransactionSvc: debtTransactionSvc,
	}
}

func (dts *debtTransactionServer) RecordNewTransaction(ctx context.Context, req *debt.RecordNewTransactionRequest) (*debt.RecordNewTransactionResponse, error) {
	userProfileID, err := ezutil.Parse[uuid.UUID](req.GetUserProfileId())
	if err != nil {
		return nil, err
	}

	friendProfileID, err := ezutil.Parse[uuid.UUID](req.GetFriendProfileId())
	if err != nil {
		return nil, err
	}

	transferMethodID, err := ezutil.Parse[uuid.UUID](req.GetTransferMethodId())
	if err != nil {
		return nil, err
	}

	request := dto.NewDebtTransactionRequest{
		UserProfileID:    userProfileID,
		FriendProfileID:  friendProfileID,
		Action:           mapper.FromProtoTransactionAction(req.GetAction()),
		Amount:           ezutil.MoneyToDecimal(req.GetAmount()),
		TransferMethodID: transferMethodID,
		Description:      req.GetDescription(),
	}

	if err = dts.validate.Struct(request); err != nil {
		return nil, err
	}

	response, err := dts.debtTransactionSvc.RecordNew(ctx, request)
	if err != nil {
		return nil, err
	}

	return &debt.RecordNewTransactionResponse{
		Transaction: mapper.ToTransactionProto(response),
	}, nil
}

func (dts *debtTransactionServer) GetTransactions(ctx context.Context, req *debt.GetTransactionsRequest) (*debt.GetTransactionsResponse, error) {
	userProfileID, err := ezutil.Parse[uuid.UUID](req.GetUserProfileId())
	if err != nil {
		return nil, err
	}

	responses, err := dts.debtTransactionSvc.GetAll(ctx, userProfileID)
	if err != nil {
		return nil, err
	}

	return &debt.GetTransactionsResponse{
		Transactions: ezutil.MapSlice(responses, mapper.ToTransactionProto),
	}, nil
}

func (dts *debtTransactionServer) ProcessConfirmedGroupExpense(ctx context.Context, req *debt.ProcessConfirmedGroupExpenseRequest) (*emptypb.Empty, error) {
	groupExpense := req.GetGroupExpense()
	if groupExpense == nil {
		return nil, ungerr.BadRequestError("groupExpense is nil")
	}

	id, err := ezutil.Parse[uuid.UUID](groupExpense.GetId())
	if err != nil {
		return nil, err
	}

	payerProfileID, err := ezutil.Parse[uuid.UUID](groupExpense.GetPayerProfileId())
	if err != nil {
		return nil, err
	}

	creatorProfileID, err := ezutil.Parse[uuid.UUID](groupExpense.GetCreatorProfileId())
	if err != nil {
		return nil, err
	}

	mapFunc := func(participant *debt.ExpenseParticipantData) (dto.ExpenseParticipantData, error) {
		if participant == nil {
			return dto.ExpenseParticipantData{}, ungerr.BadRequestError("expense participant is nil")
		}

		profileID, err := ezutil.Parse[uuid.UUID](participant.GetProfileId())
		if err != nil {
			return dto.ExpenseParticipantData{}, err
		}

		return dto.ExpenseParticipantData{
			ProfileID:   profileID,
			ShareAmount: ezutil.MoneyToDecimal(participant.GetShareAmount()),
		}, nil
	}

	participants, err := ezutil.MapSliceWithError(groupExpense.GetParticipants(), mapFunc)
	if err != nil {
		return nil, err
	}

	request := dto.GroupExpenseData{
		ID:               id,
		PayerProfileID:   payerProfileID,
		CreatorProfileID: creatorProfileID,
		Description:      groupExpense.GetDescription(),
		Participants:     participants,
	}

	if err = dts.validate.Struct(request); err != nil {
		return nil, err
	}

	err = dts.debtTransactionSvc.ProcessConfirmedGroupExpense(ctx, request)

	return nil, err
}

func (dts *debtTransactionServer) GetAllByProfileIds(ctx context.Context, req *debt.GetAllByProfileIdsRequest) (*debt.GetAllByProfileIdsResponse, error) {
	userProfileID, err := ezutil.Parse[uuid.UUID](req.GetUserProfileId())
	if err != nil {
		return nil, err
	}

	friendProfileID, err := ezutil.Parse[uuid.UUID](req.GetFriendProfileId())
	if err != nil {
		return nil, err
	}

	transactions, err := dts.debtTransactionSvc.FindAllByProfileIDs(ctx, userProfileID, friendProfileID)
	if err != nil {
		return nil, err
	}

	return &debt.GetAllByProfileIdsResponse{
		Transactions: ezutil.MapSlice(transactions, mapper.ToTransactionProto),
	}, nil
}
