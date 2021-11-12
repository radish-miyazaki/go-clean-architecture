package diary

import (
	"context"

	"github.com/radish-miyazaki/go-clean-architecture/domain/model"
	"github.com/radish-miyazaki/go-clean-architecture/usecase/repository"
)

type UpdateDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

func NewUpdateDiaryUsecase(dr repository.DiaryRepository) *UpdateDiaryUsecase {
	return &UpdateDiaryUsecase{diaryRepo: dr}
}

type UpdateDiaryInputPort struct {
	ID          int
	Title       string
	Description string
}

type UpdateDiaryOutputPort struct {
	Diary *model.Diary
}

func (du *UpdateDiaryUsecase) Execute(
	ctx context.Context,
	in *UpdateDiaryInputPort,
) (*UpdateDiaryOutputPort, error) {
	diary := &model.Diary{
		ID:          in.ID,
		Title:       in.Title,
		Description: in.Description,
	}

	diary, err := du.diaryRepo.Update(ctx, diary)
	if err != nil {
		return nil, err
	}

	return &UpdateDiaryOutputPort{Diary: diary}, nil
}
