/*
	Adapter Layer
	アプリケーションの入力部分を担当。
	外部からアプリケーションに渡ってきたRequestを、アプリケーションの処理で取り扱う
	UsecaseのInputPortに変換する。
*/

package adapter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/radish-miyazaki/go-clean-architecture/usecase/diary"
)

type createDiaryRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewCreateDiaryInputPortFromRequest(r *http.Request) (*diary.CreateDiaryInputPort, error) {
	input := createDiaryRequestBody{}
	// 渡ってきたJSONをGolangの型に変換
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	return &diary.CreateDiaryInputPort{
		Title:       input.Title,
		Description: input.Description,
	}, nil
}

type updateDiaryRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUpdateDiaryInputPortFromRequest(r *http.Request) (*diary.UpdateDiaryInputPort, error) {
	input := updateDiaryRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	return &diary.UpdateDiaryInputPort{ID: id, Title: input.Title, Description: input.Description}, nil
}

// type getDiaryRequestBody struct{}

func NewGetDiaryInputPortFromRequest(r *http.Request) (*diary.GetDiaryInputPort, error) {
	// INFO: RequestBodyから受け取る値が無いので、処理をコメントアウト
	//  input := getDiaryRequestBody{}
	//  if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
	//  	return nil, err
	//  }

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	return &diary.GetDiaryInputPort{ID: id}, nil
}

// type listDiaryRequestBody struct{}

func NewListDiaryInputPortFromRequest(r *http.Request) (*diary.ListDiaryInputPort, error) {
	// INFO: RequestBodyから受け取る値が無いので、処理をコメントアウト
	//  input := listDiaryRequestBody{}
	//  if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
	//  	return nil, err
	//  }

	return &diary.ListDiaryInputPort{}, nil
}

// type deleteDiaryRequestFromBody struct {}

func NewDeleteDiaryInputPortFromRequest(r *http.Request) (*diary.DeleteDiaryInputPort, error) {
	// INFO: RequestBodyから受け取る値が無いので、処理をコメントアウト
	//  input := deleteDiaryRequestFromBody{}
	//  if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
	//  	return nil, err
	//  }

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	return &diary.DeleteDiaryInputPort{ID: id}, nil
}
