package service

import (
	"context"
	"log"
	"net/http"

	"github.com/alwisisva/twitter-app/internal/tweet/entity"
	"github.com/alwisisva/twitter-app/pkg/helper"
)

// service implements tweet.Service.
type service struct {
	storeClient StoreClient
}

// New construts a new service.
func New(storeClient StoreClient) *service {
	s := &service{
		storeClient: storeClient,
	}

	return s
}

// TODO: implements tweet.Service with service.
func (s *service) Post(ctx context.Context, params entity.Posts) (res int64, err *helper.ErrorStruct) {
	if err := usecaseValidation[entity.Posts](params); err != nil {
		log.Println(err)
		return res, err
	}

	params.ID = helper.IDGenerator()
	res, errRep := s.storeClient.Post(ctx, params)
	if errRep != nil {
		return res, &helper.ErrorStruct{
			Err:  errRep,
			Code: http.StatusInternalServerError,
		}
	}

	return res, nil
}

func (s *service) GetList(ctx context.Context, filter entity.FilterPosts) (res entity.ListPost, err *helper.ErrorStruct) {
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	if filter.Page < 1 {
		filter.Page = 0
	} else {
		filter.Page = (filter.Page - 1) * 10
	}

	listPosts, totalRows, errRepo := s.storeClient.GetList(ctx, filter)
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	for _, v := range listPosts {
		res.Posts = append(res.Posts, entity.Posts{
			ID:        v.ID,
			Text:      v.Text,
			CreatedBy: v.CreatedBy,
			UserID:    v.UserID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	totalRowsInt := int(totalRows)

	totalPage := helper.GetMeta(&filter.Limit, &filter.Page, &totalRowsInt)
	res.Meta = entity.Meta{
		Page:        filter.Page,
		Limit:       filter.Limit,
		TotalPage:   totalPage,
		TotalRecord: totalRowsInt,
	}

	return
}
func (s *service) GetByID(ctx context.Context, id string) (res entity.Posts, err *helper.ErrorStruct) {
	res, errRepo := s.storeClient.GetByID(ctx, id)
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return
}
