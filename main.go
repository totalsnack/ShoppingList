package main

import "fmt"

func main() {
	const (
		applePrice = 5_99
		pearPrice  = 7_00
	)
	walletSize := 23_00

	basket := func(applesQuantity, pearsQuantity int) int {
		return applesQuantity*applePrice + pearsQuantity*pearPrice
	}

	shoppingList := map[string]interface{}{
		"Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?": float64(basket(9, 8)) / 100,
		"Скільки груш ми можемо купити?":                                walletSize / pearPrice,
		"Скільки яблук ми можемо купити?":                               walletSize / applePrice,
		"Чи ми можемо купити 2 груші та 2 яблука?":                      walletSize-basket(2, 2) > 0,
	}
	for k, v := range shoppingList {
		fmt.Println(k, "==>", v)
	}

}
