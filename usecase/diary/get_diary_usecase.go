/*
	Usecase Layer
	Inputがどこから来るか、Outputが何に利用されるか、データストレージに何が利用されているかなど
	一切気にせずに実装をするレイヤー。
	複雑な処理はExecuteに閉じ込めておけばOK。
*/
package diary

import (
	"context"
	"github.com/radish-miyazaki/go-clean-architecture/domain/model"
	"github.com/radish-miyazaki/go-clean-architecture/usecase/repository"
)

type GetDiaryInputPort struct {
	ID int
}

type GetDiaryOutputPort struct {
	Diary *model.Diary
}

type GetDiaryUsecase struct {
	diaryRepo repository.DiaryRepository
}

func NewGetDiaryUsecase(diaryRepo repository.DiaryRepository) *GetDiaryUsecase {
	return &GetDiaryUsecase{diaryRepo: diaryRepo}
}

func (du *GetDiaryUsecase) Execute(
	ctx context.Context,
	in *GetDiaryInputPort,
) (*GetDiaryOutputPort, error) {
	diary, err := du.diaryRepo.FindByID(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	return &GetDiaryOutputPort{Diary: diary}, nil
}
