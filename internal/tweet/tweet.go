package tweet

import (
	"context"

	"github.com/alwisisva/twitter-app/internal/tweet/entity"
	"github.com/alwisisva/twitter-app/pkg/helper"
)

type Service interface {
	// TODO: define methods to be used by HTTP handlers to
	// interact with tweet functionalities.
	Post(ctx context.Context, params entity.Posts) (res int64, err *helper.ErrorStruct)
	GetList(ctx context.Context, filter entity.FilterPosts) (entity.ListPost, *helper.ErrorStruct)
	GetByID(ctx context.Context, id string) (entity.Posts, *helper.ErrorStruct)
}
