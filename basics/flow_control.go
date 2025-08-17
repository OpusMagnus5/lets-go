package basics

import(
	"fmt"
	"math"
	"runtime"
	"time"
)

func TestFlowControl() {
	fmt.Println("--flow_control---------------------------------------------------------------------------------------")

	/* 
	 Go posiada tylko jeden typ pętli, jest nią pętla for.
	*/
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum of 0 to 9: ", sum)

	// Inicjalizacja oraz inkrementacja są opcjonalne.
	for ; sum < 1000; {
		sum += sum
	}

	// lub
	for sum < 1000 {
		sum += sum
	}

	// nieskończona pętla
	for {
		break
	}

	// Instrukcja if w Go
	if sum > 0 {
		fmt.Println("Sum is positive: ", sum)
	}

	/* 
	 Instrukcja if może rozpoczynać się krótką instrukcją umieszczoną tuż przed warunkiem boolowskim.
	 Zmienne zadeklarowane poprzez tą instrukcje są dostępne tylko do końca instrukcji if.
	 Zmienne zadeklarowane w środku krótkiej instrukcji if są również dostępne w środku każdego bloku else.
	*/
	if v := math.Pow(2, 3); v > 0 {
		fmt.Println("2^3 is positive: ", v)
	} else {
		fmt.Println("2^3 is not positive: ", v)
	}

	fmt.Println("Sqrt(77): ", Sqrt(77))

	checkOS()
	checkTime()

	testDefer()
}

// Znajduje liczbę z, taką że z² jest możliwie najbliższej liczby x.
func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 10; i++ {
		newZ := z - (z * z - x) / (2 * z)
		if newZ == z || math.Abs(newZ-z) < 1e-6 {
			println("Iteration: ", i, " z: ", z)
			return z
		}
		z = newZ
		println("Iteration: ", i, " z: ", z)
	}

	return z
}

func checkOS() {
	/* 
	 Wykonuje pierwszy przypadek dla którego wartość jego warunku jest równa wyrażaniu warunkowemu.
	*/
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func checkTime() {
	/* 
	 Switch bez głównego warunku jest tym samym co switch true.
	 Ta konstrukcja może być bardziej przejrzystym sposobem na zapisanie długiego ciągu if-then-else.
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Goood evening.")
	}
}

func testDefer() {
	/* 
	 Instrukcja defer opóźnia wykonanie funkcji do momentu gdy funkcja w której się znajduje nie zwróci wyniku.
	 Wywołane argumenty w defer są ewaluowane natychmiastowo, jednak sama funkcja jest wywołana dopiero gdy otaczająca ją funkcja zwróci wynik.
	*/
	defer fmt.Printf(" World")
	fmt.Printf("Hello")
	fmt.Printf(" Defer")

	/* 
	 Funkcje wywołane z instrukcją defer są umieszczane na stosie
	 Po tym jak główna funkcja zwróci wynik, funkcje wywołane wewnątrz niej z instrukcją defer z są wykonywane w kolejności 
	 od ostatniej do pierwszej z nich która została umieszczona na stosie
	*/
	fmt.Println("liczę")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("zrobione")
}
