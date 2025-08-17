package basics

import "fmt"

/* 
 Funkcja może przyjąć zero lub więcej argumentów.
 Typ podany jest po nazwie zmiennej
*/

func add(x int, y int) int {
	return x + y
}

/* 
 Gdy dwa lub więcej kolejnych parametrów funkcji są tego samego typu, 
 można pominąć podawanie typu po każdy z parametrów, umieszczając go tylko po ostatnim z nich.
*/

func add2(x, y int) int {
	return x + y
}

/* 
 Funkcja może zwrócić dowolną liczbę wartości.
*/

func swap(x, y string) (string, string) {
	return y, x
}

/* 
 W Go wartości zwracane mogą posiadać nazwy.
 Instrukcja return bez żadnych argumentów zwraca nazwane wartości. Mówimy wtedy o (ang. "naked" return)
 Naga instrukcja return powinna być używana tylko w krótkich funkcjach, takich jak ta w podanym przykładzie. 
 W przeciwnym razie jej użycie może pogorszyć czytelność długich funkcji.
*/

func split(sum int) (x, y int) { // tu jest deklaracja zmiennych a w funkcji są one inicjalizowane
	x = sum * 4 / 9
	y = sum - x
	return
}

func Test() {
	fmt.Println("--functions---------------------------------------")
	fmt.Println("add: ", add(42, 13))
	fmt.Println("add2: ", add2(42, 13))

	hello, world := swap("hello", "world")
	fmt.Println("swap: ", hello, world)

	split1, split2 := split(17)
	fmt.Println("split: ", split1, split2)
}