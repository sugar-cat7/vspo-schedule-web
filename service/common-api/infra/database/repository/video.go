package repository

import (
	"context"

	"github.com/sugar-cat7/vspo-portal/service/common-api/domain/model"
	"github.com/sugar-cat7/vspo-portal/service/common-api/domain/repository"
	"github.com/sugar-cat7/vspo-portal/service/common-api/infra/database"
	db_sqlc "github.com/sugar-cat7/vspo-portal/service/common-api/infra/database/internal/db"
	"github.com/sugar-cat7/vspo-portal/service/common-api/infra/database/internal/dto"
)

type video struct{}

func NewVideo() repository.Video {
	return &video{}
}

var _ repository.Video = (*video)(nil)

func (r *video) List(
	ctx context.Context,
	query repository.ListVideosQuery,
) (model.Videos, error) {
	_, err := database.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	res := model.Videos{}
	if query.Page.Valid && query.Limit.Valid {

	}
	return res, nil
}

func (r *video) Count(
	ctx context.Context,
	query repository.ListVideosQuery,
) (uint64, error) {
	// FIXME: implement
	return 0, nil
}

func (r *video) UpsertAll(
	ctx context.Context,
	m model.Videos,
) (model.Videos, error) {
	c, err := database.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	br := c.Queries.CreateVideo(ctx, dto.VideoModelsToCreateVideoParams(m))
	defer br.Close()

	var i model.Videos
	br.QueryRow(func(_ int, ch db_sqlc.Video, err error) {
		if err != nil {
			return
		}
		i = append(i, dto.VideoToModel(&ch))
	})

	return i, nil
}
