package diary

import (
	"context"

	"github.com/radish-miyazaki/go-clean-architecture/usecase/repository"
)

type DeleteDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

func NewDeleteDiaryUsecase(dr repository.DiaryRepository) *DeleteDiaryUsecase {
	return &DeleteDiaryUsecase{diaryRepo: dr}
}

type DeleteDiaryInputPort struct {
	ID int
}

type DeleteDiaryOutputPort struct {
}

func (du *DeleteDiaryUsecase) Execute(
	ctx context.Context,
	in *DeleteDiaryInputPort,
) (*DeleteDiaryOutputPort, error) {
	if err := du.diaryRepo.Delete(ctx, in.ID); err != nil {
		return nil, err
	}

	return &DeleteDiaryOutputPort{}, nil
}
