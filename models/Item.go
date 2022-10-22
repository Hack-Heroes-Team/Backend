package models

type Item struct {
	Id        int
	Receiptid int
	Owner     string
	Name      string
	Shop      string
	Place     string
	City      string
	Price     float64
}
