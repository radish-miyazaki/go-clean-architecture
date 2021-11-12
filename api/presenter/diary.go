/*
	Presenter Layer
	アプリケーションの出力部分を担当。
	UseCaseのOutputPortを出力用の構造体に変換して出力。
*/

package presenter

import (
	"github.com/radish-miyazaki/go-clean-architecture/domain/model"
	"github.com/radish-miyazaki/go-clean-architecture/usecase/diary"
)

// simpleDiaryView レスポンス内部の構造体
type simpleDiaryView struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// mapToDiarySimpleDiaryView Domainをレスポンス用の内部構造体に変換
func mapToDiaryToSimpleDiaryView(diary *model.Diary) *simpleDiaryView {
	return &simpleDiaryView{
		ID:          diary.ID,
		Title:       diary.Title,
		Description: diary.Description,
	}
}

type createDiaryResponse struct {
	Diary *simpleDiaryView `json:"diary"`
}

func NewCreateDiaryPresenter(output *diary.CreateDiaryOutputPort) *createDiaryResponse {
	return &createDiaryResponse{Diary: mapToDiaryToSimpleDiaryView(output.Diary)}
}

type listDiaryResponse struct {
	Diaries []*simpleDiaryView `json:"diaries"`
}

func NewListDiaryPresenter(output *diary.ListDiaryOutputPort) *listDiaryResponse {
	var diaries []*simpleDiaryView

	for _, row := range output.Diaries {
		diaries = append(diaries, mapToDiaryToSimpleDiaryView(row))
	}

	return &listDiaryResponse{Diaries: diaries}
}

type updateDiaryResponse struct {
	Diary *simpleDiaryView `json:"diary"`
}

func NewUpdateDiaryPresenter(output *diary.UpdateDiaryOutputPort) *updateDiaryResponse {
	return &updateDiaryResponse{Diary: mapToDiaryToSimpleDiaryView(output.Diary)}
}

type deleteDiaryResponse struct{}

func NewDeleteDiaryPresenter(_ *diary.DeleteDiaryOutputPort) *deleteDiaryResponse {
	return &deleteDiaryResponse{}
}

type getDiaryResponse struct {
	Diary *simpleDiaryView `json:"diary"`
}

func NewGetDiaryPresenter(output *diary.GetDiaryOutputPort) *getDiaryResponse {
	return &getDiaryResponse{Diary: mapToDiaryToSimpleDiaryView(output.Diary)}
}
