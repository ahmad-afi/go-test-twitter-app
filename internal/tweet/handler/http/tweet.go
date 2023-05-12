package http

import (
	"net/http"

	"github.com/alwisisva/twitter-app/internal/tweet"
	"github.com/alwisisva/twitter-app/internal/tweet/entity"
	"github.com/alwisisva/twitter-app/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

type TweetHandler struct {
	tweetSvc tweet.Service
}

func NewTweetHandler(tweetSvc tweet.Service) *TweetHandler {
	return &TweetHandler{
		tweetSvc: tweetSvc,
	}
}

func (th *TweetHandler) PostTweet(c *fiber.Ctx) error {
	params := new(entity.Posts)
	if err := c.BodyParser(params); err != nil {
		return helper.BuildResponse(c, false, "FAILED TO PARSER DATA", err.Error(), nil, http.StatusBadRequest)
	}

	ctx := c.Context()
	res, err := th.tweetSvc.Post(ctx, *params)
	if err != nil {
		return helper.BuildResponse(c, false, "FAILED TO PARSER DATA", err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(c, true, "TWEET CREATED", "", res, http.StatusCreated)
}

func (th *TweetHandler) GetListTweets(c *fiber.Ctx) error {
	params := new(entity.FilterPosts)
	if err := c.QueryParser(params); err != nil {
		return helper.BuildResponse(c, false, "FAILED TO PARSER DATA", err.Error(), nil, http.StatusBadRequest)
	}

	ctx := c.Context()
	res, err := th.tweetSvc.GetList(ctx, *params)
	if err != nil {
		return helper.BuildResponse(c, false, "FAILED TO GET LIST DATA", err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(c, true, "SUCCEED GET LIST TWEETS", "", res, http.StatusOK)
}

func (th *TweetHandler) GetTweetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return helper.BuildResponse(c, false, "path cant be empty", "path cant be empty", nil, http.StatusBadRequest)
	}

	ctx := c.Context()
	res, err := th.tweetSvc.GetByID(ctx, id)
	if err != nil {
		return helper.BuildResponse(c, false, "FAILED TO GET DATA", err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(c, true, "SUCCEED GET TWEET BY ID", "", res, http.StatusOK)
}
