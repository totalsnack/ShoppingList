package main

import "fmt"

type Fruit struct {
	price    int
	quantity int
}

func basketOfFruits(F ...Fruit) (total int) {
	for _, f := range F {
		total += f.price * f.quantity
	}
	return
}

func main() {
	const (
		applePrice = 5_99
		pearPrice  = 7_00
	)
	walletSize := 23_00

	shoppingList := map[string]interface{}{
		"Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?": float64(basketOfFruits(Fruit{applePrice, 9}, Fruit{pearPrice, 9})) / 100,
		"Скільки груш ми можемо купити?":                                walletSize / pearPrice,
		"Скільки яблук ми можемо купити?":                               walletSize / applePrice,
		"Чи ми можемо купити 2 груші та 2 яблука?":                      walletSize-basketOfFruits(Fruit{applePrice, 2}, Fruit{pearPrice, 2}) >= 0,
	}
	for k, v := range shoppingList {
		fmt.Println(k, "==>", v)
	}

}
