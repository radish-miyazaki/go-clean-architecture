/*
	Usecase内で用いるRepositoryのInterfaceを定義
*/

package repository

import (
	"context"

	"github.com/radish-miyazaki/go-clean-architecture/domain/model"
)

type DiaryRepository interface {
	Store(ctx context.Context, diary *model.Diary) (*model.Diary, error)
	Update(ctx context.Context, diary *model.Diary) (*model.Diary, error)
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (*model.Diary, error)
	FindAll(ctx context.Context) ([]*model.Diary, error)
}
