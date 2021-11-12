/*
	APIのRouter・Handlerの定義を行う。
*/
package api

import (
	"database/sql"

	"github.com/gorilla/mux"

	"github.com/radish-miyazaki/go-clean-architecture/api/handler"
	"github.com/radish-miyazaki/go-clean-architecture/interface/repository"
	"github.com/radish-miyazaki/go-clean-architecture/usecase/diary"
)

func buildProjectRoutes(r *mux.Router, db *sql.DB) {
	dr := repository.NewDiaryRepository(db)

	r.Handle(
		"/diary",
		handler.NewCreateDiaryHandler(diary.NewCreateDiaryUsecase(dr)),
	).Methods("POST")
	r.Handle(
		"/diary/{id}",
		handler.NewUpdateDiaryHandler(diary.NewUpdateDiaryUsecase(dr)),
	).Methods("PUT")
	r.Handle(
		"/diary/{id}",
		handler.NewDeleteDiaryHandler(diary.NewDeleteDiaryUsecase(dr)),
	).Methods("DELETE")
	r.Handle(
		"/diary/{id}",
		handler.NewGetDiaryHandler(diary.NewGetDiaryUsecase(dr)),
	).Methods("GET")
	r.Handle(
		"/diary",
		handler.NewListDiaryHandler(diary.NewListDiaryUsecase(dr)),
	).Methods("GET")
}

func BuildRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	buildProjectRoutes(r, db)
	return r
}
