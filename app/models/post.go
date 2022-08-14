package models

import "four-seasons-of-Kunming/app/utils"

type Post struct {
	Id            int
	Content       string
	PublisherName string
	PublishTime   utils.LocalTime
}
