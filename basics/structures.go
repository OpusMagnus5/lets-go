package basics

import "fmt"

func TestStructures() {
	fmt.Println("--structures-----------------------------------------------------------------------------------------")
	pointers()
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
