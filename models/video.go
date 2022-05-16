package models

import (
	"time"
)

type Video struct {
	ID                int64 `gorm:"primaryKey"`
	UserID            int64
	PlayUrl           string `gorm:"column:play_url;"`
	CoverUrl          string `gorm:"column:cover_url;"`
	FavoriteCount     int64  `gorm:"column:favorite_count;"`
	CommentCount      int64  `json:"column:comment_count;"`
	FavoriteUserList  []*User `gorm:"many2many:video_with_favorite_user;"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
