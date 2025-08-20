package basics

import (
	"fmt"
	"math"
)

func TestMethodsAndInterfaces() {
	fmt.Println("--methods and interfaces----------------------------------------------------------------------------")
	methods()
}

type Coordinates struct {
	X, Y float64
}

/*
Go nie ma klas, ale możemy definiować metody dla typów.
Metoda to funkcja ze specjalnym argumentem nazywanym odbiorcą
Odbiorcę podajemy w jego własnej liście argumentów pomiędzy słowem func a nazwą metody.
Metodę możesz zadeklarować tylko dla odbiorcy którego typ jest zdefiniowany w tym samym pakiecie co metoda.
*/
func (v Coordinates) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

/*
Możesz również zadeklarować metodę na typach nie będących strukturami.
*/
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methods() {
	v := Coordinates{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

