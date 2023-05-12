package http

import (
	"fmt"
	"net/http"

	"github.com/alwisisva/twitter-app/pkg/helper"
	"github.com/gofiber/fiber/v2"
)

// New construct a new Handler.
func New(app *fiber.App, th *TweetHandler) {
	app.Get("", HealthCheck)

	apiRouter := fmt.Sprintf("/api/%s", helper.EnvString("APPS_VERSION"))
	apiGroup := app.Group(apiRouter)

	apiGroup.Get("", HealthCheck)

	tg := apiGroup.Group("/tweet")
	{
		tg.Post("", th.PostTweet)
		tg.Get("", th.GetListTweets)
		tg.Get("/:id", th.GetTweetByID)
	}

}

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": "Server is up and running",
	})
}
