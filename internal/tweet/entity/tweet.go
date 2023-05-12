package entity

import "time"

type (
	Posts struct {
		ID        int64     `json:"id" db:"id"`
		Text      string    `json:"text" db:"text" validate:"required"`
		CreatedBy string    `json:"createdBy" db:"created_by" validate:"required"`
		UserID    int64     `json:"userId" db:"user_id" validate:"required"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	}

	FilterPosts struct {
		Limit int `query:"limit"`
		Page  int `query:"page"`
	}
)
