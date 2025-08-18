package basics

import "fmt"

func TestStructures() {
	fmt.Println("--structures-----------------------------------------------------------------------------------------")
	pointers()
	structs()
	arrays()
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


