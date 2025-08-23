package basics

import (
	"fmt"
)

func IO() {
	fmt.Println("--Basics I/O-----------------------------------------------------------------------------------------")
	var a, b int
	/*
	Metoda fmt.Scan() służy do wczytywania danych z wejścia standardowego.
	Przyjmuje jako argumenty wskaźniki do zmiennych, do których mają zostać wczytane wartości.
	*/
	fmt.Scan(&a, &b)
	fmt.Println(a + b)
}