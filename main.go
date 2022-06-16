package main

import "fmt"

type Fruit struct {
	price    int
	quantity int
}

type Basket struct {
	Items []*Fruit
}

func (basket *Basket) TotalPrice() (total int) {
	for _, f := range basket.Items {
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

	basketReach := &Basket{[]*Fruit{&Fruit{applePrice, 9}, &Fruit{pearPrice, 9}}}
	basketPoor := &Basket{[]*Fruit{&Fruit{applePrice, 2}, &Fruit{pearPrice, 2}}}

	shoppingList := map[string]interface{}{
		"Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?": float64(basketReach.TotalPrice()) / 100,
		"Скільки груш ми можемо купити?":                                walletSize / pearPrice,
		"Скільки яблук ми можемо купити?":                               walletSize / applePrice,
		"Чи ми можемо купити 2 груші та 2 яблука?":                      walletSize-basketPoor.TotalPrice() >= 0,
	}
	for k, v := range shoppingList {
		fmt.Println(k, "==>", v)
	}

}
