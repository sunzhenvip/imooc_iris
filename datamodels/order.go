package datamodels

type Order struct {
	ID          int64 `sql:"ID"`
	UserId      int64 `sql:UserID`
	ProductId   int64 `sql:"productID"`
	OrderStatus int64 `sql:"orderStatus"`
}

const (
	OrderWait = iota
	OrderSuccess
	OrderFailed
)
