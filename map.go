package main

import "fmt"

func main() {
	var res int
	groceryList := make(map[string]int)
	groceryList["хлеб"] = 50
	groceryList["молоко"] = 100
	groceryList["масло"] = 200
	groceryList["колбаса"] = 500
	groceryList["соль"] = 20
	groceryList["огурцы"] = 200
	groceryList["сыр"] = 600
	groceryList["ветчина"] = 700
	groceryList["буженина"] = 900
	groceryList["помидоры"] = 250
	groceryList["рыба"] = 300
	groceryList["хамон"] = 1500

	products := make(map[string]int)

	for a, b := range groceryList {
		if b > 500 {
			products[a] = b
		}
	}
	fmt.Println("MAP: ", products)
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}

	for i := 0; i < len(order); i++ {
		res += groceryList[order[i]]
	}
	fmt.Println("Result: ", res)
}
