package basics

import (
	"fmt"
)

func TestGenerics() {
	fmt.Println("--generics-------------------------------------------------------------------------------------------")
	generics()
}

/*
Funkcje Go mogą być napisane tak, aby działały na wielu typach przy użyciu parametrów typu.
Parametry typu funkcji pojawiają się między nawiasami, przed argumentami funkcji.

Deklaracja ta oznacza, że s jest wycinkiem dowolnego typu T, który spełnia wbudowane ograniczenie comparable. x jest również wartością tego samego typu.
comparable jest użytecznym ograniczeniem, które umożliwia użycie operatorów == i != na wartościach danego typu. 
W tym przykładzie używamy go do porównywania wartości ze wszystkimi elementami wycinka, aż do znalezienia dopasowania. 
Ta funkcja indeksu działa dla każdego typu, który obsługuje porównywanie.
*/
func generics() {
	si := []int{10, 20, 15, -10}
	fmt.Println(index(si, 15))

}

/*
Funkcje Go mogą być napisane tak, aby działały na wielu typach przy użyciu parametrów typu.
Parametry typu funkcji pojawiają się między nawiasami, przed argumentami funkcji.

Deklaracja ta oznacza, że s jest wycinkiem dowolnego typu T, który spełnia wbudowane ograniczenie comparable. x jest również wartością tego samego typu.
comparable jest użytecznym ograniczeniem, które umożliwia użycie operatorów == i != na wartościach danego typu. 
W tym przykładzie używamy go do porównywania wartości ze wszystkimi elementami wycinka, aż do znalezienia dopasowania. 
Ta funkcja indeksu działa dla każdego typu, który obsługuje porównywanie.
*/
func index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

/*
Oprócz funkcji generycznych, Go obsługuje również typy generyczne. 
Typ może być parametryzowany za pomocą parametru typu, co może być przydatne do implementacji ogólnych struktur danych.
Poniższy przykład demonstruje prostą deklarację typu dla pojedynczo połączonej listy przechowującej dowolny typ wartości.
*/
type List[T any] struct {
	next *List[T]
	val  T
}