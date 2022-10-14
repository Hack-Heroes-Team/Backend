package models

import "time"

type Receipt struct {
	Id    int
	Name  string
	Shop  string
	Owner string
	Items []Item
	Date  time.Time
}
