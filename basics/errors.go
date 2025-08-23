package basics

import (
	"errors"
	"fmt"
)

func Errors() {
	fmt.Println("--errors---------------------------------------------------------------------------------------------")
	_, err := f2(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}

/*
Zgodnie z konwencją, błędy są ostatnią zwracaną wartością i mają typ error, wbudowany interfejs.
errors.New konstruuje podstawową wartość błędu z podanym komunikatem o błędzie.
Wartość nil w pozycji błędu oznacza, że ​​nie wystąpił błąd.
*/
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

/*
Błąd sentinel to wstępnie zadeklarowana zmienna, która jest używana do sygnalizowania określonego stanu błędu.
*/
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		/*
		Możemy zawijać błędy z błędami wyższego poziomu, aby dodać kontekst. 
		Najprostszym sposobem na to jest użycie %w w fmt.Errorf. Zawinięte błędy tworzą logiczny łańcuch (A zawija B, który zawija C itd.), 
		który można sprawdzić za pomocą funkcji takich jak errors.Is i errors.As.
		*/
		return fmt.Errorf("making tea: %w", ErrPower)
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			/*
			errors.Is sprawdza, czy dany błąd (lub dowolny błąd w jego łańcuchu) pasuje do określonej wartości błędu. 
			Jest to szczególnie przydatne w przypadku zawiniętych lub zagnieżdżonych błędów, umożliwiając identyfikację 
			określonych typów błędów lub błędów wartowniczych w łańcuchu błędów.
			*/
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
	return nil
}

/*
Możliwe jest użycie niestandardowych typów jako błędów poprzez zaimplementowanie na nich metody Error(). 
Oto wariant powyższego przykładu, który wykorzystuje niestandardowy typ do jawnego reprezentowania błędu argumentu.
*/

	
type argError struct {
    arg     int
    message string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// Return our custom error.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
