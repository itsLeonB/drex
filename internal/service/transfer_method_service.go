package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/itsLeonB/drex/internal/dto"
	"github.com/itsLeonB/drex/internal/entity"
	"github.com/itsLeonB/drex/internal/mapper"
	"github.com/itsLeonB/drex/internal/repository"
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/go-crud"
	"github.com/itsLeonB/ungerr"
	"github.com/rotisserie/eris"
)

type transferMethodServiceImpl struct {
	transferMethodRepository repository.TransferMethodRepository
}

func NewTransferMethodService(transferMethodRepository repository.TransferMethodRepository) TransferMethodService {
	return &transferMethodServiceImpl{transferMethodRepository}
}

func (tms *transferMethodServiceImpl) GetAll(ctx context.Context) ([]dto.TransferMethodResponse, error) {
	transferMethods, err := tms.transferMethodRepository.FindAll(ctx, crud.Specification[entity.TransferMethod]{})
	if err != nil {
		return nil, err
	}

	return ezutil.MapSlice(transferMethods, mapper.TransferMethodToResponse), nil
}

func (tms *transferMethodServiceImpl) GetByID(ctx context.Context, id uuid.UUID) (entity.TransferMethod, error) {
	spec := crud.Specification[entity.TransferMethod]{}
	spec.Model.ID = id

	transferMethod, err := tms.transferMethodRepository.FindFirst(ctx, spec)
	if err != nil {
		return entity.TransferMethod{}, err
	}

	if transferMethod.IsZero() {
		return entity.TransferMethod{}, ungerr.NotFoundError(fmt.Sprintf(appconstant.ErrTransferMethodNotFound, id))
	}

	if transferMethod.IsDeleted() {
		return entity.TransferMethod{}, ungerr.UnprocessableEntityError(fmt.Sprintf(appconstant.ErrTransferMethodDeleted, id))
	}

	return transferMethod, nil
}

func (tms *transferMethodServiceImpl) GetByName(ctx context.Context, name string) (entity.TransferMethod, error) {
	spec := crud.Specification[entity.TransferMethod]{}
	spec.Model.Name = name

	transferMethod, err := tms.transferMethodRepository.FindFirst(ctx, spec)
	if err != nil {
		return entity.TransferMethod{}, err
	}

	if transferMethod.IsZero() {
		return entity.TransferMethod{}, eris.Errorf("%s transfer method not found", name)
	}

	if transferMethod.IsDeleted() {
		return entity.TransferMethod{}, eris.Errorf("%s transfer method is deleted", name)
	}

	return transferMethod, nil
}
