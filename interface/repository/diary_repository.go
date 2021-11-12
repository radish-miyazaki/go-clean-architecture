/*
	Repository Layer
	実際に、データストレージ（DB）とやり取りを行うレイヤー。
	今回は、Usecase内で定義されたInterfaceに沿って、実際の処理を定義する。
*/

package repository

import (
	"context"
	"database/sql"

	"github.com/radish-miyazaki/go-clean-architecture/domain/model"
)

type diaryRepository struct {
	db *sql.DB
}

func NewDiaryRepository(db *sql.DB) *diaryRepository {
	return &diaryRepository{db: db}
}

func (dr diaryRepository) Store(ctx context.Context, diary *model.Diary) (*model.Diary, error) {
	if err := dr.db.QueryRowContext(
		ctx,
		`INSERT INTO diaries(title, description) VALUES ($1, $2) RETURNING id`,
		diary.Title,
		diary.Description,
	).Scan(&diary.ID); err != nil {
		return nil, err
	}

	return diary, nil
}

func (dr diaryRepository) Update(ctx context.Context, diary *model.Diary) (*model.Diary, error) {
	if err := dr.db.QueryRowContext(
		ctx,
		`UPDATE diaries SET title = $1, description = $2 WHERE id = $3`,
		diary.Title,
		diary.Description,
		diary.ID,
	).Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
		// Error処理
		return nil, err
	}

	return diary, nil
}

func (dr diaryRepository) Delete(ctx context.Context, id int) error {
	if _, err := dr.db.ExecContext(
		ctx,
		`DELETE FROM diaries WHERE id = $1`,
		id,
	); err != nil {
		return err
	}

	return nil
}

func (dr diaryRepository) FindByID(ctx context.Context, id int) (*model.Diary, error) {
	diary := &model.Diary{}
	if err := dr.db.QueryRowContext(
		ctx,
		`SELECT id, title, description FROM diaries WHERE id = $1`,
		id,
	).Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
		return nil, err
	}

	return diary, nil
}

func (dr diaryRepository) FindAll(ctx context.Context) (diaries []*model.Diary, err error) {
	rows, err := dr.db.QueryContext(
		ctx,
		`SELECT id, title, description FROM diaries`,
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		var diary model.Diary
		if err := rows.Scan(
			&diary.ID,
			&diary.Title,
			&diary.Description,
		); err != nil {
			return nil, err
		}
		diaries = append(diaries, &diary)
	}

	return diaries, nil
}
