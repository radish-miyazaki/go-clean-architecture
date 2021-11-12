package diary

import (
	"context"

	"github.com/radish-miyazaki/go-clean-architecture/domain/model"
	"github.com/radish-miyazaki/go-clean-architecture/usecase/repository"
)

type CreateDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

func NewCreateDiaryUsecase(dr repository.DiaryRepository) *CreateDiaryUsecase {
	return &CreateDiaryUsecase{diaryRepo: dr}
}

type CreateDiaryInputPort struct {
	Title       string
	Description string
}

type CreateDiaryOutputPort struct {
	Diary *model.Diary
}

func (du *CreateDiaryUsecase) Execute(
	ctx context.Context,
	in *CreateDiaryInputPort,
) (*CreateDiaryOutputPort, error) {
	diary := &model.Diary{
		Title:       in.Title,
		Description: in.Description,
	}

	diary, err := du.diaryRepo.Store(ctx, diary)
	if err != nil {
		return nil, err
	}

	return &CreateDiaryOutputPort{Diary: diary}, nil
}
