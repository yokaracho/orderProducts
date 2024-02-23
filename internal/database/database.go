package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/yokaracho/orderapps/internal/logger"
	"github.com/yokaracho/orderapps/internal/logic"
)

var getLogger = logger.GetLogger()

func GetOrderProductsAndShelves(pool *pgxpool.Pool, orderNumbers []int) ([]logic.OrderProduct, []logic.ShelfProduct, error) {
	var orderProducts []logic.OrderProduct
	var shelfProducts []logic.ShelfProduct
	query := `
        SELECT id, order_id, product_id, quantity FROM orderproducts WHERE order_id IN (SELECT id FROM orders WHERE number = ANY($1))
    `
	rows, err := pool.Query(context.Background(), query, orderNumbers)
	if err != nil {
		getLogger.Error("Error scanning order product row:", err)
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderProduct logic.OrderProduct
		err := rows.Scan(&orderProduct.ID, &orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity)
		if err != nil {
			getLogger.Error("Error scanning order product row:", err)
			return nil, nil, err
		}
		orderProducts = append(orderProducts, orderProduct)
	}

	query = `
        SELECT sp.id, sp.shelf_id, s.name AS shelf_name, sp.product_id, p.name AS product_name, sp.order_id, sp.quantity, sp.is_main, sp.additional_shelves
		FROM shelfproducts sp 
		JOIN shelves s ON sp.shelf_id = s.id 
		JOIN products p ON sp.product_id = p.id 
		WHERE sp.order_id IN (SELECT id FROM orders WHERE number = ANY($1))
    `

	rows, err = pool.Query(context.Background(), query, orderNumbers)
	if err != nil {
		getLogger.Error("Error querying shelf products:", err)
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var shelfProduct logic.ShelfProduct
		err := rows.Scan(&shelfProduct.ID, &shelfProduct.ShelfID, &shelfProduct.ShelfName, &shelfProduct.ProductID, &shelfProduct.ProductName, &shelfProduct.OrderID, &shelfProduct.Quantity, &shelfProduct.IsMain, &shelfProduct.AdditionalShelves)
		if err != nil {
			getLogger.Error("Error scanning shelf product row:", err)
			return nil, nil, err
		}
		shelfProducts = append(shelfProducts, shelfProduct)
	}

	return orderProducts, shelfProducts, nil
}

func GetOrderNumbers(pool *pgxpool.Pool) ([]int, error) {
	query := "SELECT number FROM orders"

	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		getLogger.Error("Error querying order numbers: ", err)
		return nil, err
	}

	defer rows.Close()

	var orderNumbers []int

	for rows.Next() {
		var number int
		if err := rows.Scan(&number); err != nil {
			getLogger.Error("Error scanning order numbers:", err)
			return nil, err
		}

		orderNumbers = append(orderNumbers, number)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orderNumbers, nil
}
