package logic

type OrderProduct struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
}

type ShelfProduct struct {
	ID                int
	ShelfID           int
	ShelfName         string
	ProductID         int
	ProductName       string
	OrderID           int
	Quantity          int
	IsMain            bool
	AdditionalShelves string
}
