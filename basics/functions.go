package basics

import (
	"fmt"
	"iter"
)

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

/*
Variadic Functions mogą być wywoływane z dowolną liczbą argumentów końcowych.
Wewnątrz funkcji typ nums jest równoważny []int. Możemy wywołać len(nums), iterować po nim za pomocą range, itp.
*/
func sum2(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
	for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func TestFunctions() {
	fmt.Println("--functions------------------------------------------------------------------------------------------")
	fmt.Println("add: ", add(42, 13))
	fmt.Println("add2: ", add2(42, 13))

	hello, world := swap("hello", "world")
	fmt.Println("swap: ", hello, world)

	split1, split2 := split(17)
	fmt.Println("split: ", split1, split2)

	sum2(1, 2) // Variadic Functions mogą być wywoływane w zwykły sposób z indywidualnymi argumentami.
	nums := []int{1, 2, 3, 4}
	sum2(nums...) // Jeśli masz już wiele argumentów w wycinku, zastosuj je do funkcji variadic za pomocą func(slice...) w następujący sposób.

	CountTo10()
}

/*
Funkcja iteratora przyjmuje jako parametr inną funkcję, zwaną umownie yield (ale nazwa może być dowolna). 
Wywoła ona yield dla każdego elementu, nad którym chcemy iterować, i zanotuje wartość zwracaną yield dla potencjalnego wcześniejszego zakończenia.
*/
func CountTo10() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i) { // próbujemy oddać wartość
				return       // ktoś już nie chce, więc kończymy
			}
		}
	}
}