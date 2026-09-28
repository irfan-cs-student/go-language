// 2. Product Struct

// Create a Product struct with:

// name → string
// price → float64
// quantity → int

// Then:

// Create 2 products
// Print their information
// Change the quantity of one product
// Calculate and print its total value:
// total = price × quantity

package main

import "fmt"

func main() {
	type product struct {
		name     string
		price    float64
		quantity int
	}

	electric := product{
		name:     "torch",
		price:    200,
		quantity: 2,
	}
	fmt.Println(electric.name, "price:", electric.price)
	fmt.Println("quantity", electric.quantity)
	fmt.Println("total price: ", electric.quantity*int(electric.price))

	fmt.Println()

	clothes := product{
		name:     "silvat",
		price:    5000,
		quantity: 5,
	}
	fmt.Println(clothes.name, "price:", clothes.price)
	fmt.Println("quantity", clothes.quantity)
	fmt.Println("total price : ", clothes.quantity*int(clothes.price))
	fmt.Println()

}
