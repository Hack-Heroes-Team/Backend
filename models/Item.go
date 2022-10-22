package models

type Item struct {
	Id        int
	ReceiptId int
	Owner     string
	Name      string
	Shop      string
	Place     string
	City      string
	Price     float64
}
