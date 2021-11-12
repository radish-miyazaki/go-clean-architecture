package diary

import (
	"context"

	"github.com/radish-miyazaki/go-clean-architecture/domain/model"
	"github.com/radish-miyazaki/go-clean-architecture/usecase/repository"
)

type ListDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

func NewListDiaryUsecase(dr repository.DiaryRepository) *ListDiaryUsecase {
	return &ListDiaryUsecase{diaryRepo: dr}
}

type ListDiaryInputPort struct{}

type ListDiaryOutputPort struct {
	Diaries []*model.Diary
}

func (du *ListDiaryUsecase) Execute(
	ctx context.Context,
	_ *ListDiaryInputPort,
) (*ListDiaryOutputPort, error) {
	diaries, err := du.diaryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return &ListDiaryOutputPort{Diaries: diaries}, nil
}
