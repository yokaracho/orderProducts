package logic

import "fmt"

func PrintOrder(orderNumbers []int, shelfProducts []ShelfProduct) {
	fmt.Println("=+=+=+=")
	fmt.Printf("Страница сборки заказов %d\n\n", orderNumbers)

	currentShelf := ""

	for _, shelfProduct := range shelfProducts {
		if currentShelf != shelfProduct.ShelfName {
			fmt.Printf("\n===Стеллаж %s\n", shelfProduct.ShelfName)
			currentShelf = shelfProduct.ShelfName
		}

		fmt.Printf("%s (id=%d)\nзаказ %d, %d шт", shelfProduct.ProductName, shelfProduct.ProductID, shelfProduct.OrderID, shelfProduct.Quantity)

		if shelfProduct.AdditionalShelves != "" {
			fmt.Printf("\nдоп стеллаж: %s", shelfProduct.AdditionalShelves)
		}

		fmt.Println("\n")
	}
}
