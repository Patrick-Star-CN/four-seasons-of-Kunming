package models

import (
	"time"
)

type Msg struct {
	Id            int       `json:"id"`
	Content       string    `json:"content"`
	PublisherName string    `json:"publisher_name"`
	CreateAt      time.Time `json:"publish_time"`
}
