package models

import (
	"database/sql"
	"time"
)
type DeletedAt sql.NullTime

type User struct {
	ID            int64   `gorm:"primaryKey"`
	UserName      string  `gorm:"column:username;unique;"`
	Password	  string  `gorm:"column:password;"`
	FollowCount   int64   `gorm:"column:follow_count;"`
	FollowerCount int64   `gorm:"column:follower_count;"`
	VideoList     []Video `gorm:"foreignKey:UserID;"`
	Followers     []*User `gorm:"many2many:user_relation;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
