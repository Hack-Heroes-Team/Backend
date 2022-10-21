package models

type Item struct {
	Id        int
	ReceiptId int
	Owner     string
	Name      string
	Place     string
	Price     float64
}
