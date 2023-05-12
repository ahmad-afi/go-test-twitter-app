package service

import (
	"context"

	"github.com/alwisisva/twitter-app/internal/tweet/entity"
)

type StoreClient interface {
	// TODO: define methods to be used by HTTP handlers to
	// interact with tweet storage (PostgreSQL).
	Post(ctx context.Context, params entity.Posts) (res int64, err error)
	GetList(ctx context.Context, filter entity.FilterPosts) ([]entity.Posts, int64, error)
	GetByID(ctx context.Context, id string) (entity.Posts, error)
}
