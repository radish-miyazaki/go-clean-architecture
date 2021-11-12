/*
	Handler
	api.goで定義した部分の実体。
	流れとしては、http.Request -> Adapter -> Usecase -> Presenter -> ResponseWriter。
	これの流れを集約する。
*/
package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/radish-miyazaki/go-clean-architecture/api/adapter"
	"github.com/radish-miyazaki/go-clean-architecture/api/presenter"
	"github.com/radish-miyazaki/go-clean-architecture/usecase/diary"
)

func NewCreateDiaryHandler(du *diary.CreateDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adapter（リクエストをビジネスロジックにわたす処理）
		input, err := adapter.NewCreateDiaryInputPortFromRequest(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Usecase（ビジネスロジック）
		output, err := du.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// JSON以外の形式に変えたい場合でも、ビジネスロジックに触れずに変更できる。
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Presenter（ビジネスロジックの結果を外部向けに返却）
		if err := json.NewEncoder(w).Encode(presenter.NewCreateDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewUpdateDiaryHandler(du *diary.UpdateDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewUpdateDiaryInputPortFromRequest(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := du.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewUpdateDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewDeleteDiaryHandler(du *diary.DeleteDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewDeleteDiaryInputPortFromRequest(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := du.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewDeleteDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewGetDiaryHandler(du *diary.GetDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewGetDiaryInputPortFromRequest(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := du.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewGetDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewListDiaryHandler(du *diary.ListDiaryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewListDiaryInputPortFromRequest(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := du.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewListDiaryPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
