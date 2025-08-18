package main

/*
 Każdy program w Go składa się z pakietów.
 Pakiet main jest specjalnym pakietem, który oznacza, że program jest wykonywalny.
 W tym pakiecie musi znajdować się funkcja main, która jest punktem wejścia programu.
 Pakiet jest zdefiniowany przez słowo kluczowe `package`, a jego nazwa to `main`.

 Nazwa pliku czy nazwa katalogu nie są wgl powiązane z nazwą pakietu,
 ale konwencja jest taka, że plik z kodem źródłowym powinienien nazywać się `main.go`.
 Wszystki pliki w katalogu powinnuny należeć do tego samego pakietu.
*/

import (
	"fmt"
	"lets-go/basics"
	"math"
)

/*
 Można pisać każdy import w osobnej linii, ale przyjętą konwencją jest importowanie wielu pakietów w jednym imporcie.
*/

func main() {
	fmt.Printf("Teraz masz %g problemów.\n", math.Sqrt(7))

	/*
	   W Go nazwa jest eksportowana, gdy zaczyna się od dużej litery.
	   pi nie zaczynają się od dużej litery, więc nie są eksportowane.
	   Wszystkie „nieeksportowane” nazwy nie są dostępne poza pakietem w którym zostały zdefiniowane.
	*/
	//fmt.Println(math.pi)
	fmt.Println(math.Pi)

	basics.TestFunctions()
	basics.TestVariables()
	basics.TestFlowControl()
	basics.TestStructures()
}
