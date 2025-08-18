package basics

import (
	"fmt"
	"strings"
)

func TestStructures() {
	fmt.Println("--structures-----------------------------------------------------------------------------------------")
	pointers()
	structs()
	arrays()
	slices()
}

/*
 Wskaźnik przechowuje adres pamięci w której znajduje się dana wartość.
 Typ *T to wskaźnik do wartości o typie T. Jego wartość zerowa wynosi nil.
 & operator generuje wskaźnik do jego zmiennej.
 Operator * oznacza wartość wskazywaną w pamięci przez dany wskaźnik. (ang. dereferencing)
*/
func pointers() {
	i, j := 42, 2701

	p := &i     // wkaźnik do i
	fmt.Println("Wskaźnik i przez *p: ", *p) // przeczytaj i poprzez wkaźnik
	*p = 21         // ustaw wartość i poprzez wkaźnik
	fmt.Println("Wartość i: ", i)  // zobacz nową wartość i

	p = &j         // wskaźnik do j
	*p = *p / 37   // podziel j za pomocą wskaźnika
	fmt.Println("Wartość j: ", j) // zobacz nowa wartość j
}

/*
 struct jest zbiorem pól
 Aby uzyskać dostęp do pola struktury używamy kropki.
 Do pól struktury możemy się również dostać za pomocą wskaźnika do struktury.

 By uzyskać dostęp do pola X struktury którą wskazuje wskaźnik do struktury p, możemy napisać (*p).X.
 Jednakże, ta notacja jest niezbyt poręczna, więc Go pozwala użyć nam następującej składni: p.X.
 W tym przypadku Go pozwala nam opuścić jawną dereferencję wskaźnika.
*/

type Vertex struct {
	X int
	Y int
}

func structs() {
	v := Vertex{1, 2}
	fmt.Println("Nowy strucy Vertex: ", v)
	v.X = 4
	fmt.Println("Zaktualizowany strucy Vertex: ", v)
	p := &v
	p.X = 8 // nie musimy robić dereferencji, żeby dostać się do pola
	fmt.Println("Zaktualizowany strucy Vertex przez wskaźnik: ", v)

	/*
	 Struktury literalne (ang. struct literals)
	 Struktura literalna oznacza nowo utworzoną wartość zawierającą strukturę poprzez wymiennie wartości jej pól.
	 Możesz przypisać wartości tylko części pól, poprzez użycie składni Name:. Porządek nazwany pól nie ma znaczenia.
	 Prefiks & zwraca wskaźnik do wartości struktury.
	*/
	var (
		v1 = Vertex{1, 2}  // posiada typ Vertex
		v2 = Vertex{X: 1}  // Y:0 jest domniemane
		v3 = Vertex{}      // X:0 oraz Y:0
		vp  = &Vertex{1, 2} // posiada typ *Vertex
	)
	fmt.Println("Struktura literalna: ", v1, v2, v3, vp)
}

/*
 Typ [n]T jest tablicą która zawiera n wartości typu T.
 Długość tablicy jest częścią jej typu, tak więc rozmiar tablicy nie może być zmieniany.
*/
func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

