package models

type Item struct {
	Id        int
	ReceiptId int
	Owner     string
	Name      string
	Shop      string
	Price     float64
}
