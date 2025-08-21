package basics

import (
	"fmt"
	"math"
)

func TestMethodsAndInterfaces() {
	fmt.Println("--methods and interfaces----------------------------------------------------------------------------")
	methods()
	interfaces()
}

type Coordinates struct {
	X, Y float64
}

/*
Go nie ma klas, ale możemy definiować metody dla typów.
Metoda to funkcja ze specjalnym argumentem nazywanym odbiorcą
Odbiorcę podajemy w jego własnej liście argumentów pomiędzy słowem func a nazwą metody.
Metodę możesz zadeklarować tylko dla odbiorcy którego typ jest zdefiniowany w tym samym pakiecie co metoda.
Metoda nie może zmieniać wartości parametru.
*/
func (v Coordinates) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

/*
Możesz również zadeklarować metodę na typach nie będących strukturami.
*/
func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*
Możesz zadeklarować metody dla odbiorców wskaźników (ang. pointer receivers).
To oznacza, że typ odbiorcy ma składnię literalną *T (Dodatkowo T sam nie może być wskaźnikiem, takim jak na przykład *int.)
Metody z odbiorcami wskaźników mogą zmieniać wartości które wskazuje odbiorca!!!
W przeciwieństwie do funkcji tutaj zachodzi automatyczna konwersja na wskaźnik przy jej wywoływaniu.

Istnieją dwa powody dla których warto używać odbiorców wskaźników.
Pierwszy to taki że taka metoda może zmodyfikować wartość którą wskazuje odbiorca.
Drugim powód jest taki, że unikamy kopiowania wartości przy każdym wywołaniu metody.
To może być bardzo wydajna metoda pracy z danymi, jeśli na przykład odbiorca jest dużą strukturą.
*/

func (v *Coordinates) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methods() {
	v := Coordinates{3, 4}
	fmt.Println(v.abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.abs())

	v2 := Coordinates{3, 4}
	v2.scale(10)
	fmt.Println(v2.abs())
}

/*
Typ klasy interfejs jest definiowany przez zbiór sygnatur jego metody.
Typ implementuje interfejs implementując jego metody. Nie ma tutaj jawnej deklaracji tworzącej dany typ klasy interfejs, takiej jak na przykład „implements”.
Domyślne tworzenie interfejsu rozdziela definicje interfejsu od jego implementacji, które mogą się potem pojawić w dowolnym pakiecie
*/
type I interface {
	m()
}

type I2 interface {
	m2()
}

type T struct {
	s string
}

func (t T) m() {
	fmt.Println(t.s)
}

/*
Jeśli konkretna wartość w interfejsie to nil, metoda zostanie wywołana z odbiorcą nil.
*/
func (t *T) m2() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.s)
}

/*
Wartość nil interfejsu nie przetrzymuje ani konkretnej wartości ani żadnego konkretnego typu.
Wezwanie metody na wartości interfejsu będącą nil powoduje błąd w czasie wykonania programu (ang. run-time error),
*/

/*
„Pusty interfejs” to typ interfejs który nie określa żadnej metody: interface{}
Pusty interfejs może przetrzymywać wartości każdego typu. (Każdy typ implementuje co najmniej zero metod.)
Puste interfejsy używane są przez kod który obsługuje wartości nieznanych typów. Na przykład, fmt.Print przyjmuje jakakolwiek liczbę argumentów typu interface{}.
*/

func interfaces() {
	var i I = T{"witaj"}
	i.m()

	var t2 *T
	var i2 I2 = t2
	i2.m2()

	/*
	Sprawdzanie typu pozwala uzyskać dostęp do wartości o konkretnym typie która jest zawarta w interfejsie.
	Instrukcja pniżej zakłada że wartość interfejsu i zawiera konkretny typ T oraz przypisuje konkretną wartość interfejsu o typie T zmiennej t.
	Jeśli i nie zawiera typu T, instrukcja ta spowoduje wywołanie paniki (ang. panic).
	*/
	var hello interface{} = "hello"

	s := hello.(string)
	fmt.Println(s)

	/*
	By przetestować czy interfejs zawiera wartość konkretnego typu, sprawdzenie typu może zwrócić dwie wartości: 
	wartość zawartą w interfejsie oraz wartość typu bool która określa czy sprawdzenie typu dało wynik pozytywny.
	Jeśli i zawiera typ T, wtedy t będzie wartością zawartą w interfejsie a ok będzie miało wartość true.
	W przeciwnym wypadku, ok będzie miało wartość false, a t będzie miało wartością zerową typu T, zaś panika (ang. panic) nie zostanie wywołana.
	*/
	s, ok := hello.(string)
	fmt.Println(s, ok)

	/*
	Switch typów jest konstrukcją który pozwala dokonać kilku sprawdzeń typów po kolei.
	*/
	switch v := hello.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	/*
	Jednym z najbardziej wszechobecnym interfejsów jest Stringer zdefiniowany przez pakiet fmt.
	Stringer jest typem który może przedstawić samego siebie jako string. Pakiet fmt (i wiele innych) używa tego interfejsu by wypisywać wartości.
	*/

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