/*
 Wycinki (ang. slices)
 Tablice mają ustaloną długość. Natomiast wycinki (ang. slices) są elastycznym w użyciu podglądem (ang. view) na elementy
 pewnej tablicy, których długość możemy dynamicznie zmieniać.

 Type []T jest wycinkiem z elementami o typie T (bez podanej długości)
 Wycinek jest tworzony poprzez wybranie dwóch indeksów, dolnego (ang. low) oraz górnego (z wyłączeniem ostatniego elementu) (ang. high) oddzielonych dwukropkiem: a[low:high]
*/
func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // tworzenie wycinka
	fmt.Println(s)

	/*
	Wycinki są jak referencje do tablic. Wycinek nie przechowuje żadnych danych, tylko wskazuje na fragment pewnej tablicy.
 	Zmienianie wartości danego wycinka zmienia wartość elementu tablicy na którą wycinek wskazuje.
 	Inne wycinki które wskazują na tę samą tablicę, będą widziały dokonane zmiany.
	*/
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println("Wycinki a i b: ", a, b)

	b[0] = "XXX"
	fmt.Println("Zaktualizowane wycinki a i b: ", a, b)
	fmt.Println("Zaktualizowane tablica names: ", names)

	/*
	 Wycinki literalne (ang. slice literals)
	 Wycinek literalny wygląda jak tablica literalna bez podanej długości.
	*/
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Wycinek q: ", q)

	/*
	 Gdy tworzymy wycinki, można pominąć dolną lub górną granice, wtedy zostaną użyte ich wartości domyślne.
	 Wartość domyślna dla dolnej granicy wynosi zero, wartość granicy górnej jest równa długości tablicy.
	*/
	s2 := []int{2, 3, 5, 7, 11, 13}

	fmt.Println("Wycinek s2 [1:4] = ", s2[1:4])
	fmt.Println("Wycinek s2 [:2] = ", s2[:2])
	fmt.Println("Wycinek s2 [1:] = ", s2[1:])
	fmt.Println("Wycinek s2 [:] = ", s2[:])

	/*
	 Wycinek posiada zarówno długość (ang. length) jak i pojemność (ang. capacity).
	 Długość (ang. len) wycinka to liczba elementów na które wskazuje.
	 Pojemność (ang. cap) wycinka to liczba elementów tablicy na którą dany wycinek wskazuje, licząc od pierwszego elementu wycinka.
	 Długość wycinka nie może przekroczyć jego capacity. error outOfBound
	*/
	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3)

	// Nadaje wycinkowi długość zerową ale nie nie zmniejsza capacity
	s3 = s3[:0]
	printSlice(s3)

	// Powiększa długość wycinka.
	s3 = s3[:4]
	printSlice(s3)

	// Usuwa z wycinka pierwsze dwie wartości - zmienia capacity bo zmienia początek slice'a
	s3 = s3[2:]
	printSlice(s3)

	/*
	 Wartość zerowa wycinka to nil.
	 Wycinek nil ma długość oraz pojemność równą 0 oraz nie wskazuje na żadną tablicę.
	*/
	var s4 []int
	printSlice(s4)

	/*
	 Wycinki mogą być utworzone za pomocą wbudowanej funkcji make; tak właśnie tworzymy tablice o dynamicznej długości.
	 Funkcja make tworzy tablicę o wyzerowanych wartościach elementów oraz zwraca wycinek który na nią wskazuje.
	 By określić pojemność, wystarczy dodać trzeci argument do funkcji make.
	*/

	a2 := make([]int, 5) // długość 5
	printSlice(a2)

	b2 := make([]int, 0, 5) // długość 0, capacity 5
	printSlice(b2)

	c2 := b2[:2]
	printSlice(c2)

	d2 := c2[2:5]
	printSlice(d2)

	/*
	 Wycinki mogą zawierać jako swój element każdy typ, włączając w to inne wycinki.
	*/
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	/*
	Aby dodać nowy element do wycinka, Go posiada wbudowaną funkcję append
	Pierwszym parametrem funkcji append jest wycinek s o typie T, pozostałe parametry to wartości typu T które chcemy dołączyć do wycinka.
	Rezultatem funkcji append jest wycinek zawierający wszystkie elementy oryginalnego wycinka oraz dodatkowe elementy.
	Jeśli tablica na którą wskazuje s jest zbyt mała by pomieścić wszystkie wartości, nowa, większa tablica zostanie utworzona.
	Zwrócony wycinek będzie wskazywał na nowo przydzieloną tablice.
	*/

	var s5 []int
	printSlice(s5)

	// append działa na wycinkach nilowych.
	s5 = append(s5, 0)
	printSlice(s5)

	// Wycinek powiększa się w miarę potrzeb.
	s5 = append(s5, 1)
	printSlice(s5)

	// Możemy dodać więcej niż jeden element w danym czasie.
	s5 = append(s5, 2, 3, 4)
	printSlice(s5)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
