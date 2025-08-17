package basics

import "fmt"

/* 
 Instrukcja var deklaruje listę zmiennych; tak jak w liście argumentów funkcji typ podajemy na samym końcu.
 Instrukcja var może znajdować się na poziomie pakietu lub funkcji
*/

var c, python, java bool

func TestVariables() {
	fmt.Println("--variables------------------------------------------------------------------------------------------")
	var i int
	fmt.Println("var: ", i, c, python, java)

	/* 
	 Deklaracja zmiennej może zawierać inicjalizator, po jednym dla każdej zmiennej.
	 Jeśli inicjalizator jest obecny, podanie typu jest zbędne; zmienna przyjmie typ inicjalizatora.
	*/
	var c, python, java = true, false, "nie!"
	fmt.Println("initialised var: ", i, c, python, java)

	/* 
	 W środku funkcji możemy użyć składni deklaracji := zamiast var z domniemanym typem.
	 Poza funkcją, każda instrukcja rozpoczyna się od słowa kluczowego (var, func, i tak dalej), więc nie możemy tam używać składni :=.
	*/
	kotlin, rust := true, 1
	fmt.Println("short var: ", kotlin, rust)

	fmt.Println("--types----------------------------------------------------------------------------------------------")

	/* 
	 Podstawowe typy Go to:
		bool
		string
		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr
		byte // alias dla uint8
		rune // alias dla int32
		float32 float64
		complex64 complex128

	 Typy int, uint, oraz uintptr mają zazwyczaj długość 32 bitów na 32-bitowych systemach i 64 bitów na systemach 64-bitowych.
	*/
	var (
		toBe   		bool	= false
		maxInt 		uint64 	= 1
		bigString 	string	= "Hello, World!"
	)
	printValueAndType(toBe)
	printValueAndType(maxInt)
	printValueAndType(bigString)

	/* 
	 Zmienne zadeklarowane bez podania jawnej wartości początkowej przyjmują wartość zerową.
	 Wartości zerowe wynoszą:
	 	0 dla typów numerycznych,
	 	false dla typu bool, oraz
	 	"" (pusty string) dla stringów.
	*/

	/* 
	 Wyrażenie T(v) konwertuje wartość v na typ T.
	*/
	toConvert := 42
	converted := float64(toConvert)
	fmt.Println("convert: ", toConvert, converted)

	/* 
	 Stałe są deklarowane podobnie jak zmienne, jednak z użyciem słowa kluczowego const.
	 Stałe mogą być znakami tekstowymi, stringami, wartościami boolowskimi lub liczbowymi.
	 Stałe nie mogą być zadeklarowane przy pomocy składni :=
	*/
	const Pi = 3.14
	fmt.Println("const Pi: ", Pi)

	/* 
	 Stałe numeryczne (numeric constants) to stałe, które są liczbami - Nie mają typu aż do użycia w wyrażeniu
	*/
	const (
		Big = 12676506002282295
		Small = 2
	)
	fmt.Println("needInt Small: ", needInt(Small))
	fmt.Println("needFloat Small: ", needFloat(Small))
	fmt.Println("needFloat Big: ", needFloat(Big))

}

func printValueAndType(value any) {
	fmt.Printf("Typ: %T Wartość: %v\n", value, value)
}


func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}