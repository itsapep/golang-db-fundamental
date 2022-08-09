package entity

type Product struct {
	Id          string
	Name        string
	Price       string
	Description string
	Stock       int
	Shop        string `db:"shop_id"`
}
