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

func (s *service) GetList(ctx context.Context, filter entity.FilterPosts) (res []entity.Posts, err *helper.ErrorStruct) {
	return
}
func (s *service) GetByID(ctx context.Context, id string) (res entity.Posts, err *helper.ErrorStruct) {
	return
}
