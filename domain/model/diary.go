/*
	Domain Layer
	Infrastructure（データリソース等）から独立したモデルを定義する。
*/

package model

type Diary struct {
	ID          int
	Title       string
	Description string
}
