package models

import "time"

type Receipt struct {
	Id    int
	Name  string
	Shop  string
	Owner string
	Items []Item
	Price float64
	Date  time.Time
}
